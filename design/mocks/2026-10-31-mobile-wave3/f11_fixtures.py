"""One story for every F11 fixture: the Onyxia board for Elemental Shamans, US Normal, Phase 1,
as it stands on Thursday 17 December 2026 (spec §0 "One character, one number").

Lightningbolltz (Hordebreakers) is 14th of 212 Elemental Shaman characters at 398.0 DPS; his
alt Sparkwick is 204th at 210.4 DPS (F11.3). Figures fall off as a power of the rank so the
14th row reads 398.0 exactly. Gear and best-in-slot picks are real items of the active build's
items/shaman.json and its Elemental raid list; people, guilds, parses and scores are invented.

Run: python3 design/mocks/2026-10-31-mobile-wave3/f11_fixtures.py docs/contracts/fixtures/app
"""
import hashlib
import itertools
import json
import math
import sys
from datetime import datetime, timedelta, timezone
from pathlib import Path

OUT = Path(sys.argv[1])
REPO = Path(__file__).resolve().parents[3]
BUILD = json.loads((REPO / "web/src/data/active-build.json").read_text())["build"]
RID = "9f2c41d07ab3e655"
TODAY = datetime(2026, 12, 17, 20, 0, tzinfo=timezone.utc)
ONYXIA = 1084
TOP, YOURS, YOUR_RANK, TOTAL = 462.0, 398.0, 14, 212
EXPONENT = math.log(TOP / YOURS) / math.log(YOUR_RANK)


def env(data):
    return {"ok": True, "data": data, "error": None, "request_id": RID}


def err(code, message, fields=None):
    error = {"code": code, "message": message}
    if fields:
        error["fields"] = fields
    return {"ok": False, "data": None, "error": error, "request_id": RID}


def write(name, body):
    (OUT / f"{name}.json").write_text(json.dumps(body, indent=2, ensure_ascii=False) + "\n")


def iso(at):
    return at.strftime("%Y-%m-%dT%H:%M:%SZ")


def slug(name):
    return name.lower().replace(" ", "-")


GUILDS = ["Hordebreakers", "Nightfall Raiders", "OLYMPUS XXVII", "Iron Pact", "Blood & Thunder",
          "Ashen Covenant", "Thunder Bluff Irregulars", "Bloodfang"]
FIRST = ["Volt", "Storm", "Zul", "Crackle", "Thunder", "Ash", "Spark", "Flint", "Gale", "Rumble",
         "Ember", "Static", "Tide", "Cinder", "Shock"]
LAST = ["aris", "wrath", "thrax", "heim", "gaze", "fang", "mane", "jaw", "hoof", "tusk"]
BOARD_TOP = ["Voltaris", "Stormwrath", "Zulthrax", "Crackleheim", "Thundergaze"]


def names():
    seen = set(BOARD_TOP) | {"Lightningbolltz", "Sparkwick"}
    out = list(BOARD_TOP)
    combos = sorted((f + l for f, l in itertools.product(FIRST, LAST)),
                    key=lambda n: hashlib.md5(n.encode()).hexdigest())
    for name in combos:
        if name not in seen:
            seen.add(name)
            out.append(name)
    return out


def value(rank):
    return round(TOP * rank ** -EXPONENT, 1)


def row(rank, name, val, guild, days, seconds):
    fought = TODAY - timedelta(days=days, hours=(rank * 5) % 7)
    return {
        "rank": rank,
        "player": {"key": f"us/normal/{slug(name)}", "name": name, "class": "Shaman", "spec": "Elemental"},
        "guild": {"name": guild, "region": "us", "ruleset": "normal"},
        "value": val,
        "ilvl": 63 - rank % 4,
        "size": 40,
        "fought_at": iso(fought),
        "duration_ms": seconds * 1000,
        "talent_split": "30/0/21",
        "trinkets": [19120, 11302],
        "buff_count": 6 + rank % 3,
        "report_id": f"onyx{rank:04d}kill",
        "fight_index": 9,
        "state": "ok",
        "execution_score": round(0.97 - rank * 0.002, 3) if rank % 5 else None,
    }


def board():
    pool = names()
    rows = []
    for rank in range(1, 101):
        if rank == YOUR_RANK:
            rows.append(row(rank, "Lightningbolltz", YOURS, "Hordebreakers", 1, 194))
            continue
        name = pool.pop(0)
        guild = "Hordebreakers" if rank == 1 else GUILDS[1 + rank % (len(GUILDS) - 1)]
        rows.append(row(rank, name, value(rank), guild, 1 + rank % 7, 192 + rank % 23))
    return rows


def page(rows, character_rank):
    data = {"rows": rows, "total": TOTAL, "page": 1, "per_page": 100,
            "updated_at": iso(TODAY - timedelta(minutes=4)), "deduped": True}
    if character_rank:
        data["character_rank"] = character_rank
    return data


