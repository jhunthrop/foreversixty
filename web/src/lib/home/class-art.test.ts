// web/src/lib/home/class-art.test.ts
import { describe, expect, it } from 'vitest';
import { classArtUrl } from './class-art';

describe('classArtUrl', () => {
  it('maps every class to one of its tree backgrounds for the given build', () => {
    expect(classArtUrl('druid', '1.60.1.69893')).toBe('/data/1.60.1.69893/trees/druidferalcombat.webp');
    expect(classArtUrl('warrior', '1.60.1.69893')).toBe('/data/1.60.1.69893/trees/warriorarms.webp');
  });

  it('tolerates the API casing and spacing of a class name', () => {
    expect(classArtUrl('Hunter ', '1.60.1.69893')).toBe('/data/1.60.1.69893/trees/huntermarksmanship.webp');
  });

  it('is undefined for a missing or unknown class', () => {
    expect(classArtUrl(undefined)).toBeUndefined();
    expect(classArtUrl('deathknight')).toBeUndefined();
  });

  it('defaults to the active build', () => {
    expect(classArtUrl('mage')).toMatch(/^\/data\/[0-9.]+\/trees\/magefire\.webp$/);
  });
});
