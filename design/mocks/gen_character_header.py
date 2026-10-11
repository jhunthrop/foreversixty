"""The shared character header (design/specs/2026-10-11-character-header.md).

  python3 design/mocks/gen_character_header.py   -> renders/character-header-states.png

One component, three slots. FIXED: the panel, the identity (crest, label, name, descriptor, sync line)
and every state. SLOTS: the page's answer and the page's one action. gen_rankings.py and gen_logs.py
import `character_header` and the PAGE_ANSWERS table from here, so the header on every board is the same
code, not a copy. The CHARACTER (Obnoxious Yell) and every figure in PAGE_ANSWERS are EXAMPLES, labelled
in the spec and never on a board. Class names, spec names, the encounter, the faction logos and the
crests are the site's own. HTML goes to a scratch dir; only PNGs land in renders/.
"""
import subprocess
import sys
import tempfile
from pathlib import Path

import gen_nav as N
from mocklib import (
    WEB, RENDERS, GOLD, MUTED, TEXT, BODY, BORDER, SOFT, RAISED, CLASS_COLOR, PARSE,
    crest, data_uri, page,
)

EMBER = N.EMBER
PAGE_MAX = 1344
BAND_MIN = 136   # the tallest two-line state: name, descriptor and a two-line stale sync
FLAT = {f: data_uri(WEB / f'public/icons/hd/faction/{f}-logo-512.webp') for f in ('alliance', 'horde')}
FACTION_WORD = {'alliance': 'Alliance', 'horde': 'Horde'}

# EXAMPLE character: invented name on a real class and spec pair, as in gen_nav.py and gen_bis_landing.py.
EX = {'name': 'Obnoxious Yell', 'cls': 'warrior', 'race': 'Gnome', 'specname': 'Fury', 'level': 60, 'faction': 'alliance',
      'region': 'us', 'ruleset': 'normal', 'source': 'Addon', 'age': '12 minutes ago', 'characters': 2}
NORACE = {**EX, 'race': None}
LONG = {**EX, 'name': 'Bartholomew Thunderfist', 'age': '12 minutes ago'}
NAV_EX = N.example('Obnoxious Yell', 'warrior-fury', 60, 'Living Flame', 'alliance', 'Addon', '12 minutes ago', current=True)

CLAMP2 = 'display:-webkit-box;-webkit-line-clamp:2;-webkit-box-orient:vertical;overflow:hidden;'
STALE_TAIL = 'Log out in the game to sync'

# What each page fills. Figures are EXAMPLES (spec section 4). `needs_gear` pages show the no-sync sentence
# when the character has no gear on file; the others are unaffected by it.
PAGE_ANSWERS = {
    'home': {
        'label': 'Upgrades in your band', 'figure': '9', 'unit': 'upgrades', 'figure2': '+11.2', 'unit2': 'DPS',
        'note': 'Best in slot for 60, together', 'needs_gear': True, 'needs_level': 20,
        'under20': 'Best in slot starts at 20.',
        'nosync': ('Add the addon to see which slots you can upgrade.', 'Get the addon'),
        'empty': 'Every slot matches the 60 list.', 'error': 'Your upgrades did not load.',
        'action': 'See your upgrades →', 'href': '#upgrades'},
    'bis': {
        'label': 'Best-in-slot slots you wear', 'figure': '6', 'unit': 'of 17 slots', 'bar': 6 / 17,
        'note': '', 'needs_gear': True, 'needs_level': 20,
        'under20': 'Best in slot starts at 20.',
        'nosync': ('Add the addon to see how many of these slots you already wear.', 'Get the addon'),
        'stale_label': 'Slots you wore at the last sync', 'error': 'Your gear count did not load.',
        'action': 'Fury gear for 60 →', 'href': '#bis/warrior/fury'},
    'rankings': {
        'label': 'Your rank', 'figure': '14th', 'unit': 'of 212 Fury Warrior kills', 'figure_color': PARSE['purple'],
        'note': 'On Onyxia, Normal ruleset',
        'empty': 'Not ranked yet on Onyxia. Upload a log of a kill to get on the board.',
        'error': 'Your rank did not load.',
        'action': 'See your parses →', 'href': '#character'},
    'logs': {
        'label': 'Latest kill', 'figure': '612', 'unit': 'DPS',
        'note': 'Onyxia kill · 3 days ago', 'note_link': 'Open the fight', 'note2': 'Last report 1 day ago',
        'empty': 'No logs yet. Upload one, or log live with the companion.', 'error': 'Your latest kill did not load.',
        'action': 'Upload a log', 'href': '#upload'},
}

