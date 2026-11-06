import { describe, expect, it } from 'vitest';
import { REGION, sectionRegion } from './regions';

describe('parity regions', () => {
  it('slugs a section title the way the mock library does', () => {
    expect(sectionRegion('What to stack')).toBe('section-what-to-stack');
    expect(sectionRegion('  Runners-up · DPS against the pick ')).toBe(
      'section-runners-up-dps-against-the-pick',
    );
  });

  it('names every region once', () => {
    const names = Object.values(REGION);
    expect(new Set(names).size).toBe(names.length);
  });
});
