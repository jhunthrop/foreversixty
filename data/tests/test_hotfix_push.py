"""The server's hotfix push layered over the shipped spell-side tables
(sim-accuracy round 6, pipeline lane).

Round 5 read the 1.60.1.70707 push by hand: every spell generator (simconst,
trainables, spellranks, normalize, and the engine fork's relicmods,
periodiccrit and eureka) read only the shipped DB2 CSVs, so Ice Barrier's
5% / Frigid Power's 15% and Touch of the Grave's 15 s cooldown never reached
a generated file. `pipeline.hotfix_push` layers the push into `raw/` itself,
so every reader of `raw/<Table>.csv` sees the server's rows.
"""

from __future__ import annotations

import csv
import html
import json
from pathlib import Path

import httpx
import pytest

from pipeline.csvio import read_csv
from pipeline.hotfix_push import (
    LAYER_RECORD,
    PUSH_DIR,
    PUSH_RECORD,
    PUSH_TABLES,
    SHIPPED_DIR,
    STATUS_INVALID,
    STATUS_NOT_PUBLIC,
    STATUS_REMOVED,
    STATUS_VALID,
    PushEntry,
    check_push,
    fetch_push,
    layer_hotfix_push,
    layer_rows,
    push_moved,
    push_totals,
    read_push_record,
    search_terms,
)
from pipeline.wago import TABLES
from pipeline.wago_hotfixes import TABLES as ITEM_HOTFIX_TABLES

COLUMNS = ["ID", "EffectBasePointsF", "EffectMiscValue_0", "EffectMiscValue_1", "SpellID"]


def entry(
    record_id: int,
    status: int = STATUS_VALID,
    data: list | None = None,
    push_id: int = 100,
    entry_id: int = 1,
) -> PushEntry:
    return PushEntry(
        entry_id=entry_id, push_id=push_id, record_id=record_id, status=status, data=data
    )


def shipped() -> list[dict[str, str]]:
    return [
        {"ID": "1", "EffectBasePointsF": "3", "EffectMiscValue_0": "0", "EffectMiscValue_1": "0",
         "SpellID": "11426"},
        {"ID": "2", "EffectBasePointsF": "2.79999995232", "EffectMiscValue_0": "127",
         "EffectMiscValue_1": "0", "SpellID": "11426"},
        {"ID": "3", "EffectBasePointsF": "5", "EffectMiscValue_0": "0", "EffectMiscValue_1": "0",
         "SpellID": "1259691"},
    ]  # fmt: skip


# --- the layering rule ------------------------------------------------------


def test_a_pushed_row_replaces_the_shipped_row():
    rows, counts = layer_rows(shipped(), COLUMNS, [entry(1, data=[1, "5", [0, 0], 11426])])
    assert rows[0]["EffectBasePointsF"] == "5"
    assert counts.changed == 1 and counts.added == 0 and counts.removed == 0


def test_a_pushed_row_the_client_lacks_is_added_in_id_order():
    rows, counts = layer_rows(shipped(), COLUMNS, [entry(9, data=[9, "1", [4, 0], 1311703])])
    assert [row["ID"] for row in rows] == ["1", "2", "3", "9"]
    assert rows[-1] == {"ID": "9", "EffectBasePointsF": "1", "EffectMiscValue_0": "4",
                        "EffectMiscValue_1": "0", "SpellID": "1311703"}  # fmt: skip
    assert counts.added == 1


def test_a_removed_record_drops_the_shipped_row():
    rows, counts = layer_rows(shipped(), COLUMNS, [entry(3, status=STATUS_REMOVED)])
    assert [row["ID"] for row in rows] == ["1", "2"]
    assert counts.removed == 1


