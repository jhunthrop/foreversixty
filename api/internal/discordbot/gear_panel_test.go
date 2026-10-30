package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// The gear panel's reference world (spec 2026-10-29): the mock's real data, the active build's
// Beast Mastery Hunter band at level 60 Phase 1 raid, Alliance, cut to the items it names in
// testdata/gear-panel (from data/builds/1.60.1.70658: the band, the class item table's rows and
// loot.json's sources and quests for them). The payloads it is compared with are the mock's,
// design/mocks/2026-10-29-discord-gear-panel/payloads.
const (
	gearBuild    = "1.60.1.70658"
	gearSite     = "https://foreversixty.gg"
	gearAPI      = "https://api.foreversixty.gg"
	gearPayloads = "../../../design/mocks/2026-10-29-discord-gear-panel/payloads"
	gearData     = "testdata/gear-panel"
	gearSynced   = 1796767200
)

// gearGuilds is the lookup world's guilds over the reference build.
type gearGuilds struct {
	*lookupGuilds
	items map[int]trees.Item
}

func (g *gearGuilds) ActiveBuild() (string, bool) { return gearBuild, true }
func (g *gearGuilds) BisBuild() string            { return gearBuild }
func (g *gearGuilds) RaceFaction(string) string   { return "alliance" }
func (g *gearGuilds) BuildItem(version, _ string, id int) (trees.Item, bool) {
	it, ok := g.items[id]
	return it, ok && version == gearBuild
}

// stubSheets is the renderer with its answer scripted; it records the models it was given.
type stubSheets struct {
	png    []byte
	err    error
	models []sheet.Sheet
}

func (s *stubSheets) Render(m sheet.Sheet) ([]byte, error) {
	s.models = append(s.models, m)
	return s.png, s.err
}

type gearWorld struct {
	*lookupWorld
	gg     *gearGuilds
	sheets *stubSheets
	emoji  *countingEmoji
}

// mockDots are the six quality dots under the mock's emoji ids.
func mockDots() []AppEmoji {
	var out []AppEmoji
	for q := range qualityWords {
		out = append(out, AppEmoji{ID: fmt.Sprintf("13026%014d", q), Name: sheet.QualityEmojiName(q)})
	}
	return out
}

func newGearWorld(t *testing.T) *gearWorld {
	t.Helper()
	lw := newLookupWorld(t)
	raw, err := os.ReadFile(filepath.Join(gearData, gearBuild, "items.json"))
	if err != nil {
		t.Fatal(err)
	}
	var list []trees.Item
	if err := json.Unmarshal(raw, &list); err != nil {
		t.Fatal(err)
	}
	items := map[int]trees.Item{}
	for _, it := range list {
		items[it.ID] = it
	}
	dir, err := filepath.Abs(gearData)
	if err != nil {
		t.Fatal(err)
	}
	lw.lg.dir = dir
	w := &gearWorld{lookupWorld: lw, gg: &gearGuilds{lookupGuilds: lw.lg, items: items},
		sheets: &stubSheets{png: []byte("\x89PNG sheet")}, emoji: &countingEmoji{items: mockDots()}}
	w.svc = New(Deps{Store: lw.st, REST: lw.rest, Guilds: w.gg, Characters: lw.cards, Emoji: w.emoji, Sheets: w.sheets,
		PublicBaseURL: gearSite, APIBaseURL: gearAPI, Log: slog_(lw.log)})
	w.svc.now = func() time.Time { return boardClock }
	w.port.candidates = []guilds.SignupCandidate{
		{CharacterKey: keyZulmara, Name: "Zulmara", Class: "hunter", Spec: "Beast Mastery", Verified: true, Main: true},
		{CharacterKey: keyDuskstrider, Name: "Duskstrider", Class: "warrior", Spec: "Fury", Verified: true},
	}
	w.cards.cards["zulmara"] = gearZulmara()
	dusk := gearZulmara()
	dusk.Character.Key, dusk.Character.Name, dusk.Character.Class = keyDuskstrider, "Duskstrider", "warrior"
	dusk.GearSpec, dusk.Character.Level = ref("warrior-fury"), ref(58)
	w.cards.cards["duskstrider"] = dusk
	w.lg.readiness[keyZulmara] = guilds.ReadinessFacts{Word: "Ready", SyncedAt: ref(time.Unix(gearSynced, 0)), OnRoster: true}
	return w
}

