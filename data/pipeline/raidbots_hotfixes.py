"""Blizzard's hotfix push for a build, from Raidbots' public copy of the client cache.

The CDN (`pipeline.cdn_tables`) gives exactly the shipped tables; the server
then pushes hotfixes to a logged-in client, which keeps them in `DBCache.bin`.
Raidbots publishes the enUS cache its own Forever client receives at
`RAIDBOTS_URL` (the owner approved reading it on 2026-11-18; it is credited
on the site's sources page). One GET per run, conditional on the ETag and
Last-Modified of the copy already cached, with the honest User-Agent.

The cache is layered only when its header names the build being fetched: a
push is for one build, and an older client's cache replayed onto a newer
client (or the reverse) is how the item rule in `pipeline.hotfix_merge` was
burned before. Two layers, the same ones wago.tools' crawls fed:

* every table in `pipeline.hotfix_push.PUSH_TABLES` gets its pushed rows by
  `hotfix_push.layer_rows` (valid replaces or adds, removed drops, invalid and
  not-public withhold), through `hotfix_push.layer_hotfix_push`, so the layer
  record and every reader of `raw/` stay as they were;
* the item pair (`ItemSparse`, `Item`) is written to `raw/wago-hotfixes/` with
  the build in `_HotfixBuild`, the slot `pipeline.hotfix_merge` reads over the
  committed play-session dump.

A record is decoded with the table's WoWDBDefs definition (the layout the
build's own DB2 reports): a hotfix row is the uncompressed record, every
column in definition order except a non-inline id (the record id), strings
inline.
"""

from __future__ import annotations

import csv
import json
import logging
import os
from collections import defaultdict
from collections.abc import Mapping
from dataclasses import dataclass
from email.utils import formatdate
from pathlib import Path
from typing import Any

import httpx

from pipeline.blizzard_cdn import CdnClient, CdnMissing, ClientBuild, cache_root, open_build
from pipeline.cdn_tables import definition
from pipeline.client_tables import FILE_DATA_IDS
from pipeline.db2 import _BitCursor, _read_header
from pipeline.dbd import Field, read_definition
from pipeline.hotfix_cache import HotfixRecord, HotfixStatus, read_cache
from pipeline.hotfix_push import PUSH_TABLES, PushEntry, layer_hotfix_push, newest_per_record
from pipeline.hotfix_push_store import deciding_entries, read_store, write_store
from pipeline.wago_hotfixes import HOTFIX_BUILD_COLUMN, WAGO_HOTFIX_DIR
from pipeline.wago_hotfixes import TABLES as ITEM_TABLES
from pipeline.wago_http import USER_AGENT
from pipeline.wdc5 import _cell, _typed, layout_hash

logger = logging.getLogger(__name__)

RAIDBOTS_URL = "https://storage.googleapis.com/raidbots-static/wow/forever/enUS/DBCache.bin"
CACHE_NAME = "raidbots/forever-enUS-DBCache.bin"
#: Under raw/: the decoded push, committed, so a later fetch of the same build
#: (bis.yml, engine.yml) layers exactly what data.yml layered.
STORE_DIR = "raidbots"
STORE = "push.json"
#: Set to a local DBCache.bin (the owner's own client's) to layer it instead.
CACHE_FILE_ENV = "FOREVER_DBCACHE"


class HotfixDecodeError(ValueError):
    """A pushed record does not fit its table's definition."""


def download(cache: Path | None = None, http: httpx.Client | None = None) -> Path:
    """The cache file, fetched only when Raidbots' copy changed since the one kept."""
    local = os.environ.get(CACHE_FILE_ENV, "").strip()
    if local:
        return Path(local)
    path = (cache or cache_root()) / CACHE_NAME
    meta_path = path.with_suffix(".json")
    meta = json.loads(meta_path.read_text()) if meta_path.exists() and path.exists() else {}
    headers = {"User-Agent": USER_AGENT}
    if meta.get("etag"):
        headers["If-None-Match"] = meta["etag"]
    if meta.get("last_modified"):
        headers["If-Modified-Since"] = meta["last_modified"]
    own = http is None
    http = http or httpx.Client(timeout=120.0, follow_redirects=True)
    try:
        response = http.get(RAIDBOTS_URL, headers=headers)
    finally:
        if own:
            http.close()
    if response.status_code == 304:
        logger.info("raidbots DBCache.bin unchanged since %s", meta.get("last_modified"))
        return path
    response.raise_for_status()
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_bytes(response.content)
    meta_path.write_text(
        json.dumps(
            {
                "etag": response.headers.get("etag", ""),
                "last_modified": response.headers.get("last-modified", ""),
                "fetched": formatdate(usegmt=True),
                "bytes": len(response.content),
            },
            indent=1,
        )
    )
    logger.info("raidbots DBCache.bin: %d bytes", len(response.content))
    return path


def decode_record(data: bytes, record_id: int, fields: list[Field]) -> list[str]:
    """A pushed row as CSV cells in the table's column order."""
    cursor = _BitCursor(data + bytes(8))
    cells: list[str] = []
    for field in fields:
        if field.is_id and not field.inline:
            cells.append(str(record_id))
            continue
        for _ in range(field.count):
            if field.kind in ("string", "locstring"):
                cells.append(cursor.read_cstring())
            else:
                bits = field.bits if field.kind == "int" else 32
                cells.append(_cell(_typed(cursor.read_uint(bits), field)))
    if cursor.bit_pos // 8 != len(data):
        raise HotfixDecodeError(
            f"record {record_id}: {len(data)} bytes, the definition reads {cursor.bit_pos // 8}"
        )
    return cells


