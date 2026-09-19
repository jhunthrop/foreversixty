// Package adapter turns an engine result into the logs engine's per-fight
// summary, so the report page's components render a simulation with no
// second renderer.
//
// The engine accumulates its metrics across every iteration. A summary
// describes one fight. So every count and every total here is divided by
// IterationsDone before it enters the summary; the distribution stays in
// api.SimResult.DPS, which is where a range belongs.
//
// The target is summary.Summary as of logs engine 0.5.3. Four of its
// twenty fields - ThreatByTarget, Taunts, Mechanics and Phases - are
// things a simulation cannot know, and all four are set explicitly empty
// rather than left nil, because the report components render a list and a
// nil slice marshals as null.
package adapter

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/statid"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core/proto"
)

var (
	// ErrSimFailed is returned when the engine itself reported a failure.
	ErrSimFailed = errors.New("adapter: the engine reported an error")
	// ErrAborted is returned when the run was stopped on request. It is
	// not a failure: the user pressed Stop, there is nothing to report
	// and nothing to fix. The engine says so with an ErrorOutcome whose
	// Type is ErrorOutcomeAborted and whose Message is EMPTY, so a guard
	// on the message alone lets an abort through and the caller then
	// reports "iterations_done is 0" - a corrupt result - for something
	// the user asked for.
	ErrAborted = errors.New("adapter: the run was stopped")
	// ErrNoPlayer is returned when the result carries no player metrics,
	// which means the request had no player in party one.
	ErrNoPlayer = errors.New("adapter: the result has no player metrics")
	// ErrDuplicateRow is returned when two rows of the summary share the
	// key the logs engine and the report components identify a row by.
	// It cannot happen for a real fight and must not happen for a sim:
	// a repeated key is a runtime error in the report's keyed blocks.
	ErrDuplicateRow = errors.New("adapter: two summary rows share one key")
)

// playerGUID is the synthetic unit id the summary uses for the simmed
// player. A real fight's guids come from the combat log; a sim has none,
// so it gets a stable one that the report components treat identically.
const playerGUID = "sim-player"

// auraTypeBuff is the logs engine's own word for an aura on a friendly
// unit, upper case as summary.Accumulator writes it. Its counterpart,
// "DEBUFF", has no source in an engine result: sim/core reports aura
// metrics for the player only.
const auraTypeBuff = "BUFF"

// EmptySummary is the summary of no fight: every list present and
// empty, never absent. It is what a run that folded nothing still has
// to carry.
//
// JSON has two ways to say "no rows" and only one of them is the shape
// the web parses. A nil Go slice marshals as `null`, and an aborted
// result used to be built from a zero summary.Summary, so `damage_done`
// and fifteen other keys came back null where a finished result has
// `[]`. Every consumer then needs a null check per key, on the one code
// path - the user pressed Stop - least likely to be exercised.
//
// Summarize builds on this too, so the finished shape and the empty one
// cannot drift apart. TestEverySummaryListIsEmptyNotNull walks the
// struct by reflection, so a list field added to summary.Summary later
// fails here rather than reaching the web as a null.
func EmptySummary() summary.Summary {
	return summary.Summary{
		// The engine that RAN it, from the pin compiled into this
		// binary - not req.EngineVersion, which is the client's claim.
		// A cached request naming an old sha, re-run by a new build,
		// used to come back stamped with the old one, and every stored
		// row's provenance was hearsay.
		EngineVersion: "sim:" + enginever.Version,

		DamageDone:   []summary.Actor{},
		DamageTaken:  []summary.Actor{},
		Healing:      []summary.Actor{},
		HealingTaken: []summary.Actor{},

		Deaths:     []summary.Death{},
		Auras:      []summary.AuraTrack{},
		Casts:      []summary.CastRow{},
		Interrupts: []summary.ExchangeRow{},
		Dispels:    []summary.ExchangeRow{},
		Resources:  []summary.ResourceTrack{},

		// A sim has no threat model attached, no taunt, no curated
		// mechanics table for a target dummy, and no encounter phases.
		// All five are present and empty, which is exactly what the logs
		// engine produces for a fight whose encounter has no table.
		Threat:         []summary.ThreatRow{},
		ThreatByTarget: []summary.ThreatPair{},
		Taunts:         []summary.Taunt{},
		Combatants:     []summary.CombatantRow{},
		Roster:         []summary.RosterRow{},
		Mechanics:      summary.MechanicsBlock{TableFound: false, Rows: []summary.MechanicRow{}},
		Phases:         []summary.Phase{},
	}
}

