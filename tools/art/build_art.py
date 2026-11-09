"""The web's copy of the owner's round-5 art, seals and paperdoll frame (design/specs/2026-11-09-directions.md).

    python3 tools/art/build_art.py

Reads design/assets (the owner's sources) and the mock pipeline's processing (design/mocks/directions: art.py for the
periodic, band-shaped ground tile; paperdoll/ for the cleaned frame and slots.json), and writes every file the site
serves under web/public/art plus web/src/lib/ui/paperdoll-slots.json. web/src/lib/ui/assets.ts is the one module that
names these files; pages never hard-code a path. The build is deterministic: run twice, the outputs are identical.

Generated art is the owner's explicit choice (2026-11-09): background only, never captioned as a screenshot.
"""
import json
import shutil
import sys
from pathlib import Path

from PIL import Image

REPO = Path(__file__).resolve().parents[2]
DESIGN = REPO / 'design'
DIRECTIONS = DESIGN / 'mocks/directions'
OUT = REPO / 'web/public/art'
SLOTS_OUT = REPO / 'web/src/lib/ui/paperdoll-slots.json'

sys.path.insert(0, str(DIRECTIONS))
from art import ground_tile  # noqa: E402  (the mock pipeline's tile: one source for mocks, web and apps)

WEBP_QUALITY = 80
# Each art file at the widths the site asks for (lib/ui/assets.ts: the phone width, then the desktop width).
ART = {
    'band': ('art/blue-band.jpeg', (1200, 2400)),
    'hero': ('art/hero-road-camp-wide.jpeg', (1200, 2400)),
    'cave': ('art/raid-panel-cave.jpeg', (800, 1600)),
}
# One logo (owner 2026-11-09): the seal v3, cut out with real transparency, everywhere: nav, footer,
# favicons and the link preview (which reads the apple-touch icon). Its bold rim holds at 16 px.
SEAL_SOURCE = 'logo/lx-seal-v3.png'
SEALS = {'seal-88': (SEAL_SOURCE, 176), 'seal-128': ('logo/lx-seal-v3-256.png', 256)}
FAVICONS = {'favicon-32.png': ('logo/lx-seal-v3-32.png', 32), 'apple-touch-icon.png': (SEAL_SOURCE, 180)}
# favicon.ico carries the owner's sized renders, never a downscale of one size.
FAVICON_ICO = ('logo/lx-seal-v3-16.png', 'logo/lx-seal-v3-32.png', 'logo/lx-seal-v3-48.png')
OG_SIZE, OG_BOX = (1200, 630), (1180, 480, 2600, 1225)  # art_crops.py: the glow over the pass, the road's head


def source(rel: str) -> Path:
    path = DESIGN / 'assets' / rel
    if not path.exists():
        raise FileNotFoundError(f'{path} not found')
    return path


def save_webp(im: Image.Image, name: str) -> None:
    im.save(OUT / name, 'WEBP', quality=WEBP_QUALITY, method=6)


def scaled(im: Image.Image, width: int) -> Image.Image:
    return im.resize((width, round(im.height * width / im.width)), Image.LANCZOS)


def build_art() -> None:
    save_webp(ground_tile(), 'ground.webp')
    for key, (rel, widths) in ART.items():
        im = Image.open(source(rel)).convert('RGB')
        for width in widths:
            save_webp(scaled(im, width), f'{key}-{width}.webp')
    og = Image.open(source('art/social-banner-valley.jpeg')).convert('RGB').crop(OG_BOX).resize(OG_SIZE, Image.LANCZOS)
    og.save(OUT / 'og-art.jpg', 'JPEG', quality=86, optimize=True)


def build_seals() -> None:
    for name, (rel, px) in SEALS.items():
        Image.open(source(rel)).convert('RGBA').resize((px, px), Image.LANCZOS).save(OUT / f'{name}.webp', 'WEBP', quality=88, method=6)
    for name, (rel, px) in FAVICONS.items():
        Image.open(source(rel)).convert('RGBA').resize((px, px), Image.LANCZOS).save(REPO / 'web/public' / name, optimize=True)
    sized = [Image.open(source(rel)).convert('RGBA') for rel in FAVICON_ICO]
    largest = sized[-1]
    largest.save(REPO / 'web/public/favicon.ico', format='ICO', sizes=[im.size for im in sized],
                 append_images=sized[:-1])


def build_paperdoll() -> None:
    for scale in ('1x', '2x'):
        shutil.copyfile(DIRECTIONS / f'paperdoll/frame@{scale}.webp', OUT / f'paperdoll-frame@{scale}.webp')
    slots = json.loads((DIRECTIONS / 'paperdoll/slots.json').read_text())
    SLOTS_OUT.write_text(json.dumps(slots, indent=1) + '\n')


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    build_art()
    build_seals()
    build_paperdoll()
    for path in sorted(OUT.iterdir()):
        print(f'{path.relative_to(REPO)}  {path.stat().st_size // 1024} KB')


if __name__ == '__main__':
    main()
