"""Forever's own loot facts, laid over the two databases' Era ones.

Parity contract 6.1 as corrected by 10.4, and the design's risk list: the
fork's item database covers Classic Era, and Forever re-itemises. Only
2,811 of the fork's 7,553 item ids exist in the 1.60 client at all, so a
great deal of what a Forever player will actually loot has no source in
either database. This is where a sourced statement about that goes -- and
where the phase a raid opens in goes, since neither database states a date.

Every file is a `LootOverlay`: curated `sources` and `notes` like every
other hand-maintained fact here, plus `add`, `replace` and `remove`.
Files apply in filename order; within a file, add, then replace, then
remove. A stale instruction -- adding a source that exists, patching or
removing one that does not -- is an error, because the alternative is an
overlay that quietly stops doing anything.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.curated import parse_sources
from pipeline.loot.sources import KIND_ORDER
from pipeline.models import LootFile, LootOverlay
from pipeline.normalize.classes import slugify

#: The phase an undated raid source is gated to: no raid is open at launch.
RAID_DEFAULT_OPENS = "later"

#: The six Era world bosses, keyed by the classic wow npc id and naming
#: each one exactly the way the pinned cmangos/classic-db dump's own
#: `creature_template` row does (verified 2026-09-30 against the
#: committed `builds/1.60.1.70009/raw/classicdb/sources.json` cache --
#: every `quest_reward`/creature-drop record naming that npc_id carries
#: this exact `name`; the engine fork's own `assets/database/db.json`
#: `npcs` table separately confirms Lord Kazzak and Azuregos, but has no
#: row at all for Emeriss, Lethon, Taerar or Ysondre, so classic-db is
#: the only primary source for those four). None of the six carried an
#: `opens` before this dict existed, so the leveling ranker read
#: Amberseal Keeper (off Lord Kazzak) and Blazefury Medallion (off
#: Azuregos) as launch-obtainable -- data-followups-2 lane, 2026-09-30.
#: `sim/cmd/leveling-bis` hand-lists the same six npc ids as its own
#: fallback for a build that carries no curated `opens` here; this dict
#: is what makes that fallback stop being load-bearing, since it always
#: wins whenever `loot.json` states an `opens` for the source it names.
WORLD_BOSS_NPC_NAMES: dict[int, str] = {
    12397: "Lord Kazzak",
    6109: "Azuregos",
    14889: "Emeriss",
    14888: "Lethon",
    14890: "Taerar",
    14887: "Ysondre",
}

#: `world:<slug>` for each `WORLD_BOSS_NPC_NAMES` entry, slugified the
#: same way `pipeline.loot.sources`/`pipeline.loot.classicdb` mint a
#: `world:` source id from an npc's name in the first place -- so a name
#: either database changes surfaces as a source id this set no longer
#: matches, rather than silently gating nothing.
WORLD_BOSS_SOURCE_IDS = {f"world:{slugify(name)}" for name in WORLD_BOSS_NPC_NAMES.values()}

#: The phase these six default to. Not `RAID_DEFAULT_OPENS` -- that
#: sentinel means "no raid names a date at all" -- but the phase the
#: FIRST raid tier itself opens in: `data/curated/phases.json` names
#: this boundary "raids-1", and `curated/loot/forever-raid-phases.json`
#: already patches the one raid the generator emits for that tier
#: (`raid:onyxias-lair`) to the identical value. A world boss server-
#: first happening before a raid tier is live is not a thing the real
#: game allows either, so these six share the raid's own phase rather
#: than a separately invented one.
WORLD_BOSS_DEFAULT_OPENS = "raids-1"


class OverlayError(SystemExit):
    """A curated loot overlay names something the generated file does not."""


def load_overlays(overlay_dir: Path) -> list[tuple[Path, LootOverlay]]:
    """Every overlay in filename order, with its provenance validated.

    A missing directory is not an error: a build with nothing to say about
    its loot is a normal state, and the day one exists the file appears.
    """
    if not overlay_dir.is_dir():
        return []
    loaded: list[tuple[Path, LootOverlay]] = []
    for path in sorted(overlay_dir.glob("*.json")):
        document = LootOverlay(**json.loads(path.read_text(encoding="utf-8")))
        parse_sources([source.model_dump() for source in document.sources], str(path))
        if not document.notes.strip():
            raise OverlayError(f"{path} has no notes; say why the overlay exists")
        loaded.append((path, document))
    return loaded


def apply_overlays(
    document: LootFile, overlays: list[tuple[Path, LootOverlay]]
) -> LootFile:
    by_id = {source.id: source for source in document.sources}
    for path, overlay in overlays:
        for source in overlay.add:
            if source.id in by_id:
                raise OverlayError(
                    f"{path} adds source {source.id!r}, which the generated file "
                    f"already has; use `replace` to change it"
                )
            by_id[source.id] = source
        for patch in overlay.replace:
            if patch.id not in by_id:
                raise OverlayError(
                    f"{path} replaces keys on source {patch.id!r}, which no source "
                    f"has; use `add`, or delete the stale instruction"
                )
            existing = by_id[patch.id]
            changes = patch.model_dump(exclude_unset=True)
            changes.pop("id")
            kind = changes.pop("kind", existing.kind)
            if kind != existing.kind:
                raise OverlayError(
                    f"{path} would change source {patch.id!r} from kind "
                    f"{existing.kind!r} to {kind!r}; a source's kind is part of "
                    f"its id and of every `drop:` origin that names it"
                )
            # `model_copy(update=...)` assigns verbatim, without
            # re-validating -- a patch that touches a nested-model field
            # like `bosses` would then store plain dicts where `LootBoss`
            # instances belong. Merging through `model_validate` instead
            # re-parses the whole source, so a patched `bosses` list comes
            # back out as real `LootBoss` instances, not raw dicts.
            by_id[patch.id] = type(existing).model_validate(
                {**existing.model_dump(), **changes}
            )
        for source_id in overlay.remove:
            if source_id not in by_id:
                raise OverlayError(
                    f"{path} removes source {source_id!r}, which is not there"
                )
            del by_id[source_id]
    # Nothing raids at launch (curated/loot/forever-raid-phases.json's own notes):
    # a raid the generator emits that no overlay dates -- Forever's own Scarlet
    # Enclave surfaced from a fresh Map.csv on 2026-09-30 with no `opens` and would
    # have read as open from launch -- defaults to "later" rather than nothing.
    for source_id, source in list(by_id.items()):
        if source.kind == "raid" and source.opens is None:
            by_id[source_id] = source.model_copy(update={"opens": RAID_DEFAULT_OPENS})
    # The six named world bosses (WORLD_BOSS_NPC_NAMES's own doc): same
    # default-fill shape as the raid loop above, and run after every
    # overlay's own `replace` already applied, so a curated fact about
    # one of these six (a world boss killed server-first before the raid
    # tier, say) always wins over this default rather than being
    # clobbered back to it.
    for source_id in WORLD_BOSS_SOURCE_IDS:
        source = by_id.get(source_id)
        if source is not None and source.kind == "world" and source.opens is None:
            by_id[source_id] = source.model_copy(update={"opens": WORLD_BOSS_DEFAULT_OPENS})
    return LootFile(
        sources=sorted(
            by_id.values(), key=lambda source: (KIND_ORDER.index(source.kind), source.id)
        ),
        # An overlay curates `sources` only; the quests map and the faction
        # map are generated facts about items, not about sources, and no
        # overlay operation here touches either.
        quests=document.quests,
        factions=document.factions,
    )
