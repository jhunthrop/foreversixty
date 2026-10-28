"""Mock boards for the log coach's three surfaces (design/specs/2026-10-28-log-coach.md).

  python3 design/mocks/log-coach/gen_log_coach.py
    writes log-coach-panel.html            (1440, the report page, Coach tab, a model review)

Built on gen_guild_raids (G1) and gen_discord_bot: their CSS, pills, crests and Discord chrome are imported,
never copied. Boards go through mocklib.write into design/mocks/renders/; render.py screenshots them (names
with 'phone' at 390, '2000' at 2000, '1024' at 1024, the rest at 1440).
"""
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE.parent))
sys.path.insert(0, str(HERE))

import coach_copy as C  # noqa: E402
import gen_discord_bot as D  # noqa: E402
from gen_guild_raids import CSS as G1_CSS, DOT, GUILD, KILL, WIPE, crest_img, pill  # noqa: E402
from mocklib import CLASS_COLOR, PARSE, footer, nav, nav_phone, page, write  # noqa: E402

BOSS = 'Onyxia'
GOLD, MUTED, STRONG = '#e5b955', '#9a9484', '#f2eee4'
LH = 22  # body line height inside the review, px
CAPTION = ('Mock. Fury warrior Tyrim, Onyxia (40), kill, 3:52. Every value is the product spec section 5.3 invented '
           "example, not a measurement. Spell icons are the build's; the flask and food icons are stand-ins.")

# Lines the model's longest allowed text takes in each slot (spec section 4.3): desktop 720 px, phone 328 px measure.
RESERVE_LINES = dict(open=(3, 7), why=(4, 8), drill=(3, 6), then=(3, 5), well=(2, 4))

ICONS = dict([
    ('Bloodthirst', 'spell_nature_bloodlust'), ('Battle Shout', 'ability_warrior_battleshout'),
    ('Whirlwind', 'ability_whirlwind'), ('Shield Wall', 'ability_warrior_shieldwall'),
    ('Sunder Armor', 'ability_warrior_sunder'), ('Greater Heal', 'spell_holy_greaterheal'),
    ('Inner Fire', 'spell_holy_innerfire'), ('Flask', 'inv_potion_80'), ('Food', 'inv_misc_food_13'),
    ('Rage', 'spell_nature_bloodlust'), ('Shield Block', 'ability_defend'),
])


def ico(name: str, size: int = 20) -> str:
    return (f'<img class="ico" src="{D.icon_uri(ICONS[name])}" width="{size}" height="{size}" alt="" '
            f'style="width:{size}px;height:{size}px">')

