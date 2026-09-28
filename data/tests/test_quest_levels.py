# data/tests/test_quest_levels.py
"""The merged, committed `quest-levels.json` (pipeline.quest_levels):
classic-db entries plus wowhead fill-in for the ids classic-db lacks."""

import gzip
import json
from pathlib import Path

import httpx

from pipeline import quest_levels as ql
from pipeline.classic_quest_levels import SOURCE_URL
from pipeline.wowhead_quests import raw_path as wowhead_raw_path

SAMPLE_SQL = (
    "INSERT INTO `quest_template` VALUES "
    "(53,2,331,40,255,44,0,0,178,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,"
    "'Sweet Amber','','','','','','','','','',0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,"
    "0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0);"
)

FIXTURE = Path(__file__).parent / "fixtures" / "wowhead-quest-53.html"


def test_load_quest_levels_on_a_build_with_no_file_returns_empty(tmp_path: Path):
    assert ql.load_quest_levels(tmp_path / "1.60.1.70009") == {}


def test_merge_classic_db_writes_the_committed_file_tagged_classic_db(tmp_path: Path):
    payload = gzip.compress(SAMPLE_SQL.encode("utf-8"))

    def handler(request: httpx.Request) -> httpx.Response:
        assert str(request.url) == SOURCE_URL
        return httpx.Response(200, content=payload)

    client = httpx.Client(transport=httpx.MockTransport(handler))
    stats = ql.merge_classic_db("1.60.1.70009", root=tmp_path, client=client)
    assert stats.added == 1 and stats.total == 1

    entries = ql.load_quest_levels(tmp_path / "1.60.1.70009")
    assert entries[53] == ql.QuestLevelEntry(
        min_level=40, level=44, source="classic-db", fetched_at=entries[53].fetched_at
    )

    # The file itself carries a provenance header naming the exact
    # upstream repo/commit (the controller's own instruction).
    document = json.loads(ql.raw_path(tmp_path / "1.60.1.70009").read_text(encoding="utf-8"))
    assert document["sources"]["classic-db"]["repo"] == "cmangos/classic-db"
    assert document["sources"]["classic-db"]["commit"]


def test_fetch_missing_from_wowhead_only_fetches_ids_not_already_covered(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    # Quest 53 is already covered (classic-db); 54 is not.
    ql._write(build_dir, {53: ql.QuestLevelEntry(min_level=40, level=44, source="classic-db", fetched_at="x")})

    requested: list[int] = []

    def handler(request: httpx.Request) -> httpx.Response:
        quest_id = int(str(request.url).rsplit("=", 1)[-1])
        requested.append(quest_id)
        text = FIXTURE.read_text(encoding="utf-8").replace("[53]", f"[{quest_id}]")
        return httpx.Response(200, text=text)

    client = httpx.Client(transport=httpx.MockTransport(handler))
    stats = ql.fetch_missing_from_wowhead(
        "1.60.1.70009", [53, 54], max_pages=10, root=tmp_path, delay=0, client=client
    )

    assert requested == [54]  # 53 was already covered; only 54 was fetched
    assert stats == ql.WowheadMergeStats(fetched=1, still_missing=0, needed=1)

    entries = ql.load_quest_levels(build_dir)
    assert entries[53].source == "classic-db"  # untouched
    assert entries[54].source == "wowhead"


def test_fetch_missing_from_wowhead_with_nothing_missing_does_not_touch_the_file(tmp_path: Path):
    build_dir = tmp_path / "1.60.1.70009"
    ql._write(build_dir, {53: ql.QuestLevelEntry(min_level=40, level=44, source="classic-db", fetched_at="x")})
    before = ql.raw_path(build_dir).read_text(encoding="utf-8")

    stats = ql.fetch_missing_from_wowhead("1.60.1.70009", [53], max_pages=10, root=tmp_path)
    assert stats == ql.WowheadMergeStats(fetched=0, still_missing=0, needed=0)
    assert ql.raw_path(build_dir).read_text(encoding="utf-8") == before
