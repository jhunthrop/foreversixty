"""Guild parity fixtures, the Overview states the persona round left out: before raids open (an officer, the screenshot
case), a low-level member, no raid planned or logged once the raids are open, and a signed-in visitor who is in no guild.
The API answers web/tests/e2e/support/guild-parity/states.ts routes, written by guild_fixtures.py into
web/tests/e2e/fixtures/guild-parity/ (files `states-*.json`); ws_guild_states.py draws the mocks from the same values.
"""
from datetime import datetime, timedelta, timezone
from pathlib import Path

import guild_data as gd
import guild_overview_data as D
import guild_overview_fixtures as F
from guild_data import GUILD, R, REGION, RULESET

# The clocks: Sat 7 Nov 2026, 4:00 pm CST before the raids open (32 days 2 hours to Wed 9 Dec, 00:00 UTC), and Tue 15 Dec,
# 4:00 pm CST once they are open. The visitor state keeps the spec's own clock (guild_data.NOW).
PRE_NOW = '2026-11-07T22:00:00Z'
OPEN_NOW = '2026-12-15T22:00:00Z'

# Before the raids open: the officer (Dunmaro) and the low-level member (Nightlatch) are both level 23, as in the
# owner's screenshot. Every other character's level, and the professions the roster carries (slugs, as the export sends).
LOW_LEVEL = 23
LOW_MEMBER = 'Nightlatch'
PRE_LEVELS = {'Thornwicke': 60, 'Ironbrand': 60, 'Valomira': 54, 'Seraphine': 58, 'Dunmaro': LOW_LEVEL, 'Pinefeather': 47,
              'Skytalon': 41, 'Shiverknife': 60, 'Blacktide': 52, 'Hollowmend': 60, 'Dawnweaver': 44, 'Voidcall': 31,
              'Stonetide': 55, 'Emberfall': 49, 'Groundshock': 38, 'Frostquill': 57, 'Cinderglass': 42, 'Skyarcana': 27,
              'Doomwhisper': 51, 'Grimsworn': 45, 'Hexmarrow': 36, 'Mossveil': 48, 'Starnettle': None, 'Nightlatch': LOW_LEVEL}
PROFESSIONS = {'Ironbrand': ['blacksmithing', 'mining'], 'Hollowmend': ['tailoring', 'enchanting'], 'Shiverknife': ['alchemy', 'herbalism'],
               'Thornwicke': ['blacksmithing', 'mining'], 'Stonetide': ['alchemy', 'herbalism'], 'Frostquill': ['tailoring', 'enchanting'],
               'Doomwhisper': ['tailoring', 'enchanting'], 'Blacktide': ['leatherworking', 'skinning']}
# Once the raids are open every raider is 60 but two.
OPEN_LEVELS = {name: 60 for name, *_ in R} | {'Voidcall': 58, 'Skyarcana': 57}
NO_RAID_MEMBER = D.MEMBER


def stamp(at: str, minutes: int) -> str:
    """`at` less `minutes`, as the API writes an instant."""
    t = datetime.fromisoformat(at.replace('Z', '+00:00')) - timedelta(minutes=minutes)
    return t.astimezone(timezone.utc).strftime('%Y-%m-%dT%H:%M:%SZ')


def row(name: str) -> tuple:
    return D.row(name)


# ------------------------------------------------------------------ the guild's activity (GET /v1/feed?filter=guild)
def sim_row(i: int, name: str, figure: str, unit: str, rank: str, at: str) -> dict:
    """A guild sim line (home B spec P7): `<Name> simmed <figure> <unit>, #<r> of <N> <role> in <Guild> (<Spec>, <band>)`."""
    _, cls, spec, role, *_ = row(name)
    words = 'healers' if role == 'Healer' else 'tanks' if role == 'Tank' else 'damage dealers'
    return {'id': f'g{i}', 'kind': 'guild_sim', 'at': at, 'tag': 'your_guild', 'icon': None, 'crest': cls, 'quality': None,
            'href': f'/character/{REGION}/{RULESET}/{name.lower()}', 'date_only': False,
            'segments': [{'text': name, 'style': 'name', 'class': cls}, {'text': f' simmed {figure} {unit}, {rank} {words} in {GUILD} ({spec} {cls.title()}, 60)'}]}


def pre_feed() -> dict:
    return {'rows': [sim_row(1, 'Shiverknife', '486.2', 'DPS', '#1 of 4', stamp(PRE_NOW, 38)),
                     sim_row(2, 'Ironbrand', '402.7', 'HPS', '#1 of 2', stamp(PRE_NOW, 60 * 5)),
                     sim_row(3, 'Thornwicke', '311.0', 'DPS', '#2 of 4', stamp(PRE_NOW, 60 * 26))], 'next_cursor': None}


