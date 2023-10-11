package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/breno5g/rest-standard-go-api/model"
)

var (
	listTaskRe   = regexp.MustCompile(`^\/task[\/]*$`)
	getTaskRe    = regexp.MustCompile(`^\/task\/(\d+)$`)
	createTaskRe = regexp.MustCompile(`^\/task[\/]*$`)
	deleteTaskRe = regexp.MustCompile(`^\/task\/(\d+)$`)
	updateTaskRe = regexp.MustCompile(`^\/task\/(\d+)$`)
)

type TaskHandler struct{}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == "GET" && listTaskRe.MatchString(r.URL.Path):
		h.List(w, r)
	case r.Method == "GET" && getTaskRe.MatchString(r.URL.Path):
		fmt.Fprintln(w, "GET")
	case r.Method == "POST" && createTaskRe.MatchString(r.URL.Path):
		fmt.Fprintln(w, "POST")
	case r.Method == "DELETE" && deleteTaskRe.MatchString(r.URL.Path):
		fmt.Fprintln(w, "DELETE")
	case r.Method == "PUT" && updateTaskRe.MatchString(r.URL.Path):
		fmt.Fprintln(w, "PUT")
	default:
		http.NotFound(w, r)
	}
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks := model.List()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
