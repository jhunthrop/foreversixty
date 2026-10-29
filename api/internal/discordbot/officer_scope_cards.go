package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// partScopeRaid is the select on a /readiness or /report answer that changes its raid.
const partScopeRaid = "raid"

// raidLabel is `{Title} · {Ddd d Mmm}` in the guild's zone; the title is plain text in a
// select, so it is collapsed, not escaped, and cut to fit the label.
func raidLabel(title string, starts time.Time, loc *time.Location) string {
	date := starts.In(loc).Format(raidChoiceDate)
	room := maxOptionRunes - len([]rune(fill(guideCopy.SignupRaidLabel, "{Title}", "", "{Ddd d Mmm}", date)))
	return fill(guideCopy.SignupRaidLabel, "{Title}", truncateRunes(notify.Collapse(title), room), "{Ddd d Mmm}", date)
}

// scopeSelect is the one select under a /readiness or /report answer.
func scopeSelect(prefix, placeholder string, options []SelectOption) ActionRow {
	return ActionRow{Component{
		Kind: KindStringSelect, CustomID: cardCustomID(prefix, partScopeRaid),
		Placeholder: placeholder, Options: options,
	}}
}

// readinessRaidRow is the "Another raid" select, present with two or more planned raids.
func readinessRaidRow(raids []guilds.Raid, shown int64, loc *time.Location) []ActionRow {
	if len(raids) < 2 {
		return nil
	}
	options := make([]SelectOption, 0, min(len(raids), maxSelectOptions))
	for _, r := range raids[:min(len(raids), maxSelectOptions)] {
		options = append(options, SelectOption{
			Label: raidLabel(r.Title, r.StartsAt, loc), Value: strconv.FormatInt(r.ID, 10),
			Description: r.StartsAt.In(loc).Format(signupGuildTimeLayout), Default: r.ID == shown,
		})
	}
	return []ActionRow{scopeSelect(prefixReadiness, guideCopy.ReadinessPlaceholder, options)}
}

// readinessPress answers a pick on the /readiness select: that raid's board, with rank, plan
// and freeze checked again.
func (s *Service) readinessPress() Handler {
	return s.cardHandler("/readiness", map[string]cardPart{
		partScopeRaid: func(ctx context.Context, in Interaction, _ cardPress) (Reply, error) {
			guildID, viewer, bound, err := s.lookupCaller(ctx, in)
			if err != nil {
				return Reply{}, err
			}
			if !bound {
				return pressReply(commonCopy.Unbound, true), nil
			}
			return s.readinessReply(ctx, guildID, viewer, picked(in), true)
		},
	})
}

// autocompleteReadiness suggests /readiness's raid: the planned raids for an officer, nothing
// for anyone else.
func (s *Service) autocompleteReadiness(ctx context.Context, in Interaction) ([]Choice, error) {
	if focused, _ := in.Focused(); focused != optReadinessRaid {
		return []Choice{}, nil
	}
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound || viewer == nil {
		return []Choice{}, err
	}
	officer, err := s.d.Guilds.IsOfficer(ctx, guildID, *viewer)
	if err != nil || !officer {
		return []Choice{}, err
	}
	_, typed := in.Focused()
	return s.raidChoices(ctx, guildID, typed)
}

// reportNightRow is the "Another night" select under a report: the guild's last done raids the
// viewer may open, newest first, present with two or more. Each description is that night's
// line, read as the viewer.
func (s *Service) reportNightRow(ctx context.Context, guildID, viewer, shown int64) ([]ActionRow, error) {
	member, err := s.verifiedMember(ctx, guildID, &viewer)
	if err != nil {
		return nil, err
	}
	raids, err := s.d.Guilds.DoneRaids(ctx, guildID, !member, reportChoices)
	if err != nil || len(raids) < 2 {
		return nil, err
	}
	zone, err := s.d.Guilds.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("discordbot: report nights zone: %w", err)
	}
	options := make([]SelectOption, 0, len(raids))
	for _, r := range raids {
		ans, found, err := s.d.Guilds.ReportFor(ctx, guildID, viewer, &r.ID)
		if err != nil {
			return nil, fmt.Errorf("discordbot: report night: %w", err)
		}
		if !found || ans.Raid.Night == nil {
			continue
		}
		options = append(options, SelectOption{
			Label: raidLabel(r.Title, r.StartsAt, raidLocation(zone)), Value: strconv.FormatInt(r.ID, 10),
			Description: truncateRunes(nightLine(*ans.Raid.Night), maxOptionRunes), Default: r.ID == shown,
		})
	}
	if len(options) < 2 {
		return nil, nil
	}
	return []ActionRow{scopeSelect(prefixReport, guideCopy.ReportPlaceholder, options)}, nil
}

// withNightRow puts the night select under the presser's own copy of a shared report, so the
// answer keeps its control after Share.
func (s *Service) withNightRow(ctx context.Context, guildID, viewer int64, opts map[string]string, msg Message) (Message, error) {
	shown, _ := strconv.ParseInt(opts[optReportRaid], 10, 64)
	rows, err := s.reportNightRow(ctx, guildID, viewer, shown)
	if err != nil {
		return Message{}, err
	}
	msg.Components = append(append([]ActionRow(nil), msg.Components...), rows...)
	return msg, nil
}

// reportPress answers a pick on the /report select: that night's report, with a fresh Share
// when its log is public.
func (s *Service) reportPress() Handler {
	return s.cardHandler("/report", map[string]cardPart{
		partScopeRaid: func(ctx context.Context, in Interaction, _ cardPress) (Reply, error) {
			guildID, viewer, bound, err := s.lookupCaller(ctx, in)
			if err != nil {
				return Reply{}, err
			}
			if !bound {
				return pressReply(commonCopy.Unbound, true), nil
			}
			if viewer == nil {
				return pressReply(officerCopy.ReportUnlinked, true, s.linkButton()), nil
			}
			return s.reportReply(ctx, in, guildID, *viewer, picked(in), true)
		},
	})
}
