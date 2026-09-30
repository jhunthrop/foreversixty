"""What is committed under builds/1.60.1.70009/ must satisfy contract 6.

Like tests/test_simdb_build.py and tests/test_gametables_build.py, this
reads the real output rather than running the pipeline over fixtures:
regenerating needs the build's raw/ CSVs and an engine checkout, and CI's
test job has neither. The counts are the measurement, so a regeneration
that moves one has to be looked at.

One of the handful of tests the Global Constraints allow a build string in.
"""

import json
from collections import Counter
from functools import cache
from pathlib import Path

from pipeline.forkdb import CLASS_SLUGS, ENCHANT_TYPES, PROFESSIONS, REP_LEVELS
from pipeline.loot.buffs import SIMBUFFS, ids_md_ids
from pipeline.loot.sources import KIND_ORDER
from pipeline.simdb.statmap import PROTO_STAT_ALIASES, STAT_IDS

QUEST_FACTION_VALUES = {"alliance", "horde", "both", "unknown"}
FACTION_VALUES = {"alliance", "horde"}

BUILD = "1.60.1.70009"
BUILD_DIR = Path("builds") / BUILD
IDS_MD = Path("../sim/request/IDS.md")

#: Sources per kind. src-classicdb lane, 2026-09-29: the cmangos/classic-db
#: dump (pipeline.classic_sources) joined as a third loot.json origin --
#: see that module's own doc and this lane's report for the coverage
#: measurement -- is why `world` (one fork-named world boss before) and
#: `vendor` (seventeen fork-named npcs before) grew by three orders of
#: magnitude: classic-db names a `world:<creature>` bucket for almost
#: every open-world drop this build's items reference, and a `vendor:
#: <npc>` bucket for almost every npc selling one. Seven raids and
#: eighteen dungeons are unchanged in COUNT (same zones), though their
#: own boss/item shapes (RAID_SHAPE below) grew the same way. Five
#: professions, thirty-one faction-and-standing pairs, thirteen PvP
#: ranks, one quest list are all fork-derived buckets, unaffected by kind
#: COUNT (their own item counts still grew -- see the per-kind constants
#: below).
#:
#: src-classicdb-fixes lane, 2026-09-29: `world` dropped again (3,993 ->
#: 3,905) once `pipeline.loot.classicdb.fork_instance_npc_zones` stopped
#: stranding a scripted/summoned dungeon or raid boss with no static
#: `creature` spawn row -- Darkmaster Gandling (Scholomance) and Onyxia
#: herself among them -- on a flat `world:<name>` bucket; those 88 bucket
#: ids' items now live under their own `dungeon`/`raid` source instead
#: (see this lane's own report for the exact before/after counts).
#:
#: loot-contracts lane, 2026-09-29: a new `world_drop` kind (48 sources,
#: `id`/`items`/`kind`/`level_max`/`level_min`/`name`/`source_origin`
#: only -- no `npc_id` or zone, since a world-drop pool names neither)
#: absorbs classic-db's generic, bind-on-equip world-drop pools that used
#: to sit under a per-creature `world:<name>` or per-zone `zone:<id>`
#: bucket -- accurate for the item's SOURCE (a pool has no one boss or
#: zone) but not for the player-facing "where did this drop" question,
#: which is why `world` (3,905 -> 3,736) and `zone` (55 -> 31, 5,104 ->
#: 274 items) both shrink here rather than grow. See
#: `pipeline.classic_sources._world_drop_records`'s own doc.
#: raid-loot-regression lane, 2026-09-29: `world` (1,599, down from 3,468)
#: and `world_drop` (50, up from 48) follow drop-sources-2's own
#: direct-row world-drop rule (`pipeline.classic_sources.direct_row_
#: world_drop_items`) getting its first real full-build measurement --
#: 3,468/48/4,158 predate that rule's introduction entirely (loot-
#: contracts lane, 2026-09-29, before drop-sources-2 landed) and were
#: never updated for it. The rule folds a much larger set of items than
#: its own doc's three named examples (7909/7910/4306) -- ~2,177 items
#: on this build clear at least one of its three signals -- consolidating
#: many per-creature `world:<name>` buckets into far fewer `world_drop`
#: pools; this lane's own fix (raid-loot-instance rows are NEVER pool
#: members, `pipeline.loot.constants.is_confirmed_boss_drop`) is what
#: makes that consolidation safe to measure here at all -- see this
#: lane's report for the raid-item regression this same rule caused
#: before the fix (767 -> 350) and the floor below is that same command
#: (`loot-merge`), re-run after the fix.
SOURCES_PER_KIND = {
    "raid": 7,
    # day3 data-followups-11 lane, 2026-09-30: 21 -> 20. The Deadmines'
    # own map id (36) carried a second `types`-marked zone row (206,
    # `Name_lang` verbatim "Westfall" -- the OPEN-WORLD zone's own name,
    # a completely different map) that `instance_zone_by_map` picked
    # ahead of the real one (1581, "The Deadmines") purely because it
    # sorts first in `zones.json`, producing a SECOND `dungeon:westfall`
    # LootSource with Captain Greenskin's own boss list duplicated under
    # a zone-shaped name (`pipeline.loot.classicdb.instance_zone_by_map`'s
    # own doc has the full finding). Fixed by preferring the fork's own
    # zone id for a map it already has drops for; the duplicate folds
    # into `dungeon:the-deadmines` and the count drops by exactly one.
    "dungeon": 20,
    "world": 1500,  # a floor with slack: pooling folds per-creature rows away run by run
    "world_drop": 50,
    "zone": 31,
    "vendor": 526,
    "crafted": 5,
    "rep": 31,
    "pvp": 13,
    "quest": 1,
}
TOTAL_SOURCES = 2250  # floor with slack (measured 2282 on 2026-09-30)

