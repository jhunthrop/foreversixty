# data/tests/test_loot_quest_faction.py
"""quest-faction lane, 2026-09-29: a quest's own classic-db
`RequiredRaces` wins over the reward item's `factionRestriction` for
`QuestSource.faction` -- the owner defect this closes: item 270018
(Hammerbone), quest 914 (Leaders of the Fang, Horde-only via
`RequiredRaces` 178), showed as `faction: "both"` in `loot.json` because
`pipeline.loot.wowhead`/`pipeline.loot.classicdb` derived a quest
reward's faction from the ITEM's own `factionRestriction`
(unrestricted for Hammerbone), never from the quest itself.

`pipeline.classic_sources.quest_factions_from_classic_sources` and its
application in `pipeline.loot.sources.build_loot` (over every
`QuestSource`, whichever scrape produced the item-quest link) are what
this file tests -- the unit-level derivation here, the precedence
`build_loot` applies in `tests/test_loot_sources_wowhead.py`'s own
fixture engine (`tests/fixtures/loot`, item 108's quest 42, item 112's
quest 43, both unrestricted).
"""

import json
from pathlib import Path

from pipeline.classic_sources import (
    ClassicDbQuestInfo,
    ClassicDbSourceRecord,
    quest_factions_from_classic_sources,
)
from pipeline.csvio import read_csv
from pipeline.forkdb import load_fork_database
from pipeline.loot.sources import build_loot, instance_types, pvp_ranks

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"

#: The fixture engine's own two quest-reward items (see
#: tests/test_loot_sources_wowhead.py's own module doc): both fork rows
#: state `factionRestriction: None`, i.e. `QUEST_FACTION_BY_RESTRICTION`'s
#: item-derived guess is "both" for either one before classic-db is
#: consulted at all.
QUEST_42_ITEM = 108
QUEST_43_ITEM = 112


def _quest_reward_record(
    quest_id: int, faction: str, *, min_level: int = 1, level: int = 1
) -> ClassicDbSourceRecord:
    return ClassicDbSourceRecord(
        kind="quest_reward",
        name=f"Quest {quest_id}",
        quest=ClassicDbQuestInfo(
            quest_id=quest_id, min_level=min_level, level=level, faction=faction
        ),
    )


def test_quest_factions_from_classic_sources_reads_every_quest_reward_record():
    """One item id's `quest_reward` record -> its quest id keyed by
    faction, regardless of which OTHER item ids share the dict -- the
    real dump keys `ClassicDbSourceRecord` lists by item id, but a
    quest's own faction is a property of the QUEST, not the item, so a
    quest handing out several reward items must resolve to the exact
    same one entry either way."""
    items = {
        270018: [_quest_reward_record(914, "horde")],
        744: [_quest_reward_record(53, "alliance"), ClassicDbSourceRecord(kind="vendor", name="x")],
        159: [_quest_reward_record(53, "alliance")],  # same quest, second reward item
    }
    factions = quest_factions_from_classic_sources(items)
    assert factions == {914: "horde", 53: "alliance"}


def test_quest_factions_from_classic_sources_required_races_zero_is_both():
    items = {5: [_quest_reward_record(1, "both")]}
    assert quest_factions_from_classic_sources(items) == {1: "both"}


def test_quest_factions_from_classic_sources_ignores_non_quest_records():
    items = {5: [ClassicDbSourceRecord(kind="vendor", name="A Vendor")]}
    assert quest_factions_from_classic_sources(items) == {}


def _built(classic_sources=None, quest_levels=None, item_sources=None):
    fork = load_fork_database(ENGINE)
    zone_rows = json.loads((ENGINE / "zones.json").read_text(encoding="utf-8"))
    item_rows = json.loads((ENGINE / "items.json").read_text(encoding="utf-8"))
    build_items = {row["id"] for row in item_rows}
    item_inventory_types = {row["id"]: row["inventory_type"] for row in item_rows}
    return build_loot(
        fork,
        {row["id"]: row["name"] for row in zone_rows},
        instance_types(read_csv(ENGINE / "Map.csv"), zone_rows),
        pvp_ranks(read_csv(ENGINE / "ItemSparse.csv")),
        build_items,
        item_inventory_types,
        quest_levels,
        item_sources,
        classic_sources,
    )


