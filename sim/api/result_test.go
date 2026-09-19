package api

import (
	"encoding/json"
	"testing"
)

// The web mirrors these tags verbatim, so the JSON names are pinned
// here rather than left to whoever edits the struct next. A plain run's
// result must not grow a single key: every new field is omitempty, and
// a run that filled none of them marshals exactly as it did before.
func TestAPlainResultGainsNoKeys(t *testing.T) {
	b, err := json.Marshal(SimResult{})
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"combos", "equipped", "stages", "weights", "sample"} {
		if _, ok := got[key]; ok {
			t.Errorf("a plain result carries %q; every bulk field is omitempty", key)
		}
	}
}

func TestBulkResultJSONNames(t *testing.T) {
	res := SimResult{
		Combos: []Combo{{
			Substitutions: []Substitution{
				{Kind: SubstitutionItem, Slot: "main_hand", ItemID: 19352, Enchant: 2568, Origin: OriginBag,
					Name: "Vis'kag the Bloodletter", SourceName: "Lucifron"},
				{Kind: SubstitutionTalents, Name: "Deep Fury", Talents: "30305001302-05050005525010051"},
				{Kind: SubstitutionSet, Name: "my AQ set"},
				{Kind: SubstitutionConsumes, Name: "flask_of_the_titans, juju_power",
					Consumes: []string{"flask_of_the_titans", "juju_power"}},
			},
			DPS:   Estimate{Mean: 1100},
			Delta: Estimate{Mean: 41, Error: 9},
			Group: 0,
		}},
		Equipped: &Estimate{Mean: 1059},
		Stages:   []Stage{{Iterations: 100, Combos: 38}, {Iterations: 1000, Combos: 10}},
		Weights:  []StatWeight{{Stat: "crit", Weight: 1, Error: 0.04}},
		Sample:   []SampleCast{{AtMS: -1500, Action: "spell:1719", Resources: map[string]int{"rage": 0}}},
	}
	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"combos", "equipped", "stages", "weights", "sample"} {
		if _, ok := got[key]; !ok {
			t.Errorf("a bulk result does not carry %q", key)
		}
	}

	var back SimResult
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if len(back.Combos) != 1 || back.Combos[0].Delta.Mean != 41 ||
		back.Combos[0].Substitutions[0].ItemID != 19352 ||
		back.Combos[0].Substitutions[0].Name != "Vis'kag the Bloodletter" ||
		back.Combos[0].Substitutions[0].SourceName != "Lucifron" {
		t.Errorf("round trip lost something: %+v", back.Combos)
	}
	if back.Sample[0].AtMS != -1500 || back.Sample[0].Action != "spell:1719" || back.Sample[0].Resources["rage"] != 0 {
		t.Errorf("round trip lost the sample: %+v", back.Sample)
	}
}

// The progress payload is the page's partial result, so the three bulk
// fields are absent for a plain run rather than zero-valued keys the
// page has to ignore.
func TestProgressGainsTheStageFields(t *testing.T) {
	b, err := json.Marshal(Progress{IterationsRun: 100, DPS: Estimate{Mean: 900}})
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"iterations_run":100,"dps":{"mean":900,"stddev":0,"error":0,"min":0,"max":0}}` {
		t.Errorf("a plain run's progress is %s", b)
	}
	b, err = json.Marshal(Progress{IterationsRun: 100, Stage: 2, CombosDone: 31, CombosTotal: 96})
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"stage", "combos_done", "combos_total"} {
		if _, ok := got[key]; !ok {
			t.Errorf("a bulk run's progress does not carry %q", key)
		}
	}
}
