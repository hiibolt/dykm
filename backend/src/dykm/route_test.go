package main

import (
	"log"
	"testing"
)

func TestSnusbaseQueryEmail(t *testing.T) {
	request := APIRequest{
		PIIType: PIITypeEmail,
		PII:     "codins12345@example.com",
	}

	tally_result := SnusbaseQuery(request.PIIType, request.PII);

	if(tally_result.IsErr()){
		log.Fatalln(tally_result.UnwrapErr());
		log.Fatalln("Error while fetching data via snusbase query via email!");
	}

	tally := tally_result.UnwrapOk();

	log.Println(tally.String())
}
