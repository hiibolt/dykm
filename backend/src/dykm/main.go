package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func public_handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../../frontend/pages/index.html")
}
func page_handler(w http.ResponseWriter, r *http.Request) {
	path_segments := strings.Split(r.URL.Path, "/")

	// Home Directory
	if len(path_segments) < 2 || path_segments[1] == "" {
		http.ServeFile(w, r, "../../../frontend/pages/index.html")

		return
	}

	switch path_segments[1] {
	case "about", "app":
		// Serve the respective page
		http.ServeFile(w, r, "../../../frontend/pages/"+path_segments[1]+".html")
	default:
		// Serve the 404 page
		http.ServeFile(w, r, "../../../frontend/pages/404.html")
	}
}
func api_handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	var port string = ":" + os.Getenv("PORT")

	http.Handle(
		"/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("../../../frontend/public"))),
	)
	http.HandleFunc(
		"/",
		page_handler,
	)
	http.HandleFunc(
		"/api",
		api_handler,
	)

	log.Println(fmt.Sprintf("Listing for requests at http://localhost%s/!", port))
	log.Fatal(http.ListenAndServe(port, nil))
}
