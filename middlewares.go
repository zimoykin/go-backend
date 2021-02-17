package main

import (
	"log"
	"net/http"
)

func middleTest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware")
		next.ServeHTTP(w, r)
	})
}
