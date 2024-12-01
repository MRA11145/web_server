package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm();
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error %v", err);
		return;
	}
	fmt.Fprintf(w, "Post request SucessFul");
	name := r.FormValue("name");
	address := r.FormValue("address");
	fmt.Fprintf(w, "Name is %s\n", name);
	fmt.Fprintf(w, "Address is %s\n", address);
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("The request is ", r)
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!");
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port number")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}