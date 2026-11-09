"""Round 5 style: Blizzard's WoW: Forever colours exactly, soft rounded corners, the owner's art set
(design/specs/2026-11-09-directions.md).

Colours are read from their page (blizzard/palette.json: computed values; blizzard/*.png: averaged solid regions). Nothing
else of theirs: no logo, art or ornament asset. Our logo is the owner's LX seal (flat at 48 px and under, painted above).
Corners follow their nav bar (measured about 8 px): one scale, R_LG 10 px for the nav, panels and plaques, R_SM 6 px for
buttons and controls, 3 px for icons; the client tooltip keeps the client's 4. The featured-panel frame (a thin cream
inset line on the radius, with small rounded corner flourishes on the most important one) goes on two or three elements
per page, never all. The ground is the owner's painted brown tile made periodic and band-shaped (art.ground_tile); the blue band is the owner's painted
band (art.py). Fonts: Marcellus (titles) and Open Sans, both SIL OFL 1.1.
Copy rule (ai-slop-reviewer tell 22): a number takes its unit and nothing else.
"""
from html import escape
from urllib.parse import quote

import data as D
from art import GROUND_TILE_PX, band_uri, ground_uri, hero_uri
from assets import flat_logo_uri, logo_uri

FONTS = 'family=Archivo:wght@500;600;700&family=Marcellus&family=Open+Sans:wght@400;600;700'
FACTION_NOTCH = {'alliance': '#2f6fd6', 'horde': '#c0392b'}
R_LG, R_SM, R_IC = 10, 6, 3
FRAME_LINE = 'rgba(243,238,226,.42)'

# Blizzard's Forever page, exactly. c: computed (palette.json); s: sampled flat region of a capture.
TOKENS = {
    'nav': '#0a2933',         # c  rgb(10,41,51), the nav bar
    'nav-2': '#0b2f3b',       # c  rgb(11,47,59)
    'band': '#054562',        # s  the section band (#053d58 top to #044c6c foot); now the fallback under the painted band
    'navy': '#001845',        # c  rgb(0,24,69)
    'night': '#071d2e',       # s  the hero sky
    'earth': '#67563a',       # s  the brown ground (#645337 to #6a593c); the fallback under the painted tile
    'earth-stripe': '#716145',  # s  their table stripe on the ground
    'card': '#170e09',        # c  rgb(23,14,9), their footer and body brown-black
    'foot': '#170e09',        # c  rgb(23,14,9), the footer
    'cream': '#f3eee2',       # c  rgb(243,238,226), display type
    'cream-2': '#fcedcf',     # c  rgb(252,237,207), footer heading
    'text': 'rgba(255,255,255,.78)',  # c  their body is white at .7; .78 for our denser 16 px
    'muted': 'rgba(255,255,255,.5)',  # c
    'gold': '#ffec71',        # c  rgb(255,236,113), their Learn More link
    'gold-hi': '#f7b90a',     # c  rgb(247,185,10), their footnote link, used for hover
    'teal': '#0f6064',        # c  rgb(15,96,100), Buy Now
    'teal-edge': '#408076',   # c  rgb(64,128,118), Subscribe
    'blue': '#0093b2',        # s  the active tab underline
    'blue-btn': '#114f72',    # c  rgb(17,79,114), their secondary button (Upgrade Now)
    'rule': '#5a5d70',        # c  rgb(90,93,112), their border
    'rule-warm': '#82705b',   # c  rgb(130,112,91)
    'gain': '#32ae88',        # c  rgb(50,174,136), their green-teal edge, for gains
}


def _svg_url(svg: str) -> str:
    return 'url("data:image/svg+xml,' + quote(svg) + '")'


def _corner(rot: int) -> str:
    """One rounded frame corner: the inset line turning on the radius, an inner arc ending in a small scroll, and a bead."""
    return _svg_url(
        f"<svg xmlns='http://www.w3.org/2000/svg' width='26' height='26' viewBox='0 0 26 26'><g transform='rotate({rot} 13 13)' "
        f"fill='none' stroke='#f3eee2' stroke-opacity='.8' stroke-linecap='round'><path d='M6 15 Q6 6 15 6' stroke-width='1.1'/>"
        f"<path d='M15 6 q3 0 3 3 q0 2 -2 2' stroke-width='1'/><circle cx='6' cy='18.5' r='1.3' fill='#f3eee2' fill-opacity='.8' stroke='none'/>"
        f"</g></svg>")


CORNERS = (f"{_corner(0)} top left/26px no-repeat,{_corner(90)} top right/26px no-repeat,"
           f"{_corner(270)} bottom left/26px no-repeat,{_corner(180)} bottom right/26px no-repeat")

