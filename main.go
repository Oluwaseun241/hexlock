package main

import (	

	"github.com/Oluwaseun241/hexlock/cmd"
	"github.com/fatih/color"
)

func main() {	
  color.Cyan("Welcome to HexLock CLI Tool!")
  color.Yellow("Choose a mode and provide necessary arguments.")
  cmd.Execute()
}