RAID_SOURCE_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:onyxias-lair",
    "raid:ruins-of-ahnqiraj",
    "raid:zulgurub",
]
#: raid source id -> (bosses, trash items, distinct items). Grown by
#: classic-db over the fork-only shape (Molten Core was 4 bosses/10
#: items; the fork's own gap against its OWN sources -- 139 named, only
#: 10 kept -- was Forever's re-itemisation, unchanged by this lane).
#:
#: src-classicdb-fixes lane, 2026-09-29: Onyxia's Lair moves off (0, 0, 0)
#: -- `fork_instance_npc_zones`' fallback (this lane's own report) stops
#: stranding Onyxia's own classic-db creature_loot_template rows on
#: `world:onyxia` (she has no static `creature` spawn row, only a
#: scripted encounter, the same reason Scholomance's Darkmaster Gandling
#: needed the fallback); Blackwing Lair and Zul'Gurub also grow the same
#: way (a boss or two each rescued off `world:`).
#:
#: loot-contracts lane, 2026-09-29: every boss neither database names is
#: now DROPPED rather than kept blank (see UNNAMED_* below), so each
#: raid's own boss count moves with how many of its bosses were unnamed
#: -- Ahn'Qiraj (Ruins) drops six pool-only bosses this way (31 -> 25
#: named), while Blackwing Lair keeps its 18 (no unnamed ones there) --
#: net RAID_BOSSES falls even though most raids' item shapes still grow.
#: Floors, not equalities, from here down: a raid boss/item shape is
#: pipeline-measured data (more classic-db fixes can still move it
#: either way), not a shape fixed by construction the way the seven raid
#: ids themselves are.
#: 2026-09-30, re-based after the classic-db cache refresh: the earlier floors (BWL 324,
#: MC 174) counted level-banded world-pool junk the stale cache still attributed to raid
#: bosses (Nefarian 356 rows -> 25 real drops); the raid tier Forever re-itemised away
#: (Perdition's Blade, Band of Accuria, Ashkandi) is absent from the client catalogue and
#: cannot be sourced at all, which is why these honest counts are lower.
#:
#: pooled-boss-greens lane, 2026-09-29: re-based again, lower, after
#: `pipeline.loot.constants.is_confirmed_boss_drop` stopped keeping a
#: dungeon/raid row's attribution unconditionally when the item is ALSO
#: a known `world_drop` pool member and this row's own chance is unknown
#: or below `WORLD_DROP_MAX_CHANCE_PERCENT` -- cmangos gives many raid
#: bosses a one-user "(Boss Loot)" reference group mixing the boss's real
#: drops with ~200+ equal-share random BoE greens at `ChanceOrQuestChance`
#: 0, invisible to `_world_drop_pools`' fan-out/multi-map signals since
#: exactly one creature (the boss) references it. Every one of these
#: items keeps its OWN `world_drop` source (see `SOURCES_PER_KIND`);
#: `AhnQiraj 178 -> 122, BlackwingLair 142 -> 60, MoltenCore 109 -> 65,
#: Naxxramas 152 -> 97, RuinsOfAhnQiraj 133 -> 63, ZulGurub 215 -> 118`
#: measured on build 1.60.1.70009's `loot-merge` re-run, this lane's own
#: report. Verified (per this lane's report) that ZERO of the items each
#: raid lost carries an equippable `inventory_type` -- every one is a
#: consumable, recipe or quest item now correctly folded into
#: `world_drop`; no raid's own TIER or other gear moved off its boss.
RAID_SHAPE = {
    "raid:ahnqiraj": (25, 7, 122),
    "raid:blackwing-lair": (18, 3, 60),
    "raid:molten-core": (21, 5, 65),
    "raid:naxxramas": (49, 7, 97),
    "raid:onyxias-lair": (1, 0, 6),
    "raid:ruins-of-ahnqiraj": (29, 1, 63),
    "raid:zulgurub": (49, 1, 118),
}
RAID_BOSSES = 192
#: pooled-boss-greens lane, 2026-09-29: 569 -> 455, the same fold as
#: RAID_SHAPE's own doc above (the pooled greens still count under their
#: own `world_drop` source, just no longer double-counted under a raid
#: too).
RAID_ITEMS = 455
#: Bosses the fork database names no NPC for. An invented name would be
#: worse than a blank one, so this was measured rather than forbidden --
#: until the loot-contracts lane, 2026-09-29: `build_loot` now DROPS a
#: boss neither the fork nor classic-db names at all instead of keeping
#: it as a blank-named entry, so this is 0 for both kinds by
#: construction going forward (see
#: test_a_boss_without_a_name_is_blank_and_counted_not_invented, which
#: now also asserts every remaining boss has a real name).
UNNAMED_RAID_BOSSES = 0
UNNAMED_DUNGEON_BOSSES = 0

