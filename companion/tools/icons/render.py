"""Renders the Forever Sixty seal v3 into the companion's icons.

    python3 companion/tools/icons/render.py

Reads design/assets/logo/lx-seal-v3*.png (the one logo since 2026-11-09) and writes:
    internal/icon/assets/seal-{16,24,32,48,64,128,256,512}.png  window icon (Linux), tray (macOS, Linux)
    internal/icon/assets/seal.ico                               Windows exe resource, window and taskbar
    internal/icon/assets/seal.icns                              macOS app bundle (Dock, Finder)
    internal/icon/assets/seal-tray-{on,warn,bad}.{png,ico}      the tray seal with the headline's tone dot
    ui/seal-64.png                                              the window's header seal (32 px at 2x)

The seal is never redrawn. Where the owner's set has a size (16, 24, 32, 40, 48, 64, 128, 256, 512, 1024)
that file is used as it is; any other size is a Lanczos downscale of the 1801 px master. Windows and Linux
show the seal full bleed, as every round icon there is. macOS sets an app icon on Apple's grid: the body
spans 824 of 1024 px over a soft shadow, so the seal stands the size of its Dock neighbours.

The tone dots take the window's tone colours from companion/ui/tokens.css (generated from
design/tokens/tokens.json; styles.css maps --tone-on, --tone-warn and --tone-bad to the same three), so
the tray and the Status headline never disagree. The build is deterministic: run twice, the outputs match.
"""
import io
import re
import struct
from pathlib import Path

from PIL import Image, ImageDraw, ImageFilter

COMPANION = Path(__file__).resolve().parents[2]
REPO = COMPANION.parent
LOGO = REPO / 'design/assets/logo'
MASTER = LOGO / 'lx-seal-v3.png'
OUT = COMPANION / 'internal/icon/assets'
TOKENS_CSS = COMPANION / 'ui/tokens.css'

MAC_BODY = 824 / 1024  # Apple's macOS icon grid
PNG_SIZES = (16, 24, 32, 48, 64, 128, 256, 512)
ICO_SIZES = (16, 20, 24, 32, 40, 48, 64, 128, 256)
TRAY_ICO_SIZES = (16, 20, 24, 32, 40, 48)  # Windows' tray at 100, 125, 150, 200, 250 and 300 %
TRAY_PNG_SIZE = 32  # macOS draws the menu-bar image at 16 pt (getlantern/systray): this is its 2x; Linux takes it too
TONES = {'on': '--color-kill', 'warn': '--color-gold-hover', 'bad': '--color-wipe'}
RING = '--color-raised'  # the dot's ring, so it reads on a light taskbar and a dark one
ICNS_TYPES = (('icp4', 16), ('icp5', 32), ('icp6', 64), ('ic07', 128), ('ic08', 256), ('ic09', 512),
              ('ic10', 1024), ('ic11', 32), ('ic12', 64), ('ic13', 256), ('ic14', 512))  # ic10..ic14: @2x

_cache: dict[int, Image.Image] = {}


def token(name: str) -> str:
    match = re.search(rf'^\s*{re.escape(name)}:\s*(#[0-9a-f]{{6}});', TOKENS_CSS.read_text(), re.M)
    if not match:
        raise ValueError(f'{TOKENS_CSS} has no hex {name}; run node design/tokens/build.mjs')
    return match.group(1)


def seal(size: int) -> Image.Image:
    """The seal at size px, full bleed: the owner's render at that size, else a downscale of the master."""
    if size not in _cache:
        own = LOGO / f'lx-seal-v3-{size}.png'
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


def with_dot(size: int, colour: str, ring: str) -> Image.Image:
    """The seal with a status dot at its lower right (companion premium spec, section 10.3), drawn at 4x."""
    scale = 4
    big = size * scale
    # A corner badge: about 30 % of the icon across with a ring of a pixel or more, its ring touching the
    # canvas edge, so it sits on the seal's rim and leaves the LX clear.
    r, edge = size * 0.15 * scale, max(1.0, size / 24) * scale
    c = big - r - edge
    layer = Image.new('RGBA', (big, big), (0, 0, 0, 0))
    draw = ImageDraw.Draw(layer)
    draw.ellipse((c - r - edge, c - r - edge, c + r + edge, c + r + edge), fill=ring)
    draw.ellipse((c - r, c - r, c + r, c + r), fill=colour)
    out = seal(size).copy()
    out.alpha_composite(layer.reduce(scale))  # a box average: no Lanczos ringing on the flat dot
    return out


def png_bytes(im: Image.Image) -> bytes:
    buf = io.BytesIO()
    im.save(buf, 'PNG', optimize=True)
    return buf.getvalue()


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
    for size in PNG_SIZES:
        (OUT / f'seal-{size}.png').write_bytes(png_bytes(seal(size)))
    (OUT / 'seal.ico').write_bytes(ico([(s, seal(s)) for s in ICO_SIZES]))
    (OUT / 'seal.icns').write_bytes(icns())
    ring = token(RING)
    for tone, name in TONES.items():
        colour = token(name)
        (OUT / f'seal-tray-{tone}.ico').write_bytes(ico([(s, with_dot(s, colour, ring)) for s in TRAY_ICO_SIZES]))
        (OUT / f'seal-tray-{tone}.png').write_bytes(png_bytes(with_dot(TRAY_PNG_SIZE, colour, ring)))
    (COMPANION / 'ui/seal-64.png').write_bytes(png_bytes(seal(64)))
    print(f'wrote the seal v3 icons to {OUT.relative_to(REPO)} and companion/ui/seal-64.png')


if __name__ == '__main__':
    main()
