"""Render the phone-app QR code the onboarding tour shows (design/specs/2026-11-08-addon-ux.md 3.1).

Run once, from anywhere:

    uv run --with segno --with pillow python addon/tools/make_app_qr.py

It writes addon/ForeverSixty/media/app-qr.tga, 128 x 128, black on white, uncompressed 32-bit (what
the client reads). The code encodes https://foreversixty.gg/app (the one stable address: the store
on a phone, a sign-up for testers until the apps ship) at error correction H with a 4-module quiet
zone. The file is committed; the addon never builds a QR at runtime. When the web lane's
design/assets/app-qr.png lands, this script can scale that file instead.
"""

from __future__ import annotations

import io
import sys
from pathlib import Path

import segno
from PIL import Image

URL = "https://foreversixty.gg/app"
SIZE = 128


def render() -> Image.Image:
    code = segno.make(URL, error="h", micro=False)
    buffer = io.BytesIO()
    code.save(buffer, kind="png", scale=1, border=4, dark="#000000", light="#ffffff")
    buffer.seek(0)
    modules = Image.open(buffer).convert("RGBA")
    # Whole modules: the largest integer scale that fits, centred on white.
    scale = SIZE // modules.width
    sharp = modules.resize((modules.width * scale, modules.height * scale), Image.NEAREST)
    canvas = Image.new("RGBA", (SIZE, SIZE), (255, 255, 255, 255))
    offset = (SIZE - sharp.width) // 2
    canvas.paste(sharp, (offset, offset))
    return canvas


def main() -> int:
    out = Path(__file__).resolve().parents[1] / "ForeverSixty" / "media" / "app-qr.tga"
    render().save(out, format="TGA", rle=False)
    print(out)
    return 0


if __name__ == "__main__":
    sys.exit(main())
