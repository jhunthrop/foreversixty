"""The server's hotfix push for a build, layered into the build's `raw/` tables.

Blizzard changes Forever's spells mostly by hotfix: the server pushes rows to
the client at login, and the DB2 files wago.tools exports never carry them.
Until round 6 of the sim-accuracy audit every spell generator read the shipped
CSVs only, so on 1.60.1.70707 Ice Barrier's pair stayed at 3 (pushed 5),
Frigid Power at 3 (pushed 15) and Touch of the Grave's cooldown at 25 s
(pushed 15 s); round 5 read the push by hand.

wago.tools keeps every pushed row it saw at https://wago.tools/hotfixes: one
entry per (push, record) with the record's status and, for a valid record, its
values in the table's column order. `layer_rows` is the one layering rule:

* per record id, the newest push wins (push id, then wago's entry id);
* status 1 (valid) replaces the shipped row, or adds one the client lacks;
* status 2 (record removed) drops the row;
* status 3 (invalidated) and 4 (not public) carry no values: the shipped row,
  if any, stands, and the record is counted as withheld.

`layer_hotfix_push` applies it to every table in `PUSH_TABLES` IN PLACE:
`raw/<Table>.csv` becomes the server's table and the file as fetched moves to
`raw/shipped/<Table>.csv`. Every reader of `raw/` (normalize, simconst,
trainables, spellranks, simdb, gen_set_rows, and the engine fork's relicmods,
periodiccrit and eureka, which read the site's build directory) therefore sees
the pushed rows with no code of its own. The CLI `fetch` runs it straight
after the download, so a fetched build is never left shipped-only.

Only the target build's push is read. The server sends a client every hotfix
live for its build, so a build's push is the whole set; an older build's push
may have been baked into the new client differently and is not replayed
(the item rule in `pipeline.hotfix_merge` is the same: a wago row replaces a
shipped row only when it was pushed for this build).

A push that cannot be fetched fails the run (an HTTP error, a page with no
data, a search that does not add up). `--offline` re-layers from the cached
push in `raw/wago-push/`, and fails when there is none.
"""

from __future__ import annotations

import csv
import hashlib
import html
import json
import logging
import re
import shutil
from collections.abc import Callable, Iterable, Mapping
from concurrent.futures import ThreadPoolExecutor
from dataclasses import asdict, dataclass
from pathlib import Path
from typing import Any

import httpx

from pipeline.csvio import read_csv
from pipeline.wago import BASE_URL, TABLES, USER_AGENT
from pipeline.wago_hotfixes import TABLES as ITEM_HOTFIX_TABLES

logger = logging.getLogger(__name__)

STATUS_VALID = 1
STATUS_REMOVED = 2
STATUS_INVALID = 3
STATUS_NOT_PUBLIC = 4

#: Every fetched table but the item pair, whose hotfixes `pipeline.hotfix_merge`
#: layers over the play-session dump by its own rule.
PUSH_TABLES: tuple[str, ...] = tuple(t for t in TABLES if t not in ITEM_HOTFIX_TABLES)
#: Under raw/: the cached push per table and the layer record.
PUSH_DIR = "wago-push"
LAYER_RECORD = "LAYERED.json"
#: Beside the build's committed outputs: the layer record plus the search
#: totals it was taken at. `check_push` compares the totals with wago's live
#: ones, and the engine follow's skip hash reads the file.
PUSH_RECORD = "hotfix-push.json"
#: Under raw/: each layered table as it was fetched.
SHIPPED_DIR = "shipped"

_LOCALE = "enUS"
_PAGE_DATA = re.compile(r'data-page="([^"]*)"')
_TIMEOUT_SECONDS = 60
#: wago.tools' fixed page size.
_PAGE_SIZE = 25
_DIGITS = "0123456789"
#: Searches in flight at once.
_WORKERS = 6

#: (build, tables) -> the build's push per table.
PushFetcher = Callable[[str, tuple[str, ...]], Mapping[str, list["PushEntry"]]]
#: (build, tables) -> how many entries wago lists for each search term.
TotalsFetcher = Callable[[str, tuple[str, ...]], Mapping[str, int]]


@dataclass(frozen=True)
class PushEntry:
    """One pushed record: wago's entry id, the push it came in, the record id,
    its status and, for a valid record, its values in column order (array
    columns nested)."""

    entry_id: int
    push_id: int
    record_id: int
    status: int
    data: list[Any] | None


