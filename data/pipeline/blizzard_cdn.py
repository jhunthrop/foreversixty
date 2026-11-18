"""Client files straight from Blizzard's public CDN (TACT), the launcher's own source.

wago.tools refused this pipeline on 2026-11-17, so the client's DB2 tables and
single files (icons, GameTables, talent-tree art) now come from where wago got
them: `{region}.version.battle.net/{product}/versions` names a build's build
config and CDN config, and the CDN serves every file by content address.

Reading one file by its file data id:

1. the build config names the root and encoding files;
2. the root maps the file data id to a content key, the encoding file maps
   that to an encoding key;
3. one of the CDN config's archive indexes says which archive holds it and
   where, and a single ranged GET fetches it (a file in no archive is fetched
   loose by its encoding key).

Being polite: one request at a time, an honest User-Agent, retries with
pauses on a dropped connection or a 5xx, and a disk cache (FOREVER_CDN_CACHE,
default `.cdn-cache`) that the workflows keep in actions/cache:

* configs, archive indexes and fetched files are content-addressed and never
  change, so each is downloaded once, ever, across builds;
* a build's encoding and root (about 150 MB together) live under
  `builds/<build config>/` with a small `resolved.json` of the file data ids
  already looked up, so a second run of the same build reads neither;
  `prune` keeps only the builds a run names, so the cache does not grow.

Accuracy: these are exactly the shipped files. Server hotfixes are another
source (`pipeline.raidbots_hotfixes`).
"""

from __future__ import annotations

import json
import logging
import os
import shutil
import time
from collections.abc import Iterable
from dataclasses import dataclass, field
from pathlib import Path

import httpx

from pipeline.tact import (
    ArchiveIndex,
    EncodingFile,
    blte_decode,
    parse_config,
    parse_keys,
    parse_psv,
    root_content_keys,
)
from pipeline.wago_http import USER_AGENT

logger = logging.getLogger(__name__)

PRODUCT = "wowf"
REGION = "us"
VERSION_SERVICE = "https://{region}.version.battle.net/{product}/{kind}"
CACHE_ENV = "FOREVER_CDN_CACHE"
DEFAULT_CACHE = Path(".cdn-cache")
MIN_INTERVAL_ENV = "CDN_MIN_INTERVAL"
#: Seconds from the start of one CDN request to the start of the next.
DEFAULT_MIN_INTERVAL_S = 0.05
#: Pauses before retrying a dropped connection or a 5xx; then the next host.
RETRY_PAUSES_S = (2.0, 5.0, 15.0)
TIMEOUT_S = 120.0
#: Build and CDN configs for builds the version service no longer names in any
#: region. A build it does name needs no entry.
#: wowdev/TACTKeys, the community's list of the encryption keys Blizzard has
#: released (each arrives in a hotfix push when its content goes live). A DB2
#: section under a key in it decrypts, as wago.tools' exports did; one under a
#: key not yet released stays unread. Pinned for reproducible runs; bump it
#: when content behind a newly released key should appear.
TACT_KEYS_COMMIT = "1f8ba0ff780db15b1f749a4e36b8c4e2d73f44a4"
TACT_KEYS_URL = "https://raw.githubusercontent.com/wowdev/TACTKeys/{commit}/WoW.txt"
PINS_PATH = Path(__file__).resolve().parents[1] / "curated" / "cdn-builds.json"


class CdnError(RuntimeError):
    """The CDN did not give us a file it should have."""


class CdnMissing(CdnError):
    """Every CDN host answered 404, or the build has no such file data id."""


@dataclass(frozen=True)
class BuildConfigs:
    version: str
    build_config: str
    cdn_config: str


def cache_root() -> Path:
    value = os.environ.get(CACHE_ENV, "").strip()
    return Path(value) if value else DEFAULT_CACHE


def _hashed(kind: str, key: str) -> str:
    return f"{kind}/{key[0:2]}/{key[2:4]}/{key}"


def _atomic_write(path: Path, data: bytes) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    tmp = path.with_suffix(path.suffix + ".part")
    tmp.write_bytes(data)
    tmp.replace(path)


