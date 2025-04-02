package api

import (
	"fmt"
	"net/http"
	"strings"
	"todo-pet/internal/app/command"
	"todo-pet/internal/app/task"
	"todo-pet/internal/storage"
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

	serializer := storage.NewJSONSerializer()

	err := serializer.Serialize(h.Service.GetAll(), w)
	if err != nil {
		http.Error(w, "Ошибка при сериализации", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) addTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Парсим JSON с полем title
	var input struct {
		Title string `json:"title"`
	}

	serializer := storage.NewJSONSerializer()

	err := serializer.Deserialize(r.Body, &input)
	if err != nil || strings.TrimSpace(input.Title) == "" {
		http.Error(w, "Некорректный JSON или пустое название задачи", http.StatusBadRequest)
		return
	}

	cmd := &command.AddCommand{
		Service: h.Service,
		Title:   input.Title,
	}

	if err := cmd.Execute(); err != nil {
		http.Error(w, "Ошибка при добавлении задачи", http.StatusInternalServerError)
		return
	}

	todos := h.Service.GetAll()
	newTodo := todos[len(todos)-1]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = serializer.Serialize(newTodo, w)
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
