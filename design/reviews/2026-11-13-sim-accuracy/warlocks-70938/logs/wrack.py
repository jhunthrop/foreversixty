"""wrack.py <dl dir> <report:guid-suffix>...: every Wrack channel in the logs, its own tick times and amounts,
and the gaps between consecutive ticks of each DoT inside a channel, outside, and across its start and end."""
import glob, sys, collections
import pandas as pd, pyarrow.parquet as pq
WRACK = 1316697
DOTS = {7648: 'Corruption', 11711: 'Agony', 11712: 'Agony', 11713: 'Agony', 18265: 'Siphon', 18879: 'Siphon', 18880: 'Siphon', 18881: 'Siphon',
        172: 'Corruption', 6222: 'Corruption', 6223: 'Corruption', 7648: 'Corruption', 11671: 'Corruption', 11672: 'Corruption', 25311: 'Corruption'}
dl = sys.argv[1]
gaps = collections.defaultdict(list)
for arg in sys.argv[2:]:
    rep, suf = arg.split(':')
    for f in sorted(glob.glob(f'{dl}/{rep}/f*.parquet')):
        t = pq.read_table(f).to_pandas()
        t = t[t.source_guid.astype(str).str.endswith(suf)]
        if not (t.spell_id == WRACK).any():
            continue
        t = t.sort_values('time_unix_nano'); t['t'] = t.time_unix_nano / 1e9
        ap = t[(t.spell_id == WRACK) & (t.event == 'SPELL_AURA_APPLIED')].t.tolist()
        rm = t[(t.spell_id == WRACK) & (t.event == 'SPELL_AURA_REMOVED')].t.tolist()
        chans = []
        for a in ap:
            e = min([r for r in rm if r >= a] or [a + 6.0])
            chans.append((a, e))
        print(f'{rep}:{suf} {f.split("/")[-1]} channels', [(round(a % 1000, 2), round(e - a, 2)) for a, e in chans])
        wt = t[(t.spell_id == WRACK) & (t.event == 'SPELL_PERIODIC_DAMAGE')]
        for a, e in chans:
            ticks = wt[(wt.t >= a) & (wt.t <= e + 0.05)]
            print('   wrack ticks at', [round(x - a, 2) for x in ticks.t], 'base', ticks.base_amount.tolist(), 'crit', ticks.critical.tolist())
        inside = lambda x: any(a <= x <= e for a, e in chans)
        for sid, nm in DOTS.items():
            d = t[(t.spell_id == sid) & (t.event == 'SPELL_PERIODIC_DAMAGE')]
            # only gaps between ticks on the same target
            for dest, g in d.groupby('dest_guid'):
                ts = g.t.tolist()
                for p, q in zip(ts, ts[1:]):
                    if q - p > 3.6:
                        continue
                    kind = 'in' if inside(p) and inside(q) else 'out' if not inside(p) and not inside(q) else ('start' if inside(q) else 'end')
                    gaps[(nm, kind)].append(round(q - p, 3))
for k in sorted(gaps):
    v = sorted(gaps[k]); print(k, len(v), 'median', v[len(v) // 2], v)
