"""hawk3.py: lifetime of each Hawk against the same hunter's next Summon Hawk, and its target's survival (reads data/hawk_units.json)."""
import json, collections, statistics as st
U = json.load(open('data/hawk_units.json'))['units']
by = collections.defaultdict(list)
for u in U: by[(u['fp'], u['hunter'])].append(u)
rows = []
for us in by.values():
    us.sort(key=lambda u: u['ts'])
    for a, b in zip(us, us[1:] + [None]): rows.append((a, (b['ts'] - a['ts']) if b else None))
sw = [(u, n) for u, n in rows if u['swings']]
def dist(xs): return sorted(collections.Counter(xs).items())
print('hawks', len(rows), 'with >=1 swing', len(sw), '| hunters', len(by and {u['hunter'] for u in U}))
print('all swinging hawks: swings', dist(len(u['swings']) for u, _ in sw), 'last-swing dt med', round(st.median(u['swings'][-1]['dt'] for u, _ in sw), 2))
free = [(u, n) for u, n in sw if (n is None or n >= 18) and u['fight_left'] >= 18]
print('A) no recast within 18 s, fight >=18 s:', len(free), 'swings', dist(len(u['swings']) for u, _ in free))
for lab, sel in (('target alive >=18 s', lambda u: u['alive18']), ('target died <18 s', lambda u: not u['alive18'])):
    xs = [u for u, _ in free if sel(u)]
    print('  ', lab, len(xs), 'swings', dist(len(u['swings']) for u in xs), 'last dt', sorted(round(u['swings'][-1]['dt'], 1) for u in xs),
          'hawk last seen', sorted(u['hawk_last_seen'] for u in xs))
cut = [(u, n) for u, n in sw if n is not None and n < 18]
print('B) recast by same hunter <18 s:', len(cut), 'swings after the recast:', sum(u['swings'][-1]['dt'] > n + 0.3 for u, n in cut),
      'recast gap med', round(st.median(n for _, n in cut), 2))
dead = [u for u, _ in sw if u['target_dies_at'] is not None and u['target_dies_at'] < 17]
print('C) target died <17 s:', len(dead), 'kept swinging after death:', sum(any(s['dt'] > u['target_dies_at'] + 0.2 for s in u['swings']) for u in dead),
      'of which on a different target:', sum(any(s['dt'] > u['target_dies_at'] + 0.2 and not s['same_target'] for s in u['swings']) for u in dead),
      'stopped at death:', sum(not any(s['dt'] > u['target_dies_at'] + 0.2 for s in u['swings']) for u in dead))
print('   last swing dt for hawks that continued:', sorted(round(u['swings'][-1]['dt'], 1) for u in dead if any(s['dt'] > u['target_dies_at'] + 0.2 for s in u['swings'])))
print('max last-swing dt over all hawks', max(u['swings'][-1]['dt'] for u, _ in sw), '| hawks with a swing beyond 18.5 s:', sum(u['swings'][-1]['dt'] > 18.5 for u, _ in sw))
print('first swing dt (from summon):', 'med', round(st.median(u['swings'][0]['dt'] for u, _ in sw), 2), 'share <0.1 s', round(sum(u['swings'][0]['dt'] < .1 for u, _ in sw) / len(sw), 2))
