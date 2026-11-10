"""The native app frame for the mobile redesign mocks (design/specs/2026-11-07-mobile-redesign.md section 4).

Two jobs, one frame:
- `from_web(mock, platform, ...)`: a screen that matches a web page is the web page's own committed 390 px mock
  (design/mocks/web/<mock>.html) with the site header and footer hidden and the app's status bar, navigation bar and
  tab bar drawn around it. The web mock stays the one source for the screen's structure and data.
- `app_page(platform, body, ...)`: an app-only screen (widgets, Live Activity, raid night, offline data, update
  required, devices, deletion), composed from ws_prims primitives inside the same frame.

`APP_CSS` also carries the native type deltas the spec rules (N1 to N4): labels in sentence case, the 16 pt gutter.
The frame is chrome, not content: the parity gate for native compares the content area under the navigation bar.
"""
from __future__ import annotations

import json
import re
from pathlib import Path

HERE = Path(__file__).resolve().parent
WEB = HERE / 'web'
APP_OUT = HERE / 'app'

# Device canvases in points / dp (spec 4.1): iPhone 16 Pro class and Pixel 9 class.
DEVICES = {'ios': (393, 852), 'android': (412, 915)}
# Safe areas and bars (spec 4.2), in points / dp.
BARS = {
    'ios': {'status': 54, 'nav': 44, 'tab': 49, 'home': 34},
    'android': {'status': 32, 'nav': 64, 'tab': 80, 'home': 16},
}

# The four tabs are the app's jobs (design/specs/2026-11-08-app-flows.md section 3, owner 2026-11-08): the answer, the raid,
# your gear, what changed. Guides, Tiers and the other tools are in the profile menu. id, label.
TABS = [('home', 'Home'), ('guild', 'Guild'), ('gear', 'Gear'), ('activity', 'Activity')]
# The tab icons as the apps draw them (web/scripts/gen-ui-icons.mjs NATIVE_VECTORS): the owner's traced nav set, and for Guild the
# site's own Guild door (design/assets/icons/ui/guild.svg, spec 2026-11-10-guild-tab.md section 1), with thicken.json's stroke.
TAB_ICON_SOURCES = {'home': 'icons/nav/home.svg', 'guild': 'icons/ui/guild.svg', 'gear': 'icons/nav/gear.svg',
                    'activity': 'icons/nav/activity.svg'}
_ASSETS = HERE.parent / 'assets'
_THICKEN = {k: v for k, v in json.loads((_ASSETS / 'icons/nav/thicken.json').read_text()).items() if not k.startswith('$')}


def tab_icon(name: str, px: int = 24) -> str:
    """A tab's icon, tintable (currentColor), thickened where the apps' vector is."""
    body = (_ASSETS / TAB_ICON_SOURCES[name]).read_text().strip()
    if name in _THICKEN:
        body = body.replace('fill="currentColor"', f'fill="currentColor" stroke="currentColor" stroke-width="{_THICKEN[name]}" stroke-linejoin="round"', 1)
    return body.replace('<svg ', f'<svg width="{px}" height="{px}" aria-hidden="true" ', 1)
# The crest that opens the profile menu sits top left on every tab root (flows spec section 3).
CREST = '../../../web/public/icons/hd/crests/{cls}.webp'
# Battle.net avatar stand-ins (the client's race portraits, 84 px like the render CDN's avatar.jpg): the selected character's
# avatar_url on the profile button and the menu rows (owner exception to the one crest language, flows spec section 3).
AVATAR = '../data/avatars/{name}.jpg'
_TOKENS = json.loads((HERE.parent / 'tokens' / 'tokens.json').read_text())['color']['faction']
FACTION_RING = {'alliance': _TOKENS['allianceBar']['$value'], 'horde': _TOKENS['hordeBar']['$value']}
SELECTED = {'warrior': ('thandrel', 'alliance'), 'hunter': ('felsnap', 'alliance')}  # the scenario's selected character
DEFAULT = object()  # avatar: the selected character's own
BADGE = 0.4  # the faction badge, as a share of the portrait's diameter
# The tier list's faction toggle mark (web/src/lib/faction-mark.ts factionMarkSrc): the client's unit-frame emblem in HD.
FACTION_MARK = '../../../web/public/icons/hd/faction/{faction}.webp'

