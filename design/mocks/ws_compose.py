"""Before/after composites: the live page and the mock at the same scale, top aligned, both full height.
   python3 ws_compose.py bis -> renders/web/compare-bis-<w>.png"""
import sys
from pathlib import Path
from PIL import Image, ImageDraw

ROOT = Path(__file__).resolve().parents[2]
BEFORE = ROOT / 'design/reviews/2026-11-05-wide-layout/before'
AFTER = Path(__file__).resolve().parent / 'renders/web'


def compose(page: str, before_name: str, after_name: str, w: int, scale: float, max_h: int = 2400, before_label: str = 'live') -> Path:
    a = Image.open(BEFORE / f'{before_name}-{w}.png').convert('RGB')
    b = Image.open(AFTER / f'{after_name}-{w}.png').convert('RGB')
    s = lambda im: im.resize((int(im.width * scale), int(im.height * scale)), Image.LANCZOS)
    a, b = s(a), s(b)
    h = min(max(a.height, b.height), max_h)
    pad, head = 12, 34
    sheet = Image.new('RGB', (a.width + b.width + pad * 3, h + head + pad), (28, 28, 30))
    d = ImageDraw.Draw(sheet)
    d.text((pad, 10), f'BEFORE ({before_label}, {w})', fill=(229, 185, 85))
    d.text((a.width + pad * 2, 10), f'AFTER (mock, {w})', fill=(127, 212, 138))
    sheet.paste(a.crop((0, 0, a.width, h)), (pad, head))
    sheet.paste(b.crop((0, 0, b.width, h)), (a.width + pad * 2, head))
    out = AFTER / f'compare-{page}-{w}.png'
    sheet.save(out)
    return out


if __name__ == '__main__':
    page = sys.argv[1]
    for w, sc in ((1440, .6), (1920, .5), (2000, .5), (390, .9)):
        print(compose(page, page, f'web-{page}', w, sc))
