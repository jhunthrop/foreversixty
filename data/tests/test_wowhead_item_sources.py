# data/tests/test_wowhead_item_sources.py
"""Wowhead's Forever item pages as a source for an item the engine
fork's own database names no source for at all. See
pipeline.wowhead_item_sources's own module doc for why only four of
wowhead's listviews are read and why one (rep-vendor purchases) is not
covered at all.
"""

from pathlib import Path

import httpx
import pytest

from pipeline import wowhead_item_sources as wis

FIXTURES = Path(__file__).parent / "fixtures"
DROPPED_BY = FIXTURES / "wowhead-item-dropped-by.html"
SOLD_BY = FIXTURES / "wowhead-item-sold-by.html"
CREATED_BY = FIXTURES / "wowhead-item-created-by.html"
REWARD_FROM_Q = FIXTURES / "wowhead-item-reward-from-q.html"
NO_LISTVIEW = FIXTURES / "wowhead-item-no-listview.html"
NOT_FOUND = FIXTURES / "wowhead-item-notfound.html"


def test_parse_item_page_reads_dropped_by():
    """The fixture row also states `minlevel: 28, maxlevel: 29, count: 7,
    outof: 29824` -- wowhead-world-drops lane, 2026-09-29: parsed into
    `NpcSource.min_level`/`max_level`/`chance` (100 * 7/29824, rounded to
    4 decimal places), which `pipeline.loot.wowhead`'s world-drop pattern
    detector reads."""
    sources = wis.parse_item_page(720, DROPPED_BY.read_text(encoding="utf-8"))
    assert sources.dropped_by == [
        wis.NpcSource(
            npc_id=205,
            name="Nightbane Dark Runner",
            zone_ids=[10],
            min_level=28,
            max_level=29,
            chance=0.0235,
        )
    ]
    assert sources.sold_by == []
    assert sources.crafted_by == []
    assert sources.quest_rewards == []
    assert not sources.is_empty()


def test_parse_item_page_reads_sold_by():
    """A `sold-by` row states no `count`/`outof` (wowhead has a
    stock/cost there instead), so `chance` stays `None` even though
    `minlevel`/`maxlevel` are present. The fixture's own `cost` field is
    the plain-gold-price shape (`[[133081]]`, no `ItemExtendedCost`
    override), so `cost_money` is the gold price and `cost_item_ids` is
    empty -- vendor-11036 lane, 2026-09-30."""
    sources = wis.parse_item_page(16769, SOLD_BY.read_text(encoding="utf-8"))
    assert sources.sold_by == [
        wis.NpcSource(
            npc_id=11555,
            name="Gorn One Eye",
            zone_ids=[361],
            min_level=55,
            max_level=55,
            cost_money=133081,
        )
    ]


def test_parse_item_page_reads_created_by_and_maps_the_skill_line_to_a_profession():
    sources = wis.parse_item_page(2864, CREATED_BY.read_text(encoding="utf-8"))
    assert sources.crafted_by == [wis.CraftedSource(profession="blacksmithing")]


def test_parse_item_page_reads_reward_from_q_with_level_and_faction():
    sources = wis.parse_item_page(270000, REWARD_FROM_Q.read_text(encoding="utf-8"))
    assert sources.quest_rewards == [
        wis.QuestRewardSource(
            quest_id=176, name='WANTED: "Hogger"', min_level=5, level=11, faction="alliance"
        )
    ]


def test_parse_item_page_reads_reqfaction_and_reqrep_off_the_gatherer_blob():
    """vendor-11036 lane, 2026-09-30: `jsonequip.reqfaction`/`reqrep`
    live in `WH.Gatherer.addData(3, 16, {...})`, a separate blob from
    every listview this module already reads -- verified against a
    primary source (item 21200, "Signet Ring of the Bronze Dragonflight":
    wowhead states `reqfaction: 910, reqrep: 7`; the fork's OWN `rep`
    source for the identical id is `rep:brood-of-nozdormu:exalted`,
    faction id 910). A synthetic minimal page (never the network)."""
    html = (
        'new Listview({template: "item", id: "outfit", data: []});'
        'WH.Gatherer.addData(3, 16, {"21200":{"name_enus":"Signet Ring of the Bronze '
        'Dragonflight","jsonequip":{"reqfaction":910,"reqrep":7,"reqlevel":60}}});'
    )
    sources = wis.parse_item_page(21200, html)
    assert sources.required_faction_id == 910
    assert sources.required_standing_raw == 7


def test_parse_item_page_leaves_reqfaction_none_when_the_gatherer_record_states_none():
    html = (
        'new Listview({template: "item", id: "outfit", data: []});'
        'WH.Gatherer.addData(3, 16, {"239512":{"name_enus":"Lightbreaker Wrists",'
        '"jsonequip":{"buyprice":252635,"reqlevel":60}}});'
    )
    sources = wis.parse_item_page(239512, html)
    assert sources.required_faction_id is None
    assert sources.required_standing_raw is None


def test_parse_item_page_returns_empty_lists_for_a_page_with_no_usable_listview():
    """A real page (has Listview blocks) that only carries a listview
    this module does not read (contained-in-item, see the module doc) --
    a normal, cacheable "checked, nothing usable" outcome, not an
    ItemPageNotFound."""
    sources = wis.parse_item_page(270257, NO_LISTVIEW.read_text(encoding="utf-8"))
    assert sources.is_empty()


