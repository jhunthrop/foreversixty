// api/internal/guilds/loot.go
//
// GET /v1/guilds/{id}/loot, POST .../loot/awards, DELETE .../loot/awards/{award_id}
// (contract step 6, docs/contracts/2026-10-04-guild-centre-api.md): the loot council
// helper, Onyxia's own real 22-item table (data/builds/1.60.1.70009/loot.json's
// raid:onyxias-lair, npc 10184). Member read-only, officer gets the Award control (design
// spec §4.F / §10.1's own proposed ruling, recorded as "Shown, read-only" for a member).
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// onyxiaEncounterID/onyxiaZone are progression.go's own constants, reused here under this
// file's own name for readability at the call sites below.
const onyxiaEncounterID = encounterOnyxia

// lootItem is one of Onyxia's own 22 drops: id, name and quality from
// data/builds/1.60.1.70009/items.json (every item carries a real row there); icon and slot
// from that same build's data/builds/<build>/items/<class>.json where one of the nine
// per-class weighted item pools happens to carry it (sixteen of the twenty-two do - the
// other six are quest/reputation/crafting items with no inventory slot at all: Scale of
// Onyxia, Onyxia Hide Backpack, both Head of Onyxia stacks, Mature Black Dragon Sinew and
// Draconic for Dummies, which is exactly why they carry no slot below and therefore never
// gain a candidates list - a quest item has no "best in slot," correctly never one here).
type lootItem struct {
	ItemID  int
	Name    string
	Icon    string
	Quality int
	Slot    string
}

// onyxiaLoot is deliberately not read from loot.json/items.json/items/<class>.json at
// request time: those three files total over 40MB for this one build, and a per-request
// scan across them to resolve 22 fixed ids would blow the readiness-style latency budget
// for no reason - this table never changes without a data refresh this package's own tests
// would need updating for anyway, matching api/cmd/seedguild/readiness.go's own
// "a small, real subset... read by hand from that file" convention for exactly this kind
// of small, fixed catalogue.
var onyxiaLoot = []lootItem{
	{15410, "Scale of Onyxia", "", 3, ""},
	{16900, "Stormrage Cover", "inv_helmet_09", 4, "head"},
	{16908, "Bloodfang Hood", "inv_helmet_41", 4, "head"},
	{16914, "Netherwind Crown", "inv_helmet_70", 4, "head"},
	{16921, "Halo of Transcendence", "inv_helmet_24", 4, "head"},
	{16929, "Nemesis Skullcap", "inv_helmet_08", 4, "head"},
	{16939, "Dragonstalker's Helm", "inv_helmet_05", 4, "head"},
	{16947, "Helmet of Ten Storms", "inv_helmet_69", 4, "head"},
	{16955, "Judgement Crown", "inv_helmet_74", 4, "head"},
	{16963, "Helm of Wrath", "inv_helmet_71", 4, "head"},
	{17064, "Shard of the Scale", "inv_misc_monsterscales_15", 4, "trinket"},
	{17067, "Ancient Cornerstone Grimoire", "inv_misc_book_07", 4, "off_hand"},
	{17068, "Deathbringer", "inv_axe_09", 4, "main_hand"},
	{17075, "Vis'kag the Bloodletter", "inv_sword_18", 4, "main_hand"},
	{17078, "Sapphiron Drape", "inv_misc_cape_16", 4, "back"},
	{17966, "Onyxia Hide Backpack", "", 2, ""},
	{18205, "Eskhandar's Collar", "inv_belt_12", 4, "neck"},
	{18422, "Head of Onyxia", "", 4, ""},
	{18423, "Head of Onyxia", "", 4, ""},
	{18705, "Mature Black Dragon Sinew", "", 4, ""},
	{18813, "Ring of Binding", "inv_jewelry_ring_13", 4, "finger"},
	{21108, "Draconic for Dummies", "", 4, ""},
}

// AwardedTo is a loot item's own award, when one exists.
type AwardedTo struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	At           string `json:"at"`
	ByName       string `json:"by_name"`
}

