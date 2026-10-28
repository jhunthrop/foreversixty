package coach

import (
	"cmp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// uptimeKindsByRole is the aura kinds F2 judges for a role (§6.2).
var uptimeKindsByRole = map[string][]string{
	RoleDPS:    {AuraDot, AuraDebuff, AuraSelfBuff},
	RoleHealer: {AuraSelfBuff},
	RoleTank:   {AuraDefensive},
}

// auraReading is one owned aura's uptime against its reference.
type auraReading struct {
	rule  AuraRule
	upMS  int64
	denMS int64
	pct   float64
	ref   float64
	basis Basis
	// anchor is where the reading opens in the report: its aura row, or the counted time
	// when the player never had the aura.
	anchor Anchor
}

func (r auraReading) gap() float64 { return r.ref - r.pct }

// uptime is F2: owned auras' uptime from first application to the end, less time dead
// and downtime, against the top logs, else the sim (validated DPS), else "not enough
// logs yet" (§6.2).
func uptime(in Input, c Counted) outcome {
	var o outcome
	var low []auraReading
	rules := slices.DeleteFunc(slices.Clone(in.Tables.Rules.Auras), func(r AuraRule) bool {
		return !slices.Contains(uptimeKindsByRole[in.Row.Role], r.Kind)
	})
	if len(rules) == 0 { // an omitted or empty list is unknown, never "nothing to check" (§6.1)
		return outcome{CouldNotCheck: []string{CannotAuraList}}
	}
	for _, rule := range rules {
		r, ok := judgeAura(in, c, rule, &o)
		switch {
		case !ok:
		case r.gap() >= UptimeFiresPoints && r.pct < UptimeFiresUnderPct:
			low = append(low, r)
		case r.pct >= r.ref && r.basis.Kind == BasisTopLogs:
			addUptimeStrength(in, r, &o)
		}
	}
	if len(low) > 0 {
		o.Finding = uptimeFinding(in, low, &o)
	}
	return o
}

// judgeAura measures one aura; ok is false when nothing is to be compared, in which case
// any "Couldn't check" line has been added to o.
func judgeAura(in Input, c Counted, rule AuraRule, o *outcome) (auraReading, bool) {
	track, appliers := ownTrack(in, rule)
	if appliers > 1 {
		addUptimeCannot(o, Fill(SharedDebuff, map[string]string{
			"Aura": rule.Name, "n": uptimeInt(int64(appliers)), "classPlural": ClassPlural(in.Row.Class),
		}))
		return auraReading{}, false
	}
	switch {
	case track == nil && rule.Talent: // a talent never up: the log can't say they have it (§6.1)
		addUptimeCannot(o, Fill(CannotTalentAura, map[string]string{"Aura": rule.Name}))
		return auraReading{}, false
	case track == nil && appliedByOthers(in, rule): // someone else's debuff: not the player's to judge
		addUptimeCannot(o, Fill(CannotOthersDebuff, map[string]string{"Aura": rule.Name}))
		return auraReading{}, false
	}
	anchor := Anchor{Kind: AnchorAura, SpellID: rule.SpellID}
	up, den, ok := int64(0), c.CountedMS, c.CountedMS > 0 // never applied: up for none of it
	if track != nil {
		up, den, ok = measureUptime(*track, c)
		anchor.SpellID = track.SpellID
	} else {
		// No aura row exists to open, so the finding points at the time it counts.
		at, to := c.FirstCastMS, c.AliveMS
		anchor = Anchor{Kind: AnchorTime, AtMS: &at, ToMS: &to}
	}
	if !ok {
		return auraReading{}, false
	}
	ref, basis, thin := comparison(in, uptimePool(rule), uptimeSim(rule), in.Row.Role == RoleDPS)
	if thin != "" {
		addUptimeCannot(o, thin)
	}
	if basis == nil {
		return auraReading{}, false
	}
	return auraReading{rule: rule, upMS: up, denMS: den, pct: float64(up) / float64(den) * 100, ref: ref, basis: *basis, anchor: anchor}, true
}

// appliedByOthers reports whether a debuff-route aura (a DoT, a debuff, or a defensive that
// is not on the player) exists in the fight on any target though the player did not apply
// it. A self-buff never is: it is the player's or it is absent.
func appliedByOthers(in Input, rule AuraRule) bool {
	if rule.Kind == AuraSelfBuff {
		return false
	}
	return longestTrack(in, rule.IDs(), func(*summary.AuraTrack) bool { return true }) != nil
}

func addUptimeCannot(o *outcome, line string) {
	if !slices.Contains(o.CouldNotCheck, line) {
		o.CouldNotCheck = append(o.CouldNotCheck, line)
	}
}

func uptimeInt(n int64) string { return strconv.FormatInt(n, 10) }

func uptimePool(rule AuraRule) func(Profile) (float64, bool) {
	if rule.Kind == AuraDefensive {
		return func(p Profile) (float64, bool) { return p.DefensiveOf(rule.SpellID) }
	}
	return func(p Profile) (float64, bool) { return p.UptimeOf(rule.SpellID) }
}

func uptimeSim(rule AuraRule) func(SimRef) (float64, bool) {
	return func(s SimRef) (float64, bool) { v, ok := s.Uptime[rule.SpellID]; return v, ok }
}

// ownTrack is the aura's track on the player's side of the log, with the number of
// appliers when the log merges more than the player; nil when the player never had it.
// A self-buff is read on the player. A DoT or debuff is read on whichever target the
// player applied it to longest, and only a sole applier is judged (appliers > 1
// otherwise). A tank's defensive is its own buff when it sits on the player, else a
// debuff it applied to the boss (Demoralizing Shout, Thunder Clap) by the debuff rule.
func ownTrack(in Input, rule AuraRule) (track *summary.AuraTrack, appliers int) {
	guid := in.Row.GUID
	onPlayer := func(a *summary.AuraTrack) bool { return a.TargetGUID == guid }
	byPlayer := func(a *summary.AuraTrack) bool { return slices.Contains(a.Appliers, guid) }
	switch rule.Kind {
	case AuraSelfBuff:
		return longestTrack(in, rule.IDs(), onPlayer), 0
	case AuraDefensive:
		if t := longestTrack(in, rule.IDs(), onPlayer); t != nil {
			return t, 0
		}
	}
	track = longestTrack(in, rule.IDs(), byPlayer)
	if track != nil && len(track.Appliers) > 1 {
		return track, len(track.Appliers)
	}
	return track, 0
}

// longestTrack is the aura's track with the most uptime among targets whose joined track
// of any of ids matches keep, or nil. Every rank's tracks on one target are one track (a
// lower rank is still the aura): segments joined, appliers pooled across all ranks before
// keep looks at them, and the spell id is the rank that was up longest, so an anchor
// opens a row that exists.
func longestTrack(in Input, ids []int64, keep func(*summary.AuraTrack) bool) *summary.AuraTrack {
	byTarget := map[string]*summary.AuraTrack{}
	topRank := map[string]*summary.AuraTrack{}
	for i := range in.Fight.Auras {
		a := &in.Fight.Auras[i]
		if !slices.Contains(ids, a.SpellID) {
			continue
		}
		byTarget[a.TargetGUID] = unionTrack(byTarget[a.TargetGUID], a)
		if t := topRank[a.TargetGUID]; t == nil || a.UptimeMS > t.UptimeMS {
			topRank[a.TargetGUID] = a
		}
	}
	var best *summary.AuraTrack
	for target, t := range byTarget {
		t.SpellID = topRank[target].SpellID
		if !keep(t) {
			continue
		}
		if best == nil || t.UptimeMS > best.UptimeMS || (t.UptimeMS == best.UptimeMS && t.TargetGUID < best.TargetGUID) {
			best = t
		}
	}
	return best
}

// unionTrack is acc joined with a: segments merged where they overlap, appliers pooled,
// applications summed, seeded when either is. With acc nil it is a copy of a.
func unionTrack(acc, a *summary.AuraTrack) *summary.AuraTrack {
	if acc == nil {
		c := *a
		c.Segments = slices.Clone(a.Segments)
		c.Appliers = slices.Clone(a.Appliers)
		return &c
	}
	segs := append(slices.Clone(acc.Segments), a.Segments...)
	slices.SortFunc(segs, func(x, y summary.Segment) int { return cmp.Compare(x.StartMS, y.StartMS) })
	var merged []summary.Segment
	var up int64
	for _, s := range segs {
		if n := len(merged); n > 0 && s.StartMS <= merged[n-1].EndMS {
			merged[n-1].EndMS = max(merged[n-1].EndMS, s.EndMS)
			continue
		}
		merged = append(merged, s)
	}
	for _, s := range merged {
		up += s.EndMS - s.StartMS
	}
	acc.Segments, acc.UptimeMS = merged, up
	acc.Applications += a.Applications
	acc.Seeded = acc.Seeded || a.Seeded
	for _, g := range a.Appliers {
		if !slices.Contains(acc.Appliers, g) {
			acc.Appliers = append(acc.Appliers, g)
		}
	}
	return acc
}

// measureUptime is the aura's time up and the time it could have been up: from the first
// application to the player's death or the end, both less the excluded windows.
func measureUptime(t summary.AuraTrack, c Counted) (upMS, denMS int64, ok bool) {
	if len(t.Segments) == 0 {
		return 0, 0, false
	}
	from := t.Segments[0].StartMS
	for _, s := range t.Segments {
		from = min(from, s.StartMS)
	}
	denMS = countedBetween(c.Excluded, from, c.AliveMS)
	for _, s := range t.Segments {
		upMS += countedBetween(c.Excluded, s.StartMS, min(s.EndMS, c.AliveMS))
	}
	return min(upMS, denMS), denMS, denMS > 0
}

func addUptimeStrength(in Input, r auraReading, o *outcome) {
	o.Strengths = append(o.Strengths, Strength{
		Sentence: Fill(StrengthUptime, map[string]string{
			"Aura": r.rule.Name, "pct": PctNumber(r.pct), "n": uptimeInt(int64(r.basis.N)), "ref": PctNumber(r.ref),
		}),
		Facts:   uptimeFacts(in, r),
		Anchors: []string{addUptimeAnchor(o, r.anchor)},
	})
}

// uptimeFinding is the worst aura's finding; up to MoreAbilitiesInDetail more go in the
// detail.
func uptimeFinding(in Input, low []auraReading, o *outcome) *Finding {
	sort.SliceStable(low, func(i, j int) bool { return low[i].gap() > low[j].gap() })
	worst := low[0]
	more := low[1:min(len(low), 1+MoreAbilitiesInDetail)]
	detail := Fill(DetailUptime, map[string]string{
		"Aura": worst.rule.Name, "n": uptimeInt(worst.upMS / 1000), "duration": Clock(worst.denMS),
	})
	if len(more) > 0 {
		detail += " " + Fill(detailUptimeMore, map[string]string{"others": moreAurasText(more)})
	}
	sev := SeverityMinor
	if worst.gap() >= UptimeMajorPoints {
		sev = SeverityMajor
	}
	facts := uptimeFacts(in, worst)
	anchors := []string{addUptimeAnchor(o, worst.anchor)}
	for _, m := range more {
		facts = mergeUptimeFacts(facts, uptimeFacts(in, m))
		anchors = append(anchors, addUptimeAnchor(o, m.anchor))
	}
	return &Finding{
		Code: CodeUptime, Severity: sev,
		Title:  Fill(TitleUptime, map[string]string{"Aura": worst.rule.Name, "pct": PctNumber(worst.pct)}),
		Detail: detail,
		Basis: Basis{Kind: worst.basis.Kind, N: worst.basis.N, Label: worst.basis.Label +
			Fill(basisUptimeValue, map[string]string{"ref": PctNumber(worst.ref), "you": PctNumber(worst.pct)})},
		Anchors: anchors, Facts: facts,
	}
}

func moreAurasText(more []auraReading) string {
	parts := make([]string, len(more))
	for i, m := range more {
		parts[i] = Fill(uptimeMoreItem, map[string]string{"Aura": m.rule.Name, "pct": PctNumber(m.pct)})
	}
	return strings.Join(parts, ", ")
}

// uptimeFacts are the numbers, the spell and the time a reading states.
func uptimeFacts(in Input, r auraReading) Facts {
	name := r.rule.Name
	f := Facts{
		Numbers: []NumberFact{
			{V: r.pct, Text: PctNumber(r.pct), Unit: "%", Of: name + " uptime"},
			{V: r.ref, Text: PctNumber(r.ref), Unit: "%", Of: name + " uptime in the comparison"},
			{V: float64(r.upMS / 1000), Text: uptimeInt(r.upMS / 1000), Unit: "s", Of: name + " up"},
		},
		Spells: []SpellFact{{ID: r.rule.SpellID, Name: name, Rank: in.Tables.Spells[r.rule.SpellID].Rank}},
		Times:  []string{Clock(r.denMS)},
	}
	if r.basis.N > 0 {
		f.Numbers = append(f.Numbers, NumberFact{V: float64(r.basis.N), Text: uptimeInt(int64(r.basis.N)), Unit: "players", Of: "top logs"})
	}
	return f
}

func mergeUptimeFacts(a, b Facts) Facts {
	return Facts{
		Numbers: append(a.Numbers, b.Numbers...), Spells: append(a.Spells, b.Spells...),
		Items: append(a.Items, b.Items...), Times: append(a.Times, b.Times...),
	}
}

// addUptimeAnchor adds a local anchor and returns its "#n" reference.
func addUptimeAnchor(o *outcome, a Anchor) string {
	o.Anchors = append(o.Anchors, a)
	return "#" + uptimeInt(int64(len(o.Anchors)-1))
}