def open_feed() -> dict:
    return {'rows': [sim_row(1, 'Frostquill', '612.4', 'DPS', '#1 of 14', stamp(OPEN_NOW, 12)),
                     sim_row(2, 'Doomwhisper', '598.0', 'DPS', '#2 of 14', stamp(OPEN_NOW, 60 * 3))], 'next_cursor': None}


# ------------------------------------------------------------------ the session and the home
def me(name: str, officer: bool, level: int) -> dict:
    page = F.me(name, officer)
    return {**page, 'characters': [{**c, 'level': level} for c in page['characters']]}


def visitor_me() -> dict:
    """A signed-in account in no guild: Felsnap, a level 34 Fury Warrior."""
    return {'user': {'id': 9, 'battletag': 'Felsnap#1', 'email': None, 'role': 'user', 'anonymize': False},
            'characters': [{'key': f'{REGION}/{RULESET}/felsnap', 'region': REGION, 'ruleset': RULESET, 'name': 'Felsnap', 'class': 'Warrior',
                            'spec': 'Fury', 'level': 34, 'realm': 'Skyborne', 'faction': 'alliance', 'guild': None}],
            'guilds': []}


def empty_tier() -> dict:
    return {**F.tier(), 'down': 0, 'pulls_total': 0}


def home(viewer: str, officer: bool, levels: dict, updated_at: str) -> dict:
    base = F.home(viewer, officer)
    roster = [{**F.roster_row(r), 'level': levels[r[0]], 'professions': PROFESSIONS.get(r[0], []), 'attendance': {'present': 0, 'nights': 0},
               'logged_at': updated_at} for r in R if levels[r[0]] is not None]
    members = [{'name': name, 'character_key': D.key_of(name), 'state': 'approved', 'rank': gd.rank_of(name),
                'class': cls, **({} if levels[name] is None else {'level': levels[name]})} for name, cls, *_ in R]
    return {**base, 'summary': {**base['summary'], 'waiting_for_approval': 0, 'tier': empty_tier(), 'named_encounters_down': 0, 'pulls_this_tier': 0,
                                'updated_at': updated_at},
            'reports': [], 'roster': roster, 'pending': [], 'members': members, 'next_raid': None}


def officer_actions() -> dict:
    return {'approvals': {'count': 0, 'names': []}, 'applications': {'count': 0, 'names': []},
            'not_signed': {'raid_id': None, 'count': 0, 'names': []}, 'drops_to_award': None}


def checklist(name: str, level: int) -> dict:
    _, cls, spec, *_ = row(name)
    return {'character': {'character_key': D.key_of(name), 'name': name, 'class': cls, 'spec': spec, 'level': level}, 'raid': None,
            'verdict': 'getting_started', 'items': []}


def empty_progression() -> dict:
    page = F.progression()
    return {'tier': empty_tier(), 'unnamed': [], 'encounters': [{**e, 'pulls': 0, 'kills': 0, 'pulls_by_night': [], 'deaths_per_pull': 0,
                                                                 'first_kill_at': None, 'best_kill_ms': None} for e in page['encounters']]}


def readiness() -> dict:
    return {**F.readiness(False), 'rows': []}


def fixtures() -> dict[str, object]:
    return {
        'states-me-officer-pre.json': me(D.OFFICER, True, LOW_LEVEL),
        'states-me-member-pre.json': me(LOW_MEMBER, False, LOW_LEVEL),
        'states-me-member-open.json': me(NO_RAID_MEMBER, False, 60),
        'states-me-visitor.json': visitor_me(),
        'states-home-officer-pre.json': home(D.OFFICER, True, PRE_LEVELS, stamp(PRE_NOW, 60)),
        'states-home-member-pre.json': home(LOW_MEMBER, False, PRE_LEVELS, stamp(PRE_NOW, 60)),
        'states-home-member-open.json': home(NO_RAID_MEMBER, False, OPEN_LEVELS, stamp(OPEN_NOW, 60)),
        'states-feed-pre.json': pre_feed(),
        'states-feed-open.json': open_feed(),
        'states-actions.json': officer_actions(),
        'states-attendance.json': {'raids': [], 'rows': []},
        'states-checklist-pre.json': checklist(LOW_MEMBER, LOW_LEVEL),
        'states-checklist-open.json': checklist(NO_RAID_MEMBER, 60),
        'states-progression.json': empty_progression(),
        'states-raids.json': {'rows': [], 'next_cursor': None},
        'states-readiness.json': readiness(),
        'states-loot-nights.json': {'officers': 2, 'nights': []},
        'states-applications.json': {'rows': [], 'pending': 0, 'discord_role_id': None},
    }


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
