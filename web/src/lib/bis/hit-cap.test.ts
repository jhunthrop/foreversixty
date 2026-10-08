import { describe, expect, it } from 'vitest';
import { hitCapLine } from './hit-cap';

describe('hitCapLine', () => {
  it('names specials and white swings for a dual wielder', () => {
    const line = hitCapLine({ baseline: 3, specials: 6, white: 25 });
    expect(line?.text).toBe('Hit to cap: 6% for specials, 25% for white swings');
  });

  it('omits the white clause when the key is absent', () => {
    expect(hitCapLine({ baseline: 3, specials: 6 })?.text).toBe('Hit to cap: 6% for specials');
  });

  it('rounds to one place and keeps whole percents bare', () => {
    expect(hitCapLine({ baseline: 0, specials: 8.96, white: 24.04 })?.text).toBe(
      'Hit to cap: 9% for specials, 24% for white swings',
    );
    expect(hitCapLine({ baseline: 0, specials: 5.55 })?.text).toMatch(/^Hit to cap: 5\.[56]% /);
  });

  it('explains the cap and the rating units in its tooltip', () => {
    const title = hitCapLine({ baseline: 3, specials: 6 })?.title ?? '';
    expect(title).toContain('full weight until the cap');
    expect(title).toContain('10 hit rating and 14 crit rating are 1%');
  });

  it('is undefined when the band publishes nothing', () => {
    expect(hitCapLine(undefined)).toBeUndefined();
    expect(hitCapLine(null)).toBeUndefined();
  });

  describe('a caster', () => {
    it('names the spell cap and, when a talent adds hit, the school figure', () => {
      const line = hitCapLine({
        kind: 'spell',
        baseline: 0,
        spell: 16,
        school: { names: ['Fire', 'Frost'], hit: 5, to_cap: 11 },
      });
      expect(line?.text).toBe('Spell hit to cap: 16% (school 11%)');
    });

    it('omits the school clause when no talent adds hit', () => {
      expect(hitCapLine({ kind: 'spell', baseline: 6, spell: 10 })?.text).toBe('Spell hit to cap: 10%');
    });

    it('keeps a school already at the cap as 0', () => {
      const line = hitCapLine({
        kind: 'spell',
        baseline: 6,
        spell: 10,
        school: { names: ['Shadow'], hit: 16, to_cap: 0 },
      });
      expect(line?.text).toBe('Spell hit to cap: 10% (school 0%)');
    });

    it('explains the 17% miss and the rating units in its tooltip', () => {
      const title = hitCapLine({ kind: 'spell', baseline: 0, spell: 16 })?.title ?? '';
      expect(title).toContain('17%');
      expect(title).toContain('10 hit rating is 1%');
    });
  });
});
