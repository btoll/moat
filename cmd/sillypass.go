package cmd

import (
	"fmt"
	"os"

	"github.com/btoll/sillypass"
	"github.com/spf13/cobra"
)

var sillypassCmd = &cobra.Command{
	Use:   "sillypass",
	Short: "Create a sillypass password",
	Run: func(cmd *cobra.Command, args []string) {
		if length < 3 {
			fmt.Fprintln(os.Stderr, "[Error] Sillypass passwords cannot be less than three characters.")
			os.Exit(1)
		}
		fmt.Println(sillypass.Generate(length))
	},
}
