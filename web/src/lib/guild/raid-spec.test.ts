// web/src/lib/guild/raid-spec.test.ts
// The spec catalogue as the raid page reads it: the class's specs in catalogue order and each
// spec's own icon.
import { describe, expect, it } from 'vitest';
import { specIconSrc, specsForClass } from './raid-spec';

describe('specsForClass', () => {
  it('lists the class in catalogue order, with the full names', () => {
    expect(specsForClass('druid').map((spec) => spec.name)).toEqual([
      'Balance',
      'Feral',
      'Feral Bear',
      'Restoration',
    ]);
    expect(specsForClass('hunter').map((spec) => spec.name)).toContain('Beast Mastery');
  });

  it('is empty for an unknown class', () => {
    expect(specsForClass('death-knight')).toEqual([]);
  });
});

describe('specIconSrc', () => {
  it('is the spec talent-tab icon from the active build', () => {
    const holy = specsForClass('paladin').find((spec) => spec.spec === 'paladin-holy');
    expect(holy).toBeDefined();
    expect(specIconSrc(holy!)).toMatch(/^\/data\/[\d.]+\/icons\/[a-z0-9_]+\.webp$/);
  });
});
