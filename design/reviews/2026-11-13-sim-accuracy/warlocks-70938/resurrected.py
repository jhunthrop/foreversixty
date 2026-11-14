"""resurrected.py <build dir>: every item whose newest pushed ItemSparse row (raw/wago-push/items.json) leaves out a
primary stat or hit/crit that the published class items still carry, with the stale Wowhead value it came from."""
import glob, json, sys
b = sys.argv[1]
store = json.load(open(f'{b}/raw/wago-push/items.json'))['entries']['ItemSparse']
best = {}
for r in store:
    if r[3] == 1 and isinstance(r[4], list) and (r[2] not in best or (r[1], r[0]) > (best[r[2]][1], best[r[2]][0])):
        best[r[2]] = r
ST = {3: 'agility', 4: 'strength', 5: 'intellect', 6: 'spirit', 7: 'stamina', 31: 'hit', 32: 'crit', 36: 'haste',
      38: 'attack_power', 41: 'healing', 42: 'spell_power', 43: 'mp5', 45: 'spell_power'}
WATCH = {'agility', 'strength', 'intellect', 'spirit', 'stamina', 'hit', 'crit'}
wh = json.load(open(f'{b}/raw/wowhead/items.json'))['items']
items = {}
for f in glob.glob(f'{b}/items/*.json'):
    for it in json.load(open(f))['items']:
        items[it['id']] = it
rows = []
for k, r in best.items():
    if k not in items or str(k) not in wh:
        continue
    types = [t for t in r[4][17] if isinstance(t, int) and t > 0]
    native = {ST.get(t, str(t)) for t in types}
    extra = {s: v for s, v in items[k]['stats'].items() if s in WATCH and s not in native}
    if extra:
        rows.append((k, items[k]['name'], r[1], sorted(native), extra, wh[str(k)]['stats']))
print(len(best), 'pushed ItemSparse records;', len(rows), 'publish a primary or hit/crit stat their pushed row leaves out')
for x in sorted(rows):
    print(x[0], x[1], 'push', x[2], 'row', x[3], 'published extra', x[4], 'wowhead', {k: v for k, v in x[5].items() if k not in ('maxcount', 'reqlevel')})
