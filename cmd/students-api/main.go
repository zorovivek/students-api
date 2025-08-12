package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zorovivek/students-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	// route esatblishment
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("server started %s", cfg.Http_Server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
