// Cinzel is the wordmark's face and nothing else's (polish spec 2026-10-28 section 4.4,
// direction B Forge): every heading, title and headline figure is set in --font-heading. This
// walks every source file that can style text and fails on any use of the display face --
// `var(--font-display)` in CSS or the `font-display` utility class -- outside the wordmark
// lockups, so a new heading cannot quietly bring Cinzel back. The `font-display:` descriptor
// inside an @font-face rule is a different thing (the swap policy) and is not a use.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const STYLED_EXTENSIONS = new Set(['.astro', '.svelte', '.ts', '.css', '.md', '.mdx']);

/** The wordmark lockups, the only places the display face may appear, and the token itself. */
const ALLOWED = new Map<string, number>([
  ['styles/tokens.css', 1], // the --font-display definition
  ['styles/nav.css', 1], // .site-wordmark span: the header lockup (Header.astro's wordmark)
  ['components/Footer.astro', 1], // the footer's seal and wordmark lockup
]);

/** `var(--font-display)`, or the `font-display` class not followed by a descriptor colon. */
const DISPLAY_FACE_USE = /var\(--font-display\)|(?<![\w-])font-display(?![\w-])(?!\s*:)/g;

function styledFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) return entry.name === 'data' ? [] : styledFiles(full);
    if (entry.name.endsWith('.test.ts')) return [];
    return STYLED_EXTENSIONS.has(path.extname(entry.name)) ? [full] : [];
  });
}

/** Counts of display-face uses per file, keyed relative to src/. */
function displayFaceUses(): Map<string, number> {
  const uses = new Map<string, number>();
  for (const file of styledFiles(SRC)) {
    const text = readFileSync(file, 'utf8')
      // The token's own definition line names it as a property, not a use.
      .replace(/--font-display\s*:/g, '');
    const count = [...text.matchAll(DISPLAY_FACE_USE)].length;
    if (count > 0) uses.set(path.relative(SRC, file).split(path.sep).join('/'), count);
  }
  return uses;
}

describe('the display face (Cinzel)', () => {
  it('is used by the wordmark lockups only; every heading uses --font-heading', () => {
    const uses = displayFaceUses();
    const strays = [...uses].filter(([file, count]) => (ALLOWED.get(file) ?? 0) < count);
    expect(strays, 'set these in the heading face (heading-* utilities or var(--font-heading))').toEqual([]);
  });

  it('is still used by every wordmark lockup it is allowed in', () => {
    const uses = displayFaceUses();
    const lockups = [...ALLOWED.keys()].filter((file) => file !== 'styles/tokens.css');
    expect(lockups.filter((file) => !uses.has(file))).toEqual([]);
  });
});