def test_the_newest_push_wins_per_record():
    """Ice Barrier 1364979 is in two 70707 pushes: 112753 says 3, 113486 says 5."""
    older = entry(1, data=[1, "3", [0, 0], 11426], push_id=112753, entry_id=7)
    newer = entry(1, data=[1, "5", [0, 0], 11426], push_id=113486, entry_id=2)
    for order in ([older, newer], [newer, older]):
        rows, _ = layer_rows(shipped(), COLUMNS, order)
        assert rows[0]["EffectBasePointsF"] == "5"


def test_a_later_removal_beats_an_earlier_push_and_a_later_push_beats_a_removal():
    removed_last = [
        entry(9, data=[9, "1", [0, 0], 1], push_id=1),
        entry(9, STATUS_REMOVED, None, 2),
    ]
    assert [row["ID"] for row in layer_rows(shipped(), COLUMNS, removed_last)[0]] == ["1", "2", "3"]
    pushed_last = [entry(3, STATUS_REMOVED, None, 1), entry(3, data=[3, "7", [0, 0], 1], push_id=2)]
    assert layer_rows(shipped(), COLUMNS, pushed_last)[0][2]["EffectBasePointsF"] == "7"


@pytest.mark.parametrize("status", [STATUS_INVALID, STATUS_NOT_PUBLIC])
def test_an_invalidated_or_hidden_record_keeps_the_shipped_row(status):
    rows, counts = layer_rows(
        shipped(), COLUMNS, [entry(1, status=status), entry(8, status=status)]
    )
    assert rows == shipped()
    assert counts.withheld == 2


def test_an_unchanged_cell_keeps_the_shipped_print():
    """wago prints a float at another width in the push; the value is the same."""
    rows, counts = layer_rows(shipped(), COLUMNS, [entry(2, data=[2, "2.8", [127, 0], 11426])])
    assert rows[1]["EffectBasePointsF"] == "2.79999995232"
    assert counts.changed == 0


def test_a_push_row_of_the_wrong_width_fails_loudly():
    with pytest.raises(SystemExit, match="SpellEffect record 1: 3 values for 5 columns"):
        layer_rows(shipped(), COLUMNS, [entry(1, data=[1, "5", 0])], table="SpellEffect")


def test_a_push_row_whose_id_column_disagrees_fails_loudly():
    with pytest.raises(SystemExit, match="ID 2 in the data of record 1"):
        layer_rows(shipped(), COLUMNS, [entry(1, data=[2, "5", [0, 0], 1])], table="SpellEffect")


def test_layering_is_deterministic():
    entries = [entry(9, data=[9, "1", [0, 0], 1]), entry(1, data=[1, "4", [0, 0], 1], push_id=3)]
    assert layer_rows(shipped(), COLUMNS, entries) == layer_rows(shipped(), COLUMNS, entries[::-1])


# --- which tables, which searches -------------------------------------------


def test_every_fetched_table_but_the_item_hotfix_pair_is_layered():
    """The item tables keep their own merge (pipeline.hotfix_merge); every other
    table any generator reads from raw/ is fetched, so it is layered."""
    assert set(PUSH_TABLES) == set(TABLES) - set(ITEM_HOTFIX_TABLES)
    assert {"SpellEffect", "SpellMisc", "SpellAuraOptions", "SkillLineAbility", "CurvePoint",
            "TraitDefinitionEffectPoints", "SpellClassOptions"} <= set(PUSH_TABLES)  # fmt: skip


def test_one_search_covers_every_table_whose_name_contains_it():
    terms = search_terms(("Spell", "SpellEffect", "SkillLine", "SkillLineAbility", "CurvePoint"))
    assert terms == ("CurvePoint", "SkillLine", "Spell")


# --- fetching the push from wago.tools --------------------------------------


def page(entries: list[dict], total: int, next_url: str | None = None) -> str:
    payload = {"props": {"hotfixes": {"data": entries, "total": total, "next_page_url": next_url}}}
    return f'<div id="app" data-page="{html.escape(json.dumps(payload))}"></div>'


