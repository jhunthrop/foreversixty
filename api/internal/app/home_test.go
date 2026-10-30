package app

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

func (w *homeWorld) get(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, w.svc)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func (w *homeWorld) home(t *testing.T, name string) Home {
	t.Helper()
	got, found, err := w.svc.HomeOf(context.Background(), "us", "normal", name)
	if err != nil || !found {
		t.Fatalf("found %v err %v", found, err)
	}
	return got
}

// The answer for the parity character: the site's upgrades (the parity test pins them) in the
// bot's order, each with its words, and the facts around them.
func TestTheHomeAnswer(t *testing.T) {
	w := newHomeWorld(t)
	h := w.home(t, "Velaren")
	if h.State != StateOK || deref(h.Spec) != "hunter-beast-mastery" || *h.Band != 60 || *h.BandLabel != "60" ||
		*h.Preset != "raid" || *h.Unit != "dps" || *h.UnitWord != "DPS" || h.DataBuild != homeBuild {
		t.Fatalf("facts = %+v", h)
	}
	var order, words []string
	for _, u := range h.Upgrades {
		order = append(order, u.Slot)
		words = append(words, u.Slot+"="+u.GainWords)
	}
	if strings.Join(order, ",") != "hands,trinket1,ranged,chest,neck,legs,head,wrist,feet,finger1,main_hand" {
		t.Fatalf("order = %v", order)
	}
	for _, want := range []string{"hands=+11.1 DPS", "trinket1=Gain not modelled yet", "ranged=+23.1 DPS", "legs=−0.8 DPS",
		"head=Gain not modelled yet", "wrist=Gain not modelled yet", "main_hand=not sim-checked"} {
		if !strings.Contains(strings.Join(words, "|"), want) {
			t.Errorf("words lack %q: %v", want, words)
		}
	}
	if h.UpgradeCount != 11 || *h.TotalGain != 76.2 || h.NotSimCheckedCount != 1 || h.NotModelledCount != 6 {
		t.Fatalf("count %d total %v unchecked %d unmodelled %d", h.UpgradeCount, *h.TotalGain, h.NotSimCheckedCount, h.NotModelledCount)
	}
	if h.SyncedAt == nil || !h.SyncedAt.Equal(homeSynced) || deref(h.SyncedFrom) != "addon" {
		t.Fatalf("synced %v from %v", h.SyncedAt, h.SyncedFrom)
	}
	if h.Talents.State != TalentsDiffer || *h.Talents.PointsDiffer != 1 || *h.Talents.BandPoints != 51 {
		t.Fatalf("talents = %+v", h.Talents)
	}
	var best []string
	for _, b := range h.AlreadyBest {
		best = append(best, b.Slot)
	}
	if strings.Join(best, ",") != "shoulder,back,waist,finger2,trinket2" || h.AlreadyBest[0].SlotLabel != "Shoulder" ||
		deref(h.AlreadyBest[0].Item.Name) != "Field Marshal's Chain Spaulders" {
		t.Fatalf("already best = %+v", h.AlreadyBest)
	}
}

// Each row carries the item and icon ids the apps draw from, and a worn item the item data
// lacks says so.
func TestTheRowsCarryTheItemFacts(t *testing.T) {
	h := newHomeWorld(t).home(t, "velaren")
	rows := map[string]Upgrade{}
	for _, u := range h.Upgrades {
		rows[u.Slot] = u
	}
	chest := rows["chest"]
	if chest.SlotLabel != "Chest" || chest.Pick.ID != 252483 || deref(chest.Pick.Name) != "Dawn Armor" || chest.Pick.Icon == nil ||
		chest.Pick.Quality == nil || chest.Pick.ItemLevel == nil || chest.Pick.SourceLine == "" || !chest.Modelled {
		t.Fatalf("chest pick = %+v", chest.Pick)
	}
	if chest.Worn == nil || chest.Worn.ID != 16721 || !chest.Worn.Known || chest.Worn.Label != "Shadowcraft Tunic" || chest.Worn.Icon == nil {
		t.Fatalf("chest worn = %+v", chest.Worn)
	}
	wrist := rows["wrist"]
	if wrist.Worn == nil || wrist.Worn.Known || wrist.Worn.Label != WornNotInData || wrist.Worn.Name != nil || wrist.Gain != nil || wrist.Modelled {
		t.Fatalf("wrist = %+v %+v", wrist, wrist.Worn)
	}
	if rows["hands"].Worn != nil || *rows["hands"].Gain != 11.1 {
		t.Fatalf("hands = %+v", rows["hands"])
	}
	if head := rows["head"]; head.Gain == nil || head.Modelled {
		t.Fatalf("an unmodelled pick keeps its figure but is not modelled: %+v", head)
	}
}

