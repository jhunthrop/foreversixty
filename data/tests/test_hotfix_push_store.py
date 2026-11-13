"""The saved hotfix push and the incremental crawl (pipeline.hotfix_push_store).

Data run 42727001388 (1.60.1.70790) spent 58 minutes and about 31,800 wago.tools
searches crawling the whole push and still timed out. Wago never takes an entry
back, so the entries and every split search's total are saved under
raw/wago-push/ and committed, and the next crawl splits only what moved.
"""

from __future__ import annotations

import html
import json
import logging
import random
from pathlib import Path

import httpx
import pytest

from pipeline.hotfix_push import (
    STATUS_REMOVED,
    STATUS_VALID,
    PushEntry,
    layer_hotfix_push,
    layer_rows,
)
from pipeline.hotfix_push_store import (
    STORE,
    STORE_FORMAT,
    VALUES_STORE,
    fetch_saved,
    full_requested,
    read_store,
)

BUILD = "1.60.1.70790"
NUMBER = 70790
TABLES = ("Spell", "SpellEffect")


def wago_entry(
    entry_id: int,
    record_id: int,
    table: str = "SpellEffect",
    push_id: int = 113000,
    build: int = NUMBER,
    locale: str = "enUS",
    value: str = "5",
    status: int = STATUS_VALID,
) -> dict:
    data = [record_id, value] if status == STATUS_VALID else None
    text = f"{locale.lower()} {push_id} {record_id} {build} {table.lower()} "
    text += json.dumps(data) if data else ""
    return {"id": entry_id, "push_id": push_id, "record_id": record_id, "status": status,
            "build": build, "table_name": table, "locale": locale, "data": data,
            "search_text": text}  # fmt: skip


def page_html(entries: list[dict], total: int) -> str:
    payload = {"props": {"hotfixes": {"data": entries, "total": total}}}
    return f'<div id="app" data-page="{html.escape(json.dumps(payload))}"></div>'


class Wago:
    """wago.tools' hotfix search as it behaves: a trimmed, case-blind substring
    match on the search text, 25 a page in no stable order."""

    def __init__(self, entries: list[dict]) -> None:
        self.entries = list(entries)
        self.searches: list[str] = []
        self._random = random.Random(7)

    def handler(self, request: httpx.Request) -> httpx.Response:
        search = request.url.params["search"].strip().lower()
        self.searches.append(search)
        hits = [e for e in self.entries if search in e["search_text"]]
        self._random.shuffle(hits)
        return httpx.Response(200, text=page_html(hits[:25], len(hits)))

    def client(self) -> httpx.Client:
        return httpx.Client(
            base_url="https://wago.tools", transport=httpx.MockTransport(self.handler)
        )


def history() -> list[dict]:
    """A build's push so far: 600 SpellEffect records, each in enUS and deDE, a
    record (7) whose id is a suffix the split reaches, a few Spell rows and
    another build's rows."""
    entries: list[dict] = []
    for n in range(600):
        for locale in ("enUS", "deDE"):
            entries.append(wago_entry(len(entries), 1000 + n, push_id=113000 + n % 30,
                                      locale=locale))  # fmt: skip
    entries += [wago_entry(5000 + n, 7, push_id=112900 + n) for n in range(3)]
    entries += [wago_entry(5100 + n, 5 + 10 * n, "Spell", push_id=113010) for n in range(3)]
    entries += [wago_entry(5200 + n, 7 + n, build=70862) for n in range(5)]
    return entries


def ids(pushed: dict) -> list[int]:
    return sorted(e.entry_id for entries in pushed.values() for e in entries)


def enus_ids(entries: list[dict]) -> list[int]:
    return sorted(e["id"] for e in entries if e["locale"] == "enUS" and e["build"] == NUMBER)


@pytest.fixture
def store(tmp_path: Path) -> Path:
    return tmp_path / "raw" / "wago-push" / STORE


def crawl(store: Path, entries: list[dict], **kwargs) -> tuple[dict, Wago]:
    wago = Wago(entries)
    return fetch_saved(BUILD, TABLES, store, client=wago.client(), **kwargs), wago


# --- the full crawl and the saved file -------------------------------------


def test_with_no_saved_file_the_full_crawl_runs_and_is_saved(store):
    pushed, wago = crawl(store, history())
    assert ids(pushed) == enus_ids(history())
    assert len(wago.searches) > 100
    saved = read_store(store)
    assert saved["format"] == STORE_FORMAT and saved["build"] == BUILD
    assert saved["tables"] == sorted(TABLES)
    assert saved["splits"]["70790 spell"] == (len(history()) - 5, 0)
    assert ids(saved["entries"]) == enus_ids(history())


def test_the_saved_file_is_deterministic(store):
    crawl(store, history())
    first = store.read_bytes()
    crawl(store, history(), full=True)
    assert store.read_bytes() == first


