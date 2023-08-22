package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var DecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt files",
	Long:  "Decrypt files using AES-GCM decryption.",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Decrypting files...")
		
    inputPaths, _ := cmd.Flags().GetStringSlice("input")
		outputPaths, _ := cmd.Flags().GetStringSlice("output")
		key := []byte("WGcDZK7dekM06L4ORZpTcigfn6NLD9hG")
		
    for i := 0; i < len(inputPaths); i++ {
			err := DecryptFile(inputPaths[i], outputPaths[i], key)
			if err != nil {
				color.Red("Error decrypting: %v", err)
			} else {
				color.Green("Decryption successful for %s", inputPaths)
			}
		}
	},
}

func DecryptFile(inputPaths, outputPaths string, key []byte) error {
  input, err := os.ReadFile(inputPaths) 
  if err != nil {
    return err
  }

  block, err := aes.NewCipher(key)
  if err != nil {
    return err
  }

  aesGCM, err := cipher.NewGCM(block)
  if err != nil {
    return err
  }

  nonceSize := aesGCM.NonceSize()
  if len (input) < nonceSize {
    return fmt.Errorf("invalid encrypted data")
  }

  nonce, input := input[:nonceSize], input[nonceSize:]

  decryptedData, err := aesGCM.Open(nil, nonce, input, nil)
  if err != nil {
    return err
  }

  err = os.WriteFile(outputPaths, decryptedData, 0777)
  if err != nil {
    return err
  }
  return nil
}

func init() {
	DecryptCmd.Flags().StringSliceP("input", "i", []string{}, "input file paths (comma-separated)")
	DecryptCmd.Flags().StringSliceP("output", "o", []string{}, "output file paths (comma-separated)")
	DecryptCmd.MarkFlagRequired("input")
  DecryptCmd.MarkFlagRequired("output")
  rootCmd.AddCommand(DecryptCmd)
}
