package llm

import (
	"context"
	"errors"
	"sync"
)

// ErrFakeExhausted is returned when a Fake has no scripted answer left.
var ErrFakeExhausted = errors.New("llm: fake has no response left")

// Fake is a scripted Client for tests: call i answers Errs[i] when set and
// non-nil (beside Responses[i] when there is one, as a billed failure carries
// its cost), else Responses[i]. Every Request is recorded in Calls. Complete is
// safe for concurrent use; read Calls only after the calls have finished.
type Fake struct {
	mu        sync.Mutex
	Responses []Response
	Errs      []error
	Calls     []Request
}

// Complete records the request and replays the next scripted answer.
func (f *Fake) Complete(_ context.Context, req Request) (Response, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	i := len(f.Calls)
	f.Calls = append(f.Calls, req)
	if i < len(f.Errs) && f.Errs[i] != nil {
		var res Response
		if i < len(f.Responses) {
			res = f.Responses[i]
		}
		return res, f.Errs[i]
	}
	if i < len(f.Responses) {
		return f.Responses[i], nil
	}
	return Response{}, ErrFakeExhausted
}
