// web/src/components/surface-guard.test.ts
// Only the ui/ primitives draw a surface (web layout spec 2026-11-05 §3A.3, owner decision 7). A
// card border, a raised ground, the panel radius, the card gradient or shadow, outside
// web/src/components/ui/, is a page inventing its own box: use Plate, Section, Row or Overlay.
// It scans every .astro and .svelte file under src/ and every styles/*.css file.
//
// It ships as a ratchet: surface-guard.allowlist.json names each file that still draws a surface
// and how many times. The test fails when a file not on the list draws one, when a listed file
// draws more than its allowance, and when a listed file draws fewer (lower or delete its entry:
// the list only ever shrinks, and ALLOWLIST_CEILING below only ever goes down). Each page lane
// deletes its files from the list as it adopts the primitives; the last one deletes the list.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import allowlist from './surface-guard.allowlist.json';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
/** The primitives themselves, the one place a surface may be drawn. */
const PRIMITIVES = 'components/ui/';
/** The number of allowlisted files when this ratchet landed (2026-11-05); lower it, never raise it. */
const ALLOWLIST_CEILING = 191;

/** Surface classes: the card, the panel radius, the line borders, the raised and hover grounds. */
const SURFACE_CLASS =
  /(?<![\w-])(?:card-surface(?:-quiet)?|rounded-panel|rounded-plate|(?:border|divide)(?:-[trblxy])?-line(?:-soft|-warm(?:-strong)?|-card)?|bg-raised|bg-card-top|bg-tile|bg-hover|bg-selected|shadow-card|shadow-overlay|border-\[var\(--color-line[\w-]*\))(?![\w-])/g;
/** A CSS background or border declaration that reads a surface token. */
const SURFACE_DECLARATION =
  /(?:^|[\s;{])(?:background|border)[\w-]*\s*:[^;}]*var\(--(?:color-(?:raised|card-top|line[\w-]*)|surface-card)\)/g;

function scannedFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory())
      return entry.name === 'data' || entry.name === 'fixtures' ? [] : scannedFiles(full);
    if (entry.name.includes('.test.')) return [];
    const rel = relative(full);
    if (rel.startsWith(PRIMITIVES)) return [];
    if (/\.(astro|svelte)$/.test(entry.name)) return [full];
    return rel.startsWith('styles/') && entry.name.endsWith('.css') ? [full] : [];
  });
}

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

function withoutComments(text: string): string {
  return text
    .replace(/\/\*[\s\S]*?\*\//g, '')
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/^\s*\/\/.*$/gm, '');
}

/** Surface draws per file, keyed relative to src/. */
export function surfaceUses(): Record<string, number> {
  const uses: Record<string, number> = {};
  for (const file of scannedFiles(SRC)) {
    const text = withoutComments(readFileSync(file, 'utf8'));
    const count = (text.match(SURFACE_CLASS)?.length ?? 0) + (text.match(SURFACE_DECLARATION)?.length ?? 0);
    if (count > 0) uses[relative(file)] = count;
  }
  return uses;
}

const allowed: Readonly<Record<string, number>> = allowlist.files;

describe('the surface guard', () => {
  const uses = surfaceUses();

  it('no file outside the primitives draws a surface it is not allowed', () => {
    const over = Object.entries(uses)
      .filter(([file, count]) => count > (allowed[file] ?? 0))
      .map(([file, count]) => `${file}: ${count} (allowed ${allowed[file] ?? 0})`);
    expect(over, 'draw it with a ui/ primitive (Plate, Section, Row, Overlay) instead').toEqual([]);
  });

  it('the allowlist only shrinks: a file that draws less lowers or deletes its entry', () => {
    const stale = Object.entries(allowed)
      .filter(([file, count]) => (uses[file] ?? 0) < count)
      .map(([file, count]) => `${file}: allowed ${count}, draws ${uses[file] ?? 0}`);
    expect(stale, 'lower these entries in surface-guard.allowlist.json').toEqual([]);
  });

  it('never grows past its ceiling', () => {
    expect(Object.keys(allowed).length).toBeLessThanOrEqual(ALLOWLIST_CEILING);
  });

  it('lets the primitives draw their surfaces', () => {
    expect(Object.keys(uses).some((file) => file.startsWith(PRIMITIVES))).toBe(false);
  });
});
