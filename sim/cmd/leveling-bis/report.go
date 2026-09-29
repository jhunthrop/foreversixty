package main

import (
	"encoding/json"
	"fmt"
	"math"
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

// significanceErrorFraction is this command's own publication bar
// (2026-09-28 weights-effects lane; owner, looking at a level-20
// hunter's weights: "these stat weights look like garbage"). The
// general-purpose /sim/weights tool's own bar
// (sim/adapter.go's `Insignificant: errAmt >= math.Abs(weight)`) only
// catches a weight the error swallows entirely -- appropriate for a
// tool a player can re-run at higher precision on demand. This
// command's sweep runs a fixed, small budget (100 iterations per
// direction by default -- see weightsIterations's own flag doc) and
// publishes every weight on the BiS page's aside as a fact a player
// acts on without re-running anything, so it holds every row to a
// tighter bar: an error under 25% of the weight's own value, or the
// row reports "not significant" instead of a number nobody should
// trust.
const significanceErrorFraction = 0.25

// isWeightSignificant applies significanceErrorFraction to one
// engine-reported weight. A weight of exactly zero is never
// significant (there is nothing for a 25%-of-value bar to compare
// against, and a hard-capped or unmoved stat reads back as
// weight=0, error=0 the same way sim/adapter.go's own comment
// describes for its own field).
func isWeightSignificant(w api.StatWeight) bool {
	if w.Weight == 0 {
		return false
	}
	return w.Error < significanceErrorFraction*math.Abs(w.Weight)
}

// noSourceSampleSize bounds how many unsourced item names the JSON
// and markdown carry - the count is exact, the sample is just enough
// to spot-check without shipping a multi-thousand-row list every run.
const noSourceSampleSize = 15

// buildReport assembles one band+faction's report from pick() output,
// the weights this band used, verification results, and the previous
// band's picks (nil for the first band run).
func buildReport(spec specInfo, band int, faction, race, talents string, talentPoints int, weights map[string]api.StatWeight, weightOrder []string, picks map[string]slotPick, setDPS float64, swaps []swapResult, noSource []candidate, previous map[string]slotPick, weightsSeconds, verifySeconds float64, verifyErrors []string) bandReport {
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
		Spec:              spec.Spec,
		Band:              band,
		Faction:           faction,
		Race:              race,
		Talents:           talents,
		TalentPoints:      talentPoints,
		Weights:           wrows,
		Slots:             rows,
		SetDPS:            setDPS,
		NoSourceCount:     len(noSource),
		NoSourceSample:    sampleNames,
		NewAtBand:         nonNil(newAt),
		WeightsRunSeconds: weightsSeconds,
		VerifyRunSeconds:  verifySeconds,
		VerifyErrors:      verifyErrors,
	}
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

			b.WriteString("| Slot | Item | Source | Score | Verified |\n")
			b.WriteString("|---|---|---|---|---|\n")
			for _, row := range r.Slots {
				item := "-"
				source := "-"
				score := ""
				verified := ""
				if row.ItemID != 0 {
					item = fmt.Sprintf("%s (%d)", row.ItemName, row.ItemID)
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
				}
				fmt.Fprintf(&b, "| %s | %s | %s | %s | %s |\n", row.Slot, item, source, score, verified)
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
