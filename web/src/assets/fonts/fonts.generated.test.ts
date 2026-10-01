// web/src/assets/fonts/fonts.generated.test.ts
// Guards the five committed font subsets (Cinzel 700, Barlow 400/600/700, JetBrains Mono
// 500) against web/fonts/glyphs.txt, the explicit inventory they are generated from (see
// web/fonts/README.md for the exact `pyftsubset` command and web/src/styles/fonts.css for
// why self-authored @font-face rules replaced the @fontsource imports).
//
// Every inventory code point should be in every subset's cmap, with one documented
// exception: a handful of characters are used in copy but were never in the source
// @fontsource `latin` files to begin with (checked directly against those files, not just
// the subsets) -- pyftsubset cannot add a glyph the source face never had, and the site
// already rendered them through the fallback stack in tokens.css before this lane existed.
// Asserting the subset equals "inventory minus this exact list" catches both directions of
// regression: a real glyph quietly dropped during subsetting, and this exception list
// silently growing because a font update (or a bug in generation) dropped something it used
// to have.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Font } from 'fontkit';
import { openSync } from 'fontkit';
import { describe, expect, it } from 'vitest';

/** openSync's return type is `Font | FontCollection` since a .ttc can hold several faces;
 * every file here is a single-face .woff2, so it is always a Font. */
function openFont(filePath: string): Font {
  return openSync(filePath) as Font;
}

const here = path.dirname(fileURLToPath(import.meta.url));
const GLYPHS_PATH = path.join(here, '..', '..', '..', 'fonts', 'glyphs.txt');

const SUBSETS = [
  { label: 'Cinzel 700', file: 'cinzel-700-subset.woff2' },
  { label: 'Barlow 400', file: 'barlow-400-subset.woff2' },
  { label: 'Barlow 600', file: 'barlow-600-subset.woff2' },
  { label: 'Barlow 700', file: 'barlow-700-subset.woff2' },
  { label: 'JetBrains Mono 500', file: 'jetbrains-mono-500-subset.woff2' },
];

// Rightwards arrow, leftwards arrow, dagger, narrow no-break space: present in
// web/fonts/glyphs.txt (the last one is in the task's baseline punctuation list) but absent
// from all three typefaces' @fontsource `latin` source files (verified with fontTools
// against the un-subsetted files, not just these subsets) -- so no subset can contain them
// either. See glyphs.txt's own comment for the source locations of the three that are
// actually used in copy today.
const NOT_IN_ANY_SOURCE_FACE = new Set([0x2190, 0x2192, 0x2020, 0x202f]);

function inventoryCodePoints(): Set<number> {
  const raw = readFileSync(GLYPHS_PATH, 'utf8');
  const points = new Set<number>();
  for (const line of raw.split('\n')) {
    if (line.startsWith('#')) continue;
    for (const char of line) points.add(char.codePointAt(0) ?? 0);
  }
  points.delete(0x0a); // the newline split-and-rejoin artifact, not a glyph
  return points;
}

describe('font subset glyph coverage', () => {
  const inventory = inventoryCodePoints();
  const expected = [...inventory].filter((point) => !NOT_IN_ANY_SOURCE_FACE.has(point));

  it('has a non-trivial inventory to check against', () => {
    expect(inventory.size).toBeGreaterThan(100);
  });

  for (const { label, file } of SUBSETS) {
    it(`${label} subset's cmap covers every glyphs.txt code point it can`, () => {
      const font = openFont(path.join(here, file));
      const missing = expected.filter((point) => !font.hasGlyphForCodePoint(point));
      expect(
        missing.map((point) => `U+${point.toString(16).toUpperCase().padStart(4, '0')}`),
        `regenerate with the pyftsubset command in web/fonts/README.md`,
      ).toEqual([]);
    });

    it(`${label} subset does not silently carry glyphs outside the inventory`, () => {
      const font = openFont(path.join(here, file));
      const charSet: number[] = font.characterSet ?? [];
      // 0xFFFF is never a real character: it is format-4 cmap's required terminator
      // segment, which some readers (fontkit included) surface as if it were a mapped code
      // point even though it always resolves to glyph 0 (.notdef). Filtering by "does this
      // code point actually resolve to a real glyph" excludes that parsing artifact without
      // excusing an actual extra glyph, which would resolve to a non-zero glyph id.
      const extra = charSet.filter(
        (point) => !inventory.has(point) && font.glyphForCodePoint(point).id !== 0,
      );
      expect(
        extra.map((point) => `U+${point.toString(16).toUpperCase().padStart(4, '0')}`),
        'the subset has a glyph glyphs.txt does not list -- regenerate from the current glyphs.txt',
      ).toEqual([]);
    });
  }
});
