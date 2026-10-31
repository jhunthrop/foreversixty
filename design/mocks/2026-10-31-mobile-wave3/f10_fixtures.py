"""One story for every F10 fixture: Obnoxious Yell's guild Onyxia night, Monday 14 December in the guild's Chicago evening (report ewlrx3enxbeb).

The kill is fight 9, 4:28 (268,000 ms), 40 players, 3 deaths. Obnoxious Yell's 190,923 damage over
268 s is 712.40 DPS, the 712.4 DPS and 91st percentile of me.raid-night; every figure below is
derived from the rows, never typed twice. People besides Obnoxious Yell are invented.
"""
import json
import sys
from pathlib import Path

OUT = Path(sys.argv[1])
RID = "9f2c41d07ab3e655"
REPORT = "ewlrx3enxbeb"
KILL_MS = 268_000


def env(data):
    return {"ok": True, "data": data, "error": None, "request_id": RID}


def err(code, message):
    return {"ok": False, "data": None, "error": {"code": code, "message": message}, "request_id": RID}


def write(name, body):
    (OUT / f"{name}.json").write_text(json.dumps(body, indent=2, ensure_ascii=False) + "\n")


def figure(total, ms):
    return round(total / (ms / 1000) * 100) / 100


# (name, class, spec slug, spec name, role, dps or hps, damage taken per second, deaths, mine)
ROSTER = [
    ("Kraggor", "warrior", "warrior-protection", "Protection", "tank", 402.1, 1210.4, 0, False),
    ("Bonecrush", "warrior", "warrior-protection", "Protection", "tank", 318.6, 402.7, 0, False),
    ("Zulgrin", "rogue", "rogue-combat", "Combat", "dps", 788.3, 21.4, 0, False),
    ("Vexara", "mage", "mage-fire", "Fire", "dps", 761.9, 30.2, 0, False),
    ("Grimtusk", "warrior", "warrior-fury", "Fury", "dps", 744.6, 64.8, 0, False),
    ("Shadowfen", "warlock", "warlock-destruction", "Destruction", "dps", 731.2, 25.1, 0, False),
    ("Kelthar", "rogue", "rogue-combat", "Combat", "dps", 726.8, 33.0, 0, False),
    ("Obnoxious Yell", "warrior", "warrior-fury", "Fury", "dps", 712.4, 58.3, 0, True),
    ("Mordak", "mage", "mage-frost", "Frost", "dps", 698.7, 19.6, 0, False),
    ("Thrallin", "hunter", "hunter-marksmanship", "Marksmanship", "dps", 684.0, 12.2, 0, False),
    ("Ashvane", "warlock", "warlock-affliction", "Affliction", "dps", 671.5, 22.8, 0, False),
    ("Rokhan", "warrior", "warrior-fury", "Fury", "dps", 662.9, 71.0, 1, False),
    ("Sylvari", "mage", "mage-fire", "Fire", "dps", 655.3, 28.4, 0, False),
    ("Drezzik", "rogue", "rogue-combat", "Combat", "dps", 647.1, 30.9, 0, False),
    ("Gorlash", "hunter", "hunter-beast-mastery", "Beast Mastery", "dps", 633.8, 14.5, 0, False),
    ("Nightbloom", "warlock", "warlock-destruction", "Destruction", "dps", 628.4, 20.0, 0, False),
    ("Fenrik", "mage", "mage-frost", "Frost", "dps", 611.0, 18.8, 0, False),
    ("Brakka", "warrior", "warrior-arms", "Arms", "dps", 603.6, 66.1, 0, False),
    ("Venomtip", "rogue", "rogue-assassination", "Assassination", "dps", 597.2, 27.3, 0, False),
    ("Hexlash", "warlock", "warlock-affliction", "Affliction", "dps", 584.9, 23.6, 1, False),
    ("Talonstrike", "hunter", "hunter-marksmanship", "Marksmanship", "dps", 572.3, 11.9, 0, False),
    ("Embercall", "mage", "mage-fire", "Fire", "dps", 566.0, 26.5, 0, False),
    ("Skarn", "shaman", "shaman-elemental", "Elemental", "dps", 551.7, 24.0, 0, False),
    ("Quickblade", "rogue", "rogue-combat", "Combat", "dps", 543.4, 29.7, 0, False),
    ("Gutrend", "warrior", "warrior-fury", "Fury", "dps", 531.8, 69.4, 0, False),
    ("Ravenna", "hunter", "hunter-survival", "Survival", "dps", 522.6, 40.3, 0, False),
    ("Frostmaw", "mage", "mage-frost", "Frost", "dps", 509.9, 17.1, 0, False),
    ("Duskwhisper", "priest", "priest-shadow", "Shadow", "dps", 497.5, 22.0, 0, False),
    ("Kazrik", "shaman", "shaman-enhancement", "Enhancement", "dps", 486.2, 52.6, 1, False),
    ("Moonfang", "druid", "druid-feral", "Feral", "dps", 470.8, 36.4, 0, False),
    ("Lightweaver", "priest", "priest-holy", "Holy", "healer", 541.3, 8.2, 0, False),
    ("Stormcaller", "shaman", "shaman-restoration", "Restoration", "healer", 528.7, 15.9, 0, False),
    ("Solace", "priest", "priest-holy", "Holy", "healer", 503.4, 6.6, 0, False),
    ("Riverwind", "shaman", "shaman-restoration", "Restoration", "healer", 489.1, 14.3, 0, False),
    ("Mendra", "priest", "priest-discipline", "Discipline", "healer", 455.8, 9.1, 0, False),
    ("Greenhorn", "druid", "druid-restoration", "Restoration", "healer", 441.2, 12.7, 0, False),
    ("Tidemother", "shaman", "shaman-restoration", "Restoration", "healer", 418.6, 13.4, 0, False),
    ("Halcyon", "priest", "priest-holy", "Holy", "healer", 402.9, 7.8, 0, False),
    ("Earthsong", "shaman", "shaman-restoration", "Restoration", "healer", 389.5, 16.0, 0, False),
    ("Brightmane", "priest", "priest-holy", "Holy", "healer", 371.0, 5.9, 0, False),
]
assert len(ROSTER) == 40
assert sum(r[7] for r in ROSTER) == 3, "the kill row says 3 deaths"


