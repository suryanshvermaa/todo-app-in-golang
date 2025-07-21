package todo

import (
	"log/slog"
	"net/http"

	"github.com/suryanshvermaa/todo-app-in-golang/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		slog.Info("creating todo")
		response.JsonResponse(w, http.StatusCreated, "todo created successfully", nil)
	}
}
