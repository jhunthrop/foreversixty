"""Mock boards for the guild page rebuild (design/specs/2026-10-04-guild-page.md).

  python3 design/mocks/gen_guild.py
    -> renders/guild-member.html   (1440, signed in, verified member, ranked)
    -> renders/guild-officer.html  (1440, signed in, officer, Officer tools + full roster)
    -> renders/guild-public.html   (1440, signed out)
    -> renders/guild-phone.html    (390, signed in, verified member)

No real guild data exists yet -- a database seed (api/cmd/seedguild, a separate lane) is
building one from the same picture this file hand-writes: 24 raiders, Molten Core and Onyxia
nights over the last two weeks, item level 55 to 70. Every board says so in its own caption.
Field names below match api/internal/guilds/home.go's RosterRow/HomeReport and
api/internal/rankings/guilds.go's Progression/RosterBest exactly, so nothing here invents a
shape the real API does not already return.
"""
from mocklib import (
    GOLD, MUTED, TEXT, BODY, BORDER, SOFT, CLASS_COLOR, GREEN,
    crest, emblem, nav, nav_phone, footer, page, write,
)

GUILD = {'name': 'Olympus XXVII', 'region': 'us', 'ruleset': 'pvp'}
FACTION = 'horde'

# ---------------------------------------------------------------------------
# Mock roster -- 24 Horde raiders, Molten Core + Onyxia over the last two weeks, ilvl 55-70.
# Shape matches RosterRow exactly: character_key, region, ruleset, name, class, spec, rank,
# verified, logged_recently, consent, item_level?, may_remove.
# ---------------------------------------------------------------------------
ROSTER = [
    # name, class, spec, rank, verified, logged, consent, ilvl
    ('Kraggor', 'warrior', 'Protection', 'leader', True, True, 'gear', 68),
    ('Obnoxious Yell', 'warrior', 'Fury', 'officer', True, True, 'gear', 66),
    ('Sunderfel', 'warrior', 'Arms', 'officer', True, True, 'gear', 61),
    ('Grimtotem', 'warrior', 'Fury', 'member', True, False, 'gear', 58),
    ('Zulmara', 'hunter', 'Marksmanship', 'member', True, True, 'gear', 64),
    ('Windtalon', 'hunter', 'Marksmanship', 'member', True, True, 'gear', 70),
    ('Duskstrider', 'hunter', 'Marksmanship', 'member', True, False, 'gear', 55),
    ('Felsnap', 'hunter', 'Beast Mastery', 'member', True, True, 'gear', 60),
    ('Vexlash', 'rogue', 'Combat', 'officer', True, True, 'gear', 65),
    ('Shadowquill', 'rogue', 'Assassination', 'member', True, True, 'gear', 59),
    ('Nixthrottle', 'rogue', 'Combat', 'member', True, False, 'roster', None),
    ('Hexbramble', 'warlock', 'Affliction', 'member', True, True, 'gear', 63),
    ('Soulgrave', 'warlock', 'Destruction', 'member', True, True, 'gear', 57),
    ('Pyrewisp', 'mage', 'Fire', 'member', True, True, 'gear', 69),
    ('Frostnettle', 'mage', 'Frost', 'member', True, False, 'gear', 56),
    ('Arcanemoor', 'mage', 'Arcane', 'member', True, True, 'gear', 62),
    ('Lightbrand', 'priest', 'Holy', 'member', True, True, 'gear', 67),
    ('Grimvow', 'priest', 'Shadow', 'member', True, True, 'roster', None),
    ('Mendwhisper', 'priest', 'Holy', 'member', True, False, 'gear', 60),
    ('Earthhoof', 'druid', 'Restoration', 'member', True, True, 'gear', 65),
    ('Thornhide', 'druid', 'Feral', 'member', True, True, 'gear', 58),
    ('Rootgall', 'shaman', 'Restoration', 'member', False, True, 'gear', 61),
    ('Stormtusk', 'shaman', 'Elemental', 'member', False, False, 'gear', 55),
    ('Fulmintide', 'shaman', 'Enhancement', 'member', False, True, 'gear', 59),
]

