// api/cmd/seedguild/guild_centre_integration_test.go
//
// The guild control centre's own integration test, built on this tool's own realistic
// fixture: apply the seed against a fresh owner-guild (seedOwnerGuild, the exact production
// shape integration_test.go's own TestApplyThenRemoveRoundTrips already exercises), then
// hit every one of the contract's endpoints (docs/contracts/2026-10-04-guild-centre-api.md)
// through a real mux as a public visitor, a verified member, a verified officer and the
// guild's own moderator-flagged account, asserting each answers the shape the contract
// promises rather than an error.
//
// Needs TEST_DATABASE_URL (api/docker-compose.test.yml); skips otherwise, matching every
// other package's own harness.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// centreHarness mounts guilds.Mount on a real HTTP server, with DataDir pointed at this
// repo's own data/builds (api/internal/bis's band reads, exactly as api/cmd/api/main.go
// wires TreeDataDir in production) and Accounts backed by a real auth.Store so
// claimed_by_name/loot award attribution resolve real battletags.
type centreHarness struct {
	t      *testing.T
	store  *guilds.Store
	actor  auth.Actor
	server *httptest.Server
}

func newCentreHarnessFor(t *testing.T, pool *pgxpool.Pool) *centreHarness {
	t.Helper()
	root, err := repoRoot()
	if err != nil {
		t.Fatal(err)
	}
	treeData, err := trees.Load(filepath.Join(root, "data", "builds"))
	if err != nil {
		t.Fatal(err)
	}
	accounts := &auth.Store{Pool: pool}
	store := &guilds.Store{Pool: pool, Accounts: accounts, DataDir: filepath.Join(root, "data", "builds"), Trees: treeData}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	svc := &guilds.Service{Store: store, Accounts: accounts, Log: quiet}
	mux := http.NewServeMux()
	guilds.Mount(mux, svc, 0)
	h := &centreHarness{t: t, store: store}
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return h
}

func (h *centreHarness) get(path string) *http.Response {
	h.t.Helper()
	res, err := http.Get(h.server.URL + path)
	if err != nil {
		h.t.Fatal(err)
	}
	return res
}

func (h *centreHarness) post(method, path, body string) *http.Response {
	h.t.Helper()
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, h.server.URL+path, reader)
	if err != nil {
		h.t.Fatal(err)
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		h.t.Fatal(err)
	}
	return res
}

