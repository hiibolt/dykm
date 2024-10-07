package main

import (
	"encoding/json"
	"strings"
)

// Represents the final Sherlock of for the PII that was requested
type Sherlock struct {
	username *string  `json:"username"`
	sites    []string `json:"sites"`
}

// a Result containing the JSON string or an error
func (sherlock Sherlock) JSONString() string {
	var body string = "{"
	body += "\"username\": \"" + *sherlock.username + "\", "
	body += "\"sites\": \"" + strings.Join(sherlock.sites, ", ") + "\", "
	body += "}"

	return body
}

// Deserializes a Sherlock from a JSON string, returning
func SherlockFromJSON(json_string string) Result[Sherlock] {
	res := Sherlock{}

	err := json.Unmarshal([]byte(json_string), &res)

	if err != nil {
		return Err[Sherlock]("\nFailed to deserialize the following string:\n\"" + json_string + "\"\n\nRaw Error: " + err.Error())
	}

	return Ok(res)
}