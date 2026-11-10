"""Writes the dense guild world's API answers (guild_dense_data.py) into web/tests/e2e/fixtures/guild-parity/dense-*.json.

Each answer keeps the shape of the API's Go view (api/internal/guilds) as the older parity fixtures carry it; this module only
fills them with the 40 raiders and five nights. web/tests/e2e/support/guild-parity/dense.ts routes these files.

  python3 design/mocks/guild_dense_fixtures.py
"""
from __future__ import annotations

from datetime import datetime, timedelta

import guild_data as gd
import guild_dense_data as D
import json

from guild_fixtures import OUT, write

GUILD = {'id': gd.GUILD_ID, 'name': gd.GUILD, 'region': gd.REGION, 'ruleset': gd.RULESET, 'faction': gd.FACTION, 'crest_url': None}
ZONE_KEY = 'onyxias-lair'
ZONE_NAME = "Onyxia's Lair"
ENCOUNTER_ID = 1084
TIER = {'name': 'First tier', 'raids': ['Barrow Deeps', 'Hyjal Summit', ZONE_NAME], 'named_encounters': 1, 'down': 1,
        'pulls_total': len(D.NIGHTS) + 1}
CELL = {'P': 'present', 'A': 'absent', 'B': 'bench', 'N': 'none'}
PENDING = [('Brambleshot', 'hunter', 'Marksmanship', 57), ('Mirelocke', 'warlock', 'Affliction', 58),
           ('Oakhallow', 'druid', 'Restoration', 59)]


def iso(at: datetime) -> str:
    return at.strftime('%Y-%m-%dT%H:%M:%SZ')


def parse(at: str) -> datetime:
    return datetime.strptime(at, '%Y-%m-%dT%H:%M:%SZ')


def ago(days: int, hours: int = 2, minutes: int = 30) -> str:
    return iso(parse(D.NOW) - timedelta(days=days, hours=hours, minutes=minutes))


def night_date(at: str) -> str:
    """The raid's local date (8 pm CST the evening before the UTC date)."""
    return (parse(at) - timedelta(hours=6)).strftime('%Y-%m-%d')


def report_of(raid_id: int) -> str:
    return next(n[2] for n in D.NIGHTS if n[0] == raid_id)


# ---- the roster
def rating(r: D.Raider) -> dict:
    o = r.ratings[1]
    parts = {k: max(20, min(99, o + D._n(r.name, k, -12, 12))) for k in
             ('output', 'survival', 'mechanics', 'utility', 'preparation', 'activity')}
    return {'overall': o, **parts}


def best_parse(r: D.Raider) -> dict:
    fight = next(n for n in D.NIGHTS if n[0] == D.LAST_RAID_ID)
    return {'metric': r.metric, 'value': max(r.figures), 'percentile': r.best_pct, 'encounter': 'Onyxia',
            'report_id': fight[2], 'fight_index': 4}


def roster_row(r: D.Raider) -> dict:
    seen = ago(r.seen_days)
    return {'character_key': r.key, 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': r.name, 'class': r.cls, 'spec': r.spec,
            'spec_name': f'{r.spec} {r.cls.title()}', 'role': r.role, 'rank': r.rank, 'verified': True, 'logged_recently': r.seen_days == 0, 'logged_at': seen,
            'group': r.group.lower(), 'item_level': r.ilvl, 'level': 60, 'consent': 'gear_bags', 'is_main': True,
            'attendance': {'present': r.present, 'nights': r.nights},
            'attendance_30': {'present': r.present, 'nights': r.nights}, 'best_parse': best_parse(r), 'rating': rating(r),
            'professions': list(r.professions), 'account_key': f'acct-{r.name.lower()}', 'last_report_at': None,
            'sim_export': True, 'may_remove': r.rank == 'member',
            'addon': {'level': 60, 'spec': r.spec_key, 'gear_gap': None if r.gap is None else round(r.gap),
                      'build': None, 'seen_at': (parse(seen) - timedelta(hours=6)).strftime('%Y-%m-%d %H:%M'), 'snapshot_at': D.UPDATED_AT}}


