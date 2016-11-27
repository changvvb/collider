package main

import (
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/static/", server))
	http.ListenAndServe(":10000", server)
}
