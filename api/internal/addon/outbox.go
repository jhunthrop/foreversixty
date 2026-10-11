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
)

// Outbox is contract 4's format 1.
type Outbox struct {
	Format         int                   `json:"format"`
	Export         *OutboxExport         `json:"export"`
	SimRequest     *sims.AddonSimRequest `json:"sim_request"`
	RosterSnapshot *RosterSnapshot       `json:"roster_snapshot"`
}

// OutboxExport is the outbox's export: Character is Name-Realm, Code the
// FS1 export string, SavedAt the player's local time.
type OutboxExport struct {
	Character string `json:"character"`
	Code      string `json:"code"`
	SavedAt   string `json:"saved_at"`
}

// RosterSnapshot is the guild tab's live roster as one member's addon saw it.
type RosterSnapshot struct {
	Guild   string         `json:"guild"`
	Members []RosterMember `json:"members"`
}

// RosterMember is one member line: contract 2's hello fields plus name and
// seen_at.
type RosterMember struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Spec    string `json:"spec"`
	GearGap int    `json:"gear_gap"`
	Build   string `json:"build"`
	SeenAt  string `json:"seen_at"`
}

// OutboxResult says what became of each part of the outbox. A part the
// outbox did not carry is absent.
type OutboxResult struct {
	Status         string `json:"status"`
	Export         string `json:"export,omitempty"`
	SimRequest     string `json:"sim_request,omitempty"`
	RosterSnapshot string `json:"roster_snapshot,omitempty"`
}

// readOutbox reads the outbox's format and, for format 1, the outbox
// itself, writing a 400 for one that is not an outbox at all. A nil
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
	var ob Outbox
	if err := json.Unmarshal(raw, &ob); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that outbox cannot be read",
			map[string]string{"outbox": err.Error()})
		return nil, false
	}
	if problem := ob.problem(); problem != "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that outbox cannot be stored",
			map[string]string{"outbox": problem})
		return nil, false
	}
	return &ob, true
}

// problem is the first thing wrong with a format 1 outbox, or "". It also
// normalises the snapshot's guild name the way an export's guild= is.
func (ob *Outbox) problem() string {
	e := ob.Export
	if e == nil || e.Code == "" || len(e.Code) > MaxExportLen {
		return "export needs a character and a code of at most 4096 characters"
	}
	if !validUnit(e.Character) {
		return "export.character must be Name-Realm"
	}
	if sr := ob.SimRequest; sr != nil {
		if len(sr.GearCode) > MaxExportLen {
			return "sim_request.gear_code is at most 4096 characters"
		}
		if !validUnit(sr.Character) {
			return "sim_request.character must be Name-Realm"
		}
		if err := sr.Validate(); err != nil {
			return err.Error()
		}
	}
	if rs := ob.RosterSnapshot; rs != nil {
		name, ok := validateGuildName(rs.Guild)
		if !ok {
			return "roster_snapshot.guild must be a guild name"
		}
		rs.Guild = name
		if len(rs.Members) > maxRosterMembers {
			return "roster_snapshot.members holds at most 1000 members"
		}
		for _, m := range rs.Members {
			if problem := m.problem(); problem != "" {
				return problem
			}
		}
	}
	return ""
}

func (m RosterMember) problem() string {
	if !plainText(m.Name) || !validUnit(m.Name) {
		return "roster_snapshot.members[].name must be Name-Realm"
	}
	if m.Level < 1 || m.Level > simapi.MaxLevel {
		return "roster_snapshot.members[].level must be between 1 and 60"
	}
	if m.Spec != "" && !plainText(m.Spec) {
		return "roster_snapshot.members[].spec is a spec slug"
	}
	if m.Build != "" && !plainText(m.Build) {
		return "roster_snapshot.members[].build is a plain build name of at most 64 characters"
	}
	if m.GearGap < 0 {
		return "roster_snapshot.members[].gear_gap is not negative"
	}
	if _, err := time.Parse(sims.RequestedAtLayout, m.SeenAt); err != nil {
		return "roster_snapshot.members[].seen_at must be YYYY-MM-DD HH:MM"
	}
	return ""
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
// character, starts its sim request and stores its roster snapshot. A nil
// outbox is one of a format this API does not read. Anything about the
// outbox's content is answered in the result; only a store failure (or a
// claimed character) is an error.
func (s *Service) applyOutbox(ctx context.Context, userID int64, ob *Outbox) (*OutboxResult, error) {
	if ob == nil {
		return &OutboxResult{Status: OutboxIgnored}, nil
	}
	res := &OutboxResult{Status: OutboxRead}
	exports, err := s.Store.Exports(ctx, userID)
	if err != nil {
		return nil, err
	}
	own, ok := resolveOwn(exports, ob.Export.Character)
	if !ok {
		res.Export = OutboxUnknownCharacter
		if ob.SimRequest != nil {
			res.SimRequest = OutboxUnknownCharacter
		}
		if ob.RosterSnapshot != nil {
			res.RosterSnapshot = OutboxUnknownCharacter
		}
		return res, nil
	}
	own.Export = ob.Export.Code
	if err := s.Store.PutExports(ctx, userID, []Export{own}); err != nil {
		return nil, err
	}
	res.Export = OutboxStored
	key := character.Key(own.Region, character.RulesetFromRealm(own.Ruleset, ""), own.Name)

	if ob.SimRequest != nil {
		if res.SimRequest, err = s.simRequest(ctx, userID, exports, *ob.SimRequest); err != nil {
			return nil, err
		}
	}
	if ob.RosterSnapshot != nil {
		stored, err := s.Store.PutRosterSnapshot(ctx, userID, key, *ob.RosterSnapshot)
		if err != nil {
			return nil, err
		}
		res.RosterSnapshot = OutboxNotInGuild
		if stored {
			res.RosterSnapshot = OutboxStored
		}
	}
	return res, nil
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
