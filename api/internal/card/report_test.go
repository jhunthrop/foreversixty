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
		Bosses:     []Boss{{Name: "Lucifron", Kill: true}, {Name: "Magmadar"}, {Name: "Magmadar"}},
		DurationMS: 95 * 60 * 1000,
	})
	if got != "Combat log · 1 of 2 killed · 1h 35m" {
		t.Fatalf("summary = %q", got)
	}
	if got := summaryLine(ReportInput{Bosses: []Boss{{Kill: true}}}); !strings.Contains(got, "1 of 1") {
		t.Fatalf("summary = %q", got)
	}
	if got := duration(45 * 60 * 1000); got != "45 min" {
		t.Fatalf("duration = %q", got)
	}
}

func TestBossRowsSplitKillsFromWipes(t *testing.T) {
	f, err := newFaces()
	if err != nil {
		t.Fatal(err)
	}
	rows := bossRows(f, []Boss{{Name: "Kelthas", Kill: true}, {Name: "Emberdeep"}, {Name: "Emberdeep"},
		{Name: "Majordomo Executus"}, {Name: "Majordomo Executus", Kill: true}, {Name: "Garr", Kill: true}}, Width-2*marginX)
	if len(rows) != 2 || rows[0].label != copyKilled || rows[0].names != "Kelthas · Majordomo Executus · Garr" ||
		rows[1].label != copyWiped || rows[1].names != "Emberdeep" {
		t.Fatalf("rows = %+v", rows)
	}
	if got := bossRows(f, nil, 100); len(got) != 1 || got[0].names != copyNoBosses {
		t.Fatalf("empty = %+v", got)
	}
}

func TestNameListCountsWhatDoesNotFit(t *testing.T) {
	f, err := newFaces()
	if err != nil {
		t.Fatal(err)
	}
	names := []string{"Lucifron", "Magmadar", "Gehennas", "Garr", "Baron Geddon", "Shazzrah", "Sulfuron Harbinger", "Golemagg", "Majordomo Executus", "Ragnaros"}
	got := nameList(f.body, names, 600)
	if !strings.HasPrefix(got, "Lucifron · ") || !strings.Contains(got, " more") {
		t.Fatalf("list = %q", got)
	}
	if measure(f.body, got) > 600 {
		t.Fatalf("list %q runs past its room", got)
	}
	if got := nameList(f.body, []string{"Onyxia"}, 600); got != "Onyxia" {
		t.Fatalf("one name = %q", got)
	}
}
