"""Shared CSS and helpers for the web-layout redesign mocks (design/specs/2026-11-05-web-layout.md).

One responsive HTML file per page, built from one flat system: three surfaces, one control bar, one type
scale, one page column. Tokens are the shipped ones (web/src/styles/tokens.css); the new ones are
marked `proposed`. Content is sample text for layout; the build renders from data.
"""
import base64
import json
import re
from html import escape
from pathlib import Path

from mocklib import ROOT, CLASS_COLOR, crest
from wide_lib import icon, crest_img, QCOL  # same real-icon helpers

OUT = Path(__file__).resolve().parent / 'web'

# The parity gate's region names (design/parity.md; web/src/lib/ui/regions.ts is the same list). The mock
# emits them as data-region, the built primitives emit the same, and web/tests/e2e/support/parity.ts pairs them.
REGION = {'site_header': 'site-header', 'site_footer': 'site-footer', 'head': 'head', 'head_title': 'head-title', 'head_meta': 'head-meta', 'head_answer': 'head-answer',
          'controls': 'controls', 'summary': 'summary', 'notes': 'notes', 'plate': 'plate', 'slots': 'slots',
          'inspector': 'inspector', 'inspector_inline': 'inspector-inline', 'footer_note': 'footer-note'}


def region_slug(title: str) -> str:
    """A Section's region, `section-<slug>`: sectionRegion in web/src/lib/ui/regions.ts."""
    import re
    return 'section-' + re.sub(r'^-+|-+$', '', re.sub(r'[^a-z0-9]+', '-', title.lower()))
RENDERS = Path(__file__).resolve().parent / 'renders' / 'web'

