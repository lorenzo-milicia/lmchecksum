package cmd

import (
	"fmt"
	"go.lorenzomilicia.dev/lmchecksum/app"
	"os"

	"github.com/spf13/cobra"
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
		hashFuncFlag, err := cmd.Flags().GetString("hash-func")
		if err != nil {
			panic(err)
		}
		fileArg := args[0]
		checksumArg := args[1]

		file, err := os.Open(fileArg)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		appArgs := app.New(file, checksumArg, hashFuncFlag)
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
	rootCmd.Flags().StringP("hash-func", "f", "sha256", "Hash function to use for the check")
}
