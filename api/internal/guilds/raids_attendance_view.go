// api/internal/guilds/raids_attendance_view.go
//
// Lane G2: the match, night and attendance halves of a raid view (design/specs/
// 2026-10-27-guild-attendance.md, contract paragraph). The store fills every figure for every
// viewer; Service.scopeAttendance then narrows the attendance rows to what the viewer may see.
package guilds

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// MatchView is the log a raid is tied to.
type MatchView struct {
	ReportID       string    `json:"report_id"`
	ZoneName       string    `json:"zone_name"`
	UploadedAt     time.Time `json:"uploaded_at"`
	UploadedByName string    `json:"uploaded_by_name"`
	Pulls          int       `json:"pulls"`
	Kills          int       `json:"kills"`
	DurationMS     int64     `json:"duration_ms"`
	How            string    `json:"how"`
}

// NightView is the night's figures, told to everyone: they name nobody.
type NightView struct {
	Signed     int   `json:"signed"`
	Came       int   `json:"came"`
	Late       int   `json:"late"`
	NotInLog   int   `json:"not_in_log"`
	Joined     int   `json:"joined"`
	Excused    int   `json:"excused"`
	Kills      int   `json:"kills"`
	Wipes      int   `json:"wipes"`
	Deaths     int   `json:"deaths"`
	DurationMS int64 `json:"duration_ms"`
}

// AttendanceView is one character's night. Own is true for the viewer's characters.
type AttendanceView struct {
	CharacterKey     string  `json:"character_key"`
	Name             string  `json:"name"`
	Class            string  `json:"class"`
	Spec             string  `json:"spec"`
	Signed           *string `json:"signed"`
	Role             *string `json:"role"`
	Outcome          string  `json:"outcome"`
	CameAs           *string `json:"came_as"`
	FirstSeenOffsetS *int    `json:"first_seen_offset_s"`
	PullsPresent     int     `json:"pulls_present"`
	PullsTotal       int     `json:"pulls_total"`
	Excused          bool    `json:"excused"`
	Own              bool    `json:"own"`
}

// raidMatchMeta is one raid's match columns and its log's figures.
type raidMatchMeta struct {
	preDone    *string
	matchedBy  *string
	reportID   *string
	uploadedAt time.Time
	uploader   string
	pulls      int
	kills      int
	durationMS int64
	deaths     int
	postedAt   *time.Time
}

// raidMatchSQL reads the match columns of the raids in $1 with the figures of their logs.
// The uploader is the logging character's name, as the match picker lists it.
const raidMatchSQL = `
	select g.id, g.pre_done_status, g.matched_by, g.report_id,
	       coalesce(r.completed_at, r.created_at, now()),
	       coalesce(nullif(c.name, ''), split_part(coalesce(r.logging_character, ''), '/', 3)),
	       coalesce(f.pulls, 0), coalesce(f.kills, 0), coalesce(f.duration_ms, 0), coalesce(f.deaths, 0), g.report_posted_at
	from guild_raids g
	left join reports r on r.id = g.report_id
	left join characters c on c.key = r.logging_character
	left join lateral (
	  select count(*) filter (where encounter_id is not null) as pulls,
	         count(*) filter (where encounter_id is not null and kill) as kills,
	         coalesce(max(start_ms + duration_ms) - min(start_ms), 0) as duration_ms,
	         coalesce(sum(deaths) filter (where encounter_id is not null), 0) as deaths
	  from fights where report_id = g.report_id) f on true
	where g.id = any($1)`

func (s *Store) raidMatchMetas(ctx context.Context, ids []int64) (map[int64]raidMatchMeta, error) {
	rows, err := s.Pool.Query(ctx, raidMatchSQL, ids)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid matches: %w", err)
	}
	defer rows.Close()
	out := map[int64]raidMatchMeta{}
	for rows.Next() {
		var id int64
		var m raidMatchMeta
		if err := rows.Scan(&id, &m.preDone, &m.matchedBy, &m.reportID, &m.uploadedAt, &m.uploader,
			&m.pulls, &m.kills, &m.durationMS, &m.deaths, &m.postedAt); err != nil {
			return nil, fmt.Errorf("guilds: raid matches: %w", err)
		}
		out[id] = m
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: raid matches: %w", err)
	}
	return out, nil
}