CSS = f'''<style>
.crest{{--c:{GOLD}}}
.cta{{height:48px;background:linear-gradient(180deg,#f0cc6c,#c99a3a);border:1px solid #a8762a;color:#1a1408;font-size:13px;box-shadow:0 0 18px rgba(229,185,85,.18);min-width:250px}}
.cta:hover,.cta.is-hover{{background:linear-gradient(180deg,#f5d27a,#d4a444);color:#1a1408}}
.cta:focus-visible,.cta.is-focus{{outline:2px solid {GOLD};outline-offset:2px}}
.cta:active,.cta.is-active{{background:linear-gradient(180deg,#d9b55c,#b88a32)}}
.cta2{{height:44px;border:1px solid {GOLD};color:{GOLD};font-size:12px;min-width:220px;background:transparent}}
.cta2:hover,.cta2.is-hover{{border-color:#f5d27a;color:#f5d27a;background:color-mix(in srgb,{GOLD} 8%,transparent)}}
.cta2:focus-visible,.cta2.is-focus{{outline:2px solid {GOLD};outline-offset:2px}}
.cta2:active,.cta2.is-active{{background:color-mix(in srgb,{GOLD} 4%,transparent)}}
.idl{{display:flex;align-items:center;gap:16px;color:inherit;min-width:0}}
.idl:hover .idn,.idl.is-hover .idn{{text-decoration:underline;text-underline-offset:4px}}
.idl:focus-visible,.idl.is-focus{{outline:2px solid {GOLD};outline-offset:4px;border-radius:6px}}
.tl{{color:{GOLD};font-weight:700}}.tl:hover{{text-decoration:underline;text-underline-offset:4px;color:#f5d27a}}
.track{{height:6px;border-radius:3px;background:{SOFT};overflow:hidden}}.fillg{{display:block;height:100%;background:{GOLD};border-radius:3px}}
.sk{{background:linear-gradient(90deg,{SOFT},#232b3d,{SOFT});border-radius:3px;display:block}}
</style>'''


# ---------------------------------------------------------------------------- bits
def crest_img(slug: str, px: int) -> str:
    return f'<img class="crest" src="{crest(slug)}" alt="" width="{px}" height="{px}" style="width:{px}px;height:{px}px;--c:{CLASS_COLOR[slug]}">'


def gutter(width: int) -> int:
    return 18 if width < 600 else 32 if width < 1024 else 48


def inner(width: int, pad_y: str) -> str:
    return f'max-width:{PAGE_MAX}px;margin:0 auto;width:100%;padding:{pad_y};padding-inline:{gutter(width)}px'


def ordinal(n: int) -> str:
    return f'{n}{"th" if 10 <= n % 100 <= 20 else {1: "st", 2: "nd", 3: "rd"}.get(n % 10, "th")}'


def skel(w: int | str, h: int, r: int = 3) -> str:
    return f'<span class="sk" style="width:{w if isinstance(w, str) else f"{w}px"};height:{h}px;border-radius:{r}px"></span>'


