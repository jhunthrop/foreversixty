"""Lane L7 review renders: the Officer tab as built (addon/ForeverSixty/surfaces/OfficerTab.lua, Officer.lua,
OfficerLoot.lua).

    python3 design/mocks/addon_hub_l7.py            # every state
    python3 design/mocks/addon_hub_l7.py l7-raid-day

Writes design/reviews/2026-11-13-addon-hub/l7/<state>.png at device scale 1: the window at its real 740 px
width, drawn with addon_hub.py's frame and tokens, holding exactly what the built tab draws from the data it
has. Where the approved boards (hub-officer.png, hub-officer-loot.png) differ from the build, the render
follows the build: the ringed class crest before every raider (one image per character, as the Guild tab's
standings), the dot legend `heard earlier` (an addon heard this hour but not in the last 10 minutes; a
member with sharing off sends nothing, so the addon cannot tell that apart), the filters `Show all`,
`Not signed up (N)` and `Online (N)`, the readiness rows the data can read (old addon versions, attunement,
the raid's supplies by kind, the last ready check), the push's confirm drawn in place of its button, and
the session's gain column headed by its unit.

Every character and guild is made up (Velisandre and Thornwake Company, as the boards).
"""
import os
import sys

import addon_hub as hub
from addon_hub import (ADDON, CLASS_HEX, CLIENT_GOLD, GUILD, ROSTER, TIER_HEX, esc, icon, item_icon, item_name,
                       section, tier_of, window)

OUT = hub.OUT / 'l7'
COLS = (104, 40, 38, 72, 38, 36, 24)
CSS = f'''
.sh > span:first-child {{ font-family: Title, serif; font-size: 14px; font-weight: 400; color: {ADDON['cream']};
  text-shadow: 1px -1px 0 #000; }}
.ros {{ table-layout: fixed; width: 388px; border-collapse: collapse; }}
.ros td, .ros th {{ height: 18px; padding: 0 6px 0 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;
  font-weight: 400; text-align: left; }}
.ros th {{ height: 16px; color: {ADDON['muted']}; }}
.ros td:last-child, .ros th:last-child {{ padding-right: 0; }}
.ros .n {{ text-align: right; font-variant-numeric: tabular-nums; }}
.ros .on {{ color: {ADDON['gold']}; }}
.rcr {{ display: inline-block; vertical-align: -3px; margin-right: 4px; }}
.dot {{ display: inline-block; width: 6px; height: 6px; border-radius: 50%; margin-left: 4px; vertical-align: 1px; }}
.rd {{ display: flex; height: 18px; align-items: center; }}
.rd .l {{ width: 84px; color: {ADDON['muted']}; flex: none; }}
.rd .c {{ width: 48px; text-align: right; flex: none; font-variant-numeric: tabular-nums; }}
.rd .v {{ margin-left: 6px; flex: 1; min-width: 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.lt {{ table-layout: fixed; width: 388px; border-collapse: collapse; }}
.lt td, .lt th {{ height: 24px; padding: 0 6px 0 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;
  font-weight: 400; text-align: left; }}
.lt th {{ height: 16px; color: {ADDON['muted']}; }}
.lt .n {{ text-align: right; font-variant-numeric: tabular-nums; }}
.lt td:last-child {{ padding-right: 0; text-align: right; overflow: visible; }}
.tip {{ position: absolute; }}
'''

WHO = {r[0]: r for r in ROSTER}
SIGN = {'In': 'ok', 'Late': 'cg', 'Bench': '', 'Out': 'wa'}


def name(n, cls=None):
    cls = cls or WHO[n][1]
    return f'<span class="cls" style="color:{CLASS_HEX[cls]}">{esc(n)}</span>'


def dot(state):
    colour = ADDON['success'] if state == 'live' else ADDON['muted']
    return f'<span class="dot" style="background:{colour}"></span>'


def raider(n, live=None, cls=None):
    cls = cls or WHO[n][1]
    return f'<span class="rcr">{hub.crest(cls, 14)}</span>{name(n, cls)}{dot(live) if live else ""}'


