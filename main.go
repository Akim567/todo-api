package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Title: "Купить молоко", Completed: false},
	{Title: "Позвонить другу", Completed: true},
	{Title: "Выгулять собаку", Completed: false},
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
			return
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
