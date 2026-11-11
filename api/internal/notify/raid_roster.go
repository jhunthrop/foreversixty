package notify

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// rosterCopy is the roster's words on the raid post (spec 2026-10-28 §P.7, the roster.* rows).
var rosterCopy = struct {
	Tanks, Healers, Melee, Ranged string // each takes {n}
	OfTarget                      string // a group heading with its target: {heading}, {want}
	OverTarget                    string // a group past its target: {heading}, {want} the target, {over}
	In                            string // a guild event's In list; takes {n}
	Late, Bench, Out              string // each takes {n}
	Continued                     string
	Empty                         string
	More                          string // {n}
	OutJoin                       string
	Sep                           string // between a name and its spec or class in a text line
	EmojiFormat                   string // %s the emoji name, %s its id, %s the name
}{
	Tanks: copygen.DiscordRosterTanks, Healers: copygen.DiscordRosterHealers, Melee: copygen.DiscordRosterMelee, Ranged: copygen.DiscordRosterRanged,
	OfTarget: copygen.DiscordRosterOfTarget, OverTarget: copygen.DiscordRosterOverTarget,
	In: copygen.DiscordRosterIn, Late: copygen.DiscordRosterLate, Bench: copygen.DiscordRosterBench, Out: copygen.DiscordRosterOut,
	Continued: "\u200b", Empty: copygen.DiscordRaidSignedNone, More: copygen.DiscordRosterMore, OutJoin: ", ", Sep: " · ",
	EmojiFormat: "<:%s:%s> %s",
}

// The roster's Discord bounds (spec §P.9.5).
const (
	fieldLines      = 16
	groupFieldsMax  = 2
	listLines       = 10
	fieldValueMax   = 1024
	embedGuardChars = 5900
	// rosterNameRunes bounds a name (a WoW name is at most 12) and its escaped form.
	rosterNameRunes    = 12
	rosterNameOutRunes = 24
)

// The sign-up states the roster reads (guild_raid_signups.state).
const (
	rosterIn    = "in"
	rosterLate  = "late"
	rosterBench = "bench"
	rosterOut   = "out"
)

// RosterLine is one sign-up as the post lists it: the row's name, class slug, spec (its
// catalogue name, "" when unknown), role and state.
type RosterLine struct {
	Name, Class, Spec, Role, State string
}

// Marks are the application emoji ids by name (fs_{class}_{spec}); nil draws every line as text.
type Marks map[string]string

// rosterGroup is one of the four In groups: its header and the raidspec.Group it lists, the one
// mapping the site's Sign-ups columns read too (spec 2026-11-05 §4.1).
type rosterGroup struct {
	label string
	group string
}

var rosterGroups = []rosterGroup{
	{rosterCopy.Tanks, raidspec.GroupTank},
	{rosterCopy.Healers, raidspec.GroupHealer},
	{rosterCopy.Melee, raidspec.GroupMelee},
	{rosterCopy.Ranged, raidspec.GroupRanged},
}

// rosterLimits are the caps the embed guard tightens (spec §P.9.5).
type rosterLimits struct {
	group       []int // lines per group, Tanks to Ranged
	late, bench int
	outChars    int
}

func defaultLimits() rosterLimits {
	most := fieldLines * groupFieldsMax
	return rosterLimits{group: []int{most, most, most, most}, late: listLines, bench: listLines, outChars: fieldValueMax}
}

// rosterFields are the groups, Late, Bench and Out of the post (spec §P.9.2), with the guard
// keeping the whole embed under Discord's 6000: while over 5900, the tail of Out, then Bench,
// then the largest group becomes `and {n} more`.
// noRoles draws a guild event's one plain list (eventRosterFields) instead.
func rosterFields(roster []RosterLine, marks Marks, target *RoleTarget, fixed int, noRoles bool) []EmbedField {
	lim := defaultLimits()
	within := func(r []RosterLine, m Marks, l rosterLimits) []EmbedField { return rosterFieldsWithin(r, m, target, l) }
	if noRoles {
		within = eventRosterFields
	}
	for {
		fields := within(roster, marks, lim)
		if fixed+fieldsChars(fields) <= embedGuardChars || !lim.tighten(roster) {
			return fields
		}
	}
}

// tighten lowers one cap; false when nothing is left to lower.
func (l *rosterLimits) tighten(roster []RosterLine) bool {
	switch {
	case l.outChars > 0 && countState(roster, rosterOut) > 0:
		l.outChars = max(0, l.outChars-200)
	case l.bench > 1 && countState(roster, rosterBench) > 1:
		l.bench--
	case l.late > 1 && countState(roster, rosterLate) > 1:
		l.late--
	default:
		i := l.largestGroup()
		if l.group[i] <= 1 {
			return false
		}
		l.group[i]--
	}
	return true
}

