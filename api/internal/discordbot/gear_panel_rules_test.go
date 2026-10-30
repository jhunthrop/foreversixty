package discordbot

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

func TestItemRowGrammar(t *testing.T) {
	w := newGearWorld(t)
	ctx := context.Background()
	epic := ItemFacts{ID: 1, Name: "Cap [of] (the) *Bold*", Quality: 4, ItemLevel: 60, Source: "Vendor: Mokvar"}
	got := w.svc.itemRowText(ctx, itemRow{item: epic, slot: "finger1", gain: "+1.2 DPS", replaces: replacesLine("Old [Ring]", 0)})
	want := "<:fs_q_epic:1302600000000000004> **Cap \\[of\\] (the) \\*Bold\\***\nRing 1 · +1.2 DPS\n-# Epic · item level 60 · Vendor: Mokvar\n-# Replaces Old \\[Ring\\]"
	if got != want {
		t.Fatalf("row =\n%s\nwant\n%s", got, want)
	}
	if strings.Contains(got, "](") {
		t.Fatal("an item row carries a link")
	}
	if got := replacesLine("Shadowcraft Tunic", 63); got != "-# Replaces Shadowcraft Tunic, item level 63" || replacesLine("", 63) != "" {
		t.Fatalf("replaces = %q", got)
	}
	// An item the build does not know: no dot, no quality word, no item level.
	unknown := w.svc.itemRowText(ctx, itemRow{item: ItemFacts{Name: "Mystery", Quality: -1, Source: "Somewhere"}, slot: "trinket2"})
	if unknown != "**Mystery**\nTrinket 2\n-# Somewhere" {
		t.Fatalf("unknown = %q", unknown)
	}
	if got := tileAlt(ItemFacts{Name: "Mystery", Quality: -1}, "Trinket 2"); got != "Mystery, trinket 2 item" {
		t.Fatalf("unknown alt = %q", got)
	}
}

func TestAMissingQualityDotDropsOnlyThatDot(t *testing.T) {
	w := newGearWorld(t)
	w.emoji.items = nil
	for _, e := range mockDots() {
		if e.Name != "fs_q_epic" {
			w.emoji.items = append(w.emoji.items, e)
		}
	}
	ctx := context.Background()
	epic := w.svc.itemRowText(ctx, itemRow{item: ItemFacts{Name: "A", Quality: 4, ItemLevel: 60}, slot: "head"})
	rare := w.svc.itemRowText(ctx, itemRow{item: ItemFacts{Name: "B", Quality: 3, ItemLevel: 60}, slot: "head"})
	if !strings.HasPrefix(epic, "**A**\nHead\n-# Epic") || !strings.HasPrefix(rare, "<:fs_q_rare:1302600000000000003> **B**") {
		t.Fatalf("epic %q rare %q", epic, rare)
	}
	if log := w.log.String(); !strings.Contains(log, "emoji_missing") || !strings.Contains(log, "fs_q_epic") {
		t.Fatalf("log = %s", log)
	}
}

func TestCharUpgradeWords(t *testing.T) {
	if moreText(1) != "-# 1 more upgrade on the character page." || moreText(4) != "-# 4 more upgrades on the character page." || moreText(0) != "" {
		t.Fatalf("more = %q %q", moreText(1), moreText(4))
	}
	geared := upgradesText(gearZulmara(), upgradeList{byDPS: true, band: "level 60 Phase 1 raid"})
	if geared != "**Top upgrades** · level 60 Phase 1 raid\nEvery slot already wears the guide's pick for level 60 Phase 1 raid." {
		t.Fatalf("no upgrade = %q", geared)
	}
	// Ordered: empty slots in frame order, measured largest first, unmodelled in frame order.
	items := ordered([]upgrade{{slot: "head", group: groupUnmodelled, value: 11}, {slot: "chest", group: groupMeasured, value: 11},
		{slot: "wrist", group: groupMeasured, value: 4.7}, {slot: "feet", group: groupUnmodelled, value: 7.6},
		{slot: "trinket2", group: groupEmpty}, {slot: "ranged", group: groupMeasured, value: 23.1}})
	var got []string
	for _, u := range items {
		got = append(got, u.slot)
	}
	if strings.Join(got, ",") != "trinket2,ranged,chest,wrist,head,feet" {
		t.Fatalf("order = %v", got)
	}
}