# Display face: Marcellus (SIL OFL) until ITC Friz Quadrata is licensed; then this one line becomes
# "'Friz Quadrata', 'Marcellus', Georgia, serif" plus its @font-face (DISPLAY_FACE).
DISPLAY_FONT = "'Marcellus',Georgia,serif"
DISPLAY_TRACKING = '-.015em'  # Quadrata sets tighter than Marcellus; measured against their headings (blizzard/fonts.png)
DISPLAY_FACE = ''  # the licensed font's @font-face rule goes here


# Item names on our dark surfaces use the live site's lifted rare and epic (wide_lib.QCOL): the client's #0070dd reads
# dim on brown-black (ai-slop-reviewer, round 5). The client tooltip keeps the client's exact colours (.tip below).
PAGE_QUALITY = {**D.QUALITY, 'rare': '#3d94f0', 'epic': '#b866f5'}


def _tokens() -> str:
    q = ';'.join(f'--q-{k}:{v}' for k, v in PAGE_QUALITY.items())
    r = f'--r-lg:{R_LG}px;--r-sm:{R_SM}px;--r-ic:{R_IC}px'
    return ':root{' + ';'.join(f'--{k}:{v}' for k, v in TOKENS.items()) + ';' + q + ';' + r + '}'


CSS = DISPLAY_FACE + _tokens() + ''':root{--title:''' + DISPLAY_FONT + ''';--title-track:''' + DISPLAY_TRACKING + ''';}
:root{--body:'Open Sans',Arial,sans-serif;--ui:'Archivo',Arial,sans-serif;--max:1320px}
@media(min-width:1760px){:root{--max:1640px}}
*{box-sizing:border-box}html,body{margin:0}
body{background:var(--earth) url(''' + ground_uri() + ''') 0 0/''' + str(GROUND_TILE_PX) + '''px repeat;color:var(--text);font:16px/1.55 var(--body);
  min-height:100vh;display:flex;flex-direction:column}
body>main{flex:1 0 auto;width:100%}  /* the footer sits at the foot of a short page, never mid-viewport (reviewers, batch 1) */
a{color:var(--gold);text-decoration:none;font-weight:600}a:hover{color:var(--gold-hi)}
h1,h2,h3{font-family:var(--title);font-weight:400;letter-spacing:var(--title-track);color:var(--cream);text-shadow:0 1px 2px rgba(0,0,0,.45);margin:0}
p,ul,ol{margin:0;padding:0;list-style:none}
.w{max-width:var(--max);margin:0 auto;padding:0 28px}@media(max-width:700px){.w{padding:0 14px}}
.num{font-variant-numeric:tabular-nums lining-nums}.muted{color:var(--muted)}.strong{color:#fff}
.ic{display:block;border:1px solid #000;box-shadow:0 0 0 1px var(--rule);border-radius:var(--r-ic)}.ic.empty{background:#141414}
/* navigation: their floating rounded bar, inset from the edges over the band; our seal and name lead it */
.topnav{max-width:var(--max);margin:0 auto;padding:16px 28px 0}
.bar{display:flex;align-items:center;height:64px;background:var(--nav-2);border:1px solid rgba(255,255,255,.08);border-radius:var(--r-lg);font:600 17px var(--ui)}
.brand{display:flex;align-items:center;gap:10px;height:100%;padding:0 18px;border-right:1px solid rgba(255,255,255,.1);font:400 25px/1 var(--title);
letter-spacing:var(--title-track);color:var(--cream);white-space:nowrap}.brand img{width:40px;height:40px}
.nav{display:flex;gap:22px;padding:0 24px;align-self:stretch}.nav a{display:flex;align-items:center;gap:6px;color:#fff;font-weight:600;white-space:nowrap}
.nav a.cur{box-shadow:inset 0 -3px 0 var(--blue)}.chev{width:9px;height:6px;opacity:.7}
.tools{margin-left:auto;display:flex;align-items:center;gap:16px;padding-right:8px}
.srch{display:flex;align-items:center;justify-content:center;width:40px;height:40px;border-radius:var(--r-sm);color:#fff}
.acct{display:flex;align-items:center;gap:9px;color:#fff;white-space:nowrap;font-weight:600}.acct .lv{color:var(--muted);font-weight:500}
/* Premium is their Subscribe button exactly (owner 2026-11-08: "just copy the subscribe button - keep the logo"):
   flat #408076, measured 47 px tall in their 64 px bar with a 6 px radius, white semibold label */
.bar .btn{display:flex;align-items:center;height:47px;padding:0 22px;font:600 16px var(--body);color:#fff;background:var(--teal-edge);
border:0;border-radius:var(--r-sm)}
.bar .btn:hover{background:var(--teal-edge);color:#fff;filter:brightness(1.08)}
.menu{display:none;color:#fff;font-weight:600}
@media(max-width:1320px){.nav,.bar .btn{display:none}.menu{display:block}}  /* the full bar with a long name needs 1318 */
@media(max-width:700px){.topnav{padding:8px 14px 0}.bar{height:56px}.brand{font-size:19px;padding:0 12px;border-right:0}.brand img{width:34px;height:34px}
.acct .nm,.acct .lv,.srch{display:none}.tools{gap:12px}}
/* the section band: the owner's painted band, navy at the top into soft teal cloud, stretched with cover (never tiled) */
.band{background:linear-gradient(rgba(5,58,84,.38),rgba(5,58,84,.38)),var(--band) url(''' + band_uri() + ''') center 30%/cover no-repeat;border-bottom:1px solid #000}
/* home only: the owner's camp-road art behind the title, under one flat veil of their night blue so type stays AA */
.band.hero{--veil:rgba(7,29,46,.16);background:linear-gradient(var(--veil),var(--veil)),url(''' + hero_uri() + ''') var(--art-pos,0 44%)/var(--art-size,124% auto) no-repeat,var(--night)}
.band.hero .row{min-height:150px;align-items:flex-end;padding-bottom:22px}
@media(max-width:700px){.band.hero{--art-pos:84% 60%;--art-size:cover}.band.hero .row{min-height:160px;padding-bottom:16px}}
.titlebar .row{display:flex;align-items:center;gap:14px;padding:16px 0 14px}.titlebar .txt{min-width:0}
.titlebar h1{font-size:26px;line-height:30px}.titlebar .meta{font-size:14px;color:var(--text)}
.titlebar .right{margin-left:auto;color:var(--text);white-space:nowrap}.titlebar .right b{color:#fff;font:600 26px var(--body)}
.ctl{display:flex;flex-wrap:wrap;gap:2px 28px;border-top:1px solid rgba(255,255,255,.14);padding:2px 0}
.ctl .g{display:flex;flex-wrap:wrap;gap:0 16px}.ctl a{color:rgba(255,255,255,.72);padding:8px 0}.ctl a.on{color:#fff;box-shadow:inset 0 -3px 0 var(--blue)}
.ctl.sub{border-top:0;padding:0 0 2px}.ctl.sub a{font-size:14px;padding:4px 0 6px}.ctl.sub a.on{box-shadow:inset 0 -2px 0 var(--cream)}
@media(max-width:700px){.titlebar .row{gap:10px}.titlebar .medal{width:40px!important;height:40px!important}.titlebar h1{font-size:21px;line-height:25px}
.titlebar .right b{font-size:21px}.ctl{gap:0 20px}
/* phone: each control group is one line that scrolls sideways, as the client and every app tab strip do */
.ctl .g{flex-wrap:nowrap;overflow-x:auto;scrollbar-width:none;max-width:100%;-webkit-mask-image:linear-gradient(90deg,#000 88%,transparent);mask-image:linear-gradient(90deg,#000 88%,transparent)}.ctl .g a{white-space:nowrap}
/* phone: the band's figure drops under the title, so a long name and meta keep the full width */
.titlebar .row{flex-wrap:wrap}.titlebar .txt{flex:1;min-width:0}.titlebar .right{flex-basis:100%;margin-left:0;padding-top:2px}}
/* surfaces: soft corners on every panel; the frame (cream inset line on the radius) on featured panels only, with
   rounded corner flourishes (ornate) on the page's one lead panel */
.card{position:relative;background:var(--card);border:1px solid #000;box-shadow:0 0 0 1px rgba(130,112,91,.55);border-radius:var(--r-lg);overflow:hidden}
.framed::before{content:"";position:absolute;inset:6px;border:1px solid ''' + FRAME_LINE + ''';border-radius:calc(var(--r-lg) - 4px);pointer-events:none;z-index:1}
.framed.ornate::after{content:"";position:absolute;inset:4px;background:''' + CORNERS + ''';pointer-events:none;z-index:1}
.pad{padding:18px 20px}.framed.pad{padding:22px 24px}.sh{font-size:22px;margin-bottom:10px}
.medal{position:relative;flex:none}.medal img{position:absolute;border-radius:50%;object-fit:cover}.medal svg{position:absolute;inset:0}
/* actions: their flat teal with a crisp edge, gold links; soft corners */
.btn{font:600 15px var(--body);color:#fff;background:var(--teal);border:1px solid var(--teal-edge);border-radius:var(--r-sm);padding:9px 20px;display:inline-block}
.btn:hover{background:var(--teal-edge);color:#fff}
.btn2{font:600 15px var(--body);color:#fff;background:var(--blue-btn);border:1px solid #1d6a96;border-radius:var(--r-sm);padding:8px 18px;display:inline-block}
.btn2.on{border-color:var(--cream);box-shadow:inset 0 0 0 1px var(--cream)}
.btnrow{display:flex;flex-wrap:wrap;gap:12px;align-items:center}
/* gear rows */
.gear{display:grid;grid-template-columns:40px 1fr auto;gap:12px;align-items:center;padding:8px 0;border-bottom:1px solid rgba(255,255,255,.1)}
.gear:last-child{border-bottom:0}.gear .n{display:block;font-weight:600;line-height:20px}.gear .src{display:block;font-size:13px;color:var(--muted)}
.gear .d{font:600 14px var(--body);color:var(--gain);font-variant-numeric:tabular-nums;white-space:nowrap}
.gear.sel{background:rgba(255,255,255,.06);border-radius:var(--r-sm);box-shadow:-8px 0 0 rgba(255,255,255,.06),8px 0 0 rgba(255,255,255,.06)}
/* the client's tooltip */
.tip{--q-rare:#0070dd;--q-epic:#a335ee;background:#060914;border:1px solid #8b8f9c;border-radius:4px;box-shadow:0 0 0 1px #000,inset 0 0 0 1px #2b3040;padding:9px 11px;font:13.5px/1.42 var(--body);color:#fff}
.tip .tl{display:flex;justify-content:space-between;gap:14px}.tip .tl:first-child{font-size:15px}
.tip .tl.epic{color:var(--q-epic)}.tip .tl.rare{color:var(--q-rare)}.tip .tl.uncommon{color:var(--q-uncommon)}.tip .tl.green{color:var(--q-uncommon)}
.tip .y{color:#ffd100;margin-top:6px}
/* footer: their leather brown, our painted seal large */
.foot{background:var(--foot);border-top:1px solid #000;margin-top:56px;padding:28px 0 40px;font-size:14px;color:var(--muted)}
.foot .w{display:flex;flex-wrap:wrap;align-items:center;gap:12px 24px}.foot a{color:rgba(255,255,255,.72);font-weight:400}
.foot .seal{display:flex;align-items:center;gap:14px;font:400 24px var(--title);color:var(--cream-2);margin-right:20px}.foot .seal img{width:88px;height:88px}
.foot .nb{margin-left:auto}
@media(max-width:700px){.foot .w{gap:8px 14px;font-size:13px}.foot .seal{flex-basis:100%}.foot .seal img{width:64px;height:64px}.foot .nb{flex-basis:100%;margin-left:0}}
'''

