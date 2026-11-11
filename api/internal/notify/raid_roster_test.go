package notify

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// allMarks is an emoji id for every spec, as the bot holds them once setup has run.
func allMarks() Marks {
	m := Marks{}
	for i, sp := range raidspec.All() {
		m[sp.Emoji()] = fmt.Sprint(1301000000000000001 + i)
	}
	return m
}

func line(name, class, spec, role, state string) RosterLine {
	return RosterLine{Name: name, Class: class, Spec: spec, Role: role, State: state}
}

// rosterPost is the open post for these sign-ups, with the counts they make.
func rosterPost(roster []RosterLine, marks Marks) Embed {
	in := onyxiaThursday()
	in.Roster, in.Counts = roster, RaidCounts{}
	for _, r := range roster {
		switch r.State {
		case "in":
			in.Counts.In++
		case "late":
			in.Counts.Late++
		}
	}
	return RaidPostEmbed(RaidPost{RaidPosted: in, State: PostOpen, Marks: marks}, fixedFinish)
}

func fieldNames(e Embed) []string {
	out := make([]string, len(e.Fields))
	for i, f := range e.Fields {
		out[i] = f.Name
	}
	return out
}

func fieldValue(t *testing.T, e Embed, name string) EmbedField {
	t.Helper()
	for _, f := range e.Fields {
		if f.Name == name {
			return f
		}
	}
	t.Fatalf("no field %q in %v", name, fieldNames(e))
	return EmbedField{}
}

func TestRosterCopyMatchesTheSpec(t *testing.T) {
	want := map[string]string{
		"roster.tanks": rosterCopy.Tanks, "roster.healers": rosterCopy.Healers, "roster.melee": rosterCopy.Melee,
		"roster.ranged": rosterCopy.Ranged, "roster.late": rosterCopy.Late, "roster.bench": rosterCopy.Bench,
		"roster.out": rosterCopy.Out, "roster.empty": rosterCopy.Empty, "roster.more": rosterCopy.More,
	}
	for key, got := range want {
		cell := guidedCell(t, key)[0]
		if cell != got {
			t.Errorf("%s = %q, spec %q", key, got, cell)
		}
	}
	if got := guidedCell(t, "roster.outJoin")[0]; got != rosterCopy.OutJoin {
		t.Errorf("outJoin %q", got)
	}
	signed := guidedCell(t, "roster.signed")
	if signed[0] != raidCopy.SignedLead || strings.ReplaceAll(signed[1], "{n}", "{signed}") != raidCopy.SignedLeadOpen ||
		signed[2] != raidCopy.SignedNone {
		t.Errorf("roster.signed %q", signed)
	}
	text := guidedCell(t, "roster.textLine")
	holy, _ := raidspec.ByKey("paladin-holy")
	if got := textLine("Aurelith", "paladin", holy, true); got != strings.NewReplacer("{Name}", "Aurelith", "{Spec}", "Holy", "{Class}", "Paladin").Replace(text[0]) {
		t.Errorf("roster.textLine = %q, spec %q", got, text[0])
	}
	if got := textLine("Aurelith", "paladin", raidspec.Spec{}, false); got != strings.NewReplacer("{Name}", "Aurelith", "{Class}", "Paladin").Replace(text[len(text)-1]) {
		t.Errorf("roster.textLine no spec = %q, spec %q", got, text[len(text)-1])
	}
	if rosterCopy.Continued != "​" {
		t.Error("continuation name")
	}
}

func TestRosterFieldsFollowTheLayout(t *testing.T) {
	roster := []RosterLine{
		line("Kraggor", "warrior", "Protection", "tank", "in"),
		line("Zulmara", "hunter", "Beast Mastery", "dps", "in"),
		line("Corvane", "rogue", "", "dps", "in"),
		line("Pyrewisp", "mage", "", "dps", "in"),
		line("Kelthar", "hunter", "Survival", "dps", "in"),
		line("Glimmerfen", "warlock", "Destruction", "dps", "late"),
		line("Selwyn", "mage", "Frost", "dps", "bench"),
		line("ysolde", "priest", "Shadow", "dps", "out"),
		line("Garrowe", "warrior", "Arms", "dps", "out"),
	}
	e := rosterPost(roster, allMarks())
	want := "When,Signed up,Tanks · 1,Healers · 0,Melee · 2,Ranged · 2,Late · 1,Bench · 1,Out · 2"
	if got := strings.Join(fieldNames(e), ","); got != want {
		t.Fatalf("fields %s", got)
	}
	signed := fieldValue(t, e, "Signed up")
	if signed.Value != "**6 of 40**" || !signed.Inline {
		t.Fatalf("signed %+v", signed)
	}
	if v := fieldValue(t, e, "Healers · 0").Value; v != "None yet" {
		t.Fatalf("empty group %q", v)
	}
	melee := fieldValue(t, e, "Melee · 2").Value
	if melee != "<:fs_hunter_survival:1301000000000000007> Kelthar\nCorvane · Rogue" {
		t.Fatalf("melee %q", melee)
	}
	if v := fieldValue(t, e, "Ranged · 2").Value; v != "<:fs_hunter_beast_mastery:1301000000000000005> Zulmara\nPyrewisp · Mage" {
		t.Fatalf("ranged %q", v)
	}
	out := fieldValue(t, e, "Out · 2")
	if out.Value != "Garrowe, ysolde" || out.Inline {
		t.Fatalf("out %+v", out)
	}
}

