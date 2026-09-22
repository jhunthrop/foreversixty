// api/internal/bnetapi/errors.go
package bnetapi

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrNotFound is a 404: the namespace, realm, character, or guild named
// does not exist (or, for a character, does not exist under the
// namespace queried).
var ErrNotFound = errors.New("bnetapi: not found")

// ErrForbidden is a 403: the namespace does not exist for this game/
// region, or the character's profile is private.
var ErrForbidden = errors.New("bnetapi: forbidden")

// ErrRateLimited is a 429: the caller must stop the batch it is running
// and log it, per spec §3.
var ErrRateLimited = errors.New("bnetapi: rate limited")

// statusError classifies a non-200 response into one of the three typed
// errors above, or a generic wrapped status for anything else.
func statusError(op string, status int) error {
	switch status {
	case http.StatusNotFound:
		return fmt.Errorf("bnetapi: %s: %w", op, ErrNotFound)
	case http.StatusForbidden:
		return fmt.Errorf("bnetapi: %s: %w", op, ErrForbidden)
	case http.StatusTooManyRequests:
		return fmt.Errorf("bnetapi: %s: %w", op, ErrRateLimited)
	default:
		return fmt.Errorf("bnetapi: %s: status %d", op, status)
	}
}
