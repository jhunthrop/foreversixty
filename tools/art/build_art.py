"""The web's copy of the owner's round-5 art, seals and paperdoll frame (design/specs/2026-11-09-directions.md).

    python3 tools/art/build_art.py

Reads design/assets (the owner's sources) and the mock pipeline's processing (design/mocks/directions: art.py for the
periodic, band-shaped ground tile; paperdoll/ for the cleaned frame and slots.json), and writes every file the site
serves under web/public/art plus web/src/lib/ui/paperdoll-slots.json, and the desktop companion's copy of the ground
and the band (companion/ui/art, embedded in its binary). web/src/lib/ui/assets.ts is the one module that
names these files; pages never hard-code a path. The build is deterministic: run twice, the outputs are identical.

Generated art is the owner's explicit choice (2026-11-09): background only, never captioned as a screenshot.
"""
import json
import shutil
import subprocess
import sys
from pathlib import Path

from PIL import Image

REPO = Path(__file__).resolve().parents[2]
DESIGN = REPO / 'design'
DIRECTIONS = DESIGN / 'mocks/directions'
OUT = REPO / 'web/public/art'
SLOTS_OUT = REPO / 'web/src/lib/ui/paperdoll-slots.json'
COMPANION_OUT = REPO / 'companion/ui/art'
# The companion window is one 800 px window: the ground tile and the phone-width band, byte for byte the site's files.
COMPANION_ART = ('ground.webp', 'band-1200.webp')

sys.path.insert(0, str(DIRECTIONS))
sys.path.insert(0, str(Path(__file__).resolve().parent))
from art import ground_tile  # noqa: E402  (the mock pipeline's tile: one source for mocks, web and apps)
from flat_seal import SMALL_SVG, TINY_SVG, flat_seal  # noqa: E402  (the flat seal at 44 px and below)

WEBP_QUALITY = 80
# Each art file at the widths the site asks for (lib/ui/assets.ts: the phone width, then the desktop width).
ART = {
    'band': ('art/blue-band.jpeg', (1200, 2400)),
    'hero': ('art/hero-road-camp-wide.jpeg', (1200, 2400)),
    'cave': ('art/raid-panel-cave.jpeg', (800, 1600)),
}
# The painted seal v4 (tools/art/cut_seal.py cuts it out with real transparency) at 64 px and up: the footer, the
# medallion and the link preview (which reads the apple-touch icon). At 44 px and below (the nav, the favicons) the
# flat seal (owner 2026-11-10, tools/art/flat_seal.py), served as its SVG so it is crisp at every density.
SEAL_VERSION = 'v4'
SEAL_SOURCE = f'logo/lx-seal-{SEAL_VERSION}.png'
SEALS = {'seal-88': (SEAL_SOURCE, 176), 'seal-128': (f'logo/lx-seal-{SEAL_VERSION}-256.png', 256)}
# seal-400.png is the Discord webhook avatar and footer icon (api/internal/notify), the Go cards' seal
# (make sheet-assets copies it into api/internal/brand/art) and the apps' bundled seal.
FAVICONS = {'apple-touch-icon.png': (SEAL_SOURCE, 180), 'seal-400.png': (f'logo/lx-seal-{SEAL_VERSION}-512.png', 400)}
# The flat seal's SVGs, byte for byte: the nav's (art/) and the tab's (favicon.svg, the tiny bold seal at 16 px).
FLAT_SVGS = {'art/seal-flat-small.svg': SMALL_SVG, 'art/seal-flat-tiny-bold.svg': TINY_SVG, 'favicon.svg': TINY_SVG}
# favicon.ico, for browsers without SVG favicons: a render at each size (the tiny bold seal at 16, the small at 32
# and 48), never a downscale of one size.
FAVICON_ICO_PX = (16, 32, 48)
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


def compress_png(path: Path) -> None:
    """pngquant (quality 90 to 100, kept only if smaller) then lossless oxipng; both on PATH (brew install)."""
    subprocess.run(['pngquant', '--quality=90-100', '--skip-if-larger', '--force', '--strip',
                    '--output', str(path), str(path)], check=False)
    subprocess.run(['oxipng', '-o', '4', '--strip', 'safe', '-q', str(path)], check=True)


def build_seals() -> None:
    for name, (rel, px) in SEALS.items():
        Image.open(source(rel)).convert('RGBA').resize((px, px), Image.LANCZOS).save(OUT / f'{name}.webp', 'WEBP', quality=88, method=6)
    for name, (rel, px) in FAVICONS.items():
        path = REPO / 'web/public' / name
        Image.open(source(rel)).convert('RGBA').resize((px, px), Image.LANCZOS).save(path)
        compress_png(path)
    for name, svg in FLAT_SVGS.items():
        shutil.copyfile(svg, REPO / 'web/public' / name)
    sized = [flat_seal(px) for px in FAVICON_ICO_PX]
    largest = sized[-1]
    largest.save(REPO / 'web/public/favicon.ico', format='ICO', sizes=[im.size for im in sized],
                 append_images=sized[:-1])


def build_paperdoll() -> None:
    for scale in ('1x', '2x'):
        shutil.copyfile(DIRECTIONS / f'paperdoll/frame@{scale}.webp', OUT / f'paperdoll-frame@{scale}.webp')
    slots = json.loads((DIRECTIONS / 'paperdoll/slots.json').read_text())
    SLOTS_OUT.write_text(json.dumps(slots, indent=1) + '\n')


def build_companion() -> None:
    COMPANION_OUT.mkdir(parents=True, exist_ok=True)
    for name in COMPANION_ART:
        shutil.copyfile(OUT / name, COMPANION_OUT / name)


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    build_art()
    build_seals()
    build_paperdoll()
    build_companion()
    for path in sorted(OUT.iterdir()):
        print(f'{path.relative_to(REPO)}  {path.stat().st_size // 1024} KB')


if __name__ == '__main__':
    main()
