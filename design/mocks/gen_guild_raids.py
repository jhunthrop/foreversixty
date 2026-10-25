"""Mock boards for Lane G1, raids and sign-ups (design/specs/2026-10-26-guild-raids.md).

  python3 design/mocks/gen_guild_raids.py
    -> guild-raids-tab.html          (1440, officer viewing from Eastern time, Raids tab)
    -> guild-raids-tab-phone.html    (390, member in the guild's zone, Raids tab + band line)
    -> guild-raids-page.html         (1440, officer, raid page: band, strip, roster, notes)
    -> guild-raids-page-phone.html   (390, member, raid page)
    -> guild-raids-form.html         (1440, officer, the form open on the Raids tab)
    -> guild-raids-form-phone.html   (390, officer, the form)
    -> guild-raids-states.html       (1440, every other state as labelled cells)
    -> guild-raids-discord.html      (1440, the three embeds, desktop and phone widths)

Boards are written by mocklib.write into design/mocks/renders/, the ignored folder
gen_guild.py also writes to; render.py screenshots them (phone boards at 390).

Mock data, labelled as such on every board. The tier is Forever's real first tier
(data/curated/loot/forever-raid-phases.json: Barrow Deeps, Hyjal Summit, Onyxia's Lair, open
9 December 2026). Onyxia's Lair is 40 players; the other two have no sourced size, so their
counts carry no denominator (spec R6). "Now" is Tuesday 8 December 2026, 6:10 pm guild time
(America/Chicago, CST). The officer views from America/New_York, one hour ahead, so the guild
time label shows; the member views from the guild's own zone, so it does not. Every count is
computed once from SIGNUPS below, never typed per board.
"""
from mocklib import (
    GOLD, MUTED, TEXT, BORDER, SOFT, RAISED, CLASS_COLOR, SEAL,
    crest, nav, nav_phone, footer, page, write,
)

GUILD = 'Olympus XXVII'
CARD_TOP = '#131824'
KILL = '#7fd48a'
WIPE = '#f0736b'
WARM = '#3a3326'
TEXT_BODY = '#e9e4d8'
TEXT_CLASS = dict(CLASS_COLOR, shaman='#3f8fe0')
CLASS_ORDER = ['warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid']

# name, class, spec, role, state, note, account
SIGNUPS = [
    ('Kraggor', 'warrior', 'Protection', 'Tank', 'in', '', 'kraggor'),
    ('Sunderfel', 'warrior', 'Protection', 'Tank', 'in', '', 'sunderfel'),
    ('Thornhide', 'druid', 'Feral', 'Tank', 'in', 'Bear. Can off-tank the whelps.', 'thornhide'),
    ('Lightbrand', 'priest', 'Holy', 'Healer', 'in', '', 'lightbrand'),
    ('Brightvow', 'priest', 'Holy', 'Healer', 'in', '', 'brightvow'),
    ('Mendwhisper', 'priest', 'Holy', 'Healer', 'late', '20 min late, work.', 'mendwhisper'),
    ('Rootgall', 'shaman', 'Restoration', 'Healer', 'in', '', 'rootgall'),
    ('Tidecaller', 'shaman', 'Restoration', 'Healer', 'in', '', 'tidecaller'),
    ('Earthhoof', 'druid', 'Restoration', 'Healer', 'in', '', 'earthhoof'),
    ('Sagebloom', 'druid', 'Restoration', 'Healer', 'in', '', 'sagebloom'),
    ('Obnoxious Yell', 'warrior', 'Fury', 'DPS', 'in', '', 'yell'),
    ('Windtalon', 'hunter', 'Marksmanship', 'DPS', 'in', '', 'windtalon'),
    ('Zulmara', 'hunter', 'Marksmanship', 'DPS', 'in', 'Bringing Tranquilizing Shot.', 'zulmara'),
    ('Felsnap', 'hunter', 'Beast Mastery', 'DPS', 'in', '', 'felsnap'),
    ('Vexlash', 'rogue', 'Combat', 'DPS', 'in', '', 'vexlash'),
    ('Shadowquill', 'rogue', 'Assassination', 'DPS', 'in', '', 'shadowquill'),
    ('Stormtusk', 'shaman', 'Elemental', 'DPS', 'in', '', 'stormtusk'),
    ('Fulmintide', 'shaman', 'Enhancement', 'DPS', 'in', '', 'fulmintide'),
    ('Pyrewisp', 'mage', 'Fire', 'DPS', 'in', '', 'pyrewisp'),
    ('Arcanemoor', 'mage', 'Arcane', 'DPS', 'in', '', 'arcanemoor'),
    ('Frostnettle', 'mage', 'Frost', 'DPS', 'in', '', 'frostnettle'),
    ('Hexbramble', 'warlock', 'Affliction', 'DPS', 'in', '', 'hexbramble'),
    ('Soulgrave', 'warlock', 'Destruction', 'DPS', 'late', 'Flight lands 8:15.', 'soulgrave'),
    ('Grimtotem', 'warrior', 'Fury', 'DPS', 'bench', '', 'kraggor'),
    ('Duskstrider', 'hunter', 'Marksmanship', 'DPS', 'bench', '', 'zulmara'),
    ('Nixthrottle', 'rogue', 'Combat', 'DPS', 'bench', 'Happy to sit.', 'nixthrottle'),
    ('Grimvow', 'priest', 'Shadow', 'DPS', 'out', 'Away this week.', 'grimvow'),
    ('Ashmane', 'warrior', 'Arms', 'DPS', 'out', '', 'ashmane'),
    ('Embercall', 'mage', 'Fire', 'DPS', 'out', '', 'embercall'),
]
NOT_SIGNED = [
    ('Thistlewind', 'druid', 'Balance'), ('Mirestalk', 'hunter', 'Survival'),
    ('Kelgrin', 'rogue', 'Combat'), ('Vorna', 'warlock', 'Demonology'),
]
OFFICER, MEMBER = 'kraggor', 'zulmara'
ONY = "Onyxia's Lair"
DOT = ' · '


