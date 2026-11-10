"""The Mac disk image's window background, at 1x and 2x (tools/package-mac-companion.sh lays the window out on it).

    python3 tools/art/dmg_background.py

Writes companion/packaging/macos/dmg-background.png (660 x 400) and dmg-background@2x.png (1320 x 800): the cream
ground, the painted seal and "Forever Sixty" in Marcellus small at the top left, and one brass arrow from where the
app sits to where Applications sits. Nothing else.

The ground is light on purpose. Finder draws icon labels over a background picture in dark ink in light and dark
appearance alike (checked on macOS 26, 2026-11-10: a dark ground left "Forever Sixty" unreadable in both), so the
labels need a light ground, as Discord's white one. The cream is the site's display-type token, the ink its card
brown-black, the arrow the legacy deep gold: all from design/tokens/tokens.json. The build is deterministic.
"""
import json
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

REPO = Path(__file__).resolve().parents[2]
TOKENS = REPO / 'design/tokens/tokens.json'
SEAL = REPO / 'design/assets/logo/lx-seal-v4.png'
FONT = REPO / 'api/internal/brand/fonts/Marcellus-Regular.ttf'
OUT = REPO / 'companion/packaging/macos'

# The window, in points; package-mac-companion.sh places the icons on these same numbers.
WINDOW = (660, 400)
APP_X, DROP_X, ICON_Y = 180, 480, 200
ICON_PT = 128
# The wordmark: the seal and the name, small, at the top left.
MARK_AT = (24, 22)
SEAL_PT = 28
NAME_PT = 19
NAME_GAP = 9
# The arrow runs between the two icons with this much air at each end, at the icons' centre line.
ARROW_AIR = 34
ARROW_STROKE = 3.5
ARROW_HEAD = 13
SUPERSAMPLE = 4


def token(*path: str) -> str:
    """A colour token's web value (round 5 leads: platformValue.web where it differs, else $value)."""
    node = json.loads(TOKENS.read_text())['color']
    for key in path:
        node = node[key]
    web = node.get('$extensions', {}).get('gg.foreversixty', {}).get('platformValue', {}).get('web')
    return web or node['$value']


def render(scale: int) -> Image.Image:
    s = scale * SUPERSAMPLE
    w, h = WINDOW[0] * s, WINDOW[1] * s
    im = Image.new('RGB', (w, h), token('forever', 'cream'))
    draw = ImageDraw.Draw(im)

    seal = Image.open(SEAL).convert('RGBA').resize((SEAL_PT * s, SEAL_PT * s), Image.LANCZOS)
    x, y = MARK_AT[0] * s, MARK_AT[1] * s
    im.paste(seal, (x, y), seal)
    font = ImageFont.truetype(str(FONT), NAME_PT * s)
    ink = token('base', 'raised')
    draw.text((x + (SEAL_PT + NAME_GAP) * s, y + SEAL_PT * s / 2), 'Forever Sixty', font=font, fill=ink, anchor='lm')

    brass = token('legacy', 'goldDeep')
    half = ICON_PT / 2
    x0, x1 = (APP_X + half + ARROW_AIR) * s, (DROP_X - half - ARROW_AIR) * s
    cy, stroke, head = ICON_Y * s, round(ARROW_STROKE * s), ARROW_HEAD * s
    draw.line((x0, cy, x1 - stroke, cy), fill=brass, width=stroke)
    draw.line((x1 - head, cy - head, x1, cy), fill=brass, width=stroke, joint='curve')
    draw.line((x1 - head, cy + head, x1, cy), fill=brass, width=stroke, joint='curve')
    for px, py in ((x0, cy), (x1 - head, cy - head), (x1 - head, cy + head), (x1, cy)):
        r = stroke / 2
        draw.ellipse((px - r, py - r, px + r, py + r), fill=brass)  # round caps and the tip
    return im.resize((WINDOW[0] * scale, WINDOW[1] * scale), Image.LANCZOS)


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    render(1).save(OUT / 'dmg-background.png', optimize=True)
    render(2).save(OUT / 'dmg-background@2x.png', optimize=True)
    print(f'wrote {OUT.relative_to(REPO)}/dmg-background.png and @2x')


if __name__ == '__main__':
    main()
