"""Round 5 mobile: the iOS and Android apps in the web's round-5 look (design/specs/2026-11-09-directions.md, Mobile).

The flows, the content and every element are the app-flows spec's (design/specs/2026-11-08-app-flows.md, its tap table
and purpose audit): the data is imported from its mock (design/mocks/ws_app_flows.py) so both stay one source. Only the
visual language changes: Blizzard's Forever palette, Open Sans and the display face, the rounded frame on the lead
panel, the owner's art and the flat seal. Each platform keeps its own conventions: iOS navigation and tab bars, sheets
with a grabber, the lock screen Live Activity; Android Material 3 top app bar, navigation bar with the pill indicator,
ModalNavigationDrawer, modal bottom sheet, segmented buttons and the ongoing notification.

    python3 design/mocks/directions/mobile.py      -> html/m-<platform>-<screen>.html
    node design/mocks/directions/shoot_mobile.cjs <node_modules>   -> png/m-<platform>-<screen>.png
"""
import base64
import sys
from pathlib import Path

import data as D
from art import GROUND_TILE_PX, band_uri, cave_uri, ground_uri, hero_uri
from assets import flat_logo_uri
from common import write
from nav_icons_check import nav_svg
from style import CORNERS, DISPLAY_FONT, FONTS, FRAME_LINE, PAGE_QUALITY, R_LG, R_SM, TOKENS, medallion

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
import ws_app_flows as F  # noqa: E402  the app-flows mock's own data: the screens' content, single-sourced
from mocklib import emblem  # noqa: E402  the site's faction mark (web/public/icons/hd/faction)

DEVICES = {'ios': (393, 852), 'android': (412, 915)}  # iPhone 15, Pixel 7
AVATARS = Path(__file__).resolve().parents[1] / 'data/avatars'
SCROLL_SCREENS = ('home', 'raid', 'gear', 'gear-frame', 'activity', 'signup')
GATED = {'Greenhammer'}  # unlocks with the first raids (data.GATED_RAIDS_1): left out, as post-fix data would
OFF_HAND_WORN = ('Off hand', 'Persuader', 'inv_hammer_08', 4, 'Blacksmithing', None)
UNREAD = 3
TAB_ICON_PX = 24  # 24 pt on iOS, 24 dp on Android


def avatar_uri(name: str) -> str:
    return 'data:image/jpeg;base64,' + base64.b64encode((AVATARS / f'{name}.jpg').read_bytes()).decode()


def ic(name: str, px: int = 40, quality: int = 0) -> str:
    q = f' q{quality}' if quality else ''
    return f'<img class="ic{q}" src="{D.icon_uri(name)}" alt="" width="{px}" height="{px}">'


# ------------------------------------------------------------------------------------------------ the look
def _tokens() -> str:
    q = ';'.join(f'--q-{k}:{v}' for k, v in PAGE_QUALITY.items())
    return ':root{' + ';'.join(f'--{k}:{v}' for k, v in TOKENS.items()) + f';{q};--r-lg:{R_LG}px;--r-sm:{R_SM}px}}'


