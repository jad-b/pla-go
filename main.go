package api

import (
	"github.com/jad-b/pla-go/api"
	// "html/template"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Running server...")
	http.HandleFunc("/echo/", api.EchoHandler)
	http.HandleFunc("/echo/json/", api.JSONEchoHandler)
	http.ListenAndServe(":8080", nil)
	defer log.Print("Server is shutting down.")
}
