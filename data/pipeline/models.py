from dataclasses import dataclass
from typing import Any, Literal

from pydantic import (
    BaseModel,
    Field,
    SerializerFunctionWrapHandler,
    field_validator,
    model_serializer,
)


class Zone(BaseModel):
    id: int
    name: str
    map_id: int
    map_name: str
    parent_id: int | None
    level_min: int | None
    level_max: int | None


class Dungeon(BaseModel):
    id: int
    name: str
    map_id: int
    description: str
    order: int


class Item(BaseModel):
    id: int
    name: str
    quality: int
    item_level: int
    required_level: int
    class_id: int
    subclass_id: int
    inventory_type: int
    #: The random suffixes this item rolls, from the engine fork's
    #: `randomSuffixOptions` (parity contract 6.3). Empty for an item that
    #: rolls none, and for every row until `python -m pipeline loot` has run
    #: for the build -- `normalize` has no fork database to read.
    suffixes: list[int] = []
    #: "alliance_only", "horde_only", or "" for no restriction, from the
    #: fork's `factionRestriction` (contract 10.3's `SimItem` field). The
    #: client states none of this on build 1.60.1.69893, so like `suffixes`
    #: it is empty until `loot` has run.
    faction_restriction: str = ""
    #: loot-parity-2 lane, 2026-09-30: the flat catalogue's own copy of
    #: `GearItem.required_level_source` (same doc, same vocabulary) -- unset
    #: (None) for a row `pipeline.normalize.items.normalize_items` built
    #: straight off the client's own `ItemSparse`, so every `items.json`
    #: consumer (not just the per-class `items/<class>.json` files) can
    #: tell a supplement row apart from a client one without a second
    #: lookup.
    required_level_source: (
        Literal["client", "wowhead", "item_level_proxy", "none", "classic-db"] | None
    ) = None
    #: The flat catalogue's own copy of `GearItem.stats_source` (same doc):
    #: `"wowhead"` or `"classic-db"` for a row `normalize_items` itself
    #: never produced -- the client's `ItemSparse`/`Item` name no such id at
    #: all -- `None` for every row that is the client's own.
    stats_source: Literal["wowhead", "classic-db"] | None = None
    #: The flat catalogue's own copy of `GearItem.client_unconfirmed` (same
    #: doc): True for a `stats_source == "classic-db"` row, real 1.12
    #: itemization the Forever client has not (yet) shipped as a hotfix.
    client_unconfirmed: bool = False
    #: data-followups-4 lane, 2026-09-30: the flat catalogue's own copy of
    #: `GearItem.superseded_by` (same doc) -- the Forever-new item id that
    #: replaces this one as an actual reward, for a legacy client row
    #: `pipeline.loot.supersede.superseded_pairs` found co-listed with its
    #: own re-itemised copy in the same loot source. `None` for every
    #: other row. The row stays in this file either way -- `items.json`'s
    #: no-shrink gate (`pipeline.csvio.check_item_sparse_completeness`)
    #: counts it as present, since a legacy id can still be equipped by
    #: whoever already has one and needs its tooltip.
    superseded_by: int | None = None


class Spell(BaseModel):
    id: int
    name: str


class Source(BaseModel):
    label: str
    url: str
    kind: str


class ForeverChange(BaseModel):
    text: str
    sources: list[Source]


class PlayableClass(BaseModel):
    id: int
    name: str
    slug: str
    color: str
    forever_changes: list[ForeverChange] = []


class PlayableRace(BaseModel):
    id: int
    name: str
    slug: str
    faction: str
    placeholder: bool = False
    forever_changes: list[ForeverChange] = []


class Combo(BaseModel):
    race_id: int
    class_id: int
    new_in_forever: bool


class TalentNode(BaseModel):
    id: int
    tab_id: int
    tab_name: str
    class_id: int
    tier: int
    column: int
    spell_ids: list[int]
    prereq_talent_id: int | None


class TalentRank(BaseModel):
    spell_id: int
    description: str


class TalentEntry(BaseModel):
    id: int
    name: str
    icon: str
    max_rank: int
    tier: int
    column: int
    prereq_talent_id: int | None
    prereq_rank: int | None
    ranks: list[TalentRank]
    #: The client's own spell for the talent, which is what every consumer
    #: keyed by spell id (the sim, the report's planner link) will see written
    #: by the game. Appended last, like `background` below: the emitted key
    #: order is the schema's only compatibility surface, so new fields go on
    #: the end and existing ones never move.
    spell_id: int
    #: True when a spell this talent names is in none of the build's client tables:
    #: a hotfix-added spell that only the Wowhead overlay knows (the id and text are
    #: Wowhead's). Appended last and left out of the emitted JSON while False, so a
    #: build with no such talent stays byte-identical to the pre-overlay schema.
    hotfix_only: bool = False

    @model_serializer(mode="wrap")
    def _omit_false_hotfix_only(self, handler: SerializerFunctionWrapHandler) -> dict[str, Any]:
        data = handler(self)
        if not data.get("hotfix_only"):
            data.pop("hotfix_only", None)
        return data


class TalentTree(BaseModel):
    id: int
    name: str
    position: int
    talents: list[TalentEntry]
    #: The tab's `BackgroundFile`, lowercased. The site fetches the processed
    #: image at `/data/<build>/trees/<background>.webp`.
    background: str
    #: The tab's own icon (tenet 3's spec-icon counterpart, day3
    #: data-followups-11 lane): the client's `TalentTab.SpellIconID`,
    #: resolved through `ManifestInterfaceData` the same way a talent's
    #: or an item's `icon` already is (`pipeline.icons.resolve_icon`).
    #: Appended last, like `TalentEntry.spell_id`: the emitted key order
    #: is this schema's only compatibility surface.
    icon: str


class ClassTalents(BaseModel):
    build: str
    class_id: int
    class_slug: str
    trees: list[TalentTree]


class ItemReputation(BaseModel):
    faction_id: int
    standing: Literal["neutral", "friendly", "honored", "revered", "exalted"]
    #: The client's Faction name and the side that earns it
    #: (`pipeline.normalize.factions`); None without a Faction row, and side
    #: None for a reputation both sides earn. loot sources an otherwise
    #: unsourced reward from these (`pipeline.loot.client_rep`).
    faction_name: str | None = None
    side: Literal["alliance", "horde"] | None = None


