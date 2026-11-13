"""own_an.py <shaman events parquet> <min real level>: per shaman in our post-hotfix reports (numbered, never named):
real level from the highest rank cast, swing speed, Stormstrike gaps and size against white hits, Maelstrom Weapon
gains by trigger, stacks at Lightning Bolt, Windfury Weapon and Windfury Totem procs, Flametongue Totem, casts a minute,
damage by ability. Matches spells by id (Portuguese client)."""
import sys, os, json, collections, statistics as st
import pandas as pd
B = os.path.join(os.path.dirname(os.path.abspath(__file__)), '../../../../../data/builds/1.60.1.70938/spellconst/shaman.json')
SC = {int(k): v for k, v in json.load(open(B))['spells'].items()}
SS, MW, MW_READY, WFW, LB_NAME = 17364, 408505, 408501, 439440, 'Lightning Bolt'
WFT = {8516, 10608, 10610}            # Windfury Totem extra attack damage rows
FTT = {8253, 8248, 10523, 16389}      # Flametongue Totem procs
HOTFIX_NS = pd.Timestamp('2026-11-12T18:45:00Z').value
# Log clocks are local time without a zone. ctxicu4cuyek is an EU log uploaded live from 18:06 UTC whose first line reads
# 20:01 and whose last heartbeat (20:53 UTC) matches its last line (22:51): UTC+2. Our live upload ju2t73py6pib is in UTC.
LOCAL_OFFSET_NS = {'ctxicu4cuyek': 2 * 3600 * 10**9}
MELEE_GAP_S = 5.0                     # a melee bout ends after this long without a melee event
df = pd.read_parquet(sys.argv[1]); minlv = int(sys.argv[2])
df = df[df.time_unix_nano - df.rep.map(LOCAL_OFFSET_NS).fillna(0) >= HOTFIX_NS]
pd.set_option('future.no_silent_downcasting', True)
def flag(c): return c.fillna(False).astype(bool)
def q10(xs): return st.quantiles(xs, n=10)[0] if len(xs) >= 10 else min(xs) if xs else None
def real_level(c):
    lv = [SC[s]['spell_level'] for s in c[c.event == 'SPELL_CAST_SUCCESS'].spell_id if s in SC and SC[s]['rank'] > 0]
    return max(lv) if lv else 0
def name(sid): return SC[sid]['name'] if sid in SC else {0: 'white', WFW: 'Windfury Weapon'}.get(sid, str(sid))
def bouts(ts):
    """Total melee time: sum of spans of melee bouts, each padded by one swing."""
    ts = sorted(ts); tot, start, prev = 0.0, None, None
    for t in ts:
        if prev is None or t - prev > MELEE_GAP_S: tot += (prev - start) if prev is not None else 0; start = t
        prev = t
    return tot + ((prev - start) if prev is not None else 0)
def mw_gains(e):
    """Stack gains at one timestamp: a gain logs a dose (or an apply), at the cap a refresh; some gains log both."""
    n = collections.Counter(e.event)
    return max(n['SPELL_AURA_APPLIED'] + n['SPELL_AURA_APPLIED_DOSE'], n['SPELL_AURA_REFRESH'])
