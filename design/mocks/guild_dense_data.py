"""The dense guild world (design/specs/2026-11-09-guild-dense.md section 2): a 40-raider Olympus XXVII five raid nights in.

ws_guild_dense.py draws the mocks from these values and guild_dense_fixtures.py writes them as the API's answers
(web/tests/e2e/fixtures/guild-parity/dense-*.json). The first 24 raiders are guild_data.R; 16 more join them. Every figure a
row carries is set here, by hand for the raiders the pages talk about and by a fixed formula for the rest, so the mock and the
build read the same numbers.
"""
from __future__ import annotations

import hashlib
from dataclasses import dataclass

import guild_data as gd

# The clock: Sat 9 Jan 2027, 4:00 pm CST. Raids every Wednesday 8:00 pm CST since raids opened on Wed 9 Dec.
NOW = '2027-01-09T22:00:00Z'
ZONE = gd.ZONE
NIGHTS = [  # raid id, start (UTC), report id, the Onyxia kill in ms, wipes before it, deaths
    (40, '2026-12-10T02:00:00Z', 'olympusony1209', 225_000, 1, 7),
    (41, '2026-12-17T02:00:00Z', 'olympusony1216', 214_000, 0, 4),
    (42, '2026-12-24T02:00:00Z', 'olympusony1223', 206_000, 0, 3),
    (43, '2026-12-31T02:00:00Z', 'olympusony1230', 198_000, 0, 5),
    (44, '2027-01-07T02:00:00Z', 'olympusony0106', 192_000, 0, 2),
]
LAST_RAID_ID = 44
PREV_RAID_ID = 43
NEXT_RAID_ID = 45
NEXT_RAID_AT = '2027-01-14T02:00:00Z'
NEXT_POSTED_AT = '2027-01-07T18:00:00Z'
SECOND_RAID_ID = 46
SECOND_RAID_AT = '2027-01-21T02:00:00Z'
UPDATED_AT = '2027-01-09T19:30:00Z'
MEDIAN_ILVL = 64
COMP_TARGET = {'tanks': 3, 'healers': 9, 'melee': 12, 'ranged': 16}
GROUP_ROLE = {'Tank': 'tank', 'Healer': 'healer', 'Melee': 'dps', 'Ranged': 'dps'}

# The 16 who joined the first 24 (name, class, spec, group, ilvl). Survival hunters fight in melee on Forever.
JOINED = [('Ironhide', 'warrior', 'Protection', 'Tank', 61), ('Bramblehorn', 'druid', 'Feral', 'Tank', 63),
          ('Sunmantle', 'paladin', 'Holy', 'Healer', 59), ('Rainsong', 'shaman', 'Restoration', 'Healer', 62),
          ('Brightvow', 'priest', 'Discipline', 'Healer', 63), ('Gorehowl', 'warrior', 'Fury', 'Melee', 66),
          ('Axebreaker', 'warrior', 'Fury', 'Melee', 64), ('Steelmane', 'warrior', 'Arms', 'Melee', 62),
          ('Duskfang', 'rogue', 'Combat', 'Melee', 61), ('Thistlepaw', 'druid', 'Feral', 'Melee', 63),
          ('Ashblade', 'paladin', 'Retribution', 'Melee', 60), ('Rimecaller', 'mage', 'Frost', 'Ranged', 65),
          ('Glimmerfrost', 'mage', 'Frost', 'Ranged', 62), ('Felbrand', 'warlock', 'Destruction', 'Ranged', 63),
          ('Longshot', 'hunter', 'Marksmanship', 'Ranged', 64), ('Moonquill', 'druid', 'Balance', 'Ranged', 61)]

# Raid 45's answers; everyone else is In. None = not answered.
ANSWERS = {'Sunmantle': 'late', 'Pinefeather': 'late', 'Nightlatch': 'bench', 'Cinderglass': 'bench', 'Brightvow': 'out',
           'Felbrand': 'out', 'Ironhide': None, 'Rainsong': None, 'Duskfang': None, 'Voidcall': None, 'Skyarcana': None,
           'Starnettle': None, 'Moonquill': None}