def player(row):
    name, cls, spec, spec_name, role, value, dtps, deaths, mine = row
    seconds = KILL_MS / 1000
    damage = round(value * seconds) if role != "healer" else round(value * 0.04 * seconds)
    healing = round(value * seconds) if role == "healer" else 0
    taken = round(dtps * seconds)
    key = "us/pvp/" + name.lower().replace(" ", "-")
    return {
        "key": key, "name": name, "class": cls, "spec": spec, "spec_name": spec_name, "role": role,
        "active_ms": round(KILL_MS * (0.97 if role != "healer" else 0.91)),
        "damage_done": damage, "healing_done": healing, "damage_taken": taken,
        "dps": figure(damage, KILL_MS), "hps": figure(healing, KILL_MS), "dtps": figure(taken, KILL_MS),
        "deaths": deaths, "mine": mine,
    }


PLAYERS = [player(r) for r in ROSTER]
YELL = next(p for p in PLAYERS if p["mine"])
assert YELL["dps"] == 712.4, YELL["dps"]
by_damage = sorted([p for p in PLAYERS if p["damage_done"] > 0], key=lambda p: (-p["dps"], p["name"]))

write("reports.meters", env({
    "fight_index": 9, "scope": "", "metric": "damage", "kill": True, "duration_ms": KILL_MS,
    "players": by_damage,
}))
write("reports.meters.error-not-found", err("not_found", "no such fight"))

write("reports.me", env({
    "report_id": REPORT,
    "characters": [{
        "key": "us/pvp/obnoxious-yell", "name": "Obnoxious Yell", "class": "warrior", "spec": "warrior-fury",
        "spec_name": "Fury", "role": "dps",
        "fights": [{
            "index": 9, "encounter_id": 1084, "encounter": "Onyxia", "kill": True, "duration_ms": KILL_MS,
            "value": YELL["dps"], "unit": "dps", "percentile": 91, "ranked": 1840, "new_best": True,
            "deaths": YELL["deaths"],
        }],
    }],
}))
write("reports.me.empty", env({"report_id": REPORT, "characters": []}))
write("reports.me.error-not-found", err("not_found", "no such report"))

