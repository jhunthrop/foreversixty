package coach

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// coreCostMark is the word in CoreNoCost that another resource's noun replaces.
var coreCostMark = PowerNames[PowerRage]

// coreGap is one stretch where an ability sat ready and uncast: [fromMS, toMS) is the whole
// gap, readyMS the counted ready time left after the grace.
type coreGap struct{ fromMS, toMS, readyMS, lost int64 }

// coreAbility is one assessed core ability: its casts, its counted gaps and its
// comparison.
type coreAbility struct {
	family     string
	spellID    int64
	spell      SpellConst
	actual     int64
	lost       int64
	gaps       []coreGap // counted gaps (at least one lost cast), in time order
	unverified bool      // the resource could not be checked
	youCPM     float64
	refCPM     float64
	cmp        *Basis
	thin       string
}

func (a coreAbility) possible() int64 { return a.actual + a.lost }

func (a coreAbility) lostPct() float64 {
	if a.possible() == 0 {
		return 0
	}
	return float64(a.lost) / float64(a.possible()) * 100
}

func (a coreAbility) fires() bool {
	return a.lost >= CoreFiresLost && a.lostPct() >= CoreFiresPct
}

func (a coreAbility) resource() string {
	if name, ok := PowerNames[a.spell.CostType]; ok {
		return name
	}
	return coreUnknownResource
}

// resourceCheck is how a gap's resource is verified: against the power track's series at
// the spell's base cost, or not at all.
type resourceCheck struct {
	series     []int64
	cost       int64
	on         bool
	unverified bool
}

// coreCasts is F1: core cooldown abilities cast fewer times than the cooldown allowed
// (§6.2). A healer is assessed on efficiency cooldowns only (§6.3).
func coreCasts(in Input, c Counted) outcome {
	entries := coreEntries(in)
	if len(entries) == 0 { // an omitted or empty list is unknown, never "nothing to check" (§6.1)
		return outcome{CouldNotCheck: []string{CannotCoreList}}
	}
	var abilities []coreAbility
	for _, e := range entries {
		if a, ok := assessAbility(in, c, e); ok {
			abilities = append(abilities, a)
		}
	}
	sort.SliceStable(abilities, func(i, j int) bool { return worseThan(abilities[i], abilities[j]) })
	return coreOutcome(in, abilities)
}

func coreEntries(in Input) []RuleEntry {
	switch in.Row.Role {
	case RoleDPS, RoleTank:
		return in.Tables.Rules.Core
	case RoleHealer:
		return in.Tables.Rules.Efficiency
	}
	return nil
}

func worseThan(a, b coreAbility) bool {
	if a.lost != b.lost {
		return a.lost > b.lost
	}
	if a.lostPct() != b.lostPct() {
		return a.lostPct() > b.lostPct()
	}
	return a.family < b.family
}

func assessAbility(in Input, c Counted, e RuleEntry) (coreAbility, bool) {
	guid := in.Row.GUID
	seq, counts := castsOf(in.Fight, guid, e.SpellIDs)
	if e.Talent && len(seq) == 0 {
		return coreAbility{}, false
	}
	id, sc, ok := pickSpell(e.SpellIDs, in.Tables.Spells, counts)
	if !ok || sc.CooldownMS < CoreCooldownMinMS || sc.CooldownMS > CoreCooldownMaxMS {
		return coreAbility{}, false
	}
	res := newResourceCheck(in.Fight, guid, sc)
	a := coreAbility{family: e.Family, spellID: id, spell: sc, actual: int64(len(seq)), unverified: res.unverified}
	a.gaps = coreGaps(seq, sc.CooldownMS, c, res)
	for _, g := range a.gaps {
		a.lost += g.lost
	}
	a.compare(in, c)
	return a, true
}

// castsOf is the sorted successful cast instants of the spell ids by guid and by anything
// guid owns, and the count per spell id.
func castsOf(fight summary.Summary, guid string, ids []int64) ([]int64, map[int64]int) {
	var seq []int64
	counts := map[int64]int{}
	for _, row := range fight.Casts {
		if (row.GUID != guid && row.OwnerGUID != guid) || !slices.Contains(ids, row.SpellID) {
			continue
		}
		seq = append(seq, row.Sequence...)
		counts[row.SpellID] += len(row.Sequence)
	}
	slices.Sort(seq)
	return seq, counts
}

