// web/src/lib/report/death-health.test.ts
import { describe, expect, it } from 'vitest';
import { believableHealth, healthPct, type HealthEvent } from './death-health';

const hit = (hp?: number, max = 32540): HealthEvent => ({ kind: 'damage', hp_after: hp, max_hp: max });
/** The readings as whole percentages, so a test can name what a bar would draw. */
const rounded = (pcts: (number | null)[]): (number | null)[] =>
  pcts.map((pct) => (pct === null ? null : Math.round(pct)));
const heal = (hp: number, max = 32540): HealthEvent => ({ kind: 'heal', hp_after: hp, max_hp: max });
/** A swing a shield ate whole: the engine writes no reading on it at all. */
const absorbed = (): HealthEvent => ({ kind: 'damage' });

describe('healthPct', () => {
  it('is null without a full bar to read against', () => {
    expect(healthPct({ hp_after: 100 })).toBeNull();
    expect(healthPct({ hp_after: 100, max_hp: 0 })).toBeNull();
  });

  it('clamps to the bar', () => {
    expect(healthPct({ hp_after: 50, max_hp: 100 })).toBe(50);
    expect(healthPct({ hp_after: -20, max_hp: 100 })).toBe(0);
    expect(healthPct({ hp_after: 120, max_hp: 100 })).toBe(100);
  });
});

describe('believableHealth', () => {
  it('keeps a bar that only ever falls', () => {
    expect(rounded(believableHealth([hit(30000), hit(20000), hit(10000)]))).toEqual([92, 61, 31]);
  });

  it('drops a reading that rises on a damage row with no heal before it', () => {
    // Hanabanana's 1:18 death on the sample log, the shape of it: the absorbed swing
    // carries no reading, and the swing after it reads 73% off a client that thinks she
    // is full. The real ceiling in the span is the 47% two heals earlier.
    const pcts = believableHealth([
      hit(14894),
      heal(15405),
      hit(9818),
      hit(8403),
      heal(9551),
      absorbed(),
      hit(23860),
      hit(5072),
    ]);
    expect(pcts[6]).toBeNull();
    expect(pcts[5]).toBeNull();
    const peak = Math.max(...pcts.filter((pct): pct is number => pct !== null));
    expect(Math.round(peak)).toBe(47);
    // The rows after the fiction are believed again: they are below the last real reading.
    expect(Math.round(pcts[7] ?? -1)).toBe(16);
  });

  it('believes a rise a heal explains, and reads the rows after it against that', () => {
    expect(rounded(believableHealth([hit(3254), heal(29286), hit(26032)]))).toEqual([10, 90, 80]);
  });

  it('believes the first reading it is given, however high', () => {
    expect(rounded(believableHealth([hit(27418)]))).toEqual([84]);
  });

  it('does not read a hit that took nothing off as a rise', () => {
    expect(rounded(believableHealth([hit(10000), hit(10000)]))).toEqual([31, 31]);
  });

  it('leaves a row with no reading out without disturbing the ones around it', () => {
    const pcts = believableHealth([hit(20000), absorbed(), hit(15000)]);
    expect(pcts[1]).toBeNull();
    expect(Math.round(pcts[2] ?? -1)).toBe(46);
  });
});
