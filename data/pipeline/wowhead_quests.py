"""Wowhead's Forever quest pages -- SPOT-CHECK VERIFICATION ONLY.

2026-09-28 quest-levels lane, superseded as the primary quest-level
source by `pipeline.classic_quest_levels` (read that module's doc
first): the Forever client itself states no quest level anywhere
`wago.tools` exports, and wowhead throttles this pipeline's address
(this module's own `THROTTLE_STOP_AFTER` handling exists because of
that -- a 403 storm at id 873, then again at id 9248 even after backing
off to 4s between requests, matching the owner's own browser getting
rate-limited from the same address at the same time).

This module now exists ONLY for `python -m pipeline verify-wowhead-quests`
(never run in CI, never called by `loot`): a small, polite, resumable
sample fetch of live wowhead pages, compared against the committed
`classic-quest-levels.json` to spot-check that source's own agreement
rate. `data/tests/fixtures/wowhead-quest-samples/` (plus
`wowhead-quest-53.html`) are ten of the 414 pages this lane fetched
before the throttle, kept as committed regression fixtures for that
comparison and for this module's own parsing tests; the other 404 were
deleted (raw/quests/*.html is gitignored, so nothing here was ever
going to keep them anyway).

Wowhead's Forever quest pages carry both numbers a quest itself states:
`reqlevel` (the minimum level to pick the quest up) and `level` (the
level the quest is meant for / rewards scale to) -- see
`https://www.wowhead.com/forever/quest=<id>`'s own
`$.extend(g_quests[<id>], {...})` line. There is no bulk export of this
data the way the gear planner has for items (`pipeline.wowhead_items`):
a `?filter=...` listview page always truncates to (at most) 1,000 of
Forever's 5,198 quests, sorted by popularity, and accepts no by-id
filter that would reach a chosen sample directly.
"""

from __future__ import annotations

import json
import logging
import time
from collections.abc import Iterable
from dataclasses import dataclass
from pathlib import Path

import httpx
from pydantic import BaseModel

from pipeline.classic_quest_levels import ClassicQuestLevel
from pipeline.wago import USER_AGENT
from pipeline.wowhead_item_sources import QUEST_SIDE_FACTION

logger = logging.getLogger(__name__)

QUEST_URL = "https://www.wowhead.com/forever/quest={quest_id}"
CACHE_SUBDIR = "quests"

#: Politeness delay between LIVE fetches (a cache hit makes no request
#: and sleeps for none of this). wowhead_items.py's own fetch is a
#: single bulk request and states no per-item delay because it never
#: loops; this module does loop, so it needs one. This lane's own
#: probing measured the real limit: 0.4s (150 quests/minute) got every
#: quest page id=873 onward 403'd for roughly ten minutes straight; even
#: 1.5s (40/minute) eventually drew 403s again at id=9248 after 414
#: successful fetches -- the owner's own browser was rate-limited from
#: the same address at the same time, so this is a shared-IP limit, not
#: something a per-run delay alone can out-wait. The controller's own
#: floor: one page per 5s minimum.
FETCH_DELAY_SECONDS = 5.0

#: A 403 or 429 is wowhead SAYING "slow down", not "this id has no
#: page" (see QuestPageNotFound for that) -- retried with backoff
#: rather than recorded as missing.
THROTTLE_STATUS_CODES = frozenset({403, 429})

#: Backoff after a throttled (403/429) response: doubles each
#: consecutive throttle, starting at 30s, capped at 5 minutes. A
#: `Retry-After` header (429's own way of stating the wait) overrides
#: the computed value when it asks for longer.
THROTTLE_BACKOFF_START_SECONDS = 30.0
THROTTLE_BACKOFF_CAP_SECONDS = 300.0

#: Stop the whole run (returning normally -- exit 0, not an exception)
#: after this many THROTTLED RESPONSES IN A ROW, even across backoffs:
#: wowhead is telling this address to stop, and the remaining ids stay
#: uncached for the nightly (or a later manual run) to pick up once the
#: limit lifts, rather than this run hammering it for longer.
THROTTLE_STOP_AFTER = 3


