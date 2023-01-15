package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	server := &http.Server{
		Addr: ":8080",
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	log.Println("start http server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
