package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "lmchecksum",
	Short:   "CLI tool for checking the validity of a checksum",
	Long:    "CLI tool for checking the validity of a checksum",
	Args:    cobra.NoArgs,
	Version: "2.0.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(ValidateCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(HashCmd)
	rootCmd.SetVersionTemplate(fmt.Sprintf("lmchecksum v%v\n", rootCmd.Version))
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	initValidate()
	initHash()
}
