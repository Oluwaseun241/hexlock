package cmd_test

import (
	"testing"

	"github.com/Oluwaseun241/hexlock/cmd"
)

func TestCompressCmd(t *testing.T) {
  t.Run("CompressCmd", func(t *testing.T) {
    cmd := cmd.NewRootCmd()
    cmd.SetArgs([]string{"compress", "-i", "input1.txt,input2.txt", "-o", "output1,output2"})
		err := cmd.Execute()
		if err != nil {
			t.Errorf("CompressCmd failed: %v", err)
		}
  })
}
