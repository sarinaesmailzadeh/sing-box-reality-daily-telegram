package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func WriteFile(filename string, newReality RealityJson) error {

	file, err := json.MarshalIndent(newReality, "", " ")
	if err != nil {
		log.Fatal("Error during MarshalIndent(): ", err)
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Error during WriteFile(): ", err)
		return err
	}

	return nil
}

func getPublicKeyAndPrivateKey() (privateKey string, publicKey string) {
	dat, err := os.ReadFile("./key_pair.txt")
	if err != nil {
		log.Fatal("error during the ReadFile")
	}
	allData := string(dat)

	allData = strings.TrimSpace(allData)

	privateKey = strings.TrimLeft(strings.TrimRight(allData, " Public key: "), "Private key: ")
	pubAns := strings.SplitAfter(allData, " Public key: ")

	fmt.Println("_____________________________")
	fmt.Println("Public_key:", pubAns[0])
	fmt.Println("Private_key:", privateKey)
	fmt.Println("_____________________________")

	return privateKey, pubAns[0]
}
