package cmd_test

import (
	"testing"

	"github.com/Oluwaseun241/hexlock/cmd"
)

func TestDecryptCmd(t *testing.T) {
  t.Run("DecryptCmd", func(t *testing.T) {
    cmd := cmd.NewRootCmd()
    cmd.SetArgs([]string{"decrypt", "-i", "input1.txt,input2.txt", "-o", "output1,output2"})
		err := cmd.Execute()
		if err != nil {
			t.Errorf("DecryptCmd failed: %v", err)
		}
  })
}
