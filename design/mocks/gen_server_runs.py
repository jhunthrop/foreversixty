"""Server runs (design/specs/2026-10-24-server-runs.md).

  python3 design/mocks/gen_server_runs.py   -> renders/server-runs-states.png

One sheet at 1440 wide holding 1440 boards and 390 frames: the sim page for a payer (every row state, empty,
loading, error), the Run again states, the character band with the "done while away" line (one run, three,
phone), and the account Server sims row. The band is gen_character_header.character_header, so the header is
the shipped one; only the footer line is new. Every figure, name and time is an EXAMPLE. Specs, crests and
copy strings are the site's own; the copy below mirrors the spec's section 6 and is the same wording.
"""
import os
import subprocess
import sys
import tempfile
from pathlib import Path

import gen_character_header as CH
import gen_nav as N
from mocklib import WEB, RENDERS, GOLD, MUTED, TEXT, BODY, BORDER, SOFT, RAISED, CLASS_COLOR, crest, nav_phone

WIPE = '#f0736b'
CARD_TOP = '#131824'
UNREAD = f'color-mix(in srgb,{GOLD} 8%,{RAISED})'
LINE_WARM = '#4a4030'
SPEC_LABEL = dict([('mage-frost', 'Frost Mage'), ('warlock-destruction', 'Destruction Warlock'), ('hunter-marksmanship', 'Marksmanship Hunter'),
                   ('priest-shadow', 'Shadow Priest'), ('warrior-fury', 'Fury Warrior')])
KIND = dict(run='Sim', gear='Gear', talents='Talents', drops='Drops', weights='Weights')
EMPTY_TEXT = 'No server runs yet. Choose Run on our servers when you start a sim, and it keeps going after you close this tab.'

# EXAMPLE runs (the API row shape, spec section 10): state, spec, kind, title, headline, age line, extras.
RUNS = dict(
    running_est=dict(state='running', spec='mage-frost', kind='run', title='', age='Started 2 minutes ago', est='takes about 3 minutes'),
    running_late=dict(state='running', spec='warlock-destruction', kind='gear', title='Pre-raid gear', age='Started 9 minutes ago', est='taking longer than usual'),
    running_none=dict(state='running', spec='hunter-marksmanship', kind='drops', title='Molten Core drops', age='Started 2 minutes ago', est=None),
    starting=dict(state='queued', spec='mage-frost', kind='run', title='', age='Started just now'),
    done_new=dict(state='done', spec='mage-frost', kind='run', title='', head='1,204 DPS', age='Finished 14 minutes ago', unseen=True),
    done_seen=dict(state='done', spec='hunter-marksmanship', kind='drops', title='Ragnaros drops', head='3 upgrades on Ragnaros', age='Finished 3 hours ago'),
    done_gear=dict(state='done', spec='warlock-destruction', kind='gear', title='Pre-raid gear', head="+41 DPS from Vis'kag", age='Finished yesterday'),
    failed=dict(state='error', spec='priest-shadow', kind='weights', title='Weights for raid gear', reason='stopped before completion', age='Ended 2 hours ago'),
    failed_start=dict(state='error', spec='warrior-fury', kind='run', title='', reason='the job could not be started', age='Ended 1 day ago'),
)

