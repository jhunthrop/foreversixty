package pipeline

import (
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/queue"
)

// The status page names what holds the queue only once the head has
// failed; a head that has never been tried is just the next upload.
func TestHeldNamesTheFightAtTheHeadOnceItHasFailed(t *testing.T) {
	q, err := queue.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if h := heldBy(q); h != nil {
		t.Fatalf("empty queue held by %+v", h)
	}
	q.Enqueue(queue.Item{Kind: queue.Fight, ReportKey: "k", FightIndex: 3}, []byte("x"))
	if h := heldBy(q); h != nil {
		t.Fatalf("an untried head is held: %+v", h)
	}
	for range 2 {
		l, err := q.Head()
		if err != nil {
			t.Fatal(err)
		}
		if err := l.Fail(errors.New("503 Service Unavailable")); err != nil {
			t.Fatal(err)
		}
	}
	h := heldBy(q)
	if h == nil || h.Kind != string(queue.Fight) || h.Fight != 4 || h.Attempts != 2 {
		t.Fatalf("held = %+v, want fight 4 after 2 tries", h)
	}
}
