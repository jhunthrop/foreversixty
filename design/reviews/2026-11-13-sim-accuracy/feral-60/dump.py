"""Dump SpellEffect rows (non-zero columns) for spells whose name matches.
Usage: python3 -I dump.py <wago dir> <tag> <name>..."""
import csv, sys
W, T = sys.argv[1:3]; want = set(sys.argv[3:])
def rows(t):
    with open(f"{W}/{t}_{T}.csv", encoding="utf-8", newline="") as f:
        return list(csv.DictReader(f))
names = {r["ID"]: r["Name_lang"] for r in rows("SpellName")}
cls = {r["SpellID"]: r["SpellClassSet"] for r in rows("SpellClassOptions")}
keep = ["EffectIndex","Effect","EffectAura","EffectBasePointsF","EffectRealPointsPerLevel","EffectPointsPerResource","BonusCoefficientFromAP","EffectBonusCoefficient","Variance","EffectChainAmplitude","EffectMiscValue_0","EffectTriggerSpell","EffectAuraPeriod","EffectSpellClassMask_0","EffectSpellClassMask_1","EffectSpellClassMask_2","EffectSpellClassMask_3","EffectAmplitude","ResourceCoefficient","Coefficient"]
for r in sorted(rows("SpellEffect"), key=lambda r: (int(r["SpellID"]), int(r["EffectIndex"]))):
    n = names.get(r["SpellID"], "?")
    if n in want and r["DifficultyID"] in ("0",""):
        print(r["SpellID"], n, "cls="+cls.get(r["SpellID"],"-"), " ".join(f"{k}={r[k]}" for k in keep if r[k] not in ("0","","0.0")))
