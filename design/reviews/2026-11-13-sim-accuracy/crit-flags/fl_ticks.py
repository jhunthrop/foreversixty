"""foreverlogs.gg: crit over all events for chosen spell ids, every public
report. Uses fl.py (a cached GET helper beside it in the working dir).

Usage: python3 fl_ticks.py <spell id>... > fl_ticks.tsv  (run from the cache dir)
"""
import collections
import sys
import urllib.parse

import fl


def events(rid, ids, stream):
    out, cur = [], None
    while True:
        p = f"/api/reports/{rid}/meta_events?limit=2000&streams={stream}&spellIds=" + ",".join(ids)
        if cur:
            p += "&cursor=" + urllib.parse.quote(str(cur))
        d = fl.get(p)
        out += d.get("events", [])
        if not d.get("hasMore"):
            return out
        cur = d["nextCursor"]


def main(ids):
    reports = [r["id"] for r in fl.get("/api/reports/public?limit=100")["reports"]]
    n, c, who, names = collections.Counter(), collections.Counter(), collections.defaultdict(set), {}
    for rid in reports:
        for stream in ("damage", "healing"):
            try:
                evs = events(rid, ids, stream)
            except Exception as exc:
                print("fail", rid, stream, exc, file=sys.stderr)
                continue
            for e in evs:
                if not e.get("spell") or e["type"] not in ("damage", "heal"):
                    continue
                k = ((e.get("source") or {}).get("class") or "?", str(e["spell"]["id"]), e["type"])
                n[k] += 1
                c[k] += bool(e.get("isCritical"))
                who[k].add((rid, (e.get("source") or {}).get("id")))
                names[k[1]] = e["spell"]["name"]
    print(f"reports\t{len(reports)}")
    print("class\tspell\tname\tevent\tevents\tcrits\tplayers")
    for k in sorted(n):
        print(f"{k[0]}\t{k[1]}\t{names[k[1]]}\t{k[2]}\t{n[k]}\t{c[k]}\t{len(who[k])}")


if __name__ == "__main__":
    main(sys.argv[1:])
