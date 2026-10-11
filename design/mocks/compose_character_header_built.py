"""Composite of the built shared character header (shoot_character_header_built.mjs) against its mock.

  python3 design/mocks/compose_character_header_built.py

Writes to renders/ (git-ignored):
  character-header-built-vs-mock-states.png   mock left, built right, the states sheet at the same scale

The mock render comes from gen_character_header.py; set CH_MOCK_DIR to read it from another checkout's
renders directory (they are git-ignored, so a fresh worktree has none until the generator runs).
"""
import os
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

HERE = Path(__file__).resolve().parent
RENDERS = HERE / 'renders'
MOCKS = Path(os.environ.get('CH_MOCK_DIR', RENDERS))
BG = (7, 9, 13)
LABEL = (229, 185, 85)
FONT = ImageFont.load_default()
CAPTION_H = 26
GAP = 24

PAIRS = [   # (mock, built, output)
    ('character-header-states', 'character-header-built-states-1440', 'character-header-built-vs-mock-states'),
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
