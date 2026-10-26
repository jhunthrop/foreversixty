// api/internal/guilds/raids_roster.go
//
// The roster the raid views and the sign-up route read: every character row of the guild
// with the snapshot (name, class, spec) a sign-up stores, and who owns it.
package guilds

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// raidRosterChar is one guild_characters row with the snapshot a sign-up stores.
type raidRosterChar struct {
	RaidCharacter
	userID   int64
	verified bool
	role     string // the newest logged fight's role, "" when none
}

// defaultRole is the role a sign-up takes when the request names none: the newest logged
// fight's role, else the spec's.
func (c raidRosterChar) defaultRole() string {
	if c.role == RoleTank || c.role == RoleHealer || c.role == RoleDPS {
		return c.role
	}
	return roleForSpec(c.Spec)
}

// raidRoster reads every character row of the guild (invite placeholders excluded) with its
// name (characters, else the export, else the key's last part), its class (newest logged
// fight, else characters) and its spec and role (newest logged fight). A non-empty key
// narrows the read to that one row.
func (s *Store) raidRoster(ctx context.Context, guildID int64, key string) ([]raidRosterChar, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.character_key, gc.user_id, gc.verified_at is not null,
		       coalesce(nullif(c.name, ''), nullif(ae.name, ''), split_part(gc.character_key, '/', 3)),
		       lower(coalesce(fm.class, c.class, '')), coalesce(fm.spec, ''), coalesce(fm.role, '')
		from guild_characters gc
		left join characters c on c.key = gc.character_key
		left join addon_exports ae on ae.character_key = gc.character_key
		left join lateral (
		  select m.class, m.spec, m.role from fight_metrics m join reports r on r.id = m.report_id
		  where m.player_key = gc.character_key and `+presence.EvidenceReport("r")+`
		  order by m.fought_at desc limit 1
		) fm on true
		where gc.guild_id = $1 and gc.character_key not like 'account:%'
		  and ($2 = '' or gc.character_key = $2)`, guildID, key)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid roster: %w", err)
	}
	defer rows.Close()
	out := []raidRosterChar{}
	for rows.Next() {
		var c raidRosterChar
		if err := rows.Scan(&c.CharacterKey, &c.userID, &c.verified, &c.Name, &c.Class, &c.Spec, &c.role); err != nil {
			return nil, fmt.Errorf("guilds: raid roster: %w", err)
		}
		out = append(out, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: raid roster: %w", err)
	}
	return out, nil
}

// raidViewContext is what every view of one guild's raids shares.
type raidViewContext struct {
	viewerID int64
	timezone *string
	roster   []raidRosterChar
	ownerOf  map[string]int64
}

func (s *Store) loadRaidViewContext(ctx context.Context, guildID, viewerID int64) (raidViewContext, error) {
	tz, err := s.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return raidViewContext{}, err
	}
	roster, err := s.raidRoster(ctx, guildID, "")
	if err != nil {
		return raidViewContext{}, err
	}
	vc := raidViewContext{viewerID: viewerID, roster: roster, ownerOf: make(map[string]int64, len(roster))}
	if tz != "" {
		vc.timezone = &tz
	}
	for _, c := range roster {
		vc.ownerOf[c.CharacterKey] = c.userID
	}
	return vc, nil
}

// owns reports whether the viewer's account holds sg's character: its roster owner, else (a
// character since removed from the roster) the account the sign-up was written for.
func (vc raidViewContext) owns(sg Signup) bool {
	if owner, ok := vc.ownerOf[sg.CharacterKey]; ok {
		return owner == vc.viewerID
	}
	return sg.SignedBy == vc.viewerID
}
