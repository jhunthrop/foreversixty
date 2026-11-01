package livesurface

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

const (
	boltzKey = "us/pvp/lightningbolltz"
	yellKey  = "us/pvp/yellfury"
)

func ptr[T any](v T) *T { return &v }

// boardWorld is the F14 board's player: Lightningbolltz, In for Onyxia, three upgrades of eight.
func boardWorld(h *harness) {
	h.accounts.chars = []auth.Character{
		{Key: boltzKey, Name: "Lightningbolltz", Class: "shaman", Spec: "Elemental"},
		{Key: yellKey, Name: "Yellfury", Class: "hunter"},
	}
	starts := time.Date(2026, 12, 11, 2, 0, 0, 0, time.UTC)
	h.raids.byKey[boltzKey] = guilds.WidgetRaid{
		Guild: &guilds.WidgetGuild{ID: 2, Name: "OLYMPUS XXVII"},
		NextRaid: &guilds.WidgetNextRaid{
			RaidID: 7, GuildID: 2, GuildName: "OLYMPUS XXVII", Zone: "raid:onyxias-lair", ZoneName: "Onyxia's Lair",
			StartsAt: starts, PullAt: starts, EndsAt: starts.Add(3 * time.Hour), Size: ptr(40), SignedUp: 24,
			Status: guilds.WidgetRaidOpen, CanSignUp: true,
			My: guilds.WidgetMine{State: ptr("in"), Spec: ptr("Elemental"), Role: "dps"},
		},
		LastRaid: &guilds.WidgetLastRaid{ZoneName: "Onyxia's Lair", StartsAt: starts.Add(-7 * 24 * time.Hour)},
	}
	rows := []UpgradeRow{
		{Slot: "waist", SlotLabel: "Waist", Pick: UpgradePick{ID: 228190, Name: "Knowledge of the Timbermaw", Quality: 4, Icon: "inv_belt_09", ItemLevel: 65}, Gain: ptr(5.3)},
		{Slot: "wrist", SlotLabel: "Wrist", Pick: UpgradePick{ID: 19597, Name: "Dryad's Wrist Bindings", Quality: 4, Icon: "inv_bracer_13", ItemLevel: 45}, Gain: ptr(4.0)},
		{Slot: "back", SlotLabel: "Back", Pick: UpgradePick{ID: 272411, Name: "Arcanoweave Cloak", Quality: 3, Icon: "inv_misc_cape_05", ItemLevel: 65}, Gain: ptr(3.1)},
		{Slot: "head", SlotLabel: "Head", Pick: UpgradePick{ID: 1, Name: "Fourth", Quality: 3, Icon: "x", ItemLevel: 60}, Gain: nil},
	}
	h.svc.Upgrades = fakeUpgrades{
		boltzKey: {State: "ok", Count: 8, TotalGain: ptr(18.3), Unit: "dps", Rows: rows},
		yellKey:  {State: "no_gear", Sentence: ptr("Add the addon to see your upgrades."), Unit: "dps", Rows: []UpgradeRow{}},
	}
}

func (h *harness) widget(query string) Widget {
	h.t.Helper()
	res, raw := h.do(http.MethodGet, "/v1/me/widget"+query, "")
	h.want(res, raw, http.StatusOK)
	var env struct{ Data Widget }
	if err := json.Unmarshal(raw, &env); err != nil {
		h.t.Fatal(err)
	}
	return env.Data
}

func TestWidgetPicksTheAskedCharacterElseTheMainElseTheFirst(t *testing.T) {
	h := newHarness(t)
	boardWorld(h)
	if got := h.widget("?character=" + yellKey).Character.Key; got != yellKey {
		t.Fatalf("asked yellfury, got %s", got)
	}
	if got := h.widget("?character=us/pvp/someone-else").Character.Key; got != boltzKey {
		t.Fatalf("unknown key read %s, want the first character", got)
	}
	h.accounts.user.MainCharacterKey = ptr(yellKey)
	if got := h.widget("").Character.Key; got != yellKey {
		t.Fatalf("no key read %s, want the main", got)
	}
	if h.raids.asked != yellKey {
		t.Fatalf("the raid read was for %s", h.raids.asked)
	}
}

func TestWidgetCarriesTheFirstThreeUpgradesAndTheWholeTotal(t *testing.T) {
	h := newHarness(t)
	boardWorld(h)
	h.sims.running = 1
	w := h.widget("")
	if w.Upgrades == nil || len(w.Upgrades.Rows) != widgetRows || w.Upgrades.Count != 8 || *w.Upgrades.TotalGain != 18.3 {
		t.Fatalf("upgrades = %+v", w.Upgrades)
	}
	if w.Upgrades.Rows[0].Pick.Name != "Knowledge of the Timbermaw" {
		t.Fatalf("order changed: %+v", w.Upgrades.Rows)
	}
	if w.Live.SimRunsRunning != 1 || !w.GeneratedAt.Equal(testNow.Truncate(time.Minute)) {
		t.Fatalf("live = %+v, generated_at = %v", w.Live, w.GeneratedAt)
	}
}

func TestWidgetWithoutCharactersOrHomeAnswer(t *testing.T) {
	h := newHarness(t)
	w := h.widget("")
	if w.Character != nil || w.Guild != nil || w.NextRaid != nil || w.Upgrades != nil {
		t.Fatalf("empty account widget = %+v", w)
	}
	boardWorld(h)
	h.svc.Upgrades = nil
	if w := h.widget(""); w.Upgrades != nil || w.NextRaid == nil {
		t.Fatalf("without the Home answer: %+v", w)
	}
}

func TestWidgetIsPrivateAndRevalidates(t *testing.T) {
	h := newHarness(t)
	boardWorld(h)
	res, raw := h.do(http.MethodGet, "/v1/me/widget", "")
	h.want(res, raw, http.StatusOK)
	if got := res.Header.Get("Cache-Control"); got != "private, no-cache" {
		t.Fatalf("Cache-Control = %q", got)
	}
	etag := res.Header.Get("ETag")
	req, _ := http.NewRequest(http.MethodGet, h.server.URL+"/v1/me/widget", nil)
	req.Header.Set("If-None-Match", etag)
	again, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	_ = again.Body.Close()
	if etag == "" || again.StatusCode != http.StatusNotModified {
		t.Fatalf("ETag %q answered %d on revalidation, want 304", etag, again.StatusCode)
	}
}

func TestWidgetNeedsASignedInCaller(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{}
	res, raw := h.do(http.MethodGet, "/v1/me/widget", "")
	contract.AssertFixture(t, "error-unauthorized", h.want(res, raw, http.StatusUnauthorized), contract.Options{})
}

func TestAppContractWidget(t *testing.T) {
	h := newHarness(t)
	boardWorld(h)
	res, raw := h.do(http.MethodGet, "/v1/me/widget", "")
	contract.AssertFixture(t, "widget.get", h.want(res, raw, http.StatusOK), contract.Options{})
	res, raw = h.do(http.MethodGet, "/v1/me/widget?character="+yellKey, "")
	contract.AssertFixture(t, "widget.get.no-guild", h.want(res, raw, http.StatusOK), contract.Options{})
}
