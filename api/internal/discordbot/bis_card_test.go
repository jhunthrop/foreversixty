package discordbot

import (
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func (w *lookupWorld) bisCardOf(t *testing.T, user string, opts map[string]string) Message {
	t.Helper()
	m := w.command(t, user, "bis", opts).Message
	assertCardLimits(t, m)
	if !m.Ephemeral || !hasCard(m) || len(m.Components) != 4 {
		t.Fatalf("not a /bis card: %+v", m)
	}
	return m
}

func defaultOf(c Component) string {
	for _, o := range c.Options {
		if o.Default {
			return o.Value
		}
	}
	return ""
}

func buttonLabels(row ActionRow) []string {
	var out []string
	for _, c := range row {
		out = append(out, c.Label)
	}
	return out
}

// linkMain makes the pressing member's main a level 60 Beast Mastery Hunter named Zulmara.
func linkMain(w *lookupWorld) {
	w.port.candidates = []guilds.SignupCandidate{{CharacterKey: keyZulmaraLookup, Name: "Zulmara",
		Class: "hunter", Spec: "Beast Mastery", Verified: true, Main: true}}
	card := zulmaraCard()
	card.Character.Level = ref(60)
	w.cards.cards["zulmara"] = card
}

func TestBareBisUnlinked(t *testing.T) {
	w := newBisWorld(t)
	m := w.bisCardOf(t, "d-stranger", nil)
	e := headOf(t, m)
	if e.Author != "Best in slot" || e.Title != "Choose a class" || e.Description != guideCopy.BisLeadUnlinked ||
		e.Color != goldColor || e.Thumbnail != "" {
		t.Fatalf("embed = %+v", e)
	}
	spec := m.Components[1][0]
	if !spec.Disabled || len(spec.Options) != 1 || spec.Options[0].Value != "none" || spec.Options[0].Label != "Choose a class first" {
		t.Fatalf("spec select = %+v", spec)
	}
	rng := m.Components[2][0]
	if defaultOf(rng) != "p1" || rng.Options[5].Label != "Phase 1 raid" || rng.Options[5].Description != "Level 60, raid buffs and consumables" {
		t.Fatalf("range select = %+v", rng)
	}
	class := m.Components[0][0]
	if class.Options[0].Label != "Druid" || class.Options[0].Description != "Balance, Feral, Feral Bear, Restoration" || defaultOf(class) != "" {
		t.Fatalf("class select = %+v", class.Options[0])
	}
	row := m.Components[3]
	if len(row) != 2 || row[0].Label != "Show BiS" || row[0].Style != StyleSuccess || row[0].CustomID != "gbis:1:go:::p1:" ||
		row[1].Label != linkCopy.LinkButton || row[1].URL != w.svc.linkStartURL() {
		t.Fatalf("buttons = %+v", row)
	}
}

func TestBareBisLinkedWithAMain(t *testing.T) {
	w := newBisWorld(t)
	linkMain(w)
	m := w.bisCardOf(t, pressDiscord, nil)
	e := headOf(t, m)
	if e.Description != "Picked from your main, **Zulmara**. Show BiS, or change a pick." ||
		e.Title != "Beast Mastery Hunter · level 60 Phase 1 raid" || e.Color != classColorOf("hunter") ||
		!strings.HasSuffix(e.Thumbnail, "/icons/hd/crests/hunter.webp") {
		t.Fatalf("embed = %+v", e)
	}
	got := []string{defaultOf(m.Components[0][0]), defaultOf(m.Components[1][0]), defaultOf(m.Components[2][0])}
	if !slices.Equal(got, []string{"hunter", "beast-mastery", "p1"}) {
		t.Fatalf("defaults = %v", got)
	}
	if got := buttonLabels(m.Components[3]); !slices.Equal(got, []string{"Show BiS"}) {
		t.Fatalf("buttons = %v", got)
	}
}

func TestBisCardOffersTheMainWhenAPickDiffers(t *testing.T) {
	w := newBisWorld(t)
	linkMain(w)
	m := w.bisCardOf(t, pressDiscord, map[string]string{"phase": "40"})
	if got := buttonLabels(m.Components[3]); !slices.Equal(got, []string{"Show BiS", "Use Zulmara"}) {
		t.Fatalf("buttons = %v", got)
	}
	if m.Components[3][1].Style != StyleSecondary || m.Components[3][1].CustomID != "gbis:1:me:hunter:beast-mastery:40:" {
		t.Fatalf("use button = %+v", m.Components[3][1])
	}
}

func TestBareBisLinkedWithoutACharacter(t *testing.T) {
	w := newBisWorld(t)
	w.port.candidates = []guilds.SignupCandidate{}
	m := w.bisCardOf(t, pressDiscord, nil)
	if headOf(t, m).Description != guideCopy.BisLeadNone || len(m.Components[3]) != 1 {
		t.Fatalf("card = %+v", m)
	}
}

func TestBisPartials(t *testing.T) {
	w := newBisWorld(t)
	t.Run("class only", func(t *testing.T) {
		m := w.bisCardOf(t, "d-stranger", map[string]string{"class": "warrior"})
		if defaultOf(m.Components[0][0]) != "warrior" || defaultOf(m.Components[1][0]) != "" ||
			headOf(t, m).Description != "Now choose a Warrior spec." || headOf(t, m).Title != "Warrior · choose a spec" {
			t.Fatalf("card = %+v", headOf(t, m))
		}
		if got := m.Components[1][0].Options; len(got) != 3 || got[0].Description != "DPS" {
			t.Fatalf("spec options = %+v", got)
		}
	})
	t.Run("phase only", func(t *testing.T) {
		m := w.bisCardOf(t, "d-stranger", map[string]string{"phase": "40"})
		if defaultOf(m.Components[2][0]) != "40" {
			t.Fatalf("range = %+v", m.Components[2][0])
		}
	})
	t.Run("slot only", func(t *testing.T) {
		m := w.bisCardOf(t, "d-stranger", map[string]string{"slot": "head"})
		if headOf(t, m).Description != "For Head only. Choose a class, a spec and a range, then Show BiS." {
			t.Fatalf("lead = %q", headOf(t, m).Description)
		}
		for _, row := range m.Components {
			if id := row[0].CustomID; !strings.HasSuffix(id, ":head") {
				t.Fatalf("id %q lost the slot", id)
			}
		}
	})
	t.Run("unknown spec", func(t *testing.T) {
		m := w.bisCardOf(t, "d-stranger", map[string]string{"class": "warrior", "spec": "nonsense"})
		if headOf(t, m).Description != "No Warrior spec by that name. Choose one." || defaultOf(m.Components[0][0]) != "warrior" {
			t.Fatalf("card = %+v", headOf(t, m))
		}
		m = w.bisCardOf(t, "d-stranger", map[string]string{"spec": "nonsense"})
		if headOf(t, m).Description != "No spec by that name. Choose one." {
			t.Fatalf("lead = %q", headOf(t, m).Description)
		}
	})
	t.Run("a resolvable spec answers at once", func(t *testing.T) {
		m := w.command(t, "d-stranger", "bis", map[string]string{"spec": "warrior-fury", "phase": "p1"}).Message
		if head := textsOf(m)[0]; !strings.HasPrefix(head, "### Fury Warrior\nBest in slot · level 60 Phase 1 raid") {
			t.Fatalf("head = %q", head)
		}
	})
}

func TestBisPicksIgnoreForgedFields(t *testing.T) {
	got := bisPicksOf([]string{"paladin", "fury", "99", "boots", "x"})
	if got != (bisPicks{Class: "paladin", Range: "p1"}) {
		t.Fatalf("picks = %+v", got)
	}
	if got := bisPicksOf(nil); got != (bisPicks{Range: "p1"}) {
		t.Fatalf("nil picks = %+v", got)
	}
	got = bisPicksOf([]string{"warrior", "fury", "40", "head"})
	if got != (bisPicks{"warrior", "fury", "40", "head"}) {
		t.Fatalf("valid picks = %+v", got)
	}
	if key, ok := got.specKey(); !ok || key != "warrior-fury" {
		t.Fatalf("specKey = %q %v", key, ok)
	}
}

func TestBisCommandDescriptions(t *testing.T) {
	for _, c := range Commands() {
		if c.Name != "bis" {
			continue
		}
		for _, o := range c.Options {
			want := map[string]string{"class": "The class. Leave out to pick it.", "spec": "The spec. Leave out to pick it."}[o.Name]
			if want != "" && o.Description != want {
				t.Fatalf("%s = %q", o.Name, o.Description)
			}
		}
	}
}

func TestBisCardForAMainOfAnotherClassKeepsTheTypedClass(t *testing.T) {
	w := newBisWorld(t)
	linkMain(w)
	m := w.bisCardOf(t, pressDiscord, map[string]string{"class": "warrior"})
	if defaultOf(m.Components[0][0]) != "warrior" || defaultOf(m.Components[1][0]) != "" {
		t.Fatalf("card = %+v", m.Components)
	}
}