NAV = [('Best in slot', False), ('Guides', True), ('Tier list', False), ('Sim', False), ('Planner', False), ('Logs', True), ('Guild', True)]
CHEVRON = '<svg class="chev" viewBox="0 0 9 6" aria-hidden="true"><path d="M1 1l3.5 3.5L8 1" fill="none" stroke="currentColor" stroke-width="1.6"/></svg>'
SEARCH = ('<svg width="20" height="20" viewBox="0 0 20 20" aria-hidden="true"><circle cx="8.5" cy="8.5" r="6" fill="none" stroke="currentColor" '
          'stroke-width="1.8"/><path d="M13 13l5 5" stroke="currentColor" stroke-width="1.8"/></svg>')


def medallion(slug: str, size: int, faction: str = 'alliance') -> str:
    """Our class crest in a bronze ring, the faction as a corner notch."""
    inset = round(size * 14 / 96)
    notch = f'<path d="M70 12 l8 0 l0 8 z" fill="{FACTION_NOTCH[faction]}" stroke="#0a2933"/>' if size >= 40 else ''
    return (f'<span class="medal" style="width:{size}px;height:{size}px"><img src="{D.crest_uri(slug)}" alt="" '
            f'style="inset:{inset}px;width:{size - 2 * inset}px;height:{size - 2 * inset}px">'
            f'<svg viewBox="0 0 96 96" aria-hidden="true"><circle cx="48" cy="48" r="40" fill="none" stroke="#3b2a15" stroke-width="10"/>'
            f'<circle cx="48" cy="48" r="44.5" fill="none" stroke="#c9a46a" stroke-width="1.6"/><circle cx="48" cy="48" r="35.5" fill="none" stroke="#c9a46a" stroke-width="1.2"/>'
            f'{notch}</svg></span>')


