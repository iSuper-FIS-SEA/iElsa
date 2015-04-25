package main

import (
	router "./router"
	"log"
	"net/http"
)

func main() {
	router.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("webapp")))

	err := http.ListenAndServe(":8080", nil)
	log.Printf(err.Error())
}
