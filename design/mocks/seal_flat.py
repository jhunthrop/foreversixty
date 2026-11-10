#!/usr/bin/env python3
"""Flat vector Forever Sixty seal: trace, small variants, PNG renders, comparison sheet.

Needs: Pillow, numpy, scipy, cairosvg, and the `potrace` binary (brew install potrace).
Run:   python3 design/mocks/seal_flat.py
Inputs are the AI flat redraw and the painted seal; outputs go to design/assets/brand/seal-flat/.
"""
import io
import re
import subprocess
import tempfile
from pathlib import Path

import cairosvg
import numpy as np
from PIL import Image, ImageDraw, ImageFont
from scipy import ndimage as ndi

FLAT_SRC = Path("/Users/jh/Downloads/logo-new-flat.jpeg")
PAINTED_SRC = Path("/Users/jh/Desktop/forever-seal-v4/lx-seal-v4-1024.png")
OUT = Path(__file__).resolve().parents[1] / "assets/brand/seal-flat"

FIELD, GOLD, DARK_GOLD, OUTLINE = "#2C8A9E", "#C9A465", "#896840", "#170E09"
VIEW = 512
WORK = 1024  # trace resolution of the cropped seal

# Nearest-colour classes sampled from the owner's recoloured flat redraw (traced, then snapped to the hexes above).
CLASS_COLOURS = {
    "bg": (255, 255, 255),
    "teal": (46, 137, 155),
    "gold": (201, 162, 98),
    "dkgold": (115, 82, 48),
    "dark": (28, 16, 10),
}
SMALL_SRC = Path("/Users/jh/Downloads/small.jpeg")
TINY_SRC = Path("/Users/jh/Downloads/tiny.jpeg")
TRACE_COLOURS = {"bg": (255, 255, 255), "teal": (47, 131, 150), "gold": (196, 157, 96), "outline": (28, 18, 12)}
BOLD_STROKE = 12.0  # 512-unit dark outline in the bold tiny (about 2x the source)
BOLD_BAND = 0.10  # gold ring band as a fraction of the seal radius
BOLD_SPAN = 0.75  # LX width as a fraction of the field width
BOLD_EDGE_SHIFT_PX = 3  # work px: moves the stroke centre-line out to where the source outline sat
MIN_MONOGRAM_AREA = 0.01  # of WORK^2: smaller dark/gold bits near the field edge are not letters
SIZES = (16, 24, 32, 44, 64)
COMPARE_SIZES = (16, 24, 32, 44, 64)
BACKGROUNDS = (("header #16282F", "#16282F"), ("parchment #E9DFC8", "#E9DFC8"), ("white", "#FFFFFF"))


def disk(r):
    r = max(1, int(round(r)))
    y, x = np.ogrid[-r : r + 1, -r : r + 1]
    return x * x + y * y <= r * r


def classify(img):
    arr = np.asarray(img, dtype=np.int32)
    names = list(CLASS_COLOURS)
    cols = np.array([CLASS_COLOURS[n] for n in names], dtype=np.int32)
    d = ((arr[:, :, None, :] - cols[None, None]) ** 2).sum(-1)
    return names, d.argmin(-1)


def crop_square(img):
    mask = np.asarray(img.convert("L")) < 235
    ys, xs = np.where(mask)
    x0, x1, y0, y1 = xs.min(), xs.max() + 1, ys.min(), ys.max() + 1
    cx, cy, half = (x0 + x1) / 2, (y0 + y1) / 2, max(x1 - x0, y1 - y0) / 2 + 2
    box = tuple(int(round(v)) for v in (cx - half, cy - half, cx + half, cy + half))
    return img.crop(box), box


def masks_from(img):
    names, lab = classify(img)
    m = {n: lab == i for i, n in enumerate(names)}
    r2 = disk(2)
    teal = ndi.binary_opening(m["teal"], r2)
    gold = ndi.binary_opening(m["gold"] | m["dkgold"], r2)
    outline = ndi.binary_opening(m["dark"], disk(1))
    detail = ndi.binary_opening(m["dkgold"], disk(1))
    return {"teal": teal, "gold": gold, "outline": outline, "thin": detail}


