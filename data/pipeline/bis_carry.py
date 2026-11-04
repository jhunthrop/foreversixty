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

A new client can also reshape a talent tree: 1.60.1.70755 added two Fury talents and
removed one, so the carried bands' talent strings (one digit per talent, in the tree's tab
order) read their ranks against the wrong talents and the simulator refused the builds as
unreachable. The carry re-keys each band's ranks onto the new trees by talent name (the
guide builds' own rule, web/scripts/reencode-guide-builds.mjs); a ranked talent the new
tree lacks, or a rank past the talent's new maximum, is dropped with a report line and
`talent_points` counts what was kept. A tree reshaped that far can leave a build a gate
short (70755's Flurry went from five ranks to three under Bloodthirst); the ranking
re-spends it.

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


#: Each tree's talents, name to max rank, in tab order (the order a band's digits follow).
Trees = list[dict[str, int]]


def _trees(root: Path, build: str, class_slug: str) -> Trees:
    path = root / build / "talents" / f"{class_slug}.json"
    if not path.exists():
        raise BisCarryError(f"{path} is missing: the carry re-keys talents by name")
    payload = json.loads(path.read_text(encoding="utf-8"))
    return [
        {talent["name"]: talent["max_rank"] for talent in tree["talents"]}
        for tree in payload["trees"]
    ]


def _carry_tree(
    where: str, tree: int, string: str, old: dict[str, int], new: dict[str, int], lines: list[str]
) -> str:
    """One tree's digits re-keyed by name; a lost talent or a rank past the new maximum is
    dropped with a report line."""
    if len(string) > len(old):
        raise BisCarryError(f"{where}: tree {tree} has {len(string)} ranks for {len(old)} talents")
    # A string may stop at its last ranked talent; the talents past it hold no points.
    ranks = {name: int(digit, 36) for name, digit in zip(old, string, strict=False)}
    for name, rank in ranks.items():
        if rank > 0 and name not in new:
            lines.append(f"{where}: {name} {rank} is not in tree {tree}; dropped")
        elif rank > new.get(name, rank):
            lines.append(f"{where}: {name} {rank} is past its new maximum {new[name]}; capped")
    return "".join(str(min(ranks.get(name, 0), most)) for name, most in new.items())


def _carry_talents(
    band: dict[str, Any], old: Trees, new: Trees, lines: list[str]
) -> dict[str, Any]:
    """The band with its talent string re-keyed by name from `old`'s trees onto `new`'s."""
    if old == new:
        return band
    where = _where(band, "talents")
    strings = band["talents"].split("-")
    if len(strings) != len(old) or len(new) != len(old):
        raise BisCarryError(f"{where}: {len(strings)} trees, not {len(old)}")
    carried = [
        _carry_tree(where, tree, string, old_tree, new_tree, lines)
        for tree, (string, old_tree, new_tree) in enumerate(zip(strings, old, new, strict=True))
    ]
    points = sum(int(digit) for string in carried for digit in string)
    return {**band, "talents": "-".join(carried), "talent_points": points}


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


def _carry_band(
    root: Path, band: dict[str, Any], old: str, new: str, known: set[int], lines: list[str]
) -> dict[str, Any]:
    slots = [_carry_slot(band, row, known, new, lines) for row in band["slots"]]
    carried = {**band, "slots": slots}
    if "talents" not in band:
        return carried
    class_slug = band["spec"].split("-")[0]
    return _carry_talents(
        carried, _trees(root, old, class_slug), _trees(root, new, class_slug), lines
    )


def _carry_file(
    root: Path, payload: dict[str, Any], old: str, new: str, known: set[int], lines: list[str]
) -> dict[str, Any]:
    if payload.get("build") != old:
        raise BisCarryError(
            f"bis/{payload.get('spec')}.json is stamped {payload.get('build')!r}, not {old}"
        )
    bands = [_carry_band(root, band, old, new, known, lines) for band in payload.get("bands", [])]
    return {**payload, "build": new, "bands": bands}


def carry_bis(root: Path, old: str, new: str) -> list[str]:
    """Write builds/<new>/bis/ from builds/<old>/bis/ and return one report line per
    pick or alternative the new build's items.json lacks and per ranked talent its trees
    lack. Writes nothing when any file fails to carry."""
    source, target = root / old / "bis", root / new / "bis"
    if not source.is_dir():
        raise BisCarryError(f"{source} does not exist: nothing to carry")
    if target.exists():
        raise BisCarryError(f"{target} already has a ranking; the carry never overwrites one")
    known = _item_ids(root, new)
    lines: list[str] = []
    carried = {
        path.name: _carry_file(
            root, json.loads(path.read_text(encoding="utf-8")), old, new, known, lines
        )
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
