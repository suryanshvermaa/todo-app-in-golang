package todo

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/suryanshvermaa/todo-app-in-golang/internal/storage"
	"github.com/suryanshvermaa/todo-app-in-golang/internal/types"

	"github.com/suryanshvermaa/todo-app-in-golang/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newTodo types.Todo
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if errors.Is(err, io.EOF) {
			response.JsonResponse(w, 400, "empty body", nil)
			return
		}
		if err != nil {
			response.JsonResponse(w, 400, err.Error(), nil)
			return
		}
		if err := validator.New().Struct(newTodo); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.JsonResponse(w, 400, validateErrs.Error(), nil)
			return
		}
		newTodo, _ = storage.CreateTodo(newTodo.Todo, newTodo.Completed)
		slog.Info("creating todo")
		response.JsonResponse(w, http.StatusCreated, "todo created successfully", newTodo)
	}
}

func GetTodo(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.JsonResponse(w, 400, err.Error(), nil)
			return
		}
		todo, err := storage.GetTodoById(id)
		if err != nil {
			slog.Error("error getting todo", "error", err.Error())
			response.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		response.JsonResponse(w, http.StatusOK, "todo retrieved successfully", todo)
	}
}

func GetAllTodos(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, _ := storage.GetAllTodos()
		if todos == nil {
			response.JsonResponse(w, 400, "there are zero todos in the db", nil)
			return
		}
		response.JsonResponse(w, 200, "todos fetched successfully", todos)
	}
}

func DeleteTodo(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.JsonResponse(w, 400, err.Error(), nil)
			return
		}
		isDeleted, _ := storage.DeleteTodoById(id)
		if isDeleted {
			response.JsonResponse(w, 200, "todo deleted successfully", nil)
			return
		}
		response.JsonResponse(w, 400, "err in deleting todo", nil)
	}
}

func UpdateTodo(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.JsonResponse(w, 400, err.Error(), nil)
			return
		}
		var updatedTodo types.Todo
		err = json.NewDecoder(r.Body).Decode(&updatedTodo)
		if errors.Is(err, io.EOF) {
			response.JsonResponse(w, 400, "empty body", nil)
			return
		}
		if err != nil {
			response.JsonResponse(w, 400, err.Error(), nil)
			return
		}
		if err := validator.New().Struct(updatedTodo); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.JsonResponse(w, 400, validateErrs.Error(), nil)
			return
		}
		updatedTodo, _ = storage.UpdateTodoById(id, updatedTodo.Todo, updatedTodo.Completed)
		slog.Info("updating todo")
		response.JsonResponse(w, http.StatusOK, "todo updated successfully", updatedTodo)
	}
}