CSS = _tokens() + '''
:root{--title:''' + DISPLAY_FONT + ''';--body:'Open Sans',Arial,sans-serif}
*{box-sizing:border-box}html,body{margin:0}
body{background:#111 ;font:15px/1.45 var(--body);color:var(--text)}
.dev{position:relative;overflow:hidden;background:var(--earth) url(''' + ground_uri() + ''') 0 0/''' + str(GROUND_TILE_PX) + '''px repeat}
a{color:var(--gold);text-decoration:none;font-weight:600}
h1,h2,h3{font-family:var(--title);font-weight:400;letter-spacing:-.015em;color:var(--cream);margin:0;text-shadow:0 1px 2px rgba(0,0,0,.45)}
p,ul,ol{margin:0;padding:0;list-style:none}.num{font-variant-numeric:tabular-nums lining-nums}.muted{color:var(--muted)}
.ic{display:block;border:1px solid #000;box-shadow:0 0 0 1px var(--rule);border-radius:3px;flex:none}
.ic.q4{box-shadow:0 0 0 1px #6b3d8f}.ic.q3{box-shadow:0 0 0 1px #2c5b8c}
.crest{display:block;border-radius:50%;flex:none}
.medal{position:relative;flex:none}.medal img{position:absolute;border-radius:50%;object-fit:cover}.medal svg{position:absolute;inset:0}
/* bars */
.status{height:var(--sb);display:flex;align-items:center;justify-content:space-between;padding:0 26px 0 32px;color:#fff;font:600 15px var(--body);position:relative;z-index:2}
.island{position:absolute;left:50%;top:11px;width:124px;height:36px;margin-left:-62px;background:#000;border-radius:20px}
.sys{display:flex;gap:6px;align-items:center}.sys i{display:block;width:17px;height:11px;border-radius:2px;background:#fff}.sys i.b{width:25px;border-radius:3px}
.navb{height:var(--nb);display:flex;align-items:center;gap:12px;padding:0 16px;position:relative;z-index:2}
.navb .t{font:600 17px var(--body);color:#fff}.navb .sp{margin-left:auto;display:flex;gap:18px;color:#fff}
.pf{position:relative;width:36px;height:36px;flex:none}.pf img.av{width:32px;height:32px;border-radius:50%;position:absolute;inset:2px;box-shadow:0 0 0 2px var(--ring,#2f6fd6)}
.pf img.em{position:absolute;right:-3px;top:-3px;width:15px;height:15px;border-radius:50%;box-shadow:0 0 0 1px #0b2f3b;background:#0b2f3b}
.tabbar{display:flex;justify-content:space-around;background:var(--nav-2);border-top:1px solid rgba(255,255,255,.08);padding:6px 0 var(--hi)}
.tab{position:relative;display:flex;flex-direction:column;align-items:center;gap:3px;width:25%;font:600 11px var(--body);color:rgba(255,255,255,.82)}
.tab svg{width:24px;height:24px}.tab.on{color:#ffd36b}
.tab .bd{position:absolute;top:-4px;left:60%;min-width:18px;height:18px;padding:0 5px;border-radius:9px;background:#e0533f;color:#fff;font:700 11px/18px var(--body);text-align:center}
.homebar{position:absolute;left:50%;bottom:8px;width:134px;height:5px;margin-left:-67px;border-radius:3px;background:#fff;opacity:.85;z-index:5}
/* the band behind the bars: the owner's art (Home) or the painted blue band (the other tabs) */
.band{background:linear-gradient(rgba(5,58,84,.38),rgba(5,58,84,.38)),var(--band) url(''' + band_uri() + ''') center 30%/cover no-repeat;border-bottom:1px solid #000}
.band.hero{background:linear-gradient(rgba(7,29,46,.22),rgba(7,29,46,.22)),url(''' + hero_uri() + ''') 74% 60%/auto 100% no-repeat,var(--night)}
.head{display:flex;align-items:center;gap:12px;padding:6px 16px 18px}.head h1{font-size:30px;line-height:34px}.head p{font-size:14px;color:var(--text)}
.bigt{padding:4px 16px 14px}.bigt h1{font-size:32px;line-height:38px}
.body{padding:16px 14px 22px;display:flex;flex-direction:column;gap:20px}
.card{position:relative;background:var(--card);border:1px solid #000;box-shadow:0 0 0 1px rgba(130,112,91,.55);border-radius:var(--r-lg);overflow:hidden}
.pad{padding:16px}
.framed::before{content:"";position:absolute;inset:6px;border:1px solid ''' + FRAME_LINE + ''';border-radius:calc(var(--r-lg) - 4px);pointer-events:none;z-index:1}
.framed.ornate::after{content:"";position:absolute;inset:4px;background:''' + CORNERS + ''';pointer-events:none;z-index:1}
.sh{display:flex;align-items:baseline;justify-content:space-between;margin-bottom:8px}.sh h2{font-size:21px}.sh a{font-size:14px}
.row{display:flex;align-items:center;gap:12px;padding:9px 0;border-bottom:1px solid rgba(255,255,255,.09)}.row:last-child{border-bottom:0}
.row .tx{min-width:0;flex:1}.row b{display:block;color:#fff;font-weight:600;line-height:19px}.row .s{display:block;font-size:13px;color:var(--muted);line-height:18px}
.row .g{font:700 14px var(--body);color:#fff;white-space:nowrap}.row .g small{font-weight:600;color:var(--muted);font-size:11px;margin-left:2px}
.row .n{font-weight:600;display:block;line-height:19px}
.dot{position:relative}.dot::before{content:"";position:absolute;left:-60px;top:50%;width:5px;height:5px;margin-top:-3px;border-radius:50%;background:var(--gold)}
/* the raid card: the cave behind it, the ornate frame */
.raid{padding:18px 18px 16px;background:linear-gradient(rgba(14,9,6,.5),rgba(14,9,6,.5)),var(--cave) 62% 38%/auto 150% no-repeat,var(--card)}
.raid h2{font-size:24px}.raid .when{font-size:13px;color:var(--text)}
.count{display:flex;gap:18px;margin-top:10px}.count b{display:block;font:600 36px/1 var(--body);color:#fff}.count span{font-size:12px;color:var(--muted)}
.ans{margin-top:12px;font-size:13px;color:var(--text)}
.answers{display:grid;grid-template-columns:repeat(4,1fr);gap:8px;margin-top:8px}
.answers a{display:flex;align-items:center;justify-content:center;height:44px;border-radius:var(--r-sm);background:var(--blue-btn);border:1px solid #1d6a96;color:#fff;font-weight:600;font-size:14px}
.answers a.on{background:var(--teal);border-color:var(--teal-edge)}
.roles{display:grid;grid-template-columns:repeat(3,1fr);gap:10px;margin-top:14px;font-size:13px}
.roles div{display:flex;justify-content:space-between;color:#fff;font-weight:600}.roles .v.ok{color:#7fd48a}.roles .v.lo{color:#f0736b}
.meter{grid-column:span 1;height:5px;border-radius:3px;background:rgba(255,255,255,.12);margin-top:5px;overflow:hidden}.meter i{display:block;height:100%;background:var(--gold)}
.rhead .rl{display:flex;align-items:center;gap:12px;margin-bottom:12px}.rhead .rl b{color:#fff;font-weight:600;display:block}.rhead .rl .s{display:block;font-size:13px;line-height:18px;color:var(--muted)}
.achip{margin-left:auto;padding:4px 10px;border-radius:var(--r-sm);background:var(--teal);border:1px solid var(--teal-edge);color:#fff;font-weight:700;font-size:13px;white-space:nowrap}
.num.ok{color:#7fd48a}.num.lo{color:#f0736b}.mk{display:flex;align-items:center;justify-content:center;width:22px;height:22px;border-radius:50%;color:#111;font-weight:700;font-size:13px;flex:none}
.stamp{font-size:12px;color:var(--muted);margin-top:8px}
.pull{margin-top:14px;padding-top:12px;border-top:1px solid rgba(255,255,255,.12)}.pull b{color:#fff;font-weight:600}.pull p{font-size:13px;color:var(--text)}
/* android: Material 3 conventions in the same look */
.android .status{padding:0 18px;font-size:14px}.android .island{display:none}
.android .navb .t{font:400 22px var(--title)}
.android .tabbar{padding:12px 0 var(--hi);height:auto}.android .tab{gap:4px;font-size:12px}
.android .tab .pill{display:flex;align-items:center;justify-content:center;width:64px;height:32px;border-radius:16px}
.android .tab.on .pill{background:rgba(255,236,113,.18)}.android .tab .bd{top:-4px;left:63%}
.android .answers{gap:0;border:1px solid var(--rule);border-radius:24px;overflow:hidden}
.android .answers a{height:48px;border-radius:0;border:0;border-right:1px solid var(--rule);background:transparent}.android .answers a:last-child{border-right:0}
.android .answers a.on{background:var(--teal-edge);color:#fff}
/* overlays */
.scrim{position:absolute;inset:0;background:rgba(0,0,0,.62);z-index:6}
.drawer{position:absolute;top:0;bottom:0;left:0;width:85%;background:var(--card);z-index:7;overflow:hidden;border-right:1px solid #000}
.android .drawer{width:336px;border-radius:0 16px 16px 0;top:0}
.drawer .you{padding:calc(var(--sb) + 18px) 18px 16px;background:linear-gradient(rgba(5,58,84,.3),rgba(5,58,84,.3)),url(''' + band_uri() + ''') center 20%/cover}
.drawer .you .pf{width:60px;height:60px}.drawer .you .pf img.av{width:56px;height:56px}.drawer .you .pf img.em{width:22px;height:22px}
.drawer h3{font-size:16px;margin:0 18px;padding:14px 0 6px;color:var(--cream)}
.drawer .rows{padding:0 18px}.drawer .row{padding:7px 0}
.tag{display:inline-block;margin-top:5px;padding:1px 8px;border:1px solid var(--gold);border-radius:var(--r-sm);color:var(--gold);font-size:12px;font-weight:600}
.sheet{position:absolute;left:0;right:0;bottom:0;z-index:7;background:var(--card);border-radius:14px 14px 0 0;border-top:1px solid #000;box-shadow:0 -1px 0 rgba(130,112,91,.55);padding:8px 16px calc(var(--hi) + 18px)}
.android .sheet{border-radius:28px 28px 0 0}
.grab{width:36px;height:5px;border-radius:3px;background:rgba(255,255,255,.35);margin:2px auto 12px}.android .grab{width:32px;height:4px;margin:14px auto 18px}
.chips{display:flex;gap:8px;margin-top:10px}.raid .chips a{display:flex;align-items:center;height:44px;padding:0 16px}.android .raid .chips a{height:48px}.chips.scroll{overflow:hidden;white-space:nowrap;margin-right:-16px}.chips.scroll a{flex:none}.android .bigt{padding-top:0}.chips a{padding:7px 14px;border-radius:var(--r-sm);border:1px solid var(--rule);color:#fff;font-weight:600;font-size:14px}
.chips a.on{background:var(--teal);border-color:var(--teal-edge)}.android .chips a{border-radius:8px}
.tip{background:#060914;border:1px solid #8b8f9c;border-radius:4px;box-shadow:0 0 0 1px #000,inset 0 0 0 1px #2b3040;padding:9px 11px;font:13px/1.45 var(--body);color:#fff}
.tip .q4{color:#a335ee}.tip .gr{color:#1eff00}
/* lock screen */
/* the wallpaper is the user's, never ours (ai-slop-reviewer): a plain neutral system wallpaper stands in */
.lock{background:#2a2f36}
.clock{text-align:center;color:#fff;margin-top:34px}.clock .d{font:600 18px var(--body);opacity:.9}.clock .t{font:300 84px/1 var(--body);letter-spacing:-.02em}
.android .clock{text-align:left;padding:0 26px}.android .clock .t{font-size:72px}
.la{margin:22px 12px 0;padding:14px;border-radius:22px;background:rgba(10,14,20,.78);color:#fff}
.android .la{border-radius:24px;background:rgba(23,14,9,.92)}
.la .top{display:flex;align-items:center;gap:10px}.la b{font-weight:700}.la .cd{margin-left:auto;font:700 26px var(--body);color:var(--gold)}
.la .ln{font-size:13px;color:rgba(255,255,255,.75);margin-top:8px}
.la .acts{display:grid;grid-template-columns:repeat(4,1fr);gap:6px;margin-top:12px}
.la .acts a{display:flex;align-items:center;justify-content:center;height:34px;border-radius:17px;background:rgba(255,255,255,.14);color:#fff;font-size:13px}
.la .acts a.on{background:var(--teal-edge)}
.app{display:flex;align-items:center;gap:6px;font-size:12px;color:rgba(255,255,255,.7);margin-bottom:8px}.app img{width:16px;height:16px;border-radius:4px}
'''

