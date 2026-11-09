"""Round-4 study sheets, both compared at the same zoom against crops of Blizzard's Forever page captures.

  python3 design/mocks/directions/study_sheets.py
  -> blizzard/fonts.html (rendered by shoot.cjs-style capture below) and blizzard/texture-compare.png

The sheets hold crops of their page, so they live in blizzard/ (git-ignored) and travel only to the local gallery.
Fonts: their display face is BlizQuadrata Web (Blizzard's cut of ITC Friz Quadrata, not licensable); the candidates are
open-licence Google Fonts with a Quadrata-like glyphic or flared build.
"""
import base64
from io import BytesIO
from pathlib import Path

from PIL import Image

HERE = Path(__file__).resolve().parent
BZ = HERE / 'blizzard'
OURS = HERE / 'png'
CANDIDATES = ['Marcellus', 'Philosopher', 'Rosarivo', 'Cinzel', 'Quattrocento', 'Forum', 'Bellefair', 'Cormorant Garamond']
HEADINGS = [('screen-1440-2.png', (360, 255, 1080, 335), 'Undiscovered Vistas Await', 60),
            ('screen-1440-1.png', (110, 190, 450, 255), 'Latest Updates', 48)]


def crop_uri(name: str, box: tuple[int, int, int, int]) -> str:
    buf = BytesIO()
    Image.open(BZ / name).convert('RGB').crop(box).save(buf, 'PNG')
    return 'data:image/png;base64,' + base64.b64encode(buf.getvalue()).decode()


def fonts_sheet() -> Path:
    families = '&'.join('family=' + f.replace(' ', '+') + (':wght@400;600' if f == 'Cormorant Garamond' else '') for f in CANDIDATES)
    rows = []
    for name, box, text, px in HEADINGS:
        rows.append(f'<h2>Theirs: BlizQuadrata Web, {px}px</h2><img src="{crop_uri(name, box)}" alt="">')
        rows += [f'<div class="c"><span class="lab">{f}</span><span class="t" style="font-family:\'{f}\';font-size:{px}px">{text}</span></div>' for f in CANDIDATES]
    html = (f'<!doctype html><meta charset="utf-8"><link rel="stylesheet" href="https://fonts.googleapis.com/css2?{families}&display=swap">'
            '<style>body{margin:0;padding:24px;background:#6b5a3d;color:#f3eee2;font:14px Open Sans,Arial}h2{font:600 16px Arial;margin:24px 0 8px}'
            '.c{display:flex;align-items:baseline;gap:18px;padding:6px 0;border-bottom:1px solid rgba(0,0,0,.2)}'
            '.lab{width:170px;font:600 13px Arial;color:#fcedcf}.t{color:#f3eee2;text-shadow:0 2px 3px rgba(0,0,0,.45);line-height:1.1}'
            'img{display:block;border:1px solid #000}</style>' + ''.join(rows))
    out = BZ / 'fonts.html'
    out.write_text(html)
    return out


def ground_sheet() -> Path:
    """Our brown ground alone (the page CSS over an empty body), to crop at their crop's size."""
    from style import CSS  # local: style reads the logo through git, only needed here
    out = BZ / 'ground.html'
    out.write_text(f'<!doctype html><meta charset="utf-8"><style>{CSS}body{{height:400px}}</style>')
    return out


def texture_compare() -> Path:
    """Their brown ground beside ours at the same size and zoom (ground.png from ground_sheet, captured by shoot_sheet.cjs)."""
    region = (0, 100, 700, 170)
    theirs = Image.open(BZ / 'screen-1440-1.png').convert('RGB').crop(region)
    ours = Image.open(BZ / 'ground.png').convert('RGB').crop(region)
    sheet = Image.new('RGB', (1420, 70), '#000')
    sheet.paste(theirs, (0, 0))
    sheet.paste(ours, (720, 0))
    out = BZ / 'texture-compare.png'
    sheet.save(out)
    return out


if __name__ == '__main__':
    import sys
    if sys.argv[1:] == ['compare']:
        print(texture_compare())
    else:
        print(fonts_sheet())
        print(ground_sheet())
