// companion/internal/logging/logging.go
// Package logging gives the companion one slog logger that writes JSON
// to a size-capped file and to stderr. A desktop app runs for weeks, so
// the file caps itself rather than growing without bound; there is no
// log-rotation dependency for two files.
package logging

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

// MaxBytes is how large companion.log grows before it is rolled onto
// companion.log.1 and started again.
const MaxBytes = 4 << 20

// File is an io.Writer that rolls one file over at MaxBytes.
type File struct {
	path string
	max  int64

	mu sync.Mutex
	f  *os.File
	n  int64
}

// Open opens the log file in dir, creating it if it is not there.
func Open(dir string, max int64) (*File, error) {
	if max <= 0 {
		max = MaxBytes
	}
	w := &File{path: filepath.Join(dir, "companion.log"), max: max}
	if err := w.reopen(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *File) reopen() error {
	f, err := os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return err
	}
	w.f, w.n = f, fi.Size()
	return nil
}

// Write appends one record, rolling the file over first when it is full.
func (w *File) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.n+int64(len(p)) > w.max {
		if err := w.roll(); err != nil {
			return 0, err
		}
	}
	n, err := w.f.Write(p)
	w.n += int64(n)
	return n, err
}

func (w *File) roll() error {
	if err := w.f.Close(); err != nil {
		return err
	}
	if err := os.Rename(w.path, w.path+".1"); err != nil {
		return err
	}
	return w.reopen()
}

// Close closes the file.
func (w *File) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.f.Close()
}

// New returns the application logger and the file behind it. Every
// record carries the component so one grep separates the tail from the
// uploader.
func New(dir string, debug bool) (*slog.Logger, *File, error) {
	f, err := Open(dir, MaxBytes)
	if err != nil {
		return nil, nil, err
	}
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	h := slog.NewJSONHandler(io.MultiWriter(f, os.Stderr), &slog.HandlerOptions{Level: level})
	return slog.New(h), f, nil
}
