"""Mock boards for the /rankings redesign (design/specs/2026-10-11-rankings.md), round two.

  python3 design/mocks/gen_rankings.py   -> renders/rankings{,-signed-out,-2000,-1024,-phone,-states}.png

The header is the shared component (gen_character_header.py). The encounter (Onyxia), class, spec,
phase and ruleset names are the site's own; the CHARACTERS, guilds, values, sizes, lengths and dates on
the board are EXAMPLES, labelled in the spec and never on a board. HTML goes to a scratch dir; only PNGs
land in renders/.
"""
import sys

import gen_character_header as H
from mocklib import GOLD, MUTED, TEXT, BODY, BORDER, SOFT, RAISED, CLASS_COLOR, PARSE

ENCOUNTER = 'Onyxia'
TOTAL, OWN_RANK = 212, 14
UPDATED = 'Updated Dec 13'
# (rank, name, class, guild, region, size, value, executed %, days ago, split, length). EXAMPLES.
NAMES = ['Thorgrim Ashveil', 'Valkyra', 'Brannoch', 'Kaelith Rowe', 'Rhoswen', 'Dunmar', 'Tessiel Vane', 'Ironvow', 'Gaelwyn',
         'Hrafnild', 'Morvain Holt', 'Sellar', 'Orsolya', 'Obnoxious Yell', 'Bladewick', 'Draveth']
GUILDS = ['The Last Watch', 'Stonehearth Vanguard', '', 'Olympus XXVII', 'The Last Watch', '', 'Ember Court', 'Olympus XXVII', '',
          'Stonehearth Vanguard', '', 'Ember Court', '', 'Olympus XXVII', '', 'The Last Watch']
VALUES = [842, 821, 805, 790, 771, 755, 741, 728, 690, 672, 655, 641, 622, 612, 605, 598]
SIZES = [40] * 16     # Onyxia is a 40-man board; other sizes are shown on the states sheet's dungeon example
LENGTHS = ['2:05', '2:11', '1:52', '2:18', '2:24', '3:40', '2:31', '2:58', '2:40', '2:47', '3:12', '2:52', '4:05', '2:41', '2:49', '3:20']
EXEC = [97, 95, 96, 93, 94, 91, 92, 90, 89, 88, 90, 86, 87, 94, 85, 84]
DAYS = [1, 2, 4, 3, 1, 5, 2, 6, 9, 4, 11, 2, 8, 3, 12, 5]
REGION = ['US', 'EU', 'US', 'US', 'US', 'EU', 'US', 'US', 'US', 'EU', 'US', 'US', 'EU', 'US', 'US', 'US']
SPLIT = ['31/20/0', '31/20/0', '30/21/0', '31/20/0', '31/20/0', '31/20/0', '30/21/0', '31/20/0', '31/20/0', '31/20/0', '30/21/0', '31/20/0', '31/20/0', '31/20/0', '30/21/0', '31/20/0']
BOARD = [(i + 1, NAMES[i], 'warrior', GUILDS[i], REGION[i], SIZES[i], VALUES[i], EXEC[i], DAYS[i], SPLIT[i], LENGTHS[i]) for i in range(16)]
MIXED = [(1, 'Reglitch', 'priest', '', 'US', 40, 905, 96, 1, '0/31/20', '2:09'), (2, 'Mishvamp', 'shaman', 'Olympus XXVII', 'US', 40, 872, 95, 2, '31/0/20', '2:14'),
         (3, 'Thorgrim Ashveil', 'warrior', 'The Last Watch', 'US', 40, 842, 97, 1, '31/20/0', '2:05'), (4, 'Quillfeather', 'rogue', '', 'US', 40, 830, 90, 3, '31/20/0', '1:58'),
         (5, 'Hanabanana', 'priest', 'Ember Court', 'EU', 40, 801, 91, 4, '0/31/20', '2:20'), (6, 'Valkyra', 'warrior', 'Stonehearth Vanguard', 'EU', 40, 821, 95, 2, '31/20/0', '2:11')]
