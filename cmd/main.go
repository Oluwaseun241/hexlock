package main

import (
	"flag"
	"fmt"

	"github.com/Oluwaseun241/hexlock/internal"
)

func main() {
  inputFilePath := flag.String("i", "", "input file path")
  outputFilePath := flag.String("o", "", "output file path")
  mode := flag.String("mode", "encrypt", "encryption mode(encrypt, decrypt, compress)")

  flag.Parse()

  if *inputFilePath == "" || *outputFilePath == "" {
    fmt.Println("no input")
    flag.PrintDefaults()
    return
  }
  
  key := []byte("WGcDZK7dekM06L4ORZpTcigfn6NLD9hG") 
  /* key := generateRandomKey()  */
  var err error 
  switch *mode {
  case "encrypt":
    err = internal.EncryptFile(*inputFilePath, *outputFilePath, key)
  case "decrypt":
    err = internal.DecryptFile(*inputFilePath, *outputFilePath, key)
  case "compress":
    err = internal.CompressFile(*inputFilePath, *outputFilePath+".gz")
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

// func generateRandomKey() []byte {
//   key := make([]byte, 32)
//   if _, err := rand.Read(key); err != nil {
//     panic(err)
//   }
//   return key
// }
