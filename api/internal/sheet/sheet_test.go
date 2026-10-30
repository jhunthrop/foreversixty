package sheet

import (
	"crypto/sha256"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"testing"
	"time"
)

func testRenderer(t *testing.T) *renderer {
	t.Helper()
	r, err := newRenderer(testIcons())
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func fixtureSlot(t *testing.T, fixture, key string) Slot {
	t.Helper()
	return loadFixture(t, fixture).slot(key)
}

func fit(t *testing.T, kind string, sl Slot) cellText {
	t.Helper()
	ct, err := testRenderer(t).fitCell(kind, sl)
	if err != nil {
		t.Fatal(err)
	}
	return ct
}

func TestFitRuleNames(t *testing.T) {
	gloves := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "hands"))
	if len(gloves.name) != 1 || gloves.nameStyle.size != nameMaxPx || gloves.name[0] != "Raider Gloves" {
		t.Fatalf("Raider Gloves: %q at %vpx", gloves.name, gloves.nameStyle.size)
	}
	signet := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "finger1"))
	if len(signet.name) != 2 || signet.nameStyle.size != nameMinPx || signet.nameLH != nameTwoLH {
		t.Fatalf("Signet: %q at %v/%v", signet.name, signet.nameStyle.size, signet.nameLH)
	}
	if strings.Join(signet.name, " ") != "Signet Ring of the Bronze Dragonflight" || len(signet.meta) != 1 {
		t.Fatalf("Signet: name %q meta %q", signet.name, signet.meta)
	}
	belt := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "waist"))
	if len(belt.name) != 1 || belt.nameStyle.size >= nameMaxPx || belt.nameStyle.size < nameMinPx {
		t.Fatalf("Belt of Preserved Heads should shrink on one line: %q at %vpx", belt.name, belt.nameStyle.size)
	}
}

func TestFitRuleMeta(t *testing.T) {
	// Every meta line leads with the slot's label; the full source first, then the short form.
	waist := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "waist"))
	if strings.Join(waist.meta, " ") != "Waist · Quest: A Collection of Heads · requires level 58" || len(waist.meta) != 2 {
		t.Fatalf("waist meta %q", waist.meta)
	}
	head := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "head"))
	if !slices.Equal(head.meta, []string{"Head · Quest", "requires level 58"}) {
		t.Fatalf("head meta %q", head.meta)
	}
	neck := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "neck"))
	if !slices.Equal(neck.meta, []string{"Neck · Quest", "requires level 55"}) {
		t.Fatalf("neck meta %q", neck.meta)
	}
	back := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "back"))
	if !slices.Equal(back.meta, []string{"Back · Stratholme", "Baron Rivendare"}) {
		t.Fatalf("back meta %q", back.meta)
	}
	signet := fit(t, KindBis, fixtureSlot(t, "sheet-bis", "finger1"))
	if !slices.Equal(signet.meta, []string{"Ring 1 · Vendor: Anachronos"}) {
		t.Fatalf("signet meta %q", signet.meta)
	}
	// A source that fits no candidate leaves the label alone: the slot is always named.
	none := fit(t, KindBis, Slot{Slot: "head", Name: "Cap", Quality: 4, Icon: "inv_helmet_24",
		Meta: "Vendor: Sorcerer-Thane-Thaurissan-of-the-Shadowforge-Dynasty-and-All-Its-Halls"})
	if !slices.Equal(none.meta, []string{"Head"}) {
		t.Fatalf("a source that fits no candidate kept %q", none.meta)
	}
	gold := fit(t, KindChar, fixtureSlot(t, "sheet-char", "legs"))
	if !slices.Equal(gold.meta, []string{"Legs · item level 71"}) || gold.gold != "+9.3 DPS" {
		t.Fatalf("char legs %q %q", gold.meta, gold.gold)
	}
}