def pending_row(p: tuple) -> dict:
    name, cls, spec, ilvl = p
    return {'character_key': f'{gd.REGION}/{gd.RULESET}/{name.lower()}', 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': name,
            'class': cls, 'spec': spec, 'role': 'healer' if spec == 'Restoration' else 'dps', 'rank': 'member', 'verified': False,
            'logged_recently': True, 'logged_at': D.UPDATED_AT, 'item_level': ilvl, 'level': 60, 'consent': 'gear_bags',
            'attendance': None, 'best_parse': None, 'rating': None, 'account_key': f'acct-{name.lower()}', 'may_remove': True,
            'may_approve': True, 'addon': None}


def counts_of(rows: list[D.Raider]) -> dict:
    by = {g: [r for r in rows if r.group == g and r.answer == 'in'] for g in D.GROUPS}
    late = {g: sum(1 for r in rows if r.group == g and r.answer == 'late') for g in D.GROUPS}
    return {'in': len(D.answered('in')), 'late': len(D.answered('late')), 'bench': len(D.answered('bench')),
            'out': len(D.answered('out')), 'tanks': len(by['Tank']), 'healers': len(by['Healer']),
            'dps': len(by['Melee']) + len(by['Ranged']), 'melee': len(by['Melee']), 'ranged': len(by['Ranged']),
            'classes': {}, 'late_groups': {'tanks': late['Tank'], 'healers': late['Healer'], 'melee': late['Melee'],
                                           'ranged': late['Ranged']}}


def next_raid(viewer: str) -> dict:
    me = D.BY_NAME[viewer]
    c = counts_of(D.ROSTER)
    return {'id': D.NEXT_RAID_ID, 'zone': ZONE_KEY, 'zone_name': ZONE_NAME, 'zone_kind': 'raid', 'title': ZONE_NAME,
            'starts_at': D.NEXT_RAID_AT, 'status': 'planned', 'size': 40, 'guild_timezone': D.ZONE,
            'counts': {k: c[k] for k in ('in', 'late', 'tanks', 'healers', 'dps', 'melee', 'ranged')},
            'own': {'character_key': me.key, 'state': me.answer, 'role': me.role, 'note': ''}}


def reports() -> list[dict]:
    return [{'id': rep, 'title': f'{ZONE_NAME} · raid night', 'zone': ZONE_NAME, 'created_at': at, 'fight_count': 3 + wipes,
             'kill_count': 1, 'duration_ms': 3_000_000 + kill} for (_, at, rep, kill, wipes, _) in reversed(D.NIGHTS)]


def standing(viewer: str) -> dict | None:
    me = D.BY_NAME[viewer]
    peers = sorted((r for r in D.ROSTER if (r.cls, r.spec) == (me.cls, me.spec)), key=lambda r: -r.ilvl)
    return {'spec': me.spec, 'class': me.cls, 'same_spec_count': len(peers), 'rank_by_item_level': peers.index(me) + 1,
            'item_level': me.ilvl, 'needs_before_next_raid': ['Flask or elixir', 'Buff food'],
            'needs_raid': {'title': ZONE_NAME, 'starts_at': D.NEXT_RAID_AT, 'character_key': me.key}}