func TestTheHomeStates(t *testing.T) {
	for name, tc := range map[string]struct {
		edit    func(*homeWorld, *rankings.Character)
		state   string
		talents string
	}{
		"no gear":   {func(_ *homeWorld, c *rankings.Character) { c.Gear = []rankings.GearSlot{} }, StateNoGear, TalentsDiffer},
		"under 20":  {func(_ *homeWorld, c *rankings.Character) { c.Character.Level = ptr(19) }, StateUnder20, TalentsNoBuild},
		"no spec":   {func(_ *homeWorld, c *rankings.Character) { c.GearSpec = nil }, StateNoSpec, TalentsNoBuild},
		"no list":   {func(_ *homeWorld, c *rankings.Character) { c.GearSpec = ptr("warrior-fury") }, StateNoList, TalentsNoBuild},
		"no export": {func(w *homeWorld, _ *rankings.Character) { delete(w.exports.rows, homeKey) }, StateOK, TalentsNoExport},
		"every slot matches": {func(w *homeWorld, c *rankings.Character) {
			*c = velaren(map[string]int{"head": 226887, "neck": 19491, "shoulder": 272105, "back": 13340, "chest": 252483,
				"wrist": 226885, "hands": 272099, "waist": 20216, "legs": 237819, "feet": 226881, "finger1": 234202,
				"finger2": 272408, "trinket1": 249469, "trinket2": 249470, "main_hand": 250619, "ranged": 279273})
		}, StateOK, TalentsDiffer},
	} {
		t.Run(name, func(t *testing.T) {
			w := newHomeWorld(t)
			card := w.chars.cards["velaren"]
			tc.edit(w, &card)
			w.chars.cards["velaren"] = card
			h := w.home(t, "velaren")
			if h.State != tc.state || h.Talents.State != tc.talents {
				t.Fatalf("state %q talents %q", h.State, h.Talents.State)
			}
			if h.State != StateOK && (len(h.Upgrades) != 0 || len(h.AlreadyBest) != 0 || h.TotalGain != nil) {
				t.Fatalf("a state with no answer listed rows: %+v", h)
			}
			if h.Upgrades == nil || h.AlreadyBest == nil {
				t.Fatal("arrays are never null")
			}
		})
	}
}

func TestTheHomeRoute(t *testing.T) {
	w := newHomeWorld(t)
	rec := w.get(t, "/v1/characters/US/Normal/Velaren/home")
	if rec.Code != http.StatusOK || rec.Header().Get("Cache-Control") != "public, max-age=30, stale-while-revalidate=300" {
		t.Fatalf("%d %q", rec.Code, rec.Header().Get("Cache-Control"))
	}
	var body struct {
		OK   bool `json:"ok"`
		Data Home `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil || !body.OK || body.Data.Character != homeKey {
		t.Fatalf("body = %s", rec.Body)
	}
	// The character page's own visibility: an unknown or anonymised character (the card's
	// found=false) and an unknown region are a 404, never a partial answer.
	for _, path := range []string{"/v1/characters/us/normal/nobody/home", "/v1/characters/xx/normal/velaren/home"} {
		if rec := w.get(t, path); rec.Code != http.StatusNotFound || rec.Header().Get("Cache-Control") != "" {
			t.Errorf("%s = %d %q", path, rec.Code, rec.Header().Get("Cache-Control"))
		}
	}
	w.exports.err = errBoom
	if rec := w.get(t, "/v1/characters/us/normal/velaren/home"); rec.Code != http.StatusInternalServerError || strings.Contains(rec.Body.String(), "boom") {
		t.Fatalf("a failed read = %d %s", rec.Code, rec.Body)
	}
}

func ptr[T any](v T) *T { return &v }

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
