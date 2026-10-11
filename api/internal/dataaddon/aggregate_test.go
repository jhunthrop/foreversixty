// api/internal/dataaddon/aggregate_test.go
package dataaddon

import "testing"

func TestPercentileRoundsTheOverallHalfAwayFromZero(t *testing.T) {
	for overall, want := range map[float64]int{71: 71, 70.5: 71, 70.49: 70, 0: 0} {
		if got := percentile(overall); got != want {
			t.Errorf("percentile(%v) = %d, want %d", overall, got, want)
		}
	}
}

// The no-logs record's spec is the export's tree with strictly the most points, in the build
// data's own slugs; a tie, no points, or an unreadable export names none.
func TestExportSpecNamesTheTreeWithTheMostPoints(t *testing.T) {
	data := useRepoData(t)
	for export, want := range map[string]string{
		"FS1:1.60.1.70291:warrior:human:0/05320/0:|level=22": "warrior-fury",
		"FS1:1.60.1.70291:warrior:human:371000000/0/0:":      "warrior-arms",
		"FS1:1.60.1.70291:rogue:human:0/5/0:":                "rogue-combat",
		"FS1:1.60.1.70291:warrior:human:5/5/0:":              "",
		"FS1:1.60.1.70291:warrior:human:0/0/0:":              "",
		"not an export":                                      "",
		"FS1:9.9.9.99999:warrior:human:0/05320/0:|level=22":  "warrior-fury", // an unknown build reads as the newest
	} {
		if got := exportSpec(data, export); got != want {
			t.Errorf("exportSpec(%q) = %q, want %q", export, got, want)
		}
	}
	if got := exportSpec(nil, "FS1:1.60.1.70291:warrior:human:0/05320/0:"); got != "" {
		t.Errorf("no build data: %q, want none", got)
	}
}
