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
      countdown: null,
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
    expect(card.tag).toBe('');
    expect(card.countdown).toEqual({ days: 13, hours: 15, minutes: 0 });
    expect(card.bandsLine).toBe('Leveling best in slot, levels 20, 30, 40, 50, 60');
    expect(card.eyebrow).not.toMatch(/beta is live/i);
    expect(card.tag).not.toContain('cap 30');
  });

  it('counts down to the launch instant, 15:00 PST, in days, hours and minutes', () => {
    expect(heroCardFor(dates, at('2026-10-29T22:30:00Z')).countdown).toEqual({
      days: 6,
      hours: 0,
      minutes: 30,
    });
    expect(heroCardFor(dates, at('2026-11-04T10:00:00Z')).countdown).toEqual({
      days: 0,
      hours: 13,
      minutes: 0,
    });
  });

  it('after launch counts to the first raids, then says the level cap', () => {
    const launched = heroCardFor(dates, at('2026-11-05T10:00:00Z'));
    expect(launched).toMatchObject({ mode: 'launched', eyebrow: 'Launched Nov 4', tag: 'First raids Dec 9' });
    expect(launched.countdown).toEqual({ days: 33, hours: 14, minutes: 0 });
    const after = heroCardFor(dates, at('2026-12-10T10:00:00Z'));
    expect(after).toMatchObject({ mode: 'launched', tag: 'level cap 60', countdown: null });
  });

  it('keeps the beta wording when the data names no launch', () => {
    const beta = dates.filter((entry) => entry.role !== 'launch');
    expect(heroCardFor(beta, at('2026-11-20T00:00:00Z')).mode).toBe('beta');
  });
});