def _quest_entry(document, item_id: int, quest_id: int):
    entries = document.quests[str(item_id)]
    return next(e for e in entries if e.quest_id == quest_id)


def test_classic_db_faction_wins_over_the_items_own_unrestricted_guess():
    """The Hammerbone/Leaders-of-the-Fang shape, on the fixture engine's
    own quest 42/item 108 (fork-derived, factionRestriction None -- the
    item-derived guess would be "both"): classic-db's own RequiredRaces
    verdict for quest 42 is "horde", and it must win even though the
    item itself is unrestricted."""
    classic_sources = {QUEST_42_ITEM: [_quest_reward_record(42, "horde")]}
    document, _ = _built(classic_sources)
    entry = _quest_entry(document, QUEST_42_ITEM, 42)
    assert entry.faction == "horde"
    assert entry.faction_source == "classic-db"


def test_classic_db_required_races_zero_reports_both_and_classic_db_source():
    classic_sources = {QUEST_42_ITEM: [_quest_reward_record(42, "both")]}
    document, _ = _built(classic_sources)
    entry = _quest_entry(document, QUEST_42_ITEM, 42)
    assert entry.faction == "both"
    assert entry.faction_source == "classic-db"


def test_a_quest_id_absent_from_classic_db_and_wowhead_publishes_unknown():
    """Quest 43 (item 112) is a stand-in for a Forever-new quest id the
    pinned classic-db dump has never heard of (e.g. 78150 Friend of the
    Library, 79980 Scramble): with no classic-db coverage (an empty
    classic_sources dict) AND no `quest_levels` faction either (no
    wowhead page fetched this quest's `side` off, same as a build whose
    `quest-levels.json` predates this field), the item's own
    `factionRestriction` is NOT consulted any more (data-followups-3
    lane, 2026-09-30, item 2: that guess is not a fact about the quest,
    see `pipeline.loot.sources.build_loot`'s own doc) -- `faction`
    publishes the honest `"unknown"` instead of the old guessed
    `"both"`, still tagged unverified (`faction_source="item"`)."""
    document, _ = _built(classic_sources=None, quest_levels=None)
    entry = _quest_entry(document, QUEST_43_ITEM, 43)
    assert entry.faction == "unknown"
    assert entry.faction_source == "item"


def test_classic_db_coverage_of_other_quests_does_not_taint_an_uncovered_one():
    """classic_sources naming quest 42 must not accidentally mark quest
    43 as classic-db-verified too -- faction_source is per QUEST id, not
    a single build-wide flag."""
    classic_sources = {QUEST_42_ITEM: [_quest_reward_record(42, "horde")]}
    document, _ = _built(classic_sources)
    other = _quest_entry(document, QUEST_43_ITEM, 43)
    assert other.faction == "unknown"
    assert other.faction_source == "item"


def test_wowhead_side_wins_over_unknown_when_classic_db_has_no_coverage():
    """Quest 43, still absent from classic_sources, but `quest_levels`
    now carries the SAME quest id with a wowhead-fetched `faction`
    (`pipeline.quest_levels.QuestLevelEntry.faction`, `side` off the
    quest's own wowhead page) -- the second fallback, ahead of
    `"unknown"` and still behind classic-db, tagged
    `faction_source="wowhead"`."""
    from pipeline.quest_levels import QuestLevelEntry

    quest_levels = {
        43: QuestLevelEntry(
            min_level=1, level=1, source="wowhead", fetched_at="2026-09-30T00:00:00Z",
            faction="horde",
        )
    }  # fmt: skip
    document, _ = _built(classic_sources=None, quest_levels=quest_levels)
    entry = _quest_entry(document, QUEST_43_ITEM, 43)
    assert entry.faction == "horde"
    assert entry.faction_source == "wowhead"


