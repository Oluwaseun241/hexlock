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

  "github.com/Oluwaseun241/hexlock/internal"
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
    err = internal.encryptFile(*inputFilePath, *outputFilePath, key)
  case "decrypt":
    err = internal.decryptFile(*inputFilePath, *outputFilePath, key)
  case "compress":
    err = internal.compressFile(*inputFilePath, *outputFilePath+".gz")
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
