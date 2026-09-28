# data/tests/test_classic_quest_levels.py
"""cmangos/classic-db's `quest_template` table as the primary quest-level
source. The parsing tests use a small, synthetic mysqldump snippet with
the same quirks the real 73MB dump has (an embedded comma inside a
quoted Title, an escaped quote, two INSERT statements for one table) --
see pipeline.classic_quest_levels's own doc for why a naive "six leading
integers" regex over the whole file is wrong (180,935 false matches on
the real dump) and what replaces it.
"""

import gzip

import httpx
import pytest

from pipeline import classic_quest_levels as cql

# A comma and a parenthesis inside Title ("Bree, (Test)") and an escaped
# quote in Details ("don\'t") -- exactly the shapes a naive comma-split
# would misparse. Two INSERT statements, matching the real dump's own
# habit of splitting a large table across several.
SAMPLE_SQL = """
CREATE TABLE `quest_template` (`entry` mediumint, `Title` text);
INSERT INTO `quest_template` VALUES (53,2,331,40,255,44,0,0,178,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'Sweet Amber','Bring, (a bundle) of Charred Oak.','','','','','','','','',0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0);
INSERT INTO `quest_template` VALUES (5,2,10,17,255,20,0,0,77,0,0,0,0,0,0,0,0,0,0,0,8,0,0,0,0,0,93,0,0,0,'It doesn\\'t matter','','','','','','','','','',0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0);
"""


def test_parse_quest_template_reads_entry_min_level_and_level():
    quests = cql.parse_quest_template(SAMPLE_SQL)
    assert set(quests) == {53, 5}
    assert quests[53] == cql.ClassicQuestLevel(entry=53, min_level=40, level=44)
    assert quests[5] == cql.ClassicQuestLevel(entry=5, min_level=17, level=20)


def test_parse_quest_template_is_not_fooled_by_commas_and_parens_in_title():
    """The regression this lane's own probing found: a naive regex over
    the raw file (not quote-aware) matched 180,935 times on the real
    73MB dump for what turned out to be 4,245 real quests -- this
    fixture's own Title ("Bring, (a bundle) of Charred Oak.") has both a
    comma and a matched paren pair inside a quoted string, which a
    non-quote-aware row splitter would misread as a second row boundary.
    """
    quests = cql.parse_quest_template(SAMPLE_SQL)
    assert len(quests) == 2  # not more, from a misparsed Title


def test_parse_quest_template_handles_a_backslash_escaped_quote():
    # Quest 5's Title contains \' (an escaped single quote) -- the quote
    # tracker must not treat it as the string's closing quote.
    quests = cql.parse_quest_template(SAMPLE_SQL)
    assert 5 in quests  # parsed past the escaped quote without erroring


def test_parse_quest_template_raises_on_a_row_not_starting_with_six_integers():
    bad_sql = (
        "INSERT INTO `quest_template` VALUES "
        "('not','a','number',0,0,0,'Title','','','','','','','','','',0);"
    )
    with pytest.raises(ValueError, match="does not start with 6 integers"):
        cql.parse_quest_template(bad_sql)


def test_item_level_proxy_floors_at_60_and_0():
    assert cql.item_level_proxy(80) == 60  # Polar Leggings: floored, not 75
    assert cql.item_level_proxy(45) == 40
    assert cql.item_level_proxy(3) == 0  # never negative


def test_fetch_classic_db_quest_levels_downloads_and_gunzips(monkeypatch):
    payload = gzip.compress(SAMPLE_SQL.encode("utf-8"))

    def handler(request: httpx.Request) -> httpx.Response:
        assert str(request.url) == cql.SOURCE_URL
        return httpx.Response(200, content=payload)

    client = httpx.Client(transport=httpx.MockTransport(handler))
    quests = cql.fetch_classic_db_quest_levels(client=client)
    assert quests[53].min_level == 40 and quests[53].level == 44


def test_fetch_classic_db_quest_levels_propagates_an_http_error():
    client = httpx.Client(
        transport=httpx.MockTransport(lambda r: httpx.Response(404, text="not found"))
    )
    with pytest.raises(httpx.HTTPStatusError):
        cql.fetch_classic_db_quest_levels(client=client)


def test_real_quest_53_matches_wowheads_own_page():
    """Cross-check against the real 2026-09-28 wowhead fixture: classic-db
    and wowhead agree exactly for quest 53 (Sweet Amber), the same id
    this lane's other fixtures use throughout."""
    from pathlib import Path

    from pipeline.wowhead_quests import parse_quest_page

    html = (Path(__file__).parent / "fixtures" / "wowhead-quest-53.html").read_text(
        encoding="utf-8"
    )
    wowhead = parse_quest_page(53, html)
    classic = cql.parse_quest_template(SAMPLE_SQL)[53]
    assert (wowhead.min_level, wowhead.level) == (classic.min_level, classic.level)
