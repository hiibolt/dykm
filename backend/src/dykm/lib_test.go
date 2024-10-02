package main;

import (
	"log"
	"testing"
	"strconv"
);

func TestSerializationCycleEquivalence ( t *testing.T ) {
	serialized_json := "{\"usernames\": 1, \"emails\": 2, \"phones\": 3, \"hashes\": 4, \"salts\": 5, \"ips\": 6, \"names\": 7, \"passwords\": 8, \"addresses\": 9, \"companies\": 10, \"other\": 11}";
		
	tally_result := TallyFromJSON(serialized_json);
	if is_err(tally_result) {
		t.Fatal("Failed to deserialize Tally: " + tally_result.unwrap_err());
	}
	tally := tally_result.unwrap_ok();

	log.Println("Deserialized Tally: ");
	log.Println(strconv.Itoa(tally.Usernames));
	log.Println(strconv.Itoa(tally.Emails));
	log.Println(strconv.Itoa(tally.Phones));

	serialized_tally := tally.JSONString();
	log.Println("Reserialized Ttally: " + serialized_tally);

	log.Println("Are the two tallies equal?");
	if serialized_json != serialized_tally {
		t.Fatal("The two tallies are not equal");
	}
}