def home(viewer: str, officer: bool) -> dict:
    me = D.BY_NAME[viewer]
    return {'guild': GUILD,
            'viewer': {'role': 'officer' if officer else 'member', 'character_key': me.key, 'verified': True,
                       'officer_tools': officer, 'raid_guild_tools': True, 'verify': {'has_approver': True}},
            'claim': {'state': 'claimed', 'since': gd.CLAIM_SINCE, 'frozen': False, 'yours': officer,
                      'claimed_by': {'name': gd.VIEWER, 'class': 'hunter', 'character_key': D.BY_NAME[gd.VIEWER].key}},
            'summary': {'raiders': len(D.ROSTER), 'waiting_for_approval': len(PENDING) if officer else 0,
                        'below_rating_floor': sum(1 for r in D.ROSTER if r.ratings[1] < 55), 'tier': TIER,
                        'named_encounters_down': 1, 'pulls_this_tier': TIER['pulls_total'], 'updated_at': D.UPDATED_AT},
            'standing': None if officer else standing(viewer), 'reports': reports(),
            'roster': [roster_row(r) for r in D.ROSTER], 'pending': [pending_row(p) for p in PENDING] if officer else [],
            'addon_snapshot': {'received_at': D.UPDATED_AT, 'by_character': D.BY_NAME[gd.VIEWER].key, 'members': 43},
            'next_raid': next_raid(viewer), 'attendance_source': 'raids'}


# ---- raid 45 (planned)
def signup(r: D.Raider, viewer: str) -> dict:
    return {'character_key': r.key, 'name': r.name, 'class': r.cls, 'spec': r.spec, 'role': r.role, 'state': r.answer,
            'note': 'Late, work till 8:30' if r.name == 'Pinefeather' else '', 'late_minutes': 30 if r.answer == 'late' else None,
            'own': r.name == viewer, 'group': r.group.lower(), 'spec_key': r.spec_key, 'also': []}


def cover() -> list[dict]:
    """The buff and debuff table for raid 45 (keys, spells and icons from api/internal/guilds/raids_cover_catalogue.go): providers by sign-up state."""
    def prov(cls: str, spec: str | None = None) -> dict:
        rows = [r for r in D.ROSTER if r.cls == cls and (spec is None or r.spec == spec)]
        return {s: sum(1 for r in rows if r.answer == s) for s in ('in', 'late', 'bench')}
    table = [('leader_of_the_pack', 'Leader of the Pack', 'buff', 17007, 'spell_nature_unyeildingstamina', 1, 'raid', prov('druid', 'Feral')),
             ('moonkin_aura', 'Moonkin Aura', 'buff', 24858, 'spell_nature_forceofnature', 1, 'raid', prov('druid', 'Balance')),
             ('trueshot_aura', 'Trueshot Aura', 'buff', 20906, 'ability_trueshot', 1, 'raid', prov('hunter', 'Marksmanship')),
             ('windfury_totem', 'Windfury Totem', 'buff', 10614, 'spell_nature_windfury', 1, 'melee_group', prov('shaman')),
             ('battle_shout', 'Battle Shout', 'buff', 11551, 'ability_warrior_battleshout', 1, 'melee_group', prov('warrior')),
             ('blessing_of_kings', 'Blessing of Kings', 'buff', 20217, 'spell_magic_magearmor', 1, 'raid', prov('paladin')),
             ('sunder_armor', 'Sunder Armor', 'debuff', 11597, 'ability_warrior_sunder', 1, 'raid', prov('warrior')),
             ('curse_of_the_elements', 'Curse of the Elements', 'debuff', 1311680, 'warlock_curse_shadow', 1, 'raid', prov('warlock'))]
    out = []
    for key, name, kind, spell, icon, need, per, p in table:
        state = 'ok' if p['in'] >= need else 'late' if p['in'] + p['late'] >= need else 'gap'
        out.append({'key': key, 'name': name, 'kind': kind, 'spell_id': spell, 'icon': icon, 'need': need, 'per': per,
                    'providers_in': p['in'], 'providers_late': p['late'], 'providers_bench': p['bench'], 'likely': 0, 'state': state})
    return sorted(out, key=lambda c: (c['providers_in'] - c['need'], c['name']))