# ---------------------------------------------------------------------------- identity (fixed)
def identity(ex: dict, state: str, href: str, phone: bool, eyebrow: str, two_line: bool = False) -> str:
    cls = ex['cls']; col = CLASS_COLOR[cls]
    level = 12 if state == 'under20' else ex['level']
    px = 56 if phone else 64
    race = f'{ex["race"]} ' if ex.get('race') else ''
    nw = 'display:inline-flex;align-items:center;gap:6px;white-space:nowrap'
    faction = f'<img src="{FLAT[ex["faction"]]}" alt="" style="width:16px;height:16px;object-fit:contain;flex-shrink:0">{FACTION_WORD[ex["faction"]]}'
    dot = f'<span style="color:{MUTED}">&middot;</span>'
    if two_line:   # narrow containers: two fixed lines, so a separator can never start or end a line
        desc = (f'<span style="display:flex;flex-direction:column;gap:2px;font-size:13px;color:{BODY}"><span>{race}{ex["specname"]} {cls.title()}</span>'
                f'<span style="{nw}">Level {level}{dot}{faction}</span></span>')
    else:
        desc = (f'<span style="{nw};font-size:13px;color:{BODY}">{race}{ex["specname"]} {cls.title()}{dot}Level {level}{dot}{faction}</span>')
    if state == 'stale':
        sync = f'<span style="font-size:12px;color:{MUTED};{CLAMP2}">{ex["source"]} &middot; 19 days ago &middot; <span style="color:{EMBER}">{STALE_TAIL}</span></span>'
    elif state == 'nosync':
        sync = ''
    else:
        sync = f'<span style="font-size:12px;color:{MUTED}">{ex["source"]} &middot; {ex["age"]}</span>'
    # The name box reserves two lines on a phone (a header switch must not move the page); one line from 600 up.
    two = 'display:-webkit-box;-webkit-line-clamp:2;-webkit-box-orient:vertical;overflow:hidden;'
    name = (f'<span class="display idn" title="{ex["name"]}" style="font-size:{22 if phone else 24}px;font-weight:700;color:{col};line-height:1.2;{two}">{ex["name"]}</span>')
    if phone:
        name = f'<span style="display:flex;align-items:flex-end;min-height:53px">{name}</span>'
    ident = (f'<span style="display:flex;flex-direction:column;gap:3px;min-width:0"><span class="label" style="color:{GOLD}">{eyebrow}</span>{name}</span>')
    return (f'<span style="display:flex;flex-direction:column;gap:6px;min-width:0"><a class="idl" href="{href}">{crest_img(cls, px)}{ident}</a>'
            f'<span style="display:flex;flex-direction:column;gap:3px;padding-left:{px + 16}px">{desc}{sync}</span></span>')


# ---------------------------------------------------------------------------- answer (slot, fixed grammar)
def figure_line(a: dict, label: str) -> str:
    col = a.get('figure_color', TEXT)
    out = (f'<span class="label">{label}</span>'
           f'<span class="mono" style="font-size:26px;font-weight:500;line-height:1;display:flex;align-items:baseline;gap:6px 18px;flex-wrap:wrap">'
           f'<span style="color:{col}">{a["figure"]}<span style="font-size:15px;color:{MUTED}"> {a["unit"]}</span></span>'
           + (f'<span style="color:{TEXT}">{a["figure2"]}<span style="font-size:15px;color:{MUTED}"> {a["unit2"]}</span></span>' if a.get('figure2') else '') + '</span>')
    if 'bar' in a:
        out += f'<span class="track"><span class="fillg" style="width:{a["bar"] * 100:.1f}%"></span></span>'
    if a.get('note'):
        link = f' &middot; <a class="tl" href="#fight" style="font-size:12px;white-space:nowrap">{a["note_link"]}&nbsp;→</a>' if a.get('note_link') else ''
        out += f'<span style="font-size:12px;color:{MUTED}">{a["note"]}{link}</span>'
    if a.get('note2'):
        out += f'<a class="tl" href="#report" style="font-size:12px;font-weight:400;color:{MUTED};text-decoration:underline;text-underline-offset:4px">{a["note2"]}</a>'
    return out


