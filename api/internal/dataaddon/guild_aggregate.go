// api/internal/dataaddon/guild_aggregate.go
package dataaddon

import (
	"fmt"
	"sort"
	"strings"
)

// guildRow is one guild's line in Data.lua.
type guildRow struct {
	Name     string
	Progress string // "" when there is nothing to report (omitted from Data.lua)
	Nights   int
	Roster   int
	Members  []string // lowercased character-name slugs, sorted
}

// guildIdentity names one guild with at least one verified member -- the
// only kind this job publishes a row for.
type guildIdentity struct {
	ID                    int64
	Region, Ruleset, Name string
}

// memberNameOf reads the name segment back out of a character key
// ("us/normal/thoradin" -> "thoradin") -- guild_characters.character_key
// is written in api/internal/character.Key's own format
// (<region>/<ruleset>/<name-slug>), and this is that format's inverse for
// the one segment this package needs. An unparseable key (should not
// happen: character_key is not-null and every write path uses
// character.Key) is returned unchanged rather than panicking, so a bad row
// degrades to an odd-looking member name instead of stopping the run.
func memberNameOf(key string) string {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 {
		return key
	}
	return parts[2]
}

// buildGuildRow assembles one guild's Data.lua row. memberKeys are that
// guild's verified guild_characters.character_key values; killed/total
// are attempted-encounter counts (see store.go's progressionByGuild).
func buildGuildRow(g guildIdentity, memberKeys []string, nights, killed, total int) guildRow {
	members := make([]string, 0, len(memberKeys))
	for _, key := range memberKeys {
		members = append(members, slug(memberNameOf(key)))
	}
	sort.Strings(members)
	progress := ""
	if total > 0 {
		progress = fmt.Sprintf("%d/%d", killed, total)
	}
	return guildRow{
		Name: g.Name, Progress: progress, Nights: nights,
		Roster: len(members), Members: members,
	}
}
