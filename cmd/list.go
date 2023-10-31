package cmd

import (
	"crypto"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "Print list of available hash functions",
	Long:  "Print list of available hash functions",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		for _, h := range AvailableHashFunctions {
			str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(crypto.Hash(h).String(), "")
			fmt.Println(strings.ToLower(str))
		}
	},
}
