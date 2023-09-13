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
	allData = strings.ReplaceAll(allData, " ", "")

	privateKeyFirst := RemoveRightPart(allData, "Publickey:")
	privateKey = RemoveLeftPart(privateKeyFirst, "Privatekey:")

	pubAns := strings.SplitAfter(allData, "Publickey:")
	publicKey = pubAns[1]

	fmt.Println("_____________________________")
	fmt.Println("Private_key=>", privateKey)
	fmt.Println("Public_key=>", publicKey)
	fmt.Println("_____________________________")

	return privateKey, publicKey
}

func RemoveRightPart(str, substring string) string {
	return str[:strings.Index(str, substring)]
}

func RemoveLeftPart(str, substring string) string {
	return str[strings.Index(str, substring)+len(substring):]
}