CSS = '''<style>
.shell{width:100%;max-width:1344px;margin:0 auto}
.rband{padding:22px 0 0}
.rtabs{display:flex;gap:4px;align-items:flex-end;margin-top:14px;overflow:hidden}
.rtab{display:inline-flex;align-items:center;height:44px;padding:0 14px;font-size:13px;font-weight:600;color:#9a9484;border-bottom:2px solid transparent;white-space:nowrap}
.rtab.on{color:#f2eee4;border-bottom-color:#e5b955}
.fchip{display:inline-flex;align-items:center;gap:8px;height:44px;padding:0 12px;border:1px solid #262e40;border-radius:4px;font-size:13px;color:#e9e4d8;white-space:nowrap}
.fchip.on{border-color:#e5b955;background:#131824}
.ico{display:inline-block;border-radius:3px;flex-shrink:0;vertical-align:middle;box-shadow:0 0 0 1px #00000099}
.arow{border-bottom:1px solid #1c2230}
.agrid{display:grid;grid-template-columns:28px 40px minmax(120px,1.4fr) 52px minmax(0,3fr) 92px 80px 64px;align-items:center;gap:12px;min-height:44px;padding:8px}
.agrid.ph{grid-template-columns:28px auto minmax(96px,1fr) auto;row-gap:4px}
.abar{height:10px;background:#1c2230;border-radius:2px;overflow:hidden;display:flex}.abar i{display:block;height:100%}
.mono2{font-family:var(--font-mono);font-variant-numeric:tabular-nums}
.cline{display:flex;align-items:center;gap:8px;min-height:36px;padding:0 8px 0 100px;font-size:13px;color:#e9e4d8;border-top:1px solid #1c2230;background:#0b0f17}
.cline.ph{min-height:44px;padding-left:48px;padding-right:8px}
.cline .lead{color:#9a9484;white-space:nowrap}.cline b{font-weight:600;color:#f2eee4}
.cline .chev{margin-left:auto;color:#9a9484;font-size:14px;padding-left:8px}
.cline.hover{background:#131824}.cline.hover b{text-decoration:underline;text-underline-offset:3px}
.cline.focus{outline:2px solid #e5b955;outline-offset:-2px}
.cline.muted{color:#9a9484}
.cline .bar-sk{height:12px;border-radius:3px;background:#131824;width:260px;max-width:60%}
.cp{background:#0d111a;border:1px solid #262e40;border-radius:6px;box-shadow:inset 0 1px 0 #e5b95559;padding:24px}
.cp.ph{padding:18px}
.cp-h{display:flex;align-items:center;justify-content:space-between;gap:12px;min-height:44px}
.cp-t{font-family:var(--font-display);font-weight:700;font-size:17px;line-height:24px;color:#f2eee4}
.cp-stat{display:inline-flex;align-items:center;gap:8px;font-size:13px;color:#9a9484;white-space:nowrap}
.dot{width:8px;height:8px;border-radius:999px;background:#e5b955;display:inline-block}
.body15{font-size:15px;line-height:22px;color:#e9e4d8;margin:0;max-width:720px}
.fix{border-left:2px solid #e5b955;padding-left:16px;margin-top:20px}
.fix-t{display:flex;align-items:center;gap:10px;font-size:20px;line-height:28px;font-weight:700;color:#f2eee4;margin:4px 0 8px}
.drill{margin-top:14px;padding-top:12px;border-top:1px solid #1c2230}
.chips{display:flex;flex-wrap:wrap;gap:6px 8px;margin-top:12px}
.chip{display:inline-flex;align-items:center;gap:6px;min-height:36px;padding:0 10px;border:1px solid #262e40;border-radius:4px;font-size:12px;color:#e9e4d8;white-space:nowrap}
.chip.ph{min-height:44px}.chip.hover{border-color:#e5b955;color:#f2eee4}.chip .k{color:#9a9484}
.then{margin-top:24px}
.titem{padding:14px 0;border-top:1px solid #1c2230}
.titem h4{margin:0;display:flex;align-items:center;gap:8px;font-size:16px;line-height:24px;font-weight:600;color:#f2eee4}
.well{margin-top:20px}
.tick{color:#7fd48a;flex-shrink:0;width:14px;margin-top:4px}
.cc{margin:0;padding:0;list-style:none}.cc li{font-size:13px;line-height:20px;color:#9a9484;padding:2px 0}
.rail{display:flex;flex-direction:column;gap:20px}
.rcard{border:1px solid #1c2230;border-radius:6px;padding:16px;background:#0b0f17}
.foot12{font-size:12px;line-height:18px;color:#9a9484}
.lnk{display:inline-flex;align-items:center;min-height:44px;color:#e9e4d8;text-decoration:underline;text-underline-offset:3px;font-size:13px}
.slot{position:relative}
.slot.annot{outline:1px dashed #5a4d33;outline-offset:2px}
.slot.annot::after{content:attr(data-h);position:absolute;right:0;bottom:-16px;font-size:10px;color:#e5b955;font-family:var(--font-mono)}
.sk{background:#131824;border-radius:3px;height:14px;margin:8px 0}
.hrow{display:flex;gap:6px;flex-wrap:wrap}
.btn44{display:inline-flex;align-items:center;justify-content:center;height:44px;padding:0 16px;border:1px solid #e5b955;border-radius:4px;color:#e5b955;font-size:14px;font-weight:600;white-space:nowrap}
.btn44.sec{border-color:#3a3326;color:#e9e4d8}
.note{font-size:13px;line-height:20px;color:#9a9484;margin:0}
.statebox{border:1px dashed #262e40;border-radius:6px;padding:16px;background:#07090d}
.lrline{display:flex;align-items:center;gap:8px;min-height:44px;padding:8px 0;font-size:13px;color:#e9e4d8;border-bottom:1px solid #1c2230;line-height:20px}
.lrline.ph{min-height:64px;align-items:flex-start}
.lrline .m{color:#9a9484}.lrline b{font-weight:600;color:#f2eee4}.lrline u{text-underline-offset:3px;color:#e5b955}
.crow{display:flex;align-items:center;gap:12px;min-height:64px;border-bottom:1px solid #1c2230;padding:8px 0}
</style>'''

TICK = ('<svg class="tick" viewBox="0 0 14 14" width="14" height="14" fill="none" stroke="#7fd48a" stroke-width="2" '
        'stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M2 7.5l3 3 7-7"/></svg>')


def label(text: str, color: str = MUTED) -> str:
    return f'<span class="label" style="color:{color}">{text}</span>'


def cell(title: str, body: str, width: int = 640, caption: str = '') -> str:
    cap = f'<p class="ann" style="margin:0">{caption}</p>' if caption else ''
    return (f'<div class="cell" style="width:{width}px;max-width:100%"><span class="cell-t">{title}</span>'
            f'<div class="statebox">{body}</div>{cap}</div>')


# The report page chrome.

ROWS = [  # rank, parse, name, class, share, amount, per sec, active, bar %
    (1, 98, 'Vexlash', 'rogue', 14.2, '412k', '1,771', '97%', 100),
    (2, 91, 'Pyrewisp', 'mage', 12.9, '374k', '1,612', '95%', 91),
    (3, 87, 'Tyrim', 'warrior', 11.8, '342k', '1,474', '88%', 83),
    (4, 74, 'Windtalon', 'hunter', 10.6, '307k', '1,323', '93%', 75),
    (5, 62, 'Kraggor', 'warrior', 9.1, '264k', '1,138', '91%', 64),
    (6, 41, 'Thornhide', 'druid', 7.4, '215k', '927', '84%', 52),
]
OWNED = 'Tyrim'


def parse_colour(p: int) -> str:
    if p >= 99:
        return PARSE['pink']
    if p >= 95:
        return PARSE['orange']
    if p >= 75:
        return PARSE['purple']
    if p >= 50:
        return PARSE['blue']
    return PARSE['green'] if p >= 25 else PARSE['grey']


