"""The consumables a raid counts, for the addon's Data.lua.

`data/curated/raid-supplies.json` is written by hand and is the one list the
site and the addon share: `supplies` is the site's raid-prep catalogue (the
flasks, raid elixirs and buff foods a raid leader counts; the API keeps a Go
copy that a guilds test pins), `protection` the school protection potions an
encounter can ask for (encounter_needs.json's `consumables` name their kinds).
This module loads both into `AddonData.consumables` (Data.lua's `consumables`,
item id -> kind, docs/contracts/addon-v2.md contract 1) and refuses an entry
the build cannot vouch for: every id must be a row of the build's
`simconsumes.json` under exactly that name, and every kind one of KINDS.

A build with no simconsumes.json has nothing to check the ids against and
ships no catalogue, the "absent, not empty" rule encounters.py follows.
"""

from __future__ import annotations

import json
from pathlib import Path

from pydantic import BaseModel

CURATED_FILE = "raid-supplies.json"

#: The kinds a consumable counts as: the site's three raid supplies, then the
#: protection potions by school.
SUPPLY_KINDS = ("flask", "elixir", "food")
PROTECTION_KINDS = (
    "fire_protection",
    "frost_protection",
    "nature_protection",
    "shadow_protection",
    "arcane_protection",
)
KINDS = frozenset(SUPPLY_KINDS + PROTECTION_KINDS)


class RaidSuppliesError(SystemExit):
    """raid-supplies.json names something the build does not know."""


class RaidSupply(BaseModel):
    item_id: int
    name: str
    kind: str


class RaidSuppliesFile(BaseModel):
    notes: str
    supplies: list[RaidSupply]
    protection: list[RaidSupply]


def _consumable_names(path: Path) -> dict[int, str]:
    rows = json.loads(path.read_text(encoding="utf-8"))
    return {row["id"]: row["name"] for row in rows}


def _check(entry: RaidSupply, allowed: tuple[str, ...], names: dict[int, str], where: str) -> None:
    if entry.kind not in allowed:
        raise RaidSuppliesError(
            f"{where}: item {entry.item_id} has kind {entry.kind!r}; want one of {allowed}"
        )
    if names.get(entry.item_id) != entry.name:
        raise RaidSuppliesError(
            f"{where}: item {entry.item_id} {entry.name!r} is {names.get(entry.item_id)!r} in the "
            "build's simconsumes.json; check the id and the name against the build"
        )


def load_consumables(root: Path, build: str, curated_dir: Path = Path("curated")) -> dict[int, str]:
    """Item id -> kind, ids ascending, every entry checked against
    `root/build/simconsumes.json`. Empty when the build has no simconsumes.json
    or the curated file does not exist."""
    consumes_path = root / build / "simconsumes.json"
    curated_path = curated_dir / CURATED_FILE
    if not consumes_path.exists() or not curated_path.exists():
        return {}
    file = RaidSuppliesFile.model_validate_json(curated_path.read_text(encoding="utf-8"))
    names = _consumable_names(consumes_path)
    out: dict[int, str] = {}
    for section, allowed in ((file.supplies, SUPPLY_KINDS), (file.protection, PROTECTION_KINDS)):
        for entry in section:
            _check(entry, allowed, names, str(curated_path))
            if entry.item_id in out:
                raise RaidSuppliesError(f"{curated_path}: item {entry.item_id} is listed twice")
            out[entry.item_id] = entry.kind
    return dict(sorted(out.items()))
