"""The build's hotfix push, saved between runs so a run reads only what moved.

The crawl (`pipeline.hotfix_push.crawl_terms`) splits every search term by the
record id's last digits until each read fits wago's 25-row page. Read whole at
1.60.1.70790 that was about 31,800 searches (29,265 under "spell") and over an
hour (data run 42727001388). Wago never takes an entry back, so a search whose
total is the one a past run read still holds exactly the entries that run
read. This module saves, per build, the entries and the total of every split
search in `raw/wago-push/entries.json` (committed with the build), and the
next crawl splits only the searches whose totals moved: an unchanged term
costs one search, and a new push costs the paths down to the records it
touched. No order of wago's listing is assumed (it lists by ingestion time,
and one ingestion holds thousands of rows of many pushes in no stable order),
so a late copy of an old push is found the same way as a new push.

A build with no saved file (a new client build), or a run with `--full` or
HOTFIX_PUSH_FULL=1, crawls everything and saves the file. A saved file that
does not cover every table asked for is not trusted: the run logs FULL CRAWL
and crawls everything. A fetch that cannot finish still fails
(`pipeline.hotfix_push`'s rule: never leave the build shipped-only) and the
file is only written after a complete one.

The saved file reaches main only with the build's commit at the end of data.yml,
so a later step failing used to lose a finished crawl (data run 42739221308: a
66-minute full crawl, then normalize failed). A run with HOTFIX_PUSH_CACHE set
also keeps a copy at `<HOTFIX_PUSH_CACHE>/<build>/entries.json`, which data.yml
saves with actions/cache straight after the fetch step whatever happens next;
a build with no committed file starts from that copy.
"""

from __future__ import annotations

import json
import logging
import os
import shutil
from collections.abc import Callable, Iterable, Mapping
from pathlib import Path
from typing import Any

import httpx

from pipeline.hotfix_push import (
    ENUS_ONLY,
    EVERY_LOCALE,
    NodeTotals,
    PushCollector,
    PushEntry,
    PushFetcher,
    _build_number,
    crawl_terms,
    newest_per_record,
    open_client,
)

logger = logging.getLogger(__name__)

#: Under raw/wago-push/: the saved push, committed with the build.
STORE = "entries.json"
STORE_FORMAT = 1
#: Set to 1 (or true) to force the full crawl.
FULL_ENV = "HOTFIX_PUSH_FULL"
#: A directory outside the build that keeps a copy of each saved push
#: (data.yml's actions/cache path); unset, no copy is kept.
CACHE_ENV = "HOTFIX_PUSH_CACHE"
_TRUE = frozenset({"1", "true", "yes"})


def full_requested(flag: bool) -> bool:
    return flag or os.environ.get(FULL_ENV, "").strip().lower() in _TRUE


def cache_dir_from_env() -> Path | None:
    value = os.environ.get(CACHE_ENV, "").strip()
    return Path(value) if value else None


def _cached_copy(cache_dir: Path | None, build: str, name: str) -> Path | None:
    return cache_dir / build / name if cache_dir is not None else None


def _start_from_cache(store: Path, cached: Path | None, build: str) -> None:
    """Put the cached copy in place when the build has no saved file of its own."""
    if store.exists() or cached is None or not cached.exists():
        return
    logger.info("hotfix push %s: starting from the cached push %s", build, cached)
    store.parent.mkdir(parents=True, exist_ok=True)
    shutil.copyfile(cached, store)


def _keep_in_cache(store: Path, cache_dir: Path | None, build: str) -> None:
    """Copy the saved file into the cache, which holds this build's push only."""
    if cache_dir is None:
        return
    for stale in cache_dir.iterdir() if cache_dir.is_dir() else ():
        if stale.name != build:
            shutil.rmtree(stale, ignore_errors=True)
    cached = cache_dir / build / store.name
    cached.parent.mkdir(parents=True, exist_ok=True)
    shutil.copyfile(store, cached)


