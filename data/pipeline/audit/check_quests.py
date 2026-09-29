"""Check B: every `loot.json` quest link against cmangos/classic-db (the
1.12 dump `pipeline.classic_sources`/`pipeline.quest_levels` both pin) --
does the quest exist there at all, does its faction/min_level/level agree,
and is the item actually one of `RewardChoiceItemId1-6`/`RewardItemId1-4`
(a Forever-added reward on an existing 1.12 quest will not be, and that is
tracked separately from "quest not in 1.12 at all").
"""

from __future__ import annotations

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult
from pipeline.classic_sources import ClassicDbQuestInfo

CODE = "B"
LABEL = "Quests vs classic-db"
PRIMARY_SOURCE = "raw/classicdb/sources.json (cmangos/classic-db quest_template dump)"


def _classic_db_quest_info(
    classic_sources_by_item: dict[int, list],
) -> dict[int, ClassicDbQuestInfo]:
    """Every quest classic-db names a reward for, anywhere -- keyed by
    quest id regardless of which item this build's own loot.json ties it
    to (a quest with more than one reward choice still has one
    min_level/level/faction)."""
    info: dict[int, ClassicDbQuestInfo] = {}
    for records in classic_sources_by_item.values():
        for record in records:
            if record.kind == "quest_reward" and record.quest is not None:
                info.setdefault(record.quest.quest_id, record.quest)
    return info


def _item_quest_ids(records: list) -> set[int]:
    return {
        record.quest.quest_id
        for record in records
        if record.kind == "quest_reward" and record.quest is not None
    }


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    classic_sources = ctx.classic_sources_by_item
    if not classic_sources:
        result.skipped = (
            "raw/classicdb/sources.json has no entries (run `python -m pipeline "
            "fetch-classic-sources` first) -- every quest link below is unverified"
        )
        return result
    db_quests = _classic_db_quest_info(classic_sources)

    forever_new = 0
    reward_not_in_112 = 0
    for item_id_str, quest_sources in ctx.loot.get("quests", {}).items():
        item_id = int(item_id_str)
        item_records = classic_sources.get(item_id, [])
        item_quest_ids = _item_quest_ids(item_records)
        for qs in quest_sources:
            quest_id = qs["quest_id"]
            db_quest = db_quests.get(quest_id)
            if db_quest is None:
                forever_new += 1
                result.add(
                    "minor",
                    quest_id,
                    f"quest {qs['name']!r} (id {quest_id}, item {item_id}) is not in "
                    f"classic-db's quest_template (1.12) -- Forever-new; what loot.json "
                    f"itself knows: min_level={qs['min_level']} level={qs['level']} "
                    f"faction={qs['faction']} (source: {qs['level_source']})",
                )
                continue
            result.checked += 1
            if qs["level_source"] != "classic-db":
                result.add(
                    "minor",
                    quest_id,
                    f"quest {qs['name']!r} (id {quest_id}, item {item_id}) is tagged "
                    f"level_source={qs['level_source']!r} but classic-db actually covers "
                    "this quest id; quest-levels.json looks stale for it",
                    theirs=f"min_level={db_quest.min_level} level={db_quest.level}",
                )
            if qs["faction"] != db_quest.faction:
                result.add(
                    "blocker",
                    quest_id,
                    f"quest {qs['name']!r} (id {quest_id}, item {item_id}) faction disagrees "
                    "with classic-db's RequiredRaces",
                    ours=qs["faction"],
                    theirs=db_quest.faction,
                )
            if qs["min_level"] != db_quest.min_level:
                result.add(
                    "major",
                    quest_id,
                    f"quest {qs['name']!r} (id {quest_id}, item {item_id}) min_level disagrees "
                    "with classic-db's MinLevel",
                    ours=str(qs["min_level"]),
                    theirs=str(db_quest.min_level),
                )
            if qs["level"] != db_quest.level:
                result.add(
                    "major",
                    quest_id,
                    f"quest {qs['name']!r} (id {quest_id}, item {item_id}) level disagrees "
                    "with classic-db's QuestLevel",
                    ours=str(qs["level"]),
                    theirs=str(db_quest.level),
                )
            if quest_id not in item_quest_ids:
                reward_not_in_112 += 1
                wowhead = ctx.item_sources.get(item_id) or {}
                named_by_wowhead = any(
                    reward.get("quest_id") == quest_id
                    for reward in wowhead.get("quest_rewards", [])
                )
                result.add(
                    "minor",
                    item_id,
                    f"item {item_id} is not among quest {quest_id} ({qs['name']!r})'s "
                    "classic-db RewardChoiceItemId1-6/RewardItemId1-4 -- reward not in 1.12"
                    + (
                        "; wowhead's Forever quest cache does name this quest for the item"
                        if named_by_wowhead
                        else "; wowhead's Forever quest cache does not name it either "
                        "(raw/items/item-sources.json)"
                    ),
                )
    if forever_new:
        result.add(
            "minor",
            "*",
            f"{forever_new} quest link(s) reference a quest id classic-db's quest_template "
            "does not have at all (Forever-new quests)",
        )
    if reward_not_in_112:
        result.add(
            "minor",
            "*",
            f"{reward_not_in_112} quest link(s) are a Forever-added reward on a quest "
            "classic-db does have (the item is not in that quest's own 1.12 reward list)",
        )
    return result
