// companion/internal/addon/outbox.go
// The addon's outbox (docs/contracts/addon-v2.md, contract 4): what the
// player's last session left in ForeverSixtyDB.outbox at logout — the
// export, an in-game sim request, a guild roster snapshot — read out of
// the saved-variables file and sent to the site beside the character
// exports, and the plain words the site's answer is logged in.
package addon

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// OutboxFormat is the outbox format this companion reads. An outbox of
// any other format is not sent (contract 4: a reader ignores a format
// it does not know) and the file is left as it is.
const OutboxFormat = 1

// Outbox is one format 1 outbox, ready to send. Body is the whole
// outbox as JSON, every field the addon wrote carried through, so a
// field the addon adds inside format 1 reaches the site without a new
// companion. SavedAt (export.saved_at) and RequestedAt
// (sim_request.requested_at) are what make one outbox a different one
// from the last: the sync sends each outbox once.
type Outbox struct {
	SavedAt     string
	RequestedAt string
	Body        json.RawMessage
}

// key is what the sync remembers of the last outbox it sent.
func (o *Outbox) key() string { return o.SavedAt + "\x00" + o.RequestedAt }

// ScanOutbox reads ForeverSixtyDB.outbox out of one SavedVariables
// file. It answers nil, with no error, when there is nothing to send:
// no outbox, an outbox whose format is not OutboxFormat, or one with no
// export code (the site cannot store an outbox without its export). It
// only reads the file.
func ScanOutbox(path string) (*Outbox, error) {
	saved, err := scanSaved(path)
	return saved.Outbox, err
}

// outboxOf finds the outbox in a parsed SavedVariables file.
func outboxOf(top map[string]any) (*Outbox, error) {
	db, ok := top["ForeverSixtyDB"].(*Table)
	if !ok {
		return nil, nil
	}
	t, ok := db.Fields["outbox"].(*Table)
	if !ok {
		return nil, nil
	}
	if f, ok := t.Fields["format"].(float64); !ok || f != OutboxFormat {
		return nil, nil
	}
	export, _ := t.Fields["export"].(*Table)
	if export.Get("code") == "" {
		return nil, nil
	}
	body, err := json.Marshal(luaJSON(t))
	if err != nil {
		return nil, fmt.Errorf("the addon's outbox cannot be sent: %w", err)
	}
	ob := &Outbox{SavedAt: export.Get("saved_at"), Body: body}
	if sim, ok := t.Fields["sim_request"].(*Table); ok {
		ob.RequestedAt = sim.Get("requested_at")
	}
	return ob, nil
}

// luaJSON turns a decoded Lua value into the JSON value it stands for:
// a table of positional entries (or of keys 1..n) is a list, any other
// table an object, and strings, numbers, booleans and nil are
// themselves.
func luaJSON(v any) any {
	t, ok := v.(*Table)
	if !ok {
		return v
	}
	if len(t.Fields) == 0 {
		list := make([]any, 0, len(t.Items))
		for _, item := range t.Items {
			list = append(list, luaJSON(item))
		}
		return list
	}
	if len(t.Items) == 0 && isSequence(t.Fields) {
		list := make([]any, len(t.Fields))
		for k, item := range t.Fields {
			i, _ := strconv.Atoi(k)
			list[i-1] = luaJSON(item)
		}
		return list
	}
	obj := make(map[string]any, len(t.Fields)+len(t.Items))
	for i, item := range t.Items {
		obj[strconv.Itoa(i+1)] = luaJSON(item)
	}
	for k, item := range t.Fields {
		obj[k] = luaJSON(item)
	}
	return obj
}

// isSequence reports whether the keys are exactly "1".."n".
func isSequence(fields map[string]any) bool {
	for k := range fields {
		i, err := strconv.Atoi(k)
		if err != nil || i < 1 || i > len(fields) || strconv.Itoa(i) != k {
			return false
		}
	}
	return true
}

// ExportsResult is the body POST /v1/addon/exports answers with
// (api/internal/addon/addon.go ExportsResult): how many exports were
// stored and, when an outbox went up, what became of each of its parts.
type ExportsResult struct {
	Stored int           `json:"stored"`
	Outbox *OutboxResult `json:"outbox,omitempty"`
}

// OutboxResult is the API's answer for each part of an outbox
// (api/internal/addon/outbox.go OutboxResult). A part the outbox did
// not carry is empty.
type OutboxResult struct {
	Status         string `json:"status"`
	Export         string `json:"export,omitempty"`
	SimRequest     string `json:"sim_request,omitempty"`
	RosterSnapshot string `json:"roster_snapshot,omitempty"`
	DroppedMembers int    `json:"dropped_members,omitempty"`
}

