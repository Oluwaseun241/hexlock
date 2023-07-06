package main

import (
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
  inputFilePath := flag.String("input", "", "input file path")
  outputFilePath := flag.String("output", "", "output file path")
  mode := flag.String("mode", "encrypt", "encryption mode")

  flag.Parse()

  if *inputFilePath == "" || *outputFilePath == "" {
    fmt.Println("no input")
    flag.PrintDefaults()
    return
  }
  
  key := []byte("WGcDZK7dekM06L4ORZpTcigfn6NLD9hG")
  var err error 
  switch *mode {
  case "encrypt":
    err = encryptFile(*inputFilePath, *outputFilePath, key)
  case "decrypt":
    err = decryptFile(*inputFilePath, *outputFilePath, key)
  case "compress":
    err = compressFile(*inputFilePath, *outputFilePath+".gz")
  default:
    fmt.Println("Invalid mode")
    flag.PrintDefaults()
    return
  }
  if err != nil {
    fmt.Println("Error", err)
    return
  }
  fmt.Println("Done")
}

func encryptFile(inputFilePath, outputFilePath string,key []byte) error {
  input, err := ioutil.ReadFile(inputFilePath)
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

  encryptedData := aesGCM.Seal(nil, nonce, input, nil)
  err = ioutil.WriteFile(outputFilePath, encryptedData, 0644)
  if err != nil {
    return err
  }
  return nil
}

func decryptFile(inputFilePath, outputFilePath string, key []byte) error {
  input, err := ioutil.ReadFile(inputFilePath)
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

  err = ioutil.WriteFile(outputFilePath, decryptedData, 0644)
  if err != nil {
    return err
  }
  return nil
}

func compressFile(inputFilePath, outputFilePath string) error {
  input, err := ioutil.ReadFile(inputFilePath)
  if err != nil {
    return err
  }
  
  outputFile, err := os.Create(outputFilePath)
  if err != nil {
    return err
  }
  defer outputFile.Close()

  gzipWriter := gzip.NewWriter(outputFile)
  if err != nil {
    return err
  }
  defer gzipWriter.Close()
  
  _,err = gzipWriter.Write(input)
  if err != nil {
    return err
  }
  return nil
}
