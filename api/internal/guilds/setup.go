// api/internal/guilds/setup.go
//
// The guild setup checklist (design/specs/2026-11-09-guild-onboarding.md): the steps that
// unlock a claimed guild's full use, in order of value, each derived from the guild's own data.
// A mark (guild_setup_marks) records only what cannot be derived: an officer confirming the
// current value, saying the invite link was posted or the linked role made, or skipping a step.
// A step the data says is done reads done whatever its mark.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
)

// The steps, in the checklist's order.
const (
	StepOfficerRanks     = "officer_ranks"
	StepDiscordBot       = "discord_bot"
	StepRecruiting       = "recruiting"
	StepRaidSchedule     = "raid_schedule"
	StepInviteLink       = "invite_link"
	StepAddonSync        = "addon_sync"
	StepReportVisibility = "report_visibility"
	StepLinkedRole       = "linked_role"
)

// SetupSteps is every step in order; the migration's check lists the same keys.
var SetupSteps = []string{StepOfficerRanks, StepDiscordBot, StepRecruiting, StepRaidSchedule,
	StepInviteLink, StepAddonSync, StepReportVisibility, StepLinkedRole}

// confirmableSteps take a "done" mark: their done cannot be read from the data alone.
var confirmableSteps = []string{StepOfficerRanks, StepInviteLink, StepReportVisibility, StepLinkedRole}

// A step's state.
const (
	SetupDone    = "done"
	SetupSkipped = "skipped"
	SetupTodo    = "todo"
)

const (
	// defaultOfficerRank is guilds.officer_max_rank_index's column default (migration 0018):
	// a guild still on it has not said which ranks are officers.
	defaultOfficerRank = 1
	// defaultVisibility is guilds.default_visibility's column default (migration 0005).
	defaultVisibility = "public"
	// syncWindow is how recent an addon export counts as synced.
	syncWindow = 14 * 24 * time.Hour
	// rankSampleNames is how many names each rank line carries.
	rankSampleNames = 3
	// setupNamesShown is how many names the readiness preview and the unsynced list carry.
	setupNamesShown = 8
	// plan words for SetupView.Plan.
	planNone      = "none"
	planGuild     = "guild"
	planRaidGuild = "raid_guild"
)

// ErrSetupStep is an unknown step, or a done mark on a step whose done is derived only.
var ErrSetupStep = errors.New("guilds: not a setup step that takes that mark")

// SetupView is GET /v1/guilds/{id}/setup.
type SetupView struct {
	Steps []SetupStepView `json:"steps"`
	// Done counts the done steps; Finished is every step done or skipped (the checklist then
	// leaves the settings page and the claim page).
	Done     int  `json:"done"`
	Total    int  `json:"total"`
	Finished bool `json:"finished"`
	// Plan is "none", "guild" or "raid_guild".
	Plan       string          `json:"plan"`
	Ranks      SetupRanks      `json:"ranks"`
	Discord    SetupDiscord    `json:"discord"`
	Recruiting SetupRecruiting `json:"recruiting"`
	Schedule   SetupSchedule   `json:"schedule"`
	Invite     SetupInvite     `json:"invite"`
	Sync       SetupSync       `json:"sync"`
	// DefaultVisibility is the guild's default for new uploads (guilds.default_visibility), a
	// setting, not a report's visibility: the step reads no report.
	DefaultVisibility string `json:"visibility"`
	// Readiness is what the Raid guild readiness board would flag for the next raid, counts
	// only, for a guild without Raid guild; nil with Raid guild or no raid ahead.
	Readiness *SetupReadiness `json:"readiness"`
}

// SetupReadiness is the readiness board's counts for the next raid: the raiders signed In or
// Late and how many of them have something to fix.
type SetupReadiness struct {
	RaidTitle string `json:"raid_title"`
	Signed    int    `json:"signed"`
	Failing   int    `json:"failing"`
	// Names are the first failing raiders, without what they fail: the board itself is Raid guild.
	Names []string `json:"names"`
}

// SetupStepView is one step and its state; Marked is true when the state came from a mark.
type SetupStepView struct {
	Key    string `json:"key"`
	State  string `json:"state"`
	Marked bool   `json:"marked"`
}

