import { describe, expect, it } from 'vitest';
import { bisCopy } from './copy';
import { expertiseCapLine, weightsForGainedStats } from './expertise-cap';
import type { BisStatWeight } from './types';

describe('expertiseCapLine', () => {
  it('words a damage dealer by dodge alone', () => {
    const line = expertiseCapLine({ baseline: 1, dodge: 5.5 });
    expect(line?.text).toBe('Expertise to cap: 5.5% for dodge');
    expect(line?.title).toBe(bisCopy.expertiseToCapTitle);
  });

  it('adds the parry clause for a tank', () => {
    expect(expertiseCapLine({ baseline: 2, dodge: 4.5, parry: 12 })?.text).toBe(
      'Expertise to cap: 4.5% for dodge, 12% for parry',
    );
  });

  it('is undefined when nothing the character wears or talents into gives Expertise', () => {
    expect(expertiseCapLine({ baseline: 0, dodge: 6.5, parry: 14 })).toBeUndefined();
  });

  it('prints a fully capped figure as 0', () => {
    expect(expertiseCapLine({ baseline: 9, dodge: 0, parry: 5 })?.text).toBe(
      'Expertise to cap: 0% for dodge, 5% for parry',
    );
  });

  it('is undefined when the band publishes none, as a file ranked before the key does', () => {
    expect(expertiseCapLine(undefined)).toBeUndefined();
    expect(expertiseCapLine(null)).toBeUndefined();
  });
});

describe('weightsForGainedStats', () => {
  const weights = [{ stat: 'defense' }, { stat: 'expertise' }] as BisStatWeight[];

  it('drops the Expertise weight when nothing gives the character Expertise', () => {
    expect(weightsForGainedStats(weights, { baseline: 0, dodge: 6.5, parry: 14 }).map((w) => w.stat)).toEqual(
      ['defense'],
    );
    expect(weightsForGainedStats(weights, null).map((w) => w.stat)).toEqual(['defense']);
  });

  it('keeps it when the character has some', () => {
    expect(weightsForGainedStats(weights, { baseline: 2, dodge: 4.5 })).toHaveLength(2);
  });
});
