"""Subset Barlow Condensed 800 to exactly the glyphs the homepage hero headline can show.

The homepage's LCP element is the signed-out hero `<h1>` in src/pages/index.astro -- see
that file and web-home-lcp/web-home-chunks' lane reports for the measurement. The heading
face (Barlow Condensed, polish spec 2026-10-28 section 4.4, direction B) is loaded with
`font-display: swap` over a size-adjusted fallback (styles/fonts.css), so there is no layout
shift, but Chrome does not count text painted in the fallback face as an LCP candidate: the
LCP would land when the real face finishes downloading and paints.

The fix is not to bundle differently -- it is to make the hero's own glyphs available at
first paint. This script reads the exact strings the hero headline can show (currently just
`homeHeroCopy.headline` in home-landing-copy.ts; home-panel-copy.ts is also scanned for any
`*headline*`-named field so a future signed-in variant is picked up automatically), takes
their capitals (the hero is set `text-transform: uppercase`, so the capitals are the only
glyphs it paints; carrying the lowercase too would more than double the bytes inlined ahead
of the LCP), subsets the committed Barlow Condensed 800 subset (src/assets/fonts, itself cut from
@fontsource/barlow-condensed by the command in fonts/README.md) to only those characters
with fonttools, and writes a tiny inline `@font-face` for a `Barlow Condensed Hero` family
that `index.astro` declares in its own `<head>` (not Base.astro -- no other page should pay
for it) with `font-display: block`. The hero's metrics are untouched; only the font-family
list gains `'Barlow Condensed Hero'` ahead of the heading stack, so a glyph this subset is
missing (there should be none) still falls back to the same face, never a visual defect --
just back to swap timing for that one glyph.

Usage:
    python3 scripts/hero-font.py            # regenerate src/styles/hero-font.css
    python3 scripts/hero-font.py --check    # fail if the copy needs a glyph the
                                             # committed subset does not have

Run from web/. Requires fontTools and brotli (`python3 -c "import fontTools, brotli"`).
"""

from __future__ import annotations

import base64
import pathlib
import re
import subprocess
import sys
import tempfile

WEB_ROOT = pathlib.Path(__file__).resolve().parent.parent
COPY_FILES = [
    WEB_ROOT / "src/lib/home-landing-copy.ts",
    WEB_ROOT / "src/lib/home-panel-copy.ts",
]
# The only field actually painted in the hero's heading face today.
HERO_FIELD_PATTERN = re.compile(r"headline", re.IGNORECASE)
# Group 1 is the key, group 2 the quote character, group 3 the literal's body.
KEY_STRING_PATTERN = re.compile(r"""(\w+)\s*:\s*(['"])((?:\\.|(?!\2).)*)\2""")
OUTPUT_CSS = WEB_ROOT / "src/styles/hero-font.css"
SOURCE_FONT = WEB_ROOT / "src/assets/fonts/barlow-condensed-800-subset.woff2"
FONT_FAMILY = "Barlow Condensed Hero"
FONT_WEIGHT = 800

GENERATED_BANNER = (
    "/* GENERATED FILE -- do not hand-edit.\n"
    "   Run `python3 scripts/hero-font.py` from web/ to regenerate.\n"
    "   scripts/hero-font.py explains why this file exists and what it subsets. */"
)


def required_chars() -> set[str]:
    """Every character the hero headline can show, across every copy file and
    every headline-named field (signed-out today; a signed-in variant would
    be picked up the moment one is added, without touching this script)."""
    chars: set[str] = set()
    for path in COPY_FILES:
        text = path.read_text(encoding="utf-8")
        # Walk `key: 'value'` / `key: "value"` pairs; keep the ones whose key
        # matches HERO_FIELD_PATTERN. This is a regex scan, not a real parser,
        # but home-landing-copy.ts and home-panel-copy.ts are plain object
        # literals of string constants -- no template interpolation on the
        # fields this pattern can match.
        for match in KEY_STRING_PATTERN.finditer(text):
            key, _quote, value = match.group(1), match.group(2), match.group(3)
            if HERO_FIELD_PATTERN.search(key):
                # Capitals only: the hero is set text-transform: uppercase.
                chars.update(value.upper())
    if not chars:
        raise SystemExit(
            "hero-font.py found no headline-named field in "
            + ", ".join(str(p) for p in COPY_FILES)
            + " -- did homeHeroCopy.headline get renamed?"
        )
    return chars


