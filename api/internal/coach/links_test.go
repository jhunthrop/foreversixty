package coach

import (
	"testing"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

// win is one window.
func win(from, to int64) []ratingengine.Window {
	return []ratingengine.Window{{FromMS: from, ToMS: to}}
}

// capped and oom are F4 outcomes of each kind over ws.
func capped(bar string, ws []ratingengine.Window) outcome {
	return outcome{Windows: ws, Resource: resourceWaste{Kind: resourceCapped, Name: bar}}
}

func oom(ws []ratingengine.Window) outcome {
	return outcome{Windows: ws, Resource: resourceWaste{Kind: resourceOutOfMana, Name: "mana"}}
}

func TestLinks_OnlyWhenWindowsMeet(t *testing.T) {
	fs := []Finding{{ID: "f1", Code: CodeCoreCasts}, {ID: "f2", Code: CodeResource}}
	gap := win(160_000, 171_000) // 2:40–2:51
	meet := computeLinks(fs, map[string]outcome{"f1": {Windows: gap}, "f2": capped("rage", win(161_000, 170_000))})
	if len(meet) != 1 || meet[0].Kind != LinkGapOverlapsCap || meet[0].Findings != [2]string{"f1", "f2"} {
		t.Fatalf("%+v", meet)
	}
	if apart := computeLinks(fs, map[string]outcome{"f1": {Windows: gap}, "f2": capped("rage", win(60_000, 69_000))}); len(apart) != 0 {
		t.Fatalf("a cap at 1:00–1:09 never meets a gap at 2:40: %+v", apart)
	}
}

func TestLinks_ReadTheTypedResourceNotTheTitle(t *testing.T) {
	// a title that reads like a cap does not make one; the outcome's kind does
	fs := []Finding{{ID: "f1", Code: CodeCoreCasts}, {ID: "f2", Code: CodeResource, Title: "Rage full for 9 seconds"}}
	ev := map[string]outcome{"f1": {Windows: win(160_000, 171_000)}, "f2": oom(win(161_000, 170_000))}
	if got := computeLinks(fs, ev); len(got) != 0 {
		t.Fatalf("out of mana is no cap: %+v", got)
	}
}

func TestLinks_GapCapSentenceIsTheSpecs(t *testing.T) {
	fs := []Finding{
		{ID: "f1", Code: CodeCoreCasts, Facts: Facts{Spells: []SpellFact{{ID: 23894, Name: "Bloodthirst"}}}},
		{ID: "f4", Code: CodeResource},
	}
	got := computeLinks(fs, map[string]outcome{
		"f1": {Windows: []ratingengine.Window{{FromMS: 72_000, ToMS: 85_000}, {FromMS: 160_000, ToMS: 171_000}}},
		"f4": capped("rage", win(161_000, 170_000)),
	})
	if len(got) != 1 || got[0].Sentence != "The Bloodthirst gap at 2:40–2:51 overlaps the rage cap at 2:41–2:50." {
		t.Fatalf("%+v", got)
	}
}

func TestLinks_OutOfManaAtDeath(t *testing.T) {
	fs := []Finding{{ID: "f1", Code: CodeDeath}, {ID: "f2", Code: CodeResource}}
	got := computeLinks(fs, map[string]outcome{
		"f1": {Windows: win(200_000, 200_000)}, // the death instant, 3:20
		"f2": oom(win(185_000, 232_000)),       // out of mana from 3:05
	})
	if len(got) != 1 || got[0].Kind != LinkOutOfManaAtDeath || got[0].Sentence != "You died while out of mana." ||
		got[0].Findings != [2]string{"f2", "f1"} {
		t.Fatalf("%+v", got)
	}
	later := computeLinks(fs, map[string]outcome{"f1": {Windows: win(100_000, 100_000)}, "f2": oom(win(185_000, 232_000))})
	if len(later) != 0 {
		t.Fatalf("mana ran out after the death: %+v", later)
	}
}

func TestLinks_DeathInsideACastGap(t *testing.T) {
	fs := []Finding{{ID: "f1", Code: CodeDeath}, {ID: "f2", Code: CodeCoreCasts, Facts: Facts{Spells: []SpellFact{{Name: "Bloodthirst"}}}}}
	gaps := win(30_000, 45_000)
	got := computeLinks(fs, map[string]outcome{"f1": {Windows: win(40_000, 40_000)}, "f2": {Windows: gaps}})
	if len(got) != 1 || got[0].Kind != LinkDeathInGap || got[0].Sentence != "The Bloodthirst gap at 0:30–0:45 includes your death at 0:40." {
		t.Fatalf("%+v", got)
	}
	if outside := computeLinks(fs, map[string]outcome{"f1": {Windows: win(50_000, 50_000)}, "f2": {Windows: gaps}}); len(outside) != 0 {
		t.Fatalf("%+v", outside)
	}
}

func TestLinks_MissingFlaskCoversOutOfMana(t *testing.T) {
	fs := []Finding{{ID: "f1", Code: CodeResource}, {ID: "f2", Code: CodePreparation}}
	got := computeLinks(fs, map[string]outcome{"f1": oom(win(60_000, 120_000)), "f2": {MissingFlask: true}})
	if len(got) != 1 || got[0].Kind != LinkConsumableCovers || got[0].Findings != [2]string{"f2", "f1"} ||
		got[0].Sentence != "You pulled without a flask and your mana ran out from 1:00." {
		t.Fatalf("%+v", got)
	}
	// the flask is missing although another issue titles the finding: still linked
	fs[1].Title = "Main Hand has no enchant"
	if again := computeLinks(fs, map[string]outcome{"f1": oom(win(60_000, 120_000)), "f2": {MissingFlask: true}}); len(again) != 1 {
		t.Fatalf("%+v", again)
	}
	if none := computeLinks(fs, map[string]outcome{"f1": capped("rage", win(60_000, 70_000)), "f2": {MissingFlask: true}}); len(none) != 0 {
		t.Fatalf("a flask doesn't cover rage: %+v", none)
	}
	if none := computeLinks(fs, map[string]outcome{"f1": oom(win(60_000, 120_000)), "f2": {}}); len(none) != 0 {
		t.Fatalf("the flask was there: %+v", none)
	}
}

func TestLinks_OverhealOnTheMostCastHeal(t *testing.T) {
	in := holyOverheal(t, 10965, 23, 44, poolOverheal(10965, 22, 21))
	heal := Finding{ID: "f2", Code: CodeRanks, Facts: Facts{Spells: []SpellFact{{ID: 10965, Name: "Greater Heal", Rank: 4}}}}
	res := Finding{ID: "f1", Code: CodeResource}
	ev := map[string]outcome{"f1": oom(win(185_000, 232_000))}
	got := overhealLink(in, []Finding{res, heal}, ev)
	if len(got) != 1 || got[0].Kind != LinkOverhealOnMostCast || got[0].Findings != [2]string{"f2", "f1"} ||
		got[0].Sentence != "Greater Heal (Rank 4) is your most-cast heal with the most overheal." {
		t.Fatalf("%+v", got)
	}
	other := heal
	other.Facts.Spells = []SpellFact{{ID: 10964, Name: "Greater Heal", Rank: 3}}
	if none := overhealLink(in, []Finding{res, other}, ev); len(none) != 0 {
		t.Fatalf("Rank 3 is not the most-cast heal: %+v", none)
	}
	if none := overhealLink(in, []Finding{res, heal}, map[string]outcome{"f1": capped("rage", win(0, 10_000))}); len(none) != 0 {
		t.Fatalf("no out-of-mana finding: %+v", none)
	}
}

func TestLinks_NumberedInOrder(t *testing.T) {
	ls := numberLinks([]Link{{Kind: LinkGapOverlapsCap}, {Kind: LinkDeathInGap}})
	if ls[0].ID != "l1" || ls[1].ID != "l2" {
		t.Fatalf("%+v", ls)
	}
}
