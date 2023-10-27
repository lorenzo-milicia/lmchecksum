package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"go.lorenzomilicia.dev/lmchecksum/app"
	"go.lorenzomilicia.dev/lmchecksum/cmd/list"
	"os"
	"slices"
)

// flags
var (
	hfFlag = hashFunctionFlag{hashFunction: "sha256"}
)

var rootCmd = &cobra.Command{
	Use:     "lmchecksum <file> <checksum>",
	Short:   "CLI tool for checking the validity of a checksum",
	Long:    "CLI tool for checking the validity of a checksum",
	Args:    cobra.ExactArgs(2),
	Version: "1.1.0",
	Run:     lmchecksum,
}

func lmchecksum(_ *cobra.Command, args []string) {
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(list.Command)

	rootCmd.Flags().VarP(
		&hfFlag,
		"hash-func",
		"f",
		"run \"lmchecksum list\" to see the full list of available hash functions",
	)
	rootCmd.Flags().Var(
		&hfFlag,
		"algorithm",
		"run \"lmchecksum list\" to see the full list of available hash functions",
	)
	_ = rootCmd.Flags().MarkDeprecated("algorithm", "use --hash-func instead")
}

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
	ha := app.HashNamesMap[s]
	if !slices.Contains(app.AvailableHashFunctions, uint(ha)) {
		return errors.New(fmt.Sprintf("hash function %v not availalble", s))
	}
	h.hashFunction = s
	return nil
}
