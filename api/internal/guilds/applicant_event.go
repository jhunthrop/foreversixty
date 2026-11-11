// api/internal/guilds/applicant_event.go
//
// guild.applicant (spec 2026-11-08-app-flows F3): when someone applies, the guild's officers
// hear it. The push opens the applicant sheet (/app/guild/{id}/applicants/{app}) and offers
// Accept and Decline; Review is the notification itself. Published after the application
// commits, off the request; a failure never fails the application.
package guilds

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// applicantBodyMax is how many runes of the applicant's note the push carries.
const applicantBodyMax = 140

// applicantCopy is the words of guild.applicant.
var applicantCopy = struct {
	Title, NoNote, Accept, Decline string
}{
	Title:   "New applicant: %s",
	NoNote:  "Applied to %s.",
	Accept:  "Accept",
	Decline: "Decline",
}

// ApplicantPagePath is the app's applicant sheet for one application.
func ApplicantPagePath(guildID, applicationID int64) string {
	return fmt.Sprintf("/app/guild/%d/applicants/%d", guildID, applicationID)
}

// publishApplicant tells the guild's officers about application a, off the request.
func (s *Service) publishApplicant(ctx context.Context, guildID int64, a Application) {
	s.background(ctx, func(ctx context.Context) []events.Event {
		g, err := s.Store.getGuild(ctx, guildID)
		if err != nil {
			s.logger().Error("guilds", "op", "applicant_event", "guild_id", guildID, "err", err)
			return nil
		}
		return []events.Event{s.applicantEvent(g.Name, a)}
	})
}

// applicantEvent is guild.applicant for application a of the guild named guildName.
func (s *Service) applicantEvent(guildName string, a Application) events.Event {
	id := strconv.FormatInt(a.ID, 10)
	gid := strconv.FormatInt(a.GuildID, 10)
	return events.Event{
		Kind:      events.GuildApplicant,
		Audience:  events.Audience{Kind: events.ToGuildOfficers, GuildID: a.GuildID},
		DedupeKey: string(events.GuildApplicant) + ":" + id,
		Title:     applicantTitle(a.Character),
		Body:      applicantBody(guildName, a),
		Link:      s.PublicBaseURL + ApplicantPagePath(a.GuildID, a.ID),
		Thread:    "guild-" + gid + "-applicants",
		Collapse:  string(events.GuildApplicant) + "-" + id,
		Actions:   applicantActions(a.GuildID, a.ID),
		// The apps word an action's result from these (round 7 row 2), never from the title.
		Data: map[string]string{"guild_id": gid, "application_id": id,
			"applicant_name": a.Character.Name, "applicant": applicantLine(a.Character)},
	}
}

// applicantActions are Accept and Decline: each POSTs the officer's decision with no body.
func applicantActions(guildID, applicationID int64) []events.Action {
	action := func(decision, title string) events.Action {
		return events.Action{ID: "application_" + decision, Title: title, Method: "POST",
			Path: fmt.Sprintf("/v1/guilds/%d/applications/%d/%s", guildID, applicationID, decision)}
	}
	return []events.Action{action("accept", applicantCopy.Accept), action("decline", applicantCopy.Decline)}
}

// applicantTitle names the applicant: "New applicant: Mirelle, Holy Priest 60".
func applicantTitle(c ApplicantCharacter) string {
	return fmt.Sprintf(applicantCopy.Title, applicantLine(c))
}

// applicantLine is the applicant as the push words them, "Mirelle, Holy Priest 60", each part
// only when known.
func applicantLine(c ApplicantCharacter) string {
	var words []string
	for _, w := range []*string{c.SpecName, c.Class} {
		if w != nil && *w != "" {
			words = append(words, capitalize(*w))
		}
	}
	if c.Level != nil && *c.Level > 0 {
		words = append(words, strconv.Itoa(*c.Level))
	}
	if len(words) == 0 {
		return c.Name
	}
	return c.Name + ", " + strings.Join(words, " ")
}

// applicantBody is the applicant's note, cut at applicantBodyMax runes, else "Applied to {guild}.".
func applicantBody(guildName string, a Application) string {
	note := strings.Join(strings.Fields(a.Message), " ")
	if note == "" {
		return fmt.Sprintf(applicantCopy.NoNote, guildName)
	}
	if utf8.RuneCountInString(note) <= applicantBodyMax {
		return note
	}
	return strings.TrimSpace(string([]rune(note)[:applicantBodyMax])) + "…"
}

// capitalize upper-cases a word's first letter ("priest" reads "Priest").
func capitalize(w string) string {
	r, size := utf8.DecodeRuneInString(w)
	return strings.ToUpper(string(r)) + w[size:]
}
