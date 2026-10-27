"""Mock boards for Lane G2, attendance and the raid-night report
(design/specs/2026-10-27-guild-attendance.md).

  python3 design/mocks/gen_guild_attendance.py
    writes guild-attendance-page.html        (1440, officer, done raid: band, Report, Attendance)
    writes guild-attendance-page-1024.html   (1024, the officer board again, for the same-scale side by side)
    writes guild-attendance-page-phone.html  (390, member, done raid: band, Your night, Report)
    writes guild-attendance-match.html       (1440, ended with no log, picker, clear confirm, no-show)
    writes guild-attendance-record.html      (1440, roster chip, guild line, readiness next-raid filter)
    writes guild-attendance-discord.html     (1440, report_ready embed and the copied text)

Built on gen_guild_raids (G1): same tokens, chips, band pieces and SIGNUPS, imported, never
copied. Boards go through mocklib.write into design/mocks/renders/ (ignored). Mock data,
labelled on every board: the Onyxia's Lair raid of Thu 10 Dec from G1's boards, logged as 3
wipes then a kill. Outcomes are computed from SIGNUPS and LOG below, never typed per board.
Mock rule for R-A2: Mirestalk is Felsnap's alt, so Felsnap reads "Came as Mirestalk".
"""
from gen_guild_raids import (
    CSS, DOT, GUILD, KILL, ONY, SIGNUPS, NOT_SIGNED, CLASS_ORDER, TEXT_CLASS,
    pill, viewer_chip, crest_img, tab_strip, cell, embed, message, sheet, board,
)
from mocklib import nav, nav_phone

PULLS = 4
LATE_AFTER_MIN = 15
EPIC = '#a335ee'
MUTED_HEX = '#9a9484'
# name: (first seen, minutes after the first pull; pulls present). Absent names are not in the log.
LOG = {r[0]: (0, PULLS) for r in SIGNUPS if r[4] in ('in', 'late')}
LOG.pop('Hexbramble')
LOG.pop('Felsnap')
LOG.update(dict(Mendwhisper=(22, 3), Soulgrave=(9, PULLS), Nixthrottle=(31, 2),
                Thistlewind=(0, PULLS), Mirestalk=(0, PULLS)))
CAME_AS = dict(Felsnap='Mirestalk')
EXCUSED = {'Mendwhisper'}
DEATHS = [('Pyrewisp', 'mage', 3, 1), ('Windtalon', 'hunter', 3, 0), ('Vexlash', 'rogue', 2, 0),
          ('Frostnettle', 'mage', 2, 0), ('Kraggor', 'warrior', 1, 0), ('Arcanemoor', 'mage', 1, 0),
          ('Stormtusk', 'shaman', 1, 0)]
LOOT = [('Head of Onyxia', 'Kraggor', 'warrior'), ("Vis'kag the Bloodletter", 'Vexlash', 'rogue'),
        ('Netherwind Crown', 'Pyrewisp', 'mage'), ('Ring of Binding', 'Thornhide', 'druid'),
        ('Sapphiron Drape', 'Arcanemoor', 'mage'), ('Onyxia Hide Backpack', 'Lightbrand', 'priest')]
MOVES = [('Pyrewisp', 'mage', 'DPS', 238, 291), ('Lightbrand', 'priest', 'HPS', 402, 443),
         ('Windtalon', 'hunter', 'DPS', 251, 219), ('Stormtusk', 'shaman', 'DPS', 196, 224),
         ('Vexlash', 'rogue', 'DPS', 270, 289)]
CONSUMES = [('Hexbramble', 'warlock', 'Short', None), ('Windtalon', 'hunter', 'Short', (False, True, 1)),
            ('Pyrewisp', 'mage', 'Short', (True, True, 0)), ('Kraggor', 'warrior', 'Stocked', (True, True, 3)),
            ('Lightbrand', 'priest', 'Stocked', (True, True, 2)), ('Zulmara', 'hunter', 'Not shared', (False, True, 2))]
KILLS, WIPES, DURATION, KILL_TIME = 1, 3, '1 h 12 min', '4:38'
SIGNED_STATES = ('in', 'late')


def is_late(name):
    return max(LOG[name][0], LATE_AFTER_MIN) != LATE_AFTER_MIN