def arow(r: tuple, phone: bool, line: str = '', detail: str = '') -> str:
    rank, parse, name, cls, share, amount, persec, active, bar = r
    col = CLASS_COLOR[cls]
    barh = f'<div class="abar"><i style="width:{bar}%;background:{col}"></i></div>'
    pc = f'<span class="mono2" style="font-size:12px;color:{parse_colour(parse)}">{parse}</span>'
    nm = (f'<span style="display:flex;align-items:center;gap:8px;color:{col};font-weight:600">'
          f'{crest_img(cls, 20)}{name}</span>')
    rk = f'<span class="mono2" style="font-size:12px;color:{MUTED}">› {rank}</span>'
    if phone:
        grid = (f'<div class="agrid ph">{rk}{pc}{nm}<span class="mono2" style="text-align:right">{amount}</span>'
                f'<span style="grid-column:1/-1">{barh}</span>'
                f'<span class="label" style="grid-column:1/3">{active} active</span>'
                f'<span class="mono2" style="grid-column:3/-1;text-align:right;color:{MUTED};font-size:13px">'
                f'{persec} per sec</span></div>')
    else:
        grid = (f'<div class="agrid">{rk}{pc}{nm}'
                f'<span class="mono2" style="font-size:12px;color:{MUTED};text-align:right">{share}%</span>{barh}'
                f'<span class="mono2" style="text-align:right">{amount}</span>'
                f'<span class="mono2" style="text-align:right;color:{MUTED};font-size:13px">{persec}</span>'
                f'<span class="mono2" style="text-align:right;color:{MUTED};font-size:13px">{active}</span></div>')
    return f'<li class="arow">{grid}{line}{detail}</li>'


def chrome(inner: str, phone: bool, active: str = 'Coach') -> str:
    tabs = ['Summary', 'Damage Done', 'Healing', 'Damage Taken', 'Deaths']
    if phone:
        tabs = ['Summary', 'Healing', 'Deaths']
    on_more = active not in tabs
    ts = ''.join(f'<span class="rtab{" on" if t == active else ""}">{t}</span>' for t in tabs)
    more = f'<span class="rtab{" on" if on_more else ""}">{active if on_more else "More"} ▾</span>'
    pad = '18px' if phone else '48px'
    fights = f'<span class="fchip on">{BOSS}{DOT}3:52{DOT}<span style="color:{KILL}">Kill</span></span>'
    if not phone:
        fights += f'<span class="fchip">Trash{DOT}1:40</span><span class="fchip">Hyjal Summit{DOT}6:14</span>'
    h1 = 'font-size:20px' if phone else 'font-size:22px'
    band = (f'<div class="band" style="padding:0 {pad};border-bottom:1px solid #1c2230"><div class="shell rband">'
            f'{label("ONYXIA&#39;S LAIR · THU 10 DEC", GOLD)}'
            f'<h1 class="display" style="{h1};color:{STRONG};margin:4px 0 10px">Olympus XXVII raid night</h1>'
            f'<div class="hrow">{fights}</div><div class="rtabs">{ts}{more}</div></div></div>')
    return band + f'<div style="padding:20px {pad} 32px"><div class="shell">{inner}</div></div>'


# The review.

def chip(icon: str, kind: str, text: str, phone: bool, hover: bool = False) -> str:
    cls = 'chip' + (' ph' if phone else '') + (' hover' if hover else '')
    pic = ico(icon, 16) if icon else ''
    return f'<a class="{cls}">{pic}<span>{text}</span><span class="k">· {kind}</span></a>'


def slot(text: str, key: str, phone: bool, reserve: bool, annot: bool, size: int = 15) -> str:
    lines = RESERVE_LINES[key][1 if phone else 0]
    h = lines * LH
    style = f'min-height:{h}px' if reserve else ''
    a = ' annot' if annot and reserve else ''
    return (f'<div class="slot{a}" data-h="{h} px" style="{style}"><p class="body15" '
            f'style="font-size:{size}px">{text}</p></div>')


def link_line(text: str, phone: bool, reserve: bool) -> str:
    if not text:
        return f'<div style="margin-top:8px;min-height:{(4 if phone else 2) * 20}px"></div>' if reserve else ''
    h = 'min-height:%dpx;' % ((4 if phone else 2) * 20) if reserve else ''
    glyph = ('<svg width="12" height="12" viewBox="0 0 12 12" fill="none" stroke="#e5b955" stroke-width="1.5" '
             'stroke-linecap="round" style="flex-shrink:0;margin-top:4px"><path d="M4 8l4-4M3 6.5L2 7.5a2 2 0 0 0 2.8 2.8L6 9.5M9 5.5l1-1A2 2 0 0 0 7.2 1.7L6 2.5"/></svg>')
    return f'<div style="display:flex;gap:8px;margin-top:8px;{h}">{glyph}<p class="note" style="margin:0">{text}</p></div>'


def fix_block(title: str, why: str, drill: str, phone: bool, reserve: bool, annot: bool, link: str = '') -> str:
    chips = ''.join([chip('Bloodthirst', 'Casts', 'Bloodthirst', phone), chip('', 'Timelines', '1:12', phone),
                     chip('', 'Timelines', '2:40', phone), chip('', 'Resources', '2:41', phone)])
    return (f'<div class="fix">{label("Fix first", GOLD)}'
            f'<div class="fix-t">{ico("Bloodthirst", 32)}<span>{title}</span></div>'
            f'{slot(why, "why", phone, reserve, annot)}{link_line(link, phone, reserve)}'
            f'<div class="drill">{label("Practice")}{slot(drill, "drill", phone, reserve, annot, 14)}</div>'
            f'<div class="chips">{chips}</div></div>')


