"""Writes the guild parity fixtures: web/tests/e2e/fixtures/guild-parity/*.json, the API answers the build is served
(web/tests/e2e/support/guild-parity/) for the data ws_guild2.py draws the mocks from (guild_data.py).

  python3 design/mocks/guild_fixtures.py
"""
import json
import re
from pathlib import Path

import guild_data as gd

OUT = Path(__file__).resolve().parents[2] / 'web/tests/e2e/fixtures/guild-parity'


PRINT_WIDTH = 110  # web/.prettierrc printWidth: the fixtures pass `prettier --check` as written
_SCALAR_LIST = re.compile(r'^( *)("[^"\n]*": )?\[\n((?:\1  (?:"(?:[^"\\\n]|\\.)*"|-?[\d.eE+-]+|true|false|null),?\n)+)\1\]', re.M)


def _one_line(m: re.Match) -> str:
    indent, key, body = m.group(1), m.group(2) or '', m.group(3)
    line = f'{indent}{key}[{", ".join(x.strip().rstrip(",") for x in body.splitlines())}]'
    return line if len(line) + 1 <= PRINT_WIDTH else m.group(0)


def write(name: str, data: object) -> Path:
    """JSON as prettier prints it: two-space indent, a list of scalars on one line when it fits."""
    OUT.mkdir(parents=True, exist_ok=True)
    path = OUT / name
    path.write_text(_SCALAR_LIST.sub(_one_line, json.dumps(data, indent=2)) + '\n')
    return path


def write_raid_planner() -> list[Path]:
    """Lane B: Thursday's Onyxia's Lair planner as its officer (Dunmaro) reads it (guild_raid.py)."""
    import guild_raid
    return [write('raid-me.json', guild_raid.me()), write('raid-guild.json', guild_raid.guild_page()),
            write('raid-home.json', guild_raid.guild_home()), write('raid-onyxia.json', guild_raid.raid_view())]


# One writer per page group; each returns the files it wrote.
WRITERS: list = [write_raid_planner]

# The raid night report (state `report`): guild_report_data.py.
from guild_report_data import write_report_fixtures  # noqa: E402
WRITERS.append(lambda: write_report_fixtures(write))


def overview_writer() -> list[Path]:
    """Lane A: the member and officer Overviews and the Loot tab (guild_overview_fixtures.py)."""
    from guild_overview_fixtures import overview_fixtures
    return [write(name, data) for name, data in overview_fixtures().items()]


WRITERS.append(overview_writer)


# ---- Recruitment, applications and the member profile (lane D; support/guild-parity/recruit.ts routes these files).
def _char_key(name: str) -> str:
    return f'{gd.REGION}/{gd.RULESET}/{name.lower()}'


def _guild() -> dict:
    return {'id': gd.GUILD_ID, 'name': gd.GUILD, 'region': gd.REGION, 'ruleset': gd.RULESET, 'faction': gd.FACTION, 'crest_url': None}


def _ref(name: str) -> dict:
    return {'name': name, 'class': gd.R[_index(name)][1], 'character_key': _char_key(name)}


def _index(name: str) -> int:
    return next(i for i, r in enumerate(gd.R) if r[0] == name)


def _parse(metric: str, value: float, percentile: int) -> dict:
    return {'metric': metric, 'value': value, 'percentile': percentile, 'encounter': 'Onyxia', 'report_id': 'parity', 'fight_index': 3}


def recruit_card() -> dict:
    officers = [{'name': n, 'class': gd.R[_index(n)][1], 'spec': gd.R[_index(n)][2], 'rank': 'officer'} for n in gd.OFFICERS]
    return {'guild': {**_guild(), 'path': f'/guild/{gd.REGION}/{gd.RULESET}/{gd.SLUG}'}, 'members': gd.RAIDERS, 'open': True, 'pitch': gd.RECRUIT_PITCH,
            'raid_nights': {'text': gd.RECRUIT_NIGHTS, 'upcoming': gd.RECRUIT_UPCOMING, 'timezone': gd.ZONE},
            'progress': {'name': 'Tier 1', 'raids': ["Onyxia's Lair"], 'named_encounters': 1, 'down': 1, 'pulls_total': 2},
            'wanted': [{'class': c, 'spec': k, 'spec_name': n, 'note': note} for c, k, n, note in gd.RECRUIT_WANTED],
            'officers': officers, 'updated_at': '2026-12-01T18:00:00Z'}


def _applicant(name: str, cls: str, spec: str, ilvl: int | None, parse: tuple | None) -> dict:
    return {'character_key': _char_key(name), 'name': name, 'class': cls, 'spec': f'{cls}-{spec.lower().replace(" ", "-")}', 'spec_name': f'{spec} {cls.title()}',
            'level': 60, 'item_level': ilvl, 'best_parse': None if parse is None else _parse(*parse)}


