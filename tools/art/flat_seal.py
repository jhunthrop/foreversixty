"""The flat seal at small sizes (owner 2026-11-10): one source for the web, the addon and the companion.

The painted seal v4 stays at 64 px and up (app icons, Discord, store listings, the home hero, apple-touch and OG
images). At 44 px and below it turns to mud, so those sizes take the flat vector seal that design/mocks/seal_flat.py
traced into design/assets/brand/seal-flat: the small seal (with the swash) from 25 to 48 px, the tiny bold seal
(heavier outline, no swash) at 24 px and below. The SVGs are used as they are, never redrawn.

Rasters come from cairosvg (pip install cairosvg; it needs the cairo library, brew install cairo), rendered
straight at the target size so every pixel size is its own render, never a downscale.
"""
import io
import re
from pathlib import Path

from PIL import Image

REPO = Path(__file__).resolve().parents[2]
FLAT_DIR = REPO / 'design/assets/brand/seal-flat'
SMALL_SVG = FLAT_DIR / 'seal-flat-small.svg'
TINY_SVG = FLAT_DIR / 'seal-flat-tiny-bold.svg'
# The tiny bold seal up to this many pixels; the small seal above it.
TINY_MAX_PX = 24
# The tiny bold seal draws the LX with a dark stroke centred on the letters' outline, 95 units wide in the path's
# own coordinates. heavier_lx() lays a wider copy of that stroke under the letters, so the outline grows outward
# only and the gold keeps its traced shape; HEAVY_LX_OUTLINE units each side keeps the X's notches open.
LX_STROKE = re.compile(r'(<path [^>]*?)stroke-width="95\.0000"([^>]*?/>)', re.S)
HEAVY_LX_OUTLINE = 30
# The flat seal ends here: a raster above it is an icon the painted seal owns.
FLAT_MAX_PX = 48


def flat_svg(px: int) -> Path:
    """The flat seal's SVG for a size in pixels."""
    if px > FLAT_MAX_PX:
        raise ValueError(f'{px} px is the painted seal\'s size; the flat seal ends at {FLAT_MAX_PX} px')
    return TINY_SVG if px <= TINY_MAX_PX else SMALL_SVG


def heavier_lx(svg: str) -> str:
    """The tiny bold seal with its LX outline HEAVY_LX_OUTLINE units heavier: the same letters, stroke under fill."""
    match = LX_STROKE.search(svg)
    if not match:
        raise ValueError('the tiny bold seal has no 95-unit LX stroke to thicken; was the SVG regenerated?')
    under = f'{match.group(1)}stroke-width="{95 + 2 * HEAVY_LX_OUTLINE}"{match.group(2)}'
    return svg.replace(match.group(0), under + match.group(0), 1)


def flat_seal(px: int, shown_px: int | None = None, heavy: bool = False) -> Image.Image:
    """The flat seal rendered at px x px, RGBA with the SVG's own transparency.

    shown_px is the size the seal is drawn at when that differs from the raster (a WoW texture is a power of
    two, drawn smaller); the seal for that size is the one rendered. It defaults to px. heavy draws the tiny
    bold seal's LX with a heavier outline (heavier_lx), for a 16 px tray where the plain outline goes soft.
    """
    import cairosvg  # imported here so the painted-only paths of the callers run without cairo

    path = flat_svg(px if shown_px is None else shown_px)
    if heavy and path != TINY_SVG:
        raise ValueError(f'the heavy LX is a tiny bold seal variant; {px} px takes the small seal')
    svg = path.read_text()
    png = cairosvg.svg2png(bytestring=(heavier_lx(svg) if heavy else svg).encode(), output_width=px, output_height=px)
    return Image.open(io.BytesIO(png)).convert('RGBA')
