package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.lorenzomilicia.dev/libs/checksum/v2"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "lmchecksum <file> <checksum>",
	Short:   "CLI tool for checking the validity of a checksum",
	Long:    "CLI tool for checking the validity of a checksum",
	Args:    cobra.ExactArgs(2),
	Version: "1.1.0",
	RunE:    lmchecksum,
}

func lmchecksum(_ *cobra.Command, args []string) error {
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("lmchecksum v%v\n", rootCmd.Version))
	rootCmd.AddCommand(Command)
	rootCmd.Flags().Var(&hfFlag, "hash-function", "set the hash function to use")
	rootCmd.Flags().Var(&hfFlag, "algorithm", "set the hash function to use")
	_ = rootCmd.Flags().MarkDeprecated("algorithm", "use --hash-function instead")
	rootCmd.MarkFlagsMutuallyExclusive("hash-function", "algorithm")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
