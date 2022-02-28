//go:build darwin

package voice

import (
	"os/exec"
)

// TODO detect available voices
var VOICES = []string{
	"daniel",
	"fiona",
	"fred",
	"karen",
}
var VOICE = VOICES[0]
var VOICE_CMD = "say"

func say(message string) error {
	cmd := exec.Command(VOICE_CMD,
		"--voice", VOICE,
		"--interactive",
		message)

	return runCommand(cmd, true)
}
