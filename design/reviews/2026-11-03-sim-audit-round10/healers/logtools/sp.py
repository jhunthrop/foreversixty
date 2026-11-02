"""Dump client rows for spell ids or names: sp.py <id|name-substring> ..."""
import csv
import sys
from collections import defaultdict

RAW = "/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/r9/casters/raw/"


def load(t, key="SpellID"):
    out = defaultdict(list)
    with open(RAW + t + ".csv", newline="") as f:
        for r in csv.DictReader(f):
            out[r[key]].append(r)
    return out


names = {}
with open(RAW + "SpellName.csv", newline="") as f:
    for r in csv.DictReader(f):
        names[r["ID"]] = r["Name_lang"]
eff = load("SpellEffect")
cls = load("SpellClassOptions")
misc = load("SpellMisc")
cat = load("SpellCategories")
auo = load("SpellAuraOptions")

ids = []
for a in sys.argv[1:]:
    if a.isdigit():
        ids.append(a)
    else:
        ids += [k for k, v in names.items() if a.lower() == v.lower()]

for i in ids:
    c = cls.get(i, [{}])[0]
    m = misc.get(i, [{}])[0]
    ca = cat.get(i, [{}])[0]
    ao = auo.get(i, [{}])[0]
    print(f"== {i} {names.get(i)}  mask={[hex(int(c.get('SpellClassMask_%d' % k, 0) or 0)) for k in range(4)]} set={c.get('SpellClassSet')} "
          f"school={m.get('SchoolMask')} attr8={hex(int(m.get('Attributes_8', 0) or 0))} def={ca.get('DefenseType')} "
          f"proc={ao.get('ProcChance')} ch={ao.get('ProcCharges')} ptm={hex(int(ao.get('ProcTypeMask_0', 0) or 0))}")
    for e in sorted(eff.get(i, []), key=lambda r: int(r["EffectIndex"])):
        em = [hex(int(e["EffectSpellClassMask_%d" % k])) for k in range(4)]
        print(f"   e{e['EffectIndex']} eff={e['Effect']} aura={e['EffectAura']} bp={e['EffectBasePointsF']} misc={e['EffectMiscValue_0']}/{e['EffectMiscValue_1']} "
              f"per={e['EffectAuraPeriod']} coef={e['EffectBonusCoefficient']} trig={e['EffectTriggerSpell']} mask={em} tgt={e['ImplicitTarget_0']}/{e['ImplicitTarget_1']} rid={e['ID']}")
