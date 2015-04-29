package workorder

import (
	workorder "../../model/workorder"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func HdnumList(w http.ResponseWriter, r *http.Request) error {
	var worders = workorder.NewWorkorderManager()
	hdNum := mux.Vars(r)["hdnum"]

	res := struct{ Workorder []*workorder.Workorder }{worders.FindHdnum(hdNum)}
	if res.Workorder == nil {
		return errors.New("notFound")
	}
	return json.NewEncoder(w).Encode(res)
}

func OrderList(w http.ResponseWriter, r *http.Request) error {
	var orders = workorder.NewWorkorderManager()
	order := mux.Vars(r)["order"]

	res := struct{ Workorder []*workorder.Workorder }{orders.FindOrder(order)}
	if res.Workorder == nil {
		return errors.New("notFound")
	}
	return json.NewEncoder(w).Encode(res)
}

func FileList(w http.ResponseWriter, r *http.Request) error {
	var zfiles = workorder.NewWorkorderManager()
	zfile := mux.Vars(r)["file"]

	res := struct{ Workorder []*workorder.Workorder }{zfiles.FindFile(zfile)}
	if res.Workorder == nil {
		return errors.New("notFound")
	}
	return json.NewEncoder(w).Encode(res)
}