@dataclass
class CdnClient:
    """Sequential, cached GETs against the version service and the CDN hosts."""

    cache: Path = field(default_factory=cache_root)
    product: str = PRODUCT
    region: str = REGION
    http: httpx.Client | None = None
    min_interval: float = DEFAULT_MIN_INTERVAL_S
    sleep: object = time.sleep
    _hosts: list[str] | None = None
    _cdn_path: str = "tpr/wow"
    _last_start: float = 0.0
    _keys: dict[str, bytes] | None = None
    requests: int = 0
    downloaded: int = 0

    def __post_init__(self) -> None:
        if self.http is None:
            self.http = httpx.Client(
                headers={"User-Agent": USER_AGENT}, timeout=TIMEOUT_S, follow_redirects=True
            )
        raw = os.environ.get(MIN_INTERVAL_ENV, "").strip()
        if raw:
            self.min_interval = float(raw)

    def close(self) -> None:
        if self.http is not None:
            self.http.close()

    def _send(self, url: str, headers: dict[str, str] | None = None) -> httpx.Response:
        wait = self._last_start + self.min_interval - time.monotonic()
        if wait > 0:
            self.sleep(wait)  # type: ignore[operator]
        self._last_start = time.monotonic()
        self.requests += 1
        assert self.http is not None
        return self.http.get(url, headers=headers or {})

    def _get(self, url: str, headers: dict[str, str] | None = None) -> httpx.Response:
        for attempt in range(len(RETRY_PAUSES_S) + 1):
            try:
                response = self._send(url, headers)
            except httpx.TransportError as error:
                if attempt == len(RETRY_PAUSES_S):
                    raise CdnError(f"{url}: {error}") from error
                logger.warning("cdn %s: %s; retrying", url, error)
                self.sleep(RETRY_PAUSES_S[attempt])  # type: ignore[operator]
                continue
            if response.status_code >= 500 and attempt < len(RETRY_PAUSES_S):
                logger.warning("cdn %s answered %d; retrying", url, response.status_code)
                self.sleep(RETRY_PAUSES_S[attempt])  # type: ignore[operator]
                continue
            return response
        raise AssertionError("unreachable")

    def version_table(self, kind: str) -> list[dict[str, str]]:
        """The version service's `versions` or `cdns` table for the product."""
        url = VERSION_SERVICE.format(region=self.region, product=self.product, kind=kind)
        response = self._get(url)
        response.raise_for_status()
        return parse_psv(response.text)

    def hosts(self) -> list[str]:
        """HTTPS CDN base URLs for the region, the version service's order."""
        if self._hosts is None:
            rows = self.version_table("cdns")
            row = next((r for r in rows if r.get("Name") == self.region), rows[0])
            self._cdn_path = row["Path"]
            servers = [s.split("?", 1)[0].rstrip("/") for s in row.get("Servers", "").split()]
            secure = [s for s in servers if s.startswith("https://")]
            self._hosts = secure or [f"https://{h}" for h in row["Hosts"].split()]
        return self._hosts

    def _cdn_get(self, path: str, byte_range: tuple[int, int] | None = None) -> bytes:
        headers = {}
        if byte_range is not None:
            start, size = byte_range
            headers["Range"] = f"bytes={start}-{start + size - 1}"
        statuses: list[object] = []
        for host in self.hosts():
            try:
                response = self._get(f"{host}/{self._cdn_path}/{path}", headers)
            except CdnError as error:  # this host is unreachable; the next may answer
                statuses.append(str(error))
                continue
            if response.status_code in (200, 206):
                body = response.content
                if byte_range is not None and len(body) != byte_range[1]:
                    raise CdnError(f"{path}: range gave {len(body)} bytes, not {byte_range[1]}")
                self.downloaded += len(body)
                return body
            statuses.append(response.status_code)
        if all(status == 404 for status in statuses):
            raise CdnMissing(f"{path} is on no CDN host (404)")
        raise CdnError(f"{path}: CDN hosts answered {statuses}")

    def _cached(self, rel: str, fetch: object) -> bytes:
        path = self.cache / rel
        if path.exists():
            return path.read_bytes()
        data = fetch()  # type: ignore[operator]
        _atomic_write(path, data)
        return data

    def tact_keys(self) -> dict[str, bytes]:
        """The released TACT keys at TACT_KEYS_COMMIT, cached."""
        if self._keys is None:
            rel = f"keys/{TACT_KEYS_COMMIT}/WoW.txt"

            def fetch() -> bytes:
                response = self._get(TACT_KEYS_URL.format(commit=TACT_KEYS_COMMIT))
                response.raise_for_status()
                return response.content

            self._keys = parse_keys(self._cached(rel, fetch).decode("utf-8"))
        return self._keys

    def config(self, key: str) -> dict[str, list[str]]:
        rel = _hashed("config", key)
        data = self._cached(rel, lambda: self._cdn_get(rel))
        return parse_config(data.decode("utf-8"))

    def archive_index(self, archive: str) -> bytes:
        rel = _hashed("data", archive) + ".index"
        return self._cached(f"indices/{archive}.index", lambda: self._cdn_get(rel))

    def archived(self, ekey: str, archive: str, offset: int, size: int) -> bytes:
        rel = _hashed("data", archive)
        return self._cached(f"files/{ekey[:2]}/{ekey}", lambda: self._cdn_get(rel, (offset, size)))

    def loose(self, ekey: str, cache_as: str | None = None) -> bytes:
        rel = _hashed("data", ekey)
        return self._cached(cache_as or f"files/{ekey[:2]}/{ekey}", lambda: self._cdn_get(rel))


def load_pins(path: Path = PINS_PATH) -> dict[str, dict[str, str]]:
    if not path.exists():
        return {}
    raw = json.loads(path.read_text(encoding="utf-8"))
    return {k: v for k, v in raw.items() if not k.startswith("_")}


