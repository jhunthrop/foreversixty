"""Join logged per-spell crit counts with the client's crit flags.

Usage: python3 -I join.py <wago dir> <tag> <ticks.tsv>... > joined.tsv
Sums the given tick files, adds the English name, CANT_CRIT and
PERIODIC_CAN_CRIT of the logged spell id.
"""
import collections
import csv
import sys


def main(wago, tag, files):
    misc = {r["SpellID"]: r for r in csv.DictReader(open(f"{wago}/SpellMisc_{tag}.csv")) if r["DifficultyID"] in ("0", "")}
    names = {r["ID"]: r["Name_lang"] for r in csv.DictReader(open(f"{wago}/SpellName_{tag}.csv"))}
    n = collections.Counter()
    c = collections.Counter()
    p = collections.Counter()
    for f in files:
        for r in csv.DictReader(open(f), delimiter="\t"):
            k = (r["class"], r["spell"], r["event"])
            n[k] += int(r["events"])
            c[k] += int(r["crits"])
            p[k] += int(r["players"])
    print("class\tspell\tname\tevent\tevents\tcrits\tpct\tplayers\tcant\tpcc")
    for k in sorted(n):
        m = misc.get(k[1], {})
        a2 = int(m.get("Attributes_2") or 0)
        a8 = int(m.get("Attributes_8") or 0)
        print(f"{k[0]}\t{k[1]}\t{names.get(k[1], '?')}\t{k[2]}\t{n[k]}\t{c[k]}\t{100*c[k]/n[k]:.1f}\t{p[k]}\t"
              f"{'CANT' if a2 & 0x20000000 else '-'}\t{'PCC' if a8 & 0x200 else '-'}")


if __name__ == "__main__":
    main(sys.argv[1], sys.argv[2], sys.argv[3:])
