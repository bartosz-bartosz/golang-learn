package main

import (
	"goweb/hello/api"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello! This is the main page!")
	// })

	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello! Current path is %s\n", r.URL.Path)
	// 	w.Write([]byte("This is a test!"))
	// })

	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