def test_classic_db_still_wins_over_a_wowhead_side_value_when_both_are_present():
    """classic-db's own `RequiredRaces` is the PRIMARY source -- it must
    win even when `quest_levels` also carries a (here, deliberately
    conflicting) wowhead `faction` for the same quest id."""
    from pipeline.quest_levels import QuestLevelEntry

    classic_sources = {QUEST_42_ITEM: [_quest_reward_record(42, "horde")]}
    quest_levels = {
        42: QuestLevelEntry(
            min_level=1, level=1, source="wowhead", fetched_at="2026-09-30T00:00:00Z",
            faction="alliance",
        )
    }  # fmt: skip
    document, _ = _built(classic_sources, quest_levels)
    entry = _quest_entry(document, QUEST_42_ITEM, 42)
    assert entry.faction == "horde"
    assert entry.faction_source == "classic-db"


def test_wowhead_per_item_reward_side_wins_over_the_per_quest_page_side():
    """`item_sources[item_id].quest_rewards` (a real, already-committed
    fetch for any item the fork itself names no source for -- no new
    network call) is the more specific signal, so it wins over
    `quest_levels[quest_id].faction` (the per-QUEST page side) when both
    are present and disagree. This is the tier that resolves the
    overwhelming majority of this build's real `faction_source: "item"`
    quests (106 of 145 item/quest pairs, this lane's own report) without
    waiting on a future wowhead quest-page fetch at all."""
    from pipeline.item_sources import ItemSourceEntry
    from pipeline.quest_levels import QuestLevelEntry
    from pipeline.wowhead_item_sources import QuestRewardSource

    quest_levels = {
        43: QuestLevelEntry(
            min_level=1, level=1, source="wowhead", fetched_at="2026-09-30T00:00:00Z",
            faction="alliance",
        )
    }  # fmt: skip
    item_sources = {
        QUEST_43_ITEM: ItemSourceEntry(
            quest_rewards=[
                QuestRewardSource(
                    quest_id=43, name="Quest 43", min_level=1, level=1, faction="horde"
                )
            ],
            source="wowhead",
            fetched_at="2026-09-30T00:00:00Z",
        )
    }  # fmt: skip
    document, _ = _built(classic_sources=None, quest_levels=quest_levels, item_sources=item_sources)
    entry = _quest_entry(document, QUEST_43_ITEM, 43)
    assert entry.faction == "horde"
    assert entry.faction_source == "wowhead"


def test_wowhead_both_is_not_trusted_as_a_real_signal_and_falls_through_to_unknown():
    """The Friend of the Library shape itself: `item_sources` carries a
    `quest_rewards` row for this exact quest id, but its own `faction` is
    the generic "both" -- indistinguishable from "wowhead has not
    resolved this one either" (measured: both of the real quest's own
    reward items are wowhead-labelled "both" despite the observed
    one-reward-per-faction split in play) -- so it is skipped, same as if
    it were absent, and `faction` publishes "unknown"."""
    from pipeline.item_sources import ItemSourceEntry
    from pipeline.wowhead_item_sources import QuestRewardSource

    item_sources = {
        QUEST_43_ITEM: ItemSourceEntry(
            quest_rewards=[
                QuestRewardSource(
                    quest_id=43, name="Quest 43", min_level=1, level=1, faction="both"
                )
            ],
            source="wowhead",
            fetched_at="2026-09-30T00:00:00Z",
        )
    }  # fmt: skip
    document, _ = _built(classic_sources=None, quest_levels=None, item_sources=item_sources)
    entry = _quest_entry(document, QUEST_43_ITEM, 43)
    assert entry.faction == "unknown"
    assert entry.faction_source == "item"