// gearZulmara is the mock's example character: each slot the band's pick or one of its
// alternatives, Trinket 2 empty, a two-handed main hand.
func gearZulmara() rankings.Character {
	worn := []struct {
		slot  string
		id    int
		name  string
		level int
	}{
		{"head", 279253, "Outlaw's Collar", 65}, {"neck", 19491, "Amulet of the Darkmoon", 65},
		{"shoulder", 272105, "Darkspear Pauldrons", 65}, {"back", 18511, "Shifting Cloak", 62},
		{"chest", 16721, "Shadowcraft Tunic", 63}, {"wrist", 19587, "Forest Stalker's Bracers", 65},
		{"hands", 272099, "Raider Gloves", 65}, {"waist", 20216, "Belt of Preserved Heads", 70},
		{"legs", 231558, "Marshal's Chain Legplates", 71}, {"feet", 16711, "Shadowcraft Boots", 59},
		{"finger1", 234202, "Signet Ring of the Bronze Dragonflight", 80}, {"finger2", 9533, "Masons Fraternity Ring", 49},
		{"trinket1", 249469, "Frozen Heart of the Mountain", 55}, {"trinket2", 0, "", 0},
		{"main_hand", 250619, "Legionite Glaive", 63}, {"off_hand", 0, "", 0}, {"ranged", 16004, "Dark Iron Rifle", 55},
	}
	var gear []rankings.GearSlot
	for _, w := range worn {
		g := rankings.GearSlot{Slot: w.slot}
		if w.id > 0 {
			g.Worn = &rankings.GearWorn{ID: w.id, Name: ref(w.name), ItemLevel: ref(w.level)}
		}
		gear = append(gear, g)
	}
	return rankings.Character{
		Character: rankings.CharacterRef{Key: keyZulmara, Region: "us", Ruleset: "normal", Name: "Zulmara", Class: "hunter",
			Level: ref(60), Race: ref("Dwarf"), Guild: &rankings.CharacterGuild{Name: "Olympus XXVII"}},
		GearSpec: ref("hunter-beast-mastery"), Gear: gear,
	}
}

// mockMessage is a reference payload's message.
func mockMessage(t *testing.T, id string) map[string]any {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(gearPayloads, id+".json"))
	if err != nil {
		t.Fatal(err)
	}
	var frame struct {
		Message map[string]any `json:"message"`
	}
	if err := json.Unmarshal(raw, &frame); err != nil {
		t.Fatal(err)
	}
	return normalShares(frame.Message).(map[string]any)
}

// normalShares replaces every Share token with one placeholder: tokens are random.
func normalShares(v any) any {
	switch x := v.(type) {
	case map[string]any:
		for k, inner := range x {
			if s, ok := inner.(string); ok && k == "custom_id" && strings.HasPrefix(s, shareIDPrefix+":") {
				x[k] = shareIDPrefix + ":TOKEN"
				continue
			}
			x[k] = normalShares(inner)
		}
	case []any:
		for i := range x {
			x[i] = normalShares(x[i])
		}
	}
	return v
}

// assertMock fails unless body (a message or an edit body) is the reference payload field for
// field, Share tokens apart.
func assertMock(t *testing.T, id string, body any) {
	t.Helper()
	got := normalShares(wireOf(t, body)).(map[string]any)
	want := mockMessage(t, id)
	if !reflect.DeepEqual(got, want) {
		g, _ := json.MarshalIndent(got, "", " ")
		w, _ := json.MarshalIndent(want, "", " ")
		t.Fatalf("%s differs from the mock\ngot:\n%s\nwant:\n%s", id, g, w)
	}
}

func (w *gearWorld) bisPanel(t *testing.T, opts map[string]string) Message {
	t.Helper()
	m := w.command(t, pressDiscord, "bis", opts).Message
	assertCardLimits(t, m)
	return m
}

var gearBisOpts = map[string]string{"spec": "hunter-beast-mastery", "phase": "p1"}