CSS = '''
:root{--bg:#07090d;--raised:#0d111a;--hover:#131824;--line:#262e40;--line-soft:#1c2230;--line-warm:#3a3326;
--text:#e9e4d8;--strong:#f2eee4;--muted:#9a9484;--nav:#b9b3a4;--gold:#e5b955;--gold-hover:#f5d27a;--kill:#7fd48a;--wipe:#f0736b;
--font-title:'Marcellus',Georgia,serif;--font-body:'Barlow','Helvetica Neue',Arial,sans-serif;
/* proposed layout tokens */
--page-max:1344px;--page-gutter:48px;--page-inset:18px;--section-gap:48px;--gap:24px;--measure-reading:720px;--rail-w:360px;--toc-w:200px;
--bar-h:56px;--row-h:56px}
@media(min-width:768px){:root{--page-inset:0px}}
@media(min-width:1760px){:root{--page-max:clamp(1600px,calc(100vw - 160px),1792px);--page-gutter:64px;--section-gap:64px;--gap:32px;--rail-w:400px}}
*{box-sizing:border-box}
html{font-size:16px}
html,body{margin:0;background:var(--bg);color:var(--text);font-family:var(--font-body)}body{font-size:15px;line-height:24px}
a{color:var(--gold);text-decoration:none}a:hover{color:var(--gold-hover)}
h1,h2,h3,p,ul,dl,dd{margin:0}ul{padding:0;list-style:none}
:focus-visible{outline:2px solid var(--gold);outline-offset:2px}
/* type scale, decided 2026-11-05 (option 2 with class colour): Marcellus for the page title only, once per page, in the
   spec or class colour; everything else is Barlow. Item names 700 in quality colour, section heads 700 normal case,
   numbers Barlow tabular. No mono, no condensed, no all-caps headings. Nothing under 12. */
.t-title{font:400 40px/46px var(--font-title);letter-spacing:.01em;color:var(--strong)}
.t-h2{font:700 22px/28px var(--font-body);color:var(--strong)}
.t-h3{font:700 18px/24px var(--font-body);color:var(--strong)}
.t-item{font:700 18px/22px var(--font-body)}
.t-lead{font-size:17px;line-height:26px;max-width:62ch;color:var(--text)}
.t-prose{font-size:16px;line-height:26px;max-width:var(--measure-reading)}
.t-small{font-size:13px;line-height:20px;color:var(--muted)}
.t-label{font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.num{font-family:var(--font-body);font-variant-numeric:tabular-nums;font-weight:700}
.strong{color:var(--strong)}.muted{color:var(--muted)}.gold{color:var(--gold)}.kill{color:var(--kill)}.wipe{color:var(--wipe)}
@media(max-width:767px){.t-title{font-size:32px;line-height:38px}.t-h2{font-size:20px;line-height:26px}}
/* the one page column */
.pc{width:100%;max-width:var(--page-max);margin-inline:auto;padding-inline:var(--page-inset)}
@media(min-width:768px){.pc{padding-inline:var(--page-gutter)}}
.full-bleed{width:100%}
/* surfaces: S0 page, S1 plate (the page's one primary object), S2 hover. Nothing else has a box. */
.plate{background:var(--raised);border:1px solid var(--line);border-radius:6px;min-width:0}
.rule{border-top:1px solid var(--line-soft)}
.section{padding-top:var(--section-gap)}
.sh{display:flex;align-items:baseline;gap:12px;padding-bottom:12px;border-bottom:1px solid var(--line-soft);margin-bottom:16px}
.sh .t-small{margin-left:auto}
/* buttons: one primary per view */
.btn{display:inline-flex;align-items:center;justify-content:center;gap:8px;height:44px;padding:0 18px;border:1px solid var(--line);border-radius:4px;background:transparent;color:var(--strong);font:700 13px/16px var(--font-body);letter-spacing:.08em;text-transform:uppercase;white-space:nowrap}
.btn:hover{border-color:var(--gold);color:var(--gold-hover)}
.btn.primary{background:var(--gold);border-color:var(--gold);color:#17110a}.btn.primary:hover{background:var(--gold-hover);color:#17110a}
.btn.sm{height:36px;padding:0 12px;font-size:12px}
.link{font-weight:600}
.pill{display:inline-flex;align-items:center;height:20px;padding:0 6px;border-radius:3px;font:700 12px/16px var(--font-body);letter-spacing:.06em;text-transform:uppercase;border:1px solid var(--gold);color:var(--gold)}
.ic{border-radius:5px;border:1px solid #4a4030;background:#10141d;flex:none;display:inline-block;object-fit:cover}
.ic.q3{border-color:#0070dd}.ic.q4{border-color:#a335ee}.ic.q2{border-color:#1eff00}.ic.q5{border-color:#ff8000}
.crest{flex-shrink:0;object-fit:cover;border-radius:999px;box-shadow:0 0 0 2px color-mix(in srgb,var(--c) 55%,transparent);background:var(--raised)}
.bar-track{height:8px;border-radius:3px;background:var(--line-soft);overflow:hidden}.bar-track>i{display:block;height:100%;background:var(--gold);border-radius:3px}
/* page head: crest + title left, the answer right */
.head{display:grid;gap:16px;padding-block:16px 14px;grid-template-columns:minmax(0,1fr);align-items:end}
.head .id{display:flex;gap:16px;align-items:center;min-width:0}
.head .id>div{min-width:0}
.head .ans{display:flex;flex-direction:column;gap:4px}
@media(min-width:900px){.head{grid-template-columns:minmax(0,1fr) auto;padding-block:20px 16px}.head .ans{align-items:flex-end;text-align:right}}
/* THE control bar: one pattern for every page */
.cbar{position:sticky;top:0;z-index:5;background:color-mix(in srgb,var(--bg) 94%,transparent);backdrop-filter:blur(6px);border-block:1px solid var(--line-soft)}
.cbar .pc{display:flex;align-items:center;gap:20px;min-height:var(--bar-h);padding-block:6px;flex-wrap:wrap}
.cbar .grp{display:flex;align-items:center;gap:8px}
.cbar .grp>.t-label{display:none}
@media(min-width:1280px){.cbar:not([data-labels-from="1920"]) .grp>.t-label{display:block}}
@media(min-width:1920px){.cbar .grp>.t-label{display:block}}
.seg{display:inline-flex;border:1px solid var(--line);border-radius:4px;overflow:hidden}
.seg a{display:inline-flex;align-items:center;gap:8px;height:40px;padding:0 14px;font:600 14px/16px var(--font-body);color:var(--nav);border-left:1px solid var(--line-soft);white-space:nowrap}
.seg a:first-child{border-left:0}
.seg a:hover{background:var(--hover);color:var(--strong)}
.seg a.on{background:#e5b9551a;color:var(--gold);box-shadow:inset 0 -2px 0 var(--gold)}
.seg a .n{font-size:12px;color:var(--muted)}.seg a.on .n{color:var(--gold)}
.sel{display:inline-flex;align-items:center;gap:10px;height:40px;padding:0 12px;border:1px solid var(--line);border-radius:4px;font:600 14px/16px var(--font-body);color:var(--strong)}
.sel::after{content:"\\25BE";color:var(--muted);font-size:12px}
.cbar .end{margin-left:auto;display:flex;gap:16px;align-items:center}
.bar-d{display:none}.bar-m{display:block;width:100%}
@media(min-width:768px){.bar-d{display:contents}.bar-m{display:none}}
.bar-m summary{display:flex;align-items:center;gap:10px;min-height:48px;list-style:none;font:600 14px/20px var(--font-body);color:var(--strong);cursor:pointer}
.bar-m summary::-webkit-details-marker{display:none}
.bar-m summary .chg{margin-left:auto;color:var(--gold);font-size:13px}
.bar-m .stack{display:flex;flex-direction:column;gap:12px;padding-bottom:12px}
.bar-m .seg{display:flex}.bar-m .seg a{flex:1;justify-content:center;height:44px}.bar-m .sel{height:44px;width:100%;justify-content:space-between}
.cbar.static{position:static;background:none;border-top:0}
.mocknote{margin-block:32px 0;font-size:12px;color:var(--muted)}
.ic.fit,.fit>.ic{width:100%!important;height:auto!important;aspect-ratio:1}
'''

