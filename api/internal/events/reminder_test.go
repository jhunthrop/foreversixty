package events

import (
	"context"
	"net/http"
	"testing"
)

// Each event row names its group, so the apps build one switch per group; the flows spec's two
// new kinds come back in theirs, off by default.
func TestGetNamesEachEventsGroup(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}}
	code, env := call(t, h, http.MethodGet, "", 7)
	if code != http.StatusOK {
		t.Fatalf("get = %d", code)
	}
	for _, c := range []struct {
		kind  Kind
		group Group
		on    bool
	}{
		{RaidPosted, GroupRaids, true}, {GuildSimPassed, GroupLadder, false}, {TierMoved, GroupTiers, false},
		{RaidShort, GroupOfficer, true},
	} {
		v := eventView(env.Data, c.kind)
		if v.Group != c.group || v.Channels[ChannelPush] != c.on {
			t.Errorf("%s: %+v", c.kind, v)
		}
	}
	if env.Data.RaidReminder != "15:00" {
		t.Fatalf("raid reminder = %q", env.Data.RaidReminder)
	}
}

func TestPatchSetsTheRaidReminder(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}}
	code, env := call(t, h, http.MethodPatch, `{"raid_reminder":"17:30"}`, 7)
	if code != http.StatusOK || env.Data.RaidReminder != "17:30" {
		t.Fatalf("patch = %d %+v", code, env.Data)
	}
	for _, bad := range []string{`"24:00"`, `"17:20"`, `"5pm"`, `""`, `"7:30"`} {
		code, env := call(t, h, http.MethodPatch, `{"raid_reminder":`+bad+`}`, 7)
		if code != http.StatusUnprocessableEntity || env.Error == nil || env.Error.Fields["raid_reminder"] == "" {
			t.Errorf("%s = %d %+v", bad, code, env.Error)
		}
	}
}

// The store keeps the reminder per account, the default until one is saved.
func TestStoreKeepsTheRaidReminder(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "reminder")
	ctx := context.Background()
	p, err := st.Prefs(ctx, uid)
	if err != nil || p.RaidReminder != DefaultRaidReminder {
		t.Fatalf("default %+v %v", p, err)
	}
	at := 18*60 + 45
	if err := st.SaveSettings(ctx, uid, SettingsChange{RaidReminder: &at}); err != nil {
		t.Fatal(err)
	}
	if p, err = st.Prefs(ctx, uid); err != nil || p.RaidReminder != at || p.TimeZone != DefaultTimeZone {
		t.Fatalf("saved %+v %v", p, err)
	}
}
