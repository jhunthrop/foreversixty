"""Composites of the built nav captures (shoot_nav_built.mjs) for review.

  python3 design/mocks/compose_nav_built.py

Writes to renders/ (git-ignored):
  nav-built-vs-mock-*.png   the mock above the built page at 1440, side by side at 390, same scale
  nav-built-states.png      the states sheet (hover, focus, ring marks, stale, failed, one, none,
                            loading, load error, paste), each cell captioned

The mock renders come from gen_nav.py; set NAV_MOCK_DIR to read them from another checkout's
renders directory (they are git-ignored, so a fresh worktree has none until gen_nav.py runs).
"""
import os
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

HERE = Path(__file__).resolve().parent
RENDERS = HERE / 'renders'
MOCKS = Path(os.environ.get('NAV_MOCK_DIR', RENDERS))
BG = (7, 9, 13)
LABEL = (229, 185, 85)
TEXT = (154, 148, 132)
FONT = ImageFont.load_default()
CAPTION_H = 26

PAIRS = [   # (mock, built, orientation)
    ('nav', 'full-bar-1440', 'stack'),
    ('nav-open', 'full-open-1440', 'stack'),
    ('nav-phone-menu', 'full-menu-390', 'side'),
    ('nav-phone-open', 'phone-390-sheet', 'side'),
]

STATES = [   # (file, caption)
    ('state-hover', 'Closed, hover'),
    ('state-focus', 'Closed, focus-visible'),
    ('state-open', 'Closed, open'),
    ('state-mark-stale', 'Ring mark: stale, tooltip'),
    ('state-mark-failed', 'Ring mark: failed, tooltip'),
    ('state-mark-session', 'Ring mark: session ended, tooltip'),
    ('state-discord-tip', 'Discord below 1440, tooltip'),
    ('state-row-hover', 'Row, hover'),
    ('state-row-focus', 'Row, focus-visible'),
    ('state-stale-row', 'Stale row'),
    ('state-failed-row', 'Failed row'),
    ('state-one', 'One character'),
    ('state-none', 'None'),
    ('state-loading', 'Loading'),
    ('state-load-error', 'List did not load'),
    ('state-paste', 'Paste an export, open'),
    ('state-paste-error', 'Paste, not an export'),
]


def load(name: str, directory: Path) -> Image.Image:
    return Image.open(directory / f'{name}.png').convert('RGB')


def caption(text: str, width: int) -> Image.Image:
    strip = Image.new('RGB', (width, CAPTION_H), BG)
    ImageDraw.Draw(strip).text((8, 8), text, fill=LABEL, font=FONT)
    return strip


def pair(mock_name: str, built_name: str, orientation: str) -> Image.Image:
    mock, built = load(mock_name, MOCKS), load(f'nav-built-{built_name}', RENDERS)
    if orientation == 'stack':
        width, height = max(mock.width, built.width), CAPTION_H * 2 + mock.height + built.height
        sheet = Image.new('RGB', (width, height), BG)
        sheet.paste(caption(f'MOCK  {mock_name}', width), (0, 0))
        sheet.paste(mock, (0, CAPTION_H))
        sheet.paste(caption(f'BUILT  {built_name}', width), (0, CAPTION_H + mock.height))
        sheet.paste(built, (0, CAPTION_H * 2 + mock.height))
        return sheet
    gap = 24
    sheet = Image.new('RGB', (mock.width + built.width + gap, CAPTION_H + max(mock.height, built.height)), BG)
    sheet.paste(caption(f'MOCK  {mock_name}', mock.width), (0, 0))
    sheet.paste(mock, (0, CAPTION_H))
    sheet.paste(caption(f'BUILT  {built_name}', built.width), (mock.width + gap, 0))
    sheet.paste(built, (mock.width + gap, CAPTION_H))
    return sheet


def states_sheet() -> Image.Image:
    cells = [(load(f'nav-built-{name}', RENDERS), text) for name, text in STATES]
    columns, gap = 3, 24
    cell_w = max(image.width for image, _ in cells)
    rows = [cells[index:index + columns] for index in range(0, len(cells), columns)]
    row_heights = [CAPTION_H + max(image.height for image, _ in row) for row in rows]
    sheet = Image.new('RGB', (columns * cell_w + (columns + 1) * gap, sum(row_heights) + (len(rows) + 1) * gap), BG)
    y = gap
    for row, row_h in zip(rows, row_heights):
        for index, (image, text) in enumerate(row):
            x = gap + index * (cell_w + gap)
            sheet.paste(caption(text, cell_w), (x, y))
            sheet.paste(image, (x, y + CAPTION_H))
        y += row_h + gap
    return sheet


if __name__ == '__main__':
    for mock_name, built_name, orientation in PAIRS:
        out = RENDERS / f'nav-built-vs-mock-{built_name}.png'
        pair(mock_name, built_name, orientation).save(out)
        print(out)
    out = RENDERS / 'nav-built-states.png'
    states_sheet().save(out)
    print(out)
