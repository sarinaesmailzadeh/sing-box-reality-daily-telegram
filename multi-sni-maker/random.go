package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	result := string(hex.EncodeToString(b))
	fmt.Println("*************")
	fmt.Println(result) //"e97a333fee31c8d0"
	return result
}
