// api/internal/guilds/invite_shown_test.go
package guilds

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// testSealer is a real AES-256-GCM sealer under a fresh random key.
func testSealer(t *testing.T) *notify.Sealer {
	t.Helper()
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	sealer, err := notify.NewSealer(key)
	if err != nil {
		t.Fatal(err)
	}
	return sealer
}

// sealedHarness is the HTTP harness with an invite sealer wired, a guild, and its leader
// (Leadchar, a warrior on the roster) signed in.
func sealedHarness(t *testing.T) (h *httpHarness, gid, leader int64) {
	t.Helper()
	h = newHTTPHarness(t)
	h.store.InviteSealer = testSealer(t)
	gid = seedGuild(t, h.pool, "Forever")
	leader = seedUser(t, h.pool, "invite-leader@example.com")
	seedRosterChar(t, h, gid, leader, "us/hardcore/leadchar", "Leadchar", "warrior", 0)
	h.actor = auth.Actor{UserID: leader, Role: "user", Method: "session"}
	return h, gid, leader
}

type settingsInvite struct {
	Invite struct {
		RotatedAt *string         `json:"rotated_at"`
		URL       *string         `json:"url"`
		MadeBy    *ClaimHolderRef `json:"made_by"`
	} `json:"invite"`
}

func (h *httpHarness) settingsInvite(gid int64) settingsInvite {
	h.t.Helper()
	var got settingsInvite
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/settings", gid), ""), &got)
	return got
}

func (h *httpHarness) rotate(gid int64) (token string) {
	h.t.Helper()
	var got struct {
		Token string `json:"token"`
	}
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/invite/rotate", gid), ""), &got)
	return got.Token
}

// The request log redacts both routes that carry the token in their path.
func TestInviteRoutesAreRedactedInTheLogs(t *testing.T) {
	for _, p := range []string{"/v1/guild-invites/fsg_secret", "/v1/guilds/invite/fsg_secret/accept"} {
		if got := httpx.LogPath(p); strings.Contains(got, "fsg_secret") {
			t.Errorf("LogPath(%q) keeps the token: %q", p, got)
		}
	}
}

// The token is stored sealed, never in the clear, and opens back to itself; the hash stays the
// lookup key, so the sealed copy and the hash name the same token.
func TestRotateInviteSealsTheTokenAndInviteURLOpensIt(t *testing.T) {
	h, gid, leader := sealedHarness(t)
	ctx := context.Background()
	token, _, err := h.store.RotateInvite(ctx, gid, &leader)
	if err != nil {
		t.Fatal(err)
	}
	var box []byte
	var rotatedBy int64
	if err := h.pool.QueryRow(ctx, `select invite_token_enc, invite_rotated_by from guilds where id = $1`, gid).
		Scan(&box, &rotatedBy); err != nil {
		t.Fatal(err)
	}
	if len(box) == 0 || bytes.Contains(box, []byte(token)) || rotatedBy != leader {
		t.Fatalf("stored box %d bytes (holds the token in the clear: %v), rotated_by %d; want a sealed box by %d",
			len(box), bytes.Contains(box, []byte(token)), rotatedBy, leader)
	}
	url, err := h.store.InviteURL(ctx, gid)
	if err != nil || url == nil || *url != "/guild/invite/"+token {
		t.Fatalf("InviteURL = %v, %v; want the path of %s", url, err, token)
	}
}

// Rotating replaces the token, its sealed copy and its maker: the settings read shows the new
// link, and the old one no longer opens anything.
func TestRotatingTheInviteReplacesTheShownLink(t *testing.T) {
	h, gid, _ := sealedHarness(t)
	first := h.rotate(gid)
	if got := h.settingsInvite(gid).Invite; got.URL == nil || *got.URL != "/guild/invite/"+first {
		t.Fatalf("settings invite url = %v, want the first link", got.URL)
	}
	second := h.rotate(gid)
	if second == first {
		t.Fatal("rotation answered the same token")
	}
	got := h.settingsInvite(gid).Invite
	if got.URL == nil || *got.URL != "/guild/invite/"+second {
		t.Fatalf("settings invite url after rotation = %v, want the second link", got.URL)
	}
	if got.MadeBy == nil || got.MadeBy.Name != "Leadchar" || got.MadeBy.Class == nil || *got.MadeBy.Class != "warrior" {
		t.Fatalf("made_by = %+v, want Leadchar the warrior", got.MadeBy)
	}
	res := h.do(http.MethodGet, "/v1/guild-invites/"+first, "")
	if bodyOf(t, res); res.StatusCode != http.StatusNotFound {
		t.Fatalf("peek of the rotated-away link = %d, want 404", res.StatusCode)
	}
}

