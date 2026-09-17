"""What changed between the Wowhead pre-beta snapshot and the client's trees.

`data/raw-forever/wowhead-talents-2026-09-14.json` is the record of what was
believed about Forever's talents before the beta client shipped, and the
planner was built on it. The client now says; this writes down every place the
two disagree, keyed by the cell each talent sits in, so the difference is a
committed artefact rather than something a reader has to take on trust.

The snapshot keys its talents by `TalentTab` id, which is the same id the
emitted trees use, so no name matching is needed to pair up trees.
"""

from __future__ import annotations

import json
from collections.abc import Sequence
from pathlib import Path

from pipeline.models import ClassTalents, TalentEntry, TalentTree


def _prereq(talent: TalentEntry, by_id: dict[int, TalentEntry]) -> list | None:
    if talent.prereq_talent_id is None:
        return None
    prereq = by_id.get(talent.prereq_talent_id)
    return [prereq.name if prereq else str(talent.prereq_talent_id), talent.prereq_rank]


def _snapshot_prereq(talent: dict, tree: dict) -> list | None:
    requires = talent.get("requires") or []
    if not requires:
        return None
    first = requires[0]
    return [tree[str(first["id"])]["name"], first["qty"]]


def _tree_diff(snapshot_tree: dict, tree: TalentTree) -> dict:
    """Every disagreement inside one tree.

    Names are the identity, not cells: a talent that keeps its name and moves
    is a move, and a cell whose occupant's name is on neither side's shared
    list is a rename. What is left over on one side only is an addition or a
    removal.
    """
    snap_cell = {t["name"]: (t["row"], t["col"]) for t in snapshot_tree.values()}
    build_cell = {t.name: (t.tier, t.column) for t in tree.talents}
    only_snapshot = set(snap_cell) - set(build_cell)
    only_build = set(build_cell) - set(snap_cell)

    snap_by_cell = {(t["row"], t["col"]): t for t in snapshot_tree.values()}
    build_by_cell = {(t.tier, t.column): t for t in tree.talents}
    renamed = [
        {
            "cell": list(cell),
            "snapshot": snap_by_cell[cell]["name"],
            "build": build_by_cell[cell].name,
        }
        for cell in sorted(set(snap_by_cell) & set(build_by_cell))
        if snap_by_cell[cell]["name"] in only_snapshot and build_by_cell[cell].name in only_build
    ]
    was_renamed = {change["snapshot"] for change in renamed}
    now_renamed = {change["build"] for change in renamed}

    snap_by_name = {t["name"]: t for t in snapshot_tree.values()}
    build_by_name = {t.name: t for t in tree.talents}
    added = [
        {"cell": list(build_cell[name]), "name": name, "ranks": build_by_name[name].max_rank}
        for name in sorted(only_build - now_renamed)
    ]
    removed = [
        {"cell": list(snap_cell[name]), "name": name, "ranks": len(snap_by_name[name]["ranks"])}
        for name in sorted(only_snapshot - was_renamed)
    ]
    shared = sorted(set(snap_cell) & set(build_cell))
    moved = [
        {"name": name, "snapshot": list(snap_cell[name]), "build": list(build_cell[name])}
        for name in shared
        if snap_cell[name] != build_cell[name]
    ]

    by_id = {t.id: t for t in tree.talents}
    rank_changes, prerequisite_changes = [], []
    for name in shared:
        before, after = snap_by_name[name], build_by_name[name]
        if len(before["ranks"]) != after.max_rank:
            rank_changes.append(
                {"name": name, "snapshot": len(before["ranks"]), "build": after.max_rank}
            )
        was, now = _snapshot_prereq(before, snapshot_tree), _prereq(after, by_id)
        if was != now:
            prerequisite_changes.append({"name": name, "snapshot": was, "build": now})
    return {
        "tree_id": tree.id,
        "tree": tree.name,
        "added": added,
        "removed": removed,
        "renamed": renamed,
        "moved": moved,
        "rank_changes": rank_changes,
        "prerequisite_changes": prerequisite_changes,
    }


_CHANGE_KEYS = ("added", "removed", "renamed", "moved", "rank_changes", "prerequisite_changes")


def diff_snapshot(snapshot: dict, records: Sequence[ClassTalents]) -> dict:
    talents_by_tree = snapshot.get("talents") or {}
    trees = []
    build_total = 0
    for record in sorted(records, key=lambda r: r.class_id):
        for tree in sorted(record.trees, key=lambda t: t.position):
            build_total += len(tree.talents)
            snapshot_tree = talents_by_tree.get(str(tree.id))
            if snapshot_tree is None:
                trees.append(
                    {
                        "tree_id": tree.id,
                        "tree": tree.name,
                        "class_slug": record.class_slug,
                        "missing_from_snapshot": True,
                    }
                )
                continue
            diff = _tree_diff(snapshot_tree, tree)
            if any(diff[key] for key in _CHANGE_KEYS):
                trees.append({**diff, "class_slug": record.class_slug})
    return {
        "totals": {
            "snapshot": sum(len(t) for t in talents_by_tree.values()),
            "build": build_total,
        },
        "trees": trees,
    }


def write_snapshot_diff(
    snapshot: str,
    build: str,
    root: Path = Path("data/builds"),
    out: Path = Path("data/diffs"),
) -> Path:
    """Write the diff between `snapshot` and `build`'s talent trees.

    `snapshot`, `root` and `out` are all resolved the same way: relative to
    the repository root when they are relative paths (an absolute path is
    used as-is). That matches `snapshot`'s own default, which already names a
    `data/`-prefixed path, and makes this function give the same answer
    regardless of the caller's working directory -- unlike most of this
    pipeline's own `root` parameters (e.g. `normalize_build`'s), which are
    deliberately resolved against the CLI's cwd instead.
    """
    repo = Path(__file__).resolve().parents[2]
    payload = json.loads((repo / snapshot).read_text(encoding="utf-8"))
    records = [
        ClassTalents.model_validate_json(path.read_text(encoding="utf-8"))
        for path in sorted((repo / root / build / "talents").glob("*.json"))
    ]
    result = {"snapshot": snapshot, "build": build, **diff_snapshot(payload, records)}
    out_dir = repo / out
    out_dir.mkdir(parents=True, exist_ok=True)
    name = Path(snapshot).stem.replace("wowhead-talents-", "wowhead-")
    path = out_dir / f"{name}__{build}.json"
    path.write_text(json.dumps(result, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    counts = {key: sum(len(t.get(key, [])) for t in result["trees"]) for key in _CHANGE_KEYS}
    print(
        f"{result['totals']['snapshot']} snapshot talents, {result['totals']['build']} in "
        f"{build}: " + ", ".join(f"{key} {value}" for key, value in counts.items())
    )
    return path