VIEWER_MEMBER = 'Zulmara'   # verified, Marksmanship Hunter, ranked among peers
VIEWER_OFFICER = 'Kraggor'  # leader rank -- sees the full Officer tools strip

# Forever launches 4 November 2026; raids open 9 December 2026 -- no raid night, first-kill
# date or Updated stamp on this board may predate that (spec's own "Data notes"). The earliest
# night below (2026-12-10, Lucifron/Magmadar's first kills) has already rolled off the trailing
# "this week" reports list by the time of the other three.
REPORTS = [
    # title/zone, created_at, fight_count, kill_count
    ('Molten Core, week 3', '2026-12-22', 10, 8),
    ('Onyxia', '2026-12-17', 3, 1),
    ('Molten Core, week 2', '2026-12-15', 9, 6),
]
UPDATED_LABEL = 'Updated Dec 22'

PROGRESSION = [
    # encounter, pulls, first_kill_at or None
    ('Lucifron', 4, '2026-12-10'),
    ('Magmadar', 6, '2026-12-10'),
    ('Gehennas', 5, '2026-12-15'),
    ('Garr', 7, '2026-12-15'),
    ('Baron Geddon', 8, '2026-12-15'),
    ('Shazzrah', 9, '2026-12-22'),
    ('Sulfuron Harbinger', 11, '2026-12-22'),
    ('Golemagg the Incinerator', 6, None),
    ('Majordomo Executus', 2, None),
    ('Ragnaros', 0, None),
    ('Onyxia', 5, '2026-12-17'),
]

ROSTER_BESTS = [
    # name, class, spec, encounter, metric, value, execution
    ('Windtalon', 'hunter', 'Marksmanship', 'Lucifron', 'dps', 312.4, 0.94),
    ('Obnoxious Yell', 'warrior', 'Fury', 'Garr', 'dps', 341.7, 0.91),
    ('Pyrewisp', 'mage', 'Fire', 'Shazzrah', 'dps', 358.2, 0.97),
    ('Earthhoof', 'druid', 'Restoration', 'Sulfuron Harbinger', 'hps', 289.5, 0.88),
    ('Lightbrand', 'priest', 'Holy', 'Onyxia', 'hps', 301.9, 0.92),
]

METRIC_LABEL = {'dps': 'DPS', 'hps': 'HPS'}

# ---------------------------------------------------------------------------
PANEL = (f'background:var(--raised);border:1px solid {BORDER};border-radius:6px;'
         f'padding:18px 20px;display:flex;flex-direction:column;gap:12px')


def crest_img(slug: str, size: int) -> str:
    return f'<img class="crest" src="{crest(slug)}" alt="" style="width:{size}px;height:{size}px;--c:{CLASS_COLOR[slug]}">'


def pill(text: str, color: str = GOLD, bg: str | None = None, border: str | None = None) -> str:
    bg = bg or f'color-mix(in srgb, {color} 14%, transparent)'
    border = border or f'color-mix(in srgb, {color} 35%, transparent)'
    return f'<span class="pill" style="color:{color};background:{bg};border-color:{border};height:22px;padding:0 8px">{text}</span>'


def row_btn(lbl: str, h: int = 32) -> str:
    return (f'<button style="display:inline-flex;align-items:center;justify-content:center;'
            f'height:{h}px;padding:0 12px;border:1px solid #3a3326;border-radius:4px;'
            f'background:none;color:var(--strong);font-family:inherit;font-size:11px;font-weight:700;'
            f'letter-spacing:.06em;text-transform:uppercase;white-space:nowrap;cursor:pointer">{lbl}</button>')


def btn(lbl: str, h: int = 40, gold: bool = False) -> str:
    cls = 'btn' + (' btn-gold' if gold else '')
    return f'<a class="{cls}" href="#{lbl.lower().replace(" ", "-")}" style="height:{h}px">{lbl}</a>'


