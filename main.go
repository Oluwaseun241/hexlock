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

  inputPaths := splitFilePaths(*inputFilePath)
  outputPaths := splitFilePaths(*outputFilePath)

  if len(inputPaths) != len(outputPaths) {
    fmt.Println("input and output file counts do not match")
    return
  }
  
  key := []byte("WGcDZK7dekM06L4ORZpTcigfn6NLD9hG") 
  /* key := generateRandomKey()  */
  var err error 
  switch *mode {
  case "encrypt":
    for i := 0; i < len(inputPaths); i++ {
      err = cmd.EncryptFile(inputPaths[i], outputPaths[i], key)
    }
    //err = cmd.EncryptFile(inputPaths, outputPaths, key)
  case "decrypt":
    for i := 0; i < len(inputPaths); i++ {
      cmd.DecryptFile(inputPaths[i], outputPaths[i], key)
    }
    //err = cmd.DecryptFile(inputPaths, outputPaths, key)
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

func splitFilePaths(filePaths string) []string {
  return filepath.SplitList(filePaths)
}
// func generateRandomKey() []byte {
//   key := make([]byte, 32)
//   if _, err := rand.Read(key); err != nil {
//     panic(err)
//   }
//   return key
// }