@dataclass(frozen=True)
class LayerCounts:
    changed: int = 0
    added: int = 0
    removed: int = 0
    withheld: int = 0


def _flat(values: Iterable[Any]) -> list[Any]:
    out: list[Any] = []
    for value in values:
        out.extend(value if isinstance(value, list) else [value])
    return out


def _same(a: str, b: str) -> bool:
    """Equal by value: wago prints a float at another width in the push."""
    try:
        return round(float(a), 6) == round(float(b), 6)
    except ValueError:
        return a == b


def _pushed_row(
    entry: PushEntry, columns: list[str], base: dict[str, str] | None, table: str
) -> dict[str, str]:
    values = [str(v) for v in _flat(entry.data or [])]
    if len(values) != len(columns):
        raise SystemExit(
            f"{table} record {entry.record_id}: {len(values)} values for {len(columns)} "
            "columns; the push and the shipped table disagree on the schema"
        )
    row = dict(zip(columns, values, strict=True))
    if "ID" in row and row["ID"] != str(entry.record_id):
        raise SystemExit(
            f"{table}: ID {row['ID']} in the data of record {entry.record_id}; "
            "the push's column order is not the shipped table's"
        )
    if base is None:
        return row
    return {k: base[k] if _same(base[k], v) else v for k, v in row.items()}


def newest_per_record(entries: Iterable[PushEntry]) -> dict[int, PushEntry]:
    """The newest entry per record id: the higher push id, then wago's entry id."""
    newest: dict[int, PushEntry] = {}
    for e in sorted(entries, key=lambda e: (e.push_id, e.entry_id)):
        newest[e.record_id] = e
    return newest


def layer_rows(
    shipped: list[dict[str, str]],
    columns: list[str],
    entries: Iterable[PushEntry],
    table: str = "table",
) -> tuple[list[dict[str, str]], LayerCounts]:
    """`shipped` with the push layered over it by the module's rule, sorted by
    id, and how many rows changed, were added, removed or withheld."""
    by_id = {int(row["ID"]): row for row in shipped}
    changed = added = removed = withheld = 0
    for record_id, e in sorted(newest_per_record(entries).items()):
        base = by_id.get(record_id)
        if e.status == STATUS_REMOVED:
            removed += base is not None
            by_id.pop(record_id, None)
        elif e.status == STATUS_VALID and e.data:
            row = _pushed_row(e, columns, base, table)
            if base is None:
                added += 1
            elif row != base:
                changed += 1
            by_id[record_id] = row
        else:
            withheld += 1
    rows = [by_id[record_id] for record_id in sorted(by_id)]
    return rows, LayerCounts(changed, added, removed, withheld)


def search_terms(tables: Iterable[str]) -> tuple[str, ...]:
    """The fewest wago searches covering `tables`: its search is a substring
    match on the table name, so "Spell" also returns SpellEffect's entries."""
    names = sorted(set(tables))
    return tuple(t for t in names if not any(o != t and o.lower() in t.lower() for o in names))


def _build_number(build: str) -> int:
    return int(build.rsplit(".", 1)[-1])


def _read_page(text: str, url: str) -> dict[str, Any]:
    match = _PAGE_DATA.search(text)
    if not match:
        raise SystemExit(f"{url}: no data-page payload; wago.tools' hotfix page changed")
    try:
        return json.loads(html.unescape(match.group(1)))["props"]["hotfixes"]
    except (KeyError, json.JSONDecodeError) as error:
        raise SystemExit(f"{url}: unreadable hotfix payload ({error})") from error


@dataclass(frozen=True)
class _Answer:
    """One search's first page: its entries and how many the search matches."""

    entries: tuple[dict[str, Any], ...]
    total: int


def _ask(client: httpx.Client, search: str) -> _Answer:
    response = client.get("/hotfixes", params={"search": search}, timeout=_TIMEOUT_SECONDS)
    response.raise_for_status()
    payload = _read_page(response.text, str(response.url))
    return _Answer(tuple(payload["data"]), int(payload["total"]))


def _complete(answer: _Answer, search: str) -> tuple[dict[str, Any], ...]:
    if len(answer.entries) != answer.total:
        raise SystemExit(
            f"wago.tools search {search!r}: {len(answer.entries)} of {answer.total} entries "
            "on its one page"
        )
    return answer.entries