// pickSpell is the family's spell the player cast most (the higher rank on a tie) that has
// spell constants.
func pickSpell(ids []int64, spells map[int64]SpellConst, counts map[int64]int) (int64, SpellConst, bool) {
	var (
		bestID   int64
		best     SpellConst
		bestN    = -1
		resolved bool
	)
	for _, id := range ids {
		sc, ok := spells[id]
		if ok && counts[id] >= bestN {
			bestID, best, bestN, resolved = id, sc, counts[id], true
		}
	}
	return bestID, best, resolved
}

func newResourceCheck(fight summary.Summary, guid string, sc SpellConst) resourceCheck {
	needed := !sc.CostKnown || sc.Cost > 0
	if !needed {
		return resourceCheck{}
	}
	if sc.CostKnown {
		for _, t := range fight.Resources {
			if t.GUID == guid && t.PowerType == sc.CostType {
				return resourceCheck{series: t.Series, cost: sc.Cost, on: true}
			}
		}
	}
	return resourceCheck{unverified: true}
}

// coreGaps walks the casts: after each cast the ability is ready again one cooldown
// later, until the next cast (or the end of the player's life).
func coreGaps(seq []int64, cooldownMS int64, c Counted, res resourceCheck) []coreGap {
	var gaps []coreGap
	from := int64(0)
	for _, next := range append(slices.Clone(seq), c.AliveMS) {
		if g, ok := coreGapAt(from, min(next, c.AliveMS), cooldownMS, c, res); ok {
			gaps = append(gaps, g)
		}
		from = next + cooldownMS
	}
	return gaps
}

func coreGapAt(from, to, cooldownMS int64, c Counted, res resourceCheck) (coreGap, bool) {
	if to <= from {
		return coreGap{}, false
	}
	remaining := readyMSInGap(from, to, c, res) - CoreGapGraceMS
	lost := remaining / cooldownMS
	if remaining <= 0 || lost < 1 {
		return coreGap{}, false
	}
	return coreGap{fromMS: from, toMS: to, readyMS: remaining, lost: lost}, true
}

// readyMSInGap is the time in [from, to) the ability could have been cast: outside the
// excluded windows and, when the resource is checked, in a whole second with at least the
// base cost in the bar.
func readyMSInGap(from, to int64, c Counted, res resourceCheck) int64 {
	var total int64
	for _, w := range freeWindows(from, to, c.Excluded) {
		if res.on {
			total += resourceMS(w, res)
		} else {
			total += w.ToMS - w.FromMS
		}
	}
	return total
}

// freeWindows is [from, to) less the excluded windows (in time order, not overlapping).
func freeWindows(from, to int64, excluded []ratingengine.Window) []ratingengine.Window {
	var free []ratingengine.Window
	cursor := from
	for _, w := range excluded {
		if w.ToMS <= cursor {
			continue
		}
		if w.FromMS >= to {
			break
		}
		if w.FromMS > cursor {
			free = append(free, ratingengine.Window{FromMS: cursor, ToMS: w.FromMS})
		}
		cursor = max(cursor, w.ToMS)
	}
	if cursor < to {
		free = append(free, ratingengine.Window{FromMS: cursor, ToMS: to})
	}
	return free
}

// resourceMS is the part of w that falls in seconds where the bar held the cost.
func resourceMS(w ratingengine.Window, res resourceCheck) int64 {
	var total int64
	for sec := w.FromMS / ResourceWholeSecondMS; sec*ResourceWholeSecondMS < w.ToMS; sec++ {
		if hadResourceForOneSecond(res.series, sec, res.cost) {
			total += min(w.ToMS, (sec+1)*ResourceWholeSecondMS) - max(w.FromMS, sec*ResourceWholeSecondMS)
		}
	}
	return total
}