CSS_RAW = '''<style>
.sr-row{display:grid;align-items:center;gap:14px;padding:10px 14px;min-height:64px;border-bottom:1px solid SOFT_}
.sr-row:last-child{border-bottom:0}
.sr-link{color:inherit}
.sr-link:hover,.sr-link.is-hover{background:CARDTOP_}
.sr-link:hover .sr-t,.sr-link.is-hover .sr-t{text-decoration:underline;text-underline-offset:4px}
.sr-link:hover .chev,.sr-link.is-hover .chev{color:GOLD_}
.sr-link:active,.sr-link.is-active{background:#171d2b}
.sr-link:focus-visible,.sr-link.is-focus{outline:2px solid GOLD_;outline-offset:-2px}
.dot{display:inline-block;width:8px;height:8px;border-radius:50%;background:GOLD_;flex-shrink:0}
.dot.pulse{animation:pulse 1.6s ease-in-out infinite}
@keyframes pulse{50%{opacity:.35}}
@media (prefers-reduced-motion:reduce){.dot.pulse{animation:none}}
.again{display:inline-flex;align-items:center;justify-content:center;border:1px solid WARM_;border-radius:4px;color:TEXT_;background:RAISED_;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;min-width:112px;padding:0 14px;white-space:nowrap}
.again:hover,.again.is-hover{border-color:GOLD_;color:#f5d27a}
.again:focus-visible,.again.is-focus{outline:2px solid GOLD_;outline-offset:2px}
.again:active,.again.is-active{background:CARDTOP_}
.again.is-off{color:MUTED_;border-color:BORDER_;background:transparent}
.dline{color:inherit}
.dline:hover .dtxt,.dline.is-hover .dtxt{text-decoration:underline;text-underline-offset:4px}
.dline:hover .dgo,.dline.is-hover .dgo{color:#f5d27a}
.dline:focus-visible,.dline.is-focus{outline:2px solid GOLD_;outline-offset:-2px;border-radius:4px}
.dx{width:44px;height:44px;display:grid;place-items:center;color:MUTED_;background:transparent;border:0;border-radius:50%;flex-shrink:0}
.dx:hover,.dx.is-hover{color:TEXT_}
.dx:hover i,.dx.is-hover i{background:#ffffff10}
.dx i{width:32px;height:32px;border-radius:50%;display:grid;place-items:center;font-style:normal;font-size:18px;line-height:1}
.dx.is-focus{outline:2px solid GOLD_;outline-offset:-2px}
.dx.is-active i{background:#ffffff20}
.more{color:MUTED_;text-decoration:underline;text-underline-offset:3px}
.more:hover,.more.is-hover{color:TEXT_}
</style>'''
CSS = (CSS_RAW.replace('SOFT_', SOFT).replace('CARDTOP_', CARD_TOP).replace('GOLD_', GOLD).replace('WARM_', LINE_WARM)
       .replace('TEXT_', TEXT).replace('RAISED_', RAISED).replace('MUTED_', MUTED).replace('BORDER_', BORDER))


def cap(t):
    return f'<span class="label" style="color:{GOLD}">{t}</span>'


def note(t):
    return f'<span style="font-size:12px;color:{MUTED};line-height:1.45">{t}</span>'


def cell(title, html, n='', w=None):
    size = f'width:{w}px;flex-shrink:0' if w else ''
    nt = note(n) if n else ''
    return f'<div style="display:flex;flex-direction:column;gap:8px;{size}">{cap(title)}{html}{nt}</div>'


def h2(t):
    return f'<h2 class="display" style="font-size:18px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:{TEXT}">{t}</h2>'


def section_title(t):
    return f'<h2 class="display" style="font-size:15px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:{TEXT}">{t}</h2>'


def state_block(r, phone):
    """Line 1 and line 2 of the state block (spec 3.3)."""
    al = 'left' if phone else 'right'
    jc = 'flex-start' if phone else 'flex-end'
    age = r['age']
    line2 = f'<span class="mono" style="font-size:12px;color:{MUTED};text-align:{al}">{age}</span>'
    row1 = f'display:flex;align-items:center;gap:8px;justify-content:{jc};font-size:14px;color:{TEXT};'
    if r['state'] == 'queued':
        l1 = f'<span style="{row1}"><span class="dot"></span>Starting</span>'
    elif r['state'] == 'running':
        est = r.get('est')
        tail = f'<span style="color:{MUTED}"> &middot; {est}</span>' if est else ''
        l1 = f'<span style="{row1}"><span class="dot pulse"></span><span>Running{tail}</span></span>'
    elif r['state'] == 'done':
        head = r['head']
        l1 = f'<span class="mono" style="font-size:15px;color:{TEXT};text-align:{al}">{head}</span>'
    else:
        why = r.get('reason')
        reason = f': {why}' if why else ''
        l1 = f'<span style="font-size:14px;color:{WIPE};text-align:{al}">Failed{reason}</span>'
    return l1, line2


