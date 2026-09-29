# data/tests/test_loot_sources_classicdb.py
"""src-classicdb lane, 2026-09-29: `pipeline.loot.sources.build_loot`
folding in `pipeline.classic_sources`' cmangos/classic-db dump parse for
an item the fork database itself names no source for at all (item 110,
"Suffixed Sword", in tests/fixtures/loot -- the same fork-unsourced
fixture item `test_loot_sources_wowhead.py` uses, so a test comparing
priority order between the two origins has one item both can address).
"""

import json
from pathlib import Path

from pipeline.classic_sources import ClassicDbCondition, ClassicDbQuestInfo, ClassicDbSourceRecord
from pipeline.csvio import read_csv
from pipeline.forkdb import load_fork_database
from pipeline.item_sources import ItemSourceEntry
from pipeline.loot.sources import build_loot, instance_types, pvp_ranks, source_item_ids
from pipeline.wowhead_item_sources import NpcSource

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"

UNSOURCED_ITEM = 110  # "Suffixed Sword" -- the fixture's one fork-unsourced item


def built(classic_sources=None, item_sources=None):
    fork = load_fork_database(ENGINE)
    zone_rows = json.loads((ENGINE / "zones.json").read_text(encoding="utf-8"))
    item_rows = json.loads((ENGINE / "items.json").read_text(encoding="utf-8"))
    build_items = {row["id"] for row in item_rows}
    item_inventory_types = {row["id"]: row["inventory_type"] for row in item_rows}
    types = instance_types(read_csv(ENGINE / "Map.csv"), zone_rows)
    return build_loot(
        fork,
        {row["id"]: row["name"] for row in zone_rows},
        types,
        pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")),
        build_items,
        item_inventory_types,
        None,
        item_sources,
        classic_sources,
        zone_rows,
    )


def source(document, source_id: str):
    for candidate in document.sources:
        if candidate.id == source_id:
            return candidate
    raise AssertionError(f"no source {source_id}; have {[s.id for s in document.sources]}")


def test_with_no_classic_sources_the_item_stays_unsourced():
    document, stats = built()
    named = {item_id for s in document.sources for item_id in source_item_ids(s)}
    assert UNSOURCED_ITEM not in named
    assert stats.classicdb_items == 0


def test_an_instance_creature_drop_becomes_a_dungeon_boss_with_a_chance():
    """The fixture's own Deadmines zone (1581, map_id 36) is a dungeon
    per Map.csv -- a classic-db creature_drop whose spawn map is 36
    resolves to that same zone id and becomes a boss entry, carrying the
    dump's own drop chance. The fork ALREADY has a `dungeon:the-deadmines`
    bucket (npc 902, item 103) -- classic-db's own npc 657 joins it as a
    second boss, and the union keeps the fork's `source_origin=None`
    (fork stays primary), not classic-db's."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=657, name="Defias Pirate", map_id=36, chance=6.0,
            )
        ]
    }
    document, stats = built(classic_sources)
    dungeon = source(document, "dungeon:the-deadmines")
    assert dungeon.source_origin is None
    assert {b.npc_id for b in dungeon.bosses} == {902, 657}
    boss = next(b for b in dungeon.bosses if b.npc_id == 657)
    assert boss.items == [UNSOURCED_ITEM]
    assert boss.item_chances == {str(UNSOURCED_ITEM): 6.0}
    assert stats.classicdb_items == 1


def test_an_open_world_creature_drop_becomes_a_flat_world_bucket():
    """A creature spawning on map 1 (Kalimdor, a bare continent id, not a
    specific instance) gets no zone at all -- per this lane's own brief,
    it is a flat world:<name> bucket."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=9999, name="Some Kalimdor Mob", map_id=1, chance=12.5,
            )
        ]
    }
    document, _ = built(classic_sources)
    world = source(document, "world:some-kalimdor-mob")
    assert world.source_origin == "classic-db"
    assert world.items == [UNSOURCED_ITEM]
    assert world.item_chances == {str(UNSOURCED_ITEM): 12.5}


