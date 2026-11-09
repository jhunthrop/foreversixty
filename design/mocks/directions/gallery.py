"""Builds the round-4 gallery: python3 design/mocks/directions/gallery.py [out dir]

Default out dir ~/Desktop/forever-design-directions: index.html with png/, now/ and blizzard/ copied next to it
(blizzard/ holds their captures and the two study sheets; it is git-ignored and only ever travels to this local page).
Rounds 1 to 3 are retired; git holds them (7156c5f23, 0edfd08a1).
"""
import json
import shutil
import sys
from html import escape
from pathlib import Path

from style import TOKENS

HERE = Path(__file__).resolve().parent
OUT = Path(sys.argv[1]) if sys.argv[1:] else Path.home() / 'Desktop' / 'forever-design-directions'
PAGES = [('Home', 'g-home', 'home', 'screen-1440-0.png'), ('Best in slot', 'g-bis', 'bis', 'screen-1440-3.png'),
         ('Tier list', 'g-tiers', 'tiers', 'screen-1440-1.png')]
STUDY = ['screen-1440-0.png', 'screen-1440-1.png', 'screen-1440-3.png', 'screen-1440-5.png', 'screen-1440-6.png', 'screen-1440-7.png', 'screen-390-1.png']
# ai-slop-reviewer's list (.claude/agents/ai-slop-reviewer.md), checked against the live site and round 4.
TELLS = ['Glowing or gradient buttons; hover glows', 'Decorative gradients, gradient text', 'Near-black plus one gold accent doing all the work',
         'Cream, serif display and terracotta', 'Big thin serif slogan hero', 'Uppercase letter-spaced eyebrows',
         'Grey explainer under every heading', 'Same rounded card on every block', 'Icon-per-bullet grids, emoji markers',
         '01/02/03 that is not a sequence', 'Everything centred; symmetric three columns', 'Glass, blur, blobs, noise as decoration',
         'Skeletons that never resolve', 'Inter, Space Grotesk, Poppins by default', 'Generated imagery', 'Stat tiles that are not the point',
         'Pills on everything', 'Generic marketing lines', 'Vague benefit copy', 'Fake social proof', 'Em dashes, narrated obvious things',
         'Qualifiers on numbers and labels (tell 22)']
NOW_FAILS = {0: 'Glowing gold Sign in button', 1: 'Purple-gradient countdown card', 2: 'Near-black with gold doing the work',
             4: '"Play your class better." as a thin serif hero', 5: 'FIRST RAIDS OPEN, BEST IN SLOT', 6: 'Grey helper line under headings',
             7: 'Same rounded card on every BiS group', 12: 'Signed-out feed skeleton for ~1.6 s', 15: 'Countdown digits as tiles',
             16: 'Premium pill, caps feed tags', 21: '"This set, simmed 7 Nov", "Simmed on import"'}
NEW_NOTES = {7: 'Soft corners on every panel are one radius scale from the nav; the cream frame is on two or three featured panels a page, not all',
             11: 'The ground is the owner\'s painted tile (sd 2.7, inside their 2 to 5); every list sits on a flat surface',
             14: 'The art is generated, the owner\'s explicit written choice (2026-11-09); background only, never captioned'}
SHOTS = ['Westfall, Redridge or Duskwood at dusk from a road, UI hidden: the framed plate on the signed-out home and the guides',
         "The guild at Onyxia's Lair on 9 Dec, UI hidden: the raid card's history",
         'The character sheet with the addon tab open, full gear, 1x UI scale: the BiS empty state',
         'A tooltip of an item with crit or hit rating, and one weapon tooltip: to copy the Forever wording and the speed and damage lines',
         'The world map hovering zones along the Alliance road: to confirm the level ranges',
         'Wailing Caverns and Deadmines entrances: the itinerary stop plates']