def by_game_order(rows):
    return sorted(rows, key=lambda r: (CLASS_ORDER.index(r[1]), r[0]))


def plural(n, word):
    return f'{n} {word}' if n == 1 else f'{n} {word}s'


def counted(rows):
    signed = [r for r in rows if r[4] in ('in', 'late')]
    roles = dict((k, sum(1 for r in signed if r[3] == k)) for k in ('Tank', 'Healer', 'DPS'))
    return len(signed), roles


N_SIGNED, ROLE_N = counted(SIGNUPS)


def role_words(roles):
    return DOT.join([plural(roles['Tank'], 'tank'), plural(roles['Healer'], 'healer'), str(roles['DPS']) + ' DPS'])


# id, zone, title, (weekday, day, month), guild time, officer (Eastern) time, countdown, size, signed, roles, status
PLANNED = [
    ('r1', ONY, ONY, ('THU', '10', 'DEC'), '8:00 pm', '9:00 pm', 'in 2 days', 40, N_SIGNED, ROLE_N, 'planned'),
    ('r2', 'Barrow Deeps', 'Barrow Deeps, first clear', ('SAT', '12', 'DEC'), '7:00 pm', '8:00 pm', 'in 4 days', None, 18,
     dict(Tank=2, Healer=5, DPS=11), 'planned'),
    ('r3', 'Hyjal Summit', 'Hyjal Summit', ('TUE', '15', 'DEC'), '8:00 pm', '9:00 pm', 'in 7 days', None, 11,
     dict(Tank=2, Healer=3, DPS=6), 'cancelled'),
]
# Viewer chip per planned raid: (officer, member); None is not signed up.
VIEWER_CHIP = dict(r1=('In' + DOT + 'Tank', 'In' + DOT + 'DPS'), r2=('In' + DOT + 'Tank', None), r3=('Out', 'Out'))
# Past nights under the planned list: the Raids tab's existing rows, unchanged by G1.
NIGHTS = [('Hyjal Summit', 'Sat 5 Dec', '102 min', '6 pulls'), ('Barrow Deeps', 'Thu 3 Dec', '134 min', '7 pulls')]

# Board styles. Literal hex only (the tokens of DESIGN-SYSTEM.md); alpha as 8-digit hex.
CSS = '''<style>
.seg{display:inline-flex;border:1px solid #262e40;border-radius:4px;overflow:hidden;background:#0d111a}
.seg span{display:inline-flex;align-items:center;justify-content:center;min-width:72px;height:44px;padding:0 10px;font-size:13px;font-weight:600;color:#9a9484}
.seg span+span{border-left:1px solid #262e40}
.seg span.on{color:#f2eee4;background:#131824;box-shadow:inset 0 -2px 0 var(--mark)}
.seg.small span{height:36px;min-width:64px}
.seg.full{display:flex;width:100%}
.seg.full span{flex:1;min-width:0;padding:0}
.seg.off span{opacity:.5}
.field{display:flex;align-items:center;height:44px;padding:0 12px;border:1px solid #3a3326;border-radius:4px;background:#0d111a;color:#e9e4d8;font-size:14px}
.field.ph{color:#9a9484}
.field.sm{height:32px;font-size:13px;padding:0 8px}
.field.bad{border-color:#f0736b}
.chev{margin-left:auto;color:#9a9484;font-size:11px;padding-left:8px}
.card{background:#0d111a;border:1px solid #262e40;border-radius:6px;box-shadow:inset 0 1px 0 #e5b95559}
.ann{margin:0;font-size:12px;color:#9a9484;font-style:italic}
.sub{font-family:var(--font-display);font-size:15px;font-weight:700;color:#f2eee4}
.h2{font-family:var(--font-display);font-size:18px;font-weight:700;text-transform:uppercase;letter-spacing:.04em;color:#f2eee4}
.prow{display:grid;align-items:center;border-bottom:1px solid #1c2230;color:inherit}
.ulink{color:#e9e4d8;text-decoration:underline;text-underline-offset:3px}
.mlink{color:#9a9484;text-decoration:underline;text-underline-offset:3px;font-size:13px;display:inline-flex;align-items:center;min-height:44px}
.btnp{display:inline-flex;align-items:center;height:44px;padding:0 16px;border-radius:4px;background:#e5b955;color:#07090d;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;white-space:nowrap}
.btns{display:inline-flex;align-items:center;height:44px;padding:0 16px;border:1px solid #3a3326;border-radius:4px;color:#e9e4d8;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;white-space:nowrap}
.off{opacity:.5}
.bar{display:inline-block;height:8px;border-radius:999px;background:#1c2230;overflow:hidden;vertical-align:middle}
.bar i{display:block;height:100%;background:#e5b955}
.cb{display:inline-block;width:16px;height:16px;border-radius:3px;background:#e5b955;flex-shrink:0}
.cb.no{background:none;border:1px solid #9a9484}
.cell{display:flex;flex-direction:column;gap:10px}
.cell-t{font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:#e5b955}
.wipe{color:#f0736b}
</style>'''

PILL_KIND = dict(
    site=('#e5b955', '#e5b95524', '#e5b95559'),
    sample=('#9a9484', '#9a94841f', '#9a94844d'),
    kill=('#7bff5c', '#1eff001a', '#1eff0040'),
    wipe=('#f0736b', '#f0736b24', '#f0736b59'),
)
CHIP_KIND = dict(In='kill', Late='kill', Bench='sample', Out='sample')


def pill(text, kind):
    fg, bg, bd = PILL_KIND[kind]
    style = f'color:{fg};background:{bg};border-color:{bd};white-space:nowrap'
    return f'<span class="pill" style="{style}">{text}</span>'


def viewer_chip(text):
    if text is None:
        return pill('Not signed up', 'site')
    return pill(text, CHIP_KIND[text.split(' ')[0]])


def crest_img(cls, size):
    src = crest(cls)
    style = f'width:{size}px;height:{size}px;--c:{CLASS_COLOR[cls]}'
    return f'<img class="crest" src="{src}" alt="" style="{style}">'


