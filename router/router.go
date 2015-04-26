package controller

import (
	home "../controller/home"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const PathPrefix = "/iElsa/"

// Register router
func RegisterHandlers() {
	log.Printf("Go Baby go...")

	r := mux.NewRouter()
	r.HandleFunc(PathPrefix+"home", errorHandler(home.HomeHandler)).Methods("GET")
	http.Handle(PathPrefix, r)
}

type badRequest struct{ error }
type notFound struct{ error }

// Embed HandleFunc with errorHandler
func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)

		if err == nil {
			log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
			return
		}

		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "Not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "Oops", http.StatusInternalServerError)
		}
	}
}
