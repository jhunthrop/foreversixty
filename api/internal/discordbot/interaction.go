package discordbot

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// Interaction types Discord sends.
const (
	typePing         = 1
	typeCommand      = 2
	typeComponent    = 3
	typeAutocomplete = 4
)

// Interaction is the part of Discord's interaction object the bot reads. Token
// is a credential for the 15-minute follow-up window: never log it.
type Interaction struct {
	ID, ApplicationID, Token, GuildID string
	Type                              int
	Data                              json.RawMessage
	MessageID                         string
	DiscordUserID                     string
	ReceivedAt                        time.Time
}

type wireUser struct {
	ID string `json:"id"`
}

type wireInteraction struct {
	ID            string          `json:"id"`
	ApplicationID string          `json:"application_id"`
	Type          int             `json:"type"`
	Token         string          `json:"token"`
	GuildID       string          `json:"guild_id"`
	Data          json.RawMessage `json:"data"`
	Member        *struct {
		User wireUser `json:"user"`
	} `json:"member"`
	User    *wireUser `json:"user"`
	Message *struct {
		ID string `json:"id"`
	} `json:"message"`
}

// ParseInteraction reads a verified body.
func ParseInteraction(body []byte, receivedAt time.Time) (Interaction, error) {
	var w wireInteraction
	if err := json.Unmarshal(body, &w); err != nil {
		return Interaction{}, errors.New("discordbot: unreadable interaction")
	}
	in := Interaction{
		ID: w.ID, ApplicationID: w.ApplicationID, Token: w.Token, GuildID: w.GuildID,
		Type: w.Type, Data: w.Data, ReceivedAt: receivedAt,
	}
	switch {
	case w.Member != nil:
		in.DiscordUserID = w.Member.User.ID
	case w.User != nil:
		in.DiscordUserID = w.User.ID
	}
	if w.Message != nil {
		in.MessageID = w.Message.ID
	}
	return in, nil
}

type wireOption struct {
	Name    string          `json:"name"`
	Value   json.RawMessage `json:"value"`
	Focused bool            `json:"focused"`
}

type wireData struct {
	Name     string       `json:"name"`
	CustomID string       `json:"custom_id"`
	Values   []string     `json:"values"`
	Options  []wireOption `json:"options"`
}

func (i Interaction) data() wireData {
	var d wireData
	_ = json.Unmarshal(i.Data, &d) // a malformed data object reads as empty
	return d
}

// CommandName is the slash command's name.
func (i Interaction) CommandName() string { return i.data().Name }

// CustomID is the pressed or chosen component's id.
func (i Interaction) CustomID() string { return i.data().CustomID }

// Values are the options chosen in a select.
func (i Interaction) Values() []string { return i.data().Values }

// Option reads a command option as text, whatever its JSON type.
func (i Interaction) Option(name string) (string, bool) {
	for _, o := range i.data().Options {
		if o.Name == name {
			return optionText(o.Value), true
		}
	}
	return "", false
}

// Focused is the option the member is typing in an autocomplete.
func (i Interaction) Focused() (name, value string) {
	for _, o := range i.data().Options {
		if o.Focused {
			return o.Name, optionText(o.Value)
		}
	}
	return "", ""
}

func optionText(raw json.RawMessage) string {
	var s string
	if json.Unmarshal(raw, &s) == nil {
		return s
	}
	var f float64
	if json.Unmarshal(raw, &f) == nil {
		return strconv.FormatFloat(f, 'f', -1, 64)
	}
	return string(raw)
}
