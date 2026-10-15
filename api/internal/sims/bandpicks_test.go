package sims

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// Task L6.7: the band's picks the character page's gear panel shows are the data addon's:
// bis/<spec>.json's bare entry for the band and faction, through L1.4's unconfirmed walk and
// the band-wide rules. These run against the repository's real data/builds.

const picksBuild = "1.60.1.70291"

func TestBandPicksAreTheBareEntryThroughTheUnconfirmedWalk(t *testing.T) {
	useRepoGearData(t)
	picks, err := BandPicks(t.Context(), picksBuild, "warrior-arms", 60, "alliance")
	if err != nil {
		t.Fatal(err)
	}
	head := picks["head"]
	if head.ItemID != 12640 || head.Name != "Lionheart Helm" || head.Source != "Blacksmithing" || head.Chance != nil || head.Promoted {
		t.Errorf("head = %+v, want Lionheart Helm (12640) from Blacksmithing, no chance, not promoted", head)
	}
	// leveling-bis drops client-unconfirmed items before ranking (ruling 2), so the file's
	// own pick stands: Darkmoon Card: Heroism, a quest reward with no drop chance, not
	// promoted. (Before the ranker change the own pick, Hand of Justice, was unconfirmed and
	// the walk promoted Frozen Heart of the Mountain; the walk's rules are pinned on
	// hand-built rows below.)
	tr := picks["trinket2"]
	if tr.ItemID != 19287 || tr.Name != "Darkmoon Card: Heroism" || tr.Source != "Darkmoon Warlords Deck" || tr.Chance != nil || tr.Promoted {
		t.Errorf("trinket2 = %+v, want Darkmoon Card: Heroism (19287), not promoted, no chance", tr)
	}
	if p, ok := picks["off_hand"]; ok {
		t.Errorf("off_hand = %+v, want no pick (the file's row has no item_id)", p)
	}
	if len(picks) != 16 {
		t.Errorf("%d picks, want 16 (every slot but the off-hand)", len(picks))
	}

	// Band 20: the original pick keeps its drop chance, and so does the main hand: the ranker's
	// rare-spawn rule dropped the former pick (a rare spawn's drop, id 23171) and Smite's
	// Mighty Hammer, a Deadmines boss drop at 20%, stands in the file, not promoted.
	low, err := BandPicks(t.Context(), picksBuild, "warrior-arms", 20, "alliance")
	if err != nil {
		t.Fatal(err)
	}
	chest := low["chest"]
	if chest.ItemID != 6627 || chest.Source != "Wailing Caverns: Mutanus the Devourer" || chest.Chance == nil || *chest.Chance != 0.333333 {
		t.Errorf("band 20 chest = %+v, want 6627 from Mutanus with chance 0.333333", chest)
	}
	mh := low["main_hand"]
	if mh.ItemID != 7230 || mh.Name != "Smite's Mighty Hammer" || mh.Source != "The Deadmines: Mr. Smite" ||
		mh.Chance == nil || *mh.Chance != 0.2 || mh.Promoted {
		t.Errorf("band 20 main_hand = %+v, want Smite's Mighty Hammer (7230), not promoted, chance 0.2", mh)
	}
}

// Fix round 1, item 1: the pair rule refuses the mate's item by id or name, unique or not,
// and the band's own confirmed picks are fixed before any promotion is checked against them.
func TestBandPicksNeverRepeatThePairMatesItem(t *testing.T) {
	useRepoGearData(t)
	// mage-fire 60 alliance: the ranker ranks confirmed items only, so both own picks stand
	// (Talisman of Ascendance, Serenity Field) and the pair never repeats an item.
	fire, err := BandPicks(t.Context(), picksBuild, "mage-fire", 60, "alliance")
	if err != nil {
		t.Fatal(err)
	}
	if fire["trinket1"].ItemID != 22678 || fire["trinket2"].ItemID != 272439 {
		t.Errorf("mage-fire 60 alliance trinkets = %d, %d; want 22678, 272439", fire["trinket1"].ItemID, fire["trinket2"].ItemID)
	}
	// priest-holy 50 horde: Snake Hoop is finger1's own pick and Voodoo Band finger2's, each
	// kept as the file has it (before the ranker change finger2 had no confirmed candidate).
	holy, err := BandPicks(t.Context(), picksBuild, "priest-holy", 50, "horde")
	if err != nil {
		t.Fatal(err)
	}
	if p, ok := holy["finger2"]; holy["finger1"].ItemID != 6750 || !ok || p.ItemID != 1996 || p.Promoted {
		t.Errorf("priest-holy 50 horde rings = %d, %+v; want 6750 and Voodoo Band (1996), not promoted", holy["finger1"].ItemID, p)
	}
}

