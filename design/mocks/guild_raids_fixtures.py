"""Guild parity fixtures, Raids tab (list, plan form, log match) and the raid page (planner, assignments): the API answers
web/tests/e2e/support/guild-parity/raids.ts routes, written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/.

The Raids tab as its officer (Dunmaro) reads it: Thursday's Onyxia's Lair (the Overview's next raid) and the Tuesday after,
planned; Monday's logged night under them. The rest of the guild page reads the Overview's own answers (guild_overview_fixtures).
"""
import copy
from pathlib import Path

import guild_overview_data as D
import guild_overview_fixtures as F
from guild_data import ZONE

# The second planned raid: the Tuesday after, 8:00 pm CST, a third of the guild in so far, the viewer not signed up.
SECOND_RAID_ID = 42
SECOND_RAID_AT = '2026-12-16T02:00:00Z'
SECOND_RAID_IN = ['Thornwicke', 'Ironbrand', 'Hollowmend', 'Mossveil', 'Frostquill', 'Doomwhisper', 'Shiverknife', 'Groundshock']

_RAID = {'kind': 'raid', 'opens': '2026-12-09T00:00:00Z', 'min_level': None, 'max_level': None, 'key_item': None}
_EVENT = {'kind': 'event', 'name': 'Guild event', 'opens': None, 'min_level': None, 'max_level': None, 'key_item': None}
ZONES = [{**_EVENT, 'key': 'event:guild', 'size': None}, {**_EVENT, 'key': 'event:guild-10', 'size': 10},
         {**_RAID, 'key': 'raid:barrow-deeps', 'name': 'Barrow Deeps', 'size': None},
         {**_RAID, 'key': 'raid:hyjal-summit', 'name': 'Hyjal Summit', 'size': None},
         {**_RAID, 'key': 'raid:onyxias-lair', 'name': "Onyxia's Lair", 'size': 40}]
DISCORD = {'available': True, 'connected': True, 'post_default': True, 'remind_default': True, 'channel_name': 'raids'}


def second_raid() -> dict:
    view = copy.deepcopy(F.raid_view(D.OFFICER, True))
    signups = [s for s in view['signups'] if s['name'] in SECOND_RAID_IN]
    for s in signups:
        s['state'] = 'in'
    groups = {g: sum(1 for s in signups if s['group'] == g) for g in ('tank', 'healer', 'melee', 'ranged')}
    view.update({'id': SECOND_RAID_ID, 'starts_at': SECOND_RAID_AT, 'signups': signups, 'discord_post': None,
                 'counts': {'in': len(signups), 'late': 0, 'bench': 0, 'out': 0, 'tanks': groups['tank'], 'healers': groups['healer'],
                            'dps': groups['melee'] + groups['ranged'], 'melee': groups['melee'], 'ranged': groups['ranged'],
                            'late_groups': {'tanks': 0, 'healers': 0, 'melee': 0, 'ranged': 0}}})
    view['viewer']['characters'][0]['state'] = None
    return view


def raids_page() -> dict:
    """GET /v1/guilds/2/raids as the officer: the logged night, the two planned raids, the zones the plan form offers."""
    return {**F.raids_page(), 'planned': [F.raid_view(D.OFFICER, True), second_raid()], 'zones': ZONES, 'guild_timezone': ZONE,
            'discord': DISCORD}


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {'raids-page.json': raids_page()}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
