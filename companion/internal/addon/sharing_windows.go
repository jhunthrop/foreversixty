//go:build windows

package addon

import (
	"errors"
	"syscall"
)

// isSharingViolation reports a rename refused because another process
// has the target open (ERROR_ACCESS_DENIED, ERROR_SHARING_VIOLATION or
// ERROR_LOCK_VIOLATION).
func isSharingViolation(err error) bool {
	var errno syscall.Errno
	if !errors.As(err, &errno) {
		return false
	}
	return errno == 5 || errno == 32 || errno == 33
}
