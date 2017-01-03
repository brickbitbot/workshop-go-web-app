package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../static")))
	log.Println("Running...")
	http.ListenAndServe(":9000", nil)
}
