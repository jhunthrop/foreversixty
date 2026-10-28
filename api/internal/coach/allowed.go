package coach

import (
	"slices"
	"strconv"
	"strings"
)

// allowedOf is the validator's set for p (§7): every number, time and range the pack's
// findings, strengths, links and anchors state, normalised as ExtractNumbers normalises
// prose (2,140 is 2140, a percent keeps its %, a range's two ends are times too); the
// size, level and duration; every spell, item and boss name and the player's spec and
// class; every pack id.
func allowedOf(p Pack) Allowed {
	b := allowedBuilder{numbers: map[string]bool{}, names: map[string]bool{}, ids: map[string]bool{}}
	for _, f := range p.Findings {
		b.text(f.Title, f.Detail, f.Basis.Label)
		if f.Comparison != nil {
			b.text(f.Comparison.Label)
		}
		b.facts(f.Facts)
		b.ids[f.ID] = true
	}
	for _, s := range p.Strengths {
		b.text(s.Sentence)
		b.facts(s.Facts)
		b.ids[s.ID] = true
	}
	for _, l := range p.Links {
		b.text(l.Sentence)
		b.ids[l.ID] = true
	}
	for _, a := range p.Anchors {
		b.anchor(a)
	}
	b.context(p.Context)
	return Allowed{Numbers: sortedKeys(b.numbers), Names: sortedKeys(b.names), IDs: sortedKeys(b.ids)}
}

type allowedBuilder struct {
	numbers, names, ids map[string]bool
}

// text adds every number, time and range in the strings; a range's ends as times too.
func (b allowedBuilder) text(ss ...string) {
	for _, s := range ss {
		for _, t := range ExtractNumbers(s) {
			b.numbers[t.Norm] = true
			if t.Kind == KindRange {
				from, to, _ := strings.Cut(t.Norm, enDash)
				b.numbers[from], b.numbers[to] = true, true
			}
		}
	}
}

// facts adds a fact's numbers (a percent unit with and without its sign), times, spell
// and item names.
func (b allowedBuilder) facts(f Facts) {
	for _, n := range f.Numbers {
		b.text(n.Text)
		if n.Unit == "%" || n.Unit == "percent" {
			b.text(n.Text + "%")
		}
	}
	b.text(f.Times...)
	for _, s := range f.Spells {
		b.name(s.Name)
	}
	for _, item := range f.Items {
		b.name(item)
	}
}

// anchor adds an anchor's id and the times it opens at.
func (b allowedBuilder) anchor(a Anchor) {
	b.ids[a.ID] = true
	switch {
	case a.AtMS != nil && a.ToMS != nil:
		b.text(Span(*a.AtMS, *a.ToMS))
	case a.AtMS != nil:
		b.text(Clock(*a.AtMS))
	}
}

func (b allowedBuilder) context(c Context) {
	b.text(strconv.Itoa(c.Size), c.Duration)
	if c.Level > 0 {
		b.text(strconv.Itoa(c.Level))
	}
	b.name(c.Boss)
	b.name(c.SpecLabel)
	b.name(c.Class)
}

func (b allowedBuilder) name(n string) {
	if n = strings.TrimSpace(n); n != "" {
		b.names[n] = true
	}
}

func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	slices.Sort(out)
	return out
}

// HasNumber reports a number or percent in its normalised form ("2140", "44%").
func (a Allowed) HasNumber(norm string) bool { return a.hasKind(norm, KindNumber, KindPercent) }

// HasTime reports an m:ss time.
func (a Allowed) HasTime(norm string) bool { return a.hasKind(norm, KindTime) }

// HasRange reports an m:ss–m:ss range, joined by an en dash.
func (a Allowed) HasRange(norm string) bool { return a.hasKind(norm, KindRange) }

// HasName reports a spell, item, boss, spec or class name, ignoring case.
func (a Allowed) HasName(name string) bool {
	return slices.ContainsFunc(a.Names, func(n string) bool { return strings.EqualFold(n, strings.TrimSpace(name)) })
}

// HasID reports a finding, strength, link or anchor id.
func (a Allowed) HasID(id string) bool {
	_, ok := slices.BinarySearch(a.IDs, id)
	return ok
}

func (a Allowed) hasKind(norm string, kinds ...string) bool {
	toks := ExtractNumbers(norm)
	if len(toks) != 1 || toks[0].Norm != norm || !slices.Contains(kinds, toks[0].Kind) {
		return false
	}
	_, ok := slices.BinarySearch(a.Numbers, norm)
	return ok
}
