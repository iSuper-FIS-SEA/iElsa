package controller

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	b, _ := json.Marshal(map[string]string{"text": "hello"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return nil
}
