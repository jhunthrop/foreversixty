"""Merge the hand-maintained Forever facts into the client-derived data.

Blizzard states what changed in Forever in posts, not in client tables, so
`data/curated/{classes,races,combos}.json` holds those facts and this module
folds them into `classes.json`, `races.json` and `combos.json`. Every claim
must carry at least one source; an unsourced claim stops the pipeline rather
than shipping as if the client said it.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.models import Combo, ForeverChange, PlayableClass, PlayableRace, Source

SOURCE_KINDS = frozenset({"blizzard", "datamined", "community", "site"})


class CuratedError(SystemExit):
    """A curated file says something the pipeline will not publish."""


def _read(curated_dir: Path, name: str) -> list[dict]:
    path = curated_dir / f"{name}.json"
    if not path.exists():
        raise CuratedError(f"missing curated file {path}")
    return json.loads(path.read_text(encoding="utf-8"))


def _sources(raw: list[dict], where: str) -> list[Source]:
    if not raw:
        raise CuratedError(f"{where} needs at least one source")
    sources = []
    for entry in raw:
        source = Source(**entry)
        if source.kind not in SOURCE_KINDS:
            raise CuratedError(
                f"{where} has source kind {source.kind!r}; use one of {sorted(SOURCE_KINDS)}"
            )
        if not source.label.strip() or not source.url.strip():
            raise CuratedError(f"{where} has a source with an empty label or url")
        sources.append(source)
    return sources


def _changes(raw: list[dict], where: str) -> list[ForeverChange]:
    changes = []
    for index, entry in enumerate(raw):
        text = entry.get("text", "").strip()
        if not text:
            raise CuratedError(f"{where} change {index} has no text")
        sources = _sources(entry.get("sources", []), f"{where} change {index}")
        changes.append(ForeverChange(text=text, sources=sources))
    return changes


def merge_curated(
    classes: list[PlayableClass],
    races: list[PlayableRace],
    curated_dir: Path,
) -> tuple[list[PlayableClass], list[PlayableRace], list[Combo]]:
    merged_classes = _merge_classes(classes, _read(curated_dir, "classes"))
    merged_races = _merge_races(races, _read(curated_dir, "races"))
    combos = _build_combos(merged_classes, merged_races, _read(curated_dir, "combos"))
    return merged_classes, merged_races, combos


def _merge_classes(classes: list[PlayableClass], curated: list[dict]) -> list[PlayableClass]:
    by_slug = {record.slug: record for record in classes}
    changes_by_slug: dict[str, list[ForeverChange]] = {}
    for entry in curated:
        slug = entry["slug"]
        if slug not in by_slug:
            raise CuratedError(f"curated classes.json names unknown class slug {slug!r}")
        if slug in changes_by_slug:
            raise CuratedError(f"curated classes.json names class slug {slug!r} twice")
        changes_by_slug[slug] = _changes(entry.get("forever_changes", []), f"class {slug}")
    return sorted(
        (
            record.model_copy(update={"forever_changes": changes_by_slug.get(record.slug, [])})
            for record in classes
        ),
        key=lambda record: record.id,
    )


def _merge_races(races: list[PlayableRace], curated: list[dict]) -> list[PlayableRace]:
    by_slug = {record.slug: record for record in races}
    client_ids = {record.id for record in races}
    merged = {
        record.slug: record.model_copy(update={"forever_changes": [], "placeholder": False})
        for record in races
    }
    seen: set[str] = set()
    for entry in curated:
        slug = entry["slug"]
        if slug in seen:
            raise CuratedError(f"curated races.json names race slug {slug!r} twice")
        seen.add(slug)
        changes = _changes(entry.get("forever_changes", []), f"race {slug}")
        if slug in by_slug:
            merged[slug] = merged[slug].model_copy(update={"forever_changes": changes})
            continue
        if not entry.get("placeholder"):
            raise CuratedError(
                f"curated races.json names unknown race slug {slug!r} without placeholder: true"
            )
        race_id = int(entry["id"])
        if race_id in client_ids:
            raise CuratedError(
                f"placeholder race {slug!r} reuses client race id {race_id}; pick an unused id"
            )
        merged[slug] = PlayableRace(
            id=race_id,
            name=entry["name"],
            slug=slug,
            faction=entry["faction"],
            placeholder=True,
            forever_changes=changes,
        )
    return sorted(merged.values(), key=lambda record: record.id)


def _build_combos(
    classes: list[PlayableClass], races: list[PlayableRace], curated: list[dict]
) -> list[Combo]:
    class_ids = {record.id for record in classes}
    race_ids = {record.id for record in races}
    combos = []
    seen: set[tuple[int, int]] = set()
    for entry in curated:
        race_id, class_id = int(entry["race_id"]), int(entry["class_id"])
        where = f"combo race_id {race_id} class_id {class_id}"
        if race_id not in race_ids:
            raise CuratedError(f"{where} names unknown race_id {race_id}")
        if class_id not in class_ids:
            raise CuratedError(f"{where} names unknown class_id {class_id}")
        if (race_id, class_id) in seen:
            raise CuratedError(f"curated combos.json names {where} twice")
        seen.add((race_id, class_id))
        new_in_forever = bool(entry.get("new_in_forever", False))
        if new_in_forever:
            _sources(entry.get("sources", []), where)
        combos.append(Combo(race_id=race_id, class_id=class_id, new_in_forever=new_in_forever))
    return sorted(combos, key=lambda c: (c.race_id, c.class_id))