def _entry_line(e: PushEntry) -> str:
    """An enUS entry is five fields (every spell-side store as it was); another
    locale's entry adds its locale."""
    fields: list[object] = [e.entry_id, e.push_id, e.record_id, e.status, e.data]
    if e.locale != PushEntry.locale:
        fields.append(e.locale)
    return json.dumps(fields, separators=(",", ":"))


def write_store(
    path: Path,
    build: str,
    tables: Iterable[str],
    splits: NodeTotals,
    pushed: Mapping[str, list[PushEntry]],
) -> None:
    """One split and one entry a line, all sorted, so a new push is a small diff."""
    head = {"build": build, "format": STORE_FORMAT, "tables": sorted(tables)}
    lines = [json.dumps(head, sort_keys=True, separators=(",", ":"))[:-1] + ',"splits":{']
    lines.append(",\n".join(f"{json.dumps(s)}:[{t},{x}]" for s, (t, x) in sorted(splits.items())))
    lines.append('},"entries":{')
    names = sorted(pushed)
    for i, table in enumerate(names):
        entries = sorted(pushed[table], key=lambda e: (e.record_id, e.push_id, e.entry_id))
        body = ",\n".join(_entry_line(e) for e in entries)
        tail = "]" + ("," if i < len(names) - 1 else "")
        lines.append(f"{json.dumps(table)}:[" + ("\n" + body + "\n" if body else "") + tail)
    lines.append("}}")
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text("\n".join(lines) + "\n", encoding="utf-8")