def then_block(items: list, phone: bool, reserve: bool, annot: bool) -> str:
    out = ''
    for (title, _detail, why), ic in zip(items, ['Flask', 'Rage']):
        chips = chip('', 'Buffs', 'at the pull', phone) if ic == 'Flask' else chip('', 'Resources', '2:41', phone)
        out += (f'<div class="titem"><h4><span>{title}</span></h4>'
                f'{slot(why, "then", phone, reserve, annot, 14)}<div class="chips" style="margin-top:8px">{chips}</div></div>')
    return f'<div class="then">{label("Then")}{out}</div>'


def well_block(items: list, phone: bool, reserve: bool) -> str:
    lis = ''.join(f'<li style="display:flex;gap:8px;font-size:14px;line-height:22px;padding:4px 0;color:#e9e4d8">'
                  f'{TICK}<span>{t}</span></li>' for t in items)
    h = RESERVE_LINES['well'][1 if phone else 0] * LH + 8
    style = f'min-height:{h * len(items)}px;' if reserve else ''
    return (f'<div class="well">{label("Doing well", KILL)}'
            f'<ul style="list-style:none;margin:6px 0 0;padding:0;{style}">{lis}</ul></div>')


def rail(phone: bool, history: bool = True, foot: str = C.FURY_FOOT, cannot: list = None, who: bool = True) -> str:
    cannot = C.FURY_CANNOT if cannot is None else cannot
    wh = (f'<div class="rcard" style="display:flex;gap:12px;align-items:center">{crest_img("warrior", 40)}'
          f'<div><div style="color:{CLASS_COLOR["warrior"]};font-weight:600;font-size:15px">Tyrim</div>'
          f'<div class="note">Fury Warrior · {BOSS} (40) · Kill 3:52</div></div></div>') if who and not phone else ''
    cc = (f'<div class="rcard">{label("Couldn&#39;t check")}<ul class="cc" style="margin-top:6px">'
          + ''.join(f'<li>{c}</li>' for c in cannot) + '</ul></div>') if cannot else ''
    hist = '<a class="lnk">History</a>' if history else ''
    ft = f'<div class="rcard"><p class="foot12" style="margin:0">{foot}</p>{hist}</div>'
    return f'<aside class="rail">{wh}{cc}{ft}</aside>'


def header(title: str, right: str, sub: str = '') -> str:
    s = f'<p class="note" style="margin:2px 0 0">{sub}</p>' if sub else ''
    return f'<div class="cp-h"><span class="cp-t">{title}</span><span class="cp-stat">{right}</span></div>{s}'


def main_col(title: str, right: str, opening: str, why: str, drill: str, then: list, well: list, phone: bool,
             reserve: bool = False, annot: bool = False, sub: str = '', busy: bool = False, link: str = '') -> str:
    bz = (f'<p class="note" style="margin:10px 0 0;border-left:2px solid #3a3326;padding-left:12px">{C.BUSY}</p>'
          if busy else '')
    return (f'<section class="cp{" ph" if phone else ""}">{header(title, right, sub)}{bz}'
            f'<div style="margin-top:8px">{slot(opening, "open", phone, reserve, annot)}</div>'
            f'{fix_block(C.FURY_TITLE, why, drill, phone, reserve, annot, link)}'
            f'{then_block(then, phone, reserve, annot)}{well_block(well, phone, reserve)}</section>')


def two_col(main: str, side: str, width: int, phone: bool) -> str:
    if phone or width < 1100:
        return f'<div style="display:flex;flex-direction:column;gap:16px">{main}{side}</div>'
    return (f'<div style="display:grid;grid-template-columns:minmax(0,760px) 320px;gap:40px;align-items:start">'
            f'{main}{side}</div>')


def review_model(phone: bool, width: int, **kw) -> str:
    m = main_col('Your Fury coach', '', C.FURY_OPENING, C.FURY_WHY, C.FURY_DRILL, C.FURY_THEN, C.FURY_WELL, phone, **kw)
    return two_col(m, rail(phone), width, phone)


def review_rules(phone: bool, width: int, writing: bool = False, annot: bool = False, busy: bool = False,
                 swapped: bool = False) -> str:
    then = [(t, d, d) for t, d, _ in C.FURY_THEN]
    right = f'<span class="dot"></span>{C.WRITING}' if writing else ''
    m = main_col('Your Fury review', right, C.RULES_OPENING.format(boss=BOSS), C.FURY_DETAIL, C.FURY_DRILL,
                 then, C.FURY_WELL, phone, reserve=writing, annot=annot, busy=busy, link=C.FURY_LINK)
    return two_col(m, rail(phone), width, phone)


# Panel states other than a review.

def msg_panel(title: str, body: str, phone: bool, right: str = '', min_h: int = 0, sub: str = '') -> str:
    style = f'min-height:{min_h}px' if min_h else ''
    return f'<section class="cp{" ph" if phone else ""}" style="{style}">{header(title, right, sub)}{body}</section>'