MIXED = sorted(MIXED, key=lambda r: -r[6])
MIXED = [(i + 1, *r[1:]) for i, r in enumerate(MIXED)]
OWN = BOARD[OWN_RANK - 1]
DUNGEON = [(1, 'Brannoch', 'warrior', 'The Last Watch', 'US', 5, 1912, 97, 1, '31/20/0', '2:52'), (2, 'Valkyra', 'warrior', '', 'EU', 5, 1874, 95, 2, '31/20/0', '2:58'),
           (3, 'Dunmar', 'warrior', 'Ember Court', 'US', 4, 1790, 93, 4, '30/21/0', '3:20')]

STYLE = f'''<style>
.rw{{display:grid;align-items:center;gap:12px;min-height:44px;padding:0 14px 0 12px;border-bottom:1px solid {SOFT};border-left:2px solid transparent;font-size:14px}}
.rw.you{{background:color-mix(in srgb,{GOLD} 6%,transparent);border-left-color:{GOLD}}}
.rw:hover{{background:#101624}}
.rw a{{color:{TEXT}}}.rw a.lk{{color:{GOLD};font-weight:600}}
.rw a.bl{{color:{BODY};font-weight:400}}.rw a.bl:hover{{text-decoration:underline;text-underline-offset:4px}}
.col{{font-size:11px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;color:{MUTED}}}
.sel2{{display:flex;align-items:center;justify-content:space-between;gap:8px;height:44px;border:1px solid #3a3326;border-radius:4px;background:{RAISED};padding:0 12px;font-size:13px;color:{TEXT}}}
.pillb{{display:inline-flex;align-items:center;height:44px;padding:0 14px;border:1px solid {BORDER};border-radius:4px;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:#b9b3a4}}
.pillb.on{{border-color:{GOLD};color:{GOLD};background:color-mix(in srgb,{GOLD} 8%,transparent)}}
</style>'''
GRID_FULL = '44px minmax(150px,1.4fr) minmax(120px,1fr) 56px 96px 76px 76px 100px 84px 60px'
GRID_MID = '44px minmax(150px,1.5fr) 96px 76px 100px 84px 60px'
CHEV = '<svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"><path d="M6 9l6 6 6-6"/></svg>'


def band_of(rank: int) -> str:
    pct = 100 - (rank - 1) / TOTAL * 100
    return PARSE['pink'] if pct >= 99 else PARSE['orange'] if pct >= 95 else PARSE['purple'] if pct >= 75 else PARSE['blue'] if pct >= 50 else PARSE['green'] if pct >= 25 else '#8c8c8c'


def name_cell(r: tuple, multi: bool, crest_px: int = 28) -> str:
    img = H.crest_img(r[2], crest_px) if multi else ''
    return (f'<span style="display:flex;align-items:center;gap:10px;min-width:0">{img}<a class="lk" href="#c" '
            f'style="color:{CLASS_COLOR[r[2]]};white-space:nowrap;overflow:hidden;text-overflow:ellipsis">{r[1]}</a></span>')