# Attendance over the five nights, oldest first: P present, A absent, B bench, N before joining.
ATTENDANCE = {'Voidcall': 'PPAPA', 'Skyarcana': 'PAAPA', 'Starnettle': 'PPPAP', 'Cinderglass': 'BPBPB', 'Nightlatch': 'BBPBP',
              'Ironhide': 'NNPPP', 'Sunmantle': 'NNNPP', 'Moonquill': 'NPPAP', 'Duskfang': 'PPAPP', 'Felbrand': 'PPPPA',
              'Pinefeather': 'APPPP', 'Grimsworn': 'PPAPP', 'Rainsong': 'NPPPP'}
# Raid 45 readiness shortfalls (officer view): attuned, flask/food, enchant slots missing, unspent talent points.
NOT_ATTUNED = ['Sunmantle']
NO_FLASK_FOOD = ['Hollowmend', 'Grimsworn', 'Emberfall', 'Ashblade']
MISSING_ENCHANTS = {'Ashblade': ['hands', 'feet'], 'Glimmerfrost': ['chest']}
UNSPENT_TALENTS = {'Sunmantle': 3}
TANK_FIRE = {'Thornwicke': 142, 'Ironhide': 96, 'Bramblehorn': 118}
# Days since the addon export, where not today.
SEEN_DAYS = {'Skyarcana': 9, 'Voidcall': 8, 'Moonquill': 3, 'Starnettle': 2, 'Felbrand': 5}
# Rating on the last two kills, hand-set where the pages name the change; others follow the formula.
RATING_MOVES = {'Emberfall': (58, 77), 'Rimecaller': (71, 88), 'Seraphine': (64, 79), 'Doomwhisper': (91, 72),
                'Groundshock': (84, 66), 'Pinefeather': (60, 47)}
# Hand-set figures the member's page reads: the BiS-file gap and the best parse.
GAP_SET = {'Hollowmend': 11.0}
BEST_SET = {'Hollowmend': 82}
# Deaths on the Onyxia kill (raid id -> name -> deaths), as the report's deaths[] carries them.
DEATHS = {44: {'Doomwhisper': 1, 'Pinefeather': 1}, 43: {'Hollowmend': 1, 'Emberfall': 2, 'Groundshock': 1, 'Skytalon': 1}}
# The applicants (guild_data.APPLICANTS) by the group they would fill.
APPLICANT_GROUP = {'Ashenveil': 'Healer', 'Quillbright': 'Ranged'}
OFFICERS = list(gd.OFFICERS)
GUILD_MASTER = gd.GUILD_MASTER
MEMBER = 'Hollowmend'
OFFICER = gd.VIEWER
PROFESSIONS = ['alchemy', 'blacksmithing', 'enchanting', 'engineering', 'herbalism', 'leatherworking', 'mining', 'tailoring',
               'skinning']


def _n(name: str, salt: str, lo: int, hi: int) -> int:
    """A fixed number in [lo, hi] for a raider: the same on every run, different per raider and per figure."""
    h = int(hashlib.sha256(f'{name}:{salt}'.encode()).hexdigest()[:8], 16)
    return lo + h % (hi - lo + 1)


@dataclass(frozen=True)
class Raider:
    name: str
    cls: str
    spec: str
    group: str
    ilvl: int
    answer: str | None
    cells: str
    rank: str
    ratings: tuple[int, int]
    figures: tuple[float, float]
    best_pct: int
    gap: float | None
    upgrades: int
    professions: tuple[str, ...]
    seen_days: int

    @property
    def key(self) -> str:
        return f'{gd.REGION}/{gd.RULESET}/{self.name.lower()}'

    @property
    def role(self) -> str:
        return GROUP_ROLE[self.group]

    @property
    def metric(self) -> str:
        return 'hps' if self.group == 'Healer' else 'dps'

    @property
    def present(self) -> int:
        return self.cells.count('P')

    @property
    def nights(self) -> int:
        """Nights the raider counts for: every night since joining (bench counts as a night they were there)."""
        return len(self.cells) - self.cells.count('N')

    @property
    def spec_key(self) -> str:
        return f"{self.cls}-{self.spec.lower().replace(' ', '-')}"


