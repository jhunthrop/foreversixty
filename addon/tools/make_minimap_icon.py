"""Render the site's mark to the addon's minimap icon.

Run once, from anywhere, with the pipeline's own interpreter:

    /Users/jh/code/forever/data/.venv/bin/python addon/tools/make_minimap_icon.py

It writes addon/ForeverSixty/media/minimap.tga, a 32x32 uncompressed
32-bit TGA -- power of two, no RLE, which is what the client reads. The
file is committed; this script exists so the mark can be regenerated,
not because the build runs it. It lives in addon/tools/, one level above
the packaged folder, so it can never end up in the zip.

The mark is the site's own: the LX seal (design/logo/foreversixty-mark.svg,
design/DESIGN-SYSTEM.md's "Logo" section -- a dark disc, a gold-gradient
ring with a dotted inner track, and LX in Cinzel gold), exported at
64x64 (design/logo/export/foreversixty-icon-64.png) and downsampled here.
It replaces an earlier version of this script that hand-drew a gold "60"
in Georgia -- that glyph predates the seal and no longer matches any
other mark on the site or in the addon (favicon, addon icon, header,
footer all switched to the seal in the site-logo lane).
"""

from __future__ import annotations

import sys
from pathlib import Path

from PIL import Image

SIZE = 32
SOURCE = Path(__file__).resolve().parents[2] / "design" / "logo" / "export" / "foreversixty-icon-64.png"
OUTPUT = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media" / "minimap.tga"


def render() -> Image.Image:
    """The seal, downsampled to the minimap button's size.

    The export is already a circular disc on a transparent background (no
    square card to crop or mask), so resizing is the whole job.
    """
    source = Image.open(SOURCE).convert("RGBA")
    return source.resize((SIZE, SIZE), Image.LANCZOS)


def main() -> int:
    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    # compression=None is the uncompressed path: Pillow only writes RLE
    # when it is asked for "tga_rle", and the client will not read RLE.
    render().save(OUTPUT, format="TGA", compression=None)
    print(f"wrote {OUTPUT} ({OUTPUT.stat().st_size} bytes)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