func TestGearPanelMatchesTheMockPayloads(t *testing.T) {
	t.Run("bis page 1 and page 2", func(t *testing.T) {
		w := newGearWorld(t)
		p1 := w.bisPanel(t, gearBisOpts)
		assertMock(t, "bis-panel-p1", p1)
		in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": p1.Components[0][1].CustomID})
		in.Source = sourceOf(t, p1)
		p2 := w.route(t, in)
		assertCardLimits(t, p2.Message)
		// The press is a type 7 update over a V2 message: flags, components, never content.
		assertMock(t, "bis-panel-p2", p2.Message)
		if p2.Message.Components[0][0].CustomID != p1.Components[0][0].CustomID {
			t.Fatal("the page press minted another Share")
		}
	})
	t.Run("bis slot", func(t *testing.T) {
		w := newGearWorld(t)
		assertMock(t, "bis-panel-slot", w.bisPanel(t, map[string]string{"spec": "hunter-beast-mastery", "phase": "p1", "slot": "legs"}))
	})
	t.Run("char panel and no gear", func(t *testing.T) {
		w := newGearWorld(t)
		m := w.command(t, pressDiscord, "char", nil).Message
		assertCardLimits(t, m)
		assertMock(t, "char-panel", m)
		bare := gearZulmara()
		bare.Gear = []rankings.GearSlot{}
		w.cards.cards["zulmara"] = bare
		w.lg.readiness[keyZulmara] = guilds.ReadinessFacts{Word: "Not synced", OnRoster: true}
		m = w.command(t, pressDiscord, "char", nil).Message
		assertCardLimits(t, m)
		assertMock(t, "char-panel-nogear", m)
	})
	t.Run("share bis with the sheet, and its failure line", func(t *testing.T) {
		w := newGearWorld(t)
		p1 := w.bisPanel(t, gearBisOpts)
		rep := w.pressShareFrom(t, p1)
		rep.After(context.Background())
		ups := w.followUps()
		if len(ups) != 1 {
			t.Fatalf("follow-ups = %+v", ups)
		}
		assertMock(t, "share-bis", ups[0].Message)
		file := ups[0].Message.Files
		if len(file) != 1 || file[0].Name != "sheet.png" || string(file[0].Data) != "\x89PNG sheet" || file[0].ContentType != "image/png" ||
			file[0].Description != sheet.Alt(w.sheets.models[0]) {
			t.Fatalf("file = %+v", file)
		}
		if n := len(w.sheets.models[0].Slots); n != 16 || !w.sheets.models[0].TwoHand {
			t.Fatalf("the sheet holds %d cells, two-hand %v", n, w.sheets.models[0].TwoHand)
		}

		// A follow-up that fails puts the live Share back with the failure line on top.
		again := newGearWorld(t)
		again.rest.ErrFor = map[string]error{"FollowUp": &APIError{Status: 500, Message: "boom"}}
		p1 = again.bisPanel(t, gearBisOpts)
		rep = again.pressShareFrom(t, p1)
		rep.After(context.Background())
		restore := again.editsOriginal()
		if len(restore) != 1 {
			t.Fatalf("restores = %+v", restore)
		}
		m := restore[0].Message
		assertCardLimits(t, m)
		assertMock(t, "bis-panel-p1-sharefail", m)
		if n := m.componentCount(); n != maxV2Components {
			t.Fatalf("the share-failure frame holds %d components", n)
		}
	})
	t.Run("share bis when the sheet fails to render", func(t *testing.T) {
		w := newGearWorld(t)
		w.sheets.err = errors.New("no font")
		rep := w.pressShareFrom(t, w.bisPanel(t, gearBisOpts))
		rep.After(context.Background())
		ups := w.followUps()
		if len(ups) != 1 || len(ups[0].Message.Files) != 0 {
			t.Fatalf("follow-ups = %+v", ups)
		}
		assertMock(t, "share-bis-nosheet", ups[0].Message)
		if !strings.Contains(w.log.String(), "op=sheet_render") {
			t.Fatalf("no sheet_render line in %s", w.log.String())
		}
	})
	t.Run("share char with the sheet", func(t *testing.T) {
		w := newGearWorld(t)
		rep := w.pressShareFrom(t, w.command(t, pressDiscord, "char", nil).Message)
		rep.After(context.Background())
		ups := w.followUps()
		if len(ups) != 1 {
			t.Fatalf("follow-ups = %+v", ups)
		}
		assertMock(t, "share-char", ups[0].Message)
	})
}

// pressShareFrom presses the Share button of m, as that message's own press.
func (w *gearWorld) pressShareFrom(t *testing.T, m Message) Reply {
	t.Helper()
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": m.Components[0][0].CustomID})
	in.Token, in.Source = "tok-share", sourceOf(t, m)
	rep, err := w.svc.share(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	if !rep.Update || rep.After == nil {
		t.Fatalf("share = %+v", rep)
	}
	return rep
}

func (w *gearWorld) editsOriginal() []RESTCall {
	var out []RESTCall
	for _, c := range w.rest.Calls() {
		if c.Method == "EditOriginal" {
			out = append(out, c)
		}
	}
	return out
}