class QuestPageLevel(BaseModel):
    """One quest's level fields, as wowhead's own page states them."""

    quest_id: int
    name: str
    #: wowhead's `reqlevel`: the minimum level to pick the quest up.
    min_level: int
    #: wowhead's `level`: the level the quest (and its rewards) are
    #: written for.
    level: int
    #: "alliance", "horde" or "both" -- wowhead's own `side` field on the
    #: SAME `g_quests[<id>]` payload `min_level`/`level` already come
    #: from (`QUEST_SIDE_FACTION`'s own doc: `pipeline.
    #: wowhead_item_sources`'s `reward-from-q` listview row states the
    #: identical fact on a different page). data-followups-3 lane,
    #: 2026-09-30, item 2: the quest's own client-side race/faction DB2
    #: table (`QuestV2`/`QuestInfo`'s own allowable-races mask, whichever
    #: names it -- not fetched by this pipeline; no build under
    #: `data/builds/<build>/raw/quests/` carries it) is the PRIMARY
    #: source `pipeline.loot.sources.build_loot`'s own doc calls for;
    #: this is the fallback for a Forever-new quest neither that table
    #: nor classic-db's `quest_template.RequiredRaces` covers -- see
    #: `pipeline.quest_levels.QuestLevelEntry.faction`'s own doc for the
    #: third fallback (`"unknown"`) when this is absent too.
    faction: str | None = None


class QuestPageNotFound(Exception):
    """Wowhead served a page with no `g_quests[<id>]` row for this quest
    -- a 404 (this build's quest ids run well past what Forever's own
    Wowhead section has indexed) or a redirect to an unrelated page."""


def raw_path(build_dir: Path, quest_id: int) -> Path:
    return build_dir / "raw" / CACHE_SUBDIR / f"{quest_id}.html"


def parse_quest_page(quest_id: int, html: str) -> QuestPageLevel:
    """The `$.extend(g_quests[<id>], {...})` object off a quest page,
    the same marker-then-`raw_decode` technique
    `pipeline.wowhead_items.parse_page_data` uses for the gear planner's
    payload, which tolerates the object containing further nested
    braces (`envChange`) that a regex matched to the wrong `}` would
    not.
    """
    marker = f"$.extend(g_quests[{quest_id}],"
    at = html.find(marker)
    if at < 0:
        raise QuestPageNotFound(f"quest {quest_id}: no {marker!r} in page")
    start = html.index("{", at)
    try:
        data, _ = json.JSONDecoder().raw_decode(html[start:])
    except json.JSONDecodeError as exc:
        raise QuestPageNotFound(f"quest {quest_id}: unparseable g_quests payload: {exc}") from exc
    if not isinstance(data, dict):
        raise QuestPageNotFound(f"quest {quest_id}: g_quests payload is not an object")
    raw_side = data.get("side")
    return QuestPageLevel(
        quest_id=quest_id,
        name=str(data.get("name", "")),
        min_level=int(data.get("reqlevel") or 0),
        level=int(data.get("level") or 0),
        faction=QUEST_SIDE_FACTION.get(int(raw_side)) if raw_side is not None else None,
    )


def fetch_quest_page(quest_id: int, client: httpx.Client) -> str:
    response = client.get(QUEST_URL.format(quest_id=quest_id), timeout=30, follow_redirects=True)
    response.raise_for_status()
    return response.text


@dataclass(frozen=True)
class QuestLevelResult:
    """`fetch_quest_levels`/`load_cached_quest_levels`'s return: every
    quest id resolved, and every one that was not (no cached page and,
    for `fetch_quest_levels`, no live page either) -- the caller applies
    `item_level_proxy` to `missing`, since only it knows the item's own
    item level to fall back on.
    """

    levels: dict[int, QuestPageLevel]
    missing: set[int]


def _read_cached(build_dir: Path, quest_id: int) -> QuestPageLevel | None:
    path = raw_path(build_dir, quest_id)
    if not path.exists():
        return None
    try:
        return parse_quest_page(quest_id, path.read_text(encoding="utf-8"))
    except QuestPageNotFound as exc:
        logger.warning("wowhead quest cache %s: %s", path, exc)
        return None


