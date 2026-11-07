"""The guild Overview and loot council's API answers (lane A of the guild parity states), in the shapes web/src/lib/guild/api.ts,
overview-api.ts, loot-council-api.ts and recruit-api.ts declare, built from guild_data.py and guild_overview_data.py.
guild_fixtures.py writes them into web/tests/e2e/fixtures/guild-parity; web/tests/e2e/support/guild-parity/overview.ts routes them.
"""
import json
import math
from datetime import datetime, timedelta, timezone
from pathlib import Path

import guild_overview_data as D
from guild_data import ALLIANCE, GUILD, GUILD_ID, NOW, R, REGION, RULESET, SLUG, ZONE  # noqa: F401

ROOT = Path(__file__).resolve().parents[2]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
ROLE = {'Tank': 'tank', 'Healer': 'healer', 'Melee': 'dps', 'Ranged': 'dps'}
GROUP = {'Tank': 'tank', 'Healer': 'healer', 'Melee': 'melee', 'Ranged': 'ranged'}
FIRST_TIER = ['Barrow Deeps', 'Hyjal Summit', D.ZONE_NAME]
GEAR_CLASSES = {2, 4}
KIND_BY_CLASS = {1: 'bag', 7: 'skinned', 12: 'quest', 15: 'quest'}


def guild() -> dict:
    return {'id': GUILD_ID, 'name': GUILD, 'region': REGION, 'ruleset': RULESET, 'faction': 'alliance', 'crest_url': None}


def iso_minus(at: str, seconds: int) -> str:
    stamp = datetime.fromisoformat(at.replace('Z', '+00:00')) - timedelta(seconds=seconds)
    return stamp.astimezone(timezone.utc).strftime('%Y-%m-%dT%H:%M:%SZ')


def attendance(name: str) -> dict:
    cells = next((c for n, _, c in D.RECENT_ROWS if n == name), None)
    nights = len(D.RECENT_RAIDS)
    if cells is not None:
        return {'present': cells.count('present'), 'nights': nights}
    return {'present': nights - (1 if name in D.LAST_ABSENT or name == D.MEMBER else 0), 'nights': nights}


def tier() -> dict:
    counts = D.last_raid_counts()
    return {'name': 'First tier', 'raids': FIRST_TIER, 'named_encounters': 1, 'down': 1 if counts['kills'] else 0, 'pulls_total': counts['pulls']}


# ------------------------------------------------------------------ the guild page, the session and the home
def last_report() -> dict:
    counts = D.last_raid_counts()
    return {'id': D.LAST_REPORT_ID, 'title': f'{D.ZONE_NAME} · raid night', 'zone': D.ZONE_NAME, 'created_at': D.LAST_RAID_AT,
            'fight_count': counts['pulls'], 'kill_count': counts['kills']}


def guild_page() -> dict:
    counts = D.last_raid_counts()
    report = {k: v for k, v in last_report().items() if k in ('id', 'title', 'zone', 'created_at')}
    return {'guild': guild(), 'roster_best': [], 'reports': [report],
            'progression': [{'encounter': D.BOSS, 'encounter_id': D.ENCOUNTER_ID, 'difficulty': 1, 'kills': counts['kills'],
                             'pull_count': counts['wipes'] + counts['kills'], 'first_kill_at': D.LAST_RAID_AT}]}


def officer_can(officer: bool) -> list[str]:
    return ['approve', 'remove', 'settings', 'award'] if officer else []


def me(name: str, officer: bool) -> dict:
    _, cls, spec, _, _, _ = D.row(name)
    return {
        'user': {'id': 7, 'battletag': 'Fixture#1', 'email': None, 'role': 'user', 'anonymize': False},
        'characters': [{'key': D.key_of(name), 'region': REGION, 'ruleset': RULESET, 'name': name, 'class': cls.title(), 'spec': spec,
                        'level': 60, 'realm': 'Skyborne', 'faction': 'alliance', 'guild': {'id': GUILD_ID, 'name': GUILD, 'verified': True}}],
        'guilds': [{'id': GUILD_ID, 'region': REGION, 'ruleset': RULESET, 'name': GUILD, 'rank': 'officer' if officer else 'member',
                    'consent': 'gear_bags', 'verified': True, 'can': officer_can(officer)}],
    }