// fillMatchViews adds pre_done_status, match, night and every attendance row to views.
func (s *Store) fillMatchViews(ctx context.Context, vc raidViewContext, views []RaidView) error {
	ids := make([]int64, len(views))
	for i, v := range views {
		ids[i] = v.ID
	}
	metas, err := s.raidMatchMetas(ctx, ids)
	if err != nil {
		return err
	}
	for i := range views {
		meta := metas[views[i].ID]
		views[i].PreDoneStatus = meta.preDone
		if meta.reportID == nil {
			continue
		}
		views[i].Match = &MatchView{ReportID: *meta.reportID, ZoneName: views[i].ZoneName,
			UploadedAt: meta.uploadedAt.UTC(), UploadedByName: meta.uploader, Pulls: meta.pulls,
			Kills: meta.kills, DurationMS: meta.durationMS, How: derefString(meta.matchedBy)}
		rows, err := s.AttendanceRows(ctx, views[i].ID)
		if err != nil {
			return err
		}
		night := nightOf(rows, meta)
		views[i].Night = &night
		views[i].Attendance = attendanceViews(rows, vc)
		if views[i].Report, err = s.buildReport(ctx, vc, views[i], meta); err != nil {
			return err
		}
	}
	return nil
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func attendanceViews(rows []AttendanceRow, vc raidViewContext) []AttendanceView {
	if len(rows) == 0 {
		return nil
	}
	out := make([]AttendanceView, len(rows))
	for i, r := range rows {
		owner, known := vc.ownerOf[r.CharacterKey]
		out[i] = AttendanceView{CharacterKey: r.CharacterKey, Name: r.Name, Class: r.Class, Spec: r.Spec,
			Signed: r.Signed, Role: r.Role, Outcome: r.Outcome, CameAs: r.CameAs,
			FirstSeenOffsetS: r.FirstSeenOffsetS, PullsPresent: r.PullsPresent, PullsTotal: r.PullsTotal,
			Excused: r.Excused, Own: known && owner == vc.viewerID}
	}
	return out
}

func wasSigned(r AttendanceRow) bool {
	return r.Signed != nil && (*r.Signed == SignupIn || *r.Signed == SignupLate)
}

func cameToLog(r AttendanceRow) bool { return r.Outcome == OutcomePresent || r.Outcome == OutcomeLate }

// nightOf counts the night from the rows and the log's figures. Joined are the characters in
// the log with no In or Late sign-up, except one an account's signed character came as.
func nightOf(rows []AttendanceRow, meta raidMatchMeta) NightView {
	cameAs := map[string]bool{}
	for _, r := range rows {
		if r.CameAs != nil {
			cameAs[*r.CameAs] = true
		}
	}
	n := NightView{Kills: meta.kills, Wipes: meta.pulls - meta.kills, Deaths: meta.deaths, DurationMS: meta.durationMS}
	for _, r := range rows {
		if r.Excused {
			n.Excused++
		}
		switch {
		case wasSigned(r):
			n.Signed++
			countSigned(&n, r)
		case cameToLog(r) && !cameAs[r.Name]:
			n.Joined++
		}
	}
	return n
}

func countSigned(n *NightView, r AttendanceRow) {
	switch r.Outcome {
	case OutcomePresent:
		n.Came++
	case OutcomeLate:
		n.Came++
		n.Late++
	case OutcomeNoShow:
		n.NotInLog++
	}
}

// scopeAttendance narrows views' attendance rows to the viewer's: an officer of a guild with
// the plan keeps every row, anyone else their own characters' (null when they have none).
func (s *Service) scopeAttendance(ctx context.Context, guildID, viewerID int64, views []RaidView) error {
	if !anyMatched(views) {
		return nil
	}
	officer, err := s.seesEveryAttendanceRow(ctx, guildID, viewerID)
	if err != nil || officer {
		return err
	}
	for i := range views {
		views[i].Attendance = ownRows(views[i].Attendance)
		views[i].Report = views[i].Report.ownOnly()
	}
	return nil
}

func anyMatched(views []RaidView) bool {
	for _, v := range views {
		if v.Match != nil {
			return true
		}
	}
	return false
}

func ownRows(rows []AttendanceView) []AttendanceView {
	var own []AttendanceView
	for _, r := range rows {
		if r.Own {
			own = append(own, r)
		}
	}
	return own
}

func (s *Service) seesEveryAttendanceRow(ctx context.Context, guildID, viewerID int64) (bool, error) {
	rank, verified, err := s.Accounts.GuildRank(ctx, guildID, viewerID)
	if err != nil || !verified || (rank != "officer" && rank != "leader") || s.Entitlements == nil {
		return false, err
	}
	ok, _, err := s.Entitlements.CanForGuild(ctx, guildID, viewerID, entitlements.FeatureOfficerViews)
	return ok, err
}
