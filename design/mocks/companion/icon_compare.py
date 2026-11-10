"""Before and after the icon shrink: the full-colour render against the shipped (quantized) file, at 100 % and 200 %.

    python3 design/mocks/companion/icon_compare.py <out.png>

Each pair is the same image twice, left the unquantized render (companion/tools/icons/render.py's drawing before
png_bytes), right the bytes that ship, on a dark and a light ground, then the same pair enlarged 2x with nearest-
neighbour so a palette step would show. The PSNR under each pair is measured over the dark ground.
"""
import io
import math
import struct
import sys
from pathlib import Path

from PIL import Image, ImageChops, ImageDraw, ImageStat

REPO = Path(__file__).resolve().parents[3]
sys.path.insert(0, str(REPO / 'companion/tools/icons'))
import render as R  # noqa: E402

ASSETS = R.OUT
DARK, LIGHT = (30, 30, 30, 255), (236, 236, 236, 255)


def entries_icns() -> dict[int, Image.Image]:
    b, p, out = (ASSETS / 'seal.icns').read_bytes(), 8, {}
    sizes = dict(R.ICNS_TYPES)
    while p < len(b):
        kind, n = b[p:p + 4].decode(), struct.unpack('>I', b[p + 4:p + 8])[0]
        out[sizes[kind]] = Image.open(io.BytesIO(b[p + 8:p + n])).convert('RGBA')
        p += n
    return out


def entries_ico(name: str) -> dict[int, Image.Image]:
    b = (ASSETS / name).read_bytes()
    out = {}
    for i in range(struct.unpack('<H', b[4:6])[0]):
        e = b[6 + 16 * i:22 + 16 * i]
        n, off = struct.unpack('<II', e[8:16])
        out[e[0] or 256] = Image.open(io.BytesIO(b[off:off + n])).convert('RGBA')
    return out


def on(im: Image.Image, bg) -> Image.Image:
    c = Image.new('RGBA', im.size, bg)
    c.alpha_composite(im)
    return c


def psnr(a: Image.Image, b: Image.Image) -> float:
    st = ImageStat.Stat(ImageChops.difference(on(a, DARK).convert('RGB'), on(b, DARK).convert('RGB')))
    mse = sum(v * v for v in st.rms) / 3
    return 10 * math.log10(255 ** 2 / mse) if mse else 99.0


def main(out: Path) -> None:
    mac, ico = entries_icns(), entries_ico('seal.ico')
    ring = R.token(R.RING)
    pairs = [(f'macOS {s}', R.mac_seal(s), mac[s]) for s in (32, 64, 128, 256, 512)]
    pairs += [(f'Windows {s}', R.seal(s), ico[s]) for s in (16, 24, 32, 48, 256)]
    pairs += [('Tray on 32', R.with_dot(32, R.token(R.TONES['on']), ring), Image.open(ASSETS / 'seal-tray-on.png').convert('RGBA'))]
    rows = []
    for label, before, after in pairs:
        s = before.width
        tiles = [on(before, DARK), on(after, DARK), on(before, LIGHT), on(after, LIGHT)]
        tiles += [t.resize((s * 2, s * 2), Image.NEAREST) for t in tiles[:2]]
        w = sum(t.width for t in tiles) + 16 * len(tiles) + 16
        h = max(t.height for t in tiles) + 40
        row = Image.new('RGB', (w, h), (24, 24, 24))
        x = 8
        for t in tiles:
            row.paste(t.convert('RGB'), (x, 8))
            x += t.width + 16
        ImageDraw.Draw(row).text((8, h - 26), f'{label}: before | after (dark), before | after (light), then 200 %   PSNR {psnr(before, after):.1f} dB', fill=(220, 220, 220))
        rows.append(row)
    sheet = Image.new('RGB', (max(r.width for r in rows), sum(r.height for r in rows)), (24, 24, 24))
    y = 0
    for r in rows:
        sheet.paste(r, (0, y))
        y += r.height
    sheet.save(out)
    print(out)


if __name__ == '__main__':
    main(Path(sys.argv[1]))
