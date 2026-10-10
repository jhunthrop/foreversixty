"""Same-scale side-by-side of the mock and the built signed-in home at 1440 (spec 2026-10-10 section 9).

    python3 design/mocks/compose_home_built.py   -> design/mocks/renders/home-panel-built-vs-mock-1440.png

Left the mock (`gen_home_panel.py home-panel`), right the built page (`shoot_home_built.mjs 1440`), both 1440 px wide
and cropped to the same height (the first screen and the table), so every offset can be compared by eye.
"""
from pathlib import Path

from PIL import Image, ImageDraw

RENDERS = Path(__file__).parent / 'renders'
CROP_HEIGHT = 1100
LABEL_HEIGHT = 36
GUTTER = 24
BACKGROUND = (7, 9, 13)
LABEL_COLOUR = (229, 185, 85)


def load(name: str) -> Image.Image:
    image = Image.open(RENDERS / name).convert('RGB')
    return image.crop((0, 0, image.width, min(image.height, CROP_HEIGHT)))


def main() -> None:
    mock, built = load('home-panel.png'), load('home-panel-built-1440.png')
    height = max(mock.height, built.height) + LABEL_HEIGHT
    sheet = Image.new('RGB', (mock.width + built.width + GUTTER, height), BACKGROUND)
    draw = ImageDraw.Draw(sheet)
    draw.text((8, 10), 'MOCK  home-panel (1440)', fill=LABEL_COLOUR)
    draw.text((mock.width + GUTTER + 8, 10), 'BUILT  home-panel-built-1440', fill=LABEL_COLOUR)
    sheet.paste(mock, (0, LABEL_HEIGHT))
    sheet.paste(built, (mock.width + GUTTER, LABEL_HEIGHT))
    out = RENDERS / 'home-panel-built-vs-mock-1440.png'
    sheet.save(out)
    print(out)


if __name__ == '__main__':
    main()