def load_cached_quest_levels(
    build: str, quest_ids: Iterable[int], root: Path = Path("builds")
) -> QuestLevelResult:
    """Every quest id this build's cache already has a page for, with NO
    network calls -- what `pipeline.loot.sources.build_loot` reads. A
    quest id with no cached page is `missing`, not fetched: `loot`
    builds the site's committed data and must run offline in CI the same
    way `simdb`/`normalize` do; only the explicit
    `python -m pipeline fetch-wowhead-quests` step (`fetch_quest_levels`,
    below) is allowed to reach the network.
    """
    build_dir = root / build
    levels: dict[int, QuestPageLevel] = {}
    missing: set[int] = set()
    for quest_id in sorted(set(quest_ids)):
        found = _read_cached(build_dir, quest_id)
        if found is None:
            missing.add(quest_id)
        else:
            levels[quest_id] = found
    return QuestLevelResult(levels=levels, missing=missing)


def _retry_after_seconds(exc: httpx.HTTPStatusError) -> float | None:
    header = exc.response.headers.get("Retry-After")
    if header is None:
        return None
    try:
        return float(header)
    except ValueError:
        return None  # Retry-After can also be an HTTP-date; not seen from wowhead yet


def fetch_quest_levels(
    build: str,
    quest_ids: Iterable[int],
    root: Path = Path("builds"),
    client: httpx.Client | None = None,
    delay: float = FETCH_DELAY_SECONDS,
    max_pages: int | None = None,
) -> QuestLevelResult:
    """Fetch (live) every quest id this build's cache does not already
    have, caching each page under `builds/<build>/raw/quests/<id>.html`
    so a rerun -- including a later call to THIS function, which skips
    every id already cached -- picks up exactly where the last one
    stopped.

    `max_pages`, when given, caps the number of LIVE requests this call
    sends (a cache hit costs nothing toward it) -- the nightly budget
    (`python -m pipeline quest-levels --max-pages 200`, see
    `pipeline.quest_levels`): once reached, the run stops and returns
    normally with every id not yet fetched in `.missing`, the same as a
    throttle stop, for a later run to pick up.

    A quest id wowhead 404s (or otherwise serves no `g_quests[<id>]` row
    for) is recorded in `.missing` and NOT cached, so a later rerun
    retries it rather than remembering a miss forever (Forever's Wowhead
    section is actively being filled in; a quest missing today may have
    a page next week).

    A 403/429 is different: it means THIS ADDRESS is throttled, not
    that the id has no page, so it is retried with backoff
    (THROTTLE_BACKOFF_START_SECONDS, doubling, capped at
    THROTTLE_BACKOFF_CAP_SECONDS, honouring a `Retry-After` header when
    the response states one) rather than being recorded as missing.
    THROTTLE_STOP_AFTER throttled responses in a row -- even across
    backoffs, even across different ids -- stops the whole run and
    returns normally (never raises): every id not yet cached is
    `.missing`, for the nightly or a later manual run to pick up once
    the limit lifts. 2026-09-28: this is exactly what happened fetching
    this build's 720 ids -- 414 fetched, then id 9248 onward 403'd for
    the owner's own browser too (a shared-IP limit, not something this
    run alone caused or can wait out reliably).
    """
    build_dir = root / build
    own = client is None
    client = client or httpx.Client(headers={"User-Agent": USER_AGENT})
    levels: dict[int, QuestPageLevel] = {}
    missing: set[int] = set()
    pending = sorted(set(quest_ids))
    consecutive_throttled = 0
    backoff = THROTTLE_BACKOFF_START_SECONDS
    stopped_early_at: int | None = None
    live_requests = 0
    try:
        index = 0
        while index < len(pending):
            quest_id = pending[index]
            cached = _read_cached(build_dir, quest_id)
            if cached is not None:
                levels[quest_id] = cached
                index += 1
                continue
            if max_pages is not None and live_requests >= max_pages:
                logger.info(
                    "wowhead quests: stopped after the %d-page budget; %d of %d ids "
                    "remain unfetched",
                    max_pages,
                    len(pending) - index,
                    len(pending),
                )
                stopped_early_at = index
                break
            live_requests += 1
            try:
                html = fetch_quest_page(quest_id, client)
            except httpx.HTTPStatusError as exc:
                status = exc.response.status_code
                if status in THROTTLE_STATUS_CODES:
                    consecutive_throttled += 1
                    logger.warning(
                        "wowhead quest %d: throttled (%d), %d/%d consecutive",
                        quest_id,
                        status,
                        consecutive_throttled,
                        THROTTLE_STOP_AFTER,
                    )
                    if consecutive_throttled >= THROTTLE_STOP_AFTER:
                        stopped_early_at = index
                        break
                    wait = backoff
                    retry_after = _retry_after_seconds(exc)
                    if retry_after is not None:
                        wait = max(wait, retry_after)
                    time.sleep(min(wait, THROTTLE_BACKOFF_CAP_SECONDS))
                    backoff = min(backoff * 2, THROTTLE_BACKOFF_CAP_SECONDS)
                    continue  # retry the SAME quest_id, index unchanged
                logger.warning("wowhead quest %d: %s", quest_id, exc)
                missing.add(quest_id)
                consecutive_throttled = 0
                backoff = THROTTLE_BACKOFF_START_SECONDS
                time.sleep(delay)
                index += 1
                continue
            try:
                parsed = parse_quest_page(quest_id, html)
            except QuestPageNotFound as exc:
                logger.warning("wowhead quest %d: %s", quest_id, exc)
                missing.add(quest_id)
                consecutive_throttled = 0
                backoff = THROTTLE_BACKOFF_START_SECONDS
                time.sleep(delay)
                index += 1
                continue
            path = raw_path(build_dir, quest_id)
            path.parent.mkdir(parents=True, exist_ok=True)
            path.write_text(html, encoding="utf-8")
            levels[quest_id] = parsed
            consecutive_throttled = 0
            backoff = THROTTLE_BACKOFF_START_SECONDS
            time.sleep(delay)
            index += 1
    finally:
        if own:
            client.close()
    if stopped_early_at is not None:
        remaining = [q for q in pending[stopped_early_at:] if q not in levels]
        missing.update(remaining)
        if consecutive_throttled >= THROTTLE_STOP_AFTER:
            logger.warning(
                "wowhead quests: stopped after %d consecutive throttled responses; %d of %d "
                "ids remain unfetched (item_level_proxy fallback applies; rerun later to "
                "resume from the warm cache)",
                THROTTLE_STOP_AFTER,
                len(remaining),
                len(pending),
            )
        # else: the max_pages budget stop already logged its own message
        # at the point it triggered, above.
    elif missing:
        logger.info(
            "wowhead quests: %d of %d ids had no usable page (item_level_proxy fallback applies)",
            len(missing),
            len(levels) + len(missing),
        )
    return QuestLevelResult(levels=levels, missing=missing)


