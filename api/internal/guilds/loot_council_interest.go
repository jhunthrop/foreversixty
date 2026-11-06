// api/internal/guilds/loot_council_interest.go
//
// Main spec or off spec interest in a drop: a raider flags a drop of a night "ms" or "os" for
// their own character, or an officer sets it on the raider's behalf. Stored per (drop, character)
// in loot_interest. A member reads only their own flags; officers read everyone's through the
// council's candidates.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// LootInterest is one character's declared flag on one drop.
type LootInterest struct {
	ItemID       int    `json:"item_id"`
	CharacterKey string `json:"character_key"`
	MsOs         string `json:"ms_os"`
	// SetByOfficer is true when an officer set the flag on the raider's behalf.
	SetByOfficer bool `json:"set_by_officer"`
}

// interestOn is the night's flags of the guild, optionally narrowed to itemID (0 for every item)
// and to the characters userID holds (0 for everyone).
func (s *Store) interestOn(ctx context.Context, guildID int64, night awardNight, itemID int, userID int64, now time.Time) ([]LootInterest, error) {
	rows, err := s.Pool.Query(ctx, `
		select i.item_id, i.character_key, i.ms_os,
		       i.set_by is not null and not exists (
		         select 1 from guild_characters own
		         where own.guild_id = i.guild_id and own.character_key = i.character_key and own.user_id = i.set_by)
		from loot_interest i
		where i.guild_id = $1 and i.encounter_id = any($5) and ($6 = 0 or i.item_id = $6)
		  and ($7::bigint = 0 or exists (select 1 from guild_characters gc
		                                 where gc.guild_id = i.guild_id and gc.character_key = i.character_key and gc.user_id = $7))
		  and `+nightMatchSQL("i", "$2", "$3", "$4", "set_at")+`
		order by i.item_id, i.character_key`,
		append([]any{guildID}, append(night.args(now), allLoot.encounters, itemID, userID)...)...)
	if err != nil {
		return nil, fmt.Errorf("guilds: loot interest: %w", err)
	}
	defer rows.Close()
	out := []LootInterest{}
	for rows.Next() {
		var f LootInterest
		if err := rows.Scan(&f.ItemID, &f.CharacterKey, &f.MsOs, &f.SetByOfficer); err != nil {
			return nil, fmt.Errorf("guilds: loot interest: %w", err)
		}
		out = append(out, f)
	}
	return out, rows.Err()
}

// dropInterest is itemID's flags on night by character key.
func (s *Store) dropInterest(ctx context.Context, guildID int64, night awardNight, itemID int, now time.Time) (map[string]string, error) {
	flags, err := s.interestOn(ctx, guildID, night, itemID, 0, now)
	if err != nil {
		return nil, err
	}
	out := make(map[string]string, len(flags))
	for _, f := range flags {
		out[f.CharacterKey] = f.MsOs
	}
	return out, nil
}

// ownsCharacter reports whether userID holds characterKey's verified row in the guild.
func (s *Store) ownsCharacter(ctx context.Context, guildID, userID int64, characterKey string) (bool, error) {
	var own bool
	if err := s.Pool.QueryRow(ctx, `
		select exists(select 1 from guild_characters
		  where guild_id = $1 and character_key = $2 and user_id = $3 and verified_at is not null)`,
		guildID, characterKey, userID).Scan(&own); err != nil {
		return false, fmt.Errorf("guilds: loot interest: %w", err)
	}
	return own, nil
}

// SetLootInterest records characterKey's flag on itemID for night, set by setBy; msOs "" removes it.
func (s *Store) SetLootInterest(ctx context.Context, guildID int64, night awardNight, itemID int, characterKey, msOs string, setBy int64, now time.Time) error {
	if err := s.openDrop(ctx, guildID, night, itemID, characterKey, now); err != nil {
		return err
	}
	if msOs == "" {
		_, err := s.Pool.Exec(ctx, `
			delete from loot_interest
			where guild_id = $1 and coalesce(raid_id, 0) = coalesce($2::bigint, 0) and coalesce(report_id, '') = coalesce($3::text, '')
			  and encounter_id = $4 and item_id = $5 and character_key = $6`,
			guildID, night.raidID, night.reportID, dropEncounter(itemID), itemID, characterKey)
		if err != nil {
			return fmt.Errorf("guilds: loot interest: %w", err)
		}
		return nil
	}
	_, err := s.Pool.Exec(ctx, `
		insert into loot_interest (guild_id, raid_id, report_id, encounter_id, item_id, character_key, ms_os, set_by)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		on conflict (guild_id, (coalesce(raid_id, 0)), (coalesce(report_id, '')), encounter_id, item_id, character_key)
		do update set ms_os = excluded.ms_os, set_by = excluded.set_by, set_at = now()`,
		guildID, night.raidID, night.reportID, dropEncounter(itemID), itemID, characterKey, msOs, setBy)
	if err != nil {
		return fmt.Errorf("guilds: loot interest: %w", err)
	}
	return nil
}
