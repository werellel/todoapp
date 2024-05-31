package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("fail to terminate server: %v", err)
	}
}

func run(_ context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	s.ListenAndServe()
	return nil
}
