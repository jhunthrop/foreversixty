"""The plates' hit and crit read items.json's contract: a literal percent stays a percent.

    python3 -m pytest design/mocks/test_stat_percent.py
"""
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))

import addon_hub_l3 as gear  # noqa: E402

TRUEAIM = {'id': 13255, 'stats': {'hit': 3}, 'percent_stats': {'hit': 3}, 'effect_text': ''}
HOLY_MIGHT = {'id': 867, 'stats': {'crit': 14, 'attack_power': 10}, 'percent_stats': {}, 'effect_text': ''}


def test_a_literal_percent_is_not_divided_by_the_rating_factor():
    assert gear.stat_percent(TRUEAIM, 'hit') == 3


def test_rating_points_are_read_as_percent():
    assert gear.stat_percent(HOLY_MIGHT, 'crit') == 1


def test_a_stat_mixing_both_divides_only_the_rating_part():
    item = {'stats': {'hit': 21}, 'percent_stats': {'hit': 1}}
    assert gear.stat_percent(item, 'hit') == 3


def test_a_flat_stat_is_its_amount():
    assert gear.stat_percent(HOLY_MIGHT, 'attack_power') == 10


def test_the_plate_prints_trueaims_three_percent():
    import addon_hub_l10 as marks
    lines = ''.join(marks.equip_lines(TRUEAIM))
    assert 'Improves your chance to hit by 3.0%.' in lines
    assert '0.3%' not in lines
