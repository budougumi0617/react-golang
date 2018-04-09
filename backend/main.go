package main

import (
	"log"
	"net/http"

	"github.com/budougumi0617/react-golang/backend/routing"
)

func main() {
	router := routing.GetTaskRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
