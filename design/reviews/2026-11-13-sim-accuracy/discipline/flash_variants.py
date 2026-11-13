"""flash_variants.py: the tank's emergency Flash Heal line, gated three ways, for both priest cards.
Writes <spec>-<variant>.rotation.json beside this file. Run from the repo root."""
import copy, json, os

OUT = os.path.dirname(os.path.abspath(__file__))
FLASH = 10917
TANK = 5


def const(v):
    return {'const': {'val': str(v)}}


def pace(share):
    return {'cmp': {'op': 'OpGe', 'lhs': {'currentManaPercent': {}},
                    'rhs': {'math': {'op': 'OpMul', 'lhs': {'remainingTimePercent': {}}, 'rhs': const(share)}}}}


def tank_hp_below(v):
    return {'cmp': {'op': 'OpLt', 'lhs': {'currentHealthPercent': {'sourceUnit': {'type': 'Player', 'index': TANK}}}, 'rhs': const(v)}}


def is_tank_flash(line):
    cast = line['action'].get('castSpell', {})
    return cast.get('spellId', {}).get('spellId') == FLASH and cast.get('target', {}).get('index') == TANK


def gated(pl, condition):
    out = copy.deepcopy(pl)
    for line in out:
        if is_tank_flash(line):
            line['action']['condition'] = condition
    return out


VARIANTS = {
    'tankflash-paced70': {'and': {'vals': [tank_hp_below(0.35), pace(0.7)]}},
    'tankflash-paced90': {'and': {'vals': [tank_hp_below(0.35), pace(0.9)]}},
    'tankflash-under20': tank_hp_below(0.2),
}

for spec in ('priest-discipline', 'priest-holy'):
    base = json.load(open(f'data/curated/apl/{spec}.json'))
    assert sum(is_tank_flash(l) for l in base['rotation']['priorityList']) == 1, spec
    for name, cond in VARIANTS.items():
        doc = copy.deepcopy(base)
        doc['rotation']['priorityList'] = gated(doc['rotation']['priorityList'], cond)
        with open(os.path.join(OUT, f'{spec}-{name}.rotation.json'), 'w') as f:
            json.dump(doc, f, indent=1)
        print(spec, name)
