// logs/engine/lexer/lexer.go
// Package lexer turns a combat-log byte stream into timestamped parameter
// slices. It is fed arbitrary chunks, carries a partial trailing line across
// chunk boundaries, and silently ignores byte ranges it has already consumed
// so that at-least-once delivery is safe.
package lexer

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// MaxLineBytes caps the partial line held between chunks. A line longer than
// this is a corrupt file, not a combat log: the bytes are dropped and the
// lexer resynchronises at the next newline.
const MaxLineBytes = 1 << 20

// ErrGap is returned when a chunk starts past the end of everything fed so
// far, which would silently lose bytes.
var ErrGap = errors.New("lexer: chunk starts past the end of the stream")

// Line is one decoded log line.
type Line struct {
	Offset int64    // byte offset of the first byte of the line
	Number int64    // 1-based line number within the stream
	Stamp  string   // the timestamp text, empty when the line has none
	Params []string // CSV fields; Params[0] is the event name
	Raw    string   // the whole line, without the line terminator
}

// State is everything needed to resume lexing after a restart.
type State struct {
	Pending []byte `json:"pending"`
	Offset  int64  `json:"offset"`
	Number  int64  `json:"number"`
}

// Lexer splits a byte stream into Lines.
type Lexer struct {
	pending  []byte
	offset   int64 // stream offset of pending[0]
	number   int64 // lines emitted so far
	skipping bool  // dropping bytes until the next newline
	dropped  int
}

// New returns a Lexer positioned at the start of a stream.
func New() *Lexer { return &Lexer{} }

// Restore returns a Lexer that continues from s.
func Restore(s State) *Lexer {
	return &Lexer{
		pending: append([]byte(nil), s.Pending...),
		offset:  s.Offset,
		number:  s.Number,
	}
}

// State captures the lexer for serialisation.
func (l *Lexer) State() State {
	return State{Pending: append([]byte(nil), l.pending...), Offset: l.offset, Number: l.number}
}

// NextOffset is the stream offset of the next byte the lexer expects.
func (l *Lexer) NextOffset() int64 { return l.offset + int64(len(l.pending)) }

// Dropped counts overlong lines discarded so far.
func (l *Lexer) Dropped() int { return l.dropped }

// Feed consumes chunk, whose first byte sits at offset in the stream, and
// calls emit once per complete line. Bytes already consumed are skipped;
// a chunk that starts past NextOffset returns ErrGap.
func (l *Lexer) Feed(chunk []byte, offset int64, emit func(Line) error) error {
	next := l.NextOffset()
	switch {
	case offset > next:
		return fmt.Errorf("%w: chunk at %d, stream ends at %d", ErrGap, offset, next)
	case offset+int64(len(chunk)) <= next:
		return nil // wholly seen before
	case offset < next:
		chunk = chunk[next-offset:]
	}
	l.pending = append(l.pending, chunk...)
	return l.drain(emit, false)
}

// Flush emits a final line that the stream ended without terminating.
func (l *Lexer) Flush(emit func(Line) error) error { return l.drain(emit, true) }

func (l *Lexer) drain(emit func(Line) error, final bool) error {
	for {
		i := bytes.IndexByte(l.pending, '\n')
		if i < 0 {
			break
		}
		raw := l.pending[:i]
		start := l.offset
		l.offset += int64(i) + 1
		l.pending = l.pending[i+1:]
		if l.skipping {
			l.skipping = false
			continue
		}
		if err := l.emit(raw, start, emit); err != nil {
			return err
		}
	}
	if len(l.pending) > MaxLineBytes {
		l.dropped++
		l.skipping = true
		l.offset += int64(len(l.pending))
		l.pending = l.pending[:0]
		return nil
	}
	if final && len(l.pending) > 0 && !l.skipping {
		raw := l.pending
		start := l.offset
		l.offset += int64(len(raw))
		l.pending = nil
		return l.emit(raw, start, emit)
	}
	if len(l.pending) == 0 && cap(l.pending) > 64<<10 {
		l.pending = nil
	}
	return nil
}

func (l *Lexer) emit(raw []byte, start int64, emit func(Line) error) error {
	text := strings.TrimSuffix(string(raw), "\r")
	if text == "" {
		return nil
	}
	l.number++
	stamp, record := splitStamp(text)
	return emit(Line{
		Offset: start,
		Number: l.number,
		Stamp:  stamp,
		Params: SplitParams(record),
		Raw:    text,
	})
}

// splitStamp separates the leading timestamp from the CSV record. The game
// writes two spaces, or a tab on some clients. A line with neither (a header
// written by a tool, for instance) is all record.
func splitStamp(text string) (stamp, record string) {
	if i := strings.Index(text, "  "); i >= 0 {
		return text[:i], strings.TrimLeft(text[i:], " ")
	}
	if i := strings.IndexByte(text, '\t'); i >= 0 {
		return text[:i], strings.TrimLeft(text[i+1:], " \t")
	}
	return "", text
}

// SplitParams splits one CSV record. Commas inside quotes, brackets, or
// parentheses do not separate fields; a quoted field is returned unquoted.
func SplitParams(s string) []string {
	out := make([]string, 0, 16)
	var b strings.Builder
	quoted, escaped, depth := false, false, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case escaped:
			b.WriteByte(c)
			escaped = false
		case quoted && c == '\\':
			escaped = true
		case c == '"':
			quoted = !quoted
		case quoted:
			b.WriteByte(c)
		case c == '[' || c == '(':
			depth++
			b.WriteByte(c)
		case c == ']' || c == ')':
			if depth > 0 {
				depth--
			}
			b.WriteByte(c)
		case c == ',' && depth == 0:
			out = append(out, b.String())
			b.Reset()
		default:
			b.WriteByte(c)
		}
	}
	return append(out, b.String())
}

// Unquote strips one layer of surrounding quotes, for fields that arrive
// already split (the contents of a bracket group, for instance).
func Unquote(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
