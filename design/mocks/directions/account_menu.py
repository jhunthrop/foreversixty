"""Mock batch 4b: the account menu, the panel the header's account control opens (plan 2026-11-09-redesign).

Data and copy, never invented: web/tests/e2e/fixtures/home/me-60.json (Hunthrop#1894 and Thandrel), the selector e2e
fixtures in web/tests/e2e/support/selector.ts (Obnoxious Yell) and selector-restored.spec.ts (Frostbyte), and the
menu's own words in web/src/lib/character-selector/copy.ts. The icons are the shared glyph table, web/src/lib/ui/glyphs.ts,
read here so the mock and the site draw the same marks.

Round 5: the overlay panel on the plate radius with the panel edge and the overlay shadow, and the framed panel's cream inset
line; the ringed medallion and the class-coloured name on the panel; the selected row is the compact rank list's (accent
tint, 3 px accent rail); Premium is the teal edge colour flat, as a pill. No grey explainer under any row (owner: no
explanations); a row keeps a line only when it is data (the plan, the character's spec, level, ruleset, source and age).
"""
import json
import re
from html import escape

import data as D
from common import write
from style import header, medallion, page, titlebar

ME = json.loads((D.ROOT / 'web/tests/e2e/fixtures/home/me-60.json').read_text())
GLYPH_TS = (D.ROOT / 'web/src/lib/ui/glyphs.ts').read_text()


def glyph(name: str, size: int = 18) -> str:
    """One glyph from web/src/lib/ui/glyphs.ts, as ui/Glyph.svelte draws it."""
    match = re.search(rf"^\s+{name}:\s*\n?\s*'([^']+)'", GLYPH_TS, re.M)
    inner = match.group(1) if match else ''
    return (f'<svg width="{size}" height="{size}" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" '
            f'stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">{inner}</svg>')


# (name, class slug, spec, level, ruleset, source, age, current): Forever has no realms
CHARACTERS = [
    ('Thandrel', 'warrior', 'Fury', 60, 'Normal', 'Battle.net', '2 days ago', False),
    ('Obnoxious Yell', 'warrior', 'Fury', 60, 'Normal', 'Addon', '14 min ago', True),
    ('Frostbyte', 'mage', 'Frost', 42, 'Normal', 'Addon', '3 days ago', False),
]

MENU_CSS = '''
.stage{position:relative;min-height:760px}
.am{position:absolute;right:max(16px,calc((100vw - 1344px)/2 + 42px));top:84px;width:380px;z-index:5;font-family:var(--body);
  background:#170e09;border:1px solid #000;border-radius:10px;box-shadow:0 0 0 1px rgba(130,112,91,.55),0 12px 32px rgba(0,0,0,.6)}
.am::before{content:"";position:absolute;inset:6px;border:1px solid rgba(243,238,226,.42);border-radius:6px;pointer-events:none}
.am .in{position:relative;padding:8px 10px calc(10px + env(safe-area-inset-bottom))}
.am .rule{border-top:1px solid #2c241f;margin:6px 8px}
.am .head{display:flex;align-items:center;gap:12px;padding:12px 12px 10px}
.am .tile{display:grid;place-items:center;flex-shrink:0;width:32px;height:32px;border-radius:6px;background:rgba(255,236,113,.12);color:#ffec71}
.am .tile.q{background:rgba(255,255,255,.05);color:#8b8784}
.am .who{flex:1;min-width:0;display:flex;flex-direction:column;line-height:1.3}
.am .who b{font:700 15px var(--body);color:#fff}.am .who span{font-size:13px;color:#8b8784}
.am .pill{display:inline-flex;align-items:center;height:22px;padding:0 10px;border-radius:999px;background:#408076;color:#fff;font:600 12px var(--body)}
.am .lab{display:flex;align-items:center;justify-content:space-between;padding:4px 12px 4px;min-height:32px}
.am .lab span{font:600 12px var(--body);color:#8b8784}.am .lab a,.am .restored a{display:inline-flex;align-items:center;min-height:44px}.am .lab a{font:600 13px var(--body);color:#ffec71}
.am ul{list-style:none;margin:0;padding:0}
.am li.ch{position:relative;display:grid;grid-template-columns:36px minmax(0,1fr);gap:12px;align-items:center;min-height:60px;padding:6px 12px;border-radius:6px}
.am li.ch.cur{background:color-mix(in srgb,#0093b2 16%,transparent)}
.am li.ch.cur::before{content:"";position:absolute;left:0;top:6px;bottom:6px;width:3px;border-radius:2px;background:#0093b2}
.am .medal{position:relative;display:block}.am .medal img{position:absolute;border-radius:50%}.am .medal svg{position:absolute;inset:0;width:100%;height:100%}
.am .r1{font:600 15px/1.25 var(--body)}.am .r2{font-size:13px;line-height:1.4;color:#e3dfd9}.am .r3{font-size:12px;line-height:1.4;color:#8b8784}
.am .restored{display:flex;align-items:center;justify-content:space-between;min-height:40px;padding:0 12px;font-size:12px;color:#8b8784}
.am .restored a{font:600 13px var(--body);color:#ffec71}
.am .row{display:flex;align-items:center;gap:12px;min-height:44px;padding:6px 12px;border-radius:6px;color:#fff;font:600 15px var(--body)}
.am .row:hover{background:rgba(255,255,255,.05)}
.am .row .t{flex:1}.am .row .s{font:400 13px var(--body);color:#8b8784}.am .row .go{display:flex;flex-shrink:0;color:#8b8784}.am .row .go svg{width:16px;height:16px}
.am .row.quiet{color:#e3dfd9;font-weight:400}
@media(max-width:1023px){
  .stage{min-height:844px}
  .scrim{position:fixed;inset:0;background:rgba(0,0,0,.55);z-index:4}
  .am{position:fixed;left:0;right:0;bottom:0;top:auto;width:auto;border-bottom:0;border-radius:10px 10px 0 0;max-height:88vh;overflow:auto}
  .am::before{border-bottom:0;border-radius:6px 6px 0 0;bottom:-1px}
}
.am .grab{display:none}
@media(max-width:1023px){.am .grab{display:block;width:36px;height:4px;border-radius:2px;background:#3a302a;margin:2px auto 6px}}
'''


