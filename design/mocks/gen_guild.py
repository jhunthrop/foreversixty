"""Mock boards for the guild page rebuild, round 2 (design/specs/2026-10-04-guild-page.md v2).

  python3 design/mocks/gen_guild.py
    -> guild-overview-officer.html  (1440, officer, Overview tab)
    -> guild-roster.html            (1440, officer, Roster tab, filters + rating columns)
    -> guild-raids.html             (1440, officer, Raids tab, one night expanded)
    -> guild-progression.html       (1440, Progression tab)
    -> guild-readiness.html         (1440, officer, Readiness tab)
    -> guild-loot.html              (1440, officer, Loot tab, Onyxia selected)
    -> guild-member.html            (1440, verified member, Overview tab, pinned row)
    -> guild-phone.html             (390, verified member, Readiness tab)
  guild-public.html/.png (round 1, signed out, Overview) is kept unchanged -- not regenerated.

No real guild data exists yet -- a database seed (api/cmd/seedguild, a separate lane) is
building one from the same picture this file hand-writes: 24 raiders, item level 55 to 70,
raiding Barrow Deeps, Hyjal Summit and Onyxia's Lair (Forever's real first tier, opening 9
December 2026 -- data/curated/loot/forever-raid-phases.json's own notes field, read directly
for this round; Molten Core is a later, unannounced-date Era raid in that same file and is
no longer used anywhere in this mock, round 1's own accuracy defect, see the spec's round-2
amendment). Every board says so in its own caption.

Every derived stat (rating, attendance, parses, gear gap, professions, consumables, unspent
points) is computed ONCE from the roster/nights constants below by a small set of seeded
functions, never hand-typed per board -- so no two boards can disagree (the round-1 lesson,
restated by the coordinator for round 2). Onyxia's drop list (`ONYXIA_DROPS`) is the real
raid:onyxias-lair boss entry from data/builds/1.60.1.70009/loot.json, item names read from
that build's own items.json -- not invented.
"""
import hashlib

from mocklib import (
    GOLD, MUTED, TEXT, BODY, BORDER, SOFT, CLASS_COLOR, GREEN,
    crest, emblem, nav, nav_phone, footer, page, write,
)

GUILD = {'name': 'Olympus XXVII', 'region': 'us', 'ruleset': 'pvp'}
FACTION = 'horde'
RED = '#ff6b5c'
BLUE = '#6fb1ff'

# ---------------------------------------------------------------------------
# Mock roster -- 24 Horde raiders. Shape matches RosterRow (api/internal/guilds/home.go)
# exactly: character_key, region, ruleset, name, class, spec, rank, verified,
# logged_recently, consent, item_level?, may_remove. `role` and `account` are this board's
# own derived additions (round 2's Roster/Readiness filters and main/alt grouping), not new
# API fields by themselves -- role is DERIVABLE from spec (SPEC_ROLE below, the same map the
# guides/BiS pages already use), account is EXISTS already (guild_characters.user_id).
# ---------------------------------------------------------------------------
ROSTER = [
    # name, class, spec, rank, verified, logged, consent, ilvl
    ('Kraggor', 'warrior', 'Protection', 'leader', True, True, 'gear_bags', 68),
    ('Obnoxious Yell', 'warrior', 'Fury', 'officer', True, True, 'gear_bags', 66),
    ('Sunderfel', 'warrior', 'Arms', 'officer', True, True, 'gear', 61),
    ('Grimtotem', 'warrior', 'Fury', 'member', True, False, 'gear', 58),
    ('Zulmara', 'hunter', 'Marksmanship', 'member', True, True, 'gear_bags', 64),
    ('Windtalon', 'hunter', 'Marksmanship', 'member', True, True, 'gear_bags', 70),
    ('Duskstrider', 'hunter', 'Marksmanship', 'member', True, False, 'gear', 55),
    ('Felsnap', 'hunter', 'Beast Mastery', 'member', True, True, 'gear', 60),
    ('Vexlash', 'rogue', 'Combat', 'officer', True, True, 'gear_bags', 65),
    ('Shadowquill', 'rogue', 'Assassination', 'member', True, True, 'gear', 59),
    ('Nixthrottle', 'rogue', 'Combat', 'member', True, False, 'roster', None),
    ('Hexbramble', 'warlock', 'Affliction', 'member', True, True, 'gear', 63),
    ('Soulgrave', 'warlock', 'Destruction', 'member', True, True, 'gear', 57),
    ('Pyrewisp', 'mage', 'Fire', 'member', True, True, 'gear_bags', 69),
    ('Frostnettle', 'mage', 'Frost', 'member', True, False, 'gear', 56),
    ('Arcanemoor', 'mage', 'Arcane', 'member', True, True, 'gear', 62),
    ('Lightbrand', 'priest', 'Holy', 'member', True, True, 'gear_bags', 67),
    ('Grimvow', 'priest', 'Shadow', 'member', True, True, 'roster', None),
    ('Mendwhisper', 'priest', 'Holy', 'member', True, False, 'gear', 60),
    ('Earthhoof', 'druid', 'Restoration', 'member', True, True, 'gear_bags', 65),
    ('Thornhide', 'druid', 'Feral', 'member', True, True, 'gear', 58),
    ('Rootgall', 'shaman', 'Restoration', 'member', False, True, 'gear', 61),
    ('Stormtusk', 'shaman', 'Elemental', 'member', False, False, 'gear', 55),
    ('Fulmintide', 'shaman', 'Enhancement', 'member', False, True, 'gear', 59),
]

SPEC_ROLE = {
    'Protection': 'Tank', 'Fury': 'DPS', 'Arms': 'DPS',
    'Marksmanship': 'DPS', 'Beast Mastery': 'DPS',
    'Combat': 'DPS', 'Assassination': 'DPS',
    'Affliction': 'DPS', 'Destruction': 'DPS',
    'Fire': 'DPS', 'Frost': 'DPS', 'Arcane': 'DPS',
    'Holy': 'Healer', 'Shadow': 'DPS',
    'Restoration': 'Healer', 'Feral': 'DPS',
    'Elemental': 'DPS', 'Enhancement': 'DPS',
}
# guild_characters.user_id grouping (EXISTS) -- two accounts below run a second character.
# Shown so the Roster tab's "main/alt" column has something real to group, not an invented
# feature with nothing behind it.
ALT_OF = {'Duskstrider': 'Zulmara', 'Grimtotem': 'Kraggor'}

VIEWER_MEMBER = 'Zulmara'   # verified, Marksmanship Hunter, ranked among peers
VIEWER_OFFICER = 'Kraggor'  # leader rank -- sees every officer control

PROFESSIONS_BY_CLASS = {
    'warrior': ('Blacksmithing', 'Mining'), 'hunter': ('Leatherworking', 'Skinning'),
    'rogue': ('Engineering', 'Mining'), 'warlock': ('Tailoring', 'Enchanting'),
    'mage': ('Tailoring', 'Enchanting'), 'priest': ('Tailoring', 'Enchanting'),
    'druid': ('Herbalism', 'Alchemy'), 'shaman': ('Leatherworking', 'Enchanting'),
}