def trace(mask, size):
    """potrace a boolean mask; return (path d, transform) in the 512 viewBox."""
    bmp = Image.fromarray((~mask * 255).astype(np.uint8)).resize((size, size), Image.NEAREST).convert("1")
    with tempfile.TemporaryDirectory() as td:
        src, dst = Path(td, "m.pbm"), Path(td, "m.svg")
        bmp.save(src)
        subprocess.run(
            ["potrace", "-b", "svg", "-t", "30", "-a", "1.1", "-O", "1.2", "-u", "4", "-o", str(dst), str(src)],
            check=True,
        )
        svg = dst.read_text()
    paths = re.findall(r'<path d="([^"]+)"', svg)
    s = float(re.search(r"scale\(([\d.]+),", svg).group(1))
    k = VIEW / size * s
    return " ".join(p.strip() for p in paths), f"translate(0 {VIEW}) scale({k:.6g} -{k:.6g})"


def layer(mask, fill, size, extra=""):
    d, tf = trace(mask, size)
    return f'<path fill="{fill}"{extra} transform="{tf}" d="{d}"/>'


def svg_doc(body, title):
    return (
        f'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 {VIEW} {VIEW}" role="img">'
        f"<title>{title}</title>{body}</svg>\n"
    )


def build_full(m, size):
    under = disk(3)
    return svg_doc(
        layer(ndi.binary_dilation(m["teal"], under), FIELD, size)
        + layer(ndi.binary_dilation(m["gold"], under), GOLD, size)
        + layer(m["outline"], OUTLINE, size)
        + layer(m["thin"] & ~m["outline"], DARK_GOLD, size),
        "Forever Sixty seal (flat)",
    )


def snap_classes(img):
    """Snap every pixel to the nearest of bg, field, gold, outline (the three brand colours plus the white page)."""
    names = list(TRACE_COLOURS)
    cols = np.array([TRACE_COLOURS[n] for n in names], dtype=np.int32)
    d = ((np.asarray(img, dtype=np.int32)[:, :, None, :] - cols[None, None]) ** 2).sum(-1)
    lab = d.argmin(-1)
    return {n: lab == i for i, n in enumerate(names)}


def circle_crop(img):
    """Crop tight to the outer circle (fitted to the non-white area) and return it with the circle's work-px centre."""
    nonwhite = np.asarray(img.convert("L")) < 235
    ys, xs = np.where(nonwhite)
    cx, cy, r = xs.mean(), ys.mean(), np.sqrt(nonwhite.sum() / np.pi)
    box = tuple(int(round(v)) for v in (cx - r, cy - r, cx + r, cy + r))
    return img.crop(box).resize((WORK, WORK), Image.LANCZOS)


def split_monogram(cls, field_circle):
    """Pixels of the LX (and swash) lie inside the field disc, clear of its edge outline."""
    n = WORK
    yy, xx = np.mgrid[0:n, 0:n]
    fx, fy, fr = field_circle
    inside = np.hypot(xx - fx, yy - fy) < fr * 0.95
    lab, cnt = ndi.label(inside & ~cls["teal"])
    sizes = ndi.sum(inside & ~cls["teal"], lab, range(1, cnt + 1))
    return np.isin(lab, np.nonzero(sizes >= MIN_MONOGRAM_AREA * n * n)[0] + 1)  # edge slivers stay with the ring


def fit_field(cls):
    ys, xs = np.where(cls["teal"])
    return (xs.min() + xs.max()) / 2, (ys.min() + ys.max()) / 2, ((xs.max() - xs.min()) + (ys.max() - ys.min())) / 4


