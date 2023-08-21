package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexlock",
	Short: "A CLI tool for file encryption, decryption, and compression",
	Long:  "HexLock is a CLI tool that allows you to encrypt, decrypt, and compress files.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand. Run 'hexlock --help' for more information.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
