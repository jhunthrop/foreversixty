"""What is committed under builds/1.60.1.69893/ must satisfy contract 6.

Like tests/test_simdb_build.py and tests/test_gametables_build.py, this
reads the real output rather than running the pipeline over fixtures:
regenerating needs the build's raw/ CSVs and an engine checkout, and CI's
test job has neither. The counts are the measurement, so a regeneration
that moves one has to be looked at.

One of the handful of tests the Global Constraints allow a build string in.
"""

import json
from functools import cache
from pathlib import Path

from pipeline.forkdb import CLASS_SLUGS, ENCHANT_TYPES, PROFESSIONS, REP_LEVELS
from pipeline.loot.buffs import SIMBUFFS, ids_md_ids
from pipeline.loot.sources import KIND_ORDER

BUILD = "1.60.1.69893"
BUILD_DIR = Path("builds") / BUILD
IDS_MD = Path("../sim/request/IDS.md")

#: Sources per kind. Seven raids: the six the generator emits after
#: contract 10.4's build filter, plus Onyxia's Lair, which the filter
#: empties and curated/loot/forever-raid-phases.json adds back with its
#: announced phase. Five dungeons, one world boss (Lord Kazzak; Azuregos
#: keeps none of his ten drops on this client), five professions,
#: thirty-one faction-and-standing pairs, thirteen PvP ranks, one quest
#: list.
SOURCES_PER_KIND = {
    "raid": 7,
    "dungeon": 5,
    "world": 1,
    "crafted": 5,
    "rep": 31,
    "pvp": 13,
    "quest": 1,
}
TOTAL_SOURCES = 63

RAID_SOURCE_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:onyxias-lair",
    "raid:ruins-of-ahnqiraj",
    "raid:zulgurub",
]
#: raid source id -> (bosses, trash items, distinct items). The gap
#: against what the fork's sources name -- Molten Core 10 of 139,
#: Zul'Gurub 2 of 98 -- is Forever's re-itemisation, and the overlay's
#: notes carry it.
RAID_SHAPE = {
    "raid:ahnqiraj": (12, 7, 68),
    "raid:blackwing-lair": (7, 3, 21),
    "raid:molten-core": (4, 5, 10),
    "raid:naxxramas": (15, 1, 45),
    "raid:onyxias-lair": (0, 0, 0),
    "raid:ruins-of-ahnqiraj": (3, 1, 5),
    "raid:zulgurub": (1, 1, 2),
}
RAID_BOSSES = 42
RAID_ITEMS = 151
#: Bosses the fork database names no NPC for. An invented name would be
#: worse than a blank one, so this is measured rather than forbidden.
UNNAMED_RAID_BOSSES = 31
UNNAMED_DUNGEON_BOSSES = 1

DUNGEON_BOSSES = 6
DUNGEONS_WITH_TRASH = 1
WORLD_SOURCE_IDS = ["world:lord-kazzak"]
CRAFTED_ITEMS = {
    "crafted:blacksmithing": 190,
    "crafted:enchanting": 3,
    "crafted:engineering": 47,
    "crafted:leatherworking": 191,
    "crafted:tailoring": 157,
}
QUEST_ITEMS = 1058
PVP_ITEMS_PER_RANK = {5: 2, 6: 16, 7: 6, 8: 6, 9: 23, 10: 2, 11: 66, 12: 93,
                      14: 63, 15: 2, 16: 80, 17: 48, 18: 42}

#: Every distinct item id the file names. Contract 10.4: all of them are
#: the build's own, the 1,809 the fork names and this client does not
#: having been left out.
NAMED_ITEMS = 2507

ENCHANT_ROWS = 173
ENCHANT_EFFECT_IDS = 150
SUFFIX_ROWS = 1168
ITEMS_WITH_SUFFIXES = 69
ITEMS_FACTION_RESTRICTED = 819
SIMBUFF_ENTRIES = 165

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


def test_the_file_is_an_object_with_one_key():
    assert list(loot()) == ["sources"]


def test_every_kind_has_the_number_of_sources_measured():
    counted: dict[str, int] = {}
    for source in loot()["sources"]:
        counted[source["kind"]] = counted.get(source["kind"], 0) + 1
    assert counted == SOURCES_PER_KIND
    assert sum(SOURCES_PER_KIND.values()) == len(loot()["sources"]) == TOTAL_SOURCES


def test_sources_are_ordered_by_kind_then_id():
    order = [(KIND_ORDER.index(s["kind"]), s["id"]) for s in loot()["sources"]]
    assert order == sorted(order)


def test_source_ids_are_unique():
    ids = [source["id"] for source in loot()["sources"]]
    assert len(ids) == len(set(ids))


def test_the_raid_sources_are_the_seven_measured_with_their_shape():
    assert sorted(s["id"] for s in loot()["sources"] if s["kind"] == "raid") == RAID_SOURCE_IDS
    for source_id, (bosses, trash, distinct) in RAID_SHAPE.items():
        source = by_id()[source_id]
        assert len(source.get("bosses", [])) == bosses, source_id
        assert len(source.get("trash", [])) == trash, source_id
        assert len(source_items(source)) == distinct, source_id
    assert sum(
        len(s.get("bosses", [])) for s in loot()["sources"] if s["kind"] == "raid"
    ) == RAID_BOSSES
    assert len({
        item for s in loot()["sources"] if s["kind"] == "raid" for item in source_items(s)
    }) == RAID_ITEMS