@dataclass(frozen=True)
class TableSchema:
    table_hash: int
    fields: list[Field]


def table_schemas(
    build: ClientBuild, tables: list[str], http: httpx.Client
) -> dict[str, TableSchema]:
    """Each table's hash and definition, from the build's own DB2 header."""
    schemas = {}
    for table in tables:
        try:
            data = build.read(FILE_DATA_IDS[table])
        except CdnMissing:
            continue  # an optional table this client does not ship
        if _read_header(data).record_count == 0:
            continue
        header = _read_header(data)
        text = definition(table, build.client, http)
        fields = read_definition(text, layout_hash(data), build.configs.version)
        schemas[table] = TableSchema(header.table_hash, fields)
    return schemas


def push_entries(
    records: list[HotfixRecord], schemas: Mapping[str, TableSchema]
) -> dict[str, list[PushEntry]]:
    """The cache's records per table as `hotfix_push.PushEntry`s (entry id = file order)."""
    by_hash = {schema.table_hash: table for table, schema in schemas.items()}
    entries: dict[str, list[PushEntry]] = defaultdict(list)
    for order, record in enumerate(records):
        table = by_hash.get(record.table_hash)
        if table is None:
            continue
        data: list[Any] | None = None
        if record.status == HotfixStatus.VALID and record.data:
            data = decode_record(record.data, record.record_id, schemas[table].fields)
        entries[table].append(
            PushEntry(order, record.push_id, record.record_id, int(record.status), data)
        )
    return dict(entries)


def _write_item_layer(
    raw: Path, build: str, entries: Mapping[str, list[PushEntry]]
) -> dict[str, int]:
    """The item pair's valid pushed rows into `raw/wago-hotfixes/`, the build in
    `_HotfixBuild`, for `pipeline.hotfix_merge`."""
    out = raw / WAGO_HOTFIX_DIR
    out.mkdir(parents=True, exist_ok=True)
    counts = {}
    for table in ITEM_TABLES:
        columns = _header(raw / f"{table}.csv")
        rows = [
            e.data
            for _, e in sorted(newest_per_record(entries.get(table, [])).items())
            if e.status == HotfixStatus.VALID and e.data
        ]
        with (out / f"{table}.csv").open("w", newline="", encoding="utf-8") as f:
            writer = csv.writer(f, lineterminator="\n")
            writer.writerow([*columns, HOTFIX_BUILD_COLUMN])
            writer.writerows([*row, build] for row in rows)
        counts[table] = len(rows)
    return counts


def _header(path: Path) -> list[str]:
    with path.open(newline="", encoding="utf-8") as f:
        return next(csv.reader(f))


def save_push(
    version: str,
    root: Path = Path("builds"),
    cache_file: Path | None = None,
    client: CdnClient | None = None,
    require_build_match: bool = True,
) -> Path | None:
    """Decode the cache's push for `version` into the committed
    `raw/raidbots/push.json` (the deciding entry per record, every table the
    fetch reads); None, and nothing written, when the cache is for another build."""
    path = cache_file or download()
    header, records = read_cache(path)
    build_number = int(version.rsplit(".", 1)[-1])
    if header.build != build_number and require_build_match:
        logger.warning(
            "raidbots DBCache.bin is for build %d, not %s; its hotfixes are not used",
            header.build,
            version,
        )
        return None
    raw = root / version / "raw"
    build = open_build(version, client)
    http = httpx.Client(headers={"User-Agent": USER_AGENT}, timeout=60.0)
    try:
        schemas = table_schemas(build, [*PUSH_TABLES, *ITEM_TABLES], http)
    finally:
        http.close()
    present = sorted(schemas)
    entries = push_entries(records, schemas)
    store = raw / STORE_DIR / STORE
    write_store(
        store,
        version,
        present,
        {},
        {t: deciding_entries(entries.get(t, [])) for t in present},
    )
    logger.info(
        "raidbots %d: %d pushed records saved to %s",
        header.build,
        sum(len(v) for v in entries.values()),
        store,
    )
    return store


def layer(version: str, root: Path = Path("builds")) -> dict[str, Any] | None:
    """Layer the saved `raw/raidbots/push.json` into `root/<version>/raw/`: the
    spell-side tables in place (`hotfix_push.layer_hotfix_push`), the item pair
    into `raw/wago-hotfixes/`. None when the build has no saved push."""
    raw = root / version / "raw"
    saved = read_store(raw / STORE_DIR / STORE)
    if saved is None:
        return None
    entries: dict[str, list[PushEntry]] = saved["entries"]
    items = _write_item_layer(raw, version, entries)
    push = {t: entries.get(t, []) for t in PUSH_TABLES}
    summary = layer_hotfix_push(
        version,
        root,
        fetch=lambda _build, tables: {t: push[t] for t in tables},
        totals=lambda _build, tables: {t: len(push[t]) for t in tables},
    )
    logger.info("raidbots push for %s: item hotfix rows %s", version, items)
    return {**summary, "item_hotfixes": items}


def push_moved(version: str, root: Path = Path("builds")) -> bool:
    """`pipeline push-check`: whether Raidbots' cache, read now, holds a push for
    `version` other than the one committed in `raw/raidbots/push.json`."""
    import tempfile

    committed = root / version / "raw" / STORE_DIR / STORE
    with tempfile.TemporaryDirectory() as scratch:
        fresh = save_push(version, Path(scratch))
        if fresh is None:
            return False
        return not committed.exists() or fresh.read_bytes() != committed.read_bytes()
