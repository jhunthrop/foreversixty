"""Per character: spell power implied by Frostbolt, Cone of Cold and Frost Nova (client coefficients),
then Ice Lance's coefficient from its unfrozen and frozen non-crit hits at that spell power."""
import json, glob, re, statistics, collections, sys
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
C = open(SP + '/r19/mage/engine/sim/mage/constants_auto_gen.go').read()
def arr(name):
    m = re.search(r'var %s = \[[^\]]*\][^{]*\{(.*)\}\n' % name, C)
    return m.group(1)
def nums(name): return [float(x) for x in re.findall(r'-?[\d.]+', arr(name))]
def pairs(name): return [tuple(map(float, p)) for p in re.findall(r'\{([-\d.]+), ([-\d.]+)\}', arr(name))]
spells = {}
for n in ('Frostbolt', 'ConeOfCold', 'FrostNova', 'IceLance'):
    ids = [int(x) for x in nums(n + 'SpellId')]
    for r, sid in enumerate(ids):
        spells[sid] = dict(name=n, lvl=int(nums(n + 'Level')[r]), max=int(nums(n + 'MaxLevel')[r]),
                           ppl=nums(n + 'PointsPerLevel')[r], base=pairs(n + 'BaseDamage')[r], coef=nums(n + 'SpellCoeff')[r])
def basemean(s, L):
    grow = s['ppl'] * (min(L, s['max']) - s['lvl'])
    return (s['base'][0] + s['base'][1]) / 2 + grow, s['base'][1] + grow
levels = {}
names = {}
summary = []
for f in glob.glob(SP + '/fl/rep/*.roster.json'):
    rid = int(f.split('/')[-1].split('.')[0])
    for c in json.load(open(f))['roster']:
        sn = c.get('snapshots') or []
        if c.get('class') == 'Mage' and sn:
            tabs = sn[0]['talents']['talent_groups'][str(sn[0]['talents']['active_group'])]['tabs']
            levels[(rid, c['character_id'])] = sum(v['points'] for v in tabs.values()) + 9
            names[c['character_id']] = c['name']
hits = collections.defaultdict(lambda: collections.defaultdict(list))
for f in glob.glob('dmg/*.json'):
    rid, eid, cid = map(int, f.split('/')[-1].split('.')[:3])
    for e in json.load(open(f))['events']:
        sid = e['spell_id']
        if sid in spells and not e['is_critical']:
            hits[(rid, cid)][sid].append(int(e['amount']))
for key in sorted(hits):
    L = levels.get(key)
    if L is None: continue
    sps = []
    out = []
    for sid, v in sorted(hits[key].items()):
        s = spells[sid]
        if s['name'] == 'IceLance': continue
        if len(v) < 5: continue
        top = sorted(v)[len(v) // 4:]  # drop the low quarter: partial resists, glancing AoE edges
        med = statistics.median(top)
        bm, bmax = basemean(s, L)
        sp = (med - bm) / s['coef']
        sps.append((len(v), sp)); out.append(f"{s['name']}#{sid} n={len(v)} med={med} base={bm:.1f} SP={sp:.0f}")
    if not sps: continue
    SPest = sum(n * x for n, x in sps) / sum(n for n, _ in sps)
    for sid, v in hits[key].items():
        s = spells[sid]
        if s['name'] != 'IceLance' or len(v) < 3: continue
        bm, bmax = basemean(s, L)
        unf = [x for x in v if x < 2.2 * bmax + 0.6 * SPest]
        fro = [x / 4 for x in v if x >= 2.2 * bmax + 0.6 * SPest]
        def c(xs):
            xs = sorted(xs)[len(xs) // 4:] if len(xs) >= 4 else xs
            return (statistics.median(xs) - bm) / SPest if xs and SPest > 5 else None
        summary.append((names.get(key[1]), key[0], sid, len(unf), c(unf), len(fro), c(fro), round(SPest)))
        out.append(f"IceLance#{sid} unfrozen n={len(unf)} med={statistics.median(unf) if unf else '-'} frozen/4 n={len(fro)} med={statistics.median(fro) if fro else '-'} base={bm:.1f} -> coef unf={c(unf)} fro={c(fro)}")
    print(key, 'L', L, 'SPest', round(SPest), '|', ' ; '.join(out))

print()
print('player\treport\trank id\tSP\tunfrozen n\tcoef\tfrozen n\tcoef')
allc=[]
for nm, rid, sid, nu, cu, nf, cf, sp in sorted(summary, key=lambda r: (r[0] or '', r[1])):
    f = lambda x: '-' if x is None else '%.3f' % x
    print(f'{nm}\t{rid}\t{sid}\t{sp}\t{nu}\t{f(cu)}\t{nf}\t{f(cf)}')
    allc += [(nu, cu), (nf, cf)]
ok=[(n, c) for n, c in allc if c is not None and n >= 3]
print('hit-weighted mean', round(sum(n*c for n, c in ok)/sum(n for n, _ in ok), 4), 'over', sum(n for n,_ in ok), 'non-crit hits')