func TestShortSource(t *testing.T) {
	cases := map[string]string{
		"Quest: Saving the Best for Last · requires level 58": "Quest · requires level 58",
		"Quest: The Untitled":                                 "Quest",
		"PvP rank 11 · Knight-Lieutenant · Alliance":          "PvP rank 11 · Alliance",
		"Crafted: Leatherworking · Pattern: Dawn Armor":       "Crafted: Leatherworking",
		"Stratholme · Baron Rivendare":                        "Strat · Baron Rivendare",
		"12% from General Drakkisath (Upper Blackrock Spire)": "12% from General Drakkisath (UBRS)",
		"Dire Maul · Immol'thar":                              "Dire Maul · Immol'thar",
		"Vendor: Mokvar":                                      "Vendor: Mokvar",
		"Lower Blackrock Spire trash":                         "LBRS trash",
	}
	for in, want := range cases {
		if got := ShortSource(in); got != want {
			t.Errorf("ShortSource(%q) = %q, want %q", in, got, want)
		}
	}
}

// labelText is an empty cell's line as read, and the part of it drawn in gold.
func labelText(ct cellText) (all, gold string) {
	var words, golds []string
	for _, line := range ct.label {
		for _, w := range line {
			words = append(words, w.text)
			if w.col == colorGold {
				golds = append(golds, w.text)
			}
		}
	}
	return strings.Join(words, " "), strings.Join(golds, " ")
}

func TestEveryCellState(t *testing.T) {
	cases := []struct {
		name, kind string
		slot       Slot
		wantName   string
		wantMeta   []string
		wantGold   string
		wantColor  color.RGBA
	}{
		{"bis pick", KindBis, Slot{Slot: "feet", Name: "Beastmaster's Treads", Quality: 4, Icon: "inv_boots_plate_07", Meta: "Vendor: Mokvar"},
			"Beastmaster's Treads", []string{"Feet · Vendor: Mokvar"}, "", qualityText[4]},
		{"char worn, no gain", KindChar, Slot{Slot: "neck", Name: "Amulet of the Darkmoon", Quality: 4, Icon: "inv_jewelry_necklace_14", ItemLevel: 65},
			"Amulet of the Darkmoon", []string{"Neck · item level 65"}, "", qualityText[4]},
		{"char worn, gain", KindChar, Slot{Slot: "chest", Name: "Shadowcraft Tunic", Quality: 3, Icon: "inv_chest_leather_07", ItemLevel: 63, Gain: "+11.0 DPS"},
			"Shadowcraft Tunic", []string{"Chest · item level 63"}, "+11.0 DPS", qualityText[3]},
		{"char worn, tank gain", KindChar, Slot{Slot: "chest", Name: "Shadowcraft Tunic", Quality: 3, Icon: "inv_chest_leather_07", ItemLevel: 63, Gain: "+4.0 points"},
			"Shadowcraft Tunic", []string{"Chest · item level 63"}, "+4.0 points", qualityText[3]},
		{"char worn, not modelled", KindChar, Slot{Slot: "head", Name: "Outlaw's Collar", Quality: 4, Icon: "inv_misc_bandana_03", ItemLevel: 65, Gain: CopyNotModelled},
			"Outlaw's Collar", []string{"Head · item level 65"}, CopyNotModelled, qualityText[4]},
		{"unknown item", KindBis, Slot{Slot: "head", Name: "Mystery Hood", Quality: -1, Meta: "Vendor: Mokvar"},
			"Mystery Hood", []string{"Head · Vendor: Mokvar"}, "", colorText},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ct := fit(t, c.kind, c.slot)
			if strings.Join(ct.name, " ") != c.wantName || !slices.Equal(ct.meta, c.wantMeta) || ct.gold != c.wantGold || ct.nameColor != c.wantColor {
				t.Fatalf("got name %q meta %q gold %q colour %v", ct.name, ct.meta, ct.gold, ct.nameColor)
			}
			cl, err := testRenderer(t).layoutCell(c.kind, c.slot)
			if err != nil {
				t.Fatal(err)
			}
			if cl.empty || cl.icon == nil {
				t.Fatalf("socket: empty %v icon %v", cl.empty, cl.icon != nil)
			}
			if c.name == "unknown item" && cl.ring != colorEmptyRing {
				t.Fatalf("unknown item ring %v", cl.ring)
			}
		})
	}
}

