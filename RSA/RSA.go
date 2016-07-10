package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func getPrivRSAkey() *rsa.PrivateKey {
	PrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return PrivateKey
}

func getPubRSAkey(PrivateKey *rsa.PrivateKey) *rsa.PublicKey {

	PublicKey := &PrivateKey.PublicKey
	return PublicKey
}

func main() {

	user1PrivateKey := make([]byte, 2048)
	user1PublicKey := make([]byte, 2048)

	user1PrivateKey = getPrivRSAkey()
	user1PublicKey = getPubRSAkey(user1PrivateKey)

	fmt.Println("user1 Private Key : ", user1PrivateKey)
	fmt.Println("user1 Public key ", user1PublicKey)
	fmt.Println("user2 Private Key : ", user2PrivateKey)
	fmt.Println("user2 Public key ", user2PublicKey)

	//Encrypt user1's Message
	message := []byte("Hi user2")
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, user2PublicKey, message, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(message), ciphertext)
	fmt.Println()

	// Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, user1PrivateKey, newhash, hashed, &opts)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("PSS Signature : %x\n", signature)

	// Decrypt Message
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, user2PrivateKey, ciphertext, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", ciphertext, plainText)

	//Verify Signature
	err = rsa.VerifyPSS(user1PublicKey, newhash, hashed, signature, &opts)

	if err != nil {
		fmt.Println("Who are U? Verify Signature failed")
		os.Exit(1)
	} else {
		fmt.Println("Verify Signature successful")
	}

}