def p(text: str, color: str = '#e9e4d8', top: int = 12) -> str:
    return f'<p class="body15" style="color:{color};margin-top:{top}px">{text}</p>'


def loading_panel(phone: bool) -> str:
    lines = RESERVE_LINES['open'][1 if phone else 0]
    sk = ''.join('<div class="sk" style="width:%d%%"></div>' % w for w in ([96, 90, 62, 96, 84, 70, 40][:lines]))
    block = (f'<div class="fix" style="border-left-color:#262e40">{label("Fix first", MUTED)}'
             f'<div class="sk" style="height:28px;width:48%;margin:8px 0"></div>'
             + '<div class="sk" style="width:92%"></div>' * 3 + '</div>')
    return msg_panel('Your Fury coach', f'<div style="margin-top:8px">{sk}</div>{block}', phone,
                     right=f'<span class="dot"></span>{C.LOADING}', min_h=420 if phone else 520)


def teaser_panel(phone: bool, n: int = 3, first: bool = True) -> str:
    count = '1 thing to fix' if n == 1 else f'{n} things to fix'
    body = f'<div class="fix" style="margin-top:12px">{label(count, GOLD)}'
    if first:
        body += f'<div class="fix-t" style="margin-top:6px">{ico("Bloodthirst", 32)}<span>First: {C.FURY_TITLE}</span></div>'
    body += '</div>'
    body += '<div style="margin-top:16px"><a class="btn44">Read the full review with Premium</a></div>'
    return msg_panel('Your Fury coach', body, phone)


def clean_panel(phone: bool) -> str:
    body = (f'<div style="margin-top:8px">{slot(C.RULES_NOTHING, "open", phone, False, False)}</div>'
            f'<div class="well">{label("Doing well", KILL)}<ul style="list-style:none;margin:6px 0 0;padding:0">'
            + ''.join(f'<li style="display:flex;gap:8px;font-size:14px;line-height:22px;padding:4px 0">{TICK}<span>{t}</span></li>'
                      for t in C.FURY_WELL) + '</ul></div>')
    return msg_panel('Your Fury review', body, phone)


def not_reviewable_panel(phone: bool) -> str:
    body = (p('This pull is too short to review. The coach reads boss pulls long enough to show a rotation.')
            + f'<ul class="cc" style="margin-top:12px"><li>We couldn&#39;t tell this player&#39;s spec from the log.</li></ul>')
    return msg_panel('Your review', body, phone)


def preparing_panel(phone: bool) -> str:
    body = p('Top Fury logs for this boss are still loading. Check again in a few minutes.') + \
        '<div style="margin-top:16px"><a class="btn44 sec">Check again</a></div>'
    return msg_panel('Your Fury coach', body, phone)


def error_panel(phone: bool) -> str:
    body = p(C.ERROR) + '<div style="margin-top:16px"><a class="btn44 sec">Try again</a></div>'
    return msg_panel('Your Fury coach', body, phone)


def plain_panel(title: str, text: str, phone: bool, button: str = '') -> str:
    btn = f'<div style="margin-top:16px"><a class="btn44 sec">{button}</a></div>' if button else ''
    return msg_panel(title, p(text) + btn, phone)


def officer_panel(phone: bool, width: int) -> str:
    then = [(t, d, w) for t, d, w in C.FURY_THEN]
    pl = pill('Officer view', 'site')
    m = main_col("Tyrim's Fury review", pl, C.FURY_OPENING, C.FURY_WHY, C.FURY_DRILL, then, C.FURY_WELL, phone,
                 sub='Written to Tyrim. You see it as an officer of the guild.')
    return two_col(m, rail(phone, history=False), width, phone)


def starter_panel(phone: bool) -> str:
    body = (f'<p class="note" style="margin-top:2px">{C.STARTER_HINT.format(spec="Enhancement")}</p>'
            + p('…the rest as in a full review.', MUTED, 12))
    return msg_panel('Your Enhancement coach · early version', body, phone)


def healer_panel(phone: bool, width: int) -> str:
    chips = ''.join([chip('', 'Death', '3:20', phone), chip('', 'Resources', '3:05', phone)])
    fix = (f'<div class="fix">{label("Fix first", GOLD)}<div class="fix-t"><span>Died at 3:20 to Flame Breath</span></div>'
           + slot('A dead healer heals nothing for the rest of the pull, and Flame Breath is damage you can step out of. '
                  'You were also out of mana when it landed, so there was no answer left.', 'why', phone, False, False)
           + f'<div class="drill">{label("Practice")}'
           + slot('Before the pull, pick your spot at her side, never her front, and check it each time she lands.',
                  'drill', phone, False, False, 14) + f'</div><div class="chips">{chips}</div></div>')
    then = (f'<div class="then">{label("Then")}'
            f'<div class="titem"><h4><span>Out of mana from 3:05</span></h4><p class="body15" style="font-size:14px;margin-top:4px">'
            f'The mana ran out 15 s before the death; the fix is upstream, in how you spend it.</p></div>'
            f'<div class="titem"><h4>{ico("Greater Heal", 24)}<span>Greater Heal (Rank 4): 44% overheal</span></h4>'
            f'<p class="body15" style="font-size:14px;margin-top:4px">Top Holy priests here overheal 22% with this rank. '
            f'A lower rank costs less mana, and that is where your last 15 s went.</p></div></div>')
    well = (f'<div class="well">{label("Doing well", KILL)}<ul style="list-style:none;margin:6px 0 0;padding:0">'
            + ''.join(f'<li style="display:flex;gap:8px;font-size:14px;line-height:22px;padding:4px 0">{TICK}<span>{t}</span></li>'
                      for t in ['Flask of Distilled Wisdom at the pull.', 'Inner Fire up 100% of the fight.'])
            + '</ul></div>')
    body = (f'<div style="margin-top:8px">{slot("Your preparation was right: flask at the pull and Inner Fire up 100% of the fight. This one is about lasting to the end.", "open", phone, False, False)}</div>'
            + fix + then + well)
    main = f'<section class="cp{" ph" if phone else ""}">{header("Your Holy coach", "")}{body}</section>'
    return two_col(main, rail(phone, cannot=C.FURY_CANNOT, foot='Reviewed 10 Dec · Based on: Top 21 Holy priests on Onyxia (40), your log.', who=False), width, phone)


