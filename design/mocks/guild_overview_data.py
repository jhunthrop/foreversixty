"""The guild Overview and loot council's own data (lane A of the guild parity states): the next raid, the member's checklist,
the officer's waiting items, the recent attendance, the last raid and the loot council's night. Typed once here; ws_guild2.py
draws landing_member, landing_officer and loot_council from it and guild_overview_fixtures.py writes it as the API's answers.
The roster, its sign-up states and the clock are guild_data.py's.
"""
import guild_data as gd
from guild_data import LAST_RAID_AT, LAST_RAID_DATE, LAST_RAID_ID, LAST_REPORT_ID, NEXT_RAID_AT, NEXT_RAID_ID, R  # noqa: F401

ZONE_KEY = 'onyxias-lair'
ZONE_NAME = "Onyxia's Lair"
BOSS = 'Onyxia'
ENCOUNTER_ID = 1084
# Onyxia's Lair's zone default seats (40 players), and the attunement it asks for.
COMP_TARGET = {'tanks': 2, 'healers': 8, 'melee': 10, 'ranged': 20}
ATTUNEMENT_ITEM = {'id': 16309, 'name': 'Drakefire Amulet', 'quality': 3}

# The viewers (ws_guild2 build_page `who`): a member and an officer of the guild.
MEMBER = 'Hollowmend'
OFFICER = gd.VIEWER
UPDATED_AT = '2026-12-12T19:30:00Z'

# The member's checklist for the next raid (api guilds checklist items, in the API's order).
GEAR_SLOTS = ['hands', 'finger1']
FLASK_MISSING = ['flask_or_elixir', 'food']

# What is waiting on an officer: three characters waiting for approval (not raiders: no sign-ups, no attendance; roster rows
# name, class, spec, role, ilvl), the applications by name, and the newest loot night's open drops.
PENDING = [('Brambleshot', 'hunter', 'Marksmanship', 'Ranged', 57), ('Mirelocke', 'warlock', 'Affliction', 'Ranged', 58),
           ('Oakhallow', 'druid', 'Restoration', 'Healer', 59)]
APPROVALS = [p[0] for p in PENDING]
APPLICATIONS = ['Ashenveil', 'Quillbright']

# The raid nights logged so far, oldest first: the first night only (the raids opened Wed 9 Dec).
RECENT_RAIDS = [(LAST_RAID_ID, LAST_RAID_AT)]

# Raid-ready on the signed raiders: who has not finished the attunement, who carries no flask and food, the tanks' fire resistance.
# Pinefeather has no Drakefire Amulet, so they missed the first night (guild_report_data.AWAY).
NOT_ATTUNED = ['Pinefeather']
NO_FLASK_FOOD = ['Hollowmend', 'Grimsworn', 'Emberfall']
TANK_FIRE = {'Thornwicke': 142}

# The loot council's night (the first night): the drops (item id, name, icon, quality, kind, open or the award: winner,
# council member who awarded it, yes votes of gd.VOTERS).
VOTERS = gd.VOTERS
LOOT_NIGHT = f'raid:{LAST_RAID_ID}'
DROPS = [(17075, "Vis'kag the Bloodletter", 'inv_sword_18', 4, 'gear', None), (17068, 'Deathbringer', 'inv_axe_09', 4, 'gear', None),
         (18205, "Eskhandar's Collar", 'inv_belt_12', 4, 'gear', ('Shiverknife', gd.GUILD_MASTER, 3)),
         (18813, 'Ring of Binding', 'inv_jewelry_ring_13', 4, 'gear', ('Thornwicke', OFFICER, 3)),
         (17078, 'Sapphiron Drape', 'inv_misc_cape_16', 4, 'gear', ('Frostquill', OFFICER, 3)),
         (18422, 'Head of Onyxia', 'inv_misc_head_dragon_01', 4, 'quest', None)]
SELECTED_DROP = 17075
# Its candidates, the raiders who wield a one-handed sword: name, gain in DPS, the main-hand item worn now (id, name, icon,
# quality), BiS for them, MS/OS, yes votes. Attendance and items this tier come from the night (candidate_facts).
CANDIDATES = [('Blacktide', 14.2, (13361, 'Skullforge Reaver', 'inv_sword_47', 3), True, 'ms', 2),
              ('Shiverknife', 9.8, (12940, "Dal'Rend's Sacred Charge", 'inv_sword_36', 3), False, 'ms', 1),
              ('Thornwicke', 5.3, (15806, "Mirah's Song", 'inv_sword_24', 3), False, 'ms', 0)]


def awards_to(name: str) -> list[tuple]:
    """The night's drops awarded to `name`."""
    return [d for d in DROPS if d[5] is not None and d[5][0] == name]


def row(name: str) -> tuple:
    """The roster row of guild_data.R: name, class, spec, role, ilvl, sign-up state."""
    return next(r for r in R if r[0] == name)


def key_of(name: str) -> str:
    return f'us/pvp/{name.lower()}'


def spec_key(cls: str, spec: str) -> str:
    return f"{cls}-{spec.lower().replace(' ', '-')}"


def signed(state: str) -> list[tuple]:
    return [r for r in R if r[5] == state]


def in_by_group() -> dict[str, int]:
    """In rows by roster group (Tank, Healer, Melee, Ranged)."""
    out = {'Tank': 0, 'Healer': 0, 'Melee': 0, 'Ranged': 0}
    for r in signed('in'):
        out[r[3]] += 1
    return out


def last_raid_counts() -> dict[str, int]:
    """The first night's boss kills and wipes, its trash pulls and every pull."""
    boss = [p for p in gd.NIGHT_PULLS if p[2] is not None]
    kills = sum(1 for p in boss if p[3] == 'kill')
    return {'kills': kills, 'wipes': len(boss) - kills, 'trash': len(gd.NIGHT_PULLS) - len(boss), 'pulls': len(gd.NIGHT_PULLS)}


def came(name: str) -> bool:
    """Whether `name` is in the first night's log: a raider on the roster the night did not miss."""
    return any(r[0] == name for r in R) and name not in gd.NIGHT_AWAY


def expected(name: str) -> bool:
    """Whether the first night counts for `name`: signed In or Late, or came (the roster's attendance rule)."""
    return came(name) or any(r[0] == name and r[5] in ('in', 'late') for r in R)
