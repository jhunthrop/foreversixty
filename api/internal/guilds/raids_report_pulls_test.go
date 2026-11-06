package guilds

import (
	"encoding/json"
	"net/http"
	"os"
	"slices"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// digestedNight is matchedNight with the fights verified and summaries in a fake bucket for
// the first boss pull (fight 1, Thornhide and Sunderfel die there, sheets in the log) and the
// kill (fight 4, no sheet). Fights 2 and 3 have no summary.
func digestedNight(t *testing.T) (raidWorld, Raid, *memObjects) {
	t.Helper()
	w, raid := matchedNight(t)
	if _, err := w.h.pool.Exec(t.Context(), `update fights set verified = true where report_id = 'night'`); err != nil {
		t.Fatal(err)
	}
	pull1 := deathSummary(true)
	pull1.Deaths = append(pull1.Deaths, summary.Death{Name: "Thornhide", Class: "Druid", AtMS: 30_000,
		KillingBlow: ptrRef(hit(29_900, 18500, "Wing Buffet", "Onyxia", 1500, 300)),
		Last:        []summary.DamageRef{hit(29_900, 18500, "Wing Buffet", "Onyxia", 1500, 300)}})
	kill := summary.Summary{Roster: pull1.Roster}
	objects := newMemObjects()
	for n, sum := range map[int]summary.Summary{1: pull1, 4: kill} {
		b, err := json.Marshal(sum)
		if err != nil {
			t.Fatal(err)
		}
		objects.objects[store.Keys{ReportID: "night"}.FightSummary(n)] = b
	}
	w.h.svc.R2 = objects
	return w, raid, objects
}

func pullAt(t *testing.T, rep *ReportView, index int) PullView {
	t.Helper()
	for _, p := range rep.Pulls {
		if p.FightIndex == index {
			return p
		}
	}
	t.Fatalf("no pull %d in %+v", index, rep.Pulls)
	return PullView{}
}

func TestTheReportCarriesTheHeadlinePullsAndDeathCauses(t *testing.T) {
	w, raid, _ := digestedNight(t)
	w.as(w.officer)
	rep := w.getRaid(t, raid.ID).Report
	if rep == nil {
		t.Fatal("no report")
	}
	if h := rep.Headline; h.Kills != 1 || h.Wipes != 3 || h.Attempts != 4 || h.Deaths != 12 {
		t.Fatalf("headline = %+v, want 1 kill 3 wipes 4 attempts 12 deaths", h)
	}
	if len(rep.Pulls) != 5 {
		t.Fatalf("pulls = %+v", rep.Pulls)
	}
	trash := pullAt(t, rep, 0)
	if trash.Attempt != nil || trash.Result != PullWipe || trash.EncounterID != nil {
		t.Fatalf("trash = %+v", trash)
	}
	first := pullAt(t, rep, 1)
	if first.Result != PullWipe || first.Attempt == nil || *first.Attempt != 1 || len(first.DeathCauses) != 2 {
		t.Fatalf("first pull = %+v", first)
	}
	var thorn PullDeathView
	for _, d := range first.DeathCauses {
		if d.Name == "Thornhide" {
			thorn = d
		}
	}
	if thorn.CharacterKey != officerKey || thorn.Class != "druid" || thorn.KillingBlow == nil ||
		thorn.KillingBlow.Ability != "Wing Buffet" || thorn.KillingBlow.Source != "Onyxia" || len(thorn.LastHits) != 1 {
		t.Fatalf("thornhide's death = %+v", thorn)
	}
	if kill := pullAt(t, rep, 4); kill.Result != PullKill || *kill.Attempt != 4 || kill.DeathCauses == nil {
		t.Fatalf("kill = %+v, want read (no deaths in its summary)", kill)
	}
	if p := pullAt(t, rep, 2); p.DeathCauses != nil {
		t.Fatalf("pull 2 has no summary, causes = %+v, want null", p.DeathCauses)
	}
}

func TestTheReportsConsumablesByPullSayWhatTheLogCanRead(t *testing.T) {
	w, raid, _ := digestedNight(t)
	w.as(w.officer)
	rep := w.getRaid(t, raid.ID).Report
	byPull := map[int]PullConsumablesView{}
	for _, bp := range rep.Consumables.ByPull {
		byPull[bp.FightIndex] = bp
	}
	if len(byPull) != 4 || !rep.Consumables.LogHasBuffs {
		t.Fatalf("by pull = %+v has buffs %v", rep.Consumables.ByPull, rep.Consumables.LogHasBuffs)
	}
	if bp := byPull[4]; bp.Rows != nil || bp.Reason == nil || *bp.Reason != ConsumablesNoSheet {
		t.Fatalf("kill (no sheet) = %+v", bp)
	}
	if bp := byPull[2]; bp.Reason == nil || *bp.Reason != ConsumablesNotRead {
		t.Fatalf("pull 2 (unread) = %+v", bp)
	}
	rows := map[string]PullConsumableRow{}
	for _, r := range byPull[1].Rows {
		rows[r.CharacterKey] = r
	}
	if r := rows[officerKey]; r.Flask == nil || !*r.Flask {
		t.Fatalf("thornhide at pull 1 = %+v", r)
	}
	if r := rows[memberKeyB]; r.Flask == nil || *r.Flask {
		t.Fatalf("sunderfel at pull 1 = %+v", r)
	}
}

func TestAMembersReportKeepsOnlyTheirOwnDeathsParsesAndConsumables(t *testing.T) {
	w, raid, _ := digestedNight(t)
	w.as(w.member)
	rep := w.getRaid(t, raid.ID).Report
	if rep.Headline.Deaths != 12 || len(rep.Pulls) != 5 {
		t.Fatalf("a member reads every pull and the headline: %+v", rep.Headline)
	}
	first := pullAt(t, rep, 1)
	if len(first.DeathCauses) != 1 || first.DeathCauses[0].CharacterKey != memberKeyB {
		t.Fatalf("member's causes = %+v, want only Sunderfel", first.DeathCauses)
	}
	for _, p := range rep.Parses {
		if p.CharacterKey != memberKeyB {
			t.Fatalf("member reads another raider's parse: %+v", p)
		}
	}
	for _, bp := range rep.Consumables.ByPull {
		for _, r := range bp.Rows {
			if r.CharacterKey != memberKeyB {
				t.Fatalf("member reads another raider's consumables: %+v", r)
			}
		}
	}
}

func TestTheReportsParsesArePerRaiderPerKill(t *testing.T) {
	w, raid, _ := digestedNight(t)
	if _, err := w.h.pool.Exec(t.Context(), `update fight_metrics set metric_dps = 300 where player_key = $1`, memberKeyB); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	rep := w.getRaid(t, raid.ID).Report
	if len(rep.Parses) != 5 {
		t.Fatalf("parses = %+v, want the five raiders of the kill", rep.Parses)
	}
	if p := rep.Parses[0]; p.CharacterKey != memberKeyB || p.Value != 300 || p.Metric != "dps" || p.FightIndex != 4 || p.Rating != nil {
		t.Fatalf("best parse = %+v", p)
	}
}

func TestADigestIsReadOnce(t *testing.T) {
	w, raid, objects := digestedNight(t)
	w.as(w.officer)
	w.getRaid(t, raid.ID)
	delete(objects.objects, store.Keys{ReportID: "night"}.FightSummary(1))
	rep := w.getRaid(t, raid.ID).Report
	if p := pullAt(t, rep, 1); len(p.DeathCauses) != 2 {
		t.Fatalf("the stored digest was not reused: %+v", p)
	}
	if res := w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("get = %d", res.StatusCode)
	}
}

