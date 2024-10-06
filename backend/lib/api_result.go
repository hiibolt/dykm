package main;

// Represents the type of PII that is being sent to the API
type PIIType string;
const (
	PIITypeEmail    PIIType = "email";
	PIITypePhone    PIIType = "phone";
	PIITypeUsername PIIType = "username";
	PIITypeIP       PIIType = "ip";
	PIITypePassword PIIType = "password";
	PIITypeName     PIIType = "name";
	PIITypeHash     PIIType = "hash";
);

// Represents the request that is sent to the API
type APIRequest struct {
	PIIType PIIType;
	PII     string;
};