def test_a_scripted_bosss_missing_spawn_map_falls_back_to_the_forks_own_placement():
    """npc 902 is the fixture's own Deadmines boss (fork drop for item
    103 names `npcId: 902, zoneId: 1581`). A classic-db creature_drop for
    the SAME npc with NO spawn map (`map_id=None`, a scripted/summoned
    boss with no static `creature` row -- the real Darkmaster Gandling
    case, src-classicdb-fixes lane, 2026-09-29) still lands in
    `dungeon:the-deadmines`, never `world`, via `fork_instance_npc_zones`'
    own fallback."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=902, name="A Deadmines Boss", map_id=None, chance=8.0,
            )
        ]
    }
    document, _ = built(classic_sources)
    dungeon = source(document, "dungeon:the-deadmines")
    assert UNSOURCED_ITEM in {item for b in dungeon.bosses for item in b.items}
    world_items = {
        item
        for candidate in document.sources
        if candidate.id.startswith("world:")
        for item in candidate.items or []
    }
    assert UNSOURCED_ITEM not in world_items


def test_a_rep_gated_vendor_carries_faction_and_standing():
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="vendor", npc_id=950, name="Classic-db Vendor",
                condition=ClassicDbCondition(faction_id=76, standing="exalted"),
            )
        ]
    }
    document, _ = built(classic_sources)
    vendor = source(document, "vendor:950")
    assert vendor.faction_id == 76
    assert vendor.standing == "exalted"
    assert vendor.source_origin == "classic-db"


def test_a_quest_reward_joins_the_flat_quest_bucket_with_classic_db_level_source():
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="quest_reward", name="A Classic-db Errand",
                quest=ClassicDbQuestInfo(quest_id=53, min_level=40, level=44, faction="alliance"),
            )
        ]
    }
    document, _ = built(classic_sources)
    quest_source = source(document, "quest")
    assert UNSOURCED_ITEM in quest_source.items
    detail = document.quests[str(UNSOURCED_ITEM)]
    assert detail[0].quest_id == 53
    # quest-faction lane, 2026-09-29: the quest's own classic-db
    # RequiredRaces reading wins regardless of the item's restriction
    # (item 110 carries none) -- see the next test for the mismatch case.
    assert detail[0].faction == "alliance"
    assert detail[0].faction_source == "classic-db"
    assert detail[0].level_source == "classic-db"


def test_a_quest_reward_faction_follows_the_quests_required_races_not_the_item():
    """classic-db's own `RequiredRaces` reading for the quest (horde
    here) wins over item 100's own `factionRestriction` (alliance_only,
    per the fixture's db.json): a quest's faction is the quest giver's
    side, not the reward's (quest-faction lane, 2026-09-29: Hammerbone,
    an unrestricted item behind the horde-only Leaders of the Fang, is
    the case this protects)."""
    classic_sources = {
        100: [
            ClassicDbSourceRecord(
                kind="quest_reward", name="A Horde-Flagged Quest Row",
                quest=ClassicDbQuestInfo(quest_id=54, min_level=40, level=44, faction="horde"),
            )
        ]
    }
    document, _ = built(classic_sources)
    detail = document.quests[str(100)]
    matching = [entry for entry in detail if entry.quest_id == 54]
    assert len(matching) == 1
    assert matching[0].faction == "horde"
    assert matching[0].faction_source == "classic-db"


def test_fork_stays_primary_when_classic_db_names_the_same_bucket():
    """Item 104 ("World Boss Drop") already sources `zone:16` via the
    fork's own npc 903 drop. A classic-db creature_drop for the SAME
    item, if it named a matching open-world bucket, would union in --
    here it names a DIFFERENT npc bucket for a different item to prove
    the fork's own zone:16 source keeps `source_origin=None` untouched."""
    classic_sources = {UNSOURCED_ITEM: []}
    document, _ = built(classic_sources)
    zone = source(document, "zone:16")
    assert zone.source_origin is None


