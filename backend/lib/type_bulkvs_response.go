package main

import (
	"encoding/json"
	"strconv"
)

// Represents the final tally of for the PII that was requested
type BulkVS struct {
	Name   *string `json:"name"`
	Number *string `json:"number"`
	Time  *int64  `json:"time"`
}

//	a Result containing the JSON string or an error
func (bulkVS BulkVS) JSONString() string {
	var body string = "{"

	if bulkVS.Name != nil {
		body += "\"name\": \"" + *bulkVS.Name + "\", "
	}
	if bulkVS.Number != nil {
		body += "\"number\": \"" + *bulkVS.Number + "\", "
	}
	if bulkVS.Time != nil {
		body += "\"time\": " + strconv.FormatInt(*bulkVS.Time, 10)
	}

	body += "}"

	return body
}

// Deserializes a BulkVS from a JSON string, returning
func BulkVSFromJSON(json_string string) Result[BulkVS] {
	res := BulkVS{}

	err := json.Unmarshal([]byte(json_string), &res)

	if err != nil {
		return Err[BulkVS]("\nFailed to deserialize the following string:\n\"" + json_string + "\"\n\nRaw Error: " + err.Error())
	}

	return Ok(res)
}