class GearItem(BaseModel):
    id: int
    name: str
    icon: str
    slot: str
    quality: int
    required_level: int
    #: Where `required_level` came from, in the precedence
    #: `pipeline.normalize.gear.resolve_required_level` applies (normalize-levels
    #: lane, 2026-09-29): `"client"` for a shipped or hotfix-merged row that states
    #: a non-zero `RequiredLevel` itself; `"wowhead"` when the client states 0 (or
    #: has no row at all -- a wowhead-supplement item) and wowhead's Forever
    #: gear-planner payload (`pipeline.wowhead_items`) names a non-zero level for
    #: the same id; `"item_level_proxy"` for real gear (item level > 1) neither of
    #: those resolved -- `item_level - 5`, floored at 0 and capped at 60, the same
    #: formula `sim/leveling.ItemLevelProxyRequiredLevel` computes independently in
    #: Go (the two cannot share code across languages, only the number); `"none"`
    #: for an item_level-1 row (or anything else with no proxy to compute), where
    #: required_level 0 is simply the right answer, not a gap.
    #: `"classic-db"` added catalogue-universe lane (2026-09-30): the whole
    #: row -- not just `required_level` -- has no client counterpart at all,
    #: same all-or-nothing meaning `stats_source`'s own `"classic-db"` value
    #: carries; see that field's doc.
    required_level_source: Literal["client", "wowhead", "item_level_proxy", "none", "classic-db"]
    item_level: int
    armor: int
    stats: dict[str, int]
    #: Weapon damage per swing and the swing itself, from ItemSparse's
    #: ItemDamageMin_0/ItemDamageMax_0/ItemDelay. Zero for anything that is
    #: not a weapon, and zero on a build whose client computes damage from
    #: curves this pipeline does not resolve.
    damage_min: int = 0
    damage_max: int = 0
    #: Seconds, ItemDelay / 1000.
    speed: float = 0.0
    #: Derived: the mid damage over the speed, 0.0 when either is unknown.
    dps: float = 0.0
    #: True for a two-handed main-hand weapon -- one that leaves no off-hand
    #: free. Validation rule 6 refuses an off-hand item beside one. A bow, gun,
    #: crossbow, thrown weapon or wand is False: it takes the ranged slot, not
    #: the main hand (see `normalize.gear.TWO_HAND_INVENTORY_TYPES`).
    two_hand: bool = False
    #: The use or proc description, token-substituted the way talent text
    #: is. Empty for an item whose value is entirely in its stats.
    effect_text: str = ""
    #: `"wowhead"` when every field above but `id`/`name`/`slot` came from
    #: `pipeline.wowhead_items.to_gear_item` -- the client's own ItemSparse/
    #: Item carries no row for this id at all (raid gear the beta character
    #: has never seen drop, so it never reached the shipped or hotfix
    #: export), and wowhead's Forever gear-planner payload is the only
    #: source for its armor/stats/damage. `None` (the default) for a row the
    #: client itself states, same as `required_level_source`'s `"client"` --
    #: the two fields differ because a client row can still have its
    #: `required_level` alone resolved from wowhead (source `"wowhead"`)
    #: while every other field is the client's own; `stats_source` is only
    #: ever set when the *entire* row -- not just one field -- has no client
    #: counterpart (catalogue-completeness lane, 2026-09-29). The planner and
    #: the simulator's Top Gear both need this to know a number here is
    #: wowhead's scrape, not the client's own table, the same way the site
    #: already shows `required_level_source`.
    #: `"classic-db"` added catalogue-universe lane (2026-09-30): the client's
    #: `ItemSparse`/hotfix cache has never seen this id at all (the beta
    #: character never had it drop), and cmangos/classic-db's own 1.12
    #: `item_template` is the only source -- armor, stats, weapon damage and
    #: `effect_text` all come from that table (`pipeline.classicdb_items`).
    #: Distinct from `"wowhead"`: wowhead's scrape is corroborated against
    #: this same client's curve tables on every id both sources share
    #: (`pipeline.wowhead_items`'s own doc); classic-db is 1.12 data with no
    #: such corroboration, which is exactly what `client_unconfirmed` below
    #: flags.
    stats_source: Literal["wowhead", "classic-db"] | None = None
    set_id: int | None
    unique: bool
    #: True for a `stats_source == "classic-db"` row the client's own `Item`
    #: and item-effect tables do not name in classic-db's shape
    #: (`pipeline.client_evidence`; `data/README.md`, "classic-db-only items
    #: and client_unconfirmed"). A cleared row keeps `stats_source`.
    #: A future hotfix refresh that adds the id to `raw/hotfixes/ItemSparse.csv`
    #: promotes the row to client stats (`stats_source` reverts to `None`) the
    #: next time `normalize` runs, the same way any other hotfix promotion
    #: works. False (the default) for every other row, client or wowhead alike.
    client_unconfirmed: bool = False
    #: ratings lane, 2026-10-07: the part of each rating-family `stats` entry
    #: (hit, crit, dodge, parry, block, defense) that an on-equip spell states
    #: as a LITERAL PERCENT (Fury Visor's 1 hit / 1 crit). Every other
    #: rating-family amount in `stats` is a combat-rating point count (Lionheart
    #: Helm's 20 hit / 28 crit) that `gametables/combatratings.txt` divides down
    #: to a percent (`pipeline/simdb/ratings.py`). A consumer converting `stats`
    #: to percent divides `stats[key] - percent_stats[key]` by the factor and
    #: adds `percent_stats[key]` back unchanged. A wowhead row records its
    #: `critstrkpct`/`hitpct`-style keys here (`wowhead_items.PERCENT_STAT_KEYS`).
    #: Empty for every item with no percent-stated rating-family stat.
    percent_stats: dict[str, int] = {}
    #: classicdb-fidelity lane, 2026-09-30: where `icon` actually came from --
    #: `"client"` (the client's own `IconFileDataID`, resolved through
    #: `ManifestInterfaceData`), `"fork"` (the engine fork's own
    #: `assets/database/db.json`) or `"wowhead"` (wowhead's Forever
    #: gear-planner payload) -- the same three-way fallback
    #: `pipeline.icons.resolve_icon_name` already applies, now recorded per
    #: row instead of only counted in the build log. `None` for a row built
    #: before this field existed (never emitted by current code, kept
    #: optional so an old committed build still validates).
    icon_source: Literal["client", "fork", "wowhead"] | None = None
    #: classicdb-fidelity lane, 2026-09-30: the planner's weapon-type
    #: vocabulary for a weapon row (`pipeline.normalize.gear.
    #: WEAPON_TYPE_BY_SUBCLASS`) -- one of the five ranged types (`wand`,
    #: `bow`, `gun`, `crossbow`, `thrown`) for a `slot: "ranged"` row, or a
    #: melee type (`axe`, `mace`, `polearm`, `sword`, `staff`, `fist`,
    #: `dagger`) for a `main_hand`/`off_hand` row; `two_hand` above already
    #: carries the one/two-handed distinction, so `axe`/`mace`/`sword` cover
    #: both. `None` for a non-weapon row, or a weapon row whose SubclassID
    #: this table does not (yet) map -- `sim/cmd/leveling-bis`'s
    #: `weapon_requirements.go` (`bis-ranker-integrity-6` lane) gates a
    #: caster's/hunter's ranged slot by this field, treating `None` the same
    #: as "unknown, exclude it" rather than guessing a type.
    weapon_type: str | None = None
    #: data-followups-4 lane, 2026-09-30 (`pipeline.loot.supersede`'s own
    #: doc): the Forever-new item id that actually replaces this row as a
    #: reward, for a legacy client row this build's own loot data still
    #: lists side by side with its re-itemised copy in the same source
    #: (the ninth wow-player sweep's Grand Marshal's Stave/High Warlord's
    #: War Staff finding). `None` for every other row, including the
    #: replacement itself. `sim/cmd/leveling-bis` (Go) is meant to treat a
    #: row with this set as never a pick candidate -- not yet wired there,
    #: see the lane report.
    superseded_by: int | None = None
    #: Spec C6 (api/openapi.yaml `Item.bonding`, docs/contracts/app-v1.md),
    #: published by the paperdoll lane 2026-10-31: the client's
    #: `ItemSparse.Bonding` as a word (`pipeline.normalize.gear.bonding_for`),
    #: for the tooltip's bind line. Set only on a client row of a build whose
    #: ItemSparse has the column; `None` for no binding, and for a classic-db
    #: or wowhead row, whose sources state none -- never guessed.
    bonding: Literal["pickup", "equip", "use", "quest"] | None = None
    #: Spec C6 (`Item.armor_type`): the armour word from Item.ClassID 4's
    #: SubclassID (`pipeline.normalize.gear.armor_type_for`); `None` for a
    #: cloak, a neck, a ring, a trinket, a weapon or a relic.
    armor_type: Literal["cloth", "leather", "mail", "plate", "shield"] | None = None
    #: Sim-accuracy round 9 (ranker lane): the client's reputation requirement,
    #: ItemSparse.MinFactionID and MinReputation
    #: (`pipeline.normalize.gear.reputation_for`). `None` for none, and for a
    #: classic-db or wowhead row. The BiS ranker gates every source of the item
    #: on it (sim/cmd/leveling-bis/reputation.go).
    reputation: ItemReputation | None = None
    #: Round 9 (ranker lane): the client's ItemSparse.LimitCategory, the
    #: "Unique-Equipped: <family> (1)" line. `None` for none. The BiS ranker never
    #: pairs two items of one category (sim/cmd/leveling-bis/pair.go).
    limit_category: int | None = None
    #: The client's own item shape, Item.ClassID, SubclassID and
    #: InventoryType (items.json's same three columns), so a page can apply
    #: an enchant's SpellEquippedItems rule (`EnchantRecord.equipped_*`) to
    #: the item: a held-in-off-hand item (class 4, inventory type 23) takes
    #: no weapon enchant. `None` on a build written before these existed.
    class_id: int | None = None
    subclass_id: int | None = None
    inventory_type: int | None = None


