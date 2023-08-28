package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "moat",
	Short: "moat is a passphrase generator",
}

var delimiter string
var length int
var words int

func init() {
	RootCmd.AddCommand(dicewareCmd)
	RootCmd.AddCommand(sillypassCmd)

	dicewareCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "", "Separator between the words of the Diceware passphrase.")
	dicewareCmd.Flags().IntVarP(&words, "words", "w", 6, "The number of words in the Diceware passphrase.")
	sillypassCmd.Flags().IntVarP(&length, "length", "l", 12, "The number of characters in the sillypass password.")
}
