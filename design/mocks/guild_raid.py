"""The next Onyxia's Lair (Wed 16 Dec) as the API answers it (Go guilds.RaidView): the one raid view the raid planner mock
(ws_guild2.raid_planner) draws and guild_fixtures.py writes for the build (web/tests/e2e/fixtures/guild-parity/raid-*.json).

Built from guild_data.R (the sign-ups), the API's own Onyxia assignment template (api/internal/guilds/raid_assign_templates.go)
and buff-cover catalogue (data/curated/raid-cover.json, icons from its Go twin), and the zone's default comp target
(api/internal/guilds/raids_comp_target.go). Nothing here is typed twice: the mock and the fixtures both call raid_view().
"""
import json
import math
import re
from pathlib import Path

import guild_data as gd
from guild_data import GUILD, GUILD_ID, R, REGION, RULESET, SLUG, ZONE

ROOT = Path(__file__).resolve().parents[2]
GUILDS_GO = ROOT / 'api/internal/guilds'

RAID_ID = gd.NEXT_RAID_ID
RAID_ZONE_KEY = 'raid:onyxias-lair'
RAID_ZONE_NAME = "Onyxia's Lair"
ONYXIA_ENCOUNTER = 1084
# Wed 16 Dec 2026, 8:00 pm CST, three hours.
RAID_STARTS = gd.NEXT_RAID_AT
RAID_MINUTES = 180
RAID_SIZE = 40
# The raid post went out the day after the first night; nobody has pinged the non-signers yet.
POSTED_AT = gd.NEXT_POSTED_AT
VIEWER = gd.VIEWER
# Non-signers whose account has no Discord linked.
NO_DISCORD = ('Skyarcana',)
NOTES_MD = 'Fire resistance on the tanks. Flask and food for everyone.'
NOTES_HTML = '<p>Fire resistance on the tanks. Flask and food for everyone.</p>'
# The raid leader's picks so far, by the template's slot key (manual: officers name every one).
PICKS = {
    'main_tank': ['Thornwicke'],
    'main_tank_heals': ['Ironbrand', 'Hollowmend', 'Stonetide'],
    'raid_heals': ['Valomira', 'Dawnweaver', 'Mossveil'],
    'fear_cover': ['Hollowmend'],
    'whelps_west': ['Dunmaro', 'Skytalon'],
    'breath_callers': ['Frostquill'],
    'soulstones': ['Doomwhisper', 'Hexmarrow'],
}

GROUP = {'Tank': 'tank', 'Healer': 'healer', 'Melee': 'melee', 'Ranged': 'ranged'}
ROLE = {'tank': 'tank', 'healer': 'healer', 'melee': 'dps', 'ranged': 'dps'}
SIGNED = ('in', 'late', 'bench')


def key_of(name: str) -> str:
    return f'{REGION}/{RULESET}/{name.lower()}'


def spec_key(cls: str, spec: str) -> str:
    return f'{cls}-{spec.lower().replace(" ", "-")}'


def member(name: str) -> dict:
    """A roster row of guild_data.R as the API's comp fields name it."""
    _, cls, spec, role, ilvl, state = next(r for r in R if r[0] == name)
    group = GROUP[role]
    return {'character_key': key_of(name), 'name': name, 'class': cls, 'spec': spec, 'role': ROLE[group], 'group': group,
            'spec_key': spec_key(cls, spec), 'ilvl': ilvl, 'state': state}


MEMBERS = [member(r[0]) for r in R]


def _signup(m: dict) -> dict:
    return {'character_key': m['character_key'], 'name': m['name'], 'class': m['class'], 'spec': m['spec'], 'role': m['role'],
            'state': m['state'], 'note': '', 'own': m['name'] == VIEWER, 'group': m['group'], 'spec_key': m['spec_key'], 'also': []}


def _comp(m: dict) -> dict:
    return {'character_key': m['character_key'], 'name': m['name'], 'class': m['class'], 'spec': m['spec'], 'group': m['group'],
            'spec_key': m['spec_key']}


def counts() -> dict:
    by_state = {s: [m for m in MEMBERS if m['state'] == s] for s in ('in', 'late', 'bench')}
    on = by_state['in']
    per_group = {g: sum(1 for m in on if m['group'] == g) for g in GROUP.values()}
    late_groups = {f'{g}s' if g in ('tank', 'healer') else g: sum(1 for m in by_state['late'] if m['group'] == g) for g in GROUP.values()}
    classes = sorted({m['class'] for m in MEMBERS})
    return {'in': len(on), 'late': len(by_state['late']), 'bench': len(by_state['bench']), 'out': 0,
            'tanks': per_group['tank'], 'healers': per_group['healer'], 'dps': per_group['melee'] + per_group['ranged'],
            'melee': per_group['melee'], 'ranged': per_group['ranged'],
            'classes': {c: sum(1 for m in on + by_state['late'] if m['class'] == c) for c in classes}, 'late_groups': late_groups}


