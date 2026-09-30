# data/tests/test_wowhead_quests.py
"""Wowhead's Forever quest pages as the source for a quest's own
min_level/level -- the fields the client's per-item required_level 0
does not carry for a quest reward. See pipeline.wowhead_quests's own
module doc for why this loops per-quest-id rather than reading one bulk
payload the way pipeline.wowhead_items does.
"""

from pathlib import Path

import httpx
import pytest

from pipeline import wowhead_quests as wq

FIXTURE = Path(__file__).parent / "fixtures" / "wowhead-quest-53.html"
NOT_FOUND_FIXTURE = Path(__file__).parent / "fixtures" / "wowhead-quest-notfound.html"


def test_parse_quest_page_reads_reqlevel_and_level():
    html = FIXTURE.read_text(encoding="utf-8")
    quest = wq.parse_quest_page(53, html)
    # data-followups-3 lane, 2026-09-30, item 2: `faction` joins
    # min_level/level off the SAME `g_quests[53]` payload -- this
    # fixture's own `side` is 1 (alliance), see QUEST_SIDE_FACTION.
    assert quest == wq.QuestPageLevel(
        quest_id=53, name="Sweet Amber", min_level=40, level=44, faction="alliance"
    )


def test_parse_quest_page_raises_when_the_page_has_no_g_quests_row():
    html = NOT_FOUND_FIXTURE.read_text(encoding="utf-8")
    with pytest.raises(wq.QuestPageNotFound, match="53"):
        wq.parse_quest_page(53, html)


def test_parse_quest_page_raises_for_the_wrong_quest_id():
    """The fixture's line is g_quests[53]; asking for a different id finds
    no matching marker even though the page loaded fine."""
    html = FIXTURE.read_text(encoding="utf-8")
    with pytest.raises(wq.QuestPageNotFound, match="54"):
        wq.parse_quest_page(54, html)


