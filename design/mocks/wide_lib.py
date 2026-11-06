"""Shared CSS and helpers for the wide-layout mocks (design/specs/2026-11-05-wide-layout.md).

Each mock is ONE responsive HTML file. The page column, breakpoints and rail tokens are the
proposed ones; `html.legacy` switches the xl rules off so the same markup renders today's 1344 px
column ("before") from the same source. Colours, type and radii are the shipped tokens
(web/src/styles/tokens.css); headings use Barlow Condensed 800 as the live site renders them.
Content in a mock is sample text for layout only; the build renders from data.
"""
import base64
import json
from pathlib import Path

from mocklib import ROOT, SEAL, CLASS_COLOR, crest, emblem

OUT = Path(__file__).resolve().parent / 'wide'
BUILD = json.load(open(ROOT / 'web/src/data/active-build.json'))['build']
ICONS = ROOT / 'web/public/data' / BUILD / 'icons'
_cache: dict[str, str] = {}


def icon(name: str, size: int = 36, q: str = '') -> str:
    """A real game icon as an inline data URI, in the site's rounded frame; a dark tile if absent."""
    p = next((d / f'{name}.webp' for d in [ICONS, *sorted((ROOT / 'web/public/data').glob('*/icons'))] if (d / f'{name}.webp').exists()), ICONS / f'{name}.webp')
    cls = f'ic {q}'.strip()
    if not p.exists():
        return f'<span class="{cls}" style="width:{size}px;height:{size}px"></span>'
    if name not in _cache:
        _cache[name] = 'data:image/webp;base64,' + base64.b64encode(p.read_bytes()).decode()
    return f'<img class="{cls}" src="{_cache[name]}" alt="" width="{size}" height="{size}" style="width:{size}px;height:{size}px">'


def crest_img(slug: str, px: int) -> str:
    return f'<img class="crest" src="{crest(slug)}" alt="" style="width:{px}px;height:{px}px;--c:{CLASS_COLOR[slug]}">'


QCOL = {1: '#ffffff', 2: '#1eff00', 3: '#3d94f0', 4: '#b866f5', 5: '#ff8000'}

