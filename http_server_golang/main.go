package main

import (
	"net/http"
	"time"
)

func main() {

	// create the ServeMux
	m := http.NewServeMux()
	// register the handle func
	m.HandleFunc("/", testHandler)

	// create the server
	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// Call the ListenAndServe func to "start" the server
	srv.ListenAndServe()

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}