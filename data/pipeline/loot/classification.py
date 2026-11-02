"""The two client facts `loot.json` needs from raw tables, kept in one file
both loot generators read.

`loot` (data.yml) has the client's `Map.csv` and `ItemSparse.csv`;
`loot-merge` (bis.yml's nightly) does not. Each used to work out instance
types and PvP ranks its own way: `loot` from the client tables, the merge
by reading them back out of the committed loot.json. The two answers
differed (Alterac Valley and Arathi Basin are battlegrounds in Map.csv and
absent from loot.json; a classic honor reward re-itemised under a new id
lost its rank on the way back), so loot.json changed on every pair of runs
(round 8, owner ruling 3).

Now `loot` writes what it read to `raw/client/loot-classification.json`
and both commands build from that file through the same code, so they
write the same loot.json from the same inputs. The file holds the client's
facts unchanged (battlegrounds included); what loot does with a
battleground is `loot_instance_types`, one rule both commands apply.
"""

from __future__ import annotations

import json
from dataclasses import dataclass
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.loot import constants
from pipeline.loot.sources import instance_types, pvp_ranks

#: Under `raw/` so the build manifest leaves it out, like the other raw-free
#: input caches committed there (`raw/classicdb/sources.json`).
CLASSIFICATION = Path("raw/client/loot-classification.json")


@dataclass(frozen=True)
class LootClassification:
    #: AreaTable zone id -> the client's `Map.InstanceType` (1 dungeon, 2 raid,
    #: 3 battleground, ...), through `pipeline.loot.sources.instance_types`.
    instance_types: dict[int, int]
    #: Item id -> the client's `ItemSparse.RequiredPVPRank`.
    pvp_ranks: dict[int, int]


def classification_from_raw(raw: Path, zone_rows: list[dict]) -> LootClassification:
    return LootClassification(
        instance_types=instance_types(read_csv(raw / "Map.csv"), zone_rows),
        pvp_ranks=pvp_ranks(read_csv(raw / "ItemSparse.csv")),
    )


def write_classification(build_dir: Path, classification: LootClassification) -> Path:
    path = build_dir / CLASSIFICATION
    path.parent.mkdir(parents=True, exist_ok=True)
    document = {
        "instance_types": {str(k): v for k, v in sorted(classification.instance_types.items())},
        "pvp_ranks": {str(k): v for k, v in sorted(classification.pvp_ranks.items())},
    }
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


def load_classification(build_dir: Path) -> LootClassification:
    path = build_dir / CLASSIFICATION
    if not path.exists():
        raise SystemExit(
            f"no {path}; run `python -m pipeline loot --build {build_dir.name}` once with the "
            "build's raw client tables to write it"
        )
    document = json.loads(path.read_text(encoding="utf-8"))
    return LootClassification(
        instance_types={int(k): int(v) for k, v in document["instance_types"].items()},
        pvp_ranks={int(k): int(v) for k, v in document["pvp_ranks"].items()},
    )


def loot_instance_types(classification: LootClassification) -> dict[int, int]:
    """The instance types loot.json is built with. With
    `constants.BATTLEGROUND_SPAWNS_ARE_SOURCES` a battleground zone is left
    out, so its spawns are open-world sources; without it the zone keeps its
    type and gets no loot source (`classicdb.instanced_zones`)."""
    if not constants.BATTLEGROUND_SPAWNS_ARE_SOURCES:
        return dict(classification.instance_types)
    return {
        zone_id: kind
        for zone_id, kind in classification.instance_types.items()
        if kind != constants.BATTLEGROUND_INSTANCE_TYPE
    }
