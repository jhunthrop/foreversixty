// web/src/styles/font-families.test.ts
// The site's type rule (round 5, design/specs/2026-11-09-directions.md section 4): no condensed
// face anywhere, no mono in UI, and only the faces the site actually loads. Marcellus 400 sets
// every heading (--font-heading is --font-title) and the bar's name, Open Sans 400/600/700 sets
// body, labels and figures, Archivo 600 the navigation bar, and JetBrains Mono literal code only. This walks every source file that can style text and fails on a
// condensed family or stretch, on any use of the retired display face, and on a fonts.css
// family outside the allowed list, so a new heading cannot quietly bring a third family back.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const STYLED_EXTENSIONS = new Set(['.astro', '.svelte', '.ts', '.css', '.md', '.mdx']);

/** Every family fonts.css may declare: the loaded faces and their size-adjusted fallbacks. */
const ALLOWED_FAMILIES = [
  'Archivo',
  'Open Sans',
  'Open Sans Fallback',
  'JetBrains Mono',
  'Marcellus',
  'Marcellus Fallback',
  'Marcellus Fallback Times',
];

/** A font family, font shorthand, font-stretch or --font-* token declaration and its value. */
const FONT_DECLARATION = /(?:^|[\s{;"'`])(font-family|font-stretch|font|--font-[\w-]+)\s*:\s*([^;}\n]+)/g;
const CONDENSED_WORD = /\b(?:condensed|narrow|compressed)\b/i;
/** Tailwind's stretch utilities (`font-stretch-condensed`, `font-stretch-75%`, ...). */
const STRETCH_CLASS = /(?<![\w-])font-stretch-[\w%-]+/;
/** The retired display face: `var(--font-display)` or the `font-display` utility class. The
 *  `font-display:` descriptor inside an @font-face rule is the swap policy, not a use. */
const DISPLAY_FACE_USE = /var\(--font-display\b|--font-display\s*:|(?<![\w-])font-display(?![\w-])(?!\s*:)/;

function styledFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) return entry.name === 'data' ? [] : styledFiles(full);
    if (entry.name.endsWith('.test.ts')) return [];
    return STYLED_EXTENSIONS.has(path.extname(entry.name)) ? [full] : [];
  });
}

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

/** Files whose text matches `check`, keyed relative to src/. */
function offenders(check: (text: string) => boolean): string[] {
  return styledFiles(SRC)
    .filter((file) => check(readFileSync(file, 'utf8')))
    .map(relative);
}

function declaresCondensed(text: string): boolean {
  const condensedValue = [...text.matchAll(FONT_DECLARATION)].some(
    ([, property, value]) => property === 'font-stretch' || CONDENSED_WORD.test(value),
  );
  return condensedValue || STRETCH_CLASS.test(text);
}

function declaredFamilies(css: string): string[] {
  const families = [...css.matchAll(/@font-face\s*\{[^}]*?font-family:\s*['"]([^'"]+)['"]/g)].map(
    ([, f]) => f,
  );
  return [...new Set(families)].sort();
}

describe('the site type rule', () => {
  it('declares no condensed face or stretch in any style source', () => {
    expect(offenders(declaresCondensed), 'set it in --font-body instead').toEqual([]);
  });

  it('never uses the retired display face', () => {
    expect(offenders((text) => DISPLAY_FACE_USE.test(text))).toEqual([]);
  });

  it('declares only the allowed families in fonts.css', () => {
    const css = readFileSync(path.join(SRC, 'styles/fonts.css'), 'utf8');
    expect(declaredFamilies(css)).toEqual([...ALLOWED_FAMILIES].sort());
  });

  it('sets headings in the title face', () => {
    const tokens = readFileSync(path.join(SRC, 'styles/tokens.css'), 'utf8');
    expect(tokens).toMatch(/--font-heading:\s*var\(--font-title\);/);
  });

  it('draws the seal beside the name in both lockups', () => {
    for (const lockup of ['components/Header.astro', 'components/Footer.astro']) {
      const source = readFileSync(path.join(SRC, lockup), 'utf8');
      expect(source, lockup).toContain('<Seal ');
      expect(source, lockup).toContain('{SITE_NAME}');
    }
  });
});
