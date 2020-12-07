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


func serveApp(ctx context.Context) error {
	srv := http.Server{Addr: ":8081"}

	go func() {
		<-ctx.Done()
		ctx, cancelFunc := context.WithTimeout(ctx, time.Second*2)
		defer cancelFunc()
		_ = srv.Shutdown(ctx)
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

	g.Go(func() error {
		if err := serveApp(ctx); err != nil {
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
		log.Println("exit after 2 seconds")
	}
	<-time.After(time.Second * 2)
}
