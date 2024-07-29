/*
Copyright © 2024 Yannik Schiebelhut <yannik.schiebelhut@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yschiebelhut/sotp/data"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <secret name>",
	Short: "Add a secret",
	Long: `Add a secret to the secret store.
The secret will be read from stdin.`,
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Please enter the secret:")
		var secret string
		fmt.Scanln(&secret)
		err := data.AddSecret(args[0], secret)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		cmd.Println("Added secret")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
