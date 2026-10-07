"""Wowhead's live Forever talent payload, saved into a build's raw/ directory.

Blizzard ships most Forever talent-tree changes as hotfixes, which live in the client's
DBCache and never reach the CDN DB2 tables wago.tools exports, so the trait tables
`normalize` reads miss them. Wowhead's talent payload is read off the live game and is
hotfix-aware; `pipeline.normalize.wowhead_overlay` lays it over the trait-table trees.

The endpoint answers with a `WH.setPageData("...", {...});` envelope followed by other
statements, so the payload is the first JSON object in the body and the rest is ignored.
"""

from __future__ import annotations

import json
import logging
import time
from datetime import UTC, datetime
from pathlib import Path

import httpx

from pipeline.wago import USER_AGENT

logger = logging.getLogger(__name__)

WOWHEAD_TALENTS_URL = "https://nether.wowhead.com/forever/data/talents-classic"
WOWHEAD_TALENTS_PARAMS = {"dv": "100"}
RAW_TALENTS_FILE = "wowhead-talents.json"
REQUEST_TIMEOUT_SECONDS = 120
MAX_ATTEMPTS = 3
RETRY_BACKOFF_SECONDS = 5.0
RETRYABLE_STATUS = frozenset({429, 500, 502, 503, 504})


class WowheadTalentsError(ValueError):
    """The response is not the Forever talent payload."""


def unwrap_talents(text: str) -> dict:
    """The first JSON object in a `WH.setPageData(...)` body, trailing statements ignored."""
    start = text.find("{")
    if start < 0:
        raise WowheadTalentsError("response contains no JSON object")
    try:
        payload, _ = json.JSONDecoder().raw_decode(text, start)
    except json.JSONDecodeError as exc:
        raise WowheadTalentsError(f"response's first object is not valid JSON: {exc}") from exc
    missing = [key for key in ("trees", "talents") if not payload.get(key)]
    if missing:
        raise WowheadTalentsError(f"payload lacks {', '.join(missing)}")
    return payload


def raw_talents_path(build_dir: Path) -> Path:
    return build_dir / "raw" / RAW_TALENTS_FILE


def _get_with_retries(client: httpx.Client) -> httpx.Response:
    for attempt in range(1, MAX_ATTEMPTS + 1):
        try:
            response = client.get(
                WOWHEAD_TALENTS_URL, params=WOWHEAD_TALENTS_PARAMS, timeout=REQUEST_TIMEOUT_SECONDS
            )
            if response.status_code not in RETRYABLE_STATUS:
                response.raise_for_status()
                return response
            failure: Exception = httpx.HTTPStatusError(
                f"HTTP {response.status_code}", request=response.request, response=response
            )
        except httpx.TransportError as exc:
            failure = exc
        if attempt == MAX_ATTEMPTS:
            raise failure
        logger.warning("wowhead talents: attempt %d failed (%s); retrying", attempt, failure)
        time.sleep(RETRY_BACKOFF_SECONDS * attempt)
    raise AssertionError("unreachable")


def fetch_wowhead_talents(
    build: str, root: Path = Path("builds"), client: httpx.Client | None = None
) -> Path:
    """Download the live talent payload into `<root>/<build>/raw/wowhead-talents.json`.

    The file is the unwrapped payload plus a `_meta` object (fetched_at, url) that
    `normalize` drops before reading it.
    """
    own = client is None
    client = client or httpx.Client(headers={"User-Agent": USER_AGENT})
    try:
        response = _get_with_retries(client)
        payload = unwrap_talents(response.text)
    finally:
        if own:
            client.close()
    payload["_meta"] = {"url": str(response.url), "fetched_at": datetime.now(UTC).isoformat()}
    path = raw_talents_path(root / build)
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=1, sort_keys=True) + "\n", encoding="utf-8")
    logger.info("wowhead talents: %d trees saved to %s", len(payload["trees"]), path)
    return path
