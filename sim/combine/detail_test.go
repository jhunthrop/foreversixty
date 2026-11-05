package combine

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestHistogramsSumAcrossParts(t *testing.T) {
	got := poolHistograms([]api.SimResult{
		{DPSHistogram: &api.DPSHistogram{Start: 900, BinDPS: 10, Counts: []int{1, 2}}},
		{DPSHistogram: &api.DPSHistogram{Start: 890, BinDPS: 10, Counts: []int{3, 4, 5, 6}}},
	})
	want := []int{3, 5, 7, 6}
	if got.Start != 890 || got.BinDPS != 10 || len(got.Counts) != len(want) {
		t.Fatalf("got %+v", got)
	}
	for i := range want {
		if got.Counts[i] != want[i] {
			t.Errorf("counts = %v, want %v", got.Counts, want)
		}
	}
}

func TestAPartWithoutAHistogramLeavesNone(t *testing.T) {
	got := poolHistograms([]api.SimResult{{DPSHistogram: &api.DPSHistogram{Start: 1, BinDPS: 1, Counts: []int{1}}}, {}})
	if got != nil {
		t.Errorf("got %+v, want none", got)
	}
}

func TestResourceSourcesAreIterationWeighted(t *testing.T) {
	got := poolResourceSources([]api.SimResult{
		{IterationsRun: 3000, ResourceSources: []api.ResourceSources{{
			Resource: "rage",
			Gained:   []api.ResourceSourceRow{{Action: "spell:1", Amount: 100, Wasted: 10}},
			Spent:    []api.ResourceSourceRow{{Action: "spell:2", Amount: 50}},
		}}},
		{IterationsRun: 1000, ResourceSources: []api.ResourceSources{{
			Resource: "rage",
			Gained:   []api.ResourceSourceRow{{Action: "spell:1", Amount: 200, Wasted: 20}},
			Spent:    []api.ResourceSourceRow{{Action: "spell:2", Amount: 90}},
		}}},
	}, 4000)
	if len(got) != 1 || got[0].Gained[0] != (api.ResourceSourceRow{Action: "spell:1", Amount: 125, Wasted: 12.5}) ||
		got[0].Spent[0] != (api.ResourceSourceRow{Action: "spell:2", Amount: 60}) {
		t.Errorf("got %+v", got)
	}
}

func TestExtraAttacksAreIterationWeighted(t *testing.T) {
	got := poolExtraAttacks([]api.SimResult{
		{IterationsRun: 1000, ExtraAttackSources: []api.ExtraAttackSource{{Source: "windfury_totem", Count: 20}}},
		{IterationsRun: 1000, ExtraAttackSources: []api.ExtraAttackSource{{Source: "windfury_totem", Count: 30}}},
	}, 2000)
	if len(got) != 1 || got[0].Count != 25 {
		t.Errorf("got %+v", got)
	}
}
