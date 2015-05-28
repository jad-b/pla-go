package test

import (
	"bytes"
	"encoding/json"
	"github.com/jad-b/pla-go/api"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	baseURL = "127.0.0.1"
	tests   = []struct {
		in  string
		out string
	}{}
)

func TestHelloWorldHandler(t *testing.T) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	api.HelloWorldHandler(rr, req)

	var response api.HelloWorld
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Error("Failed to unmarshal JSON")

	} else if response.Msg != "Hello, world!" {
		t.Error("YOU NEED TO SAY HELLO")
	}
}

func TestEchoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	api.EchoHandler(rr, req)

	buf := &bytes.Buffer{} // Composite literal
	req.Write(buf)
	ret := rr.Body.String()
	ret = ret[strings.Index(ret, "\n")+1:]
	if ret != buf.String() {
		t.Errorf("Expected to see this:\n%s\nGot this instead:\n%s", ret, buf.String())
	} else if rr.Code != 200 {
		t.Error("Expect a return code of 200")
	}
}

func TestJSONEchoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	api.JSONEchoHandler(rr, req)

	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(req)
	ret := rr.Body.String()
	ret = ret[strings.Index(ret, "\n")+1:]
	if ret != buf.String() {
		t.Errorf("Expected to see this:\n%s\nGot this instead:\n%s", ret, buf.String())
	} else if rr.Code != 200 {
		t.Error("Expect a return code of 200")
	}
}
