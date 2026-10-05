"""Render the site's circular class crests to the addon's own media, plus
the one ring texture the header tints per class at runtime.

Run once, from anywhere, with the pipeline's own interpreter (falls back
to the system `python3` used here, which also carries Pillow):

    /Users/jh/code/forever/data/.venv/bin/python addon/tools/make_class_crests.py

It writes addon/ForeverSixty/media/crests/<class>.tga (nine files, one
per class) and addon/ForeverSixty/media/crests/ring.tga, every one a
32-bit uncompressed TGA -- the same format and the same reasoning
make_minimap_icon.py already gives for the seal: power of two, no RLE,
what the 1.60 client reads.

Design §2 / §4.5.1 (round-2 owner ruling): the character header carries
the circular ringed class crest, never a square class icon. The site's
own crests (web/public/icons/hd/crests/<class>.webp) are already a
circular disc on a transparent background -- confirmed here by sampling
a corner pixel's alpha -- so there is no square card to crop and no
MaskTexture dependency to guard: resizing is the whole job for the crest
itself, the same shortcut make_minimap_icon.py already takes for the
seal. The 2px class-colour ring is a second, class-agnostic asset (a
plain white ring, transparent inside and outside it) that Window.lua
tints at runtime with Theme.classColor via SetVertexColor, so nine ring
files are never needed for nine classes.
"""

from __future__ import annotations

import sys
from pathlib import Path

from PIL import Image, ImageDraw

SIZE = 64
SOURCE_DIR = Path(__file__).resolve().parents[2] / "web" / "public" / "icons" / "hd" / "crests"
OUTPUT_DIR = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media" / "crests"

# The nine class slugs this addon already keys everything else by
# (Talents.playerClassSlug's own lower-cased UnitClass token).
CLASSES = [
    "warrior", "paladin", "hunter", "rogue", "priest",
    "shaman", "mage", "warlock", "druid",
]

# The ring's own geometry: an outer radius just inside the 64px canvas's
# own edge (so the ring never clips when drawn at the header's 36px
# on-screen size) and an inner radius RING_THICKNESS inside that.
RING_OUTER = 31
RING_THICKNESS = 4


def render_crest(slug: str) -> Image.Image:
    source = Image.open(SOURCE_DIR / f"{slug}.webp").convert("RGBA")
    corner_alpha = source.getpixel((2, 2))[3]
    if corner_alpha > 8:
        raise ValueError(f"{slug}.webp's corner is not transparent (alpha {corner_alpha}); "
            "it is not the circular disc this script assumes, and needs an explicit crop/mask step.")
    return source.resize((SIZE, SIZE), Image.LANCZOS)


def render_ring() -> Image.Image:
    """A plain white ring, tinted per class by Theme.classColor at runtime
    (SetVertexColor) -- one asset for every class, not nine."""
    canvas = Image.new("RGBA", (SIZE, SIZE), (0, 0, 0, 0))
    draw = ImageDraw.Draw(canvas)
    center = SIZE / 2
    draw.ellipse(
        (center - RING_OUTER, center - RING_OUTER, center + RING_OUTER, center + RING_OUTER),
        outline=(255, 255, 255, 255), width=RING_THICKNESS,
    )
    return canvas


def write_tga(image: Image.Image, path: Path) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    # compression=None: Pillow only writes RLE when asked for "tga_rle",
    # and the client will not read RLE -- make_minimap_icon.py's own rule.
    image.save(path, format="TGA", compression=None)
    print(f"wrote {path} ({path.stat().st_size} bytes)")


def main() -> int:
    for slug in CLASSES:
        write_tga(render_crest(slug), OUTPUT_DIR / f"{slug}.tga")
    write_tga(render_ring(), OUTPUT_DIR / "ring.tga")
    return 0


if __name__ == "__main__":
    sys.exit(main())
