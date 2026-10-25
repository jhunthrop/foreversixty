package notify

import "strings"

// embedFooter is the footer text every message carries.
const embedFooter = "Forever Sixty"

// accountDiscordPath is where the test message's title links.
const accountDiscordPath = "/account#discord"

// embedCopy is the words of every message the channel receives (spec 6.2).
// Task 3 adds the sim and log messages to this table.
var embedCopy = struct {
	TestTitle   string
	TestIntro   string
	TestLead    string
	TestClauses map[EventName]string
	TestNone    string
}{
	TestTitle: "Test message from Forever Sixty",
	TestIntro: "Your webhook works.",
	TestLead:  "You will get a message here when",
	TestClauses: map[EventName]string{
		EventSimFinished:  "a server sim finishes",
		EventUploadParsed: "a log file you upload finishes parsing",
	},
	TestNone: "No events are on yet.",
}

// testEventOrder fixes the order the clauses read in.
var testEventOrder = []EventName{EventSimFinished, EventUploadParsed}

// TestEmbed is the "Send test message" post. Its description names only the
// events currently on.
func TestEmbed(publicBaseURL string, events map[EventName]bool) Embed {
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
	return Embed{
		Title:       embedCopy.TestTitle,
		Description: description,
		URL:         strings.TrimRight(publicBaseURL, "/") + accountDiscordPath,
		Footer:      embedFooter,
	}
}
