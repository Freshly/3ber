package common

import (
	"fmt"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd, printStdout bool) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	for {
		buf := make([]byte, 1024)
		_, err := stdout.Read(buf)
		if printStdout {
			fmt.Print(string(buf))
		}
		if err != nil {
			break
		}
	}

	return cmd.Wait()
}