def comp_target() -> dict:
    """The zone's default seats (raids_comp_target.go defaultCompTargets)."""
    src = (GUILDS_GO / 'raids_comp_target.go').read_text()
    found = re.search(r'"' + re.escape(RAID_ZONE_KEY) + r'":\s*\{Tanks: (\d+), Healers: (\d+), Melee: (\d+), Ranged: (\d+)\}', src)
    if found is None:
        raise ValueError(f'no default comp target for {RAID_ZONE_KEY} in raids_comp_target.go')
    tanks, healers, melee, ranged = (int(x) for x in found.groups())
    return {'tanks': tanks, 'healers': healers, 'melee': melee, 'ranged': ranged, 'source': 'default'}


# ------------------------------------------------------------------------------------------------ buff and debuff cover
def _cover_icons() -> dict:
    src = (GUILDS_GO / 'raids_cover_catalogue.go').read_text()
    return dict(re.findall(r'Key: "([a-z_]+)", .*?Icon: "([a-z0-9_]+)"', src))


def _provides(entry: dict, m: dict) -> str:
    """raids_cover.go provides: the class, the spec and the roster group when named, then a talent the sign-up cannot
    show is likely."""
    best = ''
    for p in entry['providers']:
        if p['class'] != m['class'] or ('spec' in p and p['spec'] != m['spec_key']) or ('group' in p and p['group'] != m['group']):
            continue
        best = 'sure' if 'talent' not in p else (best or 'likely')
    return best


def cover() -> list[dict]:
    """raids_cover.go computeCover over the sign-ups: pooled rows draw on the same players in catalogue order."""
    catalogue = json.loads((ROOT / 'data/curated/raid-cover.json').read_text())['entries']
    icons = _cover_icons()
    signed = [m for m in MEMBERS if m['state'] in SIGNED]
    melee = sum(1 for m in signed if m['state'] in ('in', 'late') and m['group'] in ('tank', 'melee'))
    caster = sum(1 for m in signed if m['state'] in ('in', 'late') and m['group'] in ('ranged', 'healer'))
    pools: dict = {}
    rows = []
    for e in catalogue:
        per = e['need']['per']
        people = {'melee_group': melee, 'caster_group': caster}.get(per)
        need = e['need']['count'] * (1 if people is None else max(1, math.ceil(people / 5)))
        tally = {s: {'sure': 0, 'likely': 0} for s in SIGNED}
        for m in signed:
            how = _provides(e, m)
            if how:
                tally[m['state']][how] += 1
        if e.get('pool'):
            left = pools.setdefault(e['pool'], tally)
            got = {}
            want = need
            for s in SIGNED:
                sure = min(left[s]['sure'], want)
                want -= sure
                likely = min(left[s]['likely'], want)
                want -= likely
                left[s] = {'sure': left[s]['sure'] - sure, 'likely': left[s]['likely'] - likely}
                got[s] = {'sure': sure, 'likely': likely}
            tally = got
        total = {s: tally[s]['sure'] + tally[s]['likely'] for s in SIGNED}
        on = total['in'] + total['late']
        rows.append({'key': e['key'], 'name': e['name'], 'kind': e['kind'], 'spell_id': e['spell_id'], 'icon': icons[e['key']],
                     'need': need, 'per': per, 'providers_in': total['in'], 'providers_late': total['late'],
                     'providers_bench': total['bench'], 'likely': tally['in']['likely'] + tally['late']['likely'],
                     'state': 'ok' if on >= need else ('short' if on > 0 else 'missing')})
    return rows


# ------------------------------------------------------------------------------------------------ assignments
def onyxia_template() -> list[dict]:
    """Onyxia's slots from raid_assign_templates.go, in the call's order."""
    src = (GUILDS_GO / 'raid_assign_templates.go').read_text()
    block = src.split(f'{ONYXIA_ENCOUNTER}: {{', 1)[1].split('\n\t},', 1)[0]
    slots = []
    for found in re.finditer(r'\{Key: "(\w+)", Label: "([^"]+)", Warning: "([^"]+)", Count: (\d+), Role: slotRole(\w+),\s*Purpose: "([^"]+)"\}', block):
        key, label, warning, count, role, purpose = found.groups()
        slots.append({'key': key, 'label': label, 'warning': warning, 'count': int(count), 'role': role.lower(), 'purpose': purpose})
    if len(slots) < 12:
        raise ValueError(f'raid_assign_templates.go: read {len(slots)} Onyxia slots, expected the 12-slot template')
    return slots


def _raider(m: dict) -> dict:
    return {'character_key': m['character_key'], 'name': m['name'], 'class': m['class'], 'group': m['group'], 'state': m['state'],
            'own': m['name'] == VIEWER}


