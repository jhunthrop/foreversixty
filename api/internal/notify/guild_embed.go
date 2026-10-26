package notify

import (
	"fmt"
	"time"
)

const (
	// guildTestTitle is the guild test message's title (spec 2026-10-26-guild-raids §6.2);
	// %s is the escaped guild name.
	guildTestTitle = "Forever Sixty will post raids for %s here."
	// guildRaidsAnchor is the guild page's Raids tab: members in the channel cannot open
	// Settings, so the guild posts link there.
	guildRaidsAnchor = "#raids"
	// guildFooterFormat is every guild post's footer: the guild, then Forever Sixty.
	guildFooterFormat = "%s · " + embedFooter
	// maxGuildNameRunes bounds the guild name a post carries (a WoW guild name is 24).
	maxGuildNameRunes = 40
)

// guildBaseEmbed is baseEmbed with the guild's footer.
func guildBaseEmbed(guildName, title, link string, at time.Time) Embed {
	e := baseEmbed(title, link, at)
	e.Footer = fmt.Sprintf(guildFooterFormat, truncateRunes(collapse(guildName), maxGuildNameRunes))
	return e
}

// GuildTestEmbed is the guild's "Send test message" post: one title linking to the guild
// page's Raids tab, no fields, no description. guildPath is the guild page's site path.
func GuildTestEmbed(publicBaseURL, guildName, guildPath string, now time.Time) Embed {
	name := escapedWithin(collapse(guildName), maxGuildNameRunes, maxTitleRunes-len([]rune(guildTestTitle)))
	return guildBaseEmbed(guildName, fmt.Sprintf(guildTestTitle, name),
		absoluteURL(publicBaseURL, guildPath+guildRaidsAnchor), now)
}