def again_button(state, phone, mode=''):
    txt = dict([('', 'Run again'), ('pending', 'Starting…'), ('started', 'Started')])[state]
    extra = f' is-{mode}' if mode else ''
    off = ' is-off' if state in ('pending', 'started') else ''
    busy = ' aria-busy="true"' if state == 'pending' else ''
    height = 44 if phone else 36
    width = 'width:100%' if phone else ''
    return f'<span class="again{extra}{off}"{busy} style="height:{height}px;{width}">{txt}</span>'


def run_row(r, phone=False, mode='', again='', again_mode='', err=''):
    """mode: '' | hover | focus | active (the row, when it is a link). again: '' | pending | started."""
    slug = r['spec'].split('-')[0]
    color = CLASS_COLOR[slug]
    img = f'<img class="crest" src="{crest(slug)}" alt="" width="32" height="32" style="width:32px;height:32px;--c:{color}">'
    label = SPEC_LABEL[r['spec']]
    kind = KIND[r['kind']]
    title = r['title'] or label
    meta = f'{label} &middot; {kind}' if r['title'] else kind
    dot = '<span class="dot" style="margin-right:8px" title="Not opened yet"></span>' if r.get('unseen') else ''
    clip = 'white-space:nowrap;overflow:hidden;text-overflow:ellipsis'
    tline = f'<span style="display:flex;align-items:center;min-width:0">{dot}<span class="sr-t" style="font-size:15px;font-weight:600;color:{TEXT};{clip}">{title}</span></span>'
    metaline = f'<span style="font-size:13px;color:{MUTED};{clip}">{meta}</span>'
    l1, l2 = state_block(r, phone)
    link = r['state'] == 'done'
    cls = 'sr-row' + (' sr-link' if link else '') + (f' is-{mode}' if mode else '')
    ground = f'background:{UNREAD};' if r.get('unseen') else ''
    tag = 'a href="#result"' if link else 'div'
    close = 'a' if link else 'div'
    btn = ''
    if r['state'] == 'error':
        btn = again_button(again, phone, again_mode)
    elif link:
        btn = f'<span class="chev" style="color:{MUTED};font-size:18px">&rsaquo;</span>'
    errline = f'<span style="font-size:13px;color:{WIPE}">{err}</span>' if err else ''
    if phone:
        extra_btn = f'<div style="grid-column:2">{btn}</div>' if r['state'] == 'error' else ''
        extra_err = f'<div style="grid-column:2">{errline}</div>' if err else ''
        body = (f'<div style="grid-column:1;grid-row:1/span 4;align-self:start;padding-top:2px">{img}</div>'
                f'<div style="grid-column:2;display:flex;flex-direction:column;gap:3px;min-width:0">{tline}{metaline}</div>'
                f'<div style="grid-column:2;display:flex;flex-direction:column;gap:2px">{l1}{l2}</div>{extra_btn}{extra_err}')
        return f'<{tag} class="{cls}" style="grid-template-columns:32px minmax(0,1fr);row-gap:8px;padding:12px 14px;{ground}">{body}</{close}>'
    body = (f'{img}<div style="display:flex;flex-direction:column;gap:3px;min-width:0">{tline}{metaline}</div>'
            f'<div style="display:flex;flex-direction:column;gap:2px;align-items:flex-end">{l1}{l2}{errline}</div>'
            f'<div style="width:112px;display:flex;justify-content:flex-end">{btn}</div>')
    return f'<{tag} class="{cls}" style="grid-template-columns:32px minmax(0,1.4fr) minmax(0,1.2fr) auto;{ground}">{body}</{close}>'


def panel(inner_html, footer=''):
    foot = f'<div style="padding:8px 14px;font-size:12px;color:{MUTED};border-top:1px solid {SOFT}">{footer}</div>' if footer else ''
    return f'<div style="border:1px solid {BORDER};border-radius:6px;background:{RAISED};overflow:hidden">{inner_html}{foot}</div>'


