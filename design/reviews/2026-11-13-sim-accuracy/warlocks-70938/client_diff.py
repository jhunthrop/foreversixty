"""client_diff.py <old build dir> <new build dir>: every warlock spellconst row that differs between two builds
(talents and trainables are compared whole), then every push entry since 10 Nov (push 113700 on) whose record
or spell id is a warlock spell."""
import json, sys
old, new = sys.argv[1:3]
a = json.load(open(f'{old}/spellconst/warlock.json'))['spells']
b = json.load(open(f'{new}/spellconst/warlock.json'))['spells']
for k in sorted(set(a) | set(b), key=int):
    if a.get(k) == b.get(k):
        continue
    if k not in a or k not in b:
        print('added' if k in b else 'removed', k); continue
    d = [f'{f}:{a[k][f]}->{b[k][f]}' for f in b[k] if f != 'effects' and a[k].get(f) != b[k][f]]
    for i, (ea, eb) in enumerate(zip(a[k]['effects'], b[k]['effects'])):
        d += [f'e{i}.{g}:{ea.get(g)}->{eb[g]}' for g in eb if ea.get(g) != eb[g]]
    print(k, b[k]['name'], b[k]['rank'], d)
for part in ('talents', 'trainables'):
    same = json.load(open(f'{old}/{part}/warlock.json')) | {'build': ''} == json.load(open(f'{new}/{part}/warlock.json')) | {'build': ''}
    print(part, 'identical' if same else 'DIFFER')
ids = {int(k) for k in b}
v = json.load(open(f'{new}/raw/wago-push/values.json'))['entries']
seen = set()
for tb, rows in v.items():
    for r in rows:
        if r[1] < 113700 or not (tb.startswith('Spell') or tb.startswith('Talent')):
            continue
        last = r[4][-1] if isinstance(r[4], list) and r[4] else None
        sid = r[2] if r[2] in ids else last if isinstance(last, int) and last in ids else None
        if sid and (r[1], tb, r[2]) not in seen:
            seen.add((r[1], tb, r[2])); print('push', r[1], tb, 'record', r[2], 'spell', sid, b[str(sid)]['name'])
