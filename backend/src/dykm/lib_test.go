package main

import (
	"log"
	"strconv"
	"testing"
)

func TestSerializationCycleEquivalence ( t *testing.T ) {
	serialized_json := "{\"usernames\": 1, \"emails\": 2, \"phones\": 3, \"hashes\": 4, \"salts\": 5, \"ips\": 6, \"names\": 7, \"passwords\": 8, \"addresses\": 9, \"companies\": 10, \"other\": 11}";
		
	tally_result := TallyFromJSON(serialized_json);
	if tally_result.IsErr() {
		t.Fatal("Failed to deserialize Tally: " + tally_result.UnwrapErr());
	}

	tally := tally_result.UnwrapOk();

	log.Println("Deserialized Tally: ")
	log.Println(strconv.Itoa(tally.Usernames))
	log.Println(strconv.Itoa(tally.Emails))
	log.Println(strconv.Itoa(tally.Phones))

	serialized_tally := tally.JSONString()
	log.Println("Reserialized Ttally: " + serialized_tally)

	log.Println("Are the two tallies equal?")
	if serialized_json != serialized_tally {
		t.Fatal("The two tallies are not equal")
	}
}

func TestPrettyStrings(t *testing.T) {
	tally := Tally{}

	expected_output := `
  Usernames:  0
  Emails:     0
  Phones:     0
  Hashes:     0
  Salts:      0
  Ips:        0
  Names:      0
  Passwords:  0
  Addresses:  0
  Companies:  0
  Other:      0`

	log.Println(expected_output)
	log.Println(tally.String())
	if expected_output != tally.String() {
		t.Fatal("Output does not match expected output")
	}

}
