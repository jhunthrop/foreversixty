"""ss_gaps.py <shaman events parquet>: every gap between a shaman's Stormstrike casts under 30 s in our reports, and
whether the shaman dodged or parried an attack in between (Improved Stormstrike's reset)."""
import sys, collections
import pandas as pd
df = pd.read_parquet(sys.argv[1]); out = collections.Counter(); short = []
for g, t in df.groupby('source_guid'):
    ss = sorted(t[(t.event == 'SPELL_CAST_SUCCESS') & (t.spell_id == 17364)].time_unix_nano)
    if len(ss) < 2: continue
    inc = df[(df.dest_guid == g) & df.event.isin(['SWING_MISSED', 'SPELL_MISSED', 'RANGE_MISSED']) & df.miss_type.isin(['DODGE', 'PARRY'])].time_unix_nano
    for a, b in zip(ss, ss[1:]):
        gap = (b - a) / 1e9
        if gap >= 30: continue
        av = int(((inc > a) & (inc < b)).sum())
        k = '<7.9' if gap < 7.9 else '7.9-8.6' if gap < 8.6 else '>=8.6'
        out[(k, 'dodge/parry between' if av else 'none between')] += 1
        if gap < 7.9: short.append((round(gap, 2), av))
for k, v in sorted(out.items()): print(k, v)
print('short gaps (gap, avoids between):', sorted(short))
