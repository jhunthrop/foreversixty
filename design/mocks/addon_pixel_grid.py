"""Before and after of the tour's page 4 as the client rasterizes it (2026-11-13, addon 2.0.45).

    python3 design/mocks/addon_pixel_grid.py [units-per-pixel ...]

Writes design/reviews/2026-11-13-addon-dialogs/pixel-grid-<scale>.png: the window's left edge, title
band (medallion, rule) and footer at the given physical pixels per UI unit (1.875 = 1440p at UI scale 1,
the default; 1.40625 = 1080p), the window's origin a fraction of a pixel off the grid as a centred
window lands. Left, 2.0.44: textures unsnapped (the client blends each edge across the pixels it
straddles), 1-unit outlines, the 64 px crest stretched, Skip tour laid out before its face loaded
(blank). Right, the fix: every texture snaps (Theme.snapTexture), outlines whole pixels
(Theme.wholePixels), the 128 px crest, the faces asked for at load (Theme.preloadFaces). Text sits on
floored pixels in both, as the client draws it. Colours from Theme.HEX, labels from Locale.lua.
"""
import sys
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

ROOT = Path(__file__).resolve().parents[2]
ADDON = ROOT / 'addon/ForeverSixty'
OUT = ROOT / 'design/reviews/2026-11-13-addon-dialogs'
HEX = {'umber': '170e09', 'edge': '3d342d', 'blue': '114f72', 'blueEdge': '1d6a96', 'teal': '0f6064',
       'tealEdge': '408076', 'ink': 'e3dfd9', 'bone': 'f3eee2', 'bg': '3a2a22'}
SS = 8              # supersampling for coverage
ORIGIN = (0.4, 0.6)  # the window's top-left, in pixels past the grid
W_UNITS, H_UNITS = 360, 190  # the part of the window drawn: left edge, band, footer
BAND, MEDALLION, INSET, BUTTON_H = 64, 44, 6, 22


def rgb(key):
    h = HEX[key]
    return tuple(int(h[i:i + 2], 16) for i in (0, 2, 4))


def label(key):
    text = (ADDON / 'Locale.lua').read_text()
    return text.split(f'L.{key} = "', 1)[1].split('"', 1)[0]


class Raster:
    """Draws rects in UI units at `ppu` pixels a unit: blended by coverage, or snapped to whole pixels."""

    def __init__(self, ppu, snapped):
        self.ppu, self.snapped = ppu, snapped
        self.w, self.h = int(W_UNITS * ppu) + 4, int(H_UNITS * ppu) + 4
        self.hi = Image.new('RGB', (self.w * SS, self.h * SS), rgb('bg'))
        self.draw = ImageDraw.Draw(self.hi)

    def px(self, x, y):
        return ORIGIN[0] + x * self.ppu, ORIGIN[1] + y * self.ppu

    def thickness(self, units):
        if not self.snapped:
            return units
        return max(1, round(units * self.ppu)) / self.ppu

    def rect(self, x, y, w, h, key):
        x0, y0 = self.px(x, y)
        x1, y1 = self.px(x + w, y + h)
        if self.snapped:
            x0, y0, x1, y1 = round(x0), round(y0), round(x1), round(y1)
        self.draw.rectangle([x0 * SS, y0 * SS, x1 * SS - 1, y1 * SS - 1], fill=rgb(key))

    def outline(self, x, y, w, h, key):
        t = self.thickness(1)
        self.rect(x, y, w, t, key)
        self.rect(x, y + h - t, w, t, key)
        self.rect(x, y, t, h, key)
        self.rect(x + w - t, y, t, h, key)

    def finish(self):
        return self.hi.resize((self.w, self.h), Image.BOX)


def crest(img, raster, x, y, size, source):
    art = Image.open(ADDON / 'media/crests' / 'warrior.tga').convert('RGBA') if source is None else source
    x0, y0 = raster.px(x, y)
    if raster.snapped:
        x0, y0 = round(x0), round(y0)
    px = round(size * raster.ppu)
    drawn = art.resize((px, px), Image.BILINEAR)
    img.paste(drawn, (int(x0), int(y0)), drawn)


def text(img, raster, x, y, value, key, size=12, face='OpenSans-SemiBold.ttf', centre_w=None):
    font = ImageFont.truetype(str(ADDON / 'media/fonts' / face), round(size * raster.ppu))
    d = ImageDraw.Draw(img)
    x0, y0 = raster.px(x, y)
    if centre_w is not None:
        width = d.textlength(value, font=font)
        x0 += (centre_w * raster.ppu - width) / 2
    d.text((int(x0), int(y0)), value, font=font, fill=rgb(key))


def window(ppu, fixed):
    r = Raster(ppu, fixed)
    r.rect(0, 0, W_UNITS + 20, H_UNITS + 20, 'umber')
    r.outline(0, 0, W_UNITS + 20, H_UNITS + 20, 'edge')
    r.rect(0, BAND, W_UNITS + 20, r.thickness(1), 'edge')
    footer_y = H_UNITS - 14 - BUTTON_H
    buttons = [(16, 79, 'tourSkip', False), (16 + 79 + 120, 64, 'tourBackButton', False),
               (16 + 79 + 120 + 72, 64, 'tourNext', True)]
    for bx, bw, _, primary in buttons:
        r.rect(bx, footer_y, bw, BUTTON_H, 'teal' if primary else 'blue')
        r.outline(bx, footer_y, bw, BUTTON_H, 'tealEdge' if primary else 'blueEdge')
    img = r.finish()
    old = Image.open(ADDON / 'media/crests/warrior.tga').convert('RGBA').resize((64, 64), Image.LANCZOS)
    inset = MEDALLION - 2 * INSET
    crest(img, r, 12 + INSET, (BAND - MEDALLION) / 2 + INSET, inset, None if fixed else old)
    text(img, r, 12 + MEDALLION + 12, 16, label('tourGuildHeading'), 'bone', 19, 'Marcellus-Regular.ttf')
    text(img, r, 26, BAND + 16, label('tourGuildLead').split('. ')[0] + '.', 'ink', 12, 'OpenSans-Regular.ttf')
    for bx, bw, key, _ in buttons:
        if key == 'tourSkip' and not fixed:
            continue
        text(img, r, bx, footer_y + 3, label(key), 'ink', 12, centre_w=bw)
    return img


def board(ppu):
    before, after = window(ppu, False), window(ppu, True)
    gap = 24
    out = Image.new('RGB', (before.width * 2 + gap, before.height), rgb('bg'))
    out.paste(before, (0, 0))
    out.paste(after, (before.width + gap, 0))
    return out


def main(scales):
    OUT.mkdir(parents=True, exist_ok=True)
    for ppu in scales:
        path = OUT / f'pixel-grid-{ppu}.png'
        img = board(ppu)
        img.save(path)
        crop = img.crop((0, img.height - 90, img.width, img.height)).resize((img.width * 2, 180), Image.NEAREST)
        crop.save(OUT / f'pixel-grid-{ppu}-footer-2x.png')
        print(path)


if __name__ == '__main__':
    main([float(s) for s in sys.argv[1:]] or [1.875, 1.40625])
