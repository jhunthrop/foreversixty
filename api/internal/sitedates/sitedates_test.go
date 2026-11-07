package sitedates

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestFirstRaidsMatchesTheSiteCalendar(t *testing.T) {
	raw, err := os.ReadFile("../../../web/src/data/dates.json")
	if err != nil {
		t.Fatal(err)
	}
	var dates []struct {
		Role string `json:"role"`
		ISO  string `json:"iso"`
	}
	if err := json.Unmarshal(raw, &dates); err != nil {
		t.Fatal(err)
	}
	for _, d := range dates {
		if d.Role == "raids" {
			if d.ISO != FirstRaids {
				t.Errorf("dates.json says the first raids open %s; FirstRaids is %s", d.ISO, FirstRaids)
			}
			return
		}
	}
	t.Error("dates.json has no raids date")
}

func TestRaidsOpen(t *testing.T) {
	if RaidsOpen(time.Date(2026, 12, 8, 23, 59, 0, 0, time.UTC)) {
		t.Error("open the day before")
	}
	if !RaidsOpen(time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC)) {
		t.Error("closed on the day")
	}
}