// Only a verified officer or leader reads the link back. A moderator reads settings without it,
// and a plain member reads no settings at all.
func TestOnlyAVerifiedOfficerReadsTheInviteLinkBack(t *testing.T) {
	h, gid, _ := sealedHarness(t)
	token := h.rotate(gid)

	officer := seedUser(t, h.pool, "invite-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/inviteofficer", "officer", true)
	syncMembership(t, h.pool, gid, officer)
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	if got := h.settingsInvite(gid).Invite.URL; got == nil || *got != "/guild/invite/"+token {
		t.Fatalf("officer reads url %v, want the link", got)
	}

	moderator := seedUser(t, h.pool, "invite-mod@example.com")
	h.actor = auth.Actor{UserID: moderator, Role: "moderator", Method: "session"}
	raw := bodyOf(t, h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/settings", gid), ""))
	if strings.Contains(raw, token) || !strings.Contains(raw, `"url":null`) {
		t.Fatalf("moderator settings carry the token or no null url: %s", raw)
	}

	member := seedUser(t, h.pool, "invite-member@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/invitemember", "member", true)
	syncMembership(t, h.pool, gid, member)
	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/settings", gid), "")
	if raw := bodyOf(t, res); res.StatusCode != http.StatusForbidden || strings.Contains(raw, token) {
		t.Fatalf("member settings = %d %s, want 403 without the token", res.StatusCode, raw)
	}
}

// Without a sealer nothing but the hash is kept, so the link cannot be shown again; a kept box
// whose hash has moved (a build without 0070 rotated it) is not shown; a box sealed under another
// key does not open and reads the same way.
func TestAnInviteLinkThatCannotBeShownAgainReadsAsNone(t *testing.T) {
	h, gid, leader := sealedHarness(t)
	ctx := context.Background()

	h.store.InviteSealer = nil
	h.rotate(gid)
	if got := h.settingsInvite(gid).Invite; got.URL != nil || got.RotatedAt == nil {
		t.Fatalf("no sealer: url %v rotated_at %v, want no url and the date", got.URL, got.RotatedAt)
	}

	h.store.InviteSealer = testSealer(t)
	if _, _, err := h.store.RotateInvite(ctx, gid, &leader); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx, `update guilds set invite_token_hash = $2 where id = $1`,
		gid, auth.TokenHash("fsg_rotatedbyanolderbuild")); err != nil {
		t.Fatal(err)
	}
	if got := h.settingsInvite(gid).Invite.URL; got != nil {
		t.Fatalf("stale box: url %v, want none", *got)
	}

	h.rotate(gid)
	h.store.InviteSealer = testSealer(t) // the key changed under the stored box
	if got := h.settingsInvite(gid).Invite.URL; got != nil {
		t.Fatalf("foreign key: url %v, want none", *got)
	}
}

// The rotate answer carries the settings shape, so the page knows whether the link it just made
// can be shown again and who made it.
func TestRotateAnswersTheInviteView(t *testing.T) {
	h, gid, _ := sealedHarness(t)
	var got struct {
		Token  string     `json:"token"`
		Invite InviteView `json:"invite"`
	}
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/invite/rotate", gid), ""), &got)
	if got.Invite.URL == nil || *got.Invite.URL != "/guild/invite/"+got.Token || got.Invite.RotatedAt == nil ||
		got.Invite.MadeBy == nil || got.Invite.MadeBy.Name != "Leadchar" {
		t.Fatalf("rotate invite = %+v, want the kept url, the date and Leadchar", got.Invite)
	}
}

// The join page names the officer who made the link by their character on the roster, signed in
// or not; a link made by a tool, or by an account with no character there, names nobody.
func TestPeekInviteNamesTheInviter(t *testing.T) {
	h, gid, leader := sealedHarness(t)
	ctx := context.Background()
	token := h.rotate(gid)

	peek := func() *ClaimHolderRef {
		t.Helper()
		var got struct {
			Inviter *ClaimHolderRef `json:"inviter"`
		}
		h.data(h.do(http.MethodGet, "/v1/guild-invites/"+token, ""), &got)
		return got.Inviter
	}
	h.actor = auth.Actor{}
	if got := peek(); got == nil || got.Name != "Leadchar" || got.Class == nil || *got.Class != "warrior" {
		t.Fatalf("inviter = %+v, want Leadchar the warrior", got)
	}

	if _, err := h.pool.Exec(ctx, `delete from guild_characters where guild_id = $1 and user_id = $2`, gid, leader); err != nil {
		t.Fatal(err)
	}
	if got := peek(); got != nil {
		t.Fatalf("inviter off the roster = %+v, want none", got)
	}

	var err error
	if token, _, err = h.store.RotateInvite(ctx, gid, nil); err != nil {
		t.Fatal(err)
	}
	if got := peek(); got != nil {
		t.Fatalf("tool-made link inviter = %+v, want none", got)
	}
}