def loading_rows(phone):
    h = 92 if phone else 64
    s1 = CH.skel(180, 15)
    s2 = CH.skel(120, 12)
    s3 = '' if phone else CH.skel(140, 14)
    one = (f'<div class="sr-row" style="grid-template-columns:32px minmax(0,1fr) auto;min-height:{h}px">'
           f'<span class="sk" style="width:32px;height:32px;border-radius:50%"></span>'
           f'<span style="display:flex;flex-direction:column;gap:8px">{s1}{s2}</span>{s3}</div>')
    return one * 3


def empty_row():
    return f'<div style="display:flex;align-items:center;min-height:64px;padding:8px 14px;font-size:14px;color:{MUTED};line-height:1.4">{EMPTY_TEXT}</div>'


def error_row():
    return (f'<div style="display:flex;align-items:center;flex-wrap:wrap;gap:12px;min-height:64px;padding:8px 14px;font-size:14px;color:{MUTED}">'
            f'Your server runs did not load.<span style="display:inline-flex;align-items:center;min-height:44px;padding:0 12px;border:1px solid #3a3326;border-radius:4px;font-size:13px;font-weight:600;color:{TEXT}">Try again</span></div>')


def server_runs(width, keys=(), rows_html=None, footer='', desktop=False):
    phone = width < 768 and not desktop
    inner_html = rows_html if rows_html is not None else ''.join(run_row(RUNS[k], phone) for k in keys)
    return f'<section style="display:flex;flex-direction:column;gap:12px">{section_title("Your server runs")}{panel(inner_html, footer)}</section>'


# ---------------------------------------------------------------------------- the example character (a Frost Mage, the run's own spec)
MAGE = dict(CH.EX)
MAGE.update(name='Frostbyte', cls='mage', race='Gnome', specname='Frost', source='Addon', age='12 minutes ago')
NAV_MAGE = N.example('Frostbyte', 'mage-frost', 60, 'Living Flame', 'alliance', 'Addon', '12 minutes ago', current=True)
SENTENCE = 'Your Frost Mage sim finished: '
MC = CLASS_COLOR['mage']


def mode_cls(mode):
    return f'is-{mode}' if mode else ''


def done_line(phone, more=0, link_mode='', dx_mode='', more_mode=''):
    """The footer row of the band (spec 3.4)."""
    dps = f'<span class="mono" style="color:{TEXT}">1,204 DPS</span>'
    lm = mode_cls(link_mode)
    mm = mode_cls(more_mode)
    if phone:
        go = f'<span class="dgo" style="color:{GOLD};font-weight:700"> →</span>'
        text = (f'<a class="dline {lm}" href="#result" style="display:flex;align-items:center;min-height:44px;font-size:13px;color:{TEXT};line-height:1.4">'
                f'<span class="dtxt">{SENTENCE}{dps}{go}</span></a>')
        more_html = f'<a class="more {mm}" href="#runs" style="display:flex;align-items:center;min-height:44px;font-size:12px">and {more} more</a>' if more else ''
        left = f'<span style="display:flex;flex-direction:column;min-width:0">{text}{more_html}</span>'
        pad = 'padding:6px 0 6px 14px;min-height:56px'
    else:
        go = f'<span class="dgo" style="color:{GOLD};font-weight:700;white-space:nowrap">Open result →</span>'
        text = (f'<a class="dline {lm}" href="#result" style="display:inline-flex;align-items:center;min-height:44px;gap:10px;font-size:13px;color:{TEXT}">'
                f'<span class="dtxt">{SENTENCE}{dps}</span>{go}</a>')
        more_html = f'<a class="more {mm}" href="#runs" style="display:inline-flex;align-items:center;min-height:44px;font-size:13px">and {more} more</a>' if more else ''
        left = f'<span style="display:flex;align-items:center;gap:12px;min-width:0">{text}{more_html}</span>'
        pad = 'padding:0 0 0 22px;min-height:44px'
    label = 'Dismiss: your Frost Mage sim finished'
    dxm = mode_cls(dx_mode)
    dx = f'<button class="dx {dxm}" aria-label="{label}"><i>&times;</i></button>'
    return (f'<div role="status" style="display:flex;align-items:center;justify-content:space-between;gap:8px;{pad};'
            f'border-top:1px solid {SOFT};background:{UNREAD}">{left}{dx}</div>')


