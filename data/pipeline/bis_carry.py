"""Carry the active build's BiS ranking onto a newly activated client build.

`tools/activate-build.sh` moves the site to a new client the morning it appears, hours
before the nightly ranking (bis.yml, dispatched with force at the end of the move) has
ranked it. Until then the new build serves the old build's `bis/` files, restamped.

A new client can drop items the old ranking picked: 1.60.1.70685's items.json lost 320
of 70658's items on 2026-10-29, 29 of them carried picks, and `addon-data` stopped the
move on the first (its Data.lua names every pick). The ruling that day: the carry drops
a removed pick with a report line. A dropped pick leaves the slot empty with
`empty_reason: "not_in_client"` (the page's fallback "no known source" line) until the
ranking reruns; a dropped alternative or runner-up simply leaves its list. Nothing is
promoted in its place: a promotion would have to re-apply the ranker's pair and
two-hand rules, and the ranking replaces the whole file within hours anyway.

Every non-JSON file (the `.md` reports) is copied unchanged.
"""

from __future__ import annotations

import json
import shutil
from pathlib import Path
from typing import Any

#: The `empty_reason` a slot whose carried pick the new client lacks is published with.
NOT_IN_CLIENT_REASON = "not_in_client"

#: The slot-row lists that hold other candidates (sim/cmd/leveling-bis/report.go).
CANDIDATE_LISTS = ("alternatives", "runners_up")


class BisCarryError(SystemExit):
    """The carry cannot stand for the new build as asked."""


def _item_ids(root: Path, build: str) -> set[int]:
    path = root / build / "items.json"
    if not path.exists():
        raise BisCarryError(f"{path} is missing: fetch {build} first")
    return {row["id"] for row in json.loads(path.read_text(encoding="utf-8"))}


def _where(band: dict[str, Any], slot: str) -> str:
    return (
        f"{band.get('spec')} band {band.get('band')} {band.get('faction')} "
        f"{band.get('preset', 'bare')} {slot}"
    )


def _carry_slot(
    band: dict[str, Any], row: dict[str, Any], known: set[int], new: str, lines: list[str]
) -> dict[str, Any]:
    where = _where(band, row["slot"])
    pick_lost = "item_id" in row and row["item_id"] not in known
    if pick_lost:
        lines.append(
            f"{where}: pick {row['item_id']} {row.get('item_name', '')} is not in {new}; "
            "slot left empty until the ranking reruns"
        )
    lines.extend(
        f"{where}: alternative {cand['item_id']} {cand.get('item_name', '')} "
        f"is not in {new}; dropped"
        for cand in row.get("alternatives", ())
        if cand["item_id"] not in known
    )
    lists = {
        key: kept
        for key in CANDIDATE_LISTS
        if (kept := [cand for cand in row.get(key, ()) if cand["item_id"] in known])
    }
    if not pick_lost:
        rest = {key: value for key, value in row.items() if key not in CANDIDATE_LISTS}
        return {**rest, **lists}
    return {
        "slot": row["slot"],
        "verified": False,
        **lists,
        "empty_reason": NOT_IN_CLIENT_REASON,
    }


def _carry_file(
    payload: dict[str, Any], old: str, new: str, known: set[int], lines: list[str]
) -> dict[str, Any]:
    if payload.get("build") != old:
        raise BisCarryError(
            f"bis/{payload.get('spec')}.json is stamped {payload.get('build')!r}, not {old}"
        )
    bands = [
        {**band, "slots": [_carry_slot(band, row, known, new, lines) for row in band["slots"]]}
        for band in payload.get("bands", [])
    ]
    return {**payload, "build": new, "bands": bands}


def carry_bis(root: Path, old: str, new: str) -> list[str]:
    """Write builds/<new>/bis/ from builds/<old>/bis/ and return one report line per
    pick or alternative the new build's items.json lacks. Writes nothing when any file
    fails to carry."""
    source, target = root / old / "bis", root / new / "bis"
    if not source.is_dir():
        raise BisCarryError(f"{source} does not exist: nothing to carry")
    if target.exists():
        raise BisCarryError(f"{target} already has a ranking; the carry never overwrites one")
    known = _item_ids(root, new)
    lines: list[str] = []
    carried = {
        path.name: _carry_file(json.loads(path.read_text(encoding="utf-8")), old, new, known, lines)
        for path in sorted(source.glob("*.json"))
    }
    target.mkdir()
    for name, payload in carried.items():
        text = json.dumps(payload, indent=2, ensure_ascii=False) + "\n"
        (target / name).write_text(text, encoding="utf-8")
    for path in sorted(source.iterdir()):
        if path.suffix != ".json" and path.is_file():
            shutil.copy2(path, target / path.name)
    return lines
