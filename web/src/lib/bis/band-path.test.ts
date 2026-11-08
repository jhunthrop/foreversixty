// web/src/lib/bis/band-path.test.ts
import { describe, expect, it } from 'vitest';
import { bandFromSlug, bandSlug, bandWords, bisBandPath, bisSpecPath } from './band-path';

describe('band paths', () => {
  it('names each band by its range, the top band as 60', () => {
    expect([20, 30, 40, 50, 60].map(bandSlug)).toEqual(['20-29', '30-39', '40-49', '50-59', '60']);
  });

  it('reads back only the slugs it writes', () => {
    for (const band of [20, 30, 40, 50, 60]) expect(bandFromSlug(bandSlug(band))).toBe(band);
    for (const slug of ['40-48', '41-50', '60-69', '70', 'abc', '']) expect(bandFromSlug(slug)).toBeNull();
  });

  it('keeps the first band on the spec’s own path, every other on its own', () => {
    expect(bisSpecPath('warrior', 'fury')).toBe('/bis/warrior/fury');
    expect(bisBandPath('warrior', 'fury', 20, 20)).toBe('/bis/warrior/fury');
    expect(bisBandPath('warrior', 'fury', 40, 20)).toBe('/bis/warrior/fury/40-49');
    expect(bisBandPath('warrior', 'fury', 60, 20)).toBe('/bis/warrior/fury/60');
  });

  it('says the band in words', () => {
    expect(bandWords(30)).toBe('Levels 30 to 39');
    expect(bandWords(60)).toBe('Level 60');
  });
});
