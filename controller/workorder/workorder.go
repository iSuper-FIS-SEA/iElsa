package workorder

import (
	workorder "../../model/workorder"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) error {
	var worders = workorder.NewWorkorderManager()
	hdNum := mux.Vars(r)["hdnum"]

	res := struct{ Workorder []*workorder.Workorder }{worders.Find(hdNum)}
	if res.Workorder == nil {
		return errors.New("notFound")
	}
	return json.NewEncoder(w).Encode(res)
}