def answer(a: dict, state: str) -> str:
    """The slot. The component owns every state's presentation; the page supplies the words."""
    if state in ('loading', 'loading-answer'):
        return f'<span style="display:flex;flex-direction:column;gap:10px">{skel(150, 11)}{skel(150, 26)}{skel("100%", 12)}</span>'
    body = None
    if state == 'under20' and a.get('under20'):
        body = f'<span style="font-size:14px;color:{TEXT}">{a["under20"]}</span>'
    elif state == 'nosync' and a.get('needs_gear'):
        sentence, link = a['nosync']
        body = (f'<span style="display:flex;flex-direction:column;align-items:flex-start;gap:2px;font-size:13px;color:{BODY}">{sentence}'
                f'<a class="tl" href="#setup" style="display:inline-flex;align-items:center;min-height:44px;font-size:13px">{link} →</a></span>')
    elif state == 'empty' and a.get('empty'):
        body = f'<span style="font-size:14px;color:{BODY}">{a["empty"]}</span>'
    elif state == 'error':
        body = (f'<span style="display:flex;flex-direction:column;align-items:flex-start;gap:2px;font-size:13px;color:{BODY}">{a["error"]}'
                f'<a class="tl" href="#retry" style="display:inline-flex;align-items:center;min-height:44px;font-size:13px">Try again</a></span>')
    if body is not None:
        return body
    label = a.get('stale_label', a['label']) if state == 'stale' else a['label']
    return f'<span style="display:flex;flex-direction:column;gap:6px">{figure_line(a, label)}</span>'


# ---------------------------------------------------------------------------- the component
def character_header(width: int, page_key: str, state: str = 'ready', *, ex: dict | None = None, layout: str | None = None,
                     emphasis: str = 'primary', with_answer: bool = True, action: str | None = None, eyebrow: str | None = None,
                     page_answers: dict | None = None, action_state: str = '', id_state: str = '') -> str:
    """state: ready | stale | nosync | under20 | empty | error | loading | loading-answer.
    layout: band (three columns, the 1344 column at 1024 and up) or card (stacked, any width below 900 or a side card).
    emphasis: primary (filled gold, a landing page's lead) or secondary (outline gold, a tool page's context card)."""
    ex = ex or EX
    a = (page_answers or PAGE_ANSWERS)[page_key]
    layout = layout or ('band' if width >= 1024 else 'card')
    phone = layout == 'card'
    eyebrow = eyebrow or 'Your character'
    cols = '' if phone else f'grid-template-columns:minmax(0,{1.15 if width >= 1280 else 1.3}fr) minmax(0,1fr) auto;'
    has_answer = with_answer
    if not phone and not has_answer:
        cols = 'grid-template-columns:minmax(0,1fr) auto;'
    box = (f'display:grid;gap:{16 if phone else 28}px;align-items:center;padding:{"16px" if phone else "18px 22px"};border:1px solid {GOLD};border-radius:6px;'
           f'background:{RAISED};min-height:{0 if phone else BAND_MIN}px;{cols}')
    label = action or a['action']
    cls = 'cta' if emphasis == 'primary' else 'cta2'
    btn = (f'<a class="btn {cls} {action_state}" href="{a["href"]}" style="{"width:100%" if phone else ("min-width:210px" if width < 1280 else "")}">{label}</a>')
    ident = identity(ex, 'ready' if state == 'loading' else state, a['href'] if page_key == 'home' else '#', phone, eyebrow, phone or width < 1280)
    ident = ident.replace('class="idl"', f'class="idl {id_state}"', 1)
    if state == 'loading':
        ready = ident + (answer(a, 'ready') if has_answer else '') + btn
        px = 56 if phone else 64
        shapes = (f'<span style="display:flex;align-items:center;gap:16px"><span class="sk" style="width:{px}px;height:{px}px;border-radius:999px"></span>'
                  f'<span style="display:flex;flex-direction:column;gap:8px">{skel(90, 11)}{skel(200, 24)}{skel(240, 13)}</span></span>')
        mid = f'<span style="display:flex;flex-direction:column;gap:10px">{skel(150, 11)}{skel(150, 26)}{skel("100%", 12)}</span>' if has_answer else ''
        over = (f'<div style="position:absolute;inset:0;display:grid;gap:{16 if phone else 28}px;align-items:center;padding:{"16px" if phone else "18px 22px"};{cols}">'
                f'{shapes}{mid}{skel("100%" if phone else "250px", 48, 4)}</div>')
        return f'<div style="{box}position:relative"><div style="display:contents;visibility:hidden">{ready}</div>{over}</div>'
    parts = ident + (answer(a, state) if has_answer else '') + btn
    return f'<div style="{box}">{parts}</div>'



