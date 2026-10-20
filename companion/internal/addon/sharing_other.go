//go:build !windows

package addon

// isSharingViolation is false off Windows: a rename over an open file
// succeeds there.
func isSharingViolation(error) bool { return false }
