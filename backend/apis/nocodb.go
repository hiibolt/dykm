package main;

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func GetUser(UserApiKey string) Result[User] {
	// Send a POST request to Snusbase API
	url := "https://osint.hiibolt.com/api/v1/users/get" 
	data := []byte("");

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data));

	if err != nil {
		return Err[User](err.Error());
	}

	//Add headers
	req.Header.Add("Content-Type", "application/json");
	req.Header.Add("Authorization", os.Getenv("API_KEY"))
	req.Header.Add("User-API-Key", UserApiKey);
	
	//Send request
	resp, err := client.Do(req);

	if err != nil {
		return Err[User](err.Error());
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Err[User](err.Error());
	}

	// Check that it was a 200 OK response
	if resp.StatusCode != 200 {
		return Err[User](string(body));
	}

	return UserFromJSON(string(body));
}