CSS = '''
:root{--bg:#141516;--fg:#e6e3dc;--muted:#9c978d;--line:#2d2f32;--pass:#7fc28a;--watch:#e0b25c;--fail:#ef7a6e}
*{box-sizing:border-box}body{margin:0;background:var(--bg);color:var(--fg);font:15px/22px -apple-system,Segoe UI,Arial,sans-serif}
main{max-width:1560px;margin:0 auto;padding:28px 24px 80px}h1{font-size:28px;line-height:34px;margin:0 0 8px}
h2{font-size:22px;margin:52px 0 10px;padding-top:18px;border-top:1px solid var(--line)}h3{font-size:16px;margin:22px 0 8px}
p{max-width:960px;margin:0 0 10px}a{color:#9cc8ff}.muted{color:var(--muted)}ul.f{max-width:960px;padding-left:18px}ul.f li{margin:4px 0}
table{border-collapse:collapse;font-size:14px;margin-top:8px}td,th{border-bottom:1px solid var(--line);padding:5px 10px;text-align:left;vertical-align:top}
.pass{color:var(--pass)}.watch{color:var(--watch)}.fail{color:var(--fail)}
.three{display:grid;grid-template-columns:repeat(3,1fr);gap:14px}.two{display:grid;grid-template-columns:repeat(4,1fr);gap:14px}
figure{margin:0}figure img{width:100%;display:block;border:1px solid var(--line)}figcaption{font-size:13px;color:var(--muted);margin:4px 0}
.rv{display:grid;grid-template-columns:repeat(3,1fr);gap:16px}.rv div{border-top:2px solid var(--line);padding-top:8px;font-size:14px}.rv b{display:block}
.study{display:grid;grid-template-columns:repeat(4,1fr);gap:12px}.sw{display:inline-block;width:18px;height:18px;border:1px solid #555;vertical-align:middle;margin-right:8px}
@media(max-width:900px){.three,.two,.rv,.study{grid-template-columns:1fr}}
'''


def fig(src: str, cap: str) -> str:
    return f'<figure><a href="{src}"><img src="{src}" alt="{escape(cap)}" loading="lazy"></a><figcaption>{escape(cap)}</figcaption></figure>'


def verdicts() -> str:
    rv = json.loads((HERE / 'reviews.json').read_text()) if (HERE / 'reviews.json').exists() else {}
    cells = []
    for who in ('ux-designer', 'wow-player', 'ai-slop-reviewer'):
        rounds = ''.join(f'<p><b>{escape(r["round"])}</b>{escape(r["text"])}</p>' for r in rv.get(who, []))
        cells.append(f'<div><b>{who}</b>{rounds or "<p class=muted>pending</p>"}</div>')
    return f'<div class="rv">{"".join(cells)}</div>'


def checklist() -> str:
    rows = []
    for i, t in enumerate(TELLS):
        now = f'<td class="fail">fail: {escape(NOW_FAILS[i])}</td>' if i in NOW_FAILS else '<td class="pass">pass</td>'
        new = f'<td class="watch">watch: {escape(NEW_NOTES[i])}</td>' if i in NEW_NOTES else '<td class="pass">pass</td>'
        rows.append(f'<tr><td>{i + 1}. {escape(t)}</td>{now}{new}</tr>')
    return f'<table><tr><th>Tell</th><th>Now (live)</th><th>Round 5</th></tr>{"".join(rows)}</table>'


def tokens() -> str:
    rows = ''.join(f'<tr><td><span class="sw" style="background:{v}"></span>{escape(k)}</td><td>{escape(v)}</td></tr>' for k, v in TOKENS.items())
    return f'<table><tr><th>Token</th><th>Value (theirs, exactly)</th></tr>{rows}</table>'


MOBILE = [('home', 'Home (raid day)'), ('raid', 'Raid'), ('gear', 'Gear'), ('activity', 'Activity'), ('menu', 'Profile drawer'),
          ('signup', 'Raid sign-up: Late, the minutes inline'), ('item', 'Item sheet'), ('lock', 'Live Activity / ongoing notification')]
MOBILE_BEFORE = 'reviews-app-flows'  # the 2026-11-08 app-flows renders, copied in by build()


def paperdoll_section() -> str:
    """The BiS paperdoll on the owner's frame: public and signed in at 1440 and 390, the mobile Gear tab, the cleaned frame."""
    shots = (f'<h3>Best in slot, public (the selected slot\'s tooltip in the centre)</h3><div class="three">{fig("png/g-bis-frame-1440.png", "1440")}'
             f'{fig("png/g-bis-frame-390.png", "390: the frame\'s header and the two-column list (slots under 36 px)")}{fig("png/g-bis-1440.png", "Round 5 before it")}</div>'
             f'<h3>Best in slot, signed in (the character\'s Battle.net render in the centre)</h3><div class="three">'
             f'{fig("png/g-bis-frame-signed-1440.png", "1440, the render placeholder standing in")}{fig("png/g-bis-frame-signed-390.png", "390")}</div>'
             f'<h3>Mobile Gear tab</h3><div class="two">{fig("png/m-ios-gear-frame.png", "iPhone 15")}{fig("png/m-android-gear-frame.png", "Pixel 7")}</div>'
             f'<h3>The frame, cleaned</h3><div class="three">{fig("paperdoll/frame@1x.png", "frame@1x.png, real alpha, clean centre")}</div>')
    return (f'<h2 id="paperdoll">Paperdoll</h2><p>The owner\'s frame (generated, the owner\'s choice) cleaned: the baked checkerboard cut to real '
            f'alpha, the fake writing in the centre replaced with the blue band. Every slot, the medallion, the title bar and the centre measured as '
            f'percentages of the frame (paperdoll/slots.json), so the item icons sit inside the slots at any width. Empty slots keep the frame\'s glyph; '
            f'the two-hander shows greyed in the off hand.</p>{shots}')