def seg(checked=None, mark=KILL, small=False, full=False, disabled=False):
    cls = 'seg' + (' small' if small else '') + (' full' if full else '') + (' off' if disabled else '')
    parts = []
    for word in ('In', 'Late', 'Bench', 'Out'):
        on = ' class="on"' if word == checked else ''
        parts.append(f'<span{on}>{word}</span>')
    colour = KILL if checked in ('In', 'Late') else MUTED
    return f'<div class="{cls}" style="--mark:{colour}">' + ''.join(parts) + '</div>'


def field(text, placeholder=False, select=False, small=False, width='', bad=False):
    cls = 'field' + (' ph' if placeholder else '') + (' sm' if small else '') + (' bad' if bad else '')
    chev = '<span class="chev">▾</span>' if select else ''
    style = f' style="width:{width}"' if width else ''
    return f'<div class="{cls}"{style}><span>{text}</span>{chev}</div>'


def label(text):
    return f'<span class="label">{text}</span>'


def mock_caption():
    return ('<p class="ann">Mock guild and roster, no real sign-ups. Raids: Barrow Deeps, Hyjal Summit, '
            + ONY + ', the first tier, open 9 December 2026. Now: Tue 8 Dec, 6:10 pm guild time (CST).</p>')


def note(text):
    return f'<p class="ann" style="margin-top:4px">{text}</p>'


def bar(width, signed, size):
    pct = round(100 * signed / size)
    return f'<span class="bar" style="width:{width}px"><i style="width:{pct}%"></i></span>'

TABS = dict(
    public=['Overview', 'Raids', 'Progression'],
    member=['Overview', 'Roster', 'Raids', 'Progression', 'Readiness', 'Loot'],
    officer=['Overview', 'Roster', 'Raids', 'Progression', 'Readiness', 'Loot', 'Settings'],
)


def tab_strip(active, role, phone):
    cells = []
    for t in TABS[role]:
        if t == active:
            st = 'border-color:#e5b955;background:#e5b95514;color:#e5b955'
        else:
            st = 'border-color:#262e40;color:#e9e4d8'
        pad = '0 12px' if phone else '0 14px'
        cells.append(f'<span style="display:inline-flex;align-items:center;height:44px;padding:{pad};border:1px solid;'
                     f'border-radius:4px;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;'
                     f'white-space:nowrap;{st}">{t}</span>')
    gut = '18px' if phone else '48px'
    ovf = 'overflow:hidden' if phone else ''
    return f'<div style="display:flex;gap:8px;padding:0 {gut};{ovf}">' + ''.join(cells) + '</div>'


VIEW = dict(
    officer=dict(name='Kraggor', cls='warrior', role='Tank', when='Thursday 9 pm (8 pm guild time)', guild_time=True),
    member=dict(name='Zulmara', cls='hunter', role='DPS', when='Thursday 8 pm', guild_time=False),
)


def raid_line(role, phone, state='In', planned=True):
    v = VIEW[role]
    link = f'<a href="#raid" style="font-size:13px;font-weight:700;white-space:nowrap;display:inline-flex;align-items:center;min-height:44px">{ONY} →</a>'
    if not planned:
        tail = ' <a href="#raids" style="font-size:13px;font-weight:700">Plan a raid</a>' if role == 'officer' else ''
        return f'<p style="margin:0;font-size:14px;color:#9a9484">No raid planned yet.{tail}</p>'
    count = f'{N_SIGNED} of 40 signed up.'
    if state is None:
        sentence = f'{v["when"]}, you have not signed up, {count}'
    else:
        who = v['name']
        sentence = f'{v["when"]}, you are <b style="color:#f2eee4">{state}</b> on {who} as {v["role"]}, {count}'
    segs = seg(state, small=not phone, full=phone)
    if phone:
        return (f'<div style="display:flex;flex-direction:column;gap:8px"><p style="margin:0;font-size:14px">{sentence}</p>'
                f'{segs}{link}</div>')
    return (f'<div style="display:flex;align-items:center;justify-content:space-between;gap:16px;min-height:44px">'
            f'<p style="margin:0;font-size:14px">{sentence}</p><span style="display:flex;align-items:center;gap:16px">{segs}{link}</span></div>')


def guild_band(role, phone, active='Raids', line_state='In', planned=True):
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    v = VIEW.get(role)
    lines = ''
    if v is not None:
        colour = CLASS_COLOR[v['cls']]
        standing = (f'<p style="margin:0;font-size:14px;font-weight:600;color:{colour}">{v["name"]}{DOT}2 of 3 '
                    f'{"Protection Warriors" if role == "officer" else "Marksmanship Hunters"} this week by item level</p>')
        rl = raid_line(role, phone, line_state, planned)
        day = 'Thursday' if planned else 'the next raid'
        needs = f'<p style="margin:0;font-size:13px;color:#ff6b5c">What the guild needs from you before {day}: no enchant: Chest; 1 unspent talent point.</p>'
        lines = standing + rl + needs
    nav_html = nav_phone(v is not None, None) if phone else nav('', v is not None, None)
    return (f'<div class="band" style="display:flex;flex-direction:column">{nav_html}'
            f'<div style="display:flex;flex-direction:column;gap:12px;padding:{pad};{maxw}">'
            f'<span class="label" style="color:#e5b955">Guild</span>'
            f'<h1 class="display" style="margin:0;font-size:22px;color:#f2eee4">{GUILD}</h1>'
            f'<p style="margin:0;font-size:13px;color:#9a9484">PvP US{DOT}1 boss down{DOT}Updated Dec 8</p>{lines}</div>'
            f'<div style="margin-top:18px;padding-bottom:0;{maxw}">{tab_strip(active, role if v else "public", phone)}</div></div>')


def date_block(wd, d, m, phone):
    w, h = (48, 52) if phone else (56, 56)
    return (f'<div style="width:{w}px;height:{h}px;border:1px solid #262e40;border-radius:4px;display:flex;flex-direction:column;'
            f'align-items:center;justify-content:center;line-height:1.05">{label(wd)}'
            f'<span class="mono" style="font-size:22px;color:#f2eee4">{d}</span>{label(m)}</div>')


