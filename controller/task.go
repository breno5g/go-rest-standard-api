package controller

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	listTaskRe   = regexp.MustCompile(`^\/task[\/]*$`)
	getTaskRe    = regexp.MustCompile(`^\/task\/(\d+)$`)
	createTaskRe = regexp.MustCompile(`^\/task[\/]*$`)
)

type TaskHandler struct{}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == "GET" && listTaskRe.MatchString(r.URL.Path):
		fmt.Fprintln(w, "GET")
	case r.Method == "GET" && getTaskRe.MatchString(r.URL.Path):
		fmt.Fprintln(w, "GET")
	case r.Method == "POST" && createTaskRe.MatchString(r.URL.Path):
		fmt.Fprintln(w, "POST")
	default:
		http.NotFound(w, r)
	}
}
