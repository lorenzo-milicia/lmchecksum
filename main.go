package main

import (
	"crypto"
	_ "crypto/sha256"
	"log"
	"os"

	"go.lorenzomilicia.dev/libs/checksum"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("Not enough arguments passed")
	}
	file, err := os.ReadFile(args[1])

	if err != nil {
		log.Fatalf("Cannot read file %v", args[1])
	}

	checksum.Checksum(file, args[2], crypto.SHA256)
}
