"""The loot classification's PvP ranks read the hotfix ItemSparse too.

Accuracy lane 2026-11-09: Forever's re-itemised honor rewards (Marshal's
Plate Legguards 231540 and 499 more) exist only as hotfix rows, and those rows
carry `RequiredPVPRank` (231540: 16, in-game rank 12). The classification read
the shipped `raw/ItemSparse.csv` alone, so every one of them reached
`unlock_title_only_ranks` as "title only" and was published as an unranked
"PvP reward", although the client ranks it.
"""

from __future__ import annotations

import csv
from pathlib import Path

from pipeline.loot.classification import classification_from_raw

_COLUMNS = ("ID", "Display_lang", "RequiredPVPRank")


def _write(path: Path, rows: list[dict[str, str]], columns: tuple[str, ...] = _COLUMNS) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", newline="", encoding="utf-8") as handle:
        writer = csv.DictWriter(handle, fieldnames=list(columns))
        writer.writeheader()
        writer.writerows(rows)


def _raw(tmp_path: Path) -> Path:
    raw = tmp_path / "raw"
    _write(raw / "Map.csv", [], ("ID", "InstanceType"))
    _write(
        raw / "ItemSparse.csv",
        [{"ID": "16479", "Display_lang": "Marshal's Plate Legguards", "RequiredPVPRank": "16"}],
    )
    return raw


def test_a_hotfix_only_honor_reward_keeps_the_clients_rank(tmp_path: Path) -> None:
    raw = _raw(tmp_path)
    _write(
        raw / "hotfixes" / "ItemSparse.csv",
        [
            {"ID": "231540", "Display_lang": "Marshal's Plate Legguards", "RequiredPVPRank": "16"},
            {"ID": "237814", "Display_lang": "Sentinel's Legguards", "RequiredPVPRank": "0"},
        ],
    )

    ranks = classification_from_raw(raw, []).pvp_ranks

    assert ranks == {16479: 16, 231540: 16}


def test_no_hotfix_table_reads_the_shipped_ranks_alone(tmp_path: Path) -> None:
    assert classification_from_raw(_raw(tmp_path), []).pvp_ranks == {16479: 16}