def assignments() -> dict:
    by_name = {m['name']: m for m in MEMBERS}
    unknown = [n for names in PICKS.values() for n in names if by_name.get(n, {}).get('state') not in SIGNED]
    if unknown:
        raise ValueError(f'assignment picks must be signed In, Late or Bench: {unknown}')
    slots, warning = [], []
    for t in onyxia_template():
        names = [{**_raider(by_name[n]), 'off_role': t['role'] != 'any' and t['role'] != by_name[n]['group']} for n in PICKS.get(t['key'], [])]
        slots.append({'key': t['key'], 'label': t['label'], 'purpose': t['purpose'], 'count': t['count'], 'role': t['role'], 'names': names})
        if names:
            warning.append(f'{t["warning"]}: {", ".join(n["name"] for n in names)}')
    encounter = {'encounter_id': ONYXIA_ENCOUNTER, 'name': 'Onyxia', 'raid_zone': RAID_ZONE_NAME, 'slots': slots, 'raid_warning': warning}
    return {'encounters': [encounter], 'not_in_client': [], 'candidates': [_raider(m) for m in MEMBERS if m['state'] in SIGNED],
            'posted_at': None}


# ------------------------------------------------------------------------------------------------ the raid view
def raid_view() -> dict:
    not_signed = [m for m in MEMBERS if m['state'] == 'none']
    viewer = member(VIEWER)
    return {
        'id': RAID_ID, 'zone': RAID_ZONE_KEY, 'zone_name': RAID_ZONE_NAME, 'zone_kind': 'raid', 'key_item': None, 'title': RAID_ZONE_NAME,
        'starts_at': RAID_STARTS, 'duration_minutes': RAID_MINUTES, 'status': 'planned', 'size': RAID_SIZE, 'guild_timezone': ZONE,
        'notes_html': NOTES_HTML, 'notes': NOTES_MD, 'counts': counts(),
        'signups': [_signup(m) for m in MEMBERS if m['state'] in SIGNED],
        'not_signed': [_comp(m) for m in not_signed],
        'viewer': {'characters': [{'character_key': viewer['character_key'], 'name': VIEWER, 'class': viewer['class'], 'spec': viewer['spec'],
                                   'verified': True, 'state': viewer['state'], 'role': viewer['role'], 'raid_spec': viewer['spec_key'],
                                   'spec_key': viewer['spec_key'], 'default_spec': viewer['spec_key']}]},
        'discord': {'available': True, 'connected': True, 'post_default': True, 'remind_default': True, 'channel_name': 'raids'},
        'discord_post': {'posted_at': POSTED_AT, 'post_at': None, 'reminded_at': None, 'remind': True},
        'cover': cover(), 'comp_target': comp_target(),
        'not_signed_raiders': [{'character_key': m['character_key'], 'name': m['name'], 'class': m['class'], 'account_key': f'account/{m["name"].lower()}',
                                'discord_linked': m['name'] not in NO_DISCORD} for m in not_signed],
        'ping': {'last_pinged_at': None, 'next_allowed_at': None},
        'assignments': assignments(),
        'pre_done_status': None, 'match': None, 'night': None, 'attendance': None, 'report': None,
    }


# ------------------------------------------------------------------------------------------------ the session around it
def me() -> dict:
    """/v1/me for the officer viewing the planner: one character, the guild's officer row."""
    v = member(VIEWER)
    character = {'key': v['character_key'], 'region': REGION, 'ruleset': RULESET, 'name': VIEWER, 'class': v['class'].title(), 'spec': v['spec'],
                 'race': 'Dwarf', 'realm': 'Skyborne', 'level': 60, 'faction': 'alliance', 'item_level': v['ilvl'], 'source': 'bnet',
                 'guild': {'id': GUILD_ID, 'name': GUILD, 'rank': 'Officer', 'verified': True}}
    guild = {'id': GUILD_ID, 'region': REGION, 'ruleset': RULESET, 'name': GUILD, 'slug': SLUG, 'rank': 'Officer', 'verified': True,
             'verified_count': len(R), 'faction': 'alliance', 'can': ['approve', 'remove', 'settings', 'award']}
    return {'user': {'id': 7, 'battletag': 'Hunthrop#1894', 'email': None, 'role': 'user', 'anonymize': False},
            'characters': [character], 'guilds': [guild], 'main_character_key': v['character_key']}


def guild_page() -> dict:
    """The public guild read (/v1/guilds/{region}/{ruleset}/{slug}): the planner reads the guild's id, name and faction."""
    return {'guild': {'id': GUILD_ID, 'name': GUILD, 'region': REGION, 'ruleset': RULESET, 'faction': 'alliance', 'crest_url': None},
            'progression': [], 'roster_best': [], 'reports': []}


def guild_home() -> dict:
    """/v1/guilds/{id}/home as the planner reads it: the viewer's role, officer tools and the claim."""
    return {'guild': {'id': GUILD_ID, 'name': GUILD, 'region': REGION, 'ruleset': RULESET, 'faction': 'alliance', 'crest_url': None},
            'viewer': {'raid_guild_tools': True, 'role': 'officer', 'character_key': key_of(VIEWER), 'verified': True, 'officer_tools': True},
            'claim': {'state': 'claimed', 'since': gd.CLAIM_SINCE, 'frozen': False, 'yours': True,
                      'claimed_by': {'name': VIEWER, 'class': member(VIEWER)['class'], 'character_key': key_of(VIEWER)}}}
