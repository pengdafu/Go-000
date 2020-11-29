package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"week02/api/http"
	"week02/internal/model"
)

func main() {
	err := model.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	errs := make(chan error)
	r := http.Route()
	go func() {
		errs <- r.Run(":8080")
	}()
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		errs <- fmt.Errorf("%v", <-ch)
	}()

	log.Fatal(<-errs)
}
