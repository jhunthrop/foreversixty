package queue

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestItemsComeBackInSequenceOrder(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for i := range 3 {
		if _, err := q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: i},
			[]byte{byte('a' + i)}); err != nil {
			t.Fatal(err)
		}
	}
	if n, _ := q.Len(); n != 3 {
		t.Fatalf("Len = %d, want 3", n)
	}
	for i := range 3 {
		l, err := q.Head()
		if err != nil {
			t.Fatal(err)
		}
		if l.Item.FightIndex != i || string(l.Body) != string([]byte{byte('a' + i)}) {
			t.Fatalf("head %d = %+v body %q", i, l.Item, l.Body)
		}
		if err := l.Ack(); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := q.Head(); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Head on an empty queue = %v", err)
	}
}

func TestAFailedItemStaysAtTheHeadAndCountsItsAttempts(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 0}, []byte("first"))
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 1}, []byte("second"))

	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if err := l.Fail(errors.New("network is unreachable")); err != nil {
		t.Fatal(err)
	}
	again, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if again.Item.FightIndex != 0 {
		t.Fatalf("head after a failure = fight %d, want 0", again.Item.FightIndex)
	}
	if again.Item.Attempts != 1 || again.Item.LastError != "network is unreachable" {
		t.Fatalf("item = %+v", again.Item)
	}
}

func TestARawItemCarriesTheDecodedHashAcrossAReopen(t *testing.T) {
	dir := t.TempDir()
	q, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	// The body is the compressed frame; the digest is over the
	// plaintext, so it has to travel with the item rather than be
	// recomputed at upload time.
	if _, err := q.Enqueue(Item{Kind: Raw, ReportKey: "local-1", Offset: 4194304,
		SHA256: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		[]byte("\x28\xb5\x2f\xfd")); err != nil {
		t.Fatal(err)
	}
	reopened, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	l, err := reopened.Head()
	if err != nil {
		t.Fatal(err)
	}
	if l.Item.SHA256 != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Fatalf("digest = %q", l.Item.SHA256)
	}
}

func TestAQueueSurvivesAReopen(t *testing.T) {
	dir := t.TempDir()
	q, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 7}, []byte("bundle"))

	reopened, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	l, err := reopened.Head()
	if err != nil {
		t.Fatal(err)
	}
	if l.Item.FightIndex != 7 || string(l.Body) != "bundle" {
		t.Fatalf("item = %+v body %q", l.Item, l.Body)
	}
	seq, err := reopened.Enqueue(Item{Kind: Complete, ReportKey: "local-1"}, []byte("{}"))
	if err != nil {
		t.Fatal(err)
	}
	if seq <= l.Item.Seq {
		t.Fatalf("new sequence %d is not after %d", seq, l.Item.Seq)
	}
}

func TestAnItemWithNoBodyIsSkippedRatherThanBlocking(t *testing.T) {
	dir := t.TempDir()
	q, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	seq, err := q.Enqueue(Item{Kind: Raw, ReportKey: "local-1", Offset: 0}, []byte("chunk"))
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Raw, ReportKey: "local-1", Offset: 4194304}, []byte("chunk2"))
	if err := os.Remove(filepath.Join(dir, name(seq)+".bin")); err != nil {
		t.Fatal(err)
	}
	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if l.Item.Offset != 4194304 {
		t.Fatalf("head = %+v, want the second chunk", l.Item)
	}
}

func TestOnlyOneLeaseIsOutAtATime(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1"}, []byte("a"))
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 1}, []byte("b"))
	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := q.Head(); !errors.Is(err, ErrEmpty) {
		t.Fatal("a second lease was handed out while one was in flight")
	}
	if err := l.Drop(); err != nil {
		t.Fatal(err)
	}
	next, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if next.Item.FightIndex != 1 {
		t.Fatalf("after a drop the head is %+v", next.Item)
	}
}

func TestOpeningAQueueInAnImpossiblePlaceIsAnError(t *testing.T) {
	file := filepath.Join(t.TempDir(), "not-a-directory")
	if err := os.WriteFile(file, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(filepath.Join(file, "queue")); err == nil {
		t.Fatal("Open accepted a path inside a file")
	}
}

func TestUnrelatedFilesInTheQueueDirectoryAreLeftAlone(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "README.txt"), []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "notanumber.json"), []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	q, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	if n, _ := q.Len(); n != 0 {
		t.Fatalf("Len = %d, want 0", n)
	}
	if _, err := q.Head(); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Head = %v", err)
	}
}

func TestTheClockCanBeReplaced(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	when := time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)
	q.SetClock(func() time.Time { return when })
	if _, err := q.Enqueue(Item{Kind: Complete, ReportKey: "k"}, []byte("{}")); err != nil {
		t.Fatal(err)
	}
	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if !l.Item.EnqueuedAt.Equal(when) {
		t.Fatalf("enqueued at %s", l.Item.EnqueuedAt)
	}
}
