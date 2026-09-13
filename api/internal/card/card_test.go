package card

import (
	"bytes"
	"image/png"
	"strings"
	"sync"
	"testing"
)

func decode(t *testing.T, b []byte) (int, int) {
	t.Helper()
	cfg, err := png.DecodeConfig(bytes.NewReader(b))
	if err != nil {
		t.Fatalf("not a PNG: %v", err)
	}
	return cfg.Width, cfg.Height
}

func sample() Input {
	return Input{
		Title:      "Arms leveling",
		Race:       "Human",
		Class:      "Warrior",
		ClassColor: "#c69b6d",
		Split:      []int{31, 0, 20},
		Level:      60,
	}
}

func TestRenderProducesA1200x630PNG(t *testing.T) {
	got, err := Render(sample())
	if err != nil {
		t.Fatal(err)
	}
	w, h := decode(t, got)
	if w != Width || h != Height {
		t.Fatalf("card is %dx%d, want %dx%d", w, h, Width, Height)
	}
}

func TestRenderIsDeterministic(t *testing.T) {
	first, err := Render(sample())
	if err != nil {
		t.Fatal(err)
	}
	second, err := Render(sample())
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, second) {
		t.Fatal("the same build rendered two different cards")
	}
}

func TestRenderHandlesOverlongAndEmptyContent(t *testing.T) {
	long := sample()
	long.Title = strings.Repeat("Holy leveling ten to thirty ", 12)
	got, err := Render(long)
	if err != nil {
		t.Fatal(err)
	}
	if w, h := decode(t, got); w != Width || h != Height {
		t.Fatalf("long title card is %dx%d", w, h)
	}

	empty, err := Render(Input{})
	if err != nil {
		t.Fatal(err)
	}
	if w, h := decode(t, empty); w != Width || h != Height {
		t.Fatalf("empty card is %dx%d", w, h)
	}
}

func TestFallbackIsA1200x630PNG(t *testing.T) {
	got := Fallback()
	if len(got) == 0 {
		t.Fatal("the fallback card is empty")
	}
	if w, h := decode(t, got); w != Width || h != Height {
		t.Fatalf("fallback card is %dx%d", w, h)
	}
	if !bytes.Equal(got, Fallback()) {
		t.Fatal("the fallback card must be the same bytes every time")
	}
}

func TestSplitTextJoinsTheTreeCounts(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want string
	}{
		{[]int{31, 0, 20}, "31 / 0 / 20"},
		{[]int{6, 1}, "6 / 1"},
		{nil, "0"},
	} {
		if got := splitText(tc.in); got != tc.want {
			t.Errorf("splitText(%v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// TestRenderIsSafeForConcurrentUse drives Render from several goroutines at
// once, the way Discord, Slack, Twitter/X and Facebook all fetch a card the
// moment one link is posted and Cloud Run serves them on one instance. Each
// card must be byte-identical to a card rendered on its own: a garbled one
// would be served 200 and cached for a week under a content-hash URL that
// cannot be busted.
func TestRenderIsSafeForConcurrentUse(t *testing.T) {
	want, err := Render(sample())
	if err != nil {
		t.Fatal(err)
	}

	const goroutines, perGoroutine = 8, 20
	var wg sync.WaitGroup
	start := make(chan struct{})
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for range perGoroutine {
				got, err := Render(sample())
				if err != nil {
					t.Errorf("concurrent render: %v", err)
					return
				}
				if !bytes.Equal(got, want) {
					t.Error("a concurrent render produced different bytes than a lone one")
					return
				}
			}
		}()
	}
	close(start)
	wg.Wait()
}
