// web/src/lib/home/upgrades.test.ts
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { BisBand, BisSlot } from '../bis/types';
import type { Item } from '../planner/types';
import { upgradesFor, weightsRecordFor } from './upgrades';

function item(id: number, name: string, stats: Partial<Record<string, number>>): Item {
  return {
    id,
    name,
    icon: 'inv_misc_questionmark',
    slot: 'head',
    quality: 3,
    required_level: 1,
    item_level: 1,
    armor: 0,
    stats: stats as Item['stats'],
    set_id: null,
    unique: false,
  };
}

function pick(slot: string, itemId: number, itemName: string, score: number | undefined): BisSlot {
  return {
    slot,
    item_id: itemId,
    item_name: itemName,
    source: 'World drop',
    source_kind: 'world_drop',
    score,
    verified: true,
  };
}

function band(overrides: Partial<BisBand> = {}): BisBand {
  return {
    spec: 'hunter-marksmanship',
    band: 20,
    faction: 'horde',
    race: 'troll',
    talents: '0000000000000000-35300000000000000-000000000000000000',
    talent_points: 11,
    weights: [{ stat: 'agility', weight: 2 }],
    slots: [],
    set_dps: 75,
    no_source_count: 0,
    coverage: {},
    new_at_band: [],
    weights_run_seconds: 0,
    verify_run_seconds: 0,
    reference_dps_per_point: 0.1,
    ...overrides,
  };
}

describe('weightsRecordFor', () => {
  it('reshapes a band weight list into a stat -> weight record', () => {
    expect(
      weightsRecordFor(
        band({
          weights: [
            { stat: 'agility', weight: 2 },
            { stat: 'hit', weight: 0.5 },
          ],
        }),
      ),
    ).toEqual({ agility: 2, hit: 0.5 });
  });
});

describe('upgradesFor', () => {
  const PICK = pick('head', 1, 'Band Pick Hood', 10); // 5 agility * weight 2 = score 10

  it('flags a slot an upgrade when the worn item differs from the pick, with the DPS gain converted via reference_dps_per_point', () => {
    const items = new Map([
      [1, item(1, 'Band Pick Hood', { agility: 5 })],
      [2, item(2, 'Starter Cap', { agility: 0 })],
    ]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { head: 2 } },
    };
    const result = upgradesFor(character, band({ slots: [PICK] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]).toMatchObject({ slot: 'head', wornItemId: 2, wornItemName: 'Starter Cap' });
    // (10 - 0) * 0.1 reference_dps_per_point = 1
    expect(result.upgrades[0]!.gainDps).toBeCloseTo(1);
    expect(result.totalGainDps).toBeCloseTo(1);
    expect(result.alreadyBis).toHaveLength(0);
  });

  it('treats wearing the pick itself as already best in slot, not an upgrade', () => {
    const items = new Map([[1, item(1, 'Band Pick Hood', { agility: 5 })]]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { head: 1 } },
    };
    const result = upgradesFor(character, band({ slots: [PICK] }), items);
    expect(result.upgrades).toHaveLength(0);
    expect(result.alreadyBis).toEqual([{ slot: 'head', itemId: 1, itemName: 'Band Pick Hood' }]);
  });

  it('treats a listed alternative with dps_delta 0 as a tie, not an upgrade', () => {
    const pickWithTie: BisSlot = {
      ...PICK,
      alternatives: [
        { item_id: 3, item_name: 'Tied Alt', source: 'World drop', source_kind: 'world_drop', dps_delta: 0 },
      ],
    };
    const items = new Map([
      [1, item(1, 'Band Pick Hood', { agility: 5 })],
      [3, item(3, 'Tied Alt', { agility: 5 })],
    ]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { head: 3 } },
    };
    const result = upgradesFor(character, band({ slots: [pickWithTie] }), items);
    expect(result.upgrades).toHaveLength(0);
    expect(result.alreadyBis).toEqual([{ slot: 'head', itemId: 3, itemName: 'Tied Alt' }]);
  });

  it("counts an empty slot (nothing equipped) as an upgrade, with the pick's full score as the gain", () => {
    const items = new Map([[1, item(1, 'Band Pick Hood', { agility: 5 })]]);
    const character: Pick<MeCharacter, 'build'> = { build: { source: 'addon', captured_at: '', gear: {} } };
    const result = upgradesFor(character, band({ slots: [PICK] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]!.wornItemId).toBeUndefined();
    expect(result.upgrades[0]!.gainDps).toBeCloseTo(1); // (10 - 0) * 0.1
  });

  it("never fabricates a gain for a worn item id this build's item table does not carry", () => {
    const items = new Map([[1, item(1, 'Band Pick Hood', { agility: 5 })]]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { head: 999 } },
    };
    const result = upgradesFor(character, band({ slots: [PICK] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]).toMatchObject({ wornItemId: 999, wornUnknown: true, gainDps: null });
    expect(result.totalGainDps).toBe(0);
  });

  it('never fabricates a gain for a sim-decided pick that carries no score', () => {
    const simDecidedPick = pick('ranged', 5, 'Sim Decided Bow', undefined);
    const items = new Map([
      [5, item(5, 'Sim Decided Bow', {})],
      [6, item(6, 'Old Bow', {})],
    ]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { ranged: 6 } },
    };
    const result = upgradesFor(character, band({ slots: [simDecidedPick] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]!.gainDps).toBeNull();
  });

  it('skips a slot the band has no known source for entirely -- neither an upgrade nor already best in slot', () => {
    // The real pipeline writes an unsourced slot as `{ slot, verified: false }` alone, with
    // no `item_id` at all (`BisSlot`'s own doc comment) -- the exact shape `hasKnownSource`
    // exists to guard against.
    const emptySlotRow = { slot: 'trinket1', verified: false } as unknown as BisSlot;
    const character: Pick<MeCharacter, 'build'> = { build: { source: 'addon', captured_at: '', gear: {} } };
    const result = upgradesFor(character, band({ slots: [emptySlotRow] }), new Map());
    expect(result.upgrades).toHaveLength(0);
    expect(result.alreadyBis).toHaveLength(0);
  });

  it('sorts upgrades by gain descending, with an unknown gain sorting last', () => {
    const headPick = pick('head', 1, 'Head Pick', 10);
    const neckPick = pick('neck', 2, 'Neck Pick', 20);
    const unknownGainPick = pick('feet', 3, 'Feet Pick', undefined);
    const items = new Map([
      [1, item(1, 'Head Pick', {})],
      [2, item(2, 'Neck Pick', {})],
      [3, item(3, 'Feet Pick', {})],
    ]);
    const character: Pick<MeCharacter, 'build'> = { build: { source: 'addon', captured_at: '', gear: {} } };
    const result = upgradesFor(character, band({ slots: [headPick, neckPick, unknownGainPick] }), items);
    expect(result.upgrades.map((u) => u.slot)).toEqual(['neck', 'head', 'feet']);
  });
});
