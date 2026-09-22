package addon

import "testing"

func TestParseFS1GuildReadsTheGuildSection(t *testing.T) {
	cases := []struct {
		name     string
		export   string
		wantName string
		wantRank int
		wantOK   bool
	}{
		{
			name:     "valid guild section after professions",
			export:   "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:2",
			wantName: "Iron Vanguard", wantRank: 2, wantOK: true,
		},
		{
			name:   "non-numeric rank refuses the section",
			export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:officer",
			wantOK: false,
		},
		{
			name:   "no guild section at all",
			export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|professions=mining,herbalism",
			wantOK: false,
		},
		{
			name:   "empty export",
			export: "",
			wantOK: false,
		},
		{
			name:     "guild master is rank index 0",
			export:   "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:0",
			wantName: "Iron Vanguard", wantRank: 0, wantOK: true,
		},
		{
			name:   "missing colon refuses the section",
			export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=IronVanguard",
			wantOK: false,
		},
		{
			name:     "guild section among several others",
			export:   "FS1:1.60.1.69893:warrior:tauren:0/0/0:|bags=1,2|sets=a=x|guild=Forever:1|loadouts=b=y",
			wantName: "Forever", wantRank: 1, wantOK: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			name, rank, ok := ParseFS1Guild(c.export)
			if ok != c.wantOK {
				t.Fatalf("ok = %v, want %v", ok, c.wantOK)
			}
			if !ok {
				return
			}
			if name != c.wantName || rank != c.wantRank {
				t.Fatalf("got (%q, %d), want (%q, %d)", name, rank, c.wantName, c.wantRank)
			}
		})
	}
}

func TestParseFS1ClassReadsTheHeadsClassSlug(t *testing.T) {
	cases := []struct {
		name      string
		export    string
		wantClass string
		wantOK    bool
	}{
		{
			name: "plain FS1 head", export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:",
			wantClass: "warrior", wantOK: true,
		},
		{
			name: "head with sections after it", export: "FS1:1.60.1.69893:mage:human:0/0/0:|guild=Forever:1",
			wantClass: "mage", wantOK: true,
		},
		{name: "not FS1", export: "FS2:1.60.1.69893:warrior:tauren:0/0/0:", wantOK: false},
		{name: "too short", export: "FS1:1.60.1.69893", wantOK: false},
		{name: "empty export", export: "", wantOK: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			class, ok := ParseFS1Class(c.export)
			if ok != c.wantOK {
				t.Fatalf("ok = %v, want %v", ok, c.wantOK)
			}
			if ok && class != c.wantClass {
				t.Fatalf("class = %q, want %q", class, c.wantClass)
			}
		})
	}
}