def count_html(signed, size, roles, bar_w):
    if size is None:
        head = f'<span class="mono" style="font-size:15px;color:#f2eee4">{signed} signed up</span>'
    else:
        b = bar(bar_w, signed, size)
        head = f'<span class="mono" style="font-size:15px;color:#f2eee4">{signed}/{size}</span> {b}'
    return head, f'<span style="font-size:12px;color:#9a9484">{role_words(roles)}</span>'


def planned_row(p, role, phone):
    rid, zone, title, date, gtime, otime, countdown, size, signed, roles, status = p
    officer = role == 'officer'
    local = otime if officer else gtime
    chip_text = VIEWER_CHIP[rid][0 if officer else 1]
    chips = viewer_chip(chip_text)
    if status == 'cancelled':
        chips = pill('Cancelled', 'wipe') + ' ' + chips
    cancelled = status == 'cancelled'
    tcol = '#9a9484' if cancelled else '#f2eee4'
    cd_col = '#9a9484' if (chip_text is None or cancelled) else '#e5b955'
    zone_line = f'<span style="font-size:13px;color:#9a9484">{zone}</span>' if title != zone else ''
    gt = f'{gtime} CST guild time' if officer else ''
    head, words = count_html(signed, size, roles, 64 if phone else 96)
    db = date_block(date[0], date[1], date[2], phone)
    ttl = f'<span style="font-size:15px;font-weight:600;color:{tcol};white-space:nowrap;overflow:hidden;text-overflow:ellipsis">{title}</span>'
    if phone:
        when_parts = [f'<span class="mono">{local}</span>'] + ([gt] if gt else []) + [f'<span class="mono" style="color:{cd_col}">{countdown}</span>']
        return (f'<a class="prow" href="#{rid}" style="grid-template-columns:48px minmax(0,1fr);gap:12px;padding:12px 14px;min-height:96px">'
                f'{db}<div style="display:flex;flex-direction:column;gap:4px;min-width:0">'
                f'<div style="display:flex;justify-content:space-between;gap:8px;align-items:center">{ttl}<span style="display:flex;gap:6px">{chips}</span></div>'
                f'<span style="font-size:13px;color:#e9e4d8">' + DOT.join(when_parts) + '</span>'
                f'<span style="display:flex;align-items:center;gap:8px;flex-wrap:wrap">{head}{words}</span></div></a>')
    gt_html = f'<span style="font-size:12px;color:#9a9484">{gt}</span>' if gt else ''
    return (f'<a class="prow" href="#{rid}" style="grid-template-columns:56px minmax(0,1.5fr) minmax(0,1fr) minmax(0,1fr) 176px 16px;gap:16px;padding:10px 14px;min-height:64px">'
            f'{db}<div style="display:flex;flex-direction:column;min-width:0">{ttl}{zone_line}</div>'
            f'<div style="display:flex;flex-direction:column"><span class="mono" style="font-size:14px">{local}</span>{gt_html}'
            f'<span class="mono" style="font-size:12px;color:{cd_col}">{countdown}</span></div>'
            f'<div style="display:flex;flex-direction:column;gap:2px"><span style="display:flex;align-items:center;gap:8px">{head}</span>{words}</div>'
            f'<span style="display:flex;gap:6px;justify-content:flex-end">{chips}</span><span style="color:#9a9484">›</span></a>')


def gate_sentence():
    return ('<p style="margin:0;font-size:13px;line-height:20px">Planning and running raids are part of the guild plan. '
            '<a href="#settings" style="text-decoration:underline">Subscribe the guild</a>.</p>')


def night_row(n):
    zone, day, dur, pulls = n
    return (f'<div style="display:flex;gap:12px;align-items:center;padding:8px 0;border-bottom:1px solid #1c2230;min-height:44px">'
            f'<span style="min-width:140px;font-weight:600">{zone}</span><span style="color:#9a9484;font-size:13px">{day}</span>'
            f'<span class="mono" style="color:#9a9484;font-size:13px">{dur}</span>'
            f'<span style="margin-left:auto;color:#9a9484;font-size:13px">{pulls}{DOT}no named encounter published</span></div>')


def raids_tab(role, phone, form_html='', empty=False, read_only=False):
    officer = role == 'officer'
    btn = ''
    if officer:
        btn = '<span class="btnp off">Plan a raid</span>' if read_only else '<span class="btnp">Plan a raid</span>'
    head = f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px"><span class="h2">Raids</span>{btn}</div>'
    gate = gate_sentence() if read_only else ''
    if empty:
        msg = 'No raid planned yet. Plan one and the guild can sign up.' if officer else 'No raid planned yet. Officers plan raids here.'
        planned = f'<div style="display:flex;align-items:center;min-height:44px;color:#9a9484;font-size:14px">{msg}</div>'
        count = '0'
    else:
        planned = ''.join(planned_row(p, role, phone) for p in PLANNED)
        count = str(len(PLANNED))
    nights = ''.join(night_row(n) for n in NIGHTS)
    return (f'<section style="display:flex;flex-direction:column;gap:12px">{head}{gate}{form_html}'
            f'<span class="label">Planned <span class="mono">{count}</span></span><div>{planned}</div>'
            f'<span class="label" style="margin-top:12px">Raid nights</span><div>{nights}</div></section>')

STATE_WORD = dict(zip(['in', 'late', 'bench', 'out'], ['In', 'Late', 'Bench', 'Out']))
STATUS_CHIP = dict(planned=('Planned', 'site'), locked=('Locked', 'sample'), cancelled=('Cancelled', 'wipe'),
                   ended=('Ended', 'sample'), done=('Done', 'kill'))


