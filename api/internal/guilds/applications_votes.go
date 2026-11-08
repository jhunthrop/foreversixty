// api/internal/guilds/applications_votes.go
//
// Officers' yes or no on an application (migration 0075 guild_application_votes), and the
// officer list that carries them.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// The vote words.
const (
	VoteYes = "yes"
	VoteNo  = "no"
)

// Votes is one application's tally: Voters is how many verified officers (guild master
// included) the guild has, Mine the viewer's own vote, null when none.
type Votes struct {
	Yes    int     `json:"yes"`
	No     int     `json:"no"`
	Voters int     `json:"voters"`
	Mine   *string `json:"mine"`
}

// OfficerApplication is one application as officers read it.
type OfficerApplication struct {
	Application
	Votes     Votes           `json:"votes"`
	DecidedBy *ClaimHolderRef `json:"decided_by"`
	// RoleOutcome and DMOutcome are what the bot did on the decision (DiscordOutcome's words),
	// null until one is made.
	RoleOutcome *string `json:"role_outcome"`
	DMOutcome   *string `json:"dm_outcome"`
	// Sheet is what the officer reads before Accept (applicant_sheet.go).
	Sheet *ApplicantSheet `json:"sheet"`
}

// officerVoters is how many verified officer and guild master accounts guildID has.
func (s *Store) officerVoters(ctx context.Context, guildID int64) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx, `select count(*) from guild_members
		where guild_id = $1 and verified_at is not null and rank in ('officer', 'leader')`, guildID).Scan(&n); err != nil {
		return 0, fmt.Errorf("guilds: officer voters: %w", err)
	}
	return n, nil
}

// votesFor is each application's tally as viewerID reads it.
func (s *Store) votesFor(ctx context.Context, guildID, viewerID int64, ids []int64) (map[int64]Votes, error) {
	voters, err := s.officerVoters(ctx, guildID)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]Votes, len(ids))
	for _, id := range ids {
		out[id] = Votes{Voters: voters}
	}
	rows, err := s.Pool.Query(ctx, `
		select application_id, count(*) filter (where yes), count(*) filter (where not yes),
		       max(case when voter_id = $2 then case when yes then 'yes' else 'no' end end)
		from guild_application_votes where application_id = any($1) group by application_id`, ids, viewerID)
	if err != nil {
		return nil, fmt.Errorf("guilds: application votes: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		v := Votes{Voters: voters}
		if err := rows.Scan(&id, &v.Yes, &v.No, &v.Mine); err != nil {
			return nil, fmt.Errorf("guilds: application votes: %w", err)
		}
		out[id] = v
	}
	return out, rows.Err()
}

// OfficerApplications is the officer list: every pending application, oldest first, then the
// latest applicationDecidedShown decided ones, newest first. Never nil.
func (s *Store) OfficerApplications(ctx context.Context, guildID, viewerID int64) ([]OfficerApplication, error) {
	apps, err := s.applications(ctx, `select `+applicationColumns+` from (
		  (select *, 0 as part, extract(epoch from created_at) as at
		   from guild_applications where guild_id = $1 and status = 'pending')
		  union all
		  (select *, 1, -extract(epoch from coalesce(decided_at, created_at))
		   from guild_applications where guild_id = $1 and status <> 'pending'
		   order by coalesce(decided_at, created_at) desc limit $2)
		) x
		order by part, at, id`, guildID, applicationDecidedShown)
	if err != nil {
		return nil, err
	}
	return s.officerViews(ctx, guildID, viewerID, apps)
}

// officerViews adds the votes and the deciding officer to each application.
func (s *Store) officerViews(ctx context.Context, guildID, viewerID int64, apps []Application) ([]OfficerApplication, error) {
	ids := make([]int64, 0, len(apps))
	for _, a := range apps {
		ids = append(ids, a.ID)
	}
	votes, err := s.votesFor(ctx, guildID, viewerID, ids)
	if err != nil {
		return nil, err
	}
	zone, err := s.nextZone(ctx, guildID)
	if err != nil {
		return nil, err
	}
	out := make([]OfficerApplication, 0, len(apps))
	for _, a := range apps {
		view := OfficerApplication{Application: a, Votes: votes[a.ID], RoleOutcome: a.roleOutcome, DMOutcome: a.dmOutcome}
		sheet, err := s.applicantSheet(ctx, guildID, a, zone)
		if err != nil {
			return nil, err
		}
		view.Sheet = &sheet
		if a.decidedBy != nil && a.Status != ApplicationWithdrawn {
			if view.DecidedBy, err = s.claimHolderRef(ctx, guildID, *a.decidedBy); err != nil {
				return nil, err
			}
		}
		out = append(out, view)
	}
	return out, nil
}

// OfficerApplication is one application of guildID with its votes as viewerID reads them.
func (s *Store) OfficerApplication(ctx context.Context, guildID, viewerID, id int64) (OfficerApplication, error) {
	a, err := s.application(ctx, guildID, id)
	if err != nil {
		return OfficerApplication{}, err
	}
	views, err := s.officerViews(ctx, guildID, viewerID, []Application{a})
	if err != nil {
		return OfficerApplication{}, err
	}
	return views[0], nil
}

// Vote sets (vote "yes" or "no") or, with vote "", removes officerID's vote on a pending
// application of guildID, answering the new tally.
func (s *Store) Vote(ctx context.Context, guildID, officerID, id int64, vote string) (Votes, error) {
	a, err := s.application(ctx, guildID, id)
	if err != nil {
		return Votes{}, err
	}
	if a.Status != ApplicationPending {
		return Votes{}, ErrApplicationClosed
	}
	if vote == "" {
		_, err = s.Pool.Exec(ctx, `delete from guild_application_votes where application_id = $1 and voter_id = $2`, id, officerID)
	} else {
		_, err = s.Pool.Exec(ctx, `insert into guild_application_votes (application_id, voter_id, yes, voted_at)
			values ($1, $2, $3, $4) on conflict (application_id, voter_id) do update
			set yes = excluded.yes, voted_at = excluded.voted_at`, id, officerID, vote == VoteYes, time.Now())
	}
	if err != nil {
		return Votes{}, fmt.Errorf("guilds: vote: %w", err)
	}
	votes, err := s.votesFor(ctx, guildID, officerID, []int64{id})
	if err != nil {
		return Votes{}, err
	}
	return votes[id], nil
}