class ClassItems(BaseModel):
    build: str
    class_slug: str
    items: list[GearItem]


class ItemSetBonus(BaseModel):
    pieces: int
    description: str


class ItemSetRecord(BaseModel):
    id: int
    name: str
    item_ids: list[int]
    bonuses: list[ItemSetBonus]


class SuffixRecord(BaseModel):
    id: int
    name: str
    stats: dict[str, float]


class EnchantRecord(BaseModel):
    #: The effect id. NOT unique: the fork's own table repeats an effect
    #: across the slots it can go in (19 of its 150 ids), which is why
    #: `spell_id` and `item_id` are here too.
    id: int
    name: str
    #: The client's own text for the enchant (`SpellItemEnchantment.Name_lang`,
    #: "+7 Strength"); `name` is the recipe's. Empty when the client has no row.
    display: str = ""
    icon: str
    #: The planner slot names this enchant can be applied in.
    slots: list[str]
    #: The shape of item it needs: normal, two_hand, shield, kit or staff.
    item_types: list[str]
    #: Class slugs allowed to use it. Empty means no restriction.
    classes: list[str]
    stats: dict[str, float]
    phase: int
    #: Appended after the contract's keys, for telling two rows with the
    #: same effect id apart. 0 where the fork states neither.
    spell_id: int
    item_id: int
    #: The client's own rule for what the enchant can go on: its spell's
    #: SpellEquippedItems row. An item takes it when its class is
    #: `equipped_item_class` (-1: any), its subclass bit is set in
    #: `equipped_item_subclass_mask` (0: any) and its inventory type bit in
    #: `equipped_inventory_type_mask` (0: any). A held-in-off-hand item
    #: (armor class 4) takes no weapon enchant (class 2).
    equipped_item_class: int = -1
    equipped_item_subclass_mask: int = 0
    equipped_inventory_type_mask: int = 0


class LootBoss(BaseModel):
    id: str
    #: Empty where the fork database names no NPC for the id. 35 of the 67
    #: raid bosses and 39 of the 230 dungeon bosses on the pinned fork are
    #: in that state; an invented name would be worse than a blank one.
    name: str
    npc_id: int
    items: list[int]
    #: item id (string key, JSON-object convention) -> percent drop chance
    #: (0-100), from cmangos/classic-db's own `ChanceOrQuestChance`
    #: (src-classicdb lane, 2026-09-29). Only classic-db states a chance at
    #: all -- the fork database and wowhead's scrape carry none -- so this
    #: is unset for every boss entry classic-db did not itself contribute
    #: an item to.
    item_chances: dict[str, float] | None = None
    #: item id (string key) -> the CLASSIC item id it was copied from,
    #: for a Forever-new item re-itemisation added to this boss
    #: (`pipeline.loot.reitemise`'s own doc). Unset for every item this
    #: boss's own real source named.
    reitemised_from: dict[str, int] | None = None
    #: item id (string key) -> `"fork"` or `"wowhead"`, drop-sources-2
    #: lane 2026-09-29: which origin attributed this boss item when
    #: cmangos/classic-db's own `creature_loot_template` (direct or
    #: reference) names NO drop for this npc at all -- tenet 8's own
    #: labelling requirement ("nothing we cannot verify is shown as
    #: fact... labelled as unverified") for a fact this pipeline cannot
    #: independently confirm. `pipeline.audit.check_drops` reads this to
    #: set the finding's own severity (a `"wowhead"`-origin attribution,
    #: already labelled unverified here, is `minor`; an unlabelled
    #: `"fork"`-origin attribution in a launch, non-raid-gated instance
    #: is the `major` case worth a person's judgment). Unset for every
    #: item classic-db DOES corroborate (verified, not unconfirmable) and
    #: for every item classic-db itself attributed (its own record IS the
    #: corroboration).
    item_source_origin: dict[str, str] | None = None
    #: The creature's own cmangos/classic-db `creature_template.Rank` (task
    #: L1.7, `pipeline.loot.npc_rank`): 1 elite, 2 rare elite, 3 boss, 4 rare
    #: (`pipeline.classic_sources.RARE_NPC_RANKS`). Unset for a normal
    #: creature (rank 0) and for an npc the dump does not rank.
    rank: int | None = None
    #: True for a rare SPAWN (task L1.7, `pipeline.classic_sources.
    #: parse_rare_spawns`): a rare or rare-elite creature with a pooled
    #: spawn, so it is not there on every run. Unset otherwise -- including
    #: a fixed boss the dump ranks rare elite (Edwin VanCleef).
    rare_spawn: bool | None = None


class VendorCost(BaseModel):
    """A vendor price in something other than gold."""

    amount: int
    currency: str


