package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewStore()
	mux := http.NewServeMux()

	// Frontend — Go serves index.html directly
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// API
	mux.HandleFunc("POST /todos", createTodo(store))
	mux.HandleFunc("GET /todos", listTodos(store))
	mux.HandleFunc("GET /todos/{id}", getTodo(store))
	mux.HandleFunc("PATCH /todos/{id}", updateTodo(store))
	mux.HandleFunc("DELETE /todos/{id}", deleteTodo(store))

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", loggingMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
