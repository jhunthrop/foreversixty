"""Writes the native renderers' faces: the site's own subset fonts, unwrapped from woff2 to TTF.

Two targets read the same subsets (web/src/assets/fonts/*-subset.woff2, the round-5 faces):
- api/internal/brand/fonts: the Discord cards and link previews draw with golang.org/x/image,
  which reads TrueType but not woff2 (Open Sans 400/600/700 and Marcellus 400).
- apps/shared/fonts: the iOS and Android apps (Open Sans 400/600/700 and Archivo 600). The apps'
  Marcellus stays the unmodified upstream file already there: its OFL reserves the name
  "Marcellus", which a subset (a Modified Version) may not carry.
Each face is copied glyph for glyph, with each face's OFL licence beside it. The timestamps are
kept, so a rerun on the same subsets writes the same bytes.

Run: make brand-fonts   (needs fontTools and brotli: pip install fonttools brotli)
"""

from dataclasses import dataclass
from pathlib import Path

from fontTools.ttLib import TTFont

REPO = Path(__file__).resolve().parents[2]
SITE = REPO / 'web/src/assets/fonts'
OFL_BODY = Path(__file__).with_name('OFL-1.1.txt')


@dataclass(frozen=True)
class Target:
    """One output folder: its faces (the site's subset name to the TTF name) and licences."""

    out: Path
    faces: dict[str, str]
    licences: dict[str, str]


TARGETS = (
    Target(
        out=REPO / 'api/internal/brand/fonts',
        faces={
            'open-sans-400': 'OpenSans-Regular',
            'open-sans-600': 'OpenSans-SemiBold',
            'open-sans-700': 'OpenSans-Bold',
            'marcellus-400': 'Marcellus-Regular',
        },
        licences={'OFL-open-sans.txt': 'open-sans-400', 'OFL-marcellus.txt': 'marcellus-400'},
    ),
    # The TTF names are the faces' PostScript names (tokens.json `faces`): iOS registers a face by
    # it, and the Android resource is the file name lowercased (ArchivoSemiBold-SemiBold.ttf to
    # archivosemibold_semibold), which Tokens.kt derives from the same PostScript name.
    Target(
        out=REPO / 'apps/shared/fonts',
        faces={
            'open-sans-400': 'OpenSans-Regular',
            'open-sans-600': 'OpenSans-SemiBold',
            'open-sans-700': 'OpenSans-Bold',
            'archivo-600': 'ArchivoSemiBold-SemiBold',
        },
        licences={'OFL-OpenSans.txt': 'open-sans-400', 'OFL-Archivo.txt': 'archivo-600'},
    ),
)


def convert(out: Path, site_name: str, out_name: str) -> Path:
    font = TTFont(SITE / f'{site_name}-subset.woff2', recalcTimestamp=False)
    font.flavor = None
    path = out / f'{out_name}.ttf'
    font.save(path, reorderTables=True)
    return path


def licence(out: Path, out_name: str, site_name: str) -> Path:
    copyright_line = TTFont(SITE / f'{site_name}-subset.woff2')['name'].getDebugName(0)
    path = out / out_name
    path.write_text(f'{copyright_line}\n\n{OFL_BODY.read_text()}')
    return path


def build(target: Target) -> list[Path]:
    target.out.mkdir(parents=True, exist_ok=True)
    faces = [convert(target.out, site, name) for site, name in target.faces.items()]
    licences = [licence(target.out, name, site) for name, site in target.licences.items()]
    return faces + licences


def main() -> None:
    for target in TARGETS:
        for path in build(target):
            print(f'{path.relative_to(REPO)}  {path.stat().st_size // 1024} KB')


if __name__ == '__main__':
    main()
