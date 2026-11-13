"""own_ssmodel.py <shaman events parquet> <guid suffix>...: unhasted swing speed (gaps with Flurry down), attack power at
the shaman's casts, and Stormstrike against white hits before armor (base_amount): the logged SS minus white median
difference against the 100% normalized model AP x (3.3 - speed) / 14 and the 150% model."""
import sys, statistics as st
import pandas as pd
df = pd.read_parquet(sys.argv[1])
for suf in sys.argv[2:]:
    t = df[df.source_guid.str.endswith(suf, na=False)].sort_values('time_unix_nano'); g = t.source_guid.iloc[0]
    fl = t[(t.spell_id == 16257) & (t.dest_guid == g)]
    up = []
    for e in fl.itertuples():
        if e.event in ('SPELL_AURA_APPLIED',): up.append([e.time_unix_nano, None])
        elif e.event == 'SPELL_AURA_REMOVED' and up and up[-1][1] is None: up[-1][1] = e.time_unix_nano
    def flurry(ts): return any(a <= ts and (b is None or ts <= b) for a, b in up)
    gaps = []
    for _, f in t[t.event.isin(['SWING_DAMAGE', 'SWING_MISSED'])].groupby('fight'):
        x = f.time_unix_nano.tolist()
        gaps += [(b - a) / 1e9 for a, b in zip(x, x[1:]) if (b - a) < 6e9 and not flurry(a)]
    ap = t[(t.event == 'SPELL_CAST_SUCCESS') & (t.adv_info_guid == g)].adv_attack_power.median()
    w = t[(t.event == 'SWING_DAMAGE') & ~t.critical.fillna(False).astype(bool) & ~t.glancing.fillna(False).astype(bool)].base_amount
    s = t[(t.event == 'SPELL_DAMAGE') & (t.spell_id == 17364) & ~t.critical.fillna(False).astype(bool)].base_amount
    spd = min(gaps) if gaps else None
    print(f'{suf}: unhasted swing gaps n {len(gaps)} min {spd} median {st.median(gaps) if gaps else None}; AP median {ap}; white base median {w.median()} (n {len(w)}), SS base median {s.median()} (n {len(s)})')
    for sp in sorted({round(x, 1) for x in gaps})[:3]:
        u = w.median() - ap * sp / 14
        print(f'   if speed {sp}: weapon roll median {u:.0f}; 100% model SS {u + ap * 3.3 / 14:.0f}, 150% model {1.5 * (u + ap * 3.3 / 14):.0f}, logged {s.median()}')