def h2(text: str) -> str:
    return f'<span class="display" style="font-size:18px;text-transform:uppercase;letter-spacing:.04em;color:{TEXT}">{text}</span>'


def label(text: str) -> str:
    return f'<span class="label">{text}</span>'


def mock_caption() -> str:
    return (f'<p style="margin:0;font-size:12px;color:{MUTED};font-style:italic">'
            f'Mock roster &mdash; 24 characters, no real guild data yet.</p>')


def rank_among_peers(viewer_name: str):
    viewer = next(r for r in ROSTER if r[0] == viewer_name)
    _, cls, spec, *_ = viewer
    peers = sorted(
        [r for r in ROSTER if r[1] == cls and r[2] == spec and r[7] is not None],
        key=lambda r: -r[7],
    )
    rank = next(i for i, r in enumerate(peers, 1) if r[0] == viewer_name)
    return rank, len(peers), viewer


# ---------------------------------------------------------------------------
def header_facts() -> str:
    killed = sum(1 for _, pulls, kill in PROGRESSION if kill is not None)
    pulls = sum(p for _, p, _ in PROGRESSION)
    return (f'<span class="mono" style="color:{TEXT}">PvP US</span>'
            f' &middot; <span class="mono" style="color:{TEXT}">{killed}</span> bosses down'
            f' &middot; <span class="mono" style="color:{TEXT}">{pulls}</span> pulls')


def standing_line(viewer_name: str) -> str:
    rank, total, viewer = rank_among_peers(viewer_name)
    name, cls, spec = viewer[0], viewer[1], viewer[2]
    color = CLASS_COLOR[cls]
    if total == 1:
        text = f'{name} &middot; {spec} {cls.capitalize()} &middot; the only {spec} {cls.capitalize()} in this guild this week'
    else:
        ilvl = viewer[7]
        text = f'{name} &middot; {rank} of {total} {spec} {cls.capitalize()}s this week by item level ({ilvl} ilvl)'
    return (f'<div style="display:flex;align-items:center;gap:12px;flex-wrap:wrap">'
            f'<span style="font-size:14px;color:{color};font-weight:600">{text}</span>'
            f'<a href="#gear-in-the-planner" style="font-size:13px;font-weight:700">Gear in the planner</a></div>')


def unverified_note() -> str:
    return (f'<p style="margin:0;font-size:14px;color:{TEXT}">You are not verified yet. '
            f'An officer can approve you from the roster below.</p>')


def signin_prompt() -> str:
    return (f'<div style="display:flex;flex-direction:column;align-items:flex-start;gap:10px">'
            f'<p style="margin:0;font-size:14px;color:{TEXT}">See where you stand in this guild. '
            f'<a href="#login">Use an email link instead.</a></p>'
            f'{btn("Sign in with Battle.net")}</div>')


def header(role_slot: str, phone: bool = False) -> str:
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    return f'''<div style="display:flex;flex-direction:column;gap:14px;padding:{pad};{maxw}">
<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Guild</span>
<h1 class="display" style="margin:0;font-size:22px;color:{TEXT}">{GUILD["name"]}</h1>
<span style="font-size:12px;color:{MUTED}">{UPDATED_LABEL}</span>
<p style="margin:0;font-size:14px;color:{BODY}">{header_facts()}</p>
{role_slot}
</div>'''


# ---------------------------------------------------------------------------
def officer_tools() -> str:
    waiting = sum(1 for r in ROSTER if not r[4])
    return f'''<div style="{PANEL};flex-direction:row;align-items:center;justify-content:space-between;gap:20px;flex-wrap:wrap">
<div style="display:flex;align-items:center;gap:10px">{pill("Claimed")}<span style="font-size:13px;color:{MUTED}">by Kraggor</span></div>
<a href="#guild-roster-unverified" style="font-size:13px;font-weight:700;color:{GOLD}">{waiting} waiting for approval</a>
<div style="display:flex;align-items:center;gap:12px;flex-wrap:wrap">
<span class="label" style="color:{MUTED}">Invite link</span>
<span style="font-size:12px;color:{MUTED}">Anyone with this link can join as a member.</span>
{btn("Rotate invite link")}
</div></div>'''