TAB_ICONS = {
    'Home': '<path d="M4 11 13 4l9 7v10h-6v-6h-6v6H4z" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linejoin="round"/>',
    'Raid': '<path d="M6 22V4m0 1h12l-2.5 4L18 13H6" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linejoin="round"/>',
    'Gear': '<path d="M13 3l8 3v6c0 5-3.5 8-8 10-4.5-2-8-5-8-10V6z" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linejoin="round"/>'
            '<path d="M9.5 12.5l2.5 2.5 4.5-5" fill="none" stroke="currentColor" stroke-width="1.8"/>',
    'Activity': '<path d="M7 17V11a6 6 0 0 1 12 0v6l2 2H5zM11 21h4" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linejoin="round"/>',
}
BADGES = {'Raid': 1, 'Activity': UNREAD}
SEARCH = ('<svg width="22" height="22" viewBox="0 0 22 22"><circle cx="9.5" cy="9.5" r="6.5" fill="none" stroke="currentColor" stroke-width="1.9"/>'
          '<path d="M14.5 14.5 20 20" stroke="currentColor" stroke-width="1.9"/></svg>')


def status(platform: str) -> str:
    island = '<span class="island"></span>' if platform == 'ios' else ''
    return f'<div class="status"><span>9:41</span>{island}<span class="sys"><i></i><i></i><i class="b"></i></span></div>'