def trace_source(path, centre_monogram=False):
    """Faithful three-colour trace of an owner redraw; only the outer silhouette is made a true circle."""
    cls = snap_classes(circle_crop(Image.open(path).convert("RGB")))
    under = disk(3)
    cls = {k: ndi.binary_opening(v, disk(1)) for k, v in cls.items()}
    fx, fy, fr = fit_field(cls)
    mono = split_monogram(cls, (fx, fy, fr))
    mono_outline = ndi.binary_dilation(mono, disk(8)) & (cls["outline"] | cls["gold"])
    field = ndi.binary_fill_holes(cls["teal"] | mono)
    ring_gold = ndi.binary_dilation(cls["gold"], under) & ~mono_outline
    ring_outline = cls["outline"] & ~mono_outline
    half = VIEW / 2
    body = (
        f'<clipPath id="c"><circle cx="{half}" cy="{half}" r="{half}"/></clipPath>'
        f'<g clip-path="url(#c)"><circle cx="{half}" cy="{half}" r="{half}" fill="{OUTLINE}"/>'
        + layer(field, FIELD, WORK)
        + layer(ring_gold, GOLD, WORK)
        + layer(ring_outline, OUTLINE, WORK)
    )
    mono_gold = ndi.binary_dilation(cls["gold"] & mono_outline, under)
    unit = VIEW / WORK
    dy = (fy - _vertical_centre(mono_outline)) * unit if centre_monogram else 0.0
    tf = f' transform="translate(0 {dy:.1f})"' if centre_monogram else ""
    body += f"<g{tf}>" + layer(mono_gold, GOLD, WORK) + layer(mono_outline & cls["outline"], OUTLINE, WORK) + "</g></g>"
    return svg_doc(body, "Forever Sixty seal")


def build_tiny_bold(path):
    """Geometric bold tiny: true-circle single-band ring and the traced LX fill, rescaled and stroked (not bloated)."""
    cls = snap_classes(circle_crop(Image.open(path).convert("RGB")))
    cls = {k: ndi.binary_opening(v, disk(1)) for k, v in cls.items()}
    fill = ndi.binary_dilation(cls["gold"] & split_monogram(cls, fit_field(cls)), disk(BOLD_EDGE_SHIFT_PX))
    d, tf = trace(fill, WORK)
    k = float(re.search(r"scale\(([-\d.e]+)", tf).group(1))
    ys, xs = np.where(fill)
    unit = VIEW / WORK
    w, h = (xs.max() - xs.min()) * unit, (ys.max() - ys.min()) * unit
    cx, cy = (xs.max() + xs.min()) / 2 * unit, (ys.max() + ys.min()) / 2 * unit
    half = VIEW / 2
    outer_c = half - BOLD_STROKE / 2  # centre-line radius of the rim stroke
    inner_c = outer_c - (BOLD_BAND * half + BOLD_STROKE)  # gold band is BOLD_BAND of the radius between the strokes
    field_w = 2 * (inner_c - BOLD_STROKE / 2)
    scale = (BOLD_SPAN * field_w - BOLD_STROKE) / w
    ring = (
        f'<circle cx="{half}" cy="{half}" r="{outer_c:.2f}" fill="{GOLD}" stroke="{OUTLINE}" stroke-width="{BOLD_STROKE}"/>'
        f'<circle cx="{half}" cy="{half}" r="{inner_c:.2f}" fill="{FIELD}" stroke="{OUTLINE}" stroke-width="{BOLD_STROKE}"/>'
    )
    group = (
        f'<g transform="translate({half} {half}) scale({scale:.5f}) translate({-cx:.2f} {-cy:.2f})">'
        f'<path fill="{GOLD}" stroke="{OUTLINE}" stroke-linejoin="round" stroke-width="{BOLD_STROKE / scale / k:.4f}" '
        f'transform="{tf}" d="{d}"/></g>'
    )
    return svg_doc(ring + group, "Forever Sixty seal (tiny, bold)")


def _vertical_centre(mask):
    rows = np.nonzero(mask.any(axis=1))[0]
    return (rows[0] + rows[-1]) / 2


def render(svg_text, px):
    png = cairosvg.svg2png(bytestring=svg_text.encode(), output_width=px, output_height=px)
    return Image.open(io.BytesIO(png)).convert("RGBA")


def painted(px):
    return Image.open(PAINTED_SRC).convert("RGBA").resize((px, px), Image.LANCZOS)


def font(sz):
    for p in ("/System/Library/Fonts/Helvetica.ttc", "/System/Library/Fonts/SFNS.ttf"):
        try:
            return ImageFont.truetype(p, sz)
        except OSError:
            continue
    return ImageFont.load_default()


