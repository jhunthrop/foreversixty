"""Side-by-side sheets for the BA1 brand gate: each addon board beside the web page it must match.

    python3 design/reviews/2026-11-10-ba1/side_by_side.py

Reads captures/*.png (design/mocks/addon_ux.py with MOCK_OUT=.../captures) and the web renders in
design/reviews/2026-11-10-unify-format/after/, writes side-by-side-*.png here.
"""
from pathlib import Path

from PIL import Image

HERE = Path(__file__).resolve().parent
WEB = HERE.parents[0] / '2026-11-10-unify-format/after'
PAIRS = {
    'panel': ('captures/panel-raider-60.png', 'character-1440.jpg'),
    'onboarding': ('captures/onboarding-5-set-up.png', 'home-1440.jpg'),
    'tooltip': ('captures/item-tooltip-after.png', 'bis-1440.jpg'),
}
HEIGHT = 1100


def scaled(path: Path) -> Image.Image:
    im = Image.open(path).convert('RGB')
    return im.resize((round(im.width * HEIGHT / im.height), HEIGHT))


def main() -> None:
    for name, (addon, web) in PAIRS.items():
        left, right = scaled(HERE / addon), scaled(WEB / web)
        sheet = Image.new('RGB', (left.width + right.width + 24, HEIGHT), (0, 0, 0))
        sheet.paste(left, (0, 0))
        sheet.paste(right, (left.width + 24, 0))
        sheet.save(HERE / f'side-by-side-{name}.png')
        print(HERE / f'side-by-side-{name}.png')


if __name__ == '__main__':
    main()
