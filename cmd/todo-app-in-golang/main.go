package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/suryanshvermaa/todo-app-in-golang/internal/config"
	"github.com/suryanshvermaa/todo-app-in-golang/internal/http/handlers/todo"
	"github.com/suryanshvermaa/todo-app-in-golang/internal/storage/data"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	// setup router
	db := &data.Database{}
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"sucess": "true", "message": "healthy"})
	})
	router.HandleFunc("POST /api/createTodo", todo.New(db))
	router.HandleFunc("GET /api/getTodo/{id}", todo.GetTodo(db))
	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Println("Server is running on", cfg.Addr)
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// graceful shutdown
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server:", err.Error())
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
