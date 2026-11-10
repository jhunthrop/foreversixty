"""Render the site's mark to the addon's minimap icon.

Run once, from the repository root, with a Python that has Pillow and cairosvg:

    python3 addon/tools/make_minimap_icon.py

It writes addon/ForeverSixty/media/minimap.tga, a 32x32 uncompressed
32-bit TGA -- power of two, no RLE, which is what the client reads. The
file is committed; this script exists so the mark can be regenerated,
not because the build runs it. It lives in addon/tools/, one level above
the packaged folder, so it can never end up in the zip.

The minimap button draws the mark at 20 px (Theme.SIZES.minimapIcon), where the painted seal turns
to mud, so the button takes the site's flat seal (owner 2026-11-10, tools/art/flat_seal.py): the
tiny bold seal, the one for 24 px and below, rendered at the texture's 32 px. The 64x64 list icon
(addon/ForeverSixty/icon.png, the TOC's IconTexture) stays the painted LX seal v4
(design/assets/logo/lx-seal-v4.png), as every mark at 64 px and up does. Rendering the flat seal
needs cairosvg (see tools/art/flat_seal.py).
"""

from __future__ import annotations

import subprocess
import sys
from pathlib import Path

from PIL import Image

sys.path.insert(0, str(Path(__file__).resolve().parents[2] / "tools" / "art"))
from flat_seal import flat_seal  # noqa: E402  (the flat seal at 44 px and below: one source for site and addon)

SIZE = 32
# The size the button draws the icon at (Theme.SIZES.minimapIcon), which picks the flat seal.
SHOWN_SIZE = 20
SOURCE = Path(__file__).resolve().parents[2] / "design" / "assets" / "logo" / "lx-seal-v4.png"
LIST_ICON = Path(__file__).resolve().parents[1] / "ForeverSixty" / "icon.png"
LIST_ICON_SIZE = 64
OUTPUT = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media" / "minimap.tga"


def render_painted(size: int) -> Image.Image:
    """The painted seal, downsampled to a square of `size` px.

    The source is already a circular disc on a transparent background (no
    square card to crop or mask), so resizing is the whole job.
    """
    source = Image.open(SOURCE).convert("RGBA")
    return source.resize((size, size), Image.LANCZOS)


def compress_png(path: Path) -> None:
    """pngquant (quality 90 to 100, kept only if smaller) then lossless oxipng; both on PATH (brew install)."""
    subprocess.run(["pngquant", "--quality=90-100", "--skip-if-larger", "--force", "--strip",
                    "--output", str(path), str(path)], check=False)
    subprocess.run(["oxipng", "-o", "4", "--strip", "safe", "-q", str(path)], check=True)


def main() -> int:
    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    # compression=None is the uncompressed path: Pillow only writes RLE
    # when it is asked for "tga_rle", and the client will not read RLE.
    flat_seal(SIZE, shown_px=SHOWN_SIZE).save(OUTPUT, format="TGA", compression=None)
    render_painted(LIST_ICON_SIZE).save(LIST_ICON, format="PNG")
    compress_png(LIST_ICON)
    for written in (OUTPUT, LIST_ICON):
        print(f"wrote {written} ({written.stat().st_size} bytes)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
