// logs/bench/bench_test.go
// Package bench measures the engine against the spec's performance budget:
// at least 10 MB per second per CPU and under 1 GB peak, memory bounded by
// the open fight rather than the file. It runs over the real retail
// sample, which is fetched into a git-ignored cache and skipped when the
// machine is offline.
package bench

import (
	"io"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
	"github.com/jhunthrop/foreversixty/logs/internal/sample"
)

const (
	// MinThroughputMBPerSecond is the spec's parse budget.
	MinThroughputMBPerSecond = 10
	// MaxPeakBytes is the spec's memory budget.
	MaxPeakBytes = 1 << 30
	chunkSize    = 1 << 20
)

func samplePath(tb testing.TB) string {
	tb.Helper()
	p, err := sample.Fetch(tb.Context())
	if err != nil {
		tb.Skipf("retail sample unavailable, skipping: %v\n"+
			"Run online once to cache it, or set %s to a local copy.", err, sample.EnvOverride)
	}
	return p
}

func options(base time.Time) session.Options {
	o := session.Options{
		ReportID: "bench",
		Base:     base,
		Infer:    true,
		Units:    units.Options{ClassBySpec: units.RetailSpecClass},
		Fight:    fight.DefaultOptions(),
		Summary:  summary.DefaultOptions(),
	}
	o.Summary.SpecNames = units.RetailSpecName
	return o
}

// parseFile streams the whole sample through one session and returns the
// bytes consumed and the number of fights closed.
func parseFile(tb testing.TB, path string) (int64, int) {
	tb.Helper()
	fi, err := os.Stat(path)
	if err != nil {
		tb.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	s := session.New(options(fi.ModTime().UTC()))
	buf := make([]byte, chunkSize)
	var offset int64
	fights := 0
	for {
		n, rerr := f.Read(buf)
		if n > 0 {
			res, ferr := s.Feed(buf[:n], offset)
			if ferr != nil {
				tb.Fatal(ferr)
			}
			offset += int64(n)
			fights += len(res.Closed)
		}
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			tb.Fatal(rerr)
		}
	}
	res, err := s.Close()
	if err != nil {
		tb.Fatal(err)
	}
	return offset, fights + len(res.Closed)
}

func TestThroughputAndMemoryMeetTheBudget(t *testing.T) {
	path := samplePath(t)

	runtime.GC()
	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	start := time.Now()
	bytes, fights := parseFile(t, path)
	elapsed := time.Since(start)

	var after runtime.MemStats
	runtime.ReadMemStats(&after)
	allocated := after.TotalAlloc - before.TotalAlloc

	mbps := float64(bytes) / elapsed.Seconds() / (1 << 20)
	t.Logf("%d bytes in %s: %.1f MB/s, %d fights, %d bytes allocated, %d bytes from the OS",
		bytes, elapsed, mbps, fights, allocated, after.Sys)

	if mbps < MinThroughputMBPerSecond {
		t.Errorf("throughput %.1f MB/s is under the %d MB/s budget", mbps, MinThroughputMBPerSecond)
	}
	if after.Sys > MaxPeakBytes {
		t.Errorf("peak %d bytes is over the %d byte budget", after.Sys, MaxPeakBytes)
	}
	if fights == 0 {
		t.Error("the sample has encounters; none were found")
	}
}

func BenchmarkParseRetailSample(b *testing.B) {
	path := samplePath(b)
	fi, err := os.Stat(path)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(fi.Size())
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		parseFile(b, path)
	}
}
