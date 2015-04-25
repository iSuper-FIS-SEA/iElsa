package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	Index(rw, req)

	if rw.Code != 200 {
		t.Errorf("Expect 200 but got %d", rw.Code)
	}

	if rw.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expect application/json but got %s", rw.Header().Get("Content-Type"))
	}

	if rw.Body.String() != `{"text":"hello"}` {
		t.Errorf(`Expect {"text":"hello"} but got %s`, rw.Body)
	}
}
