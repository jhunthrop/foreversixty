package homesections

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
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

var (
	chicago, _ = time.LoadLocation("America/Chicago")
	// 3:00 pm in Chicago, Wednesday 9 Dec 2026.
	wed3pm = time.Date(2026, 12, 9, 21, 0, 0, 0, time.UTC)
	// 8:00 pm in Chicago that day.
	raid8pm = time.Date(2026, 12, 10, 2, 0, 0, 0, time.UTC)
)

func at(t time.Time) *time.Time { return &t }

func TestMoments(t *testing.T) {
	guild := View{Level: 60, InGuild: true, Zone: chicago}
	with := func(now time.Time, next, last *time.Time) View {
		v := guild
		v.Now, v.NextRaid, v.LastRaid = now, next, last
		return v
	}
	cases := []struct {
		name string
		view View
		want Moment
	}{
		{"raid day, before the raid", with(wed3pm, at(raid8pm), nil), MomentRaidDay},
		{"the raid started 1 h ago", with(raid8pm.Add(time.Hour), at(raid8pm), at(raid8pm)), MomentRaidLive},
		{"4 h after an 8 pm start it is midnight in Chicago: the day after", with(raid8pm.Add(4*time.Hour), at(raid8pm), at(raid8pm)), MomentDayAfter},
		{"3 h 59 after the start it is still live", with(raid8pm.Add(4*time.Hour-time.Minute), at(raid8pm), at(raid8pm)), MomentRaidLive},
		{"the next day in Chicago", with(raid8pm.Add(18*time.Hour), nil, at(raid8pm)), MomentDayAfter},
		{"two days later", with(raid8pm.Add(42*time.Hour), nil, at(raid8pm)), MomentGuild},
		{"a raid tomorrow is not raid day", with(wed3pm.Add(-24*time.Hour), at(raid8pm), nil), MomentGuild},
		{"8 pm Chicago is 2 am UTC the next day; raid day follows the guild's zone", with(wed3pm, at(raid8pm), nil), MomentRaidDay},
		{"under 60 with no raid", View{Level: 34, InGuild: true, Now: wed3pm}, MomentLeveling},
		{"under 60 on raid day still answers the raid", View{Level: 34, InGuild: true, Now: wed3pm, Zone: chicago, NextRaid: at(raid8pm)}, MomentRaidDay},
		{"no guild at 60", View{Level: 60, Now: wed3pm}, MomentNoGuild},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := MomentOf(c.view, Override{}); got != c.want {
				t.Errorf("got %s, want %s", got, c.want)
			}
		})
	}
}

func TestSignedInOrders(t *testing.T) {
	cases := []struct {
		name string
		view View
		o    Override
		want string
	}{
		{"raid day: the next raid leads", View{Level: 60, InGuild: true, Now: wed3pm, Zone: chicago, NextRaid: at(raid8pm)}, Override{},
			"head next_raid:card upgrades activity"},
		{"raid live: the raid and the raid activity", View{Level: 60, InGuild: true, Now: raid8pm.Add(time.Hour), Zone: chicago, NextRaid: at(raid8pm)}, Override{},
			"head raid_live activity"},
		{"the day after: the summary", View{Level: 60, InGuild: true, Now: raid8pm.Add(18 * time.Hour), Zone: chicago, LastRaid: at(raid8pm)}, Override{},
			"head raid_summary upgrades activity"},
		{"any other day in a guild", View{Level: 60, InGuild: true, Now: wed3pm}, Override{},
			"head upgrades next_raid:line activity"},
		{"no guild at 60: find one", View{Level: 60, Now: wed3pm}, Override{}, "head upgrades activity guild_cta"},
		{"leveling", View{Level: 34, Now: wed3pm}, Override{}, "head upgrades activity"},
		{"find_a_guild switched off", View{Level: 60, Now: wed3pm}, Override{Rules: map[string]bool{RuleFindAGuild: false}}, "head upgrades activity"},
		{"raid moments switched off", View{Level: 60, InGuild: true, Now: wed3pm, Zone: chicago, NextRaid: at(raid8pm)},
			Override{Rules: map[string]bool{RuleRaidMoments: false}}, "head upgrades next_raid:line activity"},
		{"off, order and a notice", View{Level: 60, InGuild: true, Now: wed3pm},
			Override{Off: []string{Activity}, Order: []string{NextRaid}, Notice: &NoticeOverride{Text: "Servers restart at 10."}},
			"next_raid:line head notice upgrades"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := types(SignedIn(c.view, c.o)); got != c.want {
				t.Errorf("got  %s\nwant %s", got, c.want)
			}
		})
	}
}

func TestRaidLivePreviewsTheRaidActivity(t *testing.T) {
	out := SignedIn(View{Level: 60, InGuild: true, Now: raid8pm.Add(time.Hour), NextRaid: at(raid8pm)}, Override{})
	if out[2].Type != Activity || out[2].Group != "raids" || out[2].Limit != 3 {
		t.Errorf("activity = %+v", out[2])
	}
}

func TestParseOverride(t *testing.T) {
	if o, err := ParseOverride(""); err != nil || !reflect.DeepEqual(o, Override{}) {
		t.Fatalf("empty = %+v, %v", o, err)
	}
	o, err := ParseOverride(`{"rules":{"raid_moments":false},"off":["activity"],"notice":{"text":"Hi","href":"/guild"}}`)
	if err != nil || o.ruleOn(RuleRaidMoments) || !o.ruleOn(RuleLeveling) {
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
