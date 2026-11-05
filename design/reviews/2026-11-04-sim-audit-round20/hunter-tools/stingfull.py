# Serpent Sting ticks per hunter and rank in the buff-bearing beta encounters, with the
# pet's presence, Curse of the Elements on the target, Arcane Shot and the ranged weapon cadence.
import json, glob, os, collections, statistics, sys
D = sys.argv[1]
STING = {1978: 1, 13549: 2, 13550: 3, 13551: 4, 13552: 5, 13553: 6, 13554: 7, 13555: 8}
ARC = {3044: 1, 14281: 2, 14282: 3, 14283: 4, 14284: 5, 14285: 6, 14286: 7, 14287: 8}
COE_NAMES = ('Curse of the Elements', 'Curse of Shadow')
per = collections.defaultdict(lambda: {'ticks': collections.Counter(), 'arc': collections.Counter(), 'auto': [], 'autots': [], 'multi': [], 'aimed': collections.defaultdict(list), 'pet': 0, 'coe': 0, 'lw': 0, 'encs': set(), 'buffs': collections.Counter()})
for f in sorted(glob.glob(D + '/*.json')):
    rid, eid, cid = os.path.basename(f).split('.')[:3]
    try:
        ev = json.load(open(f))['events']
    except Exception:
        continue
    ev.sort(key=lambda e: e['timestamp_ms'])
    hunters = {e['source_id'] for e in ev if e.get('source_class') == 'Hunter' and e['event_type'] == 'damage'}
    coe = any(e['event_type'] == 'buff_applied' and e.get('spell_name') in COE_NAMES and e.get('target_type') == 'creature' for e in ev)
    for h in hunters:
        name = next(e['source_name'] for e in ev if e['source_id'] == h)
        p = per[(rid, name)]
        p['encs'].add(eid)
        if coe: p['coe'] += 1
        if any(e.get('source_owner_id') == h and e['event_type'] == 'damage' for e in ev): p['pet'] += 1
        for e in ev:
            if e['source_id'] == h and e['event_type'] == 'buff_applied' and e.get('target_id') == h:
                p['buffs'][(e['spell_name'], e['spell_id'])] += 1
            if e['source_id'] != h or e['event_type'] != 'damage' or e['is_critical']:
                continue
            s, a = e['spell_id'], int(e['amount'])
            if s in STING: p['ticks'][(STING[s], a)] += 1
            elif s in ARC: p['arc'][(ARC[s], a)] += 1
            elif e['spell_name'] == 'Auto Shot': p['auto'].append(a); p['autots'].append(e['timestamp_ms'])
            elif e['spell_name'] == 'Multi-Shot': p['multi'].append(a)
            elif e['spell_name'] == 'Aimed Shot': p['aimed'][s].append(a)
for k, p in sorted(per.items()):
    if not p['ticks']: continue
    gaps = [b - a for a, b in zip(p['autots'], p['autots'][1:]) if 1500 < b - a < 4000]
    print('==', k, 'encounters', len(p['encs']), 'pet in', p['pet'], 'CoE in', p['coe'])
    byrank = collections.defaultdict(list)
    for (r, a), n in p['ticks'].items(): byrank[r].append((a, n))
    for r in sorted(byrank): print('   sting rank', r, sorted(byrank[r]))
    byr = collections.defaultdict(list)
    for (r, a), n in p['arc'].items(): byr[r].append((a, n))
    for r in sorted(byr): print('   arcane rank', r, sorted(byr[r]))
    if p['auto']: print('   auto n', len(p['auto']), 'mean', round(statistics.mean(p['auto']), 1), 'swing ms mode', collections.Counter(round(g, -2) for g in gaps).most_common(2))
    if p['multi']: print('   multi n', len(p['multi']), 'mean', round(statistics.mean(p['multi']), 1))
    for s, v in p['aimed'].items(): print('   aimed', s, 'n', len(v), 'mean', round(statistics.mean(v), 1))
    print('   self buffs', [b for b in p['buffs'].most_common(12)])
