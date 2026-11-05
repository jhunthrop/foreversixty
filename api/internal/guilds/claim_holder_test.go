// api/internal/guilds/claim_holder_test.go
package guilds

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// battleTag matches an account tag anywhere in a response.
var battleTag = regexp.MustCompile(`#[0-9]{4}`)

// bodyOf reads a response whole and closes it.
func bodyOf(t *testing.T, res *http.Response) string {
	t.Helper()
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

// seedRosterChar gives userID a character (with a class) on the guild's roster at a rank number.
func seedRosterChar(t *testing.T, h *httpHarness, gid, userID int64, key, name, class string, rankIndex int) {
	t.Helper()
	ctx := context.Background()
	if _, err := h.pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id) values ($1, 'us', 'hardcore', $2, $3, $4)
		 on conflict (key) do update set name = excluded.name, class = excluded.class, user_id = excluded.user_id`,
		key, name, class, userID); err != nil {
		t.Fatal(err)
	}
	seedCharacter(t, h.pool, gid, userID, key, "leader", true)
	if _, err := h.pool.Exec(ctx,
		`update guild_characters set rank_index = $3 where guild_id = $1 and character_key = $2`,
		gid, key, rankIndex); err != nil {
		t.Fatal(err)
	}
	syncMembership(t, h.pool, gid, userID)
}

func claimedGuild(t *testing.T, h *httpHarness, name string) (gid, holder int64) {
	t.Helper()
	ctx := context.Background()
	gid = seedGuild(t, h.pool, name)
	holder = seedUser(t, h.pool, name+"-holder@example.com")
	if _, err := h.pool.Exec(ctx, `update users set battletag = 'Holder#1894' where id = $1`, holder); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1, claimed_at = now() where id = $2`, holder, gid); err != nil {
		t.Fatal(err)
	}
	return gid, holder
}

func TestClaimHolderIsTheLowestRankCharacterTiesToTheMain(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid, holder := claimedGuild(t, h, "Ranks")
	seedRosterChar(t, h, gid, holder, "us/hardcore/altchar", "Altchar", "mage", 3)
	seedRosterChar(t, h, gid, holder, "us/hardcore/bchar", "Bchar", "warrior", 0)
	seedRosterChar(t, h, gid, holder, "us/hardcore/cchar", "Cchar", "rogue", 0)

	ref, err := h.store.claimHolderRef(ctx, gid, holder)
	if err != nil || ref == nil || ref.Name != "Bchar" || *ref.Class != "warrior" {
		t.Fatalf("no main: holder = %+v, %v; want the first by name at rank 0, Bchar the warrior", ref, err)
	}
	if _, err := h.pool.Exec(ctx, `update users set main_character_key = 'us/hardcore/cchar' where id = $1`, holder); err != nil {
		t.Fatal(err)
	}
	ref, err = h.store.claimHolderRef(ctx, gid, holder)
	if err != nil || ref == nil || ref.Name != "Cchar" || ref.CharacterKey != "us/hardcore/cchar" {
		t.Fatalf("with a main: holder = %+v, %v; want the main on the tie, Cchar", ref, err)
	}
}

func TestClaimHolderWithNoCharacterOnTheRosterIsNil(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid, holder := claimedGuild(t, h, "Bare")
	seedCharacter(t, h.pool, gid, holder, syntheticKey(holder), "leader", true)
	ref, err := h.store.claimHolderRef(ctx, gid, holder)
	if err != nil || ref != nil {
		t.Fatalf("holder = %+v, %v; want nil: an invite row names no character", ref, err)
	}
}

