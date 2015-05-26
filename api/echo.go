package api

import (
	// "github.com/jad-b/pla-go/auth"
	// "html/template"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HelloWorld ...
type HelloWorld struct {
	Msg string
	Day time.Time
}

// EchoHandler echoes
func EchoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You sent the following request:\n")
	req.Write(w) // Write the request out in wire format
}

// JSONEchoHandler echoes
func JSONEchoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You sent the following request:\n")
	json.NewEncoder(w).Encode(req) // Write the request out in json
}

// HelloWorldHandler ...
func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(&HelloWorld{"Hello, world!", time.Now()})
}
