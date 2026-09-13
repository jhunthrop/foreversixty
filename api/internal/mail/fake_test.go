package mail

import (
	"context"
	"errors"
	"testing"
)

func TestFakeRecordsSent(t *testing.T) {
	f := &Fake{}
	msg1 := Message{To: "user1@example.com", Subject: "First", Text: "body1"}
	msg2 := Message{To: "user2@example.com", Subject: "Second", Text: "body2"}

	err1 := f.Send(context.Background(), msg1)
	if err1 != nil {
		t.Fatalf("first send: %v", err1)
	}
	err2 := f.Send(context.Background(), msg2)
	if err2 != nil {
		t.Fatalf("second send: %v", err2)
	}

	if len(f.Sent) != 2 {
		t.Errorf("Sent length = %d, want 2", len(f.Sent))
	}
	if f.Sent[0] != msg1 {
		t.Errorf("Sent[0] = %v, want %v", f.Sent[0], msg1)
	}
	if f.Sent[1] != msg2 {
		t.Errorf("Sent[1] = %v, want %v", f.Sent[1], msg2)
	}
}

func TestFakeErrShortCircuits(t *testing.T) {
	testErr := errors.New("test error")
	f := &Fake{Err: testErr}

	msg := Message{To: "user@example.com", Subject: "Test", Text: "body"}
	err := f.Send(context.Background(), msg)

	if err != testErr {
		t.Errorf("Send returned %v, want %v", err, testErr)
	}
	if len(f.Sent) != 0 {
		t.Errorf("Sent length = %d, want 0", len(f.Sent))
	}
}
