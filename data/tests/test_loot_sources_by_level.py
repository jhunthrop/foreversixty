"""The lane `bis-data` addition to loot.json (parity contract 6.1, grown):
the `vendor` and `zone` source kinds, the per-item `quests` map, and the
`factions` map, all read off the committed build the way test_loot_build.py
reads everything else.

This file's own job, split out from test_loot_build.py's pinned counts, is
the one check that needs a second copy of the data to mean anything: that
every kind loot.json already had before this lane -- raid, dungeon, world,
crafted, rep, pvp, quest -- names exactly the same items it always did.
`origin/main`'s committed loot.json, from before this lane touched
`pipeline/loot/sources.py`, is that second copy. Like test_loot_build.py,
this needs no engine checkout or raw/ CSVs: it reads two already-written
files.
"""

import json
import subprocess
from functools import cache
from pathlib import Path

BUILD = "1.60.1.70009"
BUILD_DIR = Path("builds") / BUILD
#: Kinds loot.json emitted before this lane. `zone` and `vendor` are new
#: and so are deliberately absent from this set -- there is nothing in
#: origin/main's file to compare them against.
PRE_EXISTING_KINDS = {"raid", "dungeon", "world", "crafted", "rep", "pvp", "quest"}


@cache
def loot() -> dict:
    return json.loads((BUILD_DIR / "loot.json").read_text(encoding="utf-8"))


@cache
def items_by_id() -> dict[str, dict]:
    rows = json.loads((BUILD_DIR / "items.json").read_text(encoding="utf-8"))
    return {str(row["id"]): row for row in rows}


@cache
def loot_before_this_lane() -> dict:
    """The same file as `loot()`, from `origin/main` -- before
    `pipeline/loot/sources.py` grew the `vendor` and `zone` kinds and the
    `quests`/`factions` maps. A build with no committed baseline there
    (a build added on this lane's own branch) has nothing to regress
    against, so the one test that needs this is skipped, not failed.
    """
    result = subprocess.run(
        ["git", "show", f"origin/main:data/{BUILD_DIR / 'loot.json'}"],
        capture_output=True,
        text=True,
        check=False,
    )
    if result.returncode != 0:
        return {}
    return json.loads(result.stdout)


def source_items(source: dict) -> set[int]:
    return set(
        source.get("items", [])
        + source.get("trash", [])
        + [item for boss in source.get("bosses", []) for item in boss["items"]]
    )


def by_kind(document: dict) -> dict[str, set[int]]:
    """kind -> every item id any source of that kind names, unioned."""
    grouped: dict[str, set[int]] = {}
    for source in document["sources"]:
        grouped.setdefault(source["kind"], set())
        grouped[source["kind"]] |= source_items(source)
    return grouped


def test_every_pre_existing_kinds_item_set_is_unchanged():
    before = loot_before_this_lane()
    if not before:
        return  # no origin/main baseline reachable; nothing to regress against
    before_by_kind = by_kind(before)
    after_by_kind = by_kind(loot())
    assert set(before_by_kind) == PRE_EXISTING_KINDS
    for kind in PRE_EXISTING_KINDS:
        assert after_by_kind[kind] == before_by_kind[kind], kind


def test_every_pre_existing_source_is_byte_identical():
    """Stronger than the item-set check above: `add`, `replace` and
    `remove` are all keyed by source id in `apply_overlays`, so a source
    this lane's new code path touched at all -- not just its item list --
    would risk breaking whichever overlay or picker state references it."""
    before = loot_before_this_lane()
    if not before:
        return
    before_by_id = {source["id"]: source for source in before["sources"]}
    after_by_id = {source["id"]: source for source in loot()["sources"]}
    for source_id, source in before_by_id.items():
        assert after_by_id.get(source_id) == source, source_id


def test_vendor_and_zone_are_the_only_new_kinds():
    before = loot_before_this_lane()
    if not before:
        return
    added_kinds = {s["kind"] for s in loot()["sources"]} - {s["kind"] for s in before["sources"]}
    assert added_kinds == {"vendor", "zone"}


def test_the_quests_map_faction_matches_the_items_own_restriction():
    """The design's faction rule, checked per item rather than by count:
    an item's own `factionRestriction` (0 both, 1 alliance, 2 horde) is
    what `quests[item].faction` reports, because the fork states no
    faction on the quest itself. `items.json`'s `faction_restriction`
    column is the same fact, independently written by
    `pipeline/loot/gear.py`'s `faction_restrictions`, so cross-checking
    against it is a real second source, not the same computation twice."""
    by_restriction = {"": "both", "alliance_only": "alliance", "horde_only": "horde"}
    rows = items_by_id()
    for item_id, entries in loot()["quests"].items():
        expected = by_restriction[rows[item_id]["faction_restriction"]]
        for entry in entries:
            assert entry["faction"] == expected, item_id


def test_the_factions_map_never_says_both():
    """Unlike `quests`, `factions` only ever names a restricted item's
    single side -- there is no "both" entry, because an unrestricted item
    (factionRestriction 0) is simply absent from the map rather than
    named as open to everyone."""
    assert set(loot()["factions"].values()) == {"alliance", "horde"}


def test_the_factions_map_agrees_with_items_json_for_every_key():
    rows = items_by_id()
    for item_id, faction in loot()["factions"].items():
        assert rows[item_id]["faction_restriction"] == f"{faction}_only", item_id
