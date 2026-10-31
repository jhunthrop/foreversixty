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
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
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
	{15410, "Scale of Onyxia", "inv_misc_monsterscales_07", 3, ""},
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
	{17966, "Onyxia Hide Backpack", "inv_misc_bag_22", 2, ""},
	{18205, "Eskhandar's Collar", "inv_belt_12", 4, "neck"},
	{18422, "Head of Onyxia", "inv_misc_head_dragon_01", 4, ""},
	{18423, "Head of Onyxia", "inv_misc_head_dragon_01", 4, ""},
	{18705, "Mature Black Dragon Sinew", "spell_holy_dispelmagic", 4, ""},
	{18813, "Ring of Binding", "inv_jewelry_ring_13", 4, "finger"},
	{21108, "Draconic for Dummies", "inv_misc_stonetablet_03", 4, ""},
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
	Replaces string `json:"replaces"`
	ItemsWon int    `json:"items_won"`
	tier     int
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
	awards, err := s.lootAwards(ctx, guildID, onyxiaEncounterID)
	if err != nil {
		return LootView{}, err
	}

	// An award to an invited member names the account's synthetic row, which no export carries a
	// name for; the picked character's name stands in for it.
	invitedName := make(map[string]string, len(invited))
	for _, row := range invited {
		invitedName[row.awardKey] = row.Name
	}
	for _, a := range awards {
		if name, ok := invitedName[a.CharacterKey]; ok && a.Name == "" {
			a.Name = name
		}
	}

	wonBy := awardCounts(awards)

	for _, item := range onyxiaLoot {
		// Two of the table's items can share a name (the Alliance and Horde "Head of Onyxia" quest
		// items, both slotless); the council lists a name once and says how many it stands for.
		if i := indexOfLootName(view.Items, item.Name); i >= 0 {
			view.Items[i].Copies++
			if view.Items[i].AwardedTo == nil {
				view.Items[i].AwardedTo = awards[item.ItemID]
			}
			continue
		}
		iv := LootItemView{ItemID: item.ItemID, Name: item.Name, Icon: item.Icon, Quality: item.Quality, Slot: item.Slot, Copies: 1}
		iv.AwardedTo = awards[item.ItemID]
		iv.Candidates = s.lootCandidates(roster, item)
		for i := range iv.Candidates {
			iv.Candidates[i].ItemsWon = wonBy[iv.Candidates[i].CharacterKey]
		}
		view.Items = append(view.Items, iv)
	}
	return view, nil
}

// indexOfLootName is the position of the item already listed under name, or -1.
func indexOfLootName(items []LootItemView, name string) int {
	for i := range items {
		if items[i].Name == name {
			return i
		}
	}
	return -1
}

// lootCandidates ranks every verified roster character this item names a candidate for,
// in two tiers (see LootCandidate's own doc comment): tier 0, this character's own BiS
// band names the item for its own slot (the pick or a listed alternative); tier 1
// (follow-up fix - loot candidates must not read empty on real data, since no BiS band in
// this repo today names any raid-tier item at all, see CONTROL_CENTRE.md), no band match
// but the character's class can use the item at all (data/builds/<build>/items/<class>.
// json, read through api/internal/trees the same way every other build-scoped reader in
// this service already loads it) and it is a real item-level upgrade over what they wear,
// or the slot is empty. A slotless item (a quest/reputation/crafting drop, six of Onyxia's
// twenty-two) never has candidates in either tier: there is no slot to compare against.
func (s *Store) lootCandidates(roster []RosterRow, item lootItem) []LootCandidate {
	out := []LootCandidate{}
	if item.Slot == "" {
		return out
	}
	for _, row := range roster {
		if !row.Verified || row.className == "" {
			continue
		}
		key := row.CharacterKey
		if row.awardKey != "" {
			key = row.awardKey
		}
		cand := LootCandidate{
			CharacterKey: key, Name: row.Name,
			Attendance: row.Attendance,
		}
		cand.Replaces = s.wornName(row, item.Slot)
		if row.Class != nil {
			cand.Class = *row.Class
		}
		if row.Spec != nil {
			cand.Spec = *row.Spec
		}

		if verdict, ok := s.bisMatch(row, item); ok {
			cand.tier = 0
			cand.GainDps = verdict.GainDps
			cand.NotSimChecked = verdict.NotSimChecked
			cand.AlreadyEquivalent = !verdict.Upgrade
			out = append(out, cand)
			continue
		}

		if delta, ok := s.fallbackMatch(row, item); ok {
			cand.tier = 1
			cand.NotSimChecked = true
			cand.IlvlDelta = &delta
			out = append(out, cand)
		}
	}
	sortLootCandidates(out)
	return out
}

// wornName is the name of the item row wears in slot, or "" when the slot is empty or the item is
// not one this class's item table can name.
func (s *Store) wornName(row RosterRow, slot string) string {
	wornID, ok := row.gear[slot]
	if !ok {
		return ""
	}
	if it, found := s.buildItemFor(row.className, wornID); found {
		return it.Name
	}
	return ""
}

// bisMatch is lootCandidates' own tier-0 test: row's spec band names item for its own slot
// (the pick or a listed alternative), scored the normal gear-gap way.
func (s *Store) bisMatch(row RosterRow, item lootItem) (bis.SlotVerdict, bool) {
	if row.specStr == "" {
		return bis.SlotVerdict{}, false
	}
	band, hasBand := s.loadBandFor(row.className, row.specStr, row.faction)
	if !hasBand {
		return bis.SlotVerdict{}, false
	}
	pick, named := band.BySlot()[item.Slot]
	if !named || pick.ItemID <= 0 || !namesItem(pick, item.ItemID) {
		return bis.SlotVerdict{}, false
	}
	wornID, hasWorn := row.gear[item.Slot]
	return bis.VerdictFor(pick, wornID, hasWorn), true
}

