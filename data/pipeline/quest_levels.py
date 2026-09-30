"""`data/builds/<build>/raw/quests/quest-levels.json`: the merged,
committed, per-quest min_level/level table `pipeline.loot.sources`
reads -- the single file two upstream sources and one runtime fallback
all write into.

2026-09-28 quest-levels lane, final shape (three controller messages'
worth of narrowing -- see this lane's own report for the full story of
why wowhead alone, then classic-db alone, were each not quite enough):

* **classic-db** (`pipeline.classic_quest_levels`, GPL-3.0,
  cmangos/classic-db's `quest_template` table): the primary source, a
  one-time pinned download covering every quest id 1.12 Classic had --
  99.7% of THIS build's 720 quest-reward ids (718/720), because Forever
  reuses old Blizzard quest ids for old content almost entirely, even
  though the owner's own count (Forever adds 1,000+ NEW quests) is
  about the full ~5,198-quest Forever catalog, not this build's actual
  quest-REWARD set.
* **wowhead** (`pipeline.wowhead_quests`): the secondary source, for
  quest ids classic-db does not have (Forever-new quests) -- fetched
  ONLY for those ids, politely (one page per 5s, 429/403 honoured with
  backoff, a hard stop after 3 consecutive throttled responses or an
  explicit `--max-pages` budget), so this pipeline never asks more of
  wowhead than the 2 (of this build's 720) ids classic-db does not
  cover. Run from the nightly workflow's own GitHub-hosted runner
  (`.github/workflows/bis.yml`), never from a contributor's own address,
  so the ids beyond classic-db's coverage fill in over successive
  nights without concentrating requests on any one source IP.
* **item_level_proxy**: the runtime, PER-ITEM fallback
  (`pipeline.loot.sources._keyed_sources`) for a quest id neither file
  source covers -- computed from the AWARDED ITEM's own item_level, not
  the quest, so it is deliberately NOT written into this merged file
  (which is keyed by quest id only): two different items the same quest
  hands out could need two different proxy values, and a quest-keyed
  cache cannot represent that. `QuestSource.level_source` on the final
  per-item `loot.json` output is where "item_level_proxy" actually
  appears; this file's own `source` field is only ever "classic-db" or
  "wowhead".

A future fourth source -- the addon exporting the Forever-native level a
player's OWN quest log states (`GetQuestLogTitle`'s level return) with
its BiS/rotation export -- would write into this same file with
`source: "addon"`, keyed the same way, probably winning over classic-db
the way classic-db already wins over item_level_proxy (it is the
Forever server's own truth, not a 1.12 approximation); see this lane's
report for where that plugs in. Not implemented here.
"""

from __future__ import annotations

import json
import logging
from collections.abc import Iterable
from dataclasses import dataclass
from datetime import UTC, datetime
from pathlib import Path
from typing import Literal

from pydantic import BaseModel

from pipeline.classic_quest_levels import SOURCE_COMMIT, SOURCE_LICENSE, SOURCE_REPO, SOURCE_URL
from pipeline.classic_quest_levels import fetch_classic_db_quest_levels as _fetch_classic_db
from pipeline.wowhead_quests import fetch_quest_levels as _fetch_wowhead

logger = logging.getLogger(__name__)

FILE_NAME = "quest-levels.json"

QuestLevelSource = Literal["classic-db", "wowhead"]


class QuestLevelEntry(BaseModel):
    min_level: int
    level: int
    source: QuestLevelSource
    fetched_at: str
    #: "alliance", "horde" or "both" -- wowhead's own `side` field off
    #: the SAME quest-page fetch that resolved `min_level`/`level`
    #: (`pipeline.wowhead_quests.QuestPageLevel.faction`'s own doc).
    #: data-followups-3 lane, 2026-09-30, item 2: the second fallback
    #: `pipeline.loot.sources.build_loot`'s own doc calls for, behind
    #: classic-db's `RequiredRaces` (`quest_factions_from_classic_
    #: sources`) and ahead of publishing `"unknown"`. Always `None` for
    #: a `source: "classic-db"` entry (that source settles faction on
    #: its own, through the OTHER path) and for every entry fetched
    #: before this field existed -- a quest already in this file from an
    #: earlier `quest-levels` run keeps `None` here until a FUTURE
    #: refetch (this lane's own report: `fetch_missing_from_wowhead`
    #: only ever fetches an id NOT already covered, so an existing
    #: `source: "wowhead"` entry is never automatically backfilled;
    #: repinning/deleting its entry to force a refetch is a controller
    #: decision, not this pipeline's to make unprompted).
    faction: str | None = None


def raw_path(build_dir: Path) -> Path:
    """`data/builds/<build>/raw/quests/quest-levels.json` -- inside the
    otherwise-gitignored `raw/` tree, but committed anyway (see this
    module's own doc and `.gitignore`'s matching exception):
    `pipeline.loot.sources.build_loot` reads it on every `loot` run,
    including in CI, which has no engine checkout to regenerate it from
    and no business hitting GitHub or wowhead for data that changes
    only when this pipeline's own fetch steps run.
    """
    return build_dir / "raw" / "quests" / FILE_NAME


