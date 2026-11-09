"""Render the bronze-ring class medallion the addon's own frames carry in their title band.

Run once, from anywhere, with the pipeline's own interpreter:

    /Users/jh/code/forever/data/.venv/bin/python addon/tools/make_medallion.py

It writes two 128 x 128 uncompressed 32-bit TGAs (power of two, no RLE, what the client reads)
to addon/ForeverSixty/media/:

- medallion-ring.tga: the site's medallion ring (design/mocks/addon_ux.py medallion()): a dark
  bronze band (#3b2a15) with a brass rim (#c9a46a) on each side, transparent inside and out, so
  the class crest drawn under it shows through the middle.
- medallion-notch.tga: the faction notch at one o'clock, a white right triangle with a dark
  outline (#0b0705). It is tinted per faction at runtime with SetVertexColor (blue Alliance,
  red Horde; Theme.HEX.notchAlliance and notchHorde): the white fill takes the tint, the dark
  outline stays dark.

Both share the mock's 96-unit viewBox scaled to the canvas, so the frame lays them over one
another with SetAllPoints. The geometry is drawn at four times the size and scaled down, so the
edges are smooth at the band's 44 px. The files are committed; this script exists so they can be
regenerated. It lives in addon/tools/, outside the packaged folder.
"""

from __future__ import annotations

import sys
from pathlib import Path

from PIL import Image, ImageDraw

SIZE = 128
SUPER = 4
VIEWBOX = 96
BRONZE = (0x3B, 0x2A, 0x15, 255)
BRASS = (0xC9, 0xA4, 0x6A, 255)
OUTLINE = (0x0B, 0x07, 0x05, 255)
WHITE = (255, 255, 255, 255)
CENTRE = 48
# medallion(): the band (r 40, stroke 10), the outer rim (r 44.5, stroke 1.6), the inner rim
# (r 35.5, stroke 1.2), in viewBox units.
BAND = (40, 10)
OUTER_RIM = (44.5, 1.6)
INNER_RIM = (35.5, 1.2)
# medallion()'s notch: M62 4 l22 0 l0 22 z, stroked 2.
NOTCH = ((62, 4), (84, 4), (84, 26))
NOTCH_STROKE = 2


def scale(value: float) -> float:
    return value * SIZE * SUPER / VIEWBOX


def ring(draw: ImageDraw.ImageDraw, radius: float, stroke: float, colour: tuple[int, int, int, int]) -> None:
    centre = scale(CENTRE)
    outer, inner = scale(radius + stroke / 2), scale(radius - stroke / 2)
    draw.ellipse((centre - outer, centre - outer, centre + outer, centre + outer), fill=colour)
    draw.ellipse((centre - inner, centre - inner, centre + inner, centre + inner), fill=(0, 0, 0, 0))


def layer() -> tuple[Image.Image, ImageDraw.ImageDraw]:
    image = Image.new("RGBA", (SIZE * SUPER, SIZE * SUPER), (0, 0, 0, 0))
    return image, ImageDraw.Draw(image)


def finish(image: Image.Image) -> Image.Image:
    return image.resize((SIZE, SIZE), Image.LANCZOS)


def draw_ring() -> Image.Image:
    # Each circle is its own layer composited in the mock's order, so a rim's transparent
    # middle never punches through the band under it.
    result = Image.new("RGBA", (SIZE * SUPER, SIZE * SUPER), (0, 0, 0, 0))
    for radius, stroke, colour in ((*BAND, BRONZE), (*OUTER_RIM, BRASS), (*INNER_RIM, BRASS)):
        image, draw = layer()
        ring(draw, radius, stroke, colour)
        result = Image.alpha_composite(result, image)
    return finish(result)


def draw_notch() -> Image.Image:
    image, draw = layer()
    points = [(scale(x), scale(y)) for x, y in NOTCH]
    draw.polygon(points, fill=WHITE)
    draw.line(points + [points[0]], fill=OUTLINE, width=round(scale(NOTCH_STROKE)), joint="curve")
    return finish(image)


def main() -> int:
    media = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media"
    for name, image in (("medallion-ring.tga", draw_ring()), ("medallion-notch.tga", draw_notch())):
        out = media / name
        image.save(out, format="TGA", rle=False)
        print(out)
    return 0


if __name__ == "__main__":
    sys.exit(main())
