import { describe, expect, it } from 'vitest';
import { factionColorVar, racePillsFor } from './race-pills';

describe('racePillsFor', () => {
  it('lists every race legal for the class, marking the recommended ones', () => {
    const pills = racePillsFor('warrior', ['human', 'troll']);
    const slugs = pills.map((p) => p.slug);
    expect(slugs).toContain('human');
    expect(slugs).toContain('orc');
    expect(pills.find((p) => p.slug === 'human')?.recommended).toBe(true);
    expect(pills.find((p) => p.slug === 'orc')?.recommended).toBe(false);
  });

  it('returns an empty list for an unknown class slug rather than throwing', () => {
    expect(racePillsFor('not-a-class', [])).toEqual([]);
  });
});

describe('factionColorVar', () => {
  it('maps alliance and horde to their tokens, and anything else to plain text', () => {
    expect(factionColorVar('alliance')).toBe('var(--color-alliance)');
    expect(factionColorVar('horde')).toBe('var(--color-horde)');
    expect(factionColorVar('neutral')).toBe('var(--color-text)');
  });
});
