package discordbot

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// The lookups' fake build: its items and its BiS files, written to a temporary data dir.
const lookupBuild = "1.60.1.99999"

// lookupGuilds is the real guilds service with the build data and the readiness read scripted.
type lookupGuilds struct {
	*pressGuilds
	items     map[int]trees.Item
	dir       string
	readiness map[string]guilds.ReadinessFacts
	// specs are the guild's specs by character key (GuildSpecOf); a key missing here reads the
	// real service's.
	specs map[string]string
}

func (l *lookupGuilds) GuildSpecOf(ctx context.Context, guildID int64, key string) (string, error) {
	if sp, ok := l.specs[key]; ok {
		return sp, nil
	}
	return l.pressGuilds.GuildSpecOf(ctx, guildID, key)
}

func (l *lookupGuilds) ActiveBuild() (string, bool) { return lookupBuild, true }
func (l *lookupGuilds) BisDataDir() string          { return l.dir }
func (l *lookupGuilds) BisBuild() string            { return lookupBuild }
func (l *lookupGuilds) RaceFaction(name string) string {
	return map[string]string{"Orc": "horde", "Night Elf": "alliance"}[name]
}
func (l *lookupGuilds) BuildItem(version, _ string, id int) (trees.Item, bool) {
	it, ok := l.items[id]
	return it, ok && version == lookupBuild
}
func (l *lookupGuilds) CharacterReadinessFacts(_ context.Context, _ int64, key string) (guilds.ReadinessFacts, error) {
	if f, ok := l.readiness[key]; ok {
		return f, nil
	}
	return guilds.ReadinessFacts{Word: "Not shared"}, nil
}

// fakeCards are public character pages by lowercase name.
type fakeCards struct{ cards map[string]rankings.Character }

func (f *fakeCards) CharacterCard(_ context.Context, region, ruleset, name string) (rankings.Character, bool, error) {
	c, ok := f.cards[strings.ToLower(name)]
	if !ok || c.Character.Region != region || c.Character.Ruleset != ruleset {
		return rankings.Character{}, false, nil
	}
	return c, true, nil
}

// The fakeGuilds stubs the install tests' port needs for the lookups' methods.
func (f *fakeGuilds) BotGuildFacts(context.Context, int64) (guilds.BotGuild, error) {
	return guilds.BotGuild{}, nil
}
func (f *fakeGuilds) PublicRosterNames(context.Context, int64, string, int) ([]string, error) {
	return nil, nil
}
func (f *fakeGuilds) GuildSpecOf(context.Context, int64, string) (string, error) { return "", nil }
func (f *fakeGuilds) CharacterReadinessFacts(context.Context, int64, string) (guilds.ReadinessFacts, error) {
	return guilds.ReadinessFacts{}, nil
}
func (f *fakeGuilds) ActiveBuild() (string, bool)                      { return "", false }
func (f *fakeGuilds) BuildItem(string, string, int) (trees.Item, bool) { return trees.Item{}, false }
func (f *fakeGuilds) BisDataDir() string                               { return "" }
func (f *fakeGuilds) BisBuild() string                                 { return "" }
func (f *fakeGuilds) RaceFaction(string) string                        { return "" }

type lookupWorld struct {
	*pressWorld
	lg    *lookupGuilds
	cards *fakeCards
}

const keyZulmaraLookup = "us/normal/zulmara"

// zulmaraCard is a level 60 Marksmanship Hunter of Olympus XXVII with four gaps: head 9.0
// (worn id 200, which the band measured at −9.2), hands 7.0, feet 5.0, wrist 2.0.
func zulmaraCard() rankings.Character {
	slot := func(name string, worn, pick int, gap float64) rankings.GearSlot {
		return rankings.GearSlot{Slot: name, Worn: &rankings.GearWorn{ID: worn, ItemLevel: ref(55)},
			Pick: &rankings.GearPick{ID: pick, Name: "Pick " + name, Source: "Source " + name}, Gap: ref(gap)}
	}
	return rankings.Character{
		Character: rankings.CharacterRef{Key: keyZulmaraLookup, Region: "us", Ruleset: "normal", Name: "Zulmara", Class: "hunter",
			Level: ref(60), Guild: &rankings.CharacterGuild{Name: "Olympus XXVII"}},
		GearSpec: ref("hunter-marksmanship"),
		Gear: []rankings.GearSlot{
			slot("feet", 201, 101, 5.0), slot("head", 200, 100, 9.0), slot("wrist", 203, 103, 2.0), slot("hands", 202, 102, 7.0),
			{Slot: "neck", Worn: &rankings.GearWorn{ID: 104, ItemLevel: ref(60)}, Pick: &rankings.GearPick{ID: 104, Name: "Same"}, Gap: ref(0.0)},
		},
	}
}

