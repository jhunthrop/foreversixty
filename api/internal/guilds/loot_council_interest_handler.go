// api/internal/guilds/loot_council_interest_handler.go
//
// GET and PUT /v1/guilds/{id}/loot/interest: a verified member reads and sets the MS/OS flags of
// their own characters; an officer may set anyone's (the officer write gate applies then).
package guilds

import (
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// LootInterestView is the body of GET /v1/guilds/{id}/loot/interest: the viewer's own flags on
// the night.
type LootInterestView struct {
	Night string         `json:"night"`
	Flags []LootInterest `json:"flags"`
}

type lootInterestInput struct {
	lootDropInput
	// MsOs is "ms" or "os"; null removes the flag.
	MsOs *string `json:"ms_os"`
}

// councilMember reads the guild id and checks the caller is a verified member.
func (s *Service) councilMember(w http.ResponseWriter, r *http.Request, op, failMessage string) (int64, bool) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return 0, false
	}
	member, err := s.Store.IsVerifiedMember(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, op, err, failMessage)
		return 0, false
	}
	if !member {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you are not a member of that guild", nil)
		return 0, false
	}
	return guildID, true
}

func (s *Service) getLootInterest(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.councilMember(w, r, "loot_interest", councilReadFail)
	if !ok {
		return
	}
	night, ok := s.councilNightFrom(w, r, guildID, r.URL.Query().Get("night"), "loot_interest", councilReadFail)
	if !ok {
		return
	}
	flags, err := s.Store.interestOn(r.Context(), guildID, night, 0, auth.ActorFrom(r.Context()).UserID, time.Now())
	if err != nil {
		s.fail(w, r, "loot_interest", err, councilReadFail)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, LootInterestView{Night: night.key(), Flags: flags})
}

func (s *Service) putLootInterest(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.councilMember(w, r, "loot_interest", councilFlagFail)
	if !ok {
		return
	}
	var in lootInterestInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	msOs := ""
	if in.MsOs != nil {
		if !validMsOs(*in.MsOs) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", `ms_os must be "ms", "os" or null`, nil)
			return
		}
		msOs = *in.MsOs
	}
	if !s.mayFlag(w, r, guildID, in.CharacterKey) {
		return
	}
	night, ok := s.councilNightFrom(w, r, guildID, in.Night, "loot_interest", councilFlagFail)
	if !ok {
		return
	}
	actor := auth.ActorFrom(r.Context())
	err := s.Store.SetLootInterest(r.Context(), guildID, night, in.ItemID, in.CharacterKey, msOs, actor.UserID, time.Now())
	if !s.councilWriteOK(w, r, err, "loot_interest", councilFlagFail) {
		return
	}
	s.logger().Info("guilds", "op", "loot_interest", "guild_id", guildID, "item_id", in.ItemID,
		"character_key", in.CharacterKey, "user_id", actor.UserID, "ms_os", msOs)
	var flag *string
	if msOs != "" {
		flag = &msOs
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"night": night.key(), "item_id": in.ItemID,
		"character_key": in.CharacterKey, "ms_os": flag})
}

// mayFlag lets the character's own account set its flag; anyone else must be an officer past the
// officer write gate.
func (s *Service) mayFlag(w http.ResponseWriter, r *http.Request, guildID int64, characterKey string) bool {
	own, err := s.Store.ownsCharacter(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, characterKey)
	if err != nil {
		s.fail(w, r, "loot_interest", err, councilFlagFail)
		return false
	}
	if own {
		return true
	}
	officer, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "loot_interest", err, councilFlagFail)
		return false
	}
	if !officer {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you can only set main or off spec for your own characters", nil)
		return false
	}
	return s.councilOfficerWrite(w, r, guildID, "loot_interest", councilFlagFail)
}
