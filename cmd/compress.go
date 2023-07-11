package cmd

import (
	"compress/gzip"
	"os"
)

func CompressFile(inputFilePath, outputFilePath string) error {
  input, err := os.ReadFile(inputFilePath)
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