def _enumerate(client: httpx.Client, number: int, term: str) -> list[dict[str, Any]]:
    """Every entry whose search text holds "<record id> <build> <term>".

    wago's search is one substring match on "enus <push> <record> <build>
    <table> <data>", paged 25 at a time in push order with ties in no stable
    order, so a crawl past page one repeats some rows and never shows others
    (CurvePoint at 70707: 182 of 212, every time). A search that fits one page
    is complete, so a wider one is split by the record id's last digits: the
    children "<d><suffix> <build> <term>" are disjoint, and what they miss of
    the parent is the records whose id IS the suffix, split by the push id's
    last digit ("<d> <suffix> <build> <term>")."""
    found: list[dict[str, Any]] = []
    frontier: list[tuple[str, str]] = [("", "")]  # (suffix, its parent's suffix)
    parent_totals: dict[str, int] = {}
    with ThreadPoolExecutor(max_workers=_WORKERS) as pool:
        while frontier:
            searches = [f"{suffix} {number} {term}".strip() for suffix, _ in frontier]
            answers = list(pool.map(lambda q: _ask(client, q), searches))
            covered: dict[str, int] = {}
            next_frontier: list[tuple[str, str]] = []
            next_totals: dict[str, int] = {}
            for (suffix, parent), search, answer in zip(frontier, searches, answers, strict=True):
                covered[parent] = covered.get(parent, 0) + answer.total
                if answer.total <= _PAGE_SIZE:
                    found.extend(_complete(answer, search))
                    continue
                next_totals[suffix] = answer.total
                next_frontier.extend((f"{d}{suffix}", suffix) for d in _DIGITS)
            for parent, total in parent_totals.items():
                missing = total - covered.get(parent, 0)
                if missing:
                    found.extend(_exact_ids(client, pool, number, term, parent, missing))
            frontier, parent_totals = next_frontier, next_totals
    return found


def _exact_ids(
    client: httpx.Client,
    pool: ThreadPoolExecutor,
    number: int,
    term: str,
    suffix: str,
    missing: int,
) -> list[dict[str, Any]]:
    """The `missing` entries whose record id is exactly `suffix`, found by the
    push id's last digit."""
    searches = [f"{d} {suffix} {number} {term}" for d in _DIGITS]
    found: list[dict[str, Any]] = []
    for search, answer in zip(searches, pool.map(lambda q: _ask(client, q), searches), strict=True):
        found.extend(e for e in _complete(answer, search) if str(e["record_id"]) == suffix)
    if len(found) != missing:
        raise SystemExit(
            f"wago.tools search for record {suffix} ({number} {term}): "
            f"{len(found)} of {missing} entries"
        )
    return found


def fetch_push(
    build: str, tables: tuple[str, ...], client: httpx.Client | None = None
) -> dict[str, list[PushEntry]]:
    """The build's pushed entries for each of `tables` (enUS)."""
    own = client is None
    client = client or httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    number = _build_number(build)
    wanted = {t.lower(): t for t in tables}
    seen: dict[int, PushEntry] = {}
    table_of: dict[int, str] = {}
    try:
        for term in search_terms(tables):
            for e in _enumerate(client, number, term.lower()):
                table = wanted.get(str(e["table_name"]).lower())
                if table is None or e["build"] != number or e.get("locale", _LOCALE) != _LOCALE:
                    continue
                entry = PushEntry(
                    int(e["id"]),
                    int(e["push_id"]),
                    int(e["record_id"]),
                    int(e["status"]),
                    e.get("data") or None,
                )
                seen[entry.entry_id] = entry
                table_of[entry.entry_id] = table
    finally:
        if own:
            client.close()
    pushed: dict[str, list[PushEntry]] = {t: [] for t in tables}
    for entry_id, entry in seen.items():
        pushed[table_of[entry_id]].append(entry)
    return {
        t: sorted(v, key=lambda e: (e.record_id, e.push_id, e.entry_id)) for t, v in pushed.items()
    }


def push_totals(
    build: str, tables: tuple[str, ...], client: httpx.Client | None = None
) -> dict[str, int]:
    """How many entries wago lists for each of `tables`' search terms: one
    request a term. Every push adds entries, so a moved total is a new push."""
    own = client is None
    client = client or httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    number = _build_number(build)
    try:
        return {
            term: _ask(client, f"{number} {term.lower()}").total for term in search_terms(tables)
        }
    finally:
        if own:
            client.close()


def push_moved(record: Mapping[str, Any] | None, live: Mapping[str, int]) -> bool:
    """Whether wago lists a push the record was not layered from (or there is
    no record yet)."""
    return record is None or dict(record.get("search_totals", {})) != dict(live)


