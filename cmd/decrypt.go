package cmd

import (
	"crypto/aes"
	"crypto/cipher"
  "github.com/spf13/cobra"
	"fmt"
	"os"
)

var DecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt files",
	Long:  "Decrypt files using AES-GCM decryption.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Decrypting files...")
		inputPaths, _ := cmd.Flags().GetStringSlice("input")
		outputPaths, _ := cmd.Flags().GetStringSlice("output")
		key := []byte("manchester1") // Replace with your encryption key
		for i := 0; i < len(inputPaths); i++ {
			err := DecryptFile(inputPaths[i], outputPaths[i], key)
			if err != nil {
				fmt.Println("Error decrypting:", err)
			} else {
				fmt.Println("Decryption successful!")
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
	rootCmd.AddCommand(DecryptCmd)
}