# The row line.

def cl(inner: str, phone: bool, mod: str = '') -> str:
    return f'<div class="cline{" ph" if phone else ""} {mod}" style="{"min-height:52px" if phone else ""}">{inner}</div>'


def row_line(state: str, phone: bool, mod: str = '') -> str:
    ic = ico('Bloodthirst', 20)
    chev = '<span class="chev">›</span>'
    if state == 'premium':
        return cl(f'<span class="lead">Fix first ·</span>{ic}<b>{C.FURY_TITLE}</b>{chev}', phone, mod)
    if state == 'premium-none':
        return cl(f'<span class="lead">3 things to fix · Open your coach&#39;s review</span>{chev}', phone, mod)
    if state == 'clean':
        return cl(f'<span class="lead">{C.ROW_CLEAN}</span>{chev}', phone, mod)
    if state == 'free':
        up = '<u style="color:#e5b955;text-underline-offset:3px">Read the full review with Premium</u>'
        if phone:
            return (cl(f'<span class="lead">3 things to fix · First: <b style="color:#f2eee4">{C.FURY_TITLE}</b></span>{chev}', phone, mod)
                    + cl(up, phone, ''))
        return cl(f'<span class="lead">3 things to fix · First: <b>{C.FURY_TITLE}</b> · </span>{up}', phone, mod)
    if state == 'free-none':
        return cl(f'<span class="lead">3 things to fix · </span><u style="color:#e5b955;text-underline-offset:3px">'
                  f'Read the full review with Premium</u>', phone, mod)
    if state == 'loading':
        return cl('<span class="bar-sk"></span>', phone, 'muted')
    if state == 'preparing':
        return cl(f'<span class="lead">{C.ROW_PREPARING.format(spec="Fury")}</span>', phone, 'muted')
    if state == 'error':
        return cl(f'<span class="lead">{C.ROW_ERROR} </span><u style="text-underline-offset:3px">Try again</u>', phone, 'muted')
    raise ValueError(state)


def officer_detail(phone: bool) -> str:
    return (f'<div style="background:#131824;padding:12px 8px 12px {"8px" if phone else "100px"};display:flex;gap:12px;align-items:center">'
            f'<a class="btn44 sec" style="font-size:12px;letter-spacing:.06em;text-transform:uppercase;font-weight:700">Coach review</a>'
            f'<span class="note">Abilities, targets and the split follow.</span></div>')


def row_list(line: str, phone: bool, detail: str = '', own: bool = True, rows: int = 6) -> str:
    out = ''
    for r in ROWS[:rows]:
        mine = r[2] == OWNED and own
        out += arow(r, phone, line if mine else '', detail if (mine and detail) else '')
    return f'<ul style="list-style:none;margin:0;padding:0">{out}</ul>'


def row_states(phone: bool) -> str:
    w = 354 if phone else 640
    out = [cell('Premium owner, ready', row_list(row_line('premium', phone), phone, rows=3), w),
           cell('Hover (desktop) / pressed (touch)', row_list(row_line('premium', phone, 'hover'), phone, rows=3), w),
           cell('Focus-visible', row_list(row_line('premium', phone, 'focus'), phone, rows=3), w),
           cell('Free owner', row_list(row_line('free', phone), phone, rows=3), w),
           cell('Free owner, findings but no first', row_list(row_line('free-none', phone), phone, rows=3), w),
           cell('Premium owner, findings but no first', row_list(row_line('premium-none', phone), phone, rows=3), w),
           cell('Premium owner, clean pull', row_list(row_line('clean', phone), phone, rows=3), w),
           cell('Loading (slot reserved)', row_list(row_line('loading', phone), phone, rows=3), w),
           cell('Preparing', row_list(row_line('preparing', phone), phone, rows=3), w),
           cell('Error', row_list(row_line('error', phone), phone, rows=3), w),
           cell("Officer, someone else's row (no fetch, no line; the button is in the opened row)",
                row_list('', phone, officer_detail(phone), rows=3), w),
           cell("Everyone else (signed out, not the owner, not an officer): no line, no slot",
                row_list('', phone, rows=3), w)]
    return '<div style="display:flex;flex-wrap:wrap;gap:24px">' + ''.join(out) + '</div>'


# The account page.