def roster_row(r: tuple) -> dict:
    name, cls, spec, role, ilvl, _ = r
    return {'character_key': D.key_of(name), 'region': REGION, 'ruleset': RULESET, 'name': name, 'class': cls, 'spec': spec,
            'role': ROLE[role], 'rank': 'officer' if name == D.OFFICER else 'member', 'verified': True, 'logged_recently': True,
            'logged_at': D.UPDATED_AT, 'item_level': ilvl, 'level': 60, 'consent': 'gear_bags', 'attendance': attendance(name),
            'best_parse': None, 'rating': None, 'account_key': f'acct-{name.lower()}', 'may_remove': False, 'addon': None}


def next_raid(viewer: str) -> dict:
    groups = D.in_by_group()
    _, _, _, role, _, state = D.row(viewer)
    return {'id': D.NEXT_RAID_ID, 'zone': D.ZONE_KEY, 'zone_name': D.ZONE_NAME, 'zone_kind': 'raid', 'title': D.ZONE_NAME,
            'starts_at': D.NEXT_RAID_AT, 'status': 'planned', 'size': 40, 'guild_timezone': ZONE,
            'counts': {'in': len(D.signed('in')), 'late': len(D.signed('late')), 'tanks': groups['Tank'], 'healers': groups['Healer'],
                       'dps': groups['Melee'] + groups['Ranged'], 'melee': groups['Melee'], 'ranged': groups['Ranged']},
            'own': {'character_key': D.key_of(viewer), 'state': state, 'role': ROLE[role], 'note': ''}}


def home(viewer: str, officer: bool) -> dict:
    return {
        'guild': guild(),
        'viewer': {'role': 'officer' if officer else 'member', 'character_key': D.key_of(viewer), 'verified': True, 'officer_tools': True,
                   'raid_guild_tools': True, 'verify': {'has_approver': True}},
        'claim': {'state': 'claimed', 'since': '2026-11-12T00:00:00Z', 'frozen': False, 'yours': officer,
                  'claimed_by': {'name': D.OFFICER, 'class': 'hunter', 'character_key': D.key_of(D.OFFICER)}},
        'summary': {'raiders': len(R), 'waiting_for_approval': len(D.APPROVALS), 'below_rating_floor': 0, 'tier': tier(),
                    'named_encounters_down': tier()['down'], 'pulls_this_tier': tier()['pulls_total'], 'updated_at': D.UPDATED_AT},
        'standing': None,
        'reports': [last_report()],
        'roster': [roster_row(r) for r in R],
        'pending': [],
        'addon_snapshot': None,
        'next_raid': next_raid(viewer),
        'attendance_source': 'raids',
    }


# ------------------------------------------------------------------ the depth reads the Overview makes
def raids_page() -> dict:
    fights = [{'index': i, 'name': name, 'encounter_id': D.ENCOUNTER_ID if name else None, 'kill': kill, 'duration_ms': secs * 1000,
               'deaths': deaths, 'players': len(R) - len(D.LAST_ABSENT)} for i, (name, kill, secs, deaths) in enumerate(D.LAST_PULLS)]
    counts = D.last_raid_counts()
    night = {**last_report(), 'duration_ms': sum(f['duration_ms'] for f in fights), 'wipe_count': counts['wipes'],
             'raiders': len(R) - len(D.LAST_ABSENT), 'deaths': sum(f['deaths'] for f in fights), 'top_hps': None, 'top_dps': None,
             'fights': fights, 'present': [], 'raid_id': D.LAST_RAID_ID}
    return {'rows': [night], 'next_cursor': None}


def progression() -> dict:
    counts = D.last_raid_counts()
    boss_pulls = counts['kills'] + counts['wipes']
    kill_ms = next(secs for name, kill, secs, _ in D.LAST_PULLS if name and kill) * 1000
    deaths = sum(d for name, _, _, d in D.LAST_PULLS if name)
    return {'tier': tier(), 'unnamed': [{'zone': D.ZONE_NAME, 'pulls': counts['trash'], 'nights': 1}],
            'encounters': [{'encounter_id': D.ENCOUNTER_ID, 'name': D.BOSS, 'zone': D.ZONE_NAME, 'first_kill_at': D.LAST_RAID_AT,
                            'pulls': boss_pulls, 'kills': counts['kills'], 'best_kill_ms': kill_ms,
                            'pulls_by_night': [{'report_id': D.LAST_REPORT_ID, 'date': D.LAST_RAID_DATE, 'pulls': boss_pulls, 'killed': True}],
                            'deaths_per_pull': deaths / boss_pulls, 'best_by_role': {}}]}


