"""Diff druid-owned spell rows between two hotfixed wago builds.
Usage: python3 -I diff.py <wago dir> <old tag> <new tag>"""
import collections, csv, sys
W, A, B = sys.argv[1:4]
def rows(t, tag):
    with open(f"{W}/{t}_{tag}.csv", encoding="utf-8", newline="") as f:
        return list(csv.DictReader(f))
def druid_ids(tag):
    own = {r["SpellID"] for r in rows("SpellClassOptions", tag) if r["SpellClassSet"] == "7"}
    for t in rows("Talent", tag):
        if t["ClassID"] == "11":
            own |= {t[f"SpellRank_{i}"] for i in range(9) if t[f"SpellRank_{i}"] not in ("0", "")}
    names = {r["ID"]: r["Name_lang"] for r in rows("SpellName", tag)}
    # trait spells: TraitDefinition SpellIDs whose name is a druid talent is hard; take all trait spells
    tds = {r["SpellID"] for r in rows("TraitDefinition", tag) if r.get("SpellID", "0") not in ("0", "")}
    eff = collections.defaultdict(list)
    for r in rows("SpellEffect", tag):
        eff[r["SpellID"]].append(r)
    own |= tds
    front = list(own)
    for _ in range(2):
        nxt = []
        for s in front:
            for e in eff.get(s, []):
                t = e["EffectTriggerSpell"]
                if t not in ("0", "") and t not in own:
                    own.add(t); nxt.append(t)
        front = nxt
    return own, names
ids, names = druid_ids(B)
ida, _ = druid_ids(A)
ids |= ida
for t in ["SpellEffect", "SpellMisc", "SpellPower", "SpellCooldowns", "SpellLevels", "SpellCategories", "SpellAuraOptions", "SpellClassOptions", "Spell", "SpellEquippedItems"]:
    key = "ID"
    ra = {r[key]: r for r in rows(t, A)}
    rb = {r[key]: r for r in rows(t, B)}
    sid = "ID" if t == "Spell" else "SpellID"
    for k in sorted(set(ra) | set(rb), key=lambda x: int(x) if x.lstrip('-').isdigit() else 0):
        a, b = ra.get(k), rb.get(k)
        s = (a or b)[sid]
        if s not in ids or a == b:
            continue
        if a is None or b is None:
            print(t, s, names.get(s, "?"), "ADDED" if a is None else "REMOVED", {kk: v for kk, v in (a or b).items() if v not in ("0", "", "0.0")})
            continue
        d = {c: (a[c], b[c]) for c in a if a[c] != b.get(c)}
        print(t, s, names.get(s, "?"), d)
