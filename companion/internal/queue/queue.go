// Package queue is the companion's durable outbox. Everything the
// uploader sends goes through it in sequence order and stays on disk
// until the server has taken it, which is what makes "fights arrive in
// order across a network drop and a restart" a property of the design
// rather than a hope.
package queue

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Kind is what one queued item asks the uploader to do.
type Kind string

// The three durable operations. Live snapshots are deliberately absent:
// a snapshot that fails is replaced by the next one five seconds later,
// and queueing them would delay the fights behind them.
const (
	Fight    Kind = "fight"
	Raw      Kind = "raw"
	Complete Kind = "complete"
)

// Item is one queued operation. The body lives beside it in a .bin
// file, so a 6 MiB bundle never passes through JSON.
type Item struct {
	Seq        int64  `json:"seq"`
	Kind       Kind   `json:"kind"`
	ReportKey  string `json:"report_key"`
	FightIndex int    `json:"fight_index,omitempty"`
	Offset     int64  `json:"offset,omitempty"`
	// SHA256 is the hex digest of a raw chunk's DECODED bytes. The
	// body beside this item is the compressed frame, so the digest
	// cannot be recomputed from it at upload time and is recorded
	// here when the chunk is queued.
	SHA256      string    `json:"sha256,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	Attempts    int       `json:"attempts"`
	LastError   string    `json:"last_error,omitempty"`
	EnqueuedAt  time.Time `json:"enqueued_at"`
}

// Queue is a directory of items, drained strictly in sequence order.
type Queue struct {
	dir string

	mu   sync.Mutex
	next int64
	// leased is the sequence number currently handed out, so a second
	// Head while one is in flight returns nothing rather than a
	// duplicate.
	leased int64
	now    func() time.Time
}

// Open opens or creates the queue directory and recovers the next
// sequence number from whatever is still pending.
func Open(dir string) (*Queue, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("create the queue directory: %w", err)
	}
	q := &Queue{dir: dir, next: 1, now: time.Now}
	seqs, err := q.pending()
	if err != nil {
		return nil, err
	}
	if len(seqs) > 0 {
		q.next = seqs[len(seqs)-1] + 1
	}
	return q, nil
}

// SetClock replaces the clock, for tests.
func (q *Queue) SetClock(now func() time.Time) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.now = now
}

func (q *Queue) meta(seq int64) string { return filepath.Join(q.dir, name(seq)+".json") }
func (q *Queue) body(seq int64) string { return filepath.Join(q.dir, name(seq)+".bin") }

func name(seq int64) string { return fmt.Sprintf("%020d", seq) }

// pending lists the sequence numbers on disk, ascending.
func (q *Queue) pending() ([]int64, error) {
	entries, err := os.ReadDir(q.dir)
	if err != nil {
		return nil, fmt.Errorf("read the queue directory: %w", err)
	}
	var seqs []int64
	for _, e := range entries {
		n := e.Name()
		if !strings.HasSuffix(n, ".json") {
			continue
		}
		seq, err := strconv.ParseInt(strings.TrimSuffix(n, ".json"), 10, 64)
		if err != nil {
			continue // not ours; leave it alone
		}
		seqs = append(seqs, seq)
	}
	sort.Slice(seqs, func(i, j int) bool { return seqs[i] < seqs[j] })
	return seqs, nil
}

// Enqueue appends one item and returns its sequence number. The body
// is written first and the metadata second, so a crash between the two
// leaves an orphan .bin that Head ignores rather than an item with no
// body.
func (q *Queue) Enqueue(it Item, body []byte) (int64, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	it.Seq = q.next
	it.EnqueuedAt = q.now().UTC()
	if err := writeFile(q.body(it.Seq), body); err != nil {
		return 0, err
	}
	if err := q.writeMeta(it); err != nil {
		os.Remove(q.body(it.Seq))
		return 0, err
	}
	q.next++
	return it.Seq, nil
}

func (q *Queue) writeMeta(it Item) error {
	b, err := json.Marshal(it)
	if err != nil {
		return err
	}
	return writeFile(q.meta(it.Seq), append(b, '\n'))
}

// writeFile writes atomically: a torn queue file would cost a fight.
func writeFile(path string, b []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), ".q-*")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(name, 0o600); err != nil {
		return err
	}
	return os.Rename(name, path)
}

// Len is how many items are waiting.
func (q *Queue) Len() (int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	seqs, err := q.pending()
	if err != nil {
		return 0, err
	}
	return len(seqs), nil
}

// ErrEmpty is returned by Head when nothing is queued.
var ErrEmpty = errors.New("queue: empty")

// Lease is the item at the head of the queue, checked out.
type Lease struct {
	Item Item
	Body []byte

	q *Queue
}

// Head checks out the oldest item. It returns ErrEmpty when there is
// nothing to do. Only one lease exists at a time: the whole point is
// that item N+1 does not go out before item N.
func (q *Queue) Head() (*Lease, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.leased != 0 {
		return nil, ErrEmpty
	}
	seqs, err := q.pending()
	if err != nil {
		return nil, err
	}
	for _, seq := range seqs {
		b, err := os.ReadFile(q.meta(seq))
		if err != nil {
			return nil, fmt.Errorf("read queue item %d: %w", seq, err)
		}
		var it Item
		if err := json.Unmarshal(b, &it); err != nil {
			return nil, fmt.Errorf("parse queue item %d: %w", seq, err)
		}
		body, err := os.ReadFile(q.body(seq))
		if errors.Is(err, fs.ErrNotExist) {
			// The body never landed. Drop the metadata rather than
			// block the queue on an item that can never be sent.
			os.Remove(q.meta(seq))
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("read queue body %d: %w", seq, err)
		}
		q.leased = seq
		return &Lease{Item: it, Body: body, q: q}, nil
	}
	return nil, ErrEmpty
}

// Ack removes the item: the server has it.
func (l *Lease) Ack() error {
	q := l.q
	q.mu.Lock()
	defer q.mu.Unlock()
	q.leased = 0
	if err := os.Remove(q.meta(l.Item.Seq)); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if err := os.Remove(q.body(l.Item.Seq)); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	return nil
}

// Fail returns the item to the head of the queue with one more attempt
// recorded. Nothing behind it moves.
func (l *Lease) Fail(cause error) error {
	q := l.q
	q.mu.Lock()
	defer q.mu.Unlock()
	q.leased = 0
	l.Item.Attempts++
	if cause != nil {
		l.Item.LastError = cause.Error()
	}
	return q.writeMeta(l.Item)
}

// Drop removes an item the server will never accept, so the queue
// behind it can move. The reason is the caller's to log.
func (l *Lease) Drop() error { return l.Ack() }
