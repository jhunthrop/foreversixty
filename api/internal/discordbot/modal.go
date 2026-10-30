package discordbot

import (
	"encoding/json"
	"fmt"
	"strings"
)

// modalInputID is the custom id of the one text input in both modals.
const modalInputID = "name"

// Component type and style of the modal's text input.
const (
	kindTextInput  = 4
	styleTextShort = 1
)

// Modal is a one-input dialog. Opening it does no I/O, so it always answers inside the window.
type Modal struct {
	CustomID, Title, InputLabel, Placeholder string
	MinLength, MaxLength                     int
}

type textInputJSON struct {
	Type        int    `json:"type"`
	CustomID    string `json:"custom_id"`
	Style       int    `json:"style"`
	Label       string `json:"label"`
	Placeholder string `json:"placeholder,omitempty"`
	MinLength   int    `json:"min_length,omitempty"`
	MaxLength   int    `json:"max_length,omitempty"`
	Required    bool   `json:"required"`
}

// MarshalJSON writes Discord's modal object.
func (m Modal) MarshalJSON() ([]byte, error) {
	if m.CustomID == "" {
		return nil, fmt.Errorf("discordbot: modal %q needs a custom_id", m.Title)
	}
	input := textInputJSON{
		Type: kindTextInput, CustomID: modalInputID, Style: styleTextShort, Label: m.InputLabel,
		Placeholder: m.Placeholder, MinLength: m.MinLength, MaxLength: m.MaxLength, Required: true,
	}
	row := struct {
		Type       int             `json:"type"`
		Components []textInputJSON `json:"components"`
	}{rowType, []textInputJSON{input}}
	return json.Marshal(struct {
		CustomID   string `json:"custom_id"`
		Title      string `json:"title"`
		Components any    `json:"components"`
	}{m.CustomID, m.Title, []any{row}})
}

type wireModalRow struct {
	Components []struct {
		CustomID string `json:"custom_id"`
		Value    string `json:"value"`
	} `json:"components"`
}

// maxModalValue bounds what ModalValue returns; callers still validate their own length.
const maxModalValue = 100

// ModalValue is what the member typed into the modal input with that id, trimmed and cut to
// maxModalValue runes; empty when absent or malformed. It reads the action-row layout only,
// not the newer Label-wrapped (type 18) one.
func (i Interaction) ModalValue(inputID string) string {
	var d struct {
		Components []wireModalRow `json:"components"`
	}
	_ = json.Unmarshal(i.Data, &d) // a malformed data object reads as empty
	for _, row := range d.Components {
		for _, c := range row.Components {
			if c.CustomID == inputID {
				return boundedValue(c.Value)
			}
		}
	}
	return ""
}

// sourceMessage is the message a component sat on, kept raw so a failed press can put it back
// exactly as it was.
type sourceMessage struct {
	Flags      int             `json:"flags"`
	Embeds     json.RawMessage `json:"embeds"`
	Components json.RawMessage `json:"components"`
}

// componentsV2 reports a source message that has the V2 flag.
func (m sourceMessage) componentsV2() bool { return m.Flags&componentsV2Flag != 0 }

// cardFailureRoutes are the card prefixes: a handler error on one keeps the card. Read-only.
var cardFailureRoutes = map[string]bool{
	prefixBis: true, prefixChar: true, prefixSignup: true,
	prefixAttendance: true, prefixReadiness: true, prefixReport: true, prefixSpec: true,
}

// cardFailureLeads are the cards whose failure is said in the embed's lead, not a content line:
// /bis says card.failure, the spec card pressCopy.Failure (spec 2026-10-28 §P.10.6).
var cardFailureLeads = map[string]string{prefixBis: guideCopy.CardFailure, prefixSpec: pressCopy.Failure}

// askPart is the card part whose press opens a modal.
const askPart = "ask"

// opensModal reports a press that only opens a modal: it reads nothing, so it skips the
// server lookup and gets the whole answer window.
func opensModal(in Interaction) bool {
	if in.Type != typeComponent {
		return false
	}
	p, ok := parseCardID(in.CustomID())
	return ok && cardFailureRoutes[p.Prefix] && p.Part == askPart
}

// cardFailureReply puts the pressed card back with every pick and says `card.failure`: as the
// lead for /bis, as the content line elsewhere. A press with no source card gets the plain
// failure sentence in place.
func cardFailureReply(in Interaction) Reply {
	src := in.Source
	if len(src.Embeds) == 0 && len(src.Components) == 0 {
		rep := failureReply()
		rep.Update = true
		return rep
	}
	msg := Message{Ephemeral: true, echo: &src}
	if lead, ok := cardFailureLeads[customPrefix(in.CustomID())]; ok && hasEmbeds(src.Embeds) {
		msg.echo = &sourceMessage{Embeds: withLead(src.Embeds, lead), Components: src.Components}
	} else if ok {
		msg.Content = lead
	} else {
		msg.Content = guideCopy.CardFailure
	}
	return Reply{Update: true, Message: msg}
}

// hasEmbeds reports a pressed message that carried at least one embed.
func hasEmbeds(embeds json.RawMessage) bool {
	var list []json.RawMessage
	return json.Unmarshal(embeds, &list) == nil && len(list) > 0
}

// withLead replaces the first embed's description, leaving every other field as sent.
func withLead(embeds json.RawMessage, lead string) json.RawMessage {
	var list []map[string]json.RawMessage
	if json.Unmarshal(embeds, &list) != nil || len(list) == 0 {
		return embeds
	}
	text, _ := json.Marshal(lead)
	list[0]["description"] = text
	out, err := json.Marshal(list)
	if err != nil {
		return embeds
	}
	return out
}

func boundedValue(raw string) string {
	runes := []rune(strings.TrimSpace(raw))
	if len(runes) > maxModalValue {
		runes = runes[:maxModalValue]
	}
	return string(runes)
}
