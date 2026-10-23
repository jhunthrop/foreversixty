// web/src/lib/guides/hit-cap-line.test.ts
import { describe, expect, it } from 'vitest';
import { guideHitCapLine } from './hit-cap-line';

describe('guideHitCapLine', () => {
  it('words a caster’s cap with the school distance: "Hit cap: 16% (13% from school)"', () => {
    const line = guideHitCapLine({
      kind: 'spell',
      baseline: 0,
      spell: 16,
      school: { names: ['Frost'], hit: 3, to_cap: 13 },
    });
    expect(line?.text).toBe('Hit cap: 16% (13% from school)');
  });

  it('leaves the school clause off when no talent adds hit', () => {
    expect(guideHitCapLine({ kind: 'spell', baseline: 0, spell: 16 })?.text).toBe('Hit cap: 16%');
  });

  it('words a melee cap with the white-swing distance for a dual wielder', () => {
    expect(guideHitCapLine({ baseline: 0, specials: 9, white: 28 })?.text).toBe(
      'Hit cap: 9% (28% for white swings)',
    );
    expect(guideHitCapLine({ baseline: 0, specials: 9 })?.text).toBe('Hit cap: 9%');
  });

  it('prints fractions to one place', () => {
    expect(guideHitCapLine({ baseline: 0, specials: 6.5 })?.text).toBe('Hit cap: 6.5%');
  });

  it('is undefined when the band publishes no cap', () => {
    expect(guideHitCapLine(null)).toBeUndefined();
    expect(guideHitCapLine(undefined)).toBeUndefined();
  });
});
