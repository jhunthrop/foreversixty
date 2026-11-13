"""The server's item push for a build, read from wago.tools' hotfix records.

Until 2026-11-13 the item tables' pushes came only from wago.tools' hotfixed
CSV export (`pipeline.wago_hotfixes`, `/db2/<Table>/csv?useHotfixes=1`).
Between the 1.60.1.70790 run (2026-11-12) and the 70938 run (2026-11-13) that
export stopped applying push 16778086 on every build, so Fiery War Axe (870)
went back to a 2.9 s swing (the server's is 3.6 s) and Blade of Hanna (2801)
to its pre-push speed and stats (issue #27). wago's hotfix page still lists
the push: its enUS copies of the record carry no values, its deDE and frFR
copies carry ItemDelay 3600.

This module reads that page the way `pipeline.hotfix_push` reads it for the
spell side, but in every locale (`EVERY_LOCALE`), and saves it in
`raw/wago-push/items.json` (committed with the build, crawled again only
where wago's totals moved: `pipeline.hotfix_push_store`). `layer_item_push`
puts it over the merged item table last, by `pipeline.hotfix_push.layer_rows`'
rule: per record the newest entry with values wins, a removal drops the row,
and a row read from another locale keeps the localised text underneath.

The page's search is a substring match, so "Item" would also match every
ItemSparse, ItemEffect and ItemSearchName entry (252,960 at 70938); ItemSparse
alone is 80,039, so only ItemSparse is read here. The Item table's pushes
still come from the CSV export.
"""

from __future__ import annotations

from pathlib import Path

import httpx

from pipeline.hotfix_push import EVERY_LOCALE, PUSH_DIR, PushEntry, layer_rows, newest_per_record
from pipeline.hotfix_push_store import cache_dir_from_env, fetch_saved, read_store

#: The item tables whose push is read from the hotfix page.
ITEM_PUSH_TABLES: tuple[str, ...] = ("ItemSparse",)
#: Under raw/wago-push/: the saved item push, committed with the build.
ITEM_STORE = "items.json"


def item_store(raw: Path) -> Path:
    return raw / PUSH_DIR / ITEM_STORE


def deciding_entries(entries: list[PushEntry]) -> list[PushEntry]:
    """The one entry per record the layer reads. Every push of an item comes
    in up to eleven locales with the same values, so the saved file keeps the
    deciding entry only (a tenth of the crawl); a later push still outranks
    it, so the incremental crawl lands on the same result."""
    newest = newest_per_record(entries)
    return [newest[record_id] for record_id in sorted(newest)]


def fetch_item_push(
    build: str,
    root: Path = Path("builds"),
    full: bool = False,
    client: httpx.Client | None = None,
) -> dict[str, int]:
    """`python -m pipeline fetch-item-push` (and the tail of `fetch`): crawl
    the build's item push into `raw/wago-push/items.json` and return how many
    records each table's push decides."""
    raw = root / build / "raw"
    if not raw.is_dir():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    store = item_store(raw)
    pushed = fetch_saved(
        build,
        ITEM_PUSH_TABLES,
        store,
        full=full,
        client=client,
        cache_dir=cache_dir_from_env(),
        locales=EVERY_LOCALE,
        keep=deciding_entries,
    )
    return {table: len(entries) for table, entries in pushed.items()}


def layer_item_push(
    rows: list[dict[str, str]], columns: list[str], raw: Path, table: str
) -> list[dict[str, str]]:
    """`rows` (the merged `table`, whose shipped header is `columns`) with the
    build's saved item push over it, or `rows` unchanged when the build has
    none for `table`."""
    saved = read_store(item_store(raw))
    if saved is None or table not in saved["entries"]:
        return rows
    if saved["build"] != raw.parent.name:
        store, build = item_store(raw), raw.parent.name
        raise SystemExit(f"{store} is the push of {saved['build']}, not {build}")
    layered, _ = layer_rows(rows, columns, saved["entries"][table], table)
    return layered
