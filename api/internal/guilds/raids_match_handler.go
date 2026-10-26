// api/internal/guilds/raids_match_handler.go
//
// Lane G2: the officer routes that tie a raid to its log and mark a character excused. Every
// route runs the officer raid gate (rank, plan, freeze) and the officer raid limiter; each
// action answers the raid's view.
package guilds

import (
	"errors"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// RaidRefView is the raid a log is already matched to.
type RaidRefView struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	StartsAt time.Time `json:"starts_at"`
}

// CandidateView is one log the picker lists.
type CandidateView struct {
	ReportID       string       `json:"report_id"`
	ZoneName       string       `json:"zone_name"`
	UploadedAt     time.Time    `json:"uploaded_at"`
	UploadedByName string       `json:"uploaded_by_name"`
	Pulls          int          `json:"pulls"`
	Kills          int          `json:"kills"`
	DurationMS     int64        `json:"duration_ms"`
	Closest        bool         `json:"closest"`
	MatchedRaid    *RaidRefView `json:"matched_raid"`
}

// MatchCandidatesView is GET …/raids/{raid}/match-candidates.
type MatchCandidatesView struct {
	ZoneKnown bool            `json:"zone_known"`
	Rows      []CandidateView `json:"rows"`
}

type matchRequest struct {
	ReportID string `json:"report_id"`
}

type excuseRequest struct {
	Excused *bool `json:"excused"`
}

func mountRaidMatchRoutes(mux *http.ServeMux, s *Service, officerLimit func(http.Handler) http.Handler) {
	const raid = "/v1/guilds/{id}/raids/{raid}"
	mux.Handle("GET "+raid+"/match-candidates", officerLimit(raidRouteSession(s.matchCandidates)))
	mux.Handle("PUT "+raid+"/match", officerLimit(raidRouteSession(s.putMatch)))
	mux.Handle("DELETE "+raid+"/match", officerLimit(raidRouteSession(s.deleteMatch)))
	mux.Handle("PUT "+raid+"/attendance/{region}/{ruleset}/{name}", officerLimit(raidRouteSession(s.putAttendance)))
}

// raidRouteSession wraps a Lane G2 raid route for R1: a signed-out caller is told 404, as the
// raid read tells them; a device token is refused as RequireSession refuses it.
func raidRouteSession(h http.HandlerFunc) http.HandlerFunc {
	session := auth.RequireSession(h)
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.ActorFrom(r.Context()).Signed() {
			writeRaidNotFound(w, r)
			return
		}
		session(w, r)
	}
}

// verifiedMemberOr404 answers 404 (R1) unless the caller is a verified member of the route's
// guild; ok false once it has answered.
func (s *Service) verifiedMemberOr404(w http.ResponseWriter, r *http.Request, op string) bool {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeRaidNotFound(w, r)
		return false
	}
	_, verified, err := s.Accounts.GuildRank(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, op, err, raidFailMessage)
		return false
	}
	if !verified {
		writeRaidNotFound(w, r)
		return false
	}
	return true
}

// matchOfficerRaid is officerRaid behind R1: 404 for anyone who is not a verified member, 403
// for a verified member who is not an officer.
func (s *Service) matchOfficerRaid(w http.ResponseWriter, r *http.Request, op string) (int64, Raid, bool) {
	if !s.verifiedMemberOr404(w, r, op) {
		return 0, Raid{}, false
	}
	return s.officerRaid(w, r, op)
}

func candidateViews(rows []Candidate) []CandidateView {
	out := make([]CandidateView, len(rows))
	for i, c := range rows {
		out[i] = CandidateView{ReportID: c.ReportID, ZoneName: c.ZoneName, UploadedAt: c.UploadedAt.UTC(),
			UploadedByName: c.UploadedByName, Pulls: c.Pulls, Kills: c.Kills, DurationMS: c.DurationMS, Closest: c.Closest}
		if c.MatchedRaid != nil {
			out[i].MatchedRaid = &RaidRefView{ID: c.MatchedRaid.ID, Title: c.MatchedRaid.Title, StartsAt: c.MatchedRaid.StartsAt.UTC()}
		}
	}
	return out
}

func (s *Service) matchCandidates(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.matchOfficerRaid(w, r, "raid_match_candidates")
	if !ok {
		return
	}
	res, err := s.Store.MatchCandidates(r.Context(), guildID, raid.ID, time.Now())
	if err != nil {
		s.writeRaidWriteErr(w, r, "raid_match_candidates", err)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, MatchCandidatesView{ZoneKnown: res.ZoneKnown, Rows: candidateViews(res.Rows)})
}

func (s *Service) putMatch(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.matchOfficerRaid(w, r, "raid_match")
	if !ok {
		return
	}
	var req matchRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	if req.ReportID == "" {
		writeInvalid(w, r, map[string]string{"report_id": "pick a log to match"})
		return
	}
	updated, err := s.Store.SetMatch(r.Context(), guildID, raid.ID, req.ReportID, MatchedHand)
	switch {
	case errors.Is(err, ErrReportMatched):
		httpx.WriteError(w, r, http.StatusConflict, "report_matched", "that log is already matched to another raid", nil)
		return
	case errors.Is(err, ErrRaidNotEnded):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "that raid has not ended yet", nil)
		return
	case errors.Is(err, ErrNotFound):
		writeInvalid(w, r, map[string]string{"report_id": "that is not a finished log of this guild"})
		return
	case !s.writeRaidWriteErr(w, r, "raid_match", err):
		return
	}
	s.logger().Info("guilds", "op", "raid_match", "guild_id", guildID, "raid_id", raid.ID,
		"report_id", req.ReportID, "user_id", auth.ActorFrom(r.Context()).UserID)
	s.writeRaidView(w, r, "raid_match", guildID, updated, nil)
}

func (s *Service) deleteMatch(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.matchOfficerRaid(w, r, "raid_match_clear")
	if !ok {
		return
	}
	userID := auth.ActorFrom(r.Context()).UserID
	updated, err := s.Store.ClearMatch(r.Context(), guildID, raid.ID, userID)
	if !s.writeRaidWriteErr(w, r, "raid_match_clear", err) {
		return
	}
	s.logger().Info("guilds", "op", "raid_match_clear", "guild_id", guildID, "raid_id", raid.ID, "user_id", userID)
	s.writeRaidView(w, r, "raid_match_clear", guildID, updated, nil)
}

func (s *Service) putAttendance(w http.ResponseWriter, r *http.Request) {
	guildID, raid, ok := s.matchOfficerRaid(w, r, "raid_excuse")
	if !ok {
		return
	}
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character in this guild", nil)
		return
	}
	var req excuseRequest
	if err := decodeJSON(w, r, &req); err != nil {
		return
	}
	if req.Excused == nil {
		writeInvalid(w, r, map[string]string{"excused": "say whether the character is excused"})
		return
	}
	err := s.Store.SetExcused(r.Context(), raid.ID, key, *req.Excused)
	if errors.Is(err, ErrNotFound) {
		writeInvalid(w, r, map[string]string{"character": "that character has no attendance row on this raid"})
		return
	}
	if !s.writeRaidWriteErr(w, r, "raid_excuse", err) {
		return
	}
	s.writeRaidView(w, r, "raid_excuse", guildID, raid, nil)
}
