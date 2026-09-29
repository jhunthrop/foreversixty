package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// slotRow is one slot's line in a band's report: the JSON and the
// markdown table share this shape.
type slotRow struct {
	Slot       string  `json:"slot"`
	ItemID     int     `json:"item_id,omitempty"`
	ItemName   string  `json:"item_name,omitempty"`
	Source     string  `json:"source,omitempty"`
	SourceKind string  `json:"source_kind,omitempty"`
	Score      float64 `json:"score,omitempty"`
	Verified   bool    `json:"verified"`
	SwapNote   string  `json:"swap_note,omitempty"`
	// EffectUnmodelled is true when the picked item carries an
	// effect_text the engine does NOT implement (effectids_generated.go)
	// -- this lane's brief, item 3: such a candidate is still scored on
	// its plain stats (score.go never saw the effect either way), but
	// the page shows "proc not simulated" on it rather than letting the
	// Score/Verified columns imply the whole item was accounted for.
	// False (omitted) for a candidate with no effect_text at all, and
	// for one whose effect IS implemented -- see rank.go's
	// hasImplementedEffect, the single predicate both this flag and the
	// effect-verification pass itself read.
	EffectUnmodelled bool `json:"effect_unmodelled,omitempty"`
	// Ties is every other candidate that scored identically to this
	// row's pick (slotPick.Ties, pick.go's own doc; this lane's brief,
	// defect 4) - "or Blackwater Cutlass" as populated alternatives the
	// page can print beside a pick that was really an arbitrary
	// lowest-id tie-break rather than a unique best. Empty for a slot
	// the trinket/effect/set-completion passes decided by a real sim
	// instead of score() (see slotPick.Ties's own doc). Kept exactly as
	// it was before Alternatives existed (below) - a consumer already
	// reading Ties keeps working unchanged.
	Ties []tieAlternative `json:"ties,omitempty"`
	// Alternatives is up to alternativesLimit candidates a player who
	// cannot get this row's pick can fall back on: this lane's brief
	// (owner defect A, 2026-09-29) - warrior-arms horde band 20's
	// main_hand published Forsaken Greataxe with no record that Smite's
	// Mighty Hammer (a Deadmines drop, item 7230) was ever considered,
	// leaving a player who cannot or will not run that quest nothing to
	// fall back on. Every Ties entry (score identical to the pick) is
	// listed first, dps_delta 0, ranked ahead of every lower-scoring
	// alternative (the wow-player review's own addendum to this lane's
	// brief) - a tie is exactly as good as the pick, and bySlot's own
	// score-sorted order does not otherwise distinguish a tie's rank
	// from a strictly-lower scorer's. The list is then filled to
	// alternativesLimit with the next-best sourced candidates by
	// score() (buildAlternatives, this file). See that function's own
	// doc for exactly what "next-best" excludes (the pick itself, its
	// pair-mate, and anything already listed as a tie).
	Alternatives []alternativeRow `json:"alternatives,omitempty"`
}

// tieAlternative is one equally-scored item slotRow.Ties names.
type tieAlternative struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

// alternativesLimit bounds how many candidates slotRow.Alternatives
// carries beyond the pick itself (this lane's brief: "the next best 3
// candidates by score after the pick") - enough for a player who
// cannot get the picked item's own source to see a genuine fallback,
// without ballooning the JSON with a slot's whole candidate pool.
const alternativesLimit = 3

// alternativeRow is one candidate slotRow.Alternatives names beyond
// the slot's own pick.
type alternativeRow struct {
	ItemID     int     `json:"item_id"`
	ItemName   string  `json:"item_name"`
	Score      float64 `json:"score"`
	SourceKind string  `json:"source_kind"`
	Source     string  `json:"source"`
	// DPSDelta is Score minus the pick's own published Score, in the
	// band's score unit (score.go's weighted-stat-plus-weapon-dps
	// total, not a measured DPS figure - buildAlternatives is built
	// from the scoring pass alone, per this lane's brief: "Reuse
	// slotPick.Ties/the existing scoring path; do not re-sim"). Exactly
	// 0 for a tie (this row came from pk.Ties, whose whole definition
	// is "scored identically to the pick"). Usually negative for every
	// other entry (bySlot's own list is score-sorted, so most
	// candidates after the pick score lower), but CAN be positive: a
	// runner-up whose own score() total is lower than the item it beat
	// can still be promoted into the pick by applySwaps' real-sim swap
	// pass (verify.go) - the demoted, higher-scoring item still belongs
	// in the published pick's alternatives list, just with a positive
	// delta.
	DPSDelta float64 `json:"dps_delta"`
}