_GLYPH = {
    'home': '<path d="M4 11 12 4l8 7v9h-5v-6H9v6H4z"/>',
    'bis': '<path d="M12 3 5 6v6c0 4 3 7 7 9 4-2 7-5 7-9V6z"/><path d="m9 12 2 2 4-4"/>',
    'guides': '<path d="M5 4h6a2 2 0 0 1 2 2v14a2 2 0 0 0-2-2H5zM19 4h-6a2 2 0 0 0-2 2v14a2 2 0 0 1 2-2h6z"/>',
    'guild': '<path d="M6 3v18M6 4h12l-3 4 3 4H6"/>',
    'raid': '<path d="M6 3v18M6 4h12l-3 4 3 4H6"/>',
    'gear': '<path d="M12 3 5 6v6c0 4 3 7 7 9 4-2 7-5 7-9V6z"/><path d="m9 12 2 2 4-4"/>',
    'activity': '<path d="M6 16V11a6 6 0 1 1 12 0v5l2 2H4z"/><path d="M10 20a2 2 0 0 0 4 0"/>',
    'you': '<circle cx="12" cy="8" r="4"/><path d="M4 21c1-4 4-6 8-6s7 2 8 6"/>',
    'back_ios': '<path d="m15 4-8 8 8 8"/>',
    'back_android': '<path d="M20 12H5m6-7-7 7 7 7"/>',
    'share_ios': '<path d="M12 3v12M8 7l4-4 4 4"/><path d="M6 11v9h12v-9"/>',
    'share_android': '<circle cx="18" cy="5" r="2.5"/><circle cx="6" cy="12" r="2.5"/><circle cx="18" cy="19" r="2.5"/><path d="m8.2 10.8 7.6-4.6M8.2 13.2l7.6 4.6"/>',
    'more': '<circle cx="5" cy="12" r="1.4"/><circle cx="12" cy="12" r="1.4"/><circle cx="19" cy="12" r="1.4"/>',
}


def glyph(name: str, size: int = 24, stroke: float = 1.8) -> str:
    return (f'<svg width="{size}" height="{size}" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="{stroke}" '
            f'stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">{_GLYPH[name]}</svg>')


