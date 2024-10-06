package main;

import (
	"net/http"
	"log"
)

func api_handler(w http.ResponseWriter, req *http.Request) {

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

	if !req.Form.Has("PIIType") || !req.Form.Has("PII") {
		ReturnErr(w, 400, "Bad Request", "Please provide a PIIType and PII Value.")
		return
	}

	apiRequest := APIRequest{
		PIIType: PIIType(req.FormValue("PIIType")),
		PII:     req.FormValue("PII"),
	}

	tally_result := TallyResults(apiRequest)

	if tally_result.IsErr() {
		ReturnErr(w, 500, "Internal Server Error", tally_result.UnwrapErr())
		return
	}

	ReturnJson(w, tally_result.UnwrapOk())
}