// Summarize maps an engine result onto the logs engine's Summary, per the
// interface contract's mapping table.
func Summarize(res *proto.RaidSimResult, req api.SimRequest) (summary.Summary, error) {
	if res == nil {
		return summary.Summary{}, fmt.Errorf("%w: nil result", ErrSimFailed)
	}
	if err := ResultError(res); err != nil {
		return summary.Summary{}, err
	}
	iters := float64(res.IterationsDone)
	if iters <= 0 {
		return summary.Summary{}, fmt.Errorf("%w: iterations_done is %d", ErrSimFailed, res.IterationsDone)
	}
	player, err := PlayerMetrics(res)
	if err != nil {
		return summary.Summary{}, err
	}

	durationMS := int64(math.Round(res.AvgIterationDuration * 1000))
	class, spec := splitSpecSlug(req.Spec)

	// Start from the empty shape and fill in what this run measured, so
	// the four lists a sim can populate are the ONLY difference between
	// a finished summary and the one an aborted run carries. Building
	// the two shapes separately is how they came to disagree.
	out := EmptySummary()
	out.FightIndex = 1
	out.DurationMS = durationMS
	out.DamageDone = actors(player, class, iters, durationMS)
	out.Auras = auras(player)
	out.Casts = casts(player, iters)
	out.Resources = resources(player, iters)
	out.Roster = roster(player, class, spec, out, durationMS)
	if err := checkRowIdentity(out); err != nil {
		return summary.Summary{}, err
	}
	return out, nil
}

// ResultError turns the engine's ErrorOutcome into one of ours, and
// reports nil when the run completed.
//
// The engine has two kinds of unhappy ending and they need different
// words in the UI: an abort is what the Stop button does, and a failure
// is a bug or a bad request. Only the second carries a message, which
// is why every caller must switch on Type rather than test the message
// for emptiness.
func ResultError(res *proto.RaidSimResult) error {
	if res == nil {
		return fmt.Errorf("%w: nil result", ErrSimFailed)
	}
	if res.Error == nil {
		return nil
	}
	if res.Error.Type == proto.ErrorOutcomeType_ErrorOutcomeAborted {
		return ErrAborted
	}
	if res.Error.Message == "" {
		return fmt.Errorf("%w: no message", ErrSimFailed)
	}
	return fmt.Errorf("%w: %s", ErrSimFailed, res.Error.Message)
}

// PlayerMetrics returns the first player of the first party, which is the
// only player an individual sim has.
func PlayerMetrics(res *proto.RaidSimResult) (*proto.UnitMetrics, error) {
	if res == nil || res.RaidMetrics == nil {
		return nil, ErrNoPlayer
	}
	for _, party := range res.RaidMetrics.Parties {
		for _, p := range party.Players {
			if p != nil {
				return p, nil
			}
		}
	}
	return nil, ErrNoPlayer
}

// DPS lifts the engine's distribution into the envelope's Estimate. Error
// is the standard error of the mean: stdev over the square root of the
// iteration count, which is the figure the sim page shows beside the DPS.
func DPS(res *proto.RaidSimResult) api.Estimate {
	if res == nil || res.RaidMetrics == nil || res.RaidMetrics.Dps == nil {
		return api.Estimate{}
	}
	d := res.RaidMetrics.Dps
	est := api.Estimate{Mean: d.Avg, StdDev: d.Stdev, Min: d.Min, Max: d.Max}
	if res.IterationsDone > 0 {
		est.Error = d.Stdev / math.Sqrt(float64(res.IterationsDone))
	}
	return est
}

// petGUID names a pet's row. It is derived from the owner's guid so the
// Casts tab's owner grouping and the damage table agree.
func petGUID(i int) string { return fmt.Sprintf("%s-pet-%d", playerGUID, i) }

