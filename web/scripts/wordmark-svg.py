"""Draw the FOREVER SIXTY wordmark as one SVG path, so no page loads the Cinzel face for it.

The header and footer lockups set the wordmark in Cinzel 700 with 0.1em tracking. Loading a
whole web font for thirteen fixed capitals cost a request on the first-paint path of every
page (fonts lane, 2026-11-06). This script shapes the same string with HarfBuzz (`hb-shape`,
so Cinzel's kerning is applied exactly as a browser applies it), draws each glyph's outline
with fontTools, and writes src/components/wordmark-path.generated.ts: the path in em units
(1 em = 100 units, baseline at y = 0) plus the box a line of that text occupies (the advance
width including the trailing tracking a browser adds after the last letter, and Cinzel's hhea
ascent + descent as the height, the content box `line-height: normal` gave the old text).
Wordmark.astro scales it with `font-size`, so each lockup keeps its existing size rules.

    python3 scripts/wordmark-svg.py            # regenerate the path module
    python3 scripts/wordmark-svg.py --check    # fail if the committed module is stale

Needs fontTools (`pip install fonttools brotli`) and HarfBuzz's hb-shape on PATH
(`brew install harfbuzz`). Run from web/.
"""

from __future__ import annotations

import json
import subprocess
import sys
import tempfile
from pathlib import Path

from fontTools.pens.svgPathPen import SVGPathPen
from fontTools.pens.transformPen import TransformPen
from fontTools.ttLib import TTFont

WEB_ROOT = Path(__file__).resolve().parents[1]
SOURCE_FONT = WEB_ROOT / "node_modules/@fontsource/cinzel/files/cinzel-latin-700-normal.woff2"
OUTPUT = WEB_ROOT / "src/components/wordmark-path.generated.ts"
TEXT = "FOREVER SIXTY"
TRACKING_EM = 0.1
UNITS_PER_EM = 100
DECIMALS = 1


def shape(ttf_path: str) -> list[dict]:
    """HarfBuzz glyph ids, advances and offsets (font units) for TEXT, kerning applied."""
    out = subprocess.run(
        ["hb-shape", "--output-format=json", "--no-glyph-names", ttf_path, TEXT],
        check=True,
        capture_output=True,
        text=True,
    ).stdout
    return json.loads(out)


def fmt(value: float) -> str:
    text = f"{value:.{DECIMALS}f}".rstrip("0").rstrip(".")
    return "0" if text in ("-0", "") else text


def build() -> str:
    font = TTFont(SOURCE_FONT)
    upm = font["head"].unitsPerEm
    scale = UNITS_PER_EM / upm
    glyph_order = font.getGlyphOrder()
    glyph_set = font.getGlyphSet()
    with tempfile.NamedTemporaryFile(suffix=".ttf") as ttf:
        font.flavor = None
        font.save(ttf.name)
        glyphs = shape(ttf.name)

    tracking = TRACKING_EM * upm
    pen_x = 0.0
    commands: list[str] = []
    for glyph in glyphs:
        pen = SVGPathPen(glyph_set, ntos=fmt)
        origin_x = (pen_x + glyph["dx"]) * scale
        origin_y = -glyph["dy"] * scale
        glyph_set[glyph_order[glyph["g"]]].draw(TransformPen(pen, (scale, 0, 0, -scale, origin_x, origin_y)))
        commands.append(pen.getCommands())
        pen_x += glyph["ax"] + tracking

    hhea = font["hhea"]
    width = pen_x * scale
    ascent = hhea.ascent * scale
    height = (hhea.ascent - hhea.descent) * scale
    return (
        "// GENERATED FILE -- do not hand-edit. Run `python3 scripts/wordmark-svg.py` from web/.\n"
        "// The FOREVER SIXTY wordmark, Cinzel 700 at 0.1em tracking, as one path in em units\n"
        f"// ({UNITS_PER_EM} units = 1 em, baseline at y = 0). See scripts/wordmark-svg.py.\n"
        f"export const WORDMARK_TEXT = {json.dumps(TEXT.title())};\n"
        f"export const WORDMARK_UNITS_PER_EM = {UNITS_PER_EM};\n"
        f"export const WORDMARK_WIDTH = {fmt(width)};\n"
        f"export const WORDMARK_HEIGHT = {fmt(height)};\n"
        f"export const WORDMARK_ASCENT = {fmt(ascent)};\n"
        f"export const WORDMARK_PATH =\n  {json.dumps(''.join(commands))};\n"
    )


def main() -> None:
    generated = build()
    if "--check" in sys.argv[1:]:
        if not OUTPUT.exists() or OUTPUT.read_text() != generated:
            raise SystemExit("wordmark-svg.py --check: stale; run `python3 scripts/wordmark-svg.py`")
        print("wordmark-svg.py --check: OK")
        return
    OUTPUT.write_text(generated)
    print(f"wordmark-svg.py: wrote {OUTPUT.relative_to(WEB_ROOT)} ({len(generated)} bytes)")


if __name__ == "__main__":
    main()
