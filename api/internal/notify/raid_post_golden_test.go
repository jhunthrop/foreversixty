package notify

import (
	"testing"
	"time"
)

const (
	raidGuildTime = `{"name":"Guild time","value":"8:00 pm CST","inline":true}`
	lockedText    = "Sign-ups are locked. An officer can still change yours."
)

func TestRaidPostStatesMatchTheCopyTable(t *testing.T) {
	post := func(state RaidPostState) RaidPost { return RaidPost{RaidPosted: onyxiaThursday(), State: state} }
	done := post(PostDone)
	done.Came, done.CameOf = 22, 23
	done.Kills = []KillLine{{Boss: "Onyxia", Duration: 102 * time.Second}}
	cases := map[string]struct {
		post RaidPost
		want string
	}{
		"open": {post(PostOpen), `{"title":"Onyxia's Lair · Thursday 10 December","description":"Sign up with the buttons below.",` +
			raidLink + `"fields":[` + raidWhen + `,` + raidGuildTime + `,` + raidSigned + `],` + raidFoot + `"color":15055189}`},
		"locked": {post(PostLocked), `{"title":"Onyxia's Lair · Thursday 10 December","description":"` + lockedText + `",` +
			raidLink + `"fields":[` + raidWhen + `,` + raidGuildTime + `,` + raidSigned + `],` + raidFoot + `"color":10130564}`},
		"ended": {post(PostEnded), `{"title":"Onyxia's Lair · Thursday 10 December","description":"This raid has ended.",` +
			raidLink + `"fields":[` + raidWhen + `,` + raidGuildTime + `,` + raidSigned + `],` + raidFoot + `"color":10130564}`},
		"done": {done, `{"title":"Done: Onyxia's Lair · Thursday 10 December","description":"Logged as a raid night.",` +
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

func TestBotStatesOmitGuildTimeWithoutAZone(t *testing.T) {
	for _, state := range []RaidPostState{PostOpen, PostLocked, PostMoved, PostEnded} {
		in := onyxiaThursday()
		in.Timezone = ""
		was := in.StartsAt
		for _, f := range RaidPostEmbed(RaidPost{RaidPosted: in, State: state, Was: &was}, fixedFinish).Fields {
			if f.Name == raidCopy.GuildTimeName {
				t.Fatalf("%s carries Guild time without a zone", state)
			}
		}
	}
}

func TestDoneWithNoKillsAndNoCame(t *testing.T) {
	e := RaidPostEmbed(RaidPost{RaidPosted: onyxiaThursday(), State: PostDone}, fixedFinish)
	if e.Fields[1].Value != "0 of 0" || e.Fields[2].Value != "No kills" {
		t.Fatalf("done fields = %+v", e.Fields)
	}
}
