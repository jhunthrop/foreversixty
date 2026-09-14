import json
import logging
import shutil
from collections.abc import Sequence
from dataclasses import dataclass
from pathlib import Path
from typing import Any

from pydantic import BaseModel

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class NormalizeResult:
    """What one normalize run produced, and what it could not.

    `skipped` is empty on a complete run and otherwise holds one human-readable
    reason per output the run could not emit. A caller that wants the soft
    behaviour (keep every other file, do not raise) just uses `build_dir` and
    ignores it; the CLI turns a non-empty `skipped` into a non-zero exit, so an
    incomplete build directory can never be committed by a green CI run.
    """

    build_dir: Path
    skipped: tuple[str, ...] = ()


def _write(payload: Any, path: Path) -> None:
    """Write one JSON payload in the pipeline's only serialization format.

    indent=2, ensure_ascii=False and a single trailing newline are the
    determinism contract, so they are defined here once and nowhere else.
    """
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")


def write_json(records: Sequence[BaseModel], path: Path) -> None:
    ordered = sorted(records, key=lambda r: r.id)  # type: ignore[attr-defined]
    _write([r.model_dump() for r in ordered], path)


def write_records(records: Sequence[BaseModel], path: Path) -> None:
    """Write a list of records in the order given, without sorting by id."""
    _write([r.model_dump() for r in records], path)


def write_model(record: BaseModel, path: Path) -> None:
    """Write a single record as a JSON object."""
    _write(record.model_dump(), path)


def normalize_build(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
) -> NormalizeResult:
    from pipeline.csvio import read_csv
    from pipeline.curated import merge_curated
    from pipeline.icons import icon_names
    from pipeline.manifest import write_manifest
    from pipeline.normalize.classes import normalize_classes, normalize_races
    from pipeline.normalize.dungeons import normalize_dungeons
    from pipeline.normalize.gear import ItemDataError, build_class_items, build_item_sets
    from pipeline.normalize.items import normalize_items
    from pipeline.normalize.spells import normalize_spells
    from pipeline.normalize.talent_trees import build_talent_trees
    from pipeline.normalize.talents import normalize_talents
    from pipeline.normalize.zones import normalize_zones
    from pipeline.spelltext import load_spell_text

    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    t = lambda name: read_csv(raw / f"{name}.csv")  # noqa: E731

    # Phase 0 flat entities.
    class_rows = t("ChrClasses")
    write_json(normalize_zones(t("AreaTable"), t("Map")), build_dir / "zones.json")
    write_json(normalize_dungeons(t("JournalInstance")), build_dir / "dungeons.json")
    write_json(normalize_items(t("ItemSparse"), t("Item")), build_dir / "items.json")
    write_json(normalize_spells(t("SpellName")), build_dir / "spells.json")
    write_json(normalize_talents(t("Talent"), t("TalentTab")), build_dir / "talents.json")

    # Phase 1 planner data. Both directories are rebuilt from scratch so a class
    # that disappears between builds does not leave a stale file behind.
    spell_text = load_spell_text(t("Spell"), t("SpellMisc"), t("SpellEffect"), t("SpellDuration"))
    icons = icon_names(t("ManifestInterfaceData"))
    spell_names = {int(r["ID"]): r["Name_lang"] for r in t("SpellName")}
    shutil.rmtree(build_dir / "talents", ignore_errors=True)
    for record in build_talent_trees(
        t("Talent"), t("TalentTab"), class_rows, spell_names, spell_text, icons, build
    ):
        write_model(record, build_dir / "talents" / f"{record.class_slug}.json")
    item_sets = build_item_sets(t("ItemSet"), t("ItemSetSpell"), spell_text)
    write_json(item_sets, build_dir / "sets.json")
    shutil.rmtree(build_dir / "items", ignore_errors=True)
    skipped: list[str] = []
    try:
        class_items = build_class_items(t("ItemSparse"), t("Item"), class_rows, icons, build)
    except ItemDataError as error:
        logger.warning("items not emitted for build %s: %s", build, error)
        skipped.append(f"items/: {error}")
    else:
        for record in class_items:
            write_model(record, build_dir / "items" / f"{record.class_slug}.json")

    # Curated Forever facts.
    classes, races, combos = merge_curated(
        normalize_classes(class_rows), normalize_races(t("ChrRaces")), curated_dir
    )
    write_json(classes, build_dir / "classes.json")
    write_json(races, build_dir / "races.json")
    write_records(combos, build_dir / "combos.json")

    meta = json.loads((raw / "_meta.json").read_text(encoding="utf-8"))
    write_manifest(
        build_dir, build=meta["build"], product=meta["product"], fetched_at=meta["fetched_at"]
    )
    return NormalizeResult(build_dir=build_dir, skipped=tuple(skipped))