def wago_entry(
    entry_id: int, record_id: int, table: str, push_id: int = 113486, build: int = 70707, **extra
) -> dict:
    data = [record_id, "5", [0, 0], 11426]
    text = f"enus {push_id} {record_id} {build} {table.lower()} {json.dumps(data)}"
    return {"id": entry_id, "push_id": push_id, "record_id": record_id, "status": 1,
            "build": build, "table_name": table, "locale": "enUS", "data": data,
            "search_text": text, **extra}  # fmt: skip


class FakeWago:
    """wago.tools' hotfix search as it behaves: one trimmed, case-blind
    substring match on the search text, 25 a page in push order, and past
    page one a page that repeats rows instead of showing the rest (the tie
    order is not stable)."""

    def __init__(self, entries: list[dict]) -> None:
        self.entries = entries
        self.searches: list[str] = []

    def handler(self, request: httpx.Request) -> httpx.Response:
        assert request.url.path == "/hotfixes"
        search = request.url.params["search"].strip().lower()
        self.searches.append(search)
        hits = sorted(
            (e for e in self.entries if search in e["search_text"]), key=lambda e: -e["push_id"]
        )
        more = "https://wago.tools/hotfixes?page=2" if len(hits) > 25 else None
        return httpx.Response(200, text=page(hits[:25], len(hits), more))

    def client(self) -> httpx.Client:
        return httpx.Client(
            base_url="https://wago.tools", transport=httpx.MockTransport(self.handler)
        )


def test_fetch_push_keeps_the_builds_own_rows_of_each_table():
    wago = FakeWago(
        [wago_entry(1, 1364979, "SpellEffect"), wago_entry(2, 5, "SpellEffect", build=70685),
         wago_entry(3, 1364981, "SpellEffectScaling"),
         wago_entry(4, 1364980, "SpellEffect", locale="deDE")]
    )  # fmt: skip
    pushed = fetch_push("1.60.1.70707", ("SpellEffect",), wago.client())
    assert [e.record_id for e in pushed["SpellEffect"]] == [1364979]
    assert wago.searches == ["70707 spelleffect"]


def test_fetch_push_splits_a_search_wider_than_one_page_by_record_id():
    """300 records in one push, ids 1000..1299, plus a record whose id is the
    suffix the split reaches (7): every one comes back exactly once."""
    entries = [wago_entry(n, 1000 + n, "CurvePoint", push_id=113283) for n in range(300)]
    entries += [wago_entry(900 + n, 7, "CurvePoint", push_id=113280 + n) for n in range(3)]
    entries += [wago_entry(950 + n, 2000 + n, "CurvePoint", push_id=113139) for n in range(40)]
    pushed = fetch_push("1.60.1.70707", ("CurvePoint",), FakeWago(entries).client())
    got = sorted(e.entry_id for e in pushed["CurvePoint"])
    assert got == sorted(e["id"] for e in entries)


def test_fetch_push_fails_loudly_on_an_http_error():
    client = httpx.Client(
        base_url="https://wago.tools", transport=httpx.MockTransport(lambda r: httpx.Response(503))
    )
    with pytest.raises(httpx.HTTPStatusError):
        fetch_push("1.60.1.70707", ("SpellEffect",), client)


def test_fetch_push_fails_loudly_on_a_page_it_cannot_read():
    client = httpx.Client(
        base_url="https://wago.tools",
        transport=httpx.MockTransport(lambda r: httpx.Response(200, text="<html>maintenance")),
    )
    with pytest.raises(SystemExit, match="no data-page"):
        fetch_push("1.60.1.70707", ("SpellEffect",), client)


def test_fetch_push_fails_loudly_when_a_page_does_not_add_up():
    short = page([wago_entry(1, 1364979, "SpellEffect")], 2)
    client = httpx.Client(
        base_url="https://wago.tools",
        transport=httpx.MockTransport(lambda r: httpx.Response(200, text=short)),
    )
    with pytest.raises(SystemExit, match="1 of 2 entries"):
        fetch_push("1.60.1.70707", ("SpellEffect",), client)


