package discordbot

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
) // Accept and Decline on the recruiting channel's application post (guild onboarding spec
// 2026-11-09 §2.3): a linked verified officer decides without leaving Discord, through the
// site's own Decide (the recruit role, the applicant's DM and the votes all follow). The answer
// is ephemeral, to the officer who pressed.

const (
	prefixApplication = "gapp"
	applicationAccept = "accept"
	applicationReject = "decline"
)

// applicationCopy is the presses' words.
var applicationCopy = struct {
	Accept, Decline, Review, Character, Unlinked, OfficersOnly, Accepted, Declined, Closed, Failed string
}{
	Accept:       copygen.DiscordApplicationAccept,
	Decline:      copygen.PushActionDecline,
	Review:       copygen.DiscordApplicationReview,
	Character:    copygen.DiscordApplicationCharacter,
	Unlinked:     copygen.DiscordApplicationUnlinked,
	OfficersOnly: copygen.DiscordApplicationOfficersOnly,
	Accepted:     copygen.DiscordApplicationAccepted,
	Declined:     copygen.DiscordApplicationDeclined,
	Closed:       copygen.DiscordApplicationClosed,
	Failed:       copygen.FirstPickMainFailed,
}

// applicationButtons are the post's row: Accept, Decline, the review link and the character's
// page (its gear and logs).
func applicationButtons(id int64, reviewURL, characterURL string) ActionRow {
	idText := strconv.FormatInt(id, 10)
	return ActionRow{
		{Kind: KindButton, Style: StyleSuccess, Label: applicationCopy.Accept,
			CustomID: strings.Join([]string{prefixApplication, applicationAccept, idText}, idSeparator)},
		{Kind: KindButton, Style: StyleSecondary, Label: applicationCopy.Decline,
			CustomID: strings.Join([]string{prefixApplication, applicationReject, idText}, idSeparator)},
		{Kind: KindButton, Style: StyleLink, Label: applicationCopy.Review, URL: reviewURL},
		{Kind: KindButton, Style: StyleLink, Label: applicationCopy.Character, URL: characterURL},
	}
}

// parseApplicationID reads gapp:accept:<id> or gapp:decline:<id>.
func parseApplicationID(customID string) (id int64, accept, ok bool) {
	parts := strings.Split(customID, idSeparator)
	if len(parts) != 3 || parts[0] != prefixApplication || (parts[1] != applicationAccept && parts[1] != applicationReject) {
		return 0, false, false
	}
	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil || id <= 0 {
		return 0, false, false
	}
	return id, parts[1] == applicationAccept, true
}

// applicationPress decides the application for a linked verified officer of the server's guild.
func (s *Service) applicationPress(ctx context.Context, in Interaction) (Reply, error) {
	id, accept, ok := parseApplicationID(in.CustomID())
	if !ok {
		return pressReply(applicationCopy.Failed, false), nil
	}
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil {
		return Reply{}, err
	}
	if !bound {
		return pressReply(commonCopy.Unbound, false), nil
	}
	if viewer == nil {
		return pressReply(applicationCopy.Unlinked, false, s.linkButton()), nil
	}
	standing, found, err := s.d.Guilds.GuildStanding(ctx, guildID, *viewer)
	if err != nil {
		return Reply{}, err
	}
	if !found || !standing.Officer {
		return pressReply(applicationCopy.OfficersOnly, false), nil
	}
	if standing.Frozen {
		return pressReply(officerCopy.Frozen, false), nil
	}
	d, err := s.d.Guilds.Decide(ctx, guildID, *viewer, id, accept)
	switch {
	case errors.Is(err, guilds.ErrNotFound), errors.Is(err, guilds.ErrApplicationClosed):
		return pressReply(applicationCopy.Closed, false), nil
	case err != nil:
		return Reply{}, err
	}
	word := applicationCopy.Declined
	if accept {
		word = applicationCopy.Accepted
	}
	return pressReply(fill(word, "{Name}", memberText(d.Application.Character.Name)), false), nil
}
