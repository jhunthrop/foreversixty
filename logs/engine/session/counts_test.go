// logs/engine/session/counts_test.go
package session

import (
	"encoding/csv"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// mobDeathsInText is the oracle for a fight's mob deaths, read straight from
// the log text rather than through the engine: the distinct non-player units
// a UNIT_DIED or PARTY_KILL line names between two offsets that were hostile
// on some line of the range, or neutral and damaged or killed by a unit
// flagged MINE, PARTY or RAID. The fixtures hold no critter, so the critter
// cutoff is not modelled here (fight/counts_test.go covers it).
func mobDeathsInText(t *testing.T, src []byte, start, end int64) map[string]bool {
	t.Helper()
	hostile, groupTouched, dead := map[string]bool{}, map[string]bool{}, map[string]bool{}
	for _, line := range strings.Split(string(src[start:min(end+1, int64(len(src)))]), "\n") {
		_, body, ok := strings.Cut(line, "  ")
		if !ok {
			continue
		}
		r := csv.NewReader(strings.NewReader(body))
		r.LazyQuotes = true
		f, err := r.Read()
		if err != nil || len(f) < 8 {
			continue
		}
		srcFlags, dstFlags := flagsOf(f[3]), flagsOf(f[7])
		for _, u := range [][2]string{{f[1], f[3]}, {f[5], f[7]}} {
			if units.Hostile(flagsOf(u[1])) && units.Parse(u[0]).Kind != units.KindPlayer {
				hostile[u[0]] = true
			}
		}
		damage := strings.HasSuffix(f[0], "_DAMAGE")
		if (damage || f[0] == "PARTY_KILL") && units.InGroup(srcFlags) {
			groupTouched[f[5]] = true
		}
		if (f[0] == "UNIT_DIED" || f[0] == "PARTY_KILL") && units.Parse(f[5]).Kind != units.KindPlayer {
			dead[f[5]] = units.Hostile(dstFlags) || units.Neutral(dstFlags)
		}
	}
	out := map[string]bool{}
	for guid, reacts := range dead {
		if reacts && (hostile[guid] || groupTouched[guid]) {
			out[guid] = true
		}
	}
	return out
}

func flagsOf(s string) uint32 {
	v, err := strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 32)
	if err != nil {
		return 0
	}
	return uint32(v)
}

// Every fixture dialect, so the oracle sees v16, v22 and a real Forever log.
var countFixtures = []string{
	"../event/testdata/v16.log", "../event/testdata/v22.log", "../event/testdata/forever-1.60.log",
	levellingFixture,
}

func TestEveryFixturesMobDeathsMatchTheText(t *testing.T) {
	for _, path := range countFixtures {
		t.Run(path, func(t *testing.T) {
			src, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			got, _ := feed(t, opts(), src, 512)
			if len(got.Closed) == 0 {
				t.Fatal("the fixture closed no fight")
			}
			for _, c := range got.Closed {
				f := c.Fight
				want := mobDeathsInText(t, src, f.StartOffset, f.EndOffset)
				if f.MobDeaths != len(want) {
					t.Errorf("fight %d: mob deaths = %d, the text names %d", f.Index, f.MobDeaths, len(want))
				}
				if f.GroupKills > f.MobDeaths {
					t.Errorf("fight %d: group kills %d exceed mob deaths %d", f.Index, f.GroupKills, f.MobDeaths)
				}
			}
		})
	}
}

// The v16 fixture's Hollow Sentinel dies on a UNIT_DIED and a PARTY_KILL.
// Engine 0.5 counted it twice; it is one death, and it is the group's.
func TestTheFixtureCountsEachMobDeathOnce(t *testing.T) {
	src := fixture(t)
	got, _ := feed(t, opts(), src, 512)
	sawSentinel := false
	for _, c := range got.Closed {
		f := c.Fight
		want := mobDeathsInText(t, src, f.StartOffset, f.EndOffset)
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

// levellingFixture is a hand-made levelling night: a pull of two neutral Plainstriders, two
// ambient NPCs (one neutral, one hostile) fighting each other far from the group, and a boss wipe
// where the hunter feigns and two players die, are never resurrected, and run back from the
// graveyard to cast again before the encounter ends.
const levellingFixture = "../event/testdata/v16-levelling.log"

// playerDeathsInText is the oracle for a fight's player deaths: the UNIT_DIED lines naming a
// player between two offsets, less any within a second of that player's own Feign Death cast.
func playerDeathsInText(t *testing.T, src []byte, start, end int64) int {
	t.Helper()
	feigned := map[string]time.Time{}
	deaths := 0
	for _, line := range strings.Split(string(src[start:min(end+1, int64(len(src)))]), "\n") {
		stamp, body, ok := strings.Cut(line, "  ")
		if !ok {
			continue
		}
		at, err := time.Parse("1/2 15:04:05.000", stamp)
		if err != nil {
			continue
		}
		r := csv.NewReader(strings.NewReader(body))
		r.LazyQuotes = true
		f, err := r.Read()
		if err != nil || len(f) < 8 {
			continue
		}
		if f[0] == "SPELL_CAST_SUCCESS" && len(f) > 9 && f[9] == "5384" {
			feigned[f[1]] = at
		}
		if f[0] == "UNIT_DIED" && units.Parse(f[5]).Kind == units.KindPlayer {
			if cast, ok := feigned[f[5]]; !ok || at.Sub(cast) > time.Second {
				deaths++
			}
		}
	}
	return deaths
}

// The levelling fixture pins structure, never counts: the neutral pull is a fight of its own,
// the ambient NPCs open none, and the wipe is one encounter whose deaths are every death but the
// feign, corpse runs included.
func TestALevellingNightSplitsAndCountsByTheRules(t *testing.T) {
	src, err := os.ReadFile(levellingFixture)
	if err != nil {
		t.Fatal(err)
	}
	got, health := feed(t, opts(), src, 512)
	if health.ParseErrors != 0 {
		t.Fatalf("the fixture has %d parse errors", health.ParseErrors)
	}
	ambientFrom := time.Date(2026, 9, 27, 10, 0, 30, 0, time.UTC)
	ambientTo := time.Date(2026, 9, 27, 10, 1, 30, 0, time.UTC)
	var neutral, wipe *fight.Fight
	for i := range got.Closed {
		f := &got.Closed[i].Fight
		start := f.Start.UTC()
		switch {
		case f.Kind == fight.Encounter:
			if wipe != nil {
				t.Fatalf("the wipe split into fights %d and %d", wipe.Index, f.Index)
			}
			wipe = f
		case start.After(ambientFrom) && start.Before(ambientTo):
			t.Errorf("fight %d opened on the ambient NPCs' fight at %s", f.Index, start)
		case neutral == nil:
			neutral = f
		}
	}
	if neutral == nil {
		t.Fatal("the neutral pull made no fight")
	}
	if neutral.MobDeaths < 1 || neutral.GroupKills < 1 || neutral.GroupKills > neutral.MobDeaths {
		t.Errorf("neutral pull: mob deaths %d, group kills %d; want the group's kills counted",
			neutral.MobDeaths, neutral.GroupKills)
	}
	if wipe == nil {
		t.Fatal("no encounter")
	}
	if wipe.Kill {
		t.Error("the wipe reads as a kill")
	}
	want := playerDeathsInText(t, src, wipe.StartOffset, wipe.EndOffset)
	if want < 1 || wipe.Deaths != want || wipe.GroupDeaths != want {
		t.Errorf("wipe: deaths %d, group deaths %d; the text has %d deaths that are no feign",
			wipe.Deaths, wipe.GroupDeaths, want)
	}
}
