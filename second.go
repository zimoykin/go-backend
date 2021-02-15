package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func second(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getBlogs())
}

func getBlogs() string {

	resp, err := http.Get(server + "api/blogs/list")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)

	var blogs = "blogs: "
	fmt.Println(blogs)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		blogs += scanner.Text()
	}

	return "somewhere: " + blogs
}

var server = "http://10.0.1.6:8000/"