func TestAnEmptySlotShowsItsRealGain(t *testing.T) {
	per := 0.14
	band := bis.Band{ReferenceDpsPerPoint: &per}
	score := 96.0
	head := bis.Slot{Slot: "head", ItemID: 1, Score: &score}
	if u, ok := dpsUpgrade(band, head, nil); !ok || u.gain != "+13.4 DPS" || u.group != groupEmpty {
		t.Fatalf("empty head = %+v", u)
	}
	for name, sl := range map[string]bis.Slot{
		"a sim-decided trinket": {Slot: "trinket2", ItemID: 2},
		"a weapon":              {Slot: "ranged", ItemID: 3, Score: &score},
		"an unmodelled effect":  {Slot: "feet", ItemID: 4, Score: &score, EffectUnmodelled: true},
	} {
		if u, ok := dpsUpgrade(band, sl, nil); !ok || u.gain != panelCopy.NotModelled {
			t.Errorf("%s = %+v", name, u)
		}
	}
}

func TestATankBandReadsInPoints(t *testing.T) {
	d := bisData{band: bis.Band{ScoreUnit: bis.TankScoreUnit}}
	if got := d.pickGain(bis.Slot{Alternatives: []bis.Alternative{{DpsDelta: -3.21}}}); got != "+3.2 points" {
		t.Fatalf("tank gain = %q", got)
	}
	if got := d.pickGain(bis.Slot{Alternatives: []bis.Alternative{{DpsDelta: -0.04}}}); got != "" {
		t.Fatalf("a tie shows %q", got)
	}
}

func TestTheTextGuardDropsSourcesLastRowFirst(t *testing.T) {
	w := newGearWorld(t)
	long := strings.Repeat("Name ", 24)
	var slots []bis.Slot
	for i, c := range slotChoices {
		slots = append(slots, bis.Slot{Slot: c.Value, ItemID: 900000 + i, ItemName: long, Source: strings.Repeat("Far away source ", 20),
			Alternatives: []bis.Alternative{{DpsDelta: -1}}})
	}
	sp, _ := specByKey("hunter-beast-mastery")
	band := bis.Band{Slots: slots}
	d := bisData{spec: sp, label: "level 60 Phase 1 raid", faction: "alliance", build: gearBuild, band: band, bySlot: band.BySlot()}
	m := w.svc.bisPagePanel(context.Background(), d, firstPage)
	texts := textsOf(m)
	if n := textLength(m.Layout); n > textBudget {
		t.Fatalf("%d characters", n)
	}
	if !strings.Contains(texts[1], "Far away source") || strings.Contains(texts[len(texts)-1], "Far away source") {
		t.Fatalf("the guard did not drop from the last row first: first %q last %q", texts[1], texts[len(texts)-1])
	}
}

// realBandGuilds serve the newest build in data/builds with its own item tables.
type realBandGuilds struct {
	*gearGuilds
	build string
	items map[string]map[int]trees.Item
}

func (g *realBandGuilds) BisBuild() string            { return g.build }
func (g *realBandGuilds) ActiveBuild() (string, bool) { return g.build, true }
func (g *realBandGuilds) BuildItem(_, class string, id int) (trees.Item, bool) {
	it, ok := g.items[class][id]
	return it, ok
}

// newestBuild is the newest build folder holding BiS files, class item tables and loot.json.
func newestBuild(t *testing.T, root string) string {
	t.Helper()
	dirs, _ := filepath.Glob(filepath.Join(root, "*", "bis"))
	sort.Strings(dirs)
	for i := len(dirs) - 1; i >= 0; i-- {
		build := filepath.Dir(dirs[i])
		if _, err := os.Stat(filepath.Join(build, "loot.json")); err == nil {
			if _, err := os.Stat(filepath.Join(build, "items")); err == nil {
				return filepath.Base(build)
			}
		}
	}
	t.Skip("no build in data/builds holds bis, items and loot.json")
	return ""
}

