package linkedrole

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

func newSyncer() (*Syncer, *memStore, *fakeDiscord) {
	store, discord := newMemStore(), &fakeDiscord{revoked: map[string]bool{}}
	return &Syncer{Store: store, Sealer: plainSealer{}, Discord: discord}, store, discord
}

func TestLinkSealsTheRefreshTokenAndPushesWithTheBattletag(t *testing.T) {
	s, store, discord := newSyncer()
	store.profiles[7] = Profile{Premium: true, Battletag: "hunthrop#1894"}
	if err := s.Link(context.Background(), 7, "d7", Tokens{Access: "a0", Refresh: "r0"}); err != nil {
		t.Fatal(err)
	}
	if got := string(store.links[7].RefreshBox); got != "sealed:r0" {
		t.Fatalf("stored box = %q, want the sealed refresh token", got)
	}
	want := push{Access: "a0", Conn: Connection{Premium: true, Username: "hunthrop#1894"}}
	if len(discord.pushes) != 1 || discord.pushes[0] != want {
		t.Fatalf("pushes = %+v, want %+v", discord.pushes, want)
	}
	if !store.pushed[7] {
		t.Fatal("the push was not recorded")
	}
}

func TestAnAnonymousAccountPushesNoUsername(t *testing.T) {
	s, store, discord := newSyncer()
	store.profiles[7] = Profile{Battletag: "hunthrop#1894", Anonymize: true}
	if err := s.Link(context.Background(), 7, "d7", Tokens{Access: "a0", Refresh: "r0"}); err != nil {
		t.Fatal(err)
	}
	if got := discord.pushes[0].Conn.Username; got != "" {
		t.Fatalf("username = %q, want none for an anonymous account", got)
	}
}

func TestRelinkingAnotherDiscordAccountClearsTheOldOne(t *testing.T) {
	s, store, discord := newSyncer()
	store.profiles[7] = Profile{Premium: true}
	store.links[7] = Link{UserID: 7, DiscordUserID: "old", RefreshBox: []byte("sealed:rOld")}
	if err := s.Link(context.Background(), 7, "new", Tokens{Access: "aNew", Refresh: "rNew"}); err != nil {
		t.Fatal(err)
	}
	if len(discord.pushes) != 2 || discord.pushes[0].Conn.Premium || discord.pushes[0].Access == "aNew" {
		t.Fatalf("pushes = %+v, want premium false to the old account first", discord.pushes)
	}
	if store.links[7].DiscordUserID != "new" {
		t.Fatal("the link did not move to the new Discord account")
	}
}

func TestSyncRotatesTheRefreshTokenAndPushesThePremiumNow(t *testing.T) {
	s, store, discord := newSyncer()
	store.links[7] = Link{UserID: 7, DiscordUserID: "d7", RefreshBox: []byte("sealed:r0")}
	store.profiles[7] = Profile{Premium: false}
	if err := s.SyncUser(context.Background(), 7); err != nil {
		t.Fatal(err)
	}
	if got := string(store.links[7].RefreshBox); got != "sealed:r0+1" {
		t.Fatalf("box = %q, want the rotated token sealed", got)
	}
	if len(discord.pushes) != 1 || discord.pushes[0].Conn.Premium {
		t.Fatalf("pushes = %+v, want one premium false", discord.pushes)
	}
}

func TestARevokedTokenRemovesTheLinkQuietly(t *testing.T) {
	s, store, discord := newSyncer()
	store.links[7] = Link{UserID: 7, DiscordUserID: "d7", RefreshBox: []byte("sealed:r0")}
	discord.revoked["r0"] = true
	if err := s.SyncUser(context.Background(), 7); err != nil {
		t.Fatalf("a revoked token is not an error: %v", err)
	}
	if _, ok := store.links[7]; ok {
		t.Fatal("the revoked link is still stored")
	}
	if len(discord.pushes) != 0 {
		t.Fatal("pushed without a token")
	}
}

func TestSyncOfAnUnlinkedAccountIsANoOp(t *testing.T) {
	s, _, discord := newSyncer()
	if err := s.SyncUser(context.Background(), 9); err != nil || len(discord.pushes) != 0 {
		t.Fatalf("err = %v, pushes = %d", err, len(discord.pushes))
	}
}

func TestAGuildChangeSyncsItsLinkedMembersOnly(t *testing.T) {
	s, store, discord := newSyncer()
	store.links[1] = Link{UserID: 1, DiscordUserID: "d1", RefreshBox: []byte("sealed:r1")}
	store.links[2] = Link{UserID: 2, DiscordUserID: "d2", RefreshBox: []byte("sealed:r2")}
	store.members[50] = []int64{1, 3}
	guild := int64(50)
	s.OnEntitlementChange(context.Background(), entitlements.Subject{GuildID: &guild})
	if len(discord.pushes) != 1 {
		t.Fatalf("pushes = %d, want only the guild's linked member", len(discord.pushes))
	}
}

func TestReconcileCountsAndFailsOnlyWhenEveryAccountFailed(t *testing.T) {
	s, store, _ := newSyncer()
	store.links[1] = Link{UserID: 1, DiscordUserID: "d1", RefreshBox: []byte("sealed:r1")}
	store.links[2] = Link{UserID: 2, DiscordUserID: "d2", RefreshBox: []byte("sealed:r2")}
	r, err := s.Reconcile(context.Background())
	if err != nil || r.Synced != 2 || r.Failed != 0 {
		t.Fatalf("result = %+v, err = %v", r, err)
	}
}

func TestRevokeClearsTheRoleRevokesTheRotatedTokenAndDeletesTheLink(t *testing.T) {
	s, store, discord := newSyncer()
	store.links[7] = Link{UserID: 7, DiscordUserID: "d7", RefreshBox: []byte("sealed:r0")}
	store.profiles[7] = Profile{Premium: true, Battletag: "hunthrop#1894"}
	if err := s.Revoke(context.Background(), 7); err != nil {
		t.Fatal(err)
	}
	want := push{Access: "a-r0-1", Conn: Connection{Premium: false}}
	if len(discord.pushes) != 1 || discord.pushes[0] != want {
		t.Fatalf("pushes = %+v, want %+v (premium false, no username)", discord.pushes, want)
	}
	if len(discord.revokedTokens) != 1 || discord.revokedTokens[0] != "r0+1" {
		t.Fatalf("revoked = %v, want the rotated refresh token", discord.revokedTokens)
	}
	if _, ok := store.links[7]; ok {
		t.Fatal("the link is still stored")
	}
}

func TestRevokeWithNoLinkDoesNothing(t *testing.T) {
	s, _, discord := newSyncer()
	if err := s.Revoke(context.Background(), 7); err != nil {
		t.Fatal(err)
	}
	if len(discord.pushes) != 0 || len(discord.revokedTokens) != 0 {
		t.Fatalf("pushes %v, revoked %v; want no Discord call", discord.pushes, discord.revokedTokens)
	}
}

func TestRevokeOfATokenDiscordAlreadyRevokedJustRemovesTheLink(t *testing.T) {
	s, store, discord := newSyncer()
	store.links[7] = Link{UserID: 7, DiscordUserID: "d7", RefreshBox: []byte("sealed:r0")}
	discord.revoked["r0"] = true
	if err := s.Revoke(context.Background(), 7); err != nil {
		t.Fatal(err)
	}
	if _, ok := store.links[7]; ok || len(discord.revokedTokens) != 0 {
		t.Fatalf("link kept %v, revoked %v; want the link removed and no revoke call", ok, discord.revokedTokens)
	}
}
