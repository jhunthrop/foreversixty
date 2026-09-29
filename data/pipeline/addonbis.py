"""Leveling best-in-slot picks for the addon's item-slot hover tooltip.

Lane bis-hover-addon's own brief (docs/superpowers/specs/2026-09-28-
rotation-accuracy-program-design.md's sibling wave): "make a hover when
someone hovers over an item slot to show them what the best in slot items
are for their band." sim/cmd/leveling-bis writes one
data/builds/<build>/bis/<spec>.json per spec, nightly
(.github/workflows/bis.yml); this module turns that flat per-(band,
faction) report into the {spec -> band -> faction -> slot} shape
pipeline.addondata hands the addon, the same way pipeline.addonrotation
turns a curated APL into the rotation card's table.

Three decisions are load-bearing:

* IDS ONLY. leveling-bis's JSON carries the item's name and a prose
  swap_note the client can resolve for free (Compat.itemInfo(id)) or has
  no use for at all; carrying them again in Data.lua would roughly double
  its size for a page a player never reads offline. `source_kind` stays
  (advanced detail wants it) but travels as one of SOURCE_KIND_CODES'
  single-letter codes in the rendered Lua -- see pipeline.addonlua.
* The bis/ directory does not exist on main yet: .github/workflows/bis.yml
  writes it nightly and had not run as of this module's introduction.
  `build_bis` answers an empty mapping for a missing directory rather than
  raising, unlike pipeline.addonrotation.build_rotations, which refuses a
  curated/apl file it expects to already exist -- the two directories are
  at different points in their rollout and get different rules.
* A slot entry with no `item_id` (leveling-bis found no eligible item for
  it at that band -- an off-hand band before the spec's weapon choice
  needs one, say) names no pick and is left out, the same way
  pipeline.addonrotation leaves an unlearned line out of its band rather
  than emitting a null.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.models import AddonBisBand, AddonBisItem

#: sim/cmd/leveling-bis/main.go's own defaultBandsFlag: 10 through 60 step
#: 5. Not the rotation ladder's seven rungs (pipeline.addonrotation.
#: LEVEL_BANDS) -- leveling-bis ranks gear far more often than the ladder
#: samples accuracy, so it keeps its own, finer band list.
BIS_LEVEL_BANDS = list(range(10, 61, 5))

#: leveling-bis's closed source_kind vocabulary (sim/cmd/leveling-bis/
#: band.go's sourceKindPriority). pipeline.addonlua encodes a value from
#: this set down to the mapped single letter to keep Data.lua's bis table
#: inside its size budget; tests/tooltip_bis_spec.lua decodes the same
#: table back, independently pinned against this same literal list.
SOURCE_KIND_CODES: dict[str, str] = {
    "quest": "Q",
    "vendor": "V",
    "dungeon": "D",
    "crafted": "C",
    "rep": "R",
    "pvp": "P",
    "world": "W",
    "raid": "A",
}

#: The site's own slot vocabulary (Gear.lua's SLOTS_BY_EQUIP_LOCATION
#: values / Export.INVENTORY_SLOTS). Closed, so an unrecognised slot in a
#: bis file is a generator or schema bug, not data to render blindly --
#: Data.lua's bis table keys slots as bare Lua identifiers (no brackets or
#: quoting) to save bytes, which only a checked vocabulary can allow.
KNOWN_SLOTS = {
    "head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
    "legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
    "main_hand", "off_hand", "ranged",
}  # fmt: skip

#: The two factions leveling-bis reports, also rendered as bare Lua
#: identifiers.
KNOWN_FACTIONS = {"alliance", "horde"}


class AddonBisError(SystemExit):
    """A bis/<spec>.json says something the addon generator will not publish."""


def _slot_lookup(slots: list[dict]) -> dict[str, dict]:
    """slot name -> that slot's row, skipping a row leveling-bis wrote with
    no item_id (nothing eligible for it at this band -- see module docstring)."""
    lookup: dict[str, dict] = {}
    for entry in slots:
        if "item_id" not in entry:
            continue
        lookup[entry["slot"]] = entry
    return lookup


def _new_item_ids(new_at_band: list[str], slot_lookup: dict[str, dict]) -> list[int]:
    """leveling-bis states `new_at_band` as "<slot>: <item name>" strings
    (report.go's own display format); the slot name is the join key back to
    this band's own `slots` list, since Data.lua carries ids only."""
    ids = []
    for entry in new_at_band:
        slot = entry.split(":", 1)[0]
        item = slot_lookup.get(slot)
        if item is not None:
            ids.append(item["item_id"])
    return ids


def _significant_weights(weights: list[dict]) -> dict[str, float]:
    """leveling-bis's own `weights` list, filtered to the entries it did not
    flag `insignificant`, rounded to three decimals -- 2.0452764470665192
    on 27 specs' worth of bands is exactly the byte-budget problem
    addonbis.py's ids-only rule already solved once for item names; three
    decimals is more precision than a percentage-difference verdict
    (Tooltip.verdictFor, one decimal place) can ever show."""
    return {
        entry["stat"]: round(float(entry["weight"]), 3)
        for entry in weights
        if not entry.get("insignificant", False)
    }


def _one_band(path: Path, level: int, by_faction: dict[str, dict]) -> AddonBisBand:
    factions: dict[str, dict[str, AddonBisItem]] = {}
    new_at_band: dict[str, list[int]] = {}
    weights: dict[str, float] = {}
    for index, (faction, band) in enumerate(sorted(by_faction.items())):
        if faction not in KNOWN_FACTIONS:
            raise AddonBisError(f"{path}: unknown faction {faction!r} at band {level}")
        slot_lookup = _slot_lookup(band.get("slots", []))
        for slot in slot_lookup:
            if slot not in KNOWN_SLOTS:
                raise AddonBisError(f"{path}: unknown slot {slot!r} at band {level}/{faction}")
        factions[faction] = {
            slot: AddonBisItem(
                item_id=entry["item_id"],
                source_kind=entry["source_kind"],
                source=entry.get("source", ""),
            )
            for slot, entry in slot_lookup.items()
        }
        new_at_band[faction] = _new_item_ids(band.get("new_at_band") or [], slot_lookup)
        # The first faction (sorted, so always the same one for a given
        # band) stands for the band as a whole -- see AddonBisBand.weights'
        # own docstring on why factions are not kept apart here.
        if index == 0:
            weights = _significant_weights(band.get("weights") or [])
    return AddonBisBand(level=level, factions=factions, new_at_band=new_at_band, weights=weights)


def _spec_bis(path: Path) -> list[AddonBisBand]:
    """One spec's bis/<spec>.json, regrouped from leveling-bis's flat list
    of (band, faction) reports into one AddonBisBand per level, both
    factions inline -- the shape the addon actually indexes by."""
    payload = json.loads(path.read_text(encoding="utf-8"))
    by_level: dict[int, dict[str, dict]] = {}
    for band in payload.get("bands", []):
        by_level.setdefault(band["band"], {})[band["faction"]] = band
    return [_one_band(path, level, by_faction) for level, by_faction in sorted(by_level.items())]


def build_bis(root: Path, build: str) -> dict[str, list[AddonBisBand]]:
    """Every spec's leveling BiS table the build directory has, keyed by
    spec slug (a spec with no bis/<spec>.json simply has no key -- see the
    module docstring's second point).

    Empty when data/builds/<build>/bis does not exist at all."""
    bis_dir = root / build / "bis"
    if not bis_dir.exists():
        return {}
    return {path.stem: _spec_bis(path) for path in sorted(bis_dir.glob("*.json"))}
