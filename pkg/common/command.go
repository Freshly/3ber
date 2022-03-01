package common

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func RunCommand(cmd *exec.Cmd, printStdout bool) error {
	if !viper.GetBool("QUIET") {
		PrintCommand(cmd)
	}
	return RunCommandNoPrint(cmd, printStdout)
}

func RunCommandNoPrint(cmd *exec.Cmd, printStdout bool) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout

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

func RunCommandOutputArray(cmd *exec.Cmd) ([]string, error) {
	if !viper.GetBool("QUIET") {
		PrintCommand(cmd)
	}

	output := []string{}

	r, _ := cmd.StdoutPipe()
	// cmd.Stderr = cmd.Stdout

	err := cmd.Start()
	if err != nil {
		return output, err
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output, cmd.Wait()
}

func PrintCommand(cmd *exec.Cmd) {
	fmt.Printf(">>> %s\n", strings.Join(cmd.Args, " "))
}
