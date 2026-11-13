"""Every class's periodic effects and their client crit flags.

Usage: python3 -I client.py <wago dir> <build tag, e.g. 70888h> > client.tsv
Reads the hotfixed SpellEffect, SpellMisc, SpellClassOptions, Talent and
SpellName tables. For every class spell, talent rank and three hops of
trigger spells with a periodic damage, leech or heal aura (3, 8, 53, 89), a
periodic trigger (23, 226) or a school damage / heal effect reached through a
periodic trigger, prints SPELL_ATTR2_CANT_CRIT (Attributes_2 0x20000000) and
SPELL_ATTR8_PERIODIC_CAN_CRIT (Attributes_8 0x200).
"""
import collections
import csv
import sys

CANT_CRIT = 0x20000000
PERIODIC_CAN_CRIT = 0x200
CHANNELED = (0x4, 0x40)  # Attributes_1 CHANNELED_1, CHANNELED_2
PERIODIC = {"3": "dmg", "8": "heal", "53": "leech", "89": "dmg%", "23": "trig", "226": "dummy-trig"}
FAMILY = {"3": "mage", "4": "warrior", "5": "warlock", "6": "priest", "7": "druid",
          "9": "hunter", "10": "paladin", "11": "shaman"}
TALENT_CLASS = {"1": "warrior", "2": "paladin", "3": "hunter", "5": "priest",
                "7": "shaman", "8": "mage", "9": "warlock", "11": "druid"}


def rows(wago, table, tag):
    with open(f"{wago}/{table}_{tag}.csv", encoding="utf-8", newline="") as f:
        return list(csv.DictReader(f))


def main(wago, tag):
    eff = collections.defaultdict(list)
    for r in rows(wago, "SpellEffect", tag):
        if r["DifficultyID"] in ("0", ""):
            eff[r["SpellID"]].append(r)
    misc = {r["SpellID"]: r for r in rows(wago, "SpellMisc", tag) if r["DifficultyID"] in ("0", "")}
    names = {r["ID"]: r["Name_lang"] for r in rows(wago, "SpellName", tag)}
    owner = {}
    for r in rows(wago, "SpellClassOptions", tag):
        if r["SpellClassSet"] in FAMILY:
            owner[r["SpellID"]] = FAMILY[r["SpellClassSet"]]
    for t in rows(wago, "Talent", tag):
        if t["ClassID"] in TALENT_CLASS:
            for i in range(9):
                v = t.get(f"SpellRank_{i}", "0")
                if v not in ("0", ""):
                    owner.setdefault(v, TALENT_CLASS[t["ClassID"]])
    parent = {}
    frontier = list(owner)
    for _ in range(3):
        nxt = []
        for s in frontier:
            for e in eff.get(s, []):
                t = e["EffectTriggerSpell"]
                if t not in ("0", "") and t not in owner:
                    owner[t] = owner[s]
                    parent[t] = s
                    nxt.append(t)
        frontier = nxt
    print("class\tspell\tname\tparent\tschool\tchannel\tcant_crit\tperiodic_can_crit\ta2\ta8\teffects")
    for s in sorted(owner, key=int):
        es = eff.get(s, [])
        per = [e for e in es if e["EffectAura"] in PERIODIC]
        via_trigger = s in parent and any(
            pe["EffectAura"] in ("23", "226") for pe in eff.get(parent[s], []))
        dmg = [e for e in es if e["Effect"] in ("2", "10", "9")] if via_trigger else []
        if not per and not dmg:
            continue
        m = misc.get(s, {})
        a1 = int(m.get("Attributes_1") or 0)
        a2 = int(m.get("Attributes_2") or 0)
        a8 = int(m.get("Attributes_8") or 0)
        desc = ";".join(
            f"E{e['Effect']}/A{e['EffectAura']} bp{e['EffectBasePointsF']} per{e['EffectAuraPeriod']} trig{e['EffectTriggerSpell']} mech{e['EffectMechanic']}"
            for e in per + dmg)
        print("\t".join([owner[s], s, names.get(s, "?"), parent.get(s, ""), m.get("SchoolMask", ""),
                         "Y" if a1 & (CHANNELED[0] | CHANNELED[1]) else "-",
                         "CANT" if a2 & CANT_CRIT else "-",
                         "PCC" if a8 & PERIODIC_CAN_CRIT else "-",
                         hex(a2), hex(a8), desc]))


if __name__ == "__main__":
    main(sys.argv[1], sys.argv[2])
