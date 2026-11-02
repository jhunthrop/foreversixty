"""Beta logs: a paladin's raw non-crit Holy Light, Flash of Light and Holy Shock heals (one character, one multiplier)."""
import collections, glob, json, statistics
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
agg = collections.defaultdict(lambda: collections.defaultdict(list))
for f in glob.glob(SP + '/r10/healers/fl/tl/*.json'):
    rid, enc, cid = map(int, f.split('/')[-1].split('.')[:3])
    try: ev = json.load(open(f))['events']
    except Exception: continue
    over = collections.defaultdict(float)
    for e in ev:
        if e['event_type'] == 'overheal': over[(e['timestamp_ms'], e['target_id'], e['spell_id'])] += float(e['amount'])
    for e in ev:
        if e['event_type'] == 'heal' and e['source_id'] == cid and e['spell_id'] in (1311605, 647, 1026, 1042, 19939, 19750, 19940):
            agg[(cid, e['source_name'])][(e['spell_id'], e['is_critical'])].append(float(e['amount']) + over.get((e['timestamp_ms'], e['target_id'], e['spell_id']), 0))
for k, d in agg.items():
    if (1311605, False) in d:
        print(k, {f"{s}{'c' if c else ''}": (len(v), round(statistics.mean(v), 1)) for (s, c), v in sorted(d.items())})