def test_full_forces_the_full_crawl_over_a_saved_file(store):
    _, first = crawl(store, history())
    _, again = crawl(store, history(), full=True)
    assert sorted(again.searches) == sorted(first.searches)


@pytest.mark.parametrize(
    ("value", "full"), [("1", True), ("true", True), ("", False), ("0", False)]
)
def test_the_full_crawl_env_var(monkeypatch, value, full):
    monkeypatch.setenv("HOTFIX_PUSH_FULL", value)
    assert full_requested(False) is full
    assert full_requested(True) is True


# --- the incremental crawl -------------------------------------------------


def test_an_unchanged_push_costs_one_search_per_term(store):
    crawl(store, history())
    pushed, wago = crawl(store, history())
    assert ids(pushed) == enus_ids(history())
    assert wago.searches == ["70790 spell"]


def test_an_incremental_run_splits_only_what_moved(store):
    """A new push of two records: the crawl walks down to them and no further."""
    _, full = crawl(store, history())
    new = [wago_entry(9000, 1234, push_id=113040), wago_entry(9001, 1234, push_id=113040,
           locale="deDE"), wago_entry(9002, 1555, push_id=113040)]  # fmt: skip
    pushed, wago = crawl(store, history() + new)
    assert ids(pushed) == enus_ids(history() + new)
    assert len(wago.searches) < len(full.searches) / 3
    # the record-id paths the split followed all end in the new records' digits
    split = [s.split()[0] for s in wago.searches if len(s.split()) == 3]
    assert split and all("1234".endswith(p[1:]) or "1555".endswith(p[1:]) for p in split)
    assert ids(read_store(store)["entries"]) == enus_ids(history() + new)


def test_a_late_copy_in_an_old_push_is_found_without_the_full_crawl(store, caplog):
    """wago still ingests old pushes (push 113477 gained deDE rows for 70790 two
    days after it went out); an enUS copy of an old push is found like a new one."""
    crawl(store, history())
    late = wago_entry(9100, 1016, push_id=113016, value="8")
    with caplog.at_level(logging.ERROR):
        pushed, _ = crawl(store, [*history(), late])
    assert 9100 in ids(pushed)
    assert "FULL CRAWL" not in caplog.text


def test_a_new_entry_of_a_record_that_is_a_split_suffix_is_read(store):
    """Record 7's entries are what the "7" split's children miss; a new one moves
    that count and is read by the exact-record search."""
    crawl(store, history())
    late = wago_entry(9200, 7, push_id=113050)
    pushed, _ = crawl(store, [*history(), late])
    assert 9200 in ids(pushed)


def test_a_record_count_skewed_by_a_row_landing_mid_read_is_read_again(store):
    """A row landing after a split's total is read but before its children are
    can make the exact-record count look unchanged; the split re-reads its total
    before it trusts that count."""
    crawl(store, history())

    class Landing(Wago):
        def handler(self, request: httpx.Request) -> httpx.Response:
            search = request.url.params["search"].strip().lower()
            if search == "07 70790 spell" and len(self.entries) == len(history()) + 1:
                self.entries.append(wago_entry(9301, 1007, push_id=113051))
            return super().handler(request)

    wago = Landing([*history(), wago_entry(9300, 7, push_id=113051)])
    pushed = fetch_saved(BUILD, TABLES, store, client=wago.client())
    assert {9300, 9301} <= set(ids(pushed))


def test_merging_keeps_the_newest_entry_per_record(store):
    old = [wago_entry(n, 1000 + n, push_id=113000, value="3") for n in range(30)]
    crawl(store, old)
    new = [wago_entry(100, 1000, push_id=113005, value="5"),
           wago_entry(101, 1001, push_id=113005, status=STATUS_REMOVED)]  # fmt: skip
    pushed, _ = crawl(store, old + new)
    shipped = [{"ID": "1000", "V": "3"}, {"ID": "1001", "V": "3"}]
    rows, counts = layer_rows(shipped, ["ID", "V"], pushed["SpellEffect"], "SpellEffect")
    assert [row for row in rows if row["ID"] in ("1000", "1001")] == [{"ID": "1000", "V": "5"}]
    assert counts.changed == 1 and counts.removed == 1


def test_a_table_the_saved_file_does_not_cover_falls_back_to_the_full_crawl(store, caplog):
    fetch_saved(BUILD, ("SpellEffect",), store, client=Wago(history()).client())
    with caplog.at_level(logging.ERROR):
        pushed, wago = crawl(store, history())
    assert ids(pushed) == enus_ids(history())
    assert "FULL CRAWL" in caplog.text
    assert len(wago.searches) > 100


def test_a_saved_file_of_another_build_is_refused(store):
    crawl(store, history())
    with pytest.raises(SystemExit, match="is the push of 1.60.1.70790"):
        fetch_saved("1.60.1.70862", TABLES, store, client=Wago(history()).client())


