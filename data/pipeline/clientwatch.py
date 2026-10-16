"""Which client build is the newest World of Warcraft: Forever build, and is it ours yet.

wago.tools lists every Blizzard client it has ingested, grouped by product key. Forever's
builds have already moved between keys once (`wow_classic_beta` in September 2026,
`wow_cn_beta` and `wowf` by October), and the version line may be renumbered when the
game goes live, so the pick is made two ways at once, from `curated/client.json`:

- a build under a product the file names is Forever, whatever its version;
- a build under any other product is Forever when its version is a 1.x at or above
  `min_version` (Classic Era stays on 1.15; Forever started at 1.60), unless the product
  is in `ignore_products`.

The newest matching build, by version then by wago's ingest time, is the candidate.
`newer_than` compares it with the active build; `interface_number` is the TOC
`## Interface:` value a client version implies. A build newer than the active one that
matches neither rule is returned separately so the workflow can warn about it instead of
silently missing a renumbering.

Pure functions over the builds JSON, so the tests never touch the network; the CLI in
`__main__.py` fetches `/api/builds` and reads `web/src/data/active-build.json`.
"""

from __future__ import annotations

import json
import re
from dataclasses import dataclass
from pathlib import Path
from typing import Any

import httpx

from pipeline.wago import BASE_URL, USER_AGENT

CONFIG_PATH = Path(__file__).resolve().parents[1] / "curated" / "client.json"
ACTIVE_BUILD_PATH = (
    Path(__file__).resolve().parents[2] / "web" / "src" / "data" / "active-build.json"
)

_VERSION = re.compile(r"^\d+(\.\d+){3}$")


@dataclass(frozen=True)
class ClientConfig:
    products: frozenset[str]
    ignore_products: frozenset[str]
    min_version: tuple[int, ...]


@dataclass(frozen=True)
class Build:
    product: str
    version: str
    created_at: str

    @property
    def key(self) -> tuple[tuple[int, ...], str]:
        return (version_tuple(self.version), self.created_at)


@dataclass(frozen=True)
class Check:
    """What `client-check` reports: the active build, the newest Forever build, and the
    builds newer than active that no rule claimed."""

    active: str
    newest: Build | None
    unmatched_newer: tuple[Build, ...]

    @property
    def newer(self) -> bool:
        return self.newest is not None and newer_than(self.newest.version, self.active)


class ClientWatchError(Exception):
    pass


def version_tuple(version: str) -> tuple[int, ...]:
    if not _VERSION.match(version):
        raise ClientWatchError(f"{version!r} is not a four-part client version")
    return tuple(int(part) for part in version.split("."))


def newer_than(candidate: str, active: str) -> bool:
    return version_tuple(candidate) > version_tuple(active)


def interface_number(version: str) -> int:
    """The `## Interface:` number a client version implies: 1.60.1 -> 16001, 1.15.9 -> 11509.

    Blizzard's TOC convention is major * 10000 + minor * 100 + patch; the fourth part (the
    build) never enters it.
    """
    major, minor, patch, _build = version_tuple(version)
    return major * 10000 + minor * 100 + patch


def load_config(path: Path = CONFIG_PATH) -> ClientConfig:
    raw = json.loads(path.read_text(encoding="utf-8"))
    return ClientConfig(
        products=frozenset(raw["products"]),
        ignore_products=frozenset(raw.get("ignore_products", ())),
        min_version=version_tuple(raw["min_version"]),
    )


def read_active_build(path: Path = ACTIVE_BUILD_PATH) -> str:
    build = json.loads(path.read_text(encoding="utf-8")).get("build")
    if not build:
        raise ClientWatchError(f"{path} names no build")
    return build


def is_forever(build: Build, config: ClientConfig) -> bool:
    if build.product in config.products:
        return True
    if build.product in config.ignore_products:
        return False
    try:
        version = version_tuple(build.version)
    except ClientWatchError:
        return False
    return version[0] == 1 and version >= config.min_version


def all_builds(builds_json: dict[str, Any]) -> list[Build]:
    """Every (product, version, created_at) wago lists, skipping rows without a version."""
    out: list[Build] = []
    for product, rows in builds_json.items():
        for row in rows or ():
            version = row.get("version")
            if not version or not _VERSION.match(version):
                continue
            out.append(Build(product, version, row.get("created_at") or ""))
    return out


def newest_forever(builds: list[Build], config: ClientConfig) -> Build | None:
    forever = [b for b in builds if is_forever(b, config)]
    return max(forever, key=lambda b: b.key) if forever else None


def check(builds_json: dict[str, Any], active: str, config: ClientConfig) -> Check:
    builds = all_builds(builds_json)
    newest = newest_forever(builds, config)
    unmatched = tuple(
        sorted(
            (
                b
                for b in builds
                if not is_forever(b, config)
                and b.product not in config.ignore_products
                and newer_than(b.version, active)
            ),
            key=lambda b: b.key,
            reverse=True,
        )
    )
    return Check(active=active, newest=newest, unmatched_newer=unmatched)


def fetch_builds(client: httpx.Client | None = None) -> dict[str, Any]:
    own = client is None
    client = client or httpx.Client(
        base_url=BASE_URL, headers={"User-Agent": USER_AGENT}, timeout=60
    )
    try:
        r = client.get("/api/builds")
        r.raise_for_status()
        return r.json()
    finally:
        if own:
            client.close()


def report(result: Check) -> dict[str, Any]:
    """The JSON the CLI prints and the workflow reads as step outputs."""
    newest = result.newest
    return {
        "active": result.active,
        "active_interface": interface_number(result.active),
        "newest_build": newest.version if newest else "",
        "newest_product": newest.product if newest else "",
        "newest_created_at": newest.created_at if newest else "",
        "newest_interface": interface_number(newest.version) if newest else 0,
        "newer": result.newer,
        "unmatched_newer": [
            {"product": b.product, "version": b.version, "created_at": b.created_at}
            for b in result.unmatched_newer
        ],
    }