def character_row(name: str, slug: str, spec: str, level: int, ruleset: str, source: str, age: str, current: bool) -> str:
    colour = D.CLASS_COLOR[slug]
    return (f'<li class="ch{" cur" if current else ""}">{medallion(slug, 36)}<span>'
            f'<span class="r1" style="color:{colour}">{escape(name)}</span><br>'
            f'<span class="r2">{escape(spec)} {escape(slug.title())} · {level} · {escape(ruleset)}</span><br>'
            f'<span class="r3">{escape(source)} · {escape(age)}</span></span></li>')


def action(icon: str, title: str, side: str = '', quiet: bool = False) -> str:
    tile = f'<span class="tile{" q" if quiet else ""}">{glyph(icon)}</span>'
    side_html = f'<span class="s">{escape(side)}</span>' if side else ''
    chev = '' if quiet else f'<span class="go">{glyph("chevron", 16)}</span>'
    return f'<a class="row{" quiet" if quiet else ""}" href="#">{tile}<span class="t">{escape(title)}</span>{side_html}{chev}</a>'


def menu() -> str:
    battletag = ME['data']['user']['battletag'] if 'data' in ME else ME['user']['battletag']
    rows = ''.join(character_row(*c) for c in CHARACTERS)
    return (
        '<div class="scrim"></div><div class="am" role="dialog"><div class="in"><span class="grab"></span>'
        f'<div class="head"><span class="tile">{glyph("user")}</span><span class="who"><b>{escape(battletag)}</b>'
        '<span>Signed in with Battle.net</span></span><span class="pill">Premium</span></div>'
        '<div class="rule"></div>'
        '<div class="lab"><span>Your characters</span><a href="#">Manage</a></div>'
        f'<ul>{rows}</ul>'
        '<div class="restored"><span>Restored your last character</span><a href="#">Forget</a></div>'
        '<div class="rule"></div>'
        + action('copy', 'Paste an export')
        + '<div class="rule"></div>'
        + action('chart', 'Character page')
        + action('gear', 'Account')
        + action('crown', 'Plan', 'Premium · $50/yr')
        + action('download', 'Get set up')
        + action('signout', 'Sign out', quiet=True)
        + '</div></div>'
    )


def menu_page() -> str:
    top = header('Obnoxious Yell', 'warrior', 60, 'Tier list')
    band = titlebar('warrior', 'Fury Warrior tier list', 'Level 60', top=top)
    body = f'<div class="stage">{band}{menu()}</div>'
    return page('Forever Sixty, account menu', MENU_CSS, body)


def build() -> None:
    write('b4-account-menu', menu_page())


if __name__ == '__main__':
    build()