// actors builds the damage table: one row for the player, then one per pet.
func actors(player *proto.UnitMetrics, class string, iters float64, durationMS int64) []summary.Actor {
	out := []summary.Actor{actorFrom(player, playerGUID, class, iters, durationMS)}
	for i, pet := range player.Pets {
		// A pet has no class of its own in the roster's sense; the
		// report colours it by its owner's.
		out = append(out, actorFrom(pet, petGUID(i), class, iters, durationMS))
	}
	return out
}

func actorFrom(u *proto.UnitMetrics, guid, class string, iters float64, durationMS int64) summary.Actor {
	a := summary.Actor{
		GUID:      guid,
		Name:      u.Name,
		Class:     class,
		ActiveMS:  durationMS,
		Abilities: []summary.Ability{},
		Targets:   []summary.Pair{},
		Series:    []int64{},
	}
	perTarget := map[int32]int64{}
	for _, am := range u.Actions {
		ab := ability(am, iters, perTarget)
		a.Abilities = append(a.Abilities, ab)
		a.Total += ab.Total
	}
	a.Abilities = foldAbilities(a.Abilities)
	a.Effective = a.Total

	idx := make([]int32, 0, len(perTarget))
	for k := range perTarget {
		idx = append(idx, k)
	}
	sort.Slice(idx, func(i, j int) bool { return idx[i] < idx[j] })
	for _, i := range idx {
		a.Targets = append(a.Targets, summary.Pair{
			GUID:  fmt.Sprintf("sim-target-%d", i),
			Name:  fmt.Sprintf("Target %d", i),
			Total: perTarget[i],
		})
	}
	sort.SliceStable(a.Abilities, func(i, j int) bool { return a.Abilities[i].Total > a.Abilities[j].Total })
	return a
}

// ability folds one ActionMetrics, which is already summed over every
// target and every iteration, into one summary row per fight.
//
// TargetedActionMetrics.damage is every outcome's damage already. Its
// crit, tick, glance, crush, block and resisted figures are breakdowns
// of that one total, not additions to it - sim/core books a crit in
// both damage and crit_damage, and its own DPS total reads damage alone
// (metrics_aggregator.go) - so adding them would count a crit twice and
// a partially resisted crit four times. The same holds for the outcome
// counters: resisted_hits is the subset of hits that partly resisted.
func ability(am *proto.ActionMetrics, iters float64, perTarget map[int32]int64) summary.Ability {
	spellID, name := ActionName(am.Id)
	ab := summary.Ability{
		SpellID: spellID,
		Name:    name,
		School:  int64(am.SpellSchool),
		Misses:  map[string]int64{},
	}
	for _, t := range am.Targets {
		dmg := per(t.Damage, iters)
		ab.Total += dmg
		ab.Effective += dmg
		perTarget[t.UnitIndex] += dmg

		// Resisted and Blocked stay zero. The summary means the damage a
		// resist or a block took away, and the engine tracks neither:
		// resisted_damage is the damage a partially resisted cast still
		// dealt - "Partial or full resists aren't tracked, at the
		// moment", says metrics_aggregator.go - and block_damage is the
		// damage a blocked swing still dealt. Either one in these fields
		// would print a number meaning the opposite of its column.
		// Hits, Crits and Ticks are the LOGS ENGINE's definitions, not
		// the engine's, because the report's components compute crit %
		// as crits/hits and a sim and a real fight of the same shape
		// have to print the same number.
		//
		// The logs engine (summary/damage.go's fold) increments Hits for
		// every non-periodic landing whatever its outcome, Ticks for
		// every periodic one, and Crits IN ADDITION whenever the landing
		// was critical - so Hits is inclusive of crits and Crits spans
		// ticks. sim/core's counters are disjoint instead: a crit never
		// increments Hits, and Glances, Crushes, Blocks and BlockedCrits
		// are each counted separately again. Copying them across as they
		// stand prints a glancing-heavy warrior as having almost no
		// hits, a dot as never critting, and a crit rate computed
		// against the wrong denominator.
		//
		// Misses, dodges and parries are NOT landings and stay out of
		// Hits; they are in the Misses map below, which is where the
		// logs engine puts them too.
		ab.Hits += per(float64(t.Hits+t.Crits+t.Glances+t.Crushes+t.Blocks+t.BlockedCrits), iters)
		ab.Crits += per(float64(t.Crits+t.BlockedCrits+t.CritTicks), iters)
		ab.Ticks += per(float64(t.Ticks+t.CritTicks), iters)
		// The keys are the combat log's own, so a sim and a real fight
		// aggregate together: MISS, DODGE, PARRY and BLOCK are its
		// MissType strings verbatim (logs/engine/summary/damage.go
		// writes e.MissType), and GLANCING and CRUSHING are its flag
		// names, which the engine reports as counted outcomes where a
		// log carries them on a landed hit.
		//
		// A glance, a crush and a partial block therefore appear here
		// AND in Hits above, which is what the report needs to print
		// "58 glancing of 134 swings": they are landings that happened
		// a particular way, not things that failed to land.
		addMiss(ab.Misses, "MISS", t.Misses, iters)
		addMiss(ab.Misses, "DODGE", t.Dodges, iters)
		addMiss(ab.Misses, "PARRY", t.Parries, iters)
		addMiss(ab.Misses, "BLOCK", t.Blocks+t.BlockedCrits, iters)
		addMiss(ab.Misses, "GLANCING", t.Glances, iters)
		addMiss(ab.Misses, "CRUSHING", t.Crushes, iters)
	}
	if len(ab.Misses) == 0 {
		ab.Misses = nil
	}
	// The engine reports no per-hit minimum or maximum, only per-action
	// totals, so a range would be invented. Both stay zero.
	return ab
}