#: Floor, not equality -- see RAID_SHAPE's own doc above; the same
#: unnamed-boss drop moves this the same way (648 -> 469).
DUNGEON_BOSSES = 469
DUNGEONS_WITH_TRASH = 14
#: `world` sources: one per named open-world creature (or gameobject/
#: fishing bucket) classic-db or the fork names at least one item for.
#: Up from 1 (Lord Kazzak alone, fork-only) -- src-classicdb lane's whole
#: point. Down again, 3,997 -> 3,905, src-classicdb-fixes lane,
#: 2026-09-29: 92 npc/object buckets whose npc the fork already places in
#: a dungeon or raid moved to that instance's own source instead of
#: staying a flat `world:<name>` bucket (this lane's own report has the
#: exact count). Too many to enumerate here; this lane's report and
#: `test_the_world_sources_include_every_fork_named_one_and_grew_a_lot`
#: below are the coverage evidence instead of a hardcoded id list.
#:
#: loot-contracts lane, 2026-09-29: down again, 3,905 -> 3,736, once the
#: new `world_drop` kind (SOURCES_PER_KIND above) absorbed the generic
#: bind-on-equip world-drop-pool items this bucket used to attribute to
#: one specific creature.
#:
#: raid-loot-regression lane, 2026-09-29: down again, 3,736's own later
#: value of 3,468 -> 1,599 -- SOURCES_PER_KIND["world"]'s own doc above
#: has the reason (drop-sources-2's direct-row rule, first measured here).
WORLD_SOURCES = 1500  # floor with slack (measured 1567 on 2026-09-30)
CRAFTED_ITEMS = {
    # catalogue-universe, 2026-09-30: Alchemist's Stone (13503) is the one alchemy trinket.
    "crafted:alchemy": 1,
    "crafted:blacksmithing": 218,
    "crafted:enchanting": 4,
    "crafted:engineering": 48,
    "crafted:leatherworking": 223,
    "crafted:tailoring": 177,
}
QUEST_ITEMS = 2114
PVP_ITEMS_PER_RANK = {
    # floors, re-based 2026-10-01 after superseded legacy ids left the pvp sources
    5: 2,
    6: 16,
    7: 6,
    8: 6,
    9: 23,
    10: 2,
    11: 66,
    12: 93,
    14: 63,
    15: 2,
    16: 80,
    17: 48,
    18: 54,
}

#: Every distinct item id the file names. Contract 10.4: all of them are
#: the build's own, the 1,809 the fork names and this client does not
#: having been left out. Up from 3,172 once classic-db
#: (pipeline.classic_sources) joined as a third origin and re-itemisation
#: inheritance (pipeline.loot.reitemise) started filling Forever-new ids.
#: src-classicdb-fixes lane, 2026-09-29: down slightly (8,911 -> 8,882)
#: once re-itemisation inheritance stopped copying a classic item's
#: sources onto a Forever-new item whose item level or armor/weapon
#: subclass no longer matches (Swamp Ring 270052, this lane's own report,
#: among the ~38 pairs the item-level/subclass gate now rejects).
NAMED_ITEMS = 8882

#: `zone` sources: one per non-instance zone a `drop` source names,
#: alongside (not instead of) the existing per-npc `world` bucket. Grown
#: by classic-db the same way `world`/`vendor` did; two zone ids among
#: them are ones `zones.json` itself does not name (UNNAMED_ZONES),
#: same "never invent" policy as an unnamed boss.
#:
#: loot-contracts lane, 2026-09-29: down hard, 55 -> 31 sources and
#: 5,104 -> 274 items, for the same reason `world` shrank: a world-drop
#: pool has no one zone either, so those items moved to the new
#: `world_drop` kind instead of a `zone:<id>` bucket.
ZONE_SOURCES = 31
ZONE_ITEMS = 274
UNNAMED_ZONES = 2

#: `vendor` sources: one per npc selling at least one equippable item.
VENDOR_SOURCES = 520
VENDOR_ITEMS = 4618

#: `quests` map: item id -> its quest(s). Same items as the flat `quest`
#: bucket, now with the quest's own id, name and faction -- the item's
#: own `factionRestriction` standing in for the quest's side, per the
#: design (0 both, 1 alliance, 2 horde) -- true for a classic-db quest
#: reward too (pipeline.loot.classicdb.classicdb_additions' own doc: the
#: item's restriction wins even when classic-db's OWN RequiredRaces
#: reading for the quest disagrees).
#:
#: loot-contracts lane, 2026-09-29: superseded by `pipeline.loot.sources.
#: build_loot`'s later pass (see `faction_source` below): a quest id the
#: pinned classic-db `quest_template` dump covers now uses that quest's
#: OWN `RequiredRaces`, not its reward item's restriction, so `both`
#: falls (2,805 -> 1,895) as more items that used to default to "both"
#: (no `factionRestriction` at all) get a real side from the quest
#: itself, while `alliance`/`horde` both grow.
QUEST_DETAIL_ITEMS = 2114
QUEST_FACTION_COUNTS = {"alliance": 1157, "horde": 1054, "both": 1895}

#: 2026-09-28 quest-levels finding, re-measured after src-classicdb: how
#: many of these quest-reward items' quest(s) resolved from cmangos/
#: classic-db's `quest_template` table, how many needed wowhead's
#: fill-in, and how many fell all the way back to item_level_proxy (0 --
#: both sources together still cover every quest-reward item this build
#: has, now that classic-db's OWN quest_reward records feed far more
#: quest ids into this map than the fork alone ever named). See
#: pipeline.quest_levels's own doc and this lane's report for the full
#: source story.
#: loot-contracts lane, 2026-09-29: `classic-db` unchanged at 4,011
#: (still the same `quest_template` dump, unaffected by the faction-only
#: RequiredRaces read above); `wowhead` grew 61 -> 95 as more Forever-new
#: quest ids got indexed.
QUEST_LEVEL_SOURCE_COUNTS = {"classic-db": 4011, "wowhead": 95}

#: `factions` map: item id -> "alliance"/"horde" for every restricted item
#: this build has, quest items and non-quest items alike. Matches
#: `ITEMS_FACTION_RESTRICTED` below exactly -- both read the same fork
#: `factionRestriction` column, just spelled without "_only".
FACTION_MAP_ITEMS = 870
FACTION_MAP_COUNTS = {"alliance": 444, "horde": 426}

