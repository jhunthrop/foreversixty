"""Writes the F16 fixtures: guild.home.roster, guild.raids.history, guild.raids.history-2,
and adds the W3-C16 fields to guild.home. Invented people; real zone sizes and bosses."""
import base64, copy, hashlib, json, sys
from datetime import datetime, timedelta, timezone

ROOT = sys.argv[1]
FX = f"{ROOT}/docs/contracts/fixtures/app"
RID = "9f2c41d07ab3e655"

SPECS = {  # key: (catalogue name, role, icon)
    "warrior-protection": ("Protection", "tank", "ability_warrior_defensivestance"),
    "warrior-fury": ("Fury", "dps", "ability_warrior_innerrage"),
    "warrior-arms": ("Arms", "dps", "ability_warrior_savageblow"),
    "druid-feral-bear": ("Feral Bear", "tank", "ability_racial_bearform"),
    "druid-restoration": ("Restoration", "healer", "spell_nature_healingtouch"),
    "druid-balance": ("Balance", "dps", "spell_nature_starfall"),
    "priest-holy": ("Holy", "healer", "spell_holy_guardianspirit"),
    "priest-shadow": ("Shadow", "dps", "spell_shadow_shadowwordpain"),
    "shaman-restoration": ("Restoration", "healer", "spell_nature_magicimmunity"),
    "shaman-elemental": ("Elemental", "dps", "spell_nature_lightning"),
    "shaman-enhancement": ("Enhancement", "dps", "spell_nature_lightningshield"),
    "rogue-combat": ("Combat", "dps", "ability_backstab"),
    "mage-frost": ("Frost", "dps", "spell_frost_frostbolt02"),
    "mage-fire": ("Fire", "dps", "spell_fire_firebolt02"),
    "warlock-affliction": ("Affliction", "dps", "spell_shadow_deathcoil"),
    "warlock-destruction": ("Destruction", "dps", "spell_shadow_rainoffire"),
    "hunter-marksmanship": ("Marksmanship", "dps", "ability_marksmanship"),
    "hunter-beast-mastery": ("Beast Mastery", "dps", "ability_hunter_beasttaming"),
}


def check_icons():
    """Icons must be the catalogue's (sim/specs/specs.go)."""
    src = open(f"{ROOT}/sim/specs/specs.go").read()
    for key, (_, _, icon) in SPECS.items():
        line = next(l for l in src.splitlines() if f'Spec: "{key}"' in l)
        real = line.split('Icon: "')[1].split('"')[0]
        if real != icon:
            SPECS[key] = (SPECS[key][0], SPECS[key][1], real)
        name = line.split('Name: "')[1].split('"')[0]
        role = line.split('Role: "')[1].split('"')[0]
        assert (name, role) == SPECS[key][:2], (key, name, role)


check_icons()

# name, spec key, level, rank, present-of-8, state, supporter
PEOPLE = [
    ("Kraggor", "warrior-protection", 60, "leader", 8, "approved", True),
    ("Mendara", "priest-holy", 60, "officer", 8, "approved", False),
    ("Draknar", "warrior-fury", 60, "officer", 7, "approved", False),
    ("Lightningbolltz", "shaman-elemental", 60, "member", 6, "approved", True),
    ("Rainspeaker", "shaman-restoration", 60, "member", 6, "approved", False),
    ("Skarn", "rogue-combat", 60, "member", 5, "approved", False),
    ("Zekra", "rogue-combat", 60, "member", 5, "approved", False),
    ("Morvane", "warlock-affliction", 60, "member", 8, "approved", False),
    ("Thokgar", "warrior-protection", 60, "officer", 8, "approved", False),
    ("Ursoka", "druid-feral-bear", 60, "member", 7, "approved", False),
    ("Hollowmere", "priest-holy", 60, "member", 8, "approved", False),
    ("Sennari", "priest-holy", 60, "member", 4, "approved", False),
    ("Tidecaller", "shaman-restoration", 60, "member", 7, "approved", False),
    ("Zul'jak", "shaman-restoration", 60, "member", 6, "approved", False),
    ("Ombrewood", "druid-restoration", 60, "member", 8, "approved", False),
    ("Fernhoof", "druid-restoration", 59, "member", 3, "approved", False),
    ("Vashkar", "shaman-restoration", 60, "member", 5, "approved", False),
    ("Grimtusk", "warrior-fury", 60, "member", 8, "approved", False),
    ("Bloodmaw", "warrior-fury", 60, "member", 6, "approved", False),
    ("Korrakh", "warrior-arms", 60, "member", 2, "approved", False),
    ("Sliverfang", "rogue-combat", 60, "member", 7, "approved", False),
    ("Nightvex", "rogue-combat", 60, "member", 8, "approved", False),
    ("Pyrelle", "mage-fire", 60, "member", 7, "approved", False),
    ("Frostmaw", "mage-frost", 60, "member", 8, "approved", False),
    ("Embersong", "mage-frost", 60, "member", 6, "approved", False),
    ("Quillara", "mage-frost", 58, "member", 1, "approved", False),
    ("Hexmire", "warlock-destruction", 60, "member", 7, "approved", False),
    ("Duskwhisper", "warlock-affliction", 60, "member", 5, "approved", False),
    ("Arrowmane", "hunter-marksmanship", 60, "member", 8, "approved", False),
    ("Talonstrike", "hunter-marksmanship", 60, "member", 6, "approved", False),
    ("Grizzlepaw", "hunter-beast-mastery", 60, "member", 4, "approved", False),
    ("Stormhide", "shaman-enhancement", 60, "member", 3, "approved", False),
    ("Voidsong", "priest-shadow", 57, "member", 0, "waiting", False),
    ("Moonveil", "druid-balance", 56, "member", 0, "waiting", False),
    ("Gruzz", "warrior-fury", 54, "member", None, "no_account", False),
    ("Shadeleaf", "priest-holy", 60, "member", None, "no_account", False),
    ("Tazzlo", "mage-fire", 52, "member", None, "no_account", False),
    ("Brokkar", "warrior-protection", 49, "member", None, "no_account", False),
]
NIGHTS = 8


