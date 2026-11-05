// api/internal/guilds/loot.go
//
// GET /v1/guilds/{id}/loot, POST .../loot/awards, DELETE .../loot/awards/{award_id}
// (contract step 6, docs/contracts/2026-10-04-guild-centre-api.md): the loot council
// helper, Onyxia's own real 22-item table (data/builds/1.60.1.70630/loot.json's
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
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
)

// onyxiaEncounterID/onyxiaZone are progression.go's own constants, reused here under this
// file's own name for readability at the call sites below.
const onyxiaEncounterID = encounterOnyxia

// lootItem is one of Onyxia's own 22 drops: id, name and quality from
// data/builds/1.60.1.70630/items.json (every item carries a real row there; the 22 are unchanged
// since 1.60.1.70009); slot from the BiS build's (Store.BisBuild)
// data/builds/<build>/items/<class>.json where one of the nine per-class weighted item pools
// carries it. Icon is the client's own: Item.IconFileDataID through ManifestInterfaceData (wago.tools,
// build 1.60.1.70532), so the six quest, reputation and crafting items with no inventory slot (Scale
// of Onyxia, Onyxia Hide Backpack, both Head of Onyxia stacks, Mature Black Dragon Sinew and Draconic
// for Dummies) show their real art too; having no slot, they never gain a candidates list.
// ItemClass, Subclass, InventoryType, RequiredLevel and ItemLevel are items.json's class_id,
// subclass_id, inventory_type, required_level and item_level; AllowableClass is classic-db's
// item_template allowable_class (-1 for no restriction). loot_items_test.go pins all six.
type lootItem struct {
	ItemID         int
	Name           string
	Icon           string
	Quality        int
	Slot           string
	ItemClass      int
	Subclass       int
	InventoryType  int
	RequiredLevel  int
	ItemLevel      int
	AllowableClass int
	// Unique is item_template's unique flag: a character holds one.
	Unique bool
}

