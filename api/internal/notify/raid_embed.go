package notify

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	// cancelColor is #f0736b, the site's wipe red.
	cancelColor = 15758187
	// mutedColor is #9a9484, a locked or ended raid.
	mutedColor = 10130564
	// guildDateLayout and guildTimeLayout render the raid in the guild's own zone.
	guildDateLayout = "Monday 2 January"
	guildTimeLayout = "3:04 pm MST"
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
	TitleFormat     string // %s the escaped title, %s the guild's date
	MovedPrefix     string
	CancelledPrefix string
	DonePrefix      string
	TomorrowPrefix  string
	Descriptions    map[RaidPostState]string
	WhenName        string
	WhenFormat      string // %d the unix start, twice
	WasName         string
	WasFormat       string // %d the old unix start
	StruckFormat    string // %s a struck-through value
	GuildTimeName   string
	SignedName      string
	SignedLead      string // %d signed up, %d the raid size (the bold first clause)
	SignedLeadOpen  string // %d signed up, size unknown
	SignedNone      string
	Tank, Tanks     string // Tanks takes %d
	Healer, Healers string // Healers takes %d
	DPSClause       string // %d
	LateClause      string // %d
	BenchClause     string // %d
	OutClause       string // %d
	ClauseJoin      string
	LineJoin        string
	CameName        string
	CameOf          string // %d came, %d of
	KillsName       string
	// The reminder line (spec 2026-10-28 §P.7).
	ReminderLead        string // %s the escaped title, %d the unix start, twice
	ReminderCount       string // %d signed up, %d the raid size
	ReminderCountOpen   string // %d signed up, size unknown
	ReminderNobody      string
	ReminderWebhookLink string // %s the raid page
}{
	TitleFormat:     "%s · %s",
	MovedPrefix:     "Moved: ",
	CancelledPrefix: "Cancelled: ",
	DonePrefix:      "Done: ",
	TomorrowPrefix:  "Tomorrow: ",
	Descriptions: map[RaidPostState]string{
		PostOpen:      "Sign up with the buttons below.",
		PostLocked:    "Sign-ups are locked. An officer can still change yours.",
		PostMoved:     "Sign-ups carry over. Check you can still make it.",
		PostCancelled: "This raid was cancelled.",
		PostEnded:     "This raid has ended.",
		PostDone:      "Logged as a raid night.",
		PostWebhook:   "Sign up on Forever Sixty.",
		PostTomorrow:  "The raid is tomorrow. Sign up with the buttons below.",
	},
	WhenName:       "When",
	WhenFormat:     "<t:%d:F> (<t:%d:R>)",
	WasName:        "Was",
	WasFormat:      "<t:%d:F>",
	StruckFormat:   "~~%s~~",
	GuildTimeName:  "Guild time",
	SignedName:     "Signed up",
	SignedLead:     "**%d of %d**",
	SignedLeadOpen: "**%d signed up**",
	SignedNone:     "None yet",
	Tank:           "1 tank",
	Tanks:          "%d tanks",
	Healer:         "1 healer",
	Healers:        "%d healers",
	DPSClause:      "%d DPS",
	LateClause:     "%d late",
	BenchClause:    "Bench %d",
	OutClause:      "Out %d",
	ClauseJoin:     " · ",
	LineJoin:       "\n",
	CameName:       "Came",
	CameOf:         "%d of %d",
	KillsName:      "Kills",

	ReminderLead:        "**%s is tomorrow**, <t:%d:F> (<t:%d:R>).",
	ReminderCount:       "%d of %d signed up so far.",
	ReminderCountOpen:   "%d signed up so far.",
	ReminderNobody:      "Nobody has signed up yet.",
	ReminderWebhookLink: "[Sign up on Forever Sixty](%s)",
}

// RaidCounts are a raid's sign-up totals; the roles count In rows only.
type RaidCounts struct {
	In, Late, Bench, Out, Tanks, Healers, DPS int
}

