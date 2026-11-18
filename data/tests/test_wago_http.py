"""The polite wago.tools client (pipeline.wago_http) and the table cache
(pipeline.wago_tables).

2026-11-17: wago.tools answered 403 to the pipeline after a data run made about
62,000 hotfix searches six at a time. Every wago request now goes one at a time,
spaced, within a budget, and a 403 stops the run.
"""

from __future__ import annotations

from pathlib import Path

import httpx
import pytest

from pipeline.wago import download_table
from pipeline.wago_http import (
    BUSY_ATTEMPTS,
    USER_AGENT,
    Pacer,
    WagoBudgetExceeded,
    WagoRefused,
    open_client,
    pacer_from_env,
    retry_after_seconds,
)
from pipeline.wago_tables import TableCache, get_table


class Clock:
    """A fake monotonic clock that `sleep` advances."""

    def __init__(self) -> None:
        self.now = 0.0
        self.slept: list[float] = []

    def __call__(self) -> float:
        return self.now

    def sleep(self, seconds: float) -> None:
        self.slept.append(seconds)
        self.now += seconds


def client_over(handler, clock: Clock | None = None, budget: int = 100, interval: float = 1.0):
    clock = clock or Clock()
    pacer = Pacer(min_interval=interval, budget=budget, clock=clock, sleep=clock.sleep)
    return open_client(inner=httpx.MockTransport(handler), pacer=pacer), pacer, clock


def ok(request: httpx.Request) -> httpx.Response:
    return httpx.Response(200, text="ID\n1\n")


def test_the_user_agent_is_honest_and_contactable():
    assert USER_AGENT.startswith("ForeverSixtyDataPipeline/")
    assert "https://foreversixty.gg" in USER_AGENT and "help@foreversixty.gg" in USER_AGENT
    assert "Mozilla" not in USER_AGENT


def test_every_request_carries_the_user_agent():
    seen: list[str] = []

    def handler(request: httpx.Request) -> httpx.Response:
        seen.append(request.headers["user-agent"])
        return ok(request)

    client, _, _ = client_over(handler)
    client.get("/api/builds")
    assert seen == [USER_AGENT]


def test_requests_are_spaced_by_the_minimum_interval():
    client, pacer, clock = client_over(ok, interval=1.0)
    for _ in range(3):
        client.get("/api/builds")
    assert pacer.sent == 3
    assert clock.slept == [1.0, 1.0]


def test_a_403_stops_the_run_without_retrying():
    calls: list[str] = []

    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        return httpx.Response(403)

    client, _, _ = client_over(handler)
    with pytest.raises(WagoRefused, match="refused this pipeline"):
        client.get("/db2/Map/csv")
    assert len(calls) == 1


def test_a_429_waits_out_retry_after_then_succeeds():
    answers = iter([httpx.Response(429, headers={"Retry-After": "7"}), httpx.Response(200)])
    client, pacer, clock = client_over(lambda r: next(answers))
    assert client.get("/hotfixes").status_code == 200
    assert 7.0 in clock.slept and pacer.sent == 2


def test_a_503_without_retry_after_backs_off_and_gives_up_after_its_attempts():
    client, pacer, clock = client_over(lambda r: httpx.Response(503), budget=100)
    assert client.get("/hotfixes").status_code == 503
    assert pacer.sent == BUSY_ATTEMPTS
    assert clock.slept[0] >= 30.0


def test_a_retry_after_longer_than_the_cap_is_not_waited_for():
    client, pacer, _ = client_over(lambda r: httpx.Response(429, headers={"Retry-After": "86400"}))
    assert client.get("/hotfixes").status_code == 429
    assert pacer.sent == 1


def test_the_budget_fails_the_run_loudly():
    client, _, _ = client_over(ok, budget=2)
    client.get("/a")
    client.get("/b")
    with pytest.raises(WagoBudgetExceeded, match="budget of 2"):
        client.get("/c")


