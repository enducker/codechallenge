package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

// encodePub returns the PEM encoding of publicKey
func encodePub(publicKey *ecdsa.PublicKey) string {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncodedPub)
}

// encodePriv returns the PEM encoding of privateKey
func encodePriv(privateKey *ecdsa.PrivateKey) string {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	return string(pemEncoded)
}

// decodePriv decodes pemEncoded and returns an ecdsa.PrivateKey
func decodePriv(pemEncoded string) (privateKey *ecdsa.PrivateKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ = x509.ParseECPrivateKey(x509Encoded)

	return
}