def resolve_build(
    client: CdnClient, version: str | None, pins: dict[str, dict[str, str]] | None = None
) -> BuildConfigs:
    """`version`'s configs (the newest build in the region when None): from the
    version service when any region still names it, else from the pins."""
    rows = client.version_table("versions")
    if version is None:
        row = next((r for r in rows if r.get("Region") == client.region), rows[0])
        return BuildConfigs(row["VersionsName"], row["BuildConfig"], row["CDNConfig"])
    for row in rows:
        if row.get("VersionsName") == version:
            return BuildConfigs(version, row["BuildConfig"], row["CDNConfig"])
    pin = (load_pins() if pins is None else pins).get(version)
    if pin:
        return BuildConfigs(version, pin["build_config"], pin["cdn_config"])
    named = sorted({r.get("VersionsName", "") for r in rows})
    raise SystemExit(
        f"{client.product} {version} is not on the version service (it names {named}) and "
        f"{PINS_PATH.name} has no pin for it; add its build and CDN config hashes there"
    )


@dataclass
class ClientBuild:
    """One build's files by file data id, through the CDN cache."""

    client: CdnClient
    configs: BuildConfigs
    _encoding: EncodingFile | None = None
    _indexes: list[tuple[str, ArchiveIndex]] | None = None
    _resolved: dict[str, list | None] = field(default_factory=dict)

    def __post_init__(self) -> None:
        self.build_config = self.client.config(self.configs.build_config)
        self.cdn_config = self.client.config(self.configs.cdn_config)
        name = self.build_config.get("build-name", [""])[0]
        build_number = self.configs.version.rsplit(".", 1)[-1]
        if build_number not in name:
            raise SystemExit(
                f"build config {self.configs.build_config} is {name!r}, not {self.configs.version}"
            )
        if self._memo_path.exists():
            self._resolved = json.loads(self._memo_path.read_text(encoding="utf-8"))

    @property
    def _build_dir(self) -> Path:
        return self.client.cache / "builds" / self.configs.build_config

    @property
    def _memo_path(self) -> Path:
        return self._build_dir / "resolved.json"

    def _encoding_file(self) -> EncodingFile:
        if self._encoding is None:
            _, ekey = self.build_config["encoding"]
            data = self.client.loose(ekey, cache_as=f"builds/{self.configs.build_config}/encoding")
            self._encoding = EncodingFile.parse(blte_decode(data))
        return self._encoding

    def _ekey(self, ckey: bytes) -> bytes:
        ekey = self._encoding_file().encoding_key(ckey)
        if ekey is None:
            raise CdnError(f"content key {ckey.hex()} is not in the encoding file")
        return ekey

    def _archive_indexes(self) -> list[tuple[str, ArchiveIndex]]:
        if self._indexes is None:
            archives = self.cdn_config["archives"]
            logger.info("cdn: reading %d archive indexes", len(archives))
            self._indexes = [
                (name, ArchiveIndex.parse(self.client.archive_index(name))) for name in archives
            ]
        return self._indexes

    def _locate(self, ekey: bytes) -> list:
        for name, index in self._archive_indexes():
            hit = index.find(ekey)
            if hit is not None:
                return [ekey.hex(), name, hit[0], hit[1]]
        return [ekey.hex(), "", 0, 0]

    def _root(self, wanted: set[int]) -> dict[int, bytes]:
        root_ckey = bytes.fromhex(self.build_config["root"][0])
        ekey = self._ekey(root_ckey).hex()
        data = self.client.loose(ekey, cache_as=f"builds/{self.configs.build_config}/root")
        return root_content_keys(blte_decode(data), wanted)

    def resolve(self, file_ids: Iterable[int]) -> None:
        """Look up every id not already in `resolved.json` and save it."""
        todo = {fid for fid in file_ids if str(fid) not in self._resolved}
        if not todo:
            return
        ckeys = self._root(todo)
        for fid in sorted(todo):
            ckey = ckeys.get(fid)
            self._resolved[str(fid)] = None if ckey is None else self._locate(self._ekey(ckey))
        _atomic_write(self._memo_path, json.dumps(self._resolved, sort_keys=True).encode())

    def read(self, file_id: int) -> bytes:
        """The plain bytes of `file_id` in this build; CdnMissing when it has none."""
        self.resolve([file_id])
        entry = self._resolved[str(file_id)]
        if entry is None:
            raise CdnMissing(f"file data id {file_id} is not in {self.configs.version}'s root")
        ekey, archive, offset, size = entry
        if archive:
            data = self.client.archived(ekey, archive, offset, size)
        else:
            data = self.client.loose(ekey)
        return blte_decode(data, self.client.tact_keys())


def prune(cache: Path, keep_build_configs: Iterable[str]) -> list[str]:
    """Drop every cached build's encoding and root but those named. Answers the dropped."""
    keep = set(keep_build_configs)
    builds = cache / "builds"
    dropped = []
    if builds.is_dir():
        for path in builds.iterdir():
            if path.name not in keep:
                shutil.rmtree(path)
                dropped.append(path.name)
    return dropped


def open_build(version: str | None, client: CdnClient | None = None) -> ClientBuild:
    """`version` (newest when None) of the product, ready to read files."""
    client = client or CdnClient()
    return ClientBuild(client, resolve_build(client, version))
