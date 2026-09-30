from typing import Literal

from pydantic import BaseModel


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


class TalentTree(BaseModel):
    id: int
    name: str
    position: int
    talents: list[TalentEntry]
    #: The tab's `BackgroundFile`, lowercased. The site fetches the processed
    #: image at `/data/<build>/trees/<background>.webp`.
    background: str


class ClassTalents(BaseModel):
    build: str
    class_id: int
    class_slug: str
    trees: list[TalentTree]


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
    required_level_source: Literal["client", "wowhead", "item_level_proxy", "none"]
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
    stats_source: Literal["wowhead"] | None = None
    set_id: int | None
    unique: bool


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
    level_min: int | None = None
    level_max: int | None = None
    #: "wowhead" or "classic-db" when this source exists ONLY because a
    #: scrape/dump named it for an item the engine fork's own database
    #: named no source for at all (night-item-sources lane, 2026-09-28;
    #: src-classicdb lane, 2026-09-29); omitted (None) for every source
    #: the fork database itself names, which is still most of them. A
    #: source more than one origin names (an item added to an existing
    #: `world:`/`vendor:`/`crafted:` bucket) keeps whichever origin found
    #: it FIRST in priority order fork > classic-db > wowhead.
    source_origin: str | None = None


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
    faction_source: Literal["classic-db", "item"] | None = None
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
    why (the curated APL's own one-line `notes`, unedited)."""

    spell_id: int
    name: str
    condition: str


class AddonRotationBand(BaseModel):
    """One level band's rotation: `lines` in priority order, each already
    resolved to the highest rank learned by `level` (pipeline.addonrotation
    does the resolving; the addon never re-derives a rank from a level)."""

    level: int
    lines: list[AddonRotationLine]


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


class AddonData(BaseModel):
    build: str
    classes: dict[str, AddonClass]
    weights: dict[str, dict[str, float]]
    #: Keyed by spec ("warrior-fury"), one band per pipeline.addonrotation.LEVEL_BANDS entry.
    rotations: dict[str, list[AddonRotationBand]] = {}
    #: Keyed by spec, one band per pipeline.addonbis.BIS_LEVEL_BANDS entry
    #: actually present in data/builds/<build>/bis/<spec>.json. Empty for
    #: every spec until the nightly workflow (.github/workflows/bis.yml)
    #: has written that directory -- pipeline.addonbis.build_bis tolerates
    #: a missing bis/ the same way build_rotations does not tolerate a
    #: missing curated/apl entry (the two directories have different
    #: maturity, hence the different rule).
    bis: dict[str, list[AddonBisBand]] = {}


class PhaseBoundary(BaseModel):
    """One content phase and the instant it opens (parity contract 10.4)."""

    name: str
    start: str


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


class SpellRanksFile(BaseModel):
    build: str
    classes: dict[str, dict[str, list[SpellRank]]]


class ConsumableRecord(BaseModel):
    id: int
    name: str
    quality: int
    required_level: int
    spell_ids: list[int]


class SimBuffEntry(BaseModel):
    name: str
    icon: str


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
