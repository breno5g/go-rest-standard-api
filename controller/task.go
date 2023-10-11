package controller

import (
	"fmt"
	"net/http"
)

type TaskHandler struct{}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == "GET":
		fmt.Fprintln(w, "GET")
	case r.Method == "POST":
		fmt.Fprintln(w, "POST")
	case r.Method == "PUT":
		fmt.Fprintln(w, "PUT")
	case r.Method == "DELETE":
		fmt.Fprintln(w, "DELETE")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
