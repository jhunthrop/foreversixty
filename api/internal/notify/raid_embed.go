package notify

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
)

// The raid post's stripe besides the accent: a cancelled raid in the wipe red, a locked or ended
// one in the card's muted ink.
var (
	cancelColor = brand.RGB(brand.ColorOutcomeWipe)
	mutedColor  = brand.RGB(brand.ColorForeverMutedCard)
)

const (
	// maxRaidTitleRunes is the raid form's title bound.
	maxRaidTitleRunes = 80
)

// RaidPostState is which of the raid post's looks a message wears.
type RaidPostState string

// The raid post's states (spec 2026-10-27-discord-bot §3.3.1).
const (
	PostOpen      RaidPostState = "open"
	PostLocked    RaidPostState = "locked"
	PostMoved     RaidPostState = "moved"
	PostCancelled RaidPostState = "cancelled"
	PostEnded     RaidPostState = "ended"
	PostDone      RaidPostState = "done"
	PostWebhook   RaidPostState = "webhook"
	// PostTomorrow is the post on the day-before reminder (spec 2026-10-28 §P.4).
	PostTomorrow RaidPostState = "tomorrow"
)

// raidCopy is the words of the raid posts (spec 2026-10-27-discord-bot §6.2).
var raidCopy = struct {
	TitleFormat     string // %s the escaped title, %s the guild's date (markup, no words)
	MovedPrefix     string
	CancelledPrefix string
	DonePrefix      string
	TomorrowPrefix  string
	Descriptions    map[RaidPostState]string
	WhenName        string
	WhenFormat      string // %d the unix start, twice
	WasName         string
	StruckFormat    string // %s a struck-through value
	SignedName      string
	SignedLead      string // {signed} signed up, {size} the raid size (roster.signed)
	SignedLeadOpen  string // {signed} signed up, size unknown
	SignedNone      string
	CameName        string
	CameOf          string // {came} came, {of} of
	KillsName       string
	BuffsName       string
	BuffsMissing    string // {buffs} the buffs nobody signed brings
	BuffsShort      string // {buffs} the buffs with too few bringers
	BuffsJoin       string
	// The reminder line (spec 2026-10-28 §P.7).
	ReminderLead        string // {title} the escaped title, {unix} the unix start
	ReminderCount       string // {signed} signed up, {size} the raid size
	ReminderCountOpen   string // {signed} signed up, size unknown
	ReminderNobody      string
	ReminderWebhookLink string // {url} the raid page
}{
	TitleFormat:     "%s · %s",
	MovedPrefix:     copygen.DiscordRaidMovedPrefix,
	CancelledPrefix: copygen.DiscordRaidCancelledPrefix,
	DonePrefix:      copygen.DiscordRaidDonePrefix,
	TomorrowPrefix:  copygen.DiscordRaidTomorrowPrefix,
	Descriptions: map[RaidPostState]string{
		PostOpen:      "",
		PostLocked:    copygen.RaidLocked,
		PostMoved:     copygen.DiscordRaidDescriptionsPostMoved,
		PostCancelled: copygen.RaidCancelled,
		PostEnded:     copygen.RaidEnded,
		PostDone:      copygen.DiscordRaidDescriptionsPostDone,
		PostWebhook:   copygen.DiscordRaidDescriptionsPostWebhook,
		PostTomorrow:  "",
	},
	WhenName:       copygen.NotificationsWhen,
	WhenFormat:     "<t:%d:t> (<t:%d:R>)",
	WasName:        copygen.DiscordRaidWasName,
	StruckFormat:   "~~%s~~",
	SignedName:     copygen.DiscordRaidSignedName,
	SignedLead:     copygen.DiscordRaidSignedLead,
	SignedLeadOpen: copygen.DiscordRaidSignedLeadOpen,
	SignedNone:     copygen.DiscordRaidSignedNone,
	CameName:       copygen.DiscordRaidCameName,
	CameOf:         copygen.DiscordRaidCameOf,
	KillsName:      copygen.DiscordRaidKillsName,
	BuffsName:      copygen.DiscordRaidBuffsName,
	BuffsMissing:   copygen.DiscordRaidBuffsMissing,
	BuffsShort:     copygen.DiscordRaidBuffsShort,
	BuffsJoin:      ", ",

	ReminderLead:        copygen.DiscordRaidReminderLead,
	ReminderCount:       copygen.DiscordRaidReminderCount,
	ReminderCountOpen:   copygen.DiscordRaidReminderCountOpen,
	ReminderNobody:      copygen.DiscordRaidReminderNobody,
	ReminderWebhookLink: copygen.DiscordRaidReminderWebhookLink,
}