func TestEveryEmptyState(t *testing.T) {
	cases := []struct {
		name, kind         string
		slot               Slot
		wantText, wantGold string
	}{
		{"char empty, upgrade", KindChar, Slot{Slot: "trinket2", Empty: true, Gain: CopyUpgrade},
			"Trinket 2 · Upgrade available", "Upgrade available"},
		{"char empty, not modelled", KindChar, Slot{Slot: "trinket2", Empty: true, Gain: CopyNotModelled},
			"Trinket 2 · Gain not modelled yet", "Gain not modelled yet"},
		{"char empty, measured gain", KindChar, Slot{Slot: "finger1", Empty: true, Gain: "+6.2 DPS"},
			"Ring 1 · +6.2 DPS", "+6.2 DPS"},
		{"char empty, no pick", KindChar, Slot{Slot: "trinket2", Empty: true},
			"Trinket 2 · Empty", ""},
		{"bis no pick", KindBis, Slot{Slot: "trinket2", Empty: true},
			"Trinket 2 · No pick", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ct := fit(t, c.kind, c.slot)
			text, gold := labelText(ct)
			if text != c.wantText || gold != c.wantGold || ct.name != nil || ct.meta != nil || ct.gold != "" {
				t.Fatalf("got %q (gold %q), name %q meta %q", text, gold, ct.name, ct.meta)
			}
			if len(ct.label) > 3 {
				t.Fatalf("%d lines overflow the cell", len(ct.label))
			}
			cl, err := testRenderer(t).layoutCell(c.kind, c.slot)
			if err != nil || !cl.empty {
				t.Fatalf("socket: empty %v err %v", cl.empty, err)
			}
		})
	}
}

func TestTwoHandLeavesTheOffHandOut(t *testing.T) {
	for _, p := range placements(true) {
		if p.key == SlotOff {
			t.Fatal("the off hand is placed under a two-handed main hand")
		}
	}
	if len(placements(true)) != 16 || len(placements(false)) != 17 {
		t.Fatalf("placements %d and %d", len(placements(true)), len(placements(false)))
	}
	s := loadFixture(t, "sheet-bis")
	s.Slots = append(s.Slots, Slot{Slot: SlotOff, Name: "Bulwark", Quality: 4, Icon: "inv_spear_05", Meta: "Vendor: Mokvar"})
	// The off-hand socket's ring sits at its top edge's midpoint.
	px, py := at(rightX+cellW-sockW/2, weaponsRow)
	ringAt := image.Pt(int(px), int(py)+1)
	for _, twoHand := range []bool{true, false} {
		s.TwoHand = twoHand
		png, err := Render(s, testIcons())
		if err != nil {
			t.Fatal(err)
		}
		got := decodePNG(t, png).At(ringAt.X, ringAt.Y)
		ring := closeColor(got, qualityRing[4])
		if ring == twoHand {
			t.Fatalf("twoHand %v: off-hand ring drawn %v (pixel %v)", twoHand, ring, got)
		}
	}
}

func TestEveryIconTopSitsOnThePitch(t *testing.T) {
	at := map[string]int{}
	for _, p := range placements(false) {
		if (p.y-row0)%pitch != 0 {
			t.Errorf("%s at y %d is off the %d px pitch", p.key, p.y, pitch)
		}
		at[p.key] = p.y
	}
	// Both columns are top-aligned, as in the game's frame: Head beside Hands, Wrist beside Ring 2.
	if at["head"] != row0 || at["head"] != at["hands"] || at["wrist"] != at["finger2"] {
		t.Fatalf("columns not top-aligned: head %d hands %d wrist %d ring 2 %d", at["head"], at["hands"], at["wrist"], at["finger2"])
	}
}