# ---------------------------------------------------------------------------
def sort_header(active: str) -> str:
    cols = [('Rank', 'rank'), ('Item level', 'ilvl'), ('Last seen', 'seen')]
    cells = []
    for lbl, key in cols:
        arrow = ' &#9662;' if key == active else ''
        color = GOLD if key == active else MUTED
        cells.append(f'<button style="background:none;border:none;height:36px;padding:0;font-size:11px;font-weight:700;'
                      f'letter-spacing:.1em;text-transform:uppercase;color:{color};cursor:pointer">{lbl}{arrow}</button>')
    return f'<div style="display:flex;gap:28px;padding:0 4px">{"".join(cells)}</div>'


def roster_row(row: tuple, officer: bool) -> str:
    name, cls, spec, rank, verified, logged, consent, ilvl = row
    color = CLASS_COLOR[cls]
    rank_pill = pill(rank.capitalize()) if rank in ('officer', 'leader') else ''
    unverified_pill = pill('Unverified', MUTED) if not verified else ''
    logged_pill = pill('Logged in the last day', GREEN, 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)') if logged else ''
    ilvl_cell = f'<span class="mono" style="font-size:13px;color:{TEXT}">ilvl {ilvl}</span>' if ilvl is not None else ''
    handoff = (f'<a href="#open-in-sim" style="font-size:13px;font-weight:700">Open in simulator</a>'
               f'<a href="#open-in-planner" style="font-size:13px;font-weight:700">Open in planner</a>') if consent != 'roster' else ''
    actions = ''
    if officer:
        if not verified:
            actions += row_btn('Approve')
        elif rank != 'leader':
            actions += row_btn('Remove')
    return (f'<div style="display:flex;align-items:center;gap:12px;flex-wrap:wrap;padding:9px 4px;'
            f'border-bottom:1px solid {SOFT}">'
            f'{crest_img(cls, 36)}'
            f'<div style="display:flex;flex-direction:column;min-width:140px">'
            f'<span class="display" style="font-size:14px;font-weight:700;color:{color}">{name}</span>'
            f'<span style="font-size:12px;color:{MUTED}">{spec}</span></div>'
            f'{rank_pill}{unverified_pill}{logged_pill}{ilvl_cell}'
            f'<div style="display:flex;gap:10px;margin-left:auto">{handoff}{actions}</div></div>')


def roster_panel(officer: bool) -> str:
    unverified = [r for r in ROSTER if not r[4]]
    verified = sorted([r for r in ROSTER if r[4]], key=lambda r: -(r[7] or 0))
    waiting = len(unverified)
    approve_all = (
        f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px">'
        f'<span style="font-size:13px;color:{MUTED}">{waiting} waiting for approval</span>'
        + row_btn('Approve all', h=36) + '</div>'
    ) if officer and waiting else ''
    unverified_rows = ''.join(roster_row(r, officer) for r in unverified)
    verified_rows = ''.join(roster_row(r, officer) for r in verified)
    return f'''<div style="{PANEL}">
{h2("Roster")}
{approve_all}
<div id="guild-roster-unverified" style="display:flex;flex-direction:column">{unverified_rows}</div>
<div style="margin-top:6px">{sort_header("ilvl")}</div>
<div style="display:flex;flex-direction:column">{verified_rows}</div>
</div>'''


# ---------------------------------------------------------------------------
def raid_nights_panel(heading: str) -> str:
    rows = []
    for title, date, fights, kills in REPORTS:
        wipes = fights - kills
        rows.append(f'<div style="display:flex;align-items:center;gap:12px;padding:9px 4px;'
                     f'border-bottom:1px solid {SOFT}">'
                     f'<a href="#report" style="font-size:14px;font-weight:600">{title}</a>'
                     f'<span class="mono" style="font-size:13px;color:{MUTED}">{date}</span>'
                     f'<span class="mono" style="font-size:13px;color:{MUTED};margin-left:auto">{kills} kills &middot; {wipes} wipes</span>'
                     f'</div>')
    return f'<div style="{PANEL}">{h2(heading)}{"".join(rows)}</div>'


