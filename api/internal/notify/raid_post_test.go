package notify

import (
	"strings"
	"testing"
	"time"
)

func TestSignedValueIsTheBoldCountAlone(t *testing.T) {
	cases := []struct {
		name      string
		c         RaidCounts
		size      int
		sizeKnown bool
		want      string
	}{
		{"full", RaidCounts{In: 20, Late: 3, Tanks: 3, Healers: 6, DPS: 12, Bench: 3, Out: 3}, 40, true, "**23 of 40**"},
		{"unknown size", RaidCounts{In: 23, Tanks: 3}, 0, false, "**23 signed up**"},
		{"only bench", RaidCounts{Bench: 2, Out: 1}, 10, true, "None yet"},
		{"none", RaidCounts{}, 10, true, "None yet"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := SignedValue(c.c, c.size, c.sizeKnown); got != c.want {
				t.Fatalf("got %q want %q", got, c.want)
			}
		})
	}
}

func TestRaidPostEmbedStates(t *testing.T) {
	was := onyxiaThursday().StartsAt
	moved := onyxiaThursday()
	moved.StartsAt = was.Add(24 * time.Hour)
	plain := []string{"When", "Guild time", "Signed up"}
	cases := []struct {
		state  RaidPostState
		post   RaidPosted
		color  int
		prefix string
		desc   string
		fields []string
	}{
		{PostOpen, onyxiaThursday(), 15055189, "", "Sign up with the buttons below.", plain},
		{PostLocked, onyxiaThursday(), 10130564, "", "Sign-ups are locked. An officer can still change yours.", plain},
		{PostMoved, moved, 15055189, "Moved: ", "Sign-ups carry over. Check you can still make it.", []string{"When", "Was", "Guild time", "Signed up"}},
		{PostCancelled, onyxiaThursday(), 15758187, "Cancelled: ", "This raid was cancelled.", []string{"When", "Signed up"}},
		{PostEnded, onyxiaThursday(), 10130564, "", "This raid has ended.", plain},
		{PostDone, onyxiaThursday(), 8377482, "Done: ", "Logged as a raid night.", []string{"When", "Came", "Kills"}},
	}
	for _, c := range cases {
		t.Run(string(c.state), func(t *testing.T) {
			e := RaidPostEmbed(RaidPost{RaidPosted: c.post, State: c.state, Was: &was, Came: 22, CameOf: 23,
				Kills: []KillLine{{Boss: "Onyxia", Duration: 102 * time.Second}}}, fixedFinish)
			if e.Color != c.color || e.Description != c.desc || !strings.HasPrefix(e.Title, c.prefix) {
				t.Fatalf("embed = %+v", e)
			}
			var names []string
			for _, f := range e.Fields {
				names = append(names, f.Name)
			}
			if strings.Join(names, ",") != strings.Join(c.fields, ",") {
				t.Fatalf("fields = %v want %v", names, c.fields)
			}
			if c.state == PostDone && (e.Fields[1].Value != "22 of 23" || e.Fields[2].Value != "Onyxia 1:42") {
				t.Fatalf("done fields = %+v", e.Fields)
			}
		})
	}
}

func TestDoneKillsCapAtSix(t *testing.T) {
	kills := make([]KillLine, 8)
	for i := range kills {
		kills[i] = KillLine{Boss: "Boss", Duration: time.Minute}
	}
	e := RaidPostEmbed(RaidPost{RaidPosted: onyxiaThursday(), State: PostDone, Kills: kills}, fixedFinish)
	lines := strings.Split(e.Fields[2].Value, "\n")
	if len(lines) != 7 || lines[6] != "+2 more" {
		t.Fatalf("kills = %q", lines)
	}
}

func TestWebhookPostListsTheRosterAsText(t *testing.T) {
	in := onyxiaThursday()
	in.Counts = RaidCounts{In: 1, Late: 1, Out: 1}
	in.Roster = []RosterLine{
		{Name: "Aurelith", Class: "paladin", Spec: "Holy", Role: "healer", State: "in"},
		{Name: "Hollowmere", Class: "rogue", Role: "dps", State: "late"},
		{Name: "Ysolde", Class: "priest", Spec: "Shadow", Role: "dps", State: "out"},
	}
	e := RaidPostedEmbed(in, fixedFinish)
	byName := map[string]string{}
	for _, f := range e.Fields {
		byName[f.Name] = f.Value
	}
	if e.Description != "Sign up on Forever Sixty." || byName["Signed up"] != "**2 of 40**" ||
		byName["Healers · 1"] != "Aurelith · Holy Paladin" || byName["Late · 1"] != "Hollowmere · Rogue" ||
		byName["Out · 1"] != "Ysolde" || byName["Tanks · 0"] != "None yet" {
		t.Fatalf("webhook post = %+v", e.Fields)
	}
}
