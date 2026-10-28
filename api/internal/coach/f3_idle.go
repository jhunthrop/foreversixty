package coach

import (
	"slices"
	"strconv"
	"strings"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

// idlePoolBasis follows the top-logs label: the pool's median active share against the
// player's own.
const idlePoolBasis = "{ref}% of counted time active. You: {you}%."

const (
	idleFactSeconds = "seconds idle"
	idleFactLongest = "longest idle window, seconds"
	idleFactPool    = "active share, top logs"
	idleFactYou     = "active share, you"
	millisPerSecond = 1000.0
)

// idle is F3 (§6.2): the sum of the idle windows (ratingengine.IdleWindows, net of
// downtime and officer assignments) less the time never counted against the player, with
// the longest windows. The title names the longest window; the detail the total. A healer gets nothing.
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
		Finding: idleFinding(in, c, idleMS, idlePct, shown, windows[0]),
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

func idleFinding(in Input, c Counted, idleMS int64, idlePct float64, shown []ratingengine.Window, longest ratingengine.Window) *Finding {
	seconds := PctNumber(float64(idleMS) / millisPerSecond)
	longestSeconds := PctNumber(float64(longest.ToMS-longest.FromMS) / millisPerSecond)
	spans := make([]string, len(shown))
	refs := make([]string, len(shown))
	for i, w := range shown {
		spans[i] = Span(w.FromMS, w.ToMS)
		refs[i] = "#" + strconv.Itoa(i)
	}
	basis, numbers := idleBasis(in, idlePct)
	numbers = append([]NumberFact{{V: float64(longest.ToMS-longest.FromMS) / millisPerSecond, Text: longestSeconds, Unit: "s", Of: idleFactLongest}, {V: float64(idleMS) / millisPerSecond, Text: seconds, Unit: "s", Of: idleFactSeconds}}, numbers...)
	severity := SeverityMinor
	if idlePct >= IdleMajorPct {
		severity = SeverityMajor
	}
	return &Finding{
		Code:     CodeIdle,
		Severity: severity,
		Title:    Fill(Drafts["title_idle"], map[string]string{"n": longestSeconds}),
		Detail: Fill(Drafts["detail_idle"], map[string]string{
			"n": seconds, "counted": Clock(c.CountedMS), "spans": strings.Join(spans, ", "),
		}),
		Basis:   basis,
		Anchors: refs,
		Facts:   Facts{Numbers: numbers, Times: spans},
	}
}

// idleBasis is the top logs' median active share when the pool has one, else the
// threshold from the player's own log.
func idleBasis(in Input, idlePct float64) (Basis, []NumberFact) {
	ref, b, _ := comparison(in, func(p Profile) (float64, bool) { return p.ActivePct, p.ActivePct > 0 }, nil, false)
	if b == nil {
		return Basis{Label: Drafts["basis_idle_threshold"], Kind: BasisOwnLog}, nil
	}
	you := percentOf - idlePct
	b.Label += Fill(idlePoolBasis, map[string]string{"ref": PctNumber(ref), "you": PctNumber(you)})
	return *b, []NumberFact{
		{V: ref, Text: PctNumber(ref), Unit: "%", Of: idleFactPool},
		{V: you, Text: PctNumber(you), Unit: "%", Of: idleFactYou},
	}
}
