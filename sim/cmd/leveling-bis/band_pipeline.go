package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// factionRace is one faction a band is ranked for, and the guide's race for it.
type factionRace struct{ name, race string }

// rankingWeights is one weight set a band can be ranked by, with everything
// the report publishes about it.
type rankingWeights struct {
	label                string
	request              api.SimRequest
	result               map[string]api.StatWeight
	weights              map[string]float64
	referenceDPSPerPoint float64
	reason               string
	lowConfidence        bool
}

// bandRanker is one band and pass's fixed inputs: everything rank needs
// besides the weights.
type bandRanker struct {
	runner           engineRunner
	spec             specInfo
	pass             string
	band             int
	talents          string // the engine layout
	items            []candidate
	lootIdx          lootIndex
	weaponSubclasses map[int]bool
	requiresDagger   bool
	castsShoot       bool
	setCat           setCatalog
	factions         []factionRace
}

// factionRun is one faction's ranked band under one weight set, up to the
// set the guard (keepVerifyBest) weighs.
type factionRun struct {
	pool            bandPool
	bySlot          map[string][]scored
	pickBySlot      map[string][]scored
	verified        verifiedBand
	considered      []consideredSet
	heldUnconfirmed int
	trinketSeconds  float64
	effectSeconds   float64
	verifySeconds   float64
}

func (r bandRanker) logf(format string, args ...any) {
	log.Printf("leveling-bis: %s band %d "+format, append([]any{r.spec.Spec, r.band}, args...)...)
}

// rank runs the band's whole pipeline under w for both factions: the
// picks, the wand credit, the trinket tournaments and their cross-faction
// reconciliation, the effect tournaments, the verify swaps, set completion
// and the confirmed-stats check.
func (r bandRanker) rank(w rankingWeights) (map[string]*factionRun, error) {
	parallel := runsConcurrently(r.runner)
	firstHalves := make([]*factionWork, len(r.factions))
	if err := forEach(len(r.factions), parallel, func(i int) error {
		fw, err := r.pickAndRankTrinkets(w, r.factions[i])
		firstHalves[i] = fw
		return err
	}); err != nil {
		return nil, err
	}
	work := make(map[string]*factionWork, len(r.factions))
	for i, f := range r.factions {
		work[f.name] = firstHalves[i]
	}
	// Both factions' trinket tournaments are finished: reconcile them
	// before either continues (faction_trinkets.go), so a faction-neutral
	// trinket's verdict never differs between the sides over noise alone.
	allianceWork, hordeWork := work["alliance"], work["horde"]
	var notes []string
	allianceWork.picks, hordeWork.picks, notes = reconcileFactionTrinkets(
		r.runner, r.spec, r.spec.ClassSlug, r.band, r.talents, r.lootIdx,
		factionTrinketInputs{Faction: "alliance", Race: allianceWork.race, BySlot: allianceWork.bySlot}, allianceWork.picks,
		factionTrinketInputs{Faction: "horde", Race: hordeWork.race, BySlot: hordeWork.bySlot}, hordeWork.picks,
	)
	for _, n := range notes {
		r.logf("faction reconcile (%s weights): %s", w.label, n)
	}
	finished := make([]*factionRun, len(r.factions))
	if err := forEach(len(r.factions), parallel, func(i int) error {
		run, err := r.finish(work[r.factions[i].name])
		finished[i] = run
		return err
	}); err != nil {
		return nil, err
	}
	out := make(map[string]*factionRun, len(r.factions))
	for i, f := range r.factions {
		out[f.name] = finished[i]
	}
	return out, nil
}

// pickAndRankTrinkets is one faction's first half: pick (twice for a spec
// that shoots a wand, the second time at the wand credit the first set
// measures; wand_credit.go) and the trinket tournaments, trinket1 first so
// trinket2 sees its real pick (clearTrinketPlaceholders' doc).
func (r bandRanker) pickAndRankTrinkets(w rankingWeights, f factionRace) (*factionWork, error) {
	wandCredit := 0.0
	if r.castsShoot {
		wandCredit = fullWandScorePerDPS(w.referenceDPSPerPoint)
	}
	first := pickBand(r.items, r.lootIdx, r.spec, r.band, f.name, w.weights, wandCredit, r.weaponSubclasses, r.requiresDagger)
	if r.castsShoot {
		credit, err := measureWandCredit(r.runner, r.spec, f.race, r.spec.ClassSlug, r.band, r.talents, first.picks, w.referenceDPSPerPoint)
		if err != nil {
			return nil, fmt.Errorf("band %d %s: %w", r.band, f.name, err)
		}
		r.logf("%s %s: wand credit %.4f score per wand DPS (whole fight %.4f, %s weights)", f.name, r.pass, credit, wandCredit, w.label)
		first = pickBand(r.items, r.lootIdx, r.spec, r.band, f.name, w.weights, credit, r.weaponSubclasses, r.requiresDagger)
	}
	picks := clearTrinketPlaceholders(first.picks)
	start := time.Now()
	for _, slot := range trinketSlots {
		var notes []string
		picks, notes = rankTrinketSlot(r.runner, r.spec, f.race, r.spec.ClassSlug, r.band, r.talents, picks, first.pickBySlot, slot, w.weights)
		for _, n := range notes {
			r.logf("%s: %s", f.name, n)
		}
	}
	return &factionWork{faction: f.name, race: f.race, pool: first.pool, bySlot: first.bySlot, pickBySlot: first.pickBySlot, picks: picks, trinketSeconds: time.Since(start).Seconds()}, nil
}

