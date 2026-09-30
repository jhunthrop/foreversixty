# data/tests/test_loot_quest_faction_contract.py
"""data-followups-2 lane, 2026-09-30, item 2: the persona flagged, then
withdrew, "Friend of the Library" (a mage-arcane neck item) as Horde-only
in practice even though nothing in the pipeline states that as fact --
the kind of silent flip a stale or re-scraped classic-db cache could
produce for any single-faction quest and no existing test would catch,
because `tests/test_loot_quest_faction.py` only exercises the DERIVATION
logic (`quest_factions_from_classic_sources`) against synthetic records,
never the real committed data.

This is the missing regression: ten quests this build's own committed
classic-db cache (`builds/1.60.1.70009/raw/classicdb/sources.json`, the
`pipeline.classic_sources.fetch_classic_db_sources`/`write_classic_
sources` cache `pipeline.loot.classicdb` reads OFFLINE) names as strictly
one-faction, cross-checked against `quest_template.RequiredRaces` two
ways:

1. Independently against real WoW Classic knowledge -- every one of
   these ten is a well-known starting-zone quest (Classic's own starting
   zones are single-race, single-faction by design; a Human/Dwarf/Gnome/
   Night Elf zone quest cannot legitimately be "horde", and vice versa),
   so the expected `faction` below is not a fact this pipeline invented
   for itself, it is the game's own contract.
2. Against the committed `loot.json`'s own `quests` map, so a future
   `fetch-classic-sources` refresh that mis-decodes or drops a
   `RequiredRaces` bit for any of these ten fails THIS test, by name,
   rather than silently flipping the site's copy on the next rebuild.

Ten quests, why each is known one-faction, and the primary source that
confirms it:
* 6    Bounty on Garrick Padfoot -- Alliance: Westfall's Deputy Willem
       chain, Sentinel Hill (Human/Alliance-only zone).
* 33   Wolves Across the Border  -- Alliance: Dun Morogh (Dwarf/Gnome
       starting zone).
* 179  Dwarven Outfitters        -- Alliance: Dun Morogh starting zone.
* 579  Stormwind Library         -- Alliance: Old Blanchy chain,
       Stormwind City (Human-only capital).
* 961  Onu is meditating         -- Alliance: Shadowglen, Teldrassil
       (Night Elf starting zone; Onu is the tutorial NPC).
* 781  Attack on Camp Narache    -- Horde: Camp Narache, Mulgore (Tauren
       starting zone).
* 788  Cutting Teeth             -- Horde: Valley of Trials, Durotar
       (Orc/Troll starting zone).
* 789  Sting of the Scorpid      -- Horde: Valley of Trials, Durotar.
* 794  Burning Blade Medallion   -- Horde: Valley of Trials, Durotar.
* 804  Sarkoth                   -- Horde: Valley of Trials, Durotar
       (the zone's own elite quest).

Every one of the ten is confirmed present in the pinned cmangos/classic-
db dump (commit ec4f596146be6467ea93c57397858e329e2db852) with exactly
the `RequiredRaces`-decoded faction named below (spot-checked directly
against `builds/1.60.1.70009/raw/classicdb/sources.json` while writing
this test, 2026-09-30).
"""

import json
from pathlib import Path

BUILD = "1.60.1.70009"
BUILD_DIR = Path("builds") / BUILD
CLASSICDB_CACHE = BUILD_DIR / "raw" / "classicdb" / "sources.json"

#: quest id -> the one faction that quest is restricted to, per this
#: module's own doc above -- the primary, independently-known fact each
#: assertion below is checked against.
ONE_FACTION_QUESTS: dict[int, str] = {
    6: "alliance",
    33: "alliance",
    179: "alliance",
    579: "alliance",
    961: "alliance",
    781: "horde",
    788: "horde",
    789: "horde",
    794: "horde",
    804: "horde",
}


def _classic_db_cache() -> dict:
    return json.loads(CLASSICDB_CACHE.read_text(encoding="utf-8"))


def _quest_records_by_id() -> dict[int, dict]:
    """quest id -> the FIRST `quest_reward` record's own `quest` object
    in the committed classic-db cache -- `pipeline.classic_sources.
    quest_factions_from_classic_sources`'s own doc: every reward item a
    quest hands out carries the identical `ClassicDbQuestInfo` for that
    quest id, so the first one found settles it."""
    by_quest: dict[int, dict] = {}
    for records in _classic_db_cache()["items"].values():
        for record in records:
            quest = record.get("quest")
            if record.get("kind") == "quest_reward" and quest is not None:
                by_quest.setdefault(quest["quest_id"], quest)
    return by_quest


def _loot_quests() -> dict[str, list[dict]]:
    loot = json.loads((BUILD_DIR / "loot.json").read_text(encoding="utf-8"))
    return loot["quests"]


