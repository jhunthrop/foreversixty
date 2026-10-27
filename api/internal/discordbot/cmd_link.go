package discordbot

import (
	"context"
	"fmt"
)

// accountPagePath is where a linked member manages the link.
const accountPagePath = "/account#you"

// linkButton is the link button that starts the account link.
func (s *Service) linkButton() Component {
	return Component{Kind: KindButton, Style: StyleLink, Label: linkCopy.LinkButton, URL: s.linkStartURL()}
}

func (s *Service) accountButton() Component {
	return Component{Kind: KindButton, Style: StyleLink, Label: linkCopy.AccountButton, URL: s.d.PublicBaseURL + accountPagePath}
}

// cmdLink answers /link ephemerally: how to link, or where to manage the link. It reads only
// whether the caller's own Discord account is linked.
func (s *Service) cmdLink(ctx context.Context, in Interaction) (Reply, error) {
	_, linked, err := s.viewer(ctx, in)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: link command: %w", err)
	}
	msg := Message{Content: linkCopy.Unlinked, Ephemeral: true, Components: []ActionRow{{s.linkButton()}}}
	if linked {
		msg = Message{Content: linkCopy.Linked, Ephemeral: true, Components: []ActionRow{{s.accountButton()}}}
	}
	return Reply{Message: msg}, nil
}
