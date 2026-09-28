// api/internal/guilds/character_state_test.go
package guilds

import (
	"context"
	"testing"
)

func TestCharacterGuildStateNoGuild(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	_, ok, err := s.CharacterGuildState(context.Background(), "us/hardcore/nobody")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("expected no guild for a character with no guild_characters row")
	}
}

func TestCharacterGuildStateForAMember(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	userID := seedUser(t, pool, "member@example.com")
	guildID := seedGuild(t, pool, "Sanguine")
	key := "us/hardcore/memberalt"
	seedCharacter(t, pool, guildID, userID, key, "member", true)

	state, ok, err := s.CharacterGuildState(context.Background(), key)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected a guild")
	}
	if state.Name != "Sanguine" || state.Rank != "member" {
		t.Fatalf("state = %+v", state)
	}
	if state.Claim.State != "unclaimed" {
		t.Fatalf("claim state = %q", state.Claim.State)
	}
	// A plain member never sees the approval queue's size, even when
	// the guild has unverified rows waiting.
	if state.PendingApprovals != 0 {
		t.Fatalf("pending approvals leaked to a member: %d", state.PendingApprovals)
	}
}

func TestCharacterGuildStateCountsPendingApprovalsForAnOfficer(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	guildID := seedGuild(t, pool, "Iron Vanguard")

	officerID := seedUser(t, pool, "officer@example.com")
	officerKey := "us/hardcore/officeralt"
	seedCharacter(t, pool, guildID, officerID, officerKey, "officer", true)

	// Two unverified characters waiting on approval, and one already
	// verified one that should not be counted.
	unverifiedA := seedUser(t, pool, "waiting-a@example.com")
	seedCharacter(t, pool, guildID, unverifiedA, "us/hardcore/waitinga", "member", false)
	unverifiedB := seedUser(t, pool, "waiting-b@example.com")
	seedCharacter(t, pool, guildID, unverifiedB, "us/hardcore/waitingb", "member", false)
	verified := seedUser(t, pool, "settled@example.com")
	seedCharacter(t, pool, guildID, verified, "us/hardcore/settled", "member", true)

	state, ok, err := s.CharacterGuildState(context.Background(), officerKey)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected a guild")
	}
	if state.Rank != "officer" || state.PendingApprovals != 2 {
		t.Fatalf("state = %+v", state)
	}
}
