package notify

import "github.com/jhunthrop/foreversixty/api/internal/auth"

// The two reasons a saved webhook is not posting (spec 10).
const (
	reasonFailedFiveTimes = "failed five times"
	reasonPlanLapsed      = "plan lapsed"
)

// blockOf is the wire shape of a saved webhook. It reads no credential: the
// Webhook it takes does not hold one.
func blockOf(w Webhook, entitled bool) *auth.DiscordBlock {
	block := &auth.DiscordBlock{
		ChannelID: w.ChannelID,
		Name:      w.Name,
		Events: auth.DiscordEvents{
			SimFinished:   w.Events[EventSimFinished],
			UploadParsed:  w.Events[EventUploadParsed],
			UploadPrivate: w.Events[EventUploadPrivate],
		},
		Failures: w.Failures,
		Source:   string(w.Source),
		GuildID:  w.GuildID,
	}
	if w.LastError != "" {
		block.LastError = &w.LastError
	}
	if reason := disabledReason(w, entitled); reason != "" {
		block.DisabledReason = &reason
	}
	return block
}

func disabledReason(w Webhook, entitled bool) string {
	switch {
	case w.DisabledAt != nil:
		return reasonFailedFiveTimes
	case !entitled:
		return reasonPlanLapsed
	}
	return ""
}