def cover_rows() -> list[dict]:
    """The buff and debuff cover the API computes from the curated catalogue over the In, Late and Bench sign-ups."""
    catalogue = json.loads((ROOT / 'data/curated/raid-cover.json').read_text())['entries']
    icons = json.loads((ROOT / f'data/builds/{BUILD}/spellicons.json').read_text())
    signed_in_late = [r for r in R if r[5] in ('in', 'late')]
    melee = sum(1 for r in signed_in_late if r[3] in ('Tank', 'Melee'))
    casters = len(signed_in_late) - melee
    out = []
    for entry in catalogue:
        classes = {p['class'] for p in entry['providers']}
        specs = {p['spec'] for p in entry['providers'] if 'spec' in p}
        talent = any('talent' in p for p in entry['providers'])

        def gives(r: tuple) -> bool:
            return r[1] in classes and (not specs or D.spec_key(r[1], r[2]) in specs)

        per = entry['need']['per']
        base = entry['need']['count']
        need = base if per == 'raid' else max(base, math.ceil((melee if per == 'melee_group' else casters) / 5))
        providers = {state: sum(1 for r in R if r[5] == state and gives(r)) for state in ('in', 'late', 'bench')}
        likely = providers['in'] + providers['late'] if talent else 0
        signed = providers['in'] + providers['late']
        state = 'missing' if signed == 0 else ('ok' if signed - likely >= need else 'short')
        icon = icons.get(str(entry.get('spell', entry['spell_id']))) or icons.get(str(entry['spell_id'])) or ''
        out.append({'key': entry['key'], 'name': entry['name'], 'kind': entry['kind'], 'spell_id': entry['spell_id'], 'icon': icon,
                    'need': need, 'per': per, 'providers_in': providers['in'], 'providers_late': providers['late'],
                    'providers_bench': providers['bench'], 'likely': likely, 'state': state})
    return out


def signup(r: tuple, viewer: str) -> dict:
    name, cls, spec, role, _, state = r
    return {'character_key': D.key_of(name), 'name': name, 'class': cls, 'spec': spec, 'role': ROLE[role], 'state': state, 'note': '',
            'own': name == viewer, 'group': GROUP[role], 'spec_key': D.spec_key(cls, spec), 'also': []}


def raid_view(viewer: str, officer: bool) -> dict:
    groups = D.in_by_group()
    late = [r for r in R if r[5] == 'late']
    late_groups = {'tanks': 0, 'healers': 0, 'melee': 0, 'ranged': 0}
    for r in late:
        late_groups[{'Tank': 'tanks', 'Healer': 'healers', 'Melee': 'melee', 'Ranged': 'ranged'}[r[3]]] += 1
    _, vcls, vspec, vrole, _, vstate = D.row(viewer)
    none = [r for r in R if r[5] == 'none']
    view = {
        'id': D.NEXT_RAID_ID, 'zone': D.ZONE_KEY, 'zone_name': D.ZONE_NAME, 'zone_kind': 'raid', 'key_item': None, 'title': D.ZONE_NAME,
        'starts_at': D.NEXT_RAID_AT, 'duration_minutes': 180, 'status': 'planned', 'size': 40, 'guild_timezone': ZONE,
        'notes_html': '', 'notes': '',
        'counts': {'in': len(D.signed('in')), 'late': len(late), 'bench': len(D.signed('bench')), 'out': 0, 'tanks': groups['Tank'],
                   'healers': groups['Healer'], 'dps': groups['Melee'] + groups['Ranged'], 'melee': groups['Melee'], 'ranged': groups['Ranged'],
                   'late_groups': late_groups},
        'signups': [signup(r, viewer) for r in R if r[5] in ('in', 'late', 'bench')],
        'not_signed': [{'character_key': D.key_of(r[0]), 'name': r[0], 'class': r[1], 'spec': r[2], 'group': GROUP[r[3]],
                        'spec_key': D.spec_key(r[1], r[2])} for r in none],
        'viewer': {'characters': [{'character_key': D.key_of(viewer), 'name': viewer, 'class': vcls, 'spec': vspec, 'verified': True,
                                   'state': vstate, 'role': ROLE[vrole], 'raid_spec': None, 'spec_key': D.spec_key(vcls, vspec),
                                   'default_spec': D.spec_key(vcls, vspec)}]},
        'discord': {'available': True, 'connected': True, 'post_default': True, 'remind_default': True, 'channel_name': 'raids'},
        'discord_post': {'posted_at': '2026-12-06T18:00:00Z', 'post_at': None, 'reminded_at': None, 'remind': True},
        'cover': cover_rows(),
        'comp_target': {**D.COMP_TARGET, 'source': 'default'},
        'assignments': None,
        'pre_done_status': None, 'match': None, 'night': None, 'attendance': None, 'report': None,
    }
    if officer:
        view['not_signed_raiders'] = [{'character_key': D.key_of(r[0]), 'name': r[0], 'class': r[1], 'account_key': f'acct-{r[0].lower()}',
                                       'discord_linked': True} for r in none]
        view['ping'] = {'last_pinged_at': None, 'next_allowed_at': None}
    return view