def raid_band(role, phone, status='planned', confirm=False):
    officer = role == 'officer'
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    word, kind = STATUS_CHIP[status]
    chip = pill(word, kind)
    times = '9:00 pm to 12:00 am' if officer else '8:00 pm to 11:00 pm'
    cd = '<span class="mono" style="color:#e5b955">in 2 days 1 h</span>'
    gt = f'<span style="color:#9a9484">8:00 pm CST guild time</span>' if officer else ''
    if phone:
        tail = DOT.join([x for x in (gt, cd) if x])
        when = (f'<p style="margin:0;font-size:15px;display:flex;flex-direction:column">'
                f'<span>Thursday 10 December</span><span class="mono">{times}</span><span>{tail}</span></p>')
    else:
        when = f'<p style="margin:0;font-size:15px">Thursday 10 December{DOT}<span class="mono">{times}</span>' + ''.join(DOT + x for x in (gt, cd) if x) + '</p>'
    b = bar(120 if phone else 160, N_SIGNED, 40)
    count = (f'<p style="margin:0;display:flex;align-items:center;gap:10px;flex-wrap:wrap">{b}'
             f'<span class="mono" style="font-size:15px;color:#f2eee4">{N_SIGNED} of 40 signed up</span>'
             f'<span style="font-size:13px;color:#9a9484">{role_words(ROLE_N)}</span></p>')
    officer_row = ''
    if officer:
        if confirm:
            officer_row = ('<div style="display:flex;flex-wrap:wrap;gap:12px;align-items:center">'
                           '<span style="font-size:14px">Cancel this raid? Members see it as cancelled.</span>'
                           '<span style="display:flex;align-items:center;gap:10px;min-height:44px;font-size:14px"><span class="cb"></span>Post the cancellation to Discord</span>'
                           '<span class="btns">Cancel the raid</span><span class="btns" style="outline:2px solid #e5b955;outline-offset:2px">Keep it</span></div>')
        else:
            officer_row = ('<div style="display:flex;flex-wrap:wrap;gap:8px">' + ''.join(
                f'<span class="btns">{b2}</span>' for b2 in ('Edit', 'Lock sign-ups', 'Copy the sign-up list', 'Cancel raid')) + '</div>')
    eyebrow = f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px;flex-wrap:wrap"><span style="display:flex;align-items:center;gap:10px"><span class="label" style="color:#e5b955">Raid</span>{chip}</span>'
    if officer and not phone:
        eyebrow += officer_row + '</div>'
        officer_row = ''
    else:
        eyebrow += '</div>'
    nav_html = nav_phone(True, None) if phone else nav('', True, None)
    back = f'<a class="mlink" href="#guild">‹ {GUILD}{DOT}Raids</a>'
    return (f'<div class="band" style="display:flex;flex-direction:column">{nav_html}'
            f'<div style="display:flex;flex-direction:column;gap:10px;padding:{pad};{maxw}">{back}{eyebrow}'
            f'<h1 class="display" style="margin:0;font-size:22px;color:#f2eee4">{ONY}</h1>{when}{count}{officer_row}</div>'
            f'<div style="margin-top:18px;{maxw}">{tab_strip("Raids", role, phone)}</div></div>')


def own_rows(role):
    acct = OFFICER if role == 'officer' else MEMBER
    return [r for r in SIGNUPS if r[6] == acct]


def strip_row(r, phone, disabled=False, status_line=''):
    name, cls, spec, rrole, state, note_text, _ = r
    who = (f'<span style="display:flex;align-items:center;gap:10px;min-width:0">{crest_img(cls, 32)}'
           f'<span style="display:flex;flex-direction:column"><span class="display" style="font-size:14px;font-weight:700;color:{TEXT_CLASS[cls]}">{name}</span>'
           f'<span style="font-size:12px;color:#9a9484">{spec}</span></span></span>')
    segs = seg(STATE_WORD[state], full=phone, disabled=disabled)
    role_sel = field(rrole, select=True, width='120px' if phone else '152px')
    note_f = field(note_text, False) if note_text else field('Note for the officers (optional)', placeholder=True)
    line = f'<p style="margin:0;min-height:20px;font-size:13px">{status_line}</p>'
    if phone:
        return (f'<div style="display:flex;flex-direction:column;gap:8px;padding:12px 0;border-bottom:1px solid #1c2230">'
                f'<div style="display:flex;justify-content:space-between;align-items:center;gap:8px">{who}{role_sel}</div>{segs}{note_f}{line}</div>')
    return (f'<div style="padding:12px 0;border-bottom:1px solid #1c2230"><div style="display:grid;grid-template-columns:minmax(180px,1fr) auto 152px minmax(200px,1.4fr);gap:16px;align-items:center">'
            f'{who}{segs}{role_sel}{note_f}</div>{line}</div>')


def strip(role, phone, locked=False, race=False):
    rows = own_rows(role)
    top = ''
    if locked:
        top = '<p id="lock" style="margin:0;font-size:13px;color:#9a9484">Sign-ups are locked. An officer can still change yours.</p>'
    out = []
    for i, r in enumerate(rows):
        status_line = ''
        if race and i == 0:
            status_line = '<span class="wipe" style="outline:2px solid #e5b955;outline-offset:2px">This raid was locked while you were signing up.</span>'
        out.append(strip_row(r, phone, disabled=locked, status_line=status_line))
    pad = '16px' if phone else '20px'
    return f'<section class="card" style="padding:{pad};display:flex;flex-direction:column;gap:4px"><span class="sub">Your sign-up</span>{top}' + ''.join(out) + '</section>'


def roster_row(r, officer, phone, viewer_acct, with_role=False):
    name, cls, spec, rrole, state, note_text, acct = r[0], r[1], r[2], r[3], r[4], r[5], r[6]
    wash = 'background:#e5b9550f;' if acct == viewer_acct else ''
    sub = spec + (DOT + rrole if with_role else '')
    note_html = f'<span style="font-size:12px;color:#e9e4d8">{note_text}</span>' if note_text else ''
    sel = ''
    if officer:
        sel = field(STATE_WORD.get(state, '\u2014'), select=True, small=not phone, width='96px')
    h = 44 if phone else 36
    return (f'<div style="display:flex;align-items:center;gap:8px;min-height:{h}px;padding:4px 6px;border-radius:4px;{wash}">'
            f'{crest_img(cls, 24)}<span style="display:flex;flex-direction:column;min-width:0;flex:1">'
            f'<span style="font-size:14px;font-weight:600;color:{TEXT_CLASS[cls]}">{name}</span>'
            f'<span style="font-size:12px;color:#9a9484">{sub}</span>{note_html}</span>{sel}</div>')


