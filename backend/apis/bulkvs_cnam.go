package main;

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func BulkVSCNAMQuery(PII string, user_api_key string) Result[BulkVS] {
	// Send a POST request to Snusbase API
	url := "https://osint.hiibolt.com/api/v1/tele/bulkvs_cnam" 
	data := []byte(PII);

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data));

	if err != nil {
		return Err[BulkVS](err.Error());
	}

	//Add headers
	req.Header.Add("Content-Type", "application/json");
	req.Header.Add("Authorization", os.Getenv("API_KEY"));
	req.Header.Add("User-API-Key", user_api_key);
	
	//Send request
	resp, err := client.Do(req);

	if err != nil {
		return Err[BulkVS](err.Error());
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Err[BulkVS](err.Error());
	}

	return BulkVSFromJSON(string(body));
}