def checklist() -> dict:
    _, cls, spec, _, _, _ = D.row(D.MEMBER)
    return {
        'character': {'character_key': D.key_of(D.MEMBER), 'name': D.MEMBER, 'class': cls, 'spec': spec, 'level': 60},
        'raid': {'id': D.NEXT_RAID_ID, 'zone_key': D.ZONE_KEY, 'zone_name': D.ZONE_NAME},
        'verdict': 'almost',
        'items': [
            {'key': 'level', 'state': 'ok', 'level': 60, 'required_level': 60},
            {'key': 'attunement', 'state': 'ok', 'attunement_kind': 'quest', 'key_item': D.ATTUNEMENT_ITEM},
            {'key': 'gear_gap', 'state': 'fail', 'upgrades': len(D.GEAR_SLOTS), 'slots': D.GEAR_SLOTS, 'unit': 'hps'},
            {'key': 'enchants', 'state': 'ok', 'missing_slots': []},
            {'key': 'flask_food', 'state': 'fail', 'supplies': [], 'missing': D.FLASK_MISSING},
            {'key': 'fire_resistance', 'state': 'not_tracked', 'reason': 'resistance_not_in_export'},
        ],
    }


def officer_actions() -> dict:
    none = [r[0] for r in R if r[5] == 'none']
    open_drops = [d for d in D.DROPS if d[5] is None]
    return {'approvals': {'count': len(D.APPROVALS), 'names': D.APPROVALS},
            'applications': {'count': len(D.APPLICATIONS), 'names': D.APPLICATIONS},
            'not_signed': {'raid_id': D.NEXT_RAID_ID, 'count': len(none), 'names': none},
            'drops_to_award': {'tracked': True, 'count': len(open_drops), 'raid_id': D.LAST_RAID_ID, 'label': f'{D.BOSS}, Mon 7 Dec',
                               'awarded': len(D.DROPS) - len(open_drops)}}


def recent_attendance() -> dict:
    return {'raids': [{'id': rid, 'starts_at': at, 'zone_name': D.ZONE_NAME} for rid, at in D.RECENT_RAIDS],
            'rows': [{'name': n, 'class': D.row(n)[1], 'account_key': f'acct-{n.lower()}', 'cells': cells, 'slipping': slipping}
                     for n, slipping, cells in D.RECENT_ROWS]}


def readiness_row(r: tuple) -> dict:
    name, cls, spec, _, ilvl, _ = r
    return {'character_key': D.key_of(name), 'name': name, 'class': cls, 'spec': spec, 'consent': 'gear_bags', 'level': 60,
            'gear_gap': {'upgrades': 0, 'gain_dps': None, 'not_sim_checked': 0, 'unit': None}, 'enchants': {'missing_slots': [], 'checked': True},
            'consumables': {'state': 'short' if name in D.NO_FLASK_FOOD else 'stocked'}, 'talent_points_unspent': 0, 'talent_points_read': True,
            'item_level': ilvl, 'item_level_delta': 0, 'logged_at': D.UPDATED_AT, 'failing': 0, 'no_export': False, 'synced_before_cutoff': True,
            'attuned': name not in D.NOT_ATTUNED,
            'resistances': {'fire': D.TANK_FIRE[name], 'nature': 0, 'frost': 0, 'shadow': 0, 'arcane': 0} if name in D.TANK_FIRE else None,
            'flask_food': {'flask_or_elixir': name not in D.NO_FLASK_FOOD, 'food': name not in D.NO_FLASK_FOOD}}