# --- layering a build's raw/ in place ---------------------------------------


FIXTURE = Path(__file__).parent / "fixtures" / "hotfix_push"
BUILD = "1.60.1.70707"


def fixture_push() -> dict[str, list[PushEntry]]:
    by_table: dict[str, list[PushEntry]] = {}
    for e in json.loads((FIXTURE / "push-70707.json").read_text()):
        by_table.setdefault(e["table_name"], []).append(
            PushEntry(e["id"], e["push_id"], e["record_id"], e["status"], e["data"])
        )
    return by_table


@pytest.fixture
def build_root(tmp_path: Path) -> Path:
    raw = tmp_path / BUILD / "raw"
    raw.mkdir(parents=True)
    for path in (FIXTURE / "raw").glob("*.csv"):
        (raw / path.name).write_bytes(path.read_bytes())
    return tmp_path


#: wago's answer to each search term's root search, the change signal.
FIXTURE_TOTALS = {"CurvePoint": 212, "Spell": 6134}


def fixture_totals(build: str, tables: tuple[str, ...]) -> dict[str, int]:
    return dict(FIXTURE_TOTALS)


def layer(root: Path, **kwargs) -> dict:
    kwargs.setdefault("totals", fixture_totals)
    return layer_hotfix_push(BUILD, root=root, fetch=lambda b, t: fixture_push(), **kwargs)


def effect_points(raw: Path) -> dict[str, str]:
    return {r["ID"]: r["EffectBasePointsF"] for r in read_csv(raw / "SpellEffect.csv")}


def test_layering_writes_the_servers_rows_into_raw_and_keeps_the_shipped_copy(build_root):
    raw = build_root / BUILD / "raw"
    summary = layer(build_root)
    assert effect_points(raw)["1364979"] == "5"
    assert effect_points(raw / SHIPPED_DIR)["1364979"] == "3"
    assert summary["tables"]["SpellEffect"]["changed"] == 4
    record = json.loads((raw / PUSH_DIR / LAYER_RECORD).read_text())
    assert record["build"] == BUILD and record["tables"]["SpellEffect"]["changed"] == 4
    assert (raw / PUSH_DIR / "SpellEffect.json").exists()


def test_layering_twice_layers_over_the_shipped_copy_not_over_itself(build_root):
    raw = build_root / BUILD / "raw"
    layer(build_root)
    first = (raw / "SpellEffect.csv").read_bytes()
    layer(build_root)
    assert (raw / "SpellEffect.csv").read_bytes() == first
    assert effect_points(raw / SHIPPED_DIR)["1364979"] == "3"


def test_a_fresh_fetch_over_a_layered_build_becomes_the_new_shipped_copy(build_root):
    raw = build_root / BUILD / "raw"
    layer(build_root)
    refetched = (FIXTURE / "raw" / "SpellEffect.csv").read_text().replace(",431,", ",432,")
    (raw / "SpellEffect.csv").write_text(refetched)
    layer(build_root)
    assert effect_points(raw / SHIPPED_DIR)["688591"] == "432"
    assert effect_points(raw)["1364979"] == "5"


def test_offline_layering_reads_the_cached_push(build_root):
    layer(build_root)
    raw = build_root / BUILD / "raw"
    (raw / "SpellEffect.csv").write_bytes((raw / SHIPPED_DIR / "SpellEffect.csv").read_bytes())
    layer_hotfix_push(BUILD, root=build_root, offline=True)
    assert effect_points(raw)["1364979"] == "5"


def test_offline_layering_without_a_cached_push_fails_loudly(build_root):
    with pytest.raises(SystemExit, match="no cached push"):
        layer_hotfix_push(BUILD, root=build_root, offline=True)