def test_load_cached_quest_levels_reads_only_what_is_already_cached(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    cache_path = wq.raw_path(build_dir, 53)
    cache_path.parent.mkdir(parents=True)
    cache_path.write_text(FIXTURE.read_text(encoding="utf-8"), encoding="utf-8")

    result = wq.load_cached_quest_levels("1.60.1.70009", [53, 999], root=tmp_path)
    assert result.levels == {53: wq.QuestPageLevel(quest_id=53, name="Sweet Amber", min_level=40, level=44, faction="alliance")}  # noqa: E501
    assert result.missing == {999}


def test_fetch_quest_levels_retries_a_throttled_response_with_backoff(tmp_path: Path, monkeypatch):
    """403 is throttling, not "no page": the request is retried (the SAME
    quest id, not skipped) after a backoff sleep, and succeeds once the
    mock stops throttling -- unlike a real 404/QuestPageNotFound, which
    goes straight to `.missing`."""
    sleeps: list[float] = []
    monkeypatch.setattr(wq.time, "sleep", lambda s: sleeps.append(s))

    attempts = {"n": 0}

    def handler(request: httpx.Request) -> httpx.Response:
        attempts["n"] += 1
        if attempts["n"] < 3:
            return httpx.Response(403, text="throttled")
        return httpx.Response(200, text=FIXTURE.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wq.fetch_quest_levels("1.60.1.70009", [53], root=tmp_path, client=client, delay=0)

    assert result.levels == {53: wq.QuestPageLevel(quest_id=53, name="Sweet Amber", min_level=40, level=44, faction="alliance")}  # noqa: E501
    assert result.missing == set()
    assert attempts["n"] == 3  # two throttled, one success -- same id retried, not skipped
    # Backoff doubled between the two throttled attempts: 30s, then 60s.
    assert sleeps[:2] == [30.0, 60.0]
    assert wq.raw_path(tmp_path / "1.60.1.70009", 53).exists()


def test_fetch_quest_levels_honours_retry_after_when_longer_than_the_backoff(
    tmp_path: Path, monkeypatch
):
    sleeps: list[float] = []
    monkeypatch.setattr(wq.time, "sleep", lambda s: sleeps.append(s))

    def handler(request: httpx.Request) -> httpx.Response:
        return httpx.Response(429, headers={"Retry-After": "120"}, text="throttled")

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wq.fetch_quest_levels("1.60.1.70009", [53], root=tmp_path, client=client, delay=0)

    # Stops after 3 consecutive throttled responses (THROTTLE_STOP_AFTER)
    # rather than retrying forever; the id is left missing for a later run.
    assert result.missing == {53}
    assert result.levels == {}
    # The 120s Retry-After beats the 30s starting backoff on every attempt
    # (it never grows past what Retry-After already demands).
    assert sleeps == [120.0, 120.0]  # only 2 sleeps: the 3rd throttle stops the run


def test_fetch_quest_levels_stopping_early_leaves_the_rest_missing_and_returns_normally(
    tmp_path: Path, monkeypatch
):
    """The controller's own requirement: stop after 3 consecutive
    throttled responses and exit normally (no exception) with a count of
    what remains, rather than hammering an address that is asking to be
    left alone."""
    monkeypatch.setattr(wq.time, "sleep", lambda s: None)

    def handler(request: httpx.Request) -> httpx.Response:
        return httpx.Response(403, text="throttled")

    client = httpx.Client(transport=httpx.MockTransport(handler))
    # Three ids, none cached: the first one alone exhausts 3 consecutive
    # throttles, so ids 2 and 3 are never even attempted -- still
    # reported missing.
    result = wq.fetch_quest_levels(
        "1.60.1.70009", [53, 54, 55], root=tmp_path, client=client, delay=0
    )
    assert result.levels == {}
    assert result.missing == {53, 54, 55}


def test_fetch_quest_levels_a_throttle_then_a_real_success_resets_the_consecutive_count(
    tmp_path: Path, monkeypatch
):
    """A successful fetch between throttled responses resets the
    consecutive-throttle counter, so an occasional 403 amid a mostly
    healthy run does not creep toward the stop threshold. Id 53 takes
    one throttle then succeeds; id 54 then takes TWO throttles (which
    would stop the run if the counter had not reset after 53's success)
    before it, too, succeeds -- proving the run does not stop early."""
    monkeypatch.setattr(wq.time, "sleep", lambda s: None)
    statuses = iter([403, 200, 403, 403, 200])

    def handler(request: httpx.Request) -> httpx.Response:
        status = next(statuses)
        if status == 200:
            # The fixture is quest 53's page; a request for id 54 gets the
            # same shape with 53 substituted for 54, so parse_quest_page
            # finds a real g_quests[54] row instead of QuestPageNotFound.
            quest_id = int(str(request.url).rsplit("=", 1)[-1])
            text = FIXTURE.read_text(encoding="utf-8").replace("[53]", f"[{quest_id}]")
            return httpx.Response(200, text=text)
        return httpx.Response(status, text="throttled")

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wq.fetch_quest_levels(
        "1.60.1.70009", [53, 54], root=tmp_path, client=client, delay=0
    )
    assert result.missing == set()
    assert set(result.levels) == {53, 54}


def test_fetch_quest_levels_stops_at_the_max_pages_budget(tmp_path: Path, monkeypatch):
    """`max_pages` caps LIVE requests only -- a cache hit costs nothing
    toward it. Of three ids, one is pre-cached and two are not: with
    max_pages=1, exactly one live request goes out and the other stays
    missing for a later run."""
    monkeypatch.setattr(wq.time, "sleep", lambda s: None)
    cached_path = wq.raw_path(tmp_path / "1.60.1.70009", 10)
    cached_path.parent.mkdir(parents=True)
    cached_path.write_text(
        FIXTURE.read_text(encoding="utf-8").replace("[53]", "[10]"), encoding="utf-8"
    )

    requested: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        quest_id = int(str(request.url).rsplit("=", 1)[-1])
        requested.append(quest_id)
        text = FIXTURE.read_text(encoding="utf-8").replace("[53]", f"[{quest_id}]")
        return httpx.Response(200, text=text)

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wq.fetch_quest_levels(
        "1.60.1.70009", [10, 20, 30], root=tmp_path, client=client, delay=0, max_pages=1
    )
    assert requested == [20]  # the cache hit (10) cost nothing; only one live request went out
    assert set(result.levels) == {10, 20}
    assert result.missing == {30}


def test_fetch_quest_levels_caches_a_live_fetch_and_skips_a_cache_hit(tmp_path: Path):
    calls: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        quest_id = int(str(request.url).rsplit("=", 1)[-1])
        calls.append(quest_id)
        if quest_id == 999:
            return httpx.Response(404, text=NOT_FOUND_FIXTURE.read_text(encoding="utf-8"))
        return httpx.Response(200, text=FIXTURE.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wq.fetch_quest_levels("1.60.1.70009", [53, 999], root=tmp_path, client=client, delay=0)
    assert result.levels == {53: wq.QuestPageLevel(quest_id=53, name="Sweet Amber", min_level=40, level=44, faction="alliance")}  # noqa: E501
    assert result.missing == {999}
    assert calls == [53, 999]
    assert wq.raw_path(tmp_path / "1.60.1.70009", 53).exists()
    # A 404 is not cached, so a rerun retries it rather than remembering
    # the miss forever.
    assert not wq.raw_path(tmp_path / "1.60.1.70009", 999).exists()

    # Rerunning with the same tmp_path makes no second request for 53 (the
    # cache is warm) -- 999 IS retried, since a miss is never cached.
    calls.clear()
    client2 = httpx.Client(transport=httpx.MockTransport(handler))
    result2 = wq.fetch_quest_levels("1.60.1.70009", [53, 999], root=tmp_path, client=client2, delay=0)  # noqa: E501
    assert calls == [999]
    assert result2.levels[53].min_level == 40
