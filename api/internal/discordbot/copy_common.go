package discordbot

// commonCopy is the bot's sentences every command shares (experience spec
// §6.2); the tests read them back out of the spec.
var commonCopy = struct {
	// Unbound answers a command in a server with no installed guild.
	Unbound string
	// Failure answers anything that did not work.
	Failure string
}{
	Unbound: "This server is not linked to a Forever Sixty guild. An officer can add the bot from the guild's settings.",
	Failure: "Could not load that. Try again.",
}
