"""variants.py: writes the Discipline rotation variants the lane measures (rotation JSON files for
spec-breakdown -rotation) from data/curated/apl/priest-discipline.json. Run from the repo root."""
import copy, json, os

SRC = 'data/curated/apl/priest-discipline.json'
OUT = os.path.dirname(os.path.abspath(__file__))
PWS = {'spellId': 10901, 'rank': 10}
FLASH = 10917


def player(i):
    return {'type': 'Player', 'index': i}


def const(v):
    return {'const': {'val': str(v)}}


def pace(share):
    return {'cmp': {'op': 'OpGe', 'lhs': {'currentManaPercent': {}},
                    'rhs': {'math': {'op': 'OpMul', 'lhs': {'remainingTimePercent': {}}, 'rhs': const(share)}}}}


def no_weakened_soul(i):
    return {'not': {'val': {'auraIsActive': {'sourceUnit': player(i), 'auraId': {'spellId': 6788, 'rank': 0}}}}}


def hp_below(i, v):
    return {'cmp': {'op': 'OpLt', 'lhs': {'currentHealthPercent': {'sourceUnit': player(i)}}, 'rhs': const(v)}}


def shield_line(i, hp, share):
    conds = [no_weakened_soul(i)]
    if hp is not None:
        conds.append(hp_below(i, hp))
    if share is not None:
        conds.append(pace(share))
    return {'action': {'castSpell': {'spellId': dict(PWS), 'target': player(i)}, 'condition': {'and': {'vals': conds}}}}


def spell_of(line):
    return line['action'].get('castSpell', {}).get('spellId', {}).get('spellId')


def target_of(line):
    return line['action'].get('castSpell', {}).get('target', {}).get('index')


def index_after_tank_shield(pl):
    return next(i for i, l in enumerate(pl) if spell_of(l) == 10901) + 1


def write(name, doc):
    with open(os.path.join(OUT, name + '.rotation.json'), 'w') as f:
        json.dump(doc, f, indent=1)


def main():
    base = json.load(open(SRC))
    variants = {}

    def mk(name, fn):
        doc = copy.deepcopy(base)
        doc['rotation']['priorityList'] = fn(doc['rotation']['priorityList'])
        variants[name] = doc

    # a: shield the party members as well, when Weakened Soul is off, on the shield's own pace
    def party_shields(hp, share):
        def fn(pl):
            at = index_after_tank_shield(pl)
            return pl[:at] + [shield_line(i, hp, share) for i in (1, 2, 3, 4)] + pl[at:]
        return fn
    mk('a-party-shield-90', party_shields(0.9, 0.9))
    mk('a-party-shield-any', party_shields(None, 0.9))
    mk('a-party-shield-nopace', party_shields(None, None))
    # b: no Flash Heal at all
    mk('b-no-flash', lambda pl: [l for l in pl if spell_of(l) != FLASH])
    # c: no Flash Heal on the party, the tank's kept
    mk('c-tank-flash-only', lambda pl: [l for l in pl if not (spell_of(l) == FLASH and target_of(l) != 5)])
    # d: party shields (any, paced) and no party Flash Heal
    mk('d-shield-any-no-party-flash', lambda pl: party_shields(None, 0.9)([l for l in pl if not (spell_of(l) == FLASH and target_of(l) != 5)]))
    # e: the tank shield ungated on health (only Weakened Soul and pace)
    def tank_any(pl):
        pl = copy.deepcopy(pl)
        i = index_after_tank_shield(pl) - 1
        pl[i] = shield_line(5, None, 0.9)
        return pl
    mk('e-tank-shield-any', tank_any)
    for name, doc in variants.items():
        write(name, doc)
        print(name)


if __name__ == '__main__':
    main()
