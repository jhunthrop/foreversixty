"""Guild parity fixtures, the guild landing (/guild), join and claim: the API answers web/tests/e2e/support/guild-parity/directory.ts
routes, written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/. ws_guild_directory.py draws the mocks from the same values.

States: `landing-out` (signed out), `landing-in` (signed in on an alt with no guild; the account is in two guilds, so the landing
stays and Your guilds leads the rail), `join` (a signed-in player with no guild opens Olympus XXVII's invite link) and `claim`
(Dunmaro, an officer, holds Olympus XXVII's claim; the guild has no plan yet).
"""
from pathlib import Path

import guild_data as gd
from guild_overview_fixtures import tier

INVITE_TOKEN = 'fixtureinvitetoken1'  # web/src/lib/report/shell-paths.ts FIXTURE_GUILD_INVITE_TOKEN

# The progression board's top five (GET /v1/rankings/guilds?kind=progress): name, region, ruleset, bosses down. The first tier
# names one boss (Onyxia), so nobody has more than one down; Olympus XXVII's count is its own tier bar's
# (guild_overview_fixtures.tier). The board ranks the earliest kill first: FOUGHT_AT, Olympus XXVII's on its first night.
TOP = [(gd.GUILD, 'us', 'pvp', tier()['down']), ('The Last Watch', 'us', 'hardcore', 1), ('Iron Vanguard', 'us', 'hardcore', 1),
       ('Ashen Covenant', 'eu', 'normal', 1), ('Pinewood Wardens', 'us', 'normal', 1)]
FOUGHT_AT = {gd.GUILD: '2026-12-10T03:45:00Z', 'The Last Watch': '2026-12-10T04:20:00Z', 'Iron Vanguard': '2026-12-11T03:10:00Z',
             'Ashen Covenant': '2026-12-11T21:40:00Z', 'Pinewood Wardens': '2026-12-12T03:30:00Z'}
RULESET_WORD = {'pvp': 'PvP', 'hardcore': 'Hardcore', 'normal': 'Normal'}

# landing-in: the account's guilds (name, region, ruleset, rank) and the alt the header points at.
MINE = [(gd.GUILD, 'us', 'pvp', 'officer'), ('The Last Watch', 'us', 'hardcore', 'member')]
ALT = ('Briarwick', 'druid', 23)

# join: the joining player's character, and the officer who made the link.
JOINER = ('Ashenveil', 'paladin', 'Human', 60)
INVITER = 'Ironbrand'

CLAIMANT = gd.VIEWER  # Dunmaro
CLAIM_SINCE = gd.CLAIM_SINCE


def _key(name: str, region: str = gd.REGION, ruleset: str = gd.RULESET) -> str:
    return f'{region}/{ruleset}/{name.lower()}'


def _class_of(name: str) -> str:
    return next(r[1] for r in gd.R if r[0] == name)


def rankings() -> dict:
    return {'rows': [{'rank': i + 1, 'guild': {'name': n, 'region': r, 'ruleset': rs}, 'value': v, 'fought_at': FOUGHT_AT[n],
                      'report_id': gd.LAST_REPORT_ID if n == gd.GUILD else f'parity{i}'} for i, (n, r, rs, v) in enumerate(TOP)]}


def _user(uid: int, tag: str) -> dict:
    return {'id': uid, 'battletag': tag, 'email': None, 'role': 'user', 'anonymize': False, 'premium': False}


def me_in() -> dict:
    name, cls, level = ALT
    alt = {'key': _key(name), 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': name, 'class': cls.title(), 'level': level, 'faction': gd.FACTION}
    main = {'key': _key(CLAIMANT), 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': CLAIMANT, 'class': 'Hunter', 'level': 60, 'faction': gd.FACTION,
            'guild': {'id': gd.GUILD_ID, 'name': gd.GUILD, 'verified': True}}
    guilds = [{'id': gd.GUILD_ID if n == gd.GUILD else 501, 'region': r, 'ruleset': rs, 'name': n, 'rank': rank, 'verified': True}
              for n, r, rs, rank in MINE]
    return {'user': _user(5, 'Dunmaro#1105'), 'characters': [alt, main], 'main_character_key': _key(CLAIMANT), 'guilds': guilds}


def me_join() -> dict:
    name, cls, _race, level = JOINER
    character = {'key': _key(name), 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': name, 'class': cls.title(), 'level': level, 'faction': gd.FACTION}
    return {'user': _user(41, 'Ashenveil#2201'), 'characters': [character], 'main_character_key': _key(name), 'guilds': []}


def invite_peek() -> dict:
    name, cls, race, level = JOINER
    return {'guild': {'id': gd.GUILD_ID, 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': gd.GUILD, 'faction': gd.FACTION, 'crest_url': None},
            'members': gd.RAIDERS, 'inviter': {'name': INVITER, 'class': _class_of(INVITER), 'character_key': _key(INVITER)},
            'joins_as': {'name': name, 'class': cls, 'race': race, 'level': level, 'on_roster': True}}


def me_claim() -> dict:
    main = {'key': _key(CLAIMANT), 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': CLAIMANT, 'class': 'Hunter', 'level': 60, 'faction': gd.FACTION,
            'guild': {'id': gd.GUILD_ID, 'name': gd.GUILD, 'verified': True}}
    guild = {'id': gd.GUILD_ID, 'region': gd.REGION, 'ruleset': gd.RULESET, 'name': gd.GUILD, 'slug': gd.SLUG, 'faction': gd.FACTION, 'rank': 'officer',
             'verified': True, 'verified_count': gd.RAIDERS, 'claim': {'state': 'claimed', 'yours': True}, 'covered': None}
    return {'user': _user(5, 'Dunmaro#1105'), 'characters': [main], 'main_character_key': _key(CLAIMANT), 'guilds': [guild]}


def claim_settings() -> dict:
    return {'default_visibility': 'guild', 'officer_max_rank_index': 1,
            'claimed_by': {'name': CLAIMANT, 'class': 'hunter', 'character_key': _key(CLAIMANT)}, 'claim_pending': None,
            'claim': {'state': 'claimed', 'since': CLAIM_SINCE, 'frozen': False, 'yours': True},
            'invite': {'rotated_at': None, 'url': None, 'made_by': None}}


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {'directory-rankings.json': rankings(), 'directory-me-in.json': me_in(), 'directory-me-join.json': me_join(),
            'directory-invite.json': invite_peek(), 'directory-me-claim.json': me_claim(), 'directory-claim-settings.json': claim_settings()}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