def slug(name):
    return name.lower().replace("'", "")


def spec_words(key):
    name, role, icon = SPECS[key]
    cls = key.split("-")[0].capitalize()
    return name, role, icon, f"{name} {cls}"


def roster_row(base, p, idx):
    name, key, level, rank, present, state, supporter = p
    sname, role, icon, full = spec_words(key)
    row = copy.deepcopy(base)
    ck = f"us/pvp/{slug(name)}"
    row.update({
        "account_key": f"u:{1800 + idx}", "character_key": ck, "class": key.split("-")[0],
        "name": name, "rank": rank, "role": role, "spec": sname, "spec_name": full, "spec_icon": icon,
        "supporter": supporter, "verified": state == "approved", "may_approve": False,
        "may_remove": False, "attendance": {"present": present, "nights": NIGHTS},
        "attendance_30": {"present": present, "nights": NIGHTS},
        "logged_at": "2026-12-17T03:10:00Z", "last_report_at": "2026-12-17T05:02:00Z" if present else None,
        "item_level": 58 + (idx * 7) % 9,
    })
    row["addon"] = {"level": level, "spec": key, "gear_gap": idx % 5, "build": None,
                    "seen_at": "2026-12-16 21:40", "snapshot_at": "2026-12-17T03:11:00Z"}
    row.pop("best_parse", None)
    row.pop("rating", None)
    row.pop("professions", None)
    return row


def member(p):
    name, key, level, rank, present, state, supporter = p
    sname, role, icon, full = spec_words(key)
    m = {"name": name, "state": state, "rank": rank, "class": key.split("-")[0], "level": level,
         "spec": key, "seen_at": "2026-12-16 21:40", "role": role, "spec_name": full, "spec_icon": icon}
    if state != "no_account":
        m["character_key"] = f"us/pvp/{slug(name)}"
    return m


def load(name):
    return json.load(open(f"{FX}/{name}.json"))


def dump(name, obj):
    with open(f"{FX}/{name}.json", "w") as f:
        json.dump(obj, f, indent=2, ensure_ascii=False)
        f.write("\n")


