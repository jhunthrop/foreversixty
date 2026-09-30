// web/src/styles/hero-font.check.test.ts
// Guards src/styles/hero-font.css, the committed Cinzel 700 subset scripts/hero-font.py
// generates for the homepage hero (see that script's module docstring for the full why).
// If a headline-named field in the home copy modules ever needs a character the committed
// subset does not have, that glyph would silently fall back to the full Cinzel face for
// anyone who hits this page before Cinzel finishes downloading -- not a visual defect
// (same face), but it would quietly reopen the LCP gap this subset exists to close. This
// test fails loudly instead, naming the fix.
//
// Pure TypeScript on purpose: CI has no fontTools, so the test reads the code points the
// generator records in the stylesheet's header rather than shelling out to the script.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import * as landingCopy from '../lib/home-landing-copy';
import * as panelCopy from '../lib/home-panel-copy';

const here = path.dirname(fileURLToPath(import.meta.url));
const cssPath = path.join(here, 'hero-font.css');

/** Every string held under a key whose name contains "headline", at any depth. */
function headlineStrings(value: unknown, keyName: string, out: Set<string>): void {
  if (typeof value === 'string') {
    if (/headline/i.test(keyName)) out.add(value);
    return;
  }
  if (Array.isArray(value)) {
    for (const entry of value) headlineStrings(entry, keyName, out);
    return;
  }
  if (value !== null && typeof value === 'object') {
    for (const [key, entry] of Object.entries(value as Record<string, unknown>)) {
      headlineStrings(entry, /headline/i.test(keyName) ? keyName : key, out);
    }
  }
}

function requiredCodePoints(): Set<number> {
  const strings = new Set<string>();
  headlineStrings(landingCopy, '', strings);
  headlineStrings(panelCopy, '', strings);
  const points = new Set<number>();
  for (const text of strings) for (const char of text) points.add(char.codePointAt(0) ?? 0);
  return points;
}

function subsetCodePoints(): Set<number> {
  const css = readFileSync(cssPath, 'utf8');
  const match = /\/\* subset-codepoints: ([0-9,]+) \*\//.exec(css);
  if (match === null) throw new Error('hero-font.css has no subset-codepoints header');
  return new Set(match[1].split(',').map((n) => Number(n)));
}

describe('hero-font subset', () => {
  it('names at least one headline to subset for', () => {
    expect(requiredCodePoints().size).toBeGreaterThan(0);
  });

  it('has every glyph the hero headline currently needs', () => {
    const subset = subsetCodePoints();
    const missing = [...requiredCodePoints()].filter((point) => !subset.has(point));
    expect(
      missing.map((point) => String.fromCodePoint(point)),
      'regenerate with `python3 scripts/hero-font.py` from web/',
    ).toEqual([]);
  });
});
