// api/internal/guilds/member_events.go
//
// The member events of the notification catalogue that the guild writes produce:
// guild.approved when an officer approves a character, loot.awarded when an officer awards
// an item, and report.ready to the raiders when a log is matched to the night's raid.
package guilds

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// memberEventCopy is the words of the member notifications.
var memberEventCopy = struct {
	ApprovedTitle, ApprovedBody, LootTitle, LootBody, LootBodyNamed, NightTitle, NightBody string
}{
	ApprovedTitle: "Approved to %s",
	ApprovedBody:  "An officer approved %s to the roster.",
	LootTitle:     "Loot awarded",
	LootBody:      "%s was awarded loot by %s.",
	LootBodyNamed: "%s was awarded %s by %s.",
	NightTitle:    "%s: the night's report is ready",
	NightBody:     "See the kills, attendance and loot.",
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

// nightReadyEvents tells the raid's In and Late players the night's report is ready.
func (s *Service) nightReadyEvents(ctx context.Context, raidID int64, reportID string) []events.Event {
	facts, ok := s.raidFacts(ctx, raidID, "event_night_ready")
	if !ok {
		return nil
	}
	p := facts.Post.RaidPosted
	e := raidEvent(events.ReportReady, facts.GuildID, p, "report.ready:"+reportID,
		events.Audience{Kind: events.ToRaidInLate, GuildID: facts.GuildID, RaidID: raidID})
	e.Title = fmt.Sprintf(memberEventCopy.NightTitle, p.Title)
	e.Body = memberEventCopy.NightBody
	e.Data["report_id"] = reportID
	return []events.Event{e}
}