def raid_view(viewer: str, officer: bool) -> dict:
    me = D.BY_NAME[viewer]
    signed = [r for r in D.ROSTER if r.answer is not None]
    not_signed = [r for r in D.ROSTER if r.answer is None]
    return {'id': D.NEXT_RAID_ID, 'zone': ZONE_KEY, 'zone_name': ZONE_NAME, 'zone_kind': 'raid', 'key_item': None,
            'title': ZONE_NAME, 'starts_at': D.NEXT_RAID_AT, 'duration_minutes': 180, 'status': 'planned', 'size': 40,
            'guild_timezone': D.ZONE, 'notes_html': '', 'notes': '', 'counts': counts_of(D.ROSTER),
            'signups': [signup(r, viewer) for r in signed],
            'not_signed': [{'character_key': r.key, 'name': r.name, 'class': r.cls, 'spec': r.spec, 'group': r.group.lower(),
                            'spec_key': r.spec_key} for r in not_signed],
            'viewer': {'characters': [{'character_key': me.key, 'name': me.name, 'class': me.cls, 'spec': me.spec,
                                       'role': me.role, 'state': me.answer}]},
            'discord': {'available': True, 'connected': True, 'post_default': True, 'remind_default': True, 'channel_name': 'raids'},
            'discord_post': {'posted_at': D.NEXT_POSTED_AT, 'post_at': None, 'reminded_at': None, 'remind': True},
            'cover': cover(), 'comp_target': {**D.COMP_TARGET, 'source': 'guild'}, 'assignments': None,
            'pre_done_status': None, 'match': None, 'night': None, 'attendance': None, 'report': None,
            'not_signed_raiders': [{'character_key': r.key, 'name': r.name, 'class': r.cls, 'account_key': f'acct-{r.name.lower()}',
                                    'discord_linked': r.name != 'Skyarcana', 'role': r.role,
                                    'last_week': CELL[r.cells[-1]] if officer else None} for r in not_signed] if officer else [],
            'ping': {'last_pinged_at': None, 'next_allowed_at': None}}


# ---- raids 43 and 44 (done): the report's parses for every raider who came
def done_view(raid_id: int) -> dict:
    idx = [n[0] for n in D.NIGHTS].index(raid_id)
    _, at, rep, kill, wipes, deaths = D.NIGHTS[idx]
    week = 1 if raid_id == D.LAST_RAID_ID else 0
    came = [r for r in D.ROSTER if r.cells[idx] == 'P']
    parses = sorted(({'character_key': r.key, 'name': r.name, 'class': r.cls, 'role': r.role, 'encounter_id': ENCOUNTER_ID,
                      'encounter': 'Onyxia', 'fight_index': 3 + wipes, 'metric': r.metric, 'value': r.figures[week],
                      'rating': r.ratings[week], 'percentile': max(5, min(99, r.best_pct - 6 + 6 * week))} for r in came), key=lambda p: (p['metric'], -p['value']))
    loot = [{'item_id': a[1], 'item_name': a[2], 'quality': 4, 'icon': a[3], 'character_key': D.BY_NAME[a[5]].key,
             'name': a[5], 'class': D.BY_NAME[a[5]].cls, 'by_name': a[6], 'at': iso(parse(at) + timedelta(minutes=61 + i))}
            for i, a in enumerate(x for x in D.AWARDS if x[0] == raid_id)]
    return {'id': raid_id, 'zone': ZONE_KEY, 'zone_name': ZONE_NAME, 'zone_kind': 'raid', 'key_item': None, 'title': ZONE_NAME,
            'starts_at': at, 'duration_minutes': 180, 'status': 'done', 'size': 40, 'guild_timezone': D.ZONE, 'notes_html': '',
            'notes': '', 'counts': counts_of(D.ROSTER), 'signups': [], 'not_signed': [], 'viewer': {'characters': []},
            'discord': None, 'discord_post': None, 'cover': [], 'comp_target': {**D.COMP_TARGET, 'source': 'guild'},
            'assignments': None, 'pre_done_status': None,
            'match': {'report_id': rep, 'zone_name': ZONE_NAME, 'uploaded_at': iso(parse(at) + timedelta(hours=3)),
                      'uploaded_by_name': gd.VIEWER, 'pulls': 3 + wipes, 'kills': 1, 'duration_ms': 3_000_000 + kill, 'how': 'auto'},
            'night': {'signed': len(came) + 2, 'came': len(came), 'late': 1, 'not_in_log': 2, 'joined': 0, 'excused': 0, 'kills': 1,
                      'wipes': wipes, 'deaths': deaths, 'duration_ms': 3_000_000 + kill},
            'attendance': [], 'report': {
                'headline': {'kills': 1, 'wipes': wipes, 'attempts': 1 + wipes, 'deaths': deaths},
                'pulls': [], 'parses': parses,
                'kill_order': [{'encounter_id': ENCOUNTER_ID, 'name': 'Onyxia', 'fight_index': 3 + wipes, 'killed': True,
                                'kill_ms': kill, 'wipes_before': wipes, 'deaths': deaths}],
                'unnamed': [], 'deaths': [{'character_key': D.BY_NAME[n].key, 'name': n, 'class': D.BY_NAME[n].cls, 'on_wipes': 0, 'on_kills': k}
                                         for n, k in D.DEATHS.get(raid_id, {}).items()], 'consumables': {'log_has_buffs': False, 'rows': [], 'by_pull': []},
                'loot': [{'encounter_id': ENCOUNTER_ID, 'boss': 'Onyxia', 'awards': loot}],
                'moves': {'previous': None, 'rows': []}, 'posted_at': None}}


