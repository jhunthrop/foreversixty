// web/src/lib/planner/estimate.test.ts
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import type { BisBand } from '../bis/types';
import type { SimCharacter } from '../sim/character';
import { bandGear, estimateCharacter } from './estimate';
import { indexTalents } from './rules';
import type { TalentFile } from './types';

const index = indexTalents(fixtureTalents as TalentFile);
const band = {
  slots: [
    { slot: 'head', item_id: 10 },
    { slot: 'main_hand', item_id: 20, enchant_id: 7 },
    { slot: 'neck', item_id: 0 },
    { slot: 'not_a_slot', item_id: 30 },
  ],
} as unknown as BisBand;
const base = {
  point_order: [1001, 1001, 1002, 2001],
  spec: 'x',
  talent_level: 13,
  level: 60,
  gear: {},
  gear_slots: [],
} as unknown as SimCharacter;

describe('bandGear', () => {
  it('keeps the band picks with their enchants and drops empty or unknown slots', () => {
    expect(bandGear(band)).toEqual({
      gear: { head: 10, main_hand: 20 },
      slots: [
        { slot: 'head', item_id: 10 },
        { slot: 'main_hand', item_id: 20, enchant: 7 },
      ],
    });
  });
});

describe('estimateCharacter', () => {
  it('sims the band gear at level 60 with every point for a visitor (P1)', () => {
    const out = estimateCharacter(base, index, band, null);
    expect(out.level).toBe(60);
    expect(out.point_order).toEqual(base.point_order);
    expect(out.gear_slots).toHaveLength(2);
  });

  it('sims a character at its own level with only the points it has earned (G7)', () => {
    const out = estimateCharacter(base, index, null, 12);
    expect(out.level).toBe(12);
    expect(out.point_order).toEqual([1001, 1001, 1002]);
    expect(out.talent_level).toBe(12);
    expect(out.gear_slots).toEqual([]);
  });
});