func TestNobodyInOrLateDrawsNoGroups(t *testing.T) {
	e := rosterPost([]RosterLine{line("Garrowe", "warrior", "Arms", "dps", "out"), line("Selwyn", "mage", "Frost", "dps", "bench")}, allMarks())
	if got := strings.Join(fieldNames(e), ","); got != "When,Signed up,Out · 1" || e.Fields[1].Value != "None yet" {
		t.Fatalf("fields %s", got)
	}
	if got := strings.Join(fieldNames(rosterPost(nil, allMarks())), ","); got != "When,Signed up" {
		t.Fatalf("empty %s", got)
	}
}

func TestRosterSortsByClassThenNameAndMarksFallBack(t *testing.T) {
	marks := allMarks()
	delete(marks, "fs_mage_frost")
	roster := []RosterLine{
		line("zed", "mage", "Frost", "dps", "in"),
		line("Abe", "warlock", "Affliction", "dps", "in"),
		line("bea", "mage", "Fire", "dps", "in"),
		line("Nobody", "", "", "dps", "in"),
		line("Moon", "druid", "Balance", "dps", "in"),
	}
	got := fieldValue(t, rosterPost(roster, marks), "Ranged · 5").Value
	want := strings.Join([]string{
		"<:fs_mage_fire:1301000000000000009> bea",
		"zed · Frost Mage",
		"<:fs_warlock_affliction:1301000000000000023> Abe",
		"<:fs_druid_balance:1301000000000000001> Moon",
		"Nobody",
	}, "\n")
	if got != want {
		t.Fatalf("ranged\n%s\nwant\n%s", got, want)
	}
}

func repeated(n int, class, spec, role, state string) []RosterLine {
	out := make([]RosterLine, n)
	for i := range out {
		out[i] = line(fmt.Sprintf("Name%03d", i), class, spec, role, state)
	}
	return out
}

func TestRosterOverflow(t *testing.T) {
	e := rosterPost(repeated(17, "mage", "Frost", "dps", "in"), allMarks())
	names := fieldNames(e)
	if names[5] != "Ranged · 17" || names[6] != "\u200b" || !e.Fields[6].Inline ||
		strings.Count(e.Fields[5].Value, "\n") != 8 || strings.Count(e.Fields[6].Value, "\n") != 7 {
		t.Fatalf("17 lines: %v", names)
	}
	e = rosterPost(repeated(40, "mage", "Frost", "dps", "in"), allMarks())
	value := e.Fields[5].Value + "\n" + e.Fields[6].Value
	ls := strings.Split(value, "\n")
	if len(ls) != 32 || ls[31] != "and 9 more" || len(e.Fields) != 7 {
		t.Fatalf("40 lines: %d, last %q, %d fields", len(ls), ls[len(ls)-1], len(e.Fields))
	}
	bench := append(repeated(1, "mage", "Frost", "dps", "in"), repeated(11, "hunter", "Survival", "dps", "bench")...)
	ls = strings.Split(fieldValue(t, rosterPost(bench, allMarks()), "Bench · 11").Value, "\n")
	if len(ls) != 10 || ls[9] != "and 2 more" {
		t.Fatalf("bench %q", ls)
	}
}

// 120 sign-ups with hostile names keep every limit: values at most 1024, at most 25 fields,
// the embed under 6000, and no name pings anyone.
func TestRosterGuardHoldsDiscordsLimits(t *testing.T) {
	var roster []RosterLine
	states := []string{"in", "late", "bench", "out"}
	specs := raidspec.All()
	for i := range 120 {
		sp := specs[i%len(specs)]
		name := fmt.Sprintf("@everyone*_%d|`>#", i)
		roster = append(roster, line(name, sp.Class, sp.Name, sp.Role, states[i%len(states)]))
	}
	for _, marks := range []Marks{allMarks(), nil} {
		e := rosterPost(roster, marks)
		if len(e.Fields) > 25 || embedChars(e) > 6000 {
			t.Fatalf("%d fields, %d chars", len(e.Fields), embedChars(e))
		}
		for _, f := range e.Fields {
			if utf8.RuneCountInString(f.Value) > 1024 || utf8.RuneCountInString(f.Name) > 256 {
				t.Fatalf("field %q is %d long", f.Name, utf8.RuneCountInString(f.Value))
			}
			if strings.Contains(f.Value, "*_") {
				t.Fatalf("unescaped markdown in %q", f.Value)
			}
		}
	}
}