n = 0
for (rep, g), t in df.groupby(['rep', 'source_guid']):
    if not str(g).startswith('Player'): continue
    t = t.sort_values(['time_unix_nano', 'line'] if 'line' in t else 'time_unix_nano')
    casts = t[t.event == 'SPELL_CAST_SUCCESS']
    if SS not in set(casts.spell_id) and MW not in set(t.spell_id): continue
    lvl = real_level(casts)
    if lvl < minlv: continue
    n += 1; ts = t.time_unix_nano / 1e9
    sw = t[t.event.isin(['SWING_DAMAGE', 'SWING_MISSED'])]
    swl = t[t.event == 'SWING_DAMAGE']
    print(f'== {rep[:4]}.. shaman #{n} real level >= {lvl} (adv {t.adv_level.dropna().max()}) fights {t.fight.nunique()} '
          f'span {pd.to_datetime(t.time_unix_nano.min() - LOCAL_OFFSET_NS.get(rep, 0))} to {pd.to_datetime(t.time_unix_nano.max() - LOCAL_OFFSET_NS.get(rep, 0))} UTC')
    # swing speed: gaps between consecutive white swings, same fight
    gaps = []
    for _, f in sw.groupby('fight'):
        x = (f.time_unix_nano / 1e9).tolist(); gaps += [b - a for a, b in zip(x, x[1:]) if b - a < 6]
    print(f'   white swings {len(sw)} (landed {len(swl)}), swing gap median {st.median(gaps) if gaps else None:.3f} p90 {st.quantiles(gaps, n=10)[-1] if len(gaps) >= 10 else None}')
    # Stormstrike
    sc = casts[casts.spell_id == SS]; sg = []
    for _, f in sc.groupby('fight'):
        x = (f.time_unix_nano / 1e9).tolist(); sg += [b - a for a, b in zip(x, x[1:]) if b - a < 30]
    sd = t[(t.event == 'SPELL_DAMAGE') & (t.spell_id == SS)]; sm = t[(t.event == 'SPELL_MISSED') & (t.spell_id == SS)]
    wn = swl[~flag(swl.critical) & ~flag(swl.glancing)]
    sn = sd[~flag(sd.critical)]
    ratio = st.median(sn.amount) / st.median(wn.amount) if len(sn) and len(wn) else None
    bratio = st.median(sn.base_amount.dropna()) / st.median(wn.base_amount.dropna()) if len(sn) and len(wn) and sn.base_amount.notna().any() and wn.base_amount.notna().any() else None
    print(f'   Stormstrike casts {len(sc)} gaps {len(sg)} min {min(sg) if sg else None} p10 {q10(sg)}; damage {len(sd)} crit {int(sd.critical.sum())} '
          f'missed {dict(collections.Counter(sm.miss_type))}; SS/white median non-crit {ratio} (base {bratio}; SS {list(sn.amount.astype(int))} white median {st.median(wn.amount) if len(wn) else None})')
    # Maelstrom Weapon gains by trigger (event at the same timestamp, within 0.15 s for Stormstrike's travel)
    mw = t[t.spell_id == MW]; gains = collections.Counter(); hits = collections.Counter()
    melee = t[(t.event.isin(['SWING_DAMAGE', 'SWING_MISSED'])) | ((t.event.isin(['SPELL_DAMAGE', 'SPELL_MISSED'])) & t.spell_id.isin({SS, WFW} | WFT))]
    for e in melee.itertuples(): hits['white' if e.event.startswith('SWING') else name(e.spell_id)] += 1
    total = 0
    for tm, e in mw[mw.event.isin(['SPELL_AURA_APPLIED', 'SPELL_AURA_APPLIED_DOSE', 'SPELL_AURA_REFRESH'])].groupby('time_unix_nano'):
        k = mw_gains(e); total += k
        near = melee[(melee.time_unix_nano - tm).abs() <= 1.5e8]; ssc = sc[(tm - sc.time_unix_nano).between(0, 1.5e8)]
        src = sorted({('white' if x.startswith('SWING') else name(s)) for x, s in zip(near.event, near.spell_id)} | ({'Stormstrike'} if len(ssc) else set()))
        gains['+'.join(src) or 'unseen'] += k
    mt = bouts(pd.concat([melee.time_unix_nano, sc.time_unix_nano]) / 1e9)
    print(f'   Maelstrom Weapon gains {total} over {mt:.0f} s of melee bouts = {60 * total / mt if mt else 0:.1f}/min; melee events {dict(hits)}; gains by trigger {dict(gains)}')
    # stacks when Lightning Bolt is cast
    lbs = casts[casts.spell_id.map(lambda s: SC.get(s, {}).get('name') == LB_NAME)]
    for e in lbs.itertuples():
        start = t[(t.event == 'SPELL_CAST_START') & (t.spell_id == e.spell_id) & (t.time_unix_nano <= e.time_unix_nano)].time_unix_nano.max()
        rem = mw[(mw.time_unix_nano - e.time_unix_nano).abs() <= 2e6]
        print(f'   Lightning Bolt {e.spell_id}: cast time {(e.time_unix_nano - start) / 1e9 if pd.notna(start) else None:.2f} s, Maelstrom events at cast {list(zip(rem.event.str[11:], rem.stacks))}, ready buff up {bool(((t.spell_id == MW_READY) & (t.event == "SPELL_AURA_APPLIED") & (t.time_unix_nano <= e.time_unix_nano)).any())}')
    # Windfury
    wfw = t[t.spell_id == WFW]; wft = t[t.spell_id.isin(WFT) & t.event.isin(['SPELL_DAMAGE', 'SPELL_MISSED'])]
    wfw_procs = sorted(set(wfw.time_unix_nano)); wft_procs = sorted(set(wft.time_unix_nano))
    icd = [(b - a) / 1e9 for a, b in zip(wfw_procs, wfw_procs[1:]) if b - a < 30e9]
    both = set(wfw_procs) & set(wft_procs)
    print(f'   Windfury Weapon procs {len(wfw_procs)} ({len(wfw)} strikes, crit {int(wfw.critical.sum())}, glancing {int(flag(wfw.glancing).sum())}); min gap {min(icd) if icd else None}; '
          f'Windfury Totem extra attacks {len(wft_procs)} ({dict(collections.Counter(wft.event))}); same-timestamp both {len(both)}; '
          f'landed main-hand hits {len(swl) + len(sd)}')
    ft = t[t.spell_id.isin(FTT)]
    print(f'   Flametongue Totem procs {dict(collections.Counter(ft.spell_id))}')
    print('   casts', {f'{name(s)}:{s}': c for s, c in casts.spell_id.value_counts().items()})
    dmg = t[t.event.isin(['SWING_DAMAGE', 'SPELL_DAMAGE', 'SPELL_PERIODIC_DAMAGE', 'RANGE_DAMAGE']) & ~t.dest_guid.str.startswith('Player', na=False)]
    tot = dmg.amount.sum(); sh = dmg.groupby(dmg.spell_id.fillna(0).astype(int)).amount.sum().sort_values(ascending=False)
    print('   damage share', {name(k): f'{100 * v / tot:.0f}%' for k, v in sh.items() if v / tot >= 0.01}, f'total {int(tot)}')
