// api/internal/dataaddon/write_split_test.go
package dataaddon

import (
	"bytes"
	"strings"
	"testing"
)

func TestRenderWithLimitSplitsByRegionAndKeepsWholeRecords(t *testing.T) {
	d := sampleData()
	d.Characters["Mörk-Pvp"] = characterRecord{Region: "eu", Spec: "mage-fire", Profile: "/character/eu/pvp/mörk"}
	single, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}
	// A limit of 1 byte guarantees the single-file render always exceeds it, forcing the
	// split path deterministically.
	files, err := renderWithLimit(d, 1)
	if err != nil {
		t.Fatal(err)
	}
	header, ok := files["Data.lua"]
	if !ok {
		t.Fatal("Data.lua (the header/init file) is missing from a split render")
	}
	if !strings.Contains(string(header), "\tformat = 2,\n\tgenerated = \"2026-11-04\",\n\tcharacters = {},\n\tguilds = {},\n") {
		t.Errorf("the header file should carry format, generated and two empty tables, got:\n%s", header)
	}
	us, eu := files["Data-us.lua"], files["Data-eu.lua"]
	if us == nil || eu == nil {
		t.Fatalf("files = %v, want Data-us.lua and Data-eu.lua", mapKeysOf(files))
	}
	for _, opening := range []string{`["Obnoxious Yell-Normal"] = {`, `["Quiet Fern-Normal"] = {`, `["OLYMPUS XXVII"] = {`} {
		if record := recordLines(t, single["Data.lua"], opening); !bytes.Contains(us, record) {
			t.Errorf("Data-us.lua lacks the whole record %s:\n%s", opening, us)
		}
	}
	if record := recordLines(t, single["Data.lua"], `["Mörk-Pvp"] = {`); !bytes.Contains(eu, record) || bytes.Contains(us, []byte("Mörk")) {
		t.Errorf("Mörk's record belongs whole in Data-eu.lua only:\n%s", eu)
	}
	if !bytes.Contains(us, []byte("ForeverSixtyData.characters[key] = row")) {
		t.Errorf("Data-us.lua should merge into the shared global, got:\n%s", us)
	}
}

func TestRegionsOfListsEveryRegionARecordBelongsTo(t *testing.T) {
	d := Data{
		Characters: map[string]characterRecord{"A-Normal": {Region: "us"}, "B-Pvp": {Region: "eu"}},
		Guilds:     map[string]guildRecord{"C": {Region: "kr"}},
	}
	got := regionsOf(d)
	want := []string{"eu", "kr", "us"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("regions = %v, want %v", got, want)
	}
}