@dataclass(frozen=True)
class ComparisonStats:
    """`compare_with_classic_db`'s report: how often a live wowhead page's
    min_level/level agree with `classic-quest-levels.json` for the same
    quest ids -- `verify-wowhead-quests`' whole reason to exist."""

    compared: int
    min_level_agrees: int
    level_agrees: int
    both_agree: int
    #: (quest_id, wowhead_min, wowhead_level, classic_min, classic_level)
    #: for every quest either field disagreed on.
    mismatches: list[tuple[int, int, int, int, int]]


def compare_with_classic_db(
    wowhead_levels: dict[int, QuestPageLevel],
    classic_levels: dict[int, ClassicQuestLevel],
) -> ComparisonStats:
    """How often `wowhead_levels` (a live/cached fetch) agrees with
    `classic_levels` (`pipeline.classic_quest_levels.load_classic_quest_levels`'s
    result) for the quest ids both name. A quest wowhead has a page for
    but classic-db does not (a Forever-new quest) is not compared --
    there is nothing to agree or disagree with.
    """
    compared = min_agree = level_agree = both_agree = 0
    mismatches: list[tuple[int, int, int, int, int]] = []
    for quest_id, wowhead in wowhead_levels.items():
        classic = classic_levels.get(quest_id)
        if classic is None:
            continue
        compared += 1
        m_ok = wowhead.min_level == classic.min_level
        l_ok = wowhead.level == classic.level
        if m_ok:
            min_agree += 1
        if l_ok:
            level_agree += 1
        if m_ok and l_ok:
            both_agree += 1
        else:
            mismatches.append(
                (quest_id, wowhead.min_level, wowhead.level, classic.min_level, classic.level)
            )
    return ComparisonStats(
        compared=compared,
        min_level_agrees=min_agree,
        level_agrees=level_agree,
        both_agree=both_agree,
        mismatches=sorted(mismatches),
    )