def mobile_section() -> str:
    """Mobile: each screen on iPhone 15 and Pixel 7 in round 5, beside the app-flows render it restyles (the same flow)."""
    rows = ''.join(
        f'<h3>{label}</h3><div class="two">{fig(f"{MOBILE_BEFORE}/ios-{before}.png", "Before, iOS")}{fig(f"png/m-ios-{key}.png", "Round 5, iPhone 15")}'
        f'{fig(f"{MOBILE_BEFORE}/android-{before}.png", "Before, Android")}{fig(f"png/m-android-{key}.png", "Round 5, Pixel 7")}</div>'
        for (key, label), before in zip(MOBILE, ('home-raid-day', 'raid', 'gear', 'activity', 'menu', 'raid-night-signup', 'item-sheet', 'live-activity')))
    return (f'<h2 id="mobile">Mobile</h2><p>The iOS and Android apps in the round-5 look. The flows, tap counts and every element are the app-flows '
            f'spec\'s (design/specs/2026-11-08-app-flows.md, its tap table and purpose audit; the content is imported from its mock). Only the '
            f'visual language changes. Each platform keeps its conventions: iOS navigation and tab bars, sheets with a grabber, the lock screen '
            f'Live Activity; Android Material 3 top app bar, navigation bar with the pill indicator, ModalNavigationDrawer, modal bottom sheet, '
            f'segmented buttons and the ongoing notification.</p>{rows}')


