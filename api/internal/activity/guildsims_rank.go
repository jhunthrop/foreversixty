package activity

import (
	"context"
	"time"
)

// GuildSimsRank is characterKey's place on guildID's sims ladder, the one Home's Guild sims block
// reads (its role, its level's band, the preset of its own latest sim there): position of of.
// ok is false for a character with no spec, a ladder under the minimum, a ladder of another
// guild, or no sim on it. The guilds package's raid badge reads it (guilds.SimsRanker).
func (h *Home) GuildSimsRank(ctx context.Context, guildID int64, characterKey string) (position, of int, ok bool, err error) {
	c, found, err := h.character(ctx, characterKey)
	if err != nil || !found || c.spec == "" {
		return 0, 0, false, err
	}
	g, err := h.Store.guildSims(ctx, c, BandOf(c.level), h.now())
	if err != nil || g == nil || g.GuildID != guildID || g.You == nil {
		return 0, 0, false, err
	}
	return g.You.Rank, g.Of, true, nil
}

func (h *Home) now() time.Time {
	if h.Now != nil {
		return h.Now()
	}
	return time.Now()
}