def row(r: tuple, width: int, you: bool = False, multi: bool = False) -> str:
    rank, name, cls, guild, region, size, val, ex, days, split, length = r
    date = H.rel_date(days)
    rk = f'<span class="mono" style="font-size:12px;color:{band_of(rank)}">{rank}</span>'
    g = f'<a href="#g" style="color:#b9b3a4">{guild}</a>' if guild else ''
    cls_ = 'rw you' if you else 'rw'
    mono = lambda t, c=MUTED, a='right': f'<span class="mono" style="text-align:{a};font-size:13px;color:{c}">{t}</span>'
    build = f'<a class="bl mono" href="#b" style="text-align:right;font-size:13px">{split}</a>'
    report = '<a class="lk" href="#r" style="text-align:right;font-size:13px">Report</a>'
    if width < 600:
        return (f'<div class="{cls_}" style="grid-template-columns:32px minmax(0,1fr) auto 52px;gap:8px;min-height:60px;padding:6px 10px 6px 8px">{rk}'
                f'<span style="display:flex;flex-direction:column;gap:1px;min-width:0">{name_cell(r, multi, 22)}'
                f'<span style="font-size:12px;color:{MUTED};white-space:nowrap;overflow:hidden;text-overflow:ellipsis">{size}-man · {length} · {ex}% · <a class="bl" href="#b" style="color:{BODY}">{split}</a></span></span>'
                f'<span class="mono" style="text-align:right">{val:,}</span><a class="lk" href="#r" style="display:inline-flex;align-items:center;justify-content:flex-end;min-height:44px;font-size:13px">Report</a></div>')
    if width < 1100:
        return (f'<div class="{cls_}" style="grid-template-columns:{GRID_MID}">{rk}<span style="display:flex;flex-direction:column;min-width:0;padding:4px 0">{name_cell(r, multi)}'
                f'<span style="font-size:12px;color:{MUTED}">{guild + " · " if guild else ""}{region}</span></span>'
                f'<span class="mono" style="text-align:right">{val:,}</span>{mono(f"{ex}%", TEXT)}{mono(date)}{build}{report}</div>')
    return (f'<div class="{cls_}" style="grid-template-columns:{GRID_FULL}">{rk}{name_cell(r, multi)}'
            f'<span style="font-size:13px;color:{MUTED};white-space:nowrap;overflow:hidden;text-overflow:ellipsis">{g}{f" · {region}" if guild else ""}</span>'
            f'{mono(size)}<span class="mono" style="text-align:right">{val:,}</span>{mono(f"{ex}%", TEXT)}{mono(length)}{mono(date)}{build}{report}</div>')


def cols(width: int, metric: str = 'Damage') -> str:
    if width < 600:
        return ''
    if width < 1100:
        labels, al, grid = ['#', 'Character', metric, 'Executed', 'Date', 'Build', 'Report'], ['', '', 'right', 'right', 'right', 'right', 'right'], GRID_MID
    else:
        labels = ['#', 'Character', 'Guild', 'Size', metric, 'Executed', 'Length', 'Date', 'Build', 'Report']
        al, grid = ['', '', '', 'right', 'right', 'right', 'right', 'right', 'right', 'right'], GRID_FULL
    cells = ''.join(f'<span class="col" style="text-align:{a or "left"}">{l}</span>' for l, a in zip(labels, al))
    return f'<div class="rw" style="grid-template-columns:{grid};min-height:36px">{cells}</div>'


def select(label: str, value: str, grow: float = 1) -> str:
    return (f'<span style="display:flex;flex-direction:column;gap:4px;flex:{grow} 1 0;min-width:0"><span class="label">{label}</span>'
            f'<span class="sel2"><span style="white-space:nowrap;overflow:hidden;text-overflow:ellipsis">{value}</span>{CHEV}</span></span>')


def filters(width: int, signed_in: bool, open_: bool = False) -> str:
    sels = [select('Encounter', ENCOUNTER, 1.3), select('Ruleset', 'Normal'), select('Region', 'All regions'), select('Phase', 'Phase 1'), select('Faction', 'Both'), select('Metric', 'Damage')]
    if not signed_in:
        sels = [select('Class', 'Every class'), select('Spec', 'Every spec')] + sels
    today = f'<label style="display:flex;align-items:center;gap:8px;min-height:44px;font-size:13px;color:{BODY}"><input type="checkbox" style="width:16px;height:16px">Today only</label>'
    if width < 600 and not open_:   # one summary line that opens the panel
        return (f'<div class="panel" style="display:flex;align-items:center;gap:10px;padding:0 6px 0 14px;min-height:52px"><span style="flex:1;min-width:0;font-size:13px;color:{BODY};padding:8px 0;line-height:1.35">'
                f'{ENCOUNTER} · Normal · All regions · Phase 1</span><span class="btn" style="height:44px;gap:6px">Filters {CHEV}</span></div>')
    if width < 1100 or open_:
        grid = 'repeat(2,1fr)' if width < 600 else 'repeat(3,1fr)'
        first = f'<div style="grid-column:1/-1">{sels[0]}</div>' if signed_in else ''
        rest = sels[1:] if signed_in else sels
        return (f'<div class="panel" style="padding:14px {10 if width < 480 else 16}px;display:grid;grid-template-columns:{grid};gap:10px;align-items:end">{first}{"".join(rest)}<div>{today}</div></div>')
    return f'<div class="panel" style="padding:14px 18px;display:flex;align-items:flex-end;gap:12px">{"".join(sels)}{today}</div>'