def group(title, rows, officer, phone, viewer_acct, with_role=False, count=None):
    shown = len(rows) if count is None else count
    head = f'<span class="label">{title} <span class="mono">{shown}</span></span>'
    body = ''.join(roster_row(r, officer, phone, viewer_acct, with_role) for r in rows)
    return f'<div style="display:flex;flex-direction:column;gap:4px">{head}{body}</div>'


def roster(role, phone):
    officer = role == 'officer'
    acct = OFFICER if officer else MEMBER
    rows = by_game_order(SIGNUPS)
    in_rows = [r for r in rows if r[4] == 'in']
    tanks = [r for r in in_rows if r[3] == 'Tank']
    heals = [r for r in in_rows if r[3] == 'Healer']
    dps = [r for r in in_rows if r[3] == 'DPS']
    late = [r for r in rows if r[4] == 'late']
    bench = [r for r in rows if r[4] == 'bench']
    out = [r for r in rows if r[4] == 'out']
    notsigned = [(n, c, s, '', 'none', '', '') for n, c, s in by_game_order(NOT_SIGNED)]
    head = f'<div style="display:flex;align-items:baseline;gap:10px"><span class="h2">Sign-ups</span><span class="mono" style="color:#9a9484">{N_SIGNED}/40</span></div>'
    if phone:
        cols = ''.join(group(t, g, officer, phone, acct) for t, g in (('Tanks', tanks), ('Healers', heals), ('DPS', dps)))
        main = f'<div style="display:flex;flex-direction:column;gap:22px">{cols}</div>'
    else:
        half = (len(dps) + 1) // 2
        d1 = group('DPS', dps[:half], officer, phone, acct, count=len(dps))
        d2 = ''.join(roster_row(r, officer, phone, acct) for r in dps[half:])
        main = (f'<div style="display:grid;grid-template-columns:repeat(4,minmax(0,1fr));gap:16px;align-items:start">'
                f'{group("Tanks", tanks, officer, phone, acct)}{group("Healers", heals, officer, phone, acct)}{d1}'
                f'<div style="display:flex;flex-direction:column;gap:4px;padding-top:21px">{d2}</div></div>')
    tail_cols = 'repeat(3,minmax(0,1fr))' if not phone else '1fr'
    tail = (f'<div style="display:grid;grid-template-columns:{tail_cols};gap:16px;margin-top:24px;align-items:start">'
            f'{group("Late", late, officer, phone, acct, True)}{group("Bench", bench, officer, phone, acct, True)}'
            f'{group("Out", out, officer, phone, acct, True)}</div>')
    if officer:
        ns = f'<div style="margin-top:16px">{group("Not signed up", notsigned, officer, phone, acct)}</div>'
    else:
        names = ', '.join(f'<span style="color:{TEXT_CLASS[c]}">{n}</span>' for n, c, s, *_ in notsigned)
        ns = (f'<div style="margin-top:16px;display:flex;flex-direction:column;gap:4px"><span class="label">Not signed up <span class="mono">{len(notsigned)}</span></span>'
              f'<p style="margin:0;font-size:13px">{names}</p></div>')
    return f'<section style="display:flex;flex-direction:column;gap:12px">{head}{main}{tail}{ns}</section>'


NOTES = ('<p style="margin:0 0 10px">Pull at <b>8:15</b>. Invites from 7:45.</p>'
         '<ul style="margin:0;padding-left:18px;line-height:1.6"><li>Flasks and food buffs on before the pull.</li>'
         '<li>Fire resistance potions for the tanks.</li><li>Phase 2: spread on the walls, no one under her.</li></ul>')


def notes(role):
    return (f'<section style="display:flex;flex-direction:column;gap:10px"><span class="h2">Notes</span>'
            f'<div style="font-size:14px;line-height:1.6;color:#e9e4d8;max-width:720px">{NOTES}</div></section>')


def raid_body(role, phone, locked=False, race=False):
    s = strip(role, phone, locked, race)
    n = notes(role)
    r = roster(role, phone)
    gap = 22 if phone else 32
    return f'<div style="display:flex;flex-direction:column;gap:{gap}px">{s}{n}{r}</div>'


def lab_field(text, inner, helper=''):
    help_html = f'<span style="font-size:13px;color:#9a9484">{helper}</span>' if helper else ''
    return f'<label style="display:flex;flex-direction:column;gap:8px;min-width:0">{label(text)}{inner}{help_html}</label>'


def check_row(text, on=True, helper=''):
    box = '<span class="cb"></span>' if on else '<span class="cb no"></span>'
    help_html = f'<span style="font-size:12px;color:#9a9484;padding-left:28px">{helper}</span>' if helper else ''
    return (f'<div style="display:flex;flex-direction:column"><span style="display:flex;align-items:center;gap:12px;min-height:44px;font-size:14px">'
            f'{box}{text}</span>{help_html}</div>')


