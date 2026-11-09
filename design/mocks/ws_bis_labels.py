"""Before and after of the BiS source lines the 2026-11-09 feedback branch changes (coordinator, 2026-11-09).

The level-60 Marksmanship Hunter page from the parity facts (data/bis-facts.json, state dps-60), Alliance and Horde,
each before and after. Only the source lines (line2) of these rows change:
  Alliance: Theramore Pauldrons and Gloves (rep, now with where to buy), Ranger's Belt (Pix Xizzix, now with place and
  its 100 Tarnished Undermine Reals); Cape of the Black Baron is the unchanged dungeon drop for contrast.
  Horde: Darkspear Pauldrons and Raider Gloves (Creeg Bothunk, now the same reputation line as the Alliance set),
  Ranger's Belt as above.
The after lines are the words web/src/lib/bis/source-line.ts prints for the new loot.json fields (its tests pin them).
    python3 design/mocks/ws_bis_labels.py   -> design/mocks/web/web-bis-labels-<state>.html
"""
import copy

from ws_bis import FACTS, render
from ws_lib import write

PIX_BEFORE = 'Vendor: Pix Xizzix'
PIX_AFTER = 'Vendor: Pix Xizzix, Booty Bay · 100 Tarnished Undermine Reals'
THERAMORE_BEFORE = 'Theramore Expeditionary Force reputation · Exalted'
THERAMORE_AFTER = 'Theramore Expeditionary Force reputation · Exalted · The Barrens'
CREEG_BEFORE = 'Vendor: Creeg Bothunk'
DARKSPEAR_AFTER = 'Darkspear Raiders reputation · Exalted · The Barrens'

# The Horde picks (data/builds/1.60.1.70790/bis/hunter-marksmanship.json, band 60 raid, horde) and their client rows.
HORDE_ROWS = {
    'shoulder': {'item': 272105, 'name': 'Darkspear Pauldrons', 'stats': '+27 Agi · +14 Crit', 'icon': 'inv_shoulder_06'},
    'hands': {'item': 272099, 'name': 'Raider Gloves', 'stats': '+27 Agi · +11 Str', 'icon': 'inv_gauntlets_27'},
}
# The rest of the Horde list's own rows, unchanged by the branch, so the Horde page shows no Alliance gear.
HORDE_UNCHANGED = {
    'neck': {'item': 19534, 'name': "Scout's Medallion", 'icon': 'inv_jewelry_necklace_14',
             'line2': 'Warsong Outriders reputation · Honored'},
    'wrist': {'line2': 'Warsong Outriders reputation · Exalted'},
    'legs': {'item': 237824, 'name': "Outrider's Chain Leggings", 'line2': 'Warsong Outriders reputation · Exalted'},
    'finger2': {'line2': 'Frostwolf Clan reputation · Exalted'},
}


def _rows(facts: dict):
    return [row for column in ('left', 'right') for row in facts['doll'][column]]


def _faction(facts: dict, side: str) -> None:
    for group in facts['controls']['groups']:
        if group['label'] == 'Faction':
            for option in group['options']:
                option['on'] = option['label'] == side


def state(side: str, after: bool) -> dict:
    facts = copy.deepcopy(FACTS['dps-60'])
    facts['id'] = f'labels-{side.lower()}-{"after" if after else "before"}'
    _faction(facts, side)
    if side == 'Horde':
        facts['answer']['trust'] = facts['answer']['trust'].replace('Night-elf', 'Troll')
    for row in _rows(facts):
        if side == 'Horde' and row['slot'] in HORDE_UNCHANGED:
            row.update(HORDE_UNCHANGED[row['slot']])
        if row['slot'] == 'waist':
            row['line2'] = PIX_AFTER if after else PIX_BEFORE
        elif row['slot'] in ('shoulder', 'hands'):
            if side == 'Horde':
                row.update(HORDE_ROWS[row['slot']])
                row['line2'] = DARKSPEAR_AFTER if after else CREEG_BEFORE
            else:
                row['line2'] = THERAMORE_AFTER if after else THERAMORE_BEFORE
    return facts


if __name__ == '__main__':
    for side in ('Alliance', 'Horde'):
        for after in (False, True):
            facts = state(side, after)
            print(write(f'web-bis-{facts["id"]}.html', render(facts)))