def read_store(path: Path) -> dict[str, Any] | None:
    """The saved push: its tables, its split totals and its entries as `PushEntry`."""
    if not path.exists():
        return None
    try:
        raw = json.loads(path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as error:
        raise SystemExit(f"{path}: unreadable saved push ({error})") from error
    if raw.get("format") != STORE_FORMAT:
        raise SystemExit(f"{path}: saved push format {raw.get('format')}, not {STORE_FORMAT}")
    return {
        **raw,
        "splits": {s: (int(t), int(x)) for s, (t, x) in raw["splits"].items()},
        "entries": {t: [PushEntry(*row) for row in rows] for t, rows in raw["entries"].items()},
    }


def _starting_point(
    saved: Mapping[str, Any] | None,
    build: str,
    tables: tuple[str, ...],
    full: bool,
    terms: tuple[str, ...] | None = None,
) -> tuple[NodeTotals, Mapping[str, list[PushEntry]], str]:
    """The split totals and entries the crawl starts from, and the run's mode.

    A table the saved file does not cover forces the full crawl when searches
    are shared between tables ("Spell" also reads SpellEffect, whose entries a
    saved split read without them would hold). With one search per table
    (`terms`) a new table's searches were never saved, so only it is crawled."""
    if saved is None:
        logger.info("hotfix push %s: full crawl (no saved push yet)", build)
        return {}, {}, "full"
    if full:
        logger.info("hotfix push %s: full crawl (asked for)", build)
        return {}, {}, "full"
    missing = sorted(set(tables) - set(saved["tables"]))
    if missing and terms is not None:
        logger.info("hotfix push %s: crawling the new tables %s", build, ", ".join(missing))
    elif missing:
        logger.error(
            "hotfix push %s: FULL CRAWL, the saved push does not cover %s",
            build,
            ", ".join(missing),
        )
        return {}, {}, "full"
    return saved["splits"], saved["entries"], "incremental"


def _counting(client: httpx.Client) -> list[int]:
    count = [0]

    def tick(_: httpx.Request) -> None:
        count[0] += 1

    hooks = client.event_hooks
    client.event_hooks = {**hooks, "request": [*hooks.get("request", []), tick]}
    return count


def fetch_saved(
    build: str,
    tables: tuple[str, ...],
    store: Path,
    full: bool = False,
    client: httpx.Client | None = None,
    cache_dir: Path | None = None,
    locales: frozenset[str] | None = ENUS_ONLY,
    keep: Callable[[list[PushEntry]], list[PushEntry]] | None = None,
    terms: tuple[str, ...] | None = None,
) -> dict[str, list[PushEntry]]:
    """The build's push for `tables` in `locales`, crawled where it moved since the
    push saved at `store` (all of it when nothing is saved or `full`), and
    saved back. With `cache_dir`, a missing `store` starts from the copy kept
    there, and the saved file is copied back into it. `keep`, when given,
    picks the entries of each table that are saved and returned, and `terms`,
    when given, are the searches crawled in place of the tables' names."""
    _start_from_cache(store, _cached_copy(cache_dir, build, store.name), build)
    saved = read_store(store)
    if saved is not None and saved["build"] != build:
        raise SystemExit(f"{store} is the push of {saved['build']}, not {build}")
    known, entries, mode = _starting_point(saved, build, tables, full, terms)
    collector = PushCollector(_build_number(build), tables, locales)
    collector.add_pushed(entries)
    own = client is None
    client = client or open_client()
    requests = _counting(client)
    try:
        observed = crawl_terms(client, collector, known, terms)
    finally:
        if own:
            client.close()
    pushed = collector.by_table()
    if keep is not None:
        pushed = {table: keep(entries) for table, entries in pushed.items()}
    write_store(store, build, tables, {**known, **observed}, pushed)
    _keep_in_cache(store, cache_dir, build)
    logger.info("hotfix push %s: %s crawl, %d wago.tools requests", build, mode, requests[0])
    return pushed


#: Under raw/wago-push/: the saved entries that carry values, every locale.
VALUES_STORE = "values.json"


def value_terms(tables: Iterable[str]) -> tuple[str, ...]:
    """One search per table for its entries that carry values: such an entry
    reads "<table> [<id>,..." (the bracket opens the values), and the build
    number before the table name keeps "Spell [" from matching SpellEffect."""
    return tuple(f"{table} [" for table in tables)


def deciding_entries(entries: list[PushEntry]) -> list[PushEntry]:
    """The one entry per record the layer reads. Every push of a record comes
    in up to eleven locales with the same values, so a values store keeps the
    deciding entry only; a later push still outranks it, so the incremental
    crawl lands on the same result."""
    newest = newest_per_record(entries)
    return [newest[record_id] for record_id in sorted(newest)]


def fetch_values(
    build: str,
    tables: tuple[str, ...],
    store: Path,
    full: bool = False,
    client: httpx.Client | None = None,
) -> dict[str, list[PushEntry]]:
    """The build's pushed entries that carry values, in every locale, one per
    record, crawled over the ones saved at `store`.

    Since 70938 (2026-11-13) wago.tools lists a push's values only on some
    locales' copies: every enUS entry of 70938's push carries `data: None`, so
    the enUS crawl layered nothing for any table (issue #27)."""
    return fetch_saved(
        build,
        tables,
        store,
        full=full,
        client=client,
        cache_dir=cache_dir_from_env(),
        locales=EVERY_LOCALE,
        keep=deciding_entries,
        terms=value_terms(tables),
    )


def values_fetcher(store_of: Callable[[str], Path], full: bool = False) -> PushFetcher:
    """A `layer_hotfix_push` values fetcher over the store at `store_of(build)`."""
    return lambda build, tables: fetch_values(build, tables, store_of(build), full=full)


def saved_fetcher(store_of: Callable[[str], Path], full: bool = False) -> PushFetcher:
    """A `layer_hotfix_push` fetcher over the saved push at `store_of(build)`,
    kept in HOTFIX_PUSH_CACHE too when that is set."""
    cache_dir = cache_dir_from_env()
    return lambda build, tables: fetch_saved(
        build, tables, store_of(build), full=full, cache_dir=cache_dir
    )