func (h *centreHarness) data(res *http.Response, into any) {
	h.t.Helper()
	defer res.Body.Close()
	var env struct {
		OK    bool                            `json:"ok"`
		Data  json.RawMessage                 `json:"data"`
		Error *struct{ Code, Message string } `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		h.t.Fatal(err)
	}
	if !env.OK {
		h.t.Fatalf("envelope failure (status %d): %+v", res.StatusCode, env.Error)
	}
	if into != nil {
		if err := json.Unmarshal(env.Data, into); err != nil {
			h.t.Fatal(err)
		}
	}
}

func TestGuildCentreAcrossEveryViewerRole(t *testing.T) {
	ctx := context.Background()
	pool := testPool(t)
	guildID, ownerID, ownerKey := seedOwnerGuild(t, ctx, pool)
	if err := ApplySeed(ctx, pool, guildID, "hunthrop#1894"); err != nil {
		t.Fatalf("ApplySeed: %v", err)
	}
	// The second officer the mock roster seeds (roster.go's own Shiverknife row) - used
	// below as the officer viewer.
	var officerID int64
	if err := pool.QueryRow(ctx,
		`select user_id from guild_characters where guild_id = $1 and rank = 'officer' and character_key <> $2 limit 1`,
		guildID, ownerKey).Scan(&officerID); err != nil {
		t.Fatalf("find a seeded officer: %v", err)
	}
	// A plain verified member - any non-officer, non-leader row.
	var memberID int64
	if err := pool.QueryRow(ctx,
		`select user_id from guild_characters where guild_id = $1 and rank = 'member' and verified_at is not null limit 1`,
		guildID).Scan(&memberID); err != nil {
		t.Fatalf("find a seeded member: %v", err)
	}
	var moderator int64
	if err := pool.QueryRow(ctx,
		`insert into users (battletag, role) values ('centremod#1111', 'moderator') returning id`,
	).Scan(&moderator); err != nil {
		t.Fatal(err)
	}

	h := newCentreHarnessFor(t, pool)

	// --- public / signed-out ---
	h.actor = auth.Actor{}
	res := h.get(fmt.Sprintf("/v1/guilds/%d/raids", guildID))
	var raids guilds.RaidsPage
	h.data(res, &raids)

	res = h.get(fmt.Sprintf("/v1/guilds/%d/progression", guildID))
	var prog guilds.ProgressionView
	h.data(res, &prog)
	if len(prog.Encounters) != 1 || prog.Encounters[0].Name != "Onyxia" {
		t.Errorf("progression.encounters = %+v, want exactly Onyxia", prog.Encounters)
	}
	if prog.Tier.Down != 1 {
		t.Errorf("tier.down = %d, want 1 (the seed's own two Onyxia kills)", prog.Tier.Down)
	}

	res = h.get(fmt.Sprintf("/v1/guilds/%d/readiness", guildID))
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("signed-out readiness = %d, want 401", res.StatusCode)
	}

	// --- verified member ---
	h.actor = auth.Actor{UserID: memberID, Role: "user", Method: "session"}
	res = h.get(fmt.Sprintf("/v1/guilds/%d/home", guildID))
	var home guilds.HomeView
	h.data(res, &home)
	if len(home.Roster) != 25 {
		t.Errorf("roster = %d rows, want 25 (owner + 24 mock)", len(home.Roster))
	}
	if len(home.Pending) != 0 {
		t.Errorf("a plain member's pending = %+v, want empty", home.Pending)
	}
	if home.Viewer.Role != "member" {
		t.Errorf("viewer.role = %q, want member", home.Viewer.Role)
	}

	res = h.get(fmt.Sprintf("/v1/guilds/%d/readiness", guildID))
	var readiness guilds.ReadinessView
	h.data(res, &readiness)
	if len(readiness.Rows) == 0 {
		t.Error("readiness.rows is empty, want one row per verified character")
	}
	for _, row := range readiness.Rows {
		if row.NudgeText != nil {
			t.Errorf("member view row %s carries nudge_text, want omitted", row.CharacterKey)
		}
	}

	res = h.get(fmt.Sprintf("/v1/guilds/%d/loot", guildID))
	var loot guilds.LootView
	h.data(res, &loot)
	if len(loot.Items) != 22 {
		t.Errorf("loot.items = %d, want 22", len(loot.Items))
	}
	// Follow-up fix: loot candidates must not read empty on real data. Every real
	// armor/weapon drop (a non-empty slot) must carry at least one candidate from this
	// 24-mock-character roster, via the fallback tier if no BiS band happens to name it
	// (api/internal/guilds/CONTROL_CENTRE.md: no BiS file in this repo names any
	// raid-tier item today, so every real candidate on this roster comes from the
	// fallback tier until raid-tier BiS/sim data exists).
	for _, item := range loot.Items {
		if item.Slot == "" {
			continue // a quest/reputation/crafting drop - no class can "wear" it
		}
		if len(item.Candidates) == 0 {
			t.Errorf("item %d (%s, slot %s) has no candidates at all on a 24-character mock roster",
				item.ItemID, item.Name, item.Slot)
		}
	}

	// --- verified officer ---
	h.actor = auth.Actor{UserID: officerID, Role: "user", Method: "session"}
	res = h.get(fmt.Sprintf("/v1/guilds/%d/home", guildID))
	h.data(res, &home)
	if len(home.Pending) != 3 {
		t.Errorf("officer pending = %d, want 3 (the seed's own unverified trio)", len(home.Pending))
	}
	if home.Viewer.Role != "officer" {
		t.Errorf("viewer.role = %q, want officer", home.Viewer.Role)
	}

	res = h.get(fmt.Sprintf("/v1/guilds/%d/readiness", guildID))
	h.data(res, &readiness)
	for _, row := range readiness.Rows {
		if row.NudgeText == nil {
			t.Errorf("officer view row %s missing nudge_text", row.CharacterKey)
		}
	}

	// --- owner/leader: claimed_by_name resolves ---
	h.actor = auth.Actor{UserID: ownerID, Role: "user", Method: "session"}
	res = h.get(fmt.Sprintf("/v1/guilds/%d/home", guildID))
	h.data(res, &home)
	if home.Claim.State != "claimed" || home.Claim.ClaimedByName == nil {
		t.Errorf("claim = %+v, want claimed with a resolved name", home.Claim)
	}

	// --- moderator: readiness/loot still need their own membership, but the account-wide
	// role check elsewhere in this package (moderationClaims) is exercised by guilds' own
	// test suite - here it is enough that a moderator with no character in this guild is
	// refused the member-only surfaces exactly like any other non-member.
	h.actor = auth.Actor{UserID: moderator, Role: "moderator", Method: "session"}
	res = h.get(fmt.Sprintf("/v1/guilds/%d/readiness", guildID))
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("a moderator with no character here reading readiness = %d, want 403", res.StatusCode)
	}

	if err := RemoveSeed(ctx, pool, guildID, "hunthrop#1894"); err != nil {
		t.Fatalf("RemoveSeed: %v", err)
	}
}
