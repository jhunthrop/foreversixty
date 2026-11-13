"""mw_rate.py <shaman events parquet>: Maelstrom Weapon (408505) gains against melee hits, by cluster. A cluster is every
melee event (white swing, Stormstrike, Windfury Weapon strike 439440, Windfury Totem attack 10608) of one shaman within
0.15 s; gains in the cluster are the doses and applies, or the refreshes at five stacks, whichever is more. Shamans are
numbered by first appearance (private report)."""
import sys, collections
import pandas as pd
WFT = {8516, 10608, 10610}
df = pd.read_parquet(sys.argv[1]); n = 0
def kind(e):
    if e.event.startswith('SWING'): return 'white'
    return {17364: 'SS', 439440: 'WFW'}.get(e.spell_id, 'WFT' if e.spell_id in WFT else None)
for g, t in df.groupby('source_guid'):
    mw = t[(t.spell_id == 408505) & t.event.isin(['SPELL_AURA_APPLIED', 'SPELL_AURA_APPLIED_DOSE', 'SPELL_AURA_REFRESH'])]
    if mw.empty: continue
    n += 1
    m = t[t.event.isin(['SWING_DAMAGE', 'SWING_MISSED']) | (t.event.isin(['SPELL_DAMAGE', 'SPELL_MISSED']) & t.spell_id.isin({17364, 439440} | WFT))]
    ss_cast = t[(t.event == 'SPELL_CAST_SUCCESS') & (t.spell_id == 17364)].time_unix_nano
    pts = sorted([(x, kind(e), e.event.endswith('MISSED')) for x, e in zip(m.time_unix_nano, m.itertuples())] + [(x + 1, 'SScast', False) for x in ss_cast])
    clusters, cur = [], []
    for p in pts:
        if cur and p[0] - cur[0][0] > 1.5e8: clusters.append(cur); cur = []
        cur.append(p)
    if cur: clusters.append(cur)
    by = collections.defaultdict(lambda: [0, 0, 0]); tot_hits = collections.Counter(); used = set()
    for c in clusters:
        a, b = c[0][0] - 2e6, c[-1][0] + 1.6e8
        ev = mw[(mw.time_unix_nano >= a) & (mw.time_unix_nano <= b)]
        used |= set(ev.index)
        gains = sum(max(((x.event != 'SPELL_AURA_REFRESH').sum()), (x.event == 'SPELL_AURA_REFRESH').sum()) for _, x in ev.groupby('time_unix_nano'))
        ks = collections.Counter(k for _, k, miss in c if k != 'SScast')
        if 'SScast' in [k for _, k, _ in c] and 'SS' not in ks: ks['SS'] += 1     # Stormstrike whose damage line the logger dropped
        for k, v in ks.items(): tot_hits[k] += v
        key = '+'.join(f'{k}x{v}' for k, v in sorted(ks.items()))
        by[key][0] += 1; by[key][1] += gains; by[key][2] += sum(ks.values())
    orphan = mw.drop(index=list(used & set(mw.index)))
    og = sum(max(((x.event != 'SPELL_AURA_REFRESH').sum()), (x.event == 'SPELL_AURA_REFRESH').sum()) for _, x in orphan.groupby('time_unix_nano'))
    print(f'shaman #{n}: melee events {dict(tot_hits)}, gains matched {sum(v[1] for v in by.values())}, gains with no logged melee event {og}')
    for k, (c, gsum, h) in sorted(by.items(), key=lambda x: -x[1][0]):
        print(f'   {k:28s} clusters {c:3d} gains {gsum:3d} per cluster {gsum / c:.2f} per hit {gsum / h:.2f}')