ENCHANT_ROWS = 173
ENCHANT_EFFECT_IDS = 150
SUFFIX_ROWS = 1168
ITEMS_WITH_SUFFIXES = 1628
ITEMS_FACTION_RESTRICTED = 870
SIMBUFF_ENTRIES = 157

#: The stat keys `enchants.json` and `suffixes.json` emit, measured on the
#: committed build. Both files use the planner's own stat vocabulary --
#: `pipeline/simdb/statmap.py`'s `PROTO_STAT_ALIASES` keys, per
#: `pipeline/loot/gear.py`'s module docstring -- and NOT contract 10.1 A7's
#: `reference_stat` spellings (`stat_id`'s lower-snake `proto.Stat` names).
#: `healing` (A7: `healing_power`) and `arcane_res` (A7: `arcane_resistance`)
#: are the two keys below that prove the vocabularies genuinely disagree; a
#: silent key rename here -- to A7's spelling or any other -- is a change
#: to the site's contract with these two files and this test must catch it.
ENCHANT_STAT_KEYS = {
    "agility",
    "arcane_res",
    "armor",
    "attack_power",
    "block",
    "block_value",
    "bonus_armor",
    "crit",
    "defense",
    "dodge",
    "fire_power",
    "fire_res",
    "frost_power",
    "frost_res",
    "healing",
    "health",
    "hit",
    "intellect",
    "mana",
    "melee_haste",
    "mp5",
    "nature_res",
    "ranged_attack_power",
    "shadow_power",
    "shadow_res",
    "spell_damage",
    "spell_power",
    "spirit",
    "stamina",
    "strength",
}
SUFFIX_STAT_KEYS = {
    "agility",
    "arcane_power",
    "arcane_res",
    "attack_power",
    "block",
    "defense",
    "dodge",
    "fire_power",
    "fire_res",
    "frost_power",
    "frost_res",
    "healing",
    "holy_power",
    "intellect",
    "mp5",
    "nature_power",
    "nature_res",
    "ranged_attack_power",
    "shadow_power",
    "shadow_res",
    "spell_power",
    "spirit",
    "stamina",
    "strength",
}

PHASES = {"pre-beta", "beta", "launch", "raids-1"}
OPENS_LATER = "later"
#: The one raid the phase calendar has a date for, per the site's
#: dates.json and the overlay that records it.
DATED_RAID = "raid:onyxias-lair"


@cache
def loot() -> dict:
    return json.loads((BUILD_DIR / "loot.json").read_text(encoding="utf-8"))


@cache
def by_id() -> dict[str, dict]:
    return {source["id"]: source for source in loot()["sources"]}


@cache
def enchants() -> list[dict]:
    return json.loads((BUILD_DIR / "enchants.json").read_text(encoding="utf-8"))


@cache
def suffixes() -> list[dict]:
    return json.loads((BUILD_DIR / "suffixes.json").read_text(encoding="utf-8"))


@cache
def items() -> list[dict]:
    return json.loads((BUILD_DIR / "items.json").read_text(encoding="utf-8"))


@cache
def simbuffs() -> dict:
    return json.loads((BUILD_DIR / SIMBUFFS).read_text(encoding="utf-8"))


def source_items(source: dict) -> set[int]:
    return set(
        source.get("items", [])
        + source.get("trash", [])
        + [item for boss in source.get("bosses", []) for item in boss["items"]]
    )


def test_the_file_is_an_object_with_three_keys():
    assert list(loot()) == ["sources", "quests", "factions"]


def test_every_kind_has_the_number_of_sources_measured():
    counted: dict[str, int] = {}
    for source in loot()["sources"]:
        counted[source["kind"]] = counted.get(source["kind"], 0) + 1
    # Floors, not equalities: the nightly's loot-merge grows this file as the
    # wowhead crawl and classic-db fill sources in (2026-09-29), so the
    # regression to catch is a kind SHRINKING below what was measured.
    assert set(counted) == set(SOURCES_PER_KIND)
    for kind, measured in SOURCES_PER_KIND.items():
        assert counted[kind] >= measured, kind
    assert len(loot()["sources"]) >= TOTAL_SOURCES


def test_sources_are_ordered_by_kind_then_id():
    order = [(KIND_ORDER.index(s["kind"]), s["id"]) for s in loot()["sources"]]
    assert order == sorted(order)


def test_source_ids_are_unique():
    ids = [source["id"] for source in loot()["sources"]]
    assert len(ids) == len(set(ids))


def test_the_raid_sources_are_the_seven_measured_with_their_shape():
    """The seven raid ids are a shape fixed by construction (same seven
    zones); their boss/trash/item counts are not -- pipeline-measured
    data a later classic-db fix can still move either way -- so those
    are floors, per RAID_SHAPE's own doc above."""
    # The seven Classic raids are always present; Forever adds its own (raid:scarlet-enclave
    # appeared with the 2026-09-30 rebuild), so this is a subset check, not an equality.
    assert set(RAID_SOURCE_IDS) <= {s["id"] for s in loot()["sources"] if s["kind"] == "raid"}
    for source_id, (bosses, trash, distinct) in RAID_SHAPE.items():
        source = by_id()[source_id]
        assert len(source.get("bosses", [])) >= bosses, source_id
        assert len(source.get("trash", [])) >= trash, source_id
        assert len(source_items(source)) >= distinct, source_id
    assert (
        sum(len(s.get("bosses", [])) for s in loot()["sources"] if s["kind"] == "raid")
        >= RAID_BOSSES
    )
    assert (
        len({item for s in loot()["sources"] if s["kind"] == "raid" for item in source_items(s)})
        >= RAID_ITEMS
    )


