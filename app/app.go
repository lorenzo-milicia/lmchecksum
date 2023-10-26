package app

import (
	"crypto"
	"fmt"
	"go.lorenzomilicia.dev/libs/checksum"
	"io"
)

var SupportedHashFunctions = []string{
	"sha1",
	"sha256",
	"sha512",
	"md5",
}

type Arguments struct {
	File      io.Reader
	Hash      string
	Algorithm string
}

func New(file io.Reader, hash string, hashFunction string) Arguments {
	return Arguments{
		File:      file,
		Hash:      hash,
		Algorithm: hashFunction,
	}
}

func VerifyChecksum(args Arguments) {
	hf := hashFunction(args.Algorithm)
	if res := checksum.Checksum(args.File, args.Hash, hf); res == true {
		fmt.Println("[✓] The checksum matches")
	} else {
		fmt.Println("[x] The checksum doesn't match")
	}
}

func hashFunction(s string) crypto.Hash {
	switch s {
	case "sha256":
		return crypto.SHA256
	case "sha1":
		return crypto.SHA1
	case "md5":
		return crypto.MD5
	case "sha512":
		return crypto.SHA512
	default:
		panic(fmt.Sprintf("Hash function %v not supported", s))
	}
}
