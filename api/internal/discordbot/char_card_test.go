package discordbot

import (
	"context"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

const (
	keyDuskstrider = "us/normal/duskstrider"
	charAskID      = "gchar:1:ask:"
	charPickID     = "gchar:1:pick:"
	charNameID     = "gchar:1:name:"
)

// linkTwoChars links the viewer with Zulmara (main) and a second verified character.
func linkTwoChars(w *lookupWorld) {
	linkMain(w)
	w.port.candidates = append(w.port.candidates, guilds.SignupCandidate{CharacterKey: keyDuskstrider,
		Name: "Duskstrider", Class: "warrior", Spec: "Fury", Verified: true})
	dusk := zulmaraCard()
	dusk.Character.Key, dusk.Character.Name, dusk.Character.Class = keyDuskstrider, "Duskstrider", "warrior"
	dusk.GearSpec, dusk.Character.Level = ref("warrior-fury"), ref(58)
	w.cards.cards["duskstrider"] = dusk
}

func rowLabels(row ActionRow) []string {
	var out []string
	for _, c := range row {
		out = append(out, c.Label)
	}
	return out
}

func TestBareCharLinkedAnswersTheMain(t *testing.T) {
	w := newLookupWorld(t)
	linkTwoChars(w)
	typed := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
	rep := w.command(t, pressDiscord, "char", nil)
	m := rep.Message
	assertCardLimits(t, m)
	if rep.Update || jsonOf(t, m.Layout) != jsonOf(t, typed.Layout) {
		t.Fatalf("bare /char differs from the typed answer: %+v", rep)
	}
	if got := strings.Join(rowLabels(m.Components[0]), "|"); got != "Share in channel|Someone else|Character page" {
		t.Fatalf("row 1 = %q", got)
	}
	if m.Components[0][1].CustomID != charAskID || m.Components[0][1].Style != StyleSecondary {
		t.Fatalf("Someone else = %+v", m.Components[0][1])
	}
	sel := m.Components[1][0]
	if sel.CustomID != charPickID || sel.Placeholder != "Another of your characters" || len(sel.Options) != 2 {
		t.Fatalf("select = %+v", sel)
	}
	if o := sel.Options[0]; o.Value != "Zulmara" || !o.Default || o.Description != "Beast Mastery Hunter · Level 60" {
		t.Fatalf("first option = %+v", o)
	}
	if sel.Options[1].Default || sel.Options[1].Description != "Fury Warrior · Level 58" {
		t.Fatalf("second option = %+v", sel.Options[1])
	}
}

func TestBareCharWithOneCharacterHasNoSelect(t *testing.T) {
	w := newLookupWorld(t)
	linkMain(w)
	m := w.command(t, pressDiscord, "char", nil).Message
	if len(m.Components) != 1 || m.Components[0][1].Label != "Someone else" {
		t.Fatalf("components = %+v", m.Components)
	}
}

func TestBareCharUnlinked(t *testing.T) {
	w := newLookupWorld(t)
	m := w.command(t, "d-stranger", "char", nil).Message
	assertCardLimits(t, m)
	if contentOf(m) != "Which character? Type a name, or link Discord to start from yours." || hasCard(m) || len(m.Components) != 1 {
		t.Fatalf("card = %+v", m)
	}
	row := m.Components[0]
	if len(row) != 2 || row[0].Label != "Type a name" || row[0].Style != StyleSuccess || row[0].CustomID != charAskID ||
		row[1].Label != "Link Discord" || row[1].Style != StyleLink {
		t.Fatalf("row = %+v", row)
	}
}

func TestBareCharLinkedWithNoCharacter(t *testing.T) {
	w := newLookupWorld(t)
	w.port.candidates = []guilds.SignupCandidate{}
	m := w.command(t, pressDiscord, "char", nil).Message
	if contentOf(m) != "Which character? Type a name." || len(m.Components[0]) != 1 {
		t.Fatalf("card = %+v", m)
	}
}

func TestTypeANameOpensTheModal(t *testing.T) {
	w := newLookupWorld(t)
	rep := w.card(t, pressDiscord, charAskID)
	m := rep.Modal
	if m == nil || m.CustomID != charNameID || m.Title != "Look up a character" || m.InputLabel != "Character name" ||
		m.Placeholder != "Zulmara" || m.MinLength != 2 || m.MaxLength != 12 {
		t.Fatalf("modal = %+v", m)
	}
}

func TestModalNameAnswersInPlace(t *testing.T) {
	w := newLookupWorld(t)
	linkTwoChars(w)
	rep := w.modal(t, pressDiscord, charNameID, "Zulmara")
	if !rep.Update || !hasCard(rep.Message) || len(rep.Message.Components) != 2 {
		t.Fatalf("answer = %+v", rep)
	}
	assertCardLimits(t, rep.Message)
	rep = w.modal(t, pressDiscord, charNameID, "Nobody")
	if !rep.Update || contentOf(rep.Message) != "No public character named Nobody. Try another name." || hasCard(rep.Message) {
		t.Fatalf("unknown = %+v", rep)
	}
	assertCardLimits(t, rep.Message)
}

func TestTypedUnknownCharOpensTheCard(t *testing.T) {
	w := newLookupWorld(t)
	linkTwoChars(w)
	rep := w.command(t, pressDiscord, "char", map[string]string{"name": "Nobody"})
	m := rep.Message
	assertCardLimits(t, m)
	if rep.Update || contentOf(m) != "No public character named Nobody. Try another name." ||
		m.Components[0][0].Label != "Type a name" || m.Components[0][0].CustomID != charAskID {
		t.Fatalf("card = %+v", rep)
	}
	if sel := m.Components[1][0]; sel.CustomID != charPickID || len(sel.Options) != 2 || defaultOf(sel) != "" {
		t.Fatalf("select = %+v", sel)
	}
}

func TestUnknownNameIsShownLiterally(t *testing.T) {
	w := newLookupWorld(t)
	rep := w.modal(t, "d-stranger", charNameID, "@everyone**x_")
	if strings.Contains(contentOf(rep.Message), "**x") {
		t.Fatalf("unescaped member text: %q", contentOf(rep.Message))
	}
	assertCardLimits(t, rep.Message)
}

func TestCharPickKeepsTheSelect(t *testing.T) {
	w := newLookupWorld(t)
	linkTwoChars(w)
	rep := w.card(t, pressDiscord, charPickID, "Duskstrider")
	if !rep.Update || headOf(t, rep.Message).Title != "Duskstrider" {
		t.Fatalf("pick = %+v", rep)
	}
	assertCardLimits(t, rep.Message)
	sel := rep.Message.Components[1][0]
	if len(sel.Options) != 2 || defaultOf(sel) != "Duskstrider" {
		t.Fatalf("select = %+v", sel)
	}
}

func TestGuidedCharShareKeepsTheControls(t *testing.T) {
	w := newLookupWorld(t)
	linkTwoChars(w)
	answer := w.command(t, pressDiscord, "char", nil).Message
	shareID := answer.Components[0][0].CustomID
	rep := w.pressShare(t, shareID)
	own := rep.Message
	if labels := strings.Join(rowLabels(own.Components[0]), "|"); labels != "Shared|Someone else|Character page" || len(own.Components) != 2 {
		t.Fatalf("shared edit = %+v", own.Components)
	}
	rep.After(context.Background())
	ups := w.followUps()
	if len(ups) != 1 {
		t.Fatalf("follow-ups = %d", len(ups))
	}
	public := ups[0].Message
	if len(public.Components) != 1 || strings.Contains(jsonOf(t, public.Components), "Someone else") ||
		strings.Contains(jsonOf(t, public.Components), "gchar") {
		t.Fatalf("public follow-up = %+v", public.Components)
	}
}

func TestCharNameIsOptional(t *testing.T) {
	for _, c := range Commands() {
		if c.Name != "char" {
			continue
		}
		o := c.Options[0]
		if o.Name != "name" || o.Required || o.Description != "A character name. Yours if left out." {
			t.Fatalf("name option = %+v", o)
		}
		return
	}
	t.Fatal("no /char command")
}
