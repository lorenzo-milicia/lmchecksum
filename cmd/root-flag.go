package cmd

import (
	"crypto"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var (
	hfFlag = hashFunctionFlag{"sha256"}
)

type hashFunctionFlag struct {
	hashFunction string
}

func (h *hashFunctionFlag) String() string {
	return h.hashFunction
}
func (h *hashFunctionFlag) Type() string {
	return ""
}

func (h *hashFunctionFlag) Set(s string) error {
	ha := HashNamesMap[s]
	if !slices.Contains(AvailableHashFunctions, uint(ha)) {
		return errors.New(fmt.Sprintf("hash function %v not availalble", s))
	}
	h.hashFunction = s
	return nil
}

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