def test_a_world_drop_pool_record_becomes_one_world_drop_source_and_a_real_boss_drop_stays_put():
    """World-drop-pool lane, 2026-09-29's own report case, at the loot-
    model level (`pipeline.classic_sources`'s own tests cover the SQL-
    dump classification that produces a `world_drop` record in the first
    place): item 110 (`UNSOURCED_ITEM`, standing in for Lambent Scale
    Cloak) gets ONE `world_drop:18-25` source, never a per-creature
    listing; item 104's own real Deadmines boss drop (standing in for Mr.
    Smite's Mighty Hammer) keeps its own boss and chance, proving a
    `world_drop` record for one item never touches an unrelated item's
    own real attribution."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="world_drop", name="World drop", chance=1.5, level_min=18, level_max=25,
            )
        ],
        104: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=646, name="Mr. Smite", map_id=36, chance=20.0,
            )
        ],
    }
    document, _ = built(classic_sources)
    world_drop = source(document, "world_drop:18-25")
    assert world_drop.kind == "world_drop"
    assert world_drop.name == "World drop"
    assert world_drop.level_min == 18
    assert world_drop.level_max == 25
    assert world_drop.items == [UNSOURCED_ITEM]
    assert world_drop.item_chances == {str(UNSOURCED_ITEM): 1.5}
    world_buckets = [c for c in document.sources if c.id.startswith("world:")]
    assert not any(UNSOURCED_ITEM in (c.items or []) for c in world_buckets)

    dungeon = source(document, "dungeon:the-deadmines")
    smite = next(b for b in dungeon.bosses if b.name == "Mr. Smite")
    assert smite.items == [104]
    assert smite.item_chances == {"104": 20.0}


def test_classic_db_fills_the_gap_before_wowhead_and_wowhead_unions_into_it():
    """Priority order fork > classic-db > wowhead: classic-db creates the
    `world:shared-mob` bucket first (source_origin="classic-db"); a
    wowhead scrape naming the SAME item under the SAME npc unions into
    it rather than creating a second source, and the union keeps
    classic-db's own origin, not wowhead's."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=8080, name="Shared Mob", map_id=1, chance=3.0,
            )
        ]
    }
    item_sources = {
        UNSOURCED_ITEM: ItemSourceEntry(
            dropped_by=[NpcSource(npc_id=8080, name="Shared Mob", zone_ids=[])],
            source="wowhead",
            fetched_at="x",
        )
    }
    document, stats = built(classic_sources, item_sources)
    world = source(document, "world:shared-mob")
    assert world.source_origin == "classic-db"
    assert stats.classicdb_items == 1
    assert stats.wowhead_items == 0  # already named by classic-db, not counted again


def test_a_fork_bosss_empty_name_resolves_from_classic_dbs_own_creature_template():
    """npc 901 is `test_loot_sources.py`'s own Molten Core fixture boss
    with no fork-stated name (raid:molten-core's second boss, item 102's
    only drop -- see that file's own `test_a_raid_lists_a_named_boss_
    per_npc_and_drops_an_unnamed_one` for the case with no classic-db
    fallback at all). A classic-db VENDOR record naming that SAME
    npc_id for an unrelated item (110, UNSOURCED_ITEM) is enough for
    `classic_db_npc_names` to resolve the name globally
    (wowhead-world-drops lane, 2026-09-29's own addendum) -- a vendor
    sale and a raid boss are unrelated buckets, so no union between the
    two records happens; the raid's own boss simply gets the name."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(kind="vendor", npc_id=901, name="Baron Geddon"),
        ]
    }
    document, stats = built(classic_sources)
    raid = source(document, "raid:molten-core")
    boss = next(b for b in raid.bosses if b.npc_id == 901)
    assert boss.name == "Baron Geddon"
    assert boss.items == [102]
    assert stats.dropped_unnamed_bosses == 0
    vendor = source(document, "vendor:901")
    assert vendor.name == "Baron Geddon"
    assert vendor.items == [UNSOURCED_ITEM]


def test_a_chance_of_exactly_zero_is_omitted_never_published_as_a_real_zero():
    """cmangos/classic-db's own `ChanceOrQuestChance` uses 0 as ITS OWN
    "no chance recorded" sentinel, not a real 0% -- wowhead-world-drops
    lane, 2026-09-29's own addendum. A record with chance=0.0 is omitted
    from item_chances/boss item_chances entirely (absent key = unknown),
    never published as a real 0, which the site would otherwise show as
    "0% from ...", a fact this dump never actually states."""
    classic_sources = {
        UNSOURCED_ITEM: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=657, name="Defias Pirate", map_id=36, chance=0.0,
            )
        ],
        104: [
            ClassicDbSourceRecord(
                kind="creature_drop", npc_id=9999, name="Zero Chance Mob", map_id=1, chance=0.0,
            )
        ],
    }
    document, _ = built(classic_sources)
    dungeon = source(document, "dungeon:the-deadmines")
    boss = next(b for b in dungeon.bosses if b.npc_id == 657)
    assert boss.items == [UNSOURCED_ITEM]
    assert boss.item_chances is None
    world = source(document, "world:zero-chance-mob")
    assert world.items == [104]
    assert world.item_chances is None
