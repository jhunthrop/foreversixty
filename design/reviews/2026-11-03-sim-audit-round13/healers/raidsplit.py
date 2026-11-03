"""Level-60 tank share of a raid's damage taken, from Classic Era Warcraft Logs (v1).

Rule: for each Molten Core, Onyxia and Blackwing Lair boss, take the same number of logged
40-man kills (raid size 30 or more), sampled from the S0 dps rankings at fixed pages and
ordered by a hash of report and fight. A kill's tank share is the damage taken by the
players Warcraft Logs' summary lists as that fight's tanks over the damage taken by every
player (pets left out). The band value is the median over every sampled kill; with the same
number of kills per boss each boss weighs the same. Trash is out: the level-60 profile is a
boss fight.

Usage: WCL_KEY=... python3 -I raidsplit.py <this dir> [kills per boss]
"""
import hashlib
import json
import statistics
import sys

sys.path.insert(0, sys.argv[1])
import wcl  # noqa: E402

PER_BOSS = int(sys.argv[2]) if len(sys.argv) > 2 else 10
PAGES = (2, 5, 10)
MIN_RAID_SIZE = 30
PLAYER_CLASSES = {"Warrior", "Paladin", "Hunter", "Rogue", "Priest", "Shaman", "Mage", "Warlock", "Druid"}
BOSSES = {
    "Molten Core": [(50663, "Lucifron"), (50664, "Magmadar"), (50665, "Gehennas"), (50666, "Garr"),
                    (50667, "Shazzrah"), (50668, "Baron Geddon"), (50669, "Sulfuron Harbinger"),
                    (50670, "Golemagg the Incinerator"), (50671, "Majordomo Executus"), (50672, "Ragnaros")],
    "Onyxia's Lair": [(51084, "Onyxia")],
    "Blackwing Lair": [(50610, "Razorgore the Untamed"), (50611, "Vaelastrasz the Corrupt"),
                       (50612, "Broodlord Lashlayer"), (50613, "Firemaw"), (50614, "Ebonroc"),
                       (50615, "Flamegor"), (50616, "Chromaggus"), (50617, "Nefarian")],
}


def sample(encounter):
    seen = {}
    for page in PAGES:
        try:
            body = wcl.get(f"rankings/encounter/{encounter}", metric="dps", partition=1, page=page)
        except RuntimeError as err:
            print(f"skip page: {err}", file=sys.stderr)
            continue
        for r in body["rankings"]:
            if r["fightID"] > 0 and r["size"] >= MIN_RAID_SIZE:
                seen[(r["reportID"], r["fightID"])] = r["startTime"]
    order = sorted(seen, key=lambda k: hashlib.sha1(f"{k[0]}:{k[1]}".encode()).hexdigest())
    return order


def fight_share(code, fight_id, encounter):
    fights = wcl.get("report/fights/" + code)
    match = [f for f in fights["fights"] if f["id"] == fight_id and f.get("boss") == encounter and f.get("kill")]
    if not match:
        return None
    f = match[0]
    summary = wcl.get("report/tables/summary/" + code, start=f["start_time"], end=f["end_time"])
    tanks = {p["id"] for p in summary["playerDetails"].get("tanks", [])}
    if not tanks:
        return None
    taken = wcl.get("report/tables/damage-taken/" + code, start=f["start_time"], end=f["end_time"])
    players = [e for e in taken["entries"] if e["type"] in PLAYER_CLASSES]
    total = sum(e["total"] for e in players)
    if total <= 0:
        return None
    tank = sum(e["total"] for e in players if e["id"] in tanks)
    return {"report": code, "fight": fight_id, "start": fights["start"], "size": f["size"],
            "seconds": round((f["end_time"] - f["start_time"]) / 1000), "tanks": len(tanks),
            "tank": tank, "total": total, "share": tank / total}


def quartiles(xs):
    q = statistics.quantiles(xs, n=4, method="inclusive")
    return q[0], statistics.median(xs), q[2]


def main():
    rows = []
    for raid, bosses in BOSSES.items():
        for encounter, name in bosses:
            got = []
            for code, fight_id in sample(encounter):
                if len(got) == PER_BOSS:
                    break
                if any(g["report"] == code for g in got):
                    continue  # one kill per report, so one raid does not weigh twice
                try:
                    share = fight_share(code, fight_id, encounter)
                except RuntimeError as err:
                    print(f"skip fight: {err}", file=sys.stderr)
                    share = None
                if share:
                    got.append({**share, "raid": raid, "boss": name, "encounter": encounter})
            rows.extend(got)
            if got:
                q25, med, q75 = quartiles([g["share"] for g in got]) if len(got) > 1 else (got[0]["share"],) * 3
                print(f"{raid:15} {name:26} kills {len(got):2} median {med:.3f} ({q25:.3f} to {q75:.3f})", flush=True)
            else:
                print(f"{raid:15} {name:26} kills  0", flush=True)
    json.dump(rows, open(sys.argv[1] + "/raidsplit-rows.json", "w"), indent=1)
    shares = [r["share"] for r in rows]
    q25, med, q75 = quartiles(shares)
    print(f"ALL kills {len(shares)} median {med:.4f} q25 {q25:.4f} q75 {q75:.4f}")
    for raid in BOSSES:
        xs = [r["share"] for r in rows if r["raid"] == raid]
        if len(xs) > 1:
            a, m, b = quartiles(xs)
            print(f"  {raid}: kills {len(xs)} median {m:.3f} ({a:.3f} to {b:.3f})")
    boss_medians = [statistics.median([r["share"] for r in rows if r["boss"] == n]) for bs in BOSSES.values() for _, n in bs if any(r["boss"] == n for r in rows)]
    print(f"  median of boss medians {statistics.median(boss_medians):.3f}")
    print(f"  pooled by damage {sum(r['tank'] for r in rows) / sum(r['total'] for r in rows):.3f}")
    print(f"  tanks per kill median {statistics.median([r['tanks'] for r in rows])}")


main()