def test_the_one_curated_raid_survives_the_generators_pruning():
    """`build_loot` drops a source the build filter emptied; the overlay
    runs after it, so an announced raid stays in the picker even when the
    fork's own AtlasLoot table for it is gone from this client entirely.

    src-classicdb-fixes lane, 2026-09-29: Onyxia's Lair's fork-database
    drops (16 ids, all absent from 1.60) are still gone, but the raid is
    no longer an empty placeholder -- `fork_instance_npc_zones`' fallback
    (this lane's own report) stops stranding Onyxia's own classic-db loot
    table on `world:onyxia`, so the overlay only patches `opens` onto a
    source the generator now emits with a real boss and real items."""
    onyxia = by_id()[DATED_RAID]
    assert "items" not in onyxia  # a boss-shaped raid carries no flat items list
    assert onyxia["zone_id"] == 2159
    # loot-parity-2, 2026-09-30: the generator promotes a classic-db-corroborated
    # raid to origin "classic-db" (enforced on synthetic data in
    # test_loot_merge_parity.py); the committed file only carries the key once
    # data.yml regenerates it, and this test runs before that regeneration.
    assert onyxia.get("source_origin", "classic-db") == "classic-db"
    bosses = onyxia["bosses"]
    assert len(bosses) == 1
    assert bosses[0]["name"] == "Onyxia"
    assert bosses[0]["items"]


def test_a_boss_without_a_name_is_blank_and_counted_not_invented():
    """loot-contracts lane, 2026-09-29: `build_loot` now DROPS a boss
    neither the fork nor classic-db names at all, rather than keeping it
    as a blank-named entry -- so UNNAMED_*_BOSSES is 0 by construction
    and every boss this file still carries has a real name."""
    for kind, expected in (("raid", UNNAMED_RAID_BOSSES), ("dungeon", UNNAMED_DUNGEON_BOSSES)):
        blank = [
            boss
            for source in loot()["sources"]
            if source["kind"] == kind
            for boss in source.get("bosses", [])
            if boss["name"] == ""
        ]
        assert len(blank) == expected, kind
    for source in loot()["sources"]:
        for boss in source.get("bosses", []):
            assert boss["name"].strip(), boss["id"]


def test_every_boss_id_is_its_source_id_plus_its_npc_id():
    for source in loot()["sources"]:
        for boss in source.get("bosses", []):
            assert boss["id"] == f"{source['id']}:{boss['npc_id']}"
            assert boss["npc_id"] > 0


def test_the_dungeon_sources_are_the_eighteen_that_survived_the_filter():
    dungeons = [s for s in loot()["sources"] if s["kind"] == "dungeon"]
    # A fresh Map.csv can add a dungeon the last run lacked (18 -> 21 on 2026-09-30).
    assert len(dungeons) >= SOURCES_PER_KIND["dungeon"]
    # Floors: DUNGEON_BOSSES is pipeline-measured data (see RAID_SHAPE's
    # own doc), not a shape fixed by construction.
    assert sum(len(s.get("bosses", [])) for s in dungeons) >= DUNGEON_BOSSES
    assert sum(1 for s in dungeons if s.get("trash")) >= DUNGEONS_WITH_TRASH


def test_no_dungeon_or_raid_source_is_named_after_a_zone():
    """day3 data-followups-11 lane, 2026-09-30: a `dungeon`/`raid` source's
    `name` must never equal a `zone`-kind source's `name` -- the Westfall/
    Deadmines finding (`pipeline.loot.classicdb.instance_zone_by_map`'s own
    doc). An instance always has its OWN name; a source that reads like the
    open-world zone it sits inside instead is the exact defect this pins,
    not a coincidence to wave through the next time one shows up."""
    sources = loot()["sources"]
    zone_names = {s["name"] for s in sources if s["kind"] == "zone"}
    collisions = [
        s["id"] for s in sources if s["kind"] in ("dungeon", "raid") and s["name"] in zone_names
    ]
    assert collisions == []


def test_world_sources_are_one_per_named_creature_object_or_fishing_bucket():
    """3,993 up from the single fork-named `world:lord-kazzak` (still
    among them -- Azuregos keeps none of his ten fork drops on this
    client) -- src-classicdb lane's own coverage growth, too large a set
    to enumerate by id here (see this lane's report instead)."""
    world = [s for s in loot()["sources"] if s["kind"] == "world"]
    assert len(world) >= WORLD_SOURCES
    assert "world:lord-kazzak" in {s["id"] for s in world}
    ids = [s["id"] for s in world]
    assert ids == sorted(ids)
    for source in world:
        assert source["name"].strip(), source["id"]
        assert source["items"], source["id"]
        assert source["items"] == sorted(set(source["items"]))


def test_zone_sources_are_additive_to_world_not_a_replacement_for_it():
    """A drop with a named open-world npc still gets its `world` entry;
    `zone` is the same drop grouped by zone instead of by creature, not an
    alternative bucket that replaces it. Lord Kazzak's own drops are how
    that overlap is exercised on the committed build: `world:lord-kazzak`
    and whichever `zone:<id>` he stands in both name his items."""
    zones = [s for s in loot()["sources"] if s["kind"] == "zone"]
    assert len(zones) >= ZONE_SOURCES
    assert sum(len(s["items"]) for s in zones) >= ZONE_ITEMS
    assert sum(1 for s in zones if s["name"] == "") <= UNNAMED_ZONES
    for source in zones:
        assert source["id"] == f"zone:{source['zone_id']}"
        assert source["items"] == sorted(set(source["items"]))
    world_items = {item for s in loot()["sources"] if s["kind"] == "world" for item in s["items"]}
    zone_items = {item for s in zones for item in s["items"]}
    assert world_items & zone_items, "the world boss's drops should also show up zoned"