def build() -> Path:
    OUT.mkdir(parents=True, exist_ok=True)
    for old in ('png', 'now', 'blizzard', 'paperdoll'):
        shutil.rmtree(OUT / old, ignore_errors=True)
        shutil.copytree(HERE / old, OUT / old)
    shutil.rmtree(OUT / MOBILE_BEFORE, ignore_errors=True)
    shutil.copytree(HERE.parents[1] / 'reviews/2026-11-08-app-flows', OUT / MOBILE_BEFORE, ignore=shutil.ignore_patterns('*.md', '*.json'))
    round5 = ''.join(
        f'<h3>{name}</h3><div class="three">{fig(f"blizzard/{bz}", "Blizzard, WoW: Forever (reference only)")}'
        f'{fig(f"png/r4/{key}-1440.png", "Round 4")}{fig(f"png/{key}-1440.png", "Round 5")}</div>'
        f'<div class="two" style="margin-top:12px">{fig(f"now/{now}-390-full.png", "Now, 390")}{fig(f"png/r4/{key}-390.png", "Round 4, 390")}'
        f'{fig(f"png/{key}-390.png", "Round 5, 390")}</div>'
        for name, key, now, bz in PAGES)
    pages = ''.join(
        f'<h3>{name}</h3><div class="three">{fig(f"blizzard/{bz}", "Blizzard, WoW: Forever (reference only)")}'
        f'{fig(f"now/{now}-1440-full.png", "Now, live site")}{fig(f"png/{key}-1440.png", "Round 5")}</div>'
        for name, key, now, bz in PAGES)
    study = ''.join(fig(f'blizzard/{s}', s) for s in STUDY)
    body = (f'<main><h1>Forever Sixty, round 5</h1><p class="muted">9 Nov 2026. Mocks only. Spec: design/specs/2026-11-09-directions.md. '
            f'Round 4 (Blizzard\'s Forever colours, fonts and floating navigation, crisp surfaces) plus the owner\'s round-5 asks: soft corners '
            f'on one radius scale taken from the nav (10 px panels, 6 px buttons), a rounded cream frame on two or three featured panels a page, '
            f'and the owner\'s art set. Every number, item, tooltip line and icon is real repo data.</p>'
            f'<p><a href="#round5">Round 5</a> · <a href="#paperdoll">Paperdoll</a> · <a href="#artset">Art set</a> · <a href="#mobile">Mobile</a> · <a href="#pages">Now and round 5</a> · <a href="#reviews">Reviewers</a> · '
            f'<a href="#tokens">Colours</a> · <a href="#fonts">Display font</a> · <a href="#tells">No AI tells</a> · <a href="#study">Blizzard study</a> · '
            f'<a href="#shots">Screenshots needed</a></p>'
            f'<h2 id="round5">Round 5: Blizzard, round 4 and round 5</h2>{round5}<p class="muted" style="margin-top:12px"><a href="png/g-home-2000.png">Home at 2000</a></p>'
            f'<h2 id="artset">The owner\'s art set</h2><p>Generated in the Classic style; the owner\'s explicit choice for art assets (2026-11-09). Background only, never '
            f'captioned or implied to be a screenshot. Hero: the camp road, its dusk glow centred and the lantern camp at the window edge, under a flat '
            f'night veil (title 12:1, the DPS figure 7.4:1 or better around it). Raid plaque: the cave mouth in the plaque\'s right edge, clear of the text '
            f'(11.9:1). Blue bands: the owner\'s painted band, stretched with cover so its edges never meet. Ground: the painted brown tile, made periodic (no seam, no repeat), its soft field clamped and its brush edges damped '
            f'(sd 2.8 over 200 px, as Blizzard\'s ground). social-banner-valley.jpeg is byte-identical to the round-4 dusk road, so the link preview is unchanged.</p>'
            f'<div class="three">{fig("png/art-band-1440.png", "Hero band at 1440")}{fig("png/art-band-2000.png", "Hero band at 2000")}'
            f'{fig("png/art-band-390.png", "Hero band at 390")}</div><div class="three" style="margin-top:12px">{fig("png/og-home.png", "Link preview, 1200 x 630")}'
            f'{fig("png/discord-banner.png", "Discord banner (also for guild pages)")}{fig("png/app-icon.png", "App icon: the flat seal on the owner\'s field")}</div>'
            f'<div class="three" style="margin-top:12px">{fig("png/ground-seam-check.png", "The ground tile, 2 x 2: no seam at the cross")}</div>'
            + paperdoll_section() + mobile_section() +
            f'<h2 id="pages">Now and round 5</h2>{pages}'
            f'<h2 id="reviews">Reviewers</h2>{verdicts()}'
            f'<h2 id="tokens">Colours, taken from their page</h2><p>Computed CSS values from the live page (blizzard/palette.json) or the mean of a flat region of '
            f'a capture, with the source of each in style.py. Their headings are gradient-filled text (#f3eee2 to #e2c2ab); ours are flat #f3eee2, '
            f'because gradient text is an AI tell.</p>{tokens()}'
            f'<h2 id="fonts">Display font: their BlizQuadrata against open stand-ins</h2><p>Marcellus is closest in build and flare once tracked to -0.015em; '
            f'their stroke is heavier. ITC Friz Quadrata (Monotype) is the licensed path: MyFonts lists single styles from US$43.99 and the family of four at '
            f'US$174.99 for desktop; the web licence is priced by monthly pageviews at checkout. In style.py the swap is one line (DISPLAY_FONT).</p>'
            f'{fig("blizzard/fonts.png", "Their heading crops, then eight Google Fonts at the same size")}'
            f'<h2 id="texture">The brown ground: theirs (left) and ours (right), same zoom</h2><p>Ours is procedural (feTurbulence, baseFrequency 0.004 and 0.6), '
            f'matched in hue and spread: luminance sd 3.1 to 4.4 against their 2.1 to 5.3. It does not reproduce the painted stone speckle of theirs. '
            f'If that matters, a CC0 photographed surface (Poly Haven or ambientCG), tinted to #67563a, gets closer without breaking the no-generated-imagery rule. '
            f'Their blue band measures flat (sd under 1), so ours is flat too.</p>{fig("blizzard/texture-compare.png", "Theirs left, ours right")}'
            f'<h2 id="tells">No AI tells</h2><p>The owner\'s rule: "It should look like a real developer made it." Tell 22, owner verbatim: "212.4 DPS, this set" '
            f'is "a dead giveaway it\'s AI. Adding unnecessary explanations. All it need to say is DPS".</p>{checklist()}'
            f'<h2 id="study">Blizzard\'s Forever page, captured</h2><p>Playwright at 1440 and 390. Reference only: none of their logos, art or ornament '
            f'appears in our pages.</p><div class="study">{study}</div>'
            f'<h2 id="shots">Screenshots the owner needs to take in game</h2><ul class="f">'
            + ''.join(f'<li>{escape(s)}</li>' for s in SHOTS) + '</ul></main>')
    html = (f'<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">'
            f'<title>Forever Sixty round 5</title><style>{CSS}</style></head><body>{body}</body></html>')
    (OUT / 'index.html').write_text(html)
    return OUT / 'index.html'


if __name__ == '__main__':
    print(build())
