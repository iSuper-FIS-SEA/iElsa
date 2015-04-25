package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("webapp")))

	log.Printf("Go Baby go ...")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}
