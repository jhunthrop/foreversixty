"""`data/builds/<build>/raw/items/item-sources.json`: the committed,
resumable cache `pipeline.loot.sources.build_loot` reads to fill in a
source for an item the engine fork's own database states none for --
see `pipeline.wowhead_item_sources`'s own doc for what gets scraped and
why only four of wowhead's item-page listviews are read.

Same committed-cache shape as `pipeline.quest_levels`' quest-levels.json
(read that module's doc first): a nightly step
(`python -m pipeline item-sources`, `.github/workflows/bis.yml`) fetches
ONLY the item ids `pipeline.loot.sources.unsourced_real_item_ids` names
for the active build that this file does not already cover, politely
and capped at `--max-pages`, and merges whatever it gets into this file.
`loot` itself never touches the network -- it reads this committed file
the same way it reads quest-levels.json, so CI (which has no business
hitting wowhead for data that changes only when the fetch step runs) can
build loot.json offline.

Nothing here decides whether a fetched-but-empty result (wowhead had a
page, named no dropped-by/sold-by/created-by/reward-from-q row) counts
as "resolved" -- it is written to the cache either way (so a later run
does not refetch it for nothing), and `pipeline.loot.sources` is what
decides an empty entry contributes no LootSource.
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

from pipeline.wowhead_item_sources import CraftedSource, NpcSource, QuestRewardSource
from pipeline.wowhead_item_sources import fetch_item_sources as _fetch_wowhead

logger = logging.getLogger(__name__)

FILE_NAME = "item-sources.json"

ItemSourceOrigin = Literal["wowhead"]


class ItemSourceEntry(BaseModel):
    """One item id's merged record in the committed cache -- the same
    four lists `pipeline.wowhead_item_sources.ItemPageSources` parses a
    single page into, plus provenance. `source` is a single-value
    `Literal` today (wowhead is the only scrape this pipeline has), kept
    for the same reason `pipeline.quest_levels.QuestLevelEntry.source`
    is one: a future source (that module's own doc names the addon
    export as the likely next one) adds a value here rather than a new
    file shape.
    """

    dropped_by: list[NpcSource] = []
    sold_by: list[NpcSource] = []
    crafted_by: list[CraftedSource] = []
    quest_rewards: list[QuestRewardSource] = []
    #: `pipeline.wowhead_item_sources.ItemPageSources`' own two fields of
    #: the same name -- carried through the merge unchanged (vendor-11036
    #: lane, 2026-09-30).
    required_faction_id: int | None = None
    required_standing_raw: int | None = None
    source: ItemSourceOrigin
    fetched_at: str


def raw_path(build_dir: Path) -> Path:
    """`data/builds/<build>/raw/items/item-sources.json` -- committed
    despite living inside the otherwise-gitignored `raw/` tree, the same
    exception `.gitignore` already carries for `raw/quests/
    quest-levels.json` and for the same reason: `loot` reads it in CI,
    which has no engine checkout or business reaching wowhead itself."""
    return build_dir / "raw" / "items" / FILE_NAME


def load_item_sources(build_dir: Path) -> dict[int, ItemSourceEntry]:
    """The committed cache for one build, or `{}` (with a warning) when
    nothing has been fetched for this build yet -- `loot` still runs,
    every item this cache does not cover simply stays unsourced, the
    same graceful degradation quest-levels.json's absence already gets
    in `pipeline.quest_levels.load_quest_levels`."""
    path = raw_path(build_dir)
    if not path.exists():
        logger.warning(
            "item-sources: no %s; run `python -m pipeline item-sources` first -- every "
            "item this cache does not cover stays unsourced until then",
            path,
        )
        return {}
    document = json.loads(path.read_text(encoding="utf-8"))
    return {int(item_id): ItemSourceEntry(**fields) for item_id, fields in document["items"].items()}  # noqa: E501


def _write(build_dir: Path, entries: dict[int, ItemSourceEntry]) -> Path:
    path = raw_path(build_dir)
    path.parent.mkdir(parents=True, exist_ok=True)
    document = {
        "generated_at": datetime.now(UTC).isoformat(),
        "source": {
            "wowhead": {
                "note": "scraped per item id, ONLY for ids the fork database names no "
                "source for at all; see pipeline.wowhead_item_sources for the URL pattern, "
                "the four listviews read, and the politeness rules.",
            },
        },
        "items": {str(item_id): entries[item_id].model_dump() for item_id in sorted(entries)},
    }
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


@dataclass(frozen=True)
class ItemSourceMergeStats:
    fetched: int
    still_missing: int
    needed: int
    #: src-crawl-order lane, 2026-09-29: of `fetched`, how many pages
    #: actually named a dropped-by/sold-by/crafted-by/quest-reward row
    #: (`ItemPageSources.is_empty()` false) rather than being cached as
    #: an empty "wowhead had a page, it named nothing this pipeline
    #: reads" result. First 440 Classic-id pages measured 425/440; this
    #: is how the crawl learns the same number for Forever-new ids.
    with_source: int


def fetch_missing_from_wowhead(
    build: str,
    item_ids: Iterable[int],
    max_pages: int | None,
    root: Path = Path("builds"),
    delay: float | None = None,
    client=None,
) -> ItemSourceMergeStats:
    """`python -m pipeline item-sources`: the nightly step. Fetches
    wowhead pages ONLY for ids in `item_ids` that `item-sources.json`
    does not already cover, capped at `max_pages` live requests, and
    merges every result (including an empty one -- see this module's own
    doc) into the committed cache. `delay`/`client` are test-only
    (override the politeness delay / inject an `httpx.Client` over a
    `MockTransport`); production callers always leave both `None`.

    `item_ids`' own order is preserved into `needed` (deduplicated, and
    filtered to ids `item-sources.json` does not already cover) rather
    than re-sorted ascending by id -- callers such as
    `pipeline.loot.wowhead.unsourced_real_item_ids` order the backlog by
    crawl priority (Forever-new before Classic, required_level
    ascending), and an ascending-id `sorted()` here would silently
    discard that ordering before it ever reached the live crawl.
    """
    build_dir = root / build
    entries = load_item_sources(build_dir)
    seen: set[int] = set()
    needed: list[int] = []
    for item_id in item_ids:
        if item_id in entries or item_id in seen:
            continue
        seen.add(item_id)
        needed.append(item_id)
    if not needed:
        logger.info("item-sources: every needed item id is already covered; nothing to fetch")
        return ItemSourceMergeStats(fetched=0, still_missing=0, needed=0, with_source=0)
    kwargs = {}
    if delay is not None:
        kwargs["delay"] = delay
    if client is not None:
        kwargs["client"] = client
    result = _fetch_wowhead(build, needed, root=root, max_pages=max_pages, **kwargs)
    now = datetime.now(UTC).isoformat()
    merged = {
        **entries,
        **{
            item_id: ItemSourceEntry(
                dropped_by=page.dropped_by,
                sold_by=page.sold_by,
                crafted_by=page.crafted_by,
                quest_rewards=page.quest_rewards,
                required_faction_id=page.required_faction_id,
                required_standing_raw=page.required_standing_raw,
                source="wowhead",
                fetched_at=now,
            )
            for item_id, page in result.sources.items()
        },
    }
    _write(build_dir, merged)
    still_missing = len(needed) - len(result.sources)
    with_source = sum(1 for page in result.sources.values() if not page.is_empty())
    logger.info(
        "item-sources: wowhead resolved a page for %d/%d needed ids (%d of those fetched "
        "pages named ANY source; %d still missing; rerun later to resume from the warm "
        "cache)",
        len(result.sources),
        len(needed),
        with_source,
        still_missing,
    )
    return ItemSourceMergeStats(
        fetched=len(result.sources),
        still_missing=still_missing,
        needed=len(needed),
        with_source=with_source,
    )
