package activity

import (
	"context"
	"testing"
	"time"
)

// The raid badge's sims_rank is the place Home's Guild sims block shows: the character's on its
// guild's ladder for its role, band and preset, none below the ladder's minimum or in another
// guild.
func TestGuildSimsRank(t *testing.T) {
	s := testStore(t)
	g, _ := guildOf(t, s, 5)
	h := &Home{Store: s, Now: func() time.Time { return testNow }}
	ctx := context.Background()
	pos, of, ok, err := h.GuildSimsRank(ctx, g, "us/normal/mc")
	if err != nil || !ok || pos != 3 || of != 5 {
		t.Fatalf("rank %d of %d (%v, %v)", pos, of, ok, err)
	}
	if _, _, ok, err := h.GuildSimsRank(ctx, g+1, "us/normal/mc"); err != nil || ok {
		t.Fatalf("another guild: %v %v", ok, err)
	}
	if _, _, ok, err := h.GuildSimsRank(ctx, g, "us/normal/nobody"); err != nil || ok {
		t.Fatalf("unknown character: %v %v", ok, err)
	}
}