def raids_page() -> dict:
    rows = []
    for i, (rid, at, rep, kill, wipes, deaths) in enumerate(reversed(D.NIGHTS)):
        idx = len(D.NIGHTS) - 1 - i
        came = [r for r in D.ROSTER if r.cells[idx] == 'P']
        top = max((r for r in came if r.metric == 'dps' and r.group != 'Tank'), key=lambda r: r.figures[1])
        rows.append({'id': rep, 'title': f'{ZONE_NAME} · raid night', 'zone': ZONE_NAME, 'created_at': at, 'fight_count': 3 + wipes,
                     'kill_count': 1, 'duration_ms': 3_000_000 + kill, 'wipe_count': wipes, 'raiders': len(came), 'deaths': deaths,
                     'top_hps': None, 'top_dps': {'name': top.name, 'class': top.cls, 'metric': 'dps', 'value': top.figures[1]},
                     'fights': [{'index': 0, 'name': None, 'encounter_id': None, 'kill': True, 'duration_ms': 60_000, 'deaths': 0, 'players': len(came)}]
                     + [{'index': 1 + w, 'name': 'Onyxia', 'encounter_id': ENCOUNTER_ID, 'kill': w == wipes, 'duration_ms': kill if w == wipes else 240_000,
                         'deaths': deaths if w == wipes else 5, 'players': len(came)} for w in range(wipes + 1)],
                     'present': [], 'raid_id': rid})
    return {'rows': rows, 'next_cursor': None}


def progression() -> dict:
    return {'tier': TIER, 'unnamed': [{'zone': 'Barrow Deeps', 'pulls': 0, 'nights': 0}, {'zone': 'Hyjal Summit', 'pulls': 0, 'nights': 0}],
            'encounters': [{'encounter_id': ENCOUNTER_ID, 'name': 'Onyxia', 'zone': ZONE_NAME, 'first_kill_at': D.NIGHTS[0][1],
                            'pulls': TIER['pulls_total'], 'kills': len(D.NIGHTS), 'best_kill_ms': min(n[3] for n in D.NIGHTS),
                            'pulls_by_night': [{'report_id': n[2], 'date': night_date(n[1]), 'pulls': 1 + n[4], 'killed': True}
                                               for n in D.NIGHTS],
                            'deaths_per_pull': round(sum(n[5] for n in D.NIGHTS) / TIER['pulls_total'], 1), 'best_by_role': {}}]}


