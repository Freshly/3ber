//go:build darwin

package voice

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

var VOICES = []string{
	"daniel",
	"fiona",
	"fred",
	"karen",
}
var VOICE = VOICES[0]

func init() {
	rand.Seed(time.Now().UnixNano())
	VOICE = randomVoice()
}

func say(message string) error {
	cmd := exec.Command("say",
		"--voice", VOICE,
		"--interactive",
		message)

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
		fmt.Print(string(buf))
		if err != nil {
			break
		}
	}

	return cmd.Wait()
}

func randomVoice() string {
	return VOICES[rand.Int()%len(VOICES)]
}