def roster(rows, pack=True, signups=True, sort='signup', shown=None, raiders=31, unsigned=16, online=24,
           filt='raiders'):
    order = {'In': 0, 'Late': 1, 'Bench': 2, 'Out': 3, '': 4}
    if sort == 'signup':
        rows = sorted(rows, key=lambda r: (order[r[5]] if signups else 0, -r[3]))
    elif sort == 'score':
        rows = sorted(rows, key=lambda r: -r[3])
    cols = ''.join(f'<col style="width:{w + (6 if i < len(COLS) - 1 else 0)}px">' for i, w in enumerate(COLS))
    heads = [('Raider', ''), ('Score', 'n'), ('Att.', 'n'), ('Sign-up', ''), ('Gap', 'n'), ('Ench.', 'n'),
             ('iLvl', 'n')]
    head = ''.join(f'<th class="{c}{" on" if h == "Score" and sort == "score" else ""}">{h}</th>' for h, c in heads)
    body = ''
    for n, cls, spec, score, att, sign, live, gap, ench, ilvl, *_ in rows:
        state = {'live': 'live', 'off': 'heard'}.get(live)
        score_cell = (f'<td class="n" style="color:{TIER_HEX[tier_of(score)]}">{score}</td>' if pack
                      else '<td class="n m">-</td>')
        att_cell = f'<td class="n">{att}%</td>' if pack else '<td class="n m">-</td>'
        if not signups:
            sign_cell = '<td class="m">-</td>'
        else:
            sign_cell = f'<td class="{SIGN.get(sign, "m")}">{sign or "Not signed"}</td>'
        dash = '<span class="m">-</span>'
        gap_cell = dash if gap is None else ('<span class="m">Best</span>' if gap == 0 else f'{gap:.1f}')
        ench_cell = dash if ench is None else (f'<span class="wa">{ench}</span>' if ench else '<span class="m">0</span>')
        body += (f'<tr><td>{raider(n, state, cls)}</td>{score_cell}{att_cell}{sign_cell}<td class="n">{gap_cell}</td>'
                 f'<td class="n">{ench_cell}</td><td class="n m">{ilvl or "-"}</td></tr>')
    count = len(rows) if shown is None else shown
    links = []
    if filt != 'all':
        links.append('<span class="link">Show all</span>')
    if signups and filt != 'unsigned':
        links.append(f'<span class="link">Not signed up ({unsigned})</span>')
    links.append(f'<span class="link">Raiders online ({online})</span>')
    foot = f'<span class="m">{count} of {raiders} shown</span> · ' + ' · '.join(links)
    legend = (f'<span class="dot" style="background:{ADDON["success"]};margin:0 4px 0 0"></span>addon live'
              f'<span class="dot" style="background:{ADDON["muted"]};margin:0 4px 0 8px"></span>heard earlier'
              ' · Gap in DPS')
    note = 'by Score' if sort == 'score' else ('by sign-up, then Score' if signups else 'sign-ups come with the companion')
    return section('Roster', f'<table class="ros"><colgroup>{cols}</colgroup><tr>{head}</tr>{body}</table>'
                   f'<div class="m" style="margin-top:4px">{foot}</div><div class="m" style="margin-top:2px">{legend}</div>',
                   note)


def ready(rows, note='Onyxia\'s Lair · 20:00'):
    html = ''
    for label, count, names_html, warn in rows:
        html += (f'<div class="rd"><span class="l">{label}</span><span class="c {"wa" if warn else "s"}">{count}</span>'
                 f'<span class="v">{names_html}</span></div>')
    return section('Tonight\'s readiness', html, note)


def names(*ns, tail=''):
    return ', '.join(name(n) for n in ns) + (f' {tail}' if tail else '')