def outcome(name, state):
    if name in CAME_AS:
        return 'present', CAME_AS[name]
    if name not in LOG:
        return ('no_show' if state in SIGNED_STATES else 'absent'), None
    return ('late' if state in SIGNED_STATES and is_late(name) else 'present'), None


ROWS = [(n, c, s, role, st) + outcome(n, st) for n, c, s, role, st, _note, _acct in SIGNUPS]
ROWS += [(n, c, s, None, None, 'present', None) for n, c, s in NOT_SIGNED if n in LOG and n not in CAME_AS.values()]
SIGNED = sum(1 for r in ROWS if r[4] in SIGNED_STATES)
CAME = sum(1 for r in ROWS if r[4] in SIGNED_STATES and r[5] in ('present', 'late'))
LATE = sum(1 for r in ROWS if r[5] == 'late')
MISSING = sum(1 for r in ROWS if r[5] == 'no_show')
JOINED = sum(1 for r in ROWS if r[4] not in SIGNED_STATES and r[0] in LOG)
N_DEATHS = sum(w + k for _, _, w, k in DEATHS)
EXTRA_CSS = '''<style>
.rrow{display:grid;align-items:center;gap:16px;border-bottom:1px solid #1c2230;min-height:44px}
.rad{width:16px;height:16px;border-radius:50%;border:1px solid #9a9484;display:inline-block;box-sizing:border-box}
.rad.on{border:5px solid #e5b955}
.icon{width:28px;height:28px;border-radius:4px;border:1px solid #a335ee;background:#131824;flex-shrink:0}
.blk{display:flex;flex-direction:column;gap:12px}
</style>'''


def name_html(name, cls, size=14):
    return f'<span style="font-size:{size}px;font-weight:600;color:{TEXT_CLASS[cls]}">{name}</span>'


def who(name, cls, spec='', crest=24):
    sub_line = f'<span style="font-size:12px;color:#9a9484">{spec}</span>' if spec else ''
    return (f'<span style="display:flex;align-items:center;gap:10px;min-width:0">{crest_img(cls, crest)}'
            f'<span style="display:flex;flex-direction:column">{name_html(name, cls)}{sub_line}</span></span>')


def sub(text, count=''):
    c = f' <span class="mono" style="font-size:13px;color:#9a9484">{count}</span>' if count != '' else ''
    return f'<span class="sub">{text}{c}</span>'


def muted(text):
    return f'<span style="font-size:13px;color:#9a9484">{text}</span>'


def night_line(phone):
    width = 120 if phone else 160
    fill = round(100 * CAME / SIGNED)
    tail = DOT.join([f'{KILLS} kill', f'{WIPES} wipes', f'{N_DEATHS} deaths', DURATION])
    return (f'<p style="margin:0;display:flex;align-items:center;gap:10px;flex-wrap:wrap">'
            f'<span class="bar" style="width:{width}px"><i style="width:{fill}%"></i></span>'
            f'<span class="mono" style="font-size:15px;color:#f2eee4">{CAME} of {SIGNED} came</span>'
            f'<span style="font-size:13px;color:#9a9484">{DOT}{tail}</span></p>')


def match_line(officer, confirm=False):
    btn = ' <a class="mlink" style="color:#e9e4d8;margin-left:8px">Not this log</a>' if officer and not confirm else ''
    gt = ' (10:42 pm guild time)' if officer else ''
    clear = ''
    if confirm:
        clear = ('<div style="display:flex;flex-wrap:wrap;gap:12px;align-items:center;margin-top:6px">'
                 "<span style=\"font-size:14px\">Clear the match? This raid's attendance is removed until another log is matched.</span>"
                 '<span class="btns" style="border-color:#f0736b;color:#f0736b">Clear the match</span>'
                 '<span class="btns" style="outline:2px solid #e5b955;outline-offset:2px">Keep it</span></div>')
    return (f'<div style="font-size:13px"><p style="margin:0">Logged as a raid night. '
            f'<a style="color:#e5b955;font-weight:700">Open the night</a></p>'
            f'<p style="margin:0;color:#9a9484;display:flex;align-items:center;flex-wrap:wrap">'
            f'Log uploaded Thu 10 Dec, 11:42 pm{gt} by Kraggor{DOT}matched by time{btn}</p>{clear}</div>')