// LootCandidate is one roster character this item's slot band names, ranked.
type LootCandidate struct {
	CharacterKey      string     `json:"character_key"`
	Name              string     `json:"name"`
	Class             string     `json:"class"`
	Spec              string     `json:"spec"`
	GainDps           float64    `json:"gain_dps"`
	NotSimChecked     bool       `json:"not_sim_checked"`
	Attendance        Attendance `json:"attendance"`
	AlreadyEquivalent bool       `json:"already_equivalent"`
}

// LootItemView is one of Onyxia's own 22 drops, as the Loot tab shows it.
type LootItemView struct {
	ItemID     int             `json:"item_id"`
	Name       string          `json:"name"`
	Icon       string          `json:"icon"`
	Quality    int             `json:"quality"`
	Slot       string          `json:"slot"`
	AwardedTo  *AwardedTo      `json:"awarded_to"`
	Candidates []LootCandidate `json:"candidates"`
}

// LootEncounterOption is one entry of the boss picker.
type LootEncounterOption struct {
	EncounterID int64  `json:"encounter_id"`
	Name        string `json:"name"`
	Zone        string `json:"zone"`
	Killed      bool   `json:"killed"`
}

// LootView is the body of GET /v1/guilds/{id}/loot.
type LootView struct {
	Encounters []LootEncounterOption `json:"encounters"`
	Selected   int64                 `json:"selected"`
	Items      []LootItemView        `json:"items"`
}

// Loot assembles the Loot tab for guildID. Onyxia is the only selectable encounter today
// (the only real, named boss this tier publishes) - any requested encounter id other than
// 1084 is ignored rather than answered with an empty page or an error, since there is
// nothing else to select yet (design spec's own "next unkilled: none, farm" copy).
func (s *Store) Loot(ctx context.Context, guildID int64) (LootView, error) {
	var kills int
	if err := s.Pool.QueryRow(ctx, `
		select count(*) filter (where f.kill) from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.encounter_id = $2`, guildID, onyxiaEncounterID).Scan(&kills); err != nil {
		return LootView{}, fmt.Errorf("guilds: loot kills: %w", err)
	}

	view := LootView{
		Encounters: []LootEncounterOption{{EncounterID: onyxiaEncounterID, Name: encounterName, Zone: zoneOnyxia, Killed: kills > 0}},
		Selected:   onyxiaEncounterID,
		Items:      []LootItemView{},
	}

	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return LootView{}, err
	}
	roster, err := s.HomeRoster(ctx, guildID, g, 0, false, false)
	if err != nil {
		return LootView{}, err
	}
	present, err := s.attendancePresent(ctx, guildID)
	if err != nil {
		return LootView{}, err
	}
	nights, err := s.attendanceWindowSize(ctx, guildID)
	if err != nil {
		return LootView{}, err
	}

	awards, err := s.lootAwards(ctx, guildID, onyxiaEncounterID)
	if err != nil {
		return LootView{}, err
	}

	for _, item := range onyxiaLoot {
		iv := LootItemView{ItemID: item.ItemID, Name: item.Name, Icon: item.Icon, Quality: item.Quality, Slot: item.Slot}
		iv.AwardedTo = awards[item.ItemID]
		iv.Candidates = s.lootCandidates(roster, item, present, nights)
		view.Items = append(view.Items, iv)
	}
	return view, nil
}