def _group(name: str, cls: str, spec: str, group: str) -> str:
    return 'Melee' if (cls, spec) == ('hunter', 'Survival') else group


def _figures(name: str, group: str, ilvl: int) -> tuple[float, float]:
    """Onyxia DPS (HPS for a healer, threat-limited DPS for a tank) the week before and last week."""
    base = {'Tank': 260, 'Healer': 520, 'Melee': 610, 'Ranged': 600}[group] + (ilvl - 60) * 18
    before = base + _n(name, 'fig0', -60, 60)
    return float(before), float(before + _n(name, 'fig1', -25, 55))


def _raider(name: str, cls: str, spec: str, group: str, ilvl: int) -> Raider:
    group = _group(name, cls, spec, group)
    r0 = _n(name, 'r0', 52, 92)
    ratings = RATING_MOVES.get(name, (r0, max(30, min(99, r0 + _n(name, 'r1', -6, 9)))))
    gap = None if group == 'Tank' else GAP_SET.get(name, round(_n(name, 'gap', 0, 260) / 10, 1))
    upgrades = 0 if gap is None or gap < 2 else 1 + int(gap // 8)
    profs = (PROFESSIONS[_n(name, 'p0', 0, 8)], PROFESSIONS[_n(name, 'p1', 0, 8)])
    return Raider(name=name, cls=cls, spec=spec, group=group, ilvl=ilvl, answer=ANSWERS.get(name, 'in'),
                  cells=ATTENDANCE.get(name, 'PPPPP'), rank=gd.rank_of(name), ratings=ratings,
                  figures=_figures(name, group, ilvl), best_pct=BEST_SET.get(name, _n(name, 'pct', 38, 97)), gap=gap, upgrades=upgrades,
                  professions=tuple(sorted(set(profs))), seen_days=SEEN_DAYS.get(name, 0))


ROSTER: list[Raider] = [_raider(r[0], r[1], r[2], r[3], r[4] + 2) for r in gd.R] + [_raider(*j) for j in JOINED]
assert len(ROSTER) == 40, len(ROSTER)
BY_NAME = {r.name: r for r in ROSTER}
GROUPS = ['Tank', 'Healer', 'Melee', 'Ranged']
GROUP_TARGET = {'Tank': 'tanks', 'Healer': 'healers', 'Melee': 'melee', 'Ranged': 'ranged'}

# The council's awards over the five nights: raid id, item id, name, icon, slot word, winner, awarded by, yes votes.
AWARDS = [
    (40, 18205, "Eskhandar's Collar", 'inv_belt_12', 'Neck', 'Shiverknife', 'Ironbrand', 3),
    (40, 18813, 'Ring of Binding', 'inv_jewelry_ring_13', 'Finger', 'Thornwicke', 'Dunmaro', 3),
    (40, 17078, 'Sapphiron Drape', 'inv_misc_cape_16', 'Back', 'Frostquill', 'Dunmaro', 3),
    (40, 16955, 'Judgement Crown', 'inv_helmet_74', 'Head', 'Ironbrand', 'Shiverknife', 2),
    (41, 16914, 'Netherwind Crown', 'inv_helmet_70', 'Head', 'Frostquill', 'Ironbrand', 3),
    (41, 16929, 'Nemesis Skullcap', 'inv_helmet_08', 'Head', 'Doomwhisper', 'Dunmaro', 3),
    (41, 17068, 'Deathbringer', 'inv_axe_09', 'One-Hand', 'Gorehowl', 'Ironbrand', 3),
    (41, 17064, 'Shard of the Scale', 'inv_misc_monsterscales_15', 'Trinket', 'Hollowmend', 'Dunmaro', 3),
    (42, 16908, 'Bloodfang Hood', 'inv_helmet_41', 'Head', 'Shiverknife', 'Ironbrand', 2),
    (42, 16947, 'Helmet of Ten Storms', 'inv_helmet_69', 'Head', 'Stonetide', 'Dunmaro', 3),
    (42, 17075, "Vis'kag the Bloodletter", 'inv_sword_18', 'One-Hand', 'Blacktide', 'Dunmaro', 3),
    (42, 18813, 'Ring of Binding', 'inv_jewelry_ring_13', 'Finger', 'Bramblehorn', 'Shiverknife', 3),
    (43, 16963, 'Helm of Wrath', 'inv_helmet_71', 'Head', 'Thornwicke', 'Ironbrand', 3),
    (43, 16921, 'Halo of Transcendence', 'inv_helmet_24', 'Head', 'Dawnweaver', 'Dunmaro', 3),
    (43, 17067, 'Ancient Cornerstone Grimoire', 'inv_misc_book_07', 'Off Hand', 'Hexmarrow', 'Dunmaro', 2),
    (43, 18205, "Eskhandar's Collar", 'inv_belt_12', 'Neck', 'Groundshock', 'Ironbrand', 3),
    (44, 16939, "Dragonstalker's Helm", 'inv_helmet_05', 'Head', 'Dunmaro', 'Ironbrand', 2),
    (44, 16900, 'Stormrage Cover', 'inv_helmet_09', 'Head', 'Mossveil', 'Shiverknife', 3),
    (44, 17078, 'Sapphiron Drape', 'inv_misc_cape_16', 'Back', 'Rimecaller', 'Dunmaro', 3),
]
# Last night's drops still open (the Loot tab's night picks Vis'kag): item id, name, icon, slot word.
OPEN_DROPS = [(17075, "Vis'kag the Bloodletter", 'inv_sword_18', 'One-Hand'), (16955, 'Judgement Crown', 'inv_helmet_74', 'Head')]
SELECTED_DROP = 17075
# Its council: name, gain in DPS, worn main hand (id, name, icon, quality), BiS for them, MS/OS, yes votes.
CANDIDATES = [('Duskfang', 16.8, (13361, 'Skullforge Reaver', 'inv_sword_12', 3), True, 'ms', 2),
              ('Shiverknife', 9.8, (12940, "Dal'Rend's Sacred Charge", 'inv_sword_43', 3), False, 'ms', 1),
              ('Gorehowl', 7.1, (15806, "Mirah's Song", 'inv_sword_34', 3), False, 'ms', 0),
              ('Axebreaker', 5.3, (15806, "Mirah's Song", 'inv_sword_34', 3), False, 'ms', 0)]


# Judgement Crown's council: the paladins without it (gain in HPS for the Holy, DPS for the Retribution).
CROWN_CANDIDATES = [('Valomira', 14.0, (16854, 'Lawbringer Helm', 'inv_helmet_05', 4), True, 'ms', 1),
                    ('Sunmantle', 11.2, (16854, 'Lawbringer Helm', 'inv_helmet_05', 4), False, 'ms', 0),
                    ('Seraphine', 6.4, (12640, 'Lionheart Helm', 'inv_helmet_36', 4), False, 'os', 0)]


def items_won(name: str) -> list[tuple]:
    return [a for a in AWARDS if a[5] == name]


def last_award_raid(name: str) -> int | None:
    won = items_won(name)
    return won[-1][0] if won else None


def answered(state: str | None) -> list[Raider]:
    return [r for r in ROSTER if r.answer == state]


def group_counts(group: str) -> dict[str, int]:
    rows = [r for r in ROSTER if r.group == group]
    return {s: sum(1 for r in rows if r.answer == s) for s in ('in', 'late', 'bench', 'out', None)}


def short_for_raid(r: Raider) -> list[str]:
    """What the readiness board fails `r` on for raid 45."""
    out = []
    if r.name in NOT_ATTUNED:
        out.append('attunement')
    if r.name in NO_FLASK_FOOD:
        out.append('flask_food')
    if r.name in MISSING_ENCHANTS:
        out.append('enchants')
    if r.name in UNSPENT_TALENTS:
        out.append('talents')
    return out


def night_at(raid_id: int) -> str:
    return next(n[1] for n in NIGHTS if n[0] == raid_id)
