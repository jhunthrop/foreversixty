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
"""

from __future__ import annotations

import json
import logging
import os
from collections.abc import Callable, Iterable, Mapping
from pathlib import Path
from typing import Any

import httpx

from pipeline.hotfix_push import (
    NodeTotals,
    PushCollector,
    PushEntry,
    PushFetcher,
    _build_number,
    crawl_terms,
    open_client,
)

logger = logging.getLogger(__name__)

#: Under raw/wago-push/: the saved push, committed with the build.
STORE = "entries.json"
STORE_FORMAT = 1
#: Set to 1 (or true) to force the full crawl.
FULL_ENV = "HOTFIX_PUSH_FULL"
_TRUE = frozenset({"1", "true", "yes"})


def full_requested(flag: bool) -> bool:
    return flag or os.environ.get(FULL_ENV, "").strip().lower() in _TRUE


def _entry_line(e: PushEntry) -> str:
    return json.dumps([e.entry_id, e.push_id, e.record_id, e.status, e.data], separators=(",", ":"))


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
    saved: Mapping[str, Any] | None, build: str, tables: tuple[str, ...], full: bool
) -> tuple[NodeTotals, Mapping[str, list[PushEntry]], str]:
    """The split totals and entries the crawl starts from, and the run's mode."""
    if saved is None:
        logger.info("hotfix push %s: full crawl (no saved push yet)", build)
        return {}, {}, "full"
    if full:
        logger.info("hotfix push %s: full crawl (asked for)", build)
        return {}, {}, "full"
    missing = sorted(set(tables) - set(saved["tables"]))
    if missing:
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
) -> dict[str, list[PushEntry]]:
    """The build's push for `tables` (enUS), crawled where it moved since the
    push saved at `store` (all of it when nothing is saved or `full`), and
    saved back."""
    saved = read_store(store)
    if saved is not None and saved["build"] != build:
        raise SystemExit(f"{store} is the push of {saved['build']}, not {build}")
    known, entries, mode = _starting_point(saved, build, tables, full)
    collector = PushCollector(_build_number(build), tables)
    collector.add_pushed(entries)
    own = client is None
    client = client or open_client()
    requests = _counting(client)
    try:
        observed = crawl_terms(client, collector, known)
    finally:
        if own:
            client.close()
    pushed = collector.by_table()
    write_store(store, build, tables, {**known, **observed}, pushed)
    logger.info("hotfix push %s: %s crawl, %d wago.tools requests", build, mode, requests[0])
    return pushed


def saved_fetcher(store_of: Callable[[str], Path], full: bool = False) -> PushFetcher:
    """A `layer_hotfix_push` fetcher over the saved push at `store_of(build)`."""
    return lambda build, tables: fetch_saved(build, tables, store_of(build), full=full)