def band(width, line=None, state='ready'):
    """The shipped band (gen_character_header) with the footer row inside the same gold border."""
    layout = 'band' if width >= 1024 else 'card'
    html = CH.character_header(width, 'home', state, ex=MAGE, layout=layout)
    if line is None:
        return html
    html = html.replace(f'border:1px solid {GOLD};border-radius:6px;', 'border:0;', 1)
    return f'<div style="border:1px solid {GOLD};border-radius:6px;overflow:hidden;background:{RAISED}">{html}{line}</div>'


# ---------------------------------------------------------------------------- the sim page for a payer
def run_block(phone):
    img = f'<img class="crest" src="{crest("mage")}" alt="" width="36" height="36" style="width:36px;height:36px;--c:{MC}">'
    desc = '' if phone else f'<span style="font-size:13px;color:{MUTED}">Gnome Frost Mage &middot; 60</span>'
    full = 'width:100%' if phone else ''
    btn = f'<span class="btn btn-fill" style="height:44px;padding:0 20px;{full}">Run sim</span>'
    return (f'<div class="panel" style="display:flex;flex-wrap:wrap;align-items:center;justify-content:space-between;gap:12px;padding:16px">'
            f'<span style="display:flex;align-items:center;gap:12px">{img}<span style="font-size:17px;font-weight:600;color:{MC}">Frostbyte</span>{desc}</span>{btn}</div>')


def landing_stub():
    img = f'<img class="crest" src="{crest("mage")}" alt="" width="32" height="32" style="width:32px;height:32px;--c:{MC}">'
    row = (f'<div style="display:flex;align-items:center;gap:12px;min-height:64px;padding:0 12px">{img}<span style="color:{MC};font-weight:600">Frostbyte</span>'
           f'<span style="font-size:13px;color:{MUTED}">Gnome Frost Mage &middot; Level 60</span></div>')
    return f'<div style="display:flex;flex-direction:column;gap:10px">{section_title("Your characters")}{panel(row)}</div>'


def sim_page(width, runs_html):
    phone = width < 768
    gut = CH.gutter(width)
    N.CURRENT_DOOR = 'Simulator'
    bar = nav_phone(True, dict([('class', 'mage')])) if phone else N.bar(NAV_MAGE)
    size = 26 if phone else 30
    h1 = (f'<div style="display:flex;flex-direction:column;gap:8px"><h1 class="display" style="font-size:{size}px;font-weight:700;color:{TEXT}">Simulator</h1>'
          f'<p style="margin:0;font-size:14px;color:{BODY};max-width:62ch">Simulates damage for the specs the engine models. Healing and tanking are not simulated yet.</p></div>')
    gap = 18 if phone else 24
    maxw = CH.PAGE_MAX
    main = (f'<div style="max-width:{maxw}px;margin:0 auto;padding:28px {gut}px 40px;display:flex;flex-direction:column;gap:{gap}px">'
            f'{h1}{runs_html}{run_block(phone)}{landing_stub()}</div>')
    return f'<div style="width:{width}px;border:1px solid {SOFT};background:var(--bg);flex-shrink:0">{bar}{main}</div>'


# ---------------------------------------------------------------------------- the account Premium rows (lane 1 look)
def account_card(width, server_value, live=True):
    def row(label, value, link=False, muted=False):
        color = MUTED if muted else TEXT
        deco = 'text-decoration:underline;' if link else ''
        return (f'<div style="display:grid;grid-template-columns:10rem minmax(0,1fr);column-gap:12px;align-items:baseline;min-height:28px;font-size:14px">'
                f'<span style="color:{MUTED}">{label}</span><span style="color:{color};{deco}">{value}</span></div>')
    rows = (row('Server sims', server_value, link=live) + row('History charts', 'Ships November', muted=True)
            + row('Discord notifications', 'Ships December', muted=True) + row('Supporter mark', 'Shown on your characters', link=True))
    title = (f'<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:7px;height:7px;border-radius:50%;background:{GOLD};display:inline-block"></i>Plan</span>')
    plan = f'<span style="font-size:14px;color:{MUTED}">Plan &middot; <span style="color:{TEXT}">Premium.</span> <span style="text-decoration:underline;color:{TEXT}">Manage billing</span></span>'
    return (f'<div style="width:{width}px;flex-shrink:0;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:22px 24px;display:flex;flex-direction:column;gap:14px">'
            f'{title}{plan}<div style="display:flex;flex-direction:column;gap:4px">{rows}</div></div>')