func TestBandPicksRefuseWhatTheDataDoesNotHave(t *testing.T) {
	useRepoGearData(t)
	if _, err := BandPicks(t.Context(), picksBuild, "warrior-arms", 25, "alliance"); !errors.Is(err, ErrNoBestInSlot) {
		t.Errorf("a band with no entry: err = %v, want ErrNoBestInSlot", err)
	}
	if _, err := BandPicks(t.Context(), picksBuild, "warrior-gladiator", 60, "alliance"); !errors.Is(err, ErrNoBestInSlot) {
		t.Errorf("a spec with no bis file: err = %v, want ErrNoBestInSlot", err)
	}
	if _, err := BandPicks(t.Context(), "0.0.0.1", "warrior-arms", 60, "alliance"); !errors.Is(err, ErrGearUnscorable) {
		t.Errorf("an unknown build: err = %v, want ErrGearUnscorable", err)
	}
}

// Every real spec, band and faction: no pick is client-unconfirmed, no unique item fills
// both slots of a pair, and no off-hand is picked beside a two-handed main hand.
func TestNoRealBandPickBreaksTheBandWideRules(t *testing.T) {
	useRepoGearData(t)
	b, ok := repoTrees.Build(picksBuild)
	if !ok {
		t.Fatal("no build")
	}
	files, err := filepath.Glob(filepath.Join(repoBuilds, picksBuild, "bis", "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("bis files: %v %v", files, err)
	}
	checked := 0
	for _, f := range files {
		spec := strings.TrimSuffix(filepath.Base(f), ".json")
		classID, ok := classIDBySlug(b, strings.SplitN(spec, "-", 2)[0])
		if !ok {
			t.Fatalf("%s: no class", spec)
		}
		for band := FirstGearBand; band <= FullGearBand; band += 10 {
			for _, faction := range []string{"alliance", "horde"} {
				picks, err := BandPicks(t.Context(), picksBuild, spec, band, faction)
				if errors.Is(err, ErrNoBestInSlot) {
					continue
				}
				if err != nil {
					t.Fatal(err)
				}
				checked++
				where := fmt.Sprintf("%s %s %d", spec, faction, band)
				for slot, p := range picks {
					if it, ok := b.Item(classID, p.ItemID); ok && it.ClientUnconfirmed {
						t.Errorf("%s %s: %d is client-unconfirmed", where, slot, p.ItemID)
					}
				}
				for _, pair := range [][2]string{{"finger1", "finger2"}, {"trinket1", "trinket2"}, {"main_hand", "off_hand"}} {
					a, ok1 := picks[pair[0]]
					z, ok2 := picks[pair[1]]
					if !ok1 || !ok2 {
						continue
					}
					ia, _ := b.Item(classID, a.ItemID)
					iz, _ := b.Item(classID, z.ItemID)
					if a.ItemID == z.ItemID || (ia.Name != "" && ia.Name == iz.Name) {
						t.Errorf("%s: %d and %d (%q) in both %s and %s", where, a.ItemID, z.ItemID, ia.Name, pair[0], pair[1])
					}
				}
				if mh, ok := picks["main_hand"]; ok {
					if it, _ := b.Item(classID, mh.ItemID); it.TwoHand {
						if oh, ok := picks["off_hand"]; ok {
							t.Errorf("%s: off-hand %d beside two-hander %d", where, oh.ItemID, mh.ItemID)
						}
					}
				}
			}
		}
	}
	if checked < 200 {
		t.Errorf("checked %d bands, want every spec's ten (about 270)", checked)
	}
}

func ptrInt(v int) *int { return &v }

// The walk's rules on hand-built rows, so each one is pinned whatever the real data holds:
// L1.4's two phases (the band's confirmed own picks are fixed, then each other slot promotes
// its first confirmed candidate the rules allow beside them), the pair rule by id or name,
// and the two-hand rule (a promoted two-hander empties the off-hand, as leveling-bis's
// confirmed.go does).
func TestTheWalkKeepsTheBandWideRules(t *testing.T) {
	items := map[int]trees.Item{
		1:  {ID: 1, Name: "Plain Ring"},
		2:  {ID: 2, Name: "Plain Ring"}, // the same item under another id
		3:  {ID: 3, Name: "Second Ring"},
		4:  {ID: 4, Name: "Unconfirmed Ring", ClientUnconfirmed: true},
		5:  {ID: 5, Name: "Unconfirmed Axe", TwoHand: true, ClientUnconfirmed: true},
		6:  {ID: 6, Name: "Confirmed Greatsword", TwoHand: true},
		7:  {ID: 7, Name: "Buckler"},
		8:  {ID: 8, Name: "Unconfirmed Cloak", ClientUnconfirmed: true},
		9:  {ID: 9, Name: "Plain Trinket"},
		10: {ID: 10, Name: "Unconfirmed Trinket", ClientUnconfirmed: true},
		11: {ID: 11, Name: "Other Trinket"},
	}
	lookup := func(id int) (trees.Item, bool) { it, ok := items[id]; return it, ok }
	chance := 0.25
	dropped := 0.05
	chances := func(id int, source string) *float64 {
		if id == 11 && source == "Boss: Drop" {
			return &dropped
		}
		return nil
	}
	rows := []bisRow{
		// finger1's own pick is unconfirmed; its first confirmed candidate is finger2's own
		// (confirmed, so already fixed) Plain Ring under another id: refused by name.
		{Slot: "finger1", ItemID: ptrInt(4), Source: "A", Chance: &chance,
			RunnersUp: &[]bisCandidate{{ItemID: 1, Source: "B"}, {ItemID: 3, Source: "C"}}},
		{Slot: "finger2", ItemID: ptrInt(2), Source: "B"},
		// trinket2's own pick is unconfirmed; its runner-up is trinket1's plain (not unique)
		// trinket: refused by id; the next one carries its own drop chance.
		{Slot: "trinket1", ItemID: ptrInt(9), Source: "D"},
		{Slot: "trinket2", ItemID: ptrInt(10), Source: "E", Chance: &chance,
			RunnersUp: &[]bisCandidate{{ItemID: 9, Source: "D"}, {ItemID: 11, Source: "Boss: Drop"}}},
		// The unconfirmed two-hander gives way to a confirmed two-hander (from the
		// alternatives: an older file with no runners_up), and the off-hand's own confirmed
		// pick goes with it.
		{Slot: "main_hand", ItemID: ptrInt(5), Source: "F",
			Alternatives: []bisCandidate{{ItemID: 6, ItemName: "Greatsword", Source: "G"}}},
		{Slot: "off_hand", ItemID: ptrInt(7), Source: "H"},
		// Every candidate unconfirmed: no pick.
		{Slot: "back", ItemID: ptrInt(8), Source: "I", RunnersUp: &[]bisCandidate{}},
		// No item_id: nothing eligible at the band.
		{Slot: "neck"},
	}
	picks := walkBand(rows, lookup, chances)
	var slots []string
	for s := range picks {
		slots = append(slots, s)
	}
	sort.Strings(slots)
	if got := strings.Join(slots, ","); got != "finger1,finger2,main_hand,trinket1,trinket2" {
		t.Fatalf("picked slots = %s, want finger1,finger2,main_hand,trinket1,trinket2", got)
	}
	if p := picks["finger1"]; p.ItemID != 3 || p.Source != "C" || p.Chance != nil || !p.Promoted {
		t.Errorf("finger1 = %+v, want the second ring, promoted, no chance", p)
	}
	if p := picks["finger2"]; p.ItemID != 2 || p.Promoted {
		t.Errorf("finger2 = %+v, want its own Plain Ring, fixed first", p)
	}
	if p := picks["trinket2"]; p.ItemID != 11 || p.Chance == nil || *p.Chance != 0.05 || !p.Promoted {
		t.Errorf("trinket2 = %+v, want the other trinket with its own 5%% chance", p)
	}
	if p := picks["main_hand"]; p.ItemID != 6 || p.Name != "Confirmed Greatsword" || p.Source != "G" || !p.Promoted {
		t.Errorf("main_hand = %+v, want the confirmed greatsword, named from the tables", p)
	}
}

// ItemGearScore is the per-item term FullGearScore sums: the band-60 mark is the sum of
// ItemGearScore over its winning set (horde raid for warrior-arms).
func TestItemGearScoreIsTheTermTheMarkSums(t *testing.T) {
	useRepoGearData(t)
	full, err := FullGearScore(t.Context(), picksBuild, "warrior-arms", FullGearBand)
	if err != nil {
		t.Fatal(err)
	}
	entries, err := loadBisEntries(repoBuilds, picksBuild, "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	best := 0.0
	for _, e := range entries {
		if e.Band != FullGearBand {
			continue
		}
		sum := 0.0
		for _, r := range e.Slots {
			if r.ItemID == nil {
				continue
			}
			v, ok, err := ItemGearScore(t.Context(), picksBuild, "warrior-arms", *r.ItemID)
			if err != nil || !ok {
				t.Fatalf("%d: %v %v", *r.ItemID, ok, err)
			}
			sum += v
		}
		best = math.Max(best, sum)
	}
	if math.Abs(best-full) > scoreTolerance {
		t.Fatalf("the best band-60 set sums to %v by ItemGearScore, FullGearScore says %v", best, full)
	}
	if _, ok, err := ItemGearScore(t.Context(), picksBuild, "warrior-arms", 999999999); ok || err != nil {
		t.Errorf("an item the tables lack: ok %v err %v, want not ok and no error", ok, err)
	}
}

func TestRaceFactionAndGearItemReadTheBuild(t *testing.T) {
	useRepoGearData(t)
	if got := RaceFaction(picksBuild, "human"); got != "alliance" {
		t.Errorf("human = %q, want alliance", got)
	}
	if got := RaceFaction(picksBuild, "orc"); got != "horde" {
		t.Errorf("orc = %q, want horde", got)
	}
	if got := RaceFaction(picksBuild, "murloc"); got != "" {
		t.Errorf("murloc = %q, want empty", got)
	}
	it, ok := GearItem(picksBuild, "warrior", 12640)
	if !ok || it.Name != "Lionheart Helm" || it.ItemLevel == 0 {
		t.Errorf("GearItem(12640) = %+v %v", it, ok)
	}
}
