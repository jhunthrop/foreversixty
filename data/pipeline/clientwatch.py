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
silently missing a renumbering; so is a 1.60+ build under an ignored product (the live
client landing on a key the config calls not-Forever). Both lists keep only builds wago
ingested after the active build, so the old 3.4-5.5 betas still listed under
`wow_classic_beta` are not reported every night. A Forever build ingested after the
active one whose number is not above it (a renumbering down) is reported as
`out_of_order`, because `newer` alone would read it as "nothing to do".

Pure functions over the builds JSON, so the tests never touch the network; the CLI in
`__main__.py` fetches `/api/builds` and reads `web/src/data/active-build.json`.
"""

from __future__ import annotations

import json
import re
from collections.abc import Iterable
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


BLIZZARD_VERSIONS_URL = "https://us.version.battle.net/{product}/versions"


@dataclass(frozen=True)
class ClientConfig:
    products: frozenset[str]
    ignore_products: frozenset[str]
    min_version: tuple[int, ...]
    blizzard_products: tuple[str, ...] = ()


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
    #: The newest client Blizzard's own version service serves for the Forever products,
    #: or None when it was not asked. wago.tools exports a build's tables hours to days
    #: after Blizzard ships it, and the fetch needs those exports, so this is the early
    #: warning, not the trigger.
    blizzard: Build | None = None
    #: Forever builds wago ingested after the active build whose version is not above it.
    out_of_order: tuple[Build, ...] = ()

    @property
    def newer(self) -> bool:
        return self.newest is not None and newer_than(self.newest.version, self.active)

    @property
    def wago_behind(self) -> bool:
        """Blizzard serves a client newer than anything wago.tools has exported."""
        if self.blizzard is None:
            return False
        if self.newest is None:
            return True
        return newer_than(self.blizzard.version, self.newest.version)


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
        blizzard_products=tuple(raw.get("blizzard_products", ())),
    )


def read_active_build(path: Path = ACTIVE_BUILD_PATH) -> str:
    build = json.loads(path.read_text(encoding="utf-8")).get("build")
    if not build:
        raise ClientWatchError(f"{path} names no build")
    return build


def in_forever_family(build: Build, config: ClientConfig) -> bool:
    """A 1.x version at or above `min_version`, whatever the product."""
    try:
        version = version_tuple(build.version)
    except ClientWatchError:
        return False
    return version[0] == 1 and version >= config.min_version


def is_forever(build: Build, config: ClientConfig) -> bool:
    if build.product in config.products:
        return True
    if build.product in config.ignore_products:
        return False
    return in_forever_family(build, config)


def first_ingested(builds: list[Build], version: str) -> str:
    """When wago first listed `version` under any product ("" when it does not list it)."""
    stamps = [b.created_at for b in builds if b.version == version and b.created_at]
    return min(stamps) if stamps else ""


def is_unclaimed(build: Build, config: ClientConfig) -> bool:
    """Not Forever by rule, yet worth a warning: an unknown product, or an ignored one
    carrying a Forever-family version."""
    if is_forever(build, config):
        return False
    return build.product not in config.ignore_products or in_forever_family(build, config)


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


def parse_blizzard_versions(product: str, text: str) -> Build | None:
    """The `us` row's VersionsName from a Blizzard `/versions` answer.

    The file is pipe-separated with a typed header line (`Region!STRING:0|...|
    BuildId!DEC:4|VersionsName!String:0|...`) and `## seqn` comment lines.
    """
    lines = [ln for ln in text.splitlines() if ln and not ln.startswith("##")]
    if not lines:
        return None
    columns = [col.split("!")[0] for col in lines[0].split("|")]
    for line in lines[1:]:
        row = dict(zip(columns, line.split("|"), strict=False))
        if row.get("Region") == "us" and _VERSION.match(row.get("VersionsName", "")):
            return Build(product, row["VersionsName"], "")
    return None


def fetch_blizzard_versions(
    products: tuple[str, ...], client: httpx.Client | None = None
) -> dict[str, str]:
    """product -> the raw `/versions` text, for each product Blizzard answers for."""
    own = client is None
    client = client or httpx.Client(headers={"User-Agent": USER_AGENT}, timeout=30)
    out: dict[str, str] = {}
    try:
        for product in products:
            r = client.get(BLIZZARD_VERSIONS_URL.format(product=product))
            if r.status_code == 200 and r.text.startswith("Region"):
                out[product] = r.text
    finally:
        if own:
            client.close()
    return out


def newest_blizzard(versions: dict[str, str]) -> Build | None:
    builds = [b for p, t in versions.items() if (b := parse_blizzard_versions(p, t))]
    return max(builds, key=lambda b: b.key) if builds else None


def check(
    builds_json: dict[str, Any],
    active: str,
    config: ClientConfig,
    blizzard_versions: dict[str, str] | None = None,
) -> Check:
    builds = all_builds(builds_json)
    newest = newest_forever(builds, config)
    blizzard = newest_blizzard(blizzard_versions) if blizzard_versions else None
    since = first_ingested(builds, active)
    later = [b for b in builds if not since or b.created_at > since]
    unmatched = _newest_first(
        b for b in later if is_unclaimed(b, config) and newer_than(b.version, active)
    )
    out_of_order = _newest_first(
        b
        for b in later
        if since
        and is_forever(b, config)
        and b.version != active
        and not newer_than(b.version, active)
    )
    return Check(
        active=active,
        newest=newest,
        unmatched_newer=unmatched,
        blizzard=blizzard,
        out_of_order=out_of_order,
    )


def _newest_first(builds: Iterable[Build]) -> tuple[Build, ...]:
    return tuple(sorted(builds, key=lambda b: b.key, reverse=True))


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
        "blizzard_build": result.blizzard.version if result.blizzard else "",
        "blizzard_product": result.blizzard.product if result.blizzard else "",
        "wago_behind": result.wago_behind,
        "unmatched_newer": [_row(b) for b in result.unmatched_newer],
        "out_of_order": [_row(b) for b in result.out_of_order],
    }


def _row(build: Build) -> dict[str, str]:
    return {"product": build.product, "version": build.version, "created_at": build.created_at}
