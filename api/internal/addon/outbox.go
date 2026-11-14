package addon

// The addon's outbox (docs/contracts/addon-v2.md, contract 4), as the
// companion posts it to POST /v1/addon/exports: the character's export, an
// optional in-game sim request, and an optional snapshot of the guild tab's
// live roster for the guild page.

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/directory"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// OutboxFormat is the outbox format this API reads. Any other is ignored
// whole, as contract 4's versioning rule asks of every reader.
const OutboxFormat = 1

const (
	// maxRosterMembers is a WoW guild's own member cap.
	maxRosterMembers = 1000
	// maxMemberText bounds a roster member's name, spec and build name,
	// in runes.
	maxMemberText = 64
)

// What applyOutbox did with each part of an outbox.
const (
	OutboxRead    = "read"
	OutboxIgnored = "ignored"

	OutboxStored = "stored"
	// OutboxUnknownCharacter: the outbox names a character this account
	// has never synced with a region, so there is no key to store it
	// under. The companion's per-character list carries the region.
	OutboxUnknownCharacter = "unknown_character"
	// OutboxNotInGuild: the snapshot names a guild the reporting
	// character is not in, and only a member's addon may describe a
	// guild's roster.
	OutboxNotInGuild = "not_in_guild"
	// OutboxInvalid: the part's content cannot be used (a sim request
	// that is not one, a snapshot with no usable guild or no usable
	// member). It is answered here, never by failing the sync the export
	// rode in.
	OutboxInvalid = "invalid"
)

// outboxWire is contract 4's format 1 as it arrives. The two optional
// parts stay raw so that their content, some of which another player's
// addon wrote, is judged part by part rather than failing the decode.
type outboxWire struct {
	Format         int             `json:"format"`
	Export         *OutboxExport   `json:"export"`
	SimRequest     json.RawMessage `json:"sim_request"`
	RosterSnapshot json.RawMessage `json:"roster_snapshot"`
	RaidAnswer     json.RawMessage `json:"raid_answer"`
}

// Outbox is a read format 1 outbox. A part the outbox carried whose
// content cannot be used is marked Invalid, so the reply can say so.
type Outbox struct {
	Export OutboxExport

	SimRequest        *sims.AddonSimRequest
	SimRequestInvalid bool

	RosterSnapshot *RosterSnapshot
	RosterInvalid  bool
	// DroppedMembers counts the snapshot's members that failed
	// validation and were left out of what is stored.
	DroppedMembers int

	// RaidAnswers is the raid_answer list (raid_answers.go), nil when the outbox carried none.
	RaidAnswers []RaidAnswerEntry
}

// OutboxExport is the outbox's export: Character is Name-Realm, Code the
// FS1 export string, SavedAt the player's local time.
type OutboxExport struct {
	Character string `json:"character"`
	Code      string `json:"code"`
	SavedAt   string `json:"saved_at"`
	// ScoreVisible is false when the player turned off "Show my Forever Score to others";
	// absent (nil) is visible (contract 4, added 2026-11-08).
	ScoreVisible *bool `json:"score_visible,omitempty"`
}

// scoreVisible is the export's choice: hidden only when it says false.
func (e OutboxExport) scoreVisible() bool { return e.ScoreVisible == nil || *e.ScoreVisible }

// RosterSnapshot is the guild tab's live roster as one member's addon saw it.
type RosterSnapshot struct {
	Guild   string         `json:"guild"`
	Members []RosterMember `json:"members"`
	// OptOutCodes are the "Remove my character" codes the addon found in members' public guild
	// notes: each member's plain name and the code alone, never the rest of the note.
	OptOutCodes []directory.NoteCode `json:"optout_codes,omitempty"`
}

// OptOutCodeVerifier removes the directory characters whose guild notes carry their codes.
type OptOutCodeVerifier interface {
	VerifyNoteCodes(ctx context.Context, region, guild string, codes []directory.NoteCode) (int, error)
}

// maxOptOutCodes bounds the codes one snapshot may carry (a guild's own member cap).
const maxOptOutCodes = maxRosterMembers

// RosterMember is one member line: contract 2's hello fields plus name and
// seen_at.
type RosterMember struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Spec    string `json:"spec"`
	GearGap int    `json:"gear_gap"`
	Build   string `json:"build"`
	SeenAt  string `json:"seen_at"`
	// Class is the client's classFilename from GetGuildRosterInfo ("WARRIOR"); absent from an
	// older addon and for a name the guild roster did not hold.
	Class string `json:"class,omitempty"`
}

