package main

import "encoding/json"

// response is returned by this application
type response struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	PubKey    string `json:"pubkey"`
}

// String prints response as JSON object
func (r response) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	check(err)
	return string(b)
}