// TestEveryRealBandHoldsTheBudget draws both pages and the sheet's alt text of every band of
// every spec in the newest build: none needs the text guard, none passes 38 components, and
// every alt is at most 1,024 characters with no ellipsis (spec §A.6, §B.7).
func TestEveryRealBandHoldsTheBudget(t *testing.T) {
	root, err := filepath.Abs("../../../data/builds")
	if err != nil {
		t.Fatal(err)
	}
	build := newestBuild(t, root)
	g := &realBandGuilds{gearGuilds: &gearGuilds{lookupGuilds: &lookupGuilds{dir: root}}, build: build, items: map[string]map[int]trees.Item{}}
	for _, c := range classChoices {
		raw, err := os.ReadFile(filepath.Join(root, build, "items", c.Value+".json"))
		if err != nil {
			t.Fatal(err)
		}
		var file struct {
			Items []trees.Item `json:"items"`
		}
		if err := json.Unmarshal(raw, &file); err != nil {
			t.Fatal(err)
		}
		g.items[c.Value] = map[int]trees.Item{}
		for _, it := range file.Items {
			g.items[c.Value][it.ID] = it
		}
	}
	s := New(Deps{Guilds: g, Emoji: &countingEmoji{items: mockDots()}, PublicBaseURL: gearSite, APIBaseURL: gearAPI})
	ctx := context.Background()
	longest, bands := 0, 0
	for _, sp := range specs.All {
		raw, err := os.ReadFile(filepath.Join(root, build, "bis", sp.Spec+".json"))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		var file struct {
			Bands []struct {
				Band    int    `json:"band"`
				Faction string `json:"faction"`
				Preset  string `json:"preset"`
			} `json:"bands"`
		}
		if err := json.Unmarshal(raw, &file); err != nil {
			t.Fatal(err)
		}
		for _, b := range file.Bands {
			band, found, err := bis.LoadPresetBand(root, build, sp.Spec, b.Band, b.Faction, b.Preset)
			if err != nil || !found {
				t.Fatalf("%s %d %s: %v", sp.Spec, b.Band, b.Faction, err)
			}
			bands++
			d := bisData{spec: sp, label: bandLabels["p1"], level: b.Band, phase: phaseRaid, faction: b.Faction, build: build,
				page: s.bisPageURL(sp, b.Faction, b.Band, b.Preset), band: band, bySlot: band.BySlot()}
			for _, page := range []int{firstPage, lastPage} {
				m := withShare(decorateOwn("bis", map[string]string{optSpec: sp.Spec, optPhase: phaseRaid}, s.bisPagePanel(ctx, d, page)),
					shareButton(strings.Repeat("t", 43), false))
				m.Ephemeral = true
				m = withLiveShare(m, strings.Repeat("t", 43))
				for _, v := range cardViolations(m) {
					t.Errorf("%s %d %s page %d: %s", sp.Spec, b.Band, b.Faction, page, v)
				}
				for _, text := range textsOf(m) {
					if strings.HasPrefix(text, "<:") && !strings.Contains(text, "\n-# ") {
						t.Errorf("%s %d %s page %d: the guard dropped a source from %q", sp.Spec, b.Band, b.Faction, page, text)
					}
				}
				longest = max(longest, textLength(m.Layout))
			}
			if alt := sheet.Alt(s.bisSheet(d).model); len([]rune(alt)) > sheet.AltMax || strings.Contains(alt, "…") {
				t.Errorf("%s %d %s: alt of %d characters", sp.Spec, b.Band, b.Faction, len([]rune(alt)))
			}
		}
	}
	if bands == 0 {
		t.Fatal("no band drawn")
	}
	t.Logf("%d bands of build %s, longest page %d characters", bands, build, longest)
}

func TestShareOfASlotAnswerPostsThePanelWithNoFile(t *testing.T) {
	w := newGearWorld(t)
	slot := w.bisPanel(t, map[string]string{"spec": "hunter-beast-mastery", "phase": "p1", "slot": "legs"})
	rep := w.pressShareFrom(t, slot)
	rep.After(context.Background())
	ups := w.followUps()
	if len(ups) != 1 || len(ups[0].Message.Files) != 0 || len(w.sheets.models) != 0 {
		t.Fatalf("follow-ups = %+v", ups)
	}
	public := ups[0].Message
	if got := strings.Join(buttonLabels(public.Components[0]), "|"); got != lookupCopy.BisOpen || public.Ephemeral ||
		textsOf(public)[0] != textsOf(slot)[0] {
		t.Fatalf("public = %+v", public)
	}
}

func TestShareFromPageTwoKeepsThePageAndSheetsAllSlots(t *testing.T) {
	w := newGearWorld(t)
	p1 := w.bisPanel(t, gearBisOpts)
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": p1.Components[0][1].CustomID})
	in.Source = sourceOf(t, p1)
	p2 := w.route(t, in).Message
	rep := w.pressShareFrom(t, p2)
	if got := textsOf(rep.Message)[0]; !strings.Contains(got, "Feet to Ranged, 8 of 17") || rep.Message.Components[0][1].Label != "Head to Legs" {
		t.Fatalf("own copy after the share = %q", got)
	}
	rep.After(context.Background())
	model := w.sheets.models[0]
	if len(model.Slots) != 16 || !model.TwoHand || model.Subtitle != "Phase 1 raid" || model.Footer != "Updated 29 Oct 2026" ||
		model.Title != "Beast Mastery Hunter" || model.Kicker != "Best in slot" || model.Level != 60 {
		t.Fatalf("sheet = %+v", model)
	}
	if head := model.Slots[0]; head.Slot != "head" || head.Meta != "Quest: Saving the Best for Last · requires level 58" || head.Icon == "" || head.Quality != 4 {
		t.Fatalf("head cell = %+v", head)
	}
}