class LootSource(BaseModel):
    """One place loot comes from (parity contract 6.1).

    Every key after `name` is optional and omitted when it does not apply
    to the kind, which is what `write_document`'s `exclude_none` is for: a
    crafted source carries `profession` and `items`, a raid carries
    `zone_id`, `bosses` and `trash`.
    """

    id: str
    kind: str
    name: str
    zone_id: int | None = None
    #: A phase name from api/internal/phase. None means open from launch.
    #: Never set by the generator -- only by a curated overlay, because the
    #: databases state no dates.
    opens: str | None = None
    profession: str | None = None
    faction_id: int | None = None
    standing: str | None = None
    rank: int | None = None
    #: True for a `world` source whose creature is a rare spawn (task L1.7,
    #: `pipeline.loot.npc_rank`): Swiftmane, Lupos. Unset otherwise.
    rare_spawn: bool | None = None
    #: The vendor kind's own npc, so the page can link it the way a boss's
    #: `npc_id` already does.
    npc_id: int | None = None
    #: "alliance" or "horde", for a `pvp` kind source only --
    #: `pipeline.loot.pvp_faction`'s own doc: an honor-rank reward's
    #: numeric `pvp:rank-N` bucket used to mix both factions' titled
    #: rewards together, which is how an Alliance-titled item
    #: ("Knight-Lieutenant's Pauldrons") reached a Horde character's
    #: list. `None` for every other kind.
    faction: Literal["alliance", "horde"] | None = None
    #: "classic-db" when `faction` came from the rank quartermaster's own
    #: classic-db `npc_vendor` row (`pipeline.loot.pvp_faction.
    #: vendor_npc_factions`); "title" when no vendor resolved it and the
    #: item's OWN display name's rank title settled it instead
    #: (`pipeline.loot.pvp_faction.title_faction`) -- the site and
    #: reports should say the vendor did not corroborate this one.
    #: `None` for every other kind.
    faction_source: Literal["classic-db", "title"] | None = None
    #: The in-game rank title (`pipeline.loot.pvp_faction.rank_title`),
    #: for a `pvp` kind source only -- fourth wow-player sweep, item 2:
    #: the bare bucket name ("Rank 9 (Alliance)") never named the
    #: quartermaster's own title ("Master Sergeant"), which
    #: `sim/cmd/leveling-bis` (Go, band.go's pvpSourceLabel) and the
    #: web (`web/src/lib/bis/copy.ts`'s own mirrored table) both need to
    #: publish "PvP rank 9 · Master Sergeant · Alliance" instead. `None`
    #: for every other kind, and for a pvp source whose rank falls
    #: outside the known 14-rank ladder (rank_title's own doc).
    title: str | None = None
    #: The Era honor rank an item's title implies ("Knight-Lieutenant's" ->
    #: 11), for a `pvp` source whose items carry no rank lock in Forever:
    #: neither the client (ItemSparse.RequiredPVPRank) nor classic-db ranks
    #: them, and Forever's server text shows only a level (round 8, owner
    #: ruling). Such a source has `rank` None and reads "PvP reward". This
    #: field is for eligibility only (the ranker's rank cap), not a lock.
    #: Exactly one of `rank` and `title_rank` is set on a pvp source.
    title_rank: int | None = None
    bosses: list[LootBoss] | None = None
    trash: list[int] | None = None
    items: list[int] | None = None
    #: item id (string key) -> percent drop chance (0-100), for a `world`
    #: or `zone` kind source's own flat `items`/`trash` list -- see
    #: `LootBoss.item_chances`' own doc for where this comes from and why
    #: it is classic-db-only.
    item_chances: dict[str, float] | None = None
    #: item id (string key) -> the CLASSIC item id it was copied from, for
    #: a Forever-new item re-itemisation added to this source's own flat
    #: `items`/`trash` list -- see `pipeline.loot.reitemise`'s own doc.
    reitemised_from: dict[str, int] | None = None
    #: `world_drop` kind only (src-classicdb world-drop-pool lane,
    #: 2026-09-29): the level range cmangos/classic-db's own dump states
    #: for this generic, bind-on-equip, auction-housable drop pool
    #: (`pipeline.classic_sources._world_drop_records`' own doc). Either
    #: side is `None` when the dump names no level for the pool at all --
    #: never invented.
    #: `dungeon` and `raid` kinds (2026-11-05): the instance's level range,
    #: the client's own Group Finder row (`pipeline.loot.instance_levels`).
    level_min: int | None = None
    level_max: int | None = None
    #: `crafted` kind only (2026-11-05): item id (string key) -> the
    #: profession skill its recipe needs (`pipeline.loot.crafted_skill`).
    #: An item with no skill in the pinned classic-db dump is absent.
    required_skill: dict[str, int] | None = None
    #: `vendor` and `rep` kinds (feedback 2026-11-09): where the seller
    #: stands ("Booty Bay", "The Barrens"), from curated/vendor-terms.json
    #: (`pipeline.loot.vendor_terms`). Absent when nothing curated says.
    place: str | None = None
    #: `vendor` kind only: item id (string key) -> its price when it is not
    #: gold ("100 Tarnished Undermine Reals"), from the same curated file.
    #: An item with no curated price is absent, never guessed.
    item_costs: dict[str, VendorCost] | None = None
    #: `vendor` kind only: the reputation a quartermaster sells for, when
    #: every item it sells needs one faction's standing in the client
    #: (Creeg Bothunk, Darkspear Raiders). With `faction_id` set beside it.
    reputation: str | None = None
    #: `vendor` kind only, beside `reputation`: item id (string key) -> the
    #: standing the client requires for it.
    item_standings: dict[str, str] | None = None
    #: "wowhead" or "classic-db" when this source exists ONLY because a
    #: scrape/dump named it for an item the engine fork's own database
    #: named no source for at all (night-item-sources lane, 2026-09-28;
    #: src-classicdb lane, 2026-09-29); omitted (None) for every source
    #: the fork database itself names, which is still most of them. A
    #: source more than one origin names (an item added to an existing
    #: `world:`/`vendor:`/`crafted:` bucket) keeps whichever origin found
    #: it FIRST in priority order fork > classic-db > wowhead --
    #: EXCEPT a `raid`/`dungeon` source classic-db's own dump fully
    #: corroborates (every item the fork itself named for it also has a
    #: classic-db `creature_loot_template` row), which is promoted to
    #: `"classic-db"` even though the fork named it first (loot-parity-2
    #: lane, 2026-09-30, `pipeline.loot.sources._classic_db_corroborated_
    #: raid_dungeon_origin`'s own doc: Onyxia's Lair is the measured
    #: case -- the fork's own 16-item AtlasLoot table for it is a subset
    #: of what classic-db's dump independently names for the same boss).
    source_origin: str | None = None
    #: A short, player-unreadable reason this source (or the `opens`
    #: gate this specific source copy carries -- `pipeline.loot.
    #: vendor_cost.apply_vendor_cost_gate`'s own `vendor:<npc_id>:later`
    #: sibling) is not simply an ordinary open-at-launch listing: never
    #: shown in player-facing copy (tenet 7), read by reports and
    #: reviewers only. `None` for every source this applies to (most of
    #: them).
    source_note: str | None = None


class QuestSource(BaseModel):
    """One quest that hands out an item, for `LootFile.quests`.

    The fork database states no faction on a quest directly; the item it
    awards carries whichever `factionRestriction` the quest's side amounts
    to in practice, so `faction` here is that restriction standing in for
    it, not a fact the quest row itself states -- UNLESS the quest id is
    covered by the pinned classic-db dump (quest-faction lane, 2026-09-29),
    in which case `faction` is the quest's own `quest_template.
    RequiredRaces` (`pipeline.classic_sources.
    quest_factions_from_classic_sources`, applied by `pipeline.loot.
    sources.build_loot` over every `QuestSource` regardless of which
    scrape -- fork, classic-db or wowhead -- produced it) and the item's
    restriction is not consulted at all: the quest's own faction can
    legitimately differ from a reward item that is itself unrestricted
    (item 270018 Hammerbone, quest 914 Leaders of the Fang, horde-only via
    `RequiredRaces` even though the item carries no `factionRestriction`).

    min_level/level (2026-09-28 quest-levels finding): the client's own
    item row for a quest reward almost always states `required_level` 0
    -- the QUEST's level gates the reward, not the item's. `min_level` is
    the quest's own minimum level (wowhead's `reqlevel`) and `level` is
    the level the quest is written for (wowhead's `level`), both from
    `pipeline.wowhead_quests` scraping
    `https://www.wowhead.com/forever/quest=<quest_id>`. `level_source` is
    `"wowhead"` when scraped, or `"item_level_proxy"` when wowhead has no
    page for the quest and both fields instead fall back to
    `wowhead_quests.item_level_proxy(item.item_level)`.

    faction_source: "classic-db" when `faction` came from the quest's own
    `RequiredRaces` (verified against a primary source); "item" when the
    quest id is absent from the pinned dump (a Forever-new quest, such as
    79980 Scramble or 92422 The Wrath of Rath'mael) and `faction` is still
    only the reward item's own restriction standing in for it, unverified
    against the quest itself -- the site and reports should say so rather
    than present it as fact. `None` only for a `QuestSource` built by a
    caller other than `build_loot` (tests, mostly) that never set it.
    """

    quest_id: int
    name: str
    faction: str
    #: data-followups-3 lane, 2026-09-30, item 2: `"wowhead"` joins
    #: `"classic-db"`/`"item"` -- wowhead's own quest-page `side` field
    #: (`pipeline.quest_levels.QuestLevelEntry.faction`), the second
    #: fallback behind classic-db's `RequiredRaces` and ahead of
    #: publishing `faction="unknown"` (`"item"`, now with no item-derived
    #: guess left to report -- see `pipeline.loot.sources.build_loot`'s
    #: own doc for why the guess was retired).
    faction_source: Literal["classic-db", "wowhead", "item"] | None = None
    min_level: int
    level: int
    level_source: str
    #: A phase name from `api/internal/phase`, same vocabulary and same
    #: meaning as `LootSource.opens` (that field's own doc) -- `None`
    #: means open from launch. Quest-gates lane, 2026-09-29:
    #: `pipeline.loot.sources.apply_quest_opens_gate` sets this (never
    #: the generator's `_keyed_sources`/`classicdb_additions`/
    #: `wowhead_additions`, which all run before every source is merged)
    #: whenever this quest's own turn-in item(s) -- classic-db's
    #: `SrcItemId`/`ReqItemId1-4`, resolved through the quest's
    #: `PrevQuestId` chain -- are themselves only obtainable from a
    #: source `opens` is already set on: a raid boss drop directly, or
    #: another such quest, recursively. `sim/cmd/leveling-bis/data.go`'s
    #: `itemSource.Opens` carries this through so the leveling ranker
    #: gates a quest-sourced legendary or raid-chain reward exactly the
    #: way it already gates a direct raid drop.
    opens: str | None = None
    #: `pipeline.classic_sources.ClassicDbQuestInfo.required_rep_faction`/
    #: `required_rep_standing` -- the reputation faction and standing this
    #: quest itself needs, when classic-db's own `quest_template` row
    #: states one directly. `pipeline.loot.sources.apply_quest_opens_gate`
    #: also fills these in as a FALLBACK, from the SAME item's own
    #: `rep`-kind `LootSource`, for a quest classic-db states no
    #: reputation requirement on directly (rep-gate lane, 2026-09-30:
    #: Earthstrike, item 21180, quest 8573 "Champion's Battlegear" -- the
    #: fork database already names that item from a
    #: `rep:cenarion-circle:exalted` source, so the quest reward inherits
    #: that gate rather than reading as available with no requirement at
    #: all). `None`/`None` for the overwhelming majority of quests, which
    #: need no reputation at all.
    required_rep_faction: int | None = None
    required_rep_standing: str | None = None
    #: `pipeline.classic_sources.ClassicDbQuestInfo.classes` -- the sorted
    #: class slugs allowed to accept this quest at all, `None` when the
    #: quest is open to any class (the overwhelming majority). Day3
    #: data-followups-7 lane, 2026-09-30: Fire Ruby (item 20036)'s own
    #: quest 8253 "Destroy Morphaz" is mage-only, so `pipeline.loot.
    #: sources.build_loot` applies this over every `QuestSource` it
    #: assembles (same "quest's own primary-source fact wins, regardless
    #: of which scrape produced the link" rule `faction` above already
    #: follows) and `pipeline.audit.check_bis.check_bis` blocks a pick or
    #: alternative whose only sources all exclude the spec's own class.
    classes: list[str] | None = None
    #: `pipeline.classic_sources.ClassicDbQuestInfo.profession`/`.skill` --
    #: a profession (or secondary skill) this quest itself requires, when
    #: classic-db's own `RequiredSkill` states one this pipeline
    #: recognises. Not a per-class filter (a profession is learnable by
    #: any class): a plain fact for the page to state ("Requires
    #: Blacksmithing 250"), same "never invented, absent rather than
    #: guessed" policy as `required_rep_faction` above.
    profession: str | None = None
    skill: int | None = None
    #: 2026-11-05: the zone the quest log files this quest under, from
    #: classic-db's `quest_template.ZoneOrSort` (`pipeline.classic_sources.
    #: ClassicDbQuestInfo.zone_id`), named from the build's zones.json.
    #: `None` for a quest the dump does not cover (Forever-new), one filed
    #: under a category rather than a zone (class, profession, holiday
    #: quests), or a zone id zones.json does not name.
    zone_id: int | None = None
    zone: str | None = None


