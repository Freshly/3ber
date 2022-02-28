package voice

import (
	"fmt"

	"github.com/spf13/viper"
)

func Say(message string) {
	if viper.GetBool("3BER_VOICE") {
		err := say(message)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(message)
	}
}
