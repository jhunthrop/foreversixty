"""Render the site's mark to the addon's minimap icon.

Run once, from anywhere, with the pipeline's own interpreter:

    /Users/jh/code/forever/data/.venv/bin/python addon/tools/make_minimap_icon.py

It writes addon/ForeverSixty/media/minimap.tga, a 32x32 uncompressed
32-bit TGA -- power of two, no RLE, which is what the client reads. The
file is committed; this script exists so the mark can be regenerated,
not because the build runs it. It lives in addon/tools/, one level above
the packaged folder, so it can never end up in the zip.

The mark is the site's own: the LX seal v3 (design/assets/logo/lx-seal-v3.png, the owner's
painted seal cut out with real transparency, 2026-11-09). The same source writes the addon's
64x64 list icon (addon/ForeverSixty/icon.png, the TOC's IconTexture), so every mark matches.
"""

from __future__ import annotations

import sys
from pathlib import Path

from PIL import Image

SIZE = 32
SOURCE = Path(__file__).resolve().parents[2] / "design" / "assets" / "logo" / "lx-seal-v3.png"
LIST_ICON = Path(__file__).resolve().parents[1] / "ForeverSixty" / "icon.png"
LIST_ICON_SIZE = 64
OUTPUT = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media" / "minimap.tga"


def render(size: int) -> Image.Image:
    """The seal, downsampled to a square of `size` px.

    The source is already a circular disc on a transparent background (no
    square card to crop or mask), so resizing is the whole job.
    """
    source = Image.open(SOURCE).convert("RGBA")
    return source.resize((size, size), Image.LANCZOS)


def main() -> int:
    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    # compression=None is the uncompressed path: Pillow only writes RLE
    # when it is asked for "tga_rle", and the client will not read RLE.
    render(SIZE).save(OUTPUT, format="TGA", compression=None)
    render(LIST_ICON_SIZE).save(LIST_ICON, format="PNG")
    for written in (OUTPUT, LIST_ICON):
        print(f"wrote {written} ({written.stat().st_size} bytes)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
