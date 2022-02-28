//go:build !darwin

package voice

func say(text string) error {
	return nil
}
