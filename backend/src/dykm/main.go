package main

import (
	"io"
	"log"
	"net/http"
)

func hello_handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	var port string = ":8000";

	http.HandleFunc("/hello", hello_handler);

    log.Println("Listing for requests at http://localhost:8000/hello");
	
	log.Fatal(http.ListenAndServe(port, nil));
}