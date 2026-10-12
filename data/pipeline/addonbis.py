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
  never resolve its id, and the slot falls silent. Since task L1.6
  leveling-bis drops such items before it ranks, so a regenerated build
  has nothing to skip; this stays as a safety net. The next confirmed
  candidate in leveling-bis's own order that the band's pair and
  two-hand rules allow takes its place, keeping its drop chance; "new at
  <band>" is recomputed from the final picks; the skips are counted per
  build (build_bis_with_skips, unconfirmed_skips_line), and the audit line
  must read `skipped 0` for a regenerated build.
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
import math
from dataclasses import dataclass, field
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

#: sim/cmd/leveling-bis/pick.go's slotOrder: the order a band's slots are
#: decided in, and the order "new at <band>" lists them.
SLOT_ORDER = [
    "head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
    "legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
    "main_hand", "off_hand", "ranged",
]  # fmt: skip

#: sim/cmd/leveling-bis/verify.go's pairSlot: slots that may never hold the
#: same item as each other.
PAIR_SLOT = {
    "finger1": "finger2",
    "finger2": "finger1",
    "trinket1": "trinket2",
    "trinket2": "trinket1",
    "main_hand": "off_hand",
    "off_hand": "main_hand",
}

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


@dataclass
class _ItemFacts:
    """What the packer needs to know about items, read once per build."""

    #: Ids items.json marks `client_unconfirmed`.
    unconfirmed: set[int] = field(default_factory=set)
    #: items.json display names, for the pair rule's same-name check.
    names: dict[int, str] = field(default_factory=dict)
    #: Ids any items/<class>.json row marks `two_hand`.
    two_hand: set[int] = field(default_factory=set)
    #: (item id, source label) -> drop chance as a fraction (_drop_chances).
    chances: dict[tuple[int, str], float] = field(default_factory=dict)


def _flat_items(root: Path, build: str) -> dict[int, dict]:
    """data/builds/<build>/items.json by id; empty when the build has none."""
    path = root / build / "items.json"
    if not path.exists():
        return {}
    return {row["id"]: row for row in json.loads(path.read_text(encoding="utf-8"))}


def _go_round_half_away(value: float) -> float:
    """Go's math.Round for a non-negative value (Python's round() is banker's)."""
    return math.floor(value + 0.5)


def _drop_chances(root: Path, build: str) -> dict[tuple[int, str], float]:
    """(item id, source label) -> drop chance, built the way leveling-bis
    builds its loot index (sim/cmd/leveling-bis/data.go's loadLootIndex: a
    source's own `items` under its `name`, a boss's under "<name>: <boss>",
    each with its `item_chances` percent; the flat `quest` source is
    skipped) and converted the way its dropChance does (report.go: percent
    to a fraction, math.Round(chance*1e4)/1e6; a chance of 0 is none)."""
    path = root / build / "loot.json"
    if not path.exists():
        return {}
    loot = json.loads(path.read_text(encoding="utf-8"))
    chances: dict[tuple[int, str], float] = {}

    def add(item_id: int, label: str, percent: float | None) -> None:
        if not percent or percent <= 0:
            return
        fraction = _go_round_half_away(percent * 1e4) / 1e6
        key = (item_id, label)
        chances[key] = max(chances.get(key, 0.0), fraction)

    for source in loot.get("sources", []):
        if source.get("kind") == "quest":
            continue
        name = source.get("name", "")
        item_chances = source.get("item_chances") or {}
        for item_id in source.get("items") or []:
            add(item_id, name, item_chances.get(str(item_id)))
        for boss in source.get("bosses") or []:
            label = f"{name}: {boss['name']}" if boss.get("name") else name
            boss_chances = boss.get("item_chances") or {}
            for item_id in boss.get("items") or []:
                add(item_id, label, boss_chances.get(str(item_id)))
    return chances


def _item_facts(root: Path, build: str) -> _ItemFacts:
    flat = _flat_items(root, build)
    two_hand: set[int] = set()
    for path in sorted((root / build / "items").glob("*.json")):
        for row in json.loads(path.read_text(encoding="utf-8"))["items"]:
            if row.get("two_hand"):
                two_hand.add(row["id"])
    return _ItemFacts(
        unconfirmed={item_id for item_id, row in flat.items() if row.get("client_unconfirmed")},
        names={item_id: row["name"] for item_id, row in flat.items() if row.get("name")},
        two_hand=two_hand,
        chances=_drop_chances(root, build),
    )


