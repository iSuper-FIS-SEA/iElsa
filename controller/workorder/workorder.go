package workorder

import (
	workorder "../../model/workorder"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var worders = workorder.NewWorkorderManager()

func List(w http.ResponseWriter, r *http.Request) error {
	_, err := worders.All()
	b, _ := json.Marshal(map[string]string{"text": "hello"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	return err
}

func GetHdNum(w http.ResponseWriter, r *http.Request) error {
	hdNum := mux.Vars(r)["hdnum"]
	log.Printf(hdNum)
	return nil
}
