//go:build linux

package voice

import (
	"fmt"
	"os/exec"

	"github.com/freshly/3ber/pkg/common"
)

// TODO detect available voices
var VOICES = []string{
	"male1",
	"male2",
	"male3",
	"female1",
	"female2",
	"female3",
}
var VOICE = VOICES[0]
var VOICE_CMD = "spd-say"

func say(message string) error {
	fmt.Println(message)

	cmd := exec.Command(VOICE_CMD,
		"-t", VOICE,
		"-w",
		message)

	return common.RunCommandNoPrint(cmd, true)
}
