// api/cmd/seedguild/lifecycle.go
//
// `seedguild --lifecycle` builds, from nothing, the small world the guild claim and
// management workflows are walked in on a local stack (tools/dev-stack): an unclaimed
// guild "Lifecycle Test" on us/pvp, its guild master, an officer and three members on three
// accounts, a fourth account in no guild, the roster snapshot the addon's outbox would have
// sent, one synced export per character, and one complete report with two boss kills
// credited to the guild master.
//
// Unlike the mock-roster seed in main.go, which dresses up a guild that already exists, this
// one creates the guild too, and it goes through the same stores a real sync uses
// (auth.Store.UpsertBnetUser, addon.Store.PutExports, addon.Store.PutRosterSnapshot), so
// the guild, the guild_characters rows, the derived guild_members rows and the factions
// are exactly what production code would have written. Only the report is written
// directly, as the mock-roster seed does, because a real one needs a log in object storage.
//
// Every account it creates carries a bnet_sub, because claiming a guild as its master
// requires a linked Battle.net identity (guilds.ErrNoBattleNetIdentity); the subs are
// placeholders that no Battle.net account can hold.
package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

const (
	lifecycleTag       = "seedguild-lifecycle"
	lifecycleGuildName = "Lifecycle Test"
	lifecycleRegion    = "us"
	lifecycleRuleset   = "pvp"
)

// lifecycleCharacter is one seeded character. RankIndex is its raw guild rank (0 is the
// guild master, 1 an officer under the default threshold, 2 and up members); a character
// with InGuild false exports no guild= section at all, like an alt in no guild.
type lifecycleCharacter struct {
	Mock      mockCharacter
	InGuild   bool
	RankIndex int
}

// lifecycleAccount is one seeded account: the sub stands in for a Battle.net subject.
type lifecycleAccount struct {
	Battletag  string
	Sub        string
	Characters []lifecycleCharacter
}

// lifecycleAccounts is the whole cast. The battletags are what tools/dev-stack/lib.sh and
// the walkthrough doc sign in with; keep the three in step.
func lifecycleAccounts() []lifecycleAccount {
	return []lifecycleAccount{
		{
			Battletag: "hunthrop#1894", Sub: "dev-lifecycle-gm",
			Characters: []lifecycleCharacter{
				{Mock: mockCharacter{Name: "Bow Jackzon", Class: "hunter", Spec: "BeastMastery", Role: roleDPS,
					RaceSlug: "dwarf", BisFile: "hunter-beast-mastery.json", Consent: "gear", ItemLevel: 66},
					InGuild: true, RankIndex: 0},
				// An alt in no guild: a character page that is not a member's.
				{Mock: mockCharacter{Name: "Pipsqueak", Class: "priest", Spec: "Holy", Role: roleHealer,
					RaceSlug: "gnome", Consent: "gear", ItemLevel: 52}},
			},
		},
		{
			Battletag: "Ashenveil#2201", Sub: "dev-lifecycle-officer",
			Characters: []lifecycleCharacter{
				{Mock: mockCharacter{Name: "Ashenveil", Class: "paladin", Spec: "Holy", Role: roleHealer,
					RaceSlug: "human", Consent: "gear", ItemLevel: 63},
					InGuild: true, RankIndex: 1},
			},
		},
		{
			Battletag: "Quillfeather#3301", Sub: "dev-lifecycle-member",
			Characters: []lifecycleCharacter{
				{Mock: mockCharacter{Name: "Quillfeather", Class: "priest", Spec: "Holy", Role: roleHealer,
					RaceSlug: "gnome", Consent: "gear", ItemLevel: 60}, InGuild: true, RankIndex: 3},
				{Mock: mockCharacter{Name: "Rootbound", Class: "druid", Spec: "Restoration", Role: roleHealer,
					RaceSlug: "night-elf", Consent: "gear", ItemLevel: 58}, InGuild: true, RankIndex: 4},
				{Mock: mockCharacter{Name: "Cindergale", Class: "mage", Spec: "Fire", Role: roleDPS,
					RaceSlug: "gnome", BisFile: "mage-fire.json", Consent: "gear", ItemLevel: 61},
					InGuild: true, RankIndex: 5},
			},
		},
		{
			// A fourth account in no guild at all: the one that opens an invite link, and the
			// signed-in stranger the claim and guild pages turn away.
			Battletag: "Wanderer#4401", Sub: "dev-lifecycle-wanderer",
			Characters: []lifecycleCharacter{
				{Mock: mockCharacter{Name: "Driftwood", Class: "rogue", Spec: "Combat", Role: roleDPS,
					RaceSlug: "human", BisFile: "rogue-combat.json", Consent: "gear", ItemLevel: 57}},
			},
		},
	}
}