CSS = '''
:root{--bg:#07090d;--raised:#0d111a;--card-top:#131824;--line:#262e40;--line-soft:#1c2230;--line-warm:#3a3326;
--text:#e9e4d8;--strong:#f2eee4;--muted:#9a9484;--nav:#b9b3a4;--gold:#e5b955;--gold-hover:#f5d27a;--kill:#7fd48a;--wipe:#f0736b;
--font-display:'Barlow Condensed','Arial Narrow',Arial,sans-serif;--font-body:'Barlow','Helvetica Neue',Arial,sans-serif;--font-mono:'JetBrains Mono',Menlo,monospace;
/* proposed tokens, below xl identical to shipped */
--page-max:1344px;--page-gutter:48px;--gap:24px;--rail-w:400px;--toc-w:200px;--measure-reading:720px;--lane-max:640px;--page-inset:18px}
@media(min-width:768px){:root{--page-inset:0px}}
@media(min-width:1760px){html:not(.legacy){--page-max:clamp(1600px,calc(100vw - 160px),1792px);--page-gutter:64px;--gap:32px}}
*{box-sizing:border-box}
html,body{margin:0;background:var(--bg);color:var(--text);font-family:var(--font-body);font-size:14px;line-height:1.45}
a{color:var(--gold);text-decoration:none}a:hover{color:var(--gold-hover)}
h1,h2,h3,h4,p,ul{margin:0}ul{padding:0;list-style:none}
:focus-visible{outline:2px solid var(--gold);outline-offset:2px}
.mono{font-family:var(--font-mono);font-variant-numeric:tabular-nums}
.muted{color:var(--muted)}.strong{color:var(--strong)}
.eyebrow{display:flex;align-items:center;gap:12px;font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--gold)}
.eyebrow::before{content:"";width:24px;height:1px;background:var(--gold)}
.h1{font-family:var(--font-display);font-weight:800;font-size:44px;line-height:1;letter-spacing:.01em;text-transform:uppercase;color:var(--strong)}
.h2{font-family:var(--font-display);font-weight:800;font-size:28px;line-height:1.1;letter-spacing:.02em;text-transform:uppercase;color:var(--strong)}
.h3{font-family:var(--font-display);font-weight:800;font-size:17px;letter-spacing:.04em;text-transform:uppercase;color:var(--strong)}
.lead{font-size:15px;line-height:1.55;color:var(--text);max-width:var(--measure-reading)}
.prose p{font-size:15px;line-height:1.65;margin:0 0 14px;max-width:var(--measure-reading)}
.label{font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--muted)}
/* the one column */
.pc{width:100%;max-width:var(--page-max);margin-inline:auto;padding-inline:var(--page-inset)}
@media(min-width:768px){.pc{padding-inline:var(--page-gutter)}}
.full-bleed{width:100%}
.hdr{border-bottom:1px solid var(--line-soft);background:var(--bg)}
.hdr .pc{display:flex;align-items:center;gap:28px;height:80px}
.logo{display:flex;align-items:center;gap:10px;font-family:'Cinzel',Georgia,serif;font-weight:700;font-size:21px;letter-spacing:.1em;color:var(--strong);white-space:nowrap}
.logo img{width:30px;height:30px}
.navl{display:none;gap:28px;margin-left:56px;flex:1}
.navl a{font-size:12px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--nav)}
.navl a.cur{color:var(--gold)}
.premium{border:1px solid var(--gold);border-radius:3px;padding:3px 8px;font-size:11px;letter-spacing:.1em}
.hsp{margin-left:auto;display:flex;gap:12px;align-items:center}
.srch{width:64px;height:48px;border:1px solid var(--line);border-radius:4px;display:flex;align-items:center;justify-content:center;color:var(--gold)}
.acct{display:none;align-items:center;gap:10px;height:48px;padding:0 18px 0 6px;border:1px solid var(--line-warm);border-radius:24px}
.acct i{width:36px;height:36px;border-radius:50%;background:#26405c55;border:1px solid var(--line)}
.acct b{display:block;font-weight:600;font-size:13px;color:var(--strong)}.acct small{font-size:11px;color:var(--muted)}
@media(min-width:1024px){.navl{display:flex}.acct{display:flex}}
@media(max-width:1023px){.hdr .pc{height:64px}}
/* panels */
.panel{background:linear-gradient(180deg,var(--card-top),var(--raised) 140px);border:1px solid var(--line-warm);border-radius:6px;min-width:0}
.panel.flat{background:var(--raised);border-color:var(--line)}
.ph{display:flex;align-items:center;gap:12px;padding:12px 16px;border-bottom:1px solid var(--line-soft)}
.ph .tile{width:28px;height:28px;border-radius:5px;background:#e5b95522;border:1px solid #4a4030;display:flex;align-items:center;justify-content:center;color:var(--gold);font-size:13px}
.ph h2,.ph h3{font-family:var(--font-display);font-weight:800;font-size:17px;letter-spacing:.06em;text-transform:uppercase;color:var(--strong)}
.ph .meta{margin-left:auto;font-size:12px;color:var(--muted)}
.pb{padding:16px}
.btn{display:inline-flex;align-items:center;justify-content:center;gap:8px;height:44px;padding:0 18px;border:1px solid var(--line-warm);border-radius:4px;background:#ffffff08;color:var(--strong);font-size:12px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;white-space:nowrap}
.btn:hover{border-color:var(--gold);color:var(--gold-hover)}
.btn.gold{background:linear-gradient(180deg,#f0c968,#d9a53f);border-color:#f0c968;color:#1b1407}
.btn.gold:hover{filter:brightness(1.08);color:#1b1407}
.btn.sm{height:36px;padding:0 12px;font-size:11px}
.btn.block{width:100%}
.tab{display:inline-flex;align-items:center;gap:8px;height:44px;padding:0 14px;border:1px solid var(--line);border-radius:4px;font-size:13px;color:var(--nav)}
.tab.cur{border-color:var(--gold);color:var(--gold);background:#e5b9550f}
.tab .n{font-size:11px;color:var(--muted)}
.tabs{display:flex;gap:8px;flex-wrap:wrap}
.strip{display:flex;gap:4px;border-bottom:1px solid var(--line);overflow-x:auto}
.strip a{padding:14px 12px;font-size:12px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;color:var(--nav);border-bottom:2px solid transparent;white-space:nowrap}
.strip a.cur{color:var(--strong);border-bottom-color:var(--gold)}
.ic{border-radius:5px;border:1px solid #4a4030;background:#10141d;flex:none;display:inline-block;object-fit:cover}
.ic.q3{border-color:#0070dd}.ic.q4{border-color:#a335ee}.ic.q2{border-color:#1eff00}.ic.q5{border-color:#ff8000}
.crest{flex-shrink:0;object-fit:cover;border-radius:999px;box-shadow:0 0 0 2px color-mix(in srgb,var(--c) 55%,transparent);background:var(--raised)}
.pill{display:inline-flex;align-items:center;height:20px;padding:0 7px;border-radius:3px;font-size:10px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;border:1px solid}
.pill.new{color:var(--gold);border-color:#e5b95588;background:#e5b95514}
.pill.sample{color:var(--muted);border-color:#9a948455;background:#9a948418}
.bar{height:8px;border-radius:3px;background:var(--line-soft);overflow:hidden}.bar>i{display:block;height:100%;border-radius:3px;background:var(--gold)}
.rowl{display:flex;align-items:center;gap:12px;padding:10px 16px;border-top:1px solid var(--line-soft);min-width:0}
.rowl:first-child{border-top:0}
.field{height:44px;border:1px solid var(--line);border-radius:4px;background:#0a0d14;display:flex;align-items:center;padding:0 14px;color:var(--muted);font-size:14px}
.note{font-size:12px;color:var(--muted)}
.mocknote{margin-block:40px 24px;font-size:12px;color:var(--muted)}
.ftr{border-top:1px solid var(--line-soft);margin-top:48px;background:#060810}
.ftr .pc{display:flex;gap:40px;padding-block:32px;flex-wrap:wrap;align-items:flex-start}
.ftr .col{display:flex;flex-direction:column;gap:8px;font-size:13px;min-width:120px}.ftr .col a{color:var(--nav)}.ftr .col .label{margin-bottom:4px}
/* bands: full-bleed backdrop, content back in the column */
.band{background:radial-gradient(1200px 260px at 50% -40%,#26405c66,transparent 70%),var(--bg);border-bottom:1px solid var(--line-soft)}
.band .pc{padding-block:28px 24px;display:flex;flex-direction:column;gap:16px}
/* ---- layout vocabulary (spec section 3) ---- */
.stack{display:flex;flex-direction:column;gap:var(--gap);min-width:0}
.cards3{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1fr)}
@media(min-width:1024px){.cards3{grid-template-columns:repeat(3,minmax(0,1fr));align-items:start}}
.pair{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1fr)}
@media(min-width:1024px){.pair{grid-template-columns:repeat(2,minmax(0,1fr));align-items:start}}
.main{min-width:0}
.rail{display:flex;flex-direction:column;gap:var(--gap);min-width:0}
.rail-layout{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1fr);align-items:start}
@media(min-width:1760px){html:not(.legacy) .rail-layout{grid-template-columns:minmax(0,1fr) var(--rail-w)}
 html:not(.legacy) .rail-layout>.main{grid-column:1;grid-row:1}html:not(.legacy) .rail-layout>.rail{grid-column:2;grid-row:1}}
@media(min-width:1760px) and (min-height:860px){html:not(.legacy) .sticky{position:sticky;top:96px}}
/* reading layout: toc | prose | lane */
.read{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1fr)}
.toc{display:none}
@media(min-width:1024px){.read{grid-template-columns:minmax(0,var(--measure-reading)) 220px}.read>.toc{display:block;grid-column:2;grid-row:1/span 12}.read>.sec{grid-column:1}}
@media(min-width:1760px){html:not(.legacy) .read{grid-template-columns:var(--toc-w) minmax(0,var(--measure-reading)) minmax(440px,var(--lane-max));justify-content:center}
 html:not(.legacy) .read>.toc{grid-column:1;grid-row:1/span 12}}
.toc{position:sticky;top:96px;align-self:start}
.toc .label{margin-bottom:10px;display:block;color:var(--strong)}
.toc a{display:block;padding:7px 12px;border-left:2px solid var(--line-soft);font-size:13px;color:var(--nav)}
.toc a.cur{border-left-color:var(--gold);color:var(--strong);font-weight:600}
.sec{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1fr);padding-block:40px 8px;border-top:1px solid var(--line-soft)}
.sec>.txt{min-width:0}.sec>.lane{display:flex;flex-direction:column;gap:20px;min-width:0}
@media(min-width:1024px){.sec{grid-template-columns:minmax(0,var(--measure-reading))}}
@media(min-width:1760px){html:not(.legacy) .sec{grid-template-columns:minmax(0,var(--measure-reading)) minmax(440px,var(--lane-max));grid-column:2/4}
 html:not(.legacy) .sec>.lane{position:sticky;top:96px;align-self:start}
 html:not(.legacy) .sec.wide{grid-template-columns:minmax(0,1fr);grid-column:2/4}}
.sec.wide{grid-column:auto}
/* data rows */
.dt{width:100%;border-collapse:collapse}
.dt th{text-align:left;font-size:12px;font-weight:600;color:var(--muted);padding:10px 12px;border-bottom:1px solid var(--line-soft)}
.dt td{padding:10px 12px;border-bottom:1px solid var(--line-soft);vertical-align:middle}
.dt tr:last-child td{border-bottom:0}
.dt .num{text-align:right;font-family:var(--font-mono);font-variant-numeric:tabular-nums}
.dt tbody tr:hover td{background:#ffffff05}
'''

