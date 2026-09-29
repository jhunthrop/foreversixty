"""Tenet 8's own instrument (docs/tenets.md, owner ruling 2026-09-29): every
published fact compared to its primary source, with counts, examples and a
machine-readable report -- so a controller can run this after every
regeneration, spawn fix lanes per category, and rerun until it is clean.

`run_audit` is the library entry point (`python -m pipeline audit` wraps
it); it never writes into `data/builds/` -- `--out` defaults to a scratch
directory -- and never fetches anything over the network. A check whose own
brief needs a source this run does not have (raw client CSVs, a classic-db
dump) says so in its own `CategoryResult.skipped` rather than fabricating a
verdict (tenet 8's own "unverifiable data is labelled or left out").
"""

from __future__ import annotations

from pathlib import Path

from pipeline.audit import (
    check_addon,
    check_bis,
    check_crafted,
    check_curated,
    check_drops,
    check_items,
    check_quests,
    check_vendors,
)
from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult
from pipeline.audit.report import summary_line, write_report

#: Run in this order (A-H); `report.render_markdown` keeps it.
_CHECKS = (
    check_items,
    check_quests,
    check_drops,
    check_vendors,
    check_crafted,
    check_bis,
    check_addon,
    check_curated,
)


def run_audit(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
    engine: Path | None = None,
    classicdb_dump: Path | None = None,
) -> list[CategoryResult]:
    ctx = AuditContext(
        build, root=root, curated_dir=curated_dir, engine=engine, classicdb_dump=classicdb_dump
    )
    return [module.check(ctx) for module in _CHECKS]


__all__ = ["run_audit", "summary_line", "write_report"]
