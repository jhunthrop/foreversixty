package discordbot

import "github.com/jhunthrop/foreversixty/api/internal/copygen"

// linkCopy is /link's text (experience spec §6.2); the tests read it back out of the spec.
var linkCopy = struct {
	// Unlinked answers a member whose Discord account is not linked.
	Unlinked string
	// LinkButton starts the link.
	LinkButton string
	// Linked answers a member whose Discord account is linked.
	Linked string
	// AccountButton opens the account page, where the link can be removed.
	AccountButton string
}{
	Unlinked:      copygen.DiscordLinkUnlinked,
	LinkButton:    copygen.DiscordLinkLinkButton,
	Linked:        copygen.DiscordLinkLinked,
	AccountButton: copygen.DiscordLinkAccountButton,
}