def lr_parts(parts: list) -> str:
    return f'<span class="m"> · </span>'.join(parts)


def lr_line(state: str, phone: bool) -> str:
    last = '<span class="m">Last raid</span>'
    zone = '<b>Molten Core, Thu 23 Oct</b>'
    open_ = '<u>Open review</u>'
    up = '<u>Read the full review with Premium</u>'
    if state == 'premium':
        body = lr_parts([last, zone, 'Fix first: Bloodthirst casts lost, on 5 of 8 bosses', open_])
    elif state == 'free':
        body = lr_parts([last, zone, '11 things to fix', 'First: Bloodthirst casts lost', up])
    elif state == 'pending':
        body = lr_parts([last, zone, '<span class="m">Reading your pulls…</span>'])
    elif state == 'clean':
        body = lr_parts([last, zone, 'Nothing to fix that the log can show', open_])
    elif state == 'none':
        body = '<span class="m">Last raid · no log in the last 14 days</span>'
    elif state == 'loading':
        body = '<span class="bar-sk" style="display:inline-block;height:12px;width:340px;max-width:70%;background:#131824;border-radius:3px"></span>'
    elif state == 'error':
        body = '<span class="m">Your last raid isn&#39;t available right now. </span><u>Try again</u>'
    else:
        raise ValueError(state)
    return f'<div class="lrline{" ph" if phone else ""}"><span style="min-width:0">{body}</span></div>'


CHARS = [('warrior', 'Tyrim', 'Level 60 Fury Warrior · Olympus XXVII', ['Main', 'Current']),
         ('hunter', 'Zulmara', 'Level 60 Marksmanship Hunter · Olympus XXVII', []),
         ('priest', 'Grimvow', 'Level 58 Shadow Priest', [])]


def char_panel(line: str, phone: bool) -> str:
    rows = ''
    for cls, name, desc, pills in CHARS:
        pl = ''.join(pill(t, 'site') for t in pills)
        rows += (f'<div class="crow">{crest_img(cls, 40)}<div style="min-width:0;flex:1"><div style="color:{CLASS_COLOR[cls]};'
                 f'font-weight:600;font-size:15px">{name}</div><div class="note">{desc}</div></div>'
                 f'<span style="display:flex;gap:6px">{pl}</span></div>')
    aside = '' if phone else '<span class="note">Imported from Battle.net 3 days ago. <u>Refresh from Battle.net</u></span>'
    return (f'<section class="cp{" ph" if phone else ""}"><div class="cp-h"><span class="label" style="color:{GOLD}">Characters</span>'
            f'{aside}</div>{line}{rows}</section>')


def account_states(phone: bool) -> str:
    w = 354 if phone else 640
    items = [('Premium', 'premium'), ('Free', 'free'), ('Pending (cold cache)', 'pending'), ('Clean night', 'clean'),
             ('No raid in 14 days', 'none'), ('Loading (slot reserved)', 'loading'), ('Error', 'error')]
    return '<div style="display:flex;flex-wrap:wrap;gap:24px">' + ''.join(
        cell(t, char_panel(lr_line(s, phone), phone), w) for t, s in items) + '</div>'


# Discord.

LINK = 'style="color:#00a8fc"'
VIEWER = dict([('class', 'warrior'), ('faction', 'alliance'), ('battletag', 'Tyrim')])


def coach_field(ph: bool, pending: bool = False, named: bool = False) -> str:
    who = '<b>Zulmara</b> · ' if named else ''
    if pending:
        return who + f'<span {LINK}>Your coach&#39;s review is on the site.</span>'
    return (who + f'Fix first: Bloodthirst casts lost — on 5 of 8 bosses. '
            f'<span {LINK}>Open your coach&#39;s review</span>')


def report_answer(ph: bool, mode: str) -> str:
    fields = [('Kill order', '1. Onyxia <b>1:42</b> after 3 wipes', False)]
    if mode == 'named':
        second = ('<br><b>Grimvow</b> · Fix first: Casts off the rotation — on 3 of 8 bosses. '
                  f'<span {LINK}>Open your coach&#39;s review</span>')
        fields.append(('Your coach', coach_field(ph, False, True) + second, False))
    elif mode != 'none':
        fields.append(('Your coach', coach_field(ph, mode == 'pending'), False))
    fields.append(('Loot', '6 items. The first 3:', False))
    main = D.embed('Report: ' + D.RAID_TITLE, D.KILL, D.NIGHT, fields=tuple(fields), phone=ph)
    loot = ''.join(D.item_embed(n, '', 'to ' + who + D.DOT + 'Onyxia', ph, foot=D.last_foot(i, len(D.LOOT), D.POST_FOOT))
                   for i, (n, who) in enumerate(D.LOOT))
    return D.message(D.embeds(main + loot, ph) + D.buttons(D.btn(D.SHARE), D.btn('Open the report', 'link')), ph,
                     used=(D.ZUL, '/report'), ephemeral=True)


def night_post(ph: bool, coach: bool) -> str:
    desc = 'Kill times, loot and the biggest changes on Forever Sixty.' + (' ' + C.RAID_NIGHT if coach else '')
    fields = (('Kills', 'Onyxia 1:42', True), ('Wipes', '3', True), ('Came', '22 of 23', True), ('Loot', '6 items', True))
    e = D.embed('Report: ' + D.RAID_TITLE, D.KILL, desc, fields=fields, phone=ph, foot=D.POST_FOOT)
    return D.message(D.embeds(e, ph), ph)


