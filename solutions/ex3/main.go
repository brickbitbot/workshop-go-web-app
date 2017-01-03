package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})

	mux.Handle("/proverbs/", http.StripPrefix("/proverbs/", http.FileServer(http.Dir("../static"))))

	log.Println("Running...")
	http.ListenAndServe(":9000", mux)
}
