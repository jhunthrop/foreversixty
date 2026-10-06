package main

import (
	"fmt"
	"sort"
)

// deepTreePoints is the depth an archetype's deep tree is filled to
// before the rest is spent anywhere: the tier-6 capstone's 30 points
// plus the capstone itself.
const deepTreePoints = 31

// maxSwapPoints is the largest point group a single swap moves.
const maxSwapPoints = 5

// minDistinct is how many points two kept candidates must differ by
// before the cap starts admitting near-duplicates.
const minDistinct = 2

// candidate is one build to sim, with where it came from.
type candidate struct {
	Label    string
	Build    build
	Estimate float64 // credit-table DPS gain over the guide
	// Structural candidates (archetypes, the guide re-spent) are kept
	// ahead of single swaps when the cap bites.
	Structural bool
}

// sortedIDs is a build's or credit table's ids in ascending order, so
// every walk is deterministic.
func sortedIDs[V any](m map[int]V) []int {
	ids := make([]int, 0, len(m))
	for id := range m {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	return ids
}

// swaps is every legal build one group move away from base: 1..5
// points out of a talent the engine does not credit with damage and
// into one it does, in any tree, at the same total.
func swaps(t talentTrees, base build, credits map[int]credit, labelPrefix string) []candidate {
	total := base.points()
	var out []candidate
	for _, src := range sortedIDs(base) {
		if credits[src].Damage {
			continue
		}
		for k := 1; k <= base[src] && k <= maxSwapPoints; k++ {
			for _, dst := range sortedIDs(credits) {
				if !credits[dst].Damage || dst == src || base[dst]+k > t.byID[dst].MaxRank {
					continue
				}
				b := base.with(src, -k).with(dst, k)
				if t.legal(b, total) != nil {
					continue
				}
				out = append(out, candidate{
					Label: fmt.Sprintf("%s%d %s -> %s", labelPrefix, k, t.byID[src].Name, t.byID[dst].Name),
					Build: b,
				})
			}
		}
	}
	return out
}

// scorer ranks single points for the greedy fills: by credit, then -
// among equal credits, which is every talent the engine gives nothing
// - a talent the guide itself takes (so a re-spent build keeps the
// guide's own utility and unmodeled picks wherever DPS does not care),
// then the lower tier (it unlocks more), then the lower id.
type scorer struct {
	credits map[int]credit
	prefer  build
}

func (s scorer) better(t talentTrees, a, b int) bool {
	if ca, cb := s.credits[a].PerPoint, s.credits[b].PerPoint; ca != cb {
		return ca > cb
	}
	if pa, pb := s.prefer[a] > 0, s.prefer[b] > 0; pa != pb {
		return pa
	}
	return t.byID[a].Tier < t.byID[b].Tier
}

// bestPoint is the legal single point s ranks highest among the
// talents allow admits.
func (s scorer) bestPoint(t talentTrees, b build, allow func(id int) bool) (int, bool) {
	best, found := 0, false
	for _, id := range sortedIDs(t.byID) {
		if !allow(id) || !t.canAdd(b, id) {
			continue
		}
		if !found || s.better(t, id, best) {
			best, found = id, true
		}
	}
	return best, found
}

// bundle is the cheapest way to put one point into target from b:
// the filler points (best-credit talents in the same tree, below
// target's tier) and prerequisite points needed first, then the point
// itself, within budget. gain is the bundle's summed credit.
func (s scorer) bundle(t talentTrees, b build, target, budget int) (build, float64, int, bool) {
	node := t.byID[target]
	if b[target] >= node.MaxRank {
		return nil, 0, 0, false
	}
	ti := t.treeOf[target]
	nb, gain := b, 0.0
	for size := 1; size <= budget; size++ {
		next := target
		switch {
		case t.canAdd(nb, target):
		case node.PrereqTalentID != 0 && nb[node.PrereqTalentID] < node.PrereqRank && t.canAdd(nb, node.PrereqTalentID):
			next = node.PrereqTalentID
		default:
			filler, ok := s.bestPoint(t, nb, func(id int) bool {
				return t.treeOf[id] == ti && t.byID[id].Tier < node.Tier
			})
			if !ok {
				return nil, 0, 0, false
			}
			next = filler
		}
		nb = nb.with(next, 1)
		gain += s.credits[next].PerPoint
		if next == target {
			return nb, gain, size, true
		}
	}
	return nil, 0, 0, false
}

// fill spends b up to total: first deep tree deepTree (if >= 0) to
// deepTreePoints by best single point, then everywhere by the best
// credit per point of any bundle that fits.
func (s scorer) fill(t talentTrees, b build, total, deepTree int) build {
	if deepTree >= 0 {
		for t.treePoints(b)[deepTree] < deepTreePoints && b.points() < total {
			id, ok := s.bestPoint(t, b, func(id int) bool { return t.treeOf[id] == deepTree })
			if !ok {
				break
			}
			b = b.with(id, 1)
		}
	}
	for b.points() < total {
		next, ok := s.bestBundle(t, b, total-b.points())
		if !ok {
			break
		}
		b = next
	}
	return b
}

func (s scorer) bestBundle(t talentTrees, b build, budget int) (build, bool) {
	var best build
	bestValue, bestSize := 0.0, 0
	for _, id := range sortedIDs(t.byID) {
		nb, gain, size, ok := s.bundle(t, b, id, budget)
		if !ok {
			continue
		}
		value := gain / float64(size)
		if best == nil || value > bestValue || (value == bestValue && size < bestSize) {
			best, bestValue, bestSize = nb, value, size
		}
	}
	return best, best != nil
}

// strip takes every point in a talent removable admits out of b, as
// far as the remaining talents' tier gates and prerequisites allow.
func strip(t talentTrees, b build, removable func(id int) bool) build {
	for changed := true; changed; {
		changed = false
		for _, id := range sortedIDs(b) {
			if !removable(id) || b[id] == 0 {
				continue
			}
			nb := b.with(id, -1)
			if t.legal(nb, nb.points()) == nil {
				b, changed = nb, true
			}
		}
	}
	return b
}

// generate is the whole candidate pool around guide (deduplicated,
// the guide itself excluded), before the cap. total is the level's
// whole budget: the structural builds spend all of it even when the
// guide, re-read against the active build, spends fewer. modeled is
// the talents the engine's code reads: the guide re-spent twice, once
// moving every non-damage point and once moving only modeled ones -
// a talent the engine never reads sims as zero, so moving it is the
// engine's blind spot rather than a finding.
func generate(t talentTrees, guide build, credits map[int]credit, modeled map[int]bool, total, refineTop int) []candidate {
	var pool []candidate
	s := scorer{credits: credits, prefer: guide}
	nonDamage := func(id int) bool { return !credits[id].Damage }
	modeledNonDamage := func(id int) bool { return modeled[id] && !credits[id].Damage }
	structural := []candidate{{
		Label:      "guide, modeled non-damage points re-spent",
		Build:      s.fill(t, strip(t, guide, modeledNonDamage), total, -1),
		Structural: true,
	}, {
		Label:      "guide, all non-damage points re-spent",
		Build:      s.fill(t, strip(t, guide, nonDamage), total, -1),
		Structural: true,
	}}
	for ti, tree := range t.trees {
		structural = append(structural, candidate{
			Label:      fmt.Sprintf("deep %s", tree.Name),
			Build:      s.fill(t, build{}, total, ti),
			Structural: true,
		})
	}
	for _, st := range structural {
		pool = append(pool, st)
		refined := swaps(t, st.Build, credits, st.Label+" + ")
		for i := range refined {
			refined[i].Estimate = estimateGain(st.Build, refined[i].Build, credits)
		}
		sort.SliceStable(refined, func(i, j int) bool { return refined[i].Estimate > refined[j].Estimate })
		if len(refined) > refineTop {
			refined = refined[:refineTop]
		}
		pool = append(pool, refined...)
	}
	pool = append(pool, swaps(t, guide, credits, "guide: ")...)
	for i := range pool {
		pool[i].Estimate = estimateGain(guide, pool[i].Build, credits)
	}
	return dedupe(t, guide, pool)
}

// dedupe drops candidates equal to the guide or to an earlier one.
func dedupe(t talentTrees, guide build, pool []candidate) []candidate {
	seen := map[string]bool{t.key(guide): true}
	var out []candidate
	for _, c := range pool {
		k := t.key(c.Build)
		if seen[k] {
			continue
		}
		seen[k] = true
		out = append(out, c)
	}
	return out
}

// capCandidates keeps at most limit candidates: structural ones first,
// then by estimated gain, admitting a candidate only if it differs from
// every kept one by at least minDistinct points; any room left is then
// filled from the near-duplicates in the same order.
func capCandidates(pool []candidate, limit int) []candidate {
	ordered := append([]candidate(nil), pool...)
	sort.SliceStable(ordered, func(i, j int) bool {
		if ordered[i].Structural != ordered[j].Structural {
			return ordered[i].Structural
		}
		return ordered[i].Estimate > ordered[j].Estimate
	})
	var kept, deferred []candidate
	for _, c := range ordered {
		if len(kept) >= limit {
			break
		}
		if distinctFromAll(c.Build, kept) {
			kept = append(kept, c)
		} else {
			deferred = append(deferred, c)
		}
	}
	for _, c := range deferred {
		if len(kept) >= limit {
			break
		}
		kept = append(kept, c)
	}
	return kept
}

func distinctFromAll(b build, kept []candidate) bool {
	for _, k := range kept {
		if distance(b, k.Build) < minDistinct {
			return false
		}
	}
	return true
}