// lootCandidates ranks every verified roster character whose own BiS band names item for
// its own slot - the pick or any listed alternative - by gain_dps desc, then attendance
// desc. A slotless item (a quest/reputation/crafting drop, six of Onyxia's twenty-two)
// never has candidates: no character's BiS band recommends a quest item.
func (s *Store) lootCandidates(roster []RosterRow, item lootItem, present map[string]int, nights int) []LootCandidate {
	out := []LootCandidate{}
	if item.Slot == "" {
		return out
	}
	for _, row := range roster {
		if !row.Verified || row.className == "" || row.specStr == "" {
			continue
		}
		band, hasBand := s.loadBandFor(row.className, row.specStr, row.faction)
		if !hasBand {
			continue
		}
		pick, named := band.BySlot()[item.Slot]
		if !named || pick.ItemID <= 0 {
			continue
		}
		if !namesItem(pick, item.ItemID) {
			continue
		}
		wornID, hasWorn := row.gear[item.Slot]
		verdict := bis.VerdictFor(pick, wornID, hasWorn)
		cand := LootCandidate{
			CharacterKey: row.CharacterKey, Name: row.Name,
			AlreadyEquivalent: !verdict.Upgrade,
			NotSimChecked:     verdict.NotSimChecked,
			Attendance:        Attendance{Present: present[row.CharacterKey], Nights: nights},
		}
		if row.Class != nil {
			cand.Class = *row.Class
		}
		if row.Spec != nil {
			cand.Spec = *row.Spec
		}
		if verdict.GainDps != nil {
			cand.GainDps = *verdict.GainDps
		}
		out = append(out, cand)
	}
	sortLootCandidates(out)
	return out
}

// namesItem reports whether itemID is pick's own top choice or one of its listed
// alternatives - the membership test a loot drop must pass to be "this character's own
// candidate" at all.
func namesItem(pick bis.Slot, itemID int) bool {
	if pick.ItemID == itemID {
		return true
	}
	for _, alt := range pick.Alternatives {
		if alt.ItemID == itemID {
			return true
		}
	}
	return false
}

func sortLootCandidates(cands []LootCandidate) {
	for i := 1; i < len(cands); i++ {
		for j := i; j > 0 && lootCandidateLess(cands[j], cands[j-1]); j-- {
			cands[j], cands[j-1] = cands[j-1], cands[j]
		}
	}
}

// lootCandidateLess reports whether a ranks ahead of b: gain_dps desc, then attendance
// (present) desc.
func lootCandidateLess(a, b LootCandidate) bool {
	if a.GainDps != b.GainDps {
		return a.GainDps > b.GainDps
	}
	return a.Attendance.Present > b.Attendance.Present
}

// lootAwards reads every current award for guildID's encounterID, keyed by item id, with
// the claimant's and the awarding officer's own names resolved.
func (s *Store) lootAwards(ctx context.Context, guildID, encounterID int64) (map[int]*AwardedTo, error) {
	rows, err := s.Pool.Query(ctx, `
		select la.item_id, la.character_key, coalesce(ae.name, ''), la.awarded_at, la.awarded_by
		from loot_awards la left join addon_exports ae on ae.character_key = la.character_key
		where la.guild_id = $1 and la.encounter_id = $2`, guildID, encounterID)
	if err != nil {
		return nil, fmt.Errorf("guilds: loot awards: %w", err)
	}
	defer rows.Close()
	out := map[int]*AwardedTo{}
	for rows.Next() {
		var itemID int
		var awardedBy *int64
		var a AwardedTo
		var awardedAt time.Time
		if err := rows.Scan(&itemID, &a.CharacterKey, &a.Name, &awardedAt, &awardedBy); err != nil {
			return nil, fmt.Errorf("guilds: loot awards: %w", err)
		}
		a.At = awardedAt.UTC().Format(time.RFC3339)
		if awardedBy != nil && s.Accounts != nil {
			if u, uerr := s.Accounts.User(ctx, *awardedBy); uerr == nil {
				a.ByName = u.PublicName()
			}
		}
		out[itemID] = &a
	}
	return out, rows.Err()
}

// ErrInvalidAward is returned for an award whose encounter/item is not one of Onyxia's
// own 22 drops, or whose character_key is not this guild's own roster.
var ErrInvalidAward = errors.New("guilds: invalid loot award")

func validOnyxiaItem(itemID int) bool {
	for _, it := range onyxiaLoot {
		if it.ItemID == itemID {
			return true
		}
	}
	return false
}

