# data/tests/test_wowhead_carry.py
"""The wowhead gear-planner carry-forward (design/reviews/2026-10-30-70707-dps-drop.md).

About 700 Forever items have no `ItemSparse` row and take their stats only from
wowhead's live feed. Between 28 and 30 Oct 2026 the feed lost 292 of them and the
items fell back to 1.12 stats. `fetch_wowhead` now merges the live feed into a
committed `raw/wowhead/items.json` seeded from the previous build: live records win,
a seeded record the feed no longer lists is kept, and it is dropped only when the
client loses the item or gains its own `ItemSparse` row for it.

Same fixture as test_wowhead_items.py: the cloak (279865) and the sword (271218) are
the planner gear the client lacks; the belt (11726) is the client's.
"""

import json
import re
from pathlib import Path

import httpx
import pytest

from pipeline import wowhead_items as wh

FIXTURE = Path(__file__).parent / "fixtures" / "wowhead-gear-planner.js"
FIXTURE_IDS = {11726, 264908, 271218, 279865, 777777, 888888}
CLOAK, SWORD, BELT = 279865, 271218, 11726


def _payload_without(*item_ids: int) -> str:
    """The fixture payload with the given item records cut out."""
    text = FIXTURE.read_text(encoding="utf-8")
    for item_id in item_ids:
        text = re.sub(rf'"{item_id}":\{{"id":{item_id},.*?"limitCategory":0[^}}]*\}},', "", text)
    return text


def _write_client(raw: Path, *, item_ids: set[int], sparse_ids: set[int]) -> None:
    raw.mkdir(parents=True, exist_ok=True)
    (raw / "Item.csv").write_text("ID\n" + "".join(f"{i}\n" for i in sorted(item_ids)))
    (raw / "ItemSparse.csv").write_text("ID\n" + "".join(f"{i}\n" for i in sorted(sparse_ids)))


def _client(text: str) -> httpx.Client:
    return httpx.Client(transport=httpx.MockTransport(lambda _r: httpx.Response(200, text=text)))


def _seed(build_dir: Path, text: str | None = None) -> None:
    records = wh.payload_records(text or FIXTURE.read_text(encoding="utf-8"))
    wh.write_records(wh.raw_path(build_dir), records, build=build_dir.name, seeded_from=None)


def test_payload_records_keep_what_the_pipeline_reads_and_drop_bookkeeping() -> None:
    records = wh.payload_records(FIXTURE.read_text(encoding="utf-8"))
    assert set(records) == FIXTURE_IDS
    cloak = records[CLOAK]
    assert "slotbak" not in cloak["stats"] and "itemSquishEraId" not in cloak["stats"]
    assert cloak["stats"]["sta"] == 5 and cloak["requiredLevel"] == 16
    assert "limitCategory" not in cloak


def test_the_carry_file_loads_exactly_as_the_payload_does(tmp_path: Path) -> None:
    _seed(tmp_path)
    assert wh.load_items(wh.raw_path(tmp_path)) == wh.load_items(FIXTURE)


def test_the_carry_file_refuses_records_that_are_not_forever_data(tmp_path: Path) -> None:
    records = wh.payload_records(FIXTURE.read_text(encoding="utf-8"))
    era = {i: {**r, "versionNum": 11300} for i, r in records.items()}
    wh.write_records(wh.raw_path(tmp_path), era, build="b", seeded_from=None)
    with pytest.raises(wh.WowheadPayloadError, match="11300"):
        wh.load_items(wh.raw_path(tmp_path))


def test_live_wins_and_a_record_the_feed_lost_is_carried() -> None:
    seed = wh.payload_records(FIXTURE.read_text(encoding="utf-8"))
    live = wh.payload_records(_payload_without(CLOAK))
    live[SWORD] = {**live[SWORD], "name": "Vileblood Scimitar (live)"}
    result = wh.carry_forward(live, seed, item_ids=FIXTURE_IDS, sparse_ids={BELT})
    assert result.records[SWORD]["name"] == "Vileblood Scimitar (live)"
    assert result.records[CLOAK] == seed[CLOAK]
    assert result.carried == (CLOAK,)


def test_a_carried_record_is_dropped_when_the_client_loses_the_item() -> None:
    seed = wh.payload_records(FIXTURE.read_text(encoding="utf-8"))
    live = wh.payload_records(_payload_without(CLOAK))
    result = wh.carry_forward(live, seed, item_ids=FIXTURE_IDS - {CLOAK}, sparse_ids={BELT})
    assert CLOAK not in result.records and result.carried == ()