def ended_line(officer, phone, late=False):
    link = '<a style="color:#9a9484;text-decoration:underline">Upload the log</a>'
    tail = f' {link}, then match it here.' if late else ' We match the log when it is uploaded.'
    btn = '<span class="btns">Match a log</span>' if officer else ''
    direction = 'column;align-items:flex-start' if phone else 'row;align-items:center'
    return (f'<div style="display:flex;flex-direction:{direction};gap:12px;flex-wrap:wrap">'
            f'<p style="margin:0;font-size:14px">Ended · no log yet<span style="color:#9a9484">{tail}</span></p>{btn}</div>')


def band(role, phone, status='done', confirm=False, late=False):
    officer = role == 'officer'
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    chip = pill('Done', 'kill') if status == 'done' else pill('Ended', 'sample')
    gt = f'{DOT}<span style="color:#9a9484">8:00 pm CST guild time</span>' if officer else ''
    times = '9:00 pm to 12:00 am' if officer else '8:00 pm to 11:00 pm'
    when = f'<p style="margin:0;font-size:15px">Thursday 10 December{DOT}<span class="mono">{times}</span>{gt}</p>'
    if status == 'done':
        body = night_line(phone) + match_line(officer, confirm)
    else:
        body = (f'<p style="margin:0"><span class="mono" style="font-size:15px;color:#9a9484">23 of 40 signed up</span></p>'
                + ended_line(officer, phone, late))
    head = f'<div style="display:flex;align-items:center;gap:10px"><span class="label" style="color:#e5b955">Raid</span>{chip}</div>'
    nav_html = nav_phone(True, None) if phone else nav('', True, None)
    return (f'<div class="band" style="display:flex;flex-direction:column">{nav_html}'
            f'<div style="display:flex;flex-direction:column;gap:10px;padding:{pad};{maxw}">'
            f'<a class="mlink">‹ {GUILD}{DOT}Raids</a>{head}'
            f'<h1 class="display" style="margin:0;font-size:22px;color:#f2eee4">{ONY}</h1>{when}{body}</div>'
            f'<div style="margin-top:18px;{maxw}">{tab_strip("Raids", role, phone)}</div></div>')


def kill_order():
    row = (f'<div class="rrow" style="grid-template-columns:20px 1fr auto">'
           f'<span class="mono" style="color:#9a9484;font-size:13px">1</span>'
           f'<span style="font-size:14px;font-weight:600;color:#f2eee4">Onyxia</span><span style="font-size:13px">'
           f'<span class="mono" style="font-size:14px">{KILL_TIME}</span> '
           f'<span style="color:#9a9484">after {WIPES} wipes{DOT}1 death</span></span></div>')
    return f'<div class="blk">{sub("Kill order", KILLS)}<div>{row}</div></div>'


def death_row(n, c, w, k, top):
    width = round(160 * (w + k) / top)
    return (f'<div class="rrow" style="grid-template-columns:minmax(0,1fr) 160px 24px 150px;min-height:40px">'
            f'<span style="display:flex;gap:8px;align-items:center">{crest_img(c, 20)}{name_html(n, c, 13)}</span>'
            f'<span class="bar" style="width:{width}px"><i style="background:#9a9484;width:100%"></i></span>'
            f'<span class="mono" style="font-size:13px;color:#f2eee4">{w + k}</span>'
            f'<span style="font-size:12px;color:#9a9484">{w} on wipes{DOT}{k} on kills</span></div>')


def deaths():
    top = max(w + k for _, _, w, k in DEATHS)
    rows = ''.join(death_row(n, c, w, k, top) for n, c, w, k in DEATHS)
    return f'<div class="blk">{sub("Deaths", N_DEATHS)}<div>{rows}</div></div>'


def loot_row(item, n, c):
    return (f'<div class="rrow" style="grid-template-columns:minmax(0,1.4fr) 16px minmax(0,1fr) auto">'
            f'<span style="display:flex;gap:10px;align-items:center"><span class="icon"></span>'
            f'<span style="font-size:14px;color:{EPIC}">{item}</span></span><span style="color:#9a9484">→</span>'
            f'<span style="display:flex;gap:8px;align-items:center">{crest_img(c, 20)}{name_html(n, c, 13)}</span>'
            f'<span style="font-size:12px;color:#9a9484">by Bow Jackzon</span></div>')