def officer_applications() -> dict:
    rows = [{'id': i, 'guild_id': gd.GUILD_ID, 'status': 'pending', 'created_at': sent, 'decided_at': None, 'message': msg, 'availability': avail, 'discord_dm': False,
             'character': _applicant(name, cls, spec, il, parse), 'votes': {'yes': y, 'no': n, 'voters': gd.VOTERS, 'mine': mine},
             'decided_by': None, 'role_outcome': None, 'dm_outcome': None}
            for i, name, cls, spec, il, parse, msg, avail, y, n, mine, sent in gd.APPLICANTS]
    rows += [{'id': i, 'guild_id': gd.GUILD_ID, 'status': 'accepted', 'created_at': '2026-11-28T19:00:00Z', 'decided_at': '2026-11-30T19:00:00Z', 'message': '', 'availability': '',
              'discord_dm': False, 'character': _applicant(name, cls, spec, None, None), 'votes': {'yes': y, 'no': 0, 'voters': gd.VOTERS, 'mine': 'yes'},
              'decided_by': _ref(by), 'role_outcome': role, 'dm_outcome': dm}
             for i, name, cls, spec, y, by, role, dm in gd.DECIDED]
    return {'rows': rows, 'pending': len(gd.APPLICANTS), 'discord_role_id': gd.DISCORD_ROLE_ID}


def member_profile() -> dict:
    name, cls, spec, role, ilvl, _state = gd.R[_index(gd.PROFILE)]
    member = {'character_key': _char_key(name), 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': name, 'class': cls, 'spec': spec, 'role': 'dps', 'rank': 'member',
              'verified': True, 'logged_recently': True, 'item_level': ilvl, 'level': 60, 'is_main': True, 'consent': 'gear_bags', 'account_key': gd.PROFILE_ACCOUNT,
              'best_parse': _parse(*gd.PROFILE_PARSE)}
    readiness = {'character_key': _char_key(name), 'name': name, 'class': cls, 'spec': spec, 'consent': 'gear_bags', 'level': 60, 'gear_gap': gd.PROFILE_GAP,
                 'enchants': {'missing_slots': [], 'checked': True}, 'consumables': {'state': 'stocked'}, 'talent_points_unspent': 0, 'talent_points_read': True,
                 'item_level': ilvl, 'item_level_delta': 0, 'logged_at': gd.NOW, 'failing': 0}
    present = sum(1 for _, o in gd.PROFILE_NIGHTS if o == 'present')
    nights = [{'raid_id': 9 + i, 'starts_at': at, 'outcome': o, 'excused': False} for i, (at, o) in enumerate(gd.PROFILE_NIGHTS)]
    loot = [{'item_id': i, 'name': n, 'icon': ic, 'quality': q, 'raid_id': 12, 'awarded_at': at} for i, n, ic, q, at in gd.PROFILE_LOOT]
    notes = [{'id': i, 'body': body, 'author': _ref(who), 'created_at': at, 'may_delete': mine} for i, body, who, at, mine in gd.PROFILE_NOTES]
    return {'account_key': gd.PROFILE_ACCOUNT, 'member': member, 'characters': [{'character_key': _char_key(name), 'name': name, 'class': cls, 'is_main': True}],
            'attendance': {'source': 'raids', 'last_4': {'present': present, 'nights': len(nights)},
                           'last_30': {'present': gd.PROFILE_LAST_30[0], 'nights': gd.PROFILE_LAST_30[1]}, 'nights': nights},
            'loot': {'count': len(loot), 'items': loot}, 'readiness': readiness, 'notes': notes}


def recruit_guild_page() -> dict:
    return {'guild': _guild(), 'progression': [], 'roster_best': [], 'reports': []}


def officer_me() -> dict:
    name, cls, spec, *_ = gd.R[_index(gd.VIEWER)]
    guild = {'id': gd.GUILD_ID, 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': gd.GUILD, 'slug': gd.SLUG, 'faction': gd.FACTION, 'rank': 'officer',
             'verified': True, 'verified_count': gd.RAIDERS, 'can': ['approve', 'remove', 'settings', 'award']}
    character = {'key': _char_key(name), 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': name, 'class': cls.title(), 'spec': spec, 'level': 60,
                 'faction': gd.FACTION, 'guild': {'id': gd.GUILD_ID, 'name': gd.GUILD, 'verified': True}}
    return {'user': {'id': 5, 'battletag': 'Dunmaro#1105', 'email': None, 'role': 'user', 'anonymize': False}, 'characters': [character], 'guilds': [guild]}


def recruit_fixtures() -> list[Path]:
    return [write('recruit-card.json', recruit_card()), write('recruit-applications.json', officer_applications()), write('recruit-profile.json', member_profile()),
            write('recruit-guild-page.json', recruit_guild_page()), write('recruit-me-officer.json', officer_me())]


WRITERS.append(recruit_fixtures)

# The GUILD-ALL lane's page groups, one fixtures module each.
import guild_directory_fixtures  # noqa: E402
import guild_raids_fixtures  # noqa: E402
import guild_settings_fixtures  # noqa: E402
import guild_states_fixtures  # noqa: E402
import guild_tabs_fixtures  # noqa: E402
WRITERS += [lambda m=m: m.writer(write) for m in (guild_states_fixtures, guild_tabs_fixtures, guild_raids_fixtures, guild_settings_fixtures, guild_directory_fixtures)]

if __name__ == '__main__':
    for writer in WRITERS:
        for path in writer():
            print(path)
