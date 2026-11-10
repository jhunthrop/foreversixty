"""The page-head contact sheet (design/reviews/2026-11-09-sim-head): every page's head crop at one
width, stacked with its page name, so a reviewer judges at a glance whether the heads read as one.

    python3 design/reviews/2026-11-09-sim-head/contact_sheet.py before 1440
    python3 design/reviews/2026-11-09-sim-head/contact_sheet.py after 390

Reads <phase>/heads/<id>-<width>.png (web/tests/captures/page-heads), writes
<phase>/contact-sheet-<width>.png. Phone sheets lay the heads out in columns.
"""
import sys
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

HERE = Path(__file__).resolve().parent
LABEL_H = 26
GAP = 10
BG = (16, 12, 10)
INK = (230, 220, 200)
PHONE_COLUMNS = 6


def heads(phase: str, width: int) -> list[tuple[str, Image.Image]]:
    folder = HERE / phase / 'heads'
    out = []
    for path in sorted(folder.glob(f'*-{width}.png')):
        out.append((path.stem.rsplit('-', 1)[0], Image.open(path).convert('RGB')))
    if not out:
        raise SystemExit(f'no heads at {width} in {folder}')
    return out


def tile(name: str, img: Image.Image, font: ImageFont.ImageFont) -> Image.Image:
    board = Image.new('RGB', (img.width, img.height + LABEL_H), BG)
    ImageDraw.Draw(board).text((8, 5), name, fill=INK, font=font)
    board.paste(img, (0, LABEL_H))
    return board


def stack(tiles: list[Image.Image], columns: int) -> Image.Image:
    rows = [tiles[i:i + columns] for i in range(0, len(tiles), columns)]
    col_w = max(t.width for t in tiles)
    heights = [max(t.height for t in row) for row in rows]
    sheet = Image.new('RGB', (columns * col_w + (columns - 1) * GAP, sum(heights) + GAP * (len(rows) - 1)), BG)
    y = 0
    for row, h in zip(rows, heights):
        for i, t in enumerate(row):
            sheet.paste(t, (i * (col_w + GAP), y))
        y += h + GAP
    return sheet


def main() -> None:
    phase, width = sys.argv[1], int(sys.argv[2])
    try:
        font = ImageFont.truetype('/System/Library/Fonts/Supplemental/Arial Bold.ttf', 15)
    except OSError:
        font = ImageFont.load_default()
    tiles = [tile(name, img, font) for name, img in heads(phase, width)]
    sheet = stack(tiles, PHONE_COLUMNS if width < 600 else 1)
    out = HERE / phase / f'contact-sheet-{width}.png'
    sheet.save(out, optimize=True)
    print(out, sheet.size)


if __name__ == '__main__':
    main()
