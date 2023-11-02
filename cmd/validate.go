package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.lorenzomilicia.dev/libs/checksum/v2"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check the validity of a checksum",
	Long:  "Check the validity of a checksum",
	Args:  cobra.ExactArgs(2),
	RunE:  validate,
}

func initValidate() {
	validateCmd.Flags().Var(&hfFlag, "hash-function", "set the hash function to use")
}

func validate(_ *cobra.Command, args []string) error {
	fileArg := args[0]
	checksumArg := args[1]

	file, err := os.Open(fileArg)
	if err != nil {
		return err
	}
	hf := HashNamesMap[hfFlag.hashFunction]
	res, err := checksum.Validate(file, checksumArg, hf)
	if err != nil {
		return err
	}
	if res == true {
		fmt.Println("[✓] The checksum matches")
	} else {
		fmt.Println("[x] The checksum doesn't match")
	}
	return nil
}