// onyxiaLoot is deliberately not read from loot.json/items.json/items/<class>.json at
// request time: those three files total over 40MB for this one build, and a per-request
// scan across them to resolve 22 fixed ids would blow the readiness-style latency budget
// for no reason - this table never changes without a data refresh this package's own tests
// would need updating for anyway, matching api/cmd/seedguild/readiness.go's own
// "a small, real subset... read by hand from that file" convention for exactly this kind
// of small, fixed catalogue.
var onyxiaLoot = []lootItem{
	{ItemID: 15410, Name: "Scale of Onyxia", Icon: "inv_misc_monsterscales_07", Quality: 3,
		ItemClass: 7, Subclass: 6, ItemLevel: 60, AllowableClass: -1},
	{ItemID: 16900, Name: "Stormrage Cover", Icon: "inv_helmet_09", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 2, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 1024},
	{ItemID: 16908, Name: "Bloodfang Hood", Icon: "inv_helmet_41", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 2, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 8},
	{ItemID: 16914, Name: "Netherwind Crown", Icon: "inv_helmet_70", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 1, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 128},
	{ItemID: 16921, Name: "Halo of Transcendence", Icon: "inv_helmet_24", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 1, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 16},
	{ItemID: 16929, Name: "Nemesis Skullcap", Icon: "inv_helmet_08", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 1, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 256},
	{ItemID: 16939, Name: "Dragonstalker's Helm", Icon: "inv_helmet_05", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 3, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 4},
	{ItemID: 16947, Name: "Helmet of Ten Storms", Icon: "inv_helmet_69", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 3, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 64},
	{ItemID: 16955, Name: "Judgement Crown", Icon: "inv_helmet_74", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 4, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 2},
	{ItemID: 16963, Name: "Helm of Wrath", Icon: "inv_helmet_71", Quality: 4, Slot: "head",
		ItemClass: 4, Subclass: 4, InventoryType: 1, RequiredLevel: 60, ItemLevel: 76, AllowableClass: 1},
	{ItemID: 17064, Name: "Shard of the Scale", Icon: "inv_misc_monsterscales_15", Quality: 4, Slot: "trinket",
		ItemClass: 4, Subclass: 0, InventoryType: 12, RequiredLevel: 60, ItemLevel: 71, AllowableClass: -1, Unique: true},
	{ItemID: 17067, Name: "Ancient Cornerstone Grimoire", Icon: "inv_misc_book_07", Quality: 4, Slot: "off_hand",
		ItemClass: 4, Subclass: 0, InventoryType: 23, RequiredLevel: 60, ItemLevel: 76, AllowableClass: -1},
	{ItemID: 17068, Name: "Deathbringer", Icon: "inv_axe_09", Quality: 4, Slot: "main_hand",
		ItemClass: 2, Subclass: 0, InventoryType: 13, RequiredLevel: 60, ItemLevel: 75, AllowableClass: -1},
	{ItemID: 17075, Name: "Vis'kag the Bloodletter", Icon: "inv_sword_18", Quality: 4, Slot: "main_hand",
		ItemClass: 2, Subclass: 7, InventoryType: 13, RequiredLevel: 60, ItemLevel: 74, AllowableClass: -1, Unique: true},
	{ItemID: 17078, Name: "Sapphiron Drape", Icon: "inv_misc_cape_16", Quality: 4, Slot: "back",
		ItemClass: 4, Subclass: 1, InventoryType: 16, RequiredLevel: 60, ItemLevel: 72, AllowableClass: -1},
	{ItemID: 17966, Name: "Onyxia Hide Backpack", Icon: "inv_misc_bag_22", Quality: 2,
		ItemClass: 1, Subclass: 0, InventoryType: 18, ItemLevel: 55, AllowableClass: -1},
	{ItemID: 18205, Name: "Eskhandar's Collar", Icon: "inv_belt_12", Quality: 4, Slot: "neck",
		ItemClass: 4, Subclass: 0, InventoryType: 2, RequiredLevel: 60, ItemLevel: 71, AllowableClass: -1, Unique: true},
	{ItemID: 18422, Name: "Head of Onyxia", Icon: "inv_misc_head_dragon_01", Quality: 4,
		ItemClass: 12, Subclass: 0, RequiredLevel: 60, ItemLevel: 60, AllowableClass: -1},
	{ItemID: 18423, Name: "Head of Onyxia", Icon: "inv_misc_head_dragon_01", Quality: 4,
		ItemClass: 12, Subclass: 0, RequiredLevel: 60, ItemLevel: 60, AllowableClass: -1},
	// classic-db's full item_template gives the Sinew AllowableClass 4 (hunter); the build's
	// narrowed raw/classicdb/item_template.json keeps equippable items only.
	{ItemID: 18705, Name: "Mature Black Dragon Sinew", Icon: "spell_holy_dispelmagic", Quality: 4,
		ItemClass: 12, Subclass: 0, RequiredLevel: 60, ItemLevel: 71, AllowableClass: 4},
	{ItemID: 18813, Name: "Ring of Binding", Icon: "inv_jewelry_ring_13", Quality: 4, Slot: "finger",
		ItemClass: 4, Subclass: 0, InventoryType: 11, RequiredLevel: 60, ItemLevel: 73, AllowableClass: -1, Unique: true},
	{ItemID: 21108, Name: "Draconic for Dummies", Icon: "inv_misc_stonetablet_03", Quality: 4,
		ItemClass: 15, Subclass: 0, ItemLevel: 1, AllowableClass: -1},
}

// onyxiaItem is the table's row for itemID.
func onyxiaItem(itemID int) (lootItem, bool) {
	for _, it := range onyxiaLoot {
		if it.ItemID == itemID {
			return it, true
		}
	}
	return lootItem{}, false
}

// AwardedTo is a loot item's own award, when one exists. AwardID is the row
// DELETE .../loot/awards/{award_id} takes, so the page can undo an award it has just read.
type AwardedTo struct {
	AwardID      int64  `json:"award_id"`
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	At           string `json:"at"`
	ByName       string `json:"by_name"`
}