func TestTheCharSheetCellsCarryTheGold(t *testing.T) {
	w := newGearWorld(t)
	rep := w.pressShareFrom(t, w.command(t, pressDiscord, "char", nil).Message)
	rep.After(context.Background())
	model := w.sheets.models[0]
	cells := map[string]sheet.Slot{}
	for _, c := range model.Slots {
		cells[c.Slot] = c
	}
	if model.Kind != sheet.KindChar || model.Title != "Zulmara" || model.Kicker != "Level 60 Beast Mastery Hunter · Olympus XXVII" ||
		model.Subtitle != "63.5" || model.Footer != "Gold: gain over what is worn, level 60 Phase 1 raid" || model.FooterRight != "" || !model.TwoHand {
		t.Fatalf("sheet = %+v", model)
	}
	for slot, want := range map[string]string{"ranged": "+23.1 DPS", "head": panelCopy.NotModelled, "trinket2": sheet.CopyNotModelled, "neck": ""} {
		if cells[slot].Gain != want {
			t.Errorf("%s gain = %q, want %q", slot, cells[slot].Gain, want)
		}
	}
	if !cells["trinket2"].Empty || cells["ranged"].Name != "Dark Iron Rifle" || cells["ranged"].ItemLevel != 55 {
		t.Fatalf("cells = %+v", cells)
	}
}

// TestFollowUpUploadsTheSheetAsMultipart: a follow-up with a file is Discord's upload form,
// payload_json then files[0], the attachment described by id; the token stays in the path.
func TestFollowUpUploadsTheSheetAsMultipart(t *testing.T) {
	var got struct {
		path, contentType string
		payload           map[string]any
		file              []byte
		fileName, auth    string
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.path, got.auth = r.URL.Path, r.Header.Get("Authorization")
		media, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		got.contentType = media
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		form := multipart.NewReader(r.Body, params["boundary"])
		for {
			part, err := form.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			body, _ := io.ReadAll(part)
			switch part.FormName() {
			case "payload_json":
				_ = json.Unmarshal(body, &got.payload)
			case "files[0]":
				got.file, got.fileName = body, part.FileName()
			}
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	rest := NewHTTPREST("app1", "bot-secret-token")
	rest.BaseURL = srv.URL
	post := &sheetPost{model: sheet.Sheet{Kind: sheet.KindBis, Class: "hunter", Faction: "alliance"}, accent: 0xaad372, head: "### Head",
		link: Component{Kind: KindButton, Style: StyleLink, Label: lookupCopy.BisAllSlots, URL: gearSite}}
	png := []byte("\x89PNG\r\n\x1a\n rest")
	if err := rest.FollowUp(context.Background(), "tok-1", sheetMessage(post, png)); err != nil {
		t.Fatal(err)
	}
	if got.contentType != "multipart/form-data" || got.path != "/webhooks/app1/tok-1" || got.auth != "" {
		t.Fatalf("request = %s %s auth %q", got.contentType, got.path, got.auth)
	}
	if !bytes.Equal(got.file, png) || got.fileName != "sheet.png" {
		t.Fatalf("file %q named %q", got.file, got.fileName)
	}
	attachments, _ := got.payload["attachments"].([]any)
	if got.payload["flags"] != float64(componentsV2Flag) || len(attachments) != 1 {
		t.Fatalf("payload = %v", got.payload)
	}
	a := attachments[0].(map[string]any)
	alt := sheet.Alt(post.model)
	if a["id"] != float64(0) || a["filename"] != "sheet.png" || a["description"] != alt || alt == "" {
		t.Fatalf("attachment = %v", a)
	}
	gallery := got.payload["components"].([]any)[0].(map[string]any)["components"].([]any)[1].(map[string]any)
	item := gallery["items"].([]any)[0].(map[string]any)
	if gallery["type"] != float64(KindMediaGallery) || item["media"].(map[string]any)["url"] != "attachment://sheet.png" || item["description"] != alt {
		t.Fatalf("gallery = %v", gallery)
	}
}

// A share whose follow-up fails logs no token: neither the interaction's nor the bot's.
func TestAFailedShareLogsNoToken(t *testing.T) {
	w := newGearWorld(t)
	w.rest.ErrFor = map[string]error{"FollowUp": &APIError{Status: 500, Message: "boom"}}
	rep := w.pressShareFrom(t, w.bisPanel(t, gearBisOpts))
	rep.After(context.Background())
	if strings.Contains(w.log.String(), "tok-share") {
		t.Fatalf("the interaction token reached the log: %s", w.log.String())
	}
}
