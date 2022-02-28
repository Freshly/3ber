//go:build windows

package voice

import "fmt"

var VOICES = []string{}
var VOICE = ""
var VOICE_CMD = ""

func say(message string) error {
	fmt.Println(message)
	return nil
}
