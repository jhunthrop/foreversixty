"""The data the guild mocks and the guild parity spec share (design/parity.md, "The guild states").

ws_guild2.py draws the mocks from these values and guild_fixtures.py writes them as the API's answers into
web/tests/e2e/fixtures/guild-parity, which web/tests/e2e/support/guild-parity/ routes to the build. One roster, typed once.
Shape: api/cmd/seedguild (24 characters) and api/internal/guilds (sign-up states In / Late / Bench).
"""
from mocklib import CLASS_COLOR

# The fixtures' clock and zone (web/tests/e2e/support/guild-parity/clock.ts): Sat 12 Dec 2026, 4:00 pm CST, four days and four
# hours before the next raid. The raids open Wed 9 Dec; the guild raids Wednesdays at 8:00 pm CST.
NOW = '2026-12-12T22:00:00Z'
ZONE = 'America/Chicago'


def house_clock(at) -> str:
    """The house clock (web/src/lib/dates.ts clock, en-US): `8 pm`, `8:30 pm`; minutes only off the hour."""
    return f"{at.hour % 12 or 12}{'' if at.minute == 0 else f':{at.minute:02d}'} {'am' if at.hour < 12 else 'pm'}"


def countdown_text(left) -> str:
    """The one countdown (web/src/lib/ui/countdown.ts, apps/shared/formats.json `countdown`):
    `4 days, 4 hours`, zero units left out, minutes rounded up; empty once passed."""
    total = -(-int(left.total_seconds()) // 60)
    if total <= 0:
        return ''
    days, hours, minutes = total // 1440, total % 1440 // 60, total % 60
    word = lambda n, one, many: f'{n} {one if n == 1 else many}'  # noqa: E731
    parts = [word(days, "day", "days")] if days else []
    parts += [word(hours, "hour", "hours")] if hours else []
    parts += [f"{minutes} min"] if minutes else []
    return ", ".join(parts)
# The guild's first raid night, Wed 9 Dec, 8:00 pm CST: Onyxia's first kill, the loot council's night and the last raid every
# surface reads (raid 40, its report and its log). Nothing is logged, killed or looted before it.
LAST_RAID_ID = 40
LAST_RAID_AT = '2026-12-10T02:00:00Z'
LAST_RAID_DATE = '2026-12-09'
LAST_REPORT_ID = 'olympusony1209'
# The first night's log: fight index, name, boss pull number (None on trash), result, duration ms, deaths as (who, killing
# ability). AWAY are the roster characters not in the log: the two who never signed, the bench, and Pinefeather (signed Late,
# not yet attuned).
NIGHT_PULLS = [
    (1, 'Onyxian Whelp pack', None, 'clear', 60_000, []),
    (2, 'Onyxia', 1, 'wipe', 240_000, [('Emberfall', 'Deep Breath'), ('Skytalon', 'Deep Breath'), ('Hexmarrow', 'Deep Breath'),
                                       ('Shiverknife', 'Tail Sweep'), ('Blacktide', 'Tail Sweep'), ('Seraphine', 'Flame Breath')]),
    (3, 'Onyxian Whelp pack', None, 'clear', 55_000, []),
    (4, 'Onyxia', 2, 'kill', 225_000, [('Groundshock', 'Flame Breath')]),
]
NIGHT_AWAY = ('Voidcall', 'Starnettle', 'Cinderglass', 'Nightlatch', 'Pinefeather')
# The next raid, Wed 16 Dec, 8:00 pm CST (raid 41, the planner), and when its Discord post went out.
NEXT_RAID_ID = 41
NEXT_RAID_AT = '2026-12-17T02:00:00Z'
NEXT_POSTED_AT = '2026-12-10T18:00:00Z'
# The raid after it, Wed 23 Dec (the Raids tab's second planned raid).
SECOND_RAID_ID = 42
SECOND_RAID_AT = '2026-12-24T02:00:00Z'
# Dunmaro has held the guild's claim since 12 Nov.
CLAIM_SINCE = '2026-11-12T00:00:00Z'
GUILD_ID = 2
REGION = 'us'
RULESET = 'pvp'
SLUG = 'olympus-xxvii'

ALLIANCE = '#6fb1ff'
GUILD = 'Olympus XXVII'
TABS = ['Overview', 'Roster', 'Raids', 'Progression', 'Readiness', 'Loot', 'Settings']
C = CLASS_COLOR
# seed roster (api/cmd/seedguild/roster.go): name, class, spec, role, ilvl, sign-up state
R = [('Thornwicke', 'warrior', 'Protection', 'Tank', 62, 'in'), ('Ironbrand', 'paladin', 'Holy', 'Healer', 63, 'in'), ('Valomira', 'paladin', 'Holy', 'Healer', 60, 'in'), ('Seraphine', 'paladin', 'Retribution', 'Melee', 65, 'in'),
     ('Dunmaro', 'hunter', 'Beast Mastery', 'Ranged', 61, 'in'), ('Pinefeather', 'hunter', 'Marksmanship', 'Ranged', 58, 'late'), ('Skytalon', 'hunter', 'Survival', 'Ranged', 59, 'in'), ('Shiverknife', 'rogue', 'Assassination', 'Melee', 66, 'in'),
     ('Blacktide', 'rogue', 'Combat', 'Melee', 57, 'in'), ('Hollowmend', 'priest', 'Holy', 'Healer', 64, 'in'), ('Dawnweaver', 'priest', 'Holy', 'Healer', 60, 'in'), ('Voidcall', 'priest', 'Shadow', 'Ranged', 56, 'none'),
     ('Stonetide', 'shaman', 'Restoration', 'Healer', 62, 'in'), ('Emberfall', 'shaman', 'Elemental', 'Ranged', 63, 'in'), ('Groundshock', 'shaman', 'Enhancement', 'Melee', 67, 'in'), ('Frostquill', 'mage', 'Frost', 'Ranged', 68, 'in'),
     ('Cinderglass', 'mage', 'Fire', 'Ranged', 59, 'bench'), ('Skyarcana', 'mage', 'Arcane', 'Ranged', 58, 'none'), ('Doomwhisper', 'warlock', 'Affliction', 'Ranged', 69, 'in'), ('Grimsworn', 'warlock', 'Demonology', 'Ranged', 60, 'late'),
     ('Hexmarrow', 'warlock', 'Destruction', 'Ranged', 61, 'in'), ('Mossveil', 'druid', 'Restoration', 'Healer', 64, 'in'), ('Starnettle', 'druid', 'Balance', 'Ranged', 70, 'none'), ('Nightlatch', 'rogue', 'Subtlety', 'Melee', 55, 'bench')]
STATE_TAG = {'in': ('In', 'kill'), 'late': ('Late', 'gold'), 'bench': ('Bench', 'muted'), 'none': ('Not signed', 'wipe')}

# ---- Recruitment, applications and the member profile (lane D; ws_guild2.recruitment / applications / member_profile,
# guild_fixtures.recruit_fixtures). The guild's faction and verified count, the public card, the open applications, and
# Frostquill's profile as an officer reads it.
FACTION = 'alliance'
RAIDERS = len(R)
# The ranks every surface reads (roster, recruit card, settings' officer rank index 1, the loot council): the guild master
# (rank index 0, `leader`) and the officers (rank index 1). The council that votes on loot and applications is all three.
GUILD_MASTER = 'Ironbrand'
VIEWER = 'Dunmaro'
OFFICERS = [VIEWER, 'Shiverknife']
COUNCIL = [GUILD_MASTER, *OFFICERS]
VOTERS = len(COUNCIL)


def rank_of(name: str) -> str:
    """The API's rank word for a roster character: leader, officer or member."""
    return 'leader' if name == GUILD_MASTER else 'officer' if name in OFFICERS else 'member'


RECRUIT_PITCH = 'We raid Wednesday nights, bring flasks, and split loot by council with the numbers shown. New to Forever is fine.'
# The API's own summary string (api recruit card raid_nights.summary, server-written; lane C).
RECRUIT_NIGHTS = 'Wed, 8:00 pm to 11:00 pm CST'
RECRUIT_UPCOMING = [NEXT_RAID_AT, SECOND_RAID_AT]
# class, spec key, spec name, note
RECRUIT_WANTED = [('warrior', 'warrior-protection', 'Protection Warrior', 'Off tank'), ('paladin', 'paladin-holy', 'Holy Paladin', 'For the heal rotation'),
                  ('priest', 'priest-holy', 'Holy Priest', '')]
# id, name, class, spec, item level, best parse (metric, value, percentile), message, availability, votes yes / no / mine, sent
APPLICANTS = [(41, 'Ashenveil', 'paladin', 'Holy', 63, ('hps', 612.4, 91), 'Two tiers in a guild that disbanded.', 'Wed', 2, 0, 'yes', '2026-12-02T19:00:00Z'),
              (42, 'Quillbright', 'warlock', 'Destruction', 60, ('dps', 488.0, 74), 'New to Forever, level 60 since launch.', '', 1, 1, None, '2026-12-04T19:00:00Z')]
# id, name, class, spec, votes yes, decided by, role outcome, dm outcome
DECIDED = [(39, 'Hexmarrow', 'warlock', 'Destruction', 3, 'Ironbrand', 'granted', 'not_opted_in')]
DISCORD_ROLE_ID = '123456789012345678'
PROFILE = 'Frostquill'
PROFILE_ACCOUNT = 'u:1016'
# The profile raider's alt on the same account (the profile's Characters rail; not on the raid roster).
PROFILE_ALT = ('Quillstep', 'priest')
# The raid nights so far (UTC starts) and the outcome of each, oldest first: the first night only.
PROFILE_NIGHTS = [(LAST_RAID_AT, 'present')]
PROFILE_LAST_30 = (1, 1)
PROFILE_PARSE = ('dps', 812.0, 96)
PROFILE_GAP = {'upgrades': 2, 'gain_dps': 14.5, 'unit': 'dps', 'not_sim_checked': 0}
# The gap named (the profile's Next upgrades): slot, item id, the pick, where it comes from, its gain, its icon. The Alliance frost
# mage's 60 raid band of 1.60.1.70790 (bis/mage-frost.json): Sentinel's Silk Leggings over Skyshroud Leggings, Dryad's Wrist
# Bindings over Netherfroth Cuffs. The gains sum to PROFILE_GAP's.
PROFILE_UPGRADES = [('legs', 237815, "Sentinel's Silk Leggings", 'Illiyana Moonblaze', 8.6, 'inv_pants_cloth_05'),
                    ('wrist', 19595, "Dryad's Wrist Bindings", 'Silverwing Sentinels', 5.9, 'inv_bracer_13')]
# id, body, author, written, the viewer may delete it
PROFILE_NOTES = [(7, 'Ask about the Wednesday conflict.', 'Dunmaro', '2026-12-05T18:00:00Z', True), (6, 'Always brings flasks.', 'Ironbrand', '2026-12-02T18:00:00Z', False)]