// SetupRanks is the officer ranks step's facts: the guild's ranks as its characters carry them.
type SetupRanks struct {
	OfficerMaxRankIndex int        `json:"officer_max_rank_index"`
	Officers            int        `json:"officers"`
	Ranks               []RankLine `json:"ranks"`
}

// RankLine is one rank: its index, how many characters hold it, how many of them are level 60
// (a level the site knows from the Blizzard import; an unknown level is not counted) and the
// first few names, level 60s first.
type RankLine struct {
	RankIndex int      `json:"rank_index"`
	Members   int      `json:"members"`
	AtSixty   int      `json:"at_sixty"`
	Names     []string `json:"names"`
}

// SetupDiscord is the Discord bot step's facts.
type SetupDiscord struct {
	BotAvailable       bool    `json:"bot_available"`
	Installed          bool    `json:"installed"`
	ServerName         *string `json:"server_name"`
	ChannelName        *string `json:"channel_name"`
	RecruitChannelName *string `json:"recruit_channel_name"`
	WebhookConnected   bool    `json:"webhook_connected"`
	LinkedMembers      int     `json:"linked_members"`
	Members            int     `json:"members"`
}

// SetupRecruiting is the recruiting step's facts.
type SetupRecruiting struct {
	Open    bool     `json:"open"`
	Pitch   bool     `json:"pitch"`
	Nights  string   `json:"nights"`
	Wanted  []string `json:"wanted"`
	RoleSet bool     `json:"role_set"`
}

// SetupSchedule is the raid schedule step's facts.
type SetupSchedule struct {
	Timezone *string    `json:"timezone"`
	Planned  int        `json:"planned"`
	Next     *SetupRaid `json:"next"`
}

// SetupRaid is the next raid as the schedule step names it; PostedAt is when its Discord post
// went out, nil while it has not.
type SetupRaid struct {
	ID       int64      `json:"id"`
	Title    string     `json:"title"`
	StartsAt time.Time  `json:"starts_at"`
	PostedAt *time.Time `json:"posted_at"`
}

// SetupInvite is the invite link step's facts: whether a link exists and who joined by one.
type SetupInvite struct {
	Made   bool `json:"made"`
	Joined int  `json:"joined"`
}

// SetupSync is the addon step's facts: accounts with a guild character, and how many of them
// synced an addon export inside the window.
type SetupSync struct {
	Synced     int `json:"synced"`
	Accounts   int `json:"accounts"`
	WindowDays int `json:"window_days"`
	// Missing names the first accounts with no export inside the window, by a character of theirs.
	Missing []string `json:"missing"`
}

// Setup is the checklist for viewerID, a verified officer or leader of guildID.
func (s *Service) Setup(ctx context.Context, guildID, viewerID int64) (SetupView, error) {
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return SetupView{}, err
	}
	facts, err := s.setupFacts(ctx, g, viewerID)
	if err != nil {
		return SetupView{}, err
	}
	marks, err := s.Store.setupMarks(ctx, guildID)
	if err != nil {
		return SetupView{}, err
	}
	return facts.withSteps(marks), nil
}

// setupFacts reads every step's facts.
func (s *Service) setupFacts(ctx context.Context, g Guild, viewerID int64) (SetupView, error) {
	v := SetupView{DefaultVisibility: g.DefaultVisibility, Plan: s.setupPlan(ctx, g.ID, viewerID)}
	v.Invite.Made = g.InviteTokenRotatedAt != nil
	v.Ranks.OfficerMaxRankIndex = g.OfficerMaxRankIndex
	steps := []func() error{
		func() (err error) { v.Ranks, err = s.Store.setupRanks(ctx, g.ID, g.OfficerMaxRankIndex); return },
		func() (err error) { v.Discord, err = s.setupDiscord(ctx, g.ID, viewerID); return },
		func() (err error) { v.Recruiting, err = s.setupRecruiting(ctx, g.ID); return },
		func() (err error) { v.Schedule, err = s.setupSchedule(ctx, g.ID); return },
		func() (err error) { v.Invite.Joined, err = s.Store.joinedByInvite(ctx, g.ID); return },
		func() (err error) { v.Sync, err = s.Store.setupSync(ctx, g.ID); return },
	}
	for _, step := range steps {
		if err := step(); err != nil {
			return SetupView{}, fmt.Errorf("guilds: setup %d: %w", g.ID, err)
		}
	}
	if v.Plan != planRaidGuild && v.Schedule.Next != nil {
		preview, err := s.setupReadiness(ctx, g.ID)
		if err != nil {
			return SetupView{}, fmt.Errorf("guilds: setup %d: %w", g.ID, err)
		}
		v.Readiness = preview
	}
	return v, nil
}

