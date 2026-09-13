import json
from collections.abc import Sequence
from pathlib import Path

from pydantic import BaseModel


def write_json(records: Sequence[BaseModel], path: Path) -> None:
    ordered = sorted(records, key=lambda r: r.id)  # type: ignore[attr-defined]
    payload = [r.model_dump() for r in ordered]
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")


def write_model(record: BaseModel, path: Path) -> None:
    """Write a single record as a JSON object."""
    path.parent.mkdir(parents=True, exist_ok=True)
    payload = json.dumps(record.model_dump(), indent=2, ensure_ascii=False) + "\n"
    path.write_text(payload, encoding="utf-8")


def write_records(records: Sequence[BaseModel], path: Path) -> None:
    """Write a list of records in the order given, without sorting by id."""
    path.parent.mkdir(parents=True, exist_ok=True)
    payload = [r.model_dump() for r in records]
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")


def normalize_build(build: str, root: Path = Path("builds")) -> Path:
    from pipeline.csvio import read_csv
    from pipeline.manifest import write_manifest
    from pipeline.normalize.classes import normalize_classes, normalize_races
    from pipeline.normalize.dungeons import normalize_dungeons
    from pipeline.normalize.items import normalize_items
    from pipeline.normalize.spells import normalize_spells
    from pipeline.normalize.talents import normalize_talents
    from pipeline.normalize.zones import normalize_zones

    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    t = lambda name: read_csv(raw / f"{name}.csv")  # noqa: E731
    write_json(normalize_zones(t("AreaTable"), t("Map")), build_dir / "zones.json")
    write_json(normalize_dungeons(t("JournalInstance")), build_dir / "dungeons.json")
    write_json(normalize_items(t("ItemSparse"), t("Item")), build_dir / "items.json")
    write_json(normalize_spells(t("SpellName")), build_dir / "spells.json")
    write_json(normalize_classes(t("ChrClasses")), build_dir / "classes.json")
    write_json(normalize_races(t("ChrRaces")), build_dir / "races.json")
    write_json(normalize_talents(t("Talent"), t("TalentTab")), build_dir / "talents.json")
    meta = json.loads((raw / "_meta.json").read_text(encoding="utf-8"))
    write_manifest(
        build_dir, build=meta["build"], product=meta["product"], fetched_at=meta["fetched_at"]
    )
    return build_dir
