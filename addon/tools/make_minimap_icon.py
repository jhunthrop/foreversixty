"""Render the site's mark to the addon's minimap icon.

Run once, from anywhere, with the pipeline's own interpreter:

    /Users/jh/code/forever/data/.venv/bin/python addon/tools/make_minimap_icon.py

It writes addon/ForeverSixty/media/minimap.tga, a 32x32 uncompressed
32-bit TGA -- power of two, no RLE, which is what the client reads. The
file is committed; this script exists so the mark can be regenerated,
not because the build runs it. It lives in addon/tools/, one level above
the packaged folder, so it can never end up in the zip.

The mark is the site's own: web/public/favicon.svg is a gold "60" in
Georgia on #07090d, and this is that at 32 pixels.
"""

from __future__ import annotations

import sys
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

SIZE = 32
CORNER_RADIUS = 6
STONE = (7, 9, 13, 255)
GOLD = (229, 185, 85, 255)
TEXT = "60"
FONT_SIZE = 18
FONT_CANDIDATES = (
    "/System/Library/Fonts/Supplemental/Georgia Bold.ttf",
    "/System/Library/Fonts/Supplemental/Georgia.ttf",
    "/usr/share/fonts/truetype/dejavu/DejaVuSerif-Bold.ttf",
)
OUTPUT = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media" / "minimap.tga"


def load_font():
    """Georgia, to match the favicon; a serif fallback, then Pillow's own."""
    for candidate in FONT_CANDIDATES:
        if Path(candidate).exists():
            return ImageFont.truetype(candidate, FONT_SIZE)
    return ImageFont.load_default(size=FONT_SIZE)


def render() -> Image.Image:
    image = Image.new("RGBA", (SIZE, SIZE), (0, 0, 0, 0))
    draw = ImageDraw.Draw(image)
    draw.rounded_rectangle(
        (0, 0, SIZE - 1, SIZE - 1), radius=CORNER_RADIUS, fill=STONE, outline=GOLD
    )
    font = load_font()
    left, top, right, bottom = draw.textbbox((0, 0), TEXT, font=font)
    draw.text(
        ((SIZE - (right - left)) / 2 - left, (SIZE - (bottom - top)) / 2 - top),
        TEXT,
        font=font,
        fill=GOLD,
    )
    return image


def main() -> int:
    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    # compression=None is the uncompressed path: Pillow only writes RLE
    # when it is asked for "tga_rle", and the client will not read RLE.
    render().save(OUTPUT, format="TGA", compression=None)
    print(f"wrote {OUTPUT} ({OUTPUT.stat().st_size} bytes)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
