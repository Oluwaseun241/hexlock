package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"

	"github.com/Oluwaseun241/hexlock/internal"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt files",
	Long:  "Encrypt files using AES-GCM encryption.",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Encrypting files...")
		
    inputPaths, _ := cmd.Flags().GetStringSlice("input")
		outputPaths, _ := cmd.Flags().GetStringSlice("output")
    key := internal.GenerateKey()    

    progress := progressbar.NewOptions(len(inputPaths),
      progressbar.OptionSetDescription("[cyan][Encrypting files...][reset]"),
      progressbar.OptionSetWriter(os.Stderr),
      progressbar.OptionShowCount(),
      progressbar.OptionShowBytes(true),
      progressbar.OptionEnableColorCodes(true),
      )

    for i := 0; i < len(inputPaths); i++ {
			err := EncryptFile(inputPaths[i], outputPaths[i], key)
			if err != nil {
				color.Red("Error encrypting: %v", err)
			} else {
        progress.Add(1)
				color.Green("\nEncryption successful for %s", inputPaths)
			}
		}
    progress.Finish()
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
	DecryptCmd.MarkFlagRequired("input")
  DecryptCmd.MarkFlagRequired("output")
  rootCmd.AddCommand(EncryptCmd)
}
