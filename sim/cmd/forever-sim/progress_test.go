package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"sync"
	"testing"
)

// The tick encoder and the engine's logger share one stream, from
// different goroutines. Every line that comes out has to be a whole
// line and has to parse, or sim/runner.Native is reading half a tick.
func TestProgressStderrKeepsEveryLineWholeAndJSON(t *testing.T) {
	var buf bytes.Buffer
	stream := newProgressStderr(&buf)
	logger := log.New(stream.LogOutput(), "", log.LstdFlags)

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			enc := json.NewEncoder(stream)
			for n := 0; n < 20; n++ {
				_ = enc.Encode(struct {
					Completed int `json:"completed"`
					Total     int `json:"total"`
				}{n, 500})
				logger.Printf("thread %d says something with a \"quote\" in it", i)
			}
		}(i)
	}
	wg.Wait()

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 8*20*2 {
		t.Fatalf("got %d lines, want %d; a write was split or lost", len(lines), 8*20*2)
	}
	var ticks, logs int
	for i, line := range lines {
		var got struct {
			Total *int    `json:"total"`
			Log   *string `json:"log"`
		}
		if err := json.Unmarshal([]byte(line), &got); err != nil {
			t.Fatalf("line %d is not JSON: %q (%v)", i, line, err)
		}
		switch {
		case got.Total != nil:
			ticks++
		case got.Log != nil:
			logs++
			if !strings.Contains(*got.Log, `"quote"`) {
				t.Errorf("line %d lost the log's own quoting: %q", i, *got.Log)
			}
		default:
			t.Errorf("line %d is neither a tick nor a log line: %q", i, line)
		}
	}
	if ticks != 160 || logs != 160 {
		t.Errorf("ticks=%d logs=%d, want 160 of each", ticks, logs)
	}
}

// A multi-line log write - the engine's panic trace is one - becomes one
// wrapped line per line, and a trailing newline adds no empty one.
func TestLogLinesWrapsEachLineOnce(t *testing.T) {
	var buf bytes.Buffer
	stream := newProgressStderr(&buf)
	if _, err := stream.LogOutput().Write([]byte("first\nsecond\n\nthird\n")); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("got %d lines, want 3: %q", len(lines), buf.String())
	}
	for i, want := range []string{"first", "second", "third"} {
		var got struct {
			Log string `json:"log"`
		}
		if err := json.Unmarshal([]byte(lines[i]), &got); err != nil {
			t.Fatal(err)
		}
		if got.Log != want {
			t.Errorf("line %d = %q, want %q", i, got.Log, want)
		}
	}
}