// buildAlternatives is slotRow.Alternatives' own builder: pk.Ties
// first (dps_delta 0, this lane's brief addendum), then list (bySlot's
// own score-sorted pool for this slot - candidatesBySlot's contract,
// unfiltered by pick()'s own per-slot narrowing, so a two-handed
// main_hand pick's alternatives still offer one-handers mixed with
// two-handers the way the brief asks for), until alternativesLimit
// total, excluding:
//   - the pick's own item (already the row's own ItemID/ItemName)
//   - the pick's pair-mate's item, by id or name (pairSlot, verify.go's
//     own doc: the same physical ring/trinket/weapon cannot also be
//     offered as a fallback for THIS slot when its pair-mate already
//     wears it)
//   - anything already listed as a tie, so a candidate never appears
//     twice
//
// list is bySlot[slot] as passed to buildReport - NOT the narrower,
// per-slot list pick() computes internally for a dual-wielder's
// off_hand (main_hand's one-handers merged in) or a dual-wielder's
// main_hand (two-handers excluded): reusing the base scoring pool
// (candidatesBySlot's own output) rather than re-deriving pick()'s
// internal, stateful narrowing is this function's own scope call - see
// this lane's report for why. off_hand of a two-handed main_hand pick
// is never reached here at all: buildReport only calls this for a
// slot whose pk.Item is non-nil, and enforceTwoHandOffHandInvariant
// (pick.go) guarantees off_hand's own Item is nil whenever main_hand
// is two-handed.
func buildAlternatives(pk slotPick, slot string, list []scored, picks map[string]slotPick) []alternativeRow {
	if pk.Item == nil {
		return nil
	}
	var mateID int
	var mateName string
	if mate, ok := pairSlot[slot]; ok {
		if mp := picks[mate].Item; mp != nil {
			mateID, mateName = mp.ID, mp.Name
		}
	}
	seen := map[int]bool{pk.Item.ID: true}
	excluded := func(c scored) bool {
		if seen[c.ID] {
			return true
		}
		return mateID != 0 && (c.ID == mateID || (mateName != "" && c.Name == mateName))
	}
	add := func(out []alternativeRow, c scored) []alternativeRow {
		seen[c.ID] = true
		return append(out, alternativeRow{
			ItemID:     c.ID,
			ItemName:   c.Name,
			Score:      c.Score,
			SourceKind: c.Source.Kind,
			Source:     c.Source.Label,
			DPSDelta:   c.Score - pk.Item.Score,
		})
	}

	out := make([]alternativeRow, 0, alternativesLimit)
	for _, tie := range pk.Ties {
		if len(out) >= alternativesLimit {
			return out
		}
		if excluded(tie) {
			continue
		}
		out = add(out, tie)
	}
	for _, c := range list {
		if len(out) >= alternativesLimit {
			return out
		}
		if excluded(c) {
			continue
		}
		out = add(out, c)
	}
	return out
}

