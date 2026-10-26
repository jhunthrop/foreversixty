package notify

import (
	"fmt"
	"net/url"
	"time"
)

const (
	// cancelColor is #f0736b, the site's wipe red.
	cancelColor = 15758187
	// guildDateLayout and guildTimeLayout render the raid in the guild's own zone.
	guildDateLayout = "Monday 2 January"
	guildTimeLayout = "3:04 pm MST"
	// maxRaidTitleRunes is the raid form's title bound.
	maxRaidTitleRunes = 80
)

// raidCopy is the words of the raid posts (spec 2026-10-26-guild-raids §6.2).
var raidCopy = struct {
	TitleFormat     string // %s the escaped title, %s the guild's date
	MovedPrefix     string
	CancelledPrefix string
	PostedText      string
	MovedText       string
	WhenName        string
	WhenFormat      string // %d the unix start, twice
	WasName         string
	WasFormat       string // %d the old unix start
	StruckFormat    string // %s a struck-through value
	GuildTimeName   string
	SignedName      string
	SignedOf        string // %d signed up, %d the raid size
	SignedCount     string // %d signed up, size unknown
	SignedNone      string
}{
	TitleFormat:     "%s · %s",
	MovedPrefix:     "Moved: ",
	CancelledPrefix: "Cancelled: ",
	PostedText:      "Sign up on Forever Sixty.",
	MovedText:       "Sign-ups carry over. Check you can still make it.",
	WhenName:        "When",
	WhenFormat:      "<t:%d:F> (<t:%d:R>)",
	WasName:         "Was",
	WasFormat:       "<t:%d:F>",
	StruckFormat:    "~~%s~~",
	GuildTimeName:   "Guild time",
	SignedName:      "Signed up",
	SignedOf:        "%d of %d",
	SignedCount:     "%d signed up",
	SignedNone:      "None yet",
}

// RaidPosted is what the raid create holds for one new raid. Counts are those at posting.
type RaidPosted struct {
	GuildName, Title string
	StartsAt         time.Time
	// Timezone is the guild's IANA zone; empty or unknown reads as UTC.
	Timezone string
	// RaidURL is the raid page; CrestURL the guild's crest_url (a thumbnail only when absolute).
	RaidURL, CrestURL string
	SignedUp, Size    int
	SizeKnown         bool
}

// RaidChanged is a raid moved (Was set) or cancelled; StartsAt is the current start.
type RaidChanged struct {
	RaidPosted
	Was       *time.Time
	Cancelled bool
}

// RaidPostedEmbed is the raid_posted message.
func RaidPostedEmbed(in RaidPosted, now time.Time) Embed {
	e := raidBaseEmbed(in, "", now)
	e.Description = raidCopy.PostedText
	e.Fields = []EmbedField{
		raidField(raidCopy.WhenName, whenValue(in.StartsAt)),
		raidField(raidCopy.GuildTimeName, in.StartsAt.In(guildLocation(in.Timezone)).Format(guildTimeLayout)),
		signedField(in),
	}
	return e
}

// RaidChangedEmbed is the raid_changed message: cancelled, or moved from Was.
func RaidChangedEmbed(in RaidChanged, now time.Time) Embed {
	if in.Cancelled {
		e := raidBaseEmbed(in.RaidPosted, raidCopy.CancelledPrefix, now)
		e.Color = cancelColor
		e.Fields = []EmbedField{
			raidField(raidCopy.WhenName, fmt.Sprintf(raidCopy.StruckFormat, whenValue(in.StartsAt))),
			signedField(in.RaidPosted),
		}
		return e
	}
	e := raidBaseEmbed(in.RaidPosted, raidCopy.MovedPrefix, now)
	e.Description = raidCopy.MovedText
	e.Fields = []EmbedField{raidField(raidCopy.WhenName, whenValue(in.StartsAt))}
	if in.Was != nil {
		e.Fields = append(e.Fields, raidField(raidCopy.WasName, fmt.Sprintf(raidCopy.WasFormat, in.Was.Unix())))
	}
	e.Fields = append(e.Fields, signedField(in.RaidPosted))
	return e
}

// raidBaseEmbed is the guild embed every raid post shares: prefix, the escaped title and
// the guild's date, linking to the raid page, with the crest when absolute.
func raidBaseEmbed(in RaidPosted, prefix string, now time.Time) Embed {
	date := in.StartsAt.In(guildLocation(in.Timezone)).Format(guildDateLayout)
	room := maxTitleRunes - len([]rune(prefix)) - len([]rune(fmt.Sprintf(raidCopy.TitleFormat, "", date)))
	title := prefix + fmt.Sprintf(raidCopy.TitleFormat, escapedWithin(collapse(in.Title), maxRaidTitleRunes, room), date)
	e := guildBaseEmbed(in.GuildName, title, in.RaidURL, now)
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

func signedField(in RaidPosted) EmbedField {
	value := raidCopy.SignedNone
	switch {
	case in.SignedUp > 0 && in.SizeKnown:
		value = fmt.Sprintf(raidCopy.SignedOf, in.SignedUp, in.Size)
	case in.SignedUp > 0:
		value = fmt.Sprintf(raidCopy.SignedCount, in.SignedUp)
	}
	return raidField(raidCopy.SignedName, value)
}

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