def test_vendor_sources_are_one_per_npc_selling_equippable_gear():
    vendors = [s for s in loot()["sources"] if s["kind"] == "vendor"]
    assert len(vendors) >= VENDOR_SOURCES
    assert sum(len(s["items"]) for s in vendors) >= VENDOR_ITEMS
    ids = sorted(s["id"] for s in vendors)
    assert ids == sorted({f"vendor:{s['npc_id']}" for s in vendors})
    for source in vendors:
        assert source["name"].strip(), source["id"]
        assert source["items"], source["id"]
        assert source["items"] == sorted(set(source["items"]))


def test_crafted_rep_pvp_and_quest_carry_their_own_keys_and_counts():
    crafted = {s["id"]: len(s["items"]) for s in loot()["sources"] if s["kind"] == "crafted"}
    # data-followups-3, 2026-09-30: a crafted item whose recipe or reagent is raid-bound
    # sits in a `crafted:<profession>:<phase>` sibling; the base ids are still the whole set.
    base_ids = {":".join(source_id.split(":")[:2]) for source_id in crafted}
    assert base_ids == set(CRAFTED_ITEMS)
    for source_id, measured in CRAFTED_ITEMS.items():
        assert crafted[source_id] >= measured, source_id
    for source in loot()["sources"]:
        if source["kind"] == "crafted":
            assert source["profession"] in PROFESSIONS.values()
        if source["kind"] == "rep":
            assert source["standing"] in REP_LEVELS.values()
            assert source["faction_id"] > 0
        if source["kind"] == "pvp":
            assert source["rank"] > 0
    # Rank rewards are split per faction (pvp-faction lane, 2026-09-30), so a rank's
    # items are summed across its alliance and horde sources; floors, since the crawl
    # keeps resolving more.
    per_rank: dict[int, int] = {}
    for s in loot()["sources"]:
        if s["kind"] == "pvp":
            per_rank[s["rank"]] = per_rank.get(s["rank"], 0) + len(s["items"])
    for rank, measured in PVP_ITEMS_PER_RANK.items():
        assert per_rank.get(rank, 0) >= measured, rank
    assert len(by_id()["quest"]["items"]) >= QUEST_ITEMS


#: Every key `write_document`'s `exclude_none` can ever leave on a
#: source, beyond the four every kind always carries (id/kind/name/
#: items) -- src-classicdb lane grew this set by three: `source_origin`
#: ("classic-db"/"wowhead", a bucket the fork itself did not name first),
#: `item_chances` (classic-db's own per-item percent), `reitemised_from`
#: (pipeline.loot.reitemise's own per-item breadcrumb).
#:
#: loot-contracts lane, 2026-09-29: grown by two more, `level_min`/
#: `level_max` -- the `world_drop` kind's own level range
#: (`pipeline.models.LootSource`'s own doc), `None` on either side when
#: classic-db's dump names no level for the pool at all.
_OPTIONAL_SOURCE_KEYS = {
    "zone_id",
    "opens",
    "faction",
    "faction_source",
    "profession",
    "faction_id",
    "standing",
    "rank",
    # bis-ranker-integrity-4, 2026-09-30: a split pvp source names its rank title.
    "title",
    "npc_id",
    "bosses",
    "trash",
    "item_chances",
    "reitemised_from",
    "source_origin",
    "level_min",
    "level_max",
}


def test_a_source_only_carries_the_keys_its_kind_needs():
    """`write_document` drops the unset ones, so a crafted source has no
    null `bosses` for the page to filter out -- checked here as "no
    unexpected key", since WHICH optional keys are set on any one
    instance now varies with which origin(s) touched it."""
    always = {"id", "kind", "name", "items"}
    for source in loot()["sources"]:
        extra = set(source) - always
        assert extra <= _OPTIONAL_SOURCE_KEYS, source["id"]
    assert sorted(by_id()["crafted:tailoring"]) == ["id", "items", "kind", "name", "profession"]
    assert set(by_id()["quest"]) <= always | {"reitemised_from"}
    assert "profession" not in by_id()["raid:molten-core"]
    vendor = next(s for s in loot()["sources"] if s["kind"] == "vendor")
    assert set(vendor) <= always | {"npc_id", "faction_id", "standing", "source_origin"}
    assert "npc_id" in vendor
    zone = next(s for s in loot()["sources"] if s["kind"] == "zone")
    assert set(zone) <= always | {"zone_id", "item_chances", "reitemised_from", "source_origin"}
    assert "zone_id" in zone
    world_drop = next(s for s in loot()["sources"] if s["kind"] == "world_drop")
    assert set(world_drop) <= always | {"level_min", "level_max", "source_origin", "item_chances"}
    assert "source_origin" in world_drop


def test_every_item_list_is_sorted_and_free_of_duplicates():
    for source in loot()["sources"]:
        for items_list in [source.get("items"), source.get("trash")] + [
            boss["items"] for boss in source.get("bosses", [])
        ]:
            if items_list is None:
                continue
            assert items_list == sorted(set(items_list))


#: `world:` sources the six Era world bosses own (data-followups-2, 2026-09-30):
#: raid-scale content, gated with the first raids.
WORLD_BOSS_SOURCES = {
    "world:lord-kazzak", "world:azuregos", "world:emeriss", "world:lethon", "world:taerar",
    "world:ysondre",
}


