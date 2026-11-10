"""Cross-build floors: what a build is compared with when it has no run of its own.

The shrink gates (`csvio.check_item_sparse_completeness`,
`normalize._check_class_items_not_shrunk`) used to compare a build only with its own
last committed run, so a build's FIRST run always passed. On 2026-10-29 wowhead's
Forever feed lost 292 of the ~700 items it alone carries, and 1.60.1.70685 and
1.60.1.70707 were both first runs: they committed vanilla stats with every gate green
(design/reviews/2026-10-30-70707-dps-drop.md).

A first run is now compared with the newest other build, by the same rule data.yml's
seed step uses (`ls builds | grep -E '^[0-9]+(\\.[0-9]+)+$' | sort -V | tail -1`).
Every threshold here is a floor, never an exact count.
"""

from __future__ import annotations

import json
import re
from collections.abc import Iterable, Mapping
from dataclasses import dataclass
from pathlib import Path

#: A build's own last run may not lose a single item id (the gate's original rule);
#: a first run may lose at most 2% against the previous build, data.yml's own
#: "No silent catalogue shrink against the previous build" tolerance.
CROSS_BUILD_ITEMS_RATIO = 0.98

#: Of the reference run's items whose stats came from wowhead's feed
#: (`stats_source == "wowhead"`), at least 90% must still take live-game stats. A
#: wowhead item that falls back to classic-db keeps the item count level while its
#: stats roll back to 1.12, so the count gate alone cannot see it. An item that now
#: takes its stats from the client's own row (`stats_source` None: a hotfix brought
#: its ItemSparse row in) is not a fallback: 1.60.1.70790's hotfixed ItemSparse
#: took over 186 of the 288 items wowhead had carried, which a bare count of
#: wowhead-sourced items read as a lost feed.
WOWHEAD_ITEMS_FLOOR_RATIO = 0.9

#: The `stats_source` values that are live-game stats: wowhead's feed, or the
#: client's own row (None).
LIVE_STATS_SOURCES = frozenset({"wowhead", None})

#: `fetch-wowhead` carries the seeded records a live feed no longer lists (70707's
#: refetch listed 397 of 689, 58%), and refuses a carried catalogue that keeps fewer
#: than 90% of the seeded supplement (the planner gear the client lacks): only
#: records whose item left the client's Item table can be lost that way.
MIN_CARRIED_COVERAGE = 0.9

_BUILD_NAME = re.compile(r"^[0-9]+(\.[0-9]+)+$")


def _version(name: str) -> tuple[int, ...]:
    return tuple(int(part) for part in name.split("."))


def previous_build_dir(build_dir: Path, having: str) -> Path | None:
    """The newest other build directory beside `build_dir` that has `having`
    (a path relative to the build directory), or None."""
    root = build_dir.parent
    if not root.is_dir():
        return None
    candidates = [
        path
        for path in root.iterdir()
        if path.is_dir()
        and _BUILD_NAME.match(path.name)
        and path.name != build_dir.name
        and (path / having).exists()
    ]
    return max(candidates, key=lambda path: _version(path.name), default=None)


@dataclass(frozen=True)
class Reference:
    """The run a build is compared with: its own last one, or the previous build's."""

    build_dir: Path
    own: bool

    @property
    def ratio(self) -> float:
        """The items floor against this reference."""
        return 1.0 if self.own else CROSS_BUILD_ITEMS_RATIO


def reference_for(build_dir: Path, having: str) -> Reference | None:
    if (build_dir / having).exists():
        return Reference(build_dir, own=True)
    previous = previous_build_dir(build_dir, having)
    return Reference(previous, own=False) if previous is not None else None


def committed_items(build_dir: Path) -> list[dict]:
    path = build_dir / "items.json"
    return json.loads(path.read_text(encoding="utf-8")) if path.exists() else []


def wowhead_item_ids(items: Iterable[dict]) -> set[int]:
    """The ids of the items that take their stats from wowhead's feed."""
    return {item["id"] for item in items if item.get("stats_source") == "wowhead"}


def wowhead_items_kept(
    previous: Iterable[dict], sources: Mapping[int, str | None]
) -> tuple[int, int]:
    """Of `previous`'s wowhead-sourced items, how many `sources` (item id ->
    `stats_source` of the new run) still gives live-game stats, and how many there
    were. An id missing from `sources` is lost."""
    ids = wowhead_item_ids(previous)
    kept = sum(1 for item_id in ids if sources.get(item_id, "missing") in LIVE_STATS_SOURCES)
    return kept, len(ids)


def stats_sources(items: Iterable[dict]) -> dict[int, str | None]:
    """Item id -> `stats_source`, for committed item dicts."""
    return {item["id"]: item.get("stats_source") for item in items}


def class_item_counts(build_dir: Path) -> dict[str, int]:
    """Class slug -> item count of each committed `items/<class>.json`."""
    return {
        path.stem: len(json.loads(path.read_text(encoding="utf-8")).get("items", []))
        for path in sorted((build_dir / "items").glob("*.json"))
    }


def floor_problems(build_dir: Path, reference_dir: Path) -> list[str]:
    """Why `build_dir`'s committed items fall below `reference_dir`'s floors, if they
    do: the item count and the wowhead-sourced items still on live-game stats, both
    with the cross-build ratios. Empty when both hold."""
    new, old = committed_items(build_dir), committed_items(reference_dir)
    problems = []
    if len(new) < len(old) * CROSS_BUILD_ITEMS_RATIO:
        problems.append(
            f"{build_dir.name} has {len(new)} items, below {CROSS_BUILD_ITEMS_RATIO:.0%} "
            f"of {reference_dir.name}'s {len(old)}"
        )
    kept, old_wowhead = wowhead_items_kept(old, stats_sources(new))
    if kept < old_wowhead * WOWHEAD_ITEMS_FLOOR_RATIO:
        problems.append(
            f"{build_dir.name} keeps live-game stats on {kept} of {reference_dir.name}'s "
            f"{old_wowhead} wowhead-sourced items, below {WOWHEAD_ITEMS_FLOOR_RATIO:.0%}"
        )
    return problems
