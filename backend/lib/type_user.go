package main

import (
	"encoding/json"
	"strconv"
)

// Represents a user
type User struct {
	ApiKey  string `json:"api_key"`
	Balance int    `json:"balance"`
}

// Serializes a User to a JSON string, returning
//
//	a Result containing the JSON string or an error
func (user User) JSONString() string {
	var body string = "{"

	body += "\"usernames\": \"" + user.ApiKey + "\", "
	body += "\"emails\": " + strconv.Itoa(user.Balance)

	body += "}"

	return body
}

// Deserializes a User from a JSON string, returning
//
//	a Result containing the User or an error
func UserFromJSON(json_string string) Result[User] {
	res := User{}

	err := json.Unmarshal([]byte(json_string), &res)

	if err != nil {
		return Err[User]("\nFailed to deserialize the following string:\n\"" + json_string + "\"\n\nRaw Error: " + err.Error())
	}

	return Ok(res)
}
