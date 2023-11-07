package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.lorenzomilicia.dev/libs/checksum/v2"
)

var HashCmd = &cobra.Command{
	Use:   "hash <file>",
	Short: "Computes the hash of the given file",
	Long:  "Computes the hash of the given file",
	Args:  cobra.ExactArgs(1),
	RunE:  hash,
}

func initHash() {
	HashCmd.Flags().Var(&hfFlag, "hash-function", "set the hash function to use")
}

func hash(_ *cobra.Command, args []string) error {
	fileArg := args[0]

	file, err := os.Open(fileArg)
	if err != nil {
		return err
	}
	hf := HashNamesMap[hfFlag.hashFunction]
	hash, err := checksum.Hash(file, hf)
	if err != nil {
		return err
	}
	fmt.Println(hash)
	return nil
}
