"""The dense world's answers for the tabs batch H redraws (Recruit, Applications, Settings), into
web/tests/e2e/fixtures/guild-parity/dense-*.json. They keep the shapes of the older parity fixtures (recruit-*.json,
settings-*.json) and re-date them to the 40-raider world's clock (guild_dense_data.NOW, Sat 9 Jan 2027).

  python3 design/mocks/guild_dense_tabs_fixtures.py
"""
from __future__ import annotations

import json
from datetime import timedelta

import guild_data as gd
import guild_dense_data as D
from guild_dense_fixtures import GUILD, TIER, ago, iso, parse
from guild_fixtures import OUT, write

RAID_NIGHTS_TEXT = 'Wed, 8:00 pm to 11:00 pm CST'
PITCH = 'We raid Wednesday nights, bring flasks, and split loot by council with the numbers shown. New to Forever is fine.'
# The open spots (guild_dense_data.COMP_TARGET against the roster): a third tank, healers, a ranged.
WANTED = [('warrior', 'warrior-protection', 'Protection Warrior', 'Off tank'),
          ('paladin', 'paladin-holy', 'Holy Paladin', 'For the heal rotation'),
          ('priest', 'priest-holy', 'Holy Priest', ''),
          ('warlock', 'warlock-destruction', 'Destruction Warlock', 'Curse of Recklessness')]
# Applications: id, name, class, spec key, spec words, ilvl, metric, figure, percentile, days ago, availability, message,
# votes yes, votes no, my vote, status, days to decision.
APPLICATIONS = [
    (48, 'Ashenveil', 'paladin', 'paladin-holy', 'Holy Paladin', 63, 'hps', 612.4, 91, 2, 'Wed', 'Two tiers in a guild that disbanded.',
     2, 0, 'yes', 'pending', None),
    (49, 'Quillbright', 'warlock', 'warlock-destruction', 'Destruction Warlock', 60, 'dps', 488.0, 74, 1, '',
     'New to Forever, level 60 since launch.', 1, 1, None, 'pending', None),
    (46, 'Thornlash', 'warrior', 'warrior-protection', 'Protection Warrior', 58, 'dps', 301.5, 38, 6, 'Wed, Sun',
     'Off tank on Fresh, can main tank.', 0, 0, None, 'pending', None),
    (44, 'Oakhallow', 'druid', 'druid-restoration', 'Restoration Druid', 59, 'hps', 540.0, 66, 12, 'Wed', '', 3, 0, None, 'accepted', 10),
    (41, 'Brambleshot', 'hunter', 'hunter-marksmanship', 'Marksmanship Hunter', 57, 'dps', 402.0, 41, 16, 'Wed',
     'Hunter main, alt warrior.', 3, 0, None, 'accepted', 14),
    (39, 'Saltmarsh', 'rogue', 'rogue-combat', 'Combat Rogue', 55, 'dps', 330.0, 22, 20, '', '', 0, 3, None, 'declined', 19),
]


def recruit_card() -> dict:
    base = json.loads((OUT / 'recruit-card.json').read_text())
    upcoming = [D.NEXT_RAID_AT, D.SECOND_RAID_AT]
    return {**base, 'guild': {**GUILD, 'path': f'/guild/{gd.REGION}/{gd.RULESET}/olympus-xxvii'}, 'members': len(D.ROSTER),
            'pitch': PITCH, 'raid_nights': {'text': RAID_NIGHTS_TEXT, 'upcoming': upcoming, 'timezone': D.ZONE},
            'progress': TIER, 'updated_at': D.UPDATED_AT,
            'wanted': [{'class': c, 'spec': s, 'spec_name': n, 'note': note} for c, s, n, note in WANTED]}


def application(a: tuple) -> dict:
    (aid, name, cls, spec, spec_name, ilvl, metric, value, pct, days, avail, message, yes, no, mine, status, decided) = a
    created = ago(days, 4, 0)
    decided_at = None if decided is None else iso(parse(created) + timedelta(days=days - decided))
    return {'id': aid, 'guild_id': gd.GUILD_ID, 'status': status, 'created_at': created, 'decided_at': decided_at,
            'message': message, 'availability': avail, 'discord_dm': False,
            'character': {'character_key': f'{gd.REGION}/{gd.RULESET}/{name.lower()}', 'name': name, 'class': cls, 'spec': spec,
                          'spec_name': spec_name, 'level': 60, 'item_level': ilvl,
                          'best_parse': {'metric': metric, 'value': value, 'percentile': pct, 'encounter': 'Onyxia',
                                         'report_id': 'parity', 'fight_index': 3}},
            'votes': {'yes': yes, 'no': no, 'voters': gd.VOTERS, 'mine': mine},
            'decided_by': None if decided is None else gd.VIEWER, 'role_outcome': None, 'dm_outcome': None}


def applications() -> dict:
    return {'rows': [application(a) for a in APPLICATIONS]}


def billing() -> dict:
    base = json.loads((OUT / 'settings-billing.json').read_text())
    return {**base, 'verified_count': len(D.ROSTER)}


def settings() -> dict:
    return json.loads((OUT / 'settings-settings.json').read_text())


def setup() -> dict:
    """The setup checklist five nights in: everything done but the linked Discord role."""
    keys = ['officer_ranks', 'discord_bot', 'recruiting', 'raid_schedule', 'invite_link', 'addon_sync', 'report_visibility',
            'linked_role']
    steps = [{'key': k, 'state': 'todo' if k == 'linked_role' else 'done', 'marked': False} for k in keys]
    synced = [r for r in D.ROSTER if r.seen_days < 7]
    return {'steps': steps, 'done': len(keys) - 1, 'total': len(keys), 'finished': False, 'plan': 'raid_guild',
            'ranks': {'officer_max_rank_index': 1, 'officers': 1 + len(gd.OFFICERS),
                      'ranks': [{'rank_index': 0, 'members': 1, 'names': [gd.GUILD_MASTER]},
                                {'rank_index': 1, 'members': len(gd.OFFICERS), 'names': list(gd.OFFICERS)},
                                {'rank_index': 2, 'members': len(D.ROSTER) - 1 - len(gd.OFFICERS),
                                 'names': [r.name for r in D.ROSTER if r.rank == 'member'][:3]}]},
            'discord': {'bot_available': True, 'installed': True, 'server_name': 'Olympus XXVII', 'channel_name': 'raid-signups',
                        'recruit_channel_name': 'recruitment', 'webhook_connected': False, 'linked_members': 31,
                        'members': len(D.ROSTER)},
            'recruiting': {'open': True, 'pitch': True, 'nights': RAID_NIGHTS_TEXT, 'wanted': [w[1] for w in WANTED], 'role_set': False},
            'schedule': {'timezone': D.ZONE, 'planned': 2, 'next': D.NEXT_RAID_AT},
            'invite': {'made': True, 'joined': len(D.ROSTER) - 1},
            'sync': {'synced': len(synced), 'accounts': len(D.ROSTER), 'window_days': 14,
                     'missing': [r.name for r in D.ROSTER if r.seen_days >= 7]},
            'visibility': 'guild', 'readiness': None}


def tabs_fixtures() -> dict[str, object]:
    return {'dense-recruit.json': recruit_card(), 'dense-applications.json': applications(),
            'dense-applications-mine.json': {'rows': []}, 'dense-billing.json': billing(), 'dense-settings.json': settings(),
            'dense-setup.json': setup()}



if __name__ == '__main__':
    for name, data in tabs_fixtures().items():
        print(write(name, data))
