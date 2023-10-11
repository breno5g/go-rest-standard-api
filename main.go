package main

import (
	"fmt"
	"net/http"

	"github.com/breno5g/rest-standard-go-api/routes"
	_ "github.com/lib/pq"
)

func main() {
	mux := routes.Init()

	fmt.Println("Server is running on port 3001")
	http.ListenAndServe(":3001", mux)
}