def readiness(for_raid: bool) -> dict:
    rows = [readiness_row(r) for r in R if (r[5] in ('in', 'late') if for_raid else True)]
    page = {'median_item_level': sorted(r[4] for r in R)[len(R) // 2], 'generated_at': D.UPDATED_AT, 'rows': rows}
    if for_raid:
        page['raid'] = {'id': D.NEXT_RAID_ID, 'title': D.ZONE_NAME, 'starts_at': D.NEXT_RAID_AT, 'cutoff_at': iso_minus(D.NEXT_RAID_AT, 3600),
                        'attunement': True, 'fire_resistance': True}
    else:
        page['next_raid'] = {'id': D.NEXT_RAID_ID, 'title': D.ZONE_NAME, 'zone_name': D.ZONE_NAME, 'starts_at': D.NEXT_RAID_AT}
    return page


def recruit_card() -> dict:
    officers = [D.OFFICER, 'Ironbrand']
    return {'guild': {**guild(), 'path': f'/guild/{REGION}/{RULESET}/{SLUG}'}, 'members': len(R), 'open': True,
            'pitch': 'We raid two nights a week and split loot by council with the numbers shown.',
            'raid_nights': {'text': 'Thu and Mon, 8 to 11 pm', 'upcoming': [D.NEXT_RAID_AT], 'timezone': ZONE},
            'progress': tier(), 'wanted': [], 'updated_at': D.UPDATED_AT,
            'officers': [{'name': n, 'class': D.row(n)[1], 'spec': D.row(n)[2], 'rank': 'leader' if n == D.OFFICER else 'officer'} for n in officers]}


# ------------------------------------------------------------------ the loot council
def decision(award_id: int, drop: tuple) -> dict:
    item_id, name, icon, quality, _, (winner, by, yes) = drop
    return {'award_id': award_id, 'item_id': item_id, 'item_name': name, 'icon': icon, 'quality': quality, 'character_key': D.key_of(winner),
            'name': winner, 'class': D.row(winner)[1], 'ms_os': 'ms', 'reason': None, 'votes_yes': yes, 'votes_no': 0,
            'votes_officers': D.OFFICERS, 'awarded_by': by, 'at': iso_minus(D.LAST_RAID_AT, -3600 - award_id * 60)}


def decisions() -> list[dict]:
    return [decision(i + 1, drop) for i, drop in enumerate(d for d in D.DROPS if d[5] is not None)]


def loot_nights() -> dict:
    awards = {d['item_id']: d for d in decisions()}
    drops = [{'item_id': item_id, 'name': name, 'icon': icon, 'quality': quality, 'kind': kind, 'encounter_id': D.ENCOUNTER_ID, 'boss': D.BOSS,
              'state': 'open' if award is None else 'awarded', 'award': awards.get(item_id), 'recorded': True, 'count': 1, 'looted_by': []}
             for item_id, name, icon, quality, kind, award in D.DROPS]
    current = {'night': D.LOOT_NIGHT, 'raid_id': D.LAST_RAID_ID, 'report_id': D.LAST_REPORT_ID, 'date': D.LAST_RAID_DATE, 'at': D.LAST_RAID_AT,
               'zone': D.ZONE_NAME, 'current': True, 'drops': drops, 'decisions': decisions()}
    earlier_id, earlier_at = D.RECENT_RAIDS[-2]
    earlier = {'night': f'raid:{earlier_id}', 'raid_id': earlier_id, 'report_id': None, 'date': '2026-12-03', 'at': earlier_at,
               'zone': D.ZONE_NAME, 'current': False, 'drops': [], 'decisions': []}
    return {'officers': D.OFFICERS, 'nights': [current, earlier]}


def worn_item(worn: tuple) -> dict:
    item_id, name, icon, quality = worn
    return {'item_id': item_id, 'name': name, 'icon': icon, 'quality': quality}


def table_candidate(c: tuple) -> dict:
    name, gain, worn, _, attended, won, _, _ = c
    return {'character_key': D.key_of(name), 'name': name, 'class': D.row(name)[1], 'spec': D.row(name)[2], 'gain_dps': gain, 'gain_unit': 'dps',
            'tier': 0, 'not_sim_checked': False, 'ilvl_delta': None, 'attendance': {'present': attended, 'nights': len(D.RECENT_RAIDS)},
            'already_equivalent': False, 'fills_empty_slot': False, 'replaces': worn[1], 'replaces_item': worn_item(worn), 'items_won': won}


def council_drop() -> dict:
    item_id, name, icon, quality, kind, _ = next(d for d in D.DROPS if d[0] == D.SELECTED_DROP)
    cands = [{**table_candidate(c), 'bis_for_them': bis, 'ms_os': ms_os, 'ms_os_source': 'declared',
              'votes': {'yes': yes, 'no': 0, 'voters': D.OFFICERS, 'mine': None}}
             for c in D.CANDIDATES for (_, _, _, bis, _, _, ms_os, yes) in [c]]
    return {'night': D.LOOT_NIGHT, 'item_id': item_id, 'name': name, 'icon': icon, 'quality': quality, 'kind': kind, 'officers': D.OFFICERS,
            'closed': False, 'award': None, 'candidates': cands}


def loot_table() -> dict:
    """GET .../loot: the first tier's loot table, Onyxia's drops from the build's raid-bosses.json, each gear row ranked for the
    night's candidates and everyone at 60 eligible for the rest."""
    raids = json.loads((ROOT / f'data/builds/{BUILD}/raid-bosses.json').read_text())['raids']
    ony = next(r for r in raids if r['name'] == D.ZONE_NAME)
    boss = ony['bosses'][0]
    awards = {d['item_id']: d for d in decisions()}
    icons = {d[0]: d[2] for d in D.DROPS}
    eligible = [{'character_key': D.key_of(r[0]), 'name': r[0], 'class': r[1]} for r in R]
    items = []
    for it in boss['loot']:
        kind = 'gear' if it['class_id'] in GEAR_CLASSES else KIND_BY_CLASS.get(it['class_id'], 'quest')
        award = awards.get(it['item_id'])
        items.append({'item_id': it['item_id'], 'name': it['name'], 'icon': icons.get(it['item_id'], ''), 'quality': it['quality'], 'slot': '',
                      'inventory_type': it['inventory_type'], 'item_class': it['class_id'], 'item_subclass': it['subclass_id'], 'kind': kind,
                      'quest_label': '', 'eligible': [] if kind == 'gear' else [{**e, 'item_id': it['item_id']} for e in eligible],
                      'awarded_to': None if award is None else {'award_id': award['award_id'], 'character_key': award['character_key'],
                                                                'name': award['name'], 'at': award['at'], 'by_name': award['awarded_by']},
                      'candidates': [table_candidate(c) for c in D.CANDIDATES] if kind == 'gear' else [],
                      'copies': 1, 'encounter_id': D.ENCOUNTER_ID, 'boss': D.BOSS, 'raid': D.ZONE_NAME})
    return {'encounters': [{'encounter_id': D.ENCOUNTER_ID, 'name': D.BOSS, 'zone': D.ZONE_NAME, 'killed': True, 'kill_nights': 1, 'on_farm': False}],
            'selected': D.ENCOUNTER_ID, 'items': items, 'at_level_cap': len(R),
            'raids': [{'key': r['key'].removeprefix('raid:'), 'name': r['name'], 'in_client': r['status'] == 'in_client',
                       'encounters': [b['encounter_id'] for b in r['bosses']]} for r in raids]}


def overview_fixtures() -> dict[str, object]:
    """Every answer the member and officer Overviews and the officer's Loot tab read, by file name."""
    return {
        'overview-page.json': guild_page(),
        'overview-me-member.json': me(D.MEMBER, False),
        'overview-me-officer.json': me(D.OFFICER, True),
        'overview-home-member.json': home(D.MEMBER, False),
        'overview-home-officer.json': home(D.OFFICER, True),
        'overview-raids.json': raids_page(),
        'overview-progression.json': progression(),
        'overview-raid-member.json': raid_view(D.MEMBER, False),
        'overview-raid-officer.json': raid_view(D.OFFICER, True),
        'overview-checklist.json': checklist(),
        'overview-actions.json': officer_actions(),
        'overview-attendance.json': recent_attendance(),
        'overview-readiness.json': readiness(False),
        'overview-readiness-raid.json': readiness(True),
        'overview-recruit.json': recruit_card(),
        'loot-nights.json': loot_nights(),
        'loot-drop.json': council_drop(),
        'loot-table.json': loot_table(),
    }
