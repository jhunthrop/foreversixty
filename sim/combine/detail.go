package combine

import (
	"cmp"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
)

// poolHistograms adds the parts' DPS histograms: the counts of a sharded run
// sum to the whole run's iterations, on the widest bin any part used. A part
// without one (a result from before the field) leaves the run without one,
// since a spread missing a shard would read as a narrower fight.
func poolHistograms(parts []api.SimResult) *api.DPSHistogram {
	bin, start, last := 0, 0, 0
	for i, p := range parts {
		h := p.DPSHistogram
		if h == nil || len(h.Counts) == 0 {
			return nil
		}
		end := h.Start + (len(h.Counts)-1)*h.BinDPS
		if i == 0 {
			start, last = h.Start, end
		}
		bin, start, last = max(bin, h.BinDPS), min(start, h.Start), max(last, end)
	}
	counts := make([]int, (last-start)/bin+1)
	for _, p := range parts {
		for i, n := range p.DPSHistogram.Counts {
			counts[(p.DPSHistogram.Start+i*p.DPSHistogram.BinDPS-start)/bin] += n
		}
	}
	return &api.DPSHistogram{Start: start, BinDPS: bin, Counts: counts}
}

// poolResourceSources is the iteration-weighted average of the parts' rows,
// the same rule weightSummaries applies to every other per-fight figure.
func poolResourceSources(parts []api.SimResult, total int) []api.ResourceSources {
	type amounts struct{ amount, wasted float64 }
	type key struct {
		resource string
		spent    bool
		action   string
	}
	sums := map[key]*amounts{}
	resources := []string{}
	add := func(p api.SimResult, resource string, spent bool, rows []api.ResourceSourceRow) {
		w := float64(p.IterationsRun) / float64(total)
		for _, row := range rows {
			k := key{resource, spent, row.Action}
			if sums[k] == nil {
				sums[k] = &amounts{}
			}
			sums[k].amount += row.Amount * w
			sums[k].wasted += row.Wasted * w
		}
	}
	for _, p := range parts {
		for _, r := range p.ResourceSources {
			if !slices.Contains(resources, r.Resource) {
				resources = append(resources, r.Resource)
			}
			add(p, r.Resource, false, r.Gained)
			add(p, r.Resource, true, r.Spent)
		}
	}
	if len(resources) == 0 {
		return nil
	}
	rows := func(resource string, spent bool) []api.ResourceSourceRow {
		out := []api.ResourceSourceRow{}
		for k, sum := range sums {
			if k.resource == resource && k.spent == spent {
				out = append(out, api.ResourceSourceRow{
					Action: k.action,
					Amount: adapter.OneDecimal(sum.amount),
					Wasted: adapter.OneDecimal(sum.wasted),
				})
			}
		}
		slices.SortFunc(out, func(a, b api.ResourceSourceRow) int {
			return cmp.Or(cmp.Compare(b.Amount, a.Amount), cmp.Compare(a.Action, b.Action))
		})
		return out
	}
	out := make([]api.ResourceSources, 0, len(resources))
	for _, resource := range resources {
		out = append(out, api.ResourceSources{Resource: resource, Gained: rows(resource, false), Spent: rows(resource, true)})
	}
	return out
}

// poolExtraAttacks is the iteration-weighted average of the parts' sources.
func poolExtraAttacks(parts []api.SimResult, total int) []api.ExtraAttackSource {
	sums := map[string]float64{}
	for _, p := range parts {
		w := float64(p.IterationsRun) / float64(total)
		for _, row := range p.ExtraAttackSources {
			sums[row.Source] += row.Count * w
		}
	}
	if len(sums) == 0 {
		return nil
	}
	out := make([]api.ExtraAttackSource, 0, len(sums))
	for source, count := range sums {
		out = append(out, api.ExtraAttackSource{Source: source, Count: adapter.OneDecimal(count)})
	}
	slices.SortFunc(out, func(a, b api.ExtraAttackSource) int {
		return cmp.Or(cmp.Compare(b.Count, a.Count), cmp.Compare(a.Source, b.Source))
	})
	return out
}
