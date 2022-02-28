//go:build !darwin

package voice

import "fmt"

func say(message string) error {
	fmt.Println(message)
	return nil
}
