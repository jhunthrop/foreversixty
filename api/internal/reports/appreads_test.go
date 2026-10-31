package reports

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// fakePlacer answers every bracket with one placement and every character with one best.
type fakePlacer struct {
	pct      float64
	ranked   int64
	placed   bool
	best     float64
	hasBest  bool
	asked    []string // the spec and metric of each percentile asked
	excluded []string // the report each best-before left out
}

func (f *fakePlacer) Percentile(_ context.Context, _, _ int64, spec, _, metric string, _ float64) (float64, int64, bool, error) {
	f.asked = append(f.asked, spec+"/"+metric)
	return f.pct, f.ranked, f.placed, nil
}

func (f *fakePlacer) BestBefore(_ context.Context, _ string, _, _ int64, _ string, _ time.Time, except string) (float64, bool, error) {
	f.excluded = append(f.excluded, except)
	return f.best, f.hasBest, nil
}

var meFightStart = time.Date(2026, 12, 15, 2, 36, 40, 0, time.UTC)

// appFight writes one fight row as the ingest does.
func (h *harness) appFight(index int, encounterID int64, name string, kill bool, durationMS int64) {
	h.t.Helper()
	start, end := int64(index*1000), int64(index*1000+999)
	rec := FightRecord{
		ReportID: h.reportID, Index: index, Name: name, Kill: kill, DurationMS: durationMS,
		StartMS: meFightStart.Add(time.Duration(index) * time.Minute).UnixMilli(), Verified: true,
		Players: []string{"Player-1", "Player-2", "Player-3"}, RawStart: &start, RawEnd: &end,
	}
	if encounterID != 0 {
		d, s := int64(8), int64(40)
		rec.EncounterID, rec.Difficulty, rec.Size = &encounterID, &d, &s
	}
	if _, err := h.store.UpsertFight(context.Background(), rec); err != nil {
		h.t.Fatal(err)
	}
}

// appSummary stores one fight's summary with the given roster.
func (h *harness) appSummary(index int, durationMS int64, roster ...summary.RosterRow) {
	h.t.Helper()
	body, err := json.Marshal(summary.Summary{FightIndex: index, DurationMS: durationMS, Roster: roster})
	if err != nil {
		h.t.Fatal(err)
	}
	if err := h.files.Put(context.Background(), Keys(h.reportID).FightSummary(index), body, store.PutOptions{}); err != nil {
		h.t.Fatal(err)
	}
}

// appReportJSON stores report.json with the given fights.
func (h *harness) appReportJSON(fights ...store.FightEntry) {
	h.t.Helper()
	body, err := json.Marshal(store.Report{ReportID: h.reportID, Fights: fights})
	if err != nil {
		h.t.Fatal(err)
	}
	if err := h.files.Put(context.Background(), Keys(h.reportID).Report(), body, store.PutOptions{}); err != nil {
		h.t.Fatal(err)
	}
}

func (h *harness) linkCharacter(name, class string) {
	h.t.Helper()
	c := auth.Character{Key: "us/hardcore/" + name, Region: "us", Ruleset: "hardcore", Name: titleName(name), Class: class}
	if err := (&auth.Store{Pool: h.store.Pool}).LinkCharacter(context.Background(), h.owner, c); err != nil {
		h.t.Fatal(err)
	}
}

func titleName(slug string) string { return string(slug[0]-'a'+'A') + slug[1:] }

func roster(name, class, spec, role string, damage, healing, taken int64, deaths int) summary.RosterRow {
	return summary.RosterRow{
		GUID: "Player-" + name, Name: titleName(name), Class: class, Spec: spec, Role: role,
		ActiveMS: 150_000, DamageDone: damage, HealingDone: healing, DamageTaken: taken, Deaths: deaths,
	}
}