NAV = ['Tools', 'Logs', 'Guild', 'Guides', 'Get set up']


def header(current: str = '') -> str:
    links = ''.join(f'<a class="{"cur" if n == current else ""}" href="#">{n}</a>' for n in NAV)
    return (f'<header class="hdr full-bleed"><div class="pc"><a class="logo" href="#"><img src="{SEAL}" alt="">FOREVER SIXTY</a>'
            f'<nav class="navl">{links}<span class="premium">PREMIUM</span></nav>'
            '<div class="hsp"><span class="srch">&#x2315;</span>'
            '<span class="acct"><i></i><span><b>Sign in</b><small>to see your characters</small></span></span></div></div></header>')


def footer() -> str:
    cols = {'Play': ['Get set up', 'Guides', 'Best in slot', 'Talent planner', 'Simulator', 'Tier list'],
            'Logs': ['Upload a log', 'Your reports', 'Rankings'],
            'Guilds': ['Find a guild', 'Claim your guild', 'Guild plans'],
            'About': ['Premium', 'About', 'Changelog', 'Data sources']}
    out = ''.join('<div class="col"><span class="label">%s</span>%s</div>' % (k, ''.join(f'<a href="#">{i}</a>' for i in v)) for k, v in cols.items())
    return (f'<footer class="ftr full-bleed"><div class="pc"><div class="col" style="min-width:300px"><a class="logo" href="#"><img src="{SEAL}" alt="">FOREVER SIXTY</a>'
            f'<span style="max-width:300px;margin-top:8px">Play your class better: gear, talents, sims and logs for World of Warcraft: Forever.</span></div>{out}</div></footer>')


