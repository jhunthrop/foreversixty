"""The lane `bis-data` addition to loot.json (parity contract 6.1, grown):
the `vendor` and `zone` source kinds, the per-item `quests` map, and the
`factions` map, all read off the committed build the way test_loot_build.py
reads everything else.

Like test_loot_build.py, this needs no engine checkout or raw/ CSVs: it
reads the already-written build. (The lane's origin/main regression
checks -- every pre-existing kind names the same items -- lived here until
the lane merged; main is now the baseline they compared against.)
"""

import json
from functools import cache
from pathlib import Path

BUILD = "1.60.1.70009"
BUILD_DIR = Path("builds") / BUILD


@cache
def loot() -> dict:
    return json.loads((BUILD_DIR / "loot.json").read_text(encoding="utf-8"))


@cache
def items_by_id() -> dict[str, dict]:
    rows = json.loads((BUILD_DIR / "items.json").read_text(encoding="utf-8"))
    return {str(row["id"]): row for row in rows}


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
