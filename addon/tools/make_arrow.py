"""Render the green upgrade arrow the tooltips, quest rewards and roll frames inline.

Run once, from anywhere, with the pipeline's own interpreter:

    /Users/jh/code/forever/data/.venv/bin/python addon/tools/make_arrow.py

It writes addon/ForeverSixty/media/arrow.tga, a 32x32 uncompressed 32-bit TGA (power of two,
no RLE, what the client reads). The file is committed; this script exists so the glyph can be
regenerated. It lives in addon/tools/, outside the packaged folder.

The glyph is the paper-doll and bag mark's own (SlotMarks.ARROW_PARTS: an up-pointing head of
six rows over a shaft, in a 12 x 12 box, Theme.HEX.success with a 1 px black outline), drawn
at twice the size so it stays sharp at the tooltip's 12 px. The client's bags-greenarrow atlas
is not used: on the live 1.60 client it drew a ~6 px dot (L4.1 fix round 6), and an inline
|A|a escape would draw the same dot inside a tooltip line.
"""

from __future__ import annotations

import sys
from pathlib import Path

from PIL import Image

SIZE = 32
SCALE = 2
# SlotMarks.ARROW_PARTS: { x, y, width, height } in the 12 x 12 box.
PARTS = [(5, 0, 2, 1), (4, 1, 4, 1), (3, 2, 6, 1), (2, 3, 8, 1), (1, 4, 10, 1), (0, 5, 12, 1), (4, 6, 4, 6)]
FILL = (0x6F, 0xCF, 0x8E, 255)
OUTLINE = (0, 0, 0, 255)
# The 12 px box, doubled, centred in the 32 px canvas.
OFFSET = (SIZE - 12 * SCALE) // 2


def draw() -> Image.Image:
    image = Image.new("RGBA", (SIZE, SIZE), (0, 0, 0, 0))
    pixels = image.load()
    for colour, grow in ((OUTLINE, SCALE), (FILL, 0)):
        for x, y, width, height in PARTS:
            left, top = OFFSET + x * SCALE - grow, OFFSET + y * SCALE - grow
            right, bottom = OFFSET + (x + width) * SCALE + grow, OFFSET + (y + height) * SCALE + grow
            for px in range(max(left, 0), min(right, SIZE)):
                for py in range(max(top, 0), min(bottom, SIZE)):
                    pixels[px, py] = colour
    return image


def main() -> int:
    out = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media" / "arrow.tga"
    draw().save(out, format="TGA", rle=False)
    print(out)
    return 0


if __name__ == "__main__":
    sys.exit(main())
