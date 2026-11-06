package discordbot

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func (w *pressWorld) raidCmdAs(t *testing.T, discordUser string) Reply {
	t.Helper()
	rep, err := w.svc.cmdRaid(context.Background(), interactionOf(typeCommand, discordUser, map[string]any{"name": "raid"}))
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func TestRaidCommandStates(t *testing.T) {
	segment := func(i int) string { return specString(t, commandHeading, "`/raid`", i, nil) }
	if segment(0) != "{Name} is In as {Role}." || segment(5) != commandCopy.RaidNone {
		t.Fatalf("spec segments moved: %q %q", segment(0), segment(5))
	}
	cases := []struct {
		state, want string
	}{
		{"in", "Zulmara is In as DPS."},
		{"late", "Zulmara is Late as DPS."},
		{"bench", "Zulmara is on the Bench."},
		{"out", "Zulmara is Out."},
	}
	for _, c := range cases {
		t.Run(c.state, func(t *testing.T) {
			w := newPressWorld(t)
			w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
			w.press(t, c.state)
			rep := w.raidCmdAs(t, pressDiscord)
			m := rep.Message
			if !m.Ephemeral || rep.Update || len(m.Embeds) != 1 || m.Embeds[0].Description != c.want {
				t.Fatalf("reply = %+v", rep)
			}
			assertRowMatchesPost(t, w, m)
		})
	}
	t.Run("unsigned", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		rep := w.raidCmdAs(t, pressDiscord)
		if got := rep.Message.Embeds[0].Description; got != segment(4) {
			t.Fatalf("description = %q", got)
		}
	})
	t.Run("unlinked", func(t *testing.T) {
		w := newPressWorld(t)
		wantReply(t, w.raidCmdAs(t, "d-stranger"), segment(6), false, wantButton{label: linkCopy.LinkButton, url: w.svc.linkStartURL()})
	})
	t.Run("non-member", func(t *testing.T) {
		w := newPressWorld(t)
		rep := w.raidCmdAs(t, pressDiscord)
		if !strings.HasPrefix(rep.Message.Content, "Sign-ups are for members of Olympus XXVII.") || len(rep.Message.Embeds) != 0 {
			t.Fatalf("reply = %+v", rep)
		}
	})
	t.Run("waiting for approval", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", false)
		rep := w.raidCmdAs(t, pressDiscord)
		want := fill(pressCopy.WaitingOne, "{Name}", "Zulmara")
		if rep.Message.Content != want || len(rep.Message.Embeds) != 0 {
			t.Fatalf("reply = %+v, want %q and no raid", rep, want)
		}
	})
	t.Run("none planned", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.exec(t, `delete from guild_raids`)
		wantReply(t, w.raidCmdAs(t, pressDiscord), segment(5), false)
	})
}

// addZoneNight posts a night of zone (a dungeon run or guild event) at starts.
func (w *pressWorld) addZoneNight(t *testing.T, zone, title string, starts time.Time) int64 {
	t.Helper()
	var id int64
	if err := w.st.Pool.QueryRow(context.Background(), `insert into guild_raids
		(guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, $2, $3, $4, 120, $5) returning id`, w.gid, zone, title, starts.Truncate(time.Minute), w.uid).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// /raid names the next raid, the one Readiness names, even with a run sooner; with no raid
// planned it shows the next run or event, led by a line that says which.
func TestRaidCommandPrefersTheNextRaid(t *testing.T) {
	t.Run("a run sooner than the raid", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.addZoneNight(t, "dungeon:ragefire-chasm", "Ragefire Chasm", time.Now().Add(2*time.Hour))
		rep := w.raidCmdAs(t, pressDiscord)
		if rep.Message.Content != "" || !strings.HasPrefix(rep.Message.Embeds[0].Title, "Onyxia's Lair") {
			t.Fatalf("reply = %q %q, want the raid", rep.Message.Content, rep.Message.Embeds[0].Title)
		}
	})
	for _, c := range []struct{ zone, title, word string }{
		{"dungeon:ragefire-chasm", "Ragefire Chasm", "run"},
		{"event:guild", "Guild meeting", "event"},
	} {
		t.Run(c.word, func(t *testing.T) {
			w := newPressWorld(t)
			w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
			w.exec(t, `delete from guild_raids`)
			w.addZoneNight(t, c.zone, c.title, time.Now().Add(2*time.Hour))
			rep := w.raidCmdAs(t, pressDiscord)
			want := fill(commandCopy.RaidFallback, "{kind}", c.word)
			if rep.Message.Content != want || !strings.HasPrefix(rep.Message.Embeds[0].Title, c.title) {
				t.Fatalf("reply = %q %q, want %q over %s", rep.Message.Content, rep.Message.Embeds[0].Title, want, c.title)
			}
			if got := specString(t, commandHeading, "`/raid`", 11, map[string]string{"kind": c.word}); got != want {
				t.Fatalf("spec = %q, want %q", got, want)
			}
		})
	}
}

// assertRowMatchesPost: the answer's embed and row are the raid post's own.
func assertRowMatchesPost(t *testing.T, w *pressWorld, m Message) {
	t.Helper()
	facts, _, err := w.gs.RaidPostFacts(context.Background(), w.raidID)
	if err != nil {
		t.Fatal(err)
	}
	post := PostMessage(postFor(facts, PostState(facts.Status, facts.Post.StartsAt, facts.EndsAt, nil, w.svc.now()), nil),
		facts.RaidURL, facts.NightURL, w.svc.now())
	if len(m.Components) != len(post.Components) || len(m.Components[0]) != len(post.Components[0]) {
		t.Fatalf("row = %+v, post row = %+v", m.Components, post.Components)
	}
	for i, c := range post.Components[0] {
		if !reflect.DeepEqual(m.Components[0][i], c) {
			t.Fatalf("button %d = %+v, want %+v", i, m.Components[0][i], c)
		}
	}
	if m.Embeds[0].Title != post.Embeds[0].Title || m.Embeds[0].Footer != post.Embeds[0].Footer {
		t.Fatalf("embed = %+v", m.Embeds[0])
	}
}

func TestRaidCommandShowsAMovedRaidAsItsPostDoes(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	ctx := context.Background()
	if _, err := w.st.EnqueuePost(ctx, w.raidID, w.gid); err != nil {
		t.Fatal(err)
	}
	w.st.MarkRaidMoved(ctx, w.raidID, time.Now().Add(24*time.Hour))
	e := w.raidCmdAs(t, pressDiscord).Message.Embeds[0]
	if !strings.HasPrefix(e.Title, "Moved: ") {
		t.Fatalf("title = %q", e.Title)
	}
}

// The nameless sentences are the /raid row's segments 7 to 10, read back from the spec.
func TestRaidCommandNamesNoCharacterItCannotFind(t *testing.T) {
	nameless := func(i int) string {
		return specString(t, commandHeading, "`/raid`", i, map[string]string{"Role": "DPS"})
	}
	for i, state := range []string{"in", "late", "bench", "out"} {
		t.Run(state, func(t *testing.T) {
			w := newPressWorld(t)
			w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
			w.press(t, state)
			w.port.candidates = []guilds.SignupCandidate{{CharacterKey: "us/normal/other", Name: "Other", Verified: true}}
			got := w.raidCmdAs(t, pressDiscord).Message.Embeds[0].Description
			if want := nameless(7 + i); got != want {
				t.Fatalf("description = %q, want %q", got, want)
			}
		})
	}
}