// OutboxResult says what became of each part of the outbox. A part the
// outbox did not carry is absent.
type OutboxResult struct {
	Status         string `json:"status"`
	Export         string `json:"export,omitempty"`
	SimRequest     string `json:"sim_request,omitempty"`
	RosterSnapshot string `json:"roster_snapshot,omitempty"`
	// DroppedMembers is how many snapshot members failed validation and
	// were not stored.
	DroppedMembers int `json:"dropped_members,omitempty"`
	// RaidAnswer is, entry by entry in the outbox's order, what became of each raid_answer:
	// accepted or refused (the companion removes it from the game's file) or deferred (kept and
	// sent again). Absent when the outbox carried none.
	RaidAnswer []RaidAnswerResult `json:"raid_answer,omitempty"`
}

// readOutbox reads the outbox's format and, for format 1, the outbox
// itself. Only an outbox that is structurally not one (not an object, a
// non-numeric format, top-level fields of the wrong type, or no usable
// export) is answered 400; the optional parts' content never is. A nil
// outbox with ok true is one of a format this API does not read.
func readOutbox(w http.ResponseWriter, r *http.Request, raw json.RawMessage) (*Outbox, bool) {
	var head struct {
		Format int `json:"format"`
	}
	if err := json.Unmarshal(raw, &head); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that outbox cannot be read",
			map[string]string{"outbox": "an object with a numeric format"})
		return nil, false
	}
	if head.Format != OutboxFormat {
		return nil, true
	}
	var wire outboxWire
	if err := json.Unmarshal(raw, &wire); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that outbox cannot be read",
			map[string]string{"outbox": err.Error()})
		return nil, false
	}
	if problem := exportProblem(wire.Export); problem != "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that outbox cannot be stored",
			map[string]string{"outbox": problem})
		return nil, false
	}
	ob := &Outbox{Export: *wire.Export}
	if present(wire.SimRequest) {
		ob.SimRequest, ob.SimRequestInvalid = readSimRequest(wire.SimRequest)
	}
	if present(wire.RosterSnapshot) {
		ob.RosterSnapshot, ob.DroppedMembers, ob.RosterInvalid = readRoster(wire.RosterSnapshot)
	}
	if present(wire.RaidAnswer) {
		ob.RaidAnswers = readRaidAnswers(wire.RaidAnswer)
	}
	return ob, true
}

// present reports whether an optional part was sent: contract 4 treats a
// missing key and nil the same.
func present(raw json.RawMessage) bool {
	return len(raw) > 0 && string(raw) != "null"
}

// exportProblem is what is wrong with the outbox's export, or "".
func exportProblem(e *OutboxExport) string {
	if e == nil || e.Code == "" || len(e.Code) > MaxExportLen {
		return "export needs a character and a code of at most 4096 characters"
	}
	if !validUnit(e.Character) {
		return "export.character must be Name-Realm"
	}
	return preLiveProblem(e.Character, e.Code)
}

// readSimRequest decodes and checks the sim request; invalid is true when
// its content cannot be used.
func readSimRequest(raw json.RawMessage) (req *sims.AddonSimRequest, invalid bool) {
	var sr sims.AddonSimRequest
	if err := json.Unmarshal(raw, &sr); err != nil {
		return nil, true
	}
	if len(sr.GearCode) > MaxExportLen || !validUnit(sr.Character) || sr.Validate() != nil {
		return nil, true
	}
	return &sr, false
}

// readRoster decodes the snapshot member by member. A member whose fields
// fail validation (another player's addon wrote them) is dropped and
// counted, and the rest are kept, at most maxRosterMembers of them. The
// snapshot is invalid when it is not an object with a guild name, or when
// it carried members and every one was dropped. The guild name is
// normalised the way an export's guild= is.
func readRoster(raw json.RawMessage) (snap *RosterSnapshot, dropped int, invalid bool) {
	var wire struct {
		Guild       string               `json:"guild"`
		Members     []json.RawMessage    `json:"members"`
		OptOutCodes []directory.NoteCode `json:"optout_codes"`
	}
	if err := json.Unmarshal(raw, &wire); err != nil {
		return nil, 0, true
	}
	guild, ok := validateGuildName(wire.Guild)
	if !ok {
		return nil, 0, true
	}
	kept := []RosterMember{}
	for _, m := range wire.Members {
		var member RosterMember
		if len(kept) == maxRosterMembers || json.Unmarshal(m, &member) != nil || !member.valid() {
			dropped++
			continue
		}
		kept = append(kept, member)
	}
	if len(wire.Members) > 0 && len(kept) == 0 {
		return nil, dropped, true
	}
	return &RosterSnapshot{Guild: guild, Members: kept, OptOutCodes: readOptOutCodes(wire.OptOutCodes)}, dropped, false
}