def form_panel(phone, errors=False, cap=False, webhook=True):
    raid = lab_field('Raid', field('Choose a raid', placeholder=True, select=True, bad=True) if errors else field(ONY, select=True),
                     '<span class="wipe">Choose a raid.</span>' if errors else '')
    title = lab_field('Title', field(ONY, placeholder=True))
    date = lab_field('Date', field('mm/dd/yyyy', placeholder=True, bad=True), '<span class="wipe">Pick a date.</span>') if errors else lab_field('Date', field('Thu, Dec 10, 2026'))
    time_f = lab_field('Time', field('8:00 PM'))
    dur = lab_field('Duration', field('3 h', select=True))
    tz = lab_field('Guild time zone', field('Central (Chicago)' + DOT + 'CST', select=True), 'For you: Thu 10 Dec, 9:00 pm.')
    notes_box = (f'<div class="field" style="height:132px;align-items:flex-start;padding:10px 12px;line-height:1.5">'
                 f'Pull at 8:15. Invites from 7:45.<br>- Fire resistance potions for the tanks.</div>')
    nts = lab_field('Notes', notes_box, '<span style="font-size:12px">Markdown: bold, lists and links. Members see this on the raid page.</span>')
    rep_help = 'Only 2 more fit: a guild plans at most 12 raids ahead.' if cap else ''
    rep = check_row('Also plan the next 3 weeks at this time', on=False, helper=rep_help)
    if webhook:
        disc = check_row('Post to Discord', on=True)
    else:
        disc = ('<p style="margin:0;font-size:13px;color:#9a9484;min-height:44px;display:flex;align-items:center">Connect a Discord channel in '
                '<a href="#settings" style="text-decoration:underline;margin-left:4px">Settings</a>&nbsp;to post raids there.</p>').replace('&nbsp;', ' ')
    actions = ('<div style="display:flex;align-items:center;gap:16px;flex-wrap:wrap"><span class="btnp">Post the raid</span>'
               '<a class="ulink" href="#cancel" style="display:inline-flex;align-items:center;min-height:44px;font-size:13px">Cancel</a></div>'
               '<p style="margin:0;min-height:20px"></p>')
    if phone:
        grid = f'<div style="display:flex;flex-direction:column;gap:16px">{raid}{title}{date}{time_f}{dur}{tz}{nts}</div>'
    else:
        grid = (f'<div style="display:grid;grid-template-columns:1fr 1fr;gap:16px">{raid}{title}</div>'
                f'<div style="display:grid;grid-template-columns:1fr 1fr 1fr;gap:16px">{date}{time_f}{dur}</div>{tz}{nts}')
    pad = '16px' if phone else '20px'
    return (f'<section class="card" style="padding:{pad};display:flex;flex-direction:column;gap:16px">'
            f'<span class="sub">Plan a raid</span>{grid}{rep}{disc}{actions}</section>')


# Discord message chrome, the same values gen_discord_notifications.py draws.
DISC_BG = '#313338'
DISC_EMBED = '#2b2d31'
DF = "'gg sans','Noto Sans','Helvetica Neue',Arial,sans-serif"


def embed(title, fields, desc, colour, width, phone):
    cols = 1 if phone else 3
    flds = ''.join(f'<div style="min-width:0"><div style="font-size:14px;font-weight:600;color:#f2f3f5">{n}</div>'
                   f'<div style="font-size:14px;color:#dbdee1">{v}</div></div>' for n, v in fields)
    ds = f'<div style="font-size:14px;color:#dbdee1;margin-top:8px">{desc}</div>' if desc else ''
    return (f'<div style="max-width:{width}px;background:{DISC_EMBED};border-left:4px solid {colour};border-radius:4px;padding:8px 16px 16px 12px">'
            f'<div style="font-size:16px;font-weight:600;color:#00a8fc;margin-top:8px">{title}</div>{ds}'
            f'<div style="display:grid;grid-template-columns:repeat({cols},minmax(0,1fr));gap:8px 16px;margin-top:8px">{flds}</div>'
            f'<div style="font-size:12px;color:#b5bac1;margin-top:12px">{GUILD}{DOT}Forever Sixty{DOT}Today at 6:12 PM</div></div>')


def message(emb, width):
    av = (f'<span style="width:40px;height:40px;border-radius:50%;background:#0d111a;display:grid;place-items:center;flex-shrink:0">'
          f'<img src="{SEAL}" width="28" height="28" alt=""></span>')
    tag = '<span style="background:#5865f2;color:#fff;font-size:10px;font-weight:600;border-radius:3px;padding:1px 5px;margin:0 6px 0 4px">APP</span>'
    return (f'<div style="width:{width}px;background:{DISC_BG};border-radius:8px;padding:16px;display:flex;gap:16px;font-family:{DF}">{av}'
            f'<div style="min-width:0;flex:1"><div><span style="font-size:16px;font-weight:500;color:#f2f3f5">Forever Sixty</span>{tag}'
            f'<span style="font-size:12px;color:#949ba4">Today at 6:12 PM</span></div><div style="margin-top:4px">{emb}</div></div></div>')


# Discord draws a <t:unix:F> stamp in each reader's own zone; shown here for a reader in Central time.
WHEN = 'Thursday, December 10, 2026 8:00 PM (in 2 days)'
EMBEDS = [
    ('raid_posted', ONY + DOT + 'Thursday 10 December', '#e5b955', 'Sign up on Forever Sixty.',
     [('When', WHEN), ('Guild time', '8:00 pm CST'), ('Signed up', f'{N_SIGNED} of 40')]),
    ('raid_changed, time', 'Moved: ' + ONY + DOT + 'Friday 11 December', '#e5b955', 'Sign-ups carry over. Check you can still make it.',
     [('When', 'Friday, December 11, 2026 8:00 PM (in 3 days)'), ('Was', 'Thursday, December 10, 2026 8:00 PM'), ('Signed up', f'{N_SIGNED} of 40')]),
    ('raid_changed, cancel', 'Cancelled: ' + ONY + DOT + 'Thursday 10 December', '#f0736b', '',
     [('When', '<s>' + WHEN + '</s>'), ('Signed up', f'{N_SIGNED} of 40')]),
]


def discord_board():
    cells = []
    for name, title, colour, desc, fields in EMBEDS:
        desk = message(embed(title, fields, desc, colour, 524, False), 600)
        ph = message(embed(title, fields, desc, colour, 278, True), 354)
        cells.append(f'<div class="cell"><span class="cell-t">{name}</span><div style="display:flex;gap:24px;align-items:flex-start">{desk}{ph}</div></div>')
    return ''.join(cells) + note('Colour 15055189 (gold) for posted and moved, 15758187 (wipe) for cancelled. No thumbnail: this mock guild has no crest. Title links to the raid page.')


def cell(title, body, caption='', width=640):
    cap = note(caption) if caption else ''
    return (f'<div class="cell" style="width:{width}px"><span class="cell-t">{title}</span>'
            f'<div style="border:1px dashed #262e40;border-radius:6px;padding:16px;background:#07090d">{body}</div>{cap}</div>')


