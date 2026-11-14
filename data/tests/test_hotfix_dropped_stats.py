"""Stats a hotfix took off an item row stay off (sim accuracy 2026-11-13).

Blizzard's 13 November item hotfixes (push 114062 and 114063) removed Blessed
Band of Light's hit (272407), turned Channeler's Ring's Spirit into spell power
(272406) and took the Bad Hat's shadow spell power away (260212). The pushed
ItemSparse rows say so; Wowhead Forever's feed still shows the old items, and
the equip fill took every stat the client row leaves out, so the three kept
1% hit, 8 Spirit and 23 shadow spell power in the planner and in simdb.
"""

from __future__ import annotations

from pipeline.normalize.dropped_stats import dropped_stat_keys


def _row(item_id: int, stat_ids: list[int]) -> dict[str, str]:
    row = {"ID": str(item_id)}
    for index in range(10):
        stat = stat_ids[index] if index < len(stat_ids) else -1
        row[f"StatModifier_bonusStat_{index}"] = str(stat)
        row[f"StatPercentEditor_{index}"] = "1000" if stat >= 0 else "0"
    return row


SHIPPED = [
    _row(272406, [5, 31, 6]),
    _row(272407, [41, 42, 7, 6, 31]),
    _row(260212, [41, 42, 5, 6, 88]),
    _row(812, [5, 6, 7]),
]
MERGED = [
    _row(272406, [5, 31, 45]),
    _row(272407, [41, 42, 7, 6]),
    _row(260212, [41, 42, 5, 6]),
    _row(812, [5, 6, 7]),
    _row(999999, [5]),
]


def test_a_stat_a_hotfix_removed_is_named_per_item():
    assert dropped_stat_keys(SHIPPED, MERGED) == {
        272406: frozenset({"spirit"}),
        272407: frozenset({"hit"}),
        260212: frozenset({"shadow_power"}),
    }


def test_an_unchanged_row_and_a_hotfix_only_row_drop_nothing():
    dropped = dropped_stat_keys(SHIPPED, MERGED)
    assert 812 not in dropped
    assert 999999 not in dropped


def test_a_row_the_hotfix_removed_entirely_drops_nothing():
    assert dropped_stat_keys([_row(1, [5])], []) == {}
