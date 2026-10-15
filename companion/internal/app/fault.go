// companion/internal/app/fault.go
package app

import (
	"errors"
	"net/http"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// Fault is a failure with a code the window can show the player a
// sentence for (copy.go). The cause stays on the error for the log and
// is never sent to the window.
type Fault struct {
	Code string
	Err  error
}

func (f *Fault) Error() string {
	if f.Err == nil {
		return f.Code
	}
	return f.Code + ": " + f.Err.Error()
}

func (f *Fault) Unwrap() error { return f.Err }

func fault(code string, err error) *Fault { return &Fault{Code: code, Err: err} }

// codeOf is the failure code of any error: a Fault's own, else the
// fallback.
func codeOf(err error, fallback string) string {
	var f *Fault
	if errors.As(err, &f) {
		return f.Code
	}
	return fallback
}

// netCode sorts an error from the API client into what the player can
// do about it. Anything that is not an API answer is the network. A
// refusal (a 4xx) is the code the caller gives for it: for a pairing
// code that is bad_code, for everything else the site is at fault.
func netCode(err error, refused string) string {
	var ae *client.Error
	if !errors.As(err, &ae) {
		return CodeOffline
	}
	if ae.Status >= 400 && ae.Status < 500 && ae.Status != http.StatusTooManyRequests &&
		ae.Status != http.StatusRequestTimeout {
		return refused
	}
	return CodeServer
}
