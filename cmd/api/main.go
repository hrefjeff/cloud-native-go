package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"myapp/config"
)

func main() {
	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server on " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	// Test breaking the connection
	// time.Sleep(3 * time.Second)
	io.WriteString(w, "Hello World!")
}
