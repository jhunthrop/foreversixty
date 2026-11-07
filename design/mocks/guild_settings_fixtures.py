"""Guild parity fixtures, the Settings tab and the guild settings page: the API answers web/tests/e2e/support/guild-parity/settings.ts
routes, written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/. The guild and its officer are the Overview's (guild_data):
Dunmaro holds the claim and pays for a founder Raid guild plan with one of its three Premium seats given; the standalone page's guild has no
plan yet, so it draws the plan chooser.
"""
from pathlib import Path

import guild_data as gd

VIEWER = 'Dunmaro'
SEAT_HOLDER = 'Frostquill'
TRANSFER_TO = ['Ironbrand', 'Shiverknife']
INVITE_MADE = '2026-12-01T18:00:00Z'
INVITE_URL = '/guild/invite/k7Qm2v9XcLp4'
VERIFIED = gd.RAIDERS
CHANNEL_ID = '1300000000000000002'
DISCORD_SERVER_ID = '1300000000000000001'


def _key(name: str) -> str:
    return f'{gd.REGION}/{gd.RULESET}/{name.lower()}'


def _class(name: str) -> str:
    return next(r[1] for r in gd.R if r[0] == name)


def _holder(name: str) -> dict:
    return {'name': name, 'class': _class(name), 'character_key': _key(name)}


def settings() -> dict:
    """GET /v1/guilds/2/settings: claimed by the viewer, the kept invite link, a connected Discord channel (webhook, no bot)."""
    return {'default_visibility': 'guild', 'officer_max_rank_index': 1, 'claimed_by': _holder(VIEWER), 'claim_pending': None,
            'claim': {'state': 'claimed', 'since': '2026-11-12T00:00:00Z', 'frozen': False},
            'invite': {'rotated_at': INVITE_MADE, 'url': INVITE_URL, 'made_by': _holder(VIEWER)},
            'discord': {'available': True, 'connected': True, 'channel_id': CHANNEL_ID, 'discord_guild_id': DISCORD_SERVER_ID, 'disabled_reason': None,
                        'last_error': None, 'failures': 0, 'post_default': True}}


def _seat(slot: int, holder: str | None) -> dict:
    seat = {'slot': slot, 'state': 'empty', 'holder': None, 'given_at': None, 'starts_at': None, 'movable_at': None, 'returned': None, 'offeree': None}
    if holder is None:
        return seat
    who = {'user_id': 1016, 'character_key': _key(holder), 'name': holder, 'class': _class(holder)}
    return {**seat, 'state': 'held', 'holder': who, 'given_at': '2026-11-20T00:00:00Z', 'starts_at': '2026-11-20T00:00:00Z', 'movable_at': '2026-12-20T00:00:00Z'}


def billing() -> dict:
    """GET /v1/guilds/2/billing: a founder Raid guild on the yearly price, billed to the viewer, one seat of three given."""
    plan = {'scope': 'plan', 'tier': 'raid_guild', 'status': 'active', 'current_period_end': '2027-11-12T00:00:00Z', 'cancel_at_period_end': False,
            'trial_end': None, 'founder': True, 'interval': 'yearly', 'amount_cents': 20000, 'billed_by': VIEWER, 'you_are_billing_contact': True}
    return {'billing': plan, 'trial_used': True, 'verified_count': VERIFIED, 'seats': [_seat(1, SEAT_HOLDER), _seat(2, None), _seat(3, None)],
            'transfer': None, 'transfer_candidates': [{'user_id': 1000 + i, 'name': n} for i, n in enumerate(TRANSFER_TO)],
            'can': ['billing', 'seats', 'transfer']}


def billing_none() -> dict:
    """The standalone page's guild before a plan: the chooser, the trial still open."""
    return {'billing': None, 'trial_used': False, 'verified_count': VERIFIED, 'seats': [], 'transfer': None, 'transfer_candidates': [],
            'can': ['billing', 'seats', 'transfer']}


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {'settings-settings.json': settings(), 'settings-billing.json': billing(), 'settings-billing-none.json': billing_none()}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
