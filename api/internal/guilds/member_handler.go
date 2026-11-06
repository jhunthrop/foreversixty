// api/internal/guilds/member_handler.go
//
// The member profile (officer or self) and the officer notes (officers only) routes. A member
// is named by account key ("u:42", the roster's account_key) or "me".
package guilds

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	profileFailMessage = "could not load that member just now"
	notesFailMessage   = "could not load those notes just now"
)

func mountMemberRoutes(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/guilds/{id}/members/{account}/profile", auth.RequireSession(s.memberProfile))
	mux.HandleFunc("GET /v1/guilds/{id}/members/{account}/notes", auth.RequireSession(s.memberNotes))
	mux.HandleFunc("POST /v1/guilds/{id}/members/{account}/notes", auth.RequireSession(s.addMemberNote))
	mux.HandleFunc("DELETE /v1/guilds/{id}/members/{account}/notes/{note}", auth.RequireSession(s.deleteMemberNote))
}

// subjectFrom reads {account}: "me" is the actor.
func subjectFrom(r *http.Request) (int64, bool) {
	raw := r.PathValue("account")
	if raw == "me" {
		return auth.ActorFrom(r.Context()).UserID, true
	}
	return parseAccountKey(raw)
}

func writeMemberNotFound(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such member of this guild", nil)
}

// officerRank is the actor's rank in guildID when they are a verified officer or guild master,
// "" otherwise.
func (s *Service) officerRank(r *http.Request, guildID int64) (string, error) {
	rank, verified, err := s.Accounts.GuildRank(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil || !verified || (rank != "officer" && rank != "leader") {
		return "", err
	}
	return rank, nil
}

// memberSubject resolves the guild and the member a member route names, writing the 404 itself
// when either is unknown or the account is not in the guild.
func (s *Service) memberSubject(w http.ResponseWriter, r *http.Request, op, failMessage string) (int64, int64, bool) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeGuildNotFound(w, r)
		return 0, 0, false
	}
	subject, ok := subjectFrom(r)
	if !ok {
		writeMemberNotFound(w, r)
		return 0, 0, false
	}
	member, err := s.Store.IsMember(r.Context(), guildID, subject)
	if err != nil {
		s.fail(w, r, op, err, failMessage)
		return 0, 0, false
	}
	if !member {
		writeMemberNotFound(w, r)
		return 0, 0, false
	}
	return guildID, subject, true
}

// memberProfile is GET /v1/guilds/{id}/members/{account}/profile?character_key=: the member
// themself or an officer; notes only for an officer.
func (s *Service) memberProfile(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeGuildNotFound(w, r)
		return
	}
	rank, err := s.officerRank(r, guildID)
	if err != nil {
		s.fail(w, r, "member_profile", err, profileFailMessage)
		return
	}
	subject, ok := subjectFrom(r)
	if ok && rank == "" && subject != auth.ActorFrom(r.Context()).UserID {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "only officers can see another member's profile", nil)
		return
	}
	guildID, subject, ok = s.memberSubject(w, r, "member_profile", profileFailMessage)
	if !ok {
		return
	}
	profile, err := s.Store.MemberProfile(r.Context(), guildID, subject, r.URL.Query().Get("character_key"))
	if err != nil {
		s.fail(w, r, "member_profile", err, profileFailMessage)
		return
	}
	if rank != "" {
		notes, err := s.Store.MemberNotes(r.Context(), guildID, subject, auth.ActorFrom(r.Context()).UserID, rank)
		if err != nil {
			s.fail(w, r, "member_profile", err, profileFailMessage)
			return
		}
		profile.Notes = &notes
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, profile)
}

// noteOfficer resolves an officer note route: the guild, the member, and the officer's rank.
func (s *Service) noteOfficer(w http.ResponseWriter, r *http.Request, write bool) (guildID, subject int64, rank string, ok bool) {
	op := "member_notes"
	if guildID, ok = s.officerGuild(w, r, op, notesFailMessage); !ok {
		return 0, 0, "", false
	}
	if write && !s.officerNotFrozen(w, r, guildID, op, notesFailMessage) {
		return 0, 0, "", false
	}
	if _, subject, ok = s.memberSubject(w, r, op, notesFailMessage); !ok {
		return 0, 0, "", false
	}
	rank, err := s.officerRank(r, guildID)
	if err != nil {
		s.fail(w, r, op, err, notesFailMessage)
		return 0, 0, "", false
	}
	return guildID, subject, rank, true
}

// memberNotes is GET /v1/guilds/{id}/members/{account}/notes: officers only, newest first.
func (s *Service) memberNotes(w http.ResponseWriter, r *http.Request) {
	guildID, subject, rank, ok := s.noteOfficer(w, r, false)
	if !ok {
		return
	}
	notes, err := s.Store.MemberNotes(r.Context(), guildID, subject, auth.ActorFrom(r.Context()).UserID, rank)
	if err != nil {
		s.fail(w, r, "member_notes", err, notesFailMessage)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": notes})
}

type noteInput struct {
	Body string `json:"body"`
}

// addMemberNote is POST /v1/guilds/{id}/members/{account}/notes.
func (s *Service) addMemberNote(w http.ResponseWriter, r *http.Request) {
	guildID, subject, rank, ok := s.noteOfficer(w, r, true)
	if !ok {
		return
	}
	var in noteInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	note, err := s.Store.AddMemberNote(r.Context(), guildID, subject, auth.ActorFrom(r.Context()).UserID, rank, in.Body)
	switch {
	case errors.Is(err, ErrInvalidNote):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "a note is 1 to 1000 characters",
			map[string]string{"body": "1 to 1000 characters"})
	case errors.Is(err, ErrNotFound):
		writeMemberNotFound(w, r)
	case err != nil:
		s.fail(w, r, "member_note_add", err, notesFailMessage)
	default:
		httpx.WriteOK(w, r, http.StatusCreated, note)
	}
}

// deleteMemberNote is DELETE /v1/guilds/{id}/members/{account}/notes/{note}: the author's own,
// or any for the guild master.
func (s *Service) deleteMemberNote(w http.ResponseWriter, r *http.Request) {
	guildID, subject, rank, ok := s.noteOfficer(w, r, true)
	if !ok {
		return
	}
	noteID, err := strconv.ParseInt(r.PathValue("note"), 10, 64)
	if err != nil || noteID <= 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such note", nil)
		return
	}
	err = s.Store.DeleteMemberNote(r.Context(), guildID, subject, noteID, auth.ActorFrom(r.Context()).UserID, rank)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such note", nil)
	case errors.Is(err, ErrNotNoteAuthor):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "only the note's author or the guild master can delete it", nil)
	case err != nil:
		s.fail(w, r, "member_note_delete", err, notesFailMessage)
	default:
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "removed"})
	}
}