// setupReadiness is the next raid's readiness counts, the board's own rule (readinessAnswerOf).
func (s *Service) setupReadiness(ctx context.Context, guildID int64) (*SetupReadiness, error) {
	now := time.Now()
	raid, found, err := s.Store.NextRaid(ctx, guildID, now)
	if err != nil || !found {
		return nil, err
	}
	view, err := s.Store.Readiness(ctx, guildID, true, &raid)
	if err != nil {
		return nil, err
	}
	ans := readinessAnswerOf(raid, view, now)
	return &SetupReadiness{RaidTitle: raid.Title, Signed: ans.Signed, Failing: ans.Failing,
		Names: failingNames(view.Rows, setupNamesShown)}, nil
}

// failingNames is the first n readiness rows with something to fix, by name.
func failingNames(rows []ReadinessRow, n int) []string {
	out := []string{}
	for _, row := range rows {
		if len(out) == n {
			break
		}
		if row.Failing > 0 || row.NoExport {
			out = append(out, row.Name)
		}
	}
	return out
}

// withSteps derives each step's state from the facts and the marks.
func (v SetupView) withSteps(marks map[string]string) SetupView {
	derived := v.derivedDone()
	v.Steps = make([]SetupStepView, 0, len(SetupSteps))
	v.Done, v.Total = 0, len(SetupSteps)
	finished := 0
	for _, key := range SetupSteps {
		step := SetupStepView{Key: key, State: SetupTodo}
		switch mark, marked := marks[key]; {
		case derived[key]:
			step.State = SetupDone
		case marked:
			step.State, step.Marked = mark, true
		}
		if step.State == SetupDone {
			v.Done++
		}
		if step.State != SetupTodo {
			finished++
		}
		v.Steps = append(v.Steps, step)
	}
	v.Finished = finished == v.Total
	return v
}

// derivedDone is which steps the guild's data says are done.
func (v SetupView) derivedDone() map[string]bool {
	r, d := v.Recruiting, v.Discord
	return map[string]bool{
		StepOfficerRanks:     v.Ranks.OfficerMaxRankIndex != defaultOfficerRank,
		StepDiscordBot:       (d.Installed && d.ChannelName != nil) || (!d.BotAvailable && d.WebhookConnected),
		StepRecruiting:       r.Open && r.Pitch && len(r.Wanted) > 0 && d.RecruitChannelName != nil,
		StepRaidSchedule:     v.Schedule.Timezone != nil && v.Schedule.Planned > 0,
		StepInviteLink:       v.Invite.Made && v.Invite.Joined > 0,
		StepAddonSync:        v.Sync.Accounts > 0 && v.Sync.Synced*2 >= v.Sync.Accounts,
		StepReportVisibility: v.DefaultVisibility != defaultVisibility,
	}
}

// setupPlan names the guild's plan as the viewer's entitlements read it; a failed read is
// logged and reads as no plan (the checklist never fails on it).
func (s *Service) setupPlan(ctx context.Context, guildID, viewerID int64) string {
	if s.raidGuildTools(ctx, guildID, viewerID) {
		return planRaidGuild
	}
	if s.officerTools(ctx, guildID, viewerID) {
		return planGuild
	}
	return planNone
}

func (s *Service) setupDiscord(ctx context.Context, guildID, viewerID int64) (SetupDiscord, error) {
	view, err := s.SettingsDiscordFor(ctx, guildID, viewerID)
	if err != nil || view == nil {
		return SetupDiscord{}, err
	}
	out := SetupDiscord{BotAvailable: view.BotAvailable, WebhookConnected: view.Connected}
	if b := view.Bot; b != nil && !b.RemovedInDiscord {
		out.Installed, out.ServerName, out.ChannelName = true, &b.ServerName, b.ChannelName
		out.RecruitChannelName, out.LinkedMembers, out.Members = b.RecruitChannelName, b.LinkedMembers, b.Members
	}
	return out, nil
}