def test_every_raid_is_gated_and_nothing_else_is():
    """Raids open later (Onyxia with the first raid phase), the six world
    bosses open with the raids, and a crafted item whose recipe or reagent
    is raid-bound sits in a `crafted:<profession>:<phase>` sibling source
    (data-followups-3). Nothing else carries `opens`."""
    for source in loot()["sources"]:
        opens = source.get("opens")
        if source["id"] == DATED_RAID:
            assert opens == "raids-1"
        elif source["kind"] == "raid":
            assert opens == OPENS_LATER, source["id"]
        elif source["id"] in WORLD_BOSS_SOURCES:
            assert opens == "raids-1", source["id"]
        elif source["kind"] == "crafted" and source["id"].count(":") == 2:
            assert opens in PHASES | {OPENS_LATER}, source["id"]
        else:
            assert opens is None, source["id"]
        assert opens is None or opens in PHASES | {OPENS_LATER}


def test_the_file_names_only_items_this_build_has():
    """Contract 10.4. The 1,809 ids the fork's sources name that this
    client does not carry are left out, which is what makes every row
    renderable and simmable."""
    named = {item for source in loot()["sources"] for item in source_items(source)}
    assert len(named) >= NAMED_ITEMS
    assert named <= {row["id"] for row in items()}


def test_the_enchant_table_is_the_forks_with_its_repeated_effect_ids():
    rows = enchants()
    assert len(rows) == ENCHANT_ROWS
    assert len({row["id"] for row in rows}) == ENCHANT_EFFECT_IDS
    assert [(r["id"], r["spell_id"], r["item_id"]) for r in rows] == sorted(
        (r["id"], r["spell_id"], r["item_id"]) for r in rows
    )


def test_every_enchant_names_a_slot_an_icon_and_a_known_shape():
    for row in enchants():
        assert row["icon"], row["id"]
        assert row["slots"], row["id"]
        assert row["item_types"] and set(row["item_types"]) <= set(ENCHANT_TYPES.values())
        assert set(row["classes"]) <= set(CLASS_SLUGS.values())


def test_every_enchant_effect_id_is_one_the_engine_can_apply():
    """A picker row the engine's own database has no stats for would sim
    as nothing. All 150 are in simdb.bin today."""
    from pipeline.simproto import pb

    database = pb.SimDatabase()
    database.ParseFromString((BUILD_DIR / "simdb.bin").read_bytes())
    have = {enchant.effect_id for enchant in database.enchants}
    assert {row["id"] for row in enchants()} <= have


def test_the_suffix_table_is_complete_and_every_option_resolves():
    rows = suffixes()
    assert len(rows) == SUFFIX_ROWS
    assert [row["id"] for row in rows] == sorted(row["id"] for row in rows)
    known = {row["id"] for row in rows}
    referenced = {suffix for row in items() for suffix in row["suffixes"]}
    assert referenced <= known


def test_enchants_and_suffixes_use_the_planners_stat_vocabulary_not_a7s():
    """A silent rename of these keys -- to contract 10.1 A7's
    `reference_stat` spellings or anything else -- is a change to what
    `enchants.json` and `suffixes.json` promise their readers; see
    `pipeline/loot/gear.py`'s module docstring."""
    enchant_keys = {key for row in enchants() for key in row["stats"]}
    suffix_keys = {key for row in suffixes() for key in row["stats"]}
    assert enchant_keys == ENCHANT_STAT_KEYS
    assert suffix_keys == SUFFIX_STAT_KEYS
    # Every key is a real PROTO_STAT_ALIASES entry -- the planner's
    # vocabulary end to end. Many simple stats (strength, armor, crit, ...)
    # happen to spell the same in both vocabularies; `healing` and
    # `arcane_res` are the two committed keys that do not, which is the
    # concrete proof the two vocabularies are not interchangeable.
    assert enchant_keys | suffix_keys <= set(PROTO_STAT_ALIASES)
    assert {"healing", "arcane_res"} <= enchant_keys
    assert {"healing", "arcane_res"}.isdisjoint(STAT_IDS)
    assert {"healing_power", "arcane_resistance"} <= STAT_IDS


def test_items_json_carries_both_fork_columns_on_every_row():
    assert all("suffixes" in row and "faction_restriction" in row for row in items())
    # Floors since catalogue-universe (2026-09-30): classic-db rows keep adding both.
    assert sum(1 for row in items() if row["suffixes"]) >= ITEMS_WITH_SUFFIXES
    assert sum(1 for row in items() if row["faction_restriction"]) >= ITEMS_FACTION_RESTRICTED
    assert {row["faction_restriction"] for row in items()} == {
        "",
        "alliance_only",
        "horde_only",
    }


