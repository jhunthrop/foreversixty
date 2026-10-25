package notify

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

const (
	// embedFooter is the footer text every message carries.
	embedFooter = "Forever Sixty"
	// embedColor is #e5b955, the gold of the seal.
	embedColor = 15055189

	accountDiscordPath = "/account#discord"
	simPathPrefix      = "/sim/"
	reportPathPrefix   = "/reports/"
	crestPathFormat    = "/icons/hd/crests/%s.webp"
	// sealPath is the raster Forever Sixty seal the web serves; Discord
	// does not draw SVG avatars.
	sealPath = "/seal-400.png"

	// Discord's limits: 256 for a title, 1024 for a field value.
	maxTitleRunes = 256
	maxValueRunes = 1024

	maxRunRunes    = 100
	maxReportTitle = 200
)

// embedCopy is the words of every message the channel receives (spec 6.2).
var embedCopy = struct {
	TestTitle   string
	TestIntro   string
	TestLead    string
	TestClauses map[EventName]string
	TestNone    string

	SimTitle         string // %s: the spec label
	SimTitleNoSpec   string
	SimResultName    string
	SimResultDPS     string // %s: the formatted DPS
	SimRunName       string
	SimTookName      string
	LogTitleSuffix   string
	LogTitleNoTitle  string
	LogZoneName      string
	LogFightsName    string
	LogFightsFormat  string // %d fights/fight, %s kills phrase
	LogOneFight      string
	LogKillsNone     string
	LogOneKill       string
	LogManyKills     string // %d
	LogVisibleName   string
	LogVisibleLabels map[string]string
}{
	TestTitle: "Test message from Forever Sixty",
	TestIntro: "Your webhook works.",
	TestLead:  "You will get a message here when",
	TestClauses: map[EventName]string{
		EventSimFinished:  "a server sim finishes",
		EventUploadParsed: "a combat log you upload finishes parsing",
	},
	TestNone: "No events are on yet.",

	SimTitle:        "Your %s sim finished",
	SimTitleNoSpec:  "Your sim finished",
	SimResultName:   "Result",
	SimResultDPS:    "%s DPS",
	SimRunName:      "Run",
	SimTookName:     "Took",
	LogTitleSuffix:  " finished parsing",
	LogTitleNoTitle: "Your log finished parsing",
	LogZoneName:     "Zone",
	LogFightsName:   "Fights",
	LogFightsFormat: "%s · %s",
	LogOneFight:     "1 fight",
	LogKillsNone:    "no kills",
	LogOneKill:      "1 kill",
	LogManyKills:    "%d kills",
	LogVisibleName:  "Visible to",
	LogVisibleLabels: map[string]string{
		"public":   "Everyone",
		"unlisted": "Anyone with the link",
		"guild":    "Your guild",
		"private":  "Only you",
	},
}

// testEventOrder fixes the order the clauses read in.
var testEventOrder = []EventName{EventSimFinished, EventUploadParsed}

// absoluteURL joins a site path to the public base, and answers "" when the
// base is not an absolute http(s) URL (Discord refuses relative links).
func absoluteURL(base, path string) string {
	u, err := url.Parse(strings.TrimSpace(base))
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
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
			strings.Join(clauses, " or ") + "."
	}
	e := baseEmbed(embedCopy.TestTitle, absoluteURL(publicBaseURL, accountDiscordPath), now)
	e.Description = description
	return e
}

// SimFinished is what the sim job holds when a server run reaches done.
type SimFinished struct {
	UserID     int64
	SimID      string
	Spec       string
	DPS        float64
	Headline   string
	Title      string
	CreatedAt  time.Time
	FinishedAt time.Time
}

// SimFinishedEmbed is the "Server sim finished" message.
func SimFinishedEmbed(publicBaseURL string, in SimFinished) Embed {
	spec, known := specs.ByKey[in.Spec]
	title := embedCopy.SimTitleNoSpec
	if known {
		title = fmt.Sprintf(embedCopy.SimTitle, spec.Name+" "+classLabel(spec.ClassSlug))
	}
	e := baseEmbed(title, absoluteURL(publicBaseURL, simPathPrefix+in.SimID), in.FinishedAt)
	result := fmt.Sprintf(embedCopy.SimResultDPS, thousands(in.DPS))
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
		Name: embedCopy.SimTookName, Value: duration(in.FinishedAt.Sub(in.CreatedAt)), Inline: true,
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
	Fights     int
	Kills      int
	FinishedAt time.Time
}

// UploadParsedEmbed is the "Log parsed" message.
func UploadParsedEmbed(publicBaseURL string, in UploadParsed) Embed {
	title := embedCopy.LogTitleNoTitle
	if name := collapse(in.Title); name != "" {
		title = escapedWithin(name, maxReportTitle, maxTitleRunes-len([]rune(embedCopy.LogTitleSuffix))) +
			embedCopy.LogTitleSuffix
	}
	e := baseEmbed(title, absoluteURL(publicBaseURL, reportPathPrefix+in.ReportID), in.FinishedAt)
	if zone := collapse(in.Zone); zone != "" {
		e.Fields = append(e.Fields, EmbedField{
			Name: embedCopy.LogZoneName, Value: escapedWithin(zone, maxValueRunes, maxValueRunes), Inline: true,
		})
	}
	e.Fields = append(e.Fields,
		EmbedField{Name: embedCopy.LogFightsName, Value: fightsLine(in.Fights, in.Kills), Inline: true},
		EmbedField{
			Name: embedCopy.LogVisibleName, Value: embedCopy.LogVisibleLabels[in.Visibility], Inline: true,
		})
	return e
}

func fightsLine(fights, kills int) string {
	fightsText := fmt.Sprintf("%d fights", fights)
	if fights == 1 {
		fightsText = embedCopy.LogOneFight
	}
	killsText := fmt.Sprintf(embedCopy.LogManyKills, kills)
	switch kills {
	case 0:
		killsText = embedCopy.LogKillsNone
	case 1:
		killsText = embedCopy.LogOneKill
	}
	return fmt.Sprintf(embedCopy.LogFightsFormat, fightsText, killsText)
}

// classLabel capitalises a class slug ("mage" -> "Mage").
func classLabel(slug string) string {
	if slug == "" {
		return slug
	}
	return strings.ToUpper(slug[:1]) + slug[1:]
}
