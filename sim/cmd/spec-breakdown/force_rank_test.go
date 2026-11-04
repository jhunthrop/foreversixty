package main

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

// Round 18 priest X4: -rank forces one rank of a spell on every reference
// the rotation makes to it (casts and conditions alike), through the spec's
// own rotation when -rotation is not given.
func TestRankFlagForcesTheRankOnEveryReference(t *testing.T) {
	forced, err := parseForcedRanks("10894=4")
	if err != nil || len(forced) != 1 || forced[0].authored != 10894 || forced[0].rank != 4 {
		t.Fatalf("parseForcedRanks = %+v, %v", forced, err)
	}
	rot, err := rotationFor(options{repoRoot: repoRoot, spec: "priest-shadow", preset: "raid"}, forced, "priest", "human", 60)
	if err != nil {
		t.Fatal(err)
	}
	if rot == nil {
		t.Fatal("-rank with no -rotation must load the spec's own rotation")
	}
	text := protojson.Format(rot)
	if strings.Contains(text, "10894") {
		t.Errorf("Shadow Word: Pain rank 8 is still referenced after forcing rank 4")
	}
	for _, bad := range []string{"10894=", "x=4", "10894=0", "10894"} {
		if _, err := parseForcedRanks(bad); err == nil {
			t.Errorf("parseForcedRanks(%q) accepted", bad)
		}
	}
}