def loot():
    rows = ''.join(loot_row(item, n, c) for item, n, c in LOOT)
    return (f'<div class="blk">{sub("Loot", len(LOOT))}<span class="label">Onyxia <span class="mono">{len(LOOT)}</span></span>'
            f'<div>{rows}</div><p class="ann">Icon squares stand in for the game icons ItemHover draws.</p></div>')


def delta(before, after):
    rose = max(before, after) == after
    colour = KILL if rose else MUTED_HEX
    sign = '+' if rose else '−'
    return f'<span class="mono" style="font-size:13px;color:{colour}">{sign}{abs(after - before)}</span>'


def move_row(n, c, m, b, a):
    return (f'<div class="rrow" style="grid-template-columns:minmax(0,1fr) auto auto 48px">'
            f'<span style="display:flex;gap:8px;align-items:center">{crest_img(c, 20)}{name_html(n, c, 13)}</span>'
            f'{muted("on Onyxia")}<span class="mono" style="font-size:14px">{m} {b} → {a}</span>{delta(b, a)}</div>')


def moves():
    rows = ''.join(move_row(*m) for m in MOVES)
    return f'<div class="blk">{sub("Biggest changes since Thu 3 Dec")}<div>{rows}</div></div>'


def used_cell(used):
    if used is None:
        return muted('No flask · No food · No potions')
    flask, food, pots = used
    bits = [pill('Flask', 'sample') if flask else muted('No flask'),
            pill('Food', 'sample') if food else muted('No food'),
            pill(f'Potions {pots}', 'sample') if pots else muted('No potions')]
    return '<span style="display:flex;gap:8px;flex-wrap:wrap">' + ''.join(bits) + '</span>'


def before_cell(state):
    kind = dict(Short='wipe', Stocked='kill').get(state)
    return pill(state, kind) if kind else muted(state)


CONSUME_GRID = 'grid-template-columns:minmax(180px,1fr) 120px minmax(0,1.5fr)'


def consumables():
    head = (f'<div class="rrow" style="{CONSUME_GRID};min-height:28px">'
            f'<span></span><span class="label">Before</span><span class="label">In the log</span></div>')
    rows = ''.join(f'<div class="rrow" style="{CONSUME_GRID}">{who(n, c, crest=20)}{before_cell(b)}{used_cell(u)}</div>'
                   for n, c, b, u in CONSUMES)
    return f'<div class="blk">{sub("Consumables", SIGNED)}<div>{head}{rows}</div></div>'


def column(*blocks):
    return '<div style="display:flex;flex-direction:column;gap:28px">' + ''.join(blocks) + '</div>'


def report(officer, phone):
    btns = ''
    if officer:
        btns = ('<span style="display:flex;gap:12px;flex-wrap:wrap"><span class="btns">Copy for Discord</span>'
                '<span class="btns">Post to Discord</span></span>')
    head = (f'<div style="display:flex;justify-content:space-between;align-items:center;gap:12px;flex-wrap:wrap">'
            f'<span class="h2">Report</span>{btns}</div>')
    if not officer:
        body = column(kill_order(), loot())
    elif phone:
        body = column(kill_order(), loot(), deaths(), moves(), consumables())
    else:
        body = (f'<div style="display:grid;grid-template-columns:minmax(0,1fr) minmax(0,1fr);gap:32px">'
                f'{column(kill_order(), deaths())}{column(loot(), moves())}</div>{consumables()}')
    return f'<section style="display:flex;flex-direction:column;gap:20px">{head}{body}</section>'


def outcome_text(r):
    name, out, came_as = r[0], r[5], r[6]
    if came_as:
        return f'Came as {came_as}'
    if out in ('no_show', 'absent'):
        return muted('Not in the log')
    if out == 'late':
        return f'Late, first seen {LOG[name][0]} min after the first pull'
    return 'Present'


def pulls(name):
    got = LOG[CAME_AS.get(name, name)][1] if CAME_AS.get(name, name) in LOG else 0
    colour = MUTED_HEX if got == PULLS else '#e9e4d8'
    return f'<span class="mono" style="font-size:13px;color:{colour}">{got}/{PULLS}</span>'


def signed_chip(state, role):
    if state in SIGNED_STATES:
        return viewer_chip(f'{state.capitalize()}{DOT}{role}')
    return viewer_chip(state.capitalize() if state else None)


