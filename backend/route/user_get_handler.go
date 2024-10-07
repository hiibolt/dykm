package main;

import (
	"net/http"
	"log"
)

func user_get_handler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		return
	}

	//Parse form from user
	err := req.ParseMultipartForm(10 << 20)

	log.Println(req.Form)

	if err != nil {
		log.Println("Error parsing form:", err)
		return
	}

	if !req.Form.Has("User-API-Key") {
		ReturnErr(w, 400, "Bad Request", "Please provide a User-API-Key value.")
		return
	}

	result := GetUser(req.FormValue("User-API-Key"))

	if result.IsErr() {
		ReturnErr(w, 500, "Internal Server Error", result.UnwrapErr())
		return
	}

	ReturnJson(w, result.UnwrapOk())
}