package notify

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const (
	// embedFooter is the footer text every message carries.
	embedFooter = copygen.AppName

	accountDiscordPath = "/account#discord"
	simPathPrefix      = "/sim/"
	reportPathPrefix   = "/reports/"
	crestPathFormat    = "/icons/hd/crests/%s.webp"
	// sealPath is the raster Forever Sixty seal the web serves; Discord
	// does not draw SVG avatars.
	sealPath = brand.SealPath

	// Discord's limits: 256 for a title, 1024 for a field value.
	maxTitleRunes = 256
	maxValueRunes = 1024

	maxRunRunes       = 100
	maxCharacterRunes = 40
	maxReportTitle    = 200
)

// embedColor is every normal post's stripe: the teal accent (round 5, color.forever.accent).
var embedColor = brand.RGB(brand.ColorForeverAccent)

// embedCopy is the words of every message the channel receives (spec 6.2).
var embedCopy = struct {
	TestTitle   string
	TestIntro   string
	TestLead    string
	TestClauses map[EventName]string
	TestNone    string
	// TestPrivateClause ends the upload clause when private logs may post.
	TestPrivateClause string

	SimTitle         string // {spec}: the spec label
	SimTitleNamed    string // {name}: the character name; {spec}: the spec label
	SimTitleNoSpec   string
	SimResultName    string
	SimResultDPS     string // {dps}: the formatted DPS
	SimRunName       string
	SimTookName      string
	LogTitle         string // {title}: the log's title
	LogTitleNoTitle  string
	LogZoneName      string
	LogFightsName    string
	LogFightsFormat  string // {n} fights, {kills} the kills phrase
	LogPushFights    string // push's fights line: {fights} the counted fights, {kills} the kills phrase
	LogBossFights    copygen.Plural
	LogKillsNone     string
	LogKills         copygen.Plural
	LogVisibleName   string
	LogVisibleLabels map[string]string
}{
	TestTitle: copygen.DiscordEmbedTestTitle,
	TestIntro: copygen.DiscordEmbedTestIntro,
	TestLead:  copygen.DiscordEmbedTestLead,
	TestClauses: map[EventName]string{
		EventSimFinished:  copygen.DiscordEmbedTestClausesEventSimFinished,
		EventUploadParsed: copygen.DiscordEmbedTestClausesEventUploadParsed,
	},
	TestNone:          copygen.DiscordEmbedTestNone,
	TestPrivateClause: copygen.DiscordEmbedTestPrivateClause,

	SimTitle:        copygen.PushEventSimTitle,
	SimTitleNamed:   copygen.PushEventSimTitleNamed,
	SimTitleNoSpec:  copygen.PushEventSimTitleNoSpec,
	SimResultName:   copygen.DiscordEmbedSimResultName,
	SimResultDPS:    copygen.RaidsNightDps,
	SimRunName:      copygen.DiscordEmbedSimRunName,
	SimTookName:     copygen.DiscordEmbedSimTookName,
	LogTitle:        copygen.PushEventLogTitle,
	LogTitleNoTitle: copygen.PushEventLogTitleNoTitle,
	LogZoneName:     copygen.ItemPageKindZone,
	LogFightsName:   copygen.DiscordEmbedLogFightsName,
	LogFightsFormat: copygen.DiscordEmbedLogFightsFormat,
	LogPushFights:   copygen.PushEventLogFights,
	LogBossFights:   copygen.PushEventBossFights,
	LogKillsNone:    copygen.PushEventLogKillsNone,
	LogKills:        copygen.PushEventKills,
	LogVisibleName:  copygen.DiscordEmbedLogVisibleName,
	LogVisibleLabels: map[string]string{
		"public":   copygen.DiscordEmbedLogVisibleLabelsPublic,
		"unlisted": copygen.DiscordEmbedLogVisibleLabelsUnlisted,
		"guild":    copygen.SharedFeedTagYourGuild,
		"private":  copygen.DiscordEmbedLogVisibleLabelsPrivate,
	},
}

// testEventOrder fixes the order the clauses read in.
var testEventOrder = []EventName{EventSimFinished, EventUploadParsed}

// absoluteURL joins a site path to the public base, and answers "" when the
// base is not an absolute http(s) URL (Discord refuses relative links).
func absoluteURL(base, path string) string {
	if !isAbsolute(strings.TrimSpace(base)) {
		return ""
	}
	return strings.TrimRight(base, "/") + path
}

// AvatarURL is the seal every message is posted with.
func AvatarURL(publicBaseURL string) string { return absoluteURL(publicBaseURL, sealPath) }

// base returns the fields every message shares.
func baseEmbed(title, link string, at time.Time) Embed {
	return Embed{Title: title, URL: link, Footer: embedFooter, Color: embedColor, Timestamp: at}
}

