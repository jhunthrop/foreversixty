package guilds

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// prepRow is a 60 at bags consent whose export reads code's raid-prep sections.
func prepRow(t *testing.T, consent, sections string) RosterRow {
	t.Helper()
	d, ok := fs1.Decode("FS1:1.60.1.70790:warrior:human:0/0/0:head=16866|level=60" + sections)
	if !ok {
		t.Fatal("decode")
	}
	return RosterRow{Consent: consent, gear: d.Gear, bags: d.Bags, hasBagsSection: d.HasBags, prep: prepOf(d)}
}

func TestReadSuppliesNeedsFlaskOrElixirAndFood(t *testing.T) {
	cases := map[string]struct {
		supplies map[int]int
		missing  []string
	}{
		"flask and food":  {map[int]int{13510: 1, 20452: 6}, []string{}},
		"elixir and food": {map[int]int{13452: 3, 287974: 10}, []string{}},
		"flask, no food":  {map[int]int{13512: 2}, []string{missingFood}},
		"food, no flask":  {map[int]int{21023: 5}, []string{missingFlaskOrElixir}},
		"only uncounted":  {map[int]int{118: 20, 13506: 1}, []string{missingFlaskOrElixir, missingFood}},
	}
	for name, tc := range cases {
		if got := readSupplies(tc.supplies); !reflect.DeepEqual(got.Missing, tc.missing) {
			t.Errorf("%s: missing %v, want %v", name, got.Missing, tc.missing)
		}
	}
	got := readSupplies(map[int]int{20452: 6, 13452: 3, 13510: 1, 118: 4})
	want := []SupplyStack{
		{ID: 13510, Name: "Flask of the Titans", Kind: supplyFlask, Count: 1},
		{ID: 13452, Name: "Elixir of the Mongoose", Kind: supplyElixir, Count: 3},
		{ID: 20452, Name: "Smoked Desert Dumplings", Kind: supplyFood, Count: 6},
	}
	if !reflect.DeepEqual(got.Stacks, want) {
		t.Fatalf("stacks = %+v", got.Stacks)
	}
}

func TestConsumablesStateNamesSuppliesAndKeepsTheOldBagsRule(t *testing.T) {
	cases := map[string]struct {
		consent string
		d       fs1.Decoded
		want    string
	}{
		"no bags share":          {"gear", fs1.Decoded{HasSupplies: true, Supplies: map[int]int{13510: 1, 20452: 1}}, consumablesUnknown},
		"supplies ready":         {"gear_bags", fs1.Decoded{HasSupplies: true, Supplies: map[int]int{13510: 1, 20452: 1}}, consumablesStocked},
		"supplies short":         {"gear_bags", fs1.Decoded{HasSupplies: true, Supplies: map[int]int{20452: 1}, HasBags: true, Bags: []int{1}}, consumablesShort},
		"old export, bags":       {"gear_bags", fs1.Decoded{HasBags: true, Bags: []int{1}}, consumablesStocked},
		"old export, empty bags": {"gear_bags", fs1.Decoded{HasBags: true}, consumablesShort},
		"old export, no bags":    {"gear_bags", fs1.Decoded{}, consumablesUnknown},
	}
	for name, tc := range cases {
		if got := consumablesState(tc.consent, tc.d); got != tc.want {
			t.Errorf("%s: %s, want %s", name, got, tc.want)
		}
	}
}

