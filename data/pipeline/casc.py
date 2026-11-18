"""One way to read a file out of the client's CASC, pinned to a build.

The bytes come from Blizzard's CDN (`pipeline.blizzard_cdn`), the build's own
root and encoding files resolving the file data id; wago.tools' `/api/casc`
served them until it refused this pipeline on 2026-11-17. The same id is not
the same bytes in two builds (`combatratings.txt` is 12,144 bytes on
1.60.1.69893 and 31,614 on 1.15.9.69722), so the cache is keyed by build as
well as by id, and a file the build does not ship raises `CascMissing`.
"""

from __future__ import annotations

import logging
from enum import Enum
from pathlib import Path
from typing import Protocol

from pipeline.blizzard_cdn import CdnMissing
from pipeline.icons import CACHE_DIR, _atomic_write

logger = logging.getLogger(__name__)


class CascError(ValueError):
    """CASC did not return a usable file for this id and build."""


class CascMissingReason(Enum):
    """Why a `CascMissing` was raised, as data rather than as prose.

    A caller that needs to tell the two apart (`pipeline/icons.py` logs a
    different warning for each) branches on this, not on `str(error)` --
    matching on message text would silently misroute the moment either
    message's wording changed.
    """

    NOT_FOUND = "not_found"
    EMPTY = "empty"


class CascMissing(CascError):
    """This build does not ship this file: a 404, or a 200 with no body.

    Its own type because the three callers differ on it. `pipeline/icons.py`
    warns and skips -- an icon the client has no art for is one item drawn
    with the placeholder, not a failed build. `pipeline/art.py` turns it into
    the ArtDataError its tests expect. `pipeline/gametables.py` lets it
    through, because a stat curve the engine needs and the client does not
    have is the whole point of failing loudly.
    """

    def __init__(self, message: str, *, reason: CascMissingReason) -> None:
        super().__init__(message)
        self.reason = reason


class FileSource(Protocol):
    """Where a client file's bytes come from: `pipeline.blizzard_cdn.ClientBuild`."""

    def read(self, file_id: int) -> bytes: ...


_SOURCES: dict[str, FileSource] = {}


def build_source(build: str) -> FileSource:
    """The build's files on Blizzard's CDN, opened once per process."""
    if build not in _SOURCES:
        from pipeline.blizzard_cdn import open_build

        _SOURCES[build] = open_build(build)
    return _SOURCES[build]


def fetch_casc_file(
    file_id: int,
    build: str,
    cache_dir: Path = CACHE_DIR,
    source: FileSource | None = None,
    suffix: str = ".bin",
) -> bytes:
    """One file out of CASC, for one build, cached under `<cache_dir>/<build>/`."""
    cached = cache_dir / build / f"{file_id}{suffix}"
    if cached.exists():
        return cached.read_bytes()
    try:
        data = (source or build_source(build)).read(file_id)
    except CdnMissing as error:
        raise CascMissing(
            f"file data id {file_id} is not in CASC at version {build}",
            reason=CascMissingReason.NOT_FOUND,
        ) from error
    if not data:
        raise CascMissing(
            f"file data id {file_id} is empty for build {build}; this build's "
            f"client does not ship that file",
            reason=CascMissingReason.EMPTY,
        )
    cached.parent.mkdir(parents=True, exist_ok=True)
    _atomic_write(cached, data)
    logger.info("casc %s @ %s: %d bytes", file_id, build, len(data))
    return data
