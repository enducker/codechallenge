package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"os"
	"reflect"
	"testing"
)

func TestGetMsg(t *testing.T) {
	expected := "your@email.com"
	os.Args = []string{"cmd", "your@email.com"}

	actual := getMsg()

	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestGetKey(t *testing.T) {
	os.RemoveAll(pemDirectory)

	privKeyInit := getKey()
	privKeySubsequent := getKey()

	if !reflect.DeepEqual(privKeyInit, privKeySubsequent) {
		t.Error("Subsequent calls to getKey() do not return the same key")
	}

}

func TestSignMessage(t *testing.T) {
	msg := "your@email.com"
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	_, err := signMessage(msg, privKey)
	if err != nil {
		t.Error("Failed to sign message:", err)
	}

}

func TestCheck(t *testing.T) {
	check(nil)
}
