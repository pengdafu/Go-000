package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func newServe(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello Golang!"))
	})
	return &http.Server{Addr: addr, Handler: mux}
}

func start(srv *http.Server) error {
	return srv.ListenAndServe()
}

func shutdown(ctx context.Context, srv *http.Server) error {
	<-ctx.Done()
	ctx, fc := context.WithTimeout(context.Background(), time.Second*5)
	defer fc()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println("shutdown fail", err)
		return err
	}
	log.Println("graceful shutdown")
	return nil
}

func handleSignal(ctx context.Context) error {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	select {
	case sig := <-ch:
		return fmt.Errorf("%v", sig)
	case <-ctx.Done():
		return nil
	}
}

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "http address")
}

func main() {
	flag.Parse()

	ctx, cancelFunc := context.WithCancel(context.Background())

	g, _ := errgroup.WithContext(ctx)

	srv := newServe(addr)

	g.Go(func() error {
		if err := start(srv); err != nil {
			cancelFunc()
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := handleSignal(ctx); err != nil {
			cancelFunc()
			return err
		}
		return nil
	})

	g.Go(func() error {
		return shutdown(ctx, srv)
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}
