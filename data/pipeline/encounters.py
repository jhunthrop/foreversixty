"""What each raid boss needs from the raid, for the addon's ready check.

`data/curated/encounter_needs.json` is written by hand: per raid boss,
the raid buffs every player should carry, the debuffs the boss should
wear, and the school the raid stacks resistance against. This module
loads it into `AddonData.encounter_needs` (Data.lua's `encounterNeeds`,
docs/contracts/addon-v2.md contract 1) and refuses to publish an entry
the build cannot vouch for:

* every npc id must be a raid boss in the build's own loot record
  (`loot.json`'s `kind == "raid"` sources). `dungeons.json` was meant to
  be that record, but it is an empty list in every committed build, so
  the raid sources in `loot.json` (which carry each boss's `npc_id`) are
  the pipeline's list of raid bosses;
* every spell id must be a known rank of some class's spell family in
  the build's `spellranks.json` -- any rank, since a raid may carry a
  lower one. Forever changed some classic ids (Curse of the Elements is
  1311680 at its top rank, not classic's 11722) and dropped others
  (there is no Curse of Shadow family), which is exactly what this
  catches;
* `resist` is a magic school name or null;
* `item`, the item a fight asks a player to carry (the Onyxia Scale Cloak for
  Nefarian's Shadow Flame), must be in the build's items.json;
* every `consumables` entry is a kind pipeline.raid_supplies knows (the
  school protection potions), which the addon counts in the bags by kind.

A build with no `loot.json` has no raid record to check against and
ships no encounter needs, the same "absent, not empty" rule
`addondata._default_builds` follows for a build the sim has not run on.
"""

from __future__ import annotations

import json
from pathlib import Path

from pydantic import BaseModel

from pipeline.models import AddonEncounterNeeds, Source
from pipeline.raid_supplies import KINDS

CURATED_FILE = "encounter_needs.json"

#: The schools a raid stacks resistance against.
RESIST_SCHOOLS = frozenset({"fire", "frost", "nature", "shadow", "arcane"})


class EncounterNeedsError(SystemExit):
    """encounter_needs.json names something the build does not know."""


class EncounterNeedsBoss(BaseModel):
    npc_id: int
    #: For the reader of the curated file only; the addon gets the boss
    #: name from the client.
    name: str
    instance: str
    buffs: list[int]
    debuffs: list[int]
    resist: str | None = None
    #: The item the fight asks each player to carry, or None.
    item: int | None = None
    #: The consumable kinds (pipeline.raid_supplies.KINDS) the fight asks for.
    consumables: list[str] = []


class EncounterNeedsFile(BaseModel):
    sources: list[Source]
    notes: str
    bosses: list[EncounterNeedsBoss]


def _raid_boss_ids(path: Path) -> set[int]:
    loot = json.loads(path.read_text(encoding="utf-8"))
    return {
        boss["npc_id"]
        for source in loot.get("sources", [])
        if source.get("kind") == "raid"
        for boss in source.get("bosses", [])
        if boss.get("npc_id") is not None
    }


def _known_spell_ids(path: Path) -> set[int]:
    if not path.exists():
        raise EncounterNeedsError(f"missing {path}")
    spellranks = json.loads(path.read_text(encoding="utf-8"))
    return {
        rank["id"]
        for families in spellranks.get("classes", {}).values()
        for ranks in families.values()
        for rank in ranks
    }


def _item_ids(path: Path) -> set[int]:
    if not path.exists():
        raise EncounterNeedsError(f"missing {path}")
    file = json.loads(path.read_text(encoding="utf-8"))
    rows = file["items"] if isinstance(file, dict) else file
    return {row["id"] for row in rows}


def _check_carried(boss: EncounterNeedsBoss, where: str, items_path: Path) -> None:
    """The item and consumables a boss asks for are ones the build and the
    catalogue know."""
    if boss.item is not None and boss.item not in _item_ids(items_path):
        raise EncounterNeedsError(
            f"{where} names item {boss.item}, which {items_path} does not carry"
        )
    unknown = [kind for kind in boss.consumables if kind not in KINDS]
    if unknown:
        raise EncounterNeedsError(
            f"{where} asks for consumables {unknown}; want kinds from {sorted(KINDS)}"
        )


def load_encounter_needs(
    root: Path, build: str, curated_dir: Path = Path("curated")
) -> dict[int, AddonEncounterNeeds]:
    """The curated needs keyed by npc id, sorted, every id checked against
    `root/build`. Empty when the build has no loot.json or the curated file
    does not exist."""
    loot_path = root / build / "loot.json"
    curated_path = curated_dir / CURATED_FILE
    if not loot_path.exists() or not curated_path.exists():
        return {}
    file = EncounterNeedsFile.model_validate_json(curated_path.read_text(encoding="utf-8"))
    bosses = _raid_boss_ids(loot_path)
    spells = _known_spell_ids(root / build / "spellranks.json")

    needs: dict[int, AddonEncounterNeeds] = {}
    for boss in file.bosses:
        where = f"{curated_path}: npc {boss.npc_id} ({boss.name})"
        if boss.npc_id in needs:
            raise EncounterNeedsError(f"{where} is listed twice")
        if boss.npc_id not in bosses:
            raise EncounterNeedsError(
                f"{where} is not a raid boss in {loot_path}; check the id against the build"
            )
        unknown = [spell for spell in boss.buffs + boss.debuffs if spell not in spells]
        if unknown:
            raise EncounterNeedsError(
                f"{where} names spell ids {unknown} that no class's spell family in "
                f"{root / build / 'spellranks.json'} carries; "
                "Forever may have changed or removed them"
            )
        if boss.resist is not None and boss.resist not in RESIST_SCHOOLS:
            raise EncounterNeedsError(
                f"{where} has resist {boss.resist!r}; want one of {sorted(RESIST_SCHOOLS)} or null"
            )
        _check_carried(boss, where, root / build / "items.json")
        needs[boss.npc_id] = AddonEncounterNeeds(
            buffs=boss.buffs,
            debuffs=boss.debuffs,
            resist=boss.resist,
            item=boss.item,
            consumables=boss.consumables,
        )
    return dict(sorted(needs.items()))