READY = [
    ('Addon live', '10 of 12', names('Orwen', 'Nessaly', tail='without'), False),
    ('Old addon', '1', names('Tamsin'), True),
    ('Raid build', '9 of 10', names('Dravik', tail='off'), False),
    ('Enchants', '5', f'{name("Kethra")} 1, {name("Velisandre")} 2 <span class="m">+1</span>', True),
    ('Not attuned', '1', names('Saffa'), True),
    ('No flask', '2', names('Dravik', 'Saffa'), True),
    ('No fire prot.', '3', names('Dravik', 'Velisandre', 'Saffa'), True),
    ('Not covered', '1', '<span class="wa">Faerie Fire</span>', True),
    ('Jobs', '9 of 10', '<span class="wa">Off tank open</span>', False),
    ('Ready check', '2', f'{name("Kethra")}, {name("Saffa")}', True),
]


def signups(site=True):
    line = ('<span class="ok">11 In</span> · <span class="cg">1 Late</span> · 2 Bench · <span class="wa">1 Out</span>'
            ' · 16 not signed up')
    roles = ' · '.join(f'<span class="{"ok" if c >= w else "wa"}">{n} {c} of {w}</span>'
                       for n, c, w in (('Tanks', 2, hub.WANTED['Tank']), ('Healers', 4, hub.WANTED['Healer']),
                                       ('Melee', 2, hub.WANTED['Melee']), ('Ranged', 4, hub.WANTED['Ranged'])))
    queue = (f'<div class="row" style="margin-top:6px"><span class="m" style="flex:1">'
             f'{"3 approvals on the site" if site else ""}</span><span class="btn">Remind 16 in chat</span></div>')
    return section('Sign-ups', f'<div>{line}</div><div>{roles}</div>{queue}', '12 of 40 coming')


def push(confirm=False, sent=True, build=True):
    if not build:
        return section('Push a build', '<div class="row"><div style="flex:1" class="m">Load a build on the Builds tab'
                       ' to send it.</div><span class="btn off">Send to guild</span></div>')
    if confirm:
        return section('Push a build', '<div>Send Beast Mastery raid build to the guild? Guildmates with Forever Sixty'
                       ' get it as a build they can choose to follow.</div><div class="row" style="margin-top:6px;gap:4px">'
                       '<span class="btn" style="min-width:64px;justify-content:center">Send</span>'
                       '<span class="btn" style="min-width:64px;justify-content:center">Cancel</span></div>')
    sub = 'sent Tue · 9 follow it' if sent else 'not sent yet'
    return section('Push a build', f'<div class="row" style="align-items:flex-start">'
                   f'{icon("ability_hunter_beasttaming", 20)}<div style="flex:1;min-width:0"><div class="one">'
                   f'Beast Mastery raid build</div><div class="m one">{sub}</div></div>'
                   '<span class="btn">Send to guild</span></div>')


def loot_line():
    return section('Master loot', '<div class="m">Opens here when you loot as master looter. Tue: 7 items, 6 to the'
                   ' biggest gain.</div>')


def win(left, right, figures=(GUILD, 'Officer'), foot=('Esc to close', 'Click a name to whisper · Shift-click to invite'),
        dots=('Gear',)):
    return window('Officer', ' · '.join(('Level 60 Hunter',) + tuple(figures)), left, right, foot, dots=dots)


def tip(lines, left, top):
    body = ''.join(f'<div>{line}</div>' for line in lines[1:])
    return (f'<div class="tip" style="left:{left}px;top:{top}px;width:auto"><div class="t">{lines[0]}</div>'
            f'<div style="color:#b8b8b8">{body}</div></div>')


def raid_day(sort='signup'):
    return win(roster(ROSTER, sort=sort), ready(READY) + signups() + push() + loot_line())


