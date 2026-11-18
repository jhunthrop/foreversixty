"""The one HTTP client every wago.tools request goes through.

On 2026-11-17 wago.tools began answering 403 to this pipeline's User-Agent:
a single data run had made about 62,000 hotfix searches with six in flight,
and the nightly ranking and the engine follow downloaded the same immutable
tables again. This module is the pipeline's side of sharing their servers:

* one request at a time, process-wide, whatever the caller's threading;
* at least `WAGO_MIN_INTERVAL` seconds (default 1.0) from the start of one
  request to the start of the next;
* a 429 or 503 waits out its Retry-After (or a doubling backoff) and tries
  again, a few times, then fails;
* a 403 stops the run at once: wago has refused us, and retrying would be
  exactly the load they asked us to stop;
* a per-run budget, `WAGO_REQUEST_BUDGET` (default `DEFAULT_BUDGET`), that
  fails the run loudly instead of quietly crawling for hours.

The User-Agent says who we are and how to reach us. It is never a browser's
or another client's.
"""

from __future__ import annotations

import logging
import os
import threading
import time
from collections.abc import Callable
from dataclasses import dataclass, field
from email.utils import parsedate_to_datetime

import httpx

logger = logging.getLogger(__name__)

BASE_URL = "https://wago.tools"
USER_AGENT = "ForeverSixtyDataPipeline/1.0 (+https://foreversixty.gg; help@foreversixty.gg)"

MIN_INTERVAL_ENV = "WAGO_MIN_INTERVAL"
BUDGET_ENV = "WAGO_REQUEST_BUDGET"
DEFAULT_MIN_INTERVAL_S = 1.0
#: A steady data run is a few hundred requests; a new build's first push read
#: is several thousand. Past this, something is wrong and the run stops.
DEFAULT_BUDGET = 12_000
#: Tries for a request wago answers 429 or 503.
BUSY_ATTEMPTS = 4
#: Seconds before the first retry of a busy answer with no Retry-After, doubled each time.
BUSY_BACKOFF_S = 30.0
#: The longest Retry-After honoured; a longer one fails the run instead of idling for it.
MAX_RETRY_AFTER_S = 600.0
BUSY_STATUSES = frozenset({429, 503})
TIMEOUT_S = 180.0


class WagoRefused(SystemExit):
    """wago.tools answered 403: the run stops and nothing retries."""


class WagoBudgetExceeded(SystemExit):
    """The run asked wago.tools for more than its request budget."""


def _env_float(name: str, default: float) -> float:
    raw = os.environ.get(name, "").strip()
    if not raw:
        return default
    try:
        value = float(raw)
    except ValueError as error:
        raise SystemExit(f"{name}={raw!r} is not a number") from error
    if value < 0:
        raise SystemExit(f"{name}={raw!r} is negative")
    return value


@dataclass
class Pacer:
    """Process-wide pacing and accounting: one request at a time, spaced by
    `min_interval`, counted against `budget`."""

    min_interval: float
    budget: int
    clock: Callable[[], float] = time.monotonic
    sleep: Callable[[float], None] = time.sleep
    sent: int = 0
    _last_start: float | None = None
    _lock: threading.Lock = field(default_factory=threading.Lock)

    def send(self, request: httpx.Request, inner: httpx.BaseTransport) -> httpx.Response:
        with self._lock:
            if self.sent >= self.budget:
                raise WagoBudgetExceeded(
                    f"wago.tools request budget of {self.budget} spent "
                    f"(next: {request.url}); raise {BUDGET_ENV} only if this run "
                    "really needs more"
                )
            if self._last_start is not None:
                wait = self._last_start + self.min_interval - self.clock()
                if wait > 0:
                    self.sleep(wait)
            self._last_start = self.clock()
            self.sent += 1
            return inner.handle_request(request)

    def pause(self, seconds: float) -> None:
        """Hold every request back for `seconds` (a busy answer's wait)."""
        with self._lock:
            self.sleep(seconds)
            self._last_start = self.clock()


def pacer_from_env() -> Pacer:
    budget = _env_float(BUDGET_ENV, DEFAULT_BUDGET)
    return Pacer(
        min_interval=_env_float(MIN_INTERVAL_ENV, DEFAULT_MIN_INTERVAL_S), budget=int(budget)
    )


_PACER: Pacer | None = None
_PACER_LOCK = threading.Lock()


def shared_pacer() -> Pacer:
    """The process's one pacer, read from the environment on first use."""
    global _PACER
    with _PACER_LOCK:
        if _PACER is None:
            _PACER = pacer_from_env()
        return _PACER


def requests_sent() -> int:
    """How many requests this process has sent to wago.tools."""
    return _PACER.sent if _PACER is not None else 0


def retry_after_seconds(
    response: httpx.Response, now: Callable[[], float] = time.time
) -> float | None:
    """A Retry-After header as seconds: delta-seconds or an HTTP date."""
    value = response.headers.get("retry-after", "").strip()
    if not value:
        return None
    if value.isdigit():
        return float(value)
    try:
        return max(0.0, parsedate_to_datetime(value).timestamp() - now())
    except (TypeError, ValueError):
        return None


class PoliteTransport(httpx.BaseTransport):
    """An httpx transport that sends through a `Pacer` and applies the rules above."""

    def __init__(
        self, inner: httpx.BaseTransport | None = None, pacer: Pacer | None = None
    ) -> None:
        self._inner = inner or httpx.HTTPTransport()
        self._pacer = pacer

    @property
    def pacer(self) -> Pacer:
        return self._pacer or shared_pacer()

    def handle_request(self, request: httpx.Request) -> httpx.Response:
        for attempt in range(BUSY_ATTEMPTS):
            response = self.pacer.send(request, self._inner)
            if response.status_code == 403:
                response.read()
                raise WagoRefused(
                    f"wago.tools answered 403 to {request.url} for User-Agent "
                    f"{request.headers.get('user-agent')!r}: they have refused this pipeline. "
                    "Stopping the run; do not retry. Reach out to wago.tools before running again."
                )
            if response.status_code not in BUSY_STATUSES or attempt == BUSY_ATTEMPTS - 1:
                return response
            wait = retry_after_seconds(response)
            wait = BUSY_BACKOFF_S * 2**attempt if wait is None else wait
            if wait > MAX_RETRY_AFTER_S:
                return response
            response.read()
            response.close()
            logger.warning(
                "wago.tools answered %d to %s; waiting %.0f s",
                response.status_code,
                request.url,
                wait,
            )
            self.pacer.pause(wait)
        raise AssertionError("unreachable")

    def close(self) -> None:
        self._inner.close()


def open_client(
    inner: httpx.BaseTransport | None = None,
    pacer: Pacer | None = None,
    base_url: str = BASE_URL,
) -> httpx.Client:
    """A wago.tools client: the honest User-Agent and the polite transport."""
    return httpx.Client(
        base_url=base_url,
        headers={"User-Agent": USER_AGENT},
        transport=PoliteTransport(inner, pacer),
        timeout=TIMEOUT_S,
    )


def log_requests(what: str) -> None:
    logger.info("%s: %d wago.tools requests this run", what, requests_sent())
