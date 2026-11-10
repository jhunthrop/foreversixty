"""The seal v3 against v4 at 512/128/64/32/16, each at 1:1, on the nav teal, the footer brown and parchment.

    python3 design/mocks/seal-v4/review_sheet.py

Writes design/reviews/2026-11-09-seal-v4/sheet.png. Every size is the sized render itself, never rescaled.
"""
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

REPO = Path(__file__).resolve().parents[3]
LOGO = REPO / 'design/assets/logo'
FONT = REPO / 'api/internal/brand/fonts/OpenSans-SemiBold.ttf'
OUT = REPO / 'design/reviews/2026-11-09-seal-v4/sheet.png'
SIZES = (512, 128, 64, 32, 16)
GROUNDS = (('nav #0b2f3b', (11, 47, 59)), ('footer #170e09', (23, 14, 9)), ('light #f3eee2', (243, 238, 226)))
VERSIONS = ('v3', 'v4')
GAP, LABEL_W, HEAD_H = 24, 170, 40
LIGHT_GROUND = 400  # channel sum above which a ground takes dark ink


def main() -> None:
    font = ImageFont.truetype(str(FONT), 18)
    row_h = SIZES[0] + 2 * GAP
    width = LABEL_W + sum(SIZES) + GAP * (len(SIZES) + 1)
    sheet = Image.new('RGB', (width, HEAD_H + row_h * len(GROUNDS) * len(VERSIONS)), (40, 40, 40))
    draw = ImageDraw.Draw(sheet)
    x = LABEL_W + GAP
    for size in SIZES:
        draw.text((x, 10), f'{size} px', fill=(230, 230, 230), font=font)
        x += size + GAP
    y = HEAD_H
    for ground_name, ground in GROUNDS:
        for version in VERSIONS:
            draw.rectangle((0, y, width, y + row_h - 1), fill=ground)
            ink = (20, 20, 20) if sum(ground) > LIGHT_GROUND else (235, 230, 220)
            draw.text((12, y + row_h // 2 - 22), f'seal {version}', fill=ink, font=font)
            draw.text((12, y + row_h // 2 + 2), ground_name, fill=ink, font=font)
            x = LABEL_W + GAP
            for size in SIZES:
                seal = Image.open(LOGO / f'lx-seal-{version}-{size}.png').convert('RGBA')
                sheet.paste(seal, (x, y + (row_h - size) // 2), seal)
                x += size + GAP
            y += row_h
    sheet.save(OUT, optimize=True)
    print(f'wrote {OUT} {sheet.size}')


if __name__ == '__main__':
    main()
