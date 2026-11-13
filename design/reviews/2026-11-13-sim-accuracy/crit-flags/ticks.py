"""Count critical over all periodic and direct damage and healing events per
class and spell, over every fight of our backfilled reports.

Usage: python3 -I ticks.py <dl dir> > ticks_ours.tsv
Players only (the unit's class from report.json); a report without
report.json is skipped for class but its spells still count under "?".
"""
import collections
import glob
import json
import os
import sys

import pyarrow.parquet as pq

EVENTS = {"SPELL_PERIODIC_DAMAGE", "SPELL_PERIODIC_HEAL", "SPELL_DAMAGE", "SPELL_HEAL", "RANGE_DAMAGE"}
COLS = ["event", "source_guid", "spell_id", "spell_name", "critical", "amount", "adv_level", "adv_info_guid", "adv_present"]


def classes(dl):
    out = {}
    for rj in glob.glob(f"{dl}/*/report.json"):
        for u in json.load(open(rj))["units"]:
            if u.get("class"):
                out[u["guid"]] = u["class"]
    return out


def main(dl):
    cls = classes(dl)
    n = collections.Counter()
    c = collections.Counter()
    units = collections.defaultdict(set)
    names = {}
    for f in glob.glob(f"{dl}/*/f*.parquet"):
        try:
            t = pq.read_table(f, columns=COLS).to_pandas()
        except Exception as exc:
            print("bad", f, exc, file=sys.stderr)
            continue
        t = t[t.event.isin(EVENTS) & t.source_guid.str.startswith("Player-")]
        for ev, src, sid, sname, crit in zip(t.event, t.source_guid, t.spell_id, t.spell_name, t.critical):
            k = (cls.get(src, "?"), int(sid), ev)
            n[k] += 1
            c[k] += bool(crit)
            units[k].add(src)
            names[int(sid)] = sname
    print("class\tspell\tname\tevent\tevents\tcrits\tplayers")
    for k in sorted(n, key=lambda k: (k[0], k[1], k[2])):
        print(f"{k[0]}\t{k[1]}\t{names[k[1]]}\t{k[2]}\t{n[k]}\t{c[k]}\t{len(units[k])}")


if __name__ == "__main__":
    main(sys.argv[1])
