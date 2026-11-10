package discordbot

import (
	"context"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// helpCommand is /help's own name, left out of its list.
const helpCommand = "help"

// cmdHelp answers /help: one accent embed, a line per command, no components. It reads nothing
// of the caller, so it answers the same in a DM or an unbound server.
func (s *Service) cmdHelp(_ context.Context, _ Interaction) (Reply, error) {
	var lines []string
	for _, c := range s.commandList() {
		if c.Name == helpCommand { // the player just ran it
			continue
		}
		lines = append(lines, fill(commandCopy.HelpLine, "{name}", c.Name, "{description}", c.Description))
	}
	embed := notify.WithSealFooter(notify.BrandEmbed(commandCopy.HelpTitle, strings.Join(lines, "\n"), s.now()), s.sealURL())
	return Reply{Message: Message{Embeds: []notify.Embed{embed}, Ephemeral: true}}, nil
}