def excused_box(r):
    if r[5] not in ('no_show', 'late'):
        return '<span></span>'
    box = 'cb' if r[0] in EXCUSED else 'cb no'
    return f'<span style="display:flex;gap:8px;align-items:center;min-height:44px;font-size:13px"><span class="{box}"></span>Excused</span>'


def full_row(r, phone):
    name, cls, spec, role, state = r[:5]
    chip = signed_chip(state, role)
    if phone:
        return (f'<div style="display:flex;flex-direction:column;gap:6px;padding:10px 0;border-bottom:1px solid #1c2230">'
                f'<div style="display:flex;justify-content:space-between;align-items:center">{who(name, cls, spec)}{chip}</div>'
                f'<span style="font-size:13px">{outcome_text(r)}{DOT}{pulls(name)}</span>{excused_box(r)}</div>')
    return (f'<div class="rrow" style="grid-template-columns:minmax(200px,1.4fr) 132px minmax(180px,1fr) 120px 112px;min-height:52px">'
            f'{who(name, cls, spec)}<span>{chip}</span><span style="font-size:13px">{outcome_text(r)}</span>{pulls(name)}{excused_box(r)}</div>')


def came_as_mark(r):
    return f'<span style="font-size:12px;color:#9a9484">as {r[6]}</span>' if r[6] else ''


def compact(rows, phone):
    cols = 1 if phone else 3
    height = 44 if phone else 36
    cells = ''.join(f'<div class="rrow" style="grid-template-columns:1fr auto;min-height:{height}px">'
                    f'<span style="display:flex;align-items:center;gap:8px">{who(r[0], r[1])}{came_as_mark(r)}</span>{pulls(r[0])}</div>'
                    for r in rows)
    return f'<div style="display:grid;grid-template-columns:repeat({cols},minmax(0,1fr));column-gap:24px">{cells}</div>'


ROLE_ORDER = dict(Tank=0, Healer=1, DPS=2)


def by_role(rows):
    return sorted(rows, key=lambda r: (ROLE_ORDER.get(r[3], 3), CLASS_ORDER.index(r[1]), r[0]))


def attendance_groups():
    return [
        ('Signed up, not in the log', [r for r in ROWS if r[5] == 'no_show'], True),
        ('Late', [r for r in ROWS if r[5] == 'late'], True),
        ('Joined from the bench or without a sign-up', [r for r in ROWS if r[4] not in SIGNED_STATES and r[0] in LOG], True),
        ('Present', [r for r in ROWS if r[4] in SIGNED_STATES and r[5] == 'present'], False),
        ('Bench and out', [r for r in ROWS if r[5] == 'absent'], False),
    ]


def attendance(phone):
    summary = DOT.join([f'{CAME} of {SIGNED} came', f'{LATE} late', f'{MISSING} not in the log',
                        f'{JOINED} more joined', f'{len(EXCUSED)} excused'])
    parts = [f'<span class="h2">Attendance</span><p style="margin:0;font-size:14px">{summary}</p>']
    for title, rows, full in attendance_groups():
        if not rows:
            continue
        ordered = by_role(rows)
        body = ''.join(full_row(r, phone) for r in ordered) if full else compact(ordered, phone)
        parts.append(f'<span class="label" style="margin-top:8px">{title} <span class="mono">{len(rows)}</span></span><div>{body}</div>')
    return '<section style="display:flex;flex-direction:column;gap:10px">' + ''.join(parts) + '</section>'


def your_night(no_show=False):
    if no_show:
        name, cls, spec = 'Hexbramble', 'warlock', 'Affliction'
        sentence = 'Signed In as DPS. You are not in this log. If you were there under another name, tell an officer.'
        facts = ''
    else:
        name, cls, spec = 'Zulmara', 'hunter', 'Marksmanship'
        sentence = f'Signed In as DPS. Present from the first pull, {PULLS} of {PULLS} pulls.'
        facts = f'0 deaths{DOT}DPS 262 on Onyxia, <span style="color:{KILL}">up 17</span> since Thu 3 Dec'
    row = (f'<div style="display:flex;flex-wrap:wrap;gap:12px 16px;align-items:center;padding:12px 0">{who(name, cls, spec, 32)}'
           f'<span style="display:flex;flex-direction:column;flex:2;min-width:240px"><span style="font-size:14px">{sentence}</span>'
           f'{muted(facts)}</span>{muted("No loot this raid.")}</div>')
    return f'<div class="card" style="padding:16px">{sub("Your night")}{row}</div>'