def panel(title: str, body: str, meta: str = '', glyph: str = '&#9670;', flat: bool = False, pad: bool = True, extra: str = '', cls: str = '') -> str:
    m = f'<span class="meta">{meta}</span>' if meta else ''
    return (f'<section class="panel{" flat" if flat else ""}{" " + cls if cls else ""}" {extra}><div class="ph"><span class="tile">{glyph}</span><h2>{title}</h2>{m}</div>'
            f'<div class="{"pb" if pad else ""}">{body}</div></section>')


def page(title: str, body: str, nav: str = '', legacy: bool = False, note: str = '') -> str:
    n = note or 'Layout mock. Sample content for layout only; the build renders every figure from data. Spec: design/specs/2026-11-05-wide-layout.md.'
    return ('<!doctype html><html lang="en"%s><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1">'
            '<title>%s</title><link rel="preconnect" href="https://fonts.googleapis.com"><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>'
            '<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Barlow+Condensed:wght@800&family=Barlow:wght@400;600;700&family=Cinzel:wght@700&family=JetBrains+Mono:wght@500&display=swap">'
            '<style>%s</style></head><body>%s<main>%s</main><p class="mocknote pc">%s</p>%s</body></html>'
            % (' class="legacy"' if legacy else '', title, CSS, header(nav), body, n, footer()))


def write(name: str, html: str) -> Path:
    OUT.mkdir(exist_ok=True)
    p = OUT / name
    p.write_text(html)
    return p

CSS += '.ic.fit,.fit>.ic{width:100%!important;height:auto!important;aspect-ratio:1}\n'