def session(armed=True, officer=True, quiet=False, waiting=False):
    item = 17069
    quiet = quiet or waiting
    if quiet:
        rows = [('Ilvar', 'hunter', None, 12651, 1, 80, False, 'live'),
                ('Velisandre', 'hunter', None, 13023, 0, 88, False, 'live')]
    else:
        rows = [('Ilvar', 'hunter', '+26.0', 12651, 1, 80, False, 'live'),
                ('Velisandre', 'hunter', '+21.1', 13023, 0, 88, True, 'live')]
    cols = '<colgroup><col style="width:92px"><col style="width:44px"><col style="width:74px"><col style="width:34px">' \
           '<col style="width:40px"><col style="width:104px"></colgroup>'
    head = ('<tr><th>Raider</th><th class="n">DPS</th><th style="padding-left:4px">Wears</th><th class="n">Won</th><th class="n">Att.</th>'
            '<th></th></tr>')
    body = ''
    for i, (n, cls, gain, worn, won, att, listed, live) in enumerate(rows):
        is_armed = armed and i == 0 and not quiet
        btn = ('<span class="btn sel">Give to Ilvar?</span>' if is_armed
               else '<span class="btn" style="min-width:40px;justify-content:center">Give</span>')
        if waiting and n == 'Ilvar':
            btn = '<span class="btn off" style="min-width:40px;justify-content:center">Give</span>'
        flag = '<span class="g" style="margin-right:6px">on list</span>' if listed else ''
        gain_cell = f'<td class="n ok">{gain}</td>' if gain else '<td class="n m">-</td>'
        wears = (f'<td style="padding-left:4px"><span class="row" style="gap:4px">{item_icon(worn, 14)}'
                 f'<span class="one">{item_name(worn)}</span></span></td>')
        tail = f'<td class="n">{won}</td><td class="n m">{att}%</td><td>{"" if is_armed else flag}{btn}</td>'
        body += f'<tr><td>{raider(n, live, cls)}</td>{gain_cell}{wears}{tail}</tr>'
    summary = ('Waiting for answers.' if waiting else 'No answers. Give still works.' if quiet else
               f'2 of 2 answered · {name("Torvald")}, {name("Kethra")} no gain · {name("Orwen")} no addon')
    note = '<span class="link">Roster</span> · item 2 of 3' if officer else 'item 2 of 3'
    left = section('Master loot', f'<div class="row" style="margin-bottom:8px">{item_icon(item, 32)}'
                   f'<div style="flex:1;min-width:0"><div class="one">{item_name(item)}</div>'
                   f'<div class="m">Ranged · Magmadar</div></div></div><table class="lt">{cols}{head}{body}</table>'
                   f'<div class="m" style="margin-top:6px">{summary}</div>', note)
    sess = [(16847, f'to {name("Ilvar")} <span class="ok">+14.8 DPS</span>'), (17069, 'open'),
            (18823, '<span class="m">no answers yet</span>')]
    items = ''.join(f'<div class="up">{item_icon(i, 20)}<div class="tx"><div class="one">{item_name(i)}</div>'
                    f'<div>{state}</div></div></div>' for i, state in sess)
    right = section('Magmadar', items, '3 items')
    hint = 'Click Give twice: master loot hands it out'
    if armed and not quiet:
        hint = 'Click Give again: Ilvar gets Striker\'s Mark'
    if waiting:
        hint = 'Master loot cannot give it to them (out of range or not eligible).'
    foot = ('Esc for the roster' if officer else 'Esc to close', hint)
    figures = (GUILD, 'Master looter')
    tabs = hub.TABS if officer else [t for t in hub.TABS if t != 'Officer'] + ['Officer']
    return window('Officer', ' · '.join(('Level 60 Hunter',) + figures), left, right, foot, tabs=tabs, dots=('Officer',))


def no_companion():
    rows = [r for r in ROSTER if r[5] != '' or True]
    online = [('Addon live', '8 of 11', names('Orwen', 'Nessaly', 'Quenna', tail='without'), False),
              ('Enchants', '5', f'{name("Kethra")} 1, {name("Velisandre")} 2 <span class="m">+1</span>', True)]
    return win(roster(rows, signups=False, online=24), ready(online, 'raiders online') + push() + loot_line())


