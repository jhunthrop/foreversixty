package coach

import (
	"testing"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

func TestLinks_OnlyWhenWindowsMeet(t *testing.T) {
	fs := []Finding{{ID: "f1", Code: CodeCoreCasts}, {ID: "f2", Code: CodeResource, Title: "Rage full for 9 seconds"}}
	gap := []ratingengine.Window{{FromMS: 160_000, ToMS: 171_000}} // 2:40–2:51
	meet := computeLinks(fs, map[string][]ratingengine.Window{"f1": gap, "f2": {{FromMS: 161_000, ToMS: 170_000}}})
	if len(meet) != 1 || meet[0].Kind != LinkGapOverlapsCap || meet[0].Findings != [2]string{"f1", "f2"} {
		t.Fatalf("%+v", meet)
	}
	if apart := computeLinks(fs, map[string][]ratingengine.Window{"f1": gap, "f2": {{FromMS: 60_000, ToMS: 69_000}}}); len(apart) != 0 {
		t.Fatalf("a cap at 1:00–1:09 never meets a gap at 2:40: %+v", apart)
	}
}

func TestLinks_GapCapSentenceIsTheSpecs(t *testing.T) {
	fs := []Finding{
		{ID: "f1", Code: CodeCoreCasts, Title: "Bloodthirst: 9 casts lost"},
		{ID: "f4", Code: CodeResource, Title: "Rage full for 14 seconds"},
	}
	got := computeLinks(fs, map[string][]ratingengine.Window{
		"f1": {{FromMS: 72_000, ToMS: 85_000}, {FromMS: 160_000, ToMS: 171_000}},
		"f4": {{FromMS: 161_000, ToMS: 170_000}},
	})
	if len(got) != 1 || got[0].Sentence != "The Bloodthirst gap at 2:40–2:51 overlaps the rage cap at 2:41–2:50." {
		t.Fatalf("%+v", got)
	}
}

func TestLinks_OutOfManaAtDeath(t *testing.T) {
	fs := []Finding{{ID: "f1", Code: CodeDeath}, {ID: "f2", Code: CodeResource, Title: "Out of mana from 3:05"}}
	got := computeLinks(fs, map[string][]ratingengine.Window{
		"f1": {{FromMS: 200_000, ToMS: 200_000}}, // the death instant, 3:20
		"f2": {{FromMS: 185_000, ToMS: 232_000}}, // out of mana from 3:05
	})
	if len(got) != 1 || got[0].Kind != LinkOutOfManaAtDeath || got[0].Sentence != "You died while out of mana." ||
		got[0].Findings != [2]string{"f2", "f1"} {
		t.Fatalf("%+v", got)
	}
	later := computeLinks(fs, map[string][]ratingengine.Window{
		"f1": {{FromMS: 100_000, ToMS: 100_000}},
		"f2": {{FromMS: 185_000, ToMS: 232_000}},
	})
	if len(later) != 0 {
		t.Fatalf("mana ran out after the death: %+v", later)
	}
}

func TestLinks_DeathInsideACastGap(t *testing.T) {
	fs := []Finding{{ID: "f1", Code: CodeDeath}, {ID: "f2", Code: CodeCoreCasts, Title: "Bloodthirst: 9 casts lost"}}
	gaps := []ratingengine.Window{{FromMS: 30_000, ToMS: 45_000}}
	got := computeLinks(fs, map[string][]ratingengine.Window{"f1": {{FromMS: 40_000, ToMS: 40_000}}, "f2": gaps})
	if len(got) != 1 || got[0].Kind != LinkDeathInGap || got[0].Sentence != "The Bloodthirst gap at 0:30–0:45 includes your death at 0:40." {
		t.Fatalf("%+v", got)
	}
	if outside := computeLinks(fs, map[string][]ratingengine.Window{"f1": {{FromMS: 50_000, ToMS: 50_000}}, "f2": gaps}); len(outside) != 0 {
		t.Fatalf("%+v", outside)
	}
}

func TestLinks_MissingFlaskCoversOutOfMana(t *testing.T) {
	oom := Finding{ID: "f1", Code: CodeResource, Title: "Out of mana from 1:00"}
	got := computeLinks([]Finding{oom, {ID: "f2", Code: CodePreparation, Title: TitleNoFlask}}, nil)
	if len(got) != 1 || got[0].Kind != LinkConsumableCovers || got[0].Findings != [2]string{"f2", "f1"} ||
		got[0].Sentence != "You pulled without a flask and your mana ran out from 1:00." {
		t.Fatalf("%+v", got)
	}
	capped := Finding{ID: "f1", Code: CodeResource, Title: "Rage full for 14 seconds"}
	if none := computeLinks([]Finding{capped, {ID: "f2", Code: CodePreparation, Title: TitleNoFlask}}, nil); len(none) != 0 {
		t.Fatalf("a flask doesn't cover rage: %+v", none)
	}
}

func TestLinks_OverhealOnTheMostCastHeal(t *testing.T) {
	in := holyOverheal(t, 10965, 23, 44, poolOverheal(10965, 22, 21))
	heal := Finding{ID: "f2", Code: CodeRanks, Title: "Greater Heal (Rank 4): 44% overheal",
		Facts: Facts{Spells: []SpellFact{{ID: 10965, Name: "Greater Heal", Rank: 4}}}}
	oom := Finding{ID: "f1", Code: CodeResource, Title: "Out of mana from 3:05"}
	got := overhealLink(in, []Finding{oom, heal})
	if len(got) != 1 || got[0].Kind != LinkOverhealOnMostCast || got[0].Findings != [2]string{"f2", "f1"} ||
		got[0].Sentence != "Greater Heal (Rank 4) is your most-cast heal with the most overheal." {
		t.Fatalf("%+v", got)
	}
	other := heal
	other.Facts.Spells = []SpellFact{{ID: 10964, Name: "Greater Heal", Rank: 3}}
	if none := overhealLink(in, []Finding{oom, other}); len(none) != 0 {
		t.Fatalf("Rank 3 is not the most-cast heal: %+v", none)
	}
}

func TestLinks_NumberedInOrder(t *testing.T) {
	ls := numberLinks([]Link{{Kind: LinkGapOverlapsCap}, {Kind: LinkDeathInGap}})
	if ls[0].ID != "l1" || ls[1].ID != "l2" {
		t.Fatalf("%+v", ls)
	}
}