// bandReport is one band's whole answer for one faction: the pick per
// slot, the verification DPS, the diff against the previous band, and
// the counts an honest reader needs (how many eligible items had no
// known source).
type bandReport struct {
	Spec              string      `json:"spec"`
	Band              int         `json:"band"`
	Faction           string      `json:"faction"`
	Race              string      `json:"race"`
	Talents           string      `json:"talents"`
	TalentPoints      int         `json:"talent_points"`
	Weights           []weightRow `json:"weights"`
	Slots             []slotRow   `json:"slots"`
	SetDPS            float64     `json:"set_dps"`
	NoSourceCount     int         `json:"no_source_count"`
	NoSourceSample    []string    `json:"no_source_sample,omitempty"`
	NewAtBand         []string    `json:"new_at_band"`
	WeightsRunSeconds float64     `json:"weights_run_seconds"`
	VerifyRunSeconds  float64     `json:"verify_run_seconds"`
	VerifyErrors      []string    `json:"verify_errors,omitempty"`
	// Coverage is band.go's buildBandPool own per-slot count (lane
	// rank-guardrails' guardrail A): planner slot -> how many items
	// eligible() passed for this band+faction, and how many of those
	// sourceFor() could actually find a source for. A slot missing from
	// this map had zero eligible candidates at all (not even an
	// unsourced one) - see coverageRow's own doc for why a cross-class
	// set item still counts here even though it never reaches Slots or
	// NoSource. web/src/lib/bis/types.ts's BisBand.coverage is this
	// field's read contract.
	Coverage map[string]coverageRow `json:"coverage"`
	// ReferenceDPSPerPoint is the measured, un-normalised DPS this
	// band's weights run found for one point of the spec's own
	// reference stat (data/curated/specs.json's reference_stat) - this
	// lane's brief, item 2: the raw number every OTHER published
	// Weights[i].Weight is a ratio against (Weight_i/scale;
	// simrun.go's runWeights/referenceStatRawWeight read scale itself
	// straight off the engine's own result, since the normalised
	// weights map alone can never recover it - the reference stat's
	// own entry is exactly 1.0 by construction). Lets the page turn
	// "Strength 1.99" into "1 Attack Power = 0.07 DPS, Strength 1.99 (=
	// 0.14 DPS per point)" instead of publishing a bare, unitless
	// ratio with nothing saying what "1" means.
	ReferenceDPSPerPoint float64 `json:"reference_dps_per_point"`
}

type weightRow struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
	// Error is this weight's standard error, in the same units as
	// Weight (sim/adapter.Weights' own StatWeight.Error, carried
	// through unchanged) -- the page shows it as "±". See
	// isWeightSignificant's own doc for why this command publishes it
	// at all when the general-purpose /sim/weights tool's own
	// significance bar (sim/adapter.go's Insignificant field) is more
	// lenient than the one this command applies below.
	Error float64 `json:"error"`
	// Insignificant is this command's OWN, stricter call (see
	// isWeightSignificant), not sim/adapter's Insignificant field: the
	// page greys this row out and the Pawn/planner consumers of this
	// JSON should not treat it as a real number.
	Insignificant bool `json:"insignificant"`
}

// significanceErrorFraction and isWeightSignificant now live in
// weights.go, alongside effectiveWeights -- the one function that
// turns a raw stat-weights result into the plain numbers score()
// ranks by, zeroing exactly the weights this file's own
// isWeightSignificant call marks Insignificant below.

// noSourceSampleSize bounds how many unsourced item names the JSON
// and markdown carry - the count is exact, the sample is just enough
// to spot-check without shipping a multi-thousand-row list every run.
const noSourceSampleSize = 15

