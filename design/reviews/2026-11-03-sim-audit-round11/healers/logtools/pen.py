"""Sub-20 coefficient penalty test: per character, ranks of one heal family.

For every logged character and heal spell id (pure direct heals and pure HoTs),
mean raw non-crit amount. Client: base mean (bp + growth), coefficient, learn level.
Model obs = (base + coef * p(level) * H) * M with M shared by a family's ranks.
Fits H, M by least squares under p = Era penalty and p = 1; prints residuals and gear H.
"""
import collections, csv, glob, json, math, sys
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
RAW = SP + '/r9/casters/raw/'
TL = SP + '/r10/healers/fl/tl'

def load(t, key='SpellID'):
    out = collections.defaultdict(list)
    with open(RAW + t + '.csv', newline='') as f:
        for r in csv.DictReader(f):
            out[r[key]].append(r)
    return out
eff = load('SpellEffect'); lev = load('SpellLevels'); names = {}
with open(RAW + 'SpellName.csv', newline='') as f:
    for r in csv.DictReader(f): names[r['ID']] = r['Name_lang']

FAM = {
 'Healing Wave': 'd', 'Lesser Healing Wave': 'd', 'Holy Light': 'd', 'Flash of Light': 'd',
 'Lesser Heal': 'd', 'Heal': 'd', 'Flash Heal': 'd', 'Greater Heal': 'd', 'Healing Touch': 'd',
 'Renew': 'h', 'Rejuvenation': 'h',
}

def client(sid, L):
    s = str(sid); kind = FAM.get(names.get(s))
    if not kind: return None
    es = eff.get(s, [])
    e = None
    for r in es:
        if kind == 'd' and r['Effect'] == '10': e = r
        if kind == 'h' and r['EffectAura'] == '8': e = r
    if not e: return None
    lv = lev.get(s, [{}])[0]
    spl = int(lv.get('SpellLevel') or 0); mx = int(lv.get('MaxLevel') or 0); base = int(lv.get('BaseLevel') or 0)
    per = float(e['EffectRealPointsPerLevel'] or 0)
    lvl = min(L, mx) if mx else L
    bp = float(e['EffectBasePointsF']) + per * max(0, lvl - max(base, spl))
    return dict(bp=bp, coef=float(e['EffectBonusCoefficient']), learn=spl, kind=kind, name=names[s])

def era(level):
    return 1 if level >= 20 or level <= 0 else 1 - (20 - level) * 0.0375

roster = {}
for f in glob.glob(SP + '/fl/rep/*.roster.json'):
    rid = int(f.split('/')[-1].split('.')[0])
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            try:
                tabs = s['talents']['talent_groups'][str(s['talents']['active_group'])]['tabs']
            except Exception:
                continue
            pts = {v['name']: v['points'] for v in tabs.values()}
            h = 0
            for g in (s.get('gear') or {}).values():
                st = ((g.get('resolved_bisbeard') or {}).get('stats') or {})
                h += st.get('healingPower', 0) + st.get('spellPower', 0)
            roster[(rid, c['character_id'])] = (s['player']['class'], sum(pts.values()) + 9, pts, h, c['name'])

def raw_heals(events):
    over = collections.defaultdict(float)
    for e in events:
        if e['event_type'] == 'overheal':
            over[(e['timestamp_ms'], e['target_id'], e['spell_id'])] += float(e['amount'])
    for e in events:
        if e['event_type'] == 'heal':
            yield e, float(e['amount']) + over.pop((e['timestamp_ms'], e['target_id'], e['spell_id']), 0.0)

agg = collections.defaultdict(list)
for f in glob.glob(TL + '/*.json'):
    rid, enc, cid = map(int, f.split('/')[-1].split('.')[:3])
    info = roster.get((rid, cid))
    if not info: continue
    try: d = json.load(open(f))
    except Exception: continue
    for e, raw in raw_heals(d.get('events') or []):
        if e['source_id'] != cid or e['is_critical']: continue
        agg[(cid, e['spell_id'])].append(raw)

chars = {}
for (rid, cid), info in roster.items(): chars[cid] = info
mn = int(sys.argv[1]) if len(sys.argv) > 1 else 8
byfam = collections.defaultdict(list)
for (cid, sid), xs in agg.items():
    if len(xs) < mn or cid not in chars: continue
    info = chars[cid]; c = client(sid, info[1])
    if not c: continue
    byfam[(cid, c['name'])].append((sid, c, sum(xs) / len(xs), len(xs)))

def fit(rows, pen):
    # obs = M*bp + (M*H)*coef*p : linear in a=M, b=M*H
    X = [(r[1]['bp'], r[1]['coef'] * pen(r[1]['learn'])) for r in rows]; y = [r[2] for r in rows]
    sxx = sum(a*a for a, _ in X); sxy = sum(a*b for a, b in X); syy = sum(b*b for _, b in X)
    tx = sum(a*v for (a, _), v in zip(X, y)); ty = sum(b*v for (_, b), v in zip(X, y))
    det = sxx*syy - sxy*sxy
    if abs(det) < 1e-9: return None
    A = (tx*syy - ty*sxy) / det; B = (sxx*ty - sxy*tx) / det
    res = [v - (A*a + B*b) for (a, b), v in zip(X, y)]
    return A, (B / A if A else float('nan')), res

for key, rows in sorted(byfam.items(), key=lambda kv: (chars[kv[0][0]][0], kv[0])):
    if len(rows) < 2: continue
    if not any(r[1]['learn'] < 20 for r in rows) or not any(r[1]['learn'] >= 20 for r in rows):
        if len(rows) < 3: continue
    info = chars[key[0]]
    print(f"== {info[0]} {info[4]} ({key[0]}) L~{info[1]} gearH {info[3]} {key[1]}")
    for sid, c, m, n in sorted(rows, key=lambda r: r[1]['learn']):
        print(f"   {sid} learn {c['learn']:2d} bp {c['bp']:7.1f} coef {c['coef']:.3f} obs {m:7.1f} n {n}")
    for lab, pen in (('era', era), ('none', lambda l: 1.0)):
        r = fit(rows, pen)
        if r: print(f"   {lab:4}: M {r[0]:.3f} H {r[1]:7.1f} resid {[round(x,1) for x in r[2]]}")
