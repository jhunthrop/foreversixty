package discordbot

import (
	"context"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// cmdHelp answers /help: one gold embed, a line per command, no components. It reads nothing
// of the caller, so it answers the same in a DM or an unbound server.
func (s *Service) cmdHelp(_ context.Context, _ Interaction) (Reply, error) {
	cmds := s.commandList()
	lines := make([]string, len(cmds))
	for i, c := range cmds {
		lines[i] = fill(commandCopy.HelpLine, "{name}", c.Name, "{description}", c.Description)
	}
	embed := notify.BrandEmbed(commandCopy.HelpTitle, strings.Join(lines, "\n"), s.now())
	return Reply{Message: Message{Embeds: []notify.Embed{embed}, Ephemeral: true}}, nil
}
