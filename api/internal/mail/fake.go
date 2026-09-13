package mail

import (
	"context"
	"sync"
)

// Fake is a test double for Mailer. It is safe for concurrent use since
// Service.Subscribe sends confirmations from a background goroutine.
type Fake struct {
	mu   sync.Mutex
	Sent []Message
	Err  error
}

func (f *Fake) Send(_ context.Context, m Message) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.Sent = append(f.Sent, m)
	return nil
}
