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
	// Default shows the option as the select's current pick when the card is drawn again.
	Default bool `json:"default,omitempty"`
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

// Message is everything the bot sends or edits. A ComponentsV2 message (spec 2026-10-29 §0) is
// drawn from Layout, then its action rows, and carries no Content or Embeds.
type Message struct {
	Content    string
	Embeds     []notify.Embed
	Components []ActionRow
	// ComponentsV2 sends the message with IS_COMPONENTS_V2; it is set explicitly, never read
	// off the layout, since a V2 update may carry rows alone.
	ComponentsV2 bool
	// Layout is a Components V2 message's top-level components before its action rows.
	Layout []Node
	// Files are uploaded with the message (multipart); a V2 message shows each one only
	// where a component points at attachment://{Name}.
	Files     []Attachment
	Ephemeral bool
	// Nonce, on a create, makes Discord answer a repeat (same nonce, same author, within a
	// few minutes) with the message it already made instead of a second one. At most 25 chars.
	Nonce string
	// ReplyTo, on a create, makes the message a reply to that message in the same channel;
	// the reply survives the original's deletion and, with the empty mention list, pings no one.
	ReplyTo string
	// echo, when set, makes an edit put a pressed card back as it was (the card failure).
	echo *sourceMessage
	// sourceV2, on an edit, says the message being replaced already has the V2 flag, so the
	// edit must not send content or embeds; zero is a legacy message (spec §A.7 rule 2).
	sourceV2 bool
}

// Attachment is a file sent with a message, with its alt text.
type Attachment struct {
	Name        string
	Description string
	ContentType string
	Data        []byte
}

type attachmentJSON struct {
	ID          int    `json:"id"`
	Filename    string `json:"filename"`
	Description string `json:"description,omitempty"`
}

// componentsV2 reports a Components V2 message.
func (m Message) componentsV2() bool { return m.ComponentsV2 }

// flags is the message's flags: ephemeral and Components V2.
func (m Message) flags() int {
	f := 0
	if m.Ephemeral {
		f |= ephemeralFlag
	}
	if m.componentsV2() {
		f |= componentsV2Flag
	}
	return f
}

// topLevel is a V2 message's components as Discord takes them: the layout, then the rows.
func (m Message) topLevel() []any {
	out := make([]any, 0, len(m.Layout)+len(m.Components))
	for _, n := range m.Layout {
		out = append(out, n)
	}
	for _, r := range m.Components {
		out = append(out, r)
	}
	return out
}

// attachmentsJSON describes Files by index, as the multipart form numbers them.
func (m Message) attachmentsJSON() []attachmentJSON {
	if len(m.Files) == 0 {
		return nil
	}
	out := make([]attachmentJSON, len(m.Files))
	for i, f := range m.Files {
		out[i] = attachmentJSON{ID: i, Filename: f.Name, Description: f.Description}
	}
	return out
}

// maxNonceRunes is Discord's nonce bound.
const maxNonceRunes = 25

// messageReference is a reply's pointer to the message it answers.
type messageReference struct {
	MessageID       string `json:"message_id"`
	FailIfNotExists bool   `json:"fail_if_not_exists"`
}

type mentions struct {
	Parse []string `json:"parse"`
}

// MarshalJSON always writes an empty mention list, so a message pings no one. A V2 message
// with content or embeds fails here: Discord answers it 400.
func (m Message) MarshalJSON() ([]byte, error) {
	if len([]rune(m.Nonce)) > maxNonceRunes {
		return nil, fmt.Errorf("discordbot: nonce over %d characters", maxNonceRunes)
	}
	var reference *messageReference
	if m.ReplyTo != "" {
		reference = &messageReference{MessageID: m.ReplyTo}
	}
	if m.componentsV2() {
		if m.Content != "" || len(m.Embeds) > 0 {
			return nil, fmt.Errorf("discordbot: a components v2 message carries content or embeds")
		}
		return json.Marshal(struct {
			Flags           int               `json:"flags"`
			Components      []any             `json:"components"`
			Attachments     []attachmentJSON  `json:"attachments,omitempty"`
			AllowedMentions mentions          `json:"allowed_mentions"`
			Nonce           string            `json:"nonce,omitempty"`
			EnforceNonce    bool              `json:"enforce_nonce,omitempty"`
			Reference       *messageReference `json:"message_reference,omitempty"`
		}{m.flags(), m.topLevel(), m.attachmentsJSON(), mentions{Parse: []string{}}, m.Nonce, m.Nonce != "", reference})
	}
	return json.Marshal(struct {
		Content         string            `json:"content,omitempty"`
		Embeds          []notify.Embed    `json:"embeds,omitempty"`
		Components      []ActionRow       `json:"components,omitempty"`
		Flags           int               `json:"flags,omitempty"`
		Attachments     []attachmentJSON  `json:"attachments,omitempty"`
		AllowedMentions mentions          `json:"allowed_mentions"`
		Nonce           string            `json:"nonce,omitempty"`
		EnforceNonce    bool              `json:"enforce_nonce,omitempty"`
		Reference       *messageReference `json:"message_reference,omitempty"`
	}{
		Content: m.Content, Embeds: m.Embeds, Components: m.Components, Flags: m.flags(),
		Attachments: m.attachmentsJSON(), AllowedMentions: mentions{Parse: []string{}},
		Nonce: m.Nonce, EnforceNonce: m.Nonce != "", Reference: reference,
	})
}

