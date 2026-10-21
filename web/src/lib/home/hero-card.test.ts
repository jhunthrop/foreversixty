import { describe, expect, it } from 'vitest';
import dates from '../../data/dates.json';
import { heroCardFor } from './hero-card';

const at = (iso: string): Date => new Date(iso);

describe('heroCardFor (W.9: the beta card follows the clock)', () => {
  it('says the beta is live, with its cap and its bands, while the beta runs', () => {
    expect(heroCardFor(dates, at('2026-10-20T12:00:00Z'))).toEqual({
      mode: 'beta',
      eyebrow: 'Beta is live',
      tag: 'level cap 30',
      bandsLine: 'Leveling best in slot, levels 10 to 30',
    });
  });

  it('stays live through the whole last day of the beta, Oct 21', () => {
    expect(heroCardFor(dates, at('2026-10-21T23:59:00Z')).mode).toBe('beta');
  });

  it('after Oct 21 reads "Launch Nov 4 · 15:00 PST" with the days left', () => {
    const card = heroCardFor(dates, at('2026-10-22T08:00:00Z'));
    expect(card.mode).toBe('launch');
    expect(card.eyebrow).toBe('Launch Nov 4 · 15:00 PST');
    expect(card.tag).toBe('in 13 days');
    expect(card.bandsLine).toBe('Leveling best in slot, levels 20, 30, 40, 50, 60');
    expect(card.eyebrow).not.toMatch(/beta is live/i);
    expect(card.tag).not.toContain('cap 30');
  });

  it('counts down to tomorrow and today, then says launched', () => {
    expect(heroCardFor(dates, at('2026-11-03T10:00:00Z')).tag).toBe('tomorrow');
    expect(heroCardFor(dates, at('2026-11-04T10:00:00Z')).tag).toBe('today');
    const launched = heroCardFor(dates, at('2026-11-05T10:00:00Z'));
    expect(launched).toMatchObject({ mode: 'launched', eyebrow: 'Launched Nov 4', tag: 'level cap 60' });
  });

  it('keeps the beta wording when the data names no launch', () => {
    const beta = dates.filter((entry) => entry.role !== 'launch');
    expect(heroCardFor(beta, at('2026-11-20T00:00:00Z')).mode).toBe('beta');
  });
});