def test_the_one_curated_raid_survives_the_generators_pruning():
    """`build_loot` drops a source the build filter emptied; the overlay
    runs after it, so an announced raid whose loot table nobody knows
    stays in the picker with an empty item list."""
    onyxia = by_id()[DATED_RAID]
    assert onyxia["items"] == []
    assert onyxia["zone_id"] == 2159


def test_a_boss_without_a_name_is_blank_and_counted_not_invented():
    for kind, expected in (("raid", UNNAMED_RAID_BOSSES), ("dungeon", UNNAMED_DUNGEON_BOSSES)):
        blank = [
            boss
            for source in loot()["sources"]
            if source["kind"] == kind
            for boss in source.get("bosses", [])
            if boss["name"] == ""
        ]
        assert len(blank) == expected, kind


def test_every_boss_id_is_its_source_id_plus_its_npc_id():
    for source in loot()["sources"]:
        for boss in source.get("bosses", []):
            assert boss["id"] == f"{source['id']}:{boss['npc_id']}"
            assert boss["npc_id"] > 0


def test_the_dungeon_sources_are_the_five_that_survived_the_filter():
    dungeons = [s for s in loot()["sources"] if s["kind"] == "dungeon"]
    assert len(dungeons) == SOURCES_PER_KIND["dungeon"]
    assert sum(len(s.get("bosses", [])) for s in dungeons) == DUNGEON_BOSSES
    assert sum(1 for s in dungeons if s.get("trash")) == DUNGEONS_WITH_TRASH


def test_the_only_world_source_is_the_one_world_boss_with_loot_left():
    assert [s["id"] for s in loot()["sources"] if s["kind"] == "world"] == WORLD_SOURCE_IDS


def test_crafted_rep_pvp_and_quest_carry_their_own_keys_and_counts():
    assert {
        s["id"]: len(s["items"]) for s in loot()["sources"] if s["kind"] == "crafted"
    } == CRAFTED_ITEMS
    for source in loot()["sources"]:
        if source["kind"] == "crafted":
            assert source["profession"] in PROFESSIONS.values()
        if source["kind"] == "rep":
            assert source["standing"] in REP_LEVELS.values()
            assert source["faction_id"] > 0
        if source["kind"] == "pvp":
            assert source["rank"] > 0
    assert {
        s["rank"]: len(s["items"]) for s in loot()["sources"] if s["kind"] == "pvp"
    } == PVP_ITEMS_PER_RANK
    assert len(by_id()["quest"]["items"]) == QUEST_ITEMS


def test_a_source_only_carries_the_keys_its_kind_needs():
    """`write_document` drops the unset ones, so a crafted source has no
    null `bosses` for the page to filter out."""
    assert sorted(by_id()["crafted:tailoring"]) == ["id", "items", "kind", "name", "profession"]
    assert sorted(by_id()["quest"]) == ["id", "items", "kind", "name"]
    assert "profession" not in by_id()["raid:molten-core"]


def test_every_item_list_is_sorted_and_free_of_duplicates():
    for source in loot()["sources"]:
        for items_list in [source.get("items"), source.get("trash")] + [
            boss["items"] for boss in source.get("bosses", [])
        ]:
            if items_list is None:
                continue
            assert items_list == sorted(set(items_list))


def test_every_raid_is_gated_and_nothing_else_is():
    for source in loot()["sources"]:
        opens = source.get("opens")
        if source["id"] == DATED_RAID:
            assert opens == "raids-1"
        elif source["kind"] == "raid":
            assert opens == OPENS_LATER, source["id"]
        else:
            assert opens is None, source["id"]
        assert opens is None or opens in PHASES | {OPENS_LATER}


def test_the_file_names_only_items_this_build_has():
    """Contract 10.4. The 1,809 ids the fork's sources name that this
    client does not carry are left out, which is what makes every row
    renderable and simmable."""
    named = {item for source in loot()["sources"] for item in source_items(source)}
    assert len(named) == NAMED_ITEMS
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


def test_items_json_carries_both_fork_columns_on_every_row():
    assert all("suffixes" in row and "faction_restriction" in row for row in items())
    assert sum(1 for row in items() if row["suffixes"]) == ITEMS_WITH_SUFFIXES
    assert sum(1 for row in items() if row["faction_restriction"]) == ITEMS_FACTION_RESTRICTED
    assert {row["faction_restriction"] for row in items()} == {
        "",
        "alliance_only",
        "horde_only",
    }


def test_simbuffs_names_every_id_the_engine_lets_a_request_send():
    entries = simbuffs()["entries"]
    assert len(entries) == SIMBUFF_ENTRIES
    assert set(entries) == set(ids_md_ids(IDS_MD.read_text(encoding="utf-8")))
    for buff_id, entry in entries.items():
        assert entry["name"].strip(), buff_id
        assert entry["icon"].strip(), buff_id
