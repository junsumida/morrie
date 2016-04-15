package main

import (
	"morrie"
	"net/http"
)

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte(morrie.Message))
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":30303", nil)
}