def test_layering_a_build_without_raw_fails_loudly(tmp_path):
    with pytest.raises(SystemExit, match="no raw data"):
        layer_hotfix_push(BUILD, root=tmp_path, fetch=lambda b, t: {}, totals=fixture_totals)


def test_a_table_the_build_does_not_have_is_skipped(build_root):
    summary = layer(build_root)
    assert "TraitNode" not in summary["tables"]
    assert not (build_root / BUILD / "raw" / "TraitNode.csv").exists()


def test_the_layered_csv_keeps_the_shipped_header_order(build_root):
    raw = build_root / BUILD / "raw"
    layer(build_root)
    with (raw / "SpellEffect.csv").open(newline="") as f:
        header = next(csv.reader(f))
    with (FIXTURE / "raw" / "SpellEffect.csv").open(newline="") as f:
        assert header == next(csv.reader(f))


# --- noticing a new push for a build that is already fetched ----------------


def test_push_totals_is_one_search_per_term():
    wago = FakeWago(
        [wago_entry(n, 1000 + n, "SpellEffect") for n in range(40)]
        + [wago_entry(100, 7, "CurvePoint"), wago_entry(101, 8, "CurvePoint", build=70685)]
    )
    totals = push_totals("1.60.1.70707", ("SpellEffect", "Spell", "CurvePoint"), wago.client())
    assert totals == {"CurvePoint": 1, "Spell": 40}
    assert sorted(wago.searches) == ["70707 curvepoint", "70707 spell"]


def test_layering_commits_a_push_record_beside_the_build(build_root):
    layer(build_root)
    record = json.loads((build_root / BUILD / PUSH_RECORD).read_text())
    assert record["build"] == BUILD
    assert record["search_totals"] == FIXTURE_TOTALS
    assert record["tables"]["SpellEffect"]["changed"] == 4
    assert read_push_record(build_root / BUILD) == record


def test_the_push_record_is_deterministic(build_root):
    layer(build_root)
    first = (build_root / BUILD / PUSH_RECORD).read_bytes()
    layer(build_root)
    assert (build_root / BUILD / PUSH_RECORD).read_bytes() == first


def test_the_totals_are_read_before_the_push(build_root):
    """A push landing between the two reads then shows as moved the next night,
    never as already layered."""
    calls: list[str] = []

    def totals(b, t):
        calls.append("totals")
        return dict(FIXTURE_TOTALS)

    def fetch(b, t):
        calls.append("fetch")
        return fixture_push()

    layer_hotfix_push(BUILD, root=build_root, fetch=fetch, totals=totals)
    assert calls == ["totals", "fetch"]


def test_offline_layering_keeps_the_recorded_totals(build_root):
    layer(build_root)
    layer_hotfix_push(BUILD, root=build_root, offline=True)
    assert read_push_record(build_root / BUILD)["search_totals"] == FIXTURE_TOTALS


@pytest.mark.parametrize(
    ("record", "live", "moved"),
    [
        (None, FIXTURE_TOTALS, True),
        ({"search_totals": FIXTURE_TOTALS}, dict(FIXTURE_TOTALS), False),
        ({"search_totals": FIXTURE_TOTALS}, {**FIXTURE_TOTALS, "Spell": 6135}, True),
        ({"search_totals": FIXTURE_TOTALS}, {"Spell": 6134}, True),
    ],
)
def test_push_moved(record, live, moved):
    assert push_moved(record, live) is moved


def test_check_push_compares_the_committed_record_with_wago(build_root):
    asked: list[tuple[str, ...]] = []

    def live(b, tables):
        asked.append(tables)
        return {**FIXTURE_TOTALS, "Spell": 6135}

    assert check_push(BUILD, root=build_root, totals=live) is True  # no record yet
    layer(build_root)
    assert check_push(BUILD, root=build_root, totals=fixture_totals) is False
    assert check_push(BUILD, root=build_root, totals=live) is True
    assert set(asked[-1]) == set(read_push_record(build_root / BUILD)["tables"])
