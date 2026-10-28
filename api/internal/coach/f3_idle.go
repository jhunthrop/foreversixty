package coach

import (
	"slices"
	"strconv"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

const (
	idleFactSeconds = "seconds idle"
	idleFactShare   = "idle share of counted time"
	idleFactLongest = "longest idle window, seconds"
	idleFactPool    = "active share, top logs"
	idleFactYou     = "active share, you"
	millisPerSecond = 1000.0
)

// idle is F3 (§6.2): the sum of the idle windows (ratingengine.IdleWindows, net of
// downtime and officer assignments) less the time never counted against the player, with
// the longest windows. The title and the detail state the total; the detail lists the
// longest windows. A healer gets nothing.
func idle(in Input, c Counted) outcome {
	if in.Row.Role != RoleDPS && in.Row.Role != RoleTank {
		return outcome{}
	}
	windows := countedIdleWindows(in, c)
	idleMS := IdleTotalMS(windows)
	if c.CountedMS <= 0 || idleMS < IdleFiresMS {
		return outcome{}
	}
	idlePct := float64(idleMS) / float64(c.CountedMS) * percentOf
	if idlePct < IdleFiresPct {
		return outcome{}
	}
	shown := longestInTimeOrder(windows, IdleWindowsShown)
	return outcome{
		Finding: idleFinding(in, idleMS, idlePct, shown, windows[0]),
		Anchors: idleAnchors(shown),
		Windows: windows,
	}
}

// countedIdleWindows is the idle windows with the time never counted against the player
// (Counted.Excluded: downtime and the lead-in before the first cast) cut out, longest first.
func countedIdleWindows(in Input, c Counted) []ratingengine.Window {
	var out []ratingengine.Window
	for _, w := range ratingengine.IdleWindows(in.Fight, in.Row.GUID, in.Tables.Mechanics, in.Tables.Assignments) {
		out = append(out, cutOut(w, c.Excluded)...)
	}
	slices.SortStableFunc(out, func(a, b ratingengine.Window) int {
		return int((b.ToMS - b.FromMS) - (a.ToMS - a.FromMS))
	})
	return out
}

// cutOut is w less every excluded window (time-ordered, non-overlapping).
func cutOut(w ratingengine.Window, excluded []ratingengine.Window) []ratingengine.Window {
	var out []ratingengine.Window
	from := w.FromMS
	for _, x := range excluded {
		if x.ToMS <= from || x.FromMS >= w.ToMS {
			continue
		}
		if x.FromMS > from {
			out = append(out, ratingengine.Window{FromMS: from, ToMS: x.FromMS})
		}
		from = max(from, x.ToMS)
	}
	if from < w.ToMS {
		out = append(out, ratingengine.Window{FromMS: from, ToMS: w.ToMS})
	}
	return out
}

// longestInTimeOrder is the n longest of windows (already longest first), earliest first.
func longestInTimeOrder(windows []ratingengine.Window, n int) []ratingengine.Window {
	top := slices.Clone(windows[:min(n, len(windows))])
	slices.SortFunc(top, func(a, b ratingengine.Window) int { return int(a.FromMS - b.FromMS) })
	return top
}

func idleAnchors(shown []ratingengine.Window) []Anchor {
	out := make([]Anchor, len(shown))
	for i, w := range shown {
		at, to := w.FromMS, w.ToMS
		out[i] = Anchor{Kind: AnchorTime, AtMS: &at, ToMS: &to}
	}
	return out
}

func idleFinding(in Input, idleMS int64, idlePct float64, shown []ratingengine.Window, longest ratingengine.Window) *Finding {
	seconds := PctNumber(float64(idleMS) / millisPerSecond)
	longestSeconds := PctNumber(float64(longest.ToMS-longest.FromMS) / millisPerSecond)
	spans := make([]string, len(shown))
	share := PctNumber(idlePct)
	refs := make([]string, len(shown))
	for i, w := range shown {
		spans[i] = Span(w.FromMS, w.ToMS)
		refs[i] = "#" + strconv.Itoa(i)
	}
	basis, numbers := idleBasis(in, idlePct)
	numbers = append([]NumberFact{
		{V: float64(longest.ToMS-longest.FromMS) / millisPerSecond, Text: longestSeconds, Unit: "s", Of: idleFactLongest},
		{V: float64(idleMS) / millisPerSecond, Text: seconds, Unit: "s", Of: idleFactSeconds},
		{V: idlePct, Text: share, Unit: "%", Of: idleFactShare},
	}, numbers...)
	severity := SeverityMinor
	if idlePct >= IdleMajorPct {
		severity = SeverityMajor
	}
	return &Finding{
		Code:     CodeIdle,
		Severity: severity,
		Title:    Fill(TitleIdle, map[string]string{"n": seconds}),
		Detail:   Fill(DetailIdle, map[string]string{"n": seconds, "pct": share, "windows": joinAnd(spans)}),
		Basis:    basis,
		Anchors:  refs,
		Facts:    Facts{Numbers: numbers, Times: spans},
	}
}

// activeShare is a profile's active share of counted time, when it records one.
func activeShare(p Profile) (float64, bool) {
	if p.ActivePct == nil {
		return 0, false
	}
	return *p.ActivePct, true
}

// countedActivePct is the player's active share of counted time (100 − F3's idle share),
// the unit the pool's median is in; nil for a role F3 never judges or no counted time.
func countedActivePct(in Input, c Counted) *float64 {
	if (in.Row.Role != RoleDPS && in.Row.Role != RoleTank) || c.CountedMS <= 0 {
		return nil
	}
	v := percentOf - float64(IdleTotalMS(countedIdleWindows(in, c)))/float64(c.CountedMS)*percentOf
	return &v
}

// idleBasis is the top logs' median active share of counted time when the pool has one,
// set beside the player's own in the same unit, else the threshold from the player's log.
func idleBasis(in Input, idlePct float64) (Basis, []NumberFact) {
	ref, b, _ := comparison(in, activeShare, nil, false)
	if b == nil {
		return Basis{Label: BasisIdle, Kind: BasisOwnLog}, nil
	}
	you := percentOf - idlePct
	fields := topFields(in, b.N)
	fields["ref"], fields["you"] = PctNumber(ref), PctNumber(you)
	b.Label = Fill(BasisIdleTop, fields)
	return *b, []NumberFact{
		{V: ref, Text: PctNumber(ref), Unit: "%", Of: idleFactPool},
		{V: you, Text: PctNumber(you), Unit: "%", Of: idleFactYou},
	}
}