// lifecycleExtraRoster are the guild-tab lines the snapshot carries for people with no
// account on the site: the roster an officer sees before everyone has installed the addon.
var lifecycleExtraRoster = []struct {
	Name, Spec string
	Level      int
}{
	{"Strayhound", "Protection", 60},
	{"Oldmarrow", "Arms", 58},
}

// lifecycleReport is the one complete report: two Onyxia kills credited to the guild master
// (Onyxia is the only boss the game data publishes, so both kills are hers; see raid.go).
func lifecycleReport() reportPlan {
	return reportPlan{
		Tag: "ony", Label: "Onyxia", Zone: zoneOnyxia, DaysAgo: 1, StartHour: 19,
		Fights: []fightPlan{
			{Name: "Onyxian Whelp", DurationMS: 60_000, Kill: true, Trash: true},
			{Name: "Onyxia", EncounterID: encOnyxia, Kill: true, DurationMS: 240_000},
			{Name: "Onyxian Whelp", DurationMS: 55_000, Kill: true, Trash: true},
			{Name: "Onyxia", EncounterID: encOnyxia, Kill: true, DurationMS: 225_000},
		},
	}
}

// lifecycleExport is one character's FS1 export as the addon writes it: the head and the
// level, bags and professions sections the mock-roster seed builds, then the guild=
// section the API reads membership from.
func lifecycleExport(c lifecycleCharacter) (string, error) {
	dir, err := bisDir()
	if err != nil {
		return "", err
	}
	slots, err := loadBisSlots(dir, c.Mock.BisFile)
	if err != nil {
		return "", err
	}
	export := buildFS1(classSlugFor(c.Mock.Class), c.Mock.RaceSlug, talentStringFor(c.Mock),
		gearWithReadiness(c.Mock, slots), professionsFor(c.Mock.Class), bagsFor(c.Mock))
	if c.InGuild {
		export += fmt.Sprintf("|guild=%s:%d", url.PathEscape(lifecycleGuildName), c.RankIndex)
	}
	return export, nil
}

func lifecycleKey(name string) string {
	return character.Key(lifecycleRegion, lifecycleRuleset, name)
}

