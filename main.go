package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
  inputFilePath := flag.String("input", "", "input file path")
  outputFilePath := flag.String("output", "", "output file path")
  
  flag.Parse()

  if *inputFilePath == "" || *outputFilePath == "" {
    fmt.Println("no input")
    flag.PrintDefaults()
    return
  }
  
  key := []byte("WGcDZK7dekM06L4ORZpTcigfn6NLD9hG")

  err := encryptFile(*inputFilePath, *outputFilePath, key)
  if err != nil {
    fmt.Println("error",err)
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
