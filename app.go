package main

import (
	rount "./controller"
	"log"
	"net/http"
)

func main() {
	rount.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("webapp")))

	err := http.ListenAndServe(":8080", nil)
	log.Printf(err.Error())
}
