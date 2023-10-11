package routes

import (
	"net/http"

	"github.com/breno5g/rest-standard-go-api/controller"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/task/", &controller.TaskHandler{})

	return mux
}
