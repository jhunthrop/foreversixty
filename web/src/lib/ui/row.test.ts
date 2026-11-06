// web/src/lib/ui/row.test.ts
import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { describe, expect, it } from 'vitest';
import { ROW_BAND_MARK, ROW_BAND_OWN, ROW_BAND_REACH, ownRowBand, ownRowBandInside, rowBand } from './row';

const GLOBAL_CSS = readFileSync(join(import.meta.dirname, '..', '..', 'styles', 'global.css'), 'utf8');

/** A class string's side margin and padding in px (Tailwind steps are 4 px). */
function sides(classes: string): { margin: number; padding: number } {
  const margin = /(?:^| )-mx-(\d+)(?: |$)/.exec(classes);
  const padding = /(?:^| )px-(\d+)(?: |$)/.exec(classes);
  return {
    margin: margin === null ? 0 : -Number(margin[1]) * 4,
    padding: padding === null ? 0 : Number(padding[1]) * 4,
  };
}

describe('row bands (row-band sweep 2026-11-01)', () => {
  it('a lit row holds its content 12 px inside the band, at the list edge an unlit row keeps', () => {
    for (const pad of [0, 1, 2] as const) {
      const lit = sides(rowBand(ROW_BAND_OWN, pad));
      const unlit = sides(rowBand(null, pad));
      expect(lit.padding, `pad ${pad}`).toBe(12);
      expect(lit.margin + lit.padding, `pad ${pad}`).toBe(unlit.padding);
    }
  });

  it('an unlit row carries only the list padding, never a band', () => {
    expect(rowBand(null, 0)).toBe('');
    expect(rowBand(null, 1)).toBe('px-1');
    expect(rowBand(null, 2)).toBe('px-2');
    expect(ownRowBand(false, 1)).toBe('px-1');
  });

  it('the own row is the own band with its reach', () => {
    expect(ownRowBand(true, 0)).toBe(`${ROW_BAND_OWN} ${ROW_BAND_REACH[0]}`);
    expect(ownRowBand(true, 0).split(' ')).toEqual(expect.arrayContaining(['row-band', 'row-band-own']));
    expect(ROW_BAND_MARK.split(' ')).toEqual(['row-band', 'row-band-mark']);
  });

  it('an inside band never reaches past its list: lit and unlit rows hold the same 12 px inset', () => {
    expect(ownRowBandInside(false)).toBe('px-3');
    expect(ownRowBandInside(true)).toBe(`${ROW_BAND_OWN} px-3`);
    expect(ownRowBandInside(true)).not.toMatch(/-mx-/);
  });

  it('global.css rounds a lit band and steps back the rules beside a lit one', () => {
    // Rounded only while lit: an unlit row's hairline runs straight to its ends.
    expect(GLOBAL_CSS).toMatch(/\.row-band \{\s*transition:/);
    expect(GLOBAL_CSS).toMatch(
      /:is\(\.row-band-own, \.row-band-mark, \.row-band-target:target, \.row-band-hover:active\) \{\s*border-radius: var\(--radius-control\);/,
    );
    expect(GLOBAL_CSS).toMatch(/\.row-band-own \{\s*background-color: var\(--row-hover\);/);
    expect(GLOBAL_CSS).toContain(
      '.row-rule:has(+ :is(.row-band-own, .row-band-mark, .row-band-target:target',
    );
    // A universal subject (a bare :has(), `+ *`) is matched against every element on every change,
    // enough to move the home page after its first layout; the neighbours are keyed to .row-rule.
    expect(GLOBAL_CSS).not.toMatch(/^\s*:has\(/m);
    expect(GLOBAL_CSS).not.toMatch(/\+ \*/);
    expect(GLOBAL_CSS).toContain('.row-band-hover:hover + .row-rule');
  });
});