// addMiss records an outcome the fight saw. A count that divides to
// zero writes no key at all: "0 dodge" is not something a report row
// should say, and a real fight's summary only carries the outcomes it
// actually had.
func addMiss(m map[string]int64, key string, count int32, iters float64) {
	if count == 0 {
		return
	}
	if n := per(float64(count), iters); n > 0 {
		m[key] += n
	}
}

// per divides an across-iterations total into a per-fight figure. It
// rounds rather than truncating, so an ability used in most iterations
// does not report zero, and it never produces a negative from a spend
// figure: callers negate first where that matters.
func per(total, iters float64) int64 {
	return int64(math.Round(total / iters))
}

func auras(u *proto.UnitMetrics) []summary.AuraTrack {
	out := make([]summary.AuraTrack, 0, len(u.Auras))
	for _, am := range u.Auras {
		spellID, name := ActionName(am.Id)
		out = append(out, summary.AuraTrack{
			TargetGUID: playerGUID,
			TargetName: u.Name,
			SpellID:    spellID,
			Name:       name,
			// The vocabulary is the logs engine's, upper case: its
			// accumulator writes "BUFF" and "DEBUFF" (summary.go's
			// seedAuras and applyAura), and the report's shared
			// AuraTable filters on exactly those two strings. A sim
			// tracks only the player's own auras - the engine's
			// UnitMetrics carries no aura row for anything on the
			// target - so every row here is a BUFF, and the constant
			// is named rather than spelled inline so the day a
			// debuff appears there is one place to change.
			Type: auraTypeBuff,
			// AuraMetrics already reports per-iteration averages, so
			// these are not divided again.
			Applications: int64(math.Round(am.ProcsAvg)),
			MaxStacks:    0,
			UptimeMS:     int64(math.Round(am.UptimeSecondsAvg * 1000)),
			// The engine reports no application timeline, so there are no
			// segments to build and none are invented.
			Segments: []summary.Segment{},
			Appliers: []string{u.Name},
		})
	}
	out = foldAuras(out)
	sort.SliceStable(out, func(i, j int) bool { return out[i].UptimeMS > out[j].UptimeMS })
	return out
}

