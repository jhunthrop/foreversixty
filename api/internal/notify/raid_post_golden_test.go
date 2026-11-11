package notify

import (
	"testing"
	"time"
)

const lockedText = "Sign-ups are locked. An officer can still change yours."

func TestRaidPostStatesMatchTheCopyTable(t *testing.T) {
	post := func(state RaidPostState) RaidPost { return RaidPost{RaidPosted: onyxiaThursday(), State: state} }
	done := post(PostDone)
	done.Came, done.CameOf = 22, 23
	done.Kills = []KillLine{{Boss: "Onyxia", Duration: 102 * time.Second}}
	cases := map[string]struct {
		post RaidPost
		want string
	}{
		"open": {post(PostOpen), `{"title":"Onyxia's Lair · Thu 10 Dec",` +
			raidLink + `"fields":[` + raidWhen + `,` + raidSigned + `],` + raidFoot + `"color":37810}`},
		"locked": {post(PostLocked), `{"title":"Onyxia's Lair · Thu 10 Dec","description":"` + lockedText + `",` +
			raidLink + `"fields":[` + raidWhen + `,` + raidSigned + `],` + raidFoot + `"color":9144196}`},
		"ended": {post(PostEnded), `{"title":"Onyxia's Lair · Thu 10 Dec","description":"This raid has ended.",` +
			raidLink + `"fields":[` + raidWhen + `,` + raidSigned + `],` + raidFoot + `"color":9144196}`},
		"done": {done, `{"title":"Done: Onyxia's Lair · Thu 10 Dec","description":"Logged as a raid night.",` +
			raidLink + `"fields":[` + raidWhen + `,{"name":"Came","value":"22 of 23","inline":true},` +
			`{"name":"Kills","value":"Onyxia 1:42"}],` + raidFoot + `"color":8377482}`},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if got, want := wireOf(t, RaidPostEmbed(c.post, fixedFinish)), wireHead+c.want+wireTail; got != want {
				t.Fatalf("\n got %s\nwant %s", got, want)
			}
		})
	}
}

// The post carries one time cell, When (Discord's own stamp in each reader's clock), and its
// title the site's raid-row date: the year only when the night falls in another year.
func TestRaidPostHasOneTimeCellAndTheRaidRowDate(t *testing.T) {
	for _, state := range []RaidPostState{PostOpen, PostLocked, PostMoved, PostEnded} {
		in := onyxiaThursday()
		was := in.StartsAt
		e := RaidPostEmbed(RaidPost{RaidPosted: in, State: state, Was: &was}, fixedFinish)
		for _, f := range e.Fields {
			if f.Name == "Guild time" {
				t.Fatalf("%s carries a second time cell", state)
			}
		}
	}
	nextYear := onyxiaThursday()
	nextYear.StartsAt = time.Date(2027, 1, 14, 2, 0, 0, 0, time.UTC)
	if got := RaidPostEmbed(RaidPost{RaidPosted: nextYear, State: PostOpen}, fixedFinish).Title; got != "Onyxia's Lair · Wed 13 Jan 2027" {
		t.Fatalf("title = %q", got)
	}
}

func TestDoneWithNoKillsAndNoCame(t *testing.T) {
	e := RaidPostEmbed(RaidPost{RaidPosted: onyxiaThursday(), State: PostDone}, fixedFinish)
	if e.Fields[1].Value != "0 of 0" || e.Fields[2].Value != "No kills" {
		t.Fatalf("done fields = %+v", e.Fields)
	}
}
