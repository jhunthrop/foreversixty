package request

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// Rows re-measured at sim-accuracy round 11 (design/reviews/2026-11-03-sim-audit-round11/ranks.md):
// each case is the id a spec's published rotation casts for one ranked spell at one level
// and preset, and the id it must no longer cast there.
func TestReMeasuredRankChoicesCastTheWinningRank(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset string
		level                     int
		want, not                 int32
	}{
		// Round 15 wave 2 (level-gap rule): rank 4 (2090, learned at 22) keeps none of the healing
		// bonus at 60, and the tank's Rejuvenation casts its authored rank 10 (9841) at bare 60.
		{"Rejuvenation at bare 60", "druid-restoration", "druid", BarePreset, 60, 9841, 2090},
		// Once the published Undead casts Dark Sacrifice when mana is low (casters lane,
		// db8fc067e), Devouring Plague rank 1 (2944) loses 4.4% at bare 60 to the top
		// rank (19280); Shadow Word: Death rank 2 still wins there.
		{"Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, 60, 19280, 2944},
		{"Shadow Word: Death at bare 60 keeps rank 2", "priest-shadow", "priest", BarePreset, 60, 1309633, 1309636},
		// Round 15 wave 2 (level-gap rule): Healing Wave rank 2 (332, learned at 6) keeps none of
		// the bonus at 60; the tank's wave is the top rank (10396), +11% at bare 60.
		{"Healing Wave at bare 60", "shaman-restoration", "shaman", BarePreset, 60, 10396, 332},
		// Round 15 wave 3 (level-gap rule): Scorch rank 1 (2948, learned at 22) loses 5% of its
		// spell power at 40 and 50% by 49; rank 2 (8444, learned at 28) keeps all of it to 45 and
		// ties or gains at bare 40 (+0.7% alliance, +0.0/+0.2% horde).
		{"Scorch at 40", "mage-fire", "mage", BarePreset, 40, 8444, 2948},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, "undead"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Rows re-measured at sim-accuracy round 12 (design/reviews/2026-11-03-sim-audit-round12/ranks.md)
// on the landing's rotations: rows are adopted beyond error at seeds 1 and 50001, and keyed by
// faction where Alliance and Horde differ beyond error (race since round 13).
func TestRound12RankChoicesCastTheWinningRank(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset, race string
		level                           int
		want, not                       int32
	}{
		// Round 15 wave 2 (level-gap rule): Mind Blast at 50 is the highest rank learned (10945):
		// ranks 2 and 3 keep 15% and 45% of spell power there, and the rows lost 3.3%. Round 19:
		// an Undead casts rank 6 (8106, whole bonus) at 50 and 51.
		{"Undead Mind Blast at 50", "priest-shadow", "priest", BarePreset, "undead", 50, 8106, 8103},
		{"Gnome Mind Blast at 50", "priest-shadow", "priest", BarePreset, "gnome", 50, 10945, 8102},
		{"Undead Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, "undead", 60, 19280, 19276},
		{"Gnome Devouring Plague at raid 60", "priest-shadow", "priest", RaidPreset, "gnome", 60, 19280, 19276},
		// Elemental's Flame Shock: rank 1 (8050) at 30 until client 1.60.1.70755 raised every
		// rank; the top rank learned (8053) wins there since (round 15, shaman).
		{"Flame Shock at 30", "shaman-elemental", "shaman", BarePreset, "orc", 30, 8053, 8050},
		// The top-rank Lightning Bolt line (15208, cast while mana is ahead) casts rank 7 (10391) at 50.
		// At bare 60 it casts rank 10 since client 1.60.1.70755's bolt retune (round 15, shaman).
		{"Lightning Bolt pacing line at 50", "shaman-elemental", "shaman", BarePreset, "orc", 50, 15207, 10391},
		{"Lightning Bolt pacing line at bare 60", "shaman-elemental", "shaman", BarePreset, "dwarf", 60, 15208, 10391},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// A request's rank choices follow its character's race: the same Shadow priest at 40 casts
// Mind Blast rank 4 as a Dwarf and rank 5 as a Gnome (round 24).
func TestBuildCastsTheRacesRankChoice(t *testing.T) {
	for race, want := range map[string]string{"dwarf": `"spellId":8104`, "gnome": `"spellId":8105`} {
		req := api.SimRequest{
			EngineVersion: enginever.Version, Spec: "priest-shadow",
			Source:    api.CharacterSource{Kind: api.SourceManual},
			Character: api.CharacterSpec{Name: "x", Race: race, Class: "priest", Level: 40},
			Encounter: api.DefaultEncounter(), Iterations: 3000, RandomSeed: 1,
			Preset: BarePreset,
		}
		built, err := Build(req)
		if err != nil {
			t.Fatal(err)
		}
		rot, err := protojson.Marshal(built.GetRaid().GetParties()[0].GetPlayers()[0].GetRotation())
		if err != nil {
			t.Fatal(err)
		}
		if got := strings.ReplaceAll(string(rot), " ", ""); !strings.Contains(got, want) {
			t.Errorf("%s at 40 does not cast %s", race, want)
		}
	}
}

// Rows re-judged per race at sim-accuracy round 13 (design/reviews/2026-11-03-sim-audit-round13/ranks.md):
// every row measured on every race of its class; a row names races only where a race loses
// with it, or another rank beats it, beyond error at seeds 1 and 50001.
func TestRound13RaceRankChoicesCastTheWinningRank(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset, race string
		level                           int
		want, not                       int32
	}{
		// Mind Blast at 40: rank 4 (8104) for a Dwarf or a Night Elf too, whom rank 3 (8103)
		// never beats; a Gnome casts the top rank, 6 (8106). Re-judged level by level at
		// sim-accuracy round 18. Round 24, on the talents at 40: a Gnome casts rank 5 (8105).
		{"Gnome Mind Blast at 40", "priest-shadow", "priest", BarePreset, "gnome", 40, 8105, 8106},
		{"Dwarf Mind Blast at 40", "priest-shadow", "priest", BarePreset, "dwarf", 40, 8104, 8103},
		{"Night Elf Mind Blast at 40", "priest-shadow", "priest", BarePreset, "night-elf", 40, 8104, 8103},
		// Round 15 wave 2 (level-gap rule): Mind Blast at 50 is the highest rank learned for every race.
		// Round 24: a Night Elf casts rank 5 (8105) at 50 and 51.
		{"Night Elf Mind Blast at 50", "priest-shadow", "priest", BarePreset, "night-elf", 50, 8105, 10945},
		{"Human Mind Blast at 50", "priest-shadow", "priest", BarePreset, "human", 50, 10945, 8103},
		// Devouring Plague at 50: rank 2 (19276) for every race since sim-accuracy round 18
		// (rank 1 for most races until then).
		{"Undead Devouring Plague at 50", "priest-shadow", "priest", BarePreset, "undead", 50, 19276, 2944},
		{"Night Elf Devouring Plague at 50", "priest-shadow", "priest", BarePreset, "night-elf", 50, 19276, 2944},
		// Devouring Plague at bare 60: rank 2 for every race but the Undead.
		{"Troll Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, "troll", 60, 19276, 19280},
		{"Night Elf Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, "night-elf", 60, 19276, 19280},
		// Shadow Word: Death at 50: rank 2 for every race, the Human too since sim-accuracy
		// round 18 (+2.2% at 50, +1.5% at 56 over the top rank learned).
		{"Human Shadow Word: Death at 50", "priest-shadow", "priest", BarePreset, "human", 50, 1309633, 1309635},
		{"Troll Shadow Word: Death at 50", "priest-shadow", "priest", BarePreset, "troll", 50, 1309633, 1309635},
		// Rows no race loses with lose their race key: Fire Blast rank 3 (2138) at 30 and Prayer
		// of Healing rank 1 at bare 60. (Corruption rank 3 at 40 was one; round 14 found it ties
		// for every race, and the row is gone: see below.)
		// Round 24 (on the talents at 30): rank 2 (2137) for every race at 30 to 35.
		{"Gnome Fire Blast at 30", "mage-fire", "mage", BarePreset, "gnome", 30, 2137, 8412},
		// Sim-accuracy round 18: rank 2 (996, whole at 60) beats rank 1 (596, 55% of its bonus) at bare 60.
		{"Gnome Prayer of Healing at bare 60", "priest-holy", "priest", BarePreset, "gnome", 60, 996, 596},
		// Protection's Holy Strike: the top rank at 40 too, since the 2026-11-11 paladins lane
		// fixed the formula (rank 1's saved mana no longer pays for its smaller flat Holy hit;
		// -4.2% to -4.7% threat at 40 for a Dwarf and an Undead).
		{"Dwarf Holy Strike at 40", "paladin-protection", "paladin", BarePreset, "dwarf", 40, 2495, 679},
		{"Holy Strike at 50 keeps the top rank", "paladin-protection", "paladin", BarePreset, "dwarf", 50, 5569, 679},
		// Lightwell rank 1 now loses or ties for every race at bare 60: the row is gone.
		{"Troll Lightwell at bare 60", "priest-holy", "priest", BarePreset, "troll", 60, 27871, 724},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Rows re-judged at sim-accuracy round 14 (design/reviews/2026-11-04-sim-audit-round14/data.md),
// each row against the default rule on its own at seeds 1 and 50001, 10000 iterations (the
// Arcane Shot rows are gone: hotfix_arcane_shot_rank_test.go); since Shadow keeps Vampiric Embrace up a Human or Gnome loses with
// Devouring Plague rank 2 at bare 60; Corruption rank 3 at 40 ties for every warlock race.
func TestRound14RankChoicesCastTheWinningRank(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset, race string
		level                           int
		want, not                       int32
	}{
		{"Human Devouring Plague at bare 60 is the top rank", "priest-shadow", "priest", BarePreset, "human", 60, 19280, 19276},
		{"Gnome Devouring Plague at bare 60 is the top rank", "priest-shadow", "priest", BarePreset, "gnome", 60, 19280, 19276},
		{"Troll Devouring Plague at bare 60 keeps rank 2", "priest-shadow", "priest", BarePreset, "troll", 60, 19276, 19280},
		{"Corruption at 40 is the default rank 4", "warlock-destruction", "warlock", BarePreset, "troll", 40, 7648, 6223},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Round 16 wave 2 (per-level sweep, spec-breakdown -level): the mage rows were measured only at a
// band's floor. Swept at every few levels, seeds 1 and 50001, both factions:
//   - Frost's Frostbolt: rank 4 (7322) beats rank 5 at 26-29 by 3 to 6%, and rank 5 (8406)
//     beats rank 6 at 32-33 by 3 to 4% (until Mage Armor at 34);
//   - Arcane's Frostbolt rank 4 row ends at 33: from 34 (Mage Armor) the highest rank wins,
//     +0.3 to +6% at 34-39 as rank 4 loses its spell power from 37;
//   - Fire's Scorch rank 2 (8444) row ends at 45: rank 3 (8445) wins from 46, +1.8% at 47 and
//     +4.3% at 49, as rank 2 loses its spell power from 45.
func TestRound16MageRankRowsFollowTheLevel(t *testing.T) {
	for _, c := range []struct {
		name, spec string
		level      int
		want, not  int32
	}{
		{"Frost Frostbolt at 26", "mage-frost", 26, 7322, 8406},
		{"Frost Frostbolt at 29", "mage-frost", 29, 8406, 7322}, // round 24: a Gnome casts rank 5 at 29
		{"Frost Frostbolt at 30", "mage-frost", 30, 8406, 7322},
		{"Frost Frostbolt at 32", "mage-frost", 32, 8406, 8407},
		{"Frost Frostbolt at 33", "mage-frost", 33, 8406, 8407},
		{"Frost Frostbolt at 34", "mage-frost", 34, 8407, 8406},
		{"Arcane Frostbolt at 33", "mage-arcane", 33, 7322, 8407},
		{"Arcane Frostbolt at 34", "mage-arcane", 34, 8407, 7322},
		{"Arcane Frostbolt at 39", "mage-arcane", 39, 8408, 7322},
		{"Fire Scorch at 45", "mage-fire", 45, 8444, 8445},
		{"Fire Scorch at 46", "mage-fire", 46, 8445, 8444},
		{"Fire Scorch at 49", "mage-fire", 49, 8445, 8444},
	} {
		ids := rotationIDs(t, c.spec, "mage", c.level, specRankChoices(c.spec, BarePreset, "gnome"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Round 17 (hunter lane, per-level sweep): since round 13 put Aimed Shot on the shared cooldown
// with Multi-Shot, it is a rotation shot, not only the pre-pull opener, and the bare bands run
// out of mana for one to two minutes. Rank 1 (19434: 75 mana, 40 damage) then beats the highest
// rank learned at every level from 28 (rank 2) at seeds 1 and 50001, 10000 iterations, Horde and
// Alliance: Beast Mastery +1.9 to +3.1% at 28-39 and +0.1 to +0.6% at 50-59 (40-49 already had
// the row), Marksmanship +1.4 to +3.9% at 28-49. Marksmanship at 50 ties (0.0 to +0.1), so its
// row ends at 49.
func TestRound17HunterAimedShotRankRowsFollowTheLevel(t *testing.T) {
	for _, c := range []struct {
		name, spec string
		level      int
		want, not  int32
	}{
		{"Beast Mastery Aimed Shot at 28", "hunter-beast-mastery", 28, 19434, 20900},
		{"Beast Mastery Aimed Shot at 36", "hunter-beast-mastery", 36, 19434, 20901},
		{"Beast Mastery Aimed Shot at 49", "hunter-beast-mastery", 49, 19434, 20902},
		{"Beast Mastery Aimed Shot at 50", "hunter-beast-mastery", 50, 19434, 20902},
		{"Beast Mastery Aimed Shot at 59", "hunter-beast-mastery", 59, 19434, 20903},
		{"Marksmanship Aimed Shot at 28", "hunter-marksmanship", 28, 19434, 20900},
		{"Marksmanship Aimed Shot at 39", "hunter-marksmanship", 39, 19434, 20901},
		{"Marksmanship Aimed Shot at 46", "hunter-marksmanship", 46, 19434, 20902},
		{"Marksmanship Aimed Shot at 50", "hunter-marksmanship", 50, 19434, 20902}, // round 22: rank 1 from 28 to 59
	} {
		ids := rotationIDs(t, c.spec, "hunter", c.level, specRankChoices(c.spec, BarePreset, "orc"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Round 18 (hunter lane, per-level sweep, guide round 15 HUNTER-R15-3): Marksmanship at 52-59,
// once rank 5 (20903, 52) is the newest, casts rank 1 again. Rank 1 gains +0.24 to +0.40% at 52,
// 53, 56 and 59, seeds 1 and 50001, Horde and Alliance, 10000 iterations (error 0.1 DPS on about
// 260). At 50 and 51, against rank 4, it tied (0.0 to +0.1) then; round 22 found rank 1 ahead there
// by about 5% once Aimed Shot stopped waiting for an auto shot (rank_choices_hunter_r22_test.go).
func TestRound18MarksmanshipAimedShotRankOneFrom52(t *testing.T) {
	for _, c := range []struct {
		name      string
		level     int
		want, not int32
	}{
		{"Marksmanship Aimed Shot at 50", 50, 19434, 20902}, // round 22
		{"Marksmanship Aimed Shot at 51", 51, 19434, 20902}, // round 22
		{"Marksmanship Aimed Shot at 52", 52, 19434, 20903},
		{"Marksmanship Aimed Shot at 56", 56, 19434, 20903},
		{"Marksmanship Aimed Shot at 59", 59, 19434, 20903},
	} {
		ids := rotationIDs(t, "hunter-marksmanship", "hunter", c.level, specRankChoices("hunter-marksmanship", BarePreset, "orc"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Round 20 (priest lane, guide round 17 R17-6): Devouring Plague at 28 and 29, where rank 2
// (19276, learned at 28) is the newest. On the ladder leveling build at that level (Mind Flay
// from 28), 10000 iterations, seeds 1 and 50001: a Gnome gains +0.2% / +0.6% with rank 1
// (2944) and a Troll +0.3% / +0.2%, both seeds, beyond the 0.03 DPS error; an Undead loses
// 0.6% with it, and a Human, Dwarf or Night Elf ties (within 0.1 DPS), so they keep the
// default rank 2. The row starts at 28 for those two races only.
func TestRound20DevouringPlagueRankOneAt28ForGnomeAndTroll(t *testing.T) {
	for _, c := range []struct {
		name, race string
		level      int
		want, not  int32
	}{
		{"Gnome Devouring Plague at 28", "gnome", 28, 2944, 19276},
		{"Gnome Devouring Plague at 29", "gnome", 29, 2944, 19276},
		{"Troll Devouring Plague at 28", "troll", 28, 2944, 19276},
		// Round 24, on the talents at 29: a Troll casts rank 2 (+0.3% / +0.4%).
		{"Troll Devouring Plague at 29", "troll", 29, 19276, 2944},
		{"Human Devouring Plague at 28 is the top rank", "human", 28, 19276, 2944},
		{"Undead Devouring Plague at 29 is the top rank", "undead", 29, 19276, 2944},
		{"Gnome Devouring Plague at 30 is rank 1 as before", "gnome", 30, 2944, 19276},
	} {
		ids := rotationIDs(t, "priest-shadow", "priest", c.level, specRankChoices("priest-shadow", BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Sim-accuracy round 24 (casters lane): the mage and warlock rows re-judged on the talents the
// player has at each level (spec-breakdown -level N -rank, frozen BiS, 10,000 iterations, seeds 1
// and 50001, every race of the class), and the rank sweep at the two levels after each new rank
// (design/reviews/2026-11-05-sim-audit-round24/casters.md).
func TestRound24CasterRankRows(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset, race string
		level                           int
		want, not                       int32
	}{
		// Fire Blast rank 2 (2137) at 30 to 35, rank 3 (2138) at 36 to 39, judged jointly with Pyroblast.
		{"Gnome Fire Blast at 30", "mage-fire", "mage", BarePreset, "gnome", 30, 2137, 2138},
		{"Undead Fire Blast at 35", "mage-fire", "mage", BarePreset, "undead", 35, 2137, 2138},
		{"Human Fire Blast at 36", "mage-fire", "mage", BarePreset, "human", 36, 2138, 2137},
		// Scorch rank 1 (2948) from 28, when rank 2 arrives.
		{"Orc Scorch at 28", "mage-fire", "mage", BarePreset, "orc", 28, 2948, 8444},
		// Pyroblast: rank 2 (12505) at 30 to 33 and 36 to 37, rank 3 (12522) at 34, 35 and 38 to 40.
		{"Human Pyroblast at 30", "mage-fire", "mage", BarePreset, "human", 30, 12505, 12522},
		{"Gnome Pyroblast at 34", "mage-fire", "mage", BarePreset, "gnome", 34, 12522, 12505},
		{"Gnome Pyroblast at 36", "mage-fire", "mage", BarePreset, "gnome", 36, 12505, 12523},
		{"Orc Pyroblast at 39", "mage-fire", "mage", BarePreset, "orc", 39, 12522, 12523},
		// Fireball rank 4 (3140) at 24 for every race but a Gnome.
		{"Undead Fireball at 24", "mage-fire", "mage", BarePreset, "undead", 24, 3140, 8400},
		{"Gnome Fireball at 24", "mage-fire", "mage", BarePreset, "gnome", 24, 8400, 3140},
		// Arcane's Frostbolt rank 4 (7322) from 29 (round 30: rank 1 at 24 to 28), when rank 5 (8406) arrives at 26, to 33.
		{"Troll Arcane Frostbolt at 29", "mage-arcane", "mage", BarePreset, "troll", 29, 7322, 8406},
		{"Human Arcane Frostbolt at 29", "mage-arcane", "mage", BarePreset, "human", 29, 7322, 8406},
		// Frost's Frostbolt at 29: rank 5 for a Gnome, Orc or Undead; rank 4 for a Human to 31.
		{"Gnome Frost Frostbolt at 29", "mage-frost", "mage", BarePreset, "gnome", 29, 8406, 7322},
		{"Troll Frost Frostbolt at 29", "mage-frost", "mage", BarePreset, "troll", 29, 7322, 8406},
		{"Human Frost Frostbolt at 30", "mage-frost", "mage", BarePreset, "human", 30, 7322, 8406},
		{"Gnome Frost Frostbolt at 30", "mage-frost", "mage", BarePreset, "gnome", 30, 8406, 7322},
		// Destruction's Shadow Bolt rank 2 (695) at 20 to 23; rank 3 (705) at 24 for a Gnome or
		// Human; rank 4 (1088) at 28 (a Troll in round 24, every race since round 28).
		{"Orc Destruction Shadow Bolt at 20", "warlock-destruction", "warlock", BarePreset, "orc", 20, 695, 705},
		{"Human Destruction Shadow Bolt at 24", "warlock-destruction", "warlock", BarePreset, "human", 24, 705, 1088},
		{"Troll Destruction Shadow Bolt at 24", "warlock-destruction", "warlock", BarePreset, "troll", 24, 1088, 705},
		{"Troll Destruction Shadow Bolt at 28", "warlock-destruction", "warlock", BarePreset, "troll", 28, 1088, 1106},
		{"Orc Destruction Shadow Bolt at 28", "warlock-destruction", "warlock", BarePreset, "orc", 28, 1088, 1106}, // every race since round 28
		// Demonology's Shadow Bolt rank 5 (1106) at 36 and 37, when rank 6 (7641) arrives.
		{"Undead Demonology Shadow Bolt at 36", "warlock-demonology", "warlock", BarePreset, "undead", 36, 1106, 7641},
		// Fire's Fireball rank 10 (10150) at bare 60; the raid set keeps rank 11 (10151).
		{"Human Fireball at bare 60", "mage-fire", "mage", BarePreset, "human", 60, 10150, 10151},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