def test_a_crawl_that_cannot_finish_fails_and_keeps_the_saved_file(store):
    crawl(store, history())
    before = store.read_bytes()
    client = httpx.Client(
        base_url="https://wago.tools",
        transport=httpx.MockTransport(lambda r: httpx.Response(200, text="<html>maintenance")),
    )
    with pytest.raises(SystemExit, match="no data-page"):
        fetch_saved(BUILD, TABLES, store, client=client)
    assert store.read_bytes() == before


# --- the cached copy (data.yml's actions/cache) -----------------------------


def test_a_finished_crawl_is_kept_in_the_cache(store, tmp_path):
    cache = tmp_path / "cache"
    (cache / "1.60.1.70755").mkdir(parents=True)  # another build's copy is pruned
    crawl(store, history(), cache_dir=cache)
    assert (cache / BUILD / STORE).read_bytes() == store.read_bytes()
    assert sorted(p.name for p in cache.iterdir()) == [BUILD]


def test_a_build_with_no_saved_file_starts_from_the_cached_copy(store, tmp_path):
    """Run 42739221308: the crawl finished, normalize failed, nothing was
    committed. The next run's fresh checkout starts from the cached copy."""
    cache = tmp_path / "cache"
    crawl(store, history(), cache_dir=cache)
    store.unlink()
    pushed, wago = crawl(store, history(), cache_dir=cache)
    assert ids(pushed) == enus_ids(history())
    assert wago.searches == ["70790 spell"]


def test_a_committed_saved_file_wins_over_the_cached_copy(store, tmp_path):
    cache = tmp_path / "cache"
    crawl(store, history()[:100], cache_dir=cache)
    stale_cache = (cache / BUILD / STORE).read_bytes()
    crawl(store, history())  # a later run that kept no copy
    (cache / BUILD / STORE).write_bytes(stale_cache)
    pushed, wago = crawl(store, history(), cache_dir=cache)
    assert ids(pushed) == enus_ids(history())
    assert wago.searches == ["70790 spell"]  # read over the committed file, not the stale copy


def test_the_cache_dir_comes_from_the_env(monkeypatch, tmp_path):
    from pipeline.hotfix_push_store import CACHE_ENV, cache_dir_from_env

    monkeypatch.delenv(CACHE_ENV, raising=False)
    assert cache_dir_from_env() is None
    monkeypatch.setenv(CACHE_ENV, str(tmp_path))
    assert cache_dir_from_env() == tmp_path


def test_the_layer_reads_the_push_over_the_saved_file_by_default(tmp_path, monkeypatch):
    raw = tmp_path / BUILD / "raw"
    raw.mkdir(parents=True)
    (raw / "SpellEffect.csv").write_text("ID,V\n1000,3\n")
    asked: list[tuple] = []

    def fake(build, tables, store, full=False, cache_dir=None, **crawl):
        asked.append((build, tables, store.name, full, cache_dir, crawl.get("terms")))
        return {"SpellEffect": []}

    monkeypatch.setattr("pipeline.hotfix_push_store.fetch_saved", fake)
    monkeypatch.setenv("HOTFIX_PUSH_CACHE", str(tmp_path / "cache"))
    layer_hotfix_push(
        BUILD, root=tmp_path, tables=("SpellEffect",), totals=lambda b, t: {}, full=True
    )
    cache = tmp_path / "cache"
    assert sorted(asked, key=str) == sorted(
        [
            (BUILD, ("SpellEffect",), STORE, True, cache, None),
            (BUILD, ("SpellEffect",), VALUES_STORE, True, cache, ("SpellEffect [",)),
        ],
        key=str,
    )


def test_the_layer_takes_the_values_another_locale_carries(tmp_path, monkeypatch):
    """1.60.1.70938: every enUS entry of the push carries `data: None`, so a
    layer of the enUS crawl alone changed no row of any table (issue #27)."""
    raw = tmp_path / BUILD / "raw"
    raw.mkdir(parents=True)
    (raw / "SpellEffect.csv").write_text("ID,V\n1000,3\n")

    def fake(build, tables, store, full=False, cache_dir=None, **crawl):
        if crawl.get("terms") is None:
            return {"SpellEffect": [PushEntry(2, 114008, 1000, STATUS_VALID, None)]}
        return {"SpellEffect": [PushEntry(1, 114008, 1000, STATUS_VALID, [1000, "5"], "deDE")]}

    monkeypatch.setattr("pipeline.hotfix_push_store.fetch_saved", fake)
    record = layer_hotfix_push(
        BUILD, root=tmp_path, tables=("SpellEffect",), totals=lambda b, t: {}
    )
    assert (raw / "SpellEffect.csv").read_text() == "ID,V\n1000,5\n"
    assert record["tables"]["SpellEffect"]["changed"] == 1
