package api

import (
	"encoding/json"
	"net/http"
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

	h.Service.Add(newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}
