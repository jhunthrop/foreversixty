"""Each loot.json boss row's creature rank and rare-spawn flag (task L1.7).

cmangos/classic-db is the only source this pipeline has for whether a
creature is a rare spawn: the fork database names its NPCs without a rank,
and so does wowhead's scrape. `pipeline.classic_sources` keeps the dump's
`creature_template.Rank` (`npc_ranks` in the committed
raw/classicdb/sources.json) and the npcs it finds to be rare spawns
(`rare_spawns`, `parse_rare_spawns`: rare-ranked AND pooled). This step
copies both onto every boss row by `npc_id`, whichever origin made the row
(fork, classic-db or wowhead) -- the npc is the same creature either way. A
normal or unknown npc keeps no `rank` key, and only a rare spawn's rows
carry `rare_spawn`.

A `world` source (one named open-world creature's drops) carries no npc id:
it is keyed by the creature's name (`world:<slug>`). It is flagged
`rare_spawn` when every classic-db creature of that name is a rare spawn
(`rare_spawn_names`) -- Swiftmane, Lupos -- never when a name is shared with a
creature that is not one.
"""

from __future__ import annotations

from collections import defaultdict

from pipeline.classic_sources import ClassicDbSourceRecord
from pipeline.models import LootFile

#: The classic-db record kinds that name a creature dropping the item.
CREATURE_KINDS = frozenset({"creature_drop", "skinning", "pickpocketing"})


def rare_spawn_names(
    classic_sources: dict[int, list[ClassicDbSourceRecord]], rare_spawns: set[int]
) -> set[str]:
    """The creature names whose every classic-db npc is a rare spawn."""
    npcs_by_name: dict[str, set[int]] = defaultdict(set)
    for records in classic_sources.values():
        for record in records:
            if record.kind in CREATURE_KINDS and record.npc_id and record.name:
                npcs_by_name[record.name].add(record.npc_id)
    return {name for name, npcs in npcs_by_name.items() if npcs <= rare_spawns}


def apply_npc_ranks(
    document: LootFile,
    npc_ranks: dict[int, int],
    rare_spawns: set[int] | None = None,
    rare_names: set[str] | None = None,
) -> LootFile:
    """`document` with every boss row's `rank` set from `npc_ranks` and its
    `rare_spawn` flag from `rare_spawns`, and every `world` source named in
    `rare_names` flagged `rare_spawn`."""
    rare_spawns = rare_spawns or set()
    rare_names = rare_names or set()
    sources = []
    for source in document.sources:
        if source.kind == "world" and source.name in rare_names:
            source = source.model_copy(update={"rare_spawn": True})
        if source.bosses:
            source = source.model_copy(
                update={
                    "bosses": [
                        boss.model_copy(
                            update={
                                "rank": npc_ranks.get(boss.npc_id),
                                "rare_spawn": True if boss.npc_id in rare_spawns else None,
                            }
                        )
                        for boss in source.bosses
                    ]
                }
            )
        sources.append(source)
    return document.model_copy(update={"sources": sources})
