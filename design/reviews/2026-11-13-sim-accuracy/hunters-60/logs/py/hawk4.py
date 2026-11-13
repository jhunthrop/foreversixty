"""hawk4.py: hawk swing base (pre-armor) vs hunter level, hawk AP, hunter AP; pre/post 11 Nov split with time-zone margin."""
import json, statistics as st, collections
U = json.load(open('data/hawk_units.json'))['units']; HOT = 1794355200
pts = [(u['hunter_lvl'], s['base'], s['amt'], s['hawk_ap'], u['hunter_ap'], u['ts']) for u in U for s in u['swings']
       if not s['miss'] and not s['crit'] and not s['glance'] and s['base'] and u['hunter_lvl']]
print('non-crit non-glance landed swings with base:', len(pts), 'hunters', len({u['hunter'] for u in U if u['swings']}))
print('base/level: med', round(st.median(b / l for l, b, *_ in pts), 3), 'p10', round(sorted(b / l for l, b, *_ in pts)[len(pts) // 10], 3), 'p90', round(sorted(b / l for l, b, *_ in pts)[9 * len(pts) // 10], 3))
print('amount/base (armor) med', round(st.median(a / b for _, b, a, *_ in pts), 3))
def corr(xs, ys):
    mx, my = st.mean(xs), st.mean(ys); sx = sum((x - mx) ** 2 for x in xs) ** .5; sy = sum((y - my) ** 2 for y in ys) ** .5
    return sum((x - mx) * (y - my) for x, y in zip(xs, ys)) / (sx * sy)
h = [(b, hap) for _, b, _, hap, *_ in pts if hap]
print('corr(base, hawk AP)', round(corr([x[0] for x in h], [x[1] for x in h]), 3), 'n', len(h))
hu = [(b, ap) for _, b, _, _, ap, _ in pts if ap]
print('corr(base, hunter AP)', round(corr([x[0] for x in hu], [x[1] for x in hu]), 3))
lv = [(l, b) for l, b, *_ in pts]
print('corr(base, level)', round(corr([x[0] for x in lv], [x[1] for x in lv]), 3))
by = collections.defaultdict(list)
for l, b, a, hap, ap, ts in pts: by[l].append(b)
print('base by hunter level (n, min, med, max, mean):', {l: (len(v), min(v), st.median(v), max(v), round(st.mean(v), 1)) for l, v in sorted(by.items())})
for lab, lo, hi in (('pre, before 10 Nov 21:00 UTC-as-logged (3 h margin)', 0, HOT - 3 * 3600), ('pre, last 3 h', HOT - 3 * 3600, HOT), ('post', HOT, 9e18)):
    xs = [(l, b) for l, b, a, hap, ap, ts in pts if lo <= ts < hi]
    if xs: print(lab, 'n', len(xs), 'base/level med', round(st.median(b / l for l, b in xs), 3), 'levels', sorted(collections.Counter(l for l, _ in xs).items()))
