package mail

import "context"

type Fake struct {
	Sent []Message
	Err  error
}

func (f *Fake) Send(_ context.Context, m Message) error {
	if f.Err != nil {
		return f.Err
	}
	f.Sent = append(f.Sent, m)
	return nil
}