// LootCandidate is one roster character this item names them a candidate for, ranked.
// Every candidate comes from one of two tiers (contract's own follow-up,
// docs/contracts/2026-10-04-guild-centre-api.md's loot section): a BiS-matched candidate
// (tier 0 - this character's own spec band names the item as its slot's pick or a listed
// alternative, GainDps a real number or nil when the band has the item but never measured
// a worn-vs-pick gap for it) or a fallback candidate (tier 1 - no band match, but the
// character's class can use the item at all and it is a real item-level upgrade over
// what they wear, or the slot is empty; GainDps always nil, NotSimChecked always true,
// IlvlDelta the item's own level advantage). tier itself is never marshaled - it is this
// struct's own sort key, read by lootCandidateLess.
type LootCandidate struct {
	CharacterKey      string     `json:"character_key"`
	Name              string     `json:"name"`
	Class             string     `json:"class"`
	Spec              string     `json:"spec"`
	GainDps           *float64   `json:"gain_dps"`
	NotSimChecked     bool       `json:"not_sim_checked"`
	IlvlDelta         *int       `json:"ilvl_delta"`
	Attendance        Attendance `json:"attendance"`
	AlreadyEquivalent bool       `json:"already_equivalent"`
	// Replaces names the item this character wears now in the drop's slot (empty when the slot is
	// empty or the item cannot be named); ItemsWon counts the awards this character has had in this
	// guild this tier (Onyxia is the only encounter an award can name, so every award row is the tier).
	// For a ring, a trinket or a one-hander a dual wielder could hold in either hand it is the weaker
	// of the two worn items, the one the drop pushes out.
	Replaces string `json:"replaces"`
	// ReplacesItem is the same worn item for the shared item link (icon, quality, tooltip); null for
	// an empty slot or one the class's item table cannot name.
	ReplacesItem *LootWornItem `json:"replaces_item"`
	ItemsWon     int           `json:"items_won"`
	// Tier is 0 for a BiS-matched candidate and 1 for one ranked by item level only.
	Tier int `json:"tier"`
	// GainUnit is what GainDps counts: "dps", "hps" (a healer band) or "tps" (a tank band); null
	// when GainDps is.
	GainUnit *string `json:"gain_unit"`
}

