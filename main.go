package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/Oluwaseun241/hexlock/cmd"
)

func main() {
  inputFilePath := flag.String("i", "", "input file path(comma-seprated)")
  outputFilePath := flag.String("o", "", "output file path(comma-seprated)")
  mode := flag.String("mode", "encrypt", "encryption mode(encrypt, decrypt, compress)")

  flag.Parse()

  if *inputFilePath == "" || *outputFilePath == "" {
    fmt.Println("no input")
    flag.PrintDefaults()
    return
  }

  inputPaths := splitFilePath(*inputFilePath)
  outputPaths := splitFilePath(*outputFilePath)

  if len(inputPaths) != len(outputPaths) {
    fmt.Println("input and output file counts do not match")
    return
  }
  
  key := []byte("WGcDZK7dekM06L4ORZpTcigfn6NLD9hG") 
  /* key := generateRandomKey()  */
  var err error 
  switch *mode {
  case "encrypt":
    err = cmd.EncryptFile(*inputFilePath, *outputFilePath, key)
  case "decrypt":
    err = cmd.DecryptFile(*inputFilePath, *outputFilePath, key)
  case "compress":
    err = cmd.CompressFile(*inputFilePath, *outputFilePath+".gz")
  default:
    fmt.Println("Invalid mode")
    flag.PrintDefaults()
    return
  }
  if err != nil {
    fmt.Println("Error", err)
    return
  }
  fmt.Println("Operation sucessfull!")
}

func splitFilePath(filepaths string) []string {
  return filepath.SplitList(filepaths)
}
// func generateRandomKey() []byte {
//   key := make([]byte, 32)
//   if _, err := rand.Read(key); err != nil {
//     panic(err)
//   }
//   return key
// }
