"""The guild Overview and loot council's own data (lane A of the guild parity states): the next raid, the member's checklist,
the officer's waiting items, the recent attendance, the last raid and the loot council's night. Typed once here; ws_guild2.py
draws landing_member, landing_officer and loot_council from it and guild_overview_fixtures.py writes it as the API's answers.
The roster, its sign-up states and the clock are guild_data.py's.
"""
from guild_data import R

# The next raid (Thu 10 Dec, 8:00 pm CST) and the last one (Mon 7 Dec, 8:00 pm CST), as the API's instants.
NEXT_RAID_ID = 41
NEXT_RAID_AT = '2026-12-11T02:00:00Z'
ZONE_KEY = 'onyxias-lair'
ZONE_NAME = "Onyxia's Lair"
BOSS = 'Onyxia'
ENCOUNTER_ID = 1084
# Onyxia's Lair's zone default seats (40 players), and the attunement it asks for.
COMP_TARGET = {'tanks': 2, 'healers': 8, 'melee': 10, 'ranged': 20}
ATTUNEMENT_ITEM = {'id': 16309, 'name': 'Drakefire Amulet', 'quality': 3}

# The viewers (ws_guild2 build_page `who`): a member and an officer of the guild.
MEMBER = 'Hollowmend'
OFFICER = 'Dunmaro'
UPDATED_AT = '2026-12-08T19:30:00Z'

# The last raid night: its report, its pulls (name, kill, seconds, deaths; a None name is trash) and its date in the guild's zone.
LAST_REPORT_ID = 'olympusony1207'
LAST_RAID_ID = 38
LAST_RAID_AT = '2026-12-08T02:00:00Z'
LAST_RAID_DATE = '2026-12-07'
LAST_PULLS = [(None, True, 60, 0), (BOSS, False, 240, 6), (None, True, 55, 0), (BOSS, True, 225, 1)]
LAST_ABSENT = ['Voidcall', 'Starnettle']

# The member's checklist for the next raid (api guilds checklist items, in the API's order).
GEAR_SLOTS = ['hands', 'finger1']
FLASK_MISSING = ['flask_or_elixir', 'food']

# What is waiting on an officer: approvals and applications by name, and the newest loot night's open drops.
APPROVALS = ['Nightlatch', 'Voidcall', 'Skyarcana']
APPLICATIONS = ['Ashenveil', 'Quillbright']

# The last four raids (Chicago dates at 8:00 pm) and the slipping and benched raiders' cells (present|absent|bench|none).
RECENT_RAIDS = [(35, '2026-11-27T02:00:00Z'), (36, '2026-11-29T02:00:00Z'), (37, '2026-12-04T02:00:00Z'), (LAST_RAID_ID, LAST_RAID_AT)]
RECENT_ROWS = [('Pinefeather', True, ['present', 'present', 'absent', 'absent']), ('Skyarcana', True, ['present', 'absent', 'absent', 'none']),
               ('Cinderglass', False, ['bench', 'bench', 'present', 'bench'])]

# Raid-ready on the signed raiders: who has not finished the attunement, who carries no flask and food, the tanks' fire resistance.
NOT_ATTUNED = ['Blacktide', 'Pinefeather']
NO_FLASK_FOOD = ['Hollowmend', 'Grimsworn', 'Emberfall']
TANK_FIRE = {'Thornwicke': 142}

# The loot council's night: the drops (item id, name, icon, quality, kind, open or the award: winner, officer, votes yes of 4).
OFFICERS = 4
LOOT_NIGHT = f'raid:{LAST_RAID_ID}'
DROPS = [(17075, "Vis'kag the Bloodletter", 'inv_sword_18', 4, 'gear', None), (17068, 'Deathbringer', 'inv_axe_09', 4, 'gear', None),
         (18205, "Eskhandar's Collar", 'inv_belt_12', 4, 'gear', ('Hollowmend', 'Ironbrand', 3)),
         (18813, 'Ring of Binding', 'inv_jewelry_ring_13', 4, 'gear', ('Doomwhisper', OFFICER, 4)),
         (17078, 'Sapphiron Drape', 'inv_misc_cape_16', 4, 'gear', None), (18422, 'Head of Onyxia', 'inv_misc_head_dragon_01', 4, 'quest', None)]
SELECTED_DROP = 17075
# Its candidates: name, gain in DPS, the item worn now (id, name, icon, quality), BiS for them, raids attended of 4, items this tier, MS/OS, yes votes.
CANDIDATES = [('Seraphine', 14.2, (15806, "Mirah's Song", 'inv_sword_24', 3), True, 4, 0, 'ms', 3),
              ('Shiverknife', 11.6, (12940, "Dal'Rend's Sacred Charge", 'inv_sword_36', 3), False, 3, 1, 'ms', 1),
              ('Groundshock', 9.1, (11684, 'Ironfoe', 'inv_mace_10', 4), False, 4, 0, 'os', 1),
              ('Blacktide', 4.0, (13361, 'Skullforge Reaver', 'inv_sword_47', 3), False, 2, 2, 'os', 1)]


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
    kills = sum(1 for name, kill, _, _ in LAST_PULLS if name is not None and kill)
    wipes = sum(1 for name, kill, _, _ in LAST_PULLS if name is not None and not kill)
    return {'kills': kills, 'wipes': wipes, 'trash': sum(1 for name, *_ in LAST_PULLS if name is None), 'pulls': len(LAST_PULLS)}