func (l *rosterLimits) largestGroup() int {
	best := 0
	for i, n := range l.group {
		if n > l.group[best] {
			best = i
		}
	}
	return best
}

func countState(roster []RosterLine, state string) int {
	n := 0
	for _, r := range roster {
		if r.State == state {
			n++
		}
	}
	return n
}

func fieldsChars(fields []EmbedField) int {
	n := 0
	for _, f := range fields {
		n += utf8.RuneCountInString(f.Name) + utf8.RuneCountInString(f.Value)
	}
	return n
}

func rosterFieldsWithin(roster []RosterLine, marks Marks, target *RoleTarget, lim rosterLimits) []EmbedField {
	var fields []EmbedField
	// Fields 4 to 6 (the groups, Late, Bench) are drawn only once someone is In or Late (§P.9.2).
	if countState(roster, rosterIn)+countState(roster, rosterLate) > 0 {
		for i, g := range rosterGroups {
			rows := inGroup(roster, g)
			heading := counted(g.label, len(rows))
			if target != nil {
				want := target.of(g.group)
				heading = copygen.Fill(rosterCopy.OfTarget, "{heading}", heading, "{want}", strconv.Itoa(want))
				if over := len(rows) - want; over > 0 {
					heading = copygen.Fill(rosterCopy.OverTarget, "{heading}", counted(g.label, len(rows)),
						"{want}", strconv.Itoa(want), "{over}", strconv.Itoa(over))
				}
			}
			fields = append(fields, groupFields(heading, lines(rows, marks), lim.group[i])...)
		}
		fields = append(fields, listFields(roster, marks, lim)...)
	}
	if out := outNames(roster); len(out) > 0 {
		fields = append(fields, EmbedField{Name: counted(rosterCopy.Out, len(out)), Value: outValue(out, lim.outChars)})
	}
	return fields
}

// eventRosterFields are a guild event's sign-ups (events spec §3.5): no role groups and no spec,
// one list grouped In, Late and Bench, each `{Name} · {Class}` and drawn only when someone is in
// it, then Out by name as on a raid.
func eventRosterFields(roster []RosterLine, _ Marks, lim rosterLimits) []EmbedField {
	var fields []EmbedField
	for _, list := range []struct {
		label, state string
		most         int
	}{{rosterCopy.In, rosterIn, lim.group[0]}, {rosterCopy.Late, rosterLate, lim.late}, {rosterCopy.Bench, rosterBench, lim.bench}} {
		rows := sortedState(roster, list.state)
		if len(rows) == 0 {
			continue
		}
		ls := make([]string, len(rows))
		for i, r := range rows {
			ls[i] = textLine(rosterName(r.Name), r.Class, raidspec.Spec{}, false)
		}
		fields = append(fields, groupFields(counted(list.label, len(rows)), ls, list.most)...)
	}
	if out := outNames(roster); len(out) > 0 {
		fields = append(fields, EmbedField{Name: counted(rosterCopy.Out, len(out)), Value: outValue(out, lim.outChars)})
	}
	return fields
}

// listFields are Late and Bench, each only when someone is in it.
func listFields(roster []RosterLine, marks Marks, lim rosterLimits) []EmbedField {
	var fields []EmbedField
	for _, list := range []struct {
		label, state string
		most         int
	}{{rosterCopy.Late, rosterLate, lim.late}, {rosterCopy.Bench, rosterBench, lim.bench}} {
		if rows := sortedState(roster, list.state); len(rows) > 0 {
			fields = append(fields, EmbedField{Name: counted(list.label, len(rows)),
				Value: fitValue(capped(lines(rows, marks), list.most)), Inline: true})
		}
	}
	return fields
}

// name is the group's word alone, its label before the count ("Healers").
func (g rosterGroup) name() string {
	name, _, _ := strings.Cut(g.label, rosterCopy.Sep)
	return name
}

// RoleShort is one role group below the raid's comp target: its name, its In count and the
// target.
type RoleShort struct {
	// Group is the raidspec group (raidspec.GroupHealer).
	Name, Group string
	Have, Want  int
}

// ShortRoles is p's role groups with fewer In than its comp target, in roster order; nil for
// a raid without a target and for a guild event.
func ShortRoles(p RaidPosted) []RoleShort {
	if p.Target == nil || p.NoRoles {
		return nil
	}
	var out []RoleShort
	for _, g := range rosterGroups {
		have, want := len(inGroup(p.Roster, g)), p.Target.of(g.group)
		if have < want {
			out = append(out, RoleShort{Name: g.name(), Group: g.group, Have: have, Want: want})
		}
	}
	return out
}

func inGroup(roster []RosterLine, g rosterGroup) []RosterLine {
	var out []RosterLine
	for _, r := range sortedState(roster, rosterIn) {
		if raidspec.Group(r.Class, r.Spec, r.Role) == g.group {
			out = append(out, r)
		}
	}
	return out
}

