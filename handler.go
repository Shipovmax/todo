package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// writeJSON writes v as a JSON response body with the given HTTP status.
// Encoding errors are logged rather than ignored: by the time Encode fails,
// headers are already flushed, so the only remaining action is to surface
// the failure for observability instead of dropping it silently.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON: failed to encode response: %v", err)
	}
}

func listTodos(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, s.GetAll())
	}
}

func getTodo(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid ID"})
			return
		}
		todo, ok := s.GetByID(id)
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "task not found"})
			return
		}
		writeJSON(w, http.StatusOK, todo)
	}
}

func createTodo(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		if req.Title == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "title must not be empty"})
			return
		}
		todo := s.Create(req.Title)
		w.Header().Set("Location", fmt.Sprintf("/todos/%d", todo.ID))
		writeJSON(w, http.StatusCreated, todo)
	}
}

func updateTodo(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid ID"})
			return
		}
		var req struct {
			Status string `json:"status"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		if req.Status != "todo" && req.Status != "in_progress" && req.Status != "done" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "status must be todo, in_progress or done"})
			return
		}
		todo, ok := s.Update(id, req.Status)
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "task not found"})
			return
		}
		writeJSON(w, http.StatusOK, todo)
	}
}

func deleteTodo(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid ID"})
			return
		}
		ok := s.Delete(id)
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "task not found"})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
