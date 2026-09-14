import json
import logging
from datetime import UTC, datetime
from pathlib import Path

import httpx

logger = logging.getLogger(__name__)

BASE_URL = "https://wago.tools"
USER_AGENT = "foreversixty-pipeline/0.1 (+https://foreversixty.gg)"
TABLES = [
    "Map",
    "AreaTable",
    "JournalInstance",
    "ItemSparse",
    "Item",
    "SpellName",
    "Spell",
    "SpellEffect",
    "SpellDuration",
    "SpellMisc",
    "ManifestInterfaceData",
    "ItemSet",
    "ItemSetSpell",
    "ChrClasses",
    "ChrRaces",
    "Talent",
    "TalentTab",
]
# Tables allowed to be missing (404) for a given product/build without aborting the
# fetch. JournalInstance (the Dungeon Journal) predates Classic Era's client, so it
# doesn't exist there. Any other 404 is treated as a real failure (e.g. a typo).
OPTIONAL_TABLES = frozenset({"JournalInstance"})


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


def download_table(table: str, build: str, dest: Path, client: httpx.Client) -> Path:
    r = client.get(f"/db2/{table}/csv", params={"build": build}, timeout=120)
    dest.mkdir(parents=True, exist_ok=True)
    path = dest / f"{table}.csv"
    if r.status_code == 404:
        if table not in OPTIONAL_TABLES:
            raise SystemExit(f"{table} not found for build {build}; check TABLES for a typo")
        # Allowlisted: this table is known to be absent for some products/builds.
        # Treat it as empty rather than aborting the whole fetch.
        logger.warning("table %s not found for build %s; writing empty table", table, build)
        path.write_text("ID\n", encoding="utf-8")
        return path
    r.raise_for_status()
    header = r.text.split("\n", 1)[0]
    if "ID" not in header.split(","):
        raise SystemExit(
            f"{table} for build {build} did not return a CSV with an ID column; "
            f"got columns: {header!r}"
        )
    path.write_text(r.text, encoding="utf-8")
    return path


def fetch_build(
    product: str,
    build: str | None,
    root: Path = Path("builds"),
    client: httpx.Client | None = None,
) -> Path:
    own = client is None
    client = client or httpx.Client(
        base_url=BASE_URL,
        headers={"User-Agent": USER_AGENT},
    )
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
