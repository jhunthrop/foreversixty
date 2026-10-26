package guilds

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

const firstTierPath = "../../../data/curated/raids/first-tier.json"

func TestZonesAreTheCuratedThree(t *testing.T) {
	raw, err := os.ReadFile(firstTierPath)
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Zones []Zone `json:"zones"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	if len(file.Zones) != 3 {
		t.Fatalf("curated file has %d zones, want 3", len(file.Zones))
	}
	if !reflect.DeepEqual(Zones(), file.Zones) {
		t.Fatalf("Zones() drifted from first-tier.json:\n got %+v\nwant %+v", Zones(), file.Zones)
	}
}

func TestZoneByKey(t *testing.T) {
	cases := []struct {
		key  string
		name string
		ok   bool
	}{
		{"raid:onyxias-lair", "Onyxia's Lair", true},
		{"raid:barrow-deeps", "Barrow Deeps", true},
		{"raid:nope", "", false},
	}
	for _, c := range cases {
		z, ok := ZoneByKey(c.key)
		if ok != c.ok || z.Name != c.name {
			t.Errorf("ZoneByKey(%q) = %+v, %v", c.key, z, ok)
		}
	}
}