def icon(name: str, size: int = 40) -> str:
    uri = D.icon_uri(name) if name else ''
    if not uri:
        return f'<span class="ic empty" style="width:{size}px;height:{size}px"></span>'
    return f'<img class="ic" src="{uri}" alt="" width="{size}" height="{size}">'


def header(name: str, slug: str, level: int | None, cur: str) -> str:
    """Their floating bar, ours inside: the flat seal and name first, menus with chevrons, search, the account on one line.
    level None leaves the level off (a real character whose level the mock does not know)."""
    lv = f'<span class="lv">{level}</span>' if level is not None else ''
    nav = ''.join(f'<a class="{"cur" if n == cur else ""}" href="#">{n}{CHEVRON if menu else ""}</a>' for n, menu in NAV)
    # no name: the signed-out bar, Sign in where the account sits
    acct = (f'<a class="acct" href="#">{medallion(slug, 32)}<span class="nm" style="color:{D.CLASS_COLOR[slug]}">{escape(name)}</span>{lv}{CHEVRON}</a>'
            if name else '<a class="acct" href="#" style="color:#fff">Sign in</a>')
    return (f'<div class="topnav"><header class="bar"><a class="brand" href="#"><img src="{flat_logo_uri(64)}" alt="">Forever Sixty</a>'
            f'<nav class="nav">{nav}</nav><div class="tools"><a class="srch" href="#" aria-label="Search">{SEARCH}</a>'
            f'{acct}<a class="btn" href="#">Premium</a><span class="menu">Menu</span></div></header></div>')


