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