def signed_out_band() -> str:
    """The one shared signed-out band: one muted line, no button. The nav selector is the only sign-in control."""
    return (f'<div style="display:flex;align-items:center;min-height:44px;font-size:13px;color:{MUTED};border-top:1px solid {SOFT};border-bottom:1px solid {SOFT}">'
            f'Sign in or paste an export to see your character here</div>')


TODAY = (2026, 12, 13)    # the mock's "today"; the example dates are relative to it


def rel_date(days_ago: int) -> str:
    """Relative inside two weeks ("3 days ago"), else "Nov 25": one rule for Rankings, Logs and the header."""
    import datetime
    if days_ago <= 0:
        return 'today'
    if days_ago < 14:
        return '1 day ago' if days_ago == 1 else f'{days_ago} days ago'
    d = datetime.date(*TODAY) - datetime.timedelta(days=days_ago)
    return d.strftime('%b ') + str(d.day)


def omitted(caption: str) -> str:
    """Signed out, no character, or /v1/me failed: the component renders nothing and reserves nothing."""
    return (f'<div style="border:1px dashed {BORDER};border-radius:6px;padding:14px 18px;font-size:13px;color:{MUTED};display:flex;align-items:center;gap:10px">'
            f'<span class="label">Nothing renders here</span>{caption}</div>')


# ---------------------------------------------------------------------------- chrome shared by the page boards
def wide_chrome(html: str) -> str:
    """Footer keeps its content in the 1344 column above 1440 (the nav bar centres itself in gen_nav.py)."""
    return html.replace('padding:28px 48px', f'padding:28px max(48px,calc((100% - {PAGE_MAX}px)/2 + 48px))')


def hero_block(width: int, eyebrow: str, h1: str, job: str, extra: str = '', stamp: str = '') -> str:
    """The band's top: eyebrow, h1, one job line, one pointer line, the one Updated stamp (BiS landing pattern)."""
    phone = width < 600
    st = f'<span style="font-size:12px;color:{MUTED}">{stamp}</span>' if stamp else ''
    head = (f'<div style="display:flex;align-items:baseline;justify-content:space-between;gap:16px">'
            f'<h1 class="display" style="font-size:{26 if phone else 30}px;font-weight:700;color:{TEXT};line-height:1.1">{h1}</h1>{"" if phone else st}</div>')
    return (f'<div style="{inner(width, "22px 0 20px" if phone else "28px 0 24px")};display:flex;flex-direction:column;gap:{10 if phone else 12}px">'
            f'<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>{eyebrow}</span>'
            f'{head}<p style="margin:0;font-size:14px;color:{BODY};max-width:62ch">{job}</p>{extra}{st if phone else ""}</div>')


def page_doc(width: int, door: str, signed_in: bool, band_html: str, main_html: str, ex_nav: dict | None = None) -> str:
    """The real header bar from gen_nav.py, the band, the main column, the footer."""
    from mocklib import footer
    N.CURRENT_DOOR = door
    bar = N.bar(ex_nav or NAV_EX) if signed_in else N.bar(None)
    phone_bar = width < 1024
    foot = footer(True) if phone_bar else wide_chrome(footer(False))
    body = (f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{bar}{band_html}</div>'
            f'<div style="{inner(width, "4px 0 40px")};display:flex;flex-direction:column;gap:{16 if width < 600 else 22}px">{main_html}</div>{foot}')
    return N.doc(CSS_ALL + body, width)


CSS_ALL = CSS