class LootFile(BaseModel):
    sources: list[LootSource]
    #: Per-item quest detail (parity contract 6.1's quest kind grew a
    #: face): item id -> the quest(s) that award it. The `quest` source
    #: above stays as the flat compatibility bucket; this is additive.
    quests: dict[str, list[QuestSource]] = {}
    #: item id -> "alliance" or "horde", for every faction-restricted item
    #: this build has -- quest items and non-quest items alike (a
    #: restricted vendor or drop item has no quest to carry the fact on).
    factions: dict[str, str] = {}


class LootSourcePatch(BaseModel):
    """Contract 10.4's `replace` entry: a loot source with every key but
    `id` optional, so gating a raid is one line instead of a restatement
    of its eleven bosses. Only the keys the file actually writes are
    applied -- `model_dump(exclude_unset=True)` is what tells "absent"
    from "explicitly null".

    `kind` is accepted but may not change: a source's kind decides the
    shape of its id, which candidates carry as `drop:<source id>`, so
    changing it would orphan every reference rather than edit one.
    """

    id: str
    kind: str | None = None
    name: str | None = None
    zone_id: int | None = None
    opens: str | None = None
    profession: str | None = None
    faction_id: int | None = None
    standing: str | None = None
    rank: int | None = None
    npc_id: int | None = None
    bosses: list[LootBoss] | None = None
    trash: list[int] | None = None
    items: list[int] | None = None

    model_config = {"extra": "forbid"}


class LootOverlay(BaseModel):
    """One `data/curated/loot/*.json`, in contract 10.4's shape.

    `sources` and `notes` are the curated provenance every hand-maintained
    fact here carries; `add`, `replace` and `remove` are the three
    operations, as flat arrays.
    """

    sources: list[Source] = []
    notes: str = ""
    add: list[LootSource] = []
    replace: list[LootSourcePatch] = []
    remove: list[str] = []

    model_config = {"extra": "forbid"}


class SpecRecord(BaseModel):
    spec: str
    class_slug: str
    spec_slug: str
    name: str
    role: str
    tree_index: int
    #: The stat `StatWeights` normalises to 1.0 for this spec (parity
    #: contract 1.4). Appended last, like `TalentEntry.spell_id`: the
    #: emitted key order is the generated files' only compatibility
    #: surface, so new fields go on the end and existing ones never move.
    reference_stat: str
    #: The stats that actually move this spec's damage or healing - the
    #: closed list `/sim/weights` offers, so a warrior is never asked
    #: about spirit and a healer never sees unexplained zeros for
    #: expertise. Follows reference_stat's own path: read here, checked
    #: against the engine's vocabulary, emitted to sim/specs/specs.go's
    #: generated Spec.WeightStats. specs.py's load_specs enforces the
    #: two content rules a physical/caster split can check without a
    #: per-spec table: reference_stat is always a member, and every id
    #: is a known stat. The full physical-vs-caster forbidden-pair rule
    #: is pinned by a Go test over specs.All (sim/specs/specs_test.go),
    #: which is what actually documents which stats belong to which
    #: kind of spec.
    weight_stats: list[str]
    #: The spell whose own icon stands for this spec, for the specs the
    #: client does not give a tree of their own (Feral is Cat Form, Feral
    #: Bear is Dire Bear Form). None for every other spec, whose icon is
    #: its tree's `TalentTab` icon. The icon itself is never curated:
    #: `pipeline.specs.spec_icons` reads it off the active build.
    form_spell: int | None = None
    #: The tree's 31-point ability as a level 60 rotation casts it, the id the
    #: engine reports. None where the tree's last tier has no pressed ability
    #: this spec's rotation casts.
    capstone_spell: int | None = None
    #: The spec as a player names it in full, where the catalogue name then the
    #: class reads wrong ("Feral Druid (Bear)", not "Feral Bear Druid"). None for
    #: every other spec: `pipeline.specs.full_name` composes theirs. The
    #: generated catalogue carries the resolved name for every spec, so no reader
    #: composes one of its own.
    full_name: str | None = None


class StatWeights(BaseModel):
    """One spec's curated stat weights, relative to its reference stat.

    Weights are opinions. Every entry carries at least one source and the
    site shows them beside the numbers; `pipeline/weights.py` refuses an
    entry without one, the way `pipeline/curated.py` refuses an unsourced
    Forever change.
    """

    spec: str
    weights: dict[str, float]
    sources: list[Source]


class AddonTalent(BaseModel):
    """One talent as the addon sees it: a cell, a ceiling and its trait node.

    tier and column are 1-based, matching `GetTalentInfo`'s own return
    values, so the addon compares what the client hands it without
    arithmetic in two places. node is the client's TraitNode id, the key
    `C_Traits.GetNodeInfo` takes: the 1.60 client keeps the Forever trees
    in the modern trait system and has no `GetTalentInfo` at all.
    """

    name: str
    tier: int
    column: int
    max_rank: int
    node: int


class AddonTab(BaseModel):
    name: str
    #: In the site's own array order. The export string encodes ranks in
    #: this order, so it is contract between Data.lua and Codec.lua.
    talents: list[AddonTalent]


class AddonClass(BaseModel):
    tabs: list[AddonTab]


