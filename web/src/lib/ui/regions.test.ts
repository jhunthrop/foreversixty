import { readFileSync } from 'node:fs';
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

  it('names the same regions as the mock library (design/mocks/ws_lib.py REGION)', () => {
    const lib = readFileSync(new URL('../../../../design/mocks/ws_lib.py', import.meta.url), 'utf8');
    const block = /^REGION = \{([\s\S]*?)\}/m.exec(lib)?.[1] ?? '';
    const mock = [...block.matchAll(/'[a-z_]+': '([a-z-]+)'/g)].map((m) => m[1]);
    expect(mock.toSorted()).toEqual(Object.values(REGION).toSorted());
  });
});