// What the API answers for an outbox and its parts.
const (
	OutboxIgnored          = "ignored"
	OutboxStored           = "stored"
	OutboxUnknownCharacter = "unknown_character"
	OutboxNotInGuild       = "not_in_guild"
	OutboxInvalid          = "invalid"

	// The sim request's own outcomes (api/internal/sims RequestFromAddon).
	SimOutcomeQueued          = "queued"
	SimOutcomeDuplicate       = "duplicate"
	SimOutcomePremiumRequired = "premium_required"
	SimOutcomeUnavailable     = "unavailable"
	SimOutcomeInvalid         = "invalid"
	SimOutcomeNotStarted      = "not_started"
)

// What became of a character's newest in-game sim request, in the
// inbox ack's words (contract 3, ack.sim_request_state). The status
// page shows these.
const (
	SimQueued       = "queued"
	SimNeedsPremium = "needs_premium"
	SimDuplicate    = "duplicate"
	SimInvalid      = "invalid"
)

// SimStateOf is a sim request outcome from the export answer in the
// ack's words, or "" for an outcome the ack has no word for.
func SimStateOf(outcome string) string {
	switch outcome {
	case SimOutcomeQueued:
		return SimQueued
	case SimOutcomePremiumRequired:
		return SimNeedsPremium
	case SimOutcomeDuplicate:
		return SimDuplicate
	case SimOutcomeInvalid:
		return SimInvalid
	}
	return ""
}

// The log's words for the site's answer, one line per part. They are
// read by players pasting companion.log, so they are plain.
const (
	CopyOutboxIgnored    = "The site does not read this outbox format, so it was ignored"
	CopyOutboxRefused    = "The site could not read the outbox, so it will not be sent again"
	CopyExportStored     = "Export stored"
	CopyExportUnknown    = "Export not stored, the site has not seen this character yet"
	CopySimQueued        = "Sim request queued"
	CopySimNeedsPremium  = "Sim request needs a premium account"
	CopySimDuplicate     = "Sim request already running this hour"
	CopySimInvalid       = "Sim request invalid, the site could not sim it"
	CopySimUnknown       = "Sim request not taken, the site has not seen this character yet"
	CopySimUnavailable   = "Sim request not taken, server sims are switched off"
	CopySimNotStarted    = "Sim request not started"
	CopyRosterStored     = "Roster snapshot stored"
	CopyRosterNotInGuild = "Roster snapshot not stored, the character is not in that guild"
	CopyRosterInvalid    = "Roster snapshot invalid, nothing in it could be stored"
	CopyRosterUnknown    = "Roster snapshot not stored, the site has not seen this character yet"
)

// Lines is the answer in plain words, one line per part the outbox
// carried.
func (r OutboxResult) Lines() []string {
	if r.Status == OutboxIgnored {
		return []string{CopyOutboxIgnored}
	}
	var out []string
	pick := func(got string, words map[string]string, what string) {
		if got == "" {
			return
		}
		if w, ok := words[got]; ok {
			out = append(out, w)
			return
		}
		out = append(out, what+" answered: "+got)
	}
	pick(r.Export, map[string]string{
		OutboxStored: CopyExportStored, OutboxUnknownCharacter: CopyExportUnknown,
	}, "Export")
	pick(r.SimRequest, map[string]string{
		SimOutcomeQueued: CopySimQueued, SimOutcomePremiumRequired: CopySimNeedsPremium,
		SimOutcomeDuplicate: CopySimDuplicate, SimOutcomeInvalid: CopySimInvalid,
		OutboxUnknownCharacter: CopySimUnknown, SimOutcomeUnavailable: CopySimUnavailable,
		SimOutcomeNotStarted: CopySimNotStarted,
	}, "Sim request")
	before := len(out)
	pick(r.RosterSnapshot, map[string]string{
		OutboxStored: CopyRosterStored, OutboxNotInGuild: CopyRosterNotInGuild,
		OutboxInvalid: CopyRosterInvalid, OutboxUnknownCharacter: CopyRosterUnknown,
	}, "Roster snapshot")
	if len(out) > before && r.DroppedMembers > 0 {
		noun := "members"
		if r.DroppedMembers == 1 {
			noun = "member"
		}
		out[len(out)-1] += fmt.Sprintf(", %d %s dropped", r.DroppedMembers, noun)
	}
	return out
}
