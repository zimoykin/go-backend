package main

import (
	"net/http"
)

func fileServer(w http.ResponseWriter, r *http.Request) {
	http.Handle("/", http.FileServer(http.Dir("/public")))

}
