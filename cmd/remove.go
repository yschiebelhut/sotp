/*
Copyright © 2024 Yannik Schiebelhut <yannik.schiebelhut@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yschiebelhut/sotp/data"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove <secret name>",
	Short:   "Remove a secret",
	Aliases: []string{"r"},
	Args:    cobra.OnlyValidArgs,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		keys, err := data.GetAllSecrets()
		if err != nil {
			cmd.PrintErrln(err)
		}
		return keys, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			err := data.RemoveSecret(arg)
			if err != nil {
				cmd.PrintErrln(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