CANDIDATES = [(ONY, 'Kraggor', 'Thu 10 Dec, 11:42 pm', '4 pulls · 1 kill', '1 h 12 min', 'closest'),
              (ONY, 'Lightbrand', 'Fri 11 Dec, 12:20 am', '2 pulls · 0 kills', '24 min', ''),
              (ONY, 'Kraggor', 'Thu 3 Dec, 11:05 pm', '6 pulls · 1 kill', '1 h 40 min', 'taken')]
PICK_GRID = 'grid-template-columns:24px minmax(0,1fr) 200px 140px 120px;min-height:56px'


def picker_row(zone, by, when, pk, dur, flag):
    taken = flag == 'taken'
    line2 = f'Matched to <a style="text-decoration:underline">{ONY}</a> · Thu 3 Dec. Clear it there first.' if taken else f'by {by}'
    style = 'opacity:.5' if taken else ''
    chip = ' ' + pill('Closest', 'site') if flag == 'closest' else ''
    radio = 'rad on' if flag == 'closest' else 'rad'
    return (f'<div class="rrow" style="{PICK_GRID};{style}"><span class="{radio}"></span>'
            f'<span style="display:flex;flex-direction:column"><span style="font-size:14px;font-weight:600;color:#f2eee4">{zone}{chip}</span>'
            f'<span style="font-size:12px;color:#9a9484">{line2}</span></span><span class="mono" style="font-size:13px">{when}</span>'
            f'<span style="font-size:13px">{pk}</span><span class="mono" style="font-size:13px">{dur}</span></div>')


def picker(empty=False, unknown=False):
    if unknown:
        helper = "Logs this guild uploaded in the last 14 days. Barrow Deeps's log name is not known yet, so every zone shows."
    else:
        helper = f'{ONY} logs this guild uploaded in the last 14 days.'
    rows = ''.join(picker_row(*c) for c in CANDIDATES)
    primary = '<span class="btnp">Match this log</span>'
    if empty:
        rows = (f'<p style="margin:0;font-size:14px">No {ONY} logs uploaded in the last 14 days. '
                f'<a style="text-decoration:underline">Upload the log</a>, then match it here.</p>')
        primary = ''
    return (f'<div class="card" style="padding:20px;display:flex;flex-direction:column;gap:12px">{sub("Match a log")}'
            f'<p style="margin:0;font-size:13px;color:#9a9484">{helper}</p><div>{rows}</div>'
            f'<div style="display:flex;gap:16px;align-items:center">{primary}<a class="mlink" style="color:#e9e4d8">Cancel</a></div></div>')


def stat_chip(text):
    return f'<span class="pill" style="color:#e9e4d8;border-color:#262e40">{text}</span>'


def filter_btn(text, on=False):
    look = 'border-color:#e5b955;color:#e5b955;background:#e5b95514' if on else 'border-color:#262e40'
    return (f'<span style="display:inline-flex;align-items:center;height:32px;padding:0 12px;border:1px solid;'
            f'border-radius:4px;font-size:12px;font-weight:600;{look}">{text}</span>')


def record_board():
    footnote = "Came: raids attended of raids signed up for or attended, over this guild's last 8 and 30 logged raids."
    roster = (f'<div style="display:flex;align-items:center;gap:12px;flex-wrap:wrap">{who("Zulmara", "hunter", "Marksmanship Hunter in game")}'
              f'{stat_chip("Item level 68")}{stat_chip("Came 7 of 8 · 26 of 30")}</div>'
              f'<p style="margin:12px 0 0;font-size:12px;color:#9a9484">{footnote}</p>')
    old = f'<div style="display:flex;gap:12px;align-items:center">{who("Zulmara", "hunter")}{stat_chip("5 of 8 nights")}</div>'
    line = (f'<span style="font-size:13px;color:#9a9484"><a style="color:#e9e4d8">{GUILD}</a> · Member '
            f'<span style="color:#f2eee4">Verified</span> · came <span class="mono">7</span> of <span class="mono">8</span> raids</span>')
    row = '<div class="rrow" style="grid-template-columns:minmax(0,1fr) auto">'
    ready = (f'<div style="display:flex;gap:8px">{filter_btn("Next raid · Thu 10 Dec", True)}{filter_btn("Role: All")}{filter_btn("Class: All")}</div>'
             f'<p style="margin:8px 0 12px;font-size:13px;color:#9a9484">Showing the {SIGNED} signed In or Late for '
             f'<a style="text-decoration:underline">{ONY}</a>, Thu 10 Dec, 8:00 pm.</p>'
             f'{row}{who("Windtalon", "hunter", "Marksmanship")}{pill("Synced 9 days before the raid", "wipe")}</div>'
             f'{row}{who("Soulgrave", "warlock", "Destruction")}{muted("No export yet")}</div>')
    cells = [cell('Roster row, raid mode, with the footnote under the roster', roster, width=660),
             cell('Roster row, report mode (unchanged)', old, width=660),
             cell('Character guild line, raid mode (own characters)', line, width=660),
             cell('Readiness, next-raid filter on: a stale export, no export', ready, width=660)]
    return '<div style="display:flex;flex-wrap:wrap;gap:24px">' + ''.join(cells) + '</div>'


