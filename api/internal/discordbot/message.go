package discordbot

import (
	"encoding/json"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// ephemeralFlag marks a message only the invoking member sees.
const ephemeralFlag = 64

// rowType is the component type of an action row.
const rowType = 1

// ComponentKind is Discord's component type number.
type ComponentKind int

// Component kinds the bot sends.
const (
	KindButton       ComponentKind = 2
	KindStringSelect ComponentKind = 3
)

// ButtonStyle is Discord's button style number.
type ButtonStyle int

// Button styles the bot sends.
const (
	StyleSecondary ButtonStyle = 2
	StyleSuccess   ButtonStyle = 3
	StyleLink      ButtonStyle = 5
)

// SelectOption is one entry of a string select.
type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
}

// Component is a button or a string select; Kind decides which fields count.
type Component struct {
	Kind        ComponentKind
	Style       ButtonStyle
	Label       string
	CustomID    string
	URL         string
	Placeholder string
	Disabled    bool
	Options     []SelectOption
}

type componentJSON struct {
	Type        ComponentKind  `json:"type"`
	Style       ButtonStyle    `json:"style,omitempty"`
	Label       string         `json:"label,omitempty"`
	CustomID    string         `json:"custom_id,omitempty"`
	URL         string         `json:"url,omitempty"`
	Placeholder string         `json:"placeholder,omitempty"`
	Disabled    bool           `json:"disabled,omitempty"`
	Options     []SelectOption `json:"options,omitempty"`
}

// MarshalJSON writes only the fields that belong to the component's kind. A
// link button needs a url and no custom_id; any other button needs a
// custom_id and no url. Anything else fails here, not at Discord.
func (c Component) MarshalJSON() ([]byte, error) {
	out := componentJSON{Type: c.Kind, CustomID: c.CustomID, Disabled: c.Disabled}
	if c.Kind == KindStringSelect {
		out.Placeholder, out.Options = c.Placeholder, c.Options
	} else {
		if err := c.checkButton(); err != nil {
			return nil, err
		}
		out.Style, out.Label, out.URL = c.Style, c.Label, c.URL
	}
	return json.Marshal(out)
}

func (c Component) checkButton() error {
	switch {
	case c.Style == StyleLink && (c.URL == "" || c.CustomID != ""):
		return fmt.Errorf("discordbot: link button %q needs a url and no custom_id", c.Label)
	case c.Style != StyleLink && (c.CustomID == "" || c.URL != ""):
		return fmt.Errorf("discordbot: button %q needs a custom_id and no url", c.Label)
	}
	return nil
}

// ActionRow is one row of components.
type ActionRow []Component

// MarshalJSON wraps the components in Discord's action-row object.
func (r ActionRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       int         `json:"type"`
		Components []Component `json:"components"`
	}{rowType, r})
}

// Message is everything the bot sends or edits.
type Message struct {
	Content    string
	Embeds     []notify.Embed
	Components []ActionRow
	Ephemeral  bool
	// Nonce, on a create, makes Discord answer a repeat (same nonce, same author, within a
	// few minutes) with the message it already made instead of a second one. At most 25 chars.
	Nonce string
}

// maxNonceRunes is Discord's nonce bound.
const maxNonceRunes = 25

type mentions struct {
	Parse []string `json:"parse"`
}

// MarshalJSON always writes an empty mention list, so a message pings no one.
func (m Message) MarshalJSON() ([]byte, error) {
	out := struct {
		Content         string         `json:"content,omitempty"`
		Embeds          []notify.Embed `json:"embeds,omitempty"`
		Components      []ActionRow    `json:"components,omitempty"`
		Flags           int            `json:"flags,omitempty"`
		AllowedMentions mentions       `json:"allowed_mentions"`
		Nonce           string         `json:"nonce,omitempty"`
		EnforceNonce    bool           `json:"enforce_nonce,omitempty"`
	}{
		Content: m.Content, Embeds: m.Embeds, Components: m.Components,
		AllowedMentions: mentions{Parse: []string{}}, Nonce: m.Nonce, EnforceNonce: m.Nonce != "",
	}
	if len([]rune(m.Nonce)) > maxNonceRunes {
		return nil, fmt.Errorf("discordbot: nonce over %d characters", maxNonceRunes)
	}
	if m.Ephemeral {
		out.Flags = ephemeralFlag
	}
	return json.Marshal(out)
}

// edit is the body of an edit. Discord leaves an omitted field unchanged, so
// an edit writes content, embeds and components explicitly: an empty one
// clears it (a locked raid post loses its buttons). The ephemeral flag is
// fixed by the first answer and is not sent.
func (m Message) edit() editBody {
	b := editBody{Content: m.Content, Embeds: m.Embeds, Components: m.Components, AllowedMentions: mentions{Parse: []string{}}}
	if b.Embeds == nil {
		b.Embeds = []notify.Embed{}
	}
	if b.Components == nil {
		b.Components = []ActionRow{}
	}
	return b
}

type editBody struct {
	Content         string         `json:"content"`
	Embeds          []notify.Embed `json:"embeds"`
	Components      []ActionRow    `json:"components"`
	AllowedMentions mentions       `json:"allowed_mentions"`
}
