package main

import (
	"log"
	"net/http"

	"todo-list/app"
)

func main() {
	log.Print("Listen And Server at http://localhost:8000")
	http.ListenAndServe(":8000", app.NewHandler())
}
