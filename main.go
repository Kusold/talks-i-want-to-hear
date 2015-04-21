package main

import (
	"log"
	"net/http"

	"github.com/Kusold/talks-i-want-to-hear/router"
)

func main() {
	router.Router()
	port := ":8080"
	log.Println("Starting server. Listening on", port)
	http.ListenAndServe(port, nil)
}
