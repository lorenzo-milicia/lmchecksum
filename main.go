package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"go.lorenzomilicia.dev/lmchecksum/app"
	"os"
)

type hashingAlgorithmFlag struct {
	algorithm crypto.Hash
}

func (h *hashingAlgorithmFlag) String() string {
	return h.algorithm.String()
}

func (h *hashingAlgorithmFlag) Set(s string) error {
	switch s {
	case "sha256":
		h.algorithm = crypto.SHA256
	case "sha1":
		h.algorithm = crypto.SHA1
	case "md5":
		h.algorithm = crypto.MD5
	case "sha512":
		h.algorithm = crypto.SHA512
	default:
		return errors.New(fmt.Sprintf("Hashing algorithm %v not supported.\n", s))
	}
	return nil
}

func main() {
	algFlag := hashingAlgorithmFlag{crypto.SHA256}
	flag.Var(&algFlag, "algorithm", "Hashing algorithm for the check (default is SHA256)")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Expecting exactly two arguments (<file> <checksum>")
		os.Exit(1)
	}

	fileArg := flag.Arg(0)
	if fileArg == "" {
		fmt.Println("Expecting a file path")
		os.Exit(1)
	}
	checksumArg := flag.Arg(1)
	if checksumArg == "" {
		fmt.Println("Expecting a hash string to compare")
		os.Exit(1)
	}
	file, err := os.Open(fileArg)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	appArgs := app.New(file, checksumArg, algFlag.algorithm)
	app.VerifyChecksum(appArgs)
	os.Exit(0)
}
