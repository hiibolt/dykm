package main

import (
	"testing"
	"os"
)

// Disabled to prevent the cost of API calls
/*
func TestSnusbaseQueryEmail(t *testing.T) {
	request := APIRequest{
		PIIType: PIITypeEmail,
		PII:     "codins12345@example.com",
	}

	tally_result := SnusbaseQuery(request.PIIType, request.PII);

	if(tally_result.IsErr()){
		log.Fatalln(tally_result.UnwrapErr());
		t.Fatal("Error while fetching data via snusbase query via email!")
	}

	tally := tally_result.UnwrapOk();

	log.Println(tally.String())
}
*/

func TestHTMLInjector(t *testing.T) {
	// Create a basic new HTML file
	html_contents := "<!--~__test.html~-->";
	html_test_contents := "<p>meow</p>";

	// Throw an error if the file already exists
	if _, err := os.Stat("../frontend/pages/__test.html"); err == nil {
		t.Fatal("File '../frontend/pages/__test.html' already exists!");
	}
	if _, err := os.Stat("../frontend/components/__test.html"); err == nil {
		t.Fatal("File '../frontend/components/__test.html' already exists!");
	}

	// Create these two files with the `os` library
	os.WriteFile("../frontend/pages/__test.html", []byte(html_contents), 0644);
	os.WriteFile("../frontend/components/__test.html", []byte(html_test_contents), 0644);

	// Inject the component
	inject_result := inject_component("__test.html");

	// Remove the test files
	os.Remove("../frontend/pages/__test.html");
	os.Remove("../frontend/components/__test.html");

	if(inject_result.IsErr()){
		t.Fatal("Error while injecting component: " + inject_result.UnwrapErr());
	}

	// Check if the injected component is equal to the test contents
	if(inject_result.UnwrapOk() != html_test_contents){
		t.Fatal("Injected component does not match test contents");
	}
}