// RaidPosted is what the raid create holds for one new raid. Counts are those at posting.
type RaidPosted struct {
	RaidID           int64
	GuildName, Title string
	StartsAt         time.Time
	// Timezone is the guild's IANA zone; empty reads as UTC and shows no Guild time.
	Timezone string
	// RaidURL is the raid page; CrestURL the guild's crest_url (a thumbnail only when absolute).
	RaidURL, CrestURL string
	Counts            RaidCounts
	Size              int
	SizeKnown         bool
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
	return e
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
			raidField(raidCopy.CameName, fmt.Sprintf(raidCopy.CameOf, p.Came, p.CameOf)),
			{Name: raidCopy.KillsName, Value: killsValue(p.Kills)},
		}
	}
	fields := []EmbedField{raidField(raidCopy.WhenName, when)}
	if p.State == PostMoved && p.Was != nil {
		fields = append(fields, raidField(raidCopy.WasName, fmt.Sprintf(raidCopy.WasFormat, p.Was.Unix())))
	}
	if in.Timezone != "" {
		fields = append(fields, raidField(raidCopy.GuildTimeName, in.StartsAt.In(guildLocation(in.Timezone)).Format(guildTimeLayout)))
	}
	return append(fields, signedField(in))
}

// raidBaseEmbed is the guild embed every raid post shares: prefix, the escaped title and
// the guild's date, linking to the raid page, with the crest when absolute.
func raidBaseEmbed(in RaidPosted, prefix string, now time.Time) Embed {
	date := in.StartsAt.In(guildLocation(in.Timezone)).Format(guildDateLayout)
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

// signedField is the full-width counts field.
func signedField(in RaidPosted) EmbedField {
	return EmbedField{Name: raidCopy.SignedName, Value: SignedUpValue(in.Counts, in.Size, in.SizeKnown)}
}

// SignedUpValue is the counts field's two lines: the G1 grammar, then Bench and Out.
func SignedUpValue(c RaidCounts, size int, sizeKnown bool) string {
	signed := c.In + c.Late
	if signed == 0 && c.Bench == 0 && c.Out == 0 {
		return raidCopy.SignedNone
	}
	lead := raidCopy.SignedNone
	switch {
	case signed > 0 && sizeKnown:
		lead = fmt.Sprintf(raidCopy.SignedLead, signed, size)
	case signed > 0:
		lead = fmt.Sprintf(raidCopy.SignedLeadOpen, signed)
	}
	lines := []string{joinClauses(append([]string{lead}, roleClauses(c)...))}
	if second := joinClauses(offClauses(c)); second != "" {
		lines = append(lines, second)
	}
	return strings.Join(lines, raidCopy.LineJoin)
}

func roleClauses(c RaidCounts) []string {
	var out []string
	if c.Tanks > 0 {
		out = append(out, plural(c.Tanks, raidCopy.Tank, raidCopy.Tanks))
	}
	if c.Healers > 0 {
		out = append(out, plural(c.Healers, raidCopy.Healer, raidCopy.Healers))
	}
	if c.DPS > 0 {
		out = append(out, fmt.Sprintf(raidCopy.DPSClause, c.DPS))
	}
	if c.Late > 0 {
		out = append(out, fmt.Sprintf(raidCopy.LateClause, c.Late))
	}
	return out
}

func offClauses(c RaidCounts) []string {
	var out []string
	if c.Bench > 0 {
		out = append(out, fmt.Sprintf(raidCopy.BenchClause, c.Bench))
	}
	if c.Out > 0 {
		out = append(out, fmt.Sprintf(raidCopy.OutClause, c.Out))
	}
	return out
}

func plural(n int, one, many string) string {
	if n == 1 {
		return one
	}
	return fmt.Sprintf(many, n)
}

func joinClauses(parts []string) string { return strings.Join(parts, raidCopy.ClauseJoin) }

// guildLocation loads the guild's zone, falling back to UTC.
func guildLocation(name string) *time.Location {
	if name == "" {
		return time.UTC
	}
	loc, err := time.LoadLocation(name)
	if err != nil {
		return time.UTC
	}
	return loc
}

// isAbsolute reports whether raw is an absolute http(s) URL (Discord refuses relative ones).
func isAbsolute(raw string) bool {
	u, err := url.Parse(raw)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}
