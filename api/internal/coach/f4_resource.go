package coach

import (
	"strconv"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// detailManaRecovered is F4's detail when mana came back after running out. The spec
// shows no such string, so it lives here until Task 33 moves it into the draft table.
const detailManaRecovered = "Your mana was under {pct}% for {n} s in all, from {t}."

// downtimeNoun fills DetailCapped's {downtime}; the spec's example is "air phase".
// Task 33 may name each encounter's downtime.
const downtimeNoun = "downtime"

// rolePowers are the bars each role is checked on (§6.3), in the order they are tried.
var rolePowers = map[string][]int64{
	RoleDPS:    {PowerRage, PowerEnergy, PowerMana},
	RoleTank:   {PowerRage, PowerMana},
	RoleHealer: {PowerMana},
}

// waste is one bar's finding with what ranks it against the player's other bars.
type waste struct {
	severity Severity
	totalMS  int64
	out      outcome
}

// resource is F4 (§6.2): rage or energy sitting at the cap, mana running out early.
// When several bars fire the worse one is the finding: major before minor, then the
// longer.
func resource(in Input, c Counted) outcome {
	if c.CountedMS <= 0 {
		return outcome{}
	}
	var best *waste
	for _, power := range rolePowers[in.Row.Role] {
		track, ok := trackOf(in.Fight, in.Row.GUID, power)
		if !ok || track.Max <= 0 {
			continue
		}
		if w := assess(in, c, track); w != nil && worse(w, best) {
			best = w
		}
	}
	if best == nil {
		return outcome{}
	}
	return best.out
}

func trackOf(fight summary.Summary, guid string, power int64) (summary.ResourceTrack, bool) {
	for _, r := range fight.Resources {
		if r.GUID == guid && r.PowerType == power {
			return r, true
		}
	}
	return summary.ResourceTrack{}, false
}

func worse(w, than *waste) bool {
	if than == nil {
		return true
	}
	if w.severity != than.severity {
		return w.severity == SeverityMajor
	}
	return w.totalMS > than.totalMS
}

func assess(in Input, c Counted, track summary.ResourceTrack) *waste {
	if track.PowerType == PowerMana {
		return assessMana(in, c, track)
	}
	return assessCap(in, c, track)
}

// assessCap counts the counted seconds the bar sat at Max.
func assessCap(in Input, c Counted, track summary.ResourceTrack) *waste {
	runs := countedRuns(track.Series, c, func(v int64) bool { return v >= track.Max })
	total := totalMS(runs)
	pct := float64(total) * 100 / float64(c.CountedMS)
	if total < CapFiresMS || pct < CapFiresPct {
		return nil
	}
	severity := SeverityMinor
	if pct >= CapMajorPct {
		severity = SeverityMajor
	}
	bar := PowerNames[track.PowerType]
	longest := longestRun(runs)
	secs := strconv.FormatInt(total/ResourceWholeSecondMS, 10)
	finding := Finding{
		Code: CodeResource, Severity: severity,
		Title:  Fill(TitleCapped, map[string]string{"Resource": capitalised(bar), "n": secs}),
		Detail: capDetail(in, c, bar, secs, longest),
		Basis:  Basis{Label: capBasis(in, c, bar), Kind: BasisOwnLog},
		Facts: Facts{
			Numbers: []NumberFact{{V: float64(total / ResourceWholeSecondMS), Text: secs, Unit: "s", Of: bar + " at its cap"}},
			Times:   []string{Clock(longest.FromMS), Clock(longest.ToMS)},
		},
	}
	return &waste{severity: severity, totalMS: total, out: outcomeOf(finding, runs, longest)}
}

func hasDowntime(in Input, c Counted) bool {
	return len(ratingengine.DowntimeWindows(in.Fight, in.Tables.Mechanics, c.AliveMS)) > 0
}

func capDetail(in Input, c Counted, bar, secs string, longest ratingengine.Window) string {
	fields := map[string]string{"resource": bar, "n": secs, "span": Span(longest.FromMS, longest.ToMS)}
	if hasDowntime(in, c) {
		fields["downtime"] = downtimeNoun
		return Fill(DetailCapped, fields)
	}
	return Fill(Drafts["detail_capped_no_downtime"], fields)
}

func capBasis(in Input, c Counted, bar string) string {
	if hasDowntime(in, c) {
		return Fill(BasisCapOwnBar, map[string]string{"resource": bar, "downtime": downtimeNoun})
	}
	return Fill(BasisOwnBar, map[string]string{"resource": bar})
}

// assessMana counts the seconds alive under ManaLowPct of Max. §6.2 sets downtime aside
// for rage and energy only, so a mana second inside downtime or before the first cast counts.
func assessMana(in Input, c Counted, track summary.ResourceTrack) *waste {
	floor := ManaLowPct * float64(track.Max)
	runs := countedRuns(track.Series, Counted{AliveMS: c.AliveMS}, func(v int64) bool { return float64(v)*100 < floor })
	total := totalMS(runs)
	if len(runs) == 0 || total < ManaFiresMS || in.Fight.DurationMS <= 0 {
		return nil
	}
	first := runs[0]
	startPct := float64(first.FromMS) * 100 / float64(in.Fight.DurationMS)
	if startPct >= ManaFiresBeforePct {
		return nil
	}
	severity := SeverityMinor
	if startPct < ManaMajorBeforePct {
		severity = SeverityMajor
	}
	secs := strconv.FormatInt(total/ResourceWholeSecondMS, 10)
	finding := Finding{
		Code: CodeResource, Severity: severity,
		Title:  Fill(TitleOutOfMana, map[string]string{"t": Clock(first.FromMS)}),
		Detail: manaDetail(in, c, runs, secs),
		Basis:  Basis{Label: Fill(BasisOwnBar, map[string]string{"resource": "mana"}), Kind: BasisOwnLog},
		Facts: Facts{
			Numbers: []NumberFact{
				{V: float64(total / ResourceWholeSecondMS), Text: secs, Unit: "s", Of: "mana under the floor"},
				{V: ManaLowPct, Text: PctNumber(ManaLowPct), Unit: "percent", Of: "of your mana bar"},
			},
			Times: []string{Clock(first.FromMS)},
		},
	}
	return &waste{severity: severity, totalMS: total, out: outcomeOf(finding, runs, first)}
}

// manaDetail words how the low stretch ended: by death, by the end of the fight, or by
// mana coming back. Only one unbroken stretch can be said to have lasted "until".
func manaDetail(in Input, c Counted, runs []ratingengine.Window, secs string) string {
	fields := map[string]string{"t": Clock(runs[0].FromMS), "n": secs, "pct": PctNumber(ManaLowPct)}
	if len(runs) == 1 && runs[0].ToMS >= c.AliveMS {
		if c.AliveMS < in.Fight.DurationMS {
			fields["death"] = Clock(c.AliveMS)
			return Fill(DetailOutOfManaDeath, fields)
		}
		return Fill(Drafts["detail_out_of_mana_to_end"], fields)
	}
	return Fill(detailManaRecovered, fields)
}

// outcomeOf wraps a finding with its one resource anchor, at the anchored window, and
// the runs as the finding's windows.
func outcomeOf(f Finding, runs []ratingengine.Window, anchored ratingengine.Window) outcome {
	at, to := anchored.FromMS, anchored.ToMS
	f.Anchors = []string{"#0"}
	return outcome{
		Finding: &f,
		Anchors: []Anchor{{Kind: AnchorResource, AtMS: &at, ToMS: &to}},
		Windows: runs,
	}
}

// countedRuns are the unbroken stretches of counted seconds whose reading satisfies
// match, in time order. Second i covers [i s, i+1 s), clipped to the time alive.
func countedRuns(series []int64, c Counted, match func(int64) bool) []ratingengine.Window {
	var runs []ratingengine.Window
	open := false
	for i, v := range series {
		from := int64(i) * ResourceWholeSecondMS
		if !isCounted(c, from) || !match(v) {
			open = false
			continue
		}
		to := min(from+ResourceWholeSecondMS, c.AliveMS)
		if open {
			runs[len(runs)-1].ToMS = to
			continue
		}
		runs = append(runs, ratingengine.Window{FromMS: from, ToMS: to})
		open = true
	}
	return runs
}

// isCounted reports whether the second starting at fromMS is alive and outside every
// excluded window.
func isCounted(c Counted, fromMS int64) bool {
	if fromMS >= c.AliveMS {
		return false
	}
	toMS := fromMS + ResourceWholeSecondMS
	for _, w := range c.Excluded {
		if w.FromMS < toMS && fromMS < w.ToMS {
			return false
		}
	}
	return true
}

func totalMS(ws []ratingengine.Window) int64 {
	var sum int64
	for _, w := range ws {
		sum += w.ToMS - w.FromMS
	}
	return sum
}

// longestRun is the longest window, the earliest on a tie.
func longestRun(ws []ratingengine.Window) ratingengine.Window {
	var best ratingengine.Window
	for _, w := range ws {
		if w.ToMS-w.FromMS > best.ToMS-best.FromMS {
			best = w
		}
	}
	return best
}

func capitalised(s string) string {
	if s == "" {
		return s
	}
	return string(s[0]-'a'+'A') + s[1:]
}
