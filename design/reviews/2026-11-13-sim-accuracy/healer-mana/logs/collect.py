"""Level-60 raid healers' mana and the damage they heal, from Classic Era Warcraft Logs (v1).

Sample: for every Molten Core boss and Onyxia, the hps rankings of a partition (default 8,
Hardcore Fresh: Phase 1 gear and no world buffs, the nearest public analogue of a Forever
Phase 1 raid), pages 2, 5 and 10, at most one kill per report, ordered by a hash of report
and fight, PER_BOSS kills a boss, raid size 30 or more.

Per kill: its length; per healer (Warcraft Logs' summary roles) the healing done and
overhealed, the mana gained by source (resources-gains, logged energizes only: potions,
runes, Mana Tide, Mana Spring, Judgement of Wisdom; Spirit, mp5 and Innervate's regen are not
logged), and the mana bar at each of its casts (classResources; on the 2026-11-13 sample every
event reads 0 of 0, so the bar is not in these logs and analyze.py leaves it out); the main tank's (the tank who took the most) damage taken a second and
its spikiness (one-second bins: the largest five-second sum over the mean five-second sum);
every other player's damage taken a second.

Usage: WCL_KEY=... WCL_CACHE=<dir> python3 -I collect.py <this dir> <out.json> [per boss] [partition]
"""
import hashlib
import json
import sys

sys.path.insert(0, sys.argv[1])
import wcl  # noqa: E402

OUT = sys.argv[2]
PER_BOSS = int(sys.argv[3]) if len(sys.argv) > 3 else 3
PARTITION = int(sys.argv[4]) if len(sys.argv) > 4 else 8
PAGES = (2, 5, 10)
MIN_RAID_SIZE = 30
BOSSES = [(50663, "Lucifron"), (50664, "Magmadar"), (50665, "Gehennas"), (50666, "Garr"),
          (50667, "Shazzrah"), (50668, "Baron Geddon"), (50669, "Sulfuron Harbinger"),
          (50670, "Golemagg the Incinerator"), (50671, "Majordomo Executus"), (50672, "Ragnaros"),
          (51084, "Onyxia")]
MANA = 0


def sample(encounter):
    seen = {}
    for page in PAGES:
        try:
            body = wcl.get(f"rankings/encounter/{encounter}", metric="hps", partition=PARTITION, page=page)
        except RuntimeError as err:
            print(f"skip page {page} of {encounter}: {err}", file=sys.stderr)
            continue
        for r in body.get("rankings", []):
            if r.get("size", 0) >= MIN_RAID_SIZE and r["reportID"] not in seen:
                seen[r["reportID"]] = r["fightID"]
    order = sorted(seen.items(), key=lambda kv: hashlib.sha1(f"{kv[0]}:{kv[1]}".encode()).hexdigest())
    return order[:PER_BOSS]


def all_events(view, code, start, end, **params):
    out, cursor = [], start
    while cursor is not None:
        body = wcl.get(f"report/events/{view}/{code}", start=cursor, end=end, **params)
        out.extend(body.get("events", []))
        cursor = body.get("nextPageTimestamp")
    return out


def mana_track(events, actor):
    track = []
    for ev in events:
        if ev.get("sourceID") != actor:
            continue
        for res in ev.get("classResources") or []:
            if res.get("type") == MANA and res.get("max"):
                track.append((ev["timestamp"], res["amount"] / res["max"]))
    return track


def spikiness(events, start, end):
    secs = max(1, (end - start) // 1000 + 1)
    bins = [0.0] * secs
    for ev in events:
        if ev.get("type") == "damage":
            bins[min(secs - 1, (ev["timestamp"] - start) // 1000)] += ev.get("amount", 0) + ev.get("absorbed", 0)
    windows = [sum(bins[i:i + 5]) for i in range(max(1, secs - 4))]
    mean = sum(windows) / len(windows)
    return sum(bins), (max(windows) / mean if mean else 0.0)


def kill(code, fid, boss):
    fights = wcl.get(f"report/fights/{code}", translate="true")
    fight = next(f for f in fights["fights"] if f["id"] == fid)
    s, e = fight["start_time"], fight["end_time"]
    secs = (e - s) / 1000
    roles = wcl.get(f"report/tables/summary/{code}", start=s, end=e)["playerDetails"]
    taken = {x["id"]: x["total"] for x in wcl.get(f"report/tables/damage-taken/{code}", start=s, end=e)["entries"]}
    heal = {x["id"]: x for x in wcl.get(f"report/tables/healing/{code}", start=s, end=e)["entries"]}
    tanks = [p["id"] for p in roles.get("tanks", [])]
    mt = max(tanks, key=lambda i: taken.get(i, 0)) if tanks else None
    others = [p["id"] for k in ("dps", "healers") for p in roles.get(k, [])]
    row = {"report": code, "fight": fid, "boss": boss, "seconds": secs,
           "players": len(others) + len(tanks), "tanks": len(tanks),
           "member_dtps": sum(taken.get(i, 0) for i in others) / max(1, len(others)) / secs,
           "healers": []}
    if mt is not None:
        total, spike = spikiness(all_events("damage-taken", code, s, e, sourceid=mt), s, e)
        row["mt_dtps"], row["mt_spike5"] = total / secs, spike
    casts = all_events("casts", code, s, e)
    for p in roles.get("healers", []):
        h = heal.get(p["id"], {})
        gains = wcl.get(f"report/tables/resources-gains/{code}", start=s, end=e, abilityid=100, sourceid=p["id"])
        track = mana_track(casts, p["id"])
        row["healers"].append({
            "class": p["type"], "spec": (p.get("specs") or [""])[0],
            "healing": h.get("total", 0), "overheal": h.get("overheal", 0),
            "gains": {str(g["guid"]): [g["name"], g["gains"]] for g in gains.get("resources", [])},
            "mana_start": track[0][1] if track else None,
            "mana_min": min(v for _, v in track) if track else None,
            "mana_end": track[-1][1] if track else None,
        })
    return row


def main():
    rows = []
    for enc, name in BOSSES:
        for code, fid in sample(enc):
            try:
                rows.append(kill(code, fid, name))
                print(f"{name} {code} {fid} {rows[-1]['seconds']:.0f}s", file=sys.stderr)
            except (RuntimeError, StopIteration, KeyError) as err:
                print(f"skip {name} {code} {fid}: {err}", file=sys.stderr)
    with open(OUT, "w") as f:
        json.dump(rows, f, indent=1)


if __name__ == "__main__":
    main()