// ApplyLifecycle writes the cast. It is a no-op, after a notice, when the lifecycle tag
// already has rows, so a second run neither duplicates nor half-reapplies anything.
func ApplyLifecycle(ctx context.Context, pool *pgxpool.Pool) error {
	applied, err := alreadyApplied(ctx, pool, lifecycleTag)
	if err != nil {
		return err
	}
	if applied {
		fmt.Printf("seedguild: %q is already seeded; nothing to do (--lifecycle --remove first to reseed)\n", lifecycleTag)
		return nil
	}

	authStore := &auth.Store{Pool: pool}
	addonStore := &addon.Store{Pool: pool}
	userIDs := map[string]int64{}
	var members []plannedCharacter

	for _, acct := range lifecycleAccounts() {
		u, err := authStore.UpsertBnetUser(ctx, acct.Sub, acct.Battletag)
		if err != nil {
			return fmt.Errorf("seedguild: create %s: %w", acct.Battletag, err)
		}
		userIDs[acct.Battletag] = u.ID
		if err := recordLifecycleRow(ctx, pool, "users", fmt.Sprint(u.ID)); err != nil {
			return err
		}

		exports := make([]addon.Export, 0, len(acct.Characters))
		for _, c := range acct.Characters {
			export, err := lifecycleExport(c)
			if err != nil {
				return err
			}
			exports = append(exports, addon.Export{
				Name: c.Mock.Name, Ruleset: lifecycleRuleset, Region: lifecycleRegion, Export: export,
			})
			if c.InGuild {
				members = append(members, plannedCharacter{
					Mock: c.Mock, Key: lifecycleKey(c.Mock.Name), Battletag: acct.Battletag,
				})
			}
		}
		// The same call the companion's upload makes: it stores the exports, creates the
		// characters and the guild, and syncs guild_characters and guild_members.
		if err := addonStore.PutExports(ctx, u.ID, exports); err != nil {
			return fmt.Errorf("seedguild: sync exports for %s: %w", acct.Battletag, err)
		}
	}

	guildID, err := lifecycleGuildID(ctx, pool)
	if err != nil {
		return err
	}
	if err := recordLifecycleRow(ctx, pool, "guilds", fmt.Sprint(guildID)); err != nil {
		return err
	}

	// An export carries no level or faction, and the API fills both from a Battle.net
	// import this stack does not have; the mock-roster seed sets the same two columns.
	if _, err := pool.Exec(ctx,
		`update characters set level = $2, faction = 'alliance' where user_id = any($1)`,
		allUserIDs(userIDs), fs1Level); err != nil {
		return fmt.Errorf("seedguild: set levels: %w", err)
	}
	if err := verifyMasterLikeBattleNet(ctx, pool, guildID, userIDs["hunthrop#1894"]); err != nil {
		return err
	}
	if err := putLifecycleSnapshot(ctx, addonStore, userIDs["hunthrop#1894"]); err != nil {
		return err
	}
	if err := writeLifecycleReport(ctx, pool, guildID, userIDs["hunthrop#1894"], members); err != nil {
		return err
	}

	fmt.Printf("seedguild: seeded %q: guild %d (%s, %s/%s), unclaimed\n",
		lifecycleTag, guildID, lifecycleGuildName, lifecycleRegion, lifecycleRuleset)
	for _, acct := range lifecycleAccounts() {
		for _, c := range acct.Characters {
			where := "in no guild"
			if c.InGuild {
				where = fmt.Sprintf("rank index %d", c.RankIndex)
			}
			fmt.Printf("  %-18s %-14s %s (%s)\n", acct.Battletag, c.Mock.Name, lifecycleKey(c.Mock.Name), where)
		}
	}
	return nil
}

func allUserIDs(m map[string]int64) []int64 {
	out := make([]int64, 0, len(m))
	for _, id := range m {
		out = append(out, id)
	}
	return out
}

func lifecycleGuildID(ctx context.Context, pool *pgxpool.Pool) (int64, error) {
	var id int64
	err := pool.QueryRow(ctx,
		`select id from guilds where region = $1 and ruleset = $2 and lower(name) = lower($3)`,
		lifecycleRegion, lifecycleRuleset, lifecycleGuildName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("seedguild: the guild %q was not created by the exports: %w", lifecycleGuildName, err)
	}
	return id, nil
}

func recordLifecycleRow(ctx context.Context, pool *pgxpool.Pool, table, rowKey string) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: record %s/%s: begin: %w", table, rowKey, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := recordRow(ctx, tx, lifecycleTag, table, rowKey, nil); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// verifyMasterLikeBattleNet marks the guild master's guild_characters row verified by
// Battle.net, which is what a real master's sign-in import leaves behind
// (bnetimport.ImportAccount, "Blizzard wins on verification"). Without it the master would
// be a plain unverified export, the site would show them as a member, and the claim page
// would refuse them ("you are a member of this guild, so you cannot claim it"): the
// walkthrough starts at a master who can claim, so it seeds the state the import produces.
// The officer and the members stay unverified, as exports alone leave them.
func verifyMasterLikeBattleNet(ctx context.Context, pool *pgxpool.Pool, guildID, masterID int64) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: verify master: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`update guild_characters set source = 'bnet', verified_by = 'bnet', verified_at = now()
		 where guild_id = $1 and user_id = $2`, guildID, masterID); err != nil {
		return fmt.Errorf("seedguild: verify master: %w", err)
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, masterID); err != nil {
		return fmt.Errorf("seedguild: verify master: recompute membership: %w", err)
	}
	return tx.Commit(ctx)
}

