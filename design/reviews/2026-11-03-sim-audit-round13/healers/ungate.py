"""Write variants of a curated healer APL with the mana-pace term removed from every line of one spell id.
usage: ungate.py <apl.json> <outdir> ; prints spell ids with gated lines."""
import json, sys, copy, os

def has_pace(o):
    s = json.dumps(o)
    return 'currentManaPercent' in s and 'remainingTimePercent' in s

def strip(cond):
    """Remove mana-pace cmp terms that sit under an 'and' (or are the whole condition)."""
    if cond is None:
        return None
    if 'cmp' in cond and has_pace(cond):
        return None
    if 'and' in cond:
        vals = [strip(v) for v in cond['and']['vals']]
        vals = [v for v in vals if v is not None]
        if not vals:
            return None
        if len(vals) == 1:
            return vals[0]
        return {'and': {'vals': vals}}
    return cond

src, out = sys.argv[1], sys.argv[2]
d = json.load(open(src))
pl = d['rotation']['priorityList']
groups = {}
for i, a in enumerate(pl):
    act = a.get('action', {})
    cs = act.get('castSpell')
    if cs and has_pace(act.get('condition', {})):
        groups.setdefault(cs['spellId']['spellId'], []).append(i)
for sid, idxs in groups.items():
    t = copy.deepcopy(d)
    for i in idxs:
        act = t['rotation']['priorityList'][i]['action']
        c = strip(act.get('condition'))
        if c is None:
            act.pop('condition', None)
        else:
            act['condition'] = c
    json.dump(t, open(os.path.join(out, f'ungate-{sid}.json'), 'w'))
    print(sid, idxs)
