"""Beta logs: each heal over time's duration (buff_applied to buff_removed, no recast between) and ticks per application."""
import collections, glob, json
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
HOTS = {'Renew', 'Rejuvenation', 'Regrowth', 'Riptide', 'Tranquility', 'Lightwell Renew', 'Healing Stream'}
ticks = collections.defaultdict(collections.Counter); durs = collections.defaultdict(collections.Counter)
for f in glob.glob(SP + '/r10/healers/fl/tl/*.json'):
    rid, enc, cid = map(int, f.split('/')[-1].split('.')[:3])
    try: ev = json.load(open(f))['events']
    except Exception: continue
    app = {}
    for e in sorted((e for e in ev if e['source_id'] == cid), key=lambda e: e['timestamp_ms']):
        if e['spell_name'] not in HOTS: continue
        k = (e['spell_name'], e['spell_id'], e['target_id'])
        if e['event_type'] == 'buff_applied': app[k] = [e['timestamp_ms'], 0, False]
        elif e['event_type'] == 'cast' and k in app: app[k][2] = True
        elif e['event_type'] == 'heal' and k in app and e['timestamp_ms'] - app[k][0] >= 500: app[k][1] += 1
        elif e['event_type'] == 'buff_removed' and k in app:
            t0, nt, ref = app.pop(k)
            if not ref:
                durs[k[:2]][round((e['timestamp_ms'] - t0) / 500) * 0.5] += 1; ticks[k[:2]][nt] += 1
for k in sorted(durs, key=lambda k: -sum(durs[k].values())):
    if sum(durs[k].values()) >= 10: print(k, 'dur', durs[k].most_common(4), 'ticks', ticks[k].most_common(4))