def test_retry_after_reads_an_http_date():
    response = httpx.Response(429, headers={"Retry-After": "Wed, 18 Nov 2026 00:00:30 GMT"})
    assert retry_after_seconds(response, now=lambda: 1_794_960_000.0) == 30.0


def test_the_pacer_reads_its_settings_from_the_env(monkeypatch):
    monkeypatch.setenv("WAGO_MIN_INTERVAL", "2.5")
    monkeypatch.setenv("WAGO_REQUEST_BUDGET", "40")
    pacer = pacer_from_env()
    assert (pacer.min_interval, pacer.budget) == (2.5, 40)
    monkeypatch.setenv("WAGO_MIN_INTERVAL", "fast")
    with pytest.raises(SystemExit, match="not a number"):
        pacer_from_env()


# --- the table cache ---------------------------------------------------------


class Exports:
    """wago.tools' table exports: a body, a last-modified, and a 304 for a
    matching If-Modified-Since."""

    LAST_MODIFIED = "Fri, 13 Nov 2026 04:27:52 GMT"

    def __init__(self) -> None:
        self.requests: list[httpx.Request] = []

    def handler(self, request: httpx.Request) -> httpx.Response:
        self.requests.append(request)
        if request.headers.get("if-modified-since") == self.LAST_MODIFIED:
            return httpx.Response(304)
        headers = {"last-modified": self.LAST_MODIFIED, "cache-control": "max-age=14400"}
        return httpx.Response(200, text="ID,Name_lang\n1,x\n", headers=headers)


def cache_over(tmp_path: Path, now: list[float]) -> TableCache:
    return TableCache(tmp_path / "cache", clock=lambda: now[0])


def test_a_shipped_table_is_downloaded_once(tmp_path: Path):
    exports = Exports()
    client, _, _ = client_over(exports.handler)
    cache = cache_over(tmp_path, [0.0])
    for _ in range(3):
        download_table("Map", "1.60.1.70938", tmp_path / "raw", client, cache=cache)
    assert len(exports.requests) == 1
    assert (tmp_path / "raw" / "Map.csv").read_text() == "ID,Name_lang\n1,x\n"


def test_a_hotfixed_table_is_reused_within_its_max_age_then_revalidated(tmp_path: Path):
    exports = Exports()
    client, _, _ = client_over(exports.handler)
    now = [0.0]
    cache = cache_over(tmp_path, now)
    first = get_table(client, "ItemSparse", "1.60.1.70938", hotfixed=True, cache=cache)
    now[0] = 14_000.0
    get_table(client, "ItemSparse", "1.60.1.70938", hotfixed=True, cache=cache)
    assert len(exports.requests) == 1
    now[0] = 15_000.0
    again = get_table(client, "ItemSparse", "1.60.1.70938", hotfixed=True, cache=cache)
    assert len(exports.requests) == 2
    assert exports.requests[1].headers["if-modified-since"] == Exports.LAST_MODIFIED
    assert exports.requests[1].url.params["useHotfixes"] == "1"
    assert again == first
    now[0] = 16_000.0  # the 304 restarted the max-age
    get_table(client, "ItemSparse", "1.60.1.70938", hotfixed=True, cache=cache)
    assert len(exports.requests) == 2


def test_a_missing_table_is_cached_too(tmp_path: Path):
    calls: list[str] = []

    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        return httpx.Response(404, json={"errors": "Table not found."})

    client, _, _ = client_over(handler)
    cache = cache_over(tmp_path, [0.0])
    for _ in range(2):
        path = download_table("JournalInstance", "1.15.9.1", tmp_path / "raw", client, cache=cache)
        assert path.read_text() == "ID\n"
    assert len(calls) == 1


def test_a_server_error_is_raised_and_not_cached(tmp_path: Path):
    client, _, _ = client_over(lambda r: httpx.Response(500))
    cache = cache_over(tmp_path, [0.0])
    with pytest.raises(httpx.HTTPStatusError):
        get_table(client, "Map", "1.60.1.70938", cache=cache)
    assert cache.read("Map", "1.60.1.70938", False) is None