def unicodes_arg(chars: set[str]) -> str:
    return ",".join(f"U+{ord(c):04X}" for c in sorted(chars))


def subset_font(chars: set[str]) -> bytes:
    if not SOURCE_FONT.exists():
        raise SystemExit(f"hero-font.py: source font not found at {SOURCE_FONT}")
    with tempfile.TemporaryDirectory() as tmp:
        out_path = pathlib.Path(tmp) / "heading-hero-subset.woff2"
        subprocess.run(
            [
                sys.executable,
                "-m",
                "fontTools.subset",
                str(SOURCE_FONT),
                f"--unicodes={unicodes_arg(chars)}",
                "--flavor=woff2",
                # Barlow Condensed 800 carries GSUB liga/locl (plus figure features the
                # hero's letters never reach) and GPOS kern (checked with fontTools.ttLib
                # against the source woff2). Dropping 'kern' would shift every glyph after
                # the first by a fraction of a pixel against the full face, so the three
                # that can touch the headline are kept explicitly.
                "--layout-features=kern,liga,locl",
                # Barlow Condensed carries TrueType hinting; the hero paints at 40 px and
                # above, where hinting changes nothing a reader can see, and the hinting
                # programs are 60% of these glyphs' bytes (3884 hinted against 1552), all of
                # it inlined ahead of the LCP.
                "--no-hinting",
                f"--output-file={out_path}",
            ],
            check=True,
        )
        return out_path.read_bytes()


def render_css(font_bytes: bytes, chars: set[str]) -> str:
    b64 = base64.b64encode(font_bytes).decode("ascii")
    char_list = "".join(sorted(chars))
    return (
        f"{GENERATED_BANNER}\n"
        f"/* Subset chars ({len(chars)}): {char_list!r} -- {len(font_bytes)} bytes woff2. */\n"
        # Machine-readable twin of the line above: src/styles/hero-font.check.test.ts reads
        # these code points so CI can verify the subset without Python or fontTools.
        f"/* subset-codepoints: {','.join(str(ord(c)) for c in sorted(chars))} */\n"
        "@font-face {\n"
        f"  font-family: '{FONT_FAMILY}';\n"
        "  font-style: normal;\n"
        f"  font-weight: {FONT_WEIGHT};\n"
        "  font-display: block;\n"
        f"  src: url(data:font/woff2;base64,{b64}) format('woff2');\n"
        "}\n"
    )


def generate() -> None:
    chars = required_chars()
    font_bytes = subset_font(chars)
    css = render_css(font_bytes, chars)
    OUTPUT_CSS.write_text(css, encoding="utf-8")
    print(f"hero-font.py: wrote {OUTPUT_CSS.relative_to(WEB_ROOT)} ({len(font_bytes)} bytes, {len(chars)} chars)")


def check() -> None:
    needed = required_chars()
    if not OUTPUT_CSS.exists():
        raise SystemExit(f"hero-font.py --check: {OUTPUT_CSS} does not exist -- run without --check first")
    css = OUTPUT_CSS.read_text(encoding="utf-8")
    match = re.search(r"base64,([A-Za-z0-9+/=]+)\)", css)
    if not match:
        raise SystemExit(f"hero-font.py --check: no base64 font payload found in {OUTPUT_CSS}")
    font_bytes = base64.b64decode(match.group(1))

    try:
        from fontTools.ttLib import TTFont
    except ImportError as exc:  # pragma: no cover - environment problem, not a code bug
        raise SystemExit("hero-font.py --check: fontTools is required (python3 -c 'import fontTools')") from exc

    with tempfile.TemporaryDirectory() as tmp:
        font_path = pathlib.Path(tmp) / "hero-subset.woff2"
        font_path.write_bytes(font_bytes)
        font = TTFont(font_path)
        cmap = font.getBestCmap()
        available = {chr(cp) for cp in cmap.keys()}

    missing = sorted(c for c in needed if c not in available)
    if missing:
        raise SystemExit(
            "hero-font.py --check: the hero headline now needs "
            f"{missing!r}, which src/styles/hero-font.css's subset does not have. "
            "Run `python3 scripts/hero-font.py` to regenerate it."
        )
    print(f"hero-font.py --check: OK ({len(needed)} chars, all present in the committed subset)")


def main() -> None:
    if "--check" in sys.argv[1:]:
        check()
    else:
        generate()


if __name__ == "__main__":
    main()