def guild_page() -> dict:
    return {'guild': GUILD, 'roster_best': [],
            'reports': [{k: r[k] for k in ('id', 'title', 'zone', 'created_at')} for r in reports()],
            'progression': [{'encounter': 'Onyxia', 'encounter_id': ENCOUNTER_ID, 'difficulty': 1, 'kills': len(D.NIGHTS),
                             'pull_count': TIER['pulls_total'], 'first_kill_at': D.NIGHTS[0][1]}]}


# ---- readiness, attendance, officer actions, checklist
def readiness_row(r: D.Raider) -> dict:
    short = D.short_for_raid(r)
    return {'character_key': r.key, 'name': r.name, 'class': r.cls, 'spec': r.spec, 'consent': 'gear_bags', 'level': 60,
            'gear_gap': {'upgrades': r.upgrades, 'gain_dps': r.gap, 'not_sim_checked': 0, 'unit': None if r.gap is None else r.metric},
            'enchants': {'missing_slots': D.MISSING_ENCHANTS.get(r.name, []), 'checked': True},
            'consumables': {'state': 'short' if r.name in D.NO_FLASK_FOOD else 'stocked'},
            'talent_points_unspent': D.UNSPENT_TALENTS.get(r.name, 0), 'talent_points_read': True,
            'item_level': r.ilvl, 'item_level_delta': r.ilvl - D.MEDIAN_ILVL, 'logged_at': ago(r.seen_days), 'failing': len(short),
            'no_export': False, 'synced_before_cutoff': r.seen_days < 7, 'attuned': r.name not in D.NOT_ATTUNED,
            'resistances': {'fire': D.TANK_FIRE.get(r.name, D._n(r.name, 'fr', 0, 40)), 'nature': 0, 'frost': 0, 'shadow': 0, 'arcane': 0},
            'flask_food': {'flask_or_elixir': r.name not in D.NO_FLASK_FOOD, 'food': r.name not in D.NO_FLASK_FOOD},
            'nudge_text': f'{r.name}, before Wednesday: ' + ', '.join(short) if short else None}


def readiness(for_raid: bool) -> dict:
    rows = [r for r in D.ROSTER if (not for_raid) or r.answer in ('in', 'late')]
    out = {'median_item_level': D.MEDIAN_ILVL, 'generated_at': D.UPDATED_AT, 'rows': [readiness_row(r) for r in rows]}
    nr = {'id': D.NEXT_RAID_ID, 'title': ZONE_NAME, 'zone_name': ZONE_NAME, 'starts_at': D.NEXT_RAID_AT}
    if for_raid:
        out['raid'] = {**nr, 'cutoff_at': D.NEXT_RAID_AT, 'attunement': {'kind': 'quest', 'key_item': {'id': 16309, 'name': 'Drakefire Amulet', 'quality': 3}},
                       'fire_resistance': None}
    else:
        out['next_raid'] = nr
    return out


def attendance() -> dict:
    return {'raids': [{'id': n[0], 'starts_at': n[1], 'zone_name': ZONE_NAME} for n in D.NIGHTS],
            'rows': sorted(({'name': r.name, 'class': r.cls, 'account_key': f'acct-{r.name.lower()}',
                             'cells': [CELL[c] for c in r.cells],
                             'slipping': sum(1 for c in r.cells[-4:] if c == 'A') >= 2} for r in D.ROSTER), key=lambda x: x['name'])}


def officer_actions() -> dict:
    not_signed = [r.name for r in D.ROSTER if r.answer is None]
    open_ = len(D.OPEN_DROPS)
    return {'approvals': {'count': len(PENDING), 'names': [p[0] for p in PENDING]},
            'applications': {'count': 2, 'names': ['Ashenveil', 'Quillbright']},
            'not_signed': {'raid_id': D.NEXT_RAID_ID, 'count': len(not_signed), 'names': not_signed},
            'drops_to_award': {'tracked': True, 'count': open_, 'raid_id': D.LAST_RAID_ID, 'label': 'Onyxia, Wed 6 Jan',
                               'awarded': sum(1 for a in D.AWARDS if a[0] == D.LAST_RAID_ID)}}