// buildReport assembles one band+faction's report from pick() output,
// the weights this band used, verification results, and the previous
// band's picks (nil for the first band run). bySlot is candidatesBySlot's
// own output for this band+faction (main.go's runSpec builds a fresh one
// per band iteration) - buildAlternatives reads it to fill each row's
// Alternatives; a nil/empty bySlot (every test but the ones this lane's
// brief adds) simply publishes no alternatives, same as an empty map
// would. referenceDPSPerPoint is runWeights' own second return
// (simrun.go) - this lane's brief, item 2.
func buildReport(spec specInfo, band int, faction, race, talents string, talentPoints int, weights map[string]api.StatWeight, weightOrder []string, picks map[string]slotPick, setDPS float64, swaps []swapResult, noSource []candidate, previous map[string]slotPick, weightsSeconds, verifySeconds float64, verifyErrors []string, coverage map[string]coverageRow, bySlot map[string][]scored, referenceDPSPerPoint float64) bandReport {
	swapBySlot := make(map[string]swapResult, len(swaps))
	for _, s := range swaps {
		swapBySlot[s.Slot] = s
	}
	erroredSlots := make(map[string]bool, len(verifyErrors))
	for _, e := range verifyErrors {
		if slot, _, ok := strings.Cut(e, ":"); ok {
			erroredSlots[slot] = true
		}
	}

	rows := make([]slotRow, 0, len(slotOrder))
	for _, slot := range slotOrder {
		pk := picks[slot]
		row := slotRow{Slot: slot}
		if pk.Item != nil {
			row.ItemID = pk.Item.ID
			row.ItemName = pk.Item.Name
			row.Score = pk.Item.Score
			row.EffectUnmodelled = pk.Item.EffectText != "" && !hasImplementedEffect(pk.Item.candidate)
			for _, tie := range pk.Ties {
				row.Ties = append(row.Ties, tieAlternative{ItemID: tie.ID, ItemName: tie.Name})
			}
			row.Alternatives = buildAlternatives(pk, slot, bySlot[slot], picks)
			if pk.Item.HasSource {
				row.Source = pk.Item.Source.Label
				row.SourceKind = pk.Item.Source.Kind
			}
			switch {
			case erroredSlots[slot]:
				row.Verified = false
				row.SwapNote = "the runner-up's verification sim failed (an engine-side error, not a scoring one - see verify_errors); the pick is unconfirmed against it"
			default:
				row.Verified = true
				if sw, ok := swapBySlot[slot]; ok && sw.Beat && pk.RunnerUp != nil {
					// applySwaps already promoted the runner-up into pk.Item and
					// demoted the scored pick to pk.RunnerUp: this row IS the
					// measured winner, verified by that very run.
					row.SwapNote = fmt.Sprintf("beat the scored pick %s (id %d) in the sim: %.1f vs %.1f set DPS", pk.RunnerUp.Name, pk.RunnerUp.ID, sw.SwapDPS, sw.BaselineDPS)
				}
			}
		}
		rows = append(rows, row)
	}

	var newAt []string
	for _, slot := range slotOrder {
		cur := picks[slot].Item
		if cur == nil {
			continue
		}
		var prevID int
		if previous != nil && previous[slot].Item != nil {
			prevID = previous[slot].Item.ID
		}
		if prevID != cur.ID {
			newAt = append(newAt, fmt.Sprintf("%s: %s", slot, cur.Name))
		}
	}

	sampleNames := make([]string, 0, noSourceSampleSize)
	for i, c := range noSource {
		if i >= noSourceSampleSize {
			break
		}
		sampleNames = append(sampleNames, fmt.Sprintf("%d %s", c.ID, c.Name))
	}

	wrows := make([]weightRow, 0, len(weightOrder))
	for _, id := range weightOrder {
		w := weights[id]
		wrows = append(wrows, weightRow{
			Stat:          id,
			Weight:        w.Weight,
			Error:         w.Error,
			Insignificant: !isWeightSignificant(w),
		})
	}

	return bandReport{
		Spec:                 spec.Spec,
		Band:                 band,
		Faction:              faction,
		Race:                 race,
		Talents:              talents,
		TalentPoints:         talentPoints,
		Weights:              wrows,
		Slots:                rows,
		SetDPS:               setDPS,
		NoSourceCount:        len(noSource),
		NoSourceSample:       sampleNames,
		NewAtBand:            nonNil(newAt),
		WeightsRunSeconds:    weightsSeconds,
		VerifyRunSeconds:     verifySeconds,
		VerifyErrors:         verifyErrors,
		Coverage:             coverage,
		ReferenceDPSPerPoint: referenceDPSPerPoint,
	}
}

