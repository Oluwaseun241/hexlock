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
  
  err = compressFile(*outputFilePath, *outputFilePath+".gz")
  if err != nil {
    return
  }
  fmt.Println("File encryption and compression completed sucessfully!")
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

func compressFile(inputFilePath, outputFilePath string) error {
  input, err := ioutil.ReadFile(inputFilePath)
  if err != nil {
    return err
  }
  
  outputFile, err := os.Create(outputFilePath)
  if err != nil {
    return err
  }

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