def checklist() -> dict:
    me = D.BY_NAME[D.MEMBER]
    return {'character': {'character_key': me.key, 'name': me.name, 'class': me.cls, 'spec': me.spec, 'level': 60},
            'raid': {'id': D.NEXT_RAID_ID, 'zone_key': ZONE_KEY, 'zone_name': ZONE_NAME}, 'verdict': 'almost',
            'items': [{'key': 'level', 'state': 'ok', 'level': 60, 'required_level': 60},
                      {'key': 'attunement', 'state': 'ok', 'attunement_kind': 'quest', 'key_item': {'id': 16309, 'name': 'Drakefire Amulet', 'quality': 3}},
                      {'key': 'gear_gap', 'state': 'fail', 'upgrades': me.upgrades, 'gain_dps': me.gap, 'slots': ['hands', 'finger1'], 'unit': 'hps'},
                      {'key': 'enchants', 'state': 'ok', 'missing_slots': []},
                      {'key': 'flask_food', 'state': 'fail', 'supplies': [], 'missing': ['flask_or_elixir', 'food']},
                      {'key': 'fire_resistance', 'state': 'not_tracked', 'reason': 'resistance_not_in_export'}]}


# ---- loot
def award(i: int, a: tuple) -> dict:
    w = D.BY_NAME[a[5]]
    return {'award_id': i + 1, 'item_id': a[1], 'item_name': a[2], 'icon': a[3], 'quality': 4, 'character_key': w.key, 'name': w.name,
            'class': w.cls, 'ms_os': 'ms', 'reason': None, 'votes_yes': a[7], 'votes_no': gd.VOTERS - a[7], 'votes_officers': gd.VOTERS,
            'awarded_by': a[6], 'at': iso(parse(D.night_at(a[0])) + timedelta(minutes=61 + i % 5))}


def loot_nights() -> dict:
    nights = []
    for rid, at, rep, *_ in reversed(D.NIGHTS):
        won = [(i, a) for i, a in enumerate(D.AWARDS) if a[0] == rid]
        drops = [{'item_id': a[1], 'name': a[2], 'icon': a[3], 'quality': 4, 'kind': 'gear', 'encounter_id': ENCOUNTER_ID, 'boss': 'Onyxia',
                  'state': 'awarded', 'award': award(i, a), 'recorded': True, 'count': 1, 'looted_by': []} for i, a in won]
        if rid == D.LAST_RAID_ID:
            drops = [{'item_id': d[0], 'name': d[1], 'icon': d[2], 'quality': 4, 'kind': 'gear', 'encounter_id': ENCOUNTER_ID, 'boss': 'Onyxia',
                      'state': 'open', 'award': None, 'recorded': True, 'count': 1, 'looted_by': []} for d in D.OPEN_DROPS] + drops
        nights.append({'night': f'raid:{rid}', 'raid_id': rid, 'report_id': rep, 'date': night_date(at), 'at': at, 'zone': ZONE_NAME,
                       'current': rid == D.LAST_RAID_ID, 'drops': drops, 'decisions': [award(i, a) for i, a in reversed(won)]})
    return {'officers': gd.VOTERS, 'nights': nights}


