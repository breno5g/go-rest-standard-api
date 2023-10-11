package main

import (
	"fmt"
	"net/http"
)

type TaskHandler struct{}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/task/", &TaskHandler{})

	fmt.Println("Server is running on port 3001")
	http.ListenAndServe(":3001", mux)
}
