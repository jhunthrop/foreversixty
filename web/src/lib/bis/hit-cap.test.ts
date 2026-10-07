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
});
