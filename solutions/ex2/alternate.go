package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Running...")
	http.ListenAndServe(":9000", http.FileServer(http.Dir("../static")))
}
