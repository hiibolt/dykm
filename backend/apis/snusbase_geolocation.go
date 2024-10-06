package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func SnusbaseGeo(PII string) Result[Tally] {
	// Send a POST request to Snusbase API
	url := "https://osint.hiibolt.com/api/v1/tally/snusbase_geolocation/ip" 
	data := []byte(PII);

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data));

	if err != nil {
		return Err[Tally](err.Error());
	}

	//Add headers
	req.Header.Add("Content-Type", "application/json");
	req.Header.Add("Authorization", os.Getenv("API_KEY"))
	
	//Send request
	resp, err := client.Do(req);

	if err != nil {
		return Err[Tally](err.Error());
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Err[Tally](err.Error());
	}

	return TallyFromJSON(string(body));
}