// RaidCounts are a raid's sign-up totals; the roles count In rows only.
type RaidCounts struct {
	In, Late, Bench, Out, Tanks, Healers, DPS int
}

// RaidPosted is what the raid create holds for one new raid. Counts are those at posting.
type RaidPosted struct {
	RaidID int64
	// Zone is the raid's zone key ("raid:onyxias-lair"); its icon is the raid events' Activity icon.
	Zone             string
	GuildName, Title string
	StartsAt         time.Time
	// Timezone is the guild's IANA zone; empty reads as UTC and shows no Guild time.
	Timezone string
	// RaidURL is the raid page; CrestURL the guild's crest_url (a thumbnail only when absolute).
	RaidURL, CrestURL string
	Counts            RaidCounts
	Size              int
	SizeKnown         bool
	// Roster is every sign-up as the post lists it (spec 2026-10-28 §P.9).
	Roster []RosterLine
	// NoRoles is a guild event (events spec §3.5): no role groups and no specs, one list of
	// names and classes grouped In, Late, Bench, Out.
	NoRoles bool
	// Target is the raid's comp target per role group; nil for a raid with none. The role
	// headings read it as "Healers · 6 of 8".
	Target *RoleTarget
	// Missing and Short are the raid buffs and debuffs nobody signed can bring, and those with
	// fewer bringers than the raid needs (the raid page's cover), by name.
	Missing, Short []string
}

// RoleTarget is a comp target: how many of each role group the raid wants.
type RoleTarget struct {
	Tanks, Healers, Melee, Ranged int
}

// of is the target for the roster group g (raidspec.Group).
func (t RoleTarget) of(group string) int {
	switch group {
	case raidspec.GroupTank:
		return t.Tanks
	case raidspec.GroupHealer:
		return t.Healers
	case raidspec.GroupMelee:
		return t.Melee
	}
	return t.Ranged
}

// RaidChanged is a raid moved (Was set) or cancelled; StartsAt is the current start.
type RaidChanged struct {
	RaidPosted
	Was       *time.Time
	Cancelled bool
}

// KillLine is one boss killed on the night.
type KillLine = ReportKill

// RaidPost is one raid post in a state: the bot's edited-for-life message, or a webhook post.
type RaidPost struct {
	RaidPosted
	State RaidPostState
	// Was is the old start of a moved raid.
	Was *time.Time
	// Came of CameOf, and the kills, are a done raid's.
	Came, CameOf int
	Kills        []KillLine
	// Marks are the bot's spec emoji; nil (every webhook post) draws the roster as text lines.
	Marks Marks
}

// RaidPostedEmbed is the raid_posted message (the webhook post, never edited).
func RaidPostedEmbed(in RaidPosted, now time.Time) Embed {
	return RaidPostEmbed(RaidPost{RaidPosted: in, State: PostWebhook}, now)
}

// RaidChangedEmbed is the raid_changed message: cancelled, or moved from Was.
func RaidChangedEmbed(in RaidChanged, now time.Time) Embed {
	state := PostMoved
	if in.Cancelled {
		state = PostCancelled
	}
	return RaidPostEmbed(RaidPost{RaidPosted: in.RaidPosted, State: state, Was: in.Was}, now)
}

// RaidPostEmbed is the one builder of every raid post, bot or webhook.
func RaidPostEmbed(p RaidPost, now time.Time) Embed {
	e := raidBaseEmbed(p.RaidPosted, statePrefix(p.State), now)
	e.Color = stateColor(p.State)
	e.Description = raidCopy.Descriptions[p.State]
	e.Fields = postFields(p)
	if hasRoster(p.State) {
		e.Fields = append(e.Fields, rosterFields(p.Roster, p.Marks, p.Target, embedChars(e), p.NoRoles)...)
	}
	return e
}

// hasRoster reports the states that list the roster: every one but Cancelled and Done (§P.9.6).
func hasRoster(s RaidPostState) bool { return s != PostCancelled && s != PostDone }

// embedChars is what Discord counts toward an embed's 6000.
func embedChars(e Embed) int {
	return utf8.RuneCountInString(e.Title+e.Description+e.Footer+e.Author) + fieldsChars(e.Fields)
}

func statePrefix(s RaidPostState) string {
	switch s {
	case PostMoved:
		return raidCopy.MovedPrefix
	case PostCancelled:
		return raidCopy.CancelledPrefix
	case PostDone:
		return raidCopy.DonePrefix
	case PostTomorrow:
		return raidCopy.TomorrowPrefix
	}
	return ""
}

