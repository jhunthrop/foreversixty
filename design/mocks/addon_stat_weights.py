"""The Stat weights block as addon/ForeverSixty/surfaces/StatWeights.lua draws it, for every hub mock that shows it
(the Gear tab's right column, addon_hub_l3.py; the Rotation tab's left, addon_hub_rotation.py): one model and one
drawing, so the two tabs never word or lay out the weights twice.

The rules are StatWeights.model's: each stat per rating point (Rules.RATING_PER_PERCENT) over the top stat's, so the
top stat reads 1.00 (the SimC convention); attack power and ranged attack power in one row where the spec weighs them
alike (the reference's); a stat under 5% of the top left out; at most six rows; rating haste apart per 1%, counted at
melee and spell haste both (Gear.score). Every string comes from Locale.lua.
"""
import re

from addon_builds import L, LOCALE, esc
from addon_ux import ADDON

RATING_PER_PERCENT = {'hit': 10, 'crit': 14, 'expertise': 10, 'dodge': 12, 'parry': 15, 'block': 5, 'defense': 1,
                      'haste': 10}
HASTE = ('melee_haste', 'spell_haste')
NO_ROW = {'haste', 'weapon_dps', *HASTE}
SAME_AS = {'attack_power': 'ranged_attack_power', 'ranged_attack_power': 'attack_power'}
MAX_ROWS, FLOOR = 6, 0.05
# StatWeights.LAYOUT: the Weight and DPS per point columns, right aligned.
WEIGHT_COLUMN, PER_POINT_COLUMN = 56, 96
STAT_NAMES = dict(re.findall(r'(\w+) = "([^"]+)"',
                             re.search(r'L\.weightsStatNames = \{(.*?)\}', LOCALE, re.S).group(1)))


def fmt(pattern, *args):
    """string.format as the addon calls it: %d, %s, %.2f and %% read the same in Python."""
    return pattern % args


def label(stat):
    name = STAT_NAMES.get(stat, stat)
    return fmt(L('weightsRating'), name) if stat in RATING_PER_PERCENT else name


def model(weights, dps, note='', reference='ranged_attack_power'):
    """StatWeights.model over a weights table, its DPS per point and the source note."""
    rows = []
    for stat, weight in weights.items():
        twin = SAME_AS.get(stat)
        if weight <= 0 or stat in NO_ROW or (twin and weights.get(twin) == weight and stat != reference):
            continue
        rows.append((stat, weight / RATING_PER_PERCENT.get(stat, 1)))
    rows.sort(key=lambda r: (-r[1], r[0]))
    top = rows[0][1]
    kept = [(s, p) for s, p in rows if p / top >= FLOOR][:MAX_ROWS]
    haste = sum(weights.get(k, 0) for k in HASTE if weights.get(k, 0) > 0)
    return {'rows': [(label(s), p / top, p * dps if dps else None) for s, p in kept],
            'haste': haste * dps if haste and dps else None, 'note': note or '', 'unit': L('weightsDps')}


def html(m):
    """StatWeights.draw: the section heading with the source, the head row, the rows, the haste line."""
    cell = 'height:18px;padding:0;white-space:nowrap;overflow:hidden;text-overflow:ellipsis'
    num = f'{cell};text-align:right;font-variant-numeric:tabular-nums'
    rows = (f'<tr style="color:{ADDON["muted"]}"><td style="{cell}">{esc(L("weightsStat"))}</td>'
            f'<td style="{num};width:{WEIGHT_COLUMN}px">{esc(L("weightsWeight"))}</td>'
            f'<td style="{num};width:{PER_POINT_COLUMN}px">{esc(fmt(L("weightsPerPoint"), m["unit"]))}</td></tr>')
    for name, scale, value in m['rows']:
        rows += (f'<tr><td style="{cell};color:{ADDON["body"]}">{esc(name)}</td>'
                 f'<td style="{num};color:{ADDON["strong"]}">{scale:.2f}</td>'
                 f'<td style="{num};color:{ADDON["muted"]}">'
                 f'{f"{value:.2f}" if value is not None else esc(L("weightsNone"))}</td></tr>')
    haste = (f'<div style="color:{ADDON["muted"]};margin-top:2px">{esc(fmt(L("weightsHaste"), m["haste"], m["unit"]))}'
             f'</div>' if m['haste'] else '')
    note = f'<span class="r">{esc(m["note"])}</span>' if m['note'] else ''
    head = (f'<div class="sh" style="font-weight:400"><span style="font-family:Title,serif;font-size:14px;'
            f'color:{ADDON["cream"]}">{esc(L("weightsTitle"))}</span>{note}</div>')
    table = f'<table style="width:100%;border-collapse:collapse;table-layout:fixed">{rows}</table>'
    return head, table + haste
