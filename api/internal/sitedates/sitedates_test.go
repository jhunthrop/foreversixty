package sitedates

import (
	"encoding/json"
	"os"
	"testing"
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
