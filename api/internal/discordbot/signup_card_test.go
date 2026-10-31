package discordbot

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

// bareSignup is /signup with no status, as the world's linked account.
func (w *pressWorld) bareSignup(t *testing.T, options ...map[string]any) Reply {
	t.Helper()
	rep := w.signupCmd(t, options...)
	assertCardLimits(t, rep.Message)
	return rep
}

// signupCardPress presses a /signup card component as the world's account.
func (w *pressWorld) signupCardPress(t *testing.T, id string, values ...string) Reply {
	t.Helper()
	data := map[string]any{"custom_id": id}
	if len(values) > 0 {
		data["values"] = values
	}
	h, ok := w.svc.components[customPrefix(id)]
	if !ok {
		t.Fatalf("no route for %q", id)
	}
	rep, err := h(context.Background(), interactionOf(typeComponent, pressDiscord, data))
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func buttonsOf(m Message) []Component {
	return m.Components[len(m.Components)-1]
}

func TestBareSignupCard(t *testing.T) {
	t.Run("one raid, one character", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		rep := w.bareSignup(t)
		if len(rep.Message.Components) != 1 || len(rep.Message.Embeds) != 1 || rep.Update {
			t.Fatalf("card = %+v", rep)
		}
		row := rep.Message.Components[0]
		want := []string{"In", "Late", "Bench", "Out", "Raid page"}
		if len(row) != len(want) {
			t.Fatalf("row = %+v", row)
		}
		for i, label := range want {
			if row[i].Label != label {
				t.Fatalf("button %d = %q, want %q", i, row[i].Label, label)
			}
		}
		for i, b := range row[:4] {
			if b.Style != StyleSecondary {
				t.Fatalf("answer %d style = %d, want Secondary: no answer reads as chosen", i, b.Style)
			}
		}
		if row[4].Style != StyleLink {
			t.Fatalf("styles = %+v", row)
		}
		if got, want := row[0].CustomID, fmt.Sprintf("gsgn:1:in:%d:%s", w.raidID, keyZulmara); got != want {
			t.Fatalf("id = %q, want %q", got, want)
		}
	})
	t.Run("two raids and two characters", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
		w.setMain(t, keyDusk)
		later := w.addRaid(t, w.gid, w.raidStart.Add(24*time.Hour))
		rep := w.bareSignup(t)
		rows := rep.Message.Components
		if len(rows) != 3 {
			t.Fatalf("rows = %+v", rows)
		}
		raids, chars := rows[0][0], rows[1][0]
		if raids.Placeholder != "Choose a raid" || len(raids.Options) != 2 || chars.Placeholder != "Choose a character" || len(chars.Options) != 2 {
			t.Fatalf("selects = %+v %+v", raids, chars)
		}
		if got, want := raids.Options[0].Label, "Onyxia's Lair · "+w.day(); got != want {
			t.Fatalf("raid label = %q, want %q", got, want)
		}
		if raids.Options[0].Value != fmt.Sprint(w.raidID) || !raids.Options[0].Default || raids.Options[1].Default || raids.Options[1].Value != fmt.Sprint(later) {
			t.Fatalf("raid options = %+v", raids.Options)
		}
		if defaultOf(chars) != keyDusk {
			t.Fatalf("default character = %q, want the main", defaultOf(chars))
		}
	})
	t.Run("typed raid and character are the defaults", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
		later := w.addRaid(t, w.gid, w.raidStart.Add(24*time.Hour))
		rep := w.bareSignup(t, opt(optRaid, fmt.Sprint(later)), opt(optCharacter, keyDusk))
		if defaultOf(rep.Message.Components[0][0]) != fmt.Sprint(later) || defaultOf(rep.Message.Components[1][0]) != keyDusk {
			t.Fatalf("card = %+v", rep.Message.Components)
		}
	})
	t.Run("no raid", func(t *testing.T) {
		w := newPressWorld(t)
		w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
		w.exec(t, `delete from guild_raids`)
		wantReply(t, w.signupCmd(t), pressCopy.NoRaid, false)
	})
}

func TestSignupCardPressSignsUp(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	card := w.bareSignup(t)
	rep := w.signupCardPress(t, buttonsOf(card.Message)[0].CustomID)
	if !rep.Update || !strings.HasPrefix(rep.Message.Content, "Zulmara is In as DPS for Onyxia's Lair, "+w.day()+".") {
		t.Fatalf("reply = %+v", rep)
	}
	if w.signups(t) != 1 || w.signedRaid(t, keyZulmara) != w.raidID {
		t.Fatal("no sign-up row")
	}
}

