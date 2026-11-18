"""A build's client tables, read off Blizzard's CDN into `raw/<Table>.csv`.

The replacement for wago.tools' CSV exports (`pipeline.wago`): each DB2 is
read by its file data id (`pipeline.client_tables.FILE_DATA_IDS`) through
`pipeline.blizzard_cdn`, decoded by `pipeline.wdc5` against its pinned
WoWDBDefs definition (`pipeline.dbd`), and written in the same CSV shape the
rest of the pipeline has always read.
"""

from __future__ import annotations

import json
import logging
from collections.abc import Iterable
from datetime import UTC, datetime
from pathlib import Path

import httpx

from pipeline.blizzard_cdn import CdnClient, CdnMissing, ClientBuild, open_build
from pipeline.client_tables import FILE_DATA_IDS, OPTIONAL_TABLES, TABLES
from pipeline.dbd import DBD_COMMIT, fetch_definition, read_definition
from pipeline.wago_http import USER_AGENT
from pipeline.wdc5 import layout_hash, table_csv

logger = logging.getLogger(__name__)

SOURCE = "blizzard-cdn"


def definition(table: str, client: CdnClient, http: httpx.Client) -> str:
    """`table`'s pinned .dbd text, from the CDN cache when it has it."""
    path = client.cache / "dbd" / DBD_COMMIT / f"{table}.dbd"
    if path.exists():
        return path.read_text(encoding="utf-8")
    text = fetch_definition(table, http)
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text, encoding="utf-8")
    return text


def table_text(table: str, build: ClientBuild, http: httpx.Client) -> str | None:
    """`table` as CSV text for the build, or None when the client does not ship it."""
    file_id = FILE_DATA_IDS.get(table)
    if file_id is None:
        raise SystemExit(f"{table} has no file data id in pipeline.client_tables.FILE_DATA_IDS")
    try:
        data = build.read(file_id)
    except CdnMissing:
        return None
    fields = read_definition(
        definition(table, build.client, http), layout_hash(data), build.configs.version
    )
    return table_csv(data, fields)


def write_table(table: str, build: ClientBuild, dest: Path, http: httpx.Client) -> Path:
    """`dest/<table>.csv`; a table the build lacks is an empty table when optional."""
    text = table_text(table, build, http)
    dest.mkdir(parents=True, exist_ok=True)
    path = dest / f"{table}.csv"
    if text is None:
        if table not in OPTIONAL_TABLES:
            raise SystemExit(f"{table} is not in {build.configs.version}'s client")
        logger.warning("table %s not in %s; writing an empty table", table, build.configs.version)
        text = "ID\n"
    path.write_text(text, encoding="utf-8")
    return path


def fetch_tables(version: str, tables: Iterable[str], dest: Path) -> list[Path]:
    """`tables` for one build into `dest/<Table>.csv` (a step that reads a few
    tables of its own: guild events, raid bosses, PlayerExpectedStat)."""
    client = CdnClient()
    http = httpx.Client(headers={"User-Agent": USER_AGENT}, timeout=60.0)
    try:
        build = open_build(version, client)
        return [write_table(table, build, dest, http) for table in tables]
    finally:
        http.close()
        client.close()


def fetch_build(product: str, version: str | None, root: Path = Path("builds")) -> Path:
    """Every table in TABLES for the build into `root/<build>/raw/`; prints the build."""
    client = CdnClient(product=product)
    http = httpx.Client(headers={"User-Agent": USER_AGENT}, timeout=60.0)
    try:
        build = open_build(version, client)
        build.resolve(FILE_DATA_IDS[t] for t in TABLES)
        raw = root / build.configs.version / "raw"
        for table in TABLES:
            write_table(table, build, raw, http)
        meta = {
            "product": product,
            "build": build.configs.version,
            "fetched_at": datetime.now(UTC).isoformat().replace("+00:00", "Z"),
            "source": SOURCE,
            "build_config": build.configs.build_config,
            "cdn_config": build.configs.cdn_config,
            "dbd_commit": DBD_COMMIT,
        }
        (raw / "_meta.json").write_text(json.dumps(meta, indent=2) + "\n")
        logger.info(
            "cdn tables %s: %d requests, %.1f MB downloaded",
            build.configs.version,
            client.requests,
            client.downloaded / 1e6,
        )
        print(build.configs.version)
        return root / build.configs.version
    finally:
        http.close()
        client.close()