// The report's JSON carries exactly the fields openapi.yaml's RaidReport requires, at the top and
// on each pull, death and hit.
func TestTheReportMatchesItsOpenAPISchema(t *testing.T) {
	raw, err := os.ReadFile("../../openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Components struct {
			Schemas map[string]struct {
				Required   []string                  `yaml:"required"`
				Properties map[string]map[string]any `yaml:"properties"`
			} `yaml:"schemas"`
		} `yaml:"components"`
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	w, raid, _ := digestedNight(t)
	w.as(w.officer)
	data := rawData(t, w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), ""))
	var rep map[string]json.RawMessage
	if err := json.Unmarshal(data["report"], &rep); err != nil {
		t.Fatal(err)
	}
	if got, want := sortedJSONKeys(rep), sorted(doc.Components.Schemas["RaidReport"].Required); !slices.Equal(got, want) {
		t.Fatalf("report keys = %v, schema requires %v", got, want)
	}
	var pulls []map[string]json.RawMessage
	if err := json.Unmarshal(rep["pulls"], &pulls); err != nil {
		t.Fatal(err)
	}
	want := []string{"attempt", "death_causes", "deaths", "duration_ms", "encounter_id", "fight_index", "name", "result"}
	if got := sortedJSONKeys(pulls[1]); !slices.Equal(got, want) {
		t.Fatalf("pull keys = %v", got)
	}
	hit := doc.Components.Schemas["RaidReportHit"].Required
	var causes []struct {
		KillingBlow map[string]json.RawMessage `json:"killing_blow"`
	}
	if err := json.Unmarshal(pulls[1]["death_causes"], &causes); err != nil {
		t.Fatal(err)
	}
	if got := sortedJSONKeys(causes[0].KillingBlow); !slices.Equal(got, sorted(hit)) {
		t.Fatalf("hit keys = %v, schema requires %v", got, hit)
	}
}

func sortedJSONKeys(m map[string]json.RawMessage) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return sorted(out)
}

func sorted(s []string) []string {
	out := append([]string{}, s...)
	slices.Sort(out)
	return out
}