// LootWornItem is the item a candidate wears in the drop's slot.
type LootWornItem struct {
	ItemID  int    `json:"item_id"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Quality int    `json:"quality"`
}

// LootItemView is one of Onyxia's own 22 drops, as the Loot tab shows it.
type LootItemView struct {
	ItemID    int        `json:"item_id"`
	Name      string     `json:"name"`
	Icon      string     `json:"icon"`
	Quality   int        `json:"quality"`
	Slot      string     `json:"slot"`
	AwardedTo *AwardedTo `json:"awarded_to"`
	// Copies is how many of the table's drops this one row stands for: two items can share a
	// name (Onyxia's two "Head of Onyxia", the Alliance and Horde quest items), and the council
	// lists a name once.
	Copies     int             `json:"copies"`
	Candidates []LootCandidate `json:"candidates"`
	// InventoryType, ItemClass and ItemSubclass are the client's own (Item.InventoryType, ClassID,
	// SubclassID): the page words the slot line from them, never from Slot.
	InventoryType int `json:"inventory_type"`
	ItemClass     int `json:"item_class"`
	ItemSubclass  int `json:"item_subclass"`
	// Kind is "gear" (a slotted drop the council ranks), "quest", "bag" or "skinned" (Scale of
	// Onyxia, skinning loot that is never a pick).
	Kind string `json:"kind"`
	// QuestLabel is a quest drop's sourced line ("Quest: A Proper String (Hunter)"); "" otherwise.
	QuestLabel string `json:"quest_label"`
	// Eligible is who an officer can hand a quest drop or the bag to: verified level-60 characters
	// the quest is for. Empty for gear and for the skinned scale.
	Eligible []LootEligible `json:"eligible"`
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
	// AtLevelCap counts the verified roster characters at level 60, whose level is on file.
	AtLevelCap int `json:"at_level_cap"`
}

// Loot assembles the Loot tab for guildID. Onyxia is the only selectable encounter today
// (the only real, named boss this tier publishes) - any requested encounter id other than
// 1084 is ignored rather than answered with an empty page or an error, since there is
// nothing else to select yet (design spec's own "next unkilled: none, farm" copy). Killed
// counts the guild's private nights too, never a flagged report (presence.NotFlagged).
func (s *Store) Loot(ctx context.Context, guildID int64) (LootView, error) {
	var kills int
	if err := s.Pool.QueryRow(ctx, `
		select count(*) filter (where f.kill) from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.encounter_id = $2 and `+presence.NotFlagged("r"), guildID, onyxiaEncounterID).Scan(&kills); err != nil {
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
	// An account that joined by invite has no character row of the guild's, but its gear is on the
	// site and it takes part in the raid, so it is a candidate like any member.
	invited, err := s.invitedRosterRows(ctx, guildID, g)
	if err != nil {
		return LootView{}, err
	}
	roster = append(roster, invited...)
	all, err := s.lootAwards(ctx, guildID, onyxiaEncounterID)
	if err != nil {
		return LootView{}, err
	}
	now := time.Now()
	nightRaid, err := s.resolveAwardRaid(ctx, guildID, nil, now)
	if err != nil {
		return LootView{}, err
	}

	// An award to an invited member names the account's synthetic row, which no export carries a
	// name for; the picked character's name stands in for it.
	invitedName := make(map[string]string, len(invited))
	for _, row := range invited {
		invitedName[row.awardKey] = row.Name
	}
	for _, a := range all {
		if name, ok := invitedName[a.to.CharacterKey]; ok && a.to.Name == "" {
			a.to.Name = name
		}
	}

	awards := tonightsAwards(all, nightRaid, now)
	wonBy := awardCounts(all)
	view.AtLevelCap = countAtLevelCap(roster)

	// Two of the table's items can share a name (the Alliance and Horde "Head of Onyxia" quest
	// items, both slotless); the council lists a name once and says how many it stands for.
	for _, group := range lootGroups(onyxiaLoot) {
		iv := s.lootItemView(roster, group, g.Faction)
		for _, item := range group {
			if iv.AwardedTo == nil {
				iv.AwardedTo = awards[item.ItemID]
			}
		}
		for i := range iv.Candidates {
			iv.Candidates[i].ItemsWon = wonBy[iv.Candidates[i].CharacterKey]
		}
		view.Items = append(view.Items, iv)
	}
	return view, nil
}

// lootGroups is items grouped by name, in the table's order.
func lootGroups(items []lootItem) [][]lootItem {
	var out [][]lootItem
	index := map[string]int{}
	for _, item := range items {
		if i, ok := index[item.Name]; ok {
			out[i] = append(out[i], item)
			continue
		}
		index[item.Name] = len(out)
		out = append(out, []lootItem{item})
	}
	return out
}

// lootItemView is one row of the Loot tab for the items listed under one name. The row's item is
// the copy for the guild's faction when the copies are a faction's quest items, else the first.
func (s *Store) lootItemView(roster []RosterRow, group []lootItem, guildFaction *string) LootItemView {
	item := group[0]
	ids := make([]int, 0, len(group))
	eligibleCopies := group
	for _, it := range group {
		ids = append(ids, it.ItemID)
		if q, ok := questFor(it.ItemID); ok && guildFaction != nil && q.Faction == *guildFaction {
			item = it
			eligibleCopies = []lootItem{it}
		}
	}
	kind := lootKind(item)
	iv := LootItemView{
		ItemID: item.ItemID, Name: item.Name, Icon: item.Icon, Quality: item.Quality, Slot: item.Slot,
		Copies: len(group), InventoryType: item.InventoryType, ItemClass: item.ItemClass,
		ItemSubclass: item.Subclass, Kind: kind, Candidates: []LootCandidate{}, Eligible: []LootEligible{},
	}
	switch kind {
	case lootKindGear:
		iv.Candidates = s.lootCandidates(roster, item)
	case lootKindQuest:
		iv.QuestLabel = questLabel(ids)
		iv.Eligible = questEligibleAll(roster, eligibleCopies)
	case lootKindBag:
		iv.Eligible = questEligible(roster, item, kind)
	}
	return iv
}

// lootAward is one award row as the Loot tab reads it: the item, the raid it was filed under
// (nil for none yet) and when, with the line the page shows.
type lootAward struct {
	itemID    int
	raidID    *int64
	awardedAt time.Time
	to        *AwardedTo
}

// raidlessAwardNight is how long an award filed under no raid still marks its item awarded on
// the Loot tab: one raid night (the longest raid's match window, an hour before to six after a
// four-hour raid, is eleven hours).
const raidlessAwardNight = 12 * time.Hour

// lootAwards reads every award row for guildID's encounterID, oldest first, with the claimant's
// and the awarding officer's own names resolved. Every row is kept: the tier's count reads them
// all, and tonightsAwards picks the ones that mark an item awarded now.
func (s *Store) lootAwards(ctx context.Context, guildID, encounterID int64) ([]lootAward, error) {
	rows, err := s.Pool.Query(ctx, `
		select la.item_id, la.raid_id, la.id, la.character_key, coalesce(ae.name, ''), la.awarded_at, la.awarded_by
		from loot_awards la left join addon_exports ae on ae.character_key = la.character_key
		where la.guild_id = $1 and la.encounter_id = $2
		order by la.awarded_at, la.id`, guildID, encounterID)
	if err != nil {
		return nil, fmt.Errorf("guilds: loot awards: %w", err)
	}
	defer rows.Close()
	var out []lootAward
	for rows.Next() {
		var row lootAward
		var awardedBy *int64
		a := &AwardedTo{}
		if err := rows.Scan(&row.itemID, &row.raidID, &a.AwardID, &a.CharacterKey, &a.Name, &row.awardedAt, &awardedBy); err != nil {
			return nil, fmt.Errorf("guilds: loot awards: %w", err)
		}
		a.At = row.awardedAt.UTC().Format(time.RFC3339)
		if awardedBy != nil && s.Accounts != nil {
			a.ByName = s.awarderName(ctx, guildID, *awardedBy)
		}
		row.to = a
		out = append(out, row)
	}
	return out, rows.Err()
}

// tonightsAwards is each item's award on the night being played, keyed by item id: an award
// filed under nightRaid (the raid whose match window holds now, nil when none), or one filed
// under no raid within the last raidlessAwardNight. Last week's award leaves the item free to
// award again; the latest award of an item wins.
func tonightsAwards(awards []lootAward, nightRaid *int64, now time.Time) map[int]*AwardedTo {
	out := map[int]*AwardedTo{}
	for _, a := range awards {
		tonight := a.raidID == nil && now.Sub(a.awardedAt) <= raidlessAwardNight
		if a.raidID != nil && nightRaid != nil {
			tonight = *a.raidID == *nightRaid
		}
		if tonight {
			out[a.itemID] = a.to
		}
	}
	return out
}

// anOfficer is what an award line says when it cannot name the officer's character.
const anOfficer = "an officer"

// awarderName is how an award line names the officer who made it: the account's main
// character in THIS guild (its highest item level, ties by name, the same pick the roster
// makes for "alt of"), never a BattleTag. An account with no character in the guild, or one
// that has asked to be shown by pseudonym (users.anonymize, which the guild pages honour by
// never printing the account behind a character), reads "an officer" - the line names no
// one the guild cannot already see, and never the "user-12" pseudonym a stranger would be told.
func (s *Store) awarderName(ctx context.Context, guildID, userID int64) string {
	names, err := s.awarderNames(ctx, guildID, []int64{userID})
	if err != nil {
		return anOfficer
	}
	return names.of(userID)
}

// awarderNameSet is awarderNames' answer; of reads one awarder, "an officer" when unnamed.
type awarderNameSet map[int64]string

func (n awarderNameSet) of(userID int64) string {
	if name := n[userID]; name != "" {
		return name
	}
	return anOfficer
}

// awarderNames is awarderName for every user in userIDs in one query.
func (s *Store) awarderNames(ctx context.Context, guildID int64, userIDs []int64) (awarderNameSet, error) {
	out := awarderNameSet{}
	if len(userIDs) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `
		select u.id, coalesce((
		  select ae.name
		  from guild_characters gc
		  join addon_exports ae on ae.character_key = gc.character_key
		  where gc.guild_id = $1 and gc.user_id = u.id and gc.character_key not like 'account:%'
		  order by coalesce((
		    select m.ilvl from fight_metrics m
		    where m.player_key = gc.character_key order by m.fought_at desc limit 1), -1) desc, ae.name
		  limit 1), '')
		from users u where u.id = any($2) and not u.anonymize`, guildID, userIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: awarder names: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("guilds: awarder names: %w", err)
		}
		out[id] = name
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: awarder names: %w", err)
	}
	return out, nil
}

// ErrInvalidAward is returned for an award whose encounter/item is not one of Onyxia's
// own 22 drops, or whose character_key is not this guild's own roster.
var ErrInvalidAward = errors.New("guilds: invalid loot award")

func validOnyxiaItem(itemID int) bool {
	_, ok := onyxiaItem(itemID)
	return ok
}

// CreateLootAward records (or replaces - the partial unique index
// loot_awards_drop_unmatched_idx on (guild_id, encounter_id, item_id) where raid_id is null,
// migration 0052, means a second Award for the same drop updates the first row rather than
// leaving two) an officer's loot decision. raidID names the done raid it was for; nil files it
// under the raid being played now (loot_raid.go), else no raid until a match attaches it.
func (s *Store) CreateLootAward(ctx context.Context, guildID, encounterID int64, itemID int, characterKey string, awardedBy int64, raidID *int64) (int64, error) {
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
	raidID, err := s.resolveAwardRaid(ctx, guildID, raidID, time.Now())
	if err != nil {
		return 0, err
	}
	return s.upsertAward(ctx, guildID, encounterID, itemID, characterKey, awardedBy, raidID)
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
// proposed ruling: "Shown, read-only"). The gate is a verified row (IsVerifiedMember):
// killed counts the guild's private nights, so an unverified row gets 403 (task L6.12).
func (s *Service) loot(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	member, err := s.Store.IsVerifiedMember(r.Context(), guildID, actor.UserID)
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
	// RaidID optionally names the done raid the award was for.
	RaidID *int64 `json:"raid_id"`
}

// requireGuildPlan writes the refusal and returns false unless the officer's guild holds
// feature. It runs after the officer check, so a member never learns the plan state. A guild
// with no plan (or no Entitlements wired) gets guild_plan_required; a guild on a plan that
// covers the officer views but not feature (the Guild tier asking for Raid guild tools) gets
// raid_guild_required, which the web maps to the Raid guild sentence. A failed lookup is a
// server error, not a refusal.
func (s *Service) requireGuildPlan(w http.ResponseWriter, r *http.Request, guildID, userID int64, feature entitlements.Feature, message, op, failMessage string) bool {
	code := "guild_plan_required"
	if s.Entitlements != nil {
		ok, _, err := s.Entitlements.CanForGuild(r.Context(), guildID, userID, feature)
		if err != nil {
			s.fail(w, r, op, err, failMessage)
			return false
		}
		if ok {
			return true
		}
		if feature != entitlements.FeatureOfficerViews {
			covered, _, err := s.Entitlements.CanForGuild(r.Context(), guildID, userID, entitlements.FeatureOfficerViews)
			if err != nil {
				s.fail(w, r, op, err, failMessage)
				return false
			}
			if covered {
				code = "raid_guild_required"
			}
		}
	}
	httpx.WriteError(w, r, http.StatusForbidden, code, message, nil)
	return false
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
	if !s.requireGuildPlan(w, r, guildID, actor.UserID, entitlements.FeatureRaidGuildTools, refusal.AwardPlanRequired, "loot_award", "could not record that award just now") {
		return
	}
	if !s.officerNotFrozen(w, r, guildID, "loot_award", "could not record that award just now") {
		return
	}
	var in lootAwardInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	id, err := s.Store.CreateLootAward(r.Context(), guildID, in.EncounterID, in.ItemID, in.CharacterKey, actor.UserID, in.RaidID)
	var badRaid *AwardRaidError
	switch {
	case errors.As(err, &badRaid):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			fmt.Sprintf("raid %d is not a done Onyxia raid of this guild", badRaid.RaidID), nil)
	case errors.Is(err, ErrInvalidAward):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a valid encounter/item/character for this guild", nil)
	case err != nil:
		s.fail(w, r, "loot_award", err, "could not record that award just now")
	default:
		s.logger().Info("guilds", "op", "loot_award", "guild_id", guildID, "item_id", in.ItemID,
			"character_key", in.CharacterKey, "user_id", actor.UserID)
		s.publishLootAwarded(r.Context(), guildID, id, in.CharacterKey, in.ItemID)
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
	if !s.requireGuildPlan(w, r, guildID, actor.UserID, entitlements.FeatureRaidGuildTools, refusal.AwardPlanRequired, "loot_award_delete", "could not remove that award just now") {
		return
	}
	if !s.officerNotFrozen(w, r, guildID, "loot_award_delete", "could not remove that award just now") {
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

// awardCounts is how many of the tier's drops each character has been awarded in this guild,
// every award row counted (two weeks of the same helm are two items). An award row names one
// encounter (Onyxia, the only one an award can name), so every row is the current tier and the
// count needs no table of its own.
func awardCounts(awards []lootAward) map[string]int {
	out := make(map[string]int, len(awards))
	for _, a := range awards {
		out[a.to.CharacterKey]++
	}
	return out
}

// attendanceRatio is present over nights, 0 for a character with no nights.
func attendanceRatio(a Attendance) float64 {
	if a.Nights == 0 {
		return 0
	}
	return float64(a.Present) / float64(a.Nights)
}
