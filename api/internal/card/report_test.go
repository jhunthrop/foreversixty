package card

import (
	"bytes"
	"image/png"
	"strings"
	"testing"
)

func TestRenderReportDrawsACardOfTheRightSize(t *testing.T) {
	out, err := RenderReport(ReportInput{
		Title: "Tuesday clear", Zone: "Blackrock Spire",
		Bosses: []Boss{
			{Name: "Warden Kelthas", Kill: true},
			{Name: "Emberdeep", Kill: false},
		},
		DurationMS: 2 * 60 * 60 * 1000,
	})
	if err != nil {
		t.Fatal(err)
	}
	img, err := png.Decode(bytes.NewReader(out))
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != Width || img.Bounds().Dy() != Height {
		t.Fatalf("card is %v, want 1200x630", img.Bounds())
	}
}

func TestRenderReportCopesWithNothingToShow(t *testing.T) {
	out, err := RenderReport(ReportInput{})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) == 0 {
		t.Fatal("an empty report should still draw a card")
	}
}

func TestRenderReportWrapsALongTitleAndCountsExtraBosses(t *testing.T) {
	bosses := make([]Boss, 0, 12)
	for range 12 {
		bosses = append(bosses, Boss{Name: "A Boss With A Long Name", Kill: true})
	}
	out, err := RenderReport(ReportInput{
		Title:  strings.Repeat("A very long report title that will not fit on one line ", 3),
		Zone:   "Blackrock Spire",
		Bosses: bosses,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) == 0 {
		t.Fatal("a long card should still draw")
	}
}

func TestSummaryLineCountsKillsAndTime(t *testing.T) {
	got := summaryLine(ReportInput{
		Bosses:     []Boss{{Kill: true}, {Kill: false}},
		DurationMS: 95 * 60 * 1000,
	})
	if !strings.Contains(got, "1 of 2 killed") || !strings.Contains(got, "1h 35m") {
		t.Fatalf("summary = %q", got)
	}
	if got := summaryLine(ReportInput{Bosses: []Boss{{Kill: true}}}); !strings.Contains(got, "1 of 1") {
		t.Fatalf("summary = %q", got)
	}
	if got := duration(45 * 60 * 1000); got != "45 min" {
		t.Fatalf("duration = %q", got)
	}
}

func TestBossLinesMarkKillsAndWipes(t *testing.T) {
	f, err := newFaces()
	if err != nil {
		t.Fatal(err)
	}
	lines := bossLines(f, []Boss{{Name: "Kelthas", Kill: true}, {Name: "Emberdeep"}}, Width-2*marginX)
	if len(lines) != 1 {
		t.Fatalf("lines = %v", lines)
	}
	if !strings.Contains(lines[0], "Kelthas +") || !strings.Contains(lines[0], "Emberdeep x") {
		t.Fatalf("line = %q", lines[0])
	}
	if got := bossLines(f, nil, 100); len(got) != 1 || !strings.Contains(got[0], "No boss") {
		t.Fatalf("empty = %v", got)
	}
}

func TestWrapSplitsOverTwoLines(t *testing.T) {
	f, err := newFaces()
	if err != nil {
		t.Fatal(err)
	}
	first, second := wrap(f.title, "Short", Width)
	if first != "Short" || second != "" {
		t.Fatalf("wrap = %q, %q", first, second)
	}
	first, second = wrap(f.title, "One two three four five six seven eight nine ten", 400)
	if first == "" || second == "" {
		t.Fatalf("wrap = %q, %q", first, second)
	}
	if first, _ := wrap(f.title, strings.Repeat("x", 200), 100); first == "" {
		t.Fatal("an unbreakable string should still be ellipsised")
	}
}