def legend() -> str:
    items = [('99+', PARSE['pink']), ('95+', PARSE['orange']), ('75+', PARSE['purple']), ('50+', PARSE['blue']), ('25+', PARSE['green']), ('under 25', '#8c8c8c')]
    return (f'<span style="font-size:12px;color:{MUTED};display:flex;gap:6px 12px;flex-wrap:wrap">Rank colour is the percentile:'
            + ''.join(f'<span class="mono" style="color:{c}">{t}</span>' for t, c in items) + '</span>')


def board_heading(signed_in: bool, scope: str) -> str:
    other = (f'<a href="#other" style="font-size:13px;color:#b9b3a4;text-decoration:underline;text-underline-offset:4px;display:inline-flex;align-items:center;min-height:44px">Another class or spec</a>' if signed_in else '')
    tabs = '<span style="display:flex;gap:6px"><span class="pillb on">Characters</span><span class="pillb">Guilds</span></span>'
    return (f'<div style="display:flex;align-items:flex-end;justify-content:space-between;gap:16px;flex-wrap:wrap"><div style="display:flex;flex-direction:column;gap:2px">'
            f'<div style="display:flex;align-items:baseline;gap:16px;flex-wrap:wrap"><h2 class="display" style="font-size:18px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:{TEXT}">{scope}</h2>{other}</div>{legend()}</div>{tabs}</div>')


def pager() -> str:
    b = lambda t, dis: f'<span class="btn" style="height:44px;opacity:{.5 if dis else 1}">{t}</span>'
    return f'<div style="display:flex;align-items:center;gap:12px">{b("Previous", True)}<span class="mono" style="font-size:13px;color:{MUTED}">Page 1 of 3</span>{b("Next", False)}</div>'


def board_panel(width: int, signed_in: bool, pin: tuple | None = None) -> str:
    rows_src = BOARD if signed_in else MIXED
    pinned = ''
    if pin is not None:   # only when the own row is off the visible page
        pinned = (f'<div style="padding:10px 14px 0"><span class="label">Your best kill</span></div>{row(pin, width, True)}<div style="height:1px;background:#4a4030"></div>')
    shown = ''.join(row(r, width, you=(signed_in and r[0] == OWN_RANK and pin is None), multi=len({x[2] for x in rows_src}) > 1) for r in rows_src)
    return f'<div class="panel" style="overflow:hidden">{cols(width)}{pinned}{shown}</div>'


def board(width: int, signed_in: bool) -> str:
    pointer = f'<span style="font-size:14px;color:{MUTED}">Comparing specs, not characters? <a class="tl" href="#tiers">Tier list →</a></span>'
    hero = H.hero_block(width, 'Characters', 'Rankings',
                        'Where your character stands among characters like it.' if signed_in else 'How characters rank on each encounter, kill by kill.', pointer, UPDATED)
    lead = H.character_header(width, 'rankings', 'ready') if signed_in else H.signed_out_band()
    scope = 'Fury Warriors' if signed_in else 'Every class and spec'
    main = f'{lead}{board_heading(signed_in, scope)}{filters(width, signed_in)}{board_panel(width, signed_in)}{pager()}'
    return H.page_doc(width, 'Rankings', signed_in, STYLE + hero, main)


