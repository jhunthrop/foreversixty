// api/internal/guilds/discord_reads.go
//
// The reads behind the Discord bot's /report (spec 2026-10-27-discord-bot §3.3.3): a done
// raid's report as a viewer may see it. The raid page's GET …/raids/{raid} answers through
// MemberRaidView, the same read, so the bot and the page never disagree. /attendance and
// /readiness are in discord_reads_attendance.go and discord_reads_readiness.go.
package guilds

import (
	"context"
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// The refusals of the bot's officer reads. ErrPlanRequired and ErrClaimContested (the sign-up
// service's) complete the set.
var (
	// ErrNotOfficer is a viewer who is not a verified officer or leader of the guild.
	ErrNotOfficer = errors.New("guilds: not a verified officer of the guild")
	// ErrRaidGuildRequired is an officer whose guild's plan does not cover the Raid guild tools.
	ErrRaidGuildRequired = errors.New("guilds: the Raid guild plan is required")
	// ErrOfficerOnly is a member naming a character that is not theirs.
	ErrOfficerOnly = errors.New("guilds: only officers can name another member")
	// ErrNoSuchMember is a name no verified character of the guild carries.
	ErrNoSuchMember = errors.New("guilds: no member by that name")
)

// ReportAnswer is a done raid's report as one viewer sees it. Officer is a verified officer
// or leader (every row of the report); Public is a matched log a signed-out visitor may open.
type ReportAnswer struct {
	Raid    RaidView
	Report  *ReportView
	Officer bool
	Public  bool
}

// MemberRaidView is raidID's view for a verified member of the guild (R1): found is false for
// anyone else and for a raid that is not the guild's.
func (s *Service) MemberRaidView(ctx context.Context, guildID, viewerID, raidID int64) (RaidView, bool, error) {
	_, verified, err := s.Accounts.GuildRank(ctx, guildID, viewerID)
	if err != nil || !verified {
		return RaidView{}, false, err
	}
	raid, found, err := s.Store.GetRaid(ctx, guildID, raidID)
	if err != nil || !found {
		return RaidView{}, false, err
	}
	view, err := s.raidViewOf(ctx, guildID, viewerID, raid)
	if err != nil {
		return RaidView{}, false, err
	}
	return view, true, nil
}

// ReportFor is the named done raid's report (raidID nil: the guild's last one) as viewerID sees
// it. A verified member sees every logged raid, scoped as the raid page scopes it; anyone else
// (viewerID 0 is a signed-out visitor) only a raid whose log is public, with no one's rows.
// found is false when there is no such report for the viewer.
func (s *Service) ReportFor(ctx context.Context, guildID, viewerID int64, raidID *int64) (ReportAnswer, bool, error) {
	member := false
	if viewerID != noViewer {
		_, verified, err := s.Accounts.GuildRank(ctx, guildID, viewerID)
		if err != nil {
			return ReportAnswer{}, false, fmt.Errorf("guilds: report viewer: %w", err)
		}
		member = verified
	}
	id, found, err := s.reportRaidID(ctx, guildID, raidID, !member)
	if err != nil || !found {
		return ReportAnswer{}, false, err
	}
	view, found, err := s.reportView(ctx, guildID, viewerID, id, member)
	if err != nil || !found || view.Status != RaidDone || view.Report == nil || view.Night == nil || view.Match == nil {
		return ReportAnswer{}, false, err
	}
	public, err := s.Store.reportIsPublic(ctx, view.Match.ReportID)
	if err != nil || (!member && !public) {
		return ReportAnswer{}, false, err
	}
	officer := false
	if member {
		if officer, err = s.seesEveryAttendanceRow(ctx, guildID, viewerID); err != nil {
			return ReportAnswer{}, false, err
		}
	}
	return ReportAnswer{Raid: view, Report: view.Report, Officer: officer, Public: public}, true, nil
}

// reportRaidID is the named raid, or the guild's last logged one (public only when asked).
func (s *Service) reportRaidID(ctx context.Context, guildID int64, raidID *int64, publicOnly bool) (int64, bool, error) {
	if raidID != nil {
		return *raidID, true, nil
	}
	raids, err := s.DoneRaids(ctx, guildID, publicOnly, 1)
	if err != nil || len(raids) == 0 {
		return 0, false, err
	}
	return raids[0].ID, true, nil
}

func (s *Service) reportView(ctx context.Context, guildID, viewerID, raidID int64, member bool) (RaidView, bool, error) {
	if member {
		return s.MemberRaidView(ctx, guildID, viewerID, raidID)
	}
	raid, found, err := s.Store.GetRaid(ctx, guildID, raidID)
	if err != nil || !found {
		return RaidView{}, false, err
	}
	view, err := s.raidViewOf(ctx, guildID, noViewer, raid)
	return view, err == nil, err
}

// DoneRaids is the guild's done raids whose matched log may be read, newest first, at most
// limit; publicOnly keeps those whose log a visitor may open.
func (s *Service) DoneRaids(ctx context.Context, guildID int64, publicOnly bool, limit int) ([]Raid, error) {
	rows, err := s.Store.Pool.Query(ctx, `
		select `+raidColumns+` from guild_raids where id in (
		  select g.id from guild_raids g join reports r on r.id = g.report_id and `+presence.EvidenceReport("r")+`
		  where g.guild_id = $1 and g.status = 'done' and (not $2 or r.visibility = 'public'))
		order by starts_at desc, id desc limit $3`, guildID, publicOnly, limit)
	if err != nil {
		return nil, fmt.Errorf("guilds: done raids: %w", err)
	}
	raids, err := scanRaids(rows)
	if err != nil {
		return nil, fmt.Errorf("guilds: done raids: %w", err)
	}
	return raids, nil
}

// reportIsPublic is whether a signed-out visitor may open the report: public and evidence.
func (s *Store) reportIsPublic(ctx context.Context, reportID string) (bool, error) {
	var ok bool
	err := s.Pool.QueryRow(ctx, `select exists (select 1 from reports r
		where r.id = $1 and r.visibility = 'public' and `+presence.EvidenceReport("r")+`)`, reportID).Scan(&ok)
	if err != nil {
		return false, fmt.Errorf("guilds: report public: %w", err)
	}
	return ok, nil
}

// GuildPagePathOf is the guild page's site path (GuildPagePath).
func (s *Service) GuildPagePathOf(ctx context.Context, guildID int64) (string, error) {
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return "", fmt.Errorf("guilds: guild path: %w", err)
	}
	return GuildPagePath(g), nil
}
