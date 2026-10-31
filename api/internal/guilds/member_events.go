// api/internal/guilds/member_events.go
//
// The member events of the notification catalogue that the guild writes produce:
// guild.approved when an officer approves a character, loot.awarded when an officer awards
// an item (off by default), and raid.summary to the raiders when a log is matched to the
// night's raid. The summary carries the night's report.ready for the raiders, so no separate
// night report.ready is sent; an upload's own report.ready (notify) is unchanged.
package guilds

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// memberEventCopy is the words of the member notifications.
var memberEventCopy = struct {
	ApprovedTitle, ApprovedBody, LootTitle, LootBody, LootBodyNamed string
}{
	ApprovedTitle: "Approved to %s",
	ApprovedBody:  "An officer approved %s to the roster.",
	LootTitle:     "Loot awarded",
	LootBody:      "%s was awarded loot by %s.",
	LootBodyNamed: "%s was awarded %s by %s.",
}

// rosterCharacter is one roster row's owner and display name.
type rosterCharacter struct {
	Key    string
	UserID int64
	Name   string
	Class  string
}

// rosterCharacters reads the owner, name and class of each of keys on the guild's roster.
func (s *Store) rosterCharacters(ctx context.Context, guildID int64, keys []string) ([]rosterCharacter, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.character_key, gc.user_id, coalesce(c.name, ''), coalesce(c.class, '')
		from guild_characters gc left join characters c on c.key = gc.character_key
		where gc.guild_id = $1 and gc.character_key = any($2) order by gc.character_key`, guildID, keys)
	if err != nil {
		return nil, fmt.Errorf("guilds: roster characters: %w", err)
	}
	defer rows.Close()
	var out []rosterCharacter
	for rows.Next() {
		var c rosterCharacter
		if err := rows.Scan(&c.Key, &c.UserID, &c.Name, &c.Class); err != nil {
			return nil, fmt.Errorf("guilds: roster characters: %w", err)
		}
		if c.Name == "" {
			c.Name = nameOfKey(c.Key)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// nameOfKey is a character key's name segment, capitalised ("us/normal/kraggor": Kraggor).
func nameOfKey(key string) string {
	name := key[strings.LastIndex(key, "/")+1:]
	if name == "" {
		return key
	}
	return strings.ToUpper(name[:1]) + name[1:]
}

// guildBase is what every member event of a guild shares: its page and ids.
func (s *Service) guildBase(g Guild) events.Event {
	return events.Event{
		Link: strings.TrimRight(s.PublicBaseURL, "/") + GuildPagePath(g),
		Data: map[string]string{"guild_id": strconv.FormatInt(g.ID, 10)},
	}
}

// publishApproved tells each approved character's account it is on the roster.
func (s *Service) publishApproved(ctx context.Context, guildID int64, keys []string) {
	if len(keys) == 0 {
		return
	}
	s.background(ctx, func(ctx context.Context) []events.Event {
		g, err := s.Store.getGuild(ctx, guildID)
		if err != nil {
			s.logger().Error("guilds", "op", "event_approved", "guild_id", guildID, "err", err)
			return nil
		}
		chars, err := s.Store.rosterCharacters(ctx, guildID, keys)
		if err != nil {
			s.logger().Error("guilds", "op", "event_approved", "guild_id", guildID, "err", err)
			return nil
		}
		out := make([]events.Event, 0, len(chars))
		for _, c := range chars {
			e := s.guildBase(g)
			e.Kind, e.Audience = events.GuildApproved, events.Users(c.UserID)
			e.DedupeKey = fmt.Sprintf("guild.approved:%d:%s", guildID, c.Key)
			e.Title = fmt.Sprintf(memberEventCopy.ApprovedTitle, g.Name)
			e.Body = fmt.Sprintf(memberEventCopy.ApprovedBody, c.Name)
			out = append(out, e)
		}
		return out
	})
}

// publishLootAwarded tells the account of the character an award went to. An award to an
// invited member's stand-in row names no roster character, so it tells nobody.
func (s *Service) publishLootAwarded(ctx context.Context, guildID, awardID int64, characterKey string, itemID int) {
	s.background(ctx, func(ctx context.Context) []events.Event {
		g, err := s.Store.getGuild(ctx, guildID)
		if err != nil {
			s.logger().Error("guilds", "op", "event_loot", "guild_id", guildID, "err", err)
			return nil
		}
		chars, err := s.Store.rosterCharacters(ctx, guildID, []string{characterKey})
		if err != nil || len(chars) == 0 {
			if err != nil {
				s.logger().Error("guilds", "op", "event_loot", "guild_id", guildID, "err", err)
			}
			return nil
		}
		c := chars[0]
		e := s.guildBase(g)
		e.Kind, e.Audience = events.LootAwarded, events.Users(c.UserID)
		e.DedupeKey = fmt.Sprintf("loot.awarded:%d", awardID)
		e.Title = memberEventCopy.LootTitle
		e.Body = fmt.Sprintf(memberEventCopy.LootBody, c.Name, g.Name)
		if item, ok := s.Store.ItemIn(s.Store.BisBuild(), c.Class, itemID); ok && item.Name != "" {
			e.Body = fmt.Sprintf(memberEventCopy.LootBodyNamed, c.Name, item.Name, g.Name)
		}
		e.Data["item_id"] = strconv.Itoa(itemID)
		return []events.Event{e}
	})
}

// summaryCopy is raid.summary's words: one line per thing the night left the player.
var summaryCopy = struct{ Title, Report, Coach, Loot, LootUnnamed string }{
	Title:       "%s: the night's summary",
	Report:      "The report is ready.",
	Coach:       "Your coach notes are ready.",
	Loot:        "You won %s.",
	LootUnnamed: "You won loot.",
}

// nightLoot is one award of a night's log to an account.
type nightLoot struct {
	UserID int64
	ItemID int
	Class  string
}

// summaryEvents is raid.summary for each In and Late account of the raid a log was matched to
// (where report.ready fires): the report, the coach notes when the account holds the log
// coach, and the loot awarded to it. Personal bests have no source yet, so none are listed.
func (s *Service) summaryEvents(ctx context.Context, raidID int64, reportID string) []events.Event {
	facts, ok := s.raidFacts(ctx, raidID, "event_raid_summary")
	if !ok {
		return nil
	}
	places, err := s.Store.accountPlaces(ctx, raidID)
	if err != nil {
		s.logger().Error("guilds", "op", "event_raid_summary", "raid_id", raidID, "err", err)
		return nil
	}
	loot, err := s.Store.nightLoot(ctx, facts.GuildID, reportID)
	if err != nil {
		s.logger().Error("guilds", "op", "event_raid_summary", "raid_id", raidID, "err", err)
	}
	var out []events.Event
	for uid, state := range places {
		if state == SignupIn || state == SignupLate {
			out = append(out, s.summaryFor(ctx, facts, uid, reportID, lootOf(loot, uid)))
		}
	}
	return out
}

func lootOf(all []nightLoot, uid int64) []nightLoot {
	var out []nightLoot
	for _, l := range all {
		if l.UserID == uid {
			out = append(out, l)
		}
	}
	return out
}

// summaryFor is one account's raid.summary.
func (s *Service) summaryFor(ctx context.Context, facts RaidPostFacts, uid int64, reportID string, loot []nightLoot) events.Event {
	p := facts.Post.RaidPosted
	e := raidEvent(events.RaidSummary, facts.GuildID, p, "raid.summary:"+reportID, events.Users(uid))
	e.Title = fmt.Sprintf(summaryCopy.Title, p.Title)
	lines := []string{summaryCopy.Report}
	e.Data["report_id"] = reportID
	e.Data["coach_ready"] = "false"
	if s.holdsCoach(ctx, facts.GuildID, uid) {
		lines = append(lines, summaryCopy.Coach)
		e.Data["coach_ready"] = "true"
	}
	if len(loot) > 0 {
		names, ids := make([]string, 0, len(loot)), make([]string, 0, len(loot))
		for _, l := range loot {
			ids = append(ids, strconv.Itoa(l.ItemID))
			if item, ok := s.Store.ItemIn(s.Store.BisBuild(), l.Class, l.ItemID); ok && item.Name != "" {
				names = append(names, item.Name)
			}
		}
		line := summaryCopy.LootUnnamed
		if len(names) == len(loot) {
			line = fmt.Sprintf(summaryCopy.Loot, strings.Join(names, ", "))
		}
		lines = append(lines, line)
		e.Data["loot_item_ids"] = strings.Join(ids, ",")
	}
	e.Body = strings.Join(lines, " ")
	return e
}

// holdsCoach is whether the account holds the log coach in this guild; a failed ask is false.
func (s *Service) holdsCoach(ctx context.Context, guildID, uid int64) bool {
	if s.Entitlements == nil {
		return false
	}
	ok, _, err := s.Entitlements.CanForGuild(ctx, guildID, uid, entitlements.FeatureLogCoach)
	if err != nil {
		s.logger().Warn("guilds", "op", "event_raid_summary_coach", "guild_id", guildID, "err", err)
		return false
	}
	return ok
}

// nightLoot is each award of the guild's attached to the report, with its winner's account.
func (s *Store) nightLoot(ctx context.Context, guildID int64, reportID string) ([]nightLoot, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.user_id, la.item_id, coalesce(c.class, '')
		from loot_awards la
		join guild_characters gc on gc.guild_id = la.guild_id and gc.character_key = la.character_key
		left join characters c on c.key = la.character_key
		where la.guild_id = $1 and la.report_id = $2 order by la.id`, guildID, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: night loot: %w", err)
	}
	return pgx.CollectRows(rows, pgx.RowToStructByPos[nightLoot])
}