func TestKeyAttunementFailsOnlyOnceTheKeysAreListed(t *testing.T) {
	scholo, _ := ZoneByKey("dungeon:scholomance")
	cases := map[string]struct {
		row  RosterRow
		want attunementReading
	}{
		"old export":         {prepRow(t, "gear_bags", "|bags=1"), attunementReading{State: CheckNotTracked, Reason: reasonKeysNotInExport}},
		"key carried":        {prepRow(t, "gear_bags", "|locks=13704"), attunementReading{State: CheckOK}},
		"keys listed, none":  {prepRow(t, "gear_bags", "|locks=12382"), attunementReading{State: CheckFail}},
		"no bags share":      {prepRow(t, "gear", "|locks=13704"), attunementReading{State: CheckNotTracked, Reason: reasonKeysNotInExport}},
		"empty keys listing": {prepRow(t, "gear_bags", "|locks="), attunementReading{State: CheckFail}},
	}
	for name, tc := range cases {
		if got := attunementOf(scholo, tc.row); got != tc.want {
			t.Errorf("%s: %+v, want %+v", name, got, tc.want)
		}
	}
	ubrs, _ := ZoneByKey("dungeon:upper-blackrock-spire")
	if got := attunementOf(ubrs, prepRow(t, "gear_bags", "|bank=12344|locks=")); got.State != CheckOK {
		t.Fatalf("the seal in the bank = %+v", got)
	}
}

func TestOnyxiaAttunementReadsTheQuestFlags(t *testing.T) {
	ony, _ := ZoneByKey(raidZoneOny)
	cases := map[string]struct {
		row  RosterRow
		want attunementReading
	}{
		"old export":       {prepRow(t, "gear_bags", ""), attunementReading{State: CheckNotTracked, Reason: reasonQuestsNotInExport}},
		"alliance done":    {prepRow(t, "gear", "|quests=6502:1,6602:0"), attunementReading{State: CheckOK}},
		"horde done":       {prepRow(t, "gear", "|quests=6502:0,6602:1"), attunementReading{State: CheckOK}},
		"asked, not done":  {prepRow(t, "gear", "|quests=6502:0,6602:0"), attunementReading{State: CheckFail}},
		"asked only one":   {prepRow(t, "gear", "|quests=6502:0"), attunementReading{State: CheckNotTracked, Reason: reasonQuestsNotInExport}},
		"amulet in bags":   {prepRow(t, "gear_bags", "|bags=16309|quests=6502:0,6602:0"), attunementReading{State: CheckOK}},
		"roster consent":   {prepRow(t, "roster", "|quests=6502:1,6602:0"), attunementReading{State: CheckNotTracked, Reason: reasonGearNotShared}},
		"amulet, no share": {prepRow(t, "gear", "|bags=16309|quests=6502:0,6602:0"), attunementReading{State: CheckFail}},
		"amulet held":      {prepRow(t, "gear", "|quests=6502:0,6602:0|owned=16309:1"), attunementReading{State: CheckOK}},
		"amulet not held":  {prepRow(t, "gear", "|quests=6502:0,6602:0|owned=16309:0"), attunementReading{State: CheckFail}},
		"amulet only":      {prepRow(t, "gear", "|owned=16309:1"), attunementReading{State: CheckOK}},
	}
	for name, tc := range cases {
		if got := attunementOf(ony, tc.row); got != tc.want {
			t.Errorf("%s: %+v, want %+v", name, got, tc.want)
		}
	}
}