APP_CSS = '''
/* ---- native deltas over the web system (spec 3.3: N1 to N4) ---- */
/* round 5 (lane M0): the apps set Open Sans, so their mocks do too */
:root{--page-inset:16px!important;--font-body:'Open Sans',Arial,sans-serif!important}
@media(min-width:0px){.pc{padding-inline:16px!important}}
.t-label,.cdn .eb,.frow .wh .k,.rgrp>h3,.menu .mg,.rprose .glabel,.gridm .h,.cdn .clk>span>span{text-transform:none!important;letter-spacing:.01em!important}
.cdn .clk>span>span,.frow .wh .k{font-size:12px!important}
.fs-chrome{display:none!important}
/* the phone paperdoll cell (spec 3.4 Paperdoll): the slot label sits above the name so the name and source get the width */
.cell{grid-template-columns:40px minmax(0,1fr) auto!important;row-gap:0}
.cell>.lb{grid-column:2;grid-row:1;flex-direction:row!important;gap:8px;font-size:12px!important;line-height:16px!important}
.cell>:nth-child(2){grid-column:1;grid-row:1/span 2;margin-top:2px}
.cell>.bd{grid-column:2;grid-row:2}
.cell>.gn,.cell>.gn0,.cell>span:last-child:not(.bd){grid-column:3;grid-row:1/span 2;width:auto!important;max-width:104px}
body{min-height:var(--dev-h)!important}
main{min-height:calc(var(--dev-h) - var(--top) - var(--bottom));flex:1 0 auto}
/* phone: the two countdown panels stack (the web mock keeps two columns at 390 and overflows; spec 5 H-1) */
.cds{grid-template-columns:minmax(0,1fr)!important}
/* ---- the frame ---- */
.app-status{position:relative;flex:none;height:var(--status);z-index:50;display:flex;align-items:center;justify-content:space-between;padding:0 28px 0 32px;
  font:600 16px/1 -apple-system,'SF Pro Text','Open Sans',sans-serif;color:var(--strong);background:var(--bg)}
.app-status .island{position:absolute;left:50%;top:11px;width:124px;height:36px;margin-left:-62px;border-radius:20px;background:#000}
.app-status .sys{display:flex;gap:6px;align-items:center}.app-status .sys i{display:block;width:17px;height:11px;border-radius:2px;background:var(--strong)}
.app-status .sys i.b{width:25px;height:12px;border-radius:3px}
[data-platform=android] .app-status{padding:0 20px;font:500 14px/1 'Roboto','Open Sans',sans-serif}
.app-nav{position:relative;flex:none;height:var(--nav);z-index:49;display:flex;align-items:center;gap:8px;padding:0 8px;background:var(--bg)}
.app-nav a{display:inline-flex;align-items:center;justify-content:center;min-width:44px;height:44px;color:var(--gold);font:400 17px/1 var(--font-body);gap:2px}
.app-nav .ttl{flex:1;text-align:center;font:600 17px/22px var(--font-body);color:var(--strong);white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
.app-nav .ttl.hidden{visibility:hidden}
[data-platform=android] .app-nav{padding:0 4px}[data-platform=android] .app-nav a{min-width:48px;height:48px;color:var(--strong)}
[data-platform=android] .app-nav .ttl{text-align:left;font-size:20px;line-height:28px;padding-left:4px}
.app-tabs{position:relative;flex:none;height:calc(var(--bottom) - var(--home));z-index:49;display:grid;grid-template-columns:repeat(4,1fr);align-items:start;
  background:color-mix(in srgb,var(--bg) 94%,transparent);backdrop-filter:blur(12px);border-top:1px solid var(--line-soft);padding-top:6px}
.app-tabs a{display:flex;flex-direction:column;align-items:center;gap:3px;min-height:44px;color:var(--nav);font:600 12px/14px var(--font-body)}
.app-tabs a.on{color:var(--gold)}
.app-tabs a{position:relative}.app-tabs .badge{position:absolute;top:-2px;left:calc(50% + 6px);min-width:18px;height:18px;padding:0 5px;border-radius:9px;
  background:var(--wipe);color:#17110a;font:700 12px/18px var(--font-body);text-align:center}
[data-platform=android] .app-tabs .badge{top:2px;left:calc(50% + 10px)}
.app-nav .menu-crest{position:relative;width:44px;height:44px}
.app-nav .menu-crest .anon{width:32px;height:32px;border-radius:50%;box-shadow:0 0 0 2px #9a9484;background:var(--raised)}
.pf{position:relative;display:inline-block;flex:none}
.pf-img{width:100%;height:100%;border-radius:50%;object-fit:cover;box-shadow:0 0 0 2px var(--ring);background:var(--raised);display:block}
.pf-anon{display:flex!important;align-items:center;justify-content:center;color:var(--muted)}
.pf-badge{position:absolute;top:-3px;right:-4px;object-fit:contain;filter:drop-shadow(0 0 1px var(--bg)) drop-shadow(0 0 1px var(--bg))}
.app-nav .menu-crest .anon{display:flex;align-items:center;justify-content:center;color:var(--muted)}
[data-platform=android] .app-tabs{padding-top:12px;border-top:0;background:var(--raised)}
[data-platform=android] .app-tabs a{gap:4px;min-height:48px}
[data-platform=android] .app-tabs a .ind{display:flex;align-items:center;justify-content:center;width:64px;height:32px;border-radius:16px}
[data-platform=android] .app-tabs a.on .ind{background:#e5b9552e}
[data-platform=android] .app-tabs a.on{color:var(--strong)}[data-platform=android] .app-tabs a.on svg{color:var(--gold)}
.app-home{display:block;flex:none;align-self:center;margin:calc((var(--home) - 5px) / 2) 0;width:134px;height:5px;border-radius:3px;background:var(--strong);z-index:51}
[data-platform=android] .app-home{width:108px;height:4px;opacity:.7}
/* pull to refresh affordance, drawn only on the states that show it */
.app-ptr{display:flex;justify-content:center;padding:10px 0 0}.app-ptr span{width:22px;height:22px;border-radius:50%;border:2px solid var(--line);border-top-color:var(--gold)}
'''