def rankings():
    rows = board()
    RANKED_ROWS.extend(rows)
    write("rankings.list", env(page(rows, {"rank": YOUR_RANK, "row": rows[YOUR_RANK - 1]})))
    spark = row(204, "Sparkwick", 210.4, "Hordebreakers", 6, 211)
    write("rankings.list.off-page", env(page(rows, {"rank": 204, "row": spark})))
    empty = {"rows": [], "total": 0, "page": 1, "per_page": 100,
             "updated_at": iso(TODAY - timedelta(minutes=4)), "deduped": True}
    write("rankings.list.empty", env(empty))
    write("rankings.percentile", env({"percentile": 93.4, "ranked": TOTAL}))
    write("rankings.encounters", env({"encounters": [
        {"id": ONYXIA, "slug": "onyxia", "name": "Onyxia", "zone": "Onyxia's Lair", "size": 40, "kills": 38},
        {"id": 227, "slug": "emperor-dagran-thaurissan", "name": "Emperor Dagran Thaurissan",
         "zone": "Blackrock Depths", "size": 5, "kills": 31},
        {"id": 267, "slug": "overlord-wyrmthalak", "name": "Overlord Wyrmthalak",
         "zone": "Lower Blackrock Spire", "size": 5, "kills": 12},
    ]}))


def watching():
    rows = [
        {"key": "us/normal/voltaris", "region": "us", "ruleset": "normal", "name": "Voltaris",
         "class": "Shaman", "level": 60, "added_at": iso(TODAY - timedelta(days=2))},
        {"key": "us/normal/stormwrath", "region": "us", "ruleset": "normal", "name": "Stormwrath",
         "class": "Shaman", "level": 60, "added_at": iso(TODAY - timedelta(days=5))},
        {"key": "us/normal/kraggor", "region": "us", "ruleset": "normal", "name": "Kraggor",
         "class": "Warrior", "level": None, "added_at": iso(TODAY - timedelta(days=6))},
    ]
    write("watching.list", env({"rows": rows}))
    write("watching.list.empty", env({"rows": []}))
    write("watching.put", env(rows[0]))
    write("watching.put.error-limit", err("limit", "you can watch up to 50 characters",
                                          {"character": "stop watching one of your 50 first"}))
    write("watching.put.error-not-found", err("not_found", "no such character"))


# (slot, worn, pick or None when the worn item is the pick, gap); the picks are the Elemental
# raid list's, with that list's source words. Voltaris (first) wears nearly the whole list;
# Lightningbolltz (14th) is a week behind him in the Elements set.
PICK_SOURCES = {"Living Crown": "Leatherworking", "Ironfeather Breastplate": "Leatherworking",
                "Ironfeather Shoulders": "Leatherworking", "Ironfeather Leggings": "Leatherworking",
                "Skycaller's Leather Waistguard": "Leatherworking", "Uther's Strength": "World drop"}
BIS = [
    ("head", "Living Crown"), ("neck", "Archlight Talisman"), ("shoulder", "Ironfeather Shoulders"),
    ("back", "Spritecaster Cape"), ("chest", "Ironfeather Breastplate"), ("wrist", "Runecloth Cuffs"),
    ("hands", "Storm Gauntlets"), ("waist", "Skycaller's Leather Waistguard"), ("legs", "Ironfeather Leggings"),
    ("feet", "Omnicast Boots"), ("finger1", "Maiden's Circle"), ("finger2", "Band of the Unicorn"),
    ("trinket1", "Rune of the Guard Captain"), ("trinket2", "Uther's Strength"),
    ("main_hand", "Hand of Edward the Odd"), ("off_hand", "Spirit of Aquementas"), ("ranged", "Totem of Thunder"),
]
# slot: (worn instead of the pick, gap)
VOLTARIS_GAPS = {"trinket2": ("Briarwood Reed", 3.0)}
YOURS_GAPS = {
    "head": ("Coif of Elements", 7.8), "shoulder": ("Pauldrons of Elements", 5.1),
    "chest": ("Vest of Elements", 9.4), "waist": ("Cord of Elements", 4.1),
    "legs": ("Kilt of Elements", 6.6), "trinket2": ("Briarwood Reed", 3.0),
}
GEARS = (VOLTARIS_GAPS, YOURS_GAPS)


def shaman_items():
    data = json.loads((REPO / f"data/builds/{BUILD}/items/shaman.json").read_text())
    by_name = {}
    for item in data["items"]:
        by_name.setdefault(item["name"], item)
    return data, by_name


def gear_panel(by_name, gaps):
    slots = []
    for slot, pick_name in BIS:
        worn_name, gap = gaps.get(slot, (pick_name, None))
        worn = by_name[worn_name]
        entry = {"slot": slot, "worn": {"id": worn["id"], "name": worn["name"], "item_level": worn["item_level"]},
                 "pick": None, "gap": gap}
        if worn_name != pick_name:
            pick = by_name[pick_name]
            entry["pick"] = {"id": pick["id"], "name": pick["name"], "source": PICK_SOURCES[pick_name], "chance": None}
        slots.append(entry)
    return slots, round(sum(g for _, g in gaps.values()), 1)


