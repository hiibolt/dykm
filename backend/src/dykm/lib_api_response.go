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

// Adds two Tallys together.
func (tally1 Tally) add(tally2 Tally) Tally {
	return Tally{
		usernames: tally1.usernames + tally2.usernames,
		emails:    tally1.emails + tally2.emails,
		phones:    tally1.phones + tally2.phones,
		hashes:    tally1.hashes + tally2.hashes,
		salts:     tally1.salts + tally2.salts,
		ips:       tally1.ips + tally2.ips,
		names:     tally1.names + tally2.names,
		passwords: tally1.passwords + tally2.passwords,
		addresses: tally1.addresses + tally2.addresses,
		companies: tally1.companies + tally2.companies,
		other:     tally1.other + tally2.other,
	};
}

// Converts an APIResponse to a human-readable string
func (response Tally) String() string {
	return "todo!();";
};