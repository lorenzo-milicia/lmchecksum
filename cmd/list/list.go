package list

import (
	"crypto"
	"fmt"
	"github.com/spf13/cobra"
	cmd2 "go.lorenzomilicia.dev/lmchecksum/cmd"
	"regexp"
	"strings"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "Print list of available hash functions",
	Long:  "Print list of available hash functions",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		for _, h := range cmd2.AvailableHashFunctions {
			str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(crypto.Hash(h).String(), "")
			fmt.Println(strings.ToLower(str))
		}
	},
}