def card(name, key, race, guild, score_parts, bests, history, slots, gap_total, history_points):
    score = sum(score_parts.values())
    tier = "epic" if score >= 1500 else "rare"
    return {
        "character": {"key": key, "region": "us", "ruleset": "normal", "name": name, "class": "Shaman",
                      "guild": {"name": guild, "rank": "member"}, "level": 60, "race": race, "realm": None},
        "best": bests, "history": history,
        "builds_seen": [{"talent_split": "30/0/21", "spec": "Elemental", "first_seen": iso(TODAY - timedelta(days=7))}],
        "score": score, "tier": tier, "as_of": "2026-12-17",
        "components": score_parts,
        "experience": {"dungeons": 14, "raid_kills": 1, "alts": 1},
        "component_caps": {"gear": 500, "experience": 700, "parses": 600, "attendance": 200, "discipline": 100},
        "score_history": history_points,
        "supporter": False,
        "gear": slots, "gear_gap_total": gap_total,
        "gear_spec": "shaman-elemental", "gear_spec_source": "talents",
    }


def fight(encounter, encounter_id, kill, val, pct, at, ms, report, index):
    out = {"encounter": encounter, "encounter_id": encounter_id, "difficulty": 8 if encounter_id == ONYXIA else 1,
           "kill": kill, "metric": "dps", "value": val, "spec": "Elemental", "talent_split": "30/0/21",
           "duration_ms": ms, "fought_at": at, "report_id": report,
           "fight_index": index, "execution_score": 0.91}
    if pct is not None:
        out["percentile"] = pct
    return out


def best(h):
    keys = ("encounter", "encounter_id", "difficulty", "metric", "value", "percentile", "spec", "fought_at",
            "report_id", "fight_index")
    return {k: h[k] for k in keys if k in h}


def score_points(final):
    points = []
    for days, score in ((7, final - 210), (5, final - 140), (3, final - 60), (1, final - 20), (0, final)):
        tier = "epic" if score >= 1500 else "rare"
        points.append({"date": (TODAY - timedelta(days=days)).strftime("%Y-%m-%d"), "score": score, "tier": tier})
    return points


def ago(days):
    return iso(TODAY - timedelta(days=days))


def kill_of(board_row, pct):
    """A character card's kill, the same fight as its board row: same value, time and report."""
    return fight("Onyxia", ONYXIA, True, board_row["value"], pct, board_row["fought_at"],
                 board_row["duration_ms"], board_row["report_id"], board_row["fight_index"])


def cards(by_name, rows):
    volt_slots, volt_gap = gear_panel(by_name, VOLTARIS_GAPS)
    volt_kill = kill_of(rows[0], 99)
    volt_hist = [
        volt_kill,
        fight("Onyxia", ONYXIA, False, 441.3, None, volt_kill["fought_at"], 74000, volt_kill["report_id"], 7),
        fight("Emperor Dagran Thaurissan", 227, True, 431.2, 81, ago(3), 96000, "brd0003run", 14),
    ]
    volt = card("Voltaris", "us/normal/voltaris", "Orc", "Hordebreakers",
                {"gear": 412, "experience": 620, "parses": 540, "attendance": 180, "discipline": 78},
                [best(volt_hist[0]), best(volt_hist[2])], volt_hist, volt_slots, volt_gap, score_points(1830))
    write("character.card.ranked", env(volt))
    your_slots, your_gap = gear_panel(by_name, YOURS_GAPS)
    your_kill = kill_of(rows[YOUR_RANK - 1], 93)
    yours_hist = [
        your_kill,
        fight("Onyxia", ONYXIA, False, 377.9, None, your_kill["fought_at"], 71000, your_kill["report_id"], 7),
    ]
    yours = card("Lightningbolltz", "us/normal/lightningbolltz", "Troll", "Hordebreakers",
                 {"gear": 371, "experience": 410, "parses": 498, "attendance": 160, "discipline": 71},
                 [best(yours_hist[0])], yours_hist, your_slots, your_gap, score_points(1510))
    write("character.card.yours", env(yours))


def items_fixture(data, by_name):
    wanted = {n for _, n in BIS} | {worn for gaps in GEARS for worn, _ in gaps.values()}
    items = sorted((by_name[n] for n in wanted), key=lambda i: i["id"])
    write("cdn.items.shaman", {"build": data["build"], "class_slug": "shaman", "items": items})


RANKED_ROWS = []


def main():
    rankings()
    watching()
    data, by_name = shaman_items()
    cards(by_name, RANKED_ROWS)
    items_fixture(data, by_name)


main()
