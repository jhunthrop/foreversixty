package discordbot

import (
	"net/url"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

// furyBands is a Fury Warrior file: a bare and a raid band at 60 for the Alliance, the raid
// band's slots each with a first alternative so the five-slot view has an order to find.
func furyBands() []map[string]any {
	slot := func(name string, id int, delta float64) map[string]any {
		return map[string]any{"slot": name, "item_id": id, "item_name": "Pick " + name, "source": "Source " + name,
			"alternatives": []map[string]any{
				{"item_id": id + 50, "item_name": "Alt one " + name, "source": "PvP rank 14 · Lieutenant Commander · Alliance", "dps_delta": delta},
				{"item_id": id + 60, "item_name": "Alt two " + name, "source": "Blacksmithing", "dps_delta": delta - 3},
				{"item_id": id + 70, "item_name": "Alt three " + name, "dps_delta": delta - 6},
			}}
	}
	raid := []map[string]any{
		slot("head", 100, -14.96), slot("neck", 101, -2), slot("finger1", 102, -20), slot("main_hand", 103, -30),
		slot("feet", 104, -1), slot("legs", 105, -9), slot("back", 106, -0.5),
		{"slot": "ranged", "item_id": 0, "item_name": ""},
	}
	return []map[string]any{
		{"band": 60, "faction": "alliance", "preset": "bare", "slots": []map[string]any{slot("head", 107, -1)}},
		{"band": 60, "faction": "alliance", "preset": "raid", "slots": raid},
	}
}

func newBisWorld(t *testing.T) *lookupWorld {
	t.Helper()
	w := newLookupWorld(t)
	writeBis(t, w.lg.dir, "warrior-fury", furyBands())
	return w
}

func TestBisWithAndWithoutSlot(t *testing.T) {
	page := pressPublic + "/bis/warrior/fury?faction=alliance&preset=raid#band-alliance-60"
	t.Run("one slot", func(t *testing.T) {
		w := newBisWorld(t)
		rep := w.command(t, pressDiscord, "bis", map[string]string{"class": "warrior", "spec": "fury", "slot": "head", "phase": "p1"})
		m := rep.Message
		assertMessageLimits(t, m)
		if m.Content != "Head for a Fury Warrior, level 60 Phase 1 raid:" || !m.Ephemeral || len(m.Embeds) != 3 {
			t.Fatalf("reply = %+v", m)
		}
		if m.Embeds[0].Author != "Best in slot" || m.Embeds[0].Description != "Head · Epic · item level 61\nSource head" ||
			!sameBandPage(m.Embeds[0].URL, page) || m.Embeds[0].Color != 10696174 {
			t.Fatalf("pick = %+v", m.Embeds[0])
		}
		if m.Embeds[1].Author != "" || m.Embeds[1].Description != "Head · Epic · item level 61\n−15.0 DPS · PvP rank 14 · Lieutenant Commander · Alliance" ||
			m.Embeds[2].Description != "Head\n−18.0 DPS · Blacksmithing" {
			t.Fatalf("alternatives = %+v", m.Embeds[1:])
		}
		assertDistinctEmbedURLs(t, m.Embeds)
		if m.Embeds[2].Footer != "Fury Warrior · level 60 Phase 1 raid · Forever Sixty" {
			t.Fatalf("footer = %q", m.Embeds[2].Footer)
		}
		row := m.Components[0]
		if len(row) != 3 || row[0].Label != lookupCopy.Share || row[1].Label != guideCopy.BisChange || row[2].Label != lookupCopy.BisOpen || row[2].URL != page {
			t.Fatalf("row = %+v", row)
		}
	})
	t.Run("five slots by the first alternative's delta", func(t *testing.T) {
		w := newBisWorld(t)
		m := w.command(t, pressDiscord, "bis", map[string]string{"spec": "warrior-fury", "phase": "p1"}).Message
		assertMessageLimits(t, m)
		if m.Content != "Fury Warrior, level 60 Phase 1 raid. The five slots where the best item gains the most over the next best:" {
			t.Fatalf("lead = %q", m.Content)
		}
		var titles []string
		for _, e := range m.Embeds {
			titles = append(titles, e.Title)
		}
		if got := strings.Join(titles, ","); got != "Pick main\\_hand,Pick finger1,Pick head,Pick legs,Pick neck" &&
			got != "Pick main_hand,Pick finger1,Pick head,Pick legs,Pick neck" {
			t.Fatalf("order = %s", got)
		}
		assertDistinctEmbedURLs(t, m.Embeds)
		if !strings.HasPrefix(m.Embeds[0].Description, "Main hand · Epic · item level 61\n+30.0 DPS · ") {
			t.Fatalf("first = %q", m.Embeds[0].Description)
		}
		if m.Components[0][2].Label != lookupCopy.BisAllSlots {
			t.Fatalf("row = %+v", m.Components[0])
		}
	})
	t.Run("defaults from the linked main", func(t *testing.T) {
		w := newBisWorld(t)
		w.port.candidates = []guilds.SignupCandidate{{CharacterKey: "us/normal/brakka", Name: "Brakka", Class: "warrior", Spec: "Fury", Verified: true, Main: true}}
		w.cards.cards["brakka"] = rankings.Character{Character: rankings.CharacterRef{Key: "us/normal/brakka", Region: "us",
			Ruleset: "normal", Name: "Brakka", Class: "warrior", Level: ref(60)}}
		m := w.command(t, pressDiscord, "bis", map[string]string{"spec": "fury", "slot": "head"}).Message
		if m.Content != "Head for a Fury Warrior, level 60 Phase 1 raid:" {
			t.Fatalf("lead = %q", m.Content)
		}
		bare := w.command(t, pressDiscord, "bis", map[string]string{"spec": "fury", "slot": "head", "phase": "60"}).Message
		if bare.Content != "Head for a Fury Warrior, level 60:" || bare.Embeds[0].Title != "Pick head" ||
			!strings.Contains(bare.Embeds[0].URL, "preset=bare") {
			t.Fatalf("bare = %+v", bare)
		}
	})
}

// assertDistinctEmbedURLs fails when two embeds share a URL: Discord folds consecutive embeds
// with one URL into a single card, and five picks once showed as one item.
func assertDistinctEmbedURLs(t *testing.T, embeds []notify.Embed) {
	t.Helper()
	seen := map[string]bool{}
	for _, e := range embeds {
		if e.URL == "" || seen[e.URL] {
			t.Fatalf("embed URL %q is empty or repeated among %d embeds", e.URL, len(embeds))
		}
		seen[e.URL] = true
	}
}

// sameBandPage reports whether link opens page: the same path, fragment and query apart from
// the per-item key.
func sameBandPage(link, page string) bool {
	l, err1 := url.Parse(link)
	p, err2 := url.Parse(page)
	if err1 != nil || err2 != nil || l.Path != p.Path || l.Fragment != p.Fragment || l.Query().Get(itemLinkParam) == "" {
		return false
	}
	q := l.Query()
	q.Del(itemLinkParam)
	return q.Encode() == p.Query().Encode()
}

func TestBisSpecAutocompleteFiltersByClass(t *testing.T) {
	w := newBisWorld(t)
	in := interactionOf(typeAutocomplete, pressDiscord, map[string]any{"name": "bis", "options": []map[string]any{
		{"name": "class", "value": "warrior"}, {"name": "spec", "value": "", "focused": true}}})
	got, err := w.svc.autocompleteSpec(t.Context(), in)
	if err != nil || len(got) != 3 || got[1].Name != "Fury Warrior" || got[1].Value != "warrior-fury" {
		t.Fatalf("choices = %+v, %v", got, err)
	}
}
