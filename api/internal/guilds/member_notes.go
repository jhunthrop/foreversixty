// api/internal/guilds/member_notes.go
//
// Private officer notes on a guild member (design/specs/2026-11-05-web-layout.md §7.2 Member
// profile; migration 0075 guild_member_notes). A note is about the member's account, not one
// character: the account survives a rename and an alt swap, and a raider is one person on the
// roster whichever character they bring. Only officers ever read them: no member-visible read
// (home, roster, readiness, profile as self) selects this table.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
)

// memberNoteMax bounds a note's body (runes), migration 0075's check.
const memberNoteMax = 1000

// ErrInvalidNote is an empty or too long note; ErrNotNoteAuthor a delete by someone who may not.
var (
	ErrInvalidNote   = errors.New("guilds: invalid note")
	ErrNotNoteAuthor = errors.New("guilds: only the note's author or the guild master may delete it")
)

// MemberNote is one officer note.
type MemberNote struct {
	ID        int64           `json:"id"`
	Body      string          `json:"body"`
	Author    *ClaimHolderRef `json:"author"`
	CreatedAt time.Time       `json:"created_at"`
	// MayDelete is whether the viewing officer may delete it: their own, or any for the guild master.
	MayDelete bool `json:"may_delete"`
	authorID  *int64
}

// MemberNotes is the notes on subjectID in guildID, newest first, as officer viewerID (whose
// rank is viewerRank) reads them. Never nil.
func (s *Store) MemberNotes(ctx context.Context, guildID, subjectID, viewerID int64, viewerRank string) ([]MemberNote, error) {
	rows, err := s.Pool.Query(ctx, `select id, body, author_id, created_at from guild_member_notes
		where guild_id = $1 and subject_user_id = $2 order by created_at desc, id desc`, guildID, subjectID)
	if err != nil {
		return nil, fmt.Errorf("guilds: member notes: %w", err)
	}
	notes, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (MemberNote, error) {
		var n MemberNote
		err := row.Scan(&n.ID, &n.Body, &n.authorID, &n.CreatedAt)
		return n, err
	})
	if err != nil {
		return nil, fmt.Errorf("guilds: member notes: %w", err)
	}
	return s.withNoteAuthors(ctx, guildID, viewerID, viewerRank, notes)
}

// withNoteAuthors names each note's author by their character on the roster, once per author.
func (s *Store) withNoteAuthors(ctx context.Context, guildID, viewerID int64, viewerRank string, notes []MemberNote) ([]MemberNote, error) {
	refs := map[int64]*ClaimHolderRef{}
	out := make([]MemberNote, 0, len(notes))
	for _, n := range notes {
		if n.authorID != nil {
			ref, seen := refs[*n.authorID]
			if !seen {
				var err error
				if ref, err = s.claimHolderRef(ctx, guildID, *n.authorID); err != nil {
					return nil, err
				}
				refs[*n.authorID] = ref
			}
			n.Author = ref
		}
		n.MayDelete = viewerRank == "leader" || (n.authorID != nil && *n.authorID == viewerID)
		out = append(out, n)
	}
	return out, nil
}

// AddMemberNote files authorID's note on subjectID, a member of guildID (ErrNotFound otherwise).
func (s *Store) AddMemberNote(ctx context.Context, guildID, subjectID, authorID int64, viewerRank, body string) (MemberNote, error) {
	body = strings.TrimSpace(body)
	if body == "" || utf8.RuneCountInString(body) > memberNoteMax {
		return MemberNote{}, ErrInvalidNote
	}
	member, err := s.IsMember(ctx, guildID, subjectID)
	if err != nil {
		return MemberNote{}, err
	}
	if !member {
		return MemberNote{}, ErrNotFound
	}
	n := MemberNote{Body: body, authorID: &authorID}
	if err := s.Pool.QueryRow(ctx, `insert into guild_member_notes (guild_id, subject_user_id, author_id, body)
		values ($1, $2, $3, $4) returning id, created_at`, guildID, subjectID, authorID, body).Scan(&n.ID, &n.CreatedAt); err != nil {
		return MemberNote{}, fmt.Errorf("guilds: add member note: %w", err)
	}
	notes, err := s.withNoteAuthors(ctx, guildID, authorID, viewerRank, []MemberNote{n})
	if err != nil {
		return MemberNote{}, err
	}
	return notes[0], nil
}

// DeleteMemberNote removes a note on subjectID in guildID: the author's own, or any for the
// guild master (viewerRank "leader"). ErrNotFound for no such note there.
func (s *Store) DeleteMemberNote(ctx context.Context, guildID, subjectID, noteID, viewerID int64, viewerRank string) error {
	var authorID *int64
	err := s.Pool.QueryRow(ctx, `select author_id from guild_member_notes
		where id = $1 and guild_id = $2 and subject_user_id = $3`, noteID, guildID, subjectID).Scan(&authorID)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("guilds: delete member note: %w", err)
	}
	if viewerRank != "leader" && (authorID == nil || *authorID != viewerID) {
		return ErrNotNoteAuthor
	}
	if _, err := s.Pool.Exec(ctx, `delete from guild_member_notes where id = $1`, noteID); err != nil {
		return fmt.Errorf("guilds: delete member note: %w", err)
	}
	return nil
}