func TestChecklistFlaskFoodAndFireResistance(t *testing.T) {
	sixty := 60
	ony, _ := ZoneByKey(raidZoneOny)
	ready := CharacterReadiness{Level: &sixty, GearChecked: true, EnchantsChecked: true}
	row := prepRow(t, "gear_bags", "|school=85:0:0:0:0|supplies=13510:1,20452:6|quests=6502:1,6602:0")
	items, verdict := buildChecklist(checklistInput{zone: &ony, readiness: ready, row: row, onRoster: true})
	if verdict != VerdictReady {
		t.Fatalf("verdict %s: %+v", verdict, items)
	}
	ff := itemOf(items, CheckFlaskFood)
	if ff.State != CheckOK || len(ff.Supplies) != 2 || ff.Supplies[0].Name != "Flask of the Titans" {
		t.Fatalf("flask_food = %+v", ff)
	}
	if fr := itemOf(items, CheckFireResistance); fr.State != CheckOK || *fr.Resistance != 85 {
		t.Fatalf("fire_resistance = %+v", fr)
	}
	if a := itemOf(items, CheckAttunement); a.State != CheckOK || a.AttunementKind != attunementQuest || a.KeyItem.ID != 16309 {
		t.Fatalf("attunement = %+v", a)
	}

	short := prepRow(t, "gear_bags", "|supplies=13510:1|quests=6502:0,6602:0")
	items, verdict = buildChecklist(checklistInput{zone: &ony, readiness: ready, row: short, onRoster: true})
	if ff := itemOf(items, CheckFlaskFood); ff.State != CheckFail || fmt.Sprint(ff.Missing) != "[food]" {
		t.Fatalf("flask, no food = %+v", ff)
	}
	if verdict != VerdictGettingStarted {
		t.Fatalf("not attuned reads %s", verdict)
	}
	if fr := itemOf(items, CheckFireResistance); fr.State != CheckNotTracked || fr.Reason != reasonNeedsResistance {
		t.Fatalf("no school section = %+v", fr)
	}

	gearOnly := prepRow(t, "gear", "|school=0:0:0:0:0|supplies=13510:1,20452:6")
	items, _ = buildChecklist(checklistInput{zone: &ony, readiness: ready, row: gearOnly, onRoster: true})
	if ff := itemOf(items, CheckFlaskFood); ff.State != CheckNotTracked || ff.Reason != reasonBagsNotShared {
		t.Fatalf("bags not shared = %+v", ff)
	}
	if fr := itemOf(items, CheckFireResistance); fr.State != CheckOK || *fr.Resistance != 0 {
		t.Fatalf("zero fire resistance = %+v", fr)
	}

	bwl := Zone{Key: "raid:barrow-deeps", Kind: ZoneRaid}
	items, _ = buildChecklist(checklistInput{zone: &bwl, readiness: ready, row: row, onRoster: true})
	if it := itemOf(items, CheckFireResistance); it.Key != "" {
		t.Fatalf("a zone with no fire asks for it: %+v", it)
	}
}

func TestOfficerReadinessCarriesTheRaidPrep(t *testing.T) {
	w := newRaidWorld(t, true)
	run := w.createRaid(t)
	export := "FS1:1.60.1.70790:priest:human:0/0/0:|level=60|school=40:5:0:0:0|supplies=20452:6|locks=|quests=6502:0,6602:0"
	if _, err := w.h.pool.Exec(t.Context(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', 'Mirelle', $3, now(), now())`, member2Key, w.member2, export); err != nil {
		t.Fatal(err)
	}
	if _, err := w.h.pool.Exec(t.Context(), `update guild_members set consent = 'gear_bags' where user_id = $1`, w.member2); err != nil {
		t.Fatal(err)
	}
	if res := w.signupAs(t, w.member2, run.ID, member2Key); res != http.StatusOK {
		t.Fatalf("signup = %d", res)
	}
	read := func(uid int64) ReadinessRow {
		w.as(uid)
		var v ReadinessView
		w.h.data(w.h.do(http.MethodGet, w.path("/readiness?raid=%d", run.ID), ""), &v)
		for _, r := range v.Rows {
			if r.CharacterKey == member2Key {
				return r
			}
		}
		t.Fatalf("no row for %s", member2Key)
		return ReadinessRow{}
	}
	got := read(w.officer)
	if got.Attuned == nil || *got.Attuned {
		t.Fatalf("officer attuned = %v, want false", got.Attuned)
	}
	if got.Resistances == nil || got.Resistances.Fire != 40 || got.Resistances.Nature != 5 {
		t.Fatalf("resistances = %+v", got.Resistances)
	}
	if got.FlaskFood == nil || got.FlaskFood.FlaskOrElixir || !got.FlaskFood.Food {
		t.Fatalf("flask_food = %+v", got.FlaskFood)
	}
	if got.Consumables.State != consumablesShort {
		t.Fatalf("consumables = %s, want short (food, no flask)", got.Consumables.State)
	}
	member := read(w.member)
	if member.Attuned != nil || member.Resistances != nil || member.FlaskFood != nil {
		t.Fatalf("a member reads the officer fields: %+v", member)
	}
}