class AddonRotationLine(BaseModel):
    """One priority-list entry the rotation card and the rotation toast
    show: which ability, at the rank this level band actually casts, and
    why (the curated APL's `panel_when` line for the step, else its own one-line
    `notes`, unedited)."""

    spell_id: int
    name: str
    condition: str
    #: The resolved rank's own icon (tenet 3: "an ability is never just a
    #: name"), the same convention as an item's or a talent's `icon` --
    #: a lowercase file name with no extension, resolved off the
    #: client's own `SpellMisc.SpellIconFileDataID` through
    #: `ManifestInterfaceData` (`pipeline.icons.resolve_icon`, the same
    #: join a talent's icon already goes through). Appended last, like
    #: `TalentEntry.spell_id`: the emitted key order is this schema's
    #: only compatibility surface, so new fields go on the end.
    icon: str
    #: Each race's own id (by the site's race slug) where a race's
    #: rank-choice row (round 13) casts a different rank from `spell_id`,
    #: which is what every other race casts; empty where every race casts
    #: `spell_id`. Appended last, like `icon`.
    race_spell_ids: dict[str, int] = {}


class AddonRotationBand(BaseModel):
    """One level band's rotation: `lines` in priority order, each already
    resolved to the highest rank learned by `level` (pipeline.addonrotation
    does the resolving; the addon never re-derives a rank from a level)."""

    level: int
    lines: list[AddonRotationLine]


class AddonBisAlternative(BaseModel):
    """One of a BiS pick's runners-up (leveling-bis's `runners_up`, the next
    two candidates by its own ranking): the item a player who cannot get
    the pick falls back on, in the same ids-only shape as the pick."""

    item_id: int
    #: One of pipeline.addonbis.SOURCE_KIND_CODES' keys, as on AddonBisItem.
    source_kind: str
    #: The short place, as on AddonBisItem.source; empty for none.
    source: str = ""


class AddonBisKindEntry(BaseModel):
    """The best item of one source kind for a slot (leveling-bis's `by_kind`, its
    chained `next` flattened, best rank first): what the slot shows in game when
    the player turned the pick's own kinds off (the addon's source filter, the
    twin of the site's; design/specs/2026-11-13-addon-source-filter.md)."""

    item_id: int
    #: One of pipeline.addonbis.SOURCE_KIND_CODES' keys: the entry's own
    #: `source_kind` when leveling-bis names one, else its by_kind key.
    source_kind: str
    #: The short place, as on AddonBisItem.source; empty for none. leveling-bis's
    #: dps_delta is not carried: the addon prices the gap from the two items' stats,
    #: as it prices every other gap, which keeps a fifth of the tail out of Data.lua.
    source: str = ""
    #: The standing a reputation or vendor item asks ("friendly", "honored",
    #: "revered", "exalted", "neutral"); empty for none.
    standing: str = ""
    #: A two-handed weapon (an off hand never sits beside it).
    two_hand: bool = False
    #: The ids the pair's first slot must not hold for this entry to be usable
    #: (leveling-bis's `not_with`, its own id included for a unique item); carried
    #: only on a pair's second slot (finger2, trinket2, off_hand), the one that reads it.
    not_with: list[int] = []


class AddonBisItem(BaseModel):
    """One slot's leveling BiS pick, ids only (pipeline.addonbis's own
    load-bearing decision: the item's name and swap_note are not carried
    again -- the client resolves the name for free from `item_id`, and the
    rest is prose a player reading a tooltip has no use for)."""

    item_id: int
    #: One of pipeline.addonbis.SOURCE_KIND_CODES' keys (leveling-bis's own
    #: sourceKindPriority list: quest, dungeon, crafted, rep, pvp, world,
    #: raid). Kept as the full word here -- render_lua is what encodes it
    #: down to one letter for size; the JSON stays readable.
    source_kind: str
    #: leveling-bis's own short place: "Wailing Caverns: Mutanus the
    #: Devourer" (instance: boss), a quest name, or a profession name for a
    #: crafted pick (report.go's own `source` string) -- the tooltip's
    #: muted source line (lane addon-tooltip-polish item 1: no reader-
    #: facing surface may show a bare "Source: Crafted" again). Empty for a
    #: pick leveling-bis wrote with no `source` at all, which the addon
    #: reads as "show no source line" rather than a placeholder.
    source: str = ""
    #: The pick's drop chance at `source` as a fraction (0.14 is 14 percent),
    #: from leveling-bis's own `chance` (the loot table's per-item drop rate,
    #: percent / 100). None when the pick is not a drop or the loot table
    #: states no rate for it, and for a bis file written before leveling-bis
    #: emitted the field.
    chance: float | None = None
    #: Up to two runners-up (leveling-bis's `runners_up`), best first. For a
    #: bis file written before leveling-bis emitted that field, the first two
    #: of the row's own `alternatives` list instead (pipeline.addonbis._pick).
    alternatives: list[AddonBisAlternative] = []
    #: The source filter's facts about the pick itself (leveling-bis's `kinds`,
    #: `standing`, `kind_standings`, `two_hand`, `not_with`): every filter kind the
    #: pick can be had from when there is more than its source_kind's own, the
    #: standing each gated kind asks, whether it is a two-hander, and the ids the
    #: pair's first slot must not hold (a pair's second slot only). Empty when the
    #: pick's source_kind says it all. Data.lua only (the packed tail): the filter
    #: facts would more than double addon-data.json, which nothing reads them from.
    kinds: list[str] = Field(default_factory=list, exclude=True)
    standings: dict[str, str] = Field(default_factory=dict, exclude=True)
    two_hand: bool = Field(default=False, exclude=True)
    not_with: list[int] = Field(default_factory=list, exclude=True)
    #: The best item of each other source kind, best rank first (`by_kind`).
    by_kind: list[AddonBisKindEntry] = Field(default_factory=list, exclude=True)


class AddonBisBand(BaseModel):
    """One level's leveling BiS, both factions inline (addonbis.py's own
    regrouping of leveling-bis's flat per-(band, faction) list -- see that
    module's docstring). `new_at_band[faction]` is the item ids newly best
    at this band, a subset of that faction's own `factions[faction]`
    values -- the tooltip's "(new at <band>)" tag reads it by id."""

    level: int
    #: faction ("alliance" | "horde") -> slot -> pick.
    factions: dict[str, dict[str, AddonBisItem]]
    #: faction -> item ids newly best at this band.
    new_at_band: dict[str, list[int]] = {}
    #: This band's own measured stat weights (leveling-bis's own `weights`
    #: list on its flat per-(band, faction) report), significant entries
    #: only, rounded for size (lane addon-tooltip-polish item 4). Weights
    #: do not vary by faction for a given spec/band in practice (race
    #: changes the sim's race, not the stat weights it measures), so
    #: pipeline.addonbis keeps the first faction's own list rather than
    #: one per faction. Empty when leveling-bis has not measured this
    #: band/spec yet -- the tooltip's verdict falls back to the curated
    #: static table (ns.Data.weights) in that case.
    weights: dict[str, float] = {}


class AddonDefaultBuild(BaseModel):
    """One default talent build the addon loads for a character with none
    saved (docs/contracts/addon-v2.md, contract 1): the guide's leveling or
    raid build as sim/cmd/addon-builds orders it -- one (tab, tier, column)
    triple per point in learning order, 1-based, the codec's own triple --
    and the guide's own short why for it, its frontmatter `buildWhy:` or
    `raidBuildWhy:` (sim/leveling.GuideBuildWhy), for the Builds pane."""

    name: str
    why: str
    order: list[tuple[int, int, int]]


class AddonDpsPerPointBand(BaseModel):
    """One leveling band's figures (task L1.5): sim/cmd/dps-per-point's
    derivatives on that band's own character -- the bare band's BiS set,
    the guide's leveling build, at the middle of the band's levels."""

    #: The spec's reference stat, repeated so a band table stands alone.
    reference: str
    #: Damage (healing for a healer) per second per reference-stat point.
    dps: float
    #: Damage per second one more point of the swung weapon's own damage
    #: per second is worth (the main hand; the ranged weapon for a
    #: ranged-attack-power spec); 0 for a spec that does not swing it.
    weapon_dps: float


