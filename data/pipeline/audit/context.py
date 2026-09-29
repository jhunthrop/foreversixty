"""Everything a check module (A-H) reads, loaded once per run.

Every loader here is tolerant of a missing input: the audit is read-only and
never fetches or regenerates anything (lane brief), so a raw table that was
never committed (raw/*.csv, raw/wowhead-gear-planner.js -- both gitignored,
see data/.gitignore) or a classic-db dump nobody pointed `--classicdb-dump`
at is simply absent, not an error. A check module reads the relevant
`AuditContext` property, finds it `None`/empty, and records
`CategoryResult.skipped` instead of guessing.
"""

from __future__ import annotations

import functools
import json
import logging
from pathlib import Path

from pipeline import classic_sources
from pipeline.csvio import read_csv

logger = logging.getLogger(__name__)


def _read_json(path: Path) -> dict | list | None:
    if not path.exists():
        return None
    return json.loads(path.read_text(encoding="utf-8"))


class AuditContext:
    """Lazily-loaded, read-only view of one build's committed data plus
    the curated/engine inputs a check needs alongside it."""

    def __init__(
        self,
        build: str,
        root: Path = Path("builds"),
        curated_dir: Path = Path("curated"),
        engine: Path | None = None,
        classicdb_dump: Path | None = None,
    ) -> None:
        self.build = build
        self.root = root
        self.build_dir = root / build
        self.curated_dir = curated_dir
        self.engine = engine
        self.classicdb_dump = classicdb_dump

    # -- client-derived, committed ------------------------------------

    @functools.cached_property
    def classes(self) -> list[dict]:
        return _read_json(self.build_dir / "classes.json") or []

    @functools.cached_property
    def class_slugs(self) -> list[str]:
        return sorted(entry["slug"] for entry in self.classes)

    @functools.cached_property
    def items_by_class(self) -> dict[str, dict]:
        """class slug -> the parsed `items/<slug>.json` document."""
        out: dict[str, dict] = {}
        for slug in self.class_slugs:
            doc = _read_json(self.build_dir / "items" / f"{slug}.json")
            if doc is not None:
                out[slug] = doc
        return out

    @functools.cached_property
    def items_by_id(self) -> dict[int, dict]:
        """item id -> (GearItem dict, class_slug) flattened across every
        class's items file -- most items appear under one class only, an
        item usable by more than one class appears once per class and
        the last one loaded wins (fine for the checks that only need
        stats/flags, which do not vary by class)."""
        out: dict[int, dict] = {}
        for doc in self.items_by_class.values():
            for row in doc.get("items", []):
                out[int(row["id"])] = row
        return out

    @functools.cached_property
    def loot(self) -> dict:
        empty: dict = {"sources": [], "quests": {}, "factions": {}}
        return _read_json(self.build_dir / "loot.json") or empty

    @functools.cached_property
    def zones(self) -> list[dict]:
        return _read_json(self.build_dir / "zones.json") or []

    @functools.cached_property
    def zone_map_id(self) -> dict[int, int]:
        return {int(z["id"]): int(z["map_id"]) for z in self.zones}

    @functools.cached_property
    def bis_specs(self) -> list[str]:
        bis_dir = self.build_dir / "bis"
        if not bis_dir.exists():
            return []
        return sorted(p.stem for p in bis_dir.glob("*.json"))

    @functools.cached_property
    def bis_by_spec(self) -> dict[str, dict]:
        out: dict[str, dict] = {}
        for spec in self.bis_specs:
            doc = _read_json(self.build_dir / "bis" / f"{spec}.json")
            if doc is not None:
                out[spec] = doc
        return out

    @functools.cached_property
    def addon_data(self) -> dict | None:
        return _read_json(self.build_dir / "addon-data.json")

    @functools.cached_property
    def item_sources(self) -> dict[int, dict]:
        doc = _read_json(self.build_dir / "raw" / "items" / "item-sources.json")
        if doc is None:
            return {}
        return {int(item_id): row for item_id, row in doc.get("items", {}).items()}

    @functools.cached_property
    def quest_levels(self) -> dict[int, dict]:
        doc = _read_json(self.build_dir / "raw" / "quests" / "quest-levels.json")
        if doc is None:
            return {}
        entries = doc.get("quests", {})
        return {int(quest_id): row for quest_id, row in entries.items()}

    @functools.cached_property
    def classic_sources_by_item(self) -> dict[int, list]:
        """item id -> `ClassicDbSourceRecord` list, from the committed
        `raw/classicdb/sources.json` cache (`pipeline.classic_sources.
        load_classic_sources`) -- `{}` (logged) when that cache was never
        fetched for this build."""
        return classic_sources.load_classic_sources(self.build_dir)

    # -- raw client tables (gitignored; usually absent) -----------------

    @functools.cached_property
    def raw_dir(self) -> Path:
        return self.build_dir / "raw"

    def _read_raw_csv(self, name: str) -> list[dict[str, str]] | None:
        path = self.raw_dir / name
        if not path.exists():
            return None
        return read_csv(path)

    @functools.cached_property
    def raw_item_sparse(self) -> list[dict[str, str]] | None:
        return self._read_raw_csv("ItemSparse.csv")

    @functools.cached_property
    def raw_item(self) -> list[dict[str, str]] | None:
        return self._read_raw_csv("Item.csv")

    @functools.cached_property
    def raw_tables_available(self) -> bool:
        return self.raw_item_sparse is not None and self.raw_item is not None

    @functools.cached_property
    def wowhead_gear_planner(self) -> dict[int, object] | None:
        """item id -> `pipeline.wowhead_items.WowheadItem`, from the
        committed gear-planner payload (`raw/wowhead-gear-planner.js` --
        gitignored, so usually absent for a build this audit sees).
        `None` when the payload is missing or is not Forever's (same
        `WowheadPayloadError` `pipeline.normalize.wowhead.load_supplement`
        tolerates)."""
        from pipeline.wowhead_items import WowheadPayloadError, load_items, raw_path

        path = raw_path(self.build_dir)
        if not path.exists():
            return None
        try:
            return {item.id: item for item in load_items(path)}
        except WowheadPayloadError:
            logger.warning("audit: %s did not parse as a Forever gear-planner payload", path)
            return None

    # -- curated / engine -------------------------------------------------

    @functools.cached_property
    def dump(self):
        """The optional `--classicdb-dump` reader (checks C, D, E's own
        item_template/spell_template/creature-spawn lookups), or `None`
        when no path was given."""
        from pipeline.audit.dumpdb import load_dump

        return load_dump(self.classicdb_dump)

    @functools.cached_property
    def specs(self) -> list[dict]:
        return _read_json(self.curated_dir / "specs.json") or []

    @functools.cached_property
    def apl_dir(self) -> Path:
        return self.curated_dir / "apl"