func (s *Service) setupRecruiting(ctx context.Context, guildID int64) (SetupRecruiting, error) {
	f, err := s.Store.recruitFields(ctx, guildID)
	if err != nil {
		return SetupRecruiting{}, err
	}
	out := SetupRecruiting{Open: f.open, Pitch: f.pitch != "", Nights: f.nights, RoleSet: f.roleID != nil && *f.roleID != "",
		Wanted: make([]string, 0, len(f.wanted))}
	for _, w := range f.wanted {
		out.Wanted = append(out.Wanted, wantedName(w))
	}
	return out, nil
}

// wantedName is a wanted line in words: "Holy Priest", or the class alone.
func wantedName(w WantedEntry) string { return specClassName(w.SpecName, w.Class) }

func (s *Service) setupSchedule(ctx context.Context, guildID int64) (SetupSchedule, error) {
	f, err := s.Store.recruitFields(ctx, guildID)
	if err != nil {
		return SetupSchedule{}, err
	}
	out := SetupSchedule{Timezone: f.timezone}
	now := time.Now()
	if out.Planned, err = s.Store.upcomingRaids(ctx, guildID, now); err != nil {
		return SetupSchedule{}, err
	}
	next, found, err := s.Store.NextRaid(ctx, guildID, now)
	if err != nil {
		return SetupSchedule{}, err
	}
	if !found {
		return out, nil
	}
	out.Next = &SetupRaid{ID: next.ID, Title: next.Title, StartsAt: next.StartsAt}
	rows, err := s.Store.SchedulesOf(ctx, []int64{next.ID})
	if err != nil {
		return SetupSchedule{}, err
	}
	if row, ok := rows[next.ID]; ok {
		out.Next.PostedAt = row.PostedAt
	}
	return out, nil
}

// MarkSetup records an officer's mark on a step: "done" (confirmable steps only) or "skipped".
func (s *Store) MarkSetup(ctx context.Context, guildID, userID int64, step, state string) error {
	if !slices.Contains(SetupSteps, step) || (state != SetupDone && state != SetupSkipped) ||
		(state == SetupDone && !slices.Contains(confirmableSteps, step)) {
		return ErrSetupStep
	}
	_, err := s.Pool.Exec(ctx, `
		insert into guild_setup_marks (guild_id, step, state, set_by) values ($1, $2, $3, $4)
		on conflict (guild_id, step) do update set state = excluded.state, set_by = excluded.set_by, set_at = now()`,
		guildID, step, state, userID)
	if err != nil {
		return fmt.Errorf("guilds: mark setup: %w", err)
	}
	return nil
}

// ClearSetupMark takes a step's mark back.
func (s *Store) ClearSetupMark(ctx context.Context, guildID int64, step string) error {
	if !slices.Contains(SetupSteps, step) {
		return ErrSetupStep
	}
	if _, err := s.Pool.Exec(ctx, `delete from guild_setup_marks where guild_id = $1 and step = $2`, guildID, step); err != nil {
		return fmt.Errorf("guilds: clear setup mark: %w", err)
	}
	return nil
}

func (s *Store) setupMarks(ctx context.Context, guildID int64) (map[string]string, error) {
	rows, err := s.Pool.Query(ctx, `select step, state from guild_setup_marks where guild_id = $1`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: setup marks: %w", err)
	}
	defer rows.Close()
	out := map[string]string{}
	for rows.Next() {
		var step, state string
		if err := rows.Scan(&step, &state); err != nil {
			return nil, fmt.Errorf("guilds: setup marks: %w", err)
		}
		out[step] = state
	}
	return out, rows.Err()
}