# The night: six trash pulls, two wipes, the kill. 9 fights, 3 boss pulls, 1 kill.
START = 1_797_300_300  # 2026-12-15T02:05:00Z: the first pull, after the 01:00 raid start and the 01:58 report
FIGHTS = [
    (1, "trash", "Trash", None, 95_000, False, 0, None),
    (2, "trash", "Trash", None, 120_000, False, 0, None),
    (3, "trash", "Trash", None, 64_000, False, 0, None),
    (4, "encounter", "Onyxia", 1084, 72_000, False, 6, 61.0),
    (5, "trash", "Trash", None, 41_000, False, 0, None),
    (6, "encounter", "Onyxia", 1084, 292_000, False, 11, 22.0),  # a phase-3 wipe, longer than the kill
    (7, "trash", "Trash", None, 88_000, False, 0, None),
    (8, "trash", "Trash", None, 52_000, False, 0, None),
    (9, "encounter", "Onyxia", 1084, KILL_MS, True, 3, None),
]
GUIDS = [f"Player-4184-{0xA1 + i:08X}" for i in range(len(ROSTER))]  # the 40 of the roster


def iso(t):
    import datetime
    return datetime.datetime.fromtimestamp(t, datetime.UTC).strftime("%Y-%m-%dT%H:%M:%SZ")


