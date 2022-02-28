package voice

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	VOICE = randomVoice()
}

func Say(message string) {
	resolve()
	if viper.GetBool("3BER_VOICE") {
		err := say(message)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(message)
	}
}

func resolve() {
	if viper.GetBool("3BER_VOICE") {
		VOICE_CMD = lookPath(VOICE_CMD)
	}
}

// lookPath returns the command filepath or disables the voice synthesizer if the command is not found
func lookPath(command string) string {
	path, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("WARNING: '%s' not found in PATH, disabling voice synthesizer\n", command)
		viper.Set("3BER_VOICE", "false")
	}
	return path
}

func randomVoice() string {
	return VOICES[rand.Int()%len(VOICES)]
}

func runCommand(cmd *exec.Cmd) error {
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
