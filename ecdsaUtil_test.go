package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"reflect"
	"testing"
)

func TestEncodeDecodePriv(t *testing.T) {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// encode privKey
	encodedPriv := encodePriv(privKey)

	// decode privKey
	decodedPriv := decodePriv(encodedPriv)

	// compare decoded privKey to original
	if !reflect.DeepEqual(privKey, decodedPriv) {
		t.Error("Decoded private key does not match generated private key")
	}

}

func TestEncodePub(t *testing.T) {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubKey := &privKey.PublicKey

	// encode pubKey
	encodedPub := encodePub(pubKey)

	// decode pubKey
	blockPub, _ := pem.Decode([]byte(encodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	decodedPubKey := genericPublicKey.(*ecdsa.PublicKey)

	// compare decoded pubKey to original
	if !reflect.DeepEqual(pubKey, decodedPubKey) {
		t.Error("Public keys do not match.")
	}
}
