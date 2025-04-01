package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"todo-pet/internal/app/task"
)

type Handler struct {
	Service *task.Service
}

func NewHandler(service *task.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) RegisterRoutes() {
	http.HandleFunc("/ping", h.ping)
	http.HandleFunc("/todos", h.getTodos)
	http.HandleFunc("/add", h.addTodo)
	http.HandleFunc("/todos/table", h.getTodosTable)
}

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (h *Handler) getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Service.GetAll())
}

func (h *Handler) addTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var newTodo task.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Ошибка при парсинге JSON", http.StatusBadRequest)
		return
	}

	todo := h.Service.Add(newTodo.Title)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) getTodosTable(w http.ResponseWriter, r *http.Request) {
	todos := h.Service.GetAll()

	var builder strings.Builder

	builder.WriteString("| ID   | Title                | Status     | Completed At         |\n")
	builder.WriteString(strings.Repeat("-", 65) + "\n")

	for _, todo := range todos {
		var completedAt string
		if todo.CompletedAt != nil && !todo.CompletedAt.IsZero() {
			completedAt = todo.CompletedAt.Format("02.01.2006 15:04")
		} else {
			completedAt = "-"
		}

		builder.WriteString(fmt.Sprintf("| %-4d | %-20s | %-10s | %-20s |\n",
			todo.ID, todo.Title, todo.Status, completedAt))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(builder.String()))
}
