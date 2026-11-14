"""clock.py <dl dir> <report:guid-suffix>...: for each DoT tick in flight when a Wrack channel starts or ends,
the observed tick time against the two models: the hotfix note's (the clock skips: the remaining time halves at
the start and doubles at the end) and a next-tick model (the in-flight tick keeps its time, the new period applies after)."""
import glob, sys
import pyarrow.parquet as pq
WRACK = 1316697
PERIOD = {7648: 3.0, 11711: 2.0, 18265: 3.0, 18879: 3.0}
NAME = {7648: 'Corruption', 11711: 'Agony', 18265: 'Siphon', 18879: 'Siphon'}
dl = sys.argv[1]
rows = []
for arg in sys.argv[2:]:
    rep, suf = arg.split(':')
    for f in sorted(glob.glob(f'{dl}/{rep}/f*.parquet')):
        t = pq.read_table(f).to_pandas()
        t = t[t.source_guid.astype(str).str.endswith(suf)].sort_values('time_unix_nano')
        if not (t.spell_id == WRACK).any():
            continue
        t['t'] = t.time_unix_nano / 1e9
        w = t[t.spell_id == WRACK]
        for _, ap in w[w.event == 'SPELL_AURA_APPLIED'].iterrows():
            a = ap.t; dest = ap.dest_guid
            rm = w[(w.event == 'SPELL_AURA_REMOVED') & (w.t >= a) & (w.dest_guid == dest)].t
            e = rm.min() if len(rm) else a + 6
            for sid, P in PERIOD.items():
                ticks = t[(t.spell_id == sid) & (t.event == 'SPELL_PERIODIC_DAMAGE') & (t.dest_guid == dest)].t.tolist()
                before = [x for x in ticks if x <= a + 0.02]; after = [x for x in ticks if x > a + 0.02]
                if before and after and a - before[-1] < P:
                    p, q = before[-1], after[0]
                    rows.append(('start', NAME[sid], round(a - p, 2), round(q - p, 3), round(p + P - p, 2), round(a + (P - (a - p)) / 2 - p, 2)))
                inch = [x for x in ticks if a < x <= e + 0.02]; post = [x for x in ticks if x > e + 0.02]
                if inch and post and e - inch[-1] < P / 2:
                    p, q = inch[-1], post[0]
                    rows.append(('end', NAME[sid], round(e - p, 2), round(q - p, 3), round(P / 2, 2), round(e + (P / 2 - (e - p)) * 2 - p, 2)))
print('edge    dot        elapsed  observed  next-tick  hotfix-note   (seconds after the last tick)')
for r in rows:
    print(f'{r[0]:6s} {r[1]:10s} {r[2]:7.2f} {r[3]:9.3f} {r[4]:10.2f} {r[5]:11.2f}  -> {"next-tick" if abs(r[3]-r[4]) < abs(r[3]-r[5]) else "hotfix-note"}')
