package llm

import (
	"context"
	"errors"
	"testing"
)

func TestFake_RecordsAndReplays(t *testing.T) {
	boom := errors.New("boom")
	f := &Fake{
		Responses: []Response{{Content: "r1"}, {}, {Content: "r3"}},
		Errs:      []error{nil, boom},
	}
	var _ Client = f
	reqs := []Request{
		{Messages: []Message{{Role: "user", Content: "a"}}},
		{Messages: []Message{{Role: "user", Content: "b"}}},
		{Messages: []Message{{Role: "user", Content: "c"}}},
		{Messages: []Message{{Role: "user", Content: "d"}}},
	}
	ctx := context.Background()
	if r, err := f.Complete(ctx, reqs[0]); err != nil || r.Content != "r1" {
		t.Fatalf("%+v %v", r, err)
	}
	if _, err := f.Complete(ctx, reqs[1]); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if r, err := f.Complete(ctx, reqs[2]); err != nil || r.Content != "r3" {
		t.Fatalf("%+v %v", r, err)
	}
	if _, err := f.Complete(ctx, reqs[3]); !errors.Is(err, ErrFakeExhausted) {
		t.Fatal(err)
	}
	if len(f.Calls) != 4 || f.Calls[1].Messages[0].Content != "b" {
		t.Fatalf("calls %+v", f.Calls)
	}
}
