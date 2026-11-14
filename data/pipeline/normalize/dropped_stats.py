"""Stats a hotfix took off an item row, which no Wowhead fill may put back.

The equip fill (`pipeline.normalize.wowhead.fill_client_equip_stats` for the
planner, `pipeline.simdb.equip_fill` for simdb) gives a client row every stat
Wowhead Forever states and the row leaves out, because a hotfix item's
on-equip link never reaches the client. Wowhead's feed lags the server's
pushes, so a stat a push removes comes back from it: Blizzard's 13 November
item hotfixes removed Blessed Band of Light's hit (272407), turned Channeler's
Ring's Spirit into spell power (272406) and took the Bad Hat's shadow spell
power away (260212), and all three kept the old stat. A stat type the shipped
ItemSparse row states and the merged (hotfixed and pushed) row no longer does
was removed on purpose, so the fill skips it for that item.
"""

from __future__ import annotations

from collections.abc import Iterable

from pipeline.normalize.gear import (
    MODIFIER_ALSO_GRANTS,
    STAT_BY_MODIFIER_ID,
    STAT_COLUMNS,
    WEAPON_DAMAGE_STAT_KEY,
)


def _stat_ids(row: dict[str, str]) -> frozenset[int]:
    ids = (row.get(f"StatModifier_bonusStat_{index}", "-1") for index in STAT_COLUMNS)
    return frozenset(int(value) for value in ids if value and int(value) >= 0)


def _stat_keys(stat_ids: Iterable[int]) -> frozenset[str]:
    keys: set[str] = set()
    for stat_id in stat_ids:
        key = STAT_BY_MODIFIER_ID.get(stat_id)
        if key is not None and key != WEAPON_DAMAGE_STAT_KEY:
            keys.add(key)
        keys.update(MODIFIER_ALSO_GRANTS.get(stat_id, ()))
    return frozenset(keys)


def dropped_stat_keys(
    shipped: Iterable[dict[str, str]], merged: Iterable[dict[str, str]]
) -> dict[int, frozenset[str]]:
    """Item id -> the planner stat keys its shipped ItemSparse row states and its
    merged row does not, for every item the merge kept whose row lost one."""
    before = {int(row["ID"]): _stat_ids(row) for row in shipped}
    dropped: dict[int, frozenset[str]] = {}
    for row in merged:
        item_id = int(row["ID"])
        if item_id not in before:
            continue
        after = _stat_ids(row)
        keys = _stat_keys(before[item_id] - after) - _stat_keys(after)
        if keys:
            dropped[item_id] = keys
    return dropped
