// The raid plaque's facts and words (lib/home/plaque.ts, round 5 home).
import { describe, expect, it } from 'vitest';
import { countdownFigures, raidWhen, roleCounts, signupFigures, signupRole } from './plaque';
import type { NextRaidView } from '../guild/api';

const NOW = new Date('2026-11-06T18:00:00Z');
const COUNTS: NextRaidView['counts'] = {
  in: 18,
  late: 3,
  tanks: 3,
  healers: 6,
  dps: 12,
  melee: 7,
  ranged: 5,
};

describe('the raid plaque', () => {
  it('reads the one countdown as big figures, nothing once passed', () => {
    expect(countdownFigures(new Date('2026-12-10T02:00:00Z'), NOW)).toEqual([
      { value: '33', label: 'days' },
      { value: '8', label: 'hours' },
    ]);
    expect(countdownFigures(new Date('2026-11-06T18:20:30Z'), NOW)).toEqual([{ value: '21', label: 'min' }]);
    expect(countdownFigures(NOW, NOW)).toEqual([]);
  });

  it('reads signed up, late and the four roles', () => {
    expect(signupFigures(COUNTS)).toEqual([
      { value: '18', label: 'signed up' },
      { value: '3', label: 'late' },
    ]);
    expect(roleCounts(COUNTS).map((r) => `${r.n} ${r.word}`)).toEqual([
      '3 tanks',
      '6 healers',
      '7 melee',
      '5 ranged',
    ]);
    expect(roleCounts({ ...COUNTS, melee: undefined, ranged: undefined })).toHaveLength(2);
  });

  it('signs up as the saved role, else the spec, else a one-role class', () => {
    const own = { character_key: 'k', state: 'late' as const, role: 'tank' as const };
    expect(signupRole(own, { class: 'warrior', spec: 'Fury' })).toBe('tank');
    expect(signupRole(null, { class: 'warrior', spec: 'Fury' })).toBe('dps');
    expect(signupRole(null, { class: 'mage' })).toBe('dps');
    expect(signupRole(null, { class: 'druid' })).toBeNull();
  });

  it("reads the raid's start whole, on the viewer's own clock with no zone", () => {
    expect(raidWhen('2026-12-10T02:00:00Z', 'America/Chicago', 'en-US')).toBe('Wednesday 9 December · 8 pm');
    expect(raidWhen('2026-12-10T02:30:00Z', 'America/Chicago', 'en-US')).toBe(
      'Wednesday 9 December · 8:30 pm',
    );
    expect(raidWhen('2026-12-10T02:00:00Z', 'Europe/London', 'en-US')).toBe('Thursday 10 December · 2 am');
    expect(raidWhen('2026-12-10T02:00:00Z', 'Europe/London', 'en-GB')).toBe('Thursday 10 December · 02:00');
  });
});