// casts builds one row per caster per spell. A pet's row stays the pet's -
// the Casts tab prints "via <pet>" - but its OwnerGUID is the player's, so
// the tab groups it under the player the way it does in a real fight.
func casts(u *proto.UnitMetrics, iters float64) []summary.CastRow {
	out := castsFor(u, playerGUID, playerGUID, iters, nil)
	for i, pet := range u.Pets {
		out = castsFor(pet, petGUID(i), playerGUID, iters, out)
	}
	out = foldCasts(out)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Succeeded > out[j].Succeeded })
	return out
}

func castsFor(u *proto.UnitMetrics, guid, owner string, iters float64, out []summary.CastRow) []summary.CastRow {
	if out == nil {
		out = make([]summary.CastRow, 0, len(u.Actions))
	}
	for _, am := range u.Actions {
		var total int32
		for _, t := range am.Targets {
			total += t.Casts
		}
		if total == 0 {
			continue
		}
		spellID, name := ActionName(am.Id)
		n := per(float64(total), iters)
		out = append(out, summary.CastRow{
			GUID:      guid,
			Name:      u.Name,
			OwnerGUID: owner,
			SpellID:   spellID,
			SpellName: name,
			Started:   n,
			Succeeded: n,
			Failed:    0,
			// A sim has no failed casts and no per-cast timing, so
			// FailReasons stays nil (it is omitempty) and CastTimeMS zero.
			CastTimeMS: 0,
			// The engine reports no cast timestamps, so the sequence the
			// report's cast timeline draws stays empty for a sim.
			Sequence: []int64{},
		})
	}
	return out
}

func resources(u *proto.UnitMetrics, iters float64) []summary.ResourceTrack {
	// The engine reports one ResourceMetrics per action per resource type;
	// the summary wants one track per resource type.
	byType := map[proto.ResourceType]*summary.ResourceTrack{}
	order := []proto.ResourceType{}
	for _, rm := range u.Resources {
		tr, ok := byType[rm.Type]
		if !ok {
			tr = &summary.ResourceTrack{
				GUID:      playerGUID,
				Name:      u.Name,
				PowerType: int64(rm.Type),
				Series:    []int64{},
				// The engine reports no per-second reading and no cap, so
				// Max, AtMaxMS and ZeroMS stay zero and the resource graph
				// draws no cap line for a sim.
			}
			byType[rm.Type] = tr
			order = append(order, rm.Type)
		}
		// Gain is negative for a spend, per the proto's own comment.
		if rm.ActualGain >= 0 {
			tr.Gained += per(rm.ActualGain, iters)
		} else {
			tr.Spent += per(-rm.ActualGain, iters)
		}
		// gain minus actual_gain is the engine's own over-cap waste, which
		// is exactly what ResourceTrack.Wasted means.
		if w := rm.Gain - rm.ActualGain; w > 0 {
			tr.Wasted += per(w, iters)
		}
	}
	out := make([]summary.ResourceTrack, 0, len(order))
	for _, t := range order {
		out = append(out, *byType[t])
	}
	return out
}

func roster(u *proto.UnitMetrics, class, spec string, s summary.Summary, durationMS int64) []summary.RosterRow {
	var damage int64
	for _, a := range s.DamageDone {
		damage += a.Total
	}
	var dps float64
	if u.Dps != nil {
		dps = u.Dps.Avg
	}
	return []summary.RosterRow{{
		GUID:        playerGUID,
		Name:        u.Name,
		Class:       class,
		ClassSource: "sim",
		Spec:        spec,
		Role:        "dps",
		ActiveMS:    durationMS,
		ActivityPct: 100,
		DamageDone:  damage,
		DPS:         dps,
	}}
}

// splitSpecSlug turns "warrior-fury" into ("warrior", "fury").
//
// The canonical pairing is sim/specs, generated from
// data/curated/specs.json, and it is consulted first so a hyphenated
// spec slug such as "hunter-beast-mastery" splits where the data lane
// says it does rather than at the first hyphen. The fallback keeps the
// function total for a slug the list has not got: the adapter's job is
// to render a result, not to police one, and sim/request already
// refused the request.
func splitSpecSlug(slug string) (class, spec string) {
	if known, ok := specs.ByKey[slug]; ok {
		return known.ClassSlug, known.SpecSlug
	}
	i := strings.Index(slug, "-")
	if i < 0 {
		return slug, ""
	}
	return slug[:i], slug[i+1:]
}

