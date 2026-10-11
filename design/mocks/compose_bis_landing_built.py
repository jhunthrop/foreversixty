"""Composite of the built /bis landing (shoot_bis_landing_built.mjs) against its mock, for review.

  python3 design/mocks/compose_bis_landing_built.py

Writes to renders/ (git-ignored):
  bis-landing-built-vs-mock-1440.png   mock left, built right, same scale, signed in at 1440
  bis-landing-built-vs-mock-signed-out-1440.png   the same, signed out

The mock renders come from gen_bis_landing.py; set BIS_MOCK_DIR to read them from another
checkout's renders directory (they are git-ignored, so a fresh worktree has none until
gen_bis_landing.py runs).
"""
import os
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

HERE = Path(__file__).resolve().parent
RENDERS = HERE / 'renders'
MOCKS = Path(os.environ.get('BIS_MOCK_DIR', RENDERS))
BG = (7, 9, 13)
LABEL = (229, 185, 85)
FONT = ImageFont.load_default()
CAPTION_H = 26
GAP = 24

PAIRS = [   # (mock, built, output)
    ('bis-landing', 'bis-landing-built-1440', 'bis-landing-built-vs-mock-1440'),
    ('bis-landing-signed-out', 'bis-landing-built-signed-out-1440', 'bis-landing-built-vs-mock-signed-out-1440'),
]


def captioned(image: Image.Image, text: str) -> Image.Image:
    out = Image.new('RGB', (image.width, image.height + CAPTION_H), BG)
    ImageDraw.Draw(out).text((8, 7), text, fill=LABEL, font=FONT)
    out.paste(image, (0, CAPTION_H))
    return out


def compose(mock: Path, built: Path, out: Path) -> None:
    left = captioned(Image.open(mock).convert('RGB'), 'MOCK (design/mocks/renders/' + mock.name + ')')
    right = captioned(Image.open(built).convert('RGB'), 'BUILT (' + built.name + ')')
    sheet = Image.new('RGB', (left.width + GAP + right.width, max(left.height, right.height)), BG)
    sheet.paste(left, (0, 0))
    sheet.paste(right, (left.width + GAP, 0))
    sheet.save(out)
    print(out)


if __name__ == '__main__':
    for mock, built, name in PAIRS:
        compose(MOCKS / f'{mock}.png', RENDERS / f'{built}.png', RENDERS / f'{name}.png')
