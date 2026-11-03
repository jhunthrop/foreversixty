"""Fetch damage timelines (every damage event) for every boss kill and a seeded sample of
trash segments (>= 10 s, up to 12 per report) in the cached foreverlogs.gg reports."""
import datetime, glob, json, os, random, subprocess, sys, time
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
OUT = SP + '/r12/healprofile/fl/dt'
FL = SP + '/r3/casters/py/fl.sh'
def dur(e):
    a = datetime.datetime.fromisoformat(e['start_time'].replace('Z', '+00:00'))
    b = datetime.datetime.fromisoformat(e['end_time'].replace('Z', '+00:00'))
    return (b - a).total_seconds()
reps = {int(f.split('/')[-1].split('.')[0]): json.load(open(f)) for f in glob.glob(SP + '/fl/rep/*[0-9].json')}
boss = {e['name'] for r in reps.values() for e in r['encounters'] if e['success']}
jobs = []
rng = random.Random(12)
for rid in sorted(reps):
    encs = reps[rid]['encounters']
    jobs += [(rid, e['id']) for e in encs if e['success']]
    trash = [e for e in encs if e['name'] not in boss and dur(e) >= 10]
    rng.shuffle(trash)
    jobs += [(rid, e['id']) for e in trash[:12]]
print(len(jobs), 'jobs', file=sys.stderr, flush=True)
for i, (rid, eid) in enumerate(jobs):
    fn = f'{OUT}/{rid}.{eid}.json'
    if os.path.exists(fn) and os.path.getsize(fn) > 0:
        continue
    for attempt in range(5):
        data = subprocess.run(['sh', FL, f'reports/{rid}/encounters/{eid}/timeline?event_types=damage'], capture_output=True, text=True).stdout
        if '"success":true' in data:
            break
        print('retry', rid, eid, data[:80], file=sys.stderr, flush=True)
        time.sleep(20 * (attempt + 1))
    else:
        continue
    open(fn, 'w').write(data)
    if i % 100 == 0:
        print(i, file=sys.stderr, flush=True)
    time.sleep(0.5)
print('done', file=sys.stderr, flush=True)
