package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

func EncryptFile(inputFilePath, outputFilePath string,key []byte) error {
  input, err := os.ReadFile(inputFilePath)
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
  err = os.WriteFile(outputFilePath, encryptedData, 0777)
  if err != nil {
    return err
  }
  return nil
}

func DecryptFile(inputFilePath, outputFilePath string, key []byte) error {
  input, err := os.ReadFile(inputFilePath) 
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

  err = os.WriteFile(outputFilePath, decryptedData, 0777)
  if err != nil {
    return err
  }
  return nil
}
