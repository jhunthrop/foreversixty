"""Renders the Forever Sixty seal into the companion's icons.

    python3 companion/tools/icons/render.py

Reads design/assets/logo/lx-seal-v4*.png (the painted seal v4, 2026-11-09) for the app icons, and the flat seal
(owner 2026-11-10: design/assets/brand/seal-flat through tools/art/flat_seal.py) for the tray and the window's
header, where the seal is 44 px or smaller and the painted one turns to mud. It writes only what something uses:
    internal/icon/assets/seal-tray-{off,on,warn,bad}.ico    Windows tray, embedded in the Windows binary (flat)
    internal/icon/assets/seal-tray-{off,on,warn,bad}.png    macOS menu bar and Linux tray, embedded there (flat)
    internal/icon/assets/seal.ico     the Windows exe's icon resource (go-winres): window, taskbar, Explorer (painted)
    internal/icon/assets/seal.icns    the macOS bundle's icon (tools/package-mac-companion.sh): Dock, Finder (painted)
    ui/seal.svg                       the window's header seal at 32 px: the site's small flat seal, byte for byte
seal.ico and seal.icns are never embedded in the binary; each platform's go:embed (internal/icon) carries its
own tray images alone.

Every PNG is palette-quantized (pngquant at quality 90 to 100, dithered) and then losslessly recompressed
(oxipng), keeping the smaller of that and the plain encode; at 16 to 64 px the 256-colour palette holds every
colour the seal has. Needs pngquant and oxipng on PATH (brew install pngquant oxipng).

The seal is never redrawn. Each tray size is the flat seal rendered at that size (the tiny bold seal at 24 px and
below, the small seal with its swash above). For the painted seal, where the owner's set has a size (16, 24, 32, 40, 48, 64, 128, 256, 512, 1024)
that file is used as it is; any other size is a Lanczos downscale of the 1768 px master. Windows and Linux
show the seal full bleed, as every round icon there is. macOS sets an app icon on Apple's grid: the body
spans 824 of 1024 px over a soft shadow, so the seal stands the size of its Dock neighbours.

The tone dots take the window's tone colours from companion/ui/tokens.css (generated from
design/tokens/tokens.json; styles.css maps --tone-on, --tone-warn and --tone-bad to the same three), so
the tray and the Status headline never disagree. The build is deterministic: run twice, the outputs match.
"""
import io
import re
import shutil
import struct
import subprocess
import sys
import tempfile
from pathlib import Path

from PIL import Image, ImageDraw, ImageFilter

COMPANION = Path(__file__).resolve().parents[2]
REPO = COMPANION.parent
sys.path.insert(0, str(REPO / 'tools/art'))
from flat_seal import SMALL_SVG, flat_seal  # noqa: E402  (the flat seal at 44 px and below: one source with the site)
LOGO = REPO / 'design/assets/logo'
SEAL_VERSION = 'v4'
MASTER = LOGO / f'lx-seal-{SEAL_VERSION}.png'
OUT = COMPANION / 'internal/icon/assets'
TOKENS_CSS = COMPANION / 'ui/tokens.css'

MAC_BODY = 824 / 1024  # Apple's macOS icon grid
# Windows draws the exe's icon at 16 to 64 px a scale step (small and large icons, 100 to 200 %) and takes 256 for
# Explorer's large views, scaling it down for 96 and 128.
ICO_SIZES = (16, 20, 24, 32, 40, 48, 64, 256)
TRAY_ICO_SIZES = (16, 20, 24, 32, 40, 48)  # Windows' tray at 100, 125, 150, 200, 250 and 300 %
# The tray's 16 px entry (100 %) draws the LX with a heavier outline: the plain one goes soft at that size.
HEAVY_LX_PX = 16
TRAY_PNG_SIZE = 32  # macOS draws the menu-bar image at 16 pt (getlantern/systray): this is its 2x; Linux takes it too
DOT_RADIUS = 0.12  # of the icon's side
TONES = {'on': '--color-kill', 'warn': '--color-gold-hover', 'bad': '--color-wipe'}
RING = '--color-raised'  # the dot's ring, so it reads on a light taskbar and a dark one
# One entry per pixel size macOS draws, 16 to 512: AppKit picks a representation by pixels, so 16 pt at 2x takes
# the 32 px entry and 256 pt at 2x the 512 px one, and the @2x types (ic11 to ic14) would only repeat the bytes.
# No 1024: nothing short of a 512 pt Finder icon on a Retina screen asks for it, and the Dock tops out at 256 pt.
ICNS_TYPES = (('icp4', 16), ('icp5', 32), ('icp6', 64), ('ic07', 128), ('ic08', 256), ('ic09', 512))

_cache: dict[int, Image.Image] = {}


def token(name: str) -> str:
    match = re.search(rf'^\s*{re.escape(name)}:\s*(#[0-9a-f]{{6}});', TOKENS_CSS.read_text(), re.M)
    if not match:
        raise ValueError(f'{TOKENS_CSS} has no hex {name}; run node design/tokens/build.mjs')
    return match.group(1)


def seal(size: int) -> Image.Image:
    """The seal at size px, full bleed: the owner's render at that size, else a downscale of the master."""
    if size not in _cache:
        own = LOGO / f'lx-seal-{SEAL_VERSION}-{size}.png'
        im = Image.open(own if own.exists() else MASTER).convert('RGBA')
        if im.size != (size, size):
            im = im.resize((size, size), Image.LANCZOS)
        _cache[size] = im
    return _cache[size]