func writeBis(t *testing.T, dir, spec string, bands any) {
	t.Helper()
	path := filepath.Join(dir, lookupBuild, "bis")
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(map[string]any{"bands": bands})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(path, spec+".json"), raw, 0o600); err != nil {
		t.Fatal(err)
	}
}

func newLookupWorld(t *testing.T) *lookupWorld {
	t.Helper()
	pw := newPressWorld(t)
	pw.exec(t, `update guilds set faction = 'alliance' where id = $1`, pw.gid)
	dir := t.TempDir()
	writeBis(t, dir, "hunter-marksmanship", []map[string]any{{"band": 60, "faction": "alliance", "slots": []map[string]any{
		{"slot": "head", "item_id": 100, "item_name": "Pick head", "alternatives": []map[string]any{
			{"item_id": 200, "item_name": "Worn head", "dps_delta": -9.2}}},
	}}})
	items := map[int]trees.Item{}
	for id := 100; id < 160; id++ {
		items[id] = trees.Item{ID: id, Name: "Item", Quality: 4, ItemLevel: 61, Icon: "inv_helmet_01"}
	}
	w := &lookupWorld{pressWorld: pw, cards: &fakeCards{cards: map[string]rankings.Character{"zulmara": zulmaraCard()}}}
	w.lg = &lookupGuilds{pressGuilds: pw.port, items: items, dir: dir, readiness: map[string]guilds.ReadinessFacts{},
		specs: map[string]string{keyZulmaraLookup: "hunter-marksmanship"}}
	w.svc = New(Deps{Store: pw.st, REST: pw.rest, Guilds: w.lg, Characters: w.cards,
		PublicBaseURL: pressPublic, APIBaseURL: pressAPI, Log: slog_(pw.log)})
	return w
}

