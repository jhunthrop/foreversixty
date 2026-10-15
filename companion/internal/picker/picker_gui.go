//go:build !nogui

package picker

import (
	"errors"

	"github.com/ncruces/zenity"
)

// Pick shows the system folder dialog and returns the chosen path. It
// blocks until the player answers, so the caller runs it off the
// window's thread.
func Pick() (string, error) {
	p, err := zenity.SelectFile(
		zenity.Directory(),
		zenity.Title("Choose your World of Warcraft folder"),
	)
	switch {
	case errors.Is(err, zenity.ErrCanceled):
		return "", ErrCancelled
	case errors.Is(err, zenity.ErrUnsupported):
		return "", ErrUnavailable
	case err != nil:
		return "", err
	}
	return p, nil
}