// hadResourceForOneSecond reports whether the reading for whole second sec was at or above
// cost. A second the series does not cover has no reading and does not count.
func hadResourceForOneSecond(series []int64, sec, cost int64) bool {
	return sec >= 0 && sec < int64(len(series)) && series[sec] >= cost
}

// compare fills the ability's casts per minute and its comparison (§11), skipping the sim
// for a WorstActions family.
func (a *coreAbility) compare(in Input, c Counted) {
	if c.CountedMS <= 0 {
		return
	}
	a.youCPM = float64(a.actual) / (float64(c.CountedMS) / float64(ratingMinuteMS))
	pool := func(p Profile) (float64, bool) { return p.CastsPerMinOf(a.family) }
	sim := func(s SimRef) (float64, bool) { v, ok := s.CastsPerMin[a.family]; return v, ok }
	a.refCPM, a.cmp, a.thin = comparison(in, pool, sim, !in.Tables.WorstActions[a.family])
}

const ratingMinuteMS = 60_000

// coreAnchors collects an outcome's local anchors and hands out their "#n" references.
type coreAnchors []Anchor

func (l *coreAnchors) add(a Anchor) string {
	*l = append(*l, a)
	return "#" + strconv.Itoa(len(*l)-1)
}

func coreOutcome(in Input, abilities []coreAbility) outcome {
	var o outcome
	var anchors coreAnchors
	var fired []coreAbility
	for _, a := range abilities {
		if a.fires() {
			fired = append(fired, a)
			continue
		}
		if s, ok := coreStrength(a, &anchors); ok {
			o.Strengths = append(o.Strengths, s)
		}
	}
	if len(fired) > 0 {
		o.Finding = coreFinding(in, fired, &anchors)
		o.Windows = coreWindows(fired[0])
		if fired[0].thin != "" {
			o.CouldNotCheck = append(o.CouldNotCheck, fired[0].thin)
		}
	}
	o.Anchors = anchors
	return o
}

func coreWindows(a coreAbility) []ratingengine.Window {
	ws := make([]ratingengine.Window, len(a.gaps))
	for i, g := range a.gaps {
		ws[i] = ratingengine.Window{FromMS: g.fromMS, ToMS: g.toMS}
	}
	return ws
}

// coreFinding is the finding for the worst of the fired abilities (worst first).
func coreFinding(in Input, fired []coreAbility, anchors *coreAnchors) *Finding {
	worst := fired[0]
	shown := longestGaps(worst.gaps, 2)
	f := &Finding{
		Code:     CodeCoreCasts,
		Severity: SeverityMinor,
		Title:    Fill(TitleCoreCasts, map[string]string{"Spell": worst.spell.Name, "n": strconv.FormatInt(worst.lost, 10)}),
		Detail:   coreDetail(worst, shown, fired[1:min(len(fired), 1+MoreAbilitiesInDetail)]),
		Basis: Basis{Kind: BasisSpellData, Label: Fill(BasisCooldown, map[string]string{
			"Spell": worst.spell.Name, "n": strconv.FormatFloat(float64(worst.spell.CooldownMS)/1000, 'f', -1, 64), "resource": worst.resource(),
		})},
		Facts: coreFacts(worst, shown, fired[1:min(len(fired), 1+MoreAbilitiesInDetail)]),
	}
	if worst.lostPct() >= CoreMajorPct {
		f.Severity = SeverityMajor
	}
	f.Anchors = append(f.Anchors, anchors.add(Anchor{Kind: AnchorCast, SpellID: worst.spellID}))
	for _, g := range shown {
		f.Anchors = append(f.Anchors, anchors.add(Anchor{Kind: AnchorTime, AtMS: ptr(g.fromMS), ToMS: ptr(g.toMS)}))
	}
	if worst.cmp != nil {
		cmp := *worst.cmp
		cmp.Label += Fill(BasisPerMin, map[string]string{"ref": PerMinute(worst.refCPM), "you": PerMinute(worst.youCPM)})
		f.Comparison = &cmp
	}
	withCeiling(f, in, worst)
	return f
}

func ptr(v int64) *int64 { return &v }