def profile(avatar: str = 'thandrel', faction: str = 'alliance') -> str:
    ring = {'alliance': '#2f6fd6', 'horde': '#c0392b'}[faction]
    return f'<span class="pf" style="--ring:{ring}"><img class="av" src="{avatar_uri(avatar)}" alt=""><img class="em" src="{emblem(faction)}" alt=""></span>'


def glyph_profile(faction: str) -> str:
    """No portrait yet: the neutral person glyph in the faction ring with the badge, never the class crest (spec 3)."""
    ring = {'alliance': '#2f6fd6', 'horde': '#c0392b'}[faction]
    glyph = ('<svg viewBox="0 0 32 32" style="position:absolute;inset:2px;width:32px;height:32px;border-radius:50%;background:#1d2a33;'
             f'box-shadow:0 0 0 2px {ring}"><circle cx="16" cy="12" r="5" fill="none" stroke="#b9c4cc" stroke-width="1.6"/>'
             '<path d="M7 26c1.5-5 5-7 9-7s7.5 2 9 7" fill="none" stroke="#b9c4cc" stroke-width="1.6"/></svg>')
    return f'<span class="pf">{glyph}<img class="em" src="{emblem(faction)}" alt=""></span>'


def navbar(platform: str, title: str = '') -> str:
    """Tab roots: the profile button at the top left (spec 3), search on Home, Raid and Gear; Android titles the bar."""
    t = f'<span class="t">{title}</span>' if title and platform == 'android' else ''
    srch = f'<span class="sp">{SEARCH}</span>' if title in ('', 'Raid', 'Gear') else ''
    return f'<div class="navb">{profile()}{t}{srch}</div>'


def tabbar(platform: str, cur: str) -> str:
    cells = []
    for name in TAB_ICONS:
        svg = nav_svg(name.lower(), TAB_ICON_PX)  # the owner's traced icons (design/assets/icons/nav), tinted by currentColor
        if platform == 'android':
            svg = f'<span class="pill">{svg}</span>'
        badge = f'<span class="bd">{BADGES[name]}</span>' if name in BADGES else ''
        cells.append(f'<span class="tab{" on" if name == cur else ""}">{svg}{badge}{name}</span>')
    return f'<nav class="tabbar">{"".join(cells)}</nav>'


def device(platform: str, inner: str, scroll: bool) -> str:
    w, h = DEVICES[platform]
    bars = {'ios': (54, 44, 34), 'android': (32, 64, 16)}[platform]
    size = f'width:{w}px;min-height:{h}px' if scroll else f'width:{w}px;height:{h}px'
    homebar = '<span class="homebar"></span>'
    return (f'<div class="dev {platform}" style="{size};--sb:{bars[0]}px;--nb:{bars[1]}px;--hi:{bars[2] + 6}px;'
            f'--cave:url({cave_uri()})">{inner}{homebar}</div>')


def page(name: str, platform: str, inner: str, scroll: bool) -> None:
    html = (f'<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">'
            f'<title>{name}</title><link rel="stylesheet" href="https://fonts.googleapis.com/css2?{FONTS}&display=swap"><style>{CSS}</style></head>'
            f'<body class="{platform}">{device(platform, inner, scroll)}</body></html>')
    write(f'm-{platform}-{name}', html)


