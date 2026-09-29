# data/tests/test_item_sources.py
"""The merged, committed `item-sources.json` (pipeline.item_sources):
wowhead's item-page scrape for items the engine fork's own database
names no source for at all."""

from pathlib import Path

import httpx

from pipeline import item_sources as isrc

FIXTURES = Path(__file__).parent / "fixtures"
DROPPED_BY = FIXTURES / "wowhead-item-dropped-by.html"
NO_LISTVIEW = FIXTURES / "wowhead-item-no-listview.html"


def test_load_item_sources_on_a_build_with_no_file_returns_empty(tmp_path: Path):
    assert isrc.load_item_sources(tmp_path / "1.60.1.70009") == {}


def test_fetch_missing_from_wowhead_only_fetches_ids_not_already_covered(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    # Item 720 is already covered; 721 is not.
    isrc._write(
        build_dir,
        {
            720: isrc.ItemSourceEntry(
                dropped_by=[{"npc_id": 205, "name": "Nightbane Dark Runner", "zone_ids": [10]}],
                source="wowhead",
                fetched_at="x",
            )
        },
    )

    requested: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        item_id = int(str(request.url).rsplit("=", 1)[-1])
        requested.append(item_id)
        return httpx.Response(200, text=DROPPED_BY.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    stats = isrc.fetch_missing_from_wowhead(
        "1.60.1.70009", [720, 721], max_pages=10, root=tmp_path, delay=0, client=client
    )

    assert requested == [721]  # 720 was already covered; only 721 was fetched
    assert stats == isrc.ItemSourceMergeStats(fetched=1, still_missing=0, needed=1, with_source=1)

    entries = isrc.load_item_sources(build_dir)
    assert entries[720].fetched_at == "x"  # untouched
    assert entries[721].dropped_by[0].npc_id == 205
    assert entries[721].source == "wowhead"


def test_fetch_missing_from_wowhead_commits_an_empty_result_so_it_is_not_refetched(
    tmp_path: Path,
):
    """A page with no usable listview (see wowhead_item_sources' own
    doc) is still written to the cache -- it was checked, there was
    nothing there, and a later run should not spend another request
    finding that out again."""

    def handler(request: httpx.Request) -> httpx.Response:
        return httpx.Response(200, text=NO_LISTVIEW.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    stats = isrc.fetch_missing_from_wowhead(
        "1.60.1.70009", [270257], max_pages=10, root=tmp_path, delay=0, client=client
    )
    assert stats == isrc.ItemSourceMergeStats(fetched=1, still_missing=0, needed=1, with_source=0)

    entries = isrc.load_item_sources(tmp_path / "1.60.1.70009")
    assert entries[270257].dropped_by == []
    assert entries[270257].sold_by == []

    # Rerunning makes no second request: the empty result is cached too.
    requested: list[int] = []

    def fail_handler(request: httpx.Request) -> httpx.Response:
        requested.append(int(str(request.url).rsplit("=", 1)[-1]))
        return httpx.Response(200, text=NO_LISTVIEW.read_text(encoding="utf-8"))

    client2 = httpx.Client(transport=httpx.MockTransport(fail_handler))
    stats2 = isrc.fetch_missing_from_wowhead(
        "1.60.1.70009", [270257], max_pages=10, root=tmp_path, delay=0, client=client2
    )
    assert requested == []
    assert stats2 == isrc.ItemSourceMergeStats(fetched=0, still_missing=0, needed=0, with_source=0)


def test_fetch_missing_from_wowhead_preserves_caller_priority_order(tmp_path: Path):
    """src-crawl-order lane, 2026-09-29: `needed` used to be
    `sorted(i for i in set(item_ids) if i not in entries)`, discarding
    whatever priority order the caller (`unsourced_real_item_ids`)
    passed in. Requests must go out in the caller's order."""
    requested: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        item_id = int(str(request.url).rsplit("=", 1)[-1])
        requested.append(item_id)
        return httpx.Response(200, text=DROPPED_BY.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    isrc.fetch_missing_from_wowhead(
        "1.60.1.70009", [900, 100, 500], max_pages=10, root=tmp_path, delay=0, client=client
    )
    assert requested == [900, 100, 500]  # NOT ascending id order


def test_fetch_missing_from_wowhead_counts_fetched_pages_that_named_any_source(tmp_path: Path):
    """`with_source` counts only pages that named a dropped-by/sold-by/
    crafted-by/quest-reward row -- 721 (DROPPED_BY fixture) counts,
    270257 (NO_LISTVIEW fixture, cacheable but empty) does not."""

    def handler(request: httpx.Request) -> httpx.Response:
        item_id = int(str(request.url).rsplit("=", 1)[-1])
        fixture = DROPPED_BY if item_id == 721 else NO_LISTVIEW
        return httpx.Response(200, text=fixture.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    stats = isrc.fetch_missing_from_wowhead(
        "1.60.1.70009", [721, 270257], max_pages=10, root=tmp_path, delay=0, client=client
    )
    assert stats == isrc.ItemSourceMergeStats(fetched=2, still_missing=0, needed=2, with_source=1)


def test_fetch_missing_from_wowhead_with_nothing_missing_does_not_touch_the_file(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    isrc._write(
        build_dir,
        {720: isrc.ItemSourceEntry(source="wowhead", fetched_at="x")},
    )
    before = isrc.raw_path(build_dir).read_text(encoding="utf-8")

    stats = isrc.fetch_missing_from_wowhead("1.60.1.70009", [720], max_pages=10, root=tmp_path)
    assert stats == isrc.ItemSourceMergeStats(fetched=0, still_missing=0, needed=0, with_source=0)
    assert isrc.raw_path(build_dir).read_text(encoding="utf-8") == before
