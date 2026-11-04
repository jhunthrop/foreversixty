"""Spell power implied by each spell's non-crit hits at the engine's coefficients, per character."""
import json, glob, re, statistics, collections, sys
exec(open('fit.py').read().split("hits = collections")[0].replace("for n in ('Frostbolt', 'ConeOfCold', 'FrostNova', 'IceLance'):", "for n in ('Frostbolt', 'ConeOfCold', 'FrostNova', 'Fireball', 'Scorch', 'ArcaneMissiles', 'ArcaneExplosion', 'BlastWave', 'Pyroblast'):"))
FB = open(SP + '/r19/mage/engine/sim/mage/fire_blast.go').read()
def fb(name): return re.search(r'var FireBlast%s = \[[^\]]*\][^{]*\{(.*)\}\n' % name, FB).group(1)
fid = [int(x) for x in re.findall(r'\d+', fb('SpellId'))]
fbase = [tuple(map(float, p)) for p in re.findall(r'\{([-\d.]+), ([-\d.]+)\}', fb('BaseDamage'))]
num = lambda n: [float(x) for x in re.findall(r'[\d.]+', fb(n))]
for r, sid in enumerate(fid):
    spells[sid] = dict(name='FireBlast', lvl=int(num('Level')[r]), max=int(num('MaxLevel')[r]), ppl=num('PointsPerLevel')[r], base=fbase[r], coef=num('SpellCoeff')[r])
want = set(int(x) for x in sys.argv[1:])
hits = collections.defaultdict(lambda: collections.defaultdict(list))
for f in glob.glob('dmg/*.json'):
    rid, eid, cid = map(int, f.split('/')[-1].split('.')[:3])
    if cid not in want: continue
    for e in json.load(open(f))['events']:
        if e['spell_id'] in spells and not e['is_critical']:
            hits[(rid, cid)][e['spell_id']].append(int(e['amount']))
for key in sorted(hits):
    L = levels.get(key)
    print(key, names.get(key[1]), 'L', L)
    for sid, v in sorted(hits[key].items()):
        s = spells[sid]
        if len(v) < 8 or s['coef'] == 0: continue
        top = sorted(v)[len(v) // 4:]
        med = statistics.median(top)
        bm, _ = basemean(s, L)
        print(f"   {s['name']}#{sid} lvl{s['lvl']} coef {s['coef']} n={len(v)} med={med} base={bm:.1f} SP={(med - bm) / s['coef']:.0f}  SP if x1.10={(med/1.10 - bm) / s['coef']:.0f}")