class AddonDpsPerPoint(BaseModel):
    """What one point of a spec's reference stat is worth (docs/contracts/
    addon-v2.md, contract 1): sim/cmd/dps-per-point's raw stat-weights
    derivative on the guide's level-60 raid character. The addon's Progress
    multiplies a score delta, which is in reference-stat points, by `dps`.
    Task L1.5: `bands` carries the same figures per leveling band, which
    differ from the level-60 one by far more than 25 percent."""

    #: The spec's reference stat (data/curated/specs.json), the one the
    #: addon's weights are normalised to 1.0 against.
    reference: str
    #: Damage per second per point; healing per second when `healing`.
    dps: float
    #: True for a healer: the addon says "healing" instead of "damage".
    healing: bool
    #: The level-60 raid character's DPS per +1 weapon DPS (see
    #: AddonDpsPerPointBand.weapon_dps).
    weapon_dps: float = 0.0
    #: Band (20 ... 60) -> that band's figures; a band the sim could not
    #: measure beyond its own error is absent.
    bands: dict[int, AddonDpsPerPointBand] = {}


class AddonEncounterNeeds(BaseModel):
    """One raid boss's ready-check needs (contract 1's encounterNeeds), from
    data/curated/encounter_needs.json through pipeline.encounters. The ids
    are spell ids the addon checks the raid's auras against and shows by the
    client's own spell name; they are never printed."""

    buffs: list[int]
    debuffs: list[int]
    #: The school a raid stacks resistance against, or None.
    resist: str | None = None


class AddonData(BaseModel):
    build: str
    classes: dict[str, AddonClass]
    weights: dict[str, dict[str, float]]
    #: Keyed by spec ("warrior-fury"): a band at each pipeline.addonrotation.LEVEL_BANDS rung
    #: and at every level between where a line changes.
    rotations: dict[str, list[AddonRotationBand]] = {}
    #: Keyed by spec, one band per pipeline.addonbis.BIS_LEVEL_BANDS entry
    #: actually present in data/builds/<build>/bis/<spec>.json. Empty for
    #: every spec until the nightly workflow (.github/workflows/bis.yml)
    #: has written that directory -- pipeline.addonbis.build_bis tolerates
    #: a missing bis/ the same way build_rotations does not tolerate a
    #: missing curated/apl entry (the two directories have different
    #: maturity, hence the different rule).
    bis: dict[str, list[AddonBisBand]] = {}
    #: Keyed by spec, then "leveling" or "raid": data/builds/<build>/
    #: addon-builds.json, written by sim/cmd/addon-builds. Empty when that
    #: file does not exist.
    default_builds: dict[str, dict[str, AddonDefaultBuild]] = {}
    #: Keyed by spec: the client icon that stands for it, the one the site's
    #: planner draws (pipeline.specs.spec_icons). Empty for a build directory
    #: with no talent trees.
    spec_icons: dict[str, str] = {}
    #: Keyed by spec: data/builds/<build>/dps-per-point.json, written by
    #: sim/cmd/dps-per-point. A spec with no entry (every tank, or one the
    #: sim has not measured) is the addon's cue to fall back to "score" units.
    dps_per_point: dict[str, AddonDpsPerPoint] = {}
    #: Keyed by raid boss npc id (pipeline.encounters). Empty for a build
    #: with no loot.json raid record to check the ids against.
    encounter_needs: dict[int, AddonEncounterNeeds] = {}
    #: Rotation spell id -> the level the first rank of its ability is
    #: learned (pipeline.addonrotation.build_ability_levels), for the
    #: rotation region's "at level N". Every id any rotation line carries.
    ability_levels: dict[int, int] = {}
    #: Every item id that is a BiS pick or alternative in any band -> its
    #: items.json display name (pipeline.addonbis.build_bis_items, task
    #: L1.4), so the tooltip can name a pick before the client's item cache
    #: fills. Once per id; the packed pick strings are unchanged.
    bis_names: dict[int, str] = {}
    #: The weapons among bis_names -> their own dps at one decimal, from
    #: items/<class>.json's `dps` (what sim/cmd/leveling-bis ranks on).
    bis_dps: dict[int, float] = {}
    #: Every distinct source place a by_kind entry names, the most used first:
    #: Data.lua's `bisSources` list, which the packed by-kind tail indexes into
    #: (one short number per entry instead of the place's text, for the idle budget).
    #: Data.lua only, like AddonBisItem.by_kind.
    bis_sources: list[str] = Field(default_factory=list, exclude=True)
    #: The id of every BiS candidate pipeline.addonbis passed over because
    #: items.json marks it client_unconfirmed (one per skip). Never written
    #: to addon-data.json or Data.lua: it feeds the addon-data command's
    #: audit line, so one run parses the bis files once.
    bis_unconfirmed_skips: list[int] = Field(default_factory=list, exclude=True)


class PhaseBoundary(BaseModel):
    """One content phase and the instant it opens (parity contract 10.4)."""

    name: str
    start: str


class AplLevelNote(BaseModel):
    """A priority step's player note for a character of `from_level` and up.

    The saved sim result's Rotation card reads the one for the run's level
    (web/scripts/rotation-apl.mjs levelNotesOf), so a level 60 result never
    reads the step's leveling clauses. Top level, like engineering_notes,
    because the engine's APLListItem rejects any other key on a step.
    """

    model_config = {"extra": "forbid"}

    from_level: int
    step: int
    notes: str


class AplDocument(BaseModel):
    spec: str
    state: str
    rotation: dict
    sources: list[Source] = []
    notes: str = ""
    #: Spell ids this rotation names on purpose that the pinned engine build
    #: cannot act on yet -- a talent-granted ability with no spell file, or an
    #: aura nothing registers. The notes say so in prose; this is the same
    #: statement in a form a test can read, and the engine lane's rotation
    #: smoke test holds the engine's own unknown-action warnings to exactly
    #: this set. Empty is the normal case.
    inert: list[int] = []
    #: Engineering commentary for a `rotation.priorityList` step, keyed by
    #: that step's own 0-based index (a string key: JSON has no integer
    #: keys) -- bis-ranker-integrity-16 lane, 2026-09-30 (fourteenth
    #: player sweep, casters.md's headline finding): a step's own `notes`
    #: is rendered VERBATIM in the live rotation drawer
    #: (web/src/lib/sim/rotations.ts's rotationNotesFor), so Go file
    #: paths, "Fix round N" narrative and internal spell/aura ids never
    #: belong there -- this field is where that commentary moved instead.
    #: Deliberately NOT a key on the step object itself
    #: (`rotation.priorityList[i]`): that dict round-trips through the
    #: engine's own strict `APLListItem` protobuf schema
    #: (`pipeline.apl.parse_rotation`, test_apl.py's own
    #: test_every_rotation_parses_as_the_engines_own_aplrotation), which
    #: has exactly three fields ("hide", "notes", "action") and rejects
    #: anything else outright. Empty for every spec this lane's cleanup
    #: never touched, which is most of them.
    engineering_notes: dict[str, str] = {}
    #: Per-level versions of step notes (AplLevelNote).
    level_notes: list[AplLevelNote] = []
    #: The addon rotation panel's short condition line for a `rotation.priorityList` step,
    #: keyed by that step's 0-based index (a string key, as `engineering_notes`): "when the
    #: target is below 20%", "on cooldown". Top level for the same reason as `level_notes`:
    #: the engine's APLListItem rejects any other key on a step. A step with no line here
    #: falls back to its note's first sentence (pipeline.addonrotation). The step's `notes`
    #: stay the guides' prose; this is only the panel's line. A value is one line for every
    #: level, or level-scoped variants ([{"from_level": 1, "text": ...}, {"from_level": 18,
    #: "text": ...}]) where a clause names a spell the character learns later
    #: (pipeline.panelnames).
    panel_when: dict[str, str | list[dict]] = {}

    @field_validator("panel_when")
    @classmethod
    def _check_panel_when(cls, value: dict[str, str | list[dict]]) -> dict[str, str | list[dict]]:
        check_panel_when(value)
        return value


#: The longest panel condition line (`AplDocument.panel_when`) a curated file may carry.
PANEL_WHEN_MAX_CHARS = 60
#: Never in a panel line: the house style has no em dashes.
PANEL_WHEN_FORBIDDEN = ("\u2014",)


@dataclass(frozen=True)
class PanelWhen:
    """One panel condition line and the first level it prints at."""

    from_level: int
    text: str


