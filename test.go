package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	keySize, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("RSA %v BENCHMARK\n", keySize)
	privateKey, _ := rsa.GenerateKey(rand.Reader, keySize)

	// Encode public key to PKIX format
	pub := privateKey.Public().(*rsa.PublicKey)

	key := []byte("XXXXXXXXXXXXXXXX")                                                    // 16 chars
	secret := []byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX") // 64 chars

	count := 0
	timeLimit := 10 * time.Second
	var encryptedKey []byte
	var encryptedSecret []byte
	for start := time.Now(); time.Since(start) < timeLimit; {
		encryptedKey, _ = rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, key, nil)
		encryptedSecret, _ = rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, secret, nil)
		count++
	}

	fmt.Printf("managed to encrypt %v keys in %v seconds; encryption rate %v keys/s \n", count, timeLimit, count/10)

	count = 0
	for start := time.Now(); time.Since(start) < timeLimit; {
		rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedKey, nil)
		rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedSecret, nil)
		count++
	}

	fmt.Printf("managed to decrypt %v keys in %v seconds; decryption rate %v keys/s \n", count, timeLimit, count/10)
}
