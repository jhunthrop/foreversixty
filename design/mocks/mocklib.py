"""Shared head, tokens and chrome for the page mock boards under design/mocks.

A board is a standalone HTML file: fonts from Google Fonts (the artifact host admits that
origin), tokens copied from web/src/styles/tokens.css, the few utility classes the boards use
(.label, .display, .mono, .panel, .btn, .pill, .crest), and the site's real HD crest, faction
emblem and seal as data URIs so one file renders anywhere. Nothing here is invented: every
colour, size and asset comes from the shipped site.
"""
import base64
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
WEB = ROOT / 'web'
DATA = Path(__file__).resolve().parent / 'data'
RENDERS = Path(__file__).resolve().parent / 'renders'

GOLD = '#e5b955'; MUTED = '#9a9484'; TEXT = '#f2eee4'; BODY = '#c9c2b2'
BORDER = '#262e40'; SOFT = '#1c2230'; BG = '#07090d'; RAISED = '#0d111a'; NIGHT = '#26405c'
GREEN = '#7fd48a'
CLASS_COLOR = {'warrior': '#c69b6d', 'paladin': '#f48cba', 'hunter': '#aad372', 'rogue': '#fff468', 'priest': '#ffffff',
               'shaman': '#3f8fe0', 'mage': '#3fc7eb', 'warlock': '#8788ee', 'druid': '#ff7c0a', 'monk': '#00ff98'}
PARSE = {'grey': '#8c8c8c', 'green': '#1eff00', 'blue': '#0070ff', 'purple': '#a335ee', 'orange': '#ff8000', 'pink': '#e268a8'}


def data_uri(path: Path, mime: str | None = None) -> str:
    mime = mime or {'webp': 'image/webp', 'svg': 'image/svg+xml', 'png': 'image/png'}[path.suffix.lstrip('.')]
    return f'data:{mime};base64,' + base64.b64encode(path.read_bytes()).decode()


def crest(slug: str) -> str:
    return data_uri(WEB / 'public/icons/hd/crests' / f'{slug}.webp')


def emblem(faction: str) -> str:
    return data_uri(WEB / 'public/icons/hd/faction' / f'{faction}.webp')


SEAL = data_uri(ROOT / 'design/logo/foreversixty-mark.svg')

HEAD = f'''<!doctype html><html lang="en"><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1"><title>Forever Sixty mock</title>
<link rel="preconnect" href="https://fonts.googleapis.com"><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Cinzel:wght@700;800&family=Barlow:wght@400;600;700&family=JetBrains+Mono:wght@500&display=swap">
<style>
:root{{--bg:{BG};--raised:{RAISED};--line:{BORDER};--line-soft:{SOFT};--text:#e9e4d8;--strong:{TEXT};--muted:{MUTED};--gold:{GOLD};--night:{NIGHT};
--font-display:'Cinzel','Trajan Pro',Georgia,serif;--font-body:'Barlow','Helvetica Neue',Arial,sans-serif;--font-mono:'JetBrains Mono','SF Mono',Menlo,monospace}}
*{{box-sizing:border-box}}html,body{{margin:0;background:var(--bg);color:var(--text);font-family:var(--font-body);font-size:14px;line-height:1.45}}
a{{color:var(--gold);text-decoration:none}}h1,h2,h3{{margin:0}}
.display{{font-family:var(--font-display)}}.mono{{font-family:var(--font-mono);font-variant-numeric:tabular-nums}}
.label{{font-size:11px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}}
.panel{{background:var(--raised);border:1px solid var(--line);border-radius:6px}}
.btn{{display:inline-flex;align-items:center;justify-content:center;height:44px;padding:0 16px;border:1px solid #3a3326;border-radius:4px;color:var(--strong);font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;white-space:nowrap}}
.btn-gold{{border-color:var(--gold);color:var(--gold)}}.btn-fill{{background:var(--gold);border-color:var(--gold);color:#0b0e14}}
.pill{{display:inline-flex;align-items:center;gap:6px;height:22px;padding:0 8px;border-radius:3px;font-size:11px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;border:1px solid}}
.pill-sample{{color:var(--muted);background:rgba(154,148,132,.12);border-color:rgba(154,148,132,.3)}}
.crest{{flex-shrink:0;object-fit:cover;border-radius:999px;box-shadow:0 0 0 2px color-mix(in srgb,var(--c) 55%,transparent);background:var(--raised)}}
.band{{position:relative;background:var(--raised)}}.band::after{{content:"";position:absolute;inset:0;background:linear-gradient(180deg,#26405c66 0%,#07090dcc 70%,#07090d 100%)}}.band>*{{position:relative;z-index:1}}
.nav a{{color:#b9b3a4;font-size:12px;font-weight:700;letter-spacing:.1em;text-transform:uppercase}}.nav a.current{{color:var(--strong)}}
</style></head><body>'''