def _candidates(path: Path, where: str, entry: dict) -> list[dict]:
    """The row's pick followed by its runners-up, best first, each carrying
    `chance` (the row's own for the pick; None for a runner-up until it is
    promoted). `chance` and `runners_up` are optional: a bis file
    leveling-bis wrote before it emitted them loads with no chance, and with
    its `alternatives` in place of runners_up."""
    chance = entry.get("chance")
    if chance is not None and not 0 < chance <= 1:
        # loot.json states percents (0-100); the contract is a fraction.
        raise AddonBisError(f"{path}: {where} chance {chance!r} is not a fraction in (0, 1]")
    return [dict(entry, chance=chance)] + [
        dict(runner_up, chance=None) for runner_up in _runners_up(path, where, entry)
    ]


def _same_item(facts: _ItemFacts, a: int, b: int) -> bool:
    """verify.go's wouldDuplicatePairMate test: the same id, or the same name."""
    if a == b:
        return True
    name = facts.names.get(a)
    return name is not None and name == facts.names.get(b)


def _usable(facts: _ItemFacts, final: dict[str, dict], slot: str, item_id: int) -> bool:
    """sim/cmd/leveling-bis/confirmed.go's usableAlternative, applied to the
    band's final picks: never the pair mate's own item (rings, trinkets,
    and the two hands -- verify.go's pairSlot), never an off hand under a
    two-handed main hand, and (fix round 1, item 1) never a two-hander
    promoted into the main hand while the off hand has a pick."""
    mate = PAIR_SLOT.get(slot)
    if mate in final and _same_item(facts, final[mate]["item_id"], item_id):
        return False
    if slot == "off_hand" and final.get("main_hand", {}).get("item_id") in facts.two_hand:
        return False
    if slot == "main_hand" and item_id in facts.two_hand and "off_hand" in final:
        return False
    return True


def _without(final: dict[str, dict], slot: str) -> dict[str, dict]:
    """The band's final picks with one slot left out."""
    return {key: row for key, row in final.items() if key != slot}


