package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/sha3"
)

func main() {
	// Read the JSON data from a file
	filePath := "./genesis.json"
	dataBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Calculate the Keccak-256 hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write(dataBytes)
	hashSum := hash.Sum(nil)

	// Print the hash
	fmt.Printf("Keccak-256 Hash: 0x%s\n", hex.EncodeToString(hashSum))
}