# ------------------------------------------------------------------------------------------------ shared blocks
ANSWER = 'In'  # one state across the set: Activity's "You're In for Onyxia" and the Live Activity's "You are In"
LATE_MINUTES = (15, 30, 60)
GUILD = 'Olympus XXVII'


def chars_in_guild(guild: str) -> int:
    return sum(guild in line for _, _, line, _, _, _ in F.CHARACTERS)


def answer_line(answer: str, late_min: int | None = None) -> str:
    """Owner ruling 2026-11-09: name the character and role only when the account has more than one character in the guild
    (Thandrel and Lightvow are both in Olympus XXVII here); otherwise the answer alone."""
    said = f'You are Late, {late_min} min' if late_min else (f'You are {answer}' if answer else 'You have not answered')
    return f'{said}, {F.AS_ME}' if chars_in_guild(GUILD) > 1 else said


def raid_card(answer: str = ANSWER, late_min: int | None = None) -> str:
    """The next raid on raid day (F1): the cave behind it, the ornate frame; the answers, the current one selected.
    Late opens the minutes under the answers, on the card (spec F1); the answer posts on the tap."""
    roles = ''.join(f'<div><span>{n}</span><span class="v {"ok" if a >= need else "lo"} num">{a}/{need}</span></div>' for _, n, a, need in F.ROLES)
    meters = ''.join(f'<div class="meter"><i style="width:{min(100, a / need * 100):.0f}%"></i></div>' for _, _, a, need in F.ROLES)
    answers = ''.join(f'<a class="{"on" if x == answer else ""}" href="#">{x}</a>' for x in ('In', 'Late', 'Bench', 'Out'))
    said = answer_line(answer, late_min)
    late = ('<div class="chips">' + ''.join(f'<a class="{"on" if m == late_min else ""}" href="#">{m} min</a>' for m in LATE_MINUTES) + '</div>'
            if answer == 'Late' else '')
    return (f'<section class="card framed ornate raid"><h2>{F.RAID}</h2><p class="when">Olympus XXVII &middot; {F.WHEN}</p>'
            f'<div class="count num"><div><b>02</b><span>hours</span></div><div><b>48</b><span>min</span></div></div>'
            f'<p class="ans">{said}</p><div class="answers">{answers}</div>{late}'
            f'<div class="roles">{roles}{meters}</div>'
            f'<div class="pull"><b>2 things before you pull</b><p>No flask in your bags; boots not enchanted &middot; from your last log-out, 9 h ago</p></div></section>')


def upgrades_block() -> str:
    """Next upgrades (F5): the top three from the gated home data, with the standing a reputation needs."""
    source = {'Stormpike Guard': 'Stormpike Guard, needs revered'}
    rows = ''.join(f'<div class="row">{ic(u["icon"], 40, 4 if u["quality"] == "epic" else 3)}<span class="tx"><span class="n" style="color:var(--q-{u["quality"]})">{u["name"]}</span>'
                   f'<span class="s">{D.SLOT_NAME[u["slot"]]} &middot; {source.get(u["source"], u["source"])}</span></span><span class="g num">+{u["gain"]:.1f}<small>DPS</small></span></div>'
                   for u in D.home('60')['upgrades'][:3])
    return f'<section><div class="sh"><h2>Next upgrades</h2><a href="#">Gear &rsaquo;</a></div><div class="card framed pad">{rows}</div></section>'


def activity_row(icon_name: str, title: str, line: str, unread: bool) -> str:
    return (f'<div class="row">{ic(icon_name, 40)}<span class="tx{" dot" if unread else ""}"><b>{title}</b><span class="s">{line}</span></span></div>')


def new_for_you() -> str:
    """The three newest Activity rows (spec 4.1), from the same list the Activity tab shows, so the two always agree."""
    newest = [r for _, rows in F.ACTIVITY for r in rows][:3]
    rows = ''.join(activity_row(i, t, f'{s} &middot; {when}', u) for i, t, s, _, when, u in newest)
    return f'<section><div class="sh"><h2>New for you</h2><a href="#">All activity ({UNREAD}) &rsaquo;</a></div><div class="card pad">{rows}</div></section>'


# ------------------------------------------------------------------------------------------------ the screens
def home(platform: str) -> str:
    head = (f'<div class="head">{medallion("warrior", 56)}<div><h1 style="color:{D.CLASS_COLOR["warrior"]}">Thandrel</h1>'
            f'<p>Level 60 Fury Warrior &middot; Olympus XXVII</p></div></div>')
    top = f'<div class="band hero">{status(platform)}{navbar(platform)}{head}</div>'
    return top + f'<div class="body">{raid_card()}{upgrades_block()}{new_for_you()}</div>' + tabbar(platform, 'Home')


def big_title(platform: str, title: str, extra: str = '') -> str:
    """iOS large title under the bar; on Android the top app bar already carries the title, so only `extra` shows."""
    h1 = f'<h1>{title}</h1>' if platform == 'ios' else ''
    return f'<div class="bigt">{h1}{extra}</div>' if h1 or extra else ''


