package guilds

import "testing"

func specWord(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

// TestResolveSpecTable pins W3-C16's words: a full spec name with the class, the catalogue's
// role, and nothing for a spec the catalogue does not name or a row with no spec.
func TestResolveSpecTable(t *testing.T) {
	cases := []struct {
		name               string
		class, spec        *string
		role               string
		wantName, wantRole string
		ok                 bool
	}{
		{"a specs.json key", nil, strp("paladin-holy"), "", "Holy Paladin", "healer", true},
		{"a log name within its class", strp("warrior"), strp("Protection"), "", "Protection Warrior", "tank", true},
		{"a logged Feral tank is Feral Bear", strp("druid"), strp("Feral"), "tank", "Feral Druid (Bear)", "tank", true},
		{"a logged Feral cat", strp("druid"), strp("Feral"), "dps", "Feral Druid", "dps", true},
		{"a log name with no class", nil, strp("Holy"), "", "", "", false},
		{"no spec", strp("priest"), nil, "", "", "", false},
		{"an unknown word", strp("mage"), strp("Spellblade"), "", "", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := resolveSpec(tc.class, tc.spec, tc.role)
			if ok != tc.ok || got.Name != tc.wantName || got.Role != tc.wantRole {
				t.Fatalf("resolveSpec = %+v %v, want %q %q %v", got, ok, tc.wantName, tc.wantRole, tc.ok)
			}
			if ok && got.Icon == "" {
				t.Error("a resolved spec has an icon")
			}
		})
	}
}

// A members[] entry takes its own (snapshot) spec first, else its roster row's, and an entry
// whose consent withholds both stays null.
func TestDescribeMembers(t *testing.T) {
	rows := []RosterRow{
		{CharacterKey: "us/pvp/kraggor", Class: strp("warrior"), Spec: strp("Protection"), Role: strp("tank")},
		{CharacterKey: "us/pvp/hidden"},
	}
	describeRosterRows(rows)
	if specWord(rows[0].SpecName) != "Protection Warrior" || specWord(rows[0].SpecIcon) == "<nil>" {
		t.Fatalf("roster row = %s %s", specWord(rows[0].SpecName), specWord(rows[0].SpecIcon))
	}
	if rows[1].SpecName != nil || rows[1].Role != nil {
		t.Fatal("a withheld row stays empty")
	}
	entries := []RosterEntry{
		{Name: "Mendara", CharacterKey: "us/pvp/mendara", Spec: strp("priest-holy")},
		{Name: "Kraggor", CharacterKey: "us/pvp/kraggor", Class: strp("warrior")},
		{Name: "Hidden", CharacterKey: "us/pvp/hidden"},
		{Name: "Strayhound", State: EntryNoAccount},
	}
	describeMembers(entries, rows)
	want := []struct{ name, role string }{
		{"Holy Priest", "healer"}, {"Protection Warrior", "tank"}, {"<nil>", "<nil>"}, {"<nil>", "<nil>"},
	}
	for i, w := range want {
		if specWord(entries[i].SpecName) != w.name || specWord(entries[i].Role) != w.role {
			t.Errorf("%s = %s %s, want %s %s", entries[i].Name, specWord(entries[i].SpecName), specWord(entries[i].Role), w.name, w.role)
		}
	}
}

// Each roster row carries the raid group it fights in (raidspec.Group), the rule the raid page's
// sign-ups use: a Survival hunter is melee, a Feral tank a tank, a row with no spec none.
func TestDescribeRosterRowsGroup(t *testing.T) {
	cases := []struct {
		name              string
		class, spec, role *string
		want              string
	}{
		{"a tank", strp("warrior"), strp("Protection"), strp("tank"), "tank"},
		{"a healer", strp("priest"), strp("Holy"), nil, "healer"},
		{"a Survival hunter fights in melee", strp("hunter"), strp("Survival"), strp("dps"), "melee"},
		{"a Marksmanship hunter is ranged", strp("hunter"), strp("Marksmanship"), strp("dps"), "ranged"},
		{"a Feral cat is melee", strp("druid"), strp("Feral"), strp("dps"), "melee"},
		{"a withheld spec has no group", nil, nil, nil, "<nil>"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rows := []RosterRow{{CharacterKey: "us/pvp/x", Class: tc.class, Spec: tc.spec, Role: tc.role}}
			describeRosterRows(rows)
			if got := specWord(rows[0].Group); got != tc.want {
				t.Fatalf("group = %s, want %s", got, tc.want)
			}
		})
	}
}