// Sample is one iteration's casts, in order: the median-DPS iteration
// the engine records when SimOptions.sample_iteration is set.
//
// Every row carries the summary's ACTION KEY - "spell:23881",
// "item:13503", "other:melee" - and no display name. That is contract
// A12, and the reason is that the page already resolves a cast row's
// name from the build's spells.json through resolveActionName; a name
// baked in here would be a second vocabulary for the same action, and
// the sample table and the cast table would disagree about what to
// call a tagged or ranked spell.
//
// The engine's own at_ms is used as it stands: it is already
// milliseconds (contract 10.3), and rounding a seconds figure here
// used to be one more place the timeline and this table could
// disagree about when something happened.
//
// No sample is not an error. A bulk stage does not ask for one, an
// aborted run has none, and neither does a result from an engine
// older than the field; the page renders the card only when there are
// rows.
func Sample(res *proto.RaidSimResult) []api.SampleCast {
	casts := res.GetSampleIteration().GetCasts()
	if len(casts) == 0 {
		return nil
	}
	out := make([]api.SampleCast, 0, len(casts))
	for _, c := range casts {
		_, action := ActionName(c.GetActionId())
		row := api.SampleCast{
			AtMS:   c.GetAtMs(),
			Action: action,
			Target: c.GetTarget(),
		}
		if len(c.GetResources()) > 0 {
			row.Resources = make(map[string]int, len(c.GetResources()))
			for k, v := range c.GetResources() {
				row.Resources[k] = int(v)
			}
		}
		out = append(out, row)
	}
	return out
}

// ErrNoWeights is returned when a stat weights result carries no DPS
// weight block. There is nothing to report and nothing to normalise.
var ErrNoWeights = errors.New("adapter: the result carries no stat weights")

// Weights lifts the engine's stat weights into the envelope's named
// rows.
//
// The engine answers with an array indexed by proto.Stat, which is a
// hundred-odd slots of mostly zero; the envelope answers with the
// stats the request asked for, in the order it asked, normalised so
// the reference stat is exactly 1. Normalising here rather than on the
// page is what makes the "copy for Pawn" string and the table the same
// numbers.
//
// The engine already normalises when asked, through EpReferenceStat,
// but it reports both the raw weights and the EP values and the two
// are easy to confuse; taking the raw weights and dividing is one
// arithmetic, in one place, that cannot pick the wrong block.
func Weights(res *proto.StatWeightsResult, req api.SimRequest) ([]api.StatWeight, error) {
	if req.Weights == nil {
		return nil, fmt.Errorf("%w: the request asked for none", ErrNoWeights)
	}
	if res == nil {
		return nil, fmt.Errorf("%w: nil result", ErrNoWeights)
	}
	if res.Error != nil && res.Error.Message != "" {
		return nil, fmt.Errorf("%w: %s", ErrSimFailed, res.Error.Message)
	}
	values := res.GetDps()
	if values.GetWeights() == nil {
		return nil, ErrNoWeights
	}
	raw := values.GetWeights().GetStats()
	stdev := values.GetWeightsStdev().GetStats()
	at := func(s proto.Stat, from []float64) float64 {
		if int(s) >= len(from) {
			return 0
		}
		return from[s]
	}

	reference, ok := statid.Parse(req.Weights.Reference)
	if !ok {
		return nil, fmt.Errorf("%w: reference %q", ErrNoWeights, req.Weights.Reference)
	}
	scale := at(reference, raw)
	if scale == 0 {
		return nil, fmt.Errorf("%w: the reference stat %q weighs nothing, so nothing can be normalised against it", ErrNoWeights, req.Weights.Reference)
	}

	out := make([]api.StatWeight, 0, len(req.Weights.Stats))
	for _, id := range req.Weights.Stats {
		s, ok := statid.Parse(id)
		if !ok {
			return nil, fmt.Errorf("%w: %q", ErrNoWeights, id)
		}
		out = append(out, api.StatWeight{
			Stat:   id,
			Weight: at(s, raw) / scale,
			Error:  at(s, stdev) / scale,
		})
	}
	return out, nil
}
