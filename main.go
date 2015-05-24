package main

import (
	"github.com/jad-b/pla-go/auth"
	// "html/template"
	"log"
	"net/http"
)

func main() {
	log.Print("Running server...")
	http.ListenAndServe(":8080", nil)
	log.Print("Server is shutting down.")
}