def frame_vars(platform: str, tabbar: bool, navbar: bool) -> str:
    w, h = DEVICES[platform]
    b = BARS[platform]
    top = b['status'] + (b['nav'] if navbar else 0)
    bottom = (b['tab'] + b['home']) if tabbar else b['home']
    return f'--dev-w:{w}px;--dev-h:{h}px;--status:{b["status"]}px;--nav:{b["nav"]}px;--home:{b["home"]}px;--top:{top}px;--bottom:{bottom}px'


def status_bar(platform: str) -> str:
    if platform == 'ios':
        return '<div class="app-status" aria-hidden="true"><span>9:41</span><span class="island"></span><span class="sys"><i></i><i></i><i class="b"></i></span></div>'
    return '<div class="app-status" aria-hidden="true"><span>9:41</span><span class="sys"><i></i><i></i><i class="b"></i></span></div>'


def portrait(faction: str, avatar: str | None, px: int) -> str:
    """A character on the profile button and in the menu (owner, 2026-11-08): the Battle.net portrait says "me"; the ring in
    the faction's colour and the faction emblem badge on its top right say "which faction" (the class is Home's header
    crest). With no avatar (a new character, a failed fetch) the default profile image, the person glyph signed out
    uses, fills the ring and the badge stays: never the class crest, so the class icon is never drawn twice on Home."""
    b = round(px * BADGE)
    face = (f'<img class="pf-img" src="{AVATAR.format(name=avatar)}" alt="">' if avatar
            else f'<span class="pf-img pf-anon">{glyph("you", round(px * 0.6))}</span>')
    badge = f'<img class="pf-badge" src="{FACTION_MARK.format(faction=faction)}" alt="" style="width:{b}px;height:{b}px">'
    return (f'<span class="pf" style="width:{px}px;height:{px}px;--ring:{FACTION_RING[faction]}">'
            f'{face}{badge}</span>')


def menu_crest(who: str | None, avatar=DEFAULT) -> str:
    """The profile menu's door, top left on every tab root inside a 44 pt / 48 dp target: the selected character's
    portrait (who is its class slug), or a person glyph signed out."""
    if who:
        av, faction = SELECTED.get(who, (None, 'alliance'))
        av = av if avatar is DEFAULT else avatar
        return f'<a href="#" class="menu-crest" aria-label="Open menu">{portrait(faction, av, 32)}</a>'
    return f'<a href="#" class="menu-crest" aria-label="Open menu"><span class="anon">{glyph("you", 20)}</span></a>'


def nav_bar(platform: str, back: str = '', title: str = '', share: bool = False, show_title: bool = False,
            who: str | None = None, menu: bool = False, avatar=DEFAULT) -> str:
    left = ''
    if menu and not back:
        left = menu_crest(who, avatar)
    elif back:
        g = glyph('back_ios' if platform == 'ios' else 'back_android', 22 if platform == 'ios' else 24, 2.2 if platform == 'ios' else 1.8)
        left = f'<a href="#" aria-label="Back to {back}">{g}{back if platform == "ios" else ""}</a>'
    else:
        left = '<a aria-hidden="true"></a>'
    right = f'<a href="#" aria-label="Share">{glyph("share_ios" if platform == "ios" else "share_android", 22)}</a>' if share else '<a aria-hidden="true"></a>'
    ttl = f'<span class="ttl{"" if show_title else " hidden"}">{title}</span>'
    return f'<nav class="app-nav" aria-label="Navigation">{left}{ttl}{right}</nav>'