// TestEmbed is the "Send test message" post. Its description names only the
// events currently on.
func TestEmbed(publicBaseURL string, events map[EventName]bool, now time.Time) Embed {
	var clauses []string
	for _, name := range testEventOrder {
		if events[name] {
			clauses = append(clauses, embedCopy.TestClauses[name])
		}
	}
	description := embedCopy.TestIntro + " " + embedCopy.TestNone
	if len(clauses) > 0 {
		description = embedCopy.TestIntro + " " + embedCopy.TestLead + " " +
			strings.Join(clauses, " or ") + privateClause(events) + "."
	}
	e := baseEmbed(embedCopy.TestTitle, absoluteURL(publicBaseURL, accountDiscordPath), now)
	e.Description = description
	return e
}

// privateClause is appended when a private log may post (and logs post at all).
func privateClause(events map[EventName]bool) string {
	if events[EventUploadParsed] && events[EventUploadPrivate] {
		return embedCopy.TestPrivateClause
	}
	return ""
}

// SimFinished is what the sim job holds when a server run reaches done.
type SimFinished struct {
	UserID int64
	SimID  string
	Spec   string
	// Character is the run's character name; empty when unknown.
	Character  string
	DPS        float64
	Headline   string
	Title      string
	CreatedAt  time.Time
	FinishedAt time.Time
}

// SimFinishedEmbed is the "Server sim finished" message.
func SimFinishedEmbed(publicBaseURL string, in SimFinished) Embed {
	spec, known := specs.ByKey[in.Spec]
	title := simTitle(in.Character, spec, known)
	e := baseEmbed(title, absoluteURL(publicBaseURL, simPathPrefix+in.SimID), in.FinishedAt)
	result := copygen.Fill(embedCopy.SimResultDPS, "{dps}", textfmt.Grouped(in.DPS))
	if headline := collapse(in.Headline); headline != "" {
		result = escapedWithin(headline, maxValueRunes, maxValueRunes)
	}
	e.Fields = append(e.Fields, EmbedField{Name: embedCopy.SimResultName, Value: result, Inline: true})
	if run := collapse(in.Title); run != "" {
		e.Fields = append(e.Fields, EmbedField{
			Name: embedCopy.SimRunName, Value: escapedWithin(run, maxRunRunes, maxValueRunes), Inline: true,
		})
	}
	e.Fields = append(e.Fields, EmbedField{
		Name: embedCopy.SimTookName, Value: textfmt.Duration(in.FinishedAt.Sub(in.CreatedAt)), Inline: true,
	})
	if known {
		e.Thumbnail = absoluteURL(publicBaseURL, fmt.Sprintf(crestPathFormat, spec.ClassSlug))
	}
	return e
}

// UploadParsed is what the parse job holds when a whole-file upload completes.
type UploadParsed struct {
	UserID     int64
	ReportID   string
	Title      string
	Zone       string
	Visibility string
	// Private is true when the report is visible to its owner only.
	Private    bool
	Fights     int
	Kills      int
	FinishedAt time.Time
}

// UploadParsedEmbed is the "Log parsed" message.
func UploadParsedEmbed(publicBaseURL string, in UploadParsed) Embed {
	title := embedCopy.LogTitleNoTitle
	if name := collapse(in.Title); name != "" {
		room := maxTitleRunes - len([]rune(copygen.Fill(embedCopy.LogTitle, "{title}", "")))
		title = copygen.Fill(embedCopy.LogTitle, "{title}", escapedWithin(name, maxReportTitle, room))
	}
	e := baseEmbed(title, absoluteURL(publicBaseURL, reportPathPrefix+in.ReportID), in.FinishedAt)
	if zone := collapse(in.Zone); zone != "" {
		e.Fields = append(e.Fields, EmbedField{
			Name: embedCopy.LogZoneName, Value: escapedWithin(zone, maxValueRunes, maxValueRunes), Inline: true,
		})
	}
	if in.Fights > 0 {
		e.Fields = append(e.Fields,
			EmbedField{Name: embedCopy.LogFightsName, Value: fightsLine(in.Fights, in.Kills), Inline: true})
	}
	e.Fields = append(e.Fields, EmbedField{
		Name: embedCopy.LogVisibleName, Value: embedCopy.LogVisibleLabels[in.Visibility], Inline: true,
	})
	return e
}

// simTitle names the character when the run has one, else addresses the member.
func simTitle(character string, spec specs.Spec, known bool) string {
	if !known {
		return embedCopy.SimTitleNoSpec
	}
	label := spec.FullName
	if name := collapse(character); name != "" {
		return copygen.Fill(embedCopy.SimTitleNamed,
			"{name}", escapedWithin(name, maxCharacterRunes, maxCharacterRunes), "{spec}", label)
	}
	return copygen.Fill(embedCopy.SimTitle, "{spec}", label)
}

func fightsLine(fights, kills int) string {
	return copygen.Fill(embedCopy.LogFightsFormat, "{n}", strconv.Itoa(fights), "{kills}", killsPhrase(kills))
}

// classLabel capitalises a class slug ("mage" -> "Mage").
func classLabel(slug string) string {
	if slug == "" {
		return slug
	}
	return strings.ToUpper(slug[:1]) + slug[1:]
}

// BrandEmbed is an accent embed with the plain "Forever Sixty" footer, for answers that belong to
// no guild (the bot's /help).
func BrandEmbed(title, description string, at time.Time) Embed {
	e := baseEmbed(title, "", at)
	e.Description = description
	return e
}