def titlebar(slug: str | None, title: str, meta: str, right: str = '', title_colour: str = '',
             ctl: list[list[tuple[str, bool]]] | None = None, top: str = '', hero: bool = False, media: str = '',
             sub: list[tuple[str, bool]] | None = None) -> str:
    """The compact title band under the floating bar (`top`; their bar floats over their band): medallion, title, one meta
    line, one number right, controls under a rule. `hero` puts the owner's camp-road art behind it (home only). `media`
    replaces the class medallion (a guild's faction emblem)."""
    med = media or (medallion(slug, 52) if slug else '')
    style = f' style="color:{title_colour}"' if title_colour else ''
    meta_html = f'<p class="meta">{meta}</p>' if meta else ''
    groups = ''.join('<div class="g">' + ''.join(f'<a class="{"on" if on else ""}" href="#">{escape(label)}</a>' for label, on in g) + '</div>'
                     for g in ctl or [])
    nav = f'<nav class="ctl" aria-label="Page controls">{groups}</nav>' if groups else ''
    if sub:  # a tab's own views, on their own row under the tabs (one item lit per row)
        nav += ('<nav class="ctl sub" aria-label="Views"><div class="g">' + ''.join(f'<a class="{"on" if on else ""}" href="#">{escape(lb)}</a>' for lb, on in sub) + '</div></nav>')
    return (f'<div class="band titlebar{" hero" if hero else ""}">{top}<div class="w"><div class="row">{med}<div class="txt"><h1{style}>{escape(title)}</h1>{meta_html}</div>'
            f'{f"<div class=right>{right}</div>" if right else ""}</div>{nav}</div></div>')


def tooltip(class_slug: str, item_id: int, ours: str = '') -> str:
    rows = ''.join(f'<div class="tl {role}"><span>{escape(left)}</span>{f"<span>{escape(right)}</span>" if right else ""}</div>'
                   for left, right, role in D.tooltip_lines(class_slug, item_id))
    extra = f'<div class="y">{escape(ours)}</div>' if ours else ''  # our line, in addon yellow: the gain alone (tell 22)
    return f'<div class="tip">{rows}{extra}</div>'


def footer() -> str:
    links = ''.join(f'<a href="#">{n}</a>' for n in ('Get set up', 'Addon', 'Discord', 'Premium', 'Privacy', 'Terms'))
    return (f'<footer class="foot"><div class="w"><span class="seal"><img src="{logo_uri(256)}" alt="">Forever Sixty</span>{links}'
            f'<span class="nb">A fan-made player tool. Not affiliated with Blizzard.</span></div></footer>')


def page(title: str, extra_css: str, body: str) -> str:
    return (f'<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">'
            f'<title>{escape(title)}</title><link rel="stylesheet" href="https://fonts.googleapis.com/css2?{FONTS}&display=swap">'
            f'<style>{CSS}{extra_css}</style></head><body>{body}</body></html>')