def candidate(c: tuple, council: bool) -> dict:
    name, gain, worn, bis, ms_os, yes = c
    r = D.BY_NAME[name]
    out = {'character_key': r.key, 'name': name, 'class': r.cls, 'spec': r.spec, 'gain_dps': gain, 'gain_unit': 'dps', 'tier': 0,
           'not_sim_checked': False, 'ilvl_delta': None, 'attendance': {'present': r.present, 'nights': r.nights},
           'already_equivalent': False, 'fills_empty_slot': False, 'replaces': worn[1],
           'replaces_item': {'item_id': worn[0], 'name': worn[1], 'icon': worn[2], 'quality': worn[3]}, 'items_won': len(D.items_won(name))}
    if council:
        out |= {'bis_for_them': bis, 'ms_os': ms_os, 'ms_os_source': 'declared',
                'votes': {'yes': yes, 'no': 0, 'voters': gd.VOTERS, 'mine': 'yes' if yes == 2 else None}}
    return out


def loot_drop() -> dict:
    d = next(x for x in D.OPEN_DROPS if x[0] == D.SELECTED_DROP)
    return {'night': f'raid:{D.LAST_RAID_ID}', 'item_id': d[0], 'name': d[1], 'icon': d[2], 'quality': 4, 'kind': 'gear',
            'officers': gd.VOTERS, 'closed': False, 'award': None, 'candidates': [candidate(c, True) for c in D.CANDIDATES]}


def loot_table() -> dict:
    """The Loot tab's boss table: the older fixture's Onyxia table, five kill nights in, for 40 raiders at level 60."""
    base = json.loads((OUT / 'loot-table.json').read_text())
    enc = [{**e, 'kill_nights': len(D.NIGHTS), 'on_farm': True} for e in base['encounters']]
    return {**base, 'encounters': enc, 'at_level_cap': len(D.ROSTER),
            'items': [{**i, 'candidates': [candidate(c, False) for c in D.CANDIDATES] if i['item_id'] == D.SELECTED_DROP else [],
                       'eligible': []} for i in base['items']]}


def member_profile(name: str) -> dict:
    """GET .../members/me/profile for a member: their row, nights, loot this tier and readiness (no notes)."""
    r = D.BY_NAME[name]
    outcome = {'P': 'present', 'A': 'absent', 'B': None, 'N': None}
    nights = [{'raid_id': n[0], 'starts_at': n[1], 'outcome': outcome[c], 'excused': False} for n, c in zip(D.NIGHTS, r.cells) if c != 'N']
    won = D.items_won(name)
    return {'account_key': f'acct-{name.lower()}', 'member': roster_row(r),
            'characters': [{'character_key': r.key, 'name': r.name, 'class': r.cls, 'is_main': True}],
            'attendance': {'source': 'raids', 'last_4': {'present': r.cells[-4:].count('P'), 'nights': 4},
                           'last_30': {'present': r.present, 'nights': r.nights}, 'nights': nights},
            'loot': {'count': len(won), 'items': [{'item_id': a[1], 'name': a[2], 'icon': a[3], 'quality': 4, 'raid_id': a[0],
                                                    'awarded_at': iso(parse(D.night_at(a[0])) + timedelta(minutes=62))} for a in won]},
            'readiness': readiness_row(r)}


def dense_fixtures() -> dict[str, object]:
    return {'dense-home-officer.json': home(D.OFFICER, True), 'dense-home-member.json': home(D.MEMBER, False),
            'dense-raid-officer.json': raid_view(D.OFFICER, True), 'dense-raid-member.json': raid_view(D.MEMBER, False),
            'dense-raid-44.json': done_view(D.LAST_RAID_ID), 'dense-raid-43.json': done_view(D.PREV_RAID_ID),
            'dense-raids.json': raids_page(), 'dense-progression.json': progression(), 'dense-page.json': guild_page(),
            'dense-readiness.json': readiness(False), 'dense-readiness-raid.json': readiness(True),
            'dense-attendance.json': attendance(), 'dense-actions.json': officer_actions(), 'dense-checklist.json': checklist(),
            'dense-loot-nights.json': loot_nights(), 'dense-loot-drop.json': loot_drop(), 'dense-loot-table.json': loot_table(),
            'dense-profile-member.json': member_profile(D.MEMBER)}


if __name__ == '__main__':
    for name, data in dense_fixtures().items():
        print(write(name, data))
