package discordbot

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
	Unlinked:      "Link this Discord account to Forever Sixty so sign-up buttons and commands know your characters.",
	LinkButton:    "Link Discord",
	Linked:        "This Discord account is linked to Forever Sixty. Unlink it on your account page.",
	AccountButton: "Account page",
}