# ---------------------------------------------------------------------------
# Raid nights -- Forever's real first tier (data/curated/loot/forever-raid-phases.json,
# read 2026-10-04: "the first tier opens on 9 December and is Barrow Deeps, Hyjal Summit and
# Onyxia", confirmed again against the rating-design spec's own "Barrow Deeps and Hyjal
# Summit run on Output/Preparation/Activity only"). Barrow Deeps and Hyjal Summit have NO
# sourced encounter names or loot yet (same file: "nothing sourced maps the announced names
# onto those ids, and neither database gives them an item") -- their nights log real pulls
# with kill=false and encounter_id=null, same as any fight the log engine cannot name; this
# spec shows that honestly (unidentified pulls) rather than inventing boss names. Onyxia's
# Lair is the one zone with a real, sourced boss (raid:onyxias-lair, npc 10184) and loot
# table, so it is the only named encounter anywhere on these boards.
# ---------------------------------------------------------------------------
NIGHTS = [
    # zone, date, duration_min, pulls, kills (named-encounter kills only)
    ('Barrow Deeps', '2026-12-10', 118, 6, 0),
    ('Hyjal Summit', '2026-12-11', 96, 5, 0),
    ('Barrow Deeps', '2026-12-13', 134, 7, 0),
    ("Onyxia's Lair", '2026-12-15', 41, 3, 1),
    ('Hyjal Summit', '2026-12-16', 102, 6, 0),
    ('Barrow Deeps', '2026-12-18', 151, 8, 0),
    ("Onyxia's Lair", '2026-12-20', 22, 2, 1),
    ('Hyjal Summit', '2026-12-22', 109, 7, 0),
]
UPDATED_LABEL = 'Updated Dec 22'
TOTAL_PULLS = sum(n[3] for n in NIGHTS)
NAMED_KILLS = sum(n[4] for n in NIGHTS)  # Onyxia, twice (first kill + a farm repeat)

# Onyxia's own pull-by-pull history across her two raid nights (the one boss this tier a
# trend can be shown for honestly).
ONYXIA_PULLS = [
    # night date, pull #, duration_s, result
    ('2026-12-15', 1, 95, 'wipe'), ('2026-12-15', 2, 210, 'wipe'), ('2026-12-15', 3, 268, 'kill'),
    ('2026-12-20', 1, 240, 'wipe'), ('2026-12-20', 2, 251, 'kill'),
]

# Onyxia's real drop table (data/builds/1.60.1.70009/loot.json, raid:onyxias-lair, npc
# 10184 -- 22 items total; the eight below are the ones that land on a class this roster
# actually has, item names from that build's own items.json, not invented).
ONYXIA_DROPS = [
    # item, slot, classes it fits (role flavour, not a hard restriction)
    ('Helm of Wrath', 'Head', ['warrior']),
    ("Dragonstalker's Helm", 'Head', ['hunter']),
    ('Netherwind Crown', 'Head', ['mage']),
    ('Halo of Transcendence', 'Head', ['priest']),
    ('Nemesis Skullcap', 'Head', ['warlock']),
    ("Stormrage Cover", 'Head', ['druid']),
    ("Eskhandar's Collar", 'Neck', ['warrior', 'hunter', 'rogue']),
    ("Vis'kag the Bloodletter", 'One-Hand', ['rogue', 'warrior']),
]

