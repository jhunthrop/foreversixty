// companion/internal/picker/picker.go
// Package picker asks the player for a folder with the operating
// system's own dialog. The real dialog is in picker_gui.go; a
// headless build has none.
package picker

import "errors"

// ErrCancelled is a dialog the player closed without choosing.
var ErrCancelled = errors.New("no folder was chosen")

// ErrUnavailable is a build or a machine with no folder dialog.
var ErrUnavailable = errors.New("there is no folder dialog on this machine")
