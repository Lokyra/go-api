package main

import (
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello-go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8000", router)
}
