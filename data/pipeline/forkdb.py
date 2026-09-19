"""The engine fork's own item database, read for the facts DB2 cannot state.

`assets/database/db.json` in the wowsims-forever checkout is the fork's
`UIDatabase`: items with their AtlasLoot and Wowhead sources, the enchants
the fork's own UI offers, random suffixes, the instance zones, the NPCs it
names, and the reputation factions. Nothing in a DB2 export says which boss
drops an item, which profession crafts it, or what "of the Bear" is worth,
so for `loot.json`, `enchants.json` and `suffixes.json` this file is the
only source there is.

It is read from a checkout at a path and never vendored here: it moves with
the engine pin, the same way `pipeline/simproto`'s bindings do. `make loot`
passes the local checkout; the data workflow clones the fork at the sha in
`sim/enginever/version.go` and passes that.

Measured on the pinned fork, 2026-09-19: 7,553 items (4,481 with sources),
1,168 random suffixes, 173 enchant rows over 150 distinct effect ids, 27
instance zones, 274 named NPCs, 18 factions.

Every enum below is transcribed from the fork's `proto/common.proto` and
`proto/ui.proto`. A value that is not in one of them raises rather than
being guessed at or dropped -- the same policy as
`pipeline/normalize/gear.py`'s `STAT_BY_MODIFIER_ID`.
"""

from __future__ import annotations

import json
from dataclasses import dataclass
from pathlib import Path

#: Where the database sits inside an engine checkout.
DB_RELATIVE = Path("assets/database/db.json")


class ForkDbError(SystemExit):
    """The fork's database is missing, or says something we cannot decode."""


#: proto/common.proto's `Profession`. 6 and 7 are unused in the fork's enum.
PROFESSIONS: dict[int, str] = {
    1: "alchemy",
    2: "blacksmithing",
    3: "enchanting",
    4: "engineering",
    5: "herbalism",
    8: "leatherworking",
    9: "mining",
    10: "skinning",
    11: "tailoring",
}

#: proto/ui.proto's `RepLevel`, minus `RepLevelUnknown`.
REP_LEVELS: dict[int, str] = {
    1: "hated",
    2: "hostile",
    3: "unfriendly",
    4: "neutral",
    5: "friendly",
    6: "honored",
    7: "revered",
    8: "exalted",
}

#: proto/common.proto's `ItemType` -> the planner's slot names, the same
#: vocabulary `pipeline/normalize/gear.py`'s SLOT_BY_INVENTORY_TYPE emits.
#: `ItemTypeWeapon` covers both hands because an enchant that can go on a
#: weapon can go on either.
ITEM_TYPE_SLOTS: dict[int, tuple[str, ...]] = {
    1: ("head",),
    2: ("neck",),
    3: ("shoulder",),
    4: ("back",),
    5: ("chest",),
    6: ("wrist",),
    7: ("hands",),
    8: ("waist",),
    9: ("legs",),
    10: ("feet",),
    11: ("finger",),
    12: ("trinket",),
    13: ("main_hand", "off_hand"),
    14: ("ranged",),
}

#: proto/common.proto's `EnchantType`: what shape of item the enchant needs,
#: which is a different question from which slot it goes in.
ENCHANT_TYPES: dict[int, str] = {
    0: "normal",
    1: "two_hand",
    2: "shield",
    3: "kit",
    4: "staff",
}

#: proto/common.proto's `Class` -> the site's class slugs (the values
#: `pipeline/normalize/classes.py`'s slugify produces from ChrClasses).
CLASS_SLUGS: dict[int, str] = {
    1: "druid",
    2: "hunter",
    3: "mage",
    4: "paladin",
    5: "priest",
    6: "rogue",
    7: "shaman",
    8: "warlock",
    9: "warrior",
}

#: proto/ui.proto's `UIItem.FactionRestriction`, minus its unspecified
#: zero. The client cannot supply this on build 1.60.1.69893 -- 19,066 of
#: its 19,171 ItemSparse rows have AllowableRace -1/-1, the 819 the fork
#: marks restricted among them -- so the fork is the only source.
FACTION_RESTRICTIONS: dict[int, str] = {1: "alliance_only", 2: "horde_only"}


def decode[T](table: dict[int, T], value: int, what: str) -> T:
    """`table[value]`, or a clear error naming the enum and the number.

    Widening a filter to make an unknown value pass is how a wrong fact
    ships quietly; this is the one place that refuses instead.
    """
    if value not in table:
        raise ForkDbError(
            f"the fork database has {what} {value}, which pipeline/forkdb.py does "
            f"not decode; add it from the fork's protos (known: {sorted(table)})"
        )
    return table[value]


@dataclass(frozen=True)
class ForkDatabase:
    """The five tables of db.json anything here reads, plus the two icon maps."""

    items: tuple[dict, ...]
    enchants: tuple[dict, ...]
    random_suffixes: tuple[dict, ...]
    zones: dict[int, str]
    npcs: dict[int, str]
    factions: dict[int, str]
    item_icons: dict[int, str]
    spell_icons: dict[int, str]
    #: The raw icon rows, kept alongside the id maps because `simbuffs`
    #: joins on the name in them and `enchants` joins on the id.
    spell_icon_rows: tuple[dict, ...]
    item_icon_rows: tuple[dict, ...]


def _named(rows: list[dict]) -> dict[int, str]:
    return {int(row["id"]): row["name"] for row in rows}


def _icons(rows: list[dict]) -> dict[int, str]:
    return {int(row["id"]): row["icon"] for row in rows if row.get("icon")}


def load_fork_database(engine_dir: Path) -> ForkDatabase:
    path = engine_dir / DB_RELATIVE
    if not path.exists():
        raise ForkDbError(
            f"no fork item database at {path}; point --engine at a "
            f"wowsims-forever checkout"
        )
    raw = json.loads(path.read_text(encoding="utf-8"))
    missing = sorted({"items", "enchants", "randomSuffixes", "zones", "npcs"} - set(raw))
    if missing:
        raise ForkDbError(f"{path} has no {', '.join(missing)}; is this the fork's db.json?")
    return ForkDatabase(
        items=tuple(raw["items"]),
        enchants=tuple(raw["enchants"]),
        random_suffixes=tuple(raw["randomSuffixes"]),
        zones=_named(raw["zones"]),
        npcs=_named(raw["npcs"]),
        factions=_named(raw.get("factions", [])),
        item_icons=_icons(raw.get("itemIcons", [])),
        spell_icons=_icons(raw.get("spellIcons", [])),
        spell_icon_rows=tuple(raw.get("spellIcons", [])),
        item_icon_rows=tuple(raw.get("itemIcons", [])),
    )