def match_board():
    cells = [cell('Ended, no log yet: officer, inside the 6 h window', band('officer', False, 'ended'), width=1344),
             cell('Ended, no log yet: member at 390, after the window', band('member', True, 'ended', late=True), width=390),
             cell('Match picker: Closest preselected, one log matched elsewhere', picker(), width=900),
             cell('Picker, empty', picker(empty=True), width=660),
             cell('Picker, log zone not known yet', picker(unknown=True), width=900),
             cell('Clear confirm (focus on Keep it)', match_line(True, confirm=True), width=900),
             cell('Your night, no-show variant: never the word, never red', your_night(no_show=True), width=900)]
    return '<div style="display:flex;flex-wrap:wrap;gap:24px">' + ''.join(cells) + '</div>'


def copied_text():
    loot_line = 'Loot: Onyxia: ' + '; '.join(f'{item} → {n}' for item, n, _c in LOOT)
    lines = [f'{ONY} · Thu 10 Dec · {GUILD}', f'{KILLS} kill, {WIPES} wipes, {DURATION}. {CAME} of {SIGNED} came.',
             f'Kills: Onyxia {KILL_TIME} (after {WIPES} wipes)', 'Most deaths: Pyrewisp 4, Windtalon 3, Vexlash 2',
             loot_line, 'Biggest changes: Pyrewisp DPS +53, Lightbrand HPS +41, Stormtusk DPS +28',
             'Report: https://foreversixty.gg/guild/us/forever/olympus-xxvii/raids/r1#report']
    return ('<div class="field" style="height:auto;display:block;padding:10px 12px;font-family:var(--font-mono);font-size:12px;line-height:1.6">'
            + '<br>'.join(lines) + '</div>')


def discord_board():
    fields = [('Kills', f'Onyxia {KILL_TIME}'), ('Wipes', str(WIPES)), ('Came', f'{CAME} of {SIGNED}'), ('Loot', f'{len(LOOT)} items')]
    title = f'Report: {ONY} · Thursday 10 December'
    desc = 'Kill times, loot and the biggest changes on Forever Sixty.'
    desk = message(embed(title, fields, desc, '#7fd48a', 524, False), 600)
    phone = message(embed(title, fields, desc, '#7fd48a', 278, True), 354)
    pair = f'<div style="display:flex;gap:24px;align-items:flex-start">{desk}{phone}</div>'
    return (cell('report_ready (colour 8377482), desktop and phone clients', pair, width=1000)
            + cell('Copy for Discord, the text', copied_text(), 'No absentee names (R-A3).', width=1000))


def page_body(role, phone):
    if role == 'officer':
        return report(True, phone) + attendance(phone)
    return your_night() + report(False, phone)


if __name__ == '__main__':
    print(board('guild-attendance-page', EXTRA_CSS + band('officer', False), page_body('officer', False), 1440, False))
    print(board('guild-attendance-page-1024', EXTRA_CSS + band('officer', False), page_body('officer', False), 1024, False))
    print(board('guild-attendance-page-phone', EXTRA_CSS + band('member', True), page_body('member', True), 390, True))
    print(sheet('guild-attendance-match', 'Attendance: match states', EXTRA_CSS + match_board()))
    print(sheet('guild-attendance-record', 'Attendance: the member record and readiness', EXTRA_CSS + record_board()))
    print(sheet('guild-attendance-discord', 'Attendance: report_ready and the copied text', EXTRA_CSS + discord_board()))
