package main

import (
  "github.com/Oluwaseun241/hexlock/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
// func main() {
//   inputFilePath := flag.String("i", "", "input file path(comma-seprated)")
//   outputFilePath := flag.String("o", "", "output file path(comma-seprated)")
//   mode := flag.String("mode", "encrypt", "encryption mode(encrypt, decrypt, compress)")
//
//   flag.Parse()
//
//   if *inputFilePath == "" || *outputFilePath == "" {
//     fmt.Println("no input")
//     flag.PrintDefaults()
//     return
//   }
//
//   inputPaths := splitFilePaths(*inputFilePath)
//   outputPaths := splitFilePaths(*outputFilePath)
//
//   if len(inputPaths) != len(outputPaths) {
//     fmt.Println("input and output file counts do not match")
//     return
//   }
//   
//   key := []byte("WGcDZK7dekM06L4ORZpTcigfn6NLD9hG") 
//  
//   totalFiles := len(inputPaths)
//   progressBar := pb.StartNew(totalFiles)
//
//   var err error 
//   switch *mode {
//   case "encrypt":
//     for i := 0; i < totalFiles; i++ {
//       err = cmd.EncryptFile(inputPaths[i], outputPaths[i], key)
//       if err != nil {
//         fmt.Println("Error", err)
//         return
//       }
//       progressBar.Increment()
//     }
//   // Decrypt
//   case "decrypt":
//     for i := 0; i < totalFiles; i++ {
//       cmd.DecryptFile(inputPaths[i], outputPaths[i], key)
//       progressBar.Increment()
//     }
//   // Compress
//   case "compress":
//   for i := 0; i < totalFiles; i++ {
//       err = cmd.CompressFile(inputPaths[i],outputPaths[i]+".gz")
//       if err != nil {
//         progressBar.Finish()
//         fmt.Println("Error", err)
//         return
//       }
//       progressBar.Increment()
//     }
//   default:
//     fmt.Println("Invalid mode")
//     flag.PrintDefaults()
//     return
//   }
//
//   progressBar.Finish()
//   fmt.Println("Operation sucessfull!")
// }
//
// func splitFilePaths(filePaths string) []string {
//   return strings.Split(filePaths, ",")
// }
