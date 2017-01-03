package main

import (
	"net/http"
)

func response(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello World."))
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":3000", nil)
}