// sortedState is the rows in a state by class in the game's order, then name (spec §P.9.3).
func sortedState(roster []RosterLine, state string) []RosterLine {
	var out []RosterLine
	for _, r := range roster {
		if r.State == state {
			out = append(out, r)
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if a, b := raidspec.ClassRank(out[i].Class), raidspec.ClassRank(out[j].Class); a != b {
			return a < b
		}
		return strings.ToLower(out[i].Name) < strings.ToLower(out[j].Name)
	})
	return out
}

// groupFields is a group as inline fields: its header on the first, U+200B on a continuation,
// the lines balanced across them; empty, `None yet`.
func groupFields(header string, all []string, most int) []EmbedField {
	if len(all) == 0 {
		return []EmbedField{{Name: header, Value: rosterCopy.Empty, Inline: true}}
	}
	ls := capped(all, most)
	n := (len(ls) + fieldLines - 1) / fieldLines
	out := make([]EmbedField, 0, n)
	for i := range n {
		name := header
		if i > 0 {
			name = rosterCopy.Continued
		}
		part := ls[ceilDiv(i*len(ls), n):ceilDiv((i+1)*len(ls), n)]
		out = append(out, EmbedField{Name: name, Value: fitValue(part), Inline: true})
	}
	return out
}

func ceilDiv(a, b int) int { return (a + b - 1) / b }

// capped keeps at most max lines, the last becoming `and {n} more`.
func capped(ls []string, most int) []string {
	if len(ls) <= most {
		return ls
	}
	keep := most - 1
	return append(append([]string{}, ls[:keep]...), counted(rosterCopy.More, len(ls)-keep))
}

// fitValue joins lines, dropping tail lines into `and {n} more` while the value would pass 1024
// (only hostile names come near it).
func fitValue(ls []string) string {
	value := strings.Join(ls, "\n")
	for utf8.RuneCountInString(value) > fieldValueMax && len(ls) > 1 {
		ls = capped(ls, len(ls)-1)
		value = strings.Join(ls, "\n")
	}
	return value
}

func lines(rows []RosterLine, marks Marks) []string {
	out := make([]string, len(rows))
	for i, r := range rows {
		out[i] = rosterLine(r, marks)
	}
	return out
}

// rosterLine is `{emoji} {Name}` when the spec's emoji is loaded, else the text line.
func rosterLine(r RosterLine, marks Marks) string {
	name := rosterName(r.Name)
	sp, ok := raidspec.Resolve(r.Class, r.Spec, r.Role)
	if id := marks[sp.Emoji()]; ok && id != "" {
		return fmt.Sprintf(rosterCopy.EmojiFormat, sp.Emoji(), id, name)
	}
	return textLine(name, r.Class, sp, ok)
}

// textLine is `{Name} · {Spec} {Class}`, `{Name} · {Class}` with no spec, `{Name}` with no class.
func textLine(name, class string, sp raidspec.Spec, ok bool) string {
	switch {
	case ok:
		return name + rosterCopy.Sep + sp.FullName
	case class != "":
		return name + rosterCopy.Sep + classWord(class)
	}
	return name
}

func rosterName(raw string) string {
	return escapedWithin(collapse(raw), rosterNameRunes, rosterNameOutRunes)
}

// classWord is the class as the game names it: "hunter" reads "Hunter".
func classWord(class string) string {
	r, size := utf8.DecodeRuneInString(class)
	if size == 0 {
		return ""
	}
	return strings.ToUpper(string(r)) + strings.ToLower(class[size:])
}

// outNames are the Out rows' names, A to Z.
func outNames(roster []RosterLine) []string {
	var out []string
	for _, r := range roster {
		if r.State == rosterOut {
			out = append(out, rosterName(r.Name))
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return strings.ToLower(out[i]) < strings.ToLower(out[j]) })
	return out
}

// outValue joins the names by `, ` until the next would pass budget, then `, and {n} more`.
func outValue(names []string, budget int) string {
	value := ""
	for i, n := range names {
		next := n
		if i > 0 {
			next = value + rosterCopy.OutJoin + n
		}
		need := utf8.RuneCountInString(next)
		if rest := len(names) - i - 1; rest > 0 {
			need += utf8.RuneCountInString(rosterCopy.OutJoin + counted(rosterCopy.More, rest))
		}
		if need > budget {
			if i == 0 {
				return counted(rosterCopy.More, len(names))
			}
			return value + rosterCopy.OutJoin + counted(rosterCopy.More, len(names)-i)
		}
		value = next
	}
	return value
}

// counted fills a label's {n}: "Tanks · 4", "and 3 more".
func counted(label string, n int) string {
	return copygen.Fill(label, "{n}", strconv.Itoa(n))
}