func TestNoEllipsisAnywhere(t *testing.T) {
	r := testRenderer(t)
	for _, name := range []string{"sheet-bis", "sheet-char"} {
		s := loadFixture(t, name)
		if strings.ContainsRune(Alt(s), '…') {
			t.Errorf("%s alt has an ellipsis", name)
		}
		for _, sl := range s.Slots {
			ct, err := r.fitCell(s.Kind, sl)
			if err != nil {
				t.Fatal(err)
			}
			all := strings.Join(append(append(slices.Clone(ct.name), ct.meta...), ct.gold), " ")
			if strings.ContainsRune(all, '…') || strings.Contains(all, "...") {
				t.Errorf("%s %s: %q", name, sl.Slot, all)
			}
		}
	}
}

func TestAltMatchesTheMock(t *testing.T) {
	var want map[string]string
	readJSON(t, filepath.Join("testdata", "alts.json"), &want)
	for kind, fixture := range map[string]string{"bis": "sheet-bis", "char": "sheet-char"} {
		got := Alt(loadFixture(t, fixture))
		if got != want[kind] {
			t.Errorf("%s alt\n got %q\nwant %q", kind, got, want[kind])
		}
		if n := len([]rune(got)); n > AltMax {
			t.Errorf("%s alt is %d characters", kind, n)
		}
	}
}

func TestAltOverTheLimitDropsGainsThenWholeSlots(t *testing.T) {
	s := loadFixture(t, "sheet-char")
	for i := range s.Slots {
		if !s.Slots[i].Empty {
			s.Slots[i].Name = strings.Repeat("Longname ", 5) + s.Slots[i].Name
		}
	}
	full, noGains := fullAltOf(s), noGainsOf(s)
	got := Alt(s)
	if len([]rune(full)) <= AltMax || len([]rune(got)) > AltMax {
		t.Fatalf("full %d, got %d", len([]rune(full)), len([]rune(got)))
	}
	if strings.Contains(got, altGainTail) || !strings.HasSuffix(got, ".") || strings.ContainsRune(got, '…') {
		t.Fatalf("alt %q", got)
	}
	if !strings.HasPrefix(noGains, strings.TrimSuffix(got, ".")) {
		t.Fatalf("the alt was not cut at a whole slot: %q", got)
	}
	// Gains alone over the limit: dropping them is enough, and every slot stays.
	s = loadFixture(t, "sheet-char")
	s.Guild += strings.Repeat("G", AltMax-len([]rune(noGainsOf(s))))
	if len([]rune(noGainsOf(s))) != AltMax || len([]rune(fullAltOf(s))) <= AltMax {
		t.Fatalf("setup: no gains %d", len([]rune(noGainsOf(s))))
	}
	if got := Alt(s); got != noGainsOf(s) {
		t.Fatalf("dropping the gains should fit:\n%q", got)
	}
}

func noGainsOf(s Sheet) string { return altLead(s) + strings.Join(altSlots(s, false), "; ") + "." }

func fullAltOf(s Sheet) string { return altLead(s) + strings.Join(altSlots(s, true), "; ") + "." }

func TestAssetsAreTheSitesBytes(t *testing.T) {
	web := filepath.Join("..", "..", "..", "web", "public")
	pairs := map[string]string{
		"assets/seal-400.png":          filepath.Join(web, "seal-400.png"),
		"assets/faction/alliance.webp": filepath.Join(web, "icons", "hd", "faction", "alliance.webp"),
		"assets/faction/horde.webp":    filepath.Join(web, "icons", "hd", "faction", "horde.webp"),
	}
	for slug := range classColors {
		pairs["assets/crests/"+slug+".webp"] = filepath.Join(web, "icons", "hd", "crests", slug+".webp")
	}
	for embedded, site := range pairs {
		a, err := assetFS.ReadFile(embedded)
		if err != nil {
			t.Fatal(err)
		}
		b, err := os.ReadFile(site)
		if err != nil {
			t.Fatal(err)
		}
		if sha256.Sum256(a) != sha256.Sum256(b) {
			t.Errorf("%s differs from %s: run make sheet-assets", embedded, site)
		}
	}
}