def write_home():
    home = load("guild.home")
    d = home["data"]
    for m in d["members"]:
        if "spec" in m:
            sname, role, icon, full = spec_words(m["spec"])
            m.update(role=role, spec_name=full, spec_icon=icon)
        else:
            m.update(role=None, spec_name=None, spec_icon=None)
    for r in d["roster"]:
        key = f'{r["class"]}-{r["spec"].lower().replace(" ", "-")}'
        sname, role, icon, full = spec_words(key)
        r.update(spec_name=full, spec_icon=icon)
    dump("guild.home", home)

    roster = copy.deepcopy(home)
    d = roster["data"]
    base = d["roster"][0]
    accounts = [p for p in PEOPLE if p[5] != "no_account"]
    d["roster"] = [roster_row(base, p, i) for i, p in enumerate(accounts)]
    d["pending"] = []
    d["members"] = [member(p) for p in PEOPLE]
    d["counts"] = {"on_roster": len(PEOPLE), "in_game_no_account": sum(p[5] == "no_account" for p in PEOPLE),
                   "waiting": sum(p[5] == "waiting" for p in PEOPLE), "not_seen": 0}
    d["viewer"] = {"role": "member", "character_key": "us/pvp/lightningbolltz", "verified": True,
                   "officer_tools": False, "raid_guild_tools": False, "verify": {"has_approver": True}}
    d["addon_snapshot"]["by_character"] = "us/pvp/kraggor"
    d["addon_snapshot"]["members"] = len(PEOPLE)
    d["addon_snapshot"]["received_at"] = "2026-12-17T03:11:00Z"
    d["summary"]["raiders"] = len(accounts)
    d["summary"]["waiting_for_approval"] = 2
    d["standing"] = None
    d.pop("standing")
    d["claim"]["claimed_by_name"] = "Kraggor#1187"
    assert d["counts"] == {"on_roster": 38, "in_game_no_account": 4, "waiting": 2, "not_seen": 0}, d["counts"]
    dump("guild.home.roster", roster)


# ---- raid history ---------------------------------------------------------------------
ONY = [("Onyxia", 1084)]
UBRS = [("Pyroguard Emberseer", 276), ("Solakar Flamewreath", 277), ("Warchief Rend Blackhand", 278),
        ("The Beast", 279), ("General Drakkisath", 280)]
LBRS = [("Highlord Omokk", 267), ("Shadow Hunter Vosh'gajin", 268), ("War Master Voone", 269),
        ("Mother Smolderweb", 270), ("Quartermaster Zigris", 272), ("Overlord Wyrmthalak", 275)]
BRD = [("Lord Roccor", 228), ("Ring of Law", 230), ("General Angerforge", 236),
       ("Golem Lord Argelmach", 237), ("Ambassador Flamelash", 242), ("Emperor Dagran Thaurissan", 245)]

ROSTER_KEYS = [(f"us/pvp/{slug(p[0])}", p[0], p[1].split("-")[0]) for p in PEOPLE if p[5] == "approved"]
ME = "us/pvp/lightningbolltz"

# (local date at the guild's 7:45 pm, zone, bosses, size, raiders, wipes on the last boss, minutes, me)
# Raids open Wed 9 Dec 2026; every night here is on or after it. Onyxia is a 40; UBRS a 10;
# LBRS and BRD are five-player dungeons. Lengths differ night to night.
NIGHTS_PLAN = [
    ("2026-12-31", "Onyxia's Lair", ONY, 40, 40, 0, 52, True),
    ("2026-12-30", "Upper Blackrock Spire", UBRS, 10, 10, 1, 108, True),
    ("2026-12-29", "Blackrock Depths", BRD, 5, 5, 0, 131, False),
    ("2026-12-27", "Onyxia's Lair", ONY, 40, 39, 1, 71, True),
    ("2026-12-26", "Lower Blackrock Spire", LBRS, 5, 5, 0, 76, True),
    ("2026-12-24", "Onyxia's Lair", ONY, 40, 38, 0, 58, False),
    ("2026-12-23", "Upper Blackrock Spire", UBRS, 10, 10, 2, 124, False),
    ("2026-12-22", "Blackrock Depths", BRD, 5, 5, 1, 142, True),
    ("2026-12-20", "Onyxia's Lair", ONY, 40, 40, 1, 66, True),
    ("2026-12-19", "Lower Blackrock Spire", LBRS, 5, 5, 0, 83, False),
    ("2026-12-17", "Onyxia's Lair", ONY, 40, 38, 0, 107, True),
    ("2026-12-16", "Upper Blackrock Spire", UBRS, 10, 10, 0, 96, True),
    ("2026-12-15", "Blackrock Depths", BRD, 5, 5, 0, 98, False),
    ("2026-12-14", "Lower Blackrock Spire", LBRS, 5, 5, 0, 72, True),
    ("2026-12-13", "Onyxia's Lair", ONY, 40, 40, 3, 112, True),
    ("2026-12-12", "Upper Blackrock Spire", UBRS, 10, 10, 1, 118, False),
    ("2026-12-12T14:00", "Blackrock Depths", BRD, 5, 5, 0, 116, True),
    ("2026-12-11", "Lower Blackrock Spire", LBRS, 5, 5, 1, 88, False),
    ("2026-12-10", "Onyxia's Lair", ONY, 40, 39, 2, 146, True),
    ("2026-12-10T13:30", "Lower Blackrock Spire", LBRS, 5, 5, 0, 79, True),
    # page two
    ("2026-12-09", "Onyxia's Lair", ONY, 40, 40, 4, 138, False),
    ("2026-12-09T15:10", "Upper Blackrock Spire", UBRS, 10, 10, 2, 131, True),
    ("2026-12-09T11:20", "Blackrock Depths", BRD, 5, 5, 1, 152, False),
]


