"""Fetch damage timelines for chosen frost mages: fetch.py rid:cid ..."""
import json, subprocess, sys, time, os
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
OUT = SP + '/r19/mage/il/dmg'
FL = SP + '/r3/casters/py/fl.sh'
os.makedirs(OUT, exist_ok=True)
for arg in sys.argv[1:]:
    rid, cid = map(int, arg.split(':'))
    rep = json.load(open(SP + f'/fl/rep/{rid}.json'))
    for e in rep['encounters']:
        fn = f'{OUT}/{rid}.{e["id"]}.{cid}.json'
        if os.path.exists(fn) and os.path.getsize(fn) > 0: continue
        q = f'reports/{rid}/encounters/{e["id"]}/timeline?sourceCharacterIds={cid}&event_types=damage'
        for attempt in range(4):
            data = subprocess.run(['sh', FL, q], capture_output=True, text=True).stdout
            if '"success":true' in data: break
            time.sleep(15 * (attempt + 1))
        else: continue
        open(fn, 'w').write(data); time.sleep(0.5)
    print('done', rid, cid, file=sys.stderr, flush=True)