def read_push_record(build_dir: Path) -> dict[str, Any] | None:
    path = build_dir / PUSH_RECORD
    return json.loads(path.read_text(encoding="utf-8")) if path.exists() else None


def check_push(
    build: str, root: Path = Path("builds"), totals: TotalsFetcher = push_totals
) -> bool:
    """`python -m pipeline push-check`: whether the build's committed push
    record trails what wago lists now."""
    record = read_push_record(root / build)
    tables = tuple(record["tables"]) if record else PUSH_TABLES
    return push_moved(record, totals(build, tables))


def _sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def _header(path: Path) -> list[str]:
    with path.open(newline="", encoding="utf-8") as f:
        return next(csv.reader(f), [])


def _write_rows(rows: list[dict[str, str]], columns: list[str], path: Path) -> None:
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=columns, lineterminator="\n")
        writer.writeheader()
        writer.writerows(rows)


def _shipped_copy(raw: Path, table: str, layered_sha: str | None) -> Path:
    """The table as fetched: the kept copy when `raw/<table>.csv` is this
    module's own output, else the file a fresh fetch just wrote."""
    current, kept = raw / f"{table}.csv", raw / SHIPPED_DIR / f"{table}.csv"
    if not (kept.exists() and layered_sha == _sha256(current)):
        kept.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(current, kept)
    return kept


def _cache(push_dir: Path, pushed: Mapping[str, list[PushEntry]]) -> None:
    push_dir.mkdir(parents=True, exist_ok=True)
    for table, entries in pushed.items():
        text = json.dumps([asdict(e) for e in entries], indent=1)
        (push_dir / f"{table}.json").write_text(text + "\n", encoding="utf-8")


def _cached(push_dir: Path, tables: tuple[str, ...]) -> dict[str, list[PushEntry]]:
    pushed: dict[str, list[PushEntry]] = {}
    for table in tables:
        path = push_dir / f"{table}.json"
        if not path.exists():
            raise SystemExit(f"no cached push at {path}; run without --offline to fetch it")
        pushed[table] = [PushEntry(**e) for e in json.loads(path.read_text(encoding="utf-8"))]
    return pushed


def layer_hotfix_push(
    build: str,
    root: Path = Path("builds"),
    fetch: PushFetcher = fetch_push,
    offline: bool = False,
    tables: tuple[str, ...] = PUSH_TABLES,
    totals: TotalsFetcher = push_totals,
) -> dict[str, Any]:
    """`python -m pipeline layer-hotfix-push` (and the tail of `fetch`): layer
    the build's push into `<root>/<build>/raw/`, write the layer record to
    raw/ and `PUSH_RECORD` beside the build, and return the record.

    The totals are read before the push, so a push landing between the two
    reads shows as moved on the next check rather than as already layered."""
    raw = root / build / "raw"
    if not raw.is_dir():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    push_dir = raw / PUSH_DIR
    record_path = push_dir / LAYER_RECORD
    previous = json.loads(record_path.read_text()) if record_path.exists() else {}
    present = tuple(t for t in tables if (raw / f"{t}.csv").exists())
    if offline:
        search = dict(previous.get("search_totals", {}))
        pushed = _cached(push_dir, present)
    else:
        search = dict(totals(build, present))
        pushed = dict(fetch(build, present))
        _cache(push_dir, {t: pushed.get(t, []) for t in present})
    summary: dict[str, Any] = {"build": build, "search_totals": search, "tables": {}}
    for table in present:
        layered_sha = previous.get("tables", {}).get(table, {}).get("sha256")
        shipped_path = _shipped_copy(raw, table, layered_sha)
        rows, counts = layer_rows(
            read_csv(shipped_path), _header(shipped_path), pushed.get(table, []), table
        )
        target = raw / f"{table}.csv"
        if counts.changed == counts.added == counts.removed == 0:
            shutil.copyfile(shipped_path, target)
        else:
            _write_rows(rows, _header(shipped_path), target)
        summary["tables"][table] = {**asdict(counts), "sha256": _sha256(target)}
        if counts != LayerCounts():
            logger.info("layered the %s push over %s: %s", build, table, asdict(counts))
    text = json.dumps(summary, indent=1, sort_keys=True) + "\n"
    record_path.write_text(text, encoding="utf-8")
    (root / build / PUSH_RECORD).write_text(text, encoding="utf-8")
    return summary
