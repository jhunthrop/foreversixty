"""own60.py <reports root> <min level>: every damage dealer at or above <min level> in our uploaded reports
(fetched read-only with ../../2026-11-13-sim-accuracy/enhancement-70938/logs/fetch_r2.py). Level is the player's own
advanced-log level (adv_level on events whose adv_info_guid is the player). Spec from signature abilities in the
fight summary. Prints per spec: players, fights, and DPS (median of per-player medians, fight-weighted median).
Player names never printed."""
import sys, glob, json, os, statistics as st, collections
import pandas as pd
root, minlv = sys.argv[1], int(sys.argv[2])
SIG = {  # spell name fragments (any locale seen: en, pt) -> spec, checked in order per class
 'Shaman': [('Stormstrike', 'Enhancement'), ('Ataque Tempestuoso', 'Enhancement'), ('', 'Elemental')],
 'Warrior': [('Bloodthirst', 'Fury'), ('Sede de Sangue', 'Fury'), ('Mortal Strike', 'Arms'), ('Golpe Mortal', 'Arms'), ('', 'Warrior other')],
 'Druid': [('Shred', 'Feral'), ('Rip', 'Feral'), ('Rake', 'Feral'), ('Retalhar', 'Feral'), ('Rasgar', 'Feral'), ('Estraçalhar', 'Feral'), ('Starfire', 'Balance'), ('Moonfire', 'Balance'), ('', 'Druid other')],
 'Rogue': [('', 'Rogue')], 'Hunter': [('', 'Hunter')], 'Paladin': [('', 'Retribution')],
 'Mage': [('', 'Mage')], 'Warlock': [('', 'Warlock')], 'Priest': [('', 'Priest')],
}
def spec_of(cls, names):
    for frag, spec in SIG.get(cls, [('', cls or '?')]):
        if frag == '' or any(frag in n for n in names):
            return spec
rows = []
for summ in glob.glob(f'{root}/*/fights/*/summary.json'):
    d = json.load(open(summ)); ev = summ.replace('summary.json', 'events.parquet')
    if not os.path.exists(ev) or d.get('duration_ms', 0) < 20000: continue
    df = pd.read_parquet(ev, columns=['adv_info_guid', 'adv_level'])
    lv = df[df.adv_level > 0].groupby('adv_info_guid').adv_level.max()
    dd = {p['guid']: p for p in d.get('damage_done') or []}
    for r in d.get('roster') or []:
        g = r['guid']
        if r.get('role') != 'dps' or g not in lv or lv[g] < minlv or g not in dd: continue
        names = [a['name'] for a in dd[g].get('abilities', [])]
        rows.append(dict(report=summ.split('/')[-4], fight=d['fight_index'], player=g, level=int(lv[g]), cls=r.get('class') or dd[g].get('class'),
                         spec=spec_of(r.get('class') or dd[g].get('class'), names), dps=r['dps'], secs=d['duration_ms'] / 1000))
t = pd.DataFrame(rows)
if t.empty: print('no rows'); sys.exit()
print(f'rows {len(t)}, reports {t.report.nunique()}, levels {t.level.min()} to {t.level.max()}')
out = []
for spec, g in t.groupby('spec'):
    per = g.groupby('player').dps.median()
    out.append((st.median(per), spec, g.player.nunique(), len(g), st.median(g.dps), g.level.min(), g.level.max()))
for m, spec, np_, nf, fm, lo, hi in sorted(out, reverse=True):
    print(f'{spec:14s} players {np_:2d} fights {nf:4d} median-of-players {m:7.1f} fight-median {fm:7.1f} levels {lo}-{hi}')