// edit is the body of an edit. Discord leaves an omitted field unchanged, so a legacy edit
// writes content, embeds and components explicitly: an empty one clears it (a locked raid post
// loses its buttons). The ephemeral flag is fixed by the first answer and is not sent.
//
// A Components V2 edit decides by the message it replaces (spec §A.7 rule 2): over a legacy
// message it sets the flag with content null and embeds empty; over a V2 message it sends
// neither key, since either is a 400 there. A legacy reply over a V2 message (a failure line)
// is sent in the V2 encoding, because the flag can never be removed.
func (m Message) edit() any {
	if m.echo != nil {
		return m.echoEdit()
	}
	if m.sourceV2 && !m.componentsV2() {
		m = asComponentsV2(m, "")
	}
	if m.componentsV2() {
		return m.v2Edit(json.RawMessage(nil))
	}
	embeds, components := m.Embeds, m.Components
	if embeds == nil {
		embeds = []notify.Embed{}
	}
	if components == nil {
		components = []ActionRow{}
	}
	return editBody{Content: m.Content, Embeds: embeds, Components: components, AllowedMentions: mentions{Parse: []string{}}}
}

// v2Edit is a V2 edit body; raw, when set, is the components exactly as the source sent them.
func (m Message) v2Edit(raw json.RawMessage) any {
	var components any = m.topLevel()
	if raw != nil {
		components = raw
	}
	body := v2EditBody{Flags: componentsV2Flag, Components: components, AllowedMentions: mentions{Parse: []string{}}}
	if m.sourceV2 {
		return body
	}
	return v2FromLegacyBody{Content: json.RawMessage("null"), Embeds: []notify.Embed{}, v2EditBody: body}
}

// echoEdit puts the pressed message back: a legacy one with Content as its lead, a V2 one
// with Content as a Text Display on top (any failure line it already had replaced).
func (m Message) echoEdit() any {
	src := m.echo
	if src.componentsV2() {
		m.sourceV2 = true
		return m.v2Edit(withLeadText(src.Components, m.Content))
	}
	return editBody{Content: m.Content, Embeds: rawOrEmpty(src.Embeds), Components: rawOrEmpty(src.Components), AllowedMentions: mentions{Parse: []string{}}}
}

// withLeadText is a V2 message's raw components with text as the first Text Display: a
// top-level Text Display already first (an earlier failure line) is replaced, not stacked.
func withLeadText(raw json.RawMessage, text string) json.RawMessage {
	var list []json.RawMessage
	if json.Unmarshal(raw, &list) != nil {
		list = nil
	}
	if len(list) > 0 {
		var first struct {
			Type ComponentKind `json:"type"`
		}
		if json.Unmarshal(list[0], &first) == nil && first.Type == KindTextDisplay {
			list = list[1:]
		}
	}
	if text != "" {
		lead, err := json.Marshal(textNode(text))
		if err == nil {
			list = append([]json.RawMessage{lead}, list...)
		}
	}
	out, err := json.Marshal(list)
	if err != nil || list == nil {
		return json.RawMessage("[]")
	}
	return out
}

// rawOrEmpty is a source field as sent, or an empty list when the message had none.
func rawOrEmpty(raw json.RawMessage) json.RawMessage {
	if len(raw) == 0 || string(raw) == "null" {
		return json.RawMessage("[]")
	}
	return raw
}

type editBody struct {
	Content         string   `json:"content"`
	Embeds          any      `json:"embeds"`
	Components      any      `json:"components"`
	AllowedMentions mentions `json:"allowed_mentions"`
}

type v2EditBody struct {
	Flags           int      `json:"flags"`
	Components      any      `json:"components"`
	AllowedMentions mentions `json:"allowed_mentions"`
}

// v2FromLegacyBody sets the V2 flag over a legacy message, which Discord accepts only with
// content null and embeds empty.
type v2FromLegacyBody struct {
	Content json.RawMessage `json:"content"`
	Embeds  []notify.Embed  `json:"embeds"`
	v2EditBody
}
