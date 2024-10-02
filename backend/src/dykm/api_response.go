package main;

// Represents the final tally of for the PII that was requested
type APIResponse struct {
	email    int8;
	phone    int8;
	username int8;
	ip       int8;
	password int8;
	name     int8;
	hash     int8;
	total    int8;
};

// Converts an APIResponse to a human-readable string
func (response APIResponse) String() string {
	return "HELLO CHAT";
};