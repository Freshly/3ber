//go:build windows

package voice

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

var VOICES = []string{""}
var VOICE = ""
var VOICE_CMD = "cmd.exe"

func say(message string) error {
	fmt.Println(message)

	tempFile := fmt.Sprintf("%s\\say_%d.bat", os.TempDir(), rand.Int()%10000)
	contents := fmt.Sprintf(
		"@echo off\nmshta.exe vbscript:Execute(\"CreateObject(\"\"SAPI.SpVoice\"\").Speak(\"\"%s\"\")(window.close)\")",
		message)

	err := os.WriteFile(tempFile, []byte(contents), 0666)
	defer os.Remove(tempFile)

	if err != nil {
		return err
	}

	cmd := exec.Command(VOICE_CMD,
		"/c", tempFile)

	return runCommand(cmd, false)
}