def pair(title: str, build) -> str:
    return (f'<div class="cell"><span class="cell-t">{title}</span>'
            f'<div style="display:flex;gap:24px;align-items:flex-start">{build(False)}{build(True)}</div></div>')


def discord_board() -> str:
    return ''.join([
        pair('/report, ephemeral: one character, a Premium viewer', lambda ph: report_answer(ph, 'one')),
        pair('/report, ephemeral: two characters (named lines, three at most)', lambda ph: report_answer(ph, 'named')),
        pair("/report, ephemeral: packs not cached within the interaction's budget", lambda ph: report_answer(ph, 'pending')),
        pair('/report, ephemeral: a free viewer, or a Premium viewer owning nothing in the log: no field',
             lambda ph: report_answer(ph, 'none')),
        pair('report_ready, in the guild channel: at least one raider holds log_coach (names and counts nobody)',
             lambda ph: night_post(ph, True)),
        pair('report_ready, in the guild channel: no raider holds log_coach', lambda ph: night_post(ph, False)),
    ])


# Boards.

def sheet(name: str, title: str, body: str, width: int) -> str:
    gutter = 18 if width < 500 else 48
    content = (f'<div style="display:flex;flex-direction:column;gap:24px;padding:32px {gutter}px">'
               f'<h1 class="display" style="margin:0;font-size:22px;color:#f2eee4">{title}</h1>{body}'
               f'<p class="ann">{CAPTION}</p></div>')
    return write(name + '.html', page(G1_CSS + D.SITE_CSS + D.DCSS + CSS + content, width)).as_posix()


def site_nav(phone: bool) -> str:
    return nav_phone(True, VIEWER) if phone else nav('Logs', True, VIEWER)


def board(name: str, inner: str, width: int, phone: bool, active: str = 'Coach') -> str:
    body = (site_nav(phone) + chrome(inner + f'<p class="ann" style="margin-top:20px">{CAPTION}</p>', phone, active)
            + footer(phone=phone))
    return write(name + '.html', page(G1_CSS + D.SITE_CSS + CSS + body, width)).as_posix()


def row_board(name: str, phone: bool, width: int) -> str:
    body = (f'<div style="margin-bottom:12px">{label("Combatants · Damage Done")}</div>'
            + row_list(row_line('premium', phone), phone))
    return board(name, body, width, phone, 'Damage Done')


def panel_states(phone: bool) -> str:
    w = 354 if phone else 640
    cells = [
        cell('Loading (Reading the pull…): the panel keeps its minimum height', loading_panel(phone), w),
        cell('Writing: the rule-only rendering in the reserved slots (dashed boxes show the reserved heights)',
             review_rules(phone, 640, writing=True, annot=True), w),
        cell('Writing, then swapped: the model review in the same slots. Titles and the rail do not move',
             review_model(phone, 640, reserve=True), w),
        cell('Busy: the rule-only rendering with one quiet line above it', review_rules(phone, 640, busy=True), w),
        cell('Preparing', preparing_panel(phone), w),
        cell('No findings', clean_panel(phone), w),
        cell('Not reviewable', not_reviewable_panel(phone), w),
        cell('Free teaser (count and the first title; nothing else is sent)', teaser_panel(phone), w),
        cell('Free teaser, one finding', teaser_panel(phone, 1), w),
        cell('Officer view', officer_panel(phone, 640), w),
        cell('Starter persona (header and hint; the body is a full review)', starter_panel(phone), w),
        cell('Healer: efficiency and survival, never "heal more"', healer_panel(phone, 640), w),
        cell('Error', error_panel(phone), w),
        cell('Signed out, from a Discord link',
             plain_panel('Your coach', "Sign in to read your coach's review.", phone, 'Sign in'), w),
        cell('Not available on this account',
             plain_panel('Your coach', "This review isn't available on this account.", phone), w),
        cell('Whole night selected',
             plain_panel('Your coach', "Coach reviews cover one pull at a time. Pick a boss pull from the list.", phone), w),
    ]
    return '<div style="display:flex;flex-wrap:wrap;gap:24px">' + ''.join(cells) + '</div>'


def main() -> None:
    out = [board('log-coach-panel', review_model(False, 1440), 1440, False),
           board('log-coach-panel-1024', review_model(False, 1024), 1024, False),
           board('log-coach-panel-2000', review_model(False, 2000), 2000, False),
           board('log-coach-panel-phone', review_model(True, 390), 390, True),
           row_board('log-coach-row', False, 1440),
           row_board('log-coach-row-phone', True, 390),
           sheet('log-coach-row-states', 'Coach line on the combatant list: states', row_states(False), 1440),
           sheet('log-coach-states', 'Coach panel: states', panel_states(False), 1440),
           sheet('log-coach-states-phone', 'Coach panel: states, phone', panel_states(True), 390),
           sheet('log-coach-account', 'Account: the last-raid line, states', account_states(False), 1440),
           sheet('log-coach-account-phone', 'Account: the last-raid line, phone', account_states(True), 390),
           sheet('log-coach-discord', 'Discord: /report and the raid-night post', discord_board(), 1440)]
    print('\n'.join(out))


if __name__ == '__main__':
    main()
