package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Replace "seed phrase" with your actual seed phrase
	seedPhrase := "candy maple cake sugar pudding cream honey rich smooth crumble sweet treat"

	// Generate a new keystore
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore("./keystore", scryptN, scryptP)

	// Create a new account with the seed phrase
	account, err := ks.NewAccount(seedPhrase)
	if err != nil {
		fmt.Printf("Error creating new account: %s", err)
		return
	}

	// Get the private key from the keystore
	privateKey, err := ks.GetPrivateKey(account, seedPhrase)
	if err != nil {
		fmt.Printf("Error getting private key: %s", err)
		return
	}

	// Generate the Ethereum address from the private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Print the address
	fmt.Printf("Address: %s\n", address.Hex())
}
