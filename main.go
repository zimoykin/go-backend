package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//fmt.Println("This is main!")
	http.HandleFunc("/", whoami)
	http.HandleFunc("/second", second)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error!  ", err)
	}
}

func whoami(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "im live here, on port 8080!")
}
