package discordbot

import (
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// The custom-id shape of a raid post button.
const (
	raidIDPrefix = "raid"
	idSeparator  = ":"
	idParts      = 3
)

// Raid statuses as the guilds store writes them.
const (
	statusCancelled = "cancelled"
	statusDone      = "done"
	statusLocked    = "locked"
)

// signUpStates are the four sign-up buttons, in row order.
var signUpStates = []struct{ State, Label string }{
	{"in", "In"},
	{"late", "Late"},
	{"bench", "Bench"},
	{"out", "Out"},
}

// postCopy is the words of the post's buttons.
var postCopy = struct{ RaidPage, OpenNight string }{
	RaidPage:  "Raid page",
	OpenNight: "Open the night",
}

// PostState is the look of a raid's post at now: cancelled, then done, then ended once the
// end has come, then locked, then moved until the raid starts, else open.
func PostState(status string, startsAt, endsAt time.Time, movedFrom *time.Time, now time.Time) notify.RaidPostState {
	switch {
	case status == statusCancelled:
		return notify.PostCancelled
	case status == statusDone:
		return notify.PostDone
	case !now.Before(endsAt):
		return notify.PostEnded
	case status == statusLocked:
		return notify.PostLocked
	case movedFrom != nil && now.Before(startsAt):
		return notify.PostMoved
	}
	return notify.PostOpen
}

// PostMessage is the bot's raid post in p.State: the one embed and its one row. Every
// state sends its row (possibly empty), so an edit clears buttons a state no longer has.
func PostMessage(p notify.RaidPost, raidURL, nightURL string, now time.Time) Message {
	return Message{Embeds: []notify.Embed{notify.RaidPostEmbed(p, now)}, Components: postRows(p, raidURL, nightURL)}
}

func postRows(p notify.RaidPost, raidURL, nightURL string) []ActionRow {
	switch p.State {
	case notify.PostDone:
		return []ActionRow{{linkButton(postCopy.OpenNight, nightURL)}}
	case notify.PostCancelled:
		return []ActionRow{{linkButton(postCopy.RaidPage, raidURL)}}
	}
	live := p.State == notify.PostOpen || p.State == notify.PostMoved
	row := make(ActionRow, 0, len(signUpStates)+1)
	for i, s := range signUpStates {
		style := StyleSecondary
		if i == 0 {
			style = StyleSuccess
		}
		row = append(row, Component{
			Kind: KindButton, Style: style, Label: s.Label,
			CustomID: raidCustomID(p.RaidID, s.State), Disabled: !live,
		})
	}
	return []ActionRow{append(row, linkButton(postCopy.RaidPage, raidURL))}
}

func linkButton(label, url string) Component {
	return Component{Kind: KindButton, Style: StyleLink, Label: label, URL: url}
}

// raidCustomID is a post button's custom id: raid:{raid}:{state}.
func raidCustomID(raidID int64, state string) string {
	return strings.Join([]string{raidIDPrefix, strconv.FormatInt(raidID, 10), state}, idSeparator)
}

// parseRaidCustomID reads a post button's custom id; ok is false for any other shape.
func parseRaidCustomID(id string) (raidID int64, state string, ok bool) {
	parts := strings.Split(id, idSeparator)
	if len(parts) != idParts || parts[0] != raidIDPrefix || parts[2] == "" {
		return 0, "", false
	}
	raidID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || raidID <= 0 {
		return 0, "", false
	}
	return raidID, parts[2], true
}