// readOptOutCodes keeps the well-formed codes, each with the member's plain name (the realm
// segment the client adds is dropped), at most maxOptOutCodes of them.
func readOptOutCodes(wire []directory.NoteCode) []directory.NoteCode {
	var out []directory.NoteCode
	for _, nc := range wire {
		if len(out) == maxOptOutCodes {
			break
		}
		name, _ := character.SplitUnit(nc.Name)
		if !directory.ValidCode(nc.Code) || !plainText(nc.Name) || !character.ValidSlug(character.Slug(name)) {
			continue
		}
		out = append(out, directory.NoteCode{Name: name, Code: nc.Code})
	}
	return out
}

// valid reports whether a member line can go on the guild page. gear_gap
// is any integer, as contract 2 types it.
func (m RosterMember) valid() bool {
	if !plainText(m.Name) || !validUnit(m.Name) {
		return false
	}
	if m.Level < 1 || m.Level > simapi.MaxLevel {
		return false
	}
	if !plainText(m.Spec) || !plainText(m.Build) || !plainText(m.Class) {
		return false
	}
	_, err := time.Parse(sims.RequestedAtLayout, m.SeenAt)
	return err == nil
}

// validUnit reports whether s is a Name-Realm whose name could be a
// character key's slug.
func validUnit(s string) bool {
	name, _ := character.SplitUnit(s)
	return name != "" && character.ValidSlug(character.Slug(name))
}

// plainText reports whether s is text another player's addon may put on the
// guild page: valid UTF-8, at most maxMemberText runes, and no control,
// format or private-use runes (validateGuildName's rule).
func plainText(s string) bool {
	if !utf8.ValidString(s) || utf8.RuneCountInString(s) > maxMemberText {
		return false
	}
	for _, r := range s {
		if unicode.Is(unicode.C, r) {
			return false
		}
	}
	return true
}

// resolveOwn finds which of the account's own exported characters a
// Name-Realm names. The outbox carries no region, so this matches the
// name's slug against the characters the account has synced, the same
// best-effort rule resolveCharacter applies to a saved sim; when two of
// them share a name, a realm segment that is a ruleset picks between them.
// Anything still ambiguous is no match rather than a guess.
func resolveOwn(exports []Export, unit string) (Export, bool) {
	name, realm := character.SplitUnit(unit)
	slug := character.Slug(name)
	var found []Export
	for _, e := range exports {
		if character.Slug(e.Name) == slug {
			found = append(found, e)
		}
	}
	if len(found) > 1 && character.ValidRuleset(strings.ToLower(realm)) {
		var same []Export
		for _, e := range found {
			if e.Ruleset == strings.ToLower(realm) {
				same = append(same, e)
			}
		}
		found = same
	}
	if len(found) != 1 {
		return Export{}, false
	}
	return found[0], true
}

// applyOutbox stores a read outbox's export under the account's own
// character, starts its sim request and stores its roster snapshot, and
// returns the key the export was stored under ("" when it was not). A nil
// outbox is one of a format this API does not read. Anything about the
// outbox's content is answered in the result; only a store failure (or a
// claimed character) is an error.
func (s *Service) applyOutbox(ctx context.Context, userID int64, ob *Outbox) (*OutboxResult, string, error) {
	if ob == nil {
		return &OutboxResult{Status: OutboxIgnored}, "", nil
	}
	res := &OutboxResult{Status: OutboxRead, DroppedMembers: ob.DroppedMembers}
	// A part whose content cannot be used says so whatever else happens.
	switch {
	case ob.SimRequestInvalid:
		res.SimRequest = OutboxInvalid
	case ob.SimRequest != nil:
		res.SimRequest = OutboxUnknownCharacter
	}
	switch {
	case ob.RosterInvalid:
		res.RosterSnapshot = OutboxInvalid
	case ob.RosterSnapshot != nil:
		res.RosterSnapshot = OutboxUnknownCharacter
	}
	exports, err := s.Store.Exports(ctx, userID)
	if err != nil {
		return nil, "", err
	}
	// The answers name their own characters, so they apply whatever becomes of the export.
	if len(ob.RaidAnswers) > 0 {
		if res.RaidAnswer, err = s.applyRaidAnswers(ctx, userID, exports, ob.RaidAnswers); err != nil {
			return nil, "", err
		}
	}
	own, ok := resolveOwn(exports, ob.Export.Character)
	if !ok {
		res.Export = OutboxUnknownCharacter
		return res, "", nil
	}
	own.Export = ob.Export.Code
	if err := s.Store.PutExports(ctx, userID, []Export{own}); err != nil {
		return nil, "", err
	}
	res.Export = OutboxStored
	key := character.Key(own.Region, character.RulesetFromRealm(own.Ruleset, ""), own.Name)
	if err := s.Store.SetScoreVisible(ctx, userID, key, ob.Export.scoreVisible()); err != nil {
		return nil, "", err
	}

	if ob.SimRequest != nil {
		if res.SimRequest, err = s.simRequest(ctx, userID, exports, *ob.SimRequest); err != nil {
			return nil, "", err
		}
	}
	if ob.RosterSnapshot != nil {
		if err := s.applyRoster(ctx, userID, own.Region, key, *ob.RosterSnapshot, res); err != nil {
			return nil, "", err
		}
	}
	return res, key, nil
}

