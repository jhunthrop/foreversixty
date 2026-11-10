"""Cut the owner's painted LX seal out of its square render: real transparency, then the sized renders.

    python3 tools/art/cut_seal.py design/assets/logo/lx-seal-source-v4.jpeg v4

Writes design/assets/logo/lx-seal-<version>.png (the master, RGBA, as large as the source allows) and
lx-seal-<version>-<size>.png for every size the site, addon and companion read (SIZES). Needs OpenCV
(pip install opencv-python-headless), Pillow and numpy; pngquant and oxipng on PATH compress the outputs.

The coin sits on a flat, near-black brown ground with a thin dark rim round its gold edge. The rim is the
coin's own and is kept; the ground goes. The cut is an ellipse fitted (cv2.fitEllipse) to the outline of
everything brighter than the ground, so JPEG speckle on the edge never bends the outline, and the mask is
that ellipse drawn analytically: inset INSET px so no ground pixel survives, with a FEATHER px ramp for the
antialiased edge. The rim band alone is lightly denoised (the JPEG blocks show there at full size), the
paint inside it is never touched. The build is deterministic: run twice, the outputs match.
"""
import argparse
import subprocess
import sys
from pathlib import Path

import cv2
import numpy as np
from PIL import Image

REPO = Path(__file__).resolve().parents[2]
LOGO = REPO / 'design/assets/logo'
# Every v3 size, kept for v4: favicons 16/32/48, nav 40 (80 at 2x), footer 88, the addon's 32/64, the
# companion's 24/40/128/256/512, store uploads 512/1024.
SIZES = (16, 24, 32, 40, 48, 64, 80, 88, 128, 256, 512, 1024)
GROUND_DISTANCE = 9.0  # RGB distance from the ground colour that counts as coin (ground ~1 sd, rim ~22 away)
INSET = 1.5  # px the cut sits inside the fitted outline, past the JPEG's ground-to-rim transition
FEATHER = 1.5  # px of antialiased ramp
EDGE_BAND = 14  # px inside the cut that the edge denoise may touch
PAD = 2  # transparent px round the coin in the master


def ground_colour(rgb: np.ndarray) -> np.ndarray:
    """The ground: the median of the four 64 px corners, which the coin never reaches."""
    k = 64
    corners = np.concatenate([rgb[:k, :k], rgb[:k, -k:], rgb[-k:, :k], rgb[-k:, -k:]]).reshape(-1, 3)
    return np.median(corners, axis=0)


def fit_outline(rgb: np.ndarray) -> tuple[tuple[float, float], tuple[float, float], float]:
    """The ellipse fitted to the outline of the largest region that is not ground."""
    distance = np.linalg.norm(rgb.astype(np.float32) - ground_colour(rgb), axis=2)
    coin = (distance > GROUND_DISTANCE).astype(np.uint8)
    coin = cv2.morphologyEx(coin, cv2.MORPH_OPEN, np.ones((5, 5), np.uint8))
    contours, _ = cv2.findContours(coin, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_NONE)
    if not contours:
        raise ValueError('no coin found on the ground')
    outline = max(contours, key=cv2.contourArea)
    return cv2.fitEllipse(outline)


def signed_distance(shape: tuple[int, int], ellipse) -> np.ndarray:
    """Approximate signed distance in px to the ellipse (negative inside), from the normalised radius."""
    (cx, cy), (w, h), angle = ellipse
    ys, xs = np.mgrid[0:shape[0], 0:shape[1]].astype(np.float64)
    # Pixel centres: OpenCV's contour coordinates are pixel indices, so the boundary pixel itself is "on".
    t = np.deg2rad(angle)
    dx, dy = xs - cx, ys - cy
    u = dx * np.cos(t) + dy * np.sin(t)
    v = -dx * np.sin(t) + dy * np.cos(t)
    a, b = w / 2, h / 2
    r = np.sqrt((u / a) ** 2 + (v / b) ** 2)
    # The local radius along the ray converts the normalised overshoot into px.
    local = np.hypot(u, v) / np.maximum(r, 1e-9)
    return (r - 1.0) * local


def alpha_from(distance: np.ndarray) -> np.ndarray:
    """Full inside the inset outline, a linear FEATHER ramp out, nothing beyond."""
    return np.clip((-(distance + INSET)) / FEATHER + 0.5, 0.0, 1.0)


def denoise_edge(rgb: np.ndarray, distance: np.ndarray) -> np.ndarray:
    """A light edge-preserving smooth, blended in only across the rim band."""
    smooth = cv2.bilateralFilter(np.ascontiguousarray(rgb), d=5, sigmaColor=18, sigmaSpace=3)
    depth = -distance - INSET  # px inside the cut
    weight = np.clip(1.0 - depth / EDGE_BAND, 0.0, 1.0)[..., None]
    return (rgb * (1 - weight) + smooth * weight).round().astype(np.uint8)


def cut(source: Path) -> Image.Image:
    rgb = np.array(Image.open(source).convert('RGB'))
    ellipse = fit_outline(rgb)
    distance = signed_distance(rgb.shape[:2], ellipse)
    clean = denoise_edge(rgb, distance)
    alpha = (alpha_from(distance) * 255).round().astype(np.uint8)
    rgba = np.dstack([clean, alpha])
    (cx, cy), (w, h), _ = ellipse
    half = int(np.ceil(max(w, h) / 2)) + PAD
    left, top = round(cx) - half, round(cy) - half
    if left < 0 or top < 0 or left + 2 * half > rgb.shape[1] or top + 2 * half > rgb.shape[0]:
        raise ValueError(f'the coin {ellipse} runs off the {rgb.shape[1]} x {rgb.shape[0]} source')
    print(f'ellipse centre ({cx:.1f}, {cy:.1f}) axes {w:.1f} x {h:.1f} angle {ellipse[2]:.1f}', file=sys.stderr)
    return Image.fromarray(rgba[top:top + 2 * half, left:left + 2 * half], 'RGBA')


def compress(path: Path, lossy: bool) -> None:
    """pngquant (sized renders only, quality 90 to 100) then oxipng, keeping the smaller result."""
    if lossy:
        subprocess.run(['pngquant', '--quality=90-100', '--skip-if-larger', '--force', '--strip',
                        '--output', str(path), str(path)], check=False)
    subprocess.run(['oxipng', '-o', '4', '--strip', 'safe', '-q', str(path)], check=True)


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__.split('\n\n')[0])
    parser.add_argument('source', type=Path)
    parser.add_argument('version', help='the file name tag, e.g. v4')
    args = parser.parse_args()
    master = cut(args.source)
    master_path = LOGO / f'lx-seal-{args.version}.png'
    master.save(master_path)
    compress(master_path, lossy=False)
    for size in SIZES:
        path = LOGO / f'lx-seal-{args.version}-{size}.png'
        master.resize((size, size), Image.LANCZOS).save(path)
        compress(path, lossy=True)
    print(f'wrote {master_path} {master.size} and {len(SIZES)} sized renders')


if __name__ == '__main__':
    main()
