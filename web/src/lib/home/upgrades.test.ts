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

  // --- rule (c): non-weapon slot, no alternative match -> scoreItem diff ------------------
  it('rule (c): flags a slot an upgrade when the worn item differs from the pick, with the DPS gain converted via reference_dps_per_point', () => {
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

  // --- rule (a): worn item is the pick itself -> not an upgrade --------------------------
  it('rule (a): treats wearing the pick itself as already best in slot, not an upgrade', () => {
    const items = new Map([[1, item(1, 'Band Pick Hood', { agility: 5 })]]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { head: 1 } },
    };
    const result = upgradesFor(character, band({ slots: [PICK] }), items);
    expect(result.upgrades).toHaveLength(0);
    expect(result.alreadyBis).toEqual([{ slot: 'head', itemId: 1, itemName: 'Band Pick Hood' }]);
  });

  it('rule (a): treats a listed alternative with dps_delta 0 as a tie, not an upgrade', () => {
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

  // --- rule (b): worn item is a listed (non-zero) alternative -> -dps_delta, never scoreItem ----
  it('rule (b): a worn item listed as an alternative uses the sim-verified -dps_delta, never scoreItem, even for a non-weapon slot', () => {
    const pickWithAlt: BisSlot = {
      ...PICK,
      alternatives: [
        {
          item_id: 4,
          item_name: 'Runner-up Hood',
          source: 'World drop',
          source_kind: 'world_drop',
          dps_delta: -0.5,
        },
      ],
    };
    // scoreItem would score this item identically to the tied alt in the test above (agility
    // 5 == the pick's own 5), which would wrongly read as a tie by stat weight alone -- the
    // real, sim-verified dps_delta (-0.5) must win instead.
    const items = new Map([
      [1, item(1, 'Band Pick Hood', { agility: 5 })],
      [4, item(4, 'Runner-up Hood', { agility: 5 })],
    ]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { head: 4 } },
    };
    const result = upgradesFor(character, band({ slots: [pickWithAlt] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]).toMatchObject({ wornItemId: 4, gainDps: 0.5, notSimChecked: false });
    expect(result.totalGainDps).toBeCloseTo(0.5);
  });

  it("rule (b): reproduces the real fixture's Ranger Bow gain from the band's own dps_delta, not a scoreItem diff", () => {
    // The exact regression fix round 1 found: data/builds/1.60.1.70009/bis/hunter-
    // marksmanship.json band 20 (horde) ranged.alternatives names Lil Timmy's Peashooter
    // (13136) at dps_delta -0.7432720066167775 against the Ranger Bow pick (score 178.445,
    // which bakes in the bow's own weapon DPS -- a figure scoreItem cannot reproduce from
    // `agility: 4` alone, which is why this must route through the alternative, not the
    // score-unit diff).
    const rangerBow: BisSlot = {
      slot: 'ranged',
      item_id: 3021,
      item_name: 'Ranger Bow',
      source: 'World drop',
      source_kind: 'world_drop',
      score: 178.44547117657908,
      verified: true,
      alternatives: [
        {
          item_id: 13136,
          item_name: "Lil Timmy's Peashooter",
          source: 'World drop',
          source_kind: 'world_drop',
          dps_delta: -0.7432720066167775,
          verified: true,
        },
      ],
    };
    const items = new Map([[13136, item(13136, "Lil Timmy's Peashooter", { agility: 4 })]]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { ranged: 13136 } },
    };
    const result = upgradesFor(
      character,
      band({ slots: [rangerBow], reference_dps_per_point: 0.059705486370807484 }),
      items,
    );
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]!.notSimChecked).toBe(false);
    expect(result.upgrades[0]!.gainDps).toBeCloseTo(0.7432720066167775, 5);
    expect(result.totalGainDps).toBeCloseTo(0.7432720066167775, 5);
  });

  it("rule (c): counts an empty slot (nothing equipped) as an upgrade, with the pick's full score as the gain", () => {
    const items = new Map([[1, item(1, 'Band Pick Hood', { agility: 5 })]]);
    const character: Pick<MeCharacter, 'build'> = { build: { source: 'addon', captured_at: '', gear: {} } };
    const result = upgradesFor(character, band({ slots: [PICK] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]!.wornItemId).toBeUndefined();
    expect(result.upgrades[0]!.gainDps).toBeCloseTo(1); // (10 - 0) * 0.1
  });

  it("rule (c): never fabricates a gain for a worn item id this build's item table does not carry", () => {
    const items = new Map([[1, item(1, 'Band Pick Hood', { agility: 5 })]]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { head: 999 } },
    };
    const result = upgradesFor(character, band({ slots: [PICK] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]).toMatchObject({
      wornItemId: 999,
      wornUnknown: true,
      gainDps: null,
      notSimChecked: false,
    });
    expect(result.totalGainDps).toBe(0);
  });

  // --- rule (d): weapon slot, no alternative match -> gain unknown, notSimChecked ---------
  it('rule (d): a weapon slot with no alternative match is still an upgrade, but the gain is unknown and excluded from the total', () => {
    const weaponPick = pick('ranged', 5, 'Sim Decided Bow', undefined);
    const items = new Map([
      [5, item(5, 'Sim Decided Bow', {})],
      [6, item(6, 'Old Bow', {})],
    ]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { ranged: 6 } },
    };
    const result = upgradesFor(character, band({ slots: [weaponPick] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]).toMatchObject({ gainDps: null, notSimChecked: true });
    expect(result.totalGainDps).toBe(0);
    expect(result.notSimCheckedCount).toBe(1);
  });

  it('rule (d): a weapon slot never uses scoreItem even when the pick does carry a score', () => {
    // If this fell through to the scoreItem path (as it did before fix round 1), it would
    // compute a large, wrong "gain" from the pick's full-sim score. It must not.
    const weaponPickWithScore = pick('main_hand', 7, 'Big Score Axe', 500);
    const items = new Map([
      [7, item(7, 'Big Score Axe', { agility: 1 })],
      [8, item(8, 'Starter Axe', { agility: 1 })],
    ]);
    const character: Pick<MeCharacter, 'build'> = {
      build: { source: 'addon', captured_at: '', gear: { main_hand: 8 } },
    };
    const result = upgradesFor(character, band({ slots: [weaponPickWithScore] }), items);
    expect(result.upgrades).toHaveLength(1);
    expect(result.upgrades[0]).toMatchObject({ gainDps: null, notSimChecked: true });
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
