"""Beta logs: Penance pulses per volley (pulse casts of 402289 within 1.3 s of each other)."""
import collections, glob, json
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
vol = collections.Counter()
for f in glob.glob(SP + '/r10/healers/fl/tl/*.112314.json'):
    ts = sorted({e['timestamp_ms'] for e in json.load(open(f))['events'] if e['spell_id'] == 402289 and e['event_type'] == 'cast'})
    grp = []
    for t in ts:
        if grp and t - grp[-1] <= 1300: grp.append(t)
        else:
            if grp: vol[len(grp)] += 1
            grp = [t]
    if grp: vol[len(grp)] += 1
print('volley sizes', sorted(vol.items()))
