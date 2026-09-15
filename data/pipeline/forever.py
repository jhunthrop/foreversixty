"""Write a pre-beta Forever build from a saved Wowhead talent snapshot.

This exists only until the beta client ships on 2026-09-17. Wowhead's Forever data
environment serves real talents but Classic Era items, so the build this writes carries
Forever talents beside the Era tables it was derived from, and its manifest says so per
file. Nothing here fetches: the snapshot is committed under `data/raw-forever/` so the
output is reproducible after the endpoint changes.
"""

from __future__ import annotations

import json
import logging
import shutil
from pathlib import Path

from pipeline.normalize.forever_talents import normalize_forever_talents

log = logging.getLogger("pipeline.forever")

#: Files copied from the source build unchanged. They are Classic Era data and the
#: manifest marks them so; the planner needs them to render a page at all.
CARRIED_OVER = (
    "classes.json",
    "races.json",
    "combos.json",
    "items.json",
    "sets.json",
    "dungeons.json",
)


def write_forever_talents(snapshot: str, from_build: str, build: str) -> Path:
    root = Path(__file__).resolve().parents[2]
    builds = root / "data" / "builds"
    src = builds / from_build
    dst = builds / build
    if not src.is_dir():
        raise FileNotFoundError(f"source build {src} does not exist")

    payload = json.loads((root / snapshot).read_text())
    classes = json.loads((src / "classes.json").read_text())

    tree_class: dict[int, int] = {}
    tree_names: dict[int, str] = {}
    for path in sorted((src / "talents").glob("*.json")):
        data = json.loads(path.read_text())
        for tree in data["trees"]:
            tree_class[tree["id"]] = data["class_id"]
            tree_names[tree["id"]] = tree["name"]

    per_class = normalize_forever_talents(
        payload,
        build=build,
        classes=classes,
        tree_class=tree_class,
        tree_names=tree_names,
    )

    (dst / "talents").mkdir(parents=True, exist_ok=True)
    for entry in per_class:
        out = dst / "talents" / f"{entry.class_slug}.json"
        out.write_text(json.dumps(entry.model_dump(), indent=1, sort_keys=True) + "\n")

    flat = [
        {
            "id": talent.id,
            "tab_id": tree.id,
            "tab_name": tree.name,
            "class_id": entry.class_id,
            "tier": talent.tier,
            "column": talent.column,
            "spell_ids": [rank.spell_id for rank in talent.ranks],
            "prereq_talent_id": talent.prereq_talent_id,
        }
        for entry in per_class
        for tree in entry.trees
        for talent in tree.talents
    ]
    (dst / "talents.json").write_text(json.dumps(flat, indent=1, sort_keys=True) + "\n")

    for name in CARRIED_OVER:
        if (src / name).is_file():
            shutil.copy2(src / name, dst / name)
    if (src / "icons").is_dir() and not (dst / "icons").exists():
        shutil.copytree(src / "icons", dst / "icons")

    talent_count = sum(len(t.talents) for e in per_class for t in e.trees)
    prereq_count = sum(
        1 for e in per_class for t in e.trees for x in t.talents if x.prereq_talent_id
    )
    (dst / "manifest.json").write_text(
        json.dumps(
            {
                "build": build,
                "kind": "forever-prebeta",
                "derived_from": from_build,
                "snapshot": snapshot,
                "provenance": {
                    "talents.json": "forever",
                    "talents/": "forever",
                    **{name: "classic-era" for name in CARRIED_OVER},
                    "icons/": "classic-era",
                },
                "counts": {
                    "classes": len(per_class),
                    "trees": sum(len(e.trees) for e in per_class),
                    "talents": talent_count,
                    "talents_with_prereq": prereq_count,
                },
                "note": (
                    "Talents are real Forever data from Wowhead's classicplus environment. "
                    "Everything else is Classic Era, carried over so the planner can render. "
                    "Replace the whole build from the beta client on 2026-09-17."
                ),
            },
            indent=1,
            sort_keys=True,
        )
        + "\n"
    )
    log.info(
        "wrote %s: %d classes, %d trees, %d talents, %d with a prerequisite",
        build,
        len(per_class),
        sum(len(e.trees) for e in per_class),
        talent_count,
        prereq_count,
    )
    return dst
