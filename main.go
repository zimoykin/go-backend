package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	app := StartApp()
	app.handleRequests()
}

func whoami(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server!")
	fmt.Fprintf(w, r.URL.Query().Encode())
}

type Application struct {
	//routes map[string]int
	logger *log.Logger
	Port   string
	Host   string
}

func StartApp() *Application {

	res := Application{
		Port: "8080",
		Host: "0.0.0.0",
	}

	return &res
}

func (app *Application) handleRequests() {

	http.HandleFunc("/", whoami)
	http.HandleFunc("/second", second)

	jsontestFunc := http.HandlerFunc(jsontest)
	http.Handle("/api/v1/jsontest", middleTest(jsontestFunc))

	err := http.ListenAndServe(app.Host+":"+app.Port, nil)
	if err != nil {
		log.Fatal("error!", err)
	}

	app.logger.Println("server started")

}
