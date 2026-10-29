package discordbot

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

func (w *pressWorld) signupCmd(t *testing.T, options ...map[string]any) Reply {
	t.Helper()
	rep, err := w.svc.cmdSignup(context.Background(), interactionOf(typeCommand, pressDiscord,
		map[string]any{"name": "signup", "options": options}))
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func opt(name, value string) map[string]any { return map[string]any{"name": name, "value": value} }

func (w *pressWorld) signedRaid(t *testing.T, key string) int64 {
	t.Helper()
	var id int64
	if err := w.st.Pool.QueryRow(context.Background(),
		`select raid_id from guild_raid_signups where character_key = $1`, key).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func TestSignupCommand(t *testing.T) {
	t.Run("no raid option signs up for the next planned raid", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.addRaid(t, w.gid, time.Now().Add(96*time.Hour))
		w.addRaid(t, w.gid, time.Now().Add(-5*time.Hour))
		rep := w.signupCmd(t, opt(optStatus, "in"))
		if !strings.HasPrefix(rep.Message.Content, "Zulmara is In as DPS for Onyxia's Lair, "+w.day()+".") || !rep.Message.Ephemeral {
			t.Fatalf("reply = %+v", rep)
		}
		if got := w.signedRaid(t, keyZulmara); got != w.raidID {
			t.Fatalf("signed raid %d, want the next one %d", got, w.raidID)
		}
	})
	t.Run("none planned", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.exec(t, `delete from guild_raids`)
		wantReply(t, w.signupCmd(t, opt(optStatus, "in")), specString(t, pressHeading, "`/raid`", 5, nil), false)
	})
	t.Run("a non-member learns nothing of the calendar", func(t *testing.T) {
		w := newPressWorld(t)
		w.exec(t, `delete from guild_raids`)
		rep := w.signupCmd(t, opt(optStatus, "in"))
		if !strings.HasPrefix(rep.Message.Content, "Sign-ups are for members of Olympus XXVII.") {
			t.Fatalf("reply = %q", rep.Message.Content)
		}
	})
	t.Run("a named raid's state is told only to a member", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", false)
		w.exec(t, `update guild_raids set status = 'cancelled' where id = $1`, w.raidID)
		want := spec(t, "Press: refusals", 5, map[string]string{"Name": "Duskstrider"})
		wantReply(t, w.signupCmd(t, opt(optStatus, "in"), opt(optRaid, fmt.Sprint(w.raidID))), want, false)
		wantReply(t, w.signupCmd(t, opt(optStatus, "in"), opt(optRaid, "999999")), want, false)
	})
	t.Run("the character option skips the select", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
		rep := w.signupCmd(t, opt(optStatus, "bench"), opt(optRaid, fmt.Sprint(w.raidID)), opt(optCharacter, keyDusk))
		wantReply(t, rep, "Duskstrider is on the Bench for Onyxia's Lair, "+w.day()+".", false, changeSpecButton(w, keyDusk))
		rep = w.signupCmd(t, opt(optStatus, "out"), opt(optCharacter, "zulmara"))
		wantReply(t, rep, "Zulmara is Out for Onyxia's Lair, "+w.day()+".", false)
	})
	t.Run("a character not the account's is refused", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		other := seedUser(t, w.st.Pool)
		w.addChar(t, w.gid, other, "us/normal/vexlash", "Vexlash", true)
		wantReply(t, w.signupCmd(t, opt(optStatus, "in"), opt(optCharacter, "us/normal/vexlash")), pressCopy.Failure, false)
		if w.signups(t) != 0 {
			t.Fatal("wrote a sign-up")
		}
	})
	t.Run("unlinked", func(t *testing.T) {
		w := newPressWorld(t)
		rep, err := w.svc.cmdSignup(context.Background(), interactionOf(typeCommand, "d-stranger",
			map[string]any{"name": "signup", "options": []any{opt(optStatus, "in")}}))
		if err != nil {
			t.Fatal(err)
		}
		wantReply(t, rep, pressCopy.NotLinked, false, wantButton{label: linkCopy.LinkButton, url: w.svc.linkStartURL()})
	})
	t.Run("an unknown status or raid fails", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		wantReply(t, w.signupCmd(t, opt(optStatus, "maybe")), pressCopy.Failure, false)
		wantReply(t, w.signupCmd(t, opt(optStatus, "in"), opt(optRaid, "Onyxia")), pressCopy.Failure, false)
	})
}

func (w *pressWorld) suggest(t *testing.T, discordUser, focused, typed string) []Choice {
	t.Helper()
	in := interactionOf(typeAutocomplete, discordUser, map[string]any{"name": "signup", "options": []any{
		opt(optStatus, "in"), map[string]any{"name": focused, "value": typed, "focused": true},
	}})
	choices, err := w.svc.autocompleteSignup(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	return choices
}

func TestSignupAutocomplete(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", false)
	other := seedUser(t, w.st.Pool)
	w.addChar(t, w.gid, other, "us/normal/vexlash", "Vexlash", true)
	w.exec(t, `update guilds set raid_timezone = 'America/Chicago' where id = $1`, w.gid)
	w.exec(t, `delete from guild_raids`)
	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		t.Fatal(err)
	}
	base := time.Now().Add(time.Hour).Truncate(time.Minute)
	ids := make([]int64, 0, 27)
	for i := 26; i >= 0; i-- { // inserted latest first: the answer sorts
		ids = append(ids, w.addRaid(t, w.gid, base.Add(time.Duration(i)*24*time.Hour)))
	}
	w.addRaid(t, w.gid, base.Add(-10*time.Hour)) // ended
	cancelled := w.addRaid(t, w.gid, base.Add(2*time.Hour))
	w.exec(t, `update guild_raids set status = 'cancelled' where id = $1`, cancelled)

	raids := w.suggest(t, pressDiscord, optRaid, "")
	if len(raids) != maxChoices {
		t.Fatalf("raid choices = %d", len(raids))
	}
	soonest := ids[len(ids)-1]
	want := "Onyxia's Lair · " + base.In(chicago).Format("Mon 2 Jan")
	if raids[0].Name != want || raids[0].Value != fmt.Sprint(soonest) {
		t.Fatalf("first raid = %+v, want %q %d", raids[0], want, soonest)
	}
	if raids[1].Value != fmt.Sprint(ids[len(ids)-2]) {
		t.Fatalf("second raid = %+v", raids[1])
	}

	chars := w.suggest(t, pressDiscord, optCharacter, "")
	if len(chars) != 1 || chars[0].Name != "Zulmara" || chars[0].Value != keyZulmara {
		t.Fatalf("character choices = %+v", chars)
	}
	if got := w.suggest(t, pressDiscord, optCharacter, "dusk"); len(got) != 0 {
		t.Fatalf("pending character offered: %+v", got)
	}
	if got := w.suggest(t, "d-stranger", optRaid, ""); len(got) != 0 {
		t.Fatalf("unlinked got %+v", got)
	}
	w.exec(t, `update guild_characters set verified_at = null where user_id = $1`, w.uid)
	w.forgetCandidates()
	if got := w.suggest(t, pressDiscord, optRaid, ""); len(got) != 0 {
		t.Fatalf("non-member got the calendar: %+v", got)
	}
}
