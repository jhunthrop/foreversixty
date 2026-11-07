"""The data the guild mocks and the guild parity spec share (design/parity.md, "The guild states").

ws_guild2.py draws the mocks from these values and guild_fixtures.py writes them as the API's answers into
web/tests/e2e/fixtures/guild-parity, which web/tests/e2e/support/guild-parity/ routes to the build. One roster, typed once.
Shape: api/cmd/seedguild (24 characters) and api/internal/guilds (sign-up states In / Late / Bench).
"""
from mocklib import CLASS_COLOR

# The fixtures' clock and zone (web/tests/e2e/support/guild-parity/clock.ts): Tue 8 Dec 2026, 4:00 pm CST, two days and four
# hours before Thursday's raid.
NOW = '2026-12-08T22:00:00Z'
ZONE = 'America/Chicago'
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
OFFICERS = ['Ironbrand', 'Shiverknife']
VIEWER = 'Dunmaro'
RECRUIT_PITCH = 'We raid two nights a week, bring flasks, and split loot by council with the numbers shown. New to Forever is fine.'
RECRUIT_NIGHTS = 'Tue and Thu, 8:00 pm to 11:00 pm CST'
RECRUIT_UPCOMING = ['2026-12-11T02:00:00Z', '2026-12-16T02:00:00Z']
# class, spec key, spec name, note
RECRUIT_WANTED = [('warrior', 'warrior-protection', 'Protection Warrior', 'Off tank'), ('paladin', 'paladin-holy', 'Holy Paladin', 'For the heal rotation'),
                  ('priest', 'priest-holy', 'Holy Priest', '')]
# id, name, class, spec, item level, best parse (metric, value, percentile), message, availability, votes yes / no / mine, sent
APPLICANTS = [(41, 'Ashenveil', 'paladin', 'Holy', 63, ('hps', 612.4, 91), 'Two tiers in a guild that disbanded.', 'Tue, Thu', 2, 0, 'yes', '2026-12-02T19:00:00Z'),
              (42, 'Quillbright', 'warlock', 'Destruction', 60, ('dps', 488.0, 74), 'New to Forever, level 60 since launch.', '', 1, 1, None, '2026-12-04T19:00:00Z')]
VOTERS = 3
# id, name, class, spec, votes yes, decided by, role outcome, dm outcome
DECIDED = [(39, 'Hexmarrow', 'warlock', 'Destruction', 3, 'Ironbrand', 'granted', 'not_opted_in')]
DISCORD_ROLE_ID = '123456789012345678'
PROFILE = 'Frostquill'
PROFILE_ACCOUNT = 'u:1016'
# The last four raid nights (UTC starts) and the outcome of each, oldest first.
PROFILE_NIGHTS = [('2026-11-24T02:00:00Z', 'present'), ('2026-11-26T02:00:00Z', 'present'), ('2026-12-01T02:00:00Z', 'absent'), ('2026-12-03T02:00:00Z', 'present')]
PROFILE_LAST_30 = (25, 30)
PROFILE_PARSE = ('dps', 812.0, 96)
PROFILE_GAP = {'upgrades': 2, 'gain_dps': 15.9, 'unit': 'dps', 'not_sim_checked': 0}
# item id, name, icon, quality, awarded
PROFILE_LOOT = [(18813, 'Ring of Binding', 'inv_jewelry_ring_13', 4, '2026-12-09T04:00:00Z')]
# id, body, author, written, the viewer may delete it
PROFILE_NOTES = [(7, 'Ask about the Thursday conflict.', 'Dunmaro', '2026-12-05T18:00:00Z', True), (6, 'Always brings flasks.', 'Ironbrand', '2026-12-02T18:00:00Z', False)]