// CreateLootAward records (or replaces - the migration's own unique index on
// (guild_id, encounter_id, item_id) means a second Award for the same drop updates the
// first row rather than leaving two) an officer's loot decision.
func (s *Store) CreateLootAward(ctx context.Context, guildID, encounterID int64, itemID int, characterKey string, awardedBy int64) (int64, error) {
	if encounterID != onyxiaEncounterID || !validOnyxiaItem(itemID) {
		return 0, ErrInvalidAward
	}
	var exists bool
	if err := s.Pool.QueryRow(ctx,
		`select exists(select 1 from guild_characters where guild_id = $1 and character_key = $2)`,
		guildID, characterKey).Scan(&exists); err != nil {
		return 0, fmt.Errorf("guilds: create loot award: %w", err)
	}
	if !exists {
		return 0, ErrInvalidAward
	}
	var id int64
	if err := s.Pool.QueryRow(ctx, `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by)
		values ($1, $2, $3, $4, $5)
		on conflict (guild_id, encounter_id, item_id) do update set
		  character_key = excluded.character_key, awarded_by = excluded.awarded_by,
		  awarded_at = now(), report_id = null
		returning id`, guildID, encounterID, itemID, characterKey, awardedBy).Scan(&id); err != nil {
		return 0, fmt.Errorf("guilds: create loot award: %w", err)
	}
	return id, nil
}

// DeleteLootAward removes one award, scoped to guildID so one guild can never delete
// another's row even if it somehow guessed the id.
func (s *Store) DeleteLootAward(ctx context.Context, guildID, awardID int64) error {
	tag, err := s.Pool.Exec(ctx, `delete from loot_awards where id = $1 and guild_id = $2`, awardID, guildID)
	if err != nil {
		return fmt.Errorf("guilds: delete loot award: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// loot is member-and-officer, identical content either way - the contract's own "member:
// read-only" is about which mutation routes an officer additionally gets, never about
// withholding any of this read's own data from a plain member (design spec §10.1's
// proposed ruling: "Shown, read-only").
func (s *Service) loot(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	member, err := s.Store.IsMember(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "loot", err, "could not load that guild's loot just now")
		return
	}
	if !member {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you are not a member of that guild", nil)
		return
	}
	view, err := s.Store.Loot(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "loot", err, "could not load that guild's loot just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

type lootAwardInput struct {
	EncounterID  int64  `json:"encounter_id"`
	ItemID       int    `json:"item_id"`
	CharacterKey string `json:"character_key"`
}

func (s *Service) createLootAward(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "loot_award", err, "could not record that award just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to award loot", nil)
		return
	}
	if frozen, err := s.freezeCheck(r.Context(), guildID, actor.IsModerator()); err != nil {
		s.fail(w, r, "loot_award", err, "could not record that award just now")
		return
	} else if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
	var in lootAwardInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	id, err := s.Store.CreateLootAward(r.Context(), guildID, in.EncounterID, in.ItemID, in.CharacterKey, actor.UserID)
	switch {
	case errors.Is(err, ErrInvalidAward):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a valid encounter/item/character for this guild", nil)
	case err != nil:
		s.fail(w, r, "loot_award", err, "could not record that award just now")
	default:
		s.logger().Info("guilds", "op", "loot_award", "guild_id", guildID, "item_id", in.ItemID,
			"character_key", in.CharacterKey, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{"award_id": id, "status": "awarded"})
	}
}

func (s *Service) deleteLootAward(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	awardID, err := strconv.ParseInt(r.PathValue("award_id"), 10, 64)
	if err != nil || awardID <= 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such award", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "loot_award_delete", err, "could not remove that award just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to remove a loot award", nil)
		return
	}
	if frozen, err := s.freezeCheck(r.Context(), guildID, actor.IsModerator()); err != nil {
		s.fail(w, r, "loot_award_delete", err, "could not remove that award just now")
		return
	} else if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
	switch err := s.Store.DeleteLootAward(r.Context(), guildID, awardID); {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such award", nil)
	case err != nil:
		s.fail(w, r, "loot_award_delete", err, "could not remove that award just now")
	default:
		s.logger().Info("guilds", "op", "loot_award_delete", "guild_id", guildID, "award_id", awardID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "removed"})
	}
}
