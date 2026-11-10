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
	mc := []Boss{
		{"Lucifron", true}, {"Magmadar", true}, {"Gehennas", true}, {"Garr", true}, {"Baron Geddon", true},
		{"Shazzrah", true}, {"Sulfuron Harbinger", true}, {"Golemagg the Incinerator", true},
		{"Majordomo Executus", false}, {"Majordomo Executus", true}, {"Ragnaros", false},
	}
	cards := map[string]func() ([]byte, error){
		"card-build": func() ([]byte, error) { return Render(sample()) },
		"card-report": func() ([]byte, error) {
			return RenderReport(ReportInput{Title: "Thursday Onyxia", Zone: "Onyxia's Lair",
				Bosses: []Boss{{"Onyxia", false}, {"Onyxia", true}}, DurationMS: int64(74 * time.Minute / time.Millisecond)})
		},
		"card-report-mc": func() ([]byte, error) {
			return RenderReport(ReportInput{Title: "Olympus XXVII Molten Core clear", Zone: "Molten Core", Bosses: mc,
				DurationMS: int64((2*time.Hour + 5*time.Minute) / time.Millisecond)})
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
