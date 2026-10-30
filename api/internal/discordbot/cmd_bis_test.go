package discordbot

import (
	"encoding/json"
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
		m := w.command(t, pressDiscord, "bis", map[string]string{"class": "warrior", "spec": "fury", "slot": "head", "phase": "p1"}).Message
		assertCardLimits(t, m)
		want := []string{"### Head · Fury Warrior\nBest in slot · level 60 Phase 1 raid · Alliance",
			"**Pick head**\nHead · +15.0 DPS\n-# Epic · item level 61 · Source head",
			panelCopy.BisSlotNext,
			"**Alt one head**\nHead · −15.0 DPS vs the best pick\n-# Epic · item level 61 · PvP rank 14 · Lieutenant Commander · Alliance",
			// An item the build does not know: no quality word or item level.
			"**Alt two head**\nHead · −18.0 DPS vs the best pick\n-# Blacksmithing"}
		if got := textsOf(m); strings.Join(got, "|") != strings.Join(want, "|") {
			t.Fatalf("texts = %q", got)
		}
		if head := headOf(t, m); head.Color != 0xc69b6d || head.Thumbnail != pressPublic+"/icons/hd/crests/warrior.webp" {
			t.Fatalf("head = %+v", head)
		}
		if tile := sectionsOf(m)[1].Accessory; tile.Media != pressAPI+"/v1/discord/tiles/"+lookupBuild+"/100.png" || tile.Description != "Pick head, Epic head item" {
			t.Fatalf("tile = %+v", tile)
		}
		if tile := sectionsOf(m)[3].Accessory; tile.Description != "Alt two head, head item" {
			t.Fatalf("unknown tile = %+v", tile)
		}
		row := m.Components[0]
		if got := strings.Join(buttonLabels(row), "|"); got != "Share in channel|Change picks|Open on Forever Sixty" || row[2].URL != page {
			t.Fatalf("row = %+v", row)
		}
	})
	t.Run("page 1 in the character frame's order", func(t *testing.T) {
		w := newBisWorld(t)
		m := w.command(t, pressDiscord, "bis", map[string]string{"spec": "warrior-fury", "phase": "p1"}).Message
		assertCardLimits(t, m)
		want := []string{"### Fury Warrior\nBest in slot · level 60 Phase 1 raid · Alliance\n-# Head to Legs, 9 of 17 slots. Each gain is over the next best item.",
			"**Pick head**\nHead · +15.0 DPS\n-# Epic · item level 61 · Source head",
			"**Pick neck**\nNeck · +2.0 DPS\n-# Epic · item level 61 · Source neck",
			"**Shoulder** · no pick for this range.",
			"**Pick back**\nBack · +0.5 DPS\n-# Epic · item level 61 · Source back",
			"**Chest** · no pick for this range.", "**Wrist** · no pick for this range.",
			"**Hands** · no pick for this range.", "**Waist** · no pick for this range.",
			"**Pick legs**\nLegs · +9.0 DPS\n-# Epic · item level 61 · Source legs"}
		if got := textsOf(m); strings.Join(got, "|") != strings.Join(want, "|") {
			t.Fatalf("texts = %q", got)
		}
		row := m.Components[0]
		if got := strings.Join(buttonLabels(row), "|"); got != "Share in channel|Feet to Ranged|Change picks|All 17 slots" ||
			row[1].CustomID != "gbis:1:pg:warrior:fury:p1::2" || row[3].URL != page {
			t.Fatalf("row = %+v", row)
		}

		// The page button answers page 2 in place, with the message's own Share copied.
		in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": row[1].CustomID})
		in.Source = sourceOf(t, m)
		two := w.route(t, in)
		if !two.Update {
			t.Fatal("the page press did not edit in place")
		}
		assertCardLimits(t, two.Message)
		texts := textsOf(two.Message)
		if !strings.HasSuffix(texts[0], "-# Feet to Ranged, 8 of 17 slots. Each gain is over the next best item.") ||
			texts[1] != "**Pick feet**\nFeet · +1.0 DPS\n-# Epic · item level 61 · Source feet" ||
			texts[2] != "**Pick finger1**\nRing 1 · +20.0 DPS\n-# Epic · item level 61 · Source finger1" ||
			texts[len(texts)-1] != "**Ranged** · no pick for this range." || len(texts) != 9 {
			t.Fatalf("page 2 = %q", texts)
		}
		back := two.Message.Components[0]
		if back[0].CustomID != row[0].CustomID || back[1].Label != "Head to Legs" || back[1].CustomID != "gbis:1:pg:warrior:fury:p1::1" {
			t.Fatalf("page 2 row = %+v", back)
		}
		var shares int
		if err := w.st.Pool.QueryRow(t.Context(), `select count(*) from discord_shares`).Scan(&shares); err != nil || shares != 1 {
			t.Fatalf("shares = %d, %v: a page press minted a token", shares, err)
		}
	})
	t.Run("defaults from the linked main", func(t *testing.T) {
		w := newBisWorld(t)
		w.port.candidates = []guilds.SignupCandidate{{CharacterKey: "us/normal/brakka", Name: "Brakka", Class: "warrior", Spec: "Fury", Verified: true, Main: true}}
		w.cards.cards["brakka"] = rankings.Character{Character: rankings.CharacterRef{Key: "us/normal/brakka", Region: "us",
			Ruleset: "normal", Name: "Brakka", Class: "warrior", Level: ref(60)}}
		m := w.command(t, pressDiscord, "bis", map[string]string{"spec": "fury", "slot": "head"}).Message
		if got := textsOf(m)[0]; got != "### Head · Fury Warrior\nBest in slot · level 60 Phase 1 raid · Alliance" {
			t.Fatalf("head = %q", got)
		}
		bare := w.command(t, pressDiscord, "bis", map[string]string{"spec": "fury", "slot": "head", "phase": "60"}).Message
		texts := textsOf(bare)
		if texts[0] != "### Head · Fury Warrior\nBest in slot · level 60 · Alliance" || !strings.HasPrefix(texts[1], "**Pick head**") ||
			!strings.Contains(bare.Components[0][len(bare.Components[0])-1].URL, "preset=bare") {
			t.Fatalf("bare = %q", texts)
		}
	})
}

// sourceOf is m as the pressed message an interaction carries: its flags and components.
func sourceOf(t *testing.T, m Message) sourceMessage {
	t.Helper()
	wire := wireOf(t, m)
	raw, err := json.Marshal(wire["components"])
	if err != nil {
		t.Fatal(err)
	}
	flags, _ := wire["flags"].(float64)
	return sourceMessage{Flags: int(flags), Components: raw}
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
