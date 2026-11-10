//go:build payloaddump

package card

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestDumpReviewCards writes the link previews for the round-5 review boards into
// $REVIEW_CARDS_DIR: a build, a raid report, a long report and the brand fallback.
func TestDumpReviewCards(t *testing.T) {
	dir := os.Getenv("REVIEW_CARDS_DIR")
	if dir == "" {
		t.Skip("REVIEW_CARDS_DIR is not set")
	}
	kill := func(name string, sec int64) Boss { return Boss{Name: name, Kill: true, DurationMS: sec * 1000} }
	wipe := func(name string) Boss { return Boss{Name: name, DurationMS: 90_000} }
	mc := []Boss{
		kill("Lucifron", 74), kill("Magmadar", 102), kill("Gehennas", 61), kill("Garr", 118), kill("Baron Geddon", 131),
		kill("Shazzrah", 58), kill("Sulfuron Harbinger", 97), kill("Golemagg the Incinerator", 112),
		wipe("Majordomo Executus"), kill("Majordomo Executus", 181), wipe("Ragnaros"), wipe("Ragnaros"),
	}
	cards := map[string]func() ([]byte, error){
		"card-build": func() ([]byte, error) { return Render(sample()) },
		"card-report": func() ([]byte, error) {
			return RenderReport(ReportInput{Title: "Thursday Onyxia", Zone: "Onyxia's Lair",
				// The night /report and the report post show: 3 wipes, a 1:42 kill, 2 h 41 min, 14 deaths.
				Bosses:     []Boss{wipe("Onyxia"), wipe("Onyxia"), wipe("Onyxia"), kill("Onyxia", 102)},
				DurationMS: int64(161 * time.Minute / time.Millisecond), Deaths: 14})
		},
		"card-report-mc": func() ([]byte, error) {
			return RenderReport(ReportInput{Title: "Olympus XXVII Thursday Molten Core", Zone: "Molten Core", Bosses: mc,
				DurationMS: int64((2*time.Hour + 5*time.Minute) / time.Millisecond), Deaths: 38})
		},
		"card-fallback": func() ([]byte, error) { return Fallback(), nil },
	}
	for name, draw := range cards {
		png, err := draw()
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, name+".png"), png, 0o600); err != nil {
			t.Fatal(err)
		}
	}
}
