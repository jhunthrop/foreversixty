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
	idleFactPool    = "active share, top logs"
	idleFactYou     = "active share, you"
	percentOf       = 100.0
	millisPerSecond = 1000.0
)

// idle is F3 (§6.2): the time the Activity measure counts against a DPS or tank, net of
// downtime and officer assignments, with its longest windows. A healer gets nothing.
func idle(in Input, c Counted) outcome {
	if in.Row.Role != RoleDPS && in.Row.Role != RoleTank {
		return outcome{}
	}
	guid, table, marks := in.Row.GUID, in.Tables.Mechanics, in.Tables.Assignments
	idleMS := ratingengine.ActivityIdleMS(in.Fight, guid, table, marks)
	if c.CountedMS <= 0 || idleMS < IdleFiresMS {
		return outcome{}
	}
	idlePct := float64(idleMS) / float64(c.CountedMS) * percentOf
	if idlePct < IdleFiresPct {
		return outcome{}
	}
	windows := ratingengine.IdleWindows(in.Fight, guid, table, marks)
	shown := longestInTimeOrder(windows, IdleWindowsShown)
	return outcome{
		Finding: idleFinding(in, c, idleMS, idlePct, shown),
		Anchors: idleAnchors(shown),
		Windows: windows,
	}
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

func idleFinding(in Input, c Counted, idleMS int64, idlePct float64, shown []ratingengine.Window) *Finding {
	seconds := PctNumber(float64(idleMS) / millisPerSecond)
	spans := make([]string, len(shown))
	refs := make([]string, len(shown))
	for i, w := range shown {
		spans[i] = Span(w.FromMS, w.ToMS)
		refs[i] = "#" + strconv.Itoa(i)
	}
	basis, numbers := idleBasis(in, idlePct)
	numbers = append([]NumberFact{{V: float64(idleMS) / millisPerSecond, Text: seconds, Unit: "s", Of: idleFactSeconds}}, numbers...)
	severity := SeverityMinor
	if idlePct >= IdleMajorPct {
		severity = SeverityMajor
	}
	return &Finding{
		Code:     CodeIdle,
		Severity: severity,
		Title:    Fill(Drafts["title_idle"], map[string]string{"n": seconds}),
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