def progression_panel() -> str:
    rows = []
    for encounter, pulls, first_kill in PROGRESSION:
        killed = first_kill if first_kill else 'not killed'
        color = TEXT if first_kill else MUTED
        rows.append(f'<div style="display:flex;align-items:center;gap:12px;padding:9px 4px;'
                     f'border-bottom:1px solid {SOFT}">'
                     f'<a href="#rankings" style="font-size:14px;font-weight:600;flex:1">{encounter}</a>'
                     f'<span class="mono" style="font-size:13px;color:{MUTED};width:90px;text-align:right">{pulls} pulls</span>'
                     f'<span class="mono" style="font-size:13px;color:{color};width:110px;text-align:right">{killed}</span>'
                     f'</div>')
    return f'<div style="{PANEL}">{h2("Progression")}{"".join(rows)}</div>'


def roster_bests_panel() -> str:
    rows = []
    for name, cls, spec, enc, metric, value, exe in ROSTER_BESTS:
        color = CLASS_COLOR[cls]
        rows.append(f'<div style="display:flex;align-items:center;gap:12px;padding:9px 4px;'
                     f'border-bottom:1px solid {SOFT}">'
                     f'<span class="display" style="font-size:14px;font-weight:700;color:{color};min-width:140px">{name}</span>'
                     f'<span style="font-size:13px;color:{MUTED};flex:1">{enc} &middot; {spec}</span>'
                     f'<span class="mono" style="font-size:13px;color:{TEXT}">{value:.1f} {METRIC_LABEL[metric]}</span>'
                     f'<span class="mono" style="font-size:12px;color:{MUTED};width:50px;text-align:right">{exe*100:.0f}%</span>'
                     f'</div>')
    return f'<div style="{PANEL}">{h2("This guild\'s best parses")}{"".join(rows)}</div>'


# ---------------------------------------------------------------------------
def build(kind: str) -> str:
    phone = kind == 'phone'
    public = kind == 'public'
    officer = kind == 'officer'
    viewer_persona = {'class': 'hunter', 'faction': FACTION, 'battletag': 'Zulmara'} if not public else None
    if officer:
        viewer_persona = {'class': 'warrior', 'faction': FACTION, 'battletag': 'Kraggor'}

    if public:
        role_slot = signin_prompt()
    else:
        viewer_name = VIEWER_OFFICER if officer else VIEWER_MEMBER
        viewer_row = next(r for r in ROSTER if r[0] == viewer_name)
        role_slot = standing_line(viewer_name) if viewer_row[4] else unverified_note()

    sections = []
    if officer:
        sections.append(officer_tools())
    if not public:
        sections.append(roster_panel(officer))
    heading = 'This week\'s raid nights' if not public else 'Recent raid nights'
    sections.append(raid_nights_panel(heading))
    sections.append(progression_panel())
    sections.append(roster_bests_panel())

    body_pad = '18px' if phone else '48px'
    nav_fn = (lambda: nav_phone(not public, viewer_persona)) if phone else (lambda: nav('', not public, viewer_persona))

    content = (f'<div style="display:flex;flex-direction:column;gap:22px;padding:24px {body_pad} 32px {body_pad};'
               + ('' if phone else 'width:100%;max-width:1344px;margin:0 auto;') + '">'
               + ''.join(sections) + mock_caption() + '</div>')

    band = (f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">'
            f'{nav_fn()}{header(role_slot, phone=phone)}</div>')

    out_html = band + content + footer(phone=phone)
    width = 390 if phone else 1440
    name = f'guild-{kind}.html'
    return write(name, page(out_html, width)).as_posix()


for kind in ['member', 'officer', 'public', 'phone']:
    print(build(kind))
