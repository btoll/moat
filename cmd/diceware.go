package cmd

import (
	"fmt"
	"os"

	"github.com/btoll/diceware"
	"github.com/spf13/cobra"
)

var dicewareCmd = &cobra.Command{
	Use:   "diceware",
	Short: "Create a Diceware passphrase",
	Run: func(cmd *cobra.Command, args []string) {
		if words <= 0 || words > 500 {
			fmt.Fprintln(os.Stderr, "Diceware passphrases must be between 1 and 500 words")
			os.Exit(1)
		}

		fmt.Println(diceware.Generate(words, delimiter))
	},
}
