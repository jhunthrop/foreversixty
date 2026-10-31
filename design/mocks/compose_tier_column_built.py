"""Side-by-side of the tier column mock and the built page (spec 2026-10-31-tier-column.md).

  python3 design/mocks/compose_tier_column_built.py  -> design/reviews/2026-10-31-tier-column/sbs-<width>.png

Left: the mock board (design/mocks/2026-10-31-tier-column/after-<width>.png). Right: the built page's
full-page capture (shoot_tier_column_built.mjs) scaled to the mock's width, so both are the same scale.
"""
from pathlib import Path
from PIL import Image

ROOT = Path(__file__).resolve().parent.parent
MOCKS = ROOT / 'mocks' / '2026-10-31-tier-column'
REVIEWS = ROOT / 'reviews' / '2026-10-31-tier-column'
GAP = 24
BG = (7, 9, 13)


def compose(width: int) -> None:
    mock = Image.open(MOCKS / f'after-{width}.png').convert('RGB')
    built = Image.open(REVIEWS / f'built-{width}.png').convert('RGB')
    built = built.resize((mock.width, round(built.height * mock.width / built.width)))
    sheet = Image.new('RGB', (mock.width * 2 + GAP, max(mock.height, built.height)), BG)
    sheet.paste(mock, (0, 0))
    sheet.paste(built, (mock.width + GAP, 0))
    sheet.save(REVIEWS / f'sbs-{width}.png')


for w in (360, 1440, 2000):
    compose(w)
