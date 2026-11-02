"""Crit size of each heal in the beta timelines: mean raw crit / mean raw non-crit, per character and spell id.

Raw = the heal event's amount plus the overheal event at the same timestamp, target and spell.
critsize.py [class] [min events]
"""
import collections
import glob
import json
import math
import sys

SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
TL = SP + '/r10/healers/fl/tl'
want = sys.argv[1].lower() if len(sys.argv) > 1 else None
mn = int(sys.argv[2]) if len(sys.argv) > 2 else 15

roster = {}
for f in glob.glob(SP + '/fl/rep/*.roster.json'):
    rid = int(f.split('/')[-1].split('.')[0])
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            tabs = s['talents']['talent_groups'][str(s['talents']['active_group'])]['tabs']
            pts = {v['name']: v['points'] for v in tabs.values()}
            roster[(rid, c['character_id'])] = (s['player']['class'], sum(pts.values()) + 9, pts)


def raw_heals(events):
    over = collections.defaultdict(float)
    for e in events:
        if e['event_type'] == 'overheal':
            over[(e['timestamp_ms'], e['target_id'], e['spell_id'])] += float(e['amount'])
    out = []
    for e in events:
        if e['event_type'] == 'heal':
            k = (e['timestamp_ms'], e['target_id'], e['spell_id'])
            out.append((e, float(e['amount']) + over.pop(k, 0.0)))
    return out


agg = collections.defaultdict(lambda: [[], []])
names = {}
for f in glob.glob(TL + '/*.json'):
    rid, enc, cid = map(int, f.split('/')[-1].split('.')[:3])
    info = roster.get((rid, cid))
    if not info or (want and info[0].lower() != want):
        continue
    try:
        d = json.load(open(f))
    except Exception:
        continue
    for e, raw in raw_heals(d.get('events') or []):
        if e['source_id'] != cid:
            continue
        key = (cid, info[1], json.dumps(info[2]), e['spell_id'])
        names[e['spell_id']] = e['spell_name']
        agg[key][1 if e['is_critical'] else 0].append(raw)

for key in sorted(agg, key=lambda k: (k[1], k[0], k[3])):
    nc, cr = agg[key]
    if len(nc) < mn or len(cr) < 5:
        continue
    m0, m1 = sum(nc) / len(nc), sum(cr) / len(cr)
    sd0 = math.sqrt(sum((x - m0) ** 2 for x in nc) / max(len(nc) - 1, 1)) / math.sqrt(len(nc))
    sd1 = math.sqrt(sum((x - m1) ** 2 for x in cr) / max(len(cr) - 1, 1)) / math.sqrt(len(cr))
    ratio = m1 / m0
    err = ratio * math.sqrt((sd0 / m0) ** 2 + (sd1 / m1) ** 2)
    print(f"char {key[0]} L{key[1]} {key[2]} {names[key[3]]} {key[3]}: n {len(nc)}/{len(cr)} crit% {100*len(cr)/(len(nc)+len(cr)):.1f} "
          f"mean {m0:.0f}/{m1:.0f} ratio {ratio:.3f} +/- {err:.3f}")
