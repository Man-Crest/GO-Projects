package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "this is not /hello route", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not get", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "your name is :%s", name)
	fmt.Fprintf(w, "your email is :%s", email)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	port := ":3000"
	fmt.Printf("Server is running on port %s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server listening error:", err)
	}
}
