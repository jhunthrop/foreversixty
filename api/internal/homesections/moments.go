package homesections

import (
	"slices"
	"time"
)

// raidLiveFor is how long after its start a raid counts as in progress (flows spec section 4).
const raidLiveFor = 4 * time.Hour

// View is what the signed-in composition reads about the viewer's character and its guild.
type View struct {
	Level int
	// InGuild: the character is a verified member of a guild.
	InGuild bool
	// Now is the time the home is composed for.
	Now time.Time
	// Zone is the guild's raid time zone, which decides what "today" and "yesterday" are; nil is UTC.
	Zone *time.Location
	// NextRaid is the start of the guild's next raid that is not over (it may have started); nil for none.
	NextRaid *time.Time
	// LastRaid is the start of the guild's most recent raid that has started; nil for none.
	LastRaid *time.Time
}

// Moment is which of section 4's Home states the viewer is in.
type Moment string

// The Home states, in the order they are tried.
const (
	MomentRaidLive Moment = "raid_live"
	MomentRaidDay  Moment = "raid_day"
	MomentDayAfter Moment = "day_after"
	MomentLeveling Moment = "leveling"
	MomentGuild    Moment = "guild"
	MomentNoGuild  Moment = "no_guild"
)

// MomentOf is the viewer's Home state under o: the raid moments first (a guild member only), then
// leveling under 60, then the guild or no-guild day.
func MomentOf(v View, o Override) Moment {
	if v.InGuild && o.ruleOn(RuleRaidMoments) {
		switch {
		case v.NextRaid != nil && !v.Now.Before(*v.NextRaid) && v.Now.Before(v.NextRaid.Add(raidLiveFor)):
			return MomentRaidLive
		case v.NextRaid != nil && v.Now.Before(*v.NextRaid) && sameDay(*v.NextRaid, v.Now, v.zone()):
			return MomentRaidDay
		case v.LastRaid != nil && sameDay(v.LastRaid.AddDate(0, 0, 1), v.Now, v.zone()):
			return MomentDayAfter
		}
	}
	switch {
	case v.Level < maxLevel && o.ruleOn(RuleLeveling):
		return MomentLeveling
	case v.InGuild:
		return MomentGuild
	default:
		return MomentNoGuild
	}
}

func (v View) zone() *time.Location {
	if v.Zone == nil {
		return time.UTC
	}
	return v.Zone
}

// sameDay reports whether a and b fall on one calendar date in loc.
func sameDay(a, b time.Time, loc *time.Location) bool {
	return a.In(loc).Format(time.DateOnly) == b.In(loc).Format(time.DateOnly)
}

// SignedIn is the signed-in home for v under o (flows spec section 4): the moment's block, then the
// next upgrades, then the three newest Activity rows; a player at 60 with no guild also gets Find a
// guild. Everything else lives on its tab.
func SignedIn(v View, o Override) []Section {
	head := Section{Type: Head}
	upgrades := Section{Type: Upgrades, Limit: upgradesShown}
	activity := Section{Type: Activity, Limit: activityShown}

	var out []Section
	switch MomentOf(v, o) {
	case MomentRaidLive:
		raidActivity := activity
		raidActivity.Group = "raids"
		out = []Section{head, {Type: RaidLive}, raidActivity}
	case MomentRaidDay:
		out = []Section{head, {Type: NextRaid, Style: StyleCard}, upgrades, activity}
	case MomentDayAfter:
		out = []Section{head, {Type: RaidSummary}, upgrades, activity}
	case MomentGuild:
		out = []Section{head, upgrades, {Type: NextRaid, Style: StyleLine}, activity}
	case MomentLeveling:
		out = []Section{head, upgrades, activity}
	default:
		out = []Section{head, upgrades, activity}
		if o.ruleOn(RuleFindAGuild) {
			out = append(out, Section{Type: GuildCTA})
		}
	}
	if o.Notice != nil {
		out = slices.Insert(out, 1, Section{Type: Notice, Text: o.Notice.Text, Href: o.Notice.Href})
	}
	return o.apply(out)
}
