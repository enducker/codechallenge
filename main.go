package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const pemDirectory = "pem"
const pemFile = pemDirectory + "/key.pem"

func main() {

	msg := getMsg()

	// get private key
	privKey := getKey()

	// get public key
	pubKey := encodePub(&privKey.PublicKey)

	// sign message
	sig, err := signMessage(msg, privKey)
	check(err)

	fmt.Println(response{msg, sig, pubKey})
}

// getMsg returns the first command line argument
func getMsg() (msg string) {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		log.Fatal("Missing command line argument: Message is required")
	} else if msg = argsWithoutProg[0]; len(msg) > 250 {
		log.Fatal("Invalid command line argument: Message length must not exceed 250 characters")
	}

	return
}

// getKey checks if keyfile exists in "pem/key.pem". If missing,
// new ECDSA key is created and persisted to file system.
// Returns *ecdsa.PrivateKey from keyfile
func getKey() *ecdsa.PrivateKey {
	// check file system for keyfile
	if _, err := ioutil.ReadFile(pemFile); err != nil {
		// generate new ECDSA key
		privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		check(err)

		encPriv := encodePriv(privKey)

		err = os.MkdirAll(pemDirectory, 600)
		check(err)

		// write PEM to file system
		ioutil.WriteFile(pemFile, []byte(encPriv), 600)
		check(err)
	}

	content, err := ioutil.ReadFile(pemFile)
	check(err)

	// decode file content to *ecdsa.PrivateKey
	privKey := decodePriv(string(content))

	return privKey
}

// signMessage signs msg with key and returns the base64 encoded string of signature or error
func signMessage(msg string, key *ecdsa.PrivateKey) (string, error) {
	// get sha256 hash
	hash := sha256.Sum256([]byte(msg))

	// sign hash
	s, err := key.Sign(rand.Reader, hash[:], crypto.SHA256)
	if err != nil {
		return "", err
	}

	// encode sig
	sig := base64.StdEncoding.EncodeToString(s)

	return sig, nil
}

// check e, exit program if not nil
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
