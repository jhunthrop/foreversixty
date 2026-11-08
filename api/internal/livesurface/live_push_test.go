package livesurface

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/push"
)

// fakeLiveAPNs records the payloads; invalid tokens answer ErrInvalidToken.
type fakeLiveAPNs struct {
	sent    []map[string]any
	tokens  []string
	invalid map[string]bool
}

func (f *fakeLiveAPNs) SendLive(_ context.Context, token string, payload map[string]any) error {
	if f.invalid[token] {
		return push.ErrInvalidToken
	}
	f.sent, f.tokens = append(f.sent, payload), append(f.tokens, token)
	return nil
}

type fakeLiveFCM struct{ sent []map[string]any }

func (f *fakeLiveFCM) SendBody(_ context.Context, body map[string]any) error {
	f.sent = append(f.sent, body)
	return nil
}

// fakeLiveRaids answers one raid's attendees, its widget raid and its raid night.
type fakeLiveRaids struct {
	attendees []guilds.RaidAttendee
	next      guilds.WidgetNextRaid
	night     guilds.RaidNightSummary
}

func (f *fakeLiveRaids) RaidAttendees(context.Context, int64) ([]guilds.RaidAttendee, error) {
	return f.attendees, nil
}
func (f *fakeLiveRaids) WidgetRaidFor(context.Context, int64, string, time.Time) (guilds.WidgetRaid, error) {
	n := f.next
	return guilds.WidgetRaid{NextRaid: &n}, nil
}
func (f *fakeLiveRaids) RaidNightOf(context.Context, int64, int64, time.Time) (guilds.RaidNightSummary, bool, error) {
	return f.night, true, nil
}

func aps(t *testing.T, payload map[string]any) map[string]any {
	t.Helper()
	a, ok := payload["aps"].(map[string]any)
	if !ok {
		t.Fatalf("payload %+v", payload)
	}
	return a
}

// raid.starting starts the countdown on an In account's start token, once: an account with
// the countdown already running gets nothing, and a refused token goes.
func TestRaidStartingPushesToStart(t *testing.T) {
	h := newHarness(t).withStore()
	ctx := context.Background()
	if err := h.svc.Tokens.SaveStartToken(ctx, h.actor.UserID, "ph", StartTokenRequest{Platform: "ios", Kind: KindRaid, Token: iosLiveToken}); err != nil {
		t.Fatal(err)
	}
	next := liveRaid()
	next.StartsAt, next.EndsAt, next.ZoneIcon, next.In, next.Healers = time.Now().Add(time.Hour), time.Now().Add(4*time.Hour), "inv_misc_head_dragon_black", 30, 5
	apns := &fakeLiveAPNs{}
	p := &LivePusher{APNs: apns, Tokens: h.svc.Tokens, Raids: &fakeLiveRaids{
		attendees: []guilds.RaidAttendee{{UserID: h.actor.UserID, CharacterKey: "us/pvp/x", Name: "Lightningbolltz"}}, next: next}}
	p.RaidStarting(ctx, 7)
	if len(apns.sent) != 1 || apns.tokens[0] != iosLiveToken {
		t.Fatalf("starts %+v", apns.sent)
	}
	a := aps(t, apns.sent[0])
	attrs := a["attributes"].(RaidAttributes)
	state := a["content-state"].(RaidState)
	if a["event"] != "start" || a["attributes-type"] != RaidActivityType || attrs.ZoneIcon != "inv_misc_head_dragon_black" ||
		state.In != 30 || state.Healers != 5 || a["alert"] == nil {
		t.Fatalf("start %+v", a)
	}
	// With the countdown registered, raid.starting sends nothing more.
	h.raids.raidEnd[7] = testNow.Add(50 * time.Hour)
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities", raidBody(7, iosLiveToken))
	h.want(res, raw, http.StatusCreated)
	p.RaidStarting(ctx, 7)
	if len(apns.sent) != 1 {
		t.Fatalf("a second start %+v", apns.sent)
	}
}

// The raid's end goes to every kept countdown: done, the night's kills, the raider's loot and
// whether the report is ready; Android's ongoing notification gets the same as data.
func TestRaidEndedPushesTheNight(t *testing.T) {
	h := newHarness(t).withStore()
	h.raids.raidEnd[7] = testNow.Add(50 * time.Hour)
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities", raidBody(7, iosLiveToken))
	h.want(res, raw, http.StatusCreated)
	kills, report := 1, "ewlrx3enxbeb"
	night := guilds.RaidNightSummary{RaidID: 7, GuildID: 2, GuildName: "OLYMPUS XXVII", ZoneName: "Onyxia's Lair",
		StartsAt: time.Unix(1796954400, 0), Kills: &kills, ReportID: &report,
		Loot: []guilds.RaidNightLoot{{Name: "Sapphiron Drape", Quality: 4, Icon: "inv_misc_cape_16"}}}
	apns := &fakeLiveAPNs{}
	p := &LivePusher{APNs: apns, Tokens: h.svc.Tokens, Raids: &fakeLiveRaids{night: night}}
	p.RaidEnded(context.Background(), 7)
	if len(apns.sent) != 1 {
		t.Fatalf("ends %+v", apns.sent)
	}
	a := aps(t, apns.sent[0])
	state := a["content-state"].(RaidState)
	if a["event"] != "end" || a["dismissal-date"] == nil || state.Status != RaidStatusDone || state.Kills == nil ||
		*state.Kills != 1 || len(state.Loot) != 1 || state.Loot[0].Name != "Sapphiron Drape" || !*state.ReportReady {
		t.Fatalf("end %+v", a)
	}
}

func TestStoreAStartToken(t *testing.T) {
	h := newHarness(t).withStore()
	res, raw := h.do(http.MethodPost, "/v1/me/live-activities/start-token", `{"platform":"ios","kind":"raid","token":"`+iosLiveToken+`"}`)
	contract.AssertFixture(t, "live.start-token", h.want(res, raw, http.StatusOK), contract.Options{})
	toks, err := h.svc.Tokens.StartTokens(context.Background(), h.actor.UserID, KindRaid)
	if err != nil || len(toks) != 1 || toks[0].Token != iosLiveToken {
		t.Fatalf("tokens %+v %v", toks, err)
	}
	res, raw = h.do(http.MethodPost, "/v1/me/live-activities/start-token", `{"platform":"ios","kind":"sim","token":"zz"}`)
	h.want(res, raw, http.StatusUnprocessableEntity)
}