def _faction_picks(
    path: Path,
    where: str,
    slot_lookup: dict[str, dict],
    facts: _ItemFacts,
    skips: list[int],
) -> dict[str, AddonBisItem]:
    """One (band, faction)'s packed picks, decided for the band as a whole.

    Task L1.4: a candidate items.json marks `client_unconfirmed` is never a
    pick or an alternative -- the client may never resolve its id, and the
    slot would fall silent. A slot whose own pick is confirmed keeps it.
    Every other slot, in leveling-bis's slot order (pick.go's slotOrder),
    promotes its first confirmed candidate that _usable allows beside the
    picks already final, so a promotion never breaks the pair or two-hand
    rules leveling-bis itself enforces -- except that when the only
    confirmed main hand left is a two-hander, it is adopted and the off hand
    empties, as leveling-bis's own confirmed.go does (task L1.6). A promoted
    pick carries its drop chance at its own source when loot.json states one
    (_drop_chances), and no chance otherwise. The
    alternatives are the next two confirmed candidates. Every unconfirmed
    candidate passed over on the way is appended to `skips`."""
    candidates = {
        slot: _candidates(path, f"{where}/{slot}", entry) for slot, entry in slot_lookup.items()
    }
    final: dict[str, dict] = {
        slot: rows[0]
        for slot, rows in candidates.items()
        if rows[0]["item_id"] not in facts.unconfirmed
    }
    for slot in SLOT_ORDER:
        if slot not in candidates or slot in final:
            continue
        confirmed = [row for row in candidates[slot] if row["item_id"] not in facts.unconfirmed]
        promoted = next(
            (row for row in confirmed if _usable(facts, final, slot, row["item_id"])), None
        )
        if promoted is None and slot == "main_hand":
            # Task L1.6, ruling 4: when the only confirmed main hand left is a
            # two-hander, adopt it and empty the off hand, exactly as
            # confirmed.go's preferConfirmedStats does (lines 91-93).
            promoted = next(
                (
                    row
                    for row in confirmed
                    if row["item_id"] in facts.two_hand
                    and _usable(facts, _without(final, "off_hand"), slot, row["item_id"])
                ),
                None,
            )
            if promoted is not None:
                final.pop("off_hand", None)
        if promoted is not None:
            chance = facts.chances.get((promoted["item_id"], promoted.get("source", "")))
            final[slot] = dict(promoted, chance=chance)
    picks: dict[str, AddonBisItem] = {}
    for slot in SLOT_ORDER:
        rows = candidates.get(slot)
        if rows is None:
            continue
        pick = final.get(slot)
        confirmed = [row for row in rows if row["item_id"] not in facts.unconfirmed]
        alternatives = [
            row for row in confirmed if pick is None or row["item_id"] != pick["item_id"]
        ][:MAX_ALTERNATIVES]
        used = ([pick] if pick is not None else []) + alternatives
        if len(used) <= MAX_ALTERNATIVES:
            # Fewer than a pick and two alternatives: every candidate was looked at.
            considered = rows
        else:
            used_ids = {row["item_id"] for row in used}
            last = max(i for i, row in enumerate(rows) if row["item_id"] in used_ids)
            considered = rows[: last + 1]
        skips.extend(row["item_id"] for row in considered if row["item_id"] in facts.unconfirmed)
        if pick is None:
            continue
        picks[slot] = AddonBisItem(
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
    return picks


def _new_item_ids(picks: dict[str, AddonBisItem], previous: dict[str, int] | None) -> list[int]:
    """The ids newly best at this band (fix round 1, item 2): a slot whose
    final pick differs from the previous band's final pick for the same
    spec, faction and slot, in slot order -- leveling-bis's own rule
    (report.go), applied to what this module actually packs rather than to
    leveling-bis's `new_at_band` strings, which name a pick a promotion may
    have replaced. With no previous band every pick is new."""
    before = previous or {}
    return [
        picks[slot].item_id
        for slot in SLOT_ORDER
        if slot in picks and before.get(slot) != picks[slot].item_id
    ]


def _one_band(
    path: Path,
    level: int,
    by_faction: dict[str, dict],
    facts: _ItemFacts,
    skips: list[int],
    previous: dict[str, dict[str, int]],
) -> AddonBisBand:
    """One level, both factions. `previous` is faction -> slot -> the
    previous band's final pick id; it is updated to this band's on return."""
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
        picks = _faction_picks(path, f"band {level}/{faction}", slot_lookup, facts, skips)
        factions[faction] = picks
        new_at_band[faction] = _new_item_ids(picks, previous.get(faction))
        previous[faction] = {slot: pick.item_id for slot, pick in picks.items()}
        # The first faction (sorted, so always the same one for a given
        # band) stands for the band as a whole -- see AddonBisBand.weights'
        # own docstring on why factions are not kept apart here.
        if index == 0:
            weights = _significant_weights(band.get("weights") or [])
    return AddonBisBand(level=level, factions=factions, new_at_band=new_at_band, weights=weights)


def _spec_bis(path: Path, facts: _ItemFacts, skips: list[int]) -> list[AddonBisBand]:
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
    previous: dict[str, dict[str, int]] = {}
    return [
        _one_band(path, level, by_faction, facts, skips, previous)
        for level, by_faction in sorted(by_level.items())
    ]


def build_bis_with_skips(root: Path, build: str) -> tuple[dict[str, list[AddonBisBand]], list[int]]:
    """build_bis, plus the id of every candidate it passed over because
    items.json marks it `client_unconfirmed` (one entry per skip, across
    every spec, band and faction) -- unconfirmed_skips_line's input."""
    bis_dir = root / build / "bis"
    if not bis_dir.exists():
        return {}, []
    facts = _item_facts(root, build)
    skips: list[int] = []
    bis = {path.stem: _spec_bis(path, facts, skips) for path in sorted(bis_dir.glob("*.json"))}
    return bis, skips


def build_bis(root: Path, build: str) -> dict[str, list[AddonBisBand]]:
    """Every spec's leveling BiS table the build directory has, keyed by
    spec slug (a spec with no bis/<spec>.json simply has no key -- see the
    module docstring's second point).

    Empty when data/builds/<build>/bis does not exist at all."""
    return build_bis_with_skips(root, build)[0]


def unconfirmed_skips_line(build: str, skips: list[int]) -> str:
    """The audit line `python -m pipeline addon-data` logs for a build:
    the skip count, and how many distinct items those skips were."""
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
