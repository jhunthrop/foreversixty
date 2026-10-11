// web/src/lib/character-profile.test.ts
import { describe, expect, it } from 'vitest';
import { bisHref, newestSpec, specClassLabel } from './character-profile';
import type { CharacterFight } from './rankings/api';

const fight = (spec?: string): CharacterFight => ({
  encounter: 'Warden Kelthas',
  encounter_id: 9001,
  difficulty: 8,
  metric: 'hps',
  value: 1840,
  spec,
  fought_at: '2026-12-09T22:10:00Z',
  report_id: 'fixture2abcd',
  fight_index: 3,
  execution_score: null,
});

describe('newestSpec', () => {
  it('reads the newest fight’s spec, then the best rows', () => {
    expect(newestSpec({ history: [fight('Holy'), fight('Discipline')], best: [] })).toBe('Holy');
    expect(newestSpec({ history: [fight()], best: [fight(), fight('Shadow')] })).toBe('Shadow');
  });

  it('is null when no fight names one', () => {
    expect(newestSpec({ history: [], best: [] })).toBeNull();
    expect(newestSpec({ history: [fight()], best: [fight()] })).toBeNull();
  });
});

describe('specClassLabel', () => {
  it('reads "Holy Priest" with a spec and "Priest" without', () => {
    expect(specClassLabel('Priest', 'Holy')).toBe('Holy Priest');
    expect(specClassLabel('Priest', null)).toBe('Priest');
    expect(specClassLabel(undefined, 'Holy')).toBe('');
  });
});

describe('bisHref', () => {
  it('links the spec’s best-in-slot page when class and spec are known', () => {
    expect(bisHref('Priest', 'Holy')).toBe('/bis/priest/holy');
    expect(bisHref('Hunter', 'Beast Mastery')).toBe('/bis/hunter/beast-mastery');
    expect(bisHref('Priest', null)).toBeNull();
    expect(bisHref(undefined, 'Holy')).toBeNull();
  });
});