// longestGaps is the n gaps with the most ready time, longest first, earliest on a tie.
func longestGaps(gaps []coreGap, n int) []coreGap {
	sorted := slices.Clone(gaps)
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].readyMS > sorted[j].readyMS })
	return sorted[:min(n, len(sorted))]
}

func coreDetail(worst coreAbility, shown []coreGap, more []coreAbility) string {
	fields := map[string]string{
		"lost": strconv.FormatInt(worst.lost, 10), "possible": strconv.FormatInt(worst.possible(), 10),
		"resource": worst.resource(), "gap1": Span(shown[0].fromMS, shown[0].toMS),
	}
	format := DetailCoreCastsOne
	if len(shown) > 1 {
		format, fields["gap2"] = DetailCoreCasts, Span(shown[1].fromMS, shown[1].toMS)
	}
	parts := []string{Fill(format, fields)}
	if len(more) > 0 {
		parts = append(parts, Fill(DetailCoreMore, map[string]string{"more": moreList(more)}))
	}
	if worst.unverified {
		parts = append(parts, noCostLine(worst.resource()))
	}
	return strings.Join(parts, " ")
}

func moreList(more []coreAbility) string {
	items := make([]string, len(more))
	for i, a := range more {
		items[i] = Fill(coreMoreEntry, map[string]string{"Spell": a.spell.Name, "n": strconv.FormatInt(a.lost, 10)})
	}
	return joinAnd(items)
}

// noCostLine is CoreNoCost for the ability's resource (the spec words it for rage).
func noCostLine(resource string) string {
	return strings.Replace(CoreNoCost, coreCostMark, resource, 1)
}

func coreNumber(v float64, text, unit, of string) NumberFact {
	return NumberFact{V: v, Text: text, Unit: unit, Of: of}
}

func castCount(n int64, of string) NumberFact {
	return coreNumber(float64(n), strconv.FormatInt(n, 10), "casts", of)
}

func spellFact(a coreAbility) SpellFact {
	return SpellFact{ID: a.spellID, Name: a.spell.Name, Rank: a.spell.Rank}
}

func coreFacts(worst coreAbility, shown []coreGap, more []coreAbility) Facts {
	f := Facts{
		Numbers: []NumberFact{castCount(worst.actual, worst.spell.Name), castCount(worst.possible(), worst.spell.Name), castCount(worst.lost, worst.spell.Name)},
		Spells:  []SpellFact{spellFact(worst)},
	}
	if worst.cmp != nil {
		f.Numbers = append(f.Numbers,
			coreNumber(worst.refCPM, PerMinute(worst.refCPM), "casts per minute", worst.spell.Name),
			coreNumber(worst.youCPM, PerMinute(worst.youCPM), "casts per minute", worst.spell.Name))
	}
	for _, a := range more {
		f.Numbers = append(f.Numbers, castCount(a.lost, a.spell.Name))
		f.Spells = append(f.Spells, spellFact(a))
	}
	for _, g := range shown {
		f.Times = append(f.Times, Span(g.fromMS, g.toMS))
	}
	return f
}

// coreStrength is a strength when the ability is cast at or above the comparison median.
func coreStrength(a coreAbility, anchors *coreAnchors) (Strength, bool) {
	if a.cmp == nil || a.youCPM < a.refCPM {
		return Strength{}, false
	}
	format := StrengthCasts
	if a.cmp.Kind == BasisSim {
		format = StrengthCastsSim
	}
	return Strength{
		Sentence: Fill(format, map[string]string{
			"Spell": a.spell.Name, "you": PerMinute(a.youCPM), "n": strconv.Itoa(a.cmp.N), "ref": PerMinute(a.refCPM),
		}),
		Facts: Facts{
			Numbers: []NumberFact{
				coreNumber(a.youCPM, PerMinute(a.youCPM), "casts per minute", a.spell.Name),
				coreNumber(a.refCPM, PerMinute(a.refCPM), "casts per minute in the comparison", a.spell.Name),
			},
			Spells: []SpellFact{spellFact(a)},
		},
		Anchors: []string{anchors.add(Anchor{Kind: AnchorCast, SpellID: a.spellID})},
	}, true
}
