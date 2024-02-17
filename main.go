package main

import (
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello-go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	router.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Created User"))
	})
	http.ListenAndServe(":8000", router)
}
