// logs/engine/session/counts_test.go
package session

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// hostileDeathsInText is the oracle for a fight's mob deaths, read straight
// from the log text rather than through the engine: the distinct hostile,
// non-player units a UNIT_DIED or PARTY_KILL line names between two
// offsets. Both line kinds put the dead unit in the dest fields (GUID at 5,
// flags at 7).
func hostileDeathsInText(t *testing.T, src []byte, start, end int64) map[string]bool {
	t.Helper()
	out := map[string]bool{}
	for _, line := range strings.Split(string(src[start:min(end+1, int64(len(src)))]), "\n") {
		_, body, ok := strings.Cut(line, "  ")
		if !ok {
			continue
		}
		f := strings.Split(body, ",")
		if len(f) < 8 || (f[0] != "UNIT_DIED" && f[0] != "PARTY_KILL") {
			continue
		}
		flags, err := strconv.ParseUint(strings.TrimPrefix(f[7], "0x"), 16, 32)
		if err != nil {
			t.Fatalf("flags %q: %v", f[7], err)
		}
		if units.Hostile(uint32(flags)) && units.Parse(f[5]).Kind != units.KindPlayer {
			out[f[5]] = true
		}
	}
	return out
}

// The v16 fixture's Hollow Sentinel dies on a UNIT_DIED and a PARTY_KILL.
// Engine 0.5 counted it twice; it is one death, and it is the group's.
func TestTheFixtureCountsEachMobDeathOnce(t *testing.T) {
	src := fixture(t)
	got, _ := feed(t, opts(), src, 512)
	sawSentinel := false
	for _, c := range got.Closed {
		f := c.Fight
		want := hostileDeathsInText(t, src, f.StartOffset, f.EndOffset)
		if f.MobDeaths != len(want) {
			t.Errorf("fight %d: mob deaths = %d, the text names %d distinct hostile units dying",
				f.Index, f.MobDeaths, len(want))
		}
		if f.GroupKills > f.MobDeaths {
			t.Errorf("fight %d: group kills %d exceed mob deaths %d", f.Index, f.GroupKills, f.MobDeaths)
		}
		if want["Creature-0-2085-2284-7855-169753-0000AA0001"] {
			sawSentinel = true
			if f.GroupKills < 1 {
				t.Errorf("fight %d: the Sentinel's PARTY_KILL did not count as the group's", f.Index)
			}
		}
	}
	if !sawSentinel {
		t.Fatal("no fight held the Sentinel's death; the fixture changed under this test")
	}
}

func TestTheFixturesGroupIsInsideItsPlayers(t *testing.T) {
	got, _ := feed(t, opts(), fixture(t), 512)
	for _, c := range got.Closed {
		f := c.Fight
		if len(f.Group) == 0 {
			t.Errorf("fight %d: no group, but the fixture's players carry party flags", f.Index)
		}
		for _, g := range f.Group {
			if !slices.Contains(f.Players, g) {
				t.Errorf("fight %d: group member %s is not among its players", f.Index, g)
			}
		}
		if f.GroupDeaths > f.Deaths {
			t.Errorf("fight %d: group deaths %d exceed deaths %d", f.Index, f.GroupDeaths, f.Deaths)
		}
	}
}