func TestSignupCardPicksRedraw(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
	later := w.addRaid(t, w.gid, w.raidStart.Add(24*time.Hour))
	id := fmt.Sprintf("gsgn:1:raid:%d:%s", w.raidID, keyZulmara)
	rep := w.signupCardPress(t, id, fmt.Sprint(later))
	assertCardLimits(t, rep.Message)
	if !rep.Update || defaultOf(rep.Message.Components[0][0]) != fmt.Sprint(later) || defaultOf(rep.Message.Components[1][0]) != keyZulmara {
		t.Fatalf("raid pick = %+v", rep)
	}
	rep = w.signupCardPress(t, fmt.Sprintf("gsgn:1:char:%d:%s", later, keyZulmara), keyDusk)
	if defaultOf(rep.Message.Components[1][0]) != keyDusk || defaultOf(rep.Message.Components[0][0]) != fmt.Sprint(later) {
		t.Fatalf("char pick = %+v", rep)
	}
	if w.signups(t) != 0 {
		t.Fatal("a pick wrote a sign-up")
	}
}

func TestSignupCardRefusesAForgedCharacter(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	other := seedUser(t, w.st.Pool)
	w.addChar(t, w.gid, other, "us/normal/someone-else", "Someone", true)
	for _, id := range []string{
		fmt.Sprintf("gsgn:1:in:%d:us/normal/someone-else", w.raidID),
		fmt.Sprintf("gsgn:1:in:%d:", w.raidID),
		"gsgn:1:in:notaraid:" + keyZulmara,
		"gsgn:1:in",
		fmt.Sprintf("gsgn:1:in:%d:%s:extra", w.raidID, keyZulmara),
	} {
		rep := w.signupCardPress(t, id)
		if rep.Message.Content != pressCopy.Failure || !rep.Update {
			t.Fatalf("%q = %+v", id, rep)
		}
	}
	for _, id := range []string{
		fmt.Sprintf("gsgn:1:raid:%d:us/normal/someone-else", w.raidID),
		fmt.Sprintf("gsgn:1:char:%d:%s", w.raidID, keyZulmara),
	} {
		value := "us/normal/someone-else"
		if strings.Contains(id, ":raid:") {
			value = "99999"
		}
		rep := w.signupCardPress(t, id, value)
		if rep.Message.Content != pressCopy.Failure || !rep.Update {
			t.Fatalf("%q = %+v", id, rep)
		}
	}
	if w.signups(t) != 0 {
		t.Fatal("a forged id wrote a sign-up")
	}
}

func TestSignupCardRaidLockedSinceDrawn(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	card := w.bareSignup(t)
	w.exec(t, `update guild_raids set status = 'locked' where id = $1`, w.raidID)
	rep := w.signupCardPress(t, buttonsOf(card.Message)[0].CustomID)
	if rep.Message.Content != pressCopy.Locked || !rep.Update {
		t.Fatalf("reply = %+v", rep)
	}
}

func TestSignupCardSurvivesHostileTitles(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
	hostile := "@everyone **x** " + strings.Repeat("é", 80)
	w.exec(t, `update guild_raids set title = $2 where id = $1`, w.raidID, string([]rune(hostile)[:80]))
	w.addRaid(t, w.gid, w.raidStart.Add(24*time.Hour))
	w.exec(t, `update guild_raids set title = $1 where title = 'Onyxia''s Lair'`, string([]rune(hostile)[:80]))
	rep := w.bareSignup(t)
	for _, o := range rep.Message.Components[0][0].Options {
		if n := len([]rune(o.Label)); n > maxOptionRunes {
			t.Fatalf("label is %d runes", n)
		}
	}
	redraw := w.signupCardPress(t, fmt.Sprintf("gsgn:1:char:%d:%s", w.raidID, keyZulmara), keyDusk)
	assertCardLimits(t, redraw.Message)
}

func TestSignupStatusIsOptional(t *testing.T) {
	for _, c := range Commands() {
		if c.Name != "signup" {
			continue
		}
		o := c.Options[0]
		if o.Name != "status" || o.Required || o.Description != "In, Late, Bench or Out. Leave out to pick on a card." {
			t.Fatalf("status = %+v", o)
		}
		return
	}
	t.Fatal("no /signup")
}
