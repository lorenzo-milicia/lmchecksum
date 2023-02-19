package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"fmt"
	"os"

	"go.lorenzomilicia.dev/libs/checksum"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("Not enough arguments passed\n")
		os.Exit(1)
	}
	file, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("Cannot read file %v\n", args[1])
		os.Exit(1)
	}

	algorithm := crypto.SHA3_256

	if len(args) > 3 {
		switch hf := args[3]; hf {
		case "sha256":
		case "sha1":
			algorithm = crypto.SHA1
		case "md5":
			algorithm = crypto.MD5
		case "sha512":
			algorithm = crypto.SHA3_512
		default:
			fmt.Printf("Hashing algorithm %v not supported.\n", hf)
			os.Exit(1)
		}
	}

	if res := checksum.Checksum(file, args[2], algorithm); res == true {
		fmt.Print("It checks out ✅\n")
	} else {
		fmt.Printf("Checksums don't match ❌\n")
	}
	os.Exit(0)
}
