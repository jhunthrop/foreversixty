# design/mocks/2026-10-29-site-alignment/compose.py
# Before | after boards for the four mocked routes: the first 1400 px at 1440 side by side, and the
# first 1800 px at 390 side by side, same scale. Run from this folder after both capture specs.
from pathlib import Path
from PIL import Image, ImageDraw

HERE = Path(__file__).parent
BEFORE = HERE.parent.parent / 'reviews' / '2026-10-29-site-alignment'
ROUTES = ['home-signed-in', 'bis-warrior-fury', 'guide-warrior-fury', 'planner']
GAP = 40
LABEL_H = 36
INK = (229, 185, 85)
GROUND = (20, 20, 24)


def board(route: str, width: int, height: int) -> None:
    before = Image.open(BEFORE / f'before-{route}-{width}.png')
    after = Image.open(HERE / f'after-{route}-{width}.png')
    out = Image.new('RGB', (width * 2 + GAP, height + LABEL_H), GROUND)
    draw = ImageDraw.Draw(out)
    for i, (name, im) in enumerate([('BEFORE (built fixture site)', before), ('AFTER (spec 2026-10-29-site-alignment)', after)]):
        x = i * (width + GAP)
        draw.text((x + 8, 10), name, fill=INK)
        out.paste(im.crop((0, 0, width, min(height, im.height))), (x, LABEL_H))
    out.save(HERE / f'compare-{route}-{width}.png')


for r in ROUTES:
    board(r, 1440, 1400)
    board(r, 390, 1800)