def sk(h, w='100%'):
    return f'<div style="height:{h}px;width:{w};border-radius:4px;background:#131824"></div>'


def states_board():
    member_signed = raid_line('member', False, 'In')
    unsigned = raid_line('member', False, None)
    none_m = raid_line('member', False, planned=False)
    none_o = raid_line('officer', False, planned=False)
    waiting = ('<div style="display:flex;align-items:center;gap:10px;opacity:.8">' + crest_img('hunter', 32)
               + '<span style="display:flex;flex-direction:column"><span class="display" style="font-size:14px;font-weight:700;color:#9a9484">Duskstrider</span>'
               '<span style="font-size:13px;color:#9a9484">Waiting for approval. You can sign it up once an officer approves it.</span></span></div>')
    lines = ''.join(f'<p style="margin:0 0 10px;font-size:14px">{t}</p>' for t in (
        pill('Cancelled', 'wipe') + '  This raid was cancelled.',
        pill('Ended', 'sample') + '  This raid has ended.',
        pill('Done', 'kill') + '  Logged as a raid night. <a href="#night" style="text-decoration:underline">Open the night</a>',
        pill('Locked', 'sample') + '  Sign-ups are locked. An officer can still change yours.'))
    visitor = (f'<p style="margin:0 0 12px;font-size:14px">Planned raids are for members of {GUILD}.</p>'
               '<p style="margin:0 0 10px;font-size:14px">See where you stand in this guild.</p><span class="btnp">Sign in with Battle.net</span>')
    unverified = '<p style="margin:0;font-size:14px">You are waiting for approval. Planned raids show once an officer approves you.</p>'
    loading = ('<div style="display:flex;flex-direction:column;gap:10px">' + sk(22, '40%') + sk(18, '60%') + sk(18, '50%')
               + '<div style="height:12px"></div>' + sk(68) + sk(68) + '<div style="height:12px"></div>' + sk(200) + '</div>')
    copy_fail = ('<p class="wipe" style="margin:0 0 8px;font-size:13px">Could not copy. Select the list and copy it.</p>'
                 '<div class="field" style="height:auto;align-items:flex-start;padding:10px 12px;font-family:var(--font-mono);font-size:12px;line-height:1.6;white-space:pre">'
                 + ONY + DOT + 'Thu 10 Dec, 8:00 pm CST<br>23 of 40: 3 tanks, 7 healers, 13 DPS<br>Tanks: Kraggor, Sunderfel, Thornhide<br>Healers: Lightbrand, Brightvow, Rootgall, ...</div>')
    gate = raids_tab('officer', False, read_only=True).split('<span class="label">Planned')[0] + '</section>'
    empty_m = raids_tab('member', False, empty=True).split('<span class="label" style="margin-top:12px">')[0] + '</section>'
    empty_o = raids_tab('officer', False, empty=True).split('<span class="label" style="margin-top:12px">')[0] + '</section>'
    locked_strip = strip('member', False, locked=True, race=True)
    cells = [
        cell('Band line, member signed In (two own characters)', member_signed, width=1344),
        cell('Band line, unsigned', unsigned, width=1344),
        cell('Band line, none planned: member', none_m, width=660),
        cell('Band line, none planned: officer', none_o, width=660),
        cell('Raids tab, empty Planned: member', empty_m, width=660),
        cell('Raids tab, empty Planned: officer', empty_o, width=660),
        cell('Raids tab, officer-read-only', gate, 'The button keeps its place, disabled; the gate sentence sits under the header row.', width=660),
        cell('Raid states (chip in the band, line in place of the strip)', lines, width=660),
        cell('Locked strip, after the lock race (status line focused)', locked_strip, width=1344),
        cell('Raid page, signed-out visitor (R1)', visitor, width=660),
        cell('Raid page, unverified viewer', unverified, width=660),
        cell('Strip, a character waiting for approval', waiting, width=660),
        cell('Copy the sign-up list failed', copy_fail, width=660),
        cell('Raid page loading (reserve heights)', loading, 'Band lines as skeleton inside the band min height; one strip row per stored own character; roster min-h 420.', width=660),
        cell('Form: errors, repeat at the cap, no webhook', form_panel(False, errors=True, cap=True, webhook=False), width=660),
        cell('Cancel confirm (officer)', raid_band('officer', True, confirm=True), width=390),
    ]
    return '<div style="display:flex;flex-wrap:wrap;gap:24px">' + ''.join(cells) + '</div>'


def board(name, band_html, body, width, phone):
    pad = '18px' if phone else '48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    content = (f'<div style="display:flex;flex-direction:column;gap:22px;padding:24px {pad} 32px {pad};{maxw}">'
               + body + mock_caption() + '</div>')
    return write(f'{name}.html', page(CSS + band_html + content + footer(phone=phone), width)).as_posix()


def sheet(name, title, body, width=1440):
    content = (f'<div style="display:flex;flex-direction:column;gap:24px;padding:32px 48px">'
               f'<h1 class="display" style="margin:0;font-size:22px;color:#f2eee4">{title}</h1>{body}{mock_caption()}</div>')
    return write(f'{name}.html', page(CSS + content, width)).as_posix()


if __name__ == '__main__':
    print(board('guild-raids-tab', guild_band('officer', False), raids_tab('officer', False), 1440, False))
    print(board('guild-raids-tab-phone', guild_band('member', True), raids_tab('member', True), 390, True))
    print(board('guild-raids-page', raid_band('officer', False), raid_body('officer', False), 1440, False))
    print(board('guild-raids-page-phone', raid_band('member', True), raid_body('member', True), 390, True))
    print(board('guild-raids-form', guild_band('officer', False), raids_tab('officer', False, form_html=form_panel(False)), 1440, False))
    print(board('guild-raids-form-phone', guild_band('officer', True), raids_tab('officer', True, form_html=form_panel(True)), 390, True))
    print(sheet('guild-raids-states', 'Raids and sign-ups: states', states_board()))
    print(sheet('guild-raids-discord', 'Raids and sign-ups: Discord embeds', discord_board()))
