package todo

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/suryanshvermaa/todo-app-in-golang/internal/types"

	"github.com/suryanshvermaa/todo-app-in-golang/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var todo types.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if errors.Is(err, io.EOF) {
			response.JsonResponse(w, 400, "empty body", nil)
			return
		}
		if err != nil {
			response.JsonResponse(w, 400, err.Error(), nil)
			return
		}
		if err := validator.New().Struct(todo); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.JsonResponse(w, 400, validateErrs.Error(), nil)
			return
		}
		todo.ID = 1
		slog.Info("creating todo")
		response.JsonResponse(w, http.StatusCreated, "todo created successfully", todo)
	}
}