# The Raid tab's content (owner ruling 2026-11-09: no second full card; Home keeps it): who is coming by role, the
# comp, and before the pull. Names, counts and the checklist are the raid-night screen's (design/mocks/ws_app.py).
COMING = [('Tanks', 2, 2, 'Grunvald, Brakka'), ('Healers', 5, 8, 'Mirethil, Selka, Hollowmend, Aenwyn, Brightwen'),
          ('Damage', 11, 30, 'Thandrel, Kaelra, Thalrin and 8 more')]
BEFORE_PULL = [('ok', 'Drakefire Amulet', 'Attuned'), ('ok', 'Elixir of the Mongoose', '3 in your bags'),
               ('miss', 'Flask of the Titans', 'None in your bags &middot; <a href="#">Auction house</a>'),
               ('warn', 'Boots not enchanted', 'Enchant Boots: Greater Agility &middot; <a href="#">Where to get it</a>')]
MARK = {'ok': ('#7fd48a', '&#10003;'), 'miss': ('#f0736b', '&times;'), 'warn': ('#e8b04a', '!')}


def raid_header(answer: str = ANSWER, late_min: int | None = None) -> str:
    """The Raid tab's top (owner ruling 2026-11-09): one compact line, the raid, its time and your answer chip; the
    answers under it change it (F1). The character and role show only when the account has more than one character in
    the guild (ruling 1)."""
    chip = f'Late &middot; {late_min} min' if late_min else answer  # lane M2a: one word for Late's minutes, the chip and the rows
    who = f'<span class="s">{F.AS_ME[0].upper()}{F.AS_ME[1:]}</span>' if chars_in_guild(GUILD) > 1 else ''
    answers = ''.join(f'<a class="{"on" if x == answer else ""}" href="#">{x}</a>' for x in ('In', 'Late', 'Bench', 'Out'))
    late = ('<div class="chips">' + ''.join(f'<a class="{"on" if m == late_min else ""}" href="#">{m} min</a>' for m in LATE_MINUTES) + '</div>'
            if answer == 'Late' else '')
    return (f'<section class="card pad rhead"><div class="rl">{ic(F.ONY, 40)}<span class="tx"><b>{F.RAID}</b>'
            f'<span class="s">8:00 pm CST &middot; pull in 2 h 48 min</span>{who}</span><span class="achip">{chip}</span></div>'
            f'<div class="answers">{answers}</div>{late}</section>')


def raid(platform: str, answer: str = ANSWER, late_min: int | None = None) -> str:
    coming = ''.join(f'<div class="row"><span class="tx"><b>{role} <span class="num {"ok" if n >= need else "lo"}">{n} of {need}</span></b>'
                     f'<span class="s">{names}{" &middot; " + str(need - n) + " short" if n < need and role == "Healers" else ""}</span></span></div>'
                     for role, n, need, names in COMING)
    pull = ''.join(f'<div class="row"><span class="mk" style="background:{MARK[k][0]}">{MARK[k][1]}</span><span class="tx"><b>{t}</b><span class="s">{s}</span></span></div>'
                   for k, t, s in BEFORE_PULL)
    pages = [('Roster', '38 members, 31 at 60'), ('Calendar', 'Onyxia Wed, Barrow Deeps Sat'), ('Loot history', 'Last raid: 4 pieces'),
             ('Guild sims', 'You are #7 of 14 DPS')]
    rows = ''.join(f'<div class="row"><span class="tx"><b>{t}</b><span class="s">{s}</span></span><span class="muted">&rsaquo;</span></div>' for t, s in pages)
    top = f'<div class="band">{status(platform)}{navbar(platform, "Raid")}{big_title(platform, "Raid")}</div>'
    return (top + f'<div class="body">{raid_header(answer, late_min)}'
            f'<section><div class="sh"><h2>Who is coming</h2><a href="#">Roster &rsaquo;</a></div><div class="card framed pad">{coming}</div></section>'
            f'<section><div class="sh"><h2>Before you pull</h2></div><div class="card pad">{pull}<p class="s stamp">From your last log-out, 9 h ago</p></div></section>'
            f'<section><div class="sh"><h2>Olympus XXVII</h2></div><div class="card pad">{rows}</div></section></div>'
            + tabbar(platform, 'Raid'))


def gear(platform: str) -> str:
    doll = [OFF_HAND_WORN if r[1] in GATED else r for r in F.DOLL]
    where = [('inv_weapon_crossbow_04', 3, 'Blackrock Spire', 'Blackcrow &middot; Shadow Hunter Vosh&rsquo;gajin, 8%'),
             ('inv_helmet_36', 4, 'Blacksmithing', 'Lionheart Helm, Ironwood Blade &middot; ask a blacksmith'),
             ('inv_jewelry_ring_35', 4, 'Stormpike Guard', 'Don Julio&rsquo;s Band, needs revered')]
    wrows = ''.join(f'<div class="row">{ic(i, 40, q)}<span class="tx"><b>{t}</b><span class="s">{s}</span></span></div>' for i, q, t, s in where)
    srows = ''.join(f'<div class="row">{ic(i, 40, q)}<span class="tx"><span class="n" style="color:var(--q-{"epic" if q == 4 else "rare"})">{name}</span>'
                    f'<span class="s">{slot} &middot; {src}</span></span><span class="g num">{"Worn" if g is None else g + "<small>DPS</small>"}</span></div>'
                    for slot, name, i, q, src, g in doll)
    summary = ('<div class="card framed pad"><div class="row"><span class="tx"><b>Fury Warrior &middot; 60 raid &middot; Alliance</b>'
               '<span class="s">4 of 17 slots on the list</span></span><a href="#">Change</a></div></div>')
    top = f'<div class="band">{status(platform)}{navbar(platform, "Gear")}{big_title(platform, "Gear")}</div>'
    return (top + f'<div class="body">{summary}<section><div class="sh"><h2>Where to go tonight</h2></div><div class="card pad">{wrows}</div></section>'
            f'<section><div class="sh"><h2>Every slot</h2></div><div class="card pad">{srows}</div></section></div>' + tabbar(platform, 'Gear'))