// putLifecycleSnapshot stores the roster the guild master's addon would have sent: every
// seeded guild member the way the addon sees them on the guild tab, plus two people with no
// account on the site.
func putLifecycleSnapshot(ctx context.Context, addonStore *addon.Store, gmUserID int64) error {
	seen := time.Now().UTC().Add(-2 * time.Hour).Format(sims.RequestedAtLayout)
	var lines []addon.RosterMember
	for _, acct := range lifecycleAccounts() {
		for _, c := range acct.Characters {
			if c.InGuild {
				lines = append(lines, addon.RosterMember{
					Name: c.Mock.Name, Level: fs1Level, Spec: c.Mock.Spec, SeenAt: seen,
				})
			}
		}
	}
	for _, x := range lifecycleExtraRoster {
		lines = append(lines, addon.RosterMember{Name: x.Name, Level: x.Level, Spec: x.Spec, SeenAt: seen})
	}
	stored, err := addonStore.PutRosterSnapshot(ctx, gmUserID, lifecycleKey("Bow Jackzon"),
		addon.RosterSnapshot{Guild: lifecycleGuildName, Members: lines})
	if err != nil {
		return err
	}
	if !stored {
		return fmt.Errorf("seedguild: the roster snapshot for %s was not stored", lifecycleGuildName)
	}
	return nil
}

// writeLifecycleReport writes the one complete report, owned by the guild master and
// logged by the guild master's character, with a fight_metrics row for each member present.
func writeLifecycleReport(ctx context.Context, pool *pgxpool.Pool, guildID, gmUserID int64, present []plannedCharacter) error {
	plan := lifecycleReport()
	now := time.Now().UTC()
	createdAt, completedAt, starts := plan.schedule(now)
	pr := plannedReport{
		Plan: plan, ID: lifecycleTag + "-report-" + plan.Tag, Title: plan.titleFor(createdAt),
		CreatedAt: createdAt, CompletedAt: completedAt, FightStarts: starts,
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seedguild: report %s: begin: %w", pr.ID, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := recordRow(ctx, tx, lifecycleTag, "reports", pr.ID, nil); err != nil {
		return err
	}
	if err := insertReport(ctx, tx, guildID, gmUserID, lifecycleKey("Bow Jackzon"), pr); err != nil {
		return err
	}
	if err := insertFightsAndMetrics(ctx, tx, pool, pr, present); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// RemoveLifecycle deletes everything ApplyLifecycle wrote: the report, the accounts (and
// with them their characters, exports and memberships) and the guild (and with it the
// roster snapshot). A no-op, after a notice, when nothing is recorded.
func RemoveLifecycle(ctx context.Context, pool *pgxpool.Pool) error {
	applied, err := alreadyApplied(ctx, pool, lifecycleTag)
	if err != nil {
		return err
	}
	if !applied {
		fmt.Printf("seedguild: %q has no recorded rows; nothing to remove\n", lifecycleTag)
		return nil
	}
	reportRows, err := readSeedRows(ctx, pool, lifecycleTag, "reports")
	if err != nil {
		return err
	}
	for _, r := range reportRows {
		if err := removeReport(ctx, pool, lifecycleTag, r.RowKey); err != nil {
			return err
		}
	}
	userRows, err := readSeedRows(ctx, pool, lifecycleTag, "users")
	if err != nil {
		return err
	}
	for _, r := range userRows {
		if err := removeMockUser(ctx, pool, lifecycleTag, r.RowKey); err != nil {
			return err
		}
	}
	guildRows, err := readSeedRows(ctx, pool, lifecycleTag, "guilds")
	if err != nil {
		return err
	}
	for _, r := range guildRows {
		tx, err := pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("seedguild: remove guild %s: begin: %w", r.RowKey, err)
		}
		if err := removeLifecycleGuild(ctx, tx, r.RowKey); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("seedguild: remove guild %s: commit: %w", r.RowKey, err)
		}
	}
	fmt.Printf("seedguild: removed %q: %d reports, %d accounts, %d guild\n",
		lifecycleTag, len(reportRows), len(userRows), len(guildRows))
	return nil
}

func removeLifecycleGuild(ctx context.Context, tx pgx.Tx, guildID string) error {
	if _, err := tx.Exec(ctx, `delete from guilds where id = $1::bigint`, guildID); err != nil {
		return fmt.Errorf("seedguild: remove guild %s: %w", guildID, err)
	}
	if _, err := tx.Exec(ctx,
		`delete from seed_rows where tag = $1 and table_name = 'guilds' and row_key = $2`, lifecycleTag, guildID); err != nil {
		return fmt.Errorf("seedguild: unrecord guild %s: %w", guildID, err)
	}
	return nil
}
