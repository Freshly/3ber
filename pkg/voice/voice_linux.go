//go:build linux

package voice

import (
	"fmt"
	"os/exec"
)

// TODO detect available voices
var VOICES = []string{
	"english",
	"english-us",
	"en-scottish",
}
var VOICE = VOICES[0]
var VOICE_CMD = "espeak"

func say(message string) error {
	fmt.Println(message)

	cmd := exec.Command(VOICE_CMD,
		"-v", VOICE,
		message)

	return runCommand(cmd)
}