def test_a_carried_record_is_dropped_when_the_client_gains_its_own_row() -> None:
    seed = wh.payload_records(FIXTURE.read_text(encoding="utf-8"))
    live = wh.payload_records(_payload_without(CLOAK))
    result = wh.carry_forward(live, seed, item_ids=FIXTURE_IDS, sparse_ids={BELT, CLOAK})
    assert CLOAK not in result.records


def test_coverage_is_measured_on_the_seeded_planner_gear_the_client_lacks() -> None:
    seed = wh.payload_records(FIXTURE.read_text(encoding="utf-8"))
    # The belt is the client's and the ring/robe/hauberk are not planner gear, so
    # losing them all says nothing about wowhead's supplement: still full coverage.
    live = wh.payload_records(_payload_without(BELT, 264908, 777777, 888888))
    full = wh.carry_forward(live, seed, item_ids=FIXTURE_IDS, sparse_ids={BELT})
    assert full.coverage == 1.0
    half = wh.carry_forward(
        wh.payload_records(_payload_without(CLOAK)), seed, item_ids=FIXTURE_IDS, sparse_ids={BELT}
    )
    assert half.coverage == 0.5
    assert wh.carry_forward(live, {}, item_ids=FIXTURE_IDS, sparse_ids=set()).coverage is None


def test_fetch_merges_the_live_feed_into_the_builds_own_carry_file(tmp_path: Path) -> None:
    build_dir = tmp_path / "1.60.1.2"
    _write_client(build_dir / "raw", item_ids=FIXTURE_IDS, sparse_ids={BELT})
    _seed(build_dir)
    # The feed lost the cloak but kept the sword: coverage 50% fails the floor...
    with pytest.raises(wh.WowheadCoverageError, match="1 of 2"):
        wh.fetch_wowhead("1.60.1.2", root=tmp_path, client=_client(_payload_without(CLOAK)))
    assert not wh.payload_path(build_dir).exists()
    # ...and a deliberate run past the floor still carries the cloak.
    path = wh.fetch_wowhead(
        "1.60.1.2", root=tmp_path, client=_client(_payload_without(CLOAK)), allow_shrink=True
    )
    assert path == build_dir / "raw" / "wowhead" / "items.json"
    document = json.loads(path.read_text(encoding="utf-8"))
    assert document["carried"] == [CLOAK] and document["live"] == 5
    assert {item.id for item in wh.load_items(path)} == FIXTURE_IDS


def test_fetch_seeds_from_the_newest_other_build_when_it_has_no_file_of_its_own(
    tmp_path: Path,
) -> None:
    _seed(tmp_path / "1.60.1.1")
    _seed(tmp_path / "1.60.1.10", _payload_without(SWORD))  # newest by version, not by name
    (tmp_path / "forever-prebeta").mkdir()
    build_dir = tmp_path / "1.60.1.11"
    _write_client(build_dir / "raw", item_ids=FIXTURE_IDS, sparse_ids={BELT})
    wh.fetch_wowhead("1.60.1.11", root=tmp_path, client=_client(FIXTURE.read_text()))
    document = json.loads(wh.raw_path(build_dir).read_text(encoding="utf-8"))
    assert document["seeded_from"] == "1.60.1.10" and document["carried"] == []


def test_fetch_with_no_seed_anywhere_writes_the_live_feed_alone(tmp_path: Path) -> None:
    path = wh.fetch_wowhead("1.60.1.1", root=tmp_path, client=_client(FIXTURE.read_text()))
    document = json.loads(path.read_text(encoding="utf-8"))
    assert document["seeded_from"] is None and len(document["items"]) == 6


def test_a_hotfix_dump_row_is_the_clients_own_row(tmp_path: Path) -> None:
    """The client ids are ItemSparse merged with the hotfix dump, as normalize and
    simdb read them: the dump outranks wowhead, so a seeded record it covers is not
    carried and does not count toward coverage."""
    build_dir = tmp_path / "1.60.1.2"
    _write_client(build_dir / "raw", item_ids=FIXTURE_IDS, sparse_ids={BELT})
    _write_client(build_dir / "raw" / "hotfixes", item_ids=set(), sparse_ids={CLOAK})
    _seed(build_dir)
    path = wh.fetch_wowhead("1.60.1.2", root=tmp_path, client=_client(_payload_without(CLOAK)))
    document = json.loads(path.read_text(encoding="utf-8"))
    assert document["carried"] == [] and str(CLOAK) not in document["items"]


def test_fetch_with_a_seed_needs_the_clients_item_tables(tmp_path: Path) -> None:
    _seed(tmp_path / "1.60.1.1")
    with pytest.raises(SystemExit, match="Item.csv"):
        wh.fetch_wowhead("1.60.1.2", root=tmp_path, client=_client(FIXTURE.read_text()))