func (w *lookupWorld) command(t *testing.T, discordUser, name string, options map[string]string) Reply {
	t.Helper()
	opts := []map[string]any{}
	for k, v := range options {
		opts = append(opts, map[string]any{"name": k, "value": v})
	}
	in := interactionOf(typeCommand, discordUser, map[string]any{"name": name, "options": opts})
	h := w.svc.commands[name]
	rep, err := h(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

// assertMessageLimits holds Discord's limits (spec §3.3.0): at most 10 embeds, 6000
// characters across them, and the footer on the last embed alone.
func assertMessageLimits(t *testing.T, m Message) {
	t.Helper()
	if len(m.Embeds) > 10 {
		t.Fatalf("%d embeds", len(m.Embeds))
	}
	total := 0
	for i, e := range m.Embeds {
		total += len([]rune(e.Title + e.Description + e.Footer + e.Author))
		for _, f := range e.Fields {
			total += len([]rune(f.Name + f.Value))
		}
		if (e.Footer != "") != (i == len(m.Embeds)-1) {
			t.Fatalf("embed %d footer %q: only the last embed carries one", i, e.Footer)
		}
	}
	if total > 6000 {
		t.Fatalf("%d characters across the embeds", total)
	}
	if _, err := json.Marshal(m); err != nil {
		t.Fatal(err)
	}
}

// TestLookupCopyMatchesTheSpec holds the bot spec's §6.2 strings the gear panel keeps (spec
// 2026-10-29 retired the embeds' lead, author, fields and footers; TestPanelCopyMatchesTheSpec
// holds their replacements).
func TestLookupCopyMatchesTheSpec(t *testing.T) {
	cell := func(label string) []string { return specCell(t, pressHeading, label) }
	for _, c := range []struct {
		name string
		got  []string
		kept []string
	}{
		{"char", cell("`/char`"), []string{lookupCopy.CharLine, lookupCopy.NoGear, lookupCopy.CharacterPage}},
		{"bis", cell("`/bis`"), []string{lookupCopy.BisOpen, lookupCopy.BisAllSlots,
			bandLabels["20"], bandLabels["30"], bandLabels["40"], bandLabels["50"], bandLabels["60"], bandLabels["p1"]}},
		{"item", cell("Item embed"), append([]string{lookupCopy.ItemLineOne, lookupCopy.ItemLineTwo}, qualityNames()...)},
		{"share", cell("Share"), []string{lookupCopy.Share, lookupCopy.Shared, lookupCopy.NoLongerPublic}},
	} {
		for _, want := range c.kept {
			if !slices.Contains(c.got, want) {
				t.Errorf("%s: spec %q lacks %q", c.name, c.got, want)
			}
		}
	}
}

// TestColoursAreTheSiteTokens: the class and quality colours are web/src/styles/tokens.css's
// --color-class-* and --color-rarity-* (never the -text variants).
func TestColoursAreTheSiteTokens(t *testing.T) {
	raw, err := os.ReadFile("../../../web/src/styles/tokens.css")
	if err != nil {
		t.Fatal(err)
	}
	token := func(name string) int {
		t.Helper()
		m := regexp.MustCompile(`--` + name + `:\s*#([0-9a-fA-F]{6});`).FindStringSubmatch(string(raw))
		if m == nil {
			t.Fatalf("tokens.css has no %s", name)
		}
		v, _ := strconv.ParseInt(m[1], 16, 32)
		return int(v)
	}
	for _, c := range classChoices {
		colour, ok := classColor(c.Value)
		if want := token("color-class-" + c.Value); !ok || colour != want {
			t.Errorf("%s = %06x, want %06x", c.Value, colour, want)
		}
	}
	for q, word := range qualityNames() {
		if want := token("color-rarity-" + strings.ToLower(word)); qualityColor(q) != want {
			t.Errorf("%s = %06x, want %06x", word, qualityColor(q), want)
		}
	}
}

func TestItemEmbedGrammar(t *testing.T) {
	it := ItemFacts{ID: 1, Name: "Champion's Plate Helm", Quality: 4, ItemLevel: 61,
		Icon: "https://site.test/data/b/icons/inv_helmet_01.webp", Source: "PvP rank 14 · Lieutenant Commander · Alliance",
		URL: "https://site.test/bis/warrior/fury"}
	e := ItemEmbed(it, "Head", signedDPS(-14.96), "Best in slot")
	if e.Color != 10696174 {
		t.Fatalf("epic colour = %d, want #a335ee", e.Color)
	}
	if e.Thumbnail != it.Icon || e.URL != it.URL || e.Author != "Best in slot" || e.Title != `Champion's Plate Helm` {
		t.Fatalf("embed = %+v", e)
	}
	if want := "Head · Epic · item level 61\n−15.0 DPS · PvP rank 14 · Lieutenant Commander · Alliance"; e.Description != want {
		t.Fatalf("description = %q, want %q", e.Description, want)
	}
	if got := ItemEmbed(ItemFacts{Name: "x", Quality: 2, ItemLevel: 30, Source: "Blacksmithing"}, "", "", "").Description; got != "Uncommon · item level 30\nBlacksmithing" {
		t.Fatalf("no slot, no delta = %q", got)
	}
	if got := ItemEmbed(ItemFacts{Name: "x", Quality: -1}, "Head", "+9.2 DPS", ""); got.Description != "Head\n+9.2 DPS" || got.Color != 0 {
		t.Fatalf("unknown item = %+v", got)
	}
	if signedDPS(9.24) != "+9.2 DPS" || signedPoints(12.35) != "+12.4 points" {
		t.Fatalf("figures %q %q", signedDPS(9.24), signedPoints(12.35))
	}
}

func TestCharCard(t *testing.T) {
	t.Run("panel and upgrades in points", func(t *testing.T) {
		w := newLookupWorld(t)
		// The band measured nothing she wears, so the list reads in points from the gear panel.
		writeBis(t, w.lg.dir, "hunter-marksmanship", []map[string]any{{"band": 60, "faction": "alliance", "slots": []map[string]any{
			{"slot": "head", "item_id": 100, "item_name": "Pick head", "alternatives": []map[string]any{
				{"item_id": 199, "item_name": "Other head", "dps_delta": -9.2}}},
		}}})
		w.port.candidates = []guilds.SignupCandidate{{CharacterKey: keyZulmaraLookup, Name: "Zulmara", Verified: true, Main: true}}
		w.lg.readiness[keyZulmaraLookup] = guilds.ReadinessFacts{Word: "Stocked", SyncedAt: ref(time.Unix(1700000000, 0)), OnRoster: true}
		m := w.command(t, pressDiscord, "char", map[string]string{"name": "zulmara"}).Message
		assertCardLimits(t, m)
		head := headOf(t, m)
		if head.Title != "Zulmara" || head.URL != pressPublic+"/character/"+keyZulmaraLookup || head.Color != 0xaad372 ||
			head.Thumbnail != pressAPI+"/v1/discord/tiles/crest/hunter.png" ||
			head.Description != "Level 60 Marksmanship Hunter · Olympus XXVII\n-# Item level 56 · Readiness: Stocked · Synced <t:1700000000:R>" {
			t.Fatalf("header = %+v", head)
		}
		texts := textsOf(m)
		if texts[1] != "**Top upgrades** · level 60\n"+panelCopy.UpgradesByScore {
			t.Fatalf("lead = %q", texts[1])
		}
		// Nothing measured, so the list reads in points, largest gap first: the top 3 of head, hands,
		// feet, wrist; the fourth rides the link's "All 4 upgrades".
		want := []string{"**Pick head**\nHead · +9.0 points\n-# item level 61 · Source head",
			"**Pick hands**\nHands · +7.0 points\n-# item level 61 · Source hands",
			"**Pick feet**\nFeet · +5.0 points\n-# item level 61 · Source feet"}
		if got := texts[2:]; strings.Join(got, "|") != strings.Join(want, "|") {
			t.Fatalf("upgrades = %q", got)
		}
		if s := sectionsOf(m)[1]; s.Accessory.Media != pressAPI+"/v1/discord/tiles/"+lookupBuild+"/100.png" || s.Accessory.Description != "Pick head, Epic head item" {
			t.Fatalf("tile = %+v", s.Accessory)
		}
		row := m.Components[0]
		if len(row) != 2 || row[0].Label != lookupCopy.Share || row[0].Style != StyleSecondary ||
			row[1].Label != "All 4 upgrades" || row[1].URL != pressPublic+"/character/"+keyZulmaraLookup {
			t.Fatalf("row = %+v", row)
		}
	})
	t.Run("DPS against the band /bis defaults to, the same gain as /bis", func(t *testing.T) {
		measured := func(slot string, pick, worn int, delta float64) map[string]any {
			return map[string]any{"slot": slot, "item_id": pick, "item_name": "Pick " + slot, "source_kind": "vendor",
				"source": "Somebody", "alternatives": []map[string]any{{"item_id": worn, "item_name": "Worn " + slot, "dps_delta": delta}}}
		}
		all := []map[string]any{measured("head", 100, 200, -9.2), measured("hands", 150, 202, -3.1), measured("feet", 101, 201, -4.4)}
		bands := func(faction string) []map[string]any {
			return []map[string]any{
				{"band": 60, "faction": faction, "preset": "bare", "slots": []map[string]any{measured("head", 100, 200, -1)}},
				{"band": 60, "faction": faction, "preset": "raid", "slots": all},
			}
		}
		w := newLookupWorld(t)
		writeBis(t, w.lg.dir, "hunter-marksmanship", bands("alliance"))
		m := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
		texts := textsOf(m)
		if texts[1] != "**Top upgrades** · level 60 Phase 1 raid" {
			t.Fatalf("lead = %q", texts[1])
		}
		// The raid band's picks (hands is the band's own 150, not the panel's 102), largest gain
		// first, each the band's measured gain over what she wears, with the page's source line.
		for i, want := range []string{"Head · +9.2 DPS\n-# item level 61 · Vendors", "Feet · +4.4 DPS\n-# item level 61 · Vendors",
			"Hands · +3.1 DPS\n-# item level 61 · Vendors"} {
			if !strings.Contains(texts[i+2], want) {
				t.Fatalf("upgrade %d = %q, want %q", i, texts[i+2], want)
			}
		}

		// The character's own faction wins over the guild's.
		writeBis(t, w.lg.dir, "hunter-marksmanship", bands("horde"))
		c := zulmaraCard()
		c.Character.Race = ref("Orc")
		w.cards.cards["zulmara"] = c
		m = w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
		if !strings.Contains(textsOf(m)[2], "+9.2 DPS") {
			t.Fatalf("horde character = %q", textsOf(m)[2])
		}
	})
	t.Run("a viewer outside the guild sees no readiness", func(t *testing.T) {
		w := newLookupWorld(t)
		m := w.command(t, "d-stranger", "char", map[string]string{"name": "Zulmara"}).Message
		if got := headOf(t, m).Description; got != "Level 60 Marksmanship Hunter · Olympus XXVII\n-# Item level 56" {
			t.Fatalf("header = %q", got)
		}
	})
	t.Run("no spec known says the bare class, as the site does", func(t *testing.T) {
		w := newLookupWorld(t)
		delete(w.lg.specs, keyZulmaraLookup) // the card's gear spec is Marksmanship; the guild knows none
		m := w.command(t, "d-stranger", "char", map[string]string{"name": "Zulmara"}).Message
		if got := headOf(t, m).Description; got != "Level 60 Hunter · Olympus XXVII\n-# Item level 56" {
			t.Fatalf("header = %q", got)
		}
	})
	t.Run("the more line counts the board's upgrades", func(t *testing.T) {
		w := newLookupWorld(t)
		w.port.candidates = []guilds.SignupCandidate{{CharacterKey: keyZulmaraLookup, Name: "Zulmara", Verified: true, Main: true}}
		w.lg.readiness[keyZulmaraLookup] = guilds.ReadinessFacts{Word: "Short: 13 gear upgrades waiting", OnRoster: true,
			GearChecked: true, GearUpgrades: 13}
		// The panel measures one upgrade (the band's head); the board counts 13: the rest ride the
		// character page link, labelled with the board's count.
		m := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
		var labels []string
		for _, row := range m.Components {
			labels = append(labels, buttonLabels(row)...)
		}
		if !slices.Contains(labels, "All 13 upgrades") {
			t.Fatalf("labels = %q (texts %q)", labels, textsOf(m))
		}
	})
	t.Run("tank specs rank by score", func(t *testing.T) {
		w := newLookupWorld(t)
		c := zulmaraCard()
		c.Character.Class, c.GearSpec = "druid", ref("druid-feral-bear")
		w.cards.cards["zulmara"] = c
		delete(w.lg.specs, keyZulmaraLookup) // with no guild spec, the gear spec still ranks the list
		texts := textsOf(w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message)
		if !strings.HasSuffix(texts[1], panelCopy.UpgradesByScore) || !strings.Contains(texts[2], "Head · +9.0 points") {
			t.Fatalf("tank = %q", texts)
		}
	})
	t.Run("healers rank by the BiS list and item level, never DPS", func(t *testing.T) {
		w := newLookupWorld(t)
		c := zulmaraCard()
		c.Character.Class, c.GearSpec = "priest", ref("priest-holy")
		w.cards.cards["zulmara"] = c
		delete(w.lg.specs, keyZulmaraLookup)
		slot := func(name string, pick int) map[string]any {
			return map[string]any{"slot": name, "item_id": pick, "item_name": "Pick " + name,
				"alternatives": []map[string]any{{"item_id": 999, "item_name": "Other", "dps_delta": -37.0}}}
		}
		writeBis(t, w.lg.dir, "priest-holy", []map[string]any{{"band": 60, "faction": "alliance", "preset": "raid", "slots": []map[string]any{
			slot("head", 100), slot("feet", 201), slot("wrist", 180), slot("trinket2", 105),
		}}})
		texts := textsOf(w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message)
		if !strings.HasSuffix(texts[1], "\n"+panelCopy.UpgradesByLevel) {
			t.Fatalf("lead = %q", texts[1])
		}
		// The empty trinket first, then the head 6 item levels up, then the wrist pick the build
		// gives no item level for; the feet already wear the pick.
		for i, want := range []string{"Trinket 2 · " + panelCopy.LevelEmpty, "Head · +6 item levels", "Wrist · " + panelCopy.LevelOnList} {
			if !strings.Contains(texts[i+2], want) {
				t.Fatalf("upgrade %d = %q, want %q", i, texts[i+2], want)
			}
		}
		if joined := strings.Join(texts, "\n"); strings.Contains(joined, "DPS gained") || strings.Contains(joined, "+37") {
			t.Fatalf("a healer list mentions DPS: %q", joined)
		}
	})
	t.Run("no gear shared", func(t *testing.T) {
		w := newLookupWorld(t)
		c := zulmaraCard()
		c.Gear = []rankings.GearSlot{}
		w.cards.cards["zulmara"] = c
		m := w.command(t, pressDiscord, "char", map[string]string{"name": "Zulmara"}).Message
		assertCardLimits(t, m)
		texts := textsOf(m)
		if len(texts) != 2 || texts[1] != "**Top upgrades** · level 60 Phase 1 raid\n"+lookupCopy.NoGear {
			t.Fatalf("no gear = %q", texts)
		}
	})
	t.Run("private or unknown", func(t *testing.T) {
		w := newLookupWorld(t)
		m := w.command(t, pressDiscord, "char", map[string]string{"name": "Nobody"}).Message
		if texts := textsOf(m); len(texts) != 1 || texts[0] != "No public character named Nobody. Try another name." || m.Content != "" {
			t.Fatalf("unknown = %+v", m)
		}
	})
}

func TestCharAutocompleteListsPublicRoster(t *testing.T) {
	w := newLookupWorld(t)
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	w.addChar(t, w.gid, w.uid, "us/normal/pending", "Pendrel", false)
	in := interactionOf(typeAutocomplete, "d-stranger", map[string]any{"name": "char",
		"options": []map[string]any{{"name": "name", "value": "", "focused": true}}})
	got, err := w.svc.autocompleteChar(context.Background(), in)
	if err != nil || len(got) != 1 || got[0].Value != "Zulmara" {
		t.Fatalf("choices = %+v, %v", got, err)
	}
}

// The "N more" counts the board's upgrades for a rostered character whose gear the board read,
// never fewer than the panel lists; anyone else counts the panel's own list.
func TestUpgradeTotal(t *testing.T) {
	four := upgradeList{items: make([]upgrade, 4)}
	board := func(checked bool, n int) *guilds.ReadinessFacts {
		return &guilds.ReadinessFacts{GearChecked: checked, GearUpgrades: n}
	}
	for _, c := range []struct {
		name  string
		facts *guilds.ReadinessFacts
		want  int
	}{
		{"the board's 13", board(true, 13), 13},
		{"never fewer than listed", board(true, 2), 4},
		{"the board read no gear", board(false, 0), 4},
		{"off the roster", nil, 4},
	} {
		if got := upgradeTotal(charData{ups: four, facts: c.facts}); got != c.want {
			t.Errorf("%s: total = %d, want %d", c.name, got, c.want)
		}
	}
	// 13 upgrades, 3 rows: the link reads "All 13 upgrades"; with every upgrade shown, the page.
	if got := characterPageLabel(maxUpgrades, upgradeTotal(charData{ups: four, facts: board(true, 13)})); got != "All 13 upgrades" {
		t.Errorf("label = %q", got)
	}
	if got := characterPageLabel(3, 3); got != lookupCopy.CharacterPage {
		t.Errorf("label = %q", got)
	}
}

// ref is a pointer to v.
func ref[T any](v T) *T { return &v }

// qualityNames are the item qualities in the game's words (sheet.QualityWord), Poor to Legendary.
func qualityNames() []string {
	out := make([]string, sheet.QualityCount)
	for q := range out {
		out[q] = sheet.QualityWord(q)
	}
	return out
}
