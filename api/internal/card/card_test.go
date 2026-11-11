package card

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"strings"
	"sync"
	"testing"
)

func decodeImage(t *testing.T, b []byte) image.Image {
	t.Helper()
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		t.Fatalf("not a PNG: %v", err)
	}
	return img
}

// The crest on the card is the ringed crest, never the bare square art: the crest square's corner
// shows the card's own ground for every class (the same pixel on a warrior's and a mage's card),
// while the ring at the top of the circle takes each class's colour.
func TestRenderDrawsTheRingedCrest(t *testing.T) {
	mage := sample()
	mage.Class, mage.ClassSlug, mage.ClassColor = "Mage", "mage", "#3fc7eb"
	warriorPNG, err := Render(sample())
	if err != nil {
		t.Fatal(err)
	}
	magePNG, err := Render(mage)
	if err != nil {
		t.Fatal(err)
	}
	warrior, mageImg := decodeImage(t, warriorPNG), decodeImage(t, magePNG)
	o := crestOrigin()
	corner := image.Pt(o.X+1, o.Y+1)
	ring := image.Pt(o.X+crestPx/2, o.Y+2)
	if color.NRGBAModel.Convert(warrior.At(corner.X, corner.Y)) != color.NRGBAModel.Convert(mageImg.At(corner.X, corner.Y)) {
		t.Fatalf("the crest square's corner differs by class: the bare art is drawn, not the ringed crest")
	}
	warriorRing := color.NRGBAModel.Convert(warrior.At(ring.X, ring.Y)).(color.NRGBA)
	mageRing := color.NRGBAModel.Convert(mageImg.At(ring.X, ring.Y)).(color.NRGBA)
	// The ring keeps the class hue over the ground: the warrior's tan is redder than blue, the
	// mage's cyan bluer than red.
	if warriorRing.R <= warriorRing.B || mageRing.B <= mageRing.R {
		t.Fatalf("the ring is not in the class colour: warrior %v, mage %v", warriorRing, mageRing)
	}
}

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
		Title:      "Arms raid build",
		Race:       "Human",
		Class:      "Warrior",
		ClassSlug:  "warrior",
		ClassColor: "#c69b6d",
		Split:      []int{31, 0, 20},
		Trees:      []string{"Arms", "Fury", "Protection"},
		Level:      60,
		DPS:        "612 DPS ±4",
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
		{[]int{31, 0, 20}, "31/0/20"},
		{[]int{6, 1}, "6/1"},
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
