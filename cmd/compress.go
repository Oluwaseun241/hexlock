package cmd

import (
	"compress/gzip"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var CompressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress files",
	Long:  "Compress files using Gzip compression.",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Compressing files...")
		
    inputPaths, _ := cmd.Flags().GetStringSlice("input")
		outputPaths, _ := cmd.Flags().GetStringSlice("output")
		
    for i := 0; i < len(inputPaths); i++ {
			err := CompressFile(inputPaths[i], outputPaths[i]+".gz")
			if err != nil {
				color.Red("Error compressing: %v", err)
			} else {
				color.Green("Compression successful for %s", inputPaths)
			}
		}
	},
}

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

func init() {
	CompressCmd.Flags().StringSliceP("input", "i", []string{}, "input file paths (comma-separated)")
	CompressCmd.Flags().StringSliceP("output", "o", []string{}, "output file paths (comma-separated)")
	CompressCmd.MarkFlagRequired("input")
  CompressCmd.MarkFlagRequired("output")
  rootCmd.AddCommand(CompressCmd)
}
