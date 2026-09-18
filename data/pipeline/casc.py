"""One way to read a file out of the client's CASC, pinned to a build.

wago.tools serves any client file by its file data id at
`/api/casc/<file data id>`, and `?version=<build>` chooses the build. It is
not optional. Measured on 2026-09-18 for `combatratings.txt`, file data id
1391669: `?version=1.60.1.69893` returns 12,144 bytes, `?version=1.15.9.69722`
returns 31,614, an unknown version answers 400 "Version not found", and a bare
request returns whichever build wago currently defaults to -- which is retail,
and which is how an earlier draft of `pipeline/gametables.py` came to read
retail's stat curves and label them Forever's.

So every caller passes the build, the cache is keyed by build as well as by
id (the same id is not the same bytes twice), and anything other than a 2xx
with a non-empty body raises: a 400 is a wrong build string, an empty body is
a file the build does not ship, and a redirect is neither.
"""

from __future__ import annotations

import logging
from enum import Enum
from pathlib import Path

import httpx

from pipeline.icons import CACHE_DIR, _atomic_write
from pipeline.wago import BASE_URL, USER_AGENT

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


def fetch_casc_file(
    file_id: int,
    build: str,
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
    suffix: str = ".bin",
) -> bytes:
    """One file out of CASC, for one build, cached under `<cache_dir>/<build>/`."""
    cached = cache_dir / build / f"{file_id}{suffix}"
    if cached.exists():
        return cached.read_bytes()
    own = client is None
    if client is None:
        client = httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    try:
        response = client.get(f"/api/casc/{file_id}", params={"version": build}, timeout=60)
        if response.is_redirect:
            raise CascError(
                f"file data id {file_id} for build {build} redirected to "
                f"{response.headers.get('location')!r}; CASC serves bytes, not redirects"
            )
        if response.status_code == 400:
            raise CascError(
                f"wago.tools does not know build {build!r} "
                f"(file data id {file_id}): {response.text.strip()}"
            )
        if response.status_code == 404:
            raise CascMissing(
                f"file data id {file_id} is not in CASC at version {build}",
                reason=CascMissingReason.NOT_FOUND,
            )
        # Anything else non-2xx stays an httpx.HTTPStatusError: that is what
        # pipeline/art.py's own test already expects from a 500, and this
        # migration does not get to change it.
        response.raise_for_status()
        if not response.content:
            raise CascMissing(
                f"file data id {file_id} is empty for build {build}; this build's "
                f"client does not ship that file",
                reason=CascMissingReason.EMPTY,
            )
        data = response.content
    finally:
        if own:
            client.close()
    cached.parent.mkdir(parents=True, exist_ok=True)
    _atomic_write(cached, data)
    logger.info("casc %s @ %s: %d bytes", file_id, build, len(data))
    return data