# ---------------------------------------------------------------------------
# Deterministic derived stats -- one seeded function per figure, so every board reads the
# same number for the same character without a hand-authored table to keep in sync.
# ---------------------------------------------------------------------------
ROSTER_BY_NAME = {r[0]: r for r in ROSTER}
ILVLS = [r[7] for r in ROSTER if r[7] is not None]
MEDIAN_ILVL = sorted(ILVLS)[len(ILVLS) // 2]

RATING_COMPONENTS = ['Output', 'Survival', 'Mechanics', 'Utility', 'Preparation', 'Activity']


def seed(name: str, salt: str) -> float:
    h = hashlib.md5(f'{name}:{salt}'.encode()).hexdigest()
    return int(h[:8], 16) / 0xFFFFFFFF


def role_for(name: str) -> str:
    return SPEC_ROLE[ROSTER_BY_NAME[name][2]]


def rating_for(name: str) -> dict:
    comps = {c: round(50 + seed(name, c) * 45) for c in RATING_COMPONENTS}
    comps['Overall'] = round(sum(comps.values()) / len(comps))
    return comps


def attendance_for(name: str) -> int:
    return 5 + int(seed(name, 'attendance') * 4)  # 5..8 of the last 8 nights


def parses_for(name: str) -> dict:
    metric = 'HPS' if role_for(name) == 'Healer' else 'DPS'
    best = 230 + seed(name, 'best') * 180
    avg = best * (0.72 + seed(name, 'avg') * 0.18)
    return {'metric': metric, 'best': best, 'avg': avg}


def gear_gap_for(name: str) -> dict:
    ilvl = ROSTER_BY_NAME[name][7]
    if ilvl is None:
        return {'upgrades': None, 'gain': None}
    deficit = max(0, 70 - ilvl)
    upgrades = min(7, round(deficit / 3) + int(seed(name, 'upgrades') * 2))
    gain = round(deficit * (2.0 + seed(name, 'gain') * 2.4), 1)
    return {'upgrades': upgrades, 'gain': gain}


ENCHANT_SLOTS = ['Weapon', 'Chest', 'Cloak', 'Boots']


def missing_enchants_for(name: str) -> list:
    n = int(seed(name, 'enchant-count') * 3)  # 0, 1 or 2
    pool = sorted(ENCHANT_SLOTS, key=lambda s: seed(name, 'enchant-' + s))
    return pool[:n]


def consumables_ok_for(name: str) -> bool:
    return seed(name, 'consumables') > 0.35


def unspent_points_for(name: str) -> int:
    return int(seed(name, 'talent-points') * 3.2)  # 0..3


def professions_for(name: str) -> tuple:
    return PROFESSIONS_BY_CLASS[ROSTER_BY_NAME[name][1]]


def last_synced_for(name: str) -> str:
    return 'Synced 40m ago' if ROSTER_BY_NAME[name][5] else 'Synced 4 days ago'


def last_log_for(name: str) -> str:
    return NIGHTS[-1][1] if ROSTER_BY_NAME[name][5] else NIGHTS[-3][1]


def readiness_fails_for(name: str) -> list:
    """The checks this one character fails right now -- feeds both the Readiness row and
    the member Overview's own "before Thursday" sentence."""
    fails = []
    gap = gear_gap_for(name)
    if gap['upgrades'] and gap['upgrades'] >= 3:
        fails.append(f"{gap['upgrades']} gear upgrades waiting ({gap['gain']} DPS)")
    missing = missing_enchants_for(name)
    if missing:
        fails.append(f"no enchant: {', '.join(missing)}")
    if not consumables_ok_for(name):
        fails.append('bags short on consumables')
    pts = unspent_points_for(name)
    if pts > 0:
        fails.append(f'{pts} unspent talent point{"s" if pts > 1 else ""}')
    return fails


def readiness_score(name: str) -> int:
    """Worst-first sort key for the Readiness tab -- more fails and a bigger gear gain sort
    first, the same "worst-first" the task asks for."""
    gap = gear_gap_for(name)
    return len(readiness_fails_for(name)) * 100 + (gap['gain'] or 0)


# ---------------------------------------------------------------------------
PANEL = (f'background:var(--raised);border:1px solid {BORDER};border-radius:6px;'
         f'padding:18px 20px;display:flex;flex-direction:column;gap:12px')

# Round-2 fix (coordinator, 2026-10-04): the Roster tab's "Open in simulator · Open in
# planner" pair was the loudest thing in the row. One shared quiet-link style, injected once
# per page by build_board, used wherever a row-end hand-off link needs to stay out of the
# way of the name/pills/figures it sits beside.
QUIET_LINK_STYLE = ('<style>.quiet-link{color:var(--muted);font-size:12px;font-weight:600;'
                     'text-decoration:none}.quiet-link:hover,.quiet-link:focus-visible{color:var(--gold)}</style>')


def crest_img(slug: str, size: int) -> str:
    return f'<img class="crest" src="{crest(slug)}" alt="" style="width:{size}px;height:{size}px;--c:{CLASS_COLOR[slug]}">'


def pill(text: str, color: str = GOLD, bg: str | None = None, border: str | None = None) -> str:
    bg = bg or f'color-mix(in srgb, {color} 14%, transparent)'
    border = border or f'color-mix(in srgb, {color} 35%, transparent)'
    return f'<span class="pill" style="color:{color};background:{bg};border-color:{border};height:22px;padding:0 8px;white-space:nowrap">{text}</span>'


def btn(lbl: str, h: int = 40, gold: bool = False) -> str:
    cls = 'btn' + (' btn-gold' if gold else '')
    return f'<a class="{cls}" href="#{lbl.lower().replace(" ", "-")}" style="height:{h}px">{lbl}</a>'


def row_btn(lbl: str, h: int = 32) -> str:
    return (f'<button style="display:inline-flex;align-items:center;justify-content:center;'
            f'height:{h}px;padding:0 12px;border:1px solid #3a3326;border-radius:4px;'
            f'background:none;color:var(--strong);font-family:inherit;font-size:11px;font-weight:700;'
            f'letter-spacing:.06em;text-transform:uppercase;white-space:nowrap;cursor:pointer">{lbl}</button>')


def h2(text: str) -> str:
    return f'<span class="display" style="font-size:18px;text-transform:uppercase;letter-spacing:.04em;color:{TEXT}">{text}</span>'


def label(text: str) -> str:
    return f'<span class="label">{text}</span>'


def annotation(label_text: str) -> str:
    """A data-source annotation for this mock board only -- round-2 fix (coordinator,
    2026-10-04): this used to be a coloured pill inside the panel header, which read as real
    page chrome rather than a board note. It is now plain grey italic caption text, the same
    style `mock_caption()` already uses for "Mock roster -- 24 characters..." -- placed under
    each panel's own content, never inside its header, and never rendered on the real page
    (stated explicitly in the spec's §11)."""
    return f'<p style="margin:0;font-size:12px;color:{MUTED};font-style:italic">{label_text}</p>'


def annotation_inline(label_text: str) -> str:
    """The same annotation, for the handful of panels that already fold their source note
    into a sentence (Readiness's and Loot's own footnotes, Progression's Onyxia panel) --
    plain italic text with no colour, inheriting the sentence's own muted styling, instead of
    a coloured pill sitting mid-sentence."""
    return f'<em style="font-style:italic">{label_text}</em>'


def panel_header(title: str, extra: str = '') -> str:
    return (f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px;flex-wrap:wrap">'
            f'{h2(title)}{extra}</div>')


def mock_caption() -> str:
    return (f'<p style="margin:0;font-size:12px;color:{MUTED};font-style:italic">'
            f'Mock roster &mdash; 24 characters, no real guild data yet. Raids: Barrow Deeps, '
            f'Hyjal Summit, Onyxia&rsquo;s Lair &mdash; Forever&rsquo;s real first tier, open '
            f'9 December 2026.</p>')


def rank_among_peers(viewer_name: str):
    viewer = ROSTER_BY_NAME[viewer_name]
    _, cls, spec, *_ = viewer
    peers = sorted(
        [r for r in ROSTER if r[1] == cls and r[2] == spec and r[7] is not None],
        key=lambda r: -r[7],
    )
    rank = next(i for i, r in enumerate(peers, 1) if r[0] == viewer_name)
    return rank, len(peers), viewer


# ---------------------------------------------------------------------------
# Header band + tab strip
# ---------------------------------------------------------------------------
TABS = ['Overview', 'Roster', 'Raids', 'Progression', 'Readiness', 'Loot', 'Settings']
# Tab visibility by viewer role -- stated once here, restated in the spec's own matrix
# (round 2 §4.0): a public visitor never reaches Roster/Readiness/Loot/Settings.
TABS_FOR_ROLE = {
    'public': ['Overview', 'Raids', 'Progression'],
    'member': ['Overview', 'Roster', 'Raids', 'Progression', 'Readiness', 'Loot'],
    'officer': TABS,
}


def tab_strip(active: str, role: str, phone: bool = False) -> str:
    tabs = TABS_FOR_ROLE[role]
    pad = '0 14px' if not phone else '0 12px'
    cells = []
    for t in tabs:
        is_active = t == active
        st = (f'border-color:{GOLD};background:rgba(229,185,85,.08);color:{GOLD}'
              if is_active else f'border-color:{BORDER};color:{TEXT}')
        suffix = ' <span style="font-size:10px;color:' + MUTED + '">&middot; read-only</span>' if (t == 'Loot' and role == 'member') else ''
        cells.append(f'<a href="#{t.lower()}" style="display:inline-flex;align-items:center;height:44px;padding:{pad};'
                      f'border:1px solid;border-radius:4px;font-size:12px;font-weight:700;letter-spacing:.06em;'
                      f'text-transform:uppercase;white-space:nowrap;{st}">{t}{suffix}</a>')
    overflow = 'overflow-x:auto' if phone else ''
    return f'<div style="display:flex;gap:8px;padding:0 {"18px" if phone else "48px"};{overflow}">{"".join(cells)}</div>'


def header_facts() -> str:
    return (f'<span class="mono" style="color:{TEXT}">PvP US</span>'
            f' &middot; <span class="mono" style="color:{TEXT}">{NAMED_KILLS}</span> named encounters down'
            f' &middot; <span class="mono" style="color:{TEXT}">{TOTAL_PULLS}</span> pulls this tier')


def standing_line(viewer_name: str, with_fails: bool = False) -> str:
    rank, total, viewer = rank_among_peers(viewer_name)
    name, cls, spec = viewer[0], viewer[1], viewer[2]
    color = CLASS_COLOR[cls]
    if total == 1:
        text = f'{name} &middot; {spec} {cls.capitalize()} &middot; the only {spec} {cls.capitalize()} in this guild this week'
    else:
        ilvl = viewer[7]
        text = f'{name} &middot; {rank} of {total} {spec} {cls.capitalize()}s this week by item level ({ilvl} ilvl)'
    line = (f'<div style="display:flex;align-items:center;gap:12px;flex-wrap:wrap">'
            f'<span style="font-size:14px;color:{color};font-weight:600">{text}</span>'
            f'<a href="#gear-in-the-planner" style="font-size:13px;font-weight:700">Gear in the planner</a></div>')
    if with_fails:
        fails = readiness_fails_for(viewer_name)
        if fails:
            needs = (f'<p style="margin:0;font-size:13px;color:{RED}">What the guild needs from you before '
                     f'Thursday: {"; ".join(fails)}.</p>')
        else:
            needs = f'<p style="margin:0;font-size:13px;color:{GREEN}">Every readiness check passes. Nothing needed before Thursday.</p>'
        line += needs
    return line


def unverified_note() -> str:
    return (f'<p style="margin:0;font-size:14px;color:{TEXT}">You are not verified yet. '
            f'An officer can approve you from the roster below.</p>')


def signin_prompt() -> str:
    return (f'<div style="display:flex;flex-direction:column;align-items:flex-start;gap:10px">'
            f'<p style="margin:0;font-size:14px;color:{TEXT}">See where you stand in this guild. '
            f'<a href="#login">Use an email link instead.</a></p>'
            f'{btn("Sign in with Battle.net")}</div>')


def header(role_slot: str, tabs_html: str, phone: bool = False) -> str:
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    return f'''<div style="display:flex;flex-direction:column;gap:14px;padding:{pad};{maxw}">
<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Guild</span>
<h1 class="display" style="margin:0;font-size:22px;color:{TEXT}">{GUILD["name"]}</h1>
<span style="font-size:12px;color:{MUTED}">{UPDATED_LABEL}</span>
<p style="margin:0;font-size:14px;color:{BODY}">{header_facts()}</p>
{role_slot}
</div>
<div style="margin-top:18px">{tabs_html}</div>'''


# ---------------------------------------------------------------------------
# Officer tools strip (Overview summary + Settings tab both use this)
# ---------------------------------------------------------------------------
def officer_tools(compact: bool = True) -> str:
    waiting = sum(1 for r in ROSTER if not r[4])
    see_all = '' if compact else ''
    return f'''<div style="{PANEL};flex-direction:row;align-items:center;justify-content:space-between;gap:20px;flex-wrap:wrap">
<div style="display:flex;align-items:center;gap:10px">{pill("Claimed")}<span style="font-size:13px;color:{MUTED}">by Kraggor</span></div>
<a href="#guild-roster-unverified" style="font-size:13px;font-weight:700;color:{GOLD}">{waiting} waiting for approval</a>
<div style="display:flex;align-items:center;gap:12px;flex-wrap:wrap">
<span class="label" style="color:{MUTED}">Invite link</span>
<span style="font-size:12px;color:{MUTED}">Anyone with this link can join as a member.</span>
{btn("Rotate invite link")}
</div></div>'''


# ---------------------------------------------------------------------------
# OVERVIEW tab -- round 1's page, trimmed to a summary of every other tab, "See all" each.
# ---------------------------------------------------------------------------
def see_all(tab: str) -> str:
    return f'<a href="#{tab.lower()}" style="font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase">See all</a>'


def overview_roster_summary() -> str:
    waiting = sum(1 for r in ROSTER if not r[4])
    below_floor = sum(1 for r in ROSTER if r[7] is not None and r[7] < MEDIAN_ILVL - 5)
    rows = (f'<span class="mono" style="font-size:13px;color:{TEXT}">{len(ROSTER)}</span> raiders &middot; '
            f'<span class="mono" style="font-size:13px;color:{GOLD}">{waiting}</span> waiting for approval &middot; '
            f'<span class="mono" style="font-size:13px;color:{TEXT}">{below_floor}</span> below the rating floor')
    return f'<div style="{PANEL}">{panel_header("Roster", see_all("Roster"))}<p style="margin:0;font-size:14px;color:{BODY}">{rows}</p>{annotation("EXISTS")}</div>' 


def overview_raids_summary() -> str:
    rows = ''.join(
        f'<div style="display:flex;align-items:center;gap:12px;padding:7px 4px;border-bottom:1px solid {SOFT}">'
        f'<span style="font-size:14px;font-weight:600;flex:1">{zone}</span>'
        f'<span class="mono" style="font-size:13px;color:{MUTED}">{date}</span>'
        f'<span class="mono" style="font-size:13px;color:{MUTED}">{kills} kill{"s" if kills != 1 else ""} &middot; {pulls - kills} wipe{"s" if (pulls - kills) != 1 else ""}</span></div>'
        for zone, date, dur, pulls, kills in NIGHTS[-3:][::-1]
    )
    return f'<div style="{PANEL}">{panel_header("This week\'s raid nights", see_all("Raids"))}{rows}{annotation("EXISTS")}</div>' 


def overview_progression_summary() -> str:
    text = (f'<span class="mono" style="color:{TEXT}">{NAMED_KILLS}</span> named encounter{"s" if NAMED_KILLS != 1 else ""} down this tier '
            f'(Onyxia). Barrow Deeps and Hyjal Summit have no published encounter list yet '
            f'&mdash; {TOTAL_PULLS - 5} pulls logged there regardless.')
    return f'<div style="{PANEL}">{panel_header("Progression", see_all("Progression"))}<p style="margin:0;font-size:14px;color:{BODY}">{text}</p>{annotation("EXISTS + DERIVABLE")}</div>' 


def overview_readiness_summary() -> str:
    worst = sorted(ROSTER, key=lambda r: -readiness_score(r[0]))[:3]
    rows = ''.join(
        f'<div style="display:flex;align-items:center;gap:10px;padding:5px 0">'
        f'<span style="font-size:13px;color:{CLASS_COLOR[r[1]]};font-weight:600">{r[0]}</span>'
        f'<span style="font-size:12px;color:{MUTED}">{len(readiness_fails_for(r[0]))} checks failing</span></div>'
        for r in worst
    )
    return f'<div style="{PANEL}">{panel_header("Readiness", see_all("Readiness"))}{rows}{annotation("DERIVABLE + NEEDS NEW DATA")}</div>' 


def overview_loot_summary() -> str:
    text = f'Onyxia&rsquo;s Lair &middot; 8 class drops across this roster &middot; next pick: {ONYXIA_DROPS[0][0]}'
    return f'<div style="{PANEL}">{panel_header("Loot", see_all("Loot"))}<p style="margin:0;font-size:14px;color:{BODY}">{text}</p>{annotation("EXISTS + DERIVABLE")}</div>' 


def overview_tab(role: str, viewer_name: str | None) -> str:
    sections = []
    if role == 'officer':
        sections.append(officer_tools())
    sections.append(overview_roster_summary() if role != 'public' else '')
    sections.append(overview_raids_summary())
    sections.append(overview_progression_summary())
    if role != 'public':
        sections.append(overview_readiness_summary())
        sections.append(overview_loot_summary())
    return ''.join(s for s in sections if s)


# ---------------------------------------------------------------------------
# ROSTER tab -- round 1's table, plus rating, attendance, parses, professions, main/alt,
# last log, filters.
# ---------------------------------------------------------------------------
def filter_bar() -> str:
    def f(lbl):
        return (f'<button style="height:32px;padding:0 12px;border:1px solid {BORDER};border-radius:4px;'
                f'background:none;color:{TEXT};font-size:12px;font-weight:600">{lbl}</button>')
    return (f'<div style="display:flex;gap:8px;flex-wrap:wrap;align-items:center">'
            f'<span class="label" style="color:{MUTED}">Filter</span>'
            f'{f("Role: All")}{f("Class: All")}{f("Verified only")}{f("Below rating floor")}</div>')


def sort_header(active: str) -> str:
    cols = [('Rank', 'rank'), ('Item level', 'ilvl'), ('Rating', 'rating'), ('Attendance', 'attendance'), ('Last seen', 'seen')]
    cells = []
    for lbl, key in cols:
        arrow = ' &#9662;' if key == active else ''
        color = GOLD if key == active else MUTED
        cells.append(f'<button style="background:none;border:none;height:36px;padding:0;font-size:11px;font-weight:700;'
                      f'letter-spacing:.1em;text-transform:uppercase;color:{color};cursor:pointer">{lbl}{arrow}</button>')
    return f'<div style="display:flex;gap:24px;padding:0 4px;flex-wrap:wrap">{"".join(cells)}</div>'


def rating_cell(name: str) -> str:
    r = rating_for(name)
    tooltip = ' &middot; '.join(f'{c} {r[c]}' for c in RATING_COMPONENTS)
    return (f'<span class="mono" style="font-size:13px;color:{TEXT};cursor:help" title="{tooltip}">'
            f'{r["Overall"]} <span style="color:{MUTED};font-size:11px">rating</span></span>')


def roster_row(row: tuple, officer: bool, pinned: bool = False, wide: bool = True) -> str:
    name, cls, spec, rank, verified, logged, consent, ilvl = row
    color = CLASS_COLOR[cls]
    rank_pill = pill(rank.capitalize()) if rank in ('officer', 'leader') else ''
    unverified_pill = pill('Unverified', MUTED) if not verified else ''
    logged_pill = pill('Logged in the last day', GREEN, 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)') if logged else ''
    ilvl_cell = f'<span class="mono" style="font-size:13px;color:{TEXT}">ilvl {ilvl}</span>' if ilvl is not None else ''
    # Round-2 fix (coordinator, 2026-10-04): these were bold 13px gold links, the loudest
    # thing in the row -- louder than the raider's own name. Quiet now: 12px, muted, gold
    # only on hover (.quiet-link, defined once in QUIET_LINK_STYLE and injected into every
    # page that uses it), so the name/pills/figures lead and this pair reads as a footnote
    # action, which is what it is on every other row-level place the site already uses it.
    handoff = (f'<a href="#open-in-sim" class="quiet-link">Open in simulator</a>'
               f'<a href="#open-in-planner" class="quiet-link">Open in planner</a>') if consent != 'roster' else ''
    actions = ''
    if officer:
        if not verified:
            actions += row_btn('Approve')
        elif rank != 'leader':
            actions += row_btn('Remove')
    alt_tag = f'<span style="font-size:11px;color:{MUTED}">alt of {ALT_OF[name]}</span>' if name in ALT_OF else ''
    extra = ''
    if wide:
        parses = parses_for(name)
        profs = professions_for(name)
        extra = (f'<span class="mono" style="font-size:12px;color:{MUTED}">{attendance_for(name)}/8 nights</span>'
                 f'<span class="mono" style="font-size:12px;color:{MUTED}">{parses["best"]:.0f} best {parses["metric"]}</span>'
                 f'<span style="font-size:11px;color:{MUTED}">{profs[0]} / {profs[1]}</span>'
                 f'{rating_cell(name)}')
    bg = 'background:rgba(229,185,85,.06);' if pinned else ''
    return (f'<div style="display:flex;align-items:center;gap:12px;flex-wrap:wrap;padding:9px 4px;{bg}'
            f'border-bottom:1px solid {SOFT}">'
            f'{crest_img(cls, 36)}'
            f'<div style="display:flex;flex-direction:column;min-width:150px">'
            f'<span class="display" style="font-size:14px;font-weight:700;color:{color}">{name}{" (you)" if pinned else ""}</span>'
            f'<span style="font-size:12px;color:{MUTED}">{spec} {alt_tag}</span></div>'
            f'{rank_pill}{unverified_pill}{logged_pill}{ilvl_cell}{extra}'
            f'<div style="display:flex;gap:10px;margin-left:auto">{handoff}{actions}</div></div>')


def roster_tab(officer: bool, pin: str | None = None, with_filters: bool = True, wide: bool = True) -> str:
    unverified = [r for r in ROSTER if not r[4]]
    rest = [r for r in ROSTER if r[4]]
    if pin:
        rest = sorted(rest, key=lambda r: (r[0] != pin, -(r[7] or 0)))
    else:
        rest = sorted(rest, key=lambda r: -(r[7] or 0))
    waiting = len(unverified)
    approve_all = (
        f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px">'
        f'<span style="font-size:13px;color:{MUTED}">{waiting} waiting for approval</span>'
        + row_btn('Approve all', h=36) + '</div>'
    ) if officer and waiting else ''
    unverified_rows = ''.join(roster_row(r, officer, wide=wide) for r in unverified)
    verified_rows = ''.join(roster_row(r, officer, pinned=(r[0] == pin), wide=wide) for r in rest)
    filters = filter_bar() if with_filters else ''
    return f'''<div style="{PANEL}">
{panel_header("Roster")}
{filters}
{approve_all}
<div id="guild-roster-unverified" style="display:flex;flex-direction:column">{unverified_rows}</div>
<div style="margin-top:6px">{sort_header("ilvl")}</div>
<div style="display:flex;flex-direction:column">{verified_rows}</div>
{annotation("EXISTS + DERIVABLE (rating/attendance/parses/professions are new queries over stored rows)")}
</div>'''


# ---------------------------------------------------------------------------
# RAIDS tab -- every night as a row; one expanded to its fight list.
# ---------------------------------------------------------------------------
def raid_row(night: tuple, expanded: bool) -> str:
    zone, date, dur, pulls, kills = night
    wipes = pulls - kills
    head = (f'<div style="display:flex;align-items:center;gap:12px;padding:9px 4px;flex-wrap:wrap;'
            f'border-bottom:1px solid {SOFT if not expanded else BORDER}">'
            f'<a href="#raid" style="font-size:14px;font-weight:600;min-width:140px">{zone}</a>'
            f'<span class="mono" style="font-size:13px;color:{MUTED}">{date}</span>'
            f'<span class="mono" style="font-size:13px;color:{MUTED}">{dur} min</span>'
            f'<span class="mono" style="font-size:13px;color:{MUTED};margin-left:auto">{kills} named kill{"s" if kills != 1 else ""} &middot; {wipes} wipe{"s" if wipes != 1 else ""}</span>'
            f'</div>')
    if not expanded:
        return head
    if zone == "Onyxia's Lair":
        pulls_for_night = [p for p in ONYXIA_PULLS if p[0] == date]
        fight_rows = ''.join(
            f'<div style="display:flex;align-items:center;gap:12px;padding:6px 12px;border-bottom:1px solid {SOFT}">'
            f'<span style="font-size:13px;color:{TEXT};min-width:100px">Onyxia &middot; pull {n}</span>'
            f'<span class="mono" style="font-size:12px;color:{MUTED}">{d // 60}:{d % 60:02d}</span>'
            f'{pill("Kill", GREEN, "rgba(30,255,0,.10)", "rgba(30,255,0,.25)") if res == "kill" else pill("Wipe", MUTED)}'
            f'</div>' for _, n, d, res in pulls_for_night
        )
    else:
        fight_rows = ''.join(
            f'<div style="display:flex;align-items:center;gap:12px;padding:6px 12px;border-bottom:1px solid {SOFT}">'
            f'<span style="font-size:13px;color:{TEXT};min-width:100px">Pull {n}</span>'
            f'<span style="font-size:12px;color:{MUTED}">no named encounter published</span>'
            f'<span class="mono" style="font-size:12px;color:{MUTED};margin-left:auto">{pill("Wipe", MUTED)}</span>'
            f'</div>' for n in range(1, pulls + 1)
        )
    present = sorted(ROSTER, key=lambda r: -(r[7] or 0))[:14]
    present_row = ' '.join(crest_img(r[1], 22) for r in present) + f' <span style="font-size:12px;color:{MUTED}">{len(present)} of {len(ROSTER)} present</span>'
    top_parse = max(ROSTER, key=lambda r: parses_for(r[0])['best'])
    top_parse_line = (f'Top parse: <span style="color:{CLASS_COLOR[top_parse[1]]};font-weight:600">{top_parse[0]}</span> '
                       f'{parses_for(top_parse[0])["best"]:.0f} {parses_for(top_parse[0])["metric"]}')
    return (head + f'<div style="padding:10px 0 4px 0;display:flex;flex-direction:column;gap:10px">'
            f'<div style="display:flex;align-items:center;flex-wrap:wrap;gap:3px">{present_row}</div>'
            f'<p style="margin:0;font-size:13px;color:{BODY}">{top_parse_line}</p>'
            f'<div style="display:flex;flex-direction:column">{fight_rows}</div>'
            f'<p style="margin:0;font-size:11px;color:{MUTED}">Attach to this guild / visibility controls are officer-only (unchanged from round 1 §4.D).</p>'
            f'</div>')


def raids_tab(expand_date: str) -> str:
    rows = ''.join(raid_row(n, n[1] == expand_date) for n in NIGHTS[::-1])
    return f'<div style="{PANEL}">{panel_header("Raids")}{rows}{annotation("EXISTS (reports + fights)")}</div>' 


# ---------------------------------------------------------------------------
# PROGRESSION tab -- tier bar, Onyxia's own pull/kill-time trend, deaths, best parse.
# ---------------------------------------------------------------------------
def progression_tab() -> str:
    bar = (f'<div style="display:flex;flex-direction:column;gap:6px">'
           f'<div style="display:flex;justify-content:space-between;font-size:13px;color:{MUTED}">'
           f'<span>Named encounters down this tier</span><span class="mono" style="color:{TEXT}">{NAMED_KILLS} &middot; Onyxia</span></div>'
           f'<div style="height:6px;background:{SOFT};border-radius:3px;overflow:hidden">'
           f'<div style="height:100%;width:60%;background:{GOLD}"></div></div>'
           f'<p style="margin:0;font-size:12px;color:{MUTED}">Barrow Deeps and Hyjal Summit have no published encounter list yet '
           f'(data/curated/loot/forever-raid-phases.json) &mdash; their pulls count toward the tier total above but cannot '
           f'be named or ranked per boss until Forever publishes them.</p></div>')

    trend_rows = ''.join(
        f'<div style="display:flex;align-items:center;gap:14px;padding:6px 4px;border-bottom:1px solid {SOFT}">'
        f'<span class="mono" style="font-size:13px;color:{MUTED};min-width:90px">{date}</span>'
        f'<span style="font-size:13px;color:{TEXT}">pull {n}</span>'
        f'<span class="mono" style="font-size:13px;color:{TEXT}">{d // 60}:{d % 60:02d}</span>'
        f'{pill("Kill", GREEN, "rgba(30,255,0,.10)", "rgba(30,255,0,.25)") if res == "kill" else pill("Wipe", MUTED)}'
        f'</div>' for date, n, d, res in ONYXIA_PULLS
    )
    onyxia_panel = f'''<div style="{PANEL}">
{panel_header("Onyxia &middot; pulls to kill")}
<p style="margin:0;font-size:13px;color:{BODY}">First kill 2026-12-15 (pull 3, 4:28); best kill time 2026-12-20 (pull 2, 4:11) &mdash; improving.</p>
{trend_rows}
<p style="margin:0;font-size:13px;color:{BODY}">Deaths per pull: 3.4 average, 1 on the best kill. Top death causes need the report's own event data
(which spell killed whom) &mdash; death <em>count</em> is stored today, the <em>cause</em> is not
({annotation_inline("NEEDS NEW DATA")}).</p>
<p style="margin:0;font-size:13px;color:{BODY}">Guild best parse, DPS role: <a href="#fight" style="font-weight:600">Pyrewisp, 358.2 DPS, 2026-12-20</a> &middot;
best HPS: <a href="#fight" style="font-weight:600">Lightbrand, 301.9 HPS, 2026-12-15</a></p>
{annotation("EXISTS (fights.duration_ms/kill)")}
</div>'''

    other_zones = ''.join(
        f'<div style="display:flex;align-items:center;gap:12px;padding:7px 4px;border-bottom:1px solid {SOFT}">'
        f'<span style="font-size:14px;font-weight:600;flex:1">{zone}</span>'
        f'<span class="mono" style="font-size:13px;color:{MUTED}">{sum(n[3] for n in NIGHTS if n[0] == zone)} pulls logged</span>'
        f'<span style="font-size:12px;color:{MUTED}">no named encounters yet</span></div>'
        for zone in ['Barrow Deeps', 'Hyjal Summit']
    )
    other_panel = f'<div style="{PANEL}">{panel_header("Barrow Deeps / Hyjal Summit")}{other_zones}{annotation("EXISTS (pulls only)")}</div>' 

    return f'<div style="{PANEL}">{panel_header("Progression")}{bar}{annotation("EXISTS")}</div>' + onyxia_panel + other_panel


# ---------------------------------------------------------------------------
# READINESS tab -- the raid leader's pre-pull board.
# ---------------------------------------------------------------------------
def readiness_header_row() -> str:
    cols = ['Raider', 'Gear gap', 'Enchants', 'Consumables', 'Talent points', 'Item level', 'Last synced', '']
    tips = {
        'Gear gap': 'Upgrades and DPS gain vs this band\'s BiS (web/src/lib/home/upgrades.ts\'s own gain rule)',
        'Enchants': 'Weapon, chest, cloak, boots checked for a non-zero enchant id in the export',
        'Consumables': 'Flask/food/potion presence in bags= -- gear_bags consent only',
        'Talent points': 'Spent points vs 51 at level 60, from the export\'s own talent string',
        'Item level': 'This character vs the roster\'s own median (' + str(MEDIAN_ILVL) + ')',
    }
    cells = []
    for c in cols:
        title = f' title="{tips[c]}"' if c in tips else ''
        cells.append(f'<span{title} style="font-size:11px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;color:{MUTED};{"cursor:help" if c in tips else ""}">{c}</span>')
    return f'<div style="display:grid;grid-template-columns:170px 170px 120px 120px 110px 100px 130px 90px;gap:8px;padding:0 4px 6px 4px">{"".join(cells)}</div>'


def _readiness_cells(name: str) -> dict:
    """The shared, un-rendered facts behind one readiness row -- both the desktop grid and
    the phone stacked card read the exact same values, so the two can never disagree."""
    row = ROSTER_BY_NAME[name]
    ilvl, consent = row[7], row[6]
    gap = gear_gap_for(name)
    has_gear_consent = consent in ('gear', 'gear_bags')
    has_bags_consent = consent == 'gear_bags'
    delta = None if ilvl is None else ilvl - MEDIAN_ILVL
    return {
        'ilvl': ilvl, 'delta': delta, 'gap': gap, 'has_gear_consent': has_gear_consent,
        'has_bags_consent': has_bags_consent,
        'missing_enchants': missing_enchants_for(name) if has_gear_consent else None,
        'consumables_ok': consumables_ok_for(name) if has_bags_consent else None,
        'unspent': unspent_points_for(name),
    }


def _stat_chip(label_text: str, value_html: str) -> str:
    return (f'<span style="display:flex;flex-direction:column;gap:2px;min-width:0">'
            f'<span class="label" style="font-size:10px;color:{MUTED}">{label_text}</span>{value_html}</span>')


def readiness_row(name: str, pinned: bool = False, phone: bool = False) -> str:
    row = ROSTER_BY_NAME[name]
    cls = row[1]
    color = CLASS_COLOR[cls]
    c = _readiness_cells(name)
    gap = c['gap']

    if c['gap']['upgrades'] is not None and c['has_gear_consent']:
        gap_val = f'<span class="mono" style="font-size:12px;color:{TEXT}">{gap["upgrades"]} upgrade{"s" if gap["upgrades"] != 1 else ""} &middot; +{gap["gain"]:.0f} DPS</span>'
    else:
        gap_val = f'<span style="font-size:12px;color:{MUTED}">no gear consent</span>'

    if c['missing_enchants'] is None:
        enchant_val = f'<span style="font-size:12px;color:{MUTED}">consent needed</span>'
    elif c['missing_enchants']:
        enchant_val = pill(', '.join(c['missing_enchants']), RED, 'rgba(255,107,92,.12)', 'rgba(255,107,92,.32)')
    else:
        enchant_val = pill('All enchanted', GREEN, 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)')

    if c['consumables_ok'] is None:
        consumable_val = f'<span style="font-size:12px;color:{MUTED}">consent needed</span>'
    else:
        consumable_val = (pill('Stocked', GREEN, 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)') if c['consumables_ok']
                           else pill('Short', RED, 'rgba(255,107,92,.12)', 'rgba(255,107,92,.32)'))

    pts = c['unspent']
    pts_val = pill(f'{pts} unspent', GOLD) if pts > 0 else f'<span style="font-size:12px;color:{MUTED}">0</span>'

    if c['ilvl'] is None:
        ilvl_val = f'<span style="font-size:12px;color:{MUTED}">&mdash;</span>'
    else:
        dcolor = GREEN if c['delta'] >= 0 else RED
        ilvl_val = f'<span class="mono" style="font-size:12px;color:{dcolor}">{c["ilvl"]} ({"+" if c["delta"] >= 0 else ""}{c["delta"]})</span>'

    bg = 'background:rgba(229,185,85,.06);' if pinned else ''
    name_html = f'<span style="font-size:13px;font-weight:700;color:{color};white-space:nowrap">{name}{" (you)" if pinned else ""}</span>'

    if phone:
        # Round-2 fix: the desktop 8-column grid below does not reflow at 390px -- it simply
        # overflows the viewport (found on this board's own round-2 screenshot review). The
        # phone board instead wraps each check as its own labelled chip, the same
        # flex-wrap idiom the round-1 roster row already uses successfully at this width.
        chips = (f'<div style="display:flex;flex-wrap:wrap;gap:14px 20px">'
                 f'{_stat_chip("Gear gap", gap_val)}{_stat_chip("Enchants", enchant_val)}'
                 f'{_stat_chip("Consumables", consumable_val)}{_stat_chip("Talent pts", pts_val)}'
                 f'{_stat_chip("Item level", ilvl_val)}</div>')
        return (f'<div style="display:flex;flex-direction:column;gap:10px;padding:10px 4px;{bg}border-bottom:1px solid {SOFT}">'
                f'<div style="display:flex;align-items:center;justify-content:space-between;gap:10px">'
                f'<span style="display:flex;align-items:center;gap:8px;min-width:0">{crest_img(cls, 26)}{name_html}</span>'
                f'<span style="font-size:11px;color:{MUTED}">{last_synced_for(name)}</span></div>'
                f'{chips}{row_btn("Nudge", h=32)}</div>')

    return (f'<div style="display:grid;grid-template-columns:170px 170px 120px 120px 110px 100px 130px 90px;'
            f'gap:8px;align-items:center;padding:8px 4px;{bg}border-bottom:1px solid {SOFT}">'
            f'<span style="display:flex;align-items:center;gap:8px;min-width:0">{crest_img(cls, 26)}{name_html}</span>'
            f'{gap_val}{enchant_val}{consumable_val}{pts_val}{ilvl_val}'
            f'<span style="font-size:11px;color:{MUTED}">{last_synced_for(name)}</span>'
            f'{row_btn("Nudge", h=28)}</div>')


def readiness_tab(officer: bool, pin: str | None = None, phone: bool = False) -> str:
    verified = [r for r in ROSTER if r[4]]
    ordered = sorted(verified, key=lambda r: -readiness_score(r[0]))
    if pin:
        ordered = sorted(ordered, key=lambda r: r[0] != pin)
    rows = ''.join(readiness_row(r[0], pinned=(r[0] == pin), phone=phone) for r in ordered)
    note = (f'<p style="margin:0;font-size:12px;color:{MUTED}">Sorted worst-first. "Nudge" copies a message to the clipboard '
            f'&mdash; there is no network path to the addon, so this is copy only, never a push. Enchant slot-has-an-enchant is '
            f'{annotation_inline("EXISTS")} (the export carries a per-slot enchant id); naming which enchant is <em>recommended</em> for a spec/slot is '
            f'{annotation_inline("NEEDS NEW DATA")} (no curated recommended-enchant table exists yet).</p>')
    head = '' if phone else readiness_header_row()
    return f'<div style="{PANEL}">{panel_header("Readiness")}{head}{rows}{note}{annotation("DERIVABLE + NEEDS NEW DATA (see note above)")}</div>' 


# ---------------------------------------------------------------------------
# LOOT tab -- the loot council helper.
# ---------------------------------------------------------------------------
def loot_candidate_row(item_name: str, classes: list, awarded: bool) -> str:
    eligible = [r for r in ROSTER if r[1] in classes and r[4] and gear_gap_for(r[0])['gain'] is not None]
    candidates = sorted(eligible, key=lambda r: -gear_gap_for(r[0])['gain'])[:4]
    # Round-2 fix (coordinator, 2026-10-04): an awarded item has exactly one awardee -- the
    # ranking's own top candidate, since nothing else here decides who got it. Every OTHER
    # candidate row shows no button, just a muted "Awarded to {name}" note; only the
    # awardee's own row carries the green "Awarded" tag. The board previously put an
    # "Awarded" button on all four candidate rows, which read as four separate awards.
    awardee = candidates[0][0] if awarded and candidates else None
    rows = []
    for r in candidates:
        name, cls = r[0], r[1]
        gap = gear_gap_for(name)
        att = attendance_for(name)
        already = gap['upgrades'] == 0
        if awarded and name == awardee:
            action = pill('Awarded', GREEN, 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)')
        elif awarded:
            action = f'<span style="font-size:12px;color:{MUTED}">Awarded to {awardee}</span>'
        else:
            action = row_btn('Award', h=28)
        rows.append(
            f'<div style="display:flex;align-items:center;gap:12px;padding:6px 4px;border-bottom:1px solid {SOFT}">'
            f'{crest_img(cls, 26)}<span style="font-size:13px;color:{CLASS_COLOR[cls]};font-weight:600;min-width:120px">{name}</span>'
            f'<span class="mono" style="font-size:12px;color:{TEXT}">+{gap["gain"]:.0f} DPS</span>'
            f'<span class="mono" style="font-size:12px;color:{MUTED}">{att}/8 nights</span>'
            f'{pill("Already holds equivalent", MUTED) if already else ""}'
            f'<span style="margin-left:auto">{action}</span></div>'
        )
    return (f'<div style="{PANEL}" id="loot-{item_name.lower().replace(chr(39), "").replace(" ", "-")}">'
            f'<div style="display:flex;align-items:center;justify-content:space-between;gap:10px">'
            f'<span class="display" style="font-size:15px;font-weight:700;color:{TEXT}">{item_name}</span>'
            f'{pill("Awarded", GREEN, "rgba(30,255,0,.10)", "rgba(30,255,0,.25)") if awarded else ""}</div>'
            f'<div style="display:flex;flex-direction:column">{"".join(rows)}</div></div>')


def loot_tab(officer: bool) -> str:
    boss_picker = (f'<div style="display:flex;align-items:center;gap:10px">'
                   f'<span class="label" style="color:{MUTED}">Boss</span>'
                   f'<span style="display:inline-flex;align-items:center;height:36px;padding:0 14px;border:1px solid {GOLD};'
                   f'border-radius:4px;color:{GOLD};font-size:13px;font-weight:700">Onyxia &middot; next unkilled: none, farm</span></div>')
    drops = ''.join(loot_candidate_row(item, classes, awarded=(item == ONYXIA_DROPS[0][0]))
                     for item, slot, classes in ONYXIA_DROPS[:6])
    note = (f'<p style="margin:0;font-size:12px;color:{MUTED}">Drops {annotation_inline("EXISTS")} (Onyxia&rsquo;s own 22-item table, '
            f'data/builds/1.60.1.70009/loot.json) &middot; ranked gain {annotation_inline("DERIVABLE")} (the planner&rsquo;s own gain rule, '
            f'web/src/lib/home/upgrades.ts, reused unchanged) &middot; &ldquo;Awarded&rdquo; {annotation_inline("NEEDS NEW DATA")} '
            f'(no table persists a loot decision today &mdash; proposed in the spec&rsquo;s §9 as a new <code>loot_awards</code> row).</p>')
    return f'<div style="{PANEL}">{panel_header("Loot")}{boss_picker}</div>' + drops + f'<div style="{PANEL}">{note}{annotation("EXISTS + DERIVABLE + NEEDS NEW DATA (see note above)")}</div>' 


# ---------------------------------------------------------------------------
# SETTINGS tab -- GuildSettings.svelte's existing fields, restyled into the page.
# ---------------------------------------------------------------------------
def settings_tab() -> str:
    rows = [
        ('Default report visibility', 'Guild only'),
        ('Officer rank threshold', 'Rank 0-1 (GM and first officer rank)'),
        ('Invite link', 'Rotated 2026-12-10 &middot; anyone with this link can join as a member'),
    ]
    body = ''.join(
        f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px;padding:9px 4px;border-bottom:1px solid {SOFT}">'
        f'<span style="font-size:13px;color:{MUTED}">{k}</span><span style="font-size:13px;color:{TEXT}">{v}</span></div>'
        for k, v in rows
    )
    return f'<div style="{PANEL}">{panel_header("Settings")}{officer_tools()}{body}<div style="display:flex;gap:10px">{btn("Save changes", gold=True)}{btn("Remove a member")}</div>{annotation("EXISTS (GuildSettings.svelte, restyled into the page, unchanged fields)")}</div>' 


# ---------------------------------------------------------------------------
TAB_BUILDERS = {
    'Overview': overview_tab,
    'Roster': roster_tab,
    'Raids': raids_tab,
    'Progression': progression_tab,
    'Readiness': readiness_tab,
    'Loot': loot_tab,
    'Settings': settings_tab,
}


def build_board(name: str, active_tab: str, role: str, phone: bool = False, **kw) -> str:
    public = role == 'public'
    officer = role == 'officer'
    viewer_persona = None
    if role == 'officer':
        viewer_persona = {'class': 'warrior', 'faction': FACTION, 'battletag': 'Kraggor'}
    elif role == 'member':
        viewer_persona = {'class': 'hunter', 'faction': FACTION, 'battletag': 'Zulmara'}

    if public:
        role_slot = signin_prompt()
    else:
        viewer_name = VIEWER_OFFICER if officer else VIEWER_MEMBER
        viewer_row = ROSTER_BY_NAME[viewer_name]
        role_slot = standing_line(viewer_name, with_fails=(active_tab == 'Overview' and not officer)) if viewer_row[4] else unverified_note()

    tabs_html = tab_strip(active_tab, role, phone=phone)

    if active_tab == 'Overview':
        body = overview_tab(role, None)
    elif active_tab == 'Roster':
        pin = VIEWER_MEMBER if role == 'member' else None
        body = roster_tab(officer, pin=pin, with_filters=True, wide=not phone)
    elif active_tab == 'Raids':
        body = raids_tab(kw.get('expand_date', NIGHTS[-1][1]))
    elif active_tab == 'Progression':
        body = progression_tab()
    elif active_tab == 'Readiness':
        pin = VIEWER_MEMBER if role == 'member' else None
        body = readiness_tab(officer, pin=pin, phone=phone)
    elif active_tab == 'Loot':
        body = loot_tab(officer)
    elif active_tab == 'Settings':
        body = settings_tab()
    else:
        body = ''

    body_pad = '18px' if phone else '48px'
    nav_fn = (lambda: nav_phone(not public, viewer_persona)) if phone else (lambda: nav('', not public, viewer_persona))

    content = (f'<div style="display:flex;flex-direction:column;gap:22px;padding:24px {body_pad} 32px {body_pad};'
               + ('' if phone else 'width:100%;max-width:1344px;margin:0 auto;') + '">'
               + body + mock_caption() + '</div>')

    band = (f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">'
            f'{nav_fn()}{header(role_slot, tabs_html, phone=phone)}</div>')

    out_html = QUIET_LINK_STYLE + band + content + footer(phone=phone)
    width = 390 if phone else 1440
    return write(f'{name}.html', page(out_html, width)).as_posix()


print(build_board('guild-overview-officer', 'Overview', 'officer'))
print(build_board('guild-roster', 'Roster', 'officer'))
print(build_board('guild-raids', 'Raids', 'officer', expand_date="2026-12-15"))
print(build_board('guild-progression', 'Progression', 'officer'))
print(build_board('guild-readiness', 'Readiness', 'officer'))
print(build_board('guild-loot', 'Loot', 'officer'))
print(build_board('guild-member', 'Overview', 'member'))
print(build_board('guild-phone', 'Readiness', 'member', phone=True))
