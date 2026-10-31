package livesurface

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// assertBuilt holds a builder's payload to its fixture exactly: the live-surface pushes have
// no handler between builder and wire, so the fixture is the builder's own output.
func assertBuilt(t *testing.T, name string, built any) {
	t.Helper()
	raw, err := json.Marshal(built)
	if err != nil {
		t.Fatal(err)
	}
	var got any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	wantRaw, err := os.ReadFile(contract.FixturePath(name))
	if err != nil {
		t.Fatal(err)
	}
	var want any
	if err := json.Unmarshal(wantRaw, &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s:\n got %s\nwant %s", name, raw, wantRaw)
	}
}

// liveRaid is the board's raid an hour before the pull: 38 of 40 signed up.
func liveRaid() guilds.WidgetNextRaid {
	starts := time.Unix(1796954400, 0).UTC()
	return guilds.WidgetNextRaid{
		RaidID: 7, GuildID: 2, GuildName: "OLYMPUS XXVII", ZoneName: "Onyxia's Lair", StartsAt: starts, PullAt: starts,
		Size: ptr(40), SignedUp: 38, Status: guilds.WidgetRaidOpen, My: guilds.WidgetMine{State: ptr("in"), Role: "dps"},
	}
}

func TestRaidStartPayloads(t *testing.T) {
	now := time.Unix(1796950800, 0)
	stale := time.Unix(1796955300, 0)
	attrs, state := RaidLive(liveRaid(), "Lightningbolltz")
	assertBuilt(t, "surface.live-raid-start.apns", APNsLive("start", now, RaidActivityType, attrs, state,
		&Alert{Title: "Onyxia's Lair", Body: "Invites in 60 minutes."}, &stale, nil))
	assertBuilt(t, "surface.live-raid-start.fcm", FCMData("fcm-registration-token", EventRaidStart, RaidData(attrs, state), true))
}

func TestSimUpdatePayloads(t *testing.T) {
	now := time.Unix(1796950800, 0)
	stale := time.Unix(1796951700, 0)
	state := SimState{State: SimRunning, Done: 40, Total: ptr(64), EtaSec: ptr(120)}
	assertBuilt(t, "surface.live-sim-update.apns", APNsLive("update", now, SimActivityType, nil, state, nil, &stale, nil))
	attrs := SimAttributes{SimID: "k3v9q2m8x1ab", Title: "Elemental · Top Gear"}
	assertBuilt(t, "surface.live-sim-update.fcm", FCMData("fcm-registration-token", EventSimUpdate, SimData(attrs, state), false))
}

func TestWidgetRefreshPayloads(t *testing.T) {
	assertBuilt(t, "surface.widget-refresh.apns", APNsWidgetRefresh())
	assertBuilt(t, "surface.widget-refresh.fcm", FCMData("fcm-registration-token", EventWidgetRefresh, nil, false))
}

func TestRaidLiveCarriesTheInviteTimeWhenSet(t *testing.T) {
	raid := liveRaid()
	invite := raid.StartsAt.Add(-15 * time.Minute)
	raid.InviteAt = &invite
	_, state := RaidLive(raid, "Lightningbolltz")
	if state.InviteAt == nil || *state.InviteAt != invite.Unix() {
		t.Fatalf("invite_at = %v", state.InviteAt)
	}
	if d := RaidData(RaidAttributes{}, state); d["invite_at"] != "1796953500" {
		t.Fatalf("fcm invite_at = %q", d["invite_at"])
	}
}
