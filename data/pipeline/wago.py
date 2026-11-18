import json
import logging
from datetime import UTC, datetime
from pathlib import Path

import httpx

from pipeline.client_tables import OPTIONAL_TABLES, TABLES
from pipeline.wago_http import open_client
from pipeline.wago_tables import (
    MISSING_TABLE_STATUSES,
    TableCache,
    get_table,
    table_cache_from_env,
)

logger = logging.getLogger(__name__)


def latest_build(product: str, client: httpx.Client) -> str:
    r = client.get("/api/builds")
    r.raise_for_status()
    builds = r.json().get(product)
    if not builds:
        known = sorted(r.json())
        raise SystemExit(
            f"wago.tools lists no builds for product {product!r}; known products: {known}"
        )
    return builds[0]["version"]


def download_table(
    table: str,
    build: str,
    dest: Path,
    client: httpx.Client,
    cache: TableCache | None = None,
) -> Path:
    """`table`'s shipped export for `build` into `dest/<table>.csv`, through `cache`
    (WAGO_TABLE_CACHE when not given): a build's shipped tables never change, so a
    cached one is never downloaded again. A missing table (MISSING_TABLE_STATUSES)
    is tolerated only for OPTIONAL_TABLES."""
    answer = get_table(client, table, build, cache=cache or table_cache_from_env())
    dest.mkdir(parents=True, exist_ok=True)
    path = dest / f"{table}.csv"
    if answer.status in MISSING_TABLE_STATUSES:
        if table not in OPTIONAL_TABLES:
            raise SystemExit(f"{table} not found for build {build}; check TABLES for a typo")
        # Allowlisted: this table is known to be absent for some products/builds.
        # Treat it as empty rather than aborting the whole fetch.
        logger.warning("table %s not found for build %s; writing empty table", table, build)
        path.write_text("ID\n", encoding="utf-8")
        return path
    header = answer.text.split("\n", 1)[0]
    if "ID" not in header.split(","):
        raise SystemExit(
            f"{table} for build {build} did not return a CSV with an ID column; "
            f"got columns: {header!r}"
        )
    path.write_text(answer.text, encoding="utf-8")
    return path


def fetch_build(
    product: str,
    build: str | None,
    root: Path = Path("builds"),
    client: httpx.Client | None = None,
) -> Path:
    own = client is None
    client = client or open_client()
    try:
        build = build or latest_build(product, client)
        raw = root / build / "raw"
        for table in TABLES:
            download_table(table, build, raw, client)
        meta = {
            "product": product,
            "build": build,
            "fetched_at": datetime.now(UTC).isoformat().replace("+00:00", "Z"),
        }
        (raw / "_meta.json").write_text(json.dumps(meta, indent=2) + "\n")
        print(build)
        return root / build
    finally:
        if own:
            client.close()