# ---------------------------------------------------------------------------- the sheet
def runs_panel(width, keys, **kw):
    return server_runs(width, keys, **kw)


def frame(width, inner_html):
    return f'<div style="width:{width}px;flex-shrink:0;display:flex;flex-direction:column;gap:18px">{inner_html}</div>'


def row_cell(title, row_html, n='', w=1248):
    return cell(title, panel(row_html), n, w)


def desktop_states():
    R = RUNS
    p = False
    lapsed = (f'Server sims are part of Premium. <span style="text-decoration:underline;color:{TEXT}">See plans</span>')
    cells = [
        row_cell('Running, with an estimate', run_row(R['running_est'], p) + run_row(R['starting'], p),
                 'Pulsing dot, Running in strong, the estimate muted. Starting lasts seconds and has no estimate. No bar, no percentage.'),
        row_cell('Running: no estimate yet, and overdue', run_row(R['running_none'], p) + run_row(R['running_late'], p),
                 'Fewer than three finished runs for the spec: Running alone. Past 1.5 times the estimate: taking longer than usual.'),
        row_cell('Done, unread and read', run_row(R['done_new'], p) + run_row(R['done_seen'], p) + run_row(R['done_gear'], p),
                 'Unread: gold dot and a 8 percent gold ground. The headline is the API sentence.'),
        row_cell('Done row: hover, focus-visible, active', run_row(R['done_seen'], p, mode='hover') + run_row(R['done_seen'], p, mode='focus') + run_row(R['done_seen'], p, mode='active'),
                 'Hover: card ground, title underlines, the chevron goes gold. Focus: inset 2 px gold outline.'),
        row_cell('Failed: Run again at rest, hover, focus, active', run_row(R['failed'], p) + run_row(R['failed'], p, again_mode='hover') + run_row(R['failed'], p, again_mode='focus') + run_row(R['failed'], p, again_mode='active'),
                 'The reason is the job log in plain words after Failed. The row is not a link; the button is the control.'),
        row_cell('Failed: pending, started, could not start, plan lapsed',
                 run_row(R['failed'], p, again='pending') + run_row(R['failed'], p, again='started')
                 + run_row(R['failed_start'], p, err='Could not start it again. Try again.') + run_row(R['failed_start'], p, err=lapsed),
                 'Pending and started disable the button. A failed start keeps the button at rest and says why under the state.'),
        cell('Empty (payer, no runs yet)', server_runs(660, rows_html=empty_row()), 'One 64 px row so the Run sim block below barely moves.', 660),
        cell('Loading (the reserved box)', server_runs(660, rows_html=loading_rows(False)), 'Three skeleton rows, 64 px each; the reserve is the last known row count.', 660),
        cell('Error', server_runs(660, rows_html=error_row()), 'The run block and everything below render as usual.', 660),
        cell('More than five runs', server_runs(1248, ['running_est', 'done_new', 'failed', 'done_seen', 'done_gear'], footer='Older finished runs are in Your sims below.', desktop=True),
             'Five rows, newest first, and one 12 px line.', 1248),
    ]
    return '<div style="display:flex;flex-wrap:wrap;gap:24px 24px;align-items:flex-start">' + ''.join(cells) + '</div>'


def phone_frames():
    keys = ['running_est', 'running_none', 'done_new', 'done_seen', 'failed']
    pop = sim_page(390, server_runs(390, keys))
    emp = sim_page(390, server_runs(390, rows_html=empty_row()))
    ld = server_runs(390, rows_html=loading_rows(True))
    er = server_runs(390, rows_html=error_row())
    both = sim_page(390, ld + er)
    return (f'<div style="display:flex;gap:24px;align-items:flex-start">'
            f'{cell("390: five runs, newest first", pop, "", 390)}{cell("390: empty", emp, "", 390)}{cell("390: loading, then error (stacked for the sheet)", both, "", 390)}</div>')


