"""The Mac disk image's window background, at 1x and 2x (tools/package-mac-companion.sh lays the window out on it).

    python3 tools/art/dmg_background.py

Writes companion/packaging/macos/dmg-background.png (660 x 400) and dmg-background@2x.png (1320 x 800): the
companion window's own head, the site's painted band under its veil with a black edge, the painted seal and
"Forever Sixty" in Marcellus cream on it; under it the cream ground, and one short teal arrow from where the app
sits to where Applications sits. Nothing else.

The ground under the icons is light on purpose. Finder draws icon labels over a background picture in dark ink in
light and dark appearance alike (checked on macOS 26, 2026-11-10: a dark ground left "Forever Sixty" unreadable in
both), so the labels sit on the brand's cream (--color-cream, every display heading's ink), and the dark brand
lives in the band above them, as the companion window opens. Every colour is a token from
design/tokens/tokens.json (band, veilBand, edge, cream, teal) and the band art is the window's own
(companion/ui/art/band-1200.webp, the site's). The build is deterministic.

Also writes companion/packaging/macos/dmg-layout.sh, the window and icon places in the shell's terms, so the image
and the layout share one set of numbers.
"""
import json
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

REPO = Path(__file__).resolve().parents[2]
TOKENS = REPO / 'design/tokens/tokens.json'
SEAL = REPO / 'design/assets/logo/lx-seal-v4.png'
BAND = REPO / 'companion/ui/art/band-1200.webp'
FONT = REPO / 'api/internal/brand/fonts/Marcellus-Regular.ttf'
OUT = REPO / 'companion/packaging/macos'

# The window's content, in points. The icons sit a little above the middle so each one with its label below
# reads centred. dmg-layout.sh hands these numbers to package-mac-companion.sh, which adds the title bar.
WINDOW = (660, 400)
TITLE_BAR_PT = 32
APP_X, DROP_X, ICON_Y = 180, 480, 188
ICON_PT = 128
# The band across the top, as the window's (--band-h is 60 px there), and the wordmark on it: the seal and the
# name at the window's 24 px gutter, the name at the window's wordmark size (--text-title-phone).
BAND_PT = 60
BAND_FOCUS_Y = 0.3  # the window's `center 30%`
EDGE_PT = 1
MARK_X = 24
SEAL_PT = 32
NAME_PT = 21
NAME_GAP = 10
# The arrow runs between the two icons with this much air at each end, at the icons' centre line.
ARROW_AIR = 70
ARROW_STROKE = 4
ARROW_HEAD = 12
SUPERSAMPLE = 4


def token(*path: str) -> str:
    """A colour token's web value (round 5 leads: platformValue.web where it differs, else $value)."""
    node = json.loads(TOKENS.read_text())['color']
    for key in path:
        node = node[key]
    web = node.get('$extensions', {}).get('gg.foreversixty', {}).get('platformValue', {}).get('web')
    return web or node['$value']


def rgba(value: str) -> tuple[int, int, int, int]:
    """A token colour, '#rrggbb' or 'rgba(r, g, b, a)', as an RGBA tuple."""
    if value.startswith('#'):
        return (*(int(value[i:i + 2], 16) for i in (1, 3, 5)), 255)
    r, g, b, a = (p.strip() for p in value[value.index('(') + 1:-1].split(','))
    return int(r), int(g), int(b), round(float(a) * 255)


def band(w: int, h: int) -> Image.Image:
    """The painted band cropped to w by h around its focus, under the band veil, as the window draws it."""
    art = Image.open(BAND).convert('RGB')
    scale = w / art.width
    art = art.resize((w, round(art.height * scale)), Image.LANCZOS)
    top = max(0, min(art.height - h, round((art.height - h) * BAND_FOCUS_Y)))
    out = Image.new('RGB', (w, h), token('forever', 'band'))
    out.paste(art.crop((0, top, w, top + h)), (0, 0))
    veil = Image.new('RGBA', (w, h), rgba(token('forever', 'veilBand')))
    return Image.alpha_composite(out.convert('RGBA'), veil).convert('RGB')


def render(scale: int) -> Image.Image:
    s = scale * SUPERSAMPLE
    w, h = WINDOW[0] * s, WINDOW[1] * s
    im = Image.new('RGB', (w, h), token('forever', 'cream'))
    draw = ImageDraw.Draw(im)

    band_h = BAND_PT * s
    im.paste(band(w, band_h), (0, 0))
    draw.rectangle((0, band_h, w, band_h + EDGE_PT * s - 1), fill=token('forever', 'edge'))

    seal = Image.open(SEAL).convert('RGBA').resize((SEAL_PT * s, SEAL_PT * s), Image.LANCZOS)
    x, y = MARK_X * s, (band_h - SEAL_PT * s) // 2
    im.paste(seal, (x, y), seal)
    font = ImageFont.truetype(str(FONT), NAME_PT * s)
    draw.text((x + (SEAL_PT + NAME_GAP) * s, band_h / 2), 'Forever Sixty', font=font, fill=token('forever', 'cream'), anchor='lm')

    arrow = token('forever', 'teal')
    half = ICON_PT / 2
    x0, x1 = (APP_X + half + ARROW_AIR) * s, (DROP_X - half - ARROW_AIR) * s
    cy, stroke, head = ICON_Y * s, round(ARROW_STROKE * s), ARROW_HEAD * s
    draw.line((x0, cy, x1 - stroke, cy), fill=arrow, width=stroke)
    draw.line((x1 - head, cy - head, x1, cy), fill=arrow, width=stroke, joint='curve')
    draw.line((x1 - head, cy + head, x1, cy), fill=arrow, width=stroke, joint='curve')
    for px, py in ((x0, cy), (x1 - head, cy - head), (x1 - head, cy + head), (x1, cy)):
        r = stroke / 2
        draw.ellipse((px - r, py - r, px + r, py + r), fill=arrow)  # round caps and the tip
    return im.resize((WINDOW[0] * scale, WINDOW[1] * scale), Image.LANCZOS)


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    render(1).save(OUT / 'dmg-background.png', optimize=True)
    render(2).save(OUT / 'dmg-background@2x.png', optimize=True)
    (OUT / 'dmg-layout.sh').write_text(
        '# Written by tools/art/dmg_background.py; package-mac-companion.sh sources it. Points; the window\n'
        '# height adds the title bar to the background\'s, since Finder\'s bounds include it.\n'
        f'DMG_WINDOW=({WINDOW[0]} {WINDOW[1] + TITLE_BAR_PT})\n'
        f'DMG_APP=({APP_X} {ICON_Y})\n'
        f'DMG_DROP=({DROP_X} {ICON_Y})\n')
    print(f'wrote {OUT.relative_to(REPO)}/dmg-background.png and @2x')


if __name__ == '__main__':
    main()