def night_row(i, plan):
    date, zone, bosses, size, raiders, wipes, minutes, me = plan
    local = datetime.fromisoformat((date if "T" in date else date + "T19:45") + ":00-06:00")
    created = (local + timedelta(minutes=minutes + 6)).astimezone(timezone.utc)
    fights, idx, deaths = [], 1, 0
    per_boss = max(4, minutes * 60_000 // (len(bosses) * 6))
    for b, (name, enc) in enumerate(bosses):
        tries = wipes if b == len(bosses) - 1 else 0
        for t in range(tries + 1):
            kill = t == tries
            d = (1 + (i + b) % 4) if size > 10 else (i + b) % 2
            if not kill:
                d = raiders // 2 if size > 10 else size - 1
            deaths += d
            players = raiders - (1 if (size == 40 and not kill and i % 2) else 0)
            dur = (190_000 + (i * 37_000) % 120_000) if name == "Onyxia" else (60_000 + ((i + b) * 23_000) % 150_000)
            if not kill:
                dur = dur // 2
            fights.append({"index": idx, "name": name, "encounter_id": enc, "kill": kill,
                           "duration_ms": dur, "deaths": d, "players": players})
            idx += 2  # trash between pulls takes the other indexes
    pool = [r for r in ROSTER_KEYS if r[0] != ME]
    take = raiders - (1 if me else 0)
    present = [pool[(i * 3 + k) % len(pool)] for k in range(take)]
    if me:
        present.insert(0, next(r for r in ROSTER_KEYS if r[0] == ME))
    present = [{"character_key": k, "name": n, "class": c} for k, n, c in dict.fromkeys(present)]
    while len(present) < raiders:  # pugs: a dungeon or a 40 fills from outside the guild
        n = ["Brakka", "Velith", "Orrin", "Juzzi", "Kelda"][len(present) % 5]
        present.append({"character_key": f"us/pvp/{n.lower()}", "name": n, "class": "warrior"})
    rid = base64.b32encode(hashlib.sha1(f"{date}{zone}".encode()).digest()).decode().lower()[:12]
    top_class = ["mage", "warlock", "rogue", "warrior"][i % 4]
    return {
        "id": rid, "title": f"{zone}, {local.strftime('%A')}", "zone": zone,
        "created_at": created.strftime("%Y-%m-%dT%H:%M:%SZ"), "duration_ms": minutes * 60_000,
        "fight_count": len(fights) + 3 * len(bosses), "kill_count": len(bosses), "wipe_count": wipes,
        "raiders": raiders, "deaths": deaths + (2 if size == 40 else 0),
        "top_parse": {"name": next(p[0] for p in PEOPLE if p[1].startswith(top_class)), "class": top_class,
                      "metric": "dps", "value": round(480 + (i * 53) % 260 + 0.4, 1)},
        "fights": fights, "present": present,
        "raid_id": 40 - i if zone == "Onyxia's Lair" else None,
    }


def write_history():
    base = load("guild.raids")
    rows = [night_row(i, p) for i, p in enumerate(NIGHTS_PLAN)]
    for r in rows:
        assert r["created_at"] >= "2026-12-09T06:00:00Z", r
        assert all(f["players"] <= r["raiders"] for f in r["fights"]), r
        assert r["raiders"] in (5, 10, 38, 39, 40)
    page1, page2 = copy.deepcopy(base), copy.deepcopy(base)
    page1["data"]["rows"] = rows[:20]
    last = rows[19]
    page1["data"]["next_cursor"] = base64.b64encode(f'{last["created_at"]}|{last["id"]}'.encode()).decode()
    page2["data"]["rows"] = rows[20:]
    page2["data"].pop("next_cursor")
    for p in (page1, page2):
        p["data"]["planned"] = []
        p["data"].pop("waiting_for_approval", None)
    dump("guild.raids.history", page1)
    dump("guild.raids.history-2", page2)


write_home()
write_history()
print("ok")
