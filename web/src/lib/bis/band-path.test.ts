// web/src/lib/bis/band-path.test.ts
import { describe, expect, it } from 'vitest';
import {
  bandFromHash,
  bandFromSlug,
  bandHash,
  bandSlug,
  bandWords,
  bisBandHref,
  bisBandPath,
  bisSpecPath,
  isFoldedBand,
  specPageBand,
} from './band-path';

describe('band paths', () => {
  it('names each band by its range, the top band as 60', () => {
    expect([20, 30, 40, 50, 60].map(bandSlug)).toEqual(['20-29', '30-39', '40-49', '50-59', '60']);
  });

  it('reads back only the slugs it writes', () => {
    for (const band of [20, 30, 40, 50, 60]) expect(bandFromSlug(bandSlug(band))).toBe(band);
    for (const slug of ['40-48', '41-50', '60-69', '70', 'abc', '']) expect(bandFromSlug(slug)).toBeNull();
  });

  it('keeps the first and the folded bands on the spec’s own path, level 60 on its own', () => {
    expect(bisSpecPath('warrior', 'fury')).toBe('/bis/warrior/fury');
    for (const band of [20, 30, 40, 50])
      expect(bisBandPath('warrior', 'fury', band, 20)).toBe('/bis/warrior/fury');
    expect(bisBandPath('warrior', 'fury', 60, 20)).toBe('/bis/warrior/fury/60');
  });

  it('folds every band below 60 (the spec page opens on 60)', () => {
    expect([20, 30, 40, 50, 60].filter((band) => isFoldedBand(band))).toEqual([20, 30, 40, 50]);
  });

  it('opens the spec page on 60, else on a short spec’s first band', () => {
    expect(specPageBand([20, 30, 40, 50, 60])).toBe(60);
    expect(specPageBand([20, 30])).toBe(20);
    expect(specPageBand([])).toBeUndefined();
  });

  it('links a folded band by its hash after the query, the rest by path', () => {
    expect(bisBandHref('warrior', 'fury', 20, 20)).toBe('/bis/warrior/fury#levels-20-29');
    expect(bisBandHref('warrior', 'fury', 40, 20)).toBe('/bis/warrior/fury#levels-40-49');
    expect(bisBandHref('warrior', 'fury', 40, 20, '?faction=horde')).toBe(
      '/bis/warrior/fury?faction=horde#levels-40-49',
    );
    expect(bisBandHref('warrior', 'fury', 60, 20, '?faction=horde')).toBe(
      '/bis/warrior/fury/60?faction=horde',
    );
  });

  it('reads back only the band hashes it writes', () => {
    for (const band of [20, 30, 40, 50]) expect(bandFromHash(bandHash(band))).toBe(band);
    for (const hash of ['#levels-60-69', '#levels-41-50', '#band-alliance-40', '#leveling', '']) {
      expect(bandFromHash(hash)).toBeNull();
    }
  });

  it('says the band in words', () => {
    expect(bandWords(30)).toBe('Levels 30 to 39');
    expect(bandWords(60)).toBe('Level 60');
  });
});