def seg(items: list[tuple], extra: str = '', label: str = '') -> str:
    """items: (label, on, count|None, icon_html). A radiogroup of radios, as the built ControlBar and Segmented are."""
    a = ''.join(f'<a href="#" role="radio" aria-checked="{"true" if on else "false"}" class="{"on" if on else ""}">{ic}{lab}{f" <span class=\"n\">{n}</span>" if n is not None else ""}</a>' for lab, on, n, ic in items)
    name = f' aria-label="{label}"' if label else ''
    return f'<span class="seg {extra}" role="radiogroup"{name}>{a}</span>'


def sel(label: str) -> str:
    return f'<span class="sel" role="combobox" aria-label="{label}">{label}</span>'


def control_bar(groups: list[tuple[str, str]], summary: str, end: str = '', open_m: bool = False, sticky: bool = True,
                label: str = 'Controls', labels_from: int = 1280) -> str:
    """groups: (label, html). Desktop: one row. Phone: one summary line that expands in place. Group labels show
    from `labels_from` (1280; the BiS bar 1920, spec 5.13)."""
    d = ''.join(f'<span class="grp"><span class="t-label">{g[0]}</span>{g[1]}</span>' for g in groups)
    m = ''.join(f'<span><span class="t-label" style="display:block;margin-bottom:4px">{g[0]}</span>{g[2] if len(g) > 2 else g[1]}</span>' for g in groups)
    return (f'<div class="cbar full-bleed{"" if sticky else " static"}" role="group" aria-label="{label}" data-region="{REGION["controls"]}" data-labels-from="{labels_from}"><div class="pc"><div class="bar-d">{d}<span class="end">{end}</span></div>'
            f'<details class="bar-m"{" open" if open_m else ""}><summary><span>{summary}</span><span class="chg">Change &#9662;</span></summary><div class="stack">{m}</div></details></div></div>')


# ---- the site chrome: the shipped header and footer, captured from the build ----------------------------------
# web/tests/e2e/chrome-capture.spec.ts writes data/chrome/: chrome.css (the site CSS that styles them, scoped under
# .fs-chrome) and one <mock>.json per parity state (the exact header and footer the build draws in that state) plus
# four templates. A mock with no capture of its own takes a template: signed out, or signed in with its `who`.
CHROME = Path(__file__).resolve().parent / 'data' / 'chrome'
_CHROME_MARK = '<!--fs-chrome:%s-->'
_TEMPLATE_WHO = ('Felsnap', 'hunter', '34')
# Base.astro's body: a full-height column with the footer pushed to the bottom of a short page.
_BODY_CSS = 'body{display:flex;flex-direction:column;min-height:100vh}.fs-foot{margin-top:auto}'   # the character template-in.json was captured with (home fixtures)