def test_parse_item_page_raises_for_a_true_notfound_page():
    with pytest.raises(wis.ItemPageNotFound, match="27000"):
        wis.parse_item_page(27000, NOT_FOUND.read_text(encoding="utf-8"))


def test_load_cached_item_sources_reads_only_what_is_already_cached(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    cache_path = wis.raw_path(build_dir, 720)
    cache_path.parent.mkdir(parents=True)
    cache_path.write_text(DROPPED_BY.read_text(encoding="utf-8"), encoding="utf-8")

    result = wis.load_cached_item_sources("1.60.1.70009", [720, 999], root=tmp_path)
    assert set(result.sources) == {720}
    assert result.sources[720].dropped_by[0].npc_id == 205
    assert result.missing == {999}


def test_fetch_item_sources_retries_a_throttled_response_with_backoff(
    tmp_path: Path, monkeypatch
):
    sleeps: list[float] = []
    monkeypatch.setattr(wis.time, "sleep", lambda s: sleeps.append(s))
    attempts = {"n": 0}

    def handler(request: httpx.Request) -> httpx.Response:
        attempts["n"] += 1
        if attempts["n"] < 3:
            return httpx.Response(403, text="throttled")
        return httpx.Response(200, text=DROPPED_BY.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wis.fetch_item_sources("1.60.1.70009", [720], root=tmp_path, client=client, delay=0)

    assert set(result.sources) == {720}
    assert result.missing == set()
    assert attempts["n"] == 3
    assert sleeps[:2] == [30.0, 60.0]
    assert wis.raw_path(tmp_path / "1.60.1.70009", 720).exists()


def test_fetch_item_sources_stops_after_three_consecutive_throttles(tmp_path: Path, monkeypatch):
    monkeypatch.setattr(wis.time, "sleep", lambda s: None)

    def handler(request: httpx.Request) -> httpx.Response:
        return httpx.Response(403, text="throttled")

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wis.fetch_item_sources(
        "1.60.1.70009", [720, 721, 722], root=tmp_path, client=client, delay=0
    )
    assert result.sources == {}
    assert result.missing == {720, 721, 722}


def test_fetch_item_sources_stops_at_the_max_pages_budget(tmp_path: Path, monkeypatch):
    monkeypatch.setattr(wis.time, "sleep", lambda s: None)
    cached_path = wis.raw_path(tmp_path / "1.60.1.70009", 10)
    cached_path.parent.mkdir(parents=True)
    cached_path.write_text(DROPPED_BY.read_text(encoding="utf-8"), encoding="utf-8")

    requested: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        item_id = int(str(request.url).rsplit("=", 1)[-1])
        requested.append(item_id)
        return httpx.Response(200, text=DROPPED_BY.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wis.fetch_item_sources(
        "1.60.1.70009", [10, 20, 30], root=tmp_path, client=client, delay=0, max_pages=1
    )
    assert requested == [20]  # id 10's cache hit cost nothing toward the budget
    assert set(result.sources) == {10, 20}
    assert result.missing == {30}


def test_fetch_item_sources_preserves_caller_order_rather_than_sorting_ascending(
    tmp_path: Path, monkeypatch
):
    """src-crawl-order lane, 2026-09-29: `pending` used to be
    `sorted(set(item_ids))`, which discarded a crawl-priority-ordered
    backlog (`pipeline.loot.wowhead.unsourced_real_item_ids`) right
    before the live requests went out. Requests must follow the CALLER's
    order, not ascending item id."""
    monkeypatch.setattr(wis.time, "sleep", lambda s: None)
    requested: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        item_id = int(str(request.url).rsplit("=", 1)[-1])
        requested.append(item_id)
        return httpx.Response(200, text=DROPPED_BY.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    wis.fetch_item_sources(
        "1.60.1.70009", [300, 100, 200, 100], root=tmp_path, client=client, delay=0
    )
    assert requested == [300, 100, 200]  # NOT ascending id order; duplicate 100 fetched once


def test_fetch_item_sources_caches_a_live_fetch_and_a_true_miss_is_not_cached(tmp_path: Path):
    calls: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        item_id = int(str(request.url).rsplit("=", 1)[-1])
        calls.append(item_id)
        if item_id == 999:
            return httpx.Response(404, text=NOT_FOUND.read_text(encoding="utf-8"))
        return httpx.Response(200, text=DROPPED_BY.read_text(encoding="utf-8"))

    client = httpx.Client(transport=httpx.MockTransport(handler))
    result = wis.fetch_item_sources("1.60.1.70009", [720, 999], root=tmp_path, client=client, delay=0)  # noqa: E501
    assert set(result.sources) == {720}
    assert result.missing == {999}
    assert calls == [720, 999]
    assert wis.raw_path(tmp_path / "1.60.1.70009", 720).exists()
    assert not wis.raw_path(tmp_path / "1.60.1.70009", 999).exists()

    calls.clear()
    client2 = httpx.Client(transport=httpx.MockTransport(handler))
    wis.fetch_item_sources("1.60.1.70009", [720, 999], root=tmp_path, client=client2, delay=0)
    assert calls == [999]  # 720's cache is warm; only the miss is retried
