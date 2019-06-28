package main

import (
	"encoding/json"
	"testing"
)

func TestString(t *testing.T) {
	response := response{"message", "signature", "pubKey"}

	valid := json.Valid([]byte(response.String()))

	if !valid {
		t.Error("String method returns invalid JSON")
	}
}
