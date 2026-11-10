"""Writes the Go renderers' faces: the site's own subset fonts, unwrapped from woff2 to TTF.

The Discord cards and link previews (api/internal/brand) draw with golang.org/x/image, which
reads TrueType but not woff2, so this copies web/src/assets/fonts/*-subset.woff2 (the round-5
faces: Open Sans 400/600/700 and Marcellus 400) into api/internal/brand/fonts as TTF, glyph for
glyph, with each face's OFL licence beside it. The timestamps are kept, so a rerun on the same
subsets writes the same bytes.

Run: make brand-fonts   (needs fontTools and brotli: pip install fonttools brotli)
"""

from pathlib import Path

from fontTools.ttLib import TTFont

REPO = Path(__file__).resolve().parents[2]
SITE = REPO / 'web/src/assets/fonts'
OUT = REPO / 'api/internal/brand/fonts'
# Every face the Go renderers draw, by the site's subset name, and the licence it ships under.
FACES = {
    'open-sans-400': 'OpenSans-Regular',
    'open-sans-600': 'OpenSans-SemiBold',
    'open-sans-700': 'OpenSans-Bold',
    'marcellus-400': 'Marcellus-Regular',
}
LICENCES = {'OFL-open-sans.txt': 'open-sans-400', 'OFL-marcellus.txt': 'marcellus-400'}
OFL_BODY = Path(__file__).with_name('OFL-1.1.txt')


def convert(site_name: str, out_name: str) -> None:
    font = TTFont(SITE / f'{site_name}-subset.woff2', recalcTimestamp=False)
    font.flavor = None
    font.save(OUT / f'{out_name}.ttf', reorderTables=True)


def licence(out_name: str, site_name: str) -> None:
    copyright_line = TTFont(SITE / f'{site_name}-subset.woff2')['name'].getDebugName(0)
    (OUT / out_name).write_text(f'{copyright_line}\n\n{OFL_BODY.read_text()}')


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    for site_name, out_name in FACES.items():
        convert(site_name, out_name)
    for out_name, site_name in LICENCES.items():
        licence(out_name, site_name)
    for path in sorted(OUT.iterdir()):
        print(f'{path.relative_to(REPO)}  {path.stat().st_size // 1024} KB')


if __name__ == '__main__':
    main()