def grid(rows, bg, label, zoom=2, pad=10, label_w=150):
    sizes = COMPARE_SIZES
    cell_w = max(sizes) * zoom + pad * 2
    row_h = max(sizes) * zoom + pad * 2
    w, h = label_w + cell_w * len(sizes), 28 + row_h * len(rows) + 24
    sheet = Image.new("RGB", (w, h), bg)
    d = ImageDraw.Draw(sheet)
    lum = sum(int(bg[i : i + 2], 16) for i in (1, 3, 5)) / 3
    ink = "#FFFFFF" if lum < 110 else "#111111"
    d.text((8, 6), f"background: {label}", fill=ink, font=font(14))
    for j, s in enumerate(sizes):
        d.text((label_w + j * cell_w + pad, 28 + row_h * len(rows) + 2), f"{s} px", fill=ink, font=font(13))
    for i, (name, fn) in enumerate(rows):
        y0 = 28 + i * row_h
        d.text((8, y0 + row_h // 2 - 8), name, fill=ink, font=font(13))
        for j, s in enumerate(sizes):
            tile = fn(s).resize((s * zoom, s * zoom), Image.NEAREST)
            sheet.paste(tile, (label_w + j * cell_w + pad, y0 + (row_h - s * zoom) // 2), tile)
    return sheet


def write_small_check(variants, zoom=4):
    """4x nearest-neighbour crops on the header colour: small at 32/44 px, tiny and tiny-centred at 16/24 px."""
    picks = [("seal-flat-small", 32), ("seal-flat-small", 44), ("seal-flat-tiny", 16), ("seal-flat-tiny", 24),
             ("seal-flat-tiny-centred", 16), ("seal-flat-tiny-centred", 24),
             ("seal-flat-tiny-bold", 16), ("seal-flat-tiny-bold", 24), ("seal-flat-tiny-bold", 32)]
    tiles = [render(variants[n], s).resize((s * zoom, s * zoom), Image.NEAREST) for n, s in picks]
    sheet = Image.new("RGB", (sum(t.width + 10 for t in tiles) + 10, max(t.height for t in tiles) + 20), BACKGROUNDS[0][1])
    x = 10
    for t in tiles:
        sheet.paste(t, (x, 10), t)
        x += t.width + 10
    sheet.save(OUT / "small-check.png")


def main():
    """seal-flat.svg is traced from FLAT_SRC; small/tiny are faithful traces of the owner's own small redraws.
    The earlier hand-built small/tiny live on as seal-flat-*-built.svg (reference only, not generated here)."""
    OUT.mkdir(parents=True, exist_ok=True)
    img, _ = crop_square(Image.open(FLAT_SRC).convert("RGB"))
    m = masks_from(img.resize((WORK, WORK), Image.LANCZOS))
    variants = {
        "seal-flat": build_full(m, WORK),
        "seal-flat-small": trace_source(SMALL_SRC),
        "seal-flat-tiny": trace_source(TINY_SRC),
        "seal-flat-tiny-centred": trace_source(TINY_SRC, centre_monogram=True),
        "seal-flat-tiny-bold": build_tiny_bold(TINY_SRC),
    }
    for name, text in variants.items():
        (OUT / f"{name}.svg").write_text(text)
        for s in SIZES:
            render(text, s).save(OUT / f"{name}-{s}.png")
    rows = [("painted seal", painted)] + [
        (n.replace("seal-flat-", ""), lambda s, n=n: render(variants[n], s))
        for n in ("seal-flat-small", "seal-flat-tiny", "seal-flat-tiny-centred", "seal-flat-tiny-bold")
    ]
    grids = [grid(rows, c, n) for n, c in BACKGROUNDS]
    sheet = Image.new("RGB", (max(g.width for g in grids), sum(g.height for g in grids)), "#888888")
    y = 0
    for g in grids:
        sheet.paste(g, (0, y))
        y += g.height
    sheet.save(OUT / "comparison.png")
    write_small_check(variants)
    side = Image.new("RGB", (128 * 2 + 30, 128 + 20), "#FFFFFF")
    side.paste(painted(128), (10, 10))
    side.paste(render(variants["seal-flat"], 128), (148, 10))
    side.save(OUT / "side-by-side-128.png")


if __name__ == "__main__":
    main()