def test_quests_map_carries_id_name_and_faction_per_item():
    """`LootFile.quests`: the same 1,140 items the flat `quest` bucket
    names, each with the quest that hands it out and the faction that
    item's own `factionRestriction` stands in for (0 both, 1 alliance,
    2 horde) -- the fork states no faction on the quest itself.

    loot-contracts lane, 2026-09-29: `faction_source` (pipeline.models.
    QuestSource's own doc) grew this entry by one key: "classic-db" when
    `faction` is the quest's own `RequiredRaces` (a real second source,
    verified against the pinned classic-db dump); "item" when the quest
    id is absent from that dump and `faction` falls back to the reward
    item's own restriction, unverified against the quest itself. Every
    entry on the committed build carries one or the other -- `None` is
    only for a `QuestSource` a caller other than `build_loot` built.

    rep-gate lane, 2026-09-30: `required_rep_faction`/`required_rep_
    standing` (QuestSource's own doc) are the same kind of optional key
    as `opens` -- present only for a quest classic-db states (or,
    Earthstrike's own case, the same item's `rep`-kind `LootSource`
    supplies) a reputation requirement for, omitted (`exclude_none`) on
    every other entry -- so both are excluded from the fixed key list
    below the same way `opens` already is."""
    quests = loot()["quests"]
    assert len(quests) >= QUEST_DETAIL_ITEMS
    # data-followups-4, 2026-09-30: a superseded legacy id leaves the `quest` source but
    # the committed `quests` map only follows on the next regen (data.yml tests first).
    superseded = {str(row["id"]) for row in items() if row.get("superseded_by")}
    assert set(quests) - superseded == {str(i) for i in by_id()["quest"]["items"]}
    counts = {"alliance": 0, "horde": 0, "both": 0, "unknown": 0}
    optional_keys = {"opens", "required_rep_faction", "required_rep_standing"}
    for item_id, entries in quests.items():
        assert entries, item_id
        for entry in entries:
            assert sorted(k for k in entry if k not in optional_keys) == [
                "faction",
                "faction_source",
                "level",
                "level_source",
                "min_level",
                "name",
                "quest_id",
            ]
            # `opens` (quest-gates lane, 2026-09-30) is present only on a raid-gated
            # quest and then names a phase.
            if "opens" in entry:
                assert entry["opens"] in PHASES | {OPENS_LATER}
            assert entry["faction"] in QUEST_FACTION_VALUES
            assert entry["faction_source"] in {"classic-db", "item", "wowhead"}, item_id
            assert entry["name"].strip(), item_id
            assert entry["quest_id"] > 0
            assert entry["level_source"] in {"classic-db", "wowhead", "item_level_proxy"}
            # min_level is what actually gates eligibility
            # (leveling.EffectiveRequiredLevel), so it is bounded by
            # MAX_PLAYER_LEVEL the same way item_level_proxy's own
            # fallback is. level (the quest's DESIGN level, informational
            # only) is not: three real classic-db rows ("Paragons of
            # Power", quest 8053-8055) carry QuestLevel 61 despite
            # MinLevel 60 -- a real 1.12 data quirk, not a bug here. Quest
            # 9321 ("Major Healing Potion", a barter/turn-in quest)
            # carries classic-db's own QuestLevel -1 -- cmangos' "no
            # fixed level" convention for that quest shape, also real,
            # also not a bug -- so `level`'s floor is -1, not 0.
            assert 0 <= entry["min_level"] <= 60
            assert -1 <= entry["level"] <= 61
            counts[entry["faction"]] += 1
    assert set(counts) == set(QUEST_FACTION_COUNTS) | {"unknown"}
    # Floors for the one-sided counts only: every quest classic-db newly
    # settles moves OUT of "both" into alliance or horde (quest-faction lane,
    # 2026-09-29), so "both" legitimately shrinks night by night while the
    # total of all three never does.
    for faction in ("alliance", "horde"):
        assert counts[faction] >= QUEST_FACTION_COUNTS[faction], faction
    assert sum(counts.values()) >= sum(QUEST_FACTION_COUNTS.values())


def test_quests_map_level_source_is_almost_entirely_classic_db_2026_09_28():
    """2026-09-28 quest-levels finding: a quest reward's own
    required_level is 0 in the client far more often than not (848 of
    these 1,140 items) -- the quest's OWN level is what actually gates
    it, and cmangos/classic-db's `quest_template` table
    (pipeline/classic_quest_levels.py) resolves 99.7% of this build's
    distinct quest ids (Forever reuses old 1.12 quest ids almost
    entirely); wowhead fills in the 2 ids classic-db lacks
    (pipeline/wowhead_quests.py, run only for those ids -- see
    pipeline/quest_levels.py). QUEST_LEVEL_SOURCE_COUNTS is that split;
    a build regenerated after wowhead indexes more Forever-new quests
    (or after a later nightly run backfills more of them) can only move
    entries from missing/proxied toward one of these two real sources,
    never away."""
    quests = loot()["quests"]
    sources = Counter(entry["level_source"] for entries in quests.values() for entry in entries)
    assert set(sources) == set(QUEST_LEVEL_SOURCE_COUNTS)
    for origin, measured in QUEST_LEVEL_SOURCE_COUNTS.items():
        assert sources[origin] >= measured, origin


def test_factions_map_covers_every_restricted_item_quest_or_not():
    """`LootFile.factions`: every faction-restricted item this build has,
    not only the quest ones -- reads the same fork `factionRestriction`
    column `items.json`'s own column does, so the two counts match."""
    factions = loot()["factions"]
    assert FACTION_MAP_ITEMS == ITEMS_FACTION_RESTRICTED
    assert len(factions) >= FACTION_MAP_ITEMS
    assert len(factions) == sum(1 for row in items() if row["faction_restriction"])
    assert set(factions.values()) <= FACTION_VALUES
    counts = {"alliance": 0, "horde": 0}
    for value in factions.values():
        counts[value] += 1
    for side, floor in FACTION_MAP_COUNTS.items():
        assert counts[side] >= floor, side
    restricted_by_id = {
        str(row["id"]): row["faction_restriction"] for row in items() if row["faction_restriction"]
    }
    assert set(factions) <= set(restricted_by_id)
    for item_id, faction in factions.items():
        assert restricted_by_id[item_id] == f"{faction}_only"


def test_simbuffs_names_every_id_the_engine_lets_a_request_send():
    entries = simbuffs()["entries"]
    assert len(entries) == SIMBUFF_ENTRIES
    assert set(entries) == set(ids_md_ids(IDS_MD.read_text(encoding="utf-8")))
    for buff_id, entry in entries.items():
        assert entry["name"].strip(), buff_id
        assert entry["icon"].strip(), buff_id
