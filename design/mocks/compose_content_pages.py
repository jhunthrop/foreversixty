"""Side-by-side sheets for lane P15: the round-5 mock (left) and the build (right), same scale, per page and width.

    python3 design/mocks/compose_content_pages.py design/reviews/2026-11-09-f-content
reads mock-<page>-<width>.png and build-<page>-<width>.png there (shoot_content_pages.mjs) and writes side-by-side-<page>-<width>.png.
"""
import sys
from pathlib import Path

from PIL import Image, ImageDraw

GAP = 24
LABEL_H = 28


def compose(left: Path, right: Path, out: Path) -> None:
    a, b = Image.open(left).convert('RGB'), Image.open(right).convert('RGB')
    w, h = a.width + GAP + b.width, max(a.height, b.height) + LABEL_H
    sheet = Image.new('RGB', (w, h), (24, 24, 24))
    sheet.paste(a, (0, LABEL_H))
    sheet.paste(b, (a.width + GAP, LABEL_H))
    draw = ImageDraw.Draw(sheet)
    draw.text((8, 8), 'mock (directions b4)', fill=(240, 240, 240))
    draw.text((a.width + GAP + 8, 8), 'build', fill=(240, 240, 240))
    sheet.save(out, optimize=True)


def main() -> None:
    folder = Path(sys.argv[1])
    for build in sorted(folder.glob('build-*.png')):
        mock = folder / build.name.replace('build-', 'mock-', 1)
        if mock.exists():
            compose(mock, build, folder / build.name.replace('build-', 'side-by-side-', 1))


if __name__ == '__main__':
    main()
