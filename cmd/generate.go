/*
Copyright © 2024 Yannik Schiebelhut <yannik.schiebelhut@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yschiebelhut/sotp/data"
	"github.com/yschiebelhut/sotp/otp"
	"golang.design/x/clipboard"
)

var generateNoClipboardCopy bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate <secret name>",
	Aliases: []string{"gen", "g"},
	Short:   "Generate an OTP code",
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		keys, err := data.GetAllSecrets()
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		return keys, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		secret, err := data.LookupSecret(args[0])
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		code, err := otp.GenerateOTP(secret)
		if err != nil {
			cmd.PrintErrln("error generating code:", err)
		}

		cmd.Println("OTP code:", code)

		if !generateNoClipboardCopy {
			err = clipboard.Init()
			if err != nil {
				cmd.PrintErrln("failed initializing clipboard:", err)
				os.Exit(1)
			}
			clipboard.Write(clipboard.FmtText, []byte(code))
			cmd.Println("Copied to clipboard")
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().BoolVarP(&generateNoClipboardCopy, "no-clipboard", "n", false, "Do not copy the generated code to the clipboard")
}