def bands():
    one = band(1440, done_line(False))
    three = band(1440, done_line(False, more=2))
    ready = band(1440)
    hov = band(1440, done_line(False, link_mode='hover', dx_mode='hover'))
    foc = band(1440, done_line(False, link_mode='focus'))
    dxf = band(1440, done_line(False, dx_mode='focus'))
    desk = (f'<div style="display:flex;flex-direction:column;gap:22px">'
            f'{cell("1440, one run finished while away", one, "The footer row is inside the same gold border: 44 px, edge to edge, 8 percent gold ground. The filled gold button stays the only filled element.")}'
            f'{cell("1440, three runs: the newest and and 2 more", three)}'
            f'{cell("1440, dismissed: the band returns to its ready height", ready)}'
            f'{cell("1440, link hover and dismiss hover", hov)}'
            f'{cell("1440, link focus-visible", foc)}'
            f'{cell("1440, dismiss focus-visible", dxf)}</div>')
    p1 = band(390, done_line(True))
    p3 = band(390, done_line(True, more=2))
    p0 = band(390)
    ph = (f'<div style="display:flex;gap:24px;align-items:flex-start">'
          f'{cell("390, one run (56 px footer)", p1, "", 390)}{cell("390, three runs (80 px footer)", p3, "", 390)}{cell("390, dismissed", p0, "", 390)}</div>')
    return desk, ph


def account_rows():
    on = account_card(395, 'On · 4 runs this month')
    one = account_card(395, 'On · 1 run this month')
    none = account_card(395, 'On · no runs yet this month')
    old = account_card(395, 'On')
    off = account_card(395, 'Browser lane only', live=False)
    ph = account_card(354, 'On · 4 runs this month')
    items = [cell('1440, four runs', on), cell('One run', one), cell('No runs yet', none), cell('No count sent (older /v1/me)', old), cell('Free member: unchanged', off), cell('390 (card 354 px)', ph)]
    return '<div style="display:flex;flex-wrap:wrap;gap:22px 24px;align-items:flex-start">' + ''.join(items) + '</div>'


def sheet():
    keys = ['running_est', 'running_none', 'done_new', 'done_seen', 'failed']
    page1440 = sim_page(1440, server_runs(1440, keys))
    desk_band, phone_band = bands()
    gut = CH.gutter(1440)
    maxw = CH.PAGE_MAX
    title = f'<h1 class="display" style="font-size:22px;color:{TEXT}">Server runs and done while away: every state</h1>'
    body = (f'<div style="max-width:{maxw}px;margin:0 auto;padding:36px {gut}px;display:flex;flex-direction:column;gap:24px">{title}'
            f'{h2("Sim page, payer, 1440")}</div>{page1440}'
            f'<div style="max-width:{maxw}px;margin:0 auto;padding:36px {gut}px;display:flex;flex-direction:column;gap:24px">'
            f'{h2("Sim page, 390")}{phone_frames()}'
            f'{h2("Row and list states (desktop grid)")}{desktop_states()}'
            f'{h2("The done line in the character band")}{desk_band}{phone_band}'
            f'{h2("Account page: the Server sims row")}{account_rows()}</div>')
    return N.doc(CH.CSS + CSS + body, 1440)


def shoot(html, name, width):
    scratch = Path(tempfile.mkdtemp(prefix='gen_sr_'))
    src = scratch / f'{name}.html'
    src.write_text(html)
    RENDERS.mkdir(exist_ok=True)
    out = RENDERS / f'{name}.png'
    local = WEB / 'node_modules/.bin/playwright'
    override = os.environ.get('FS_PLAYWRIGHT')
    runner = [override] if override else [str(local)] if local.exists() else ['npx', 'playwright']
    subprocess.run(runner + ['screenshot', '--browser=chromium', f'--viewport-size={width},900', '--full-page', '--wait-for-timeout=3000', src.as_uri(), str(out)],
                   check=True, capture_output=True, cwd=WEB)
    return out


if __name__ == '__main__':
    print(shoot(sheet(), 'server-runs-states', 1440))