// finish is one faction's second half: the effect tournaments (score()
// cannot see a proc), the two-hand/off-hand rule re-asserted after them,
// the verify swaps, set completion and the confirmed-stats check, keeping
// every set it passes through for the guard.
func (r bandRanker) finish(fw *factionWork) (*factionRun, error) {
	picks := fw.picks
	effectStart := time.Now()
	for _, slot := range slotsNeedingEffectVerification(fw.pickBySlot) {
		var notes []string
		picks, notes = rankSlotWithEffects(r.runner, r.spec, fw.race, r.spec.ClassSlug, r.band, r.talents, picks, fw.pickBySlot, slot)
		for _, n := range notes {
			r.logf("%s: %s", fw.faction, n)
		}
	}
	if r.castsShoot {
		var notes []string
		picks, notes = rankWandSlot(r.runner, r.spec, fw.race, r.spec.ClassSlug, r.band, r.talents, picks, fw.pickBySlot)
		for _, n := range notes {
			r.logf("%s: %s", fw.faction, n)
		}
	}
	effectSeconds := time.Since(effectStart).Seconds()
	picks = fillEmptyOffHand(r.spec.Spec, fw.pickBySlot, enforceTwoHandOffHandInvariant(picks))

	verifyStart := time.Now()
	verified, err := verifyAndSwap(r.runner, r.spec, fw.race, r.spec.ClassSlug, r.band, r.talents, picks, fw.pickBySlot)
	if err != nil {
		return nil, fmt.Errorf("band %d %s: %w", r.band, fw.faction, err)
	}
	considered := []consideredSet{
		{Label: consideredBeforeVerify, Picks: picks},
		{Label: consideredBeforeSets, Picks: verified.picks},
	}
	var setNotes []string
	verified, setNotes, err = completeSets(r.runner, r.spec, fw.race, r.spec.ClassSlug, r.band, r.talents, verified, fw.pickBySlot, r.setCat)
	if err != nil {
		return nil, fmt.Errorf("band %d %s: %w", r.band, fw.faction, err)
	}
	for _, n := range setNotes {
		r.logf("%s: %s", fw.faction, n)
	}
	considered = append(considered, consideredSet{Label: consideredBeforeConfirmed, Picks: verified.picks})
	verified, held, err := preferConfirmedStats(r.runner, r.spec, fw.race, r.spec.ClassSlug, r.band, r.talents, verified, fw.pickBySlot)
	if err != nil {
		return nil, fmt.Errorf("band %d %s: %w", r.band, fw.faction, err)
	}
	return &factionRun{
		pool: fw.pool, bySlot: fw.bySlot, pickBySlot: fw.pickBySlot,
		verified: verified, considered: considered, heldUnconfirmed: held,
		trinketSeconds: fw.trinketSeconds, effectSeconds: effectSeconds,
		verifySeconds: time.Since(verifyStart).Seconds(),
	}, nil
}

// chooseRun is the index of the run whose final set a faction publishes:
// the one that measures highest at the guard harness, runs[0] (the band's
// preferred weights) on a tie. Every set is simmed at the same seed and
// iteration count, so the difference between two sets carries far less
// noise than either set's own error.
//
// Round 8: ranking by a geared sweep alone left band rows below the ladder
// sweep's own ranking (Shadow band 60, Arms raid band 60, the tanks), because
// a geared sweep reads the margins of one set and the picks can move away
// from it. Every band ranked by a geared sweep is ranked by the ladder sweep
// too, and the better verified set is published.
func chooseRun(runner engineRunner, spec specInfo, race string, level int, talents string, finals []map[string]slotPick) (int, []float64, error) {
	scores := make([]float64, len(finals))
	best := 0
	for i, picks := range finals {
		score, _, err := measureSet(runner, spec, race, spec.ClassSlug, level, talents, picks)
		if err != nil {
			return 0, nil, fmt.Errorf("measuring run %d's final set: %w", i, err)
		}
		scores[i] = score
		if score > scores[best] {
			best = i
		}
	}
	return best, scores, nil
}