// coverageTotals sums every slot's coverageRow into one band-wide
// figure: how many eligible items this band+faction saw across every
// slot, and how many of those had a source. Used by main.go's own
// one-line nightly log summary (this lane's brief, guardrail A) -
// giving a reader watching the run the same "N of M sourced" honesty
// the page's per-band coverage table will show, without having to
// open the JSON.
func coverageTotals(coverage map[string]coverageRow) (eligible, sourced int) {
	for _, row := range coverage {
		eligible += row.Eligible
		sourced += row.Sourced
	}
	return eligible, sourced
}

// worstCoveredSlot names the slot with the lowest sourced/eligible
// ratio (ties broken by the most eligible items, then by slot name,
// so the result is deterministic across runs) - the one line's own
// "and worst is X" clause points a reader straight at the slot most
// worth a data lane's attention, rather than making them scan the
// full per-slot map for it. A slot with zero eligible items is
// excluded: 0/0 is not a coverage gap, it is "nothing exists here to
// gate at all".
func worstCoveredSlot(coverage map[string]coverageRow) (slot string, row coverageRow, ok bool) {
	for _, s := range slotOrder {
		r, present := coverage[s]
		if !present || r.Eligible == 0 {
			continue
		}
		if !ok || ratio(r) < ratio(row) {
			slot, row, ok = s, r, true
		}
	}
	return slot, row, ok
}

func ratio(r coverageRow) float64 {
	if r.Eligible == 0 {
		return 1
	}
	return float64(r.Sourced) / float64(r.Eligible)
}

// coverageSummary is the one-line-per-band nightly log string this
// lane's brief asks for: the band-wide total, plus the single
// worst-covered slot so a reader does not have to open the JSON to
// see which slot most needs a data lane's attention.
func coverageSummary(coverage map[string]coverageRow) string {
	eligible, sourced := coverageTotals(coverage)
	slot, row, ok := worstCoveredSlot(coverage)
	if !ok {
		return fmt.Sprintf("%d/%d eligible items sourced", sourced, eligible)
	}
	return fmt.Sprintf("%d/%d eligible items sourced; worst slot %s (%d/%d)", sourced, eligible, slot, row.Sourced, row.Eligible)
}

// titleCase upper-cases a single lower-kebab word's first letter -
// "horde" -> "Horde" - for the markdown's faction headings. strings.
// Title is deprecated (it does not handle multi-word input correctly,
// which is not a problem here, but this project's lint config flags
// its use regardless), so this is the one-word case written out.
func titleCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// specReport is data/builds/<build>/bis/<spec>.json's whole shape -
// lane bis-web's read contract (this lane's brief, step 3): every
// bandReport field name is kept exactly as the prototype defined it;
// engine_version is the one field this wrapper adds beyond what wraps
// the array (spec, build, generated_at, bands).
type specReport struct {
	Spec          string       `json:"spec"`
	Build         string       `json:"build"`
	EngineVersion string       `json:"engine_version"`
	GeneratedAt   string       `json:"generated_at"`
	Bands         []bandReport `json:"bands"`
}

