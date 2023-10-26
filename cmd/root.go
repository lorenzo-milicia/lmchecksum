package cmd

import (
	"errors"
	"fmt"
	"go.lorenzomilicia.dev/lmchecksum/app"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

// flags
var (
	hfFlag hashFunctionFlag = hashFunctionFlag{hashFunction: "sha256"}
)

var rootCmd = &cobra.Command{
	Use:     "lmchecksum <file> <checksum>",
	Short:   "CLI tool for checking the validity of a checksum",
	Long:    "CLI tool for checking the validity of a checksum",
	Args:    cobra.ExactArgs(2),
	Version: "1.1.0",
	Run:     lmchecksum(),
}

func lmchecksum() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fileArg := args[0]
		checksumArg := args[1]

		file, err := os.Open(fileArg)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		appArgs := app.New(file, checksumArg, hfFlag.hashFunction)
		app.VerifyChecksum(appArgs)
		os.Exit(0)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().VarP(
		&hfFlag,
		"hash-func",
		"f",
		fmt.Sprintf("supported: %v", strings.Join(app.SupportedHashFunctions, ", ")),
	)
	rootCmd.Flags().String("algorithm", "", "Hash function to use for the check")
	_ = rootCmd.Flags().MarkDeprecated("algorithm", "use --hash-func instead")
}

type hashFunctionFlag struct {
	hashFunction string
}

func (h *hashFunctionFlag) String() string {
	return h.hashFunction
}
func (h *hashFunctionFlag) Type() string {
	return "enum"
}

func (h *hashFunctionFlag) Set(s string) error {
	if !slices.Contains(app.SupportedHashFunctions, s) {
		return errors.New(fmt.Sprintf("must be one of %v", strings.Join(app.SupportedHashFunctions, ", ")))
	}
	h.hashFunction = s
	return nil
}