def nobody_addon():
    rows = [(n, c, s, sc, a, 'Out' if n == 'Velisandre' else sg, 'live' if n == 'Velisandre' else None,
             g if n == 'Velisandre' else None, e, il, *rest) for n, c, s, sc, a, sg, _, g, e, il, *rest in ROSTER]
    body = ('<div class="m">Nobody coming runs the addon.</div>'
            f'<div class="rd"><span class="l">Enchants</span><span class="c wa">3</span><span class="v">'
            f'{name("Kethra")} 1, {name("Saffa")} 2</span></div>'
            '<div class="rd"><span class="l">Jobs</span><span class="c s">9 of 10</span><span class="v">'
            '<span class="wa">Off tank open</span></span></div>')
    return win(roster(rows), section('Tonight\'s readiness', body, 'Onyxia\'s Lair · 20:00') + signups() + push()
               + loot_line())


def no_pack():
    return win(roster(ROSTER, pack=False), ready(READY) + signups() + push(build=False) + loot_line())


def hover():
    html = raid_day(sort='score')
    lines = [name('Kethra'), 'Level 60 Combat', f'<span style="color:{TIER_HEX["rare"]}">Forever Score 1377 · Rare · as of'
             ' 9 Dec</span>', 'Best parse: 91st percentile on Onyxia', 'Build: Combat raid', 'No enchant: Cloak',
             'Last heard from their addon 2 min ago.']
    return html.replace('<div class="ft">', tip(lines, 250, 230) + '<div class="ft">')


STATES = {
    'l7-raid-day': ('Raid day as an officer: the roster by sign-up, Tonight\'s readiness, Sign-ups with the site\'s '
                    'queue and Remind, Push a build, the night\'s master loot line.', raid_day),
    'l7-sorted-hover': ('A Score heading click sorts the roster by Score (the heading gold); the mouse on Kethra: '
                        'the hover with the score, best parse, build, the bare slot and the report\'s age.',
                        hover),
    'l7-push-confirm': ('Send to guild clicked: the confirm in its place, Send and Cancel.',
                        lambda: win(roster(ROSTER), ready(READY) + signups() + push(confirm=True) + loot_line())),
    'l7-loot-session': ('Looting Magmadar as master looter: the session on the left, Give armed for Ilvar, the boss\'s '
                        'items on the right; Esc and the gold Roster go back.', session),
    'l7-loot-waiting': ('A drop asked a few seconds ago, nobody answered yet: the members who could wear it, gain a '
                        'muted dash, Waiting for answers; Ilvar out of range: Give muted, the reason in the footer '
                        'after a click.', lambda: session(armed=False, waiting=True)),
    'l7-loot-quiet': ('A drop nobody answered after 30 s: the members who could wear it, no gain, Give still there.',
                      lambda: session(armed=False, quiet=True)),
    'l7-sent': ('Remind clicked: the one line went to guild chat; the footer says so for 2 s.',
                lambda: win(roster(ROSTER), ready(READY) + signups() + push() + loot_line(),
                            foot=('Esc to close', 'Reminder sent to guild chat'))),
    'l7-master-looter': ('A master looter without the officer gate: the session only, no Roster link, Esc closes.',
                         lambda: session(armed=False, officer=False)),
    'l7-no-companion': ('No companion (no sign-ups by name): Sign-up shows a dash, readiness counts the raiders online, '
                        'Sign-ups left out.', no_companion),
    'l7-nobody-addon': ('Nobody coming runs the addon: readiness says so, Enchants from inspect and Jobs stay.',
                        nobody_addon),
    'l7-no-pack': ('No ratings pack: Score and Att. a muted dash; no build loaded: Send to guild muted with why.',
                   no_pack),
}


def main(names_):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        page = browser.new_context(device_scale_factor=1, viewport={'width': 900, 'height': 900}).new_page()
        for key in names_:
            caption, build = STATES[key]
            html = hub.board2(f'<b>Officer, as built</b> · {caption}', build()).replace('</style>', CSS + '</style>', 1)
            page.set_content(html)
            page.evaluate('document.fonts.ready')
            page.wait_for_timeout(150)
            page.locator('#board').screenshot(path=str(OUT / f'{key}.png'))
            print(OUT / f'{key}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(STATES))
