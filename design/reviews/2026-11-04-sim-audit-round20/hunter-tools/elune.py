# Pet crit inside and outside Elune's Light (1259799) windows; Heart of the Lion (409580/409583) on pets.
import json, glob, os, collections, sys
D = sys.argv[1]
ELUNE = 1259799
HOTL = (409580, 409583)
evtypes = collections.Counter()
hotl_pet = collections.Counter(); hotl_any = collections.Counter()
elune_hunters = collections.Counter()
inside = collections.Counter(); outside = collections.Counter()
per_hunter = collections.defaultdict(lambda: {'in': collections.Counter(), 'out': collections.Counter()})
files = sorted(glob.glob(D + '/*.json'))
for f in files:
    try:
        ev = json.load(open(f))['events']
    except Exception:
        continue
    ev.sort(key=lambda e: e['timestamp_ms'])
    for e in ev:
        evtypes[e['event_type']] += 1
    # Elune windows per hunter
    windows = collections.defaultdict(list)
    for e in ev:
        if e['spell_id'] == ELUNE and e['event_type'] in ('buff_applied', 'cast') and e['source_type'] == 'player':
            windows[e['source_id']].append(e['timestamp_ms'])
            elune_hunters[(e['source_name'], e.get('source_class'))] += 1
        if e['spell_id'] in HOTL and e['event_type'] == 'buff_applied':
            hotl_any[(e['spell_id'], e.get('target_type'))] += 1
            if e.get('target_type') == 'pet':
                hotl_pet[(e['spell_id'], e['target_name'], e.get('target_owner_name'))] += 1
    for e in ev:
        if e['event_type'] != 'damage' or e.get('source_type') != 'pet' or e.get('spell_school') != 'Physical':
            continue
        if e.get('is_glancing'):
            continue
        own = e.get('source_owner_id')
        if own not in windows:
            continue
        t = e['timestamp_ms']
        inwin = any(0 <= t - w <= 15000 for w in windows[own])
        key = 'crit' if e['is_critical'] else 'hit'
        (inside if inwin else outside)[key] += 1
        per_hunter[(e.get('source_owner_name'))]['in' if inwin else 'out'][key] += 1
print('files', len(files), 'event types', dict(evtypes))
print('Elune casters', dict(elune_hunters))
print('HotL buff_applied by target type', dict(hotl_any))
print('HotL on pets', dict(hotl_pet))
def rate(c): n = c['crit'] + c['hit']; return (n, round(100 * c['crit'] / n, 1) if n else None)
print('pet swings of Elune hunters: inside', rate(inside), 'outside', rate(outside))
for h, d in per_hunter.items():
    print('  ', h, 'in', rate(d['in']), 'out', rate(d['out']))
