package discordbot

import (
	"context"
	"encoding/json"
	"slices"
	"strings"
	"testing"
)

// pressBis presses a /bis card component and checks the answer is an in-place edit.
func (w *lookupWorld) pressBis(t *testing.T, user, id string, values ...string) Reply {
	t.Helper()
	rep := w.card(t, user, id, values...)
	if !rep.Update {
		t.Fatalf("press %q did not edit in place: %+v", id, rep)
	}
	assertCardLimits(t, rep.Message)
	return rep
}

func leadOf(m Message) string { return m.Embeds[0].Description }

func jsonOf(t *testing.T, v any) string {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

func TestBisCardPicks(t *testing.T) {
	w := newBisWorld(t)
	t.Run("class then spec", func(t *testing.T) {
		m := w.pressBis(t, "d-stranger", "gbis:1:c:::p1:", "hunter").Message
		if m.Embeds[0].Title != "Hunter · choose a spec" || leadOf(m) != "Now choose a Hunter spec." ||
			defaultOf(m.Components[0][0]) != "hunter" || defaultOf(m.Components[2][0]) != "p1" || m.Components[1][0].Disabled {
			t.Fatalf("card = %+v", m.Embeds[0])
		}
		if opts := m.Components[1][0].Options; len(opts) != 3 || opts[0].Description != "DPS" {
			t.Fatalf("spec options = %+v", opts)
		}
		m = w.pressBis(t, "d-stranger", "gbis:1:s:hunter::p1:", "beast-mastery").Message
		if leadOf(m) != "Ready. Show BiS, or change a pick." || defaultOf(m.Components[1][0]) != "beast-mastery" {
			t.Fatalf("lead = %q", leadOf(m))
		}
	})
	t.Run("the nudge keeps every pick", func(t *testing.T) {
		m := w.pressBis(t, "d-stranger", "gbis:1:go:hunter::p1:").Message
		if leadOf(m) != "**Choose a spec first.** Then Show BiS." || defaultOf(m.Components[0][0]) != "hunter" || defaultOf(m.Components[2][0]) != "p1" {
			t.Fatalf("card = %+v", m.Embeds[0])
		}
		m = w.pressBis(t, "d-stranger", "gbis:1:go:::p1:").Message
		if leadOf(m) != "**Choose a class and a spec first.** Then Show BiS." {
			t.Fatalf("lead = %q", leadOf(m))
		}
	})
	t.Run("a range pick lets the nudge give way", func(t *testing.T) {
		m := w.pressBis(t, "d-stranger", "gbis:1:r:hunter::p1:", "50").Message
		if leadOf(m) != "Now choose a Hunter spec." || defaultOf(m.Components[2][0]) != "50" {
			t.Fatalf("card = %+v", m.Embeds[0])
		}
	})
	t.Run("a forged pick reads as empty", func(t *testing.T) {
		m := w.pressBis(t, "d-stranger", "gbis:1:c:::p1:", "paladin-nonsense").Message
		if defaultOf(m.Components[0][0]) != "" || leadOf(m) != guideCopy.BisLeadUnlinked {
			t.Fatalf("card = %+v", m.Embeds[0])
		}
	})
	t.Run("the main's spec comes back with the main's class", func(t *testing.T) {
		w := newBisWorld(t)
		linkMain(w)
		m := w.pressBis(t, pressDiscord, "gbis:1:c:hunter:beast-mastery:p1:", "warrior").Message
		if defaultOf(m.Components[1][0]) != "" || leadOf(m) != "Now choose a Warrior spec." {
			t.Fatalf("warrior = %+v", m.Embeds[0])
		}
		m = w.pressBis(t, pressDiscord, "gbis:1:c:warrior::p1:", "hunter").Message
		if defaultOf(m.Components[1][0]) != "beast-mastery" || leadOf(m) != "Picked from your main, **Zulmara**. Show BiS, or change a pick." {
			t.Fatalf("hunter = %+v", m.Embeds[0])
		}
	})
}

func TestShowBisEqualsTheTypedAnswer(t *testing.T) {
	w := newBisWorld(t)
	typed := w.command(t, pressDiscord, "bis", map[string]string{"spec": "warrior-fury", "phase": "p1"}).Message
	card := w.pressBis(t, pressDiscord, "gbis:1:go:warrior:fury:p1:")
	if card.Message.Content != typed.Content || jsonOf(t, card.Message.Embeds) != jsonOf(t, typed.Embeds) {
		t.Fatalf("card answer differs:\n%q\n%q", card.Message.Content, typed.Content)
	}
	if !card.Message.Ephemeral {
		t.Fatal("the answer is not ephemeral")
	}
	want := []string{lookupCopy.Share, "Change picks", lookupCopy.BisAllSlots}
	if got := buttonLabels(card.Message.Components[0]); !slices.Equal(got, want) {
		t.Fatalf("card row = %v", got)
	}
	if card.Message.Components[0][1].CustomID != "gbis:1:chg:warrior:fury:p1:" {
		t.Fatalf("change id = %q", card.Message.Components[0][1].CustomID)
	}
	if got := buttonLabels(typed.Components[0]); !slices.Equal(got, want) {
		t.Fatalf("typed row = %v", got)
	}
	for i, c := range typed.Components[0] {
		if i != 0 && jsonOf(t, c) != jsonOf(t, card.Message.Components[0][i]) {
			t.Fatalf("button %d differs", i)
		}
	}
	slotted := w.pressBis(t, pressDiscord, "gbis:1:go:warrior:fury:p1:head").Message
	if slotted.Content != "Head for a Fury Warrior, level 60 Phase 1 raid:" || len(slotted.Embeds) != 3 {
		t.Fatalf("slot answer = %+v", slotted)
	}
}

func TestUseMainReadsTheLinkNotTheID(t *testing.T) {
	w := newBisWorld(t)
	writeBis(t, w.lg.dir, "hunter-beast-mastery", furyBands())
	linkMain(w)
	rep := w.pressBis(t, pressDiscord, "gbis:1:me:warrior:fury:50:")
	if !strings.HasPrefix(rep.Message.Content, "Beast Mastery Hunter, level 60 Phase 1 raid.") {
		t.Fatalf("content = %q", rep.Message.Content)
	}
	stranger := w.pressBis(t, "d-stranger", "gbis:1:me:warrior:fury:50:")
	if stranger.Message.Content != commonCopy.Failure {
		t.Fatalf("stranger = %+v", stranger.Message)
	}
}

func TestChangePicksReopensTheCard(t *testing.T) {
	w := newBisWorld(t)
	m := w.pressBis(t, "d-stranger", "gbis:1:chg:warrior:fury:p1:head").Message
	if len(m.Components) != 4 || leadOf(m) != guideCopy.BisLeadReady ||
		defaultOf(m.Components[0][0]) != "warrior" || defaultOf(m.Components[1][0]) != "fury" || defaultOf(m.Components[2][0]) != "p1" {
		t.Fatalf("card = %+v", m.Embeds[0])
	}
	if !strings.HasSuffix(m.Components[3][0].CustomID, ":head") {
		t.Fatalf("slot lost: %q", m.Components[3][0].CustomID)
	}
}

func TestUseMainOnlyWhenAPickDiffers(t *testing.T) {
	w := newBisWorld(t)
	linkMain(w)
	m := w.pressBis(t, pressDiscord, "gbis:1:chg:hunter:beast-mastery:p1:").Message
	if got := buttonLabels(m.Components[3]); !slices.Equal(got, []string{"Show BiS"}) {
		t.Fatalf("at main picks = %v", got)
	}
	m = w.pressBis(t, pressDiscord, "gbis:1:r:hunter:beast-mastery:p1:", "50").Message
	if got := buttonLabels(m.Components[3]); !slices.Equal(got, []string{"Show BiS", "Use Zulmara"}) {
		t.Fatalf("range changed = %v", got)
	}
}

func TestShareFromACardKeepsTheControls(t *testing.T) {
	w := newBisWorld(t)
	answer := w.pressBis(t, pressDiscord, "gbis:1:go:warrior:fury:p1:").Message
	shareID := answer.Components[0][0].CustomID
	rep := w.pressShare(t, shareID)
	if !rep.Update || rep.After == nil {
		t.Fatalf("share = %+v", rep)
	}
	got := buttonLabels(rep.Message.Components[0])
	if !slices.Equal(got, []string{lookupCopy.Shared, "Change picks", lookupCopy.BisAllSlots}) {
		t.Fatalf("own copy = %v", got)
	}
	rep.After(context.Background())
	ups := w.followUps()
	if len(ups) != 1 {
		t.Fatalf("follow-ups = %d", len(ups))
	}
	if strings.Contains(jsonOf(t, ups[0].Message), "Change picks") || strings.Contains(jsonOf(t, ups[0].Message), "gbis:") {
		t.Fatal("the public follow-up carries the private controls")
	}
}