func TestArtIsNeverUpscaled(t *testing.T) {
	crest, err := EmbeddedCrest("hunter")
	if err != nil {
		t.Fatal(err)
	}
	if crest.Bounds().Dx() < int(d(64)) {
		t.Fatalf("the crest is %d px, under its drawn %d", crest.Bounds().Dx(), int(d(64)))
	}
	icon, err := testIcons().Icon("inv_helmet_24")
	if err != nil {
		t.Fatal(err)
	}
	if icon.Bounds().Dx() != int(d(iconW)) {
		t.Fatalf("icons draw 1:1 at %d, the file is %d", int(d(iconW)), icon.Bounds().Dx())
	}
}

func TestRenderRefusesABadModel(t *testing.T) {
	good := loadFixture(t, "sheet-bis")
	for name, mutate := range map[string]func(*Sheet){
		"kind":      func(s *Sheet) { s.Kind = "raid" },
		"class":     func(s *Sheet) { s.Class = "monk" },
		"faction":   func(s *Sheet) { s.Faction = "neutral" },
		"slot":      func(s *Sheet) { s.Slots = append(s.Slots, Slot{Slot: "shirt"}) },
		"duplicate": func(s *Sheet) { s.Slots = append(s.Slots, s.Slots[0]) },
	} {
		s := good
		s.Slots = slices.Clone(good.Slots)
		mutate(&s)
		if _, err := Render(s, testIcons()); err == nil {
			t.Errorf("%s: rendered", name)
		}
	}
	if _, err := Render(good, nil); err == nil {
		t.Error("rendered without icons")
	}
}

func TestBudget(t *testing.T) {
	s := loadFixture(t, "sheet-bis")
	for i := range s.Slots {
		s.Slots[i].Name = "Hyper Deluxe Sniper Rifle Mk XVII of the Bronze Dragonflight"
		s.Slots[i].Meta = "Quest: The Longest Quest Title In The Whole Data Set · requires level 58"
	}
	png, err := Render(s, testIcons())
	if err != nil {
		t.Fatal(err)
	}
	if len(png) >= 1500<<10 {
		t.Fatalf("the sheet is %d bytes, over 1.5 MB", len(png))
	}
	if raceOn || testing.Short() {
		t.Skip("the timing budget is measured without -race")
	}
	const runs = 50
	times := make([]time.Duration, runs)
	for i := range runs {
		start := time.Now()
		if _, err := Render(s, testIcons()); err != nil {
			t.Fatal(err)
		}
		times[i] = time.Since(start)
	}
	sort.Slice(times, func(i, j int) bool { return times[i] < times[j] })
	p95 := times[runs*95/100-1]
	t.Logf("render p95 %v over %d (PNG %d KB)", p95, runs, len(png)>>10)
	if p95 > 250*time.Millisecond {
		t.Fatalf("render p95 %v, over 250 ms", p95)
	}
}

func TestRendererCachesBySheet(t *testing.T) {
	icons := &countingIcons{IconSource: testIcons()}
	r := NewRenderer(icons)
	s := loadFixture(t, "sheet-char")
	a, err := r.Render(s)
	if err != nil {
		t.Fatal(err)
	}
	loads := icons.n
	b, err := r.Render(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(a) != string(b) || icons.n != loads {
		t.Fatalf("second render loaded %d more icons", icons.n-loads)
	}
	s.Title = "Someone"
	if _, err := r.Render(s); err != nil || icons.n == loads {
		t.Fatalf("a different sheet came from the cache (err %v)", err)
	}
}

type countingIcons struct {
	IconSource
	n int
}

func (c *countingIcons) Icon(name string) (image.Image, error) {
	c.n++
	return c.IconSource.Icon(name)
}

func TestSlotWords(t *testing.T) {
	if SlotLabel("finger1") != "Ring 1" || SlotLabel("trinket2") != "Trinket 2" || SlotLabel("shirt") != "" {
		t.Fatal("slot labels")
	}
	if len(SlotKeys()) != 17 || SlotKeys()[0] != "head" || SlotKeys()[16] != "ranged" {
		t.Fatalf("slot keys %v", SlotKeys())
	}
}
