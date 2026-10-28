package llm

import (
	"context"
	"errors"
)

// ErrFakeExhausted is returned when a Fake has no scripted answer left.
var ErrFakeExhausted = errors.New("llm: fake has no response left")

// Fake is a scripted Client for tests: call i answers Errs[i] when set and
// non-nil, else Responses[i]. Every Request is recorded in Calls.
type Fake struct {
	Responses []Response
	Errs      []error
	Calls     []Request
}

// Complete records the request and replays the next scripted answer.
func (f *Fake) Complete(_ context.Context, req Request) (Response, error) {
	i := len(f.Calls)
	f.Calls = append(f.Calls, req)
	if i < len(f.Errs) && f.Errs[i] != nil {
		return Response{}, f.Errs[i]
	}
	if i < len(f.Responses) {
		return f.Responses[i], nil
	}
	return Response{}, ErrFakeExhausted
}