// setupRanks is the guild's ranks as its characters carry them (the addon export's rank
// index), each with its count and first names, and how many accounts the site counts as
// officers (verified officers and the leader).
func (s *Store) setupRanks(ctx context.Context, guildID int64, maxRank int) (SetupRanks, error) {
	out := SetupRanks{OfficerMaxRankIndex: maxRank, Ranks: []RankLine{}}
	rows, err := s.Pool.Query(ctx, `
		select gc.rank_index, count(*)::int, (count(*) filter (where c.level >= $3))::int,
		       (array_agg(ae.name order by c.level >= $3 desc nulls last, ae.name)
		          filter (where ae.name is not null))[1:$2]
		from guild_characters gc
		left join addon_exports ae on ae.character_key = gc.character_key
		left join characters c on c.key = gc.character_key
		where gc.guild_id = $1 and gc.rank_index is not null and gc.character_key not like 'account:%'
		group by gc.rank_index order by gc.rank_index`, guildID, rankSampleNames, raidLevel)
	if err != nil {
		return SetupRanks{}, fmt.Errorf("guilds: setup ranks: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var line RankLine
		if err := rows.Scan(&line.RankIndex, &line.Members, &line.AtSixty, &line.Names); err != nil {
			return SetupRanks{}, fmt.Errorf("guilds: setup ranks: %w", err)
		}
		if line.Names == nil {
			line.Names = []string{}
		}
		out.Ranks = append(out.Ranks, line)
	}
	if err := rows.Err(); err != nil {
		return SetupRanks{}, fmt.Errorf("guilds: setup ranks: %w", err)
	}
	err = s.Pool.QueryRow(ctx, `select count(distinct user_id)::int from guild_characters
		where guild_id = $1 and verified_at is not null and rank in ('officer', 'leader')`, guildID).Scan(&out.Officers)
	if err != nil {
		return SetupRanks{}, fmt.Errorf("guilds: setup officers: %w", err)
	}
	return out, nil
}

// joinedByInvite counts the accounts that joined the guild by an invite link.
func (s *Store) joinedByInvite(ctx context.Context, guildID int64) (int, error) {
	var n int
	err := s.Pool.QueryRow(ctx, `select count(distinct user_id)::int from guild_characters
		where guild_id = $1 and source = 'invite'`, guildID).Scan(&n)
	if err != nil {
		return 0, fmt.Errorf("guilds: joined by invite: %w", err)
	}
	return n, nil
}

// setupSync counts the guild's verified accounts and those with an addon export inside the window.
func (s *Store) setupSync(ctx context.Context, guildID int64) (SetupSync, error) {
	out := SetupSync{WindowDays: int(syncWindow / (24 * time.Hour))}
	err := s.Pool.QueryRow(ctx, `
		select count(distinct gc.user_id)::int,
		       count(distinct gc.user_id) filter (where ae.updated_at >= now() - make_interval(secs => $2))::int
		from guild_characters gc left join addon_exports ae on ae.character_key = gc.character_key
		where gc.guild_id = $1 and gc.verified_at is not null`, guildID, syncWindow.Seconds()).
		Scan(&out.Accounts, &out.Synced)
	if err != nil {
		return SetupSync{}, fmt.Errorf("guilds: setup sync: %w", err)
	}
	out.Missing, err = s.unsyncedNames(ctx, guildID)
	return out, err
}

// unsyncedNames is one character name per verified account with no addon export inside the
// window, alphabetical, at most setupNamesShown. An invited account with no character of the
// guild's own has no name to show and is left out.
func (s *Store) unsyncedNames(ctx context.Context, guildID int64) ([]string, error) {
	rows, err := s.Pool.Query(ctx, `
		select initcap(split_part(min(gc.character_key), '/', 3))
		from guild_characters gc
		left join addon_exports ae on ae.character_key = gc.character_key
		  and ae.updated_at >= now() - make_interval(secs => $2)
		where gc.guild_id = $1 and gc.verified_at is not null and gc.character_key not like 'account:%'
		group by gc.user_id
		having count(ae.character_key) = 0
		order by 1 limit $3`, guildID, syncWindow.Seconds(), setupNamesShown)
	if err != nil {
		return nil, fmt.Errorf("guilds: unsynced names: %w", err)
	}
	names, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, fmt.Errorf("guilds: unsynced names: %w", err)
	}
	return names, nil
}

// upcomingRaids counts the guild's planned or locked raids still to start.
func (s *Store) upcomingRaids(ctx context.Context, guildID int64, now time.Time) (int, error) {
	var n int
	err := s.Pool.QueryRow(ctx, `select count(*)::int from guild_raids
		where guild_id = $1 and status in ('planned', 'locked') and starts_at > $2`, guildID, now).Scan(&n)
	if err != nil {
		return 0, fmt.Errorf("guilds: upcoming raids: %w", err)
	}
	return n, nil
}
