package main;

// Represents the final tally of for the PII that was requested
type Tally struct {
	usernames int32;
	emails    int32;
	phones    int32;
	hashes    int32;
	salts     int32;
	ips       int32;
	names     int32;
	passwords int32;
	addresses int32;
	companies int32;
	other     int32;
};

// Converts an APIResponse to a human-readable string
func (response Tally) String() string {
	return "todo!();";
};