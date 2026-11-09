"""The owner's round-5 art set, prepared for the mocks (design/specs/2026-11-09-directions.md, round 5).

The files live in design/assets (committed, the owner's sources); everything here works on them in
memory and embeds the result as a data URI. Generated art is the owner's explicit choice for art assets (2026-11-09),
used as background only and never captioned as a screenshot.
"""
import base64
from functools import lru_cache
from io import BytesIO

import numpy as np
from PIL import Image

from assets import owner_asset

# The painted ground at 1:1 on a 1x screen: large enough that the soft field never reads as a repeat.
GROUND_TILE_PX = 1600
JPEG_QUALITY = 84
HERO, BAND, CAVE = 'hero-road-camp-wide.jpeg', 'blue-band.jpeg', 'raid-panel-cave.jpeg'
HERO_PX, BAND_PX, CAVE_PX = 2400, 2400, 1600


def _jpeg_uri(im: Image.Image, quality: int = JPEG_QUALITY) -> str:
    buf = BytesIO()
    im.convert('RGB').save(buf, 'JPEG', quality=quality, optimize=True)
    return 'data:image/jpeg;base64,' + base64.b64encode(buf.getvalue()).decode()


LUMA = np.array([.2126, .7152, .0722])


def periodic(u: np.ndarray) -> np.ndarray:
    """Moisan's periodic component: the image minus the smooth field that carries its border mismatch, so it tiles with
    no seam and, unlike offset-and-feather, repeats nothing inside the tile."""
    h, w = u.shape
    v = np.zeros_like(u)
    v[0, :] += u[-1, :] - u[0, :]
    v[-1, :] += u[0, :] - u[-1, :]
    v[:, 0] += u[:, -1] - u[:, 0]
    v[:, -1] += u[:, 0] - u[:, -1]
    den = 2 * np.cos(2 * np.pi * np.arange(h)[:, None] / h) + 2 * np.cos(2 * np.pi * np.arange(w)[None, :] / w) - 4
    den[0, 0] = 1
    s = np.fft.fft2(v) / den
    s[0, 0] = 0
    return u - np.real(np.fft.ifft2(s))


def blur(u: np.ndarray, sigma: float) -> np.ndarray:
    """Circular gaussian blur in the frequency domain, so a tileable input stays tileable."""
    fy = np.fft.fftfreq(u.shape[0])[:, None]
    fx = np.fft.fftfreq(u.shape[1])[None, :]
    return np.real(np.fft.ifft2(np.fft.fft2(u) * np.exp(-2 * (np.pi * sigma) ** 2 * (fx ** 2 + fy ** 2))))


# Owner 2026-11-08: the ground needs the depth of Blizzard's site; reviewers: no stains (dark patches with edges), no
# visible repeat. The tile's luminance is split into three bands and rebuilt:
#   field (blur sigma GROUND_FIELD_SIGMA): the large soft mottling, gained, then soft-clamped both ways at GROUND_FIELD_CAP
#     so no patch sits far darker or lighter than the ground around it;
#   strokes (sigma 3 to the field): the brush edges that read as stains, damped;
#   grain (under sigma 3): the fine surface their ground has, lifted.
# Measured at 1x: sd 2.8 over 200 px windows (Blizzard's clean ground 2.1 to 4.9), darkest 200 px window 6.1 below the
# mean (8.8 before), half-tile correlation -0.07 (0.88 with offset-and-feather), seam steps half the interior step.
GROUND_FIELD_SIGMA, GROUND_FIELD_GAIN, GROUND_FIELD_CAP = 30, 3.0, 6.0
GROUND_STROKE_GAIN, GROUND_GRAIN_GAIN, GROUND_GRAIN_SIGMA = .6, 2.2, 3


def soften(im: Image.Image, k: float) -> Image.Image:
    """Pull every pixel toward the tile's mean colour by (1 - k): same hue, a fraction of the variation."""
    a = np.asarray(im.convert('RGB')).astype(float)
    mean = a.reshape(-1, 3).mean(axis=0)
    return Image.fromarray(np.clip(mean + (a - mean) * k, 0, 255).astype(np.uint8))


@lru_cache(maxsize=None)
def ground_tile() -> Image.Image:
    src = Image.open(owner_asset('art/brown-ground-tile.jpeg')).convert('RGB').resize((GROUND_TILE_PX, GROUND_TILE_PX), Image.LANCZOS)
    a = np.asarray(src).astype(float)
    lum = periodic(a @ LUMA)
    mean = lum.mean()
    d = lum - mean
    field = blur(d, GROUND_FIELD_SIGMA)
    fine = blur(d, GROUND_GRAIN_SIGMA)
    shaped = (GROUND_FIELD_CAP * np.tanh(GROUND_FIELD_GAIN * field / GROUND_FIELD_CAP)
              + GROUND_STROKE_GAIN * (fine - field) + GROUND_GRAIN_GAIN * (d - fine))
    # the luminance change rides on the tile's mean colour, so the hue stays the ground's
    rgb = a.reshape(-1, 3).mean(axis=0) * (1 + shaped / mean)[..., None]
    return Image.fromarray(np.clip(rgb, 0, 255).astype(np.uint8))


@lru_cache(maxsize=None)
def ground_uri() -> str:
    return _jpeg_uri(ground_tile())


@lru_cache(maxsize=None)
def art_uri(name: str, width: int, box: tuple[int, int, int, int] | None = None) -> str:
    """One art file, optionally cropped (source pixels), scaled to `width`."""
    im = Image.open(owner_asset(f'art/{name}')).convert('RGB')
    if box:
        im = im.crop(box)
    return _jpeg_uri(im.resize((width, round(im.height * width / im.width)), Image.LANCZOS))


def hero_uri() -> str:
    return art_uri(HERO, HERO_PX)


def band_uri() -> str:
    return art_uri(BAND, BAND_PX)


def cave_uri() -> str:
    return art_uri(CAVE, CAVE_PX)


def seam_check(path) -> None:
    """A 2x2 of the prepared tile, for the gallery: no seam may show at the cross."""
    t = ground_tile()
    sheet = Image.new('RGB', (t.width * 2, t.height * 2))
    for x in (0, t.width):
        for y in (0, t.height):
            sheet.paste(t, (x, y))
    sheet.resize((t.width, t.height)).save(path)