// onyxiaNight is a report with an Onyxia wipe, an Onyxia kill (fight 3, 192 s) and trash, and
// the caller's Baelgrim in the kill.
func (h *harness) onyxiaNight(visibility string) (string, *fakePlacer) {
	h.t.Helper()
	h.reportID = h.createReport(visibility)
	h.appFight(1, 0, "Trash", false, 40_000)
	h.appFight(2, 1084, "Onyxia", false, 72_000)
	h.appFight(3, 1084, "Onyxia", true, 192_000)
	h.appSummary(3, 192_000,
		roster("baelgrim", "Warrior", "Fury", "dps", 79_104, 0, 40_000, 1), // 412 DPS
		roster("sunwick", "Priest", "Holy", "healer", 0, 96_000, 2_000, 0), // 500 HPS
		roster("morrowlyn", "Mage", "Fire", "dps", 96_000, 0, 1_000, 0))    // 500 DPS
	h.linkCharacter("baelgrim", "warrior")
	placer := &fakePlacer{pct: 77.6, ranked: 212, placed: true, best: 401.5, hasBest: true}
	h.service.Objects = dirGetter{root: h.dir}
	h.service.Places = placer
	return h.reportID, placer
}

func decodeData[T any](t *testing.T, raw []byte) T {
	t.Helper()
	var env struct {
		Data T `json:"data"`
	}
	if err := json.Unmarshal(raw, &env); err != nil {
		t.Fatal(err)
	}
	return env.Data
}

func TestMe_YourKillsWithTheParseAndANewBest(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, placer := h.onyxiaNight(Public)

	got := decodeData[MeView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil), http.StatusOK))
	if len(got.Characters) != 1 {
		t.Fatalf("characters = %+v, want Baelgrim alone", got.Characters)
	}
	c := got.Characters[0]
	if c.Key != "us/hardcore/baelgrim" || c.Class != "warrior" || c.Spec == nil || *c.Spec != "warrior-fury" || c.SpecName != "Fury" {
		t.Errorf("character = %+v", c)
	}
	if len(c.Fights) != 1 {
		t.Fatalf("fights = %+v, want the kill only (parses count kills)", c.Fights)
	}
	f := c.Fights[0]
	if f.Index != 3 || f.Encounter != "Onyxia" || !f.Kill || f.DurationMS != 192_000 || f.Unit != "dps" || f.Deaths != 1 {
		t.Errorf("fight = %+v", f)
	}
	if f.Value != 412 {
		t.Errorf("value = %v, want 412 (79,104 over 192 s)", f.Value)
	}
	if f.Percentile == nil || *f.Percentile != 78 || f.Ranked == nil || *f.Ranked != 212 {
		t.Errorf("percentile = %v of %v, want 78 (77.6 rounded) of 212", f.Percentile, f.Ranked)
	}
	if !f.NewBest {
		t.Error("412 over an earlier best of 401.5 is a new best")
	}
	if len(placer.asked) != 1 || placer.asked[0] != "Fury/dps" {
		t.Errorf("asked = %v, want the row's spec on its role metric", placer.asked)
	}
	if len(placer.excluded) != 1 || placer.excluded[0] != id {
		t.Errorf("best-before left out %v, want this report", placer.excluded)
	}
}

func TestMe_AgreesWithTheMeters(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, _ := h.onyxiaNight(Public)

	me := decodeData[MeView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil), http.StatusOK))
	meters := decodeData[MetersView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters", "", nil), http.StatusOK))
	for _, p := range meters.Players {
		if p.Mine && p.DPS != me.Characters[0].Fights[0].Value {
			t.Fatalf("meters say %v DPS, Your parses %v: one figure, two numbers", p.DPS, me.Characters[0].Fights[0].Value)
		}
	}
}

func TestMe_NoParseWithoutARankedBracketAndNoBestOnAFirstKill(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, placer := h.onyxiaNight(Public)
	placer.placed, placer.hasBest = false, false

	got := decodeData[MeView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil), http.StatusOK))
	f := got.Characters[0].Fights[0]
	if f.Percentile != nil || f.Ranked != nil || f.NewBest {
		t.Errorf("fight = %+v, want no parse and no new best", f)
	}
}

func TestMe_NoCharactersOfYoursIsEmpty(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, _ := h.onyxiaNight(Public)
	if _, err := h.store.Pool.Exec(context.Background(), `update characters set user_id = null`); err != nil {
		t.Fatal(err)
	}
	got := decodeData[MeView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil), http.StatusOK))
	if got.ReportID != id || got.Characters == nil || len(got.Characters) != 0 {
		t.Errorf("got %+v, want an empty list", got)
	}
}

