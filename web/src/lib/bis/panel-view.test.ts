// web/src/lib/bis/panel-view.test.ts
import { describe, expect, it } from 'vitest';
import type { ItemTooltipModel } from '../items/tooltip';
import type { LootFile } from '../sim/loot';
import { bandInfosFor, collectModelsInto, type PanelViewDeps } from './panel-view';
import type { BisAlternative, BisBand, BisFile, BisSlot, ItemDetail, LootQuestsFile } from './types';

function slot(overrides: Partial<BisSlot> = {}): BisSlot {
  return {
    slot: 'head',
    item_id: 1,
    item_name: 'Test Helm',
    source: 'Vendor: Someone',
    source_kind: 'vendor',
    score: 1,
    verified: true,
    alternatives: [],
    ...overrides,
  };
}

function missingSlot(name: string): BisSlot {
  // The pipeline's own "no pick" shape (`hasKnownSource`'s own doc comment): every item
  // field simply absent, never a placeholder id.
  return { slot: name, verified: false, alternatives: [] } as unknown as BisSlot;
}

function band(overrides: Partial<BisBand> = {}): BisBand {
  return {
    spec: 'hunter-marksmanship',
    band: 20,
    faction: 'alliance',
    race: 'dwarf',
    talents: '',
    talent_points: 1,
    weights: [
      { stat: 'ranged_attack_power', weight: 1, error: 0 },
      { stat: 'agility', weight: 2, error: 0.1 },
      { stat: 'melee_haste', weight: 10, error: 8, insignificant: true },
    ],
    slots: [],
    set_dps: 100,
    no_source_count: 0,
    new_at_band: [],
    coverage: {},
    weights_run_seconds: 0,
    verify_run_seconds: 0,
    reference_dps_per_point: null,
    ...overrides,
  };
}

function fileWith(bands: BisBand[]): BisFile {
  return {
    spec: 'hunter-marksmanship',
    build: 'test',
    engine_version: 'test',
    generated_at: new Date(0).toISOString(),
    bands,
  };
}

const EMPTY_LOOT: LootFile & LootQuestsFile = { sources: [], quests: {} };

function depsWith(overrides: Partial<PanelViewDeps> = {}): PanelViewDeps {
  return {
    itemDetails: new Map<number, ItemDetail>(),
    loot: EMPTY_LOOT,
    tooltipFor: () => undefined,
    referenceStat: 'ranged_attack_power',
    spec: 'hunter-marksmanship',
    ...overrides,
  };
}

