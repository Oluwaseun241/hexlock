package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
  "github.com/spf13/cobra"
	"fmt"
	"os"
)

var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt files",
	Long:  "Encrypt files using AES-GCM encryption.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Encrypting files...")
		inputPaths, _ := cmd.Flags().GetStringSlice("input")
		outputPaths, _ := cmd.Flags().GetStringSlice("output")
		key := []byte("machester1") // Replace with your encryption key
		for i := 0; i < len(inputPaths); i++ {
			err := EncryptFile(inputPaths[i], outputPaths[i], key)
			if err != nil {
				fmt.Println("Error encrypting:", err)
			} else {
				fmt.Println("Encryption successful!")
			}
		}
	},
}

func EncryptFile(inputPaths, outputPaths string,key []byte) error {
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

  nonce := make([]byte, aesGCM.NonceSize())
  if _, err := rand.Read(nonce); err != nil {
    return err
  }
  
  if _, err := rand.Read(nonce); err != nil {
    return err
  }

  encryptedData := aesGCM.Seal(nonce, nonce, input, nil)
  err = os.WriteFile(outputPaths, encryptedData, 0777)
  if err != nil {
    return err
  }
  return nil
}

func init() {
	EncryptCmd.Flags().StringSliceP("input", "i", []string{}, "input file paths (comma-separated)")
	EncryptCmd.Flags().StringSliceP("output", "o", []string{}, "output file paths (comma-separated)")
	rootCmd.AddCommand(EncryptCmd)
}
