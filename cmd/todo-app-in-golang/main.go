package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suryanshvermaa/todo-app-in-golang/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Todo App!"))
	})
	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server:", err.Error())
	}

	fmt.Println("Server is running on", cfg.Addr)
}