# ---------------------------------------------------------------------------- the states sheet
def states_sheet() -> str:
    cap = lambda t: f'<span class="label" style="color:{GOLD}">{t}</span>'
    note = lambda t: f'<span style="font-size:12px;color:{MUTED};max-width:{PAGE_MAX}px;line-height:1.45">{t}</span>'
    cell = lambda t, h, n='': f'<div style="display:flex;flex-direction:column;gap:8px">{cap(t)}{h}{note(n) if n else ""}</div>'
    desk = lambda t, key, state, n='', **kw: cell(t, character_header(1440, key, state, **kw), n)
    ph = lambda t, key, state, n='', **kw: cell(t, character_header(390, key, state, layout='card', **kw), n)

    rows_desktop = [
        desk('Ready (the Rankings page fills the slots)', 'rankings', 'ready',
             'Identity is fixed: crest, label, name in class colour, descriptor with the flat faction logo, sync line. Answer: label, figure with unit, one note. Action: one 48 px gold button.'),
        desk('Ready, race not known (no Battle.net import): the descriptor drops the race, never a placeholder', 'logs', 'ready', ex=NORACE),
        desk('Ready, home: two figures at the same size (the count and the DPS gain), both read at a glance', 'home', 'ready'),
        desk('Loading, nothing known yet (/v1/me pending): the ready box, hidden, with skeletons drawn over it, so the height is exactly the ready height', 'rankings', 'loading'),
        desk('Loading the answer only (the character is known, the page is fetching its answer; also what a header-selector switch looks like: identity swaps in place, the answer re-fetches)', 'rankings', 'loading-answer'),
        desk('Stale sync (19 days): only the words past the age go ember, the sync line is capped at two lines; the band is 136 px so this is the tallest state and nothing else moves', 'bis', 'stale'),
        desk('Stale sync, a page that does not read gear (Logs): identity only', 'logs', 'stale'),
        desk('No gear on file, a page that needs gear (BiS): the sync line goes, the answer is one sentence and the link that fixes it; the action still works', 'bis', 'nosync'),
        desk('No gear on file, a page that does not need gear (Logs): nothing changes except the sync line is gone', 'logs', 'nosync'),
        desk('Under 20, a page whose answer starts at 20 (BiS): the page supplies the sentence; the action still opens the first band', 'bis', 'under20', action='Fury gear for 20 to 29 →'),
        desk('Under 20, a page with no level floor (Rankings): identical to ready', 'rankings', 'under20'),
        desk('Empty answer (the page has nothing yet): one sentence from the page; the action stays', 'rankings', 'empty'),
        desk('Empty answer, Logs: No logs yet; the action is the way out', 'logs', 'empty'),
        desk('Answer failed (identity and action stay): one sentence and a 44 px Try again', 'rankings', 'error'),
        desk('A long two-word name: the name clamps at two lines, the title holds the whole name, the panel grows once', 'rankings', 'ready', ex=LONG),
        cell('Signed out, no character, or /v1/me failed: the one shared band, the same on every page', signed_out_band(),
             'One 44 px muted line, no button and no link: the nav selector (Sign in or paste an export) is the only sign-in control and pages add none of their own.'),
    ]
    rows_phone = [
        ph('Ready', 'rankings', 'ready'), ph('Loading (exactly the ready height)', 'rankings', 'loading'), ph('Home, two figures', 'home', 'ready'), ph('Signed out band', 'rankings', 'ready') if False else cell('Signed out', signed_out_band()), ph('Stale (BiS)', 'bis', 'stale'),
        ph('No gear on file (BiS)', 'bis', 'nosync'), ph('Under 20 (BiS)', 'bis', 'under20', action='Fury gear for 20 to 29 →'),
        ph('Empty (Logs)', 'logs', 'empty'), ph('Error', 'rankings', 'error'),
        ph('Long name: two lines reserved on a phone', 'rankings', 'ready', ex=LONG),
        ph('Logs, ready', 'logs', 'ready'),
    ]
    ctrls = ''.join(cell(t, html) for t, html in (
        ('Action, rest', f'<a class="btn cta" href="#x">See your parses →</a>'),
        ('Action, hover', f'<a class="btn cta is-hover" href="#x">See your parses →</a>'),
        ('Action, focus-visible', f'<a class="btn cta is-focus" href="#x">See your parses →</a>'),
        ('Action, active', f'<a class="btn cta is-active" href="#x">See your parses →</a>'),
        ('Action, secondary (tool-page card), rest', f'<a class="btn cta2" href="#x">Send this build to the addon</a>'),
        ('Action, secondary, hover', f'<a class="btn cta2 is-hover" href="#x">Send this build to the addon</a>'),
        ('Action, secondary, focus-visible', f'<a class="btn cta2 is-focus" href="#x">Send this build to the addon</a>'),
        ('Identity link, hover: the name underlines', f'<a class="idl is-hover" href="#x"><span class="display idn" style="font-size:24px;font-weight:700;color:{CLASS_COLOR["warrior"]}">{EX["name"]}</span></a>'),
        ('Identity link, focus-visible', f'<a class="idl is-focus" href="#x"><span class="display idn" style="font-size:24px;font-weight:700;color:{CLASS_COLOR["warrior"]}">{EX["name"]}</span></a>'),
    ))
    pages = [
        desk('Home: answer is the gear gap, action jumps to the upgrade table (loses the 34 px name and 84 px crest; the Talents and Simulator cards move under it)', 'home', 'ready'),
        desk('BiS landing: answer is the worn count with its bar, action opens the spec list at the band', 'bis', 'ready'),
        desk('Rankings: answer is the rank in the cohort, action opens the character page', 'rankings', 'ready'),
        desk('Logs: answer is the latest parse, action opens the file picker', 'logs', 'ready'),
    ]
    cards = [
        cell('BiS spec page card, 440 px wide (replaces CharacterCard.svelte): stacked, secondary action, the same count', f'<div style="width:440px">{character_header(440, "bis", "ready", layout="card", emphasis="secondary", action="Send this list to the addon")}</div>'),
        cell('Planner card, 440 px wide (replaces PlannerCharacterCard.svelte): no answer slot, secondary action', f'<div style="width:440px">{character_header(440, "bis", "ready", layout="card", emphasis="secondary", with_answer=False, action="Send this build to the addon")}</div>'),
    ]
    h2 = lambda t: f'<h2 class="display" style="font-size:18px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:{TEXT}">{t}</h2>'
    body = (f'<div style="{inner(1440, "36px 0")};display:flex;flex-direction:column;gap:22px">'
            f'<h1 class="display" style="font-size:22px;color:{TEXT}">Character header: every state</h1>'
            f'{h2("1440, every state")}{"".join(rows_desktop)}'
            f'{h2("390, every state")}<div style="display:flex;flex-wrap:wrap;gap:18px 24px;align-items:flex-start">{"".join(f"<div style=width:390px>{r}</div>" for r in rows_phone)}</div>'
            f'{h2("Controls")}<div style="display:flex;flex-wrap:wrap;gap:18px 24px;align-items:flex-start">{ctrls}</div>'
            f'{h2("What each page fills")}{"".join(pages)}'
            f'<div style="display:flex;flex-wrap:wrap;gap:18px 24px;align-items:flex-start">{"".join(cards)}</div></div>')
    return page(CSS + body, 1440)


BOARDS = {'character-header-states': (states_sheet, 1440)}


def shoot(html: str, name: str, width: int, height: int = 900) -> Path:
    scratch = Path(tempfile.mkdtemp(prefix='gen_ch_'))
    src = scratch / f'{name}.html'
    src.write_text(html)
    RENDERS.mkdir(exist_ok=True)
    out = RENDERS / f'{name}.png'
    subprocess.run(['npx', 'playwright', 'screenshot', '--browser=chromium', f'--viewport-size={width},{height}', '--full-page',
                    '--wait-for-timeout=3000', src.as_uri(), str(out)], check=True, capture_output=True, cwd=WEB)
    return out


def main(names: list[str]) -> None:
    for name in names:
        build, width = BOARDS[name]
        print(shoot(build(), name, width))


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))