func TestMe_SomeoneElsesPrivateReportIs404(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, _ := h.onyxiaNight(Private)
	stranger, err := (&auth.Store{Pool: h.store.Pool}).UpsertEmailUser(context.Background(), "stranger@example.com")
	if err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: stranger.ID, Role: "user", Method: "session"}
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil), http.StatusNotFound)
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters", "", nil), http.StatusNotFound)
}

func TestMe_NeedsTheAccountHolder(t *testing.T) {
	h := newHarness(t)
	id, _ := h.onyxiaNight(Public)
	h.actor = auth.Actor{}
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/me", "", nil), http.StatusUnauthorized)
}

func TestMeters_DamageSortedWithYourRowAndPages(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	id, _ := h.onyxiaNight(Public)

	got := decodeData[MetersView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters", "", nil), http.StatusOK))
	if got.FightIndex != 3 || !got.Kill || got.DurationMS != 192_000 || got.Metric != "damage" {
		t.Errorf("head = %+v", got)
	}
	if len(got.Players) != 2 || got.Players[0].Name != "Morrowlyn" || got.Players[1].Name != "Baelgrim" {
		t.Fatalf("players = %+v, want Morrowlyn then Baelgrim, the healer left off", got.Players)
	}
	b := got.Players[1]
	if !b.Mine || b.DPS != 412 || b.Key == nil || *b.Key != "us/hardcore/baelgrim" || b.Spec == nil || *b.Spec != "warrior-fury" {
		t.Errorf("Baelgrim = %+v", b)
	}
	if m := got.Players[0]; m.Mine || m.Key != nil {
		t.Errorf("Morrowlyn = %+v, want not mine and no page", m)
	}
}

func TestMeters_HealingAndTaken(t *testing.T) {
	h := newHarness(t)
	id, _ := h.onyxiaNight(Public)
	h.actor = auth.Actor{}

	heal := decodeData[MetersView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters?metric=healing", "", nil), http.StatusOK))
	if len(heal.Players) != 1 || heal.Players[0].Name != "Sunwick" || heal.Players[0].HPS != 500 {
		t.Errorf("healing = %+v", heal.Players)
	}
	taken := decodeData[MetersView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters?metric=taken", "", nil), http.StatusOK))
	if len(taken.Players) != 3 || taken.Players[0].Name != "Baelgrim" || taken.Players[2].Name != "Morrowlyn" {
		t.Errorf("taken = %+v", taken.Players)
	}
	if taken.Players[0].Key != nil && taken.Players[0].Mine {
		t.Error("a signed-out reader has no rows of their own")
	}
}

func TestMeters_TrashFoldsTheZone(t *testing.T) {
	h := newHarness(t)
	id, _ := h.onyxiaNight(Public)
	h.appFight(4, 0, "Trash", false, 60_000)
	h.appFight(5, 0, "Trash", false, 20_000)
	h.appSummary(1, 40_000, roster("baelgrim", "Warrior", "Fury", "dps", 20_000, 0, 0, 0))
	h.appSummary(5, 20_000, roster("baelgrim", "Warrior", "Fury", "dps", 10_000, 0, 0, 1))
	h.appReportJSON(
		store.FightEntry{Index: 1, Kind: "trash", Name: "Trash", Zone: "Onyxia's Lair", DurationMS: 40_000},
		store.FightEntry{Index: 2, Kind: "encounter", Name: "Onyxia", Zone: "Onyxia's Lair", DurationMS: 72_000},
		store.FightEntry{Index: 3, Kind: "encounter", Name: "Onyxia", Zone: "Onyxia's Lair", DurationMS: 192_000},
		store.FightEntry{Index: 4, Kind: "trash", Name: "Trash", Zone: "Molten Core", DurationMS: 60_000},
		store.FightEntry{Index: 5, Kind: "trash", Name: "Trash", Zone: "Onyxia's Lair", DurationMS: 20_000},
	)

	got := decodeData[MetersView](t, reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/1/meters?scope=trash", "", nil), http.StatusOK))
	if got.Scope != "trash" || got.Kill || got.DurationMS != 60_000 {
		t.Errorf("head = %+v, want the zone's two pulls, 60 s", got)
	}
	if len(got.Players) != 1 || got.Players[0].DPS != 500 || got.Players[0].Deaths != 1 {
		t.Errorf("players = %+v, want 30,000 over 60 s and the one death", got.Players)
	}
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/2/meters?scope=trash", "", nil), http.StatusNotFound)

	var view View
	h.data(h.do(http.MethodGet, "/v1/reports/"+id+"?fold=trash", "", nil), &view)
	if len(view.TrashGroups) != 2 || view.TrashGroups[0].Key != 1 || len(view.TrashGroups[0].Fights) != 2 ||
		view.TrashGroups[1].Zone != "Molten Core" {
		t.Errorf("trash groups = %+v", view.TrashGroups)
	}
	if raw := reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id, "", nil), http.StatusOK); json.Valid(raw) && containsKey(raw, "trash_groups") {
		t.Error("trash_groups is sent only when asked for")
	}
}

func containsKey(raw []byte, key string) bool {
	var env struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	_ = json.Unmarshal(raw, &env)
	_, ok := env.Data[key]
	return ok
}

func TestMeters_Refusals(t *testing.T) {
	h := newHarness(t)
	id, _ := h.onyxiaNight(Public)
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/0/meters", "", nil), http.StatusBadRequest)
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters?metric=threat", "", nil), http.StatusBadRequest)
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters?scope=night", "", nil), http.StatusBadRequest)
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/9/meters", "", nil), http.StatusNotFound)
	reportBody(t, h.do(http.MethodGet, "/v1/reports/nosuchreport/fights/3/meters", "", nil), http.StatusNotFound)

	if err := h.store.Flag(context.Background(), id, "metrics_mismatch"); err != nil {
		t.Fatal(err)
	}
	raw := reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters", "", nil), http.StatusNotFound)
	var env struct {
		Error struct{ Code string } `json:"error"`
	}
	if err := json.Unmarshal(raw, &env); err != nil || env.Error.Code != "flagged" {
		t.Errorf("flagged report answered %s", raw)
	}
}

func TestMeters_NoBucketIs503(t *testing.T) {
	h := newHarness(t)
	id, _ := h.onyxiaNight(Public)
	h.service.Objects = nil
	reportBody(t, h.do(http.MethodGet, "/v1/reports/"+id+"/fights/3/meters", "", nil), http.StatusServiceUnavailable)
}

func TestTrashGroups_ZoneByZoneInTheNightsOrder(t *testing.T) {
	got := TrashGroups([]store.FightEntry{
		{Index: 1, Kind: "trash", Zone: "Molten Core", DurationMS: 10},
		{Index: 2, Kind: "encounter", Zone: "Molten Core", DurationMS: 99},
		{Index: 3, Kind: "trash", DurationMS: 5},
		{Index: 4, Kind: "trash", Zone: "Onyxia's Lair", DurationMS: 20},
		{Index: 5, Kind: "trash", Zone: "Molten Core", DurationMS: 30},
	})
	if len(got) != 2 {
		t.Fatalf("groups = %+v, want two zones; zone-less trash keeps its own row", got)
	}
	if got[0].Key != 1 || got[0].Zone != "Molten Core" || got[0].DurationMS != 40 || len(got[0].Fights) != 2 {
		t.Errorf("first = %+v", got[0])
	}
	if got[1].Key != 4 || got[1].Fights[0] != 4 {
		t.Errorf("second = %+v", got[1])
	}
}

func TestFigure_RoundsToHundredths(t *testing.T) {
	for _, c := range []struct {
		total, ms int64
		want      float64
	}{{79_104, 192_000, 412}, {1_000, 3_000, 333.33}, {5, 0, 0}} {
		if got := Figure(c.total, c.ms); got != c.want {
			t.Errorf("Figure(%d, %d) = %v, want %v", c.total, c.ms, got, c.want)
		}
	}
}