// applyRoster stores the snapshot's members as the guild's roster and hands its opt-out codes
// to the directory, both only when the sender's character is in the guild the snapshot names.
// A snapshot of codes alone leaves the stored roster as it was.
func (s *Service) applyRoster(ctx context.Context, userID int64, region, key string, snap RosterSnapshot, res *OutboxResult) error {
	inGuild := false
	if len(snap.Members) > 0 || len(snap.OptOutCodes) == 0 {
		stored, err := s.Store.PutRosterSnapshot(ctx, userID, key, snap)
		if err != nil {
			return err
		}
		inGuild = stored
	} else {
		member, err := s.Store.InGuildNamed(ctx, userID, key, snap.Guild)
		if err != nil {
			return err
		}
		inGuild = member
	}
	res.RosterSnapshot = OutboxNotInGuild
	if inGuild {
		res.RosterSnapshot = OutboxStored
	}
	if !inGuild || len(snap.OptOutCodes) == 0 || s.OptOutCodes == nil {
		return nil
	}
	if _, err := s.OptOutCodes.VerifyNoteCodes(ctx, region, snap.Guild, snap.OptOutCodes); err != nil {
		s.Store.logger().Warn("addon", "op", "opt_out_codes", "err", err)
	}
	return nil
}

// InGuildNamed reports whether the account's character is in the guild of that name.
func (s *Store) InGuildNamed(ctx context.Context, userID int64, characterKey, guild string) (bool, error) {
	var member bool
	err := s.Pool.QueryRow(ctx,
		`select exists (select 1 from guild_characters gc join guilds g on g.id = gc.guild_id
		  where gc.character_key = $1 and gc.user_id = $2 and lower(g.name) = lower($3))`,
		characterKey, userID, guild).Scan(&member)
	if err != nil {
		return false, fmt.Errorf("addon: guild membership of %s: %w", characterKey, err)
	}
	return member, nil
}

// simRequest hands the request to SimRequests for the account's own
// character it names.
func (s *Service) simRequest(ctx context.Context, userID int64, exports []Export, req sims.AddonSimRequest) (string, error) {
	own, ok := resolveOwn(exports, req.Character)
	if !ok {
		return OutboxUnknownCharacter, nil
	}
	if s.SimRequests == nil {
		return sims.RequestUnavailable, nil
	}
	key := character.Key(own.Region, character.RulesetFromRealm(own.Ruleset, ""), own.Name)
	return s.SimRequests.RequestFromAddon(ctx, userID, key, req)
}

// SetScoreVisible stores whether the account's own character shows its Forever Score to others
// (characters.score_visible, migration 0092); the score itself is gated where it is read
// (api/internal/rating score_visibility.go).
func (s *Store) SetScoreVisible(ctx context.Context, userID int64, characterKey string, visible bool) error {
	if _, err := s.Pool.Exec(ctx,
		`update characters set score_visible = $3 where key = $1 and user_id = $2 and score_visible <> $3`,
		characterKey, userID, visible); err != nil {
		return fmt.Errorf("addon: store score visibility for %s: %w", characterKey, err)
	}
	return nil
}

// PutRosterSnapshot stores snap as the newest roster of the guild
// characterKey belongs to, when that guild is the one the snapshot names;
// it reports whether it stored anything. One row per guild: the newest
// snapshot from any member replaces the last.
func (s *Store) PutRosterSnapshot(ctx context.Context, userID int64, characterKey string, snap RosterSnapshot) (bool, error) {
	if snap.Members == nil {
		snap.Members = []RosterMember{}
	}
	members, err := json.Marshal(snap.Members)
	if err != nil {
		return false, fmt.Errorf("addon: encode roster snapshot for %s: %w", characterKey, err)
	}
	tag, err := s.Pool.Exec(ctx,
		`insert into guild_roster_snapshots (guild_id, character_key, members, received_at)
		 select gc.guild_id, gc.character_key, $3::jsonb, now()
		 from guild_characters gc join guilds g on g.id = gc.guild_id
		 where gc.character_key = $1 and gc.user_id = $2 and lower(g.name) = lower($4)
		 on conflict (guild_id) do update set
		   character_key = excluded.character_key, members = excluded.members,
		   received_at = excluded.received_at`,
		characterKey, userID, string(members), snap.Guild)
	if err != nil {
		return false, fmt.Errorf("addon: store roster snapshot for %s: %w", characterKey, err)
	}
	return tag.RowsAffected() > 0, nil
}