#: The level cap: a variant's `from_level` is a level a character can be.
PANEL_WHEN_MAX_LEVEL = 60

#: The keys of one level-scoped `panel_when` variant.
PANEL_WHEN_VARIANT_KEYS = frozenset({"from_level", "text"})


def check_panel_when(lines: dict[str, str | list[dict]]) -> dict[int, tuple[PanelWhen, ...]]:
    """`panel_when` keyed by step index, each step's lines in ascending `from_level` (a plain
    string is one line from level 1), or ValueError naming the first bad entry: a key that is
    not a step index, a variant list that is empty, malformed or not strictly ascending, or a
    line that is empty, padded, over PANEL_WHEN_MAX_CHARS or holding an em dash."""
    checked: dict[int, tuple[PanelWhen, ...]] = {}
    for key, value in lines.items():
        if not key.isdigit():
            raise ValueError(f"panel_when key {key!r} is not a priorityList index")
        variants = _panel_when_variants(key, value)
        for variant in variants:
            _check_panel_line(key, variant.text)
        levels = [variant.from_level for variant in variants]
        if levels != sorted(set(levels)):
            raise ValueError(f"panel_when[{key}] from_level {levels} is not strictly ascending")
        checked[int(key)] = variants
    return checked


def _panel_when_variants(key: str, value: object) -> tuple[PanelWhen, ...]:
    """One step's `panel_when` value as its variants: a string is one line from level 1."""
    if isinstance(value, str):
        return (PanelWhen(1, value),)
    if not isinstance(value, list) or not value:
        raise ValueError(f"panel_when[{key}] is empty")
    variants = []
    for variant in value:
        if not isinstance(variant, dict) or set(variant) != PANEL_WHEN_VARIANT_KEYS:
            raise ValueError(
                f"panel_when[{key}] variant {variant!r} needs exactly from_level and text"
            )
        level = variant["from_level"]
        if (
            isinstance(level, bool)
            or not isinstance(level, int)
            or not 1 <= level <= PANEL_WHEN_MAX_LEVEL
        ):
            raise ValueError(
                f"panel_when[{key}] from_level {level!r} is not a level 1 to {PANEL_WHEN_MAX_LEVEL}"
            )
        variants.append(PanelWhen(level, variant["text"]))
    return tuple(variants)


def _check_panel_line(key: str, line: object) -> None:
    """ValueError when one panel line is empty, padded, too long or holds an em dash."""
    if not isinstance(line, str) or not line.strip():
        raise ValueError(f"panel_when[{key}] is empty")
    if line != line.strip():
        raise ValueError(f"panel_when[{key}] {line!r} has leading or trailing space")
    if len(line) > PANEL_WHEN_MAX_CHARS:
        raise ValueError(
            f"panel_when[{key}] {line!r} is {len(line)} characters; "
            f"the panel line is at most {PANEL_WHEN_MAX_CHARS}"
        )
    if any(mark in line for mark in PANEL_WHEN_FORBIDDEN):
        raise ValueError(f"panel_when[{key}] {line!r} has an em dash")


class SpellEffectConstant(BaseModel):
    index: int
    effect: int
    aura: int
    #: The amount the client shows, via pipeline.spelltext.effect_amount: the
    #: 1.60 client states it outright, Classic Era states it one short with the
    #: last point in EffectDieSides. One number, not two columns to recombine.
    amount: int
    sp_coefficient: float
    ap_coefficient: float
    period_ms: int
    misc_value: int
    trigger_spell: int
    #: SpellEffect.Variance: the roll is amount x (1 - variance/2) through
    #: amount x (1 + variance/2). Appended last (key order is the generated
    #: files' compatibility surface); 0 when the effect does not roll.
    variance: float = 0.0
    #: SpellEffect.EffectRealPointsPerLevel: added to the amount per caster
    #: level above the spell's own level, up to the spell's `max_level`.
    points_per_level: float = 0.0


class SpellConstant(BaseModel):
    name: str
    rank: int
    school_mask: int
    cast_time_ms: int
    gcd_ms: int
    cooldown_ms: int
    category_cooldown_ms: int
    duration_ms: int
    cost: int
    cost_type: int
    spell_level: int
    family_mask: list[int]
    effects: list[SpellEffectConstant]
    #: SpellPower.PowerCostPct: the cost as a percentage of base mana for the
    #: spells the client prices that way (Arcane Blast 15, Judgement 6, the
    #: warlock summons 80/100) instead of through the flat `cost` column,
    #: which reads 0 for them. Appended last: the emitted key order is the
    #: generated files' only compatibility surface, so new fields go on the
    #: end and existing ones never move. 0 when the spell has a flat cost or
    #: none.
    cost_pct: float = 0.0
    #: SpellLevels.MaxLevel: the highest caster level that still adds
    #: `points_per_level` to an effect. 0 when the client states no cap.
    max_level: int = 0


class ClassSpellConstants(BaseModel):
    build: str
    class_slug: str
    family: int
    spells: dict[str, SpellConstant]


class LevelStats(BaseModel):
    level: int
    health: int
    mana: int
    agility: int
    strength: int
    intellect: int
    spirit: int
    stamina: int
    #: Intellect needed for 1% spell crit at this level, from wowhead's
    #: `critSpell`. None for a class the payload carries no spell-crit curve
    #: for at all (warrior, rogue): "no data" rather than "zero".
    spell_crit_per_int: float | None


class ClassLevels(BaseModel):
    levels: list[LevelStats]


class RaceStatOffsets(BaseModel):
    agility: int
    strength: int
    intellect: int
    spirit: int
    stamina: int


class LevelsSource(BaseModel):
    url: str
    fetched_at: str


class LevelsFile(BaseModel):
    build: str
    source: LevelsSource
    classes: dict[str, ClassLevels]
    race_offsets: dict[str, RaceStatOffsets]


class SpellRank(BaseModel):
    id: int
    rank: int
    level: int
    #: An Ahn'Qiraj book rank (curated/book-ranks.json). Written only when true.
    book: bool = False
    #: A rank weaker than the one before it (curated/inferior-ranks.json). Written only when true.
    inferior: bool = False


class SpellRanksFile(BaseModel):
    build: str
    classes: dict[str, dict[str, list[SpellRank]]]


class ConsumableRecord(BaseModel):
    id: int
    name: str
    quality: int
    required_level: int
    spell_ids: list[int]


class SimBuffExclusive(BaseModel):
    """The group of debuffs that do not stack: the lowest `rank` holds."""

    group: str
    rank: int

    model_config = {"extra": "forbid"}


class SimBuffEffect(BaseModel):
    """What a buff does, for the saved result page's "No effect on your damage" rule.
    Curated in `curated/simbuffs.json`'s `effects`."""

    #: The stat block the engine measures on a level 60 sheet, in the stat
    #: weights' vocabulary, plus `target_armor` and `spell_crit`.
    stats: dict[str, float] = {}
    #: Schools whose damage it raises, whoever it lands on.
    school_damage: list[str] = []
    #: It works through a weapon swing (Flametongue Totem's proc).
    melee_effect: bool = False
    exclusive_with: SimBuffExclusive | None = None

    model_config = {"extra": "forbid"}


class SimBuffEntry(BaseModel):
    name: str
    icon: str
    #: It lands on the target, not on the player. Read off IDS.md's Debuffs rows.
    target: bool = False
    stats: dict[str, float] = {}
    school_damage: list[str] = []
    melee_effect: bool = False
    exclusive_with: SimBuffExclusive | None = None


class SimBuffsFile(BaseModel):
    """`simbuffs.json` (parity contract 10.4): the display name and icon
    for every id `sim/request/IDS.md` lets a request name, so the settings
    bar's full buff panel can render one."""

    entries: dict[str, SimBuffEntry]


class BuffOverride(BaseModel):
    """A curated IDS.md entry: the display name, and the client row whose
    icon it takes. Exactly one of the two ids. An explicit icon is not
    accepted -- art the build does not ship would reach the site as a
    404, and the point of naming a row is that the build resolves it."""

    name: str
    item_id: int = 0
    spell_id: int = 0

    model_config = {"extra": "forbid"}
