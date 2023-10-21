package app

import (
	"crypto"
	"fmt"
	"go.lorenzomilicia.dev/libs/checksum"
	"io"
)

type Arguments struct {
	File      io.Reader
	Hash      string
	Algorithm crypto.Hash
}

func New(file io.Reader, hash string, algorithm crypto.Hash) Arguments {
	return Arguments{
		File:      file,
		Hash:      hash,
		Algorithm: algorithm,
	}
}

func VerifyChecksum(args Arguments) {
	if res := checksum.Checksum(args.File, args.Hash, args.Algorithm); res == true {
		fmt.Print("It checks out ✅\n")
	} else {
		fmt.Printf("Checksums don't match ❌\n")
	}
}