DOLL_SLOT = {'Head': 'head', 'Neck': 'neck', 'Shoulder': 'shoulder', 'Back': 'back', 'Chest': 'chest', 'Wrist': 'wrist', 'Hands': 'hands',
             'Waist': 'waist', 'Legs': 'legs', 'Feet': 'feet', 'Main hand': 'main_hand', 'Off hand': 'off_hand', 'Ranged': 'ranged'}
FURY_60_DPS = '1004.2'  # Fury Warrior's set at 60 raid, the tier list's figure (bis file set_dps)


def gear_cells() -> dict[str, dict]:
    """The Gear tab's list (the app-flows DOLL, Greenhammer gated) as paperdoll cells; numbered fingers and trinkets."""
    cells, seen = {}, {}
    for slot, name, icon_name, q, _, gain in [OFF_HAND_WORN if r[1] in GATED else r for r in F.DOLL]:
        key = DOLL_SLOT.get(slot)
        if key is None:
            seen[slot] = seen.get(slot, 0) + 1
            key = f'{slot.lower()}{seen[slot]}'
        cells[key] = {'state': 'item', 'name': name.replace('&rsquo;', '’'), 'icon': icon_name, 'quality': 'epic' if q == 4 else 'rare',
                      'vs_worn': 'worn' if gain is None else gain}
    return cells


def gear_frame(platform: str) -> str:
    """Gear on the paperdoll (owner, 2026-11-09): at phone width the frame's slots would fall under 36 px, so the frame's
    header (crest, title, band, the set's DPS) and its two-column slot list lead; Where to go tonight follows."""
    import paperdoll as PD
    cells = gear_cells()
    lead = '60 raid &middot; Alliance &middot; 4 of 17 slots on the list &middot; <a href="#">Change</a>'  # the Gear summary names the list, first (spec 4.3)
    doll = PD.list_form(cells, 'head', False, 'Fury Warrior', FURY_60_DPS, '', lead, avatar_uri('thandrel'))  # the profile button's portrait
    where = [('inv_weapon_crossbow_04', 3, 'Blackrock Spire', 'Blackcrow &middot; Shadow Hunter Vosh&rsquo;gajin, 8%'),
             ('inv_helmet_36', 4, 'Blacksmithing', 'Lionheart Helm, Ironwood Blade &middot; ask a blacksmith'),
             ('inv_jewelry_ring_35', 4, 'Stormpike Guard', 'Don Julio&rsquo;s Band, needs revered')]
    wrows = ''.join(f'<div class="row">{ic(i, 40, q)}<span class="tx"><b>{t}</b><span class="s">{s}</span></span></div>' for i, q, t, s in where)
    top = f'<div class="band">{status(platform)}{navbar(platform, "Gear")}{big_title(platform, "Gear")}</div>'
    return (f'<style>{PD.CSS}.pdl{{display:block}}</style>' + top + f'<div class="body" style="--frame:url({PD.frame_uri("1x")})">{doll}'
            f'<section><div class="sh"><h2>Where to go tonight</h2></div><div class="card pad">{wrows}</div></section></div>' + tabbar(platform, 'Gear'))


def activity(platform: str) -> str:
    chips = ''.join(f'<a class="{"on" if g == "All" else ""}" href="#">{g}</a>' for g in F.GROUPS)
    groups = ''.join(f'<section><div class="sh"><h2>{day}</h2>{"<a href=#>Mark all read</a>" if day == "Today" else ""}</div><div class="card pad">'
                     + ''.join(f'<div class="row">{ic(i, 40)}<span class="tx{" dot" if u else ""}"><b>{t}</b><span class="s">{s}</span>'
                               f'<span class="s">{when}</span></span></div>' for i, t, s, _, when, u in rows) + '</div></section>'
                     for day, rows in F.ACTIVITY)
    top = f'<div class="band">{status(platform)}{navbar(platform, "Activity")}{big_title(platform, "Activity", f"<div class=\"chips scroll\">{chips}</div>")}</div>'
    return top + f'<div class="body">{groups}</div>' + tabbar(platform, 'Activity')


