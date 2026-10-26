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
	mux.Handle("GET "+raid+"/match-candidates", officerLimit(auth.RequireSession(s.matchCandidates)))
	mux.Handle("PUT "+raid+"/match", officerLimit(auth.RequireSession(s.putMatch)))
	mux.Handle("DELETE "+raid+"/match", officerLimit(auth.RequireSession(s.deleteMatch)))
	mux.Handle("PUT "+raid+"/attendance/{region}/{ruleset}/{name}", officerLimit(auth.RequireSession(s.putAttendance)))
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
	guildID, raid, ok := s.officerRaid(w, r, "raid_match_candidates")
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
	guildID, raid, ok := s.officerRaid(w, r, "raid_match")
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
	guildID, raid, ok := s.officerRaid(w, r, "raid_match_clear")
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
	guildID, raid, ok := s.officerRaid(w, r, "raid_excuse")
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
