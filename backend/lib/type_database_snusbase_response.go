package main

import (
	"encoding/json"
	"fmt"
)

// Represents the final SnusbaseDB of for the PII that was requested
type SnusbaseDB struct {
	took int32  `json:"took"`
	size int32 `json:"size"`
	results    map[string][]map[string]interface{} `json:"results"`
}

// a Result containing the JSON string or an error
func (snusbaseDB SnusbaseDB) JSONString() string {
	var jsonStr string;
	var body string = "{"
	body += "\"took\": \"" + string(snusbaseDB.took) + "\", "
	body += "\"size\": \"" + string(snusbaseDB.size) + "\", "

	jsonData, err := json.Marshal(snusbaseDB.results)
	if err != nil {
		fmt.Println("Error marshaling results:", err)
	} else {
		jsonStr = string(jsonData)
	}

	body += "\"results\": \"" + jsonStr + "\" "
	body += "}"

	return body
}

// Deserializes a SnusbaseDB from a JSON string, returning
func SnusbaseDBFromJSON(json_string string) Result[SnusbaseDB] {
	res := SnusbaseDB{}

	err := json.Unmarshal([]byte(json_string), &res)

	if err != nil {
		return Err[SnusbaseDB]("\nFailed to deserialize the following string:\n\"" + json_string + "\"\n\nRaw Error: " + err.Error())
	}

	return Ok(res)
}