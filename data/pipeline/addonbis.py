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

Four decisions are load-bearing:

* IDS ONLY in the pick. leveling-bis's JSON carries the item's name and a
  prose swap_note; neither goes into the packed pick string. Task L1.4
  added the name back, once per distinct id, in a side table
  (`bisNames`, with `bisDps` for a weapon's own dps -- build_bis_items),
  because the tooltip must name a pick before the client's item cache
  fills; one entry per id rather than one per pick keeps the size cost
  small and leaves the packed string, and so Gear.decodePick, unchanged.
  `source_kind` stays (advanced detail wants it) but travels as one of
  SOURCE_KIND_CODES' single-letter codes in the rendered Lua -- see
  pipeline.addonlua.
* CLIENT-CONFIRMED ONLY (task L1.4). A candidate items.json marks
  `client_unconfirmed` is never a pick or an alternative: the client may
  never resolve its id, and the slot falls silent. The next confirmed
  candidate in leveling-bis's own order takes its place; the skips are
  counted per build (bis_unconfirmed_skips).
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

from pipeline.models import AddonBisAlternative, AddonBisBand, AddonBisItem

#: The `preset` tag of a band entry measured on the bare character (sim/cmd/leveling-bis/preset.go).
BARE_PRESET = "bare"

#: sim/cmd/leveling-bis/main.go's own defaultBandsFlag: 20 through 60 step
#: 10 (owner ruling 2026-09-29). Not the rotation ladder's seven rungs
#: (pipeline.addonrotation.LEVEL_BANDS) -- the two lists are kept
#: separately on purpose and only happen to share 20/30/40/50/60.
BIS_LEVEL_BANDS = list(range(20, 61, 10))

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
    #: `world_drop` (world-drop-pool lane, 2026-09-29): loot.json's own
    #: generic, bind-on-equip, auction-housable world-drop pool bucket
    #: (sim/cmd/leveling-bis/band.go's own sourceKindPriority).
    "world_drop": "B",
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

#: How many runners-up a pick carries as alternatives (contract 1 in
#: docs/contracts/addon-v2.md: "up to two").
MAX_ALTERNATIVES = 2


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


def _new_item_ids(new_at_band: list[str], picks: dict[str, AddonBisItem]) -> list[int]:
    """leveling-bis states `new_at_band` as "<slot>: <item name>" strings
    (report.go's own display format); the slot name is the join key back to
    this band's own packed picks, since Data.lua carries ids only (a slot
    whose candidates were all client-unconfirmed has no pick and is
    skipped; one whose pick was replaced names the replacement)."""
    ids = []
    for entry in new_at_band:
        slot = entry.split(":", 1)[0]
        pick = picks.get(slot)
        if pick is not None:
            ids.append(pick.item_id)
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


def _runners_up(path: Path, where: str, entry: dict) -> list[dict]:
    """The row's own runners-up, best first: leveling-bis's `runners_up`,
    or for a bis file written before it emitted them, its `alternatives`
    list (up to three, best first -- runners_up is cut from the same list).
    `alternatives` entries past the runners-up are kept on the end: they are
    the candidates a client-unconfirmed skip (task L1.4) falls through to."""
    runners_up = entry.get("runners_up")
    alternatives = entry.get("alternatives") or []
    if runners_up is None:
        return list(alternatives)
    if len(runners_up) > MAX_ALTERNATIVES:
        raise AddonBisError(
            f"{path}: {where} has {len(runners_up)} runners_up, at most {MAX_ALTERNATIVES}"
        )
    seen = {runner_up["item_id"] for runner_up in runners_up}
    return list(runners_up) + [alt for alt in alternatives if alt["item_id"] not in seen]


def _pick(
    path: Path, where: str, entry: dict, unconfirmed: set[int], skips: list[int]
) -> AddonBisItem | None:
    """One slot row as the addon's pick. `chance` and `runners_up` are
    optional: a bis file leveling-bis wrote before it emitted them loads
    with no chance, and with the first two of its `alternatives` (if any)
    in place of runners_up.

    Task L1.4: a candidate whose items.json row is `client_unconfirmed` is
    never the pick or an alternative -- the client never resolves such an
    id, so the slot would fall silent. The next confirmed candidate, in
    leveling-bis's own order, takes its place (a promoted runner-up carries
    no `chance`: the row's chance was the original pick's). Every candidate
    passed over is appended to `skips`. None when no candidate is confirmed."""
    chance = entry.get("chance")
    if chance is not None and not 0 < chance <= 1:
        # loot.json states percents (0-100); the contract is a fraction.
        raise AddonBisError(f"{path}: {where} chance {chance!r} is not a fraction in (0, 1]")
    candidates = [dict(entry, chance=chance)] + [
        dict(runner_up, chance=None) for runner_up in _runners_up(path, where, entry)
    ]
    kept: list[dict] = []
    for candidate in candidates:
        if len(kept) > MAX_ALTERNATIVES:
            break
        if candidate["item_id"] in unconfirmed:
            skips.append(candidate["item_id"])
            continue
        kept.append(candidate)
    if not kept:
        return None
    pick, alternatives = kept[0], kept[1:]
    return AddonBisItem(
        item_id=pick["item_id"],
        source_kind=pick["source_kind"],
        source=pick.get("source", ""),
        chance=pick["chance"],
        alternatives=[
            AddonBisAlternative(
                item_id=alternative["item_id"],
                source_kind=alternative["source_kind"],
                source=alternative.get("source", ""),
            )
            for alternative in alternatives
        ],
    )


def _one_band(
    path: Path,
    level: int,
    by_faction: dict[str, dict],
    unconfirmed: set[int],
    skips: list[int],
) -> AddonBisBand:
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
        picks = {
            slot: _pick(path, f"band {level}/{faction}/{slot}", entry, unconfirmed, skips)
            for slot, entry in slot_lookup.items()
        }
        factions[faction] = {slot: pick for slot, pick in picks.items() if pick is not None}
        new_at_band[faction] = _new_item_ids(band.get("new_at_band") or [], factions[faction])
        # The first faction (sorted, so always the same one for a given
        # band) stands for the band as a whole -- see AddonBisBand.weights'
        # own docstring on why factions are not kept apart here.
        if index == 0:
            weights = _significant_weights(band.get("weights") or [])
    return AddonBisBand(level=level, factions=factions, new_at_band=new_at_band, weights=weights)


def _spec_bis(path: Path, unconfirmed: set[int], skips: list[int]) -> list[AddonBisBand]:
    """One spec's bis/<spec>.json, regrouped from leveling-bis's flat list
    of (band, faction) reports into one AddonBisBand per level, both
    factions inline -- the shape the addon actually indexes by."""
    payload = json.loads(path.read_text(encoding="utf-8"))
    by_level: dict[int, dict[str, dict]] = {}
    for band in payload.get("bands", []):
        # Level 60 also carries a "raid" entry per faction (leveling-bis's
        # presets); the addon's table is the bare character's, and a later
        # entry would otherwise overwrite the bare one.
        if band.get("preset", BARE_PRESET) != BARE_PRESET:
            continue
        by_level.setdefault(band["band"], {})[band["faction"]] = band
    return [
        _one_band(path, level, by_faction, unconfirmed, skips)
        for level, by_faction in sorted(by_level.items())
    ]


def _flat_items(root: Path, build: str) -> dict[int, dict]:
    """data/builds/<build>/items.json by id; empty when the build has none."""
    path = root / build / "items.json"
    if not path.exists():
        return {}
    return {row["id"]: row for row in json.loads(path.read_text(encoding="utf-8"))}


def _unconfirmed_ids(root: Path, build: str) -> set[int]:
    """Every id items.json marks `client_unconfirmed` (task L1.4): stats the
    client has not confirmed, an id the client may never resolve at all."""
    return {
        item_id
        for item_id, row in _flat_items(root, build).items()
        if row.get("client_unconfirmed")
    }


def _build_bis(root: Path, build: str) -> tuple[dict[str, list[AddonBisBand]], list[int]]:
    bis_dir = root / build / "bis"
    if not bis_dir.exists():
        return {}, []
    unconfirmed = _unconfirmed_ids(root, build)
    skips: list[int] = []
    bis = {
        path.stem: _spec_bis(path, unconfirmed, skips) for path in sorted(bis_dir.glob("*.json"))
    }
    return bis, skips


def build_bis(root: Path, build: str) -> dict[str, list[AddonBisBand]]:
    """Every spec's leveling BiS table the build directory has, keyed by
    spec slug (a spec with no bis/<spec>.json simply has no key -- see the
    module docstring's second point).

    Empty when data/builds/<build>/bis does not exist at all."""
    return _build_bis(root, build)[0]


def bis_unconfirmed_skips(root: Path, build: str) -> int:
    """How many candidates build_bis passed over because items.json marks
    them `client_unconfirmed`, across every spec, band and faction of the
    build -- the audit line `python -m pipeline addon-data` prints."""
    return len(_build_bis(root, build)[1])


def unconfirmed_skips_line(root: Path, build: str) -> str:
    """The audit line `python -m pipeline addon-data` logs for a build:
    the skip count, and how many distinct items those skips were."""
    skips = _build_bis(root, build)[1]
    return (
        f"bis {build}: skipped {len(skips)} client_unconfirmed candidates "
        f"({len(set(skips))} distinct items)"
    )


#: Data.lua's bisNames values are shown in WoW tooltips, where `|` starts an
#: escape sequence; a name never carries one (task L1.4, ruling 1).
NAME_SEPARATOR_STAND_IN = "/"


def _class_dps(root: Path, build: str) -> dict[int, float]:
    """Item id -> weapon dps from data/builds/<build>/items/<class>.json, the
    rows sim/cmd/leveling-bis/data.go reads `candidate.DPS` from
    (pipeline.normalize.gear computes it, two decimals). Every class file
    states the same dps for an id, so the union is taken; 0 (not a weapon)
    is left out."""
    dps: dict[int, float] = {}
    for path in sorted((root / build / "items").glob("*.json")):
        for row in json.loads(path.read_text(encoding="utf-8"))["items"]:
            if row.get("dps"):
                dps[row["id"]] = float(row["dps"])
    return dps


def build_bis_items(
    root: Path, build: str, bis: dict[str, list[AddonBisBand]]
) -> tuple[dict[int, str], dict[int, float]]:
    """Data.lua's `bisNames` and `bisDps` side tables (task L1.4): every item
    id that is a pick or an alternative in any band, to its items.json
    display name; and, for a weapon, to its own dps at one decimal. The
    tooltip names a pick and prices a weapon pick before the client's item
    cache fills; the packed pick strings stay exactly as they were, so
    Gear.decodePick reads them unchanged."""
    ids = sorted(
        {
            item_id
            for bands in bis.values()
            for band in bands
            for slots in band.factions.values()
            for pick in slots.values()
            for item_id in [pick.item_id] + [alt.item_id for alt in pick.alternatives]
        }
    )
    if not ids:
        return {}, {}
    items = _flat_items(root, build)
    class_dps = _class_dps(root, build)
    names: dict[int, str] = {}
    dps: dict[int, float] = {}
    for item_id in ids:
        row = items.get(item_id)
        if row is None or not row.get("name"):
            raise AddonBisError(f"items.json for {build} has no name for bis item {item_id}")
        names[item_id] = row["name"].replace("|", NAME_SEPARATOR_STAND_IN)
        if item_id in class_dps:
            dps[item_id] = round(class_dps[item_id], 1)
    return names, dps
