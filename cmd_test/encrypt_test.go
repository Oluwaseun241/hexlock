package cmd_test

import (
	"testing"

	"github.com/Oluwaseun241/hexlock/cmd"
)

func TestEncryptCmd(t *testing.T) {
  t.Run("EncryptCmd", func(t *testing.T) {
    cmd := cmd.NewRootCmd()
    cmd.SetArgs([]string{"encrypt", "-i", "input1.txt,input2.txt", "-o", "output1,output2"})
		err := cmd.Execute()
		if err != nil {
			t.Errorf("EncryptCmd failed: %v", err)
		}
  })
}
