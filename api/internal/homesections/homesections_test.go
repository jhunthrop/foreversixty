package homesections

import (
	"encoding/json"
	"os"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func types(ss []Section) string {
	out := make([]string, len(ss))
	for i, s := range ss {
		out[i] = s.Type
		if s.Style != "" {
			out[i] += ":" + s.Style
		}
	}
	return strings.Join(out, " ")
}

func TestSignedInRules(t *testing.T) {
	cases := []struct {
		name string
		view View
		o    Override
		want string
	}{
		{"60 in a guild (R2)", View{Level: 60, InGuild: true}, Override{},
			"head upgrades countdowns:full feed guild_sims spec_week movers next"},
		{"under 60: the countdown is one line after the feed (R4)", View{Level: 34, InGuild: true}, Override{},
			"head upgrades feed countdowns:line guild_sims spec_week movers next"},
		{"no guild: the one-line Join a guild row (R4)", View{Level: 60}, Override{},
			"head upgrades countdowns:full feed guild_sims_cta spec_week movers next"},
		{"raid day once raids open: countdowns second (R2)", View{Level: 34, InGuild: true, RaidsOpen: true, RaidDay: true}, Override{},
			"head countdowns:line upgrades feed guild_sims spec_week movers next"},
		{"raid day before raids open changes nothing", View{Level: 60, InGuild: true, RaidDay: true}, Override{},
			"head upgrades countdowns:full feed guild_sims spec_week movers next"},
		{"rules switched off", View{Level: 34, RaidsOpen: true, RaidDay: true},
			Override{Rules: map[string]bool{RuleUnder60Line: false, RuleNoGuildCTA: false}},
			"head upgrades countdowns:full feed guild_sims spec_week movers next"},
		{"off, order and a notice", View{Level: 60, InGuild: true},
			Override{Off: []string{Movers}, Order: []string{Feed, Head}, Notice: &NoticeOverride{Text: "Servers restart at 10."}},
			"feed head notice upgrades countdowns:full guild_sims spec_week next"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := types(SignedIn(c.view, c.o)); got != c.want {
				t.Errorf("got  %s\nwant %s", got, c.want)
			}
		})
	}
}

func TestSignedInFeedAndLimits(t *testing.T) {
	in := SignedIn(View{Level: 60, InGuild: true}, Override{})
	feed := in[slices.IndexFunc(in, func(s Section) bool { return s.Type == Feed })]
	if !reflect.DeepEqual(feed.Filters, []string{"all", "spec", "guild"}) || !reflect.DeepEqual(feed.MoreFilters, []string{"logs"}) || feed.Limit != 7 {
		t.Errorf("feed = %+v", feed)
	}
	if in[1].Limit != 3 {
		t.Errorf("upgrades limit = %d", in[1].Limit)
	}
	out := SignedIn(View{Level: 60}, Override{})
	feed = out[slices.IndexFunc(out, func(s Section) bool { return s.Type == Feed })]
	if slices.Contains(feed.Filters, FilterGuild) {
		t.Errorf("no guild still filters by guild: %v", feed.Filters)
	}
}

func TestParseOverride(t *testing.T) {
	if o, err := ParseOverride(""); err != nil || !reflect.DeepEqual(o, Override{}) {
		t.Fatalf("empty = %+v, %v", o, err)
	}
	o, err := ParseOverride(`{"rules":{"raid_day_first":false},"off":["movers"],"notice":{"text":"Hi","href":"/guild"}}`)
	if err != nil || o.ruleOn(RuleRaidDayFirst) || !o.ruleOn(RuleUnder60Line) {
		t.Fatalf("parsed %+v, %v", o, err)
	}
	for _, bad := range []string{
		`[]`, `{"colour":1}`, `{"rules":{"nope":true}}`, `{"off":["hero"]}`, `{"order":["head_out"]}`,
		`{"notice":{"text":""}}`, `{"notice":{"text":"x","href":"https://evil"}}`, `{"signed_out":[{"type":""}]}`,
	} {
		if _, err := ParseOverride(bad); err == nil {
			t.Errorf("%s parsed", bad)
		}
	}
}

func TestSignedOut(t *testing.T) {
	if got := types(SignedOut(Override{})); got != "head_out class_picker countdowns:full feed how_to_start_link" {
		t.Errorf("bundled = %s", got)
	}
	custom := []Section{{Type: HeadOut}}
	if got := SignedOut(Override{SignedOut: custom}); !reflect.DeepEqual(got, custom) {
		t.Errorf("override = %+v", got)
	}
}

// The API's copy is the apps' bundled file (apps/shared/home-sections.json), byte for byte.
func TestSignedOutMatchesTheAppsFile(t *testing.T) {
	shared, err := os.ReadFile("../../../apps/shared/home-sections.json")
	if err != nil {
		t.Fatal(err)
	}
	if string(shared) != string(signedOutJSON) {
		t.Error("api/internal/homesections/signed_out.json differs from apps/shared/home-sections.json; copy the apps' file over it")
	}
	var probe map[string]any
	if err := json.Unmarshal(shared, &probe); err != nil {
		t.Fatal(err)
	}
}