func stateColor(s RaidPostState) int {
	switch s {
	case PostLocked, PostEnded:
		return mutedColor
	case PostCancelled:
		return cancelColor
	case PostDone:
		return reportColor
	}
	return embedColor
}

func postFields(p RaidPost) []EmbedField {
	in := p.RaidPosted
	when := whenValue(in.StartsAt)
	switch p.State {
	case PostCancelled:
		return []EmbedField{raidField(raidCopy.WhenName, fmt.Sprintf(raidCopy.StruckFormat, when)), signedField(in)}
	case PostDone:
		return []EmbedField{
			raidField(raidCopy.WhenName, when),
			raidField(raidCopy.CameName, copygen.Fill(raidCopy.CameOf, "{came}", strconv.Itoa(p.Came), "{of}", strconv.Itoa(p.CameOf))),
			{Name: raidCopy.KillsName, Value: killsValue(p.Kills)},
		}
	}
	fields := []EmbedField{raidField(raidCopy.WhenName, when)}
	if p.State == PostMoved && p.Was != nil {
		fields = append(fields, raidField(raidCopy.WasName, textfmt.RaidWhen(*p.Was, textfmt.Location(in.Timezone))))
	}
	fields = append(fields, signedField(in))
	if buffs := buffsValue(in); buffs != "" {
		fields = append(fields, EmbedField{Name: raidCopy.BuffsName, Value: buffs})
	}
	return fields
}

// buffsValue is the cover the officer counts before the pull: what nobody signed brings, then
// what too few bring; "" when every buff and debuff is covered.
func buffsValue(in RaidPosted) string {
	var lines []string
	if len(in.Missing) > 0 {
		lines = append(lines, copygen.Fill(raidCopy.BuffsMissing, "{buffs}", strings.Join(in.Missing, raidCopy.BuffsJoin)))
	}
	if len(in.Short) > 0 {
		lines = append(lines, copygen.Fill(raidCopy.BuffsShort, "{buffs}", strings.Join(in.Short, raidCopy.BuffsJoin)))
	}
	return truncateRunes(strings.Join(lines, "\n"), maxValueRunes)
}

// raidBaseEmbed is the guild embed every raid post shares: prefix, the escaped title and
// the guild's date as the site's raid row writes it ("Wed 13 Jan", the year only in another
// year), linking to the raid page, with the crest when absolute.
func raidBaseEmbed(in RaidPosted, prefix string, now time.Time) Embed {
	date := textfmt.WeekdayDayMonth(in.StartsAt, textfmt.Location(in.Timezone), now)
	room := maxTitleRunes - len([]rune(prefix)) - len([]rune(fmt.Sprintf(raidCopy.TitleFormat, "", date)))
	title := prefix + fmt.Sprintf(raidCopy.TitleFormat, escapedWithin(collapse(in.Title), maxRaidTitleRunes, room), date)
	// Discord answers 400 to a relative link (which would count toward the five failures),
	// so the title links only when the URL is absolute.
	link := ""
	if isAbsolute(in.RaidURL) {
		link = in.RaidURL
	}
	e := guildBaseEmbed(in.GuildName, title, link, now)
	if isAbsolute(in.CrestURL) {
		e.Thumbnail = in.CrestURL
	}
	return e
}

func raidField(name, value string) EmbedField {
	return EmbedField{Name: name, Value: value, Inline: true}
}

func whenValue(at time.Time) string {
	return fmt.Sprintf(raidCopy.WhenFormat, at.Unix(), at.Unix())
}

// signedField is the bold count alone, inline beside When (spec 2026-10-28 §P.9.2).
func signedField(in RaidPosted) EmbedField {
	return raidField(raidCopy.SignedName, SignedValue(in.Counts, in.Size, in.SizeKnown))
}

// SignedValue is roster.signed: `**33 of 40**` (In plus Late), `**33 signed up**` for a raid of
// unknown size, `None yet` with nobody In or Late.
func SignedValue(c RaidCounts, size int, sizeKnown bool) string {
	signed := c.In + c.Late
	switch {
	case signed == 0:
		return raidCopy.SignedNone
	case sizeKnown:
		return copygen.Fill(raidCopy.SignedLead, "{signed}", strconv.Itoa(signed), "{size}", strconv.Itoa(size))
	}
	return copygen.Fill(raidCopy.SignedLeadOpen, "{signed}", strconv.Itoa(signed))
}

// isAbsolute reports whether raw is an absolute http(s) URL (Discord refuses relative ones).
func isAbsolute(raw string) bool {
	u, err := url.Parse(raw)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}
