"""Wowhead's Forever item pages -- the "Dropped by"/"Sold by"/"Created
by"/"Reward from" listviews for an item the engine fork's own
`assets/database/db.json` (`pipeline.forkdb`) states NO source for at
all: 6,183 distinct real (non-placeholder) items on build 1.60.1.70009
(night-item-sources lane, 2026-09-28), most of them pre-existing Classic
ids the fork's own AtlasLoot-derived database simply never itemised,
not only the 1,797 Forever-new (27xxxx-28xxxx) ids the fork's database
cannot possibly know about.

Same shape as `pipeline.wowhead_quests` (read that module's doc first
for why the politeness rules exist and are non-negotiable): one request
per `FETCH_DELAY_SECONDS`, a cache hit costs nothing, a 403/429 means
THIS ADDRESS is throttled and is retried with backoff, and
`THROTTLE_STOP_AFTER` throttled responses in a row stops the whole run
(returning normally) rather than hammering wowhead for longer. The
`--max-pages` budget and resumable `raw/items/<id>.html` cache are
identical in spirit; only the URL and the parsed shape differ.

Parsing technique: each listview wowhead cares about here appears
exactly once on the page as `new Listview({... id: '<listview-id>',
... data: [...]})`. Rather than regex-matching the whole object (whose
body can nest further braces the way `pipeline.wowhead_quests`'s own
`g_quests` payload does with `envChange`), this finds `id: '<id>',` as
an anchor, finds the `data: [` that follows it, and hands
`json.JSONDecoder().raw_decode` the substring starting at that `[` --
the same marker-then-raw_decode technique `pipeline.wowhead_items.
parse_page_data` and `pipeline.wowhead_quests.parse_quest_page` both
use, which tolerates arbitrarily nested brackets inside the array.

Four listviews are read, chosen because each maps onto an EXISTING
`loot.json` source kind (`pipeline.loot.sources`) with no invention
needed:

* `dropped-by` (template `npc`) -> `dungeon`/`raid`/`world`+`zone`,
  exactly like the fork's own per-drop `{"drop": {"zoneId", "npcId"}}`
  entries: each row's own `location` field is a list of AreaTable zone
  ids, the SAME numbering `zones.json` and the fork's `assets/database/
  db.json` both use (verified against build 1.60.1.70009: item 16769's
  wowhead `sold-by` row states npc id 11555, and the fork's own db.json
  sources that exact item to `vendor:11555` -- same npc id, same
  numbering, two independent sources). A mob with more than one
  `location` entry (roams several zones) is filed under only the
  FIRST -- the fork's own per-drop shape has no way to record more than
  one zone per drop either, so this loses nothing the fork model could
  have kept.
* `sold-by` (template `npc`) -> `vendor`.
* `created-by-spell` (template `spell`) -> `crafted`, via
  `SKILL_LINE_PROFESSIONS` (wowhead's numeric skill-line id -> this
  pipeline's profession vocabulary, `pipeline.forkdb.PROFESSIONS`'
  string keys). A recipe naming a skill line not in that map (Cooking,
  First Aid, Fishing -- none of which `loot.json`'s `crafted` kind
  covers for any item, fork-sourced or not) is simply not resolved to a
  profession; the item stays unsourced rather than being filed under a
  guessed one.
* `reward-from-q` (template `quest`) -> `quest`, and MORE directly than
  the fork path: wowhead's own row states `reqlevel`/`level` (this
  pipeline's `min_level`/`level`, see `pipeline.quest_levels`) and
  `side` (1 alliance, 2 horde, 3 or absent both) on the SAME row that
  names the item, so no separate quest-levels.json lookup or
  factionRestriction fallback is needed the way `_keyed_sources` needs
  for a fork-sourced quest item; `level_source` is written `"wowhead"`
  either way, since it did come from this scrape.

Two listviews wowhead shows that are NOT read here, on purpose, because
neither maps cleanly onto an existing kind without guessing:

* `contained-in-item`/`contained-in-object` (a chest or crate): the
  listview names the CONTAINER item, not a zone or npc -- resolving one
  would mean fetching the container's own page too (and the container
  might itself be a drop, a vendor sale, or worse, unsourced), which
  this module does not attempt. An item found only inside a container
  stays unsourced.
* Reputation-vendor purchases: wowhead's `reward-from-q` row's own
  `reprewards` field states the REP GAINED for turning the quest in,
  not a standing GATE on a vendor sale, and no `sold-by` row this lane
  observed carried a standing requirement either (2026-09-28: checked
  against three already fork-sourced `rep`-kind items, e.g. item 12185
  -- Bloodsail Buccaneers, faction id 87 in wowhead's own numbering vs.
  21 in the fork's -- which wowhead's page states only as a QUEST
  reward, never as a rep-gated vendor sale). This pipeline has no
  reliable wowhead signal for the `rep` kind and does not guess one;
  see this lane's report for the recommendation (a future pass keyed on
  the fork's OWN rep data, not wowhead, since wowhead's faction ids do
  not agree with the fork's).
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

from pipeline.wago import USER_AGENT

logger = logging.getLogger(__name__)

ITEM_URL = "https://www.wowhead.com/forever/item={item_id}"
CACHE_SUBDIR = "items"

#: See `pipeline.wowhead_quests`'s own constant for the measurement this
#: is pinned to; the same address is shared between both crawlers, so
#: the same floor applies.
FETCH_DELAY_SECONDS = 5.0
THROTTLE_STATUS_CODES = frozenset({403, 429})
THROTTLE_BACKOFF_START_SECONDS = 30.0
THROTTLE_BACKOFF_CAP_SECONDS = 300.0
THROTTLE_STOP_AFTER = 3

#: wowhead's numeric skill-line id -> `pipeline.forkdb.PROFESSIONS`'
#: string vocabulary. Only the 9 professions `loot.json`'s `crafted`
#: kind ever names (contract 6.1); Cooking (185), First Aid (129) and
#: Fishing (356) are deliberately absent -- see this module's own doc.
SKILL_LINE_PROFESSIONS: dict[int, str] = {
    164: "blacksmithing",
    171: "alchemy",
    182: "herbalism",
    186: "mining",
    197: "tailoring",
    202: "engineering",
    333: "enchanting",
    393: "skinning",
    165: "leatherworking",
}

#: wowhead's `side` field on a quest row: 1 Alliance, 2 Horde, 3 both.
#: Matches `pipeline.loot.sources.QUEST_FACTION_BY_RESTRICTION`'s own
#: vocabulary (minus the numbering, which is wowhead's own and unrelated
#: to the fork's `factionRestriction` column). Public (not `_`-prefixed):
#: data-followups-3 lane, 2026-09-30, item 2: `pipeline.wowhead_quests`
#: reads the SAME `side` field off the quest PAGE's own `g_quests[<id>]`
#: payload (not this module's `reward-from-q` listview row, a different
#: place on a different page wowhead states the identical fact), so this
#: is shared rather than duplicated.
QUEST_SIDE_FACTION = {1: "alliance", 2: "horde", 3: "both"}


class NpcSource(BaseModel):
    """One `dropped-by` or `sold-by` row: an npc, where it is, and the
    level range and (dropped-by only, most rows) drop chance wowhead
    states for it.

    wowhead-world-drops lane, 2026-09-29: `min_level`/`max_level`
    (the row's own `minlevel`/`maxlevel`) and `chance` (`100 *
    count/outof`, wowhead's own "N of M recorded kills" tally, rounded
    to 4 decimal places) feed `pipeline.loot.wowhead`'s world-drop
    pattern detector -- a `dropped-by` list naming dozens of creatures
    across a dozen zones, each with a chance under 1%, is a generic BoE
    world drop, not a source a player recognises (tenet 7). `sold-by`
    rows carry `minlevel`/`maxlevel` too (harmless to keep) but never
    `count`/`outof` (wowhead states a vendor's stock/cost there
    instead), so `chance` is always `None` for one. Every committed
    `item-sources.json` entry fetched before this lane predates these
    three fields and simply has them `None` -- the pattern detector's
    other two signals (creature count, zone count) still fire without
    them."""

    npc_id: int
    name: str
    zone_ids: list[int]
    min_level: int | None = None
    max_level: int | None = None
    chance: float | None = None
    #: data-followups-9 lane, 2026-09-30: a `sold-by` row's own
    #: `ItemExtendedCost`-shaped `cost` field -- `[money_copper,
    #: currency_pairs, item_pairs]` on wowhead's own listview row. `0`
    #: when the row states no gold price (every row this lane measured
    #: on vendor 11036/240248 is `0`: the whole price is the item cost
    #: below). `dropped-by` rows never carry `cost` at all, so this is
    #: always `0` for one.
    cost_money: int = 0
    #: The item id half of every `[item_id, qty]` pair in the same
    #: `cost` field's item-cost list (quantity is not kept -- only
    #: WHETHER the id itself resolves to a real, obtainable item matters
    #: to `pipeline.loot.wowhead`'s gate). Empty for a row with no
    #: ItemExtendedCost override (an ordinary gold buy) or no `cost` at
    #: all (`dropped-by`).
    cost_item_ids: list[int] = []


class CraftedSource(BaseModel):
    profession: str


class QuestRewardSource(BaseModel):
    quest_id: int
    name: str
    min_level: int
    level: int
    #: "alliance", "horde" or "both" -- wowhead's own `side` field, see
    #: `QUEST_SIDE_FACTION`.
    faction: str


class ItemPageSources(BaseModel):
    """Everything this module could resolve off one item's wowhead page.
    All four lists empty is a normal, cacheable outcome -- wowhead HAD a
    page for the item but named no dropped-by/sold-by/created-by/
    reward-from-q row this module reads (its `contained-in-*` might be
    the only thing on the page, see the module doc); the caller commits
    it anyway so a later run does not refetch it for nothing.
    """

    item_id: int
    dropped_by: list[NpcSource] = []
    sold_by: list[NpcSource] = []
    crafted_by: list[CraftedSource] = []
    quest_rewards: list[QuestRewardSource] = []
    #: vendor-11036 lane, 2026-09-30: this item's own `jsonequip.
    #: reqfaction` -- a real Blizzard `Faction.dbc` id, verified against
    #: a primary source to agree with the fork database's OWN numbering
    #: one for one (item 21200 "Signet Ring of the Bronze Dragonflight":
    #: wowhead states `reqfaction: 910`, the fork's own rep source for
    #: the identical id is `rep:brood-of-nozdormu:*`, faction id 910) --
    #: unlike `QuestRewardSource`'s `side`/`reprewards` fields, which
    #: `pipeline.wowhead_item_sources`' own module doc already warned
    #: disagree with the fork's numbering. `None` when the item states
    #: no reputation requirement at all (wowhead omits `jsonequip.
    #: reqfaction` entirely rather than a sentinel).
    required_faction_id: int | None = None
    #: `jsonequip.reqrep` -- wowhead's OWN standing enum, 0 (Hated)
    #: through 7 (Exalted), one less than `pipeline.forkdb.REP_LEVELS`'
    #: 1-8 at every level checked (item 21200: wowhead `reqrep: 7`, the
    #: fork's own source is `rep:brood-of-nozdormu:exalted`, REP_LEVELS'
    #: `8`; item 21197: wowhead `reqrep: 4`, fork's own source
    #: `rep:brood-of-nozdormu:friendly`, REP_LEVELS' `5`). Kept RAW
    #: (wowhead's own number, not yet offset into REP_LEVELS) the same
    #: way this module leaves `QuestRewardSource.faction` as wowhead's
    #: own `side` encoding rather than decoding early -- decoding one
    #: extra place would be one more place to keep in sync with
    #: `pipeline.forkdb.REP_LEVELS` if it ever changes.
    required_standing_raw: int | None = None

    def is_empty(self) -> bool:
        return not (self.dropped_by or self.sold_by or self.crafted_by or self.quest_rewards)


class ItemPageNotFound(Exception):
    """Wowhead served no item page for this id -- a 404 (redirected to
    `.../items?notFound=<id>`, itself served with a 404 status) or a
    page with no `new Listview(` at all, which a real item page always
    has (even an item with zero sources still carries `same-model-as`/
    `outfits`/`comments` listviews)."""


def raw_path(build_dir: Path, item_id: int) -> Path:
    return build_dir / "raw" / CACHE_SUBDIR / f"{item_id}.html"


def _listview_data(html: str, listview_id: str) -> list | None:
    """The `data: [...]` array for `new Listview({... id: '<listview_id>',
    ...})`, or None when that listview is not on the page at all. See
    the module doc for why this is a marker-then-`raw_decode` walk
    rather than a single regex over the whole object.
    """
    marker = f"id: '{listview_id}',"
    at = html.find(marker)
    if at < 0:
        return None
    data_at = html.find("data: [", at)
    if data_at < 0:
        return None
    start = data_at + len("data: ")
    data, _ = json.JSONDecoder().raw_decode(html[start:])
    return data


def _row_chance(row: dict) -> float | None:
    """`100 * count/outof` -- wowhead's own "N of M recorded kills"
    tally on a `dropped-by` row, or `None` for a `sold-by` row (never
    states `outof`) or a `dropped-by` row wowhead itself states no
    tally for."""
    count, outof = row.get("count"), row.get("outof")
    if not outof:
        return None
    return round(100 * count / outof, 4)


def _row_cost(row: dict) -> tuple[int, list[int]]:
    """`(money, item_ids)` from a `sold-by` row's own `cost` field.
    `(0, [])` for a `dropped-by` row (never carries `cost`) or a
    `sold-by` row wowhead states no cost at all for (an empty `cost`
    list). Otherwise `cost[0]` is either `[money_copper]` -- a plain
    gold price, no `ItemExtendedCost` override at all (item 16769's own
    `sold-by` row, "Gorn One Eye": `cost: [[133081]]`) -- or the fuller
    `[money_copper, currency_pairs, item_pairs]` an `ItemExtendedCost`
    override carries (vendor 11036's own rows, `cost: [[0, [],
    [[239759, 1]]]]`); either shape's first element is always the gold
    price."""
    cost = row.get("cost") or []
    if not cost:
        return 0, []
    entry = cost[0]
    money = int(entry[0] or 0) if entry else 0
    items = entry[2] if len(entry) > 2 else []
    return money, [int(pair[0]) for pair in items]


def _npc_rows(html: str, listview_id: str) -> list[NpcSource]:
    rows = _listview_data(html, listview_id) or []
    out = []
    for row in rows:
        money, item_ids = _row_cost(row)
        out.append(
            NpcSource(
                npc_id=int(row["id"]),
                name=str(row.get("name") or row.get("displayName") or ""),
                zone_ids=[int(z) for z in (row.get("location") or [])],
                min_level=row.get("minlevel"),
                max_level=row.get("maxlevel"),
                chance=_row_chance(row),
                cost_money=money,
                cost_item_ids=item_ids,
            )
        )
    return out


def _item_equip_requirement(html: str, item_id: int) -> tuple[int | None, int | None]:
    """`(reqfaction, reqrep)` off THIS item's own `jsonequip` record in
    the page's `WH.Gatherer.addData(3, 16, {...})` blob -- the same
    per-item tooltip data every wowhead item page embeds regardless of
    class/subtype (verified against item 21200, "Signet Ring of the
    Bronze Dragonflight": see `ItemPageSources.required_faction_id`'s
    own doc for the primary-source cross-check). `(None, None)` when the
    blob is missing (a malformed/notFound page `parse_item_page`'s own
    `ItemPageNotFound` guard already rejects before this runs) or the
    item's own record states no `reqfaction` -- the overwhelming
    majority of items, which need no reputation at all.
    """
    marker = "WH.Gatherer.addData(3, 16, "
    at = html.find(marker)
    if at < 0:
        return None, None
    start = at + len(marker)
    try:
        blob, _ = json.JSONDecoder().raw_decode(html[start:])
    except ValueError:
        return None, None
    record = blob.get(str(item_id))
    if not isinstance(record, dict):
        return None, None
    equip = record.get("jsonequip") or {}
    faction = equip.get("reqfaction")
    standing = equip.get("reqrep")
    if faction is None or standing is None:
        return None, None
    return int(faction), int(standing)


def _crafted_rows(html: str) -> list[CraftedSource]:
    rows = _listview_data(html, "created-by-spell") or []
    professions: list[str] = []
    for row in rows:
        for skill_line in row.get("skill") or []:
            profession = SKILL_LINE_PROFESSIONS.get(int(skill_line))
            if profession and profession not in professions:
                professions.append(profession)
    return [CraftedSource(profession=profession) for profession in professions]


def _quest_reward_rows(html: str) -> list[QuestRewardSource]:
    rows = _listview_data(html, "reward-from-q") or []
    return [
        QuestRewardSource(
            quest_id=int(row["id"]),
            name=str(row.get("name") or ""),
            min_level=int(row.get("reqlevel") or 0),
            level=int(row.get("level") or 0),
            faction=QUEST_SIDE_FACTION.get(int(row.get("side") or 3), "both"),
        )
        for row in rows
    ]


def parse_item_page(item_id: int, html: str) -> ItemPageSources:
    if "new Listview(" not in html:
        raise ItemPageNotFound(f"item {item_id}: no Listview blocks in page")
    required_faction_id, required_standing_raw = _item_equip_requirement(html, item_id)
    return ItemPageSources(
        item_id=item_id,
        dropped_by=_npc_rows(html, "dropped-by"),
        sold_by=_npc_rows(html, "sold-by"),
        crafted_by=_crafted_rows(html),
        quest_rewards=_quest_reward_rows(html),
        required_faction_id=required_faction_id,
        required_standing_raw=required_standing_raw,
    )


def fetch_item_page(item_id: int, client: httpx.Client) -> str:
    response = client.get(ITEM_URL.format(item_id=item_id), timeout=30, follow_redirects=True)
    response.raise_for_status()
    return response.text


@dataclass(frozen=True)
class ItemSourceResult:
    """`fetch_item_sources`/`load_cached_item_sources`'s return, the same
    shape as `pipeline.wowhead_quests.QuestLevelResult`."""

    sources: dict[int, ItemPageSources]
    missing: set[int]


def _read_cached(build_dir: Path, item_id: int) -> ItemPageSources | None:
    path = raw_path(build_dir, item_id)
    if not path.exists():
        return None
    try:
        return parse_item_page(item_id, path.read_text(encoding="utf-8"))
    except ItemPageNotFound as exc:
        logger.warning("wowhead item cache %s: %s", path, exc)
        return None


def load_cached_item_sources(
    build: str, item_ids: Iterable[int], root: Path = Path("builds")
) -> ItemSourceResult:
    """Every item id this build's cache already has a page for, with NO
    network calls -- the offline half of `fetch_item_sources`, split out
    the same way `pipeline.wowhead_quests.load_cached_quest_levels` is,
    for tests that should never touch the network."""
    build_dir = root / build
    sources: dict[int, ItemPageSources] = {}
    missing: set[int] = set()
    for item_id in sorted(set(item_ids)):
        found = _read_cached(build_dir, item_id)
        if found is None:
            missing.add(item_id)
        else:
            sources[item_id] = found
    return ItemSourceResult(sources=sources, missing=missing)


def _retry_after_seconds(exc: httpx.HTTPStatusError) -> float | None:
    header = exc.response.headers.get("Retry-After")
    if header is None:
        return None
    try:
        return float(header)
    except ValueError:
        return None


def fetch_item_sources(
    build: str,
    item_ids: Iterable[int],
    root: Path = Path("builds"),
    client: httpx.Client | None = None,
    delay: float = FETCH_DELAY_SECONDS,
    max_pages: int | None = None,
) -> ItemSourceResult:
    """Fetch (live) every item id this build's cache does not already
    have, caching each page under `builds/<build>/raw/items/<id>.html`
    so a rerun resumes exactly where the last one stopped -- politeness
    rules identical to `pipeline.wowhead_quests.fetch_quest_levels` (read
    that function's own doc for the full throttle/backoff/max_pages
    contract, which this mirrors line for line, only over items instead
    of quests).

    `item_ids`' own order is preserved into `pending` (deduplicated)
    rather than re-sorted ascending by id -- src-crawl-order lane,
    2026-09-29: `pipeline.item_sources.fetch_missing_from_wowhead` hands
    this a crawl-priority-ordered backlog, and a plain `sorted()` here
    would silently put it back into ascending-item-id order right before
    the live requests actually go out, which is the one ordering this
    whole lane exists to stop happening. A cached id is still skipped in
    place (no live request, `index` advances) without disturbing the
    order the rest of `pending` is walked in, so resuming a
    budget-capped run picks up exactly where the last one left off.
    """
    build_dir = root / build
    own = client is None
    client = client or httpx.Client(headers={"User-Agent": USER_AGENT})
    sources: dict[int, ItemPageSources] = {}
    missing: set[int] = set()
    pending: list[int] = []
    _pending_seen: set[int] = set()
    for item_id in item_ids:
        if item_id in _pending_seen:
            continue
        _pending_seen.add(item_id)
        pending.append(item_id)
    consecutive_throttled = 0
    backoff = THROTTLE_BACKOFF_START_SECONDS
    stopped_early_at: int | None = None
    live_requests = 0
    try:
        index = 0
        while index < len(pending):
            item_id = pending[index]
            cached = _read_cached(build_dir, item_id)
            if cached is not None:
                sources[item_id] = cached
                index += 1
                continue
            if max_pages is not None and live_requests >= max_pages:
                logger.info(
                    "wowhead items: stopped after the %d-page budget; %d of %d ids "
                    "remain unfetched",
                    max_pages,
                    len(pending) - index,
                    len(pending),
                )
                stopped_early_at = index
                break
            live_requests += 1
            try:
                html = fetch_item_page(item_id, client)
            except httpx.HTTPStatusError as exc:
                status = exc.response.status_code
                if status in THROTTLE_STATUS_CODES:
                    consecutive_throttled += 1
                    logger.warning(
                        "wowhead item %d: throttled (%d), %d/%d consecutive",
                        item_id,
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
                    continue  # retry the SAME item_id, index unchanged
                logger.warning("wowhead item %d: %s", item_id, exc)
                missing.add(item_id)
                consecutive_throttled = 0
                backoff = THROTTLE_BACKOFF_START_SECONDS
                time.sleep(delay)
                index += 1
                continue
            try:
                parsed = parse_item_page(item_id, html)
            except ItemPageNotFound as exc:
                logger.warning("wowhead item %d: %s", item_id, exc)
                missing.add(item_id)
                consecutive_throttled = 0
                backoff = THROTTLE_BACKOFF_START_SECONDS
                time.sleep(delay)
                index += 1
                continue
            path = raw_path(build_dir, item_id)
            path.parent.mkdir(parents=True, exist_ok=True)
            path.write_text(html, encoding="utf-8")
            sources[item_id] = parsed
            consecutive_throttled = 0
            backoff = THROTTLE_BACKOFF_START_SECONDS
            time.sleep(delay)
            index += 1
    finally:
        if own:
            client.close()
    if stopped_early_at is not None:
        remaining = [i for i in pending[stopped_early_at:] if i not in sources]
        missing.update(remaining)
        if consecutive_throttled >= THROTTLE_STOP_AFTER:
            logger.warning(
                "wowhead items: stopped after %d consecutive throttled responses; %d of %d "
                "ids remain unfetched (stay unsourced; rerun later to resume from the warm "
                "cache)",
                THROTTLE_STOP_AFTER,
                len(remaining),
                len(pending),
            )
    elif missing:
        logger.info(
            "wowhead items: %d of %d ids had no usable page (stay unsourced)",
            len(missing),
            len(sources) + len(missing),
        )
    return ItemSourceResult(sources=sources, missing=missing)
