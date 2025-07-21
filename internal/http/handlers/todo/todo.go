package todo

import (
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		slog.Info("creating todo")
		w.Write([]byte("creted todo"))
	}
}
