"""Write a pre-beta Forever build from a saved Wowhead talent snapshot.

This exists only until the beta client ships on 2026-09-17. Wowhead's Forever data
environment serves real talents but Classic Era items, so the build this writes carries
Forever talents beside the Era tables it was derived from, and its manifest says so per
file. Nothing here fetches: the snapshot is committed under `data/raw-forever/` so the
output is reproducible after the endpoint changes.
"""

from __future__ import annotations

import csv
import hashlib
import io
import json
import logging
import shutil
from pathlib import Path

import httpx

from pipeline.icons import CACHE_DIR, PLACEHOLDER_ICON, download_icons, icon_names
from pipeline.normalize.forever_talents import normalize_forever_talents
from pipeline.wago import BASE_URL, USER_AGENT

log = logging.getLogger("pipeline.forever")

#: Forever's talents reuse icons Blizzard shipped in later expansions, which a Classic
#: Era client has never held. Their art is unchanged, so the retail manifest resolves the
#: names and CASC serves the same files. Icons Forever invented for itself carry a
#: `classic_` prefix and exist in no shipped client, so they wait for the beta.
ICON_FALLBACK_BUILD = "12.1.0.69814"

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
    # The planner's gear picker reads the per-class item files, not items.json alone.
    if (src / "items").is_dir() and not (dst / "items").exists():
        shutil.copytree(src / "items", dst / "items")

    # The web's sync-data reads manifest["files"] and refuses to publish a build whose
    # manifest lists no talents/*.json, falling back to the checked-in fixture. The keys
    # are what it looks for; the digests let a later diff see a file move.
    files: dict[str, str] = {}
    for path in sorted(dst.rglob("*.json")):
        if path.name == "manifest.json":
            continue
        rel = path.relative_to(dst).as_posix()
        files[rel] = hashlib.sha256(path.read_bytes()).hexdigest()

    talent_count = sum(len(t.talents) for e in per_class for t in e.trees)
    prereq_count = sum(
        1 for e in per_class for t in e.trees for x in t.talents if x.prereq_talent_id
    )
    (dst / "manifest.json").write_text(
        json.dumps(
            {
                "build": build,
                "product": "wow_forever_prebeta",
                "files": files,
                "kind": "forever-prebeta",
                "derived_from": from_build,
                "snapshot": snapshot,
                "provenance": {
                    "talents.json": "forever",
                    "talents/": "forever",
                    "items/": "classic-era",
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


def _manifest(build: str, client: httpx.Client) -> list[dict[str, str]]:
    response = client.get(f"/db2/ManifestInterfaceData/csv?build={build}", timeout=240)
    response.raise_for_status()
    return list(csv.DictReader(io.StringIO(response.text)))


CLASS_ICONS = (
    "warrior",
    "paladin",
    "hunter",
    "rogue",
    "priest",
    "shaman",
    "mage",
    "warlock",
    "druid",
)


def fetch_missing_icons(
    build: str,
    fallback_build: str = ICON_FALLBACK_BUILD,
    client: httpx.Client | None = None,
    cache_dir: Path = CACHE_DIR,
) -> tuple[int, list[str]]:
    """Download the icons a Forever build names but the source build never had.

    Returns the number written and the names that no shipped client holds, which the
    site renders as the placeholder until the beta client has them.
    """
    root = Path(__file__).resolve().parents[2]
    build_dir = root / "data" / "builds" / build
    icons_dir = build_dir / "icons"

    referenced: set[str] = set()
    for path in sorted((build_dir / "talents").glob("*.json")):
        payload = json.loads(path.read_text())
        for tree in payload["trees"]:
            for talent in tree["talents"]:
                if talent["icon"]:
                    referenced.add(talent["icon"].lower())
    # The class icons the report view puts beside every player's name, which no talent
    # references and the source build only partly carried.
    for klass in CLASS_ICONS:
        referenced.add(f"classicon_{klass}")
    have = {path.stem.lower() for path in icons_dir.glob("*.webp")}
    missing = referenced - have
    if not missing:
        return 0, []

    own = client is None
    if client is None:
        client = httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    try:
        names = icon_names(_manifest(fallback_build, client))
        wanted = {file_id: name for file_id, name in names.items() if name in missing}
        # Lowest file id wins a shared name, the rule download_icons already applies.
        by_name: dict[str, int] = {}
        for file_id, name in sorted(wanted.items()):
            by_name.setdefault(name, file_id)
        written = download_icons(
            {file_id: name for name, file_id in by_name.items()},
            icons_dir,
            cache_dir=cache_dir,
            client=client,
        )
    finally:
        if own:
            client.close()

    unresolved = sorted(missing - set(by_name))
    if unresolved:
        # The talent keeps no name we cannot draw: an icon file that never arrives would
        # render as a broken image on every planner page. The beta client brings the real
        # art and the next build picks it up.
        repointed = _repoint_to_placeholder(build_dir, set(unresolved))
        log.warning(
            "%d icon names are in no shipped client; %d talents now render as %s "
            "until the beta: %s",
            len(unresolved),
            repointed,
            PLACEHOLDER_ICON,
            unresolved,
        )
    log.info("wrote %d icons into %s", written, icons_dir)
    return written, unresolved


def _repoint_to_placeholder(build_dir: Path, names: set[str]) -> int:
    """Point every talent whose icon is in `names` at the placeholder. Returns the count."""
    changed = 0
    for path in sorted((build_dir / "talents").glob("*.json")):
        payload = json.loads(path.read_text())
        dirty = False
        for tree in payload["trees"]:
            for talent in tree["talents"]:
                if talent["icon"].lower() in names:
                    talent["icon"] = PLACEHOLDER_ICON
                    dirty = True
                    changed += 1
        if dirty:
            path.write_text(json.dumps(payload, indent=1, sort_keys=True) + "\n")
    return changed
