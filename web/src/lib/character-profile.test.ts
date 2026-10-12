// web/src/lib/character-profile.test.ts
import { describe, expect, it } from 'vitest';
import {
  amountHeader,
  bestRepeatsHistory,
  bisHref,
  characterTableCopy,
  newestSpec,
  specClassLabel,
} from './character-profile';
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

describe('the fight tables (fix round 2 item 6)', () => {
  it('heads the amount column with the metric every row shares, else both', () => {
    expect(amountHeader([{ metric: 'hps' }, { metric: 'hps' }])).toBe('HPS');
    expect(amountHeader([{ metric: 'dps' }])).toBe('DPS');
    expect(amountHeader([{ metric: 'dps' }, { metric: 'hps' }])).toBe('DPS or HPS');
  });

  it('knows when "Best per encounter" only repeats every ranked fight', () => {
    const a = { report_id: 'r1', fight_index: 3 };
    const b = { report_id: 'r2', fight_index: 1 };
    const c = { report_id: 'r3', fight_index: 2 };
    expect(bestRepeatsHistory([a, b], [b, a])).toBe(true);
    expect(bestRepeatsHistory([a], [a, b])).toBe(false);
    expect(bestRepeatsHistory([a, c], [a, b])).toBe(false);
  });

  it('names the columns Boss, Spec, Parse %, Date, Execution', () => {
    expect(Object.values(characterTableCopy)).toEqual(['Boss', 'Spec', 'Parse %', 'Date', 'Execution']);
  });
});