NAV_ITEMS = ['Planner', 'BiS', 'Simulator', 'Logs', 'Rankings', 'Tier List', 'Guides', 'Get set up']


def nav(current: str, signed_in: bool, character: dict | None = None) -> str:
    links = ''.join(f'<a href="#{l.lower().replace(" ", "-")}" class="{"current" if l == current else ""}">{l}</a>' for l in NAV_ITEMS)
    if signed_in and character:
        chip = (f'<span style="display:flex;align-items:center;gap:8px;height:44px;padding:0 12px;border:1px solid var(--line);border-radius:4px;">'
                f'<img class="crest" src="{crest(character["class"])}" alt="" style="width:32px;height:32px;--c:{CLASS_COLOR[character["class"]]}">'
                f'<img src="{emblem(character["faction"])}" alt="" style="width:16px;height:16px;object-fit:contain">'
                f'<span style="font-size:13px;color:var(--strong)">{character["battletag"]}</span></span>')
    else:
        chip = '<a href="#signin" style="font-size:13px;color:#b9b3a4">Sign in</a>'
    discord = ('<span class="btn" style="height:40px;padding:0 14px;gap:8px;text-transform:none;letter-spacing:0;font-size:13px">'
               '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M8 12h.01M16 12h.01M5 18l2-2h10l2 2V8a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4z"/></svg>Discord</span>')
    return (f'<div class="nav" style="display:flex;align-items:center;justify-content:space-between;gap:24px;padding:18px 48px;border-bottom:1px solid var(--line-soft)">'
            f'<a href="#home" style="display:flex;align-items:center;gap:10px"><img src="{SEAL}" alt="" style="width:32px;height:32px"><span class="display" style="font-size:18px;font-weight:700;letter-spacing:.1em;color:var(--strong)">FOREVER SIXTY</span></a>'
            f'<span style="display:flex;gap:22px">{links}</span><span style="display:flex;align-items:center;gap:12px">{chip}{discord}</span></div>')


def nav_phone(signed_in: bool, character: dict | None = None) -> str:
    chip = (f'<img class="crest" src="{crest(character["class"])}" alt="" style="width:32px;height:32px;--c:{CLASS_COLOR[character["class"]]}">' if signed_in and character else '')
    return (f'<div style="display:flex;align-items:center;justify-content:space-between;padding:14px 18px;border-bottom:1px solid var(--line-soft)">'
            f'<span style="display:flex;align-items:center;gap:8px"><img src="{SEAL}" alt="" style="width:28px;height:28px"><span class="display" style="font-size:14px;font-weight:700;letter-spacing:.1em;color:var(--strong)">FOREVER SIXTY</span></span>'
            f'<span style="display:flex;align-items:center;gap:10px">{chip}<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="{TEXT}" stroke-width="2" stroke-linecap="round"><path d="M4 7h16M4 12h16M4 17h16"/></svg></span></div>')


def footer(phone: bool = False) -> str:
    links = ''.join(f'<a href="#{l.lower()}" style="font-size:12px;font-weight:700;letter-spacing:.08em;text-transform:uppercase">{l}</a>' for l in ['About', 'Sources', 'Discord', 'GitHub'])
    if phone:
        return (f'<div style="margin-top:auto;border-top:1px solid var(--line-soft);padding:20px 18px;display:flex;flex-direction:column;gap:10px">'
                f'<span style="display:flex;gap:16px">{links}</span><span style="font-size:12px;color:var(--muted)">Not affiliated with Blizzard Entertainment.</span></div>')
    return (f'<div style="margin-top:auto;border-top:1px solid var(--line-soft);padding:28px 48px;display:flex;align-items:center;justify-content:space-between">'
            f'<span style="display:flex;align-items:center;gap:10px"><img src="{SEAL}" alt="" style="width:24px;height:24px"><span class="display" style="font-size:14px;font-weight:700;letter-spacing:.1em;color:var(--strong)">FOREVER SIXTY</span></span>'
            f'<span style="display:flex;gap:20px">{links}</span><span style="font-size:12px;color:var(--muted)">Not affiliated with Blizzard Entertainment.</span></div>')


def page(body: str, width: int) -> str:
    return HEAD + f'<div style="width:{width}px;min-height:100vh;display:flex;flex-direction:column;background:var(--bg)">{body}</div></body></html>'


def write(name: str, html: str) -> Path:
    RENDERS.mkdir(exist_ok=True)
    out = RENDERS / name
    out.write_text(html)
    return out
