package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexlock",
	Short: "A CLI tool for file encryption, decryption, and compression",
	Long:  "HexLock is a CLI tool that allows you to encrypt, decrypt, and compress files.",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Use a subcommand. Run 'hexlock --help' for more information.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
