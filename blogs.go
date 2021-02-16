package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsontest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	p := BlogResponse{"42", "something here"}
	jdata, err := json.Marshal(p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("something went wrong")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jdata))

}

func getbloglist() {

}

type BlogResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
