"""Fetch every encounter timeline (heals, overheals, casts, buffs) for chosen healers."""
import json, subprocess, sys, time, os, glob
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
OUT = SP + '/r10/healers/fl/tl'
FL = SP + '/r3/casters/py/fl.sh'
want = []
for f in glob.glob(SP + '/fl/rep/*.roster.json'):
    rid = int(f.split('/')[-1].split('.')[0])
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            cls = s['player']['class']
            try:
                tabs = s['talents']['talent_groups'][str(s['talents']['active_group'])]['tabs']
            except Exception:
                continue
            pts = {v['name']: v['points'] for v in tabs.values()}
            L = sum(pts.values()) + 9
            heal = {'PALADIN': 'Holy', 'DRUID': 'Restoration', 'PRIEST': None, 'SHAMAN': 'Restoration'}.get(cls, 'x')
            if heal == 'x':
                continue
            if heal is None:
                ok = pts.get('Shadow', 0) < 10 and L >= 24
            else:
                ok = pts.get(heal, 0) >= 15
            if ok:
                want.append((rid, c['character_id']))
PRI={'PALADIN':0,'DRUID':1,'SHAMAN':2,'PRIEST':3}
cls_of={}
for f in glob.glob(SP + '/fl/rep/*.roster.json'):
    rid = int(f.split('/')[-1].split('.')[0])
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            cls_of[(rid, c['character_id'])] = s['player']['class']
want.sort(key=lambda k: PRI.get(cls_of.get(k), 9))
print(len(want), 'chars', file=sys.stderr)
for rid, cid in want:
    rep = json.load(open(SP + f'/fl/rep/{rid}.json'))
    for e in rep['encounters']:
        fn = f'{OUT}/{rid}.{e["id"]}.{cid}.json'
        if os.path.exists(fn) and os.path.getsize(fn) > 0:
            continue
        q = f'reports/{rid}/encounters/{e["id"]}/timeline?sourceCharacterIds={cid}&event_types=heal,overheal,cast,buff_applied,buff_removed,absorb'
        for attempt in range(8):
            data = subprocess.run(['sh', FL, q], capture_output=True, text=True).stdout
            if '"success":true' in data:
                break
            print('retry', rid, e['id'], cid, data[:60], file=sys.stderr, flush=True)
            time.sleep(30 * (attempt + 1))
        else:
            continue
        open(fn, 'w').write(data)
        time.sleep(0.6)
print('done', file=sys.stderr)
