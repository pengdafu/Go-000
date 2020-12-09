package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func serveApp(ctx context.Context, quit chan struct{}) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 5)
		_, _ = w.Write([]byte("Hello Golang!"))
	})
	srv := http.Server{Addr: ":8081", Handler: mux}

	go func() {
		defer close(quit)
		<-ctx.Done()
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10) // 不要用传来的ctx
		defer cancelFunc()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("srv shutdown fail: %+v", err)
			return
		}
		log.Println("graceful shutdown")
	}()
	return srv.ListenAndServe()
}

func handleSignal(ctx context.Context) error {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGKILL)
	select {
	case err := <-ch:
		return fmt.Errorf("%v", err)
	case <-ctx.Done():
		return nil
	}
}

func main() {
	ctx, fc := context.WithCancel(context.Background())

	g, _ := errgroup.WithContext(ctx)
	quit := make(chan struct{})
	g.Go(func() error {
		if err := serveApp(ctx, quit); err != nil {
			fc()
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := handleSignal(ctx); err != nil {
			fc()
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("%+v", err)
	}
	<-quit
}