fights, t = [], START
for index, kind, name, enc, ms, kill, deaths, health in FIGHTS:
    f = {"index": index, "kind": kind, "name": name}
    if enc:
        f.update({"encounter_id": enc, "difficulty": 8, "size": 40})
    f.update({
        "kill": kill, "in_progress": False, "zone": "Onyxia's Lair", "start": iso(t), "end": iso(t + ms // 1000),
        "duration_ms": ms, "players": GUIDS, "deaths": deaths, "mob_deaths": 1 if kill else 0,
        "group_kills": 1 if kill else 0, "group_deaths": deaths, "group": GUIDS, "verified": True,
    })
    if health is not None:
        f["boss_health_pct"] = health
    fights.append(f)
    t += ms // 1000 + 240
trash = [f for f in fights if f["kind"] == "trash"]
write("reports.get", env({
    "id": REPORT, "title": "Onyxia, Monday", "visibility": "guild",
    "owner": {"id": 1812, "battletag": "Yell#1604"},
    "guild": {"id": 2, "name": "OLYMPUS XXVII", "region": "us", "ruleset": "pvp"},
    "zone": "Onyxia's Lair", "status": "complete", "engine_version": "1f542def3",
    "fights": fights, "players": GUIDS, "created_at": "2026-12-15T01:58:00Z",
    "last_upload_at": "2026-12-15T02:55:40Z", "heartbeat_at": "2026-12-15T02:55:30Z",
    "data_base_url": f"https://api.foreversixty.gg/v1/reports/{REPORT}/files",
    "logging_character": {"key": "us/pvp/obnoxious-yell", "name": "Obnoxious Yell"},
    "trash_groups": [{
        "key": trash[0]["index"], "zone": "Onyxia's Lair", "fights": [f["index"] for f in trash],
        "duration_ms": sum(f["duration_ms"] for f in trash),
    }],
}))

YELL_REF = {"key": "us/pvp/obnoxious-yell", "name": "Obnoxious Yell"}
write("reports.list", env({
    "rows": [
        {"id": "m4xk9d2pqz7a", "title": "", "zone": "Blackrock Spire", "status": "live", "visibility": "private",
         "created_at": "2026-12-16T00:10:00Z", "fight_count": 4, "kill_count": 0, "boss_count": 0,
         "logging_character": YELL_REF, "last_upload_at": "2026-12-16T00:31:12Z", "heartbeat_at": "2026-12-16T00:31:40Z"},
        {"id": REPORT, "title": "Onyxia, Monday", "zone": "Onyxia's Lair", "status": "complete", "visibility": "guild",
         "created_at": "2026-12-15T01:58:00Z", "fight_count": len(fights),
         "kill_count": sum(1 for f in fights if f["kill"]),
         "boss_count": len({f["name"] for f in fights if f["kill"]}),
         "logging_character": YELL_REF, "last_upload_at": "2026-12-15T02:55:40Z", "heartbeat_at": "2026-12-15T02:55:30Z"},
        {"id": "q7c2v9tk4mzh", "title": "", "zone": "Molten Core", "status": "processing", "visibility": "public",
         "created_at": "2026-12-12T01:30:00Z", "fight_count": 0, "kill_count": 0, "boss_count": 0,
         "logging_character": YELL_REF, "last_upload_at": "2026-12-12T04:02:10Z", "heartbeat_at": None},
        {"id": "t3n8b5xw2kqe", "title": "", "zone": "Stratholme", "status": "failed", "visibility": "unlisted",
         "created_at": "2026-11-29T20:15:00Z", "fight_count": 0, "kill_count": 0, "boss_count": 0,
         "logging_character": YELL_REF, "last_upload_at": "2026-11-29T21:40:00Z", "heartbeat_at": None},
    ],
    "total": 4, "page": 1, "per_page": 100,
}))

# The trash group's combined meters: the same 40 over the six trash pulls (460,000 ms, the
# group's summed length), each at a little over half their boss rate, as AoE trash reads.
TRASH_MS = sum(f["duration_ms"] for f in trash)
assert TRASH_MS == 460_000


def trash_player(row):
    p = player(row)
    scale = 0.55
    for total in ("damage_done", "healing_done", "damage_taken"):
        p[total] = round(p[total] / KILL_MS * TRASH_MS * scale)
    p["active_ms"] = round(TRASH_MS * 0.82)
    p["dps"], p["hps"], p["dtps"] = (figure(p["damage_done"], TRASH_MS), figure(p["healing_done"], TRASH_MS),
                                     figure(p["damage_taken"], TRASH_MS))
    p["deaths"] = 0
    return p


TRASH_PLAYERS = sorted([trash_player(r) for r in ROSTER], key=lambda p: (-p["dps"], p["name"]))
write("reports.meters.trash", env({
    "fight_index": trash[0]["index"], "scope": "trash", "metric": "damage", "kill": False,
    "duration_ms": TRASH_MS, "players": [p for p in TRASH_PLAYERS if p["damage_done"] > 0],
}))

# The live report on the list (m4xk9d2pqz7a): a private Blackrock Spire run still being logged,
# four trash pulls so far, no boss. Its list row reads 4 fights, 0 bosses, 0 kills.
BRS_START = START + 22 * 3600 + 420  # 2026-12-16T00:10Z created; pulls from 00:12Z
brs, t = [], BRS_START
for index, ms in ((1, 48_000), (2, 71_000), (3, 39_000), (4, 66_000)):
    brs.append({"index": index, "kind": "trash", "name": "Trash", "kill": False, "in_progress": False,
                "zone": "Blackrock Spire", "start": iso(t), "end": iso(t + ms // 1000), "duration_ms": ms,
                "players": GUIDS[:5], "deaths": 0, "mob_deaths": 6, "group_kills": 6, "group_deaths": 0,
                "group": GUIDS[:5], "verified": True})
    t += ms // 1000 + 180
write("reports.get.live", env({
    "id": "m4xk9d2pqz7a", "title": "", "visibility": "private",
    "owner": {"id": 1812, "battletag": "Yell#1604"},
    "zone": "Blackrock Spire", "status": "live", "engine_version": "1f542def3",
    "fights": brs, "players": GUIDS[:5], "created_at": "2026-12-16T00:10:00Z",
    "last_upload_at": "2026-12-16T00:31:12Z", "heartbeat_at": "2026-12-16T00:31:40Z",
    "data_base_url": "https://api.foreversixty.gg/v1/reports/m4xk9d2pqz7a/files",
    "logging_character": {"key": "us/pvp/obnoxious-yell", "name": "Obnoxious Yell"},
    "trash_groups": [{"key": 1, "zone": "Blackrock Spire", "fights": [1, 2, 3, 4],
                      "duration_ms": sum(f["duration_ms"] for f in brs)}],
}))
assert len(brs) == 4  # the list row's fight_count
