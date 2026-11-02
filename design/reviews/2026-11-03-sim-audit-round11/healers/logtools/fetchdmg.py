import json, subprocess, sys, time, os, datetime
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
OUT = SP + '/r11/healers/fl/dmg'
FL = SP + '/r3/casters/py/fl.sh'
chars = [(2747,124305),(2747,122644),(2747,112314),(2700,78866),(2712,104302),(2763,143935),(2785,157083),(2706,91404),(2764,148830),(2745,120434),(2749,129798),(2726,108492)]
def dur(e):
    a = datetime.datetime.fromisoformat(e['start_time'].replace('Z','+00:00')); b = datetime.datetime.fromisoformat(e['end_time'].replace('Z','+00:00'))
    return (b-a).total_seconds()
for rid, cid in chars:
    rep = json.load(open(SP + f'/fl/rep/{rid}.json'))
    encs = sorted(rep['encounters'], key=dur, reverse=True)[:14]
    for e in encs:
        fn = f'{OUT}/{rid}.{e["id"]}.{cid}.json'
        if os.path.exists(fn) and os.path.getsize(fn) > 0: continue
        q = f'reports/{rid}/encounters/{e["id"]}/timeline?sourceCharacterIds={cid}&event_types=damage'
        for attempt in range(5):
            data = subprocess.run(['sh', FL, q], capture_output=True, text=True).stdout
            if '"success":true' in data: break
            print('retry', rid, e['id'], data[:80], file=sys.stderr, flush=True); time.sleep(20*(attempt+1))
        else: continue
        open(fn, 'w').write(data); time.sleep(0.7)
print('done', file=sys.stderr)
