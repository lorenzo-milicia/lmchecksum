package app

import (
	"crypto"
	"fmt"
	"go.lorenzomilicia.dev/libs/checksum"
	"io"
	"regexp"
	"strings"
)

var AvailableHashFunctions = func() []uint {
	var hashes []uint
	for i := uint(1); i < uint(crypto.BLAKE2b_512); i++ {
		if crypto.Hash(i).Available() {
			hashes = append(hashes, i)
		}
	}
	return hashes
}()

var HashNamesMap = func() map[string]crypto.Hash {
	m := make(map[string]crypto.Hash)
	for i := uint(1); i < uint(crypto.BLAKE2b_512); i++ {
		c := crypto.Hash(i)
		name := c.String()
		str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, "")
		m[strings.ToLower(str)] = c
	}
	return m
}()

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
	hf := HashNamesMap[args.Algorithm]
	if res := checksum.Checksum(args.File, args.Hash, hf); res == true {
		fmt.Println("[✓] The checksum matches")
	} else {
		fmt.Println("[x] The checksum doesn't match")
	}
}