def states() -> str:
    from mocklib import page
    cap = lambda t: f'<span class="label" style="color:{GOLD}">{t}</span>'
    cell = lambda t, h, w=1248: f'<div style="display:flex;flex-direction:column;gap:8px;width:{w}px">{cap(t)}{h}</div>'
    panel = lambda inner: f'<div class="panel" style="padding:18px;font-size:14px;color:{BODY};display:flex;flex-direction:column;gap:10px">{inner}</div>'
    link = lambda t: f'<a class="tl" href="#x" style="display:inline-flex;align-items:center;min-height:44px">{t}</a>'
    sk = ''.join(f'<div class="rw" style="grid-template-columns:44px 1fr 90px">{H.skel(20, 12)}{H.skel("40%", 14)}{H.skel(70, 14)}</div>' for _ in range(5))
    own_far = (204, 'Thornwick Pell', 'warrior', 'Ember Court', 'US', 40, 410, 71, 20, '31/20/0', '3:58')
    items = [
        cell('Own row on the visible page: no pin, the gold edge marks it in place (the main render)', board_panel(1440, True).replace(cols(1440), cols(1440), 1)),
        cell('Own row off the visible page (a different example character, rank 204): pinned at its true rank, then the board from 1 (older kills show a date, not "N days ago")', f'<div class="panel" style="overflow:hidden">{cols(1440)}<div style="padding:10px 14px 0"><span class="label">Your best kill</span></div>{row(own_far, 1440, True)}<div style="height:1px;background:#4a4030"></div>{row(BOARD[0], 1440)}{row(BOARD[1], 1440)}</div>'),
        cell('Other board, other sizes: a dungeon encounter (Nalthor the Rimebinder) is a 5-man board; a short group shows 4. Sizes and lengths come from the row', f'<div class="panel" style="overflow:hidden">{cols(1440)}{"".join(row(r, 1440) for r in DUNGEON)}</div>'),
        cell('Row hover (ground only, nothing moves); Build link underlines on hover, Report stays gold', f'<div class="panel" style="overflow:hidden"><div style="background:#101624">{row(BOARD[2], 1440)}</div></div>'),
        cell('Removed or at-risk kill: one plain line under the name, no pill', f'<div class="panel" style="overflow:hidden">{row(BOARD[4], 1440)}<div style="padding:0 14px 8px 58px;font-size:12px;color:{MUTED}">Removed by a moderator</div></div>'),
        cell('Board loading: eight rows at the reserved height', f'<div class="panel" style="overflow:hidden;min-height:260px">{sk}</div>'),
        cell('Board empty for this filter', panel(f'No ranked fights yet for this filter. Rankings fill in as reports are uploaded.<span>{link("Upload a log →")}</span>')),
        cell('Board failed', panel(f'Rankings did not load.{link("Try again")}')),
        cell('Header, empty: no ranked kill on this encounter. The board still shows; no pinned row', H.character_header(1440, 'rankings', 'empty')),
        cell('Header, another class or spec chosen: the answer stays about the character; the board is someone else\'s',
             f'{H.character_header(1440, "rankings", "ready")}<div style="font-size:13px;color:{MUTED};margin-top:8px">Showing Frost Mages. Your rank is on the Fury Warrior board. <a class="tl" href="#back">Back to Fury Warriors</a></div>'),
        cell('Phone filters, closed (the main render) and open', f'<div style="display:flex;gap:18px;align-items:flex-start"><div style="width:390px">{filters(390, True)}</div><div style="width:390px">{filters(390, True, True)}</div></div>'),
        cell('Encounter picker (nothing to default to): one list, the filters above stay', panel(f'<span class="label">Choose an encounter</span>{"".join(f"<div style=border-bottom:1px solid {SOFT};padding:8px+0>{e}</div>" for e in ("Onyxia", "Nalthor the Rimebinder", "Warden Kelthas"))}')),
    ]
    body = f'<div style="{H.inner(1440, "36px 0")};display:flex;flex-direction:column;gap:20px"><h1 class="display" style="font-size:22px;color:{TEXT}">Rankings: states</h1>{"".join(items)}</div>'
    return page(H.CSS + STYLE + body, 1440)


BOARDS = {
    'rankings': (lambda: board(1440, True), 1440),
    'rankings-signed-out': (lambda: board(1440, False), 1440),
    'rankings-2000': (lambda: board(2000, True), 2000),
    'rankings-1024': (lambda: board(1024, True), 1024),
    'rankings-phone': (lambda: board(390, True), 390),
    'rankings-states': (states, 1440),
}

if __name__ == '__main__':
    for name in sys.argv[1:] or list(BOARDS):
        build, width = BOARDS[name]
        print(H.shoot(build(), name, width))
