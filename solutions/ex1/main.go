package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	w.WriteHeader(http.StatusTeapot)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Running...")
	http.ListenAndServe(":9000", nil)
}
