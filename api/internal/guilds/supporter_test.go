package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// seedVerifiedMember seeds a verified guild character with its export and derived membership.
func seedVerifiedMember(t *testing.T, h *httpHarness, guildID, userID int64, slug, rank string) string {
	t.Helper()
	key := "us/hardcore/" + slug
	seedCharacter(t, h.pool, guildID, userID, key, rank, true)
	seedExport(t, h.pool, userID, key, "us", "hardcore", slug)
	recomputeMembership(t, h.pool, guildID, userID)
	return key
}

func (h *httpHarness) roster(guildID, viewerID int64) map[string]RosterRow {
	h.t.Helper()
	h.actor = auth.Actor{UserID: viewerID, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", guildID), "")
	var view HomeView
	h.data(res, &view)
	byKey := map[string]RosterRow{}
	for _, row := range view.Roster {
		byKey[row.CharacterKey] = row
	}
	return byKey
}

func TestRosterRowsCarryTheSupporterMark(t *testing.T) {
	h := newHTTPHarness(t)
	h.svc.Supporters = &entitlements.Store{Pool: h.pool}
	gid := seedGuild(t, h.pool, "Forever")
	leader := seedUser(t, h.pool, "supporter-leader@example.com")
	member := seedUser(t, h.pool, "supporter-member@example.com")
	leaderKey := seedVerifiedMember(t, h, gid, leader, "supleader", "leader")
	memberKey := seedVerifiedMember(t, h, gid, member, "supmember", "member")
	if _, err := h.pool.Exec(context.Background(),
		`insert into entitlements (guild_id, tier, plan, source, status) values ($1, 'raid_guild', 'guild', 'stripe', 'active')`, gid); err != nil {
		t.Fatal(err)
	}

	rows := h.roster(gid, leader)
	for _, key := range []string{leaderKey, memberKey} {
		if !rows[key].Supporter {
			t.Errorf("%s under an active guild plan: supporter = false, want true", key)
		}
	}

	if _, err := h.pool.Exec(context.Background(),
		`update entitlements set status = 'canceled' where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	rows = h.roster(gid, leader)
	for _, key := range []string{leaderKey, memberKey} {
		if rows[key].Supporter {
			t.Errorf("%s under a canceled guild plan: supporter = true, want false", key)
		}
	}
}

func TestRosterSupporterIsFalseWithoutAService(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	leader := seedUser(t, h.pool, "nosvc-leader@example.com")
	key := seedVerifiedMember(t, h, gid, leader, "nosvcleader", "leader")
	if h.roster(gid, leader)[key].Supporter {
		t.Error("a nil Supporters must read as false")
	}
}

// countingSupporters records every account it is asked about; err, when set, fails each ask.
type countingSupporters struct {
	asked []int64
	err   error
}

func (c *countingSupporters) IsSupporter(_ context.Context, userID int64) (bool, error) {
	c.asked = append(c.asked, userID)
	return c.err == nil, c.err
}

func TestRosterAsksOncePerDistinctAccount(t *testing.T) {
	h := newHTTPHarness(t)
	counter := &countingSupporters{}
	h.svc.Supporters = counter
	gid := seedGuild(t, h.pool, "Forever")
	leader := seedUser(t, h.pool, "once-leader@example.com")
	member := seedUser(t, h.pool, "once-member@example.com")
	seedVerifiedMember(t, h, gid, leader, "onceleader", "leader")
	seedVerifiedMember(t, h, gid, member, "oncemember", "member")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/oncealt", "member", true)
	seedExport(t, h.pool, member, "us/hardcore/oncealt", "us", "hardcore", "oncealt")

	rows := h.roster(gid, leader)
	if len(rows) != 3 {
		t.Fatalf("roster has %d rows, want 3", len(rows))
	}
	if len(counter.asked) != 2 {
		t.Errorf("IsSupporter asked %v, want one ask per distinct account (2)", counter.asked)
	}
}

func TestRosterSurvivesASupporterError(t *testing.T) {
	h := newHTTPHarness(t)
	h.svc.Supporters = &countingSupporters{err: errors.New("entitlements down")}
	gid := seedGuild(t, h.pool, "Forever")
	leader := seedUser(t, h.pool, "err-leader@example.com")
	key := seedVerifiedMember(t, h, gid, leader, "errleader", "leader")
	if h.roster(gid, leader)[key].Supporter {
		t.Error("an IsSupporter error must read as false")
	}
}