describe('bandInfosFor: empty slots', () => {
  it('gives an ordinary empty slot the plain "no known source" copy', () => {
    const file = fileWith([band({ slots: [missingSlot('off_hand')] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const offHand = infos[0].rows.find((r) => r.slot === 'off_hand');
    expect(offHand?.empty).toBe(true);
    expect(offHand?.emptyCopy).toBe('No sourced item at this level yet');
  });

  it('gives an empty off-hand its own copy when the main hand is a two-hander (ruling 4), never the generic no-source line', () => {
    const twoHandModel: ItemTooltipModel = {
      id: 99,
      name: 'Big Staff',
      quality: 3,
      icon: 'inv_staff_25',
      slotLabel: 'Main Hand',
      typeLabel: 'Two-Handed Weapon',
      itemLevel: 20,
      requiredLevel: 18,
      armor: null,
      weapon: null,
      stats: [],
      effectText: null,
      setName: null,
      sourceLines: [],
      unique: false,
    };
    const file = fileWith([
      band({ slots: [slot({ slot: 'main_hand', item_id: 99 }), missingSlot('off_hand')] }),
    ]);
    const infos = bandInfosFor(
      file,
      [20],
      'alliance',
      depsWith({ tooltipFor: (id) => (id === 99 ? twoHandModel : undefined) }),
    );
    const offHand = infos[0].rows.find((r) => r.slot === 'off_hand');
    expect(offHand?.empty).toBe(true);
    expect(offHand?.emptyCopy).toBe('Two-hander equipped');
  });

  it('never applies the two-hander copy to a slot other than off-hand', () => {
    const file = fileWith([band({ slots: [missingSlot('finger1')] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const finger = infos[0].rows.find((r) => r.slot === 'finger1');
    expect(finger?.emptyCopy).toBe('No sourced item at this level yet');
  });
});

describe('bandInfosFor: alternatives', () => {
  const alt: BisAlternative = {
    item_id: 42,
    item_name: 'Runner-up Cap',
    score: 20,
    source_kind: 'vendor',
    source: 'Vendor: Someone Else',
    score_delta: -13.3,
    dps_delta: -0.8,
  };

  it('resolves every alternative’s source cell and carries its dps_delta through', () => {
    const loot: LootFile & LootQuestsFile = {
      sources: [{ id: 'vendor:someone-else', kind: 'vendor', name: 'Someone Else', items: [42] }],
      quests: {},
    };
    const file = fileWith([band({ slots: [slot({ alternatives: [alt] })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith({ loot }));
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.alternatives).toHaveLength(1);
    expect(head?.alternatives?.[0]).toMatchObject({ itemId: 42, itemName: 'Runner-up Cap', dpsDelta: -0.8 });
    expect(head?.alternatives?.[0].sourceKind).toBe('vendor');
    expect(head?.alternatives?.[0].sourceDetail).toBe('Vendor: Someone Else');
  });

  it('falls back to the generic per-kind label, never a throw, when loot.json has nothing for the alternative’s id', () => {
    const file = fileWith([band({ slots: [slot({ alternatives: [alt] })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.alternatives?.[0].sourceKind).toBe('unknown');
  });

  it('is an empty array, not undefined, when the pick has none', () => {
    const file = fileWith([band({ slots: [slot()] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.alternatives).toEqual([]);
  });

  it('carries verified through from the alternative, undefined when the ranker never simmed it', () => {
    const file = fileWith([
      band({
        slots: [
          slot({
            alternatives: [
              { ...alt, verified: true },
              { ...alt, item_id: 43 },
            ],
          }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const [first, second] = infos[0].rows.find((r) => r.slot === 'head')?.alternatives ?? [];
    expect(first?.verified).toBe(true);
    expect(second?.verified).toBeUndefined();
  });

  it('builds metaLabel from the alternative’s own tooltip model, "needs" only above the band', () => {
    const belowBandModel: ItemTooltipModel = {
      id: 42,
      name: 'Runner-up Cap',
      quality: 2,
      icon: 'inv_helmet_01',
      slotLabel: 'Head',
      typeLabel: undefined,
      itemLevel: 24,
      requiredLevel: 18,
      armor: null,
      weapon: null,
      stats: [],
      effectText: null,
      setName: null,
      sourceLines: [],
      unique: false,
    };
    const aboveBandModel: ItemTooltipModel = { ...belowBandModel, id: 43, requiredLevel: 25 };
    const file = fileWith([band({ slots: [slot({ alternatives: [alt, { ...alt, item_id: 43 }] })] })]);
    const tooltipFor = (id: number): ItemTooltipModel | undefined =>
      id === 42 ? belowBandModel : id === 43 ? aboveBandModel : undefined;
    const infos = bandInfosFor(file, [20], 'alliance', depsWith({ tooltipFor }));
    const [first, second] = infos[0].rows.find((r) => r.slot === 'head')?.alternatives ?? [];
    expect(first?.metaLabel).toBe('ilvl 24');
    expect(second?.metaLabel).toBe('ilvl 24 · needs 25');
  });

  it('leaves metaLabel undefined when the alternative’s id has no tooltip model', () => {
    const file = fileWith([band({ slots: [slot({ alternatives: [alt] })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.alternatives?.[0].metaLabel).toBeUndefined();
  });
});

describe('collectModelsInto', () => {
  it('adds an alternative’s model to the shared tooltip map, not just the main pick’s', () => {
    const altModel: ItemTooltipModel = {
      id: 42,
      name: 'Runner-up Cap',
      quality: 2,
      icon: 'inv_helmet_01',
      slotLabel: 'Head',
      typeLabel: undefined,
      itemLevel: 15,
      requiredLevel: 18,
      armor: 10,
      weapon: null,
      stats: [],
      effectText: null,
      setName: null,
      sourceLines: [],
      unique: false,
    };
    const alt: BisAlternative = {
      item_id: 42,
      item_name: 'Runner-up Cap',
      score: 20,
      source_kind: 'vendor',
      source: 'Vendor: Someone Else',
      score_delta: 0,
      dps_delta: 0,
    };
    const file = fileWith([band({ slots: [slot({ alternatives: [alt] })] })]);
    const tooltipFor = (id: number): ItemTooltipModel | undefined => (id === 42 ? altModel : undefined);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith({ tooltipFor }));
    const models = new Map<number, ItemTooltipModel>();
    collectModelsInto(models, infos, tooltipFor);
    expect(models.get(42)).toEqual(altModel);
  });
});

describe('bandInfosFor: weight rail', () => {
  it('leaves every dpsPerPoint undefined when the band carries no reference_dps_per_point (never fabricated)', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].weightBars.every((bar) => bar.dpsPerPoint === undefined)).toBe(true);
    expect(infos[0].referenceSentenceLine.toLowerCase()).toContain('reference');
  });

  it('computes a significant row’s dpsPerPoint as weight * reference_dps_per_point when present', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: 2.5 })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const agility = infos[0].weightBars.find((bar) => bar.row.stat === 'agility');
    expect(agility?.dpsPerPoint).toBeCloseTo(2 * 2.5);
    expect(infos[0].referenceSentenceLine).toBe('1 Ranged attack power = 2.50 DPS');
  });

  it('never computes a dpsPerPoint for an insignificant row, which shows "No effect" instead', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: 2.5 })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const meleeHaste = infos[0].weightBars.find((bar) => bar.row.stat === 'melee_haste');
    expect(meleeHaste?.row.significant).toBe(false);
    expect(meleeHaste?.dpsPerPoint).toBeUndefined();
  });

  it('never computes a dpsPerPoint for the reference row itself (always "= 1")', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: 2.5 })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const reference = infos[0].weightBars.find((bar) => bar.row.isReference);
    expect(reference?.dpsPerPoint).toBeUndefined();
  });
});