// TestGuildResponsesNeverCarryABattleTag is spec 2026-11-05 §4.3 over the wire, for the holder and
// for another officer: home, settings (read and patched) and the claim confirm.
func TestGuildResponsesNeverCarryABattleTag(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid, holder := claimedGuild(t, h, "Tagless")
	seedRosterChar(t, h, gid, holder, "us/hardcore/thornhide", "Thornhide", "druid", 0)
	officer := seedUser(t, h.pool, "tagless-officer@example.com")
	if _, err := h.pool.Exec(ctx, `update users set battletag = 'Officer#2222' where id = $1`, officer); err != nil {
		t.Fatal(err)
	}
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/taglessofficer", "officer", true)
	syncMembership(t, h.pool, gid, officer)

	for _, viewer := range []struct {
		id    int64
		yours bool
	}{{holder, true}, {officer, false}} {
		h.actor = auth.Actor{UserID: viewer.id, Role: "user", Method: "session"}
		var home HomeView
		res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
		h.data(res, &home)
		if home.Claim.ClaimedBy == nil || home.Claim.ClaimedBy.Name != "Thornhide" || home.Claim.Yours != viewer.yours {
			t.Fatalf("viewer %d home claim = %+v, want Thornhide, yours %v", viewer.id, home.Claim, viewer.yours)
		}
		for _, read := range []struct{ method, path, body string }{
			{http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), ""},
			{http.MethodGet, fmt.Sprintf("/v1/guilds/%d/settings", gid), ""},
			{http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/settings", gid), `{"default_visibility":"guild"}`},
		} {
			if body := bodyOf(t, h.do(read.method, read.path, read.body)); battleTag.MatchString(body) {
				t.Errorf("%s %s printed an account tag: %s", read.method, read.path, body)
			}
		}
		var patched SettingsView
		res = h.do(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/settings", gid), `{"default_visibility":"guild"}`)
		h.data(res, &patched)
		if patched.ClaimedBy == nil || patched.ClaimedBy.Name != "Thornhide" || patched.Claim.Yours != viewer.yours {
			t.Fatalf("a patch lost the holder: %+v", patched)
		}
	}
}

func TestConfirmClaimNamesTheHolderByCharacter(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Confirmed")
	pending := seedUser(t, h.pool, "confirmed-pending@example.com")
	if _, err := h.pool.Exec(ctx, `update users set battletag = 'Pending#3333' where id = $1`, pending); err != nil {
		t.Fatal(err)
	}
	seedRosterChar(t, h, gid, pending, "us/hardcore/pendingchar", "Pendingchar", "priest", 1)
	confirmer := seedUser(t, h.pool, "confirmed-confirmer@example.com")
	seedCharacter(t, h.pool, gid, confirmer, "us/hardcore/confirmerchar", "officer", true)
	syncMembership(t, h.pool, gid, confirmer)
	if _, err := h.pool.Exec(ctx, `update guilds set claim_pending_by = $1, claim_requested_at = now() where id = $2`, pending, gid); err != nil {
		t.Fatal(err)
	}

	// The pending claim names its claimant by character on the settings read too.
	h.actor = auth.Actor{UserID: confirmer, Role: "user", Method: "session"}
	var settings SettingsView
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/settings", gid), "")
	h.data(res, &settings)
	if settings.ClaimPending == nil || settings.ClaimPending.By == nil || settings.ClaimPending.By.Name != "Pendingchar" {
		t.Fatalf("claim_pending = %+v, want Pendingchar", settings.ClaimPending)
	}

	body := bodyOf(t, h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/claim/confirm", gid), ""))
	if battleTag.MatchString(body) || !regexp.MustCompile(`"name":"Pendingchar"`).MatchString(body) {
		t.Fatalf("confirm = %s, want the character and no account tag", body)
	}
}

func TestInviteeWithNoCharacterIsNeverNamedByAccountTag(t *testing.T) {
	if got := inviteeName(12, nil); got != "Invited member 12" {
		t.Fatalf("inviteeName = %q", got)
	}
	empty, named := "", "Thornhide"
	if got := inviteeName(12, &empty); got != "Invited member 12" {
		t.Fatalf("inviteeName(empty) = %q", got)
	}
	if got := inviteeName(12, &named); got != named {
		t.Fatalf("inviteeName(named) = %q", got)
	}
}

func TestSnapshotClassReadsTheTokenThenAnUnambiguousSpec(t *testing.T) {
	for _, tc := range []struct {
		name  string
		line  snapshotMember
		class string
	}{
		{"the addon's token wins", snapshotMember{Class: "WARRIOR", Spec: "protection"}, "warrior"},
		{"a token the catalogue does not know falls through", snapshotMember{Class: "ROBOT", Spec: "arms"}, "warrior"},
		{"a specs.json key", snapshotMember{Spec: "paladin-holy"}, "paladin"},
		{"a bare name one class plays", snapshotMember{Spec: "Fury"}, "warrior"},
		{"a slug one class plays", snapshotMember{Spec: "beast-mastery"}, "hunter"},
		{"survival", snapshotMember{Spec: "survival"}, "hunter"},
		{"protection is two classes", snapshotMember{Spec: "protection"}, ""},
		{"holy is two classes", snapshotMember{Spec: "Holy"}, ""},
		{"restoration is two classes", snapshotMember{Spec: "restoration"}, ""},
		{"the addon's no-spec dash", snapshotMember{Spec: "-"}, ""},
		{"nothing", snapshotMember{}, ""},
	} {
		got := snapshotClass(tc.line)
		switch {
		case tc.class == "" && got != nil:
			t.Errorf("%s: class = %q, want unknown", tc.name, *got)
		case tc.class != "" && (got == nil || *got != tc.class):
			t.Errorf("%s: class = %v, want %q", tc.name, got, tc.class)
		}
	}
}