def tab_bar(platform: str, active: str, badges: dict | None = None) -> str:
    """badges: the open count per tab (Guild: your unanswered raid or the officer's open items; Activity: unread)."""
    badges = badges or {}
    items = ''
    for tid, label in TABS:
        on = ' on' if tid == active else ''
        cur = ' aria-current="page"' if on else ''
        g = tab_icon(tid, 24)
        inner = f'<span class="ind">{g}</span>' if platform == 'android' else g
        n = badges.get(tid, 0)
        mark = f'<span class="badge" aria-label="{n} new">{n}</span>' if n else ''
        items += f'<a href="#" class="{on.strip()}"{cur}>{inner}{mark}<span>{label}</span></a>'
    return f'<nav class="app-tabs" aria-label="Tabs">{items}</nav>'


def chrome(platform: str, tab: str | None, back: str = '', title: str = '', share: bool = False, navbar: bool = True,
           who: str | None = 'warrior', badges: dict | None = None, avatar=DEFAULT) -> tuple[str, str, str]:
    """(body style vars, top html, bottom html). A tab root (a tab, no back) carries the menu crest of `who` (the selected
    character's class; the person glyph signed out), on all four tabs in the same place (flows spec section 3). `badges`
    are the tabs' counts."""
    menu = bool(tab)
    top = status_bar(platform) + (nav_bar(platform, back, title, share, who=who, menu=menu, avatar=avatar) if navbar else '')
    bottom = (tab_bar(platform, tab, badges) if tab else '') + '<span class="app-home" aria-hidden="true"></span>'
    return frame_vars(platform, bool(tab), navbar), top, bottom


def _inject(html: str, platform: str, tab: str | None, back: str, title: str, share: bool, navbar: bool, extra_css: str = '',
            who: str | None = 'warrior', badges: dict | None = None, avatar=DEFAULT) -> str:
    style, top, bottom = chrome(platform, tab, back, title, share, navbar, who, badges, avatar)
    # The apps keep the flat system until the mobile lane adopts round 5: the web mocks' round-5 layer stays out.
    html = re.sub(r'<style data-system="round5">.*?</style>', '', html, count=1, flags=re.S)
    html = html.replace('</head>', f'<style>{APP_CSS}{extra_css}</style></head>', 1)
    html = re.sub(r'<body([^>]*)>', lambda m: f'<body{m.group(1)} data-platform="{platform}" style="{style}">{top}', html, count=1)
    return html.replace('</body>', f'{bottom}</body>', 1)


def from_web(mock: str, platform: str, tab: str | None, back: str = '', title: str = '', share: bool = False,
             navbar: bool = True, extra_css: str = '', who: str | None = 'warrior', badges: dict | None = None) -> str:
    """The committed web mock as an app screen. Relative asset paths still resolve: design/mocks/app is a sibling of web."""
    src = (WEB / f'{mock}.html').read_text()
    return _inject(src, platform, tab, back, title, share, navbar, extra_css, who, badges)


def app_page(platform: str, html: str, tab: str | None, back: str = '', title: str = '', share: bool = False, navbar: bool = True,
             extra_css: str = '', who: str | None = 'warrior', badges: dict | None = None, avatar=DEFAULT) -> str:
    """An app-only screen built with ws_prims.build_page (its HTML already carries the system CSS)."""
    return _inject(html, platform, tab, back, title, share, navbar, extra_css, who, badges, avatar)


def write_app(name: str, html: str) -> Path:
    APP_OUT.mkdir(parents=True, exist_ok=True)
    p = APP_OUT / name
    p.write_text(html)
    return p
