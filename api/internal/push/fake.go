package push

import (
	"context"
	"sync"
)

// Fake is a Sender for tests: it records every message and answers ErrInvalidToken for the
// tokens in Invalid, and Err for every other send when set.
type Fake struct {
	mu      sync.Mutex
	sent    []Message
	Invalid map[string]bool
	Err     error
}

var _ Sender = (*Fake)(nil)

// Send records m.
func (f *Fake) Send(_ context.Context, m Message) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sent = append(f.sent, m)
	if f.Invalid[m.Token] {
		return ErrInvalidToken
	}
	return f.Err
}

// Sent is a copy of every message sent so far.
func (f *Fake) Sent() []Message {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]Message(nil), f.sent...)
}