def menu(platform: str) -> str:
    """The profile menu (spec 3): over Home, the drawer from the left; iOS a modal drawer at 85%, Android M3's."""
    me = F.CHARACTERS[0]
    chars = ''.join(f'<div class="row">{profile(av, fa) if av else glyph_profile(fa)}<span class="tx"><b style="color:{D.CLASS_COLOR[cls]}">{n}</b>'
                    f'<span class="s">{line}</span></span></div>' for n, cls, line, _, av, fa in F.CHARACTERS[1:])
    tools = ''.join(f'<div class="row"><span class="tx"><b>{t}</b></span></div>' for t, _, _ in F.TOOLS)  # labels alone (owner ruling 2026-11-09)
    account = ''.join(f'<div class="row"><span class="tx"><b>{t}</b><span class="s">{s}</span></span></div>' for t, s in F.ACCOUNT[:3])
    you = (f'<div class="you"><div style="display:flex;gap:14px;align-items:center">{profile()}<div><h2 style="font-size:26px;color:{D.CLASS_COLOR["warrior"]}">{me[0]}</h2>'
           f'<p style="font-size:13px">{me[2]}</p><span class="tag">Main</span></div></div>'
           f'<p style="font-size:12px;margin-top:10px;color:var(--text)">Updated 2 h ago from Battle.net &middot; <a href="#">Refresh</a></p></div>')
    drawer = (f'<aside class="drawer">{you}<h3>Switch character</h3><div class="rows">{chars}<p style="padding:8px 0"><a href="#">Show low-level characters ({len(F.LOW)})</a></p></div>'
              f'<h3>Tools</h3><div class="rows">{tools}</div><h3>Account</h3><div class="rows">{account}</div></aside>')
    return home(platform) + '<div class="scrim"></div>' + drawer


def signup(platform: str) -> str:
    """F1 change the answer, on the Raid tab: Late selected opens the 15/30/60 chips under the answers, on the card (spec
    F1, two taps); no sheet, the answer posts on the tap."""
    return raid(platform, 'Late', 30)


def item_sheet(platform: str) -> str:
    """F5 item detail: Deathbringer from Onyxia (8%, the item file): the client tooltip, where, runner-up, who needs it."""
    tip = ('<div class="tip"><div class="q4" style="font-size:15px">Deathbringer</div>Binds when picked up<br>'
           '<span style="display:flex;justify-content:space-between"><span>One-Hand</span><span>Axe</span></span>'
           '<span style="display:flex;justify-content:space-between"><span>114 - 213 Damage</span><span>Speed 2.90</span></span>(56.4 damage per second)<br>'
           '<span class="gr">Chance on hit: Sends a shadowy bolt at the enemy causing 125 Shadow damage.</span><br>Requires Level 60'
           '<div style="color:#ffd100;margin-top:6px">+12.6 DPS</div></div>')
    rows = (f'<div class="row">{ic(F.ONY, 40)}<span class="tx"><b>Onyxia, {F.RAID}</b><span class="s">8% drop chance &middot; Wed with Olympus XXVII</span></span></div>'
            f'<div class="row">{ic("inv_sword_18", 40, 4)}<span class="tx"><span class="n" style="color:var(--q-epic)">Vis&rsquo;kag the Bloodletter</span>'
            f'<span class="s">Runner-up &middot; Onyxia &middot; 8% &middot; &minus;1.8 DPS</span></span></div>'
            f'<div class="row">{medallion("warrior", 40)}<span class="tx"><b style="color:{D.CLASS_COLOR["warrior"]}">Grunvald</b><span class="s">Protection Warrior &middot; needs it</span></span></div>'
            f'<div class="row">{medallion("shaman", 40)}<span class="tx"><b style="color:{D.CLASS_COLOR["shaman"]}">Torvak</b><span class="s">Enhancement Shaman &middot; has it</span></span></div>')
    sheet = f'<div class="sheet"><div class="grab"></div>{tip}<div style="margin-top:12px">{rows}</div></div>'
    return gear(platform) + '<div class="scrim"></div>' + sheet


def lock(platform: str) -> str:
    """F2 raid night at 0 taps: iOS the lock-screen Live Activity, Android the ongoing notification, the four answers on it."""
    acts = ''.join(f'<a class="{"on" if x == "In" else ""}" href="#">{x}</a>' for x in ('In', 'Late', 'Bench', 'Out'))
    app = f'<div class="app"><img src="{flat_logo_uri(32)}" alt="">Forever Sixty &middot; now</div>'
    la = (f'<div class="la">{app}<div class="top">{ic(F.ONY, 40)}<div><b>{F.RAID}</b><div style="font-size:13px;opacity:.8">Olympus XXVII &middot; You are In</div></div>'
          f'<span class="cd num">42:00</span></div><p class="ln">Starts 8:00 pm CST &middot; {F.COUNTS} &middot; 2 tanks, 5 heals</p><div class="acts">{acts}</div></div>')
    date = 'Wednesday 6 January'
    return f'<div class="lock" style="position:absolute;inset:0">{status(platform)}<div class="clock"><div class="d">{date}</div><div class="t">7:18</div></div>{la}</div>'


SCREENS = {'home': home, 'raid': raid, 'gear': gear, 'gear-frame': gear_frame, 'activity': activity, 'menu': menu, 'signup': signup, 'item': item_sheet, 'lock': lock}


def build() -> None:
    for platform in DEVICES:
        for name, fn in SCREENS.items():
            page(name, platform, fn(platform), name in SCROLL_SCREENS)


if __name__ == '__main__':
    build()
