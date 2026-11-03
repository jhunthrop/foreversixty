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
		// Lava Burst rank 2 (1238299) beats rank 1 (408490) from 50, beyond error at both seeds.
		{"Lava Burst at 50", "shaman-elemental", "shaman", BarePreset, 50, 1238299, 408490},
		{"Lava Burst at 59", "shaman-elemental", "shaman", BarePreset, 59, 1238299, 408490},
		{"Lava Burst at 45 keeps rank 1", "shaman-elemental", "shaman", BarePreset, 45, 408490, 1238299},
		// Since Swiftmend leaves the effect up (fork 9badd243f), Rejuvenation rank 10 (9841)
		// loses to the rotation's own rank 4 (2090) at bare 60.
		{"Rejuvenation at bare 60", "druid-restoration", "druid", BarePreset, 60, 2090, 9841},
		// Once the published Undead casts Dark Sacrifice when mana is low (casters lane,
		// db8fc067e), Devouring Plague rank 1 (2944) loses 4.4% at bare 60 to the top
		// rank (19280); Shadow Word: Death rank 2 still wins there.
		{"Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, 60, 19280, 2944},
		{"Shadow Word: Death at bare 60 keeps rank 2", "priest-shadow", "priest", BarePreset, 60, 1309633, 1309636},
		// Without the level penalty on heals (healers lane, 074edfadf), Healing Wave rank 2
		// (332) beats rank 3 (547) by 5.6% at bare 60, and the fight still runs dry.
		{"Healing Wave at bare 60", "shaman-restoration", "shaman", BarePreset, 60, 332, 547},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, HordeFaction))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Rows re-measured at sim-accuracy round 12 (design/reviews/2026-11-03-sim-audit-round12/ranks.md)
// on the landing's rotations: rows are adopted beyond error at seeds 1 and 50001, and keyed by
// faction where Alliance and Horde differ beyond error.
func TestRound12RankChoicesCastTheWinningRank(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset, faction string
		level                              int
		want, not                          int32
	}{
		// Mind Blast at 50: rank 3 (8103) for Horde (Undead with Dark Sacrifice), rank 2 (8102) for Alliance.
		{"Horde Mind Blast at 50", "priest-shadow", "priest", BarePreset, HordeFaction, 50, 8103, 8102},
		{"Alliance Mind Blast at 50", "priest-shadow", "priest", BarePreset, AllianceFaction, 50, 8102, 8103},
		// Devouring Plague at bare 60: Alliance gains with rank 2 (19276); Horde casts the top rank.
		{"Alliance Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, AllianceFaction, 60, 19276, 19280},
		{"Horde Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, HordeFaction, 60, 19280, 19276},
		{"Alliance Devouring Plague at raid 60", "priest-shadow", "priest", RaidPreset, AllianceFaction, 60, 19280, 19276},
		// Beast Mastery's Arcane Shot: rank 3 (14282) at 40, rank 1 (3044) at 50.
		{"Arcane Shot at 40", "hunter-beast-mastery", "hunter", BarePreset, HordeFaction, 40, 14282, 14284},
		{"Arcane Shot at 50", "hunter-beast-mastery", "hunter", BarePreset, AllianceFaction, 50, 3044, 14285},
		// Fire's Fire Blast rank 3 (2138) at 30 for Horde only; Pyroblast keeps its rank there.
		{"Horde Fire Blast at 30", "mage-fire", "mage", BarePreset, HordeFaction, 30, 2138, 8412},
		{"Alliance Fire Blast at 30", "mage-fire", "mage", BarePreset, AllianceFaction, 30, 8412, 2138},
		// Elemental's Flame Shock rank 1 (8050) at 30.
		{"Flame Shock at 30", "shaman-elemental", "shaman", BarePreset, HordeFaction, 30, 8050, 8053},
		// The top-rank Lightning Bolt line (15208, cast while mana is ahead) casts rank 7 (10391) at 50 and bare 60.
		{"Lightning Bolt pacing line at 50", "shaman-elemental", "shaman", BarePreset, HordeFaction, 50, 10391, 15207},
		{"Lightning Bolt pacing line at bare 60", "shaman-elemental", "shaman", BarePreset, AllianceFaction, 60, 10391, 15208},
		// Bare-60 healer rows the factions split: Binding Heal rank 1 only for Alliance,
		// Prayer of Healing and Lightwell rank 1 only for Horde.
		{"Alliance Binding Heal at bare 60", "priest-discipline", "priest", BarePreset, AllianceFaction, 60, 401937, 1240774},
		{"Horde Binding Heal at bare 60", "priest-discipline", "priest", BarePreset, HordeFaction, 60, 1240774, 401937},
		{"Horde Prayer of Healing at bare 60", "priest-holy", "priest", BarePreset, HordeFaction, 60, 596, 10961},
		{"Alliance Prayer of Healing at bare 60", "priest-holy", "priest", BarePreset, AllianceFaction, 60, 10961, 596},
		{"Horde Lightwell at bare 60", "priest-holy", "priest", BarePreset, HordeFaction, 60, 724, 27871},
		{"Alliance Lightwell at bare 60", "priest-holy", "priest", BarePreset, AllianceFaction, 60, 27871, 724},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, c.faction))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// A request's rank choices follow its character's faction: the same Shadow priest at 50 casts
// Mind Blast rank 3 as an Undead and rank 2 as a Human.
func TestBuildCastsTheFactionsRankChoice(t *testing.T) {
	for race, want := range map[string]string{"undead": `"spellId":8103`, "human": `"spellId":8102`} {
		req := api.SimRequest{
			EngineVersion: enginever.Version, Spec: "priest-shadow",
			Source:    api.CharacterSource{Kind: api.SourceManual},
			Character: api.CharacterSpec{Name: "x", Race: race, Class: "priest", Level: 50},
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
			t.Errorf("%s at 50 does not cast %s", race, want)
		}
	}
}
