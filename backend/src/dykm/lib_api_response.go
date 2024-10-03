package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Represents the final tally of for the PII that was requested
type Tally struct {
	Usernames int `json:"usernames"`
	Emails    int `json:"emails"`
	Phones    int `json:"phones"`
	Hashes    int `json:"hashes"`
	Salts     int `json:"salts"`
	Ips       int `json:"ips"`
	Names     int `json:"names"`
	Passwords int `json:"passwords"`
	Addresses int `json:"addresses"`
	Companies int `json:"companies"`
	Other     int `json:"other"`
}

// Serializes a Tally to a JSON string, returning
//
//	a Result containing the JSON string or an error
func (tally Tally) JSONString() string {
	var body string = "{"

	body += "\"usernames\": " + strconv.Itoa(tally.Usernames) + ", "
	body += "\"emails\": " + strconv.Itoa(tally.Emails) + ", "
	body += "\"phones\": " + strconv.Itoa(tally.Phones) + ", "
	body += "\"hashes\": " + strconv.Itoa(tally.Hashes) + ", "
	body += "\"salts\": " + strconv.Itoa(tally.Salts) + ", "
	body += "\"ips\": " + strconv.Itoa(tally.Ips) + ", "
	body += "\"names\": " + strconv.Itoa(tally.Names) + ", "
	body += "\"passwords\": " + strconv.Itoa(tally.Passwords) + ", "
	body += "\"addresses\": " + strconv.Itoa(tally.Addresses) + ", "
	body += "\"companies\": " + strconv.Itoa(tally.Companies) + ", "
	body += "\"other\": " + strconv.Itoa(tally.Other)

	body += "}"

	return body
}

// Deserializes a Tally from a JSON string, returning
//
//	a Result containing the Tally or an error
func TallyFromJSON(json_string string) Result[Tally] {
	res := Tally{}

	err := json.Unmarshal([]byte(json_string), &res)

	if err != nil {
		return Err[Tally]("\nFailed to deserialize the following string:\n\"" + json_string + "\"\n\nRaw Error: " + err.Error())
	}

	return Ok(res)
}

// Adds two Tallys together.
func (tally1 *Tally) add(tally2 Tally) {
	(*tally1).Usernames += tally2.Usernames;
	(*tally1).Emails += tally2.Emails;
	(*tally1).Phones += tally2.Phones;
	(*tally1).Hashes += tally2.Hashes;
	(*tally1).Salts += tally2.Salts;
	(*tally1).Ips += tally2.Ips;
	(*tally1).Names += tally2.Names;
	(*tally1).Passwords += tally2.Passwords;
	(*tally1).Addresses += tally2.Addresses;
	(*tally1).Companies += tally2.Companies;
	(*tally1).Other += tally2.Other;
}

// Converts an APIResponse to a human-readable string
func (response Tally) String() string {
	str := fmt.Sprintf(`
  Usernames:  %d
  Emails:     %d
  Phones:     %d
  Hashes:     %d
  Salts:      %d
  Ips:        %d
  Names:      %d
  Passwords:  %d
  Addresses:  %d
  Companies:  %d
  Other:      %d`,
		response.Usernames, response.Emails, response.Phones, response.Hashes, response.Salts, response.Ips,
		response.Names, response.Passwords, response.Addresses, response.Companies, response.Other,
	)
	return str
}
