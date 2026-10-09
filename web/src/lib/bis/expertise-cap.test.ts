import { describe, expect, it } from 'vitest';
import { bisCopy } from './copy';
import { expertiseCapLine } from './expertise-cap';

describe('expertiseCapLine', () => {
  it('words a damage dealer by dodge alone', () => {
    const line = expertiseCapLine({ baseline: 1, dodge: 5.5 });
    expect(line?.text).toBe('Expertise to cap: 5.5% for dodge');
    expect(line?.title).toBe(bisCopy.expertiseToCapTitle);
  });

  it('adds the parry clause for a tank', () => {
    expect(expertiseCapLine({ baseline: 0, dodge: 6.5, parry: 14 })?.text).toBe(
      'Expertise to cap: 6.5% for dodge, 14% for parry',
    );
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
