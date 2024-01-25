package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("alpapie"))
	fmt.Printf("%x", h.Sum([]byte("a")))
	fmt.Println()
	sum := sha256.Sum256([]byte("alpapie"))
	fmt.Printf("%x", sum)
	fmt.Println()
}