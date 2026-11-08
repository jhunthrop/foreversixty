package events

import (
	"context"
	"net/http"
	"testing"
)

// The officer DM opt-in is off by default and changes with PATCH.
func TestOfficerDMsIsAnOptIn(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}}
	if code, env := call(t, h, http.MethodGet, "", 7); code != http.StatusOK || env.Data.OfficerDMs {
		t.Fatalf("default %d %+v", code, env.Data.OfficerDMs)
	}
	if code, env := call(t, h, http.MethodPatch, `{"officer_dms":true}`, 7); code != http.StatusOK || !env.Data.OfficerDMs {
		t.Fatalf("patch %d %+v", code, env.Data.OfficerDMs)
	}
}

func TestStoreKeepsTheOfficerDMsOptIn(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "dms")
	ctx := context.Background()
	on := true
	if err := st.SaveSettings(ctx, uid, SettingsChange{OfficerDMs: &on}); err != nil {
		t.Fatal(err)
	}
	p, err := st.Prefs(ctx, uid)
	if err != nil || !p.OfficerDMs || p.RaidReminder != DefaultRaidReminder {
		t.Fatalf("prefs %+v %v", p, err)
	}
}