// fallbackMatch is lootCandidates' own tier-1 test: row's class can use item at all (read
// from the build's own per-class item table), and item.ItemLevel either outranks the
// worn item in item's own slot or that slot carries nothing - delta mirrors
// fallbackVerdict's own semantics, computed from real, build-resolved item levels.
func (s *Store) fallbackMatch(row RosterRow, item lootItem) (int, bool) {
	dropItem, ok := s.buildItemFor(row.className, item.ItemID)
	if !ok {
		return 0, false
	}
	wornID, hasWorn := row.gear[item.Slot]
	if !hasWorn {
		return fallbackVerdict(dropItem.ItemLevel, nil)
	}
	wornItem, ok := s.buildItemFor(row.className, wornID)
	if !ok {
		// The worn item's own level cannot be resolved (an id this class's item table
		// does not carry) - conservatively not a fallback candidate rather than guessing.
		return 0, false
	}
	level := wornItem.ItemLevel
	return fallbackVerdict(dropItem.ItemLevel, &level)
}

// fallbackVerdict is pure: given the drop's own item level and the worn item's level (nil
// for an empty slot), decides the fallback tier's own item-level-upgrade test. Class
// eligibility is the caller's own, separate check (buildItemFor's ok return already
// answers it) - this function only ever answers the item-level half.
func fallbackVerdict(dropItemLevel int, wornItemLevel *int) (delta int, qualifies bool) {
	if wornItemLevel == nil {
		return dropItemLevel, true
	}
	d := dropItemLevel - *wornItemLevel
	if d <= 0 {
		return 0, false
	}
	return d, true
}

// buildItemFor resolves itemID through classSlug's own per-class item table
// (data/builds/<build>/items/<classSlug>.json via api/internal/trees) - presence there is
// itself the class-usability test (armor type and class restrictions the data pipeline
// already applied when it built that per-class file), not a second check this package
// would otherwise have to reimplement.
func (s *Store) buildItemFor(classSlug string, itemID int) (trees.Item, bool) {
	return s.ItemIn(s.BisBuild(), classSlug, itemID)
}

// ItemIn is itemID in version's per-class item table, the one the loot tab and the BiS page
// read (name, quality, item level, icon); false for an unknown build, class or item.
func (s *Store) ItemIn(version, classSlug string, itemID int) (trees.Item, bool) {
	if s.Trees == nil {
		return trees.Item{}, false
	}
	build, ok := s.Trees.Build(version)
	if !ok {
		return trees.Item{}, false
	}
	classID, ok := spec.New(build).ClassID(classSlug)
	if !ok {
		return trees.Item{}, false
	}
	return build.Item(classID, itemID)
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

// lootCandidateLess reports whether a ranks ahead of b: tier 0 (BiS-matched) always ahead
// of tier 1 (fallback); within tier 0, gain_dps desc (a nil gain sorts last within the
// tier, never ahead of a measured one); within tier 1, ilvl_delta desc; either tier,
// attendance (present over the candidate's own nights, then present) desc breaks the remaining tie.
func lootCandidateLess(a, b LootCandidate) bool {
	if a.tier != b.tier {
		return a.tier < b.tier
	}
	if a.tier == 0 {
		ag, bg := gainOrLowest(a.GainDps), gainOrLowest(b.GainDps)
		if ag != bg {
			return ag > bg
		}
	} else {
		ad, bd := deltaOrZero(a.IlvlDelta), deltaOrZero(b.IlvlDelta)
		if ad != bd {
			return ad > bd
		}
	}
	if ar, br := attendanceRatio(a.Attendance), attendanceRatio(b.Attendance); ar != br {
		return ar > br
	}
	return a.Attendance.Present > b.Attendance.Present
}

func gainOrLowest(g *float64) float64 {
	if g == nil {
		return math.Inf(-1)
	}
	return *g
}

func deltaOrZero(d *int) int {
	if d == nil {
		return 0
	}
	return *d
}

// lootAwards reads every current award for guildID's encounterID, keyed by item id, with
// the claimant's and the awarding officer's own names resolved.
func (s *Store) lootAwards(ctx context.Context, guildID, encounterID int64) (map[int]*AwardedTo, error) {
	rows, err := s.Pool.Query(ctx, `
		select la.item_id, la.id, la.character_key, coalesce(ae.name, ''), la.awarded_at, la.awarded_by
		from loot_awards la left join addon_exports ae on ae.character_key = la.character_key
		where la.guild_id = $1 and la.encounter_id = $2
		order by la.awarded_at, la.id`, guildID, encounterID)
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
		if err := rows.Scan(&itemID, &a.AwardID, &a.CharacterKey, &a.Name, &awardedAt, &awardedBy); err != nil {
			return nil, fmt.Errorf("guilds: loot awards: %w", err)
		}
		a.At = awardedAt.UTC().Format(time.RFC3339)
		if awardedBy != nil && s.Accounts != nil {
			a.ByName = s.awarderName(ctx, guildID, *awardedBy)
		}
		out[itemID] = &a
	}
	return out, rows.Err()
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
	for _, it := range onyxiaLoot {
		if it.ItemID == itemID {
			return true
		}
	}
	return false
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

// awardCounts is how many of the tier's drops each character has been awarded in this guild. An
// award row names one encounter (Onyxia, the only one an award can name), so every row is the
// current tier and the count needs no table of its own.
func awardCounts(awards map[int]*AwardedTo) map[string]int {
	out := make(map[string]int, len(awards))
	for _, a := range awards {
		out[a.CharacterKey]++
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