// writeSpecReport writes path per the specReport contract above.
// GeneratedAt is RFC3339, in UTC so two runs on different machines (a
// dev's laptop, the nightly workflow's runner) produce comparable
// timestamps rather than each in its own local zone.
func writeSpecReport(path, spec, build string, reports []bandReport) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	out := specReport{
		Spec:          spec,
		Build:         build,
		EngineVersion: enginever.Version,
		GeneratedAt:   time.Now().UTC().Format(time.RFC3339),
		Bands:         reports,
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

// writeMarkdown renders every band+faction report for one spec, most
// recent band last, grouped by faction, in the shape this lane's
// brief asks for: a table per band (slot, item, source with faction
// badge, score, verified) and a "New at L" list.
func writeMarkdown(path string, spec specInfo, reports []bandReport) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Leveling BiS: %s\n\n", spec.Name)
	fmt.Fprintf(&b, "Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.\n\n")

	byFaction := map[string][]bandReport{}
	var factions []string
	for _, r := range reports {
		if _, ok := byFaction[r.Faction]; !ok {
			factions = append(factions, r.Faction)
		}
		byFaction[r.Faction] = append(byFaction[r.Faction], r)
	}
	sort.Strings(factions)

	for _, faction := range factions {
		fmt.Fprintf(&b, "## %s\n\n", titleCase(faction))
		for _, r := range byFaction[faction] {
			fmt.Fprintf(&b, "### Band %d (%s, %s)\n\n", r.Band, r.Race, r.Talents)
			fmt.Fprintf(&b, "Set DPS (verified): %.1f. Weights run: %.1fs. Verify run: %.1fs. %d eligible items had no known source.\n\n",
				r.SetDPS, r.WeightsRunSeconds, r.VerifyRunSeconds, r.NoSourceCount)

			fmt.Fprintf(&b, "Stat weights (normalized to %s = 1.0, error under %.0f%% of the weight to publish - see report.go's isWeightSignificant): ", spec.ReferenceStat, significanceErrorFraction*100)
			parts := make([]string, len(r.Weights))
			for i, w := range r.Weights {
				if w.Insignificant {
					parts[i] = fmt.Sprintf("%s=not significant (%.3f ± %.3f)", w.Stat, w.Weight, w.Error)
				} else {
					parts[i] = fmt.Sprintf("%s=%.3f ± %.3f", w.Stat, w.Weight, w.Error)
				}
			}
			fmt.Fprintln(&b, strings.Join(parts, ", "))
			b.WriteString("\n")

			b.WriteString("| Slot | Item | Source | Score | Verified | Alternatives |\n")
			b.WriteString("|---|---|---|---|---|---|\n")
			for _, row := range r.Slots {
				item := "-"
				source := "-"
				score := ""
				verified := ""
				alternatives := ""
				if row.ItemID != 0 {
					item = fmt.Sprintf("%s (%d)", row.ItemName, row.ItemID)
					if len(row.Ties) > 0 {
						alts := make([]string, len(row.Ties))
						for i, t := range row.Ties {
							alts[i] = fmt.Sprintf("%s (%d)", t.ItemName, t.ItemID)
						}
						item += " (or " + strings.Join(alts, ", ") + ")"
					}
					score = strconv.FormatFloat(row.Score, 'f', 1, 64)
					verified = "yes"
					if !row.Verified {
						verified = "no - " + row.SwapNote
					}
					if row.Source != "" {
						source = fmt.Sprintf("%s [%s]", row.Source, row.SourceKind)
					} else {
						source = "no known source"
					}
					// Alternatives (this lane's brief, item 1): the
					// next-best sourced candidates after the pick, so a
					// player reading the human-facing markdown table -
					// not only a script reading the JSON - sees a real
					// fallback when the pick's own source is out of
					// reach.
					alternatives = "-"
					if len(row.Alternatives) > 0 {
						alts := make([]string, len(row.Alternatives))
						for i, a := range row.Alternatives {
							alts[i] = fmt.Sprintf("%s (%d, %+.1f) [%s]", a.ItemName, a.ItemID, a.DPSDelta, a.SourceKind)
						}
						alternatives = strings.Join(alts, "; ")
					}
				}
				fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %s |\n", row.Slot, item, source, score, verified, alternatives)
			}
			b.WriteString("\n")

			if len(r.NewAtBand) > 0 {
				fmt.Fprintf(&b, "**New at %d:** %s\n\n", r.Band, strings.Join(r.NewAtBand, "; "))
			} else {
				b.WriteString("**New at this band:** nothing changed from the previous band.\n\n")
			}

			if len(r.NoSourceSample) > 0 {
				fmt.Fprintf(&b, "No-known-source sample (%d of %d, see the JSON for more): %s\n\n",
					len(r.NoSourceSample), r.NoSourceCount, strings.Join(r.NoSourceSample, "; "))
			}
		}
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

// nonNil is xs, or an empty slice for nil: the JSON contract promises an
// array, and encoding/json writes a nil slice as null, which the site's
// loader and the addon pipeline would then have to special-case.
func nonNil(xs []string) []string {
	if xs == nil {
		return []string{}
	}
	return xs
}