def test_the_pinned_classic_db_dump_still_names_these_ten_quests_one_faction():
    """The dump itself, independent of anything `build_loot` does with
    it: a `fetch-classic-sources` refresh that quietly re-decodes
    `RequiredRaces` differently for one of these ten fails here first."""
    by_quest = _quest_records_by_id()
    for quest_id, expected_faction in ONE_FACTION_QUESTS.items():
        assert quest_id in by_quest, quest_id
        assert by_quest[quest_id]["faction"] == expected_faction, quest_id


def test_the_committed_loot_json_agrees_with_classic_db_for_all_ten():
    """The site-facing fact: `loot.json`'s own `quests` map, for the
    reward item each of these ten hands out, states the SAME faction as
    classic-db's own `RequiredRaces` -- verified, not the item's own
    (possibly disagreeing) `factionRestriction` guess."""
    quests = _loot_quests()
    seen: set[int] = set()
    for entries in quests.values():
        for entry in entries:
            quest_id = entry["quest_id"]
            if quest_id in ONE_FACTION_QUESTS:
                seen.add(quest_id)
                assert entry["faction"] == ONE_FACTION_QUESTS[quest_id], quest_id
                assert entry["faction_source"] == "classic-db", (
                    quest_id,
                    "not verified against classic-db's own RequiredRaces",
                )
    missing = set(ONE_FACTION_QUESTS) - seen
    assert not missing, f"quest ids missing from loot.json entirely: {sorted(missing)}"


# data-followups-3 lane, 2026-09-30, item 2: the regression this module's
# own doc says was "flagged, then withdrawn" -- quest 78150 "Friend of the
# Library" (reward items 277203 Scholarly Pendant / 277204 Erudite's
# Amulet). Confirmed here against the SAME committed, offline inputs
# `pipeline.loot.sources.build_loot`/`resolve_quest_faction` read in CI, so
# a future refresh of any one of them that quietly starts (or stops)
# resolving this quest fails here by name, the same contract the ten
# one-faction starting quests above already get.
QUEST_78150 = 78150
ITEM_SCHOLARLY_PENDANT = 277203
ITEM_ERUDITES_AMULET = 277204
ITEMS_JSON = BUILD_DIR / "items.json"
ITEM_SOURCES_CACHE = BUILD_DIR / "raw" / "items" / "item-sources.json"


def test_neither_reward_item_carries_a_real_faction_restriction():
    """The root cause: unlike Hammerbone/Leaders of the Fang (a real
    quest-vs-item disagreement), Friend of the Library's own two reward
    items state NO `factionRestriction` at all -- `item_factions`'s own
    result never gets a key for either one, so there is no item-level
    fact to fall back on, only wowhead's and classic-db's."""
    items = {item["id"]: item for item in json.loads(ITEMS_JSON.read_text(encoding="utf-8"))}
    for item_id in (ITEM_SCHOLARLY_PENDANT, ITEM_ERUDITES_AMULET):
        assert items[item_id]["faction_restriction"] == "", item_id


def test_quest_78150_has_no_classic_db_row():
    """Confirms the brief's own claim independently: a Forever-new quest,
    absent from the pinned classic-db dump entirely -- `faction_source`
    can never be `"classic-db"` for it until/unless a future re-pin of
    `SOURCE_COMMIT` adds one."""
    by_quest = _quest_records_by_id()
    assert QUEST_78150 not in by_quest


def test_wowhead_states_both_for_this_quest_and_is_correctly_not_trusted():
    """The committed `item-sources.json` scrape (already fetched, no
    network needed here) states `faction: "both"` for both reward items
    -- confirmed real data, not a stand-in -- and `resolve_quest_faction`
    (`pipeline.loot.sources`) correctly treats this as uninformative
    (indistinguishable from "wowhead has not resolved this one either")
    rather than republishing it as fact, per that function's own doc."""
    from pipeline.item_sources import load_item_sources
    from pipeline.loot.sources import resolve_quest_faction

    item_sources = load_item_sources(BUILD_DIR)
    for item_id in (ITEM_SCHOLARLY_PENDANT, ITEM_ERUDITES_AMULET):
        entry = item_sources[item_id]
        rewards = [r for r in entry.quest_rewards if r.quest_id == QUEST_78150]
        assert len(rewards) == 1, item_id
        assert rewards[0].faction == "both", item_id

        faction, faction_source = resolve_quest_faction(
            QUEST_78150,
            item_id,
            quest_factions={},  # test_quest_78150_has_no_classic_db_row, above
            quest_levels={},  # no quest-page fetch has covered 78150 yet
            item_sources=item_sources,
            item_faction_restrictions={},  # test_neither_reward_item_..., above
        )
        assert faction == "unknown", item_id
        assert faction_source == "item", item_id