def mac_seal(size: int) -> Image.Image:
    """The seal on Apple's grid: the body centred at 824/1024, a drop shadow under it from 64 px up."""
    body = round(size * MAC_BODY)
    inset = (size - body) // 2
    shape = Image.open(MASTER).convert('RGBA').resize((body, body), Image.LANCZOS)
    canvas = Image.new('RGBA', (size, size), (0, 0, 0, 0))
    if size >= 64:
        alpha = Image.new('L', (size, size), 0)
        alpha.paste(shape.getchannel('A').point(lambda a: a * 3 // 10), (inset, inset + round(size * 0.012)))
        alpha = alpha.filter(ImageFilter.GaussianBlur(size * 0.012))
        canvas = Image.merge('RGBA', (*Image.new('RGB', (size, size)).split(), alpha))
    canvas.alpha_composite(shape, (inset, inset))
    return canvas


def tray_seal(size: int) -> Image.Image:
    """The flat seal for the tray at size px, its LX heavier at 16 px."""
    return flat_seal(size, heavy=size == HEAVY_LX_PX)


def with_dot(size: int, colour: str, ring: str) -> Image.Image:
    """The seal with a status dot at its lower right (companion premium spec, section 10.3), drawn at 4x."""
    scale = 4
    big = size * scale
    # A corner badge: about a quarter of the icon across with a ring of at least 3/4 px, its ring touching the
    # canvas edge, so it sits over the seal's rim at the bottom right and leaves the X clear.
    r, edge = size * DOT_RADIUS * scale, max(0.75, size / 24) * scale
    c = big - r - edge
    layer = Image.new('RGBA', (big, big), (0, 0, 0, 0))
    draw = ImageDraw.Draw(layer)
    draw.ellipse((c - r - edge, c - r - edge, c + r + edge, c + r + edge), fill=ring)
    draw.ellipse((c - r, c - r, c + r, c + r), fill=colour)
    out = tray_seal(size)
    out.alpha_composite(layer.reduce(scale))  # a box average: no Lanczos ringing on the flat dot
    return out


PNGQUANT = ('pngquant', '--quality=90-100', '--speed=1', '--strip', '--force', '--output')
OXIPNG = ('oxipng', '--opt', 'max', '--strip', 'all', '--quiet')


def png_bytes(im: Image.Image) -> bytes:
    """The smallest faithful PNG: quantized when pngquant meets its quality floor, then recompressed losslessly."""
    buf = io.BytesIO()
    im.save(buf, 'PNG', optimize=True)
    plain = buf.getvalue()
    with tempfile.TemporaryDirectory() as tmp:
        src, out = Path(tmp) / 'in.png', Path(tmp) / 'out.png'
        src.write_bytes(plain)
        quantized = subprocess.run([*PNGQUANT, str(out), str(src)], capture_output=True).returncode == 0
        if not quantized:  # 99: the palette would miss the quality floor; keep every colour
            out.write_bytes(plain)
        subprocess.run([*OXIPNG, str(out)], check=True)
        best = out.read_bytes()
    return best if len(best) < len(plain) else plain


def ico(images: list[tuple[int, Image.Image]]) -> bytes:
    """ICONDIR, one ICONDIRENTRY per image, then PNG blobs (Windows reads PNG entries since Vista)."""
    blobs = [png_bytes(im) for _, im in images]
    head = struct.pack('<HHH', 0, 1, len(images))
    offset = 6 + 16 * len(images)
    for (size, _), blob in zip(images, blobs):
        side = 0 if size == 256 else size
        head += struct.pack('<BBBBHHII', side, side, 0, 0, 1, 32, len(blob), offset)
        offset += len(blob)
    return head + b''.join(blobs)


def icns() -> bytes:
    elems = b''
    for kind, size in ICNS_TYPES:
        blob = png_bytes(mac_seal(size))
        elems += kind.encode('ascii') + struct.pack('>I', 8 + len(blob)) + blob
    return b'icns' + struct.pack('>I', 8 + len(elems)) + elems


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    for old in OUT.glob('seal-[0-9]*.png'):  # the per-size PNGs nothing reads any more
        old.unlink()
    (OUT / 'seal.ico').write_bytes(ico([(s, seal(s)) for s in ICO_SIZES]))
    (OUT / 'seal.icns').write_bytes(icns())
    ring = token(RING)
    (OUT / 'seal-tray-off.ico').write_bytes(ico([(s, tray_seal(s)) for s in TRAY_ICO_SIZES]))
    (OUT / 'seal-tray-off.png').write_bytes(png_bytes(tray_seal(TRAY_PNG_SIZE)))
    for tone, name in TONES.items():
        colour = token(name)
        (OUT / f'seal-tray-{tone}.ico').write_bytes(ico([(s, with_dot(s, colour, ring)) for s in TRAY_ICO_SIZES]))
        (OUT / f'seal-tray-{tone}.png').write_bytes(png_bytes(with_dot(TRAY_PNG_SIZE, colour, ring)))
    shutil.copyfile(SMALL_SVG, COMPANION / 'ui/seal.svg')
    print(f'wrote the seal icons to {OUT.relative_to(REPO)} and companion/ui/seal.svg')


if __name__ == '__main__':
    main()