func TestOutListCutsAt1024(t *testing.T) {
	roster := repeated(150, "mage", "Frost", "dps", "out")
	e := rosterPost(roster, nil)
	out := fieldValue(t, e, "Out · 150").Value
	if utf8.RuneCountInString(out) > 1024 || !strings.HasSuffix(out, " more") || !strings.Contains(out, ", and ") {
		t.Fatalf("out %d %q", utf8.RuneCountInString(out), out[len(out)-30:])
	}
}

// The webhook reminder carries the post in the Tomorrow state with the roster in text lines.
func TestWebhookReminderEmbedCarriesTheTextRoster(t *testing.T) {
	in := onyxiaThursday()
	in.Roster = []RosterLine{line("Aurelith", "paladin", "Holy", "healer", "in")}
	in.Counts = RaidCounts{In: 1}
	e := WebhookReminderEmbed(in, fixedFinish)
	if !strings.HasPrefix(e.Title, "Tomorrow: ") || e.Description != "Sign up on Forever Sixty." {
		t.Fatalf("embed %q %q", e.Title, e.Description)
	}
	if v := fieldValue(t, e, "Healers · 1").Value; v != "Aurelith · Holy Paladin" {
		t.Fatalf("healers %q", v)
	}
}

// The states: Cancelled and Done carry no roster; Tomorrow, Locked, Moved and Ended keep it.
func TestRosterStates(t *testing.T) {
	in := onyxiaThursday()
	in.Roster = []RosterLine{line("Kraggor", "warrior", "Protection", "tank", "in")}
	in.Counts = RaidCounts{In: 1}
	was := in.StartsAt
	for _, state := range []RaidPostState{PostTomorrow, PostLocked, PostMoved, PostEnded, PostCancelled, PostDone} {
		e := RaidPostEmbed(RaidPost{RaidPosted: in, State: state, Was: &was, Marks: allMarks()}, fixedFinish)
		has := strings.Contains(strings.Join(fieldNames(e), ","), "Tanks · 1")
		if has == (state == PostCancelled || state == PostDone) {
			t.Errorf("%s roster %v: %v", state, has, fieldNames(e))
		}
	}
}

// A guild event's post lists no role groups and no specs (events spec §3.5): In, Late and Bench
// by name and class, then Out.
func TestEventPostIsOnePlainList(t *testing.T) {
	in := onyxiaThursday()
	in.NoRoles = true
	in.Roster = []RosterLine{line("Garrowe", "warrior", "Protection", "tank", "in"), line("Selwyn", "mage", "Frost", "dps", "in"),
		line("Pip", "priest", "Holy", "healer", "late"), line("Bo", "rogue", "", "dps", "bench"), line("Ash", "mage", "", "", "out")}
	in.Counts = RaidCounts{In: 2, Late: 1}
	e := RaidPostEmbed(RaidPost{RaidPosted: in, State: PostOpen, Marks: allMarks()}, fixedFinish)
	got := map[string]string{}
	for _, f := range e.Fields {
		got[f.Name] = f.Value
	}
	want := map[string]string{"In · 2": "Garrowe · Warrior\nSelwyn · Mage", "Late · 1": "Pip · Priest", "Bench · 1": "Bo · Rogue", "Out · 1": "Ash"}
	for name, value := range want {
		if got[name] != value {
			t.Errorf("%s = %q, want %q (fields %v)", name, got[name], value, fieldNames(e))
		}
	}
	for _, name := range fieldNames(e) {
		if strings.HasPrefix(name, "Tanks") || strings.HasPrefix(name, "Healers") || strings.HasPrefix(name, "Melee") || strings.HasPrefix(name, "Ranged") {
			t.Errorf("an event post has the role field %q", name)
		}
	}
}

func TestShortRolesNamesOnlyTheRolesBelowTarget(t *testing.T) {
	roster := []RosterLine{
		line("Brakka", "warrior", "Protection", "tank", "in"),
		line("Mirelle", "priest", "Holy", "healer", "in"),
		line("Dawnrae", "paladin", "Holy", "healer", "late"),
		line("Gorvash", "rogue", "Combat", "dps", "in"),
	}
	target := &RoleTarget{Tanks: 1, Healers: 3, Melee: 1, Ranged: 2}
	cases := []struct {
		name string
		in   RaidPosted
		want []RoleShort
	}{
		{"below target, Late not counted", RaidPosted{Roster: roster, Target: target},
			[]RoleShort{{Name: "Healers", Group: raidspec.GroupHealer, Have: 1, Want: 3}, {Name: "Ranged", Group: raidspec.GroupRanged, Have: 0, Want: 2}}},
		{"no target", RaidPosted{Roster: roster}, nil},
		{"guild event", RaidPosted{Roster: roster, Target: target, NoRoles: true}, nil},
		{"every role covered", RaidPosted{Roster: roster, Target: &RoleTarget{Tanks: 1, Healers: 1, Melee: 1}}, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ShortRoles(c.in); fmt.Sprint(got) != fmt.Sprint(c.want) {
				t.Fatalf("ShortRoles = %v, want %v", got, c.want)
			}
		})
	}
}