def _chrome_file(name: str) -> dict:
    return json.loads((CHROME / f'{name}.json').read_text())


def _with_who(chrome: dict, who: tuple) -> dict:
    """template-in with another character: the chip's name, class crest and "Class · level" line."""
    name, cls, sub = who
    level = (re.findall(r'\d+', sub) or [_TEMPLATE_WHO[2]])[-1]
    tname, tcls, tlevel = _TEMPLATE_WHO
    swap = {tname: name, f'{tcls.capitalize()} · {tlevel}': f'{cls.capitalize()} · {level}',
            f'{tcls.capitalize()}, level {tlevel}': f'{cls.capitalize()}, level {level}'}
    header = chrome['header']
    for old, new in swap.items():
        header = header.replace(old, new)
    # the chip's own crest(s) come before the menu panels' nine
    head, cut, tail = header.partition('class="site-menu"')
    head = head.replace(f'crests/{tcls}.webp', f'crests/{cls}.webp')
    html = {**chrome['html'], 'data-pointer-class': cls}
    return {**chrome, 'header': head + cut + tail, 'html': html}


def chrome_for(mock: str, who: tuple | None, state: str = '') -> dict:
    """The capture for this mock when the parity gate checks it, else a template (state: '', 'new' or 'pending')."""
    own = CHROME / f'{mock.removesuffix(".html")}.json'
    if own.exists():
        return json.loads(own.read_text())
    if state:
        return _chrome_file(f'template-{state}')
    return _with_who(_chrome_file('template-in'), who) if who else _chrome_file('template-out')


def page(title: str, body: str, nav: str = '', who: tuple | None = None, note: str = '', chrome_state: str = '') -> str:
    """The mock's page. Its header and footer are filled in by `write`, which knows the mock's name."""
    n = note or 'Layout mock. Sample content for layout only; the build renders every figure from data. Spec: design/specs/2026-11-05-web-layout.md.'
    mark = _CHROME_MARK % json.dumps({'who': list(who) if who else None, 'state': chrome_state, 'nav': nav})
    return ('<!doctype html><html lang="en"%%HTML_ATTRS%%><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>%s</title>'
            '<link rel="preconnect" href="https://fonts.googleapis.com"><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>'
            '<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Barlow:wght@400;600;700&family=Cinzel:wght@700&family=Marcellus&display=swap">'
            '<style>%%CHROME_CSS%%</style><style>%s%s</style></head><body>%s<div class="fs-chrome"><div data-fs-body%%BODY_ATTRS%%>%%HEADER%%</div></div>'
            '<main>%s</main><!-- %s --><div class="fs-chrome fs-foot"><div data-fs-body%%BODY_ATTRS%%>%%FOOTER%%</div></div></body></html>'
            % (title, CSS, _BODY_CSS, mark, body, n))


def _fill_chrome(name: str, html: str) -> str:
    found = re.search(r'<!--fs-chrome:(\{.*?\})-->', html)
    if found is None:
        return html
    spec = json.loads(found.group(1))
    chrome = chrome_for(name, tuple(spec['who']) if spec['who'] else None, spec['state'])
    if spec.get('nav') and not (CHROME / f'{name.removesuffix(".html")}.json').exists():
        # a template marks no door: mark the mock's own, as the build does for its path
        door = f'data-door="{spec["nav"].lower()}"'
        chrome = {**chrome, 'header': chrome['header'].replace(door, f'{door} aria-current="page"', 1)}
    attrs = ''.join(f' {k}="{escape(v, quote=True)}"' for k, v in chrome['html'].items())
    body_attrs = ''.join(f' {k}="{escape(v, quote=True)}"' for k, v in chrome.get('body', {}).items())
    css = (CHROME / 'chrome.css').read_text()
    return (html.replace(found.group(0), '').replace('%HTML_ATTRS%', attrs).replace('%BODY_ATTRS%', body_attrs).replace('%CHROME_CSS%', css)
            .replace('%HEADER%', chrome['header']).replace('%FOOTER%', chrome['footer']))


def write(name: str, html: str) -> Path:
    OUT.mkdir(exist_ok=True)
    p = OUT / name
    p.write_text(_fill_chrome(name, html))
    return p