def load_quest_levels(build_dir: Path) -> dict[int, QuestLevelEntry]:
    """The committed merged table for one build, or `{}` (with a
    warning) when nothing has been fetched for this build yet -- `loot`
    still runs, every quest simply falls back to item_level_proxy, the
    same graceful degradation an individual missing id already gets."""
    path = raw_path(build_dir)
    if not path.exists():
        logger.warning(
            "quest-levels: no %s; run `python -m pipeline fetch-classic-quest-levels "
            "--build %s` first -- every quest falls back to item_level_proxy until then",
            path,
            build_dir.name,
        )
        return {}
    document = json.loads(path.read_text(encoding="utf-8"))
    return {int(entry): QuestLevelEntry(**fields) for entry, fields in document["quests"].items()}


def _write(build_dir: Path, entries: dict[int, QuestLevelEntry]) -> Path:
    path = raw_path(build_dir)
    path.parent.mkdir(parents=True, exist_ok=True)
    document = {
        "generated_at": datetime.now(UTC).isoformat(),
        "sources": {
            "classic-db": {
                "repo": SOURCE_REPO,
                "commit": SOURCE_COMMIT,
                "url": SOURCE_URL,
                "license": SOURCE_LICENSE,
                "table": "quest_template",
                "columns": ["entry", "MinLevel", "QuestLevel"],
            },
            "wowhead": {
                "note": "fetched per quest id for ids classic-db does not cover; see "
                "pipeline.wowhead_quests for the URL pattern and politeness rules.",
            },
        },
        "quests": {str(entry): entries[entry].model_dump() for entry in sorted(entries)},
    }
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


@dataclass(frozen=True)
class MergeStats:
    added: int
    total: int


def merge_classic_db(build: str, root: Path = Path("builds"), client=None) -> MergeStats:
    """`fetch-classic-quest-levels`: download the pinned classic-db dump
    and merge every entry into the committed `quest-levels.json`,
    tagged `source: "classic-db"`. A one-time (or occasional, only when
    `SOURCE_COMMIT` is deliberately repinned) step -- NOT part of the
    nightly workflow, which only ever runs `fetch_missing_from_wowhead`
    below. `client` is test-only (an `httpx.Client` over a
    `MockTransport`); production callers always leave it `None`.
    """
    build_dir = root / build
    entries = load_quest_levels(build_dir)
    classic = _fetch_classic_db(client=client)
    now = datetime.now(UTC).isoformat()
    added = 0
    for entry, quest in classic.items():
        entries[entry] = QuestLevelEntry(
            min_level=quest.min_level, level=quest.level, source="classic-db", fetched_at=now
        )
        added += 1
    _write(build_dir, entries)
    logger.info("quest-levels: merged %d classic-db entries into %s", added, raw_path(build_dir))
    return MergeStats(added=added, total=len(entries))


@dataclass(frozen=True)
class WowheadMergeStats:
    fetched: int
    still_missing: int
    needed: int


def fetch_missing_from_wowhead(
    build: str,
    quest_ids: Iterable[int],
    max_pages: int | None,
    root: Path = Path("builds"),
    delay: float | None = None,
    client=None,
) -> WowheadMergeStats:
    """`python -m pipeline quest-levels`: the nightly step. Fetches
    wowhead pages ONLY for ids in `quest_ids` that `quest-levels.json`
    does not already cover (from EITHER source), capped at `max_pages`
    live requests, and merges any it gets tagged `source: "wowhead"`.
    Ids still missing after this call simply fall back to
    item_level_proxy at `loot` time, same as always -- there is no
    error path here, by design (see `pipeline.wowhead_quests.
    fetch_quest_levels`'s own doc: this always returns normally,
    whether it fetched everything, hit the page budget, or got
    throttled). `delay`/`client` are test-only (override the real
    politeness delay / inject an `httpx.Client` over a `MockTransport`);
    production callers always leave both `None`.
    """
    build_dir = root / build
    entries = load_quest_levels(build_dir)
    needed = sorted(q for q in set(quest_ids) if q not in entries)
    if not needed:
        logger.info("quest-levels: every needed quest id is already covered; nothing to fetch")
        return WowheadMergeStats(fetched=0, still_missing=0, needed=0)
    kwargs = {}
    if delay is not None:
        kwargs["delay"] = delay
    if client is not None:
        kwargs["client"] = client
    result = _fetch_wowhead(build, needed, root=root, max_pages=max_pages, **kwargs)
    now = datetime.now(UTC).isoformat()
    for entry, quest in result.levels.items():
        entries[entry] = QuestLevelEntry(
            min_level=quest.min_level,
            level=quest.level,
            source="wowhead",
            fetched_at=now,
            faction=quest.faction,
        )
    _write(build_dir, entries)
    still_missing = len(needed) - len(result.levels)
    logger.info(
        "quest-levels: wowhead filled %d/%d needed ids (%d still missing, item_level_proxy "
        "fallback applies)",
        len(result.levels),
        len(needed),
        still_missing,
    )
    return WowheadMergeStats(
        fetched=len(result.levels), still_missing=still_missing, needed=len(needed)
    )
