package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type TallyErrResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ReturnErr(w http.ResponseWriter, statusCode int, statusString string, statusMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := TallyErrResponse{
		Status:  statusString,
		Message: statusMsg,
	}

	json.NewEncoder(w).Encode(response)
}

func ReturnJson[T any](w http.ResponseWriter, value T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(value)
}

func main() {
	var port string = ":" + os.Getenv("PORT")

	http.Handle(
		"/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("../frontend/public"))),
	)
	http.HandleFunc(
		"/",
		page_handler,
	)
	http.HandleFunc(
		"/api/tally",
		api_handler,
	)

	log.Printf("Listing for requests at http://localhost%s/!", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
