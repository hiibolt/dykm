package main

import (
	"encoding/json"
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
		http.StripPrefix("/public/", http.FileServer(http.Dir("../../../frontend/public"))),
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
