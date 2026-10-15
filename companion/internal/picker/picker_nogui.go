//go:build nogui

package picker

// Pick has no dialog to show in the headless build.
func Pick() (string, error) { return "", ErrUnavailable }
