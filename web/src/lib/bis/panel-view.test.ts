// web/src/lib/bis/panel-view.test.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import type { ItemTooltipModel } from '../items/tooltip';
import type { LootFile } from '../sim/loot';
import { bisCopy } from './copy';
import { bandInfosFor, collectModelsInto, parseSwapNote, type PanelViewDeps } from './panel-view';
import type {
  BisAlternative,
  BisBand,
  BisFile,
  BisHealProfile,
  BisSlot,
  ItemDetail,
  LootQuestsFile,
} from './types';
import { SLOTS } from '../planner/types';

const REAL_HUNTER_MARKSMANSHIP = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../data/builds/1.60.1.70009/bis/hunter-marksmanship.json',
);

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
      clientUnconfirmed: false,
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
    expect(offHand?.emptyCopy).toBe('Test Helm is a two-hander; the off hand is taken.');
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
    source_kind: 'vendor',
    source: 'Vendor: Someone Else',
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
      clientUnconfirmed: false,
    };
    const aboveBandModel: ItemTooltipModel = { ...belowBandModel, id: 43, requiredLevel: 25 };
    const file = fileWith([band({ slots: [slot({ alternatives: [alt, { ...alt, item_id: 43 }] })] })]);
    const tooltipFor = (id: number): ItemTooltipModel | undefined =>
      id === 42 ? belowBandModel : id === 43 ? aboveBandModel : undefined;
    const infos = bandInfosFor(file, [20], 'alliance', depsWith({ tooltipFor }));
    const [first, second] = infos[0].rows.find((r) => r.slot === 'head')?.alternatives ?? [];
    expect(first?.metaLabel).toBe('ilvl 24');
    expect(second?.metaLabel).toBe('ilvl 24 · needs 25');
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
      clientUnconfirmed: false,
    };
    const alt: BisAlternative = {
      item_id: 42,
      item_name: 'Runner-up Cap',
      source_kind: 'vendor',
      source: 'Vendor: Someone Else',
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

describe('bandInfosFor: scale rail', () => {
  // band()'s own default weights: ranged_attack_power weight 1, agility weight 2 (the
  // largest per-point row -> the anchor), melee_haste weight 10 insignificant -- none of
  // them publish `scale_factor` (the fixture predates this lane), so every assertion below
  // exercises the client-side fallback (`computeScaleFactors`'s own "no published fields"
  // branch), the exact case this lane's brief calls out: "the rail must fall back
  // gracefully on a JSON that lacks the new fields".

  it('leaves every dpsPerPoint undefined when the band carries no reference_dps_per_point (never fabricated)', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].scaleRows.every((row) => row.dpsPerPoint === undefined)).toBe(true);
  });

  it('computes a row’s dpsPerPoint as weight * reference_dps_per_point when present, regardless of significance', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: 2.5 })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const agility = infos[0].scaleRows.find((row) => row.stat === 'agility');
    expect(agility?.dpsPerPoint).toBeCloseTo(2 * 2.5);
  });

  it('normalizes every scale factor against the band’s own top per-point stat (agility, weight 2 here), never the engine’s old reference stat', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const agility = infos[0].scaleRows.find((row) => row.stat === 'agility');
    const rap = infos[0].scaleRows.find((row) => row.stat === 'ranged_attack_power');
    expect(agility?.scaleFactor).toBe(1);
    expect(rap?.scaleFactor).toBeCloseTo(0.5);
  });

  it('sorts the table strictly descending by scaleFactor', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].scaleRows.map((row) => row.stat)).toEqual(['agility', 'ranged_attack_power']);
  });

  it('never puts a haste stat in the table (owner correction, 2026-09-30, after player review)', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].scaleRows.some((row) => row.stat === 'melee_haste')).toBe(false);
  });

  it('gives haste a plain caption by default (haste_on_items absent defaults to true), even though band()’s own melee_haste row is insignificant', () => {
    // Owner fix, 2026-09-30, found on screenshot review: a haste row's own `insignificant`
    // flag is a STATISTICAL question (did the sweep's sample clear its own noise bar), never
    // the plain inventory question `haste_on_items` answers -- conflating the two is exactly
    // what produced the previous, doubled "Haste: 1.58 per 1%, per 1%" caption. band()'s own
    // default fixture marks melee_haste insignificant but never sets haste_on_items, so the
    // caption must still read plain (default true, no clause) here.
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    // melee_haste weight 10 / agility's own anchor weight 2 = 5.00.
    expect(infos[0].hasteCaptionLine).toBe(bisCopy.weightsHasteCaption(5, false));
    expect(infos[0].hasteCaptionLine).toBe('Haste: 5.00 per 1%');
    expect(infos[0].hasteCaptionLine).not.toContain('not in the table');
  });

  it('adds the "not in the table" clause only when the band’s own haste_on_items is explicitly false', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null, haste_on_items: false })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].hasteCaptionLine).toBe(bisCopy.weightsHasteCaption(5, true));
    expect(infos[0].hasteCaptionLine).toBe(
      'Haste: 5.00 per 1%, not in the table because no item at this band has it',
    );
  });

  it('gives haste the plain caption when its own row is significant and haste_on_items is true', () => {
    const file = fileWith([
      band({
        weights: [
          { stat: 'ranged_attack_power', weight: 1, error: 0 },
          { stat: 'agility', weight: 2, error: 0.1 },
          { stat: 'melee_haste', weight: 3, error: 0.2 },
        ],
        slots: [slot()],
        reference_dps_per_point: null,
        haste_on_items: true,
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].hasteCaptionLine).toBe(bisCopy.weightsHasteCaption(1.5, false));
    expect(infos[0].hasteCaptionLine).toBe('Haste: 1.50 per 1%');
  });

  it('gives no haste caption when the spec carries no haste weight_stat at all', () => {
    const file = fileWith([
      band({
        weights: [
          { stat: 'spell_power', weight: 1, error: 0 },
          { stat: 'intellect', weight: 0.5, error: 0.02 },
        ],
        slots: [slot()],
        reference_dps_per_point: null,
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].hasteCaptionLine).toBeUndefined();
  });

  it('carries the hit-to-cap line, with the white clause only when published', () => {
    const file = fileWith([
      band({ slots: [slot()], hit_to_cap: { baseline: 3, specials: 6, white: 25 } }),
      band({ band: 30, slots: [slot()], hit_to_cap: { baseline: 3, specials: 6.5 } }),
    ]);
    const [dual, single] = bandInfosFor(file, [20, 30], 'alliance', depsWith());
    expect(dual.hitCap?.text).toBe('Hit to cap: 6% for specials, 25% for white swings');
    expect(single.hitCap?.text).toBe('Hit to cap: 6.5% for specials');
    expect(dual.hitCap?.title).toBe(bisCopy.hitToCapTitle);
  });

  it('carries no hit-to-cap line when the band publishes none', () => {
    const file = fileWith([band({ slots: [slot()] })]);
    expect(bandInfosFor(file, [20], 'alliance', depsWith())[0].hitCap).toBeUndefined();
    const nulled = fileWith([band({ slots: [slot()], hit_to_cap: null })]);
    expect(bandInfosFor(nulled, [20], 'alliance', depsWith())[0].hitCap).toBeUndefined();
  });

  it('states the band’s own top stat in scaleNoteLine', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: null })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].scaleNoteLine).toBe(bisCopy.weightsScaleNote('Agility'));
  });

  it('renders an empty table and the unmeasured-weights line when the band carries weights_reason, even with a real reference_dps_per_point', () => {
    const file = fileWith([
      band({
        slots: [slot()],
        reference_dps_per_point: 2.5,
        weights_reason: 'reference stat spell_power measured -0.1893 ± 0.6199 DPS per point',
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].scaleRows).toEqual([]);
    expect(infos[0].scaleNoteLine).toBe(bisCopy.weightsUnmeasuredLine);
    expect(infos[0].hasteCaptionLine).toBeUndefined();
  });

  it('never shows the unmeasured-weights line when weights_reason is absent', () => {
    const file = fileWith([band({ slots: [slot()], reference_dps_per_point: 2.5 })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].scaleNoteLine).not.toBe(bisCopy.weightsUnmeasuredLine);
  });

  it('reads the ranker’s own published scale_factor/dps_per_point/scale_error instead of recomputing them when present', () => {
    const file = fileWith([
      band({
        weights: [
          {
            stat: 'ranged_attack_power',
            weight: 1,
            error: 0,
            scale_factor: 0.5,
            dps_per_point: 1.25,
            scale_error: 0,
          },
          { stat: 'agility', weight: 2, error: 0.1, scale_factor: 1, dps_per_point: 2.5, scale_error: 0.05 },
        ],
        slots: [slot()],
        reference_dps_per_point: 2.5,
        scale_reference_stat: 'agility',
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const agility = infos[0].scaleRows.find((row) => row.stat === 'agility');
    expect(agility?.scaleFactor).toBe(1);
    expect(agility?.dpsPerPoint).toBe(2.5);
    expect(agility?.scaleError).toBe(0.05);
    expect(infos[0].scaleNoteLine).toBe(bisCopy.weightsScaleNote('Agility'));
  });

  // Spec addendum 2 (weight rail rows go per rating point): a `unit: 'rating'` row's label
  // gains a quiet " rating" suffix and its native `title` reads the client's own
  // rating-per-percent conversion, while a plain, non-rating row (e.g. Agility here) keeps
  // its plain label and no title override.
  it('gives a rating-family row a " rating" label and its rating-factor title', () => {
    const file = fileWith([
      band({
        weights: [
          { stat: 'ranged_attack_power', weight: 1, error: 0 },
          { stat: 'agility', weight: 2, error: 0.1 },
          {
            stat: 'crit',
            weight: 1.44,
            error: 0.07,
            unit: 'rating',
            rating_factor: 14,
            weight_per_percent: 20.16,
          },
        ],
        slots: [slot()],
        reference_dps_per_point: 2.5,
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const crit = infos[0].scaleRows.find((row) => row.stat === 'crit');
    expect(crit?.label).toBe('Crit rating');
    expect(crit?.ratingFactorTitle).toBe('14 Crit rating = 1% Crit');
  });

  it('leaves a non-rating row’s label and title untouched (no suffix, no override)', () => {
    const file = fileWith([
      band({
        weights: [
          { stat: 'ranged_attack_power', weight: 1, error: 0 },
          { stat: 'agility', weight: 2, error: 0.1 },
          {
            stat: 'crit',
            weight: 1.44,
            error: 0.07,
            unit: 'rating',
            rating_factor: 14,
            weight_per_percent: 20.16,
          },
        ],
        slots: [slot()],
        reference_dps_per_point: 2.5,
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const agility = infos[0].scaleRows.find((row) => row.stat === 'agility');
    expect(agility?.label).toBe('Agility');
    expect(agility?.label).not.toContain('rating');
    expect(agility?.ratingFactorTitle).toBeUndefined();
  });

  it('keeps an insignificant rating row in the table, marked significant: false, with the " rating" label', () => {
    const file = fileWith([
      band({
        weights: [
          { stat: 'ranged_attack_power', weight: 1, error: 0 },
          { stat: 'agility', weight: 2, error: 0.1 },
          {
            stat: 'hit',
            weight: 0.06,
            error: 0.3,
            insignificant: true,
            unit: 'rating',
            rating_factor: 10,
            weight_per_percent: 0.6,
          },
        ],
        slots: [slot()],
        reference_dps_per_point: 2.5,
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const hit = infos[0].scaleRows.find((row) => row.stat === 'hit');
    expect(hit?.label).toBe('Hit rating');
    expect(hit?.significant).toBe(false);
    expect(hit?.dpsPerPoint).toBeCloseTo(0.06 * 2.5);
  });
});

describe('parseSwapNote', () => {
  it('parses the "confirmed by the sim against" template, pick DPS first', () => {
    expect(
      parseSwapNote(
        'confirmed by the sim against Diamond Hammer (id 2194): kept the pick, 45.8 vs 38.0 set DPS',
      ),
    ).toEqual({ itemName: 'Diamond Hammer', pickDps: 45.8, altDps: 38.0 });
  });

  it('parses the "beat the scored pick" template, the same way', () => {
    expect(
      parseSwapNote('beat the scored pick Wolfmaster Cape (id 6314) in the sim: 70.9 vs 69.9 set DPS'),
    ).toEqual({ itemName: 'Wolfmaster Cape', pickDps: 70.9, altDps: 69.9 });
  });

  it('is undefined for a swap_note format it does not recognise, never a guessed reading', () => {
    expect(
      parseSwapNote(
        "the runner-up's verification sim failed (an engine-side error, not a scoring one - see verify_errors); the pick is unconfirmed against it",
      ),
    ).toBeUndefined();
    expect(parseSwapNote('')).toBeUndefined();
  });

  it('matches every real swap_note in the 2026-09-29 build fixture shape (a multi-word item name, a decimal id-adjacent DPS)', () => {
    expect(
      parseSwapNote(
        "confirmed by the sim against Ironspine's Fist (id 7687): kept the pick, 69.9 vs 70.7 set DPS",
      ),
    ).toEqual({ itemName: "Ironspine's Fist", pickDps: 69.9, altDps: 70.7 });
  });
});

describe('bandInfosFor: evidence line and verified-glyph title', () => {
  it('renders the evidence line in player words when swap_note matches a known ranker template', () => {
    const file = fileWith([
      band({
        slots: [
          slot({
            swap_note:
              'confirmed by the sim against Diamond Hammer (id 2194): kept the pick, 45.8 vs 38.0 set DPS',
          }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.evidenceLine).toBe('Sim-checked against Diamond Hammer: 45.8 vs 38.0 DPS');
  });

  it('prefers the row’s own dps_delta over swap_note’s two absolute numbers when both are present', () => {
    const file = fileWith([
      band({
        slots: [
          slot({
            swap_note:
              'confirmed by the sim against Diamond Hammer (id 2194): kept the pick, 45.8 vs 38.0 set DPS',
            dps_delta: 7.8,
          }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.evidenceLine).toBe('Sim-checked against Diamond Hammer: +7.8 DPS');
  });

  it('prefers dps_delta for the bis-ranker-integrity-5 "over it" swap_note suffix, which the old two-number parse cannot match', () => {
    const file = fileWith([
      band({
        slots: [
          slot({
            swap_note:
              "confirmed by the sim against Ironspine's Fist (id 7687): kept the pick, +4.5 DPS over it",
            dps_delta: 4.5,
          }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.evidenceLine).toBe("Sim-checked against Ironspine's Fist: +4.5 DPS");
  });

  it('prefers dps_delta for the "beat the scored pick" template too', () => {
    const file = fileWith([
      band({
        slots: [
          slot({
            swap_note: 'beat the scored pick Diamond Hammer (id 2194) in the sim: 40.2 vs 38.0 set DPS',
            dps_delta: 2.2,
          }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.evidenceLine).toBe('Sim-checked against Diamond Hammer: +2.2 DPS');
  });

  it('falls back to the old two-number parse when dps_delta is absent (a file published before this field existed)', () => {
    const file = fileWith([
      band({
        slots: [
          slot({
            swap_note:
              'confirmed by the sim against Diamond Hammer (id 2194): kept the pick, 45.8 vs 38.0 set DPS',
            dps_delta: null,
          }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.evidenceLine).toBe('Sim-checked against Diamond Hammer: 45.8 vs 38.0 DPS');
  });

  it('falls back to the raw swap_note text when dps_delta is present but the swap_note format is unrecognised', () => {
    const file = fileWith([
      band({ slots: [slot({ swap_note: 'something the parser has never seen', dps_delta: 3.1 })] }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.evidenceLine).toBe(
      'something the parser has never seen',
    );
  });

  it('falls back to the raw swap_note text when the format is unrecognised', () => {
    const file = fileWith([band({ slots: [slot({ swap_note: 'something the parser has never seen' })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.evidenceLine).toBe(
      'something the parser has never seen',
    );
  });

  it('leaves evidenceLine undefined when the pick has no swap_note at all', () => {
    const file = fileWith([band({ slots: [slot()] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.evidenceLine).toBeUndefined();
  });

  it('gives the verified glyph a sim-DPS title when the pick has sim_dps and no swap_note', () => {
    const file = fileWith([band({ slots: [slot({ sim_dps: 78.6, swap_note: undefined })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.verifiedGlyphTitle).toBe(
      'Confirmed by a full sim: 78.6 DPS with this item',
    );
  });

  it('leaves the verified glyph title undefined when a swap_note is already telling the story', () => {
    const file = fileWith([
      band({
        slots: [
          slot({
            sim_dps: 78.6,
            swap_note:
              'confirmed by the sim against Diamond Hammer (id 2194): kept the pick, 45.8 vs 38.0 set DPS',
          }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.verifiedGlyphTitle).toBeUndefined();
  });

  it('leaves the verified glyph title undefined for an ordinary weight-ranked pick (no sim_dps, no swap_note)', () => {
    const file = fileWith([band({ slots: [slot()] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.verifiedGlyphTitle).toBeUndefined();
  });
});

describe('bandInfosFor: effect_unmodelled and low_value flags', () => {
  it('carries effect_unmodelled through on a pick', () => {
    const file = fileWith([band({ slots: [slot({ effect_unmodelled: true })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.effectUnmodelled).toBe(true);
  });

  it('is undefined on a pick that carries no effect_unmodelled flag', () => {
    const file = fileWith([band({ slots: [slot()] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.effectUnmodelled).toBeUndefined();
  });

  it('carries effect_unmodelled through on an alternative independently of the pick', () => {
    const alt: BisAlternative = {
      item_id: 42,
      item_name: 'Weakness Analyzer',
      source_kind: 'vendor',
      source: 'Vendor: Someone',
      dps_delta: -4.5,
      effect_unmodelled: true,
    };
    const file = fileWith([band({ slots: [slot({ alternatives: [alt] })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.effectUnmodelled).toBeUndefined();
    expect(head?.alternatives?.[0].effectUnmodelled).toBe(true);
  });

  it('carries low_value through on a weapon row', () => {
    const file = fileWith([band({ slots: [slot({ slot: 'ranged', low_value: true })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'ranged')?.lowValue).toBe(true);
  });
});

describe('bandInfosFor: sim_status and set_dps_partial (spec addendum 3)', () => {
  it('sets notSimChecked on a pick carrying sim_status "not_in_sim"', () => {
    const file = fileWith([band({ slots: [slot({ sim_status: 'not_in_sim' })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.notSimChecked).toBe(true);
  });

  it('is undefined on a pick that carries no sim_status flag', () => {
    const file = fileWith([band({ slots: [slot()] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'head')?.notSimChecked).toBe(false);
  });

  it('carries both notSimChecked and effectUnmodelled independently on one row', () => {
    const file = fileWith([band({ slots: [slot({ sim_status: 'not_in_sim', effect_unmodelled: true })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect(head?.notSimChecked).toBe(true);
    expect(head?.effectUnmodelled).toBe(true);
  });

  it('never flags an alternative -- BisAlternative carries no sim_status field at all', () => {
    const alt: BisAlternative = {
      item_id: 42,
      item_name: 'Weakness Analyzer',
      source_kind: 'vendor',
      source: 'Vendor: Someone',
      dps_delta: -4.5,
    };
    const file = fileWith([band({ slots: [slot({ sim_status: 'not_in_sim', alternatives: [alt] })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const head = infos[0].rows.find((r) => r.slot === 'head');
    expect((head?.alternatives?.[0] as { notSimChecked?: boolean }).notSimChecked).toBeUndefined();
  });

  it('reports setDpsPartial false and a zero count when the band carries no flag', () => {
    const file = fileWith([band({ slots: [slot()] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].setDpsPartial).toBe(false);
    expect(infos[0].setDpsPartialCount).toBe(0);
  });

  it('reports setDpsPartial true with a count matching the notSimChecked rows', () => {
    const file = fileWith([
      band({
        set_dps_partial: true,
        slots: [
          slot({ sim_status: 'not_in_sim' }),
          slot({ slot: 'neck', sim_status: 'not_in_sim' }),
          slot({ slot: 'shoulder' }),
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].setDpsPartial).toBe(true);
    expect(infos[0].setDpsPartialCount).toBe(2);
  });

  it('computes setDpsPartialCount from the same rows array, never a second counter', () => {
    // Even when the band's own flag disagrees with reality (should never happen from a
    // real pipeline file, but the view layer must never trust a second backend count it
    // did not itself derive from `rows`), the count always matches the rows the template
    // actually renders.
    const file = fileWith([band({ set_dps_partial: true, slots: [slot({ sim_status: 'not_in_sim' })] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    const notSimCheckedRows = infos[0].rows.filter((r) => r.notSimChecked).length;
    expect(infos[0].setDpsPartialCount).toBe(notSimCheckedRows);
  });
});

describe('bandInfosFor: empty_reason copy', () => {
  it('reads the thin-pool line for no_dps_value', () => {
    const file = fileWith([
      band({ slots: [{ ...missingSlot('trinket1'), empty_reason: 'no_dps_value' } as BisSlot] }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe(
      'Nothing sourced at this level helps your DPS',
    );
  });

  it('reads the unmodelled-relic line for effect_not_modelled', () => {
    const file = fileWith([
      band({ slots: [{ ...missingSlot('ranged'), empty_reason: 'effect_not_modelled' } as BisSlot] }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'ranged')?.emptyCopy).toBe(
      "Relic effects aren't simulated yet",
    );
  });

  it('computes the "no sourced item" sentence for no_sourced_item, and falls back to the plain line for any unrecognised value', () => {
    const file = fileWith([
      band({
        slots: [
          { ...missingSlot('finger1'), empty_reason: 'no_sourced_item' } as BisSlot,
          { ...missingSlot('finger2'), empty_reason: 'some_future_reason' } as BisSlot,
        ],
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    // finger1/finger2 both collapse to the "Ring" display label (bis rebuild spec §4.D) --
    // no later band in this one-band fixture to name, so the second sentence is dropped.
    expect(infos[0].rows.find((r) => r.slot === 'finger1')?.emptyCopy).toBe(
      'No ring you can get at 20 to 29 raises your damage.',
    );
    expect(infos[0].rows.find((r) => r.slot === 'finger2')?.emptyCopy).toBe(
      'No sourced item at this level yet',
    );
  });

  it('a truly missing slot (no empty_reason field at all) still gets the plain no-source line', () => {
    const file = fileWith([band({ slots: [missingSlot('finger1')] })]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'finger1')?.emptyCopy).toBe(
      'No sourced item at this level yet',
    );
  });

  it('the off-hand-under-a-two-hander rule still wins over any empty_reason the row carries', () => {
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
      clientUnconfirmed: false,
    };
    const file = fileWith([
      band({
        slots: [
          slot({ slot: 'main_hand', item_id: 99 }),
          { ...missingSlot('off_hand'), empty_reason: 'no_dps_value' } as BisSlot,
        ],
      }),
    ]);
    const infos = bandInfosFor(
      file,
      [20],
      'alliance',
      depsWith({ tooltipFor: (id) => (id === 99 ? twoHandModel : undefined) }),
    );
    expect(infos[0].rows.find((r) => r.slot === 'off_hand')?.emptyCopy).toBe(
      'Test Helm is a two-hander; the off hand is taken.',
    );
  });
});

describe('bandInfosFor: no_sourced_item across bands (bis rebuild spec §4.D)', () => {
  const noSourcedTrinket = { ...missingSlot('trinket1'), empty_reason: 'no_sourced_item' } as BisSlot;

  it('names the real "comes at N" band by scanning forward, never the next band by default', () => {
    const file = fileWith([
      band({ band: 20, slots: [noSourcedTrinket] }),
      band({ band: 30, slots: [noSourcedTrinket] }),
      band({ band: 40, slots: [slot({ slot: 'trinket1' })] }),
    ]);
    const infos = bandInfosFor(file, [20, 30, 40], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe(
      'No trinket you can get at 20 to 29 raises your damage. The first that does comes at 40.',
    ); // "comes at N" names the plain band number, never its own "N to N+9" range
    expect(infos[1].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe('Nothing here either until 40.');
    expect(infos[2].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBeUndefined();
  });

  it('drops the second sentence entirely when no later band ever sources the slot', () => {
    const file = fileWith([
      band({ band: 20, slots: [noSourcedTrinket] }),
      band({ band: 30, slots: [noSourcedTrinket] }),
    ]);
    const infos = bandInfosFor(file, [20, 30], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe(
      'No trinket you can get at 20 to 29 raises your damage.',
    );
    expect(infos[1].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe(
      'Nothing here helps at this level either.',
    );
  });

  it('plain 60 never reads "60 to 69"', () => {
    const file = fileWith([band({ band: 60, slots: [noSourcedTrinket] })]);
    const infos = bandInfosFor(file, [60], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe(
      'No trinket you can get at 60 raises your damage.',
    );
  });

  it('the second empty trinket row in the SAME band defers to the short sentence (wow-player review round 1): never repeats the first row’s full sentence verbatim', () => {
    const noSourcedTrinket2 = { ...missingSlot('trinket2'), empty_reason: 'no_sourced_item' } as BisSlot;
    const file = fileWith([
      band({ band: 20, slots: [noSourcedTrinket, noSourcedTrinket2] }),
      band({ band: 40, slots: [slot({ slot: 'trinket1' }), slot({ slot: 'trinket2' })] }),
    ]);
    const infos = bandInfosFor(file, [20, 40], 'alliance', depsWith());
    expect(infos[0].rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe(
      'No trinket you can get at 20 to 29 raises your damage. The first that does comes at 40.',
    );
    expect(infos[0].rows.find((r) => r.slot === 'trinket2')?.emptyCopy).toBe('Nothing here either until 40.');
  });
});

describe('bandInfosFor: no_sourced_item against the real published file (ux-designer review round 1)', () => {
  // A regression guard against the real data, not a synthetic fixture (tenet 8): both
  // trinket slots are `no_sourced_item` through band 40 in the 2026-09-30 nightly regen for
  // BOTH factions, first sourced at band 50 -- this pins that real number so a future regen
  // that moves it fails loudly here rather than silently changing the page's own sentence.
  const real = JSON.parse(readFileSync(REAL_HUNTER_MARKSMANSHIP, 'utf8')) as BisFile;
  const bands = [...new Set(real.bands.map((b) => b.band))].sort((a, b) => a - b);

  it.each(['alliance', 'horde'] as const)('names band 50 as where trinkets first help, %s', (faction) => {
    const infos = bandInfosFor(real, bands, faction, depsWith({ spec: 'hunter-marksmanship' }));
    const band20 = infos.find((info) => info.band === 20)!;
    expect(band20.rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe(
      'No trinket you can get at 20 to 29 raises your damage. The first that does comes at 50.',
    );
    expect(band20.rows.find((r) => r.slot === 'trinket2')?.emptyCopy).toBe('Nothing here either until 50.');
    const band40 = infos.find((info) => info.band === 40)!;
    expect(band40.rows.find((r) => r.slot === 'trinket1')?.emptyCopy).toBe('Nothing here either until 50.');
    const band50 = infos.find((info) => info.band === 50)!;
    expect(band50.rows.find((r) => r.slot === 'trinket1')?.empty).toBe(false);
  });
});

describe('bandInfosFor: totalSlots (spec §4.B/§4.E denominator)', () => {
  it('is 17 with nothing empty', () => {
    const file = fileWith([
      band({
        slots: SLOTS.map((slotName) => slot({ slot: slotName, item_id: 1 + SLOTS.indexOf(slotName) })),
      }),
    ]);
    const infos = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(infos[0].totalSlots).toBe(17);
  });

  it('drops one for a two-hander and one per still-empty trinket', () => {
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
      clientUnconfirmed: false,
    };
    const file = fileWith([
      band({
        slots: [
          slot({ slot: 'main_hand', item_id: 99 }),
          missingSlot('off_hand'),
          { ...missingSlot('trinket1'), empty_reason: 'no_sourced_item' } as BisSlot,
          { ...missingSlot('trinket2'), empty_reason: 'no_sourced_item' } as BisSlot,
        ],
      }),
    ]);
    const infos = bandInfosFor(
      file,
      [20],
      'alliance',
      depsWith({ tooltipFor: (id) => (id === 99 ? twoHandModel : undefined) }),
    );
    // 17 - 1 (off hand, two-hander) - 2 (both trinkets empty) = 14, matching the worked
    // example's verified real number (hunter-marksmanship band 20 horde).
    expect(infos[0].totalSlots).toBe(14);
  });
});

describe('bandInfosFor: presets', () => {
  const file = fileWith([
    band({ band: 50, set_dps: 100 }),
    band({ band: 60, set_dps: 200 }),
    band({ band: 60, set_dps: 300, preset: 'raid' }),
  ]);

  it('reads the raid entry by default and the bare entry on request', () => {
    expect(bandInfosFor(file, [50, 60], 'alliance', depsWith())[1].setDps).toBe(300);
    expect(bandInfosFor(file, [50, 60], 'alliance', depsWith({ preset: 'bare' }))[1].setDps).toBe(200);
  });

  it('only reports a DPS delta against the previous band when both are the same preset', () => {
    const raid = bandInfosFor(file, [50, 60], 'alliance', depsWith())[1];
    const bare = bandInfosFor(file, [50, 60], 'alliance', depsWith({ preset: 'bare' }))[1];
    expect(raid.dpsDelta).toBeUndefined();
    expect(bare.dpsDelta).toBe(100);
  });
});

describe('bandInfosFor: healer hook and unit threading', () => {
  const healerFields: Partial<BisBand> = {
    role: 'healer',
    profile: 'onyxia-sized',
    metrics: { hps: 300, raw_hps: 360, overheal_pct: 0.16, mana_lasts_sec: 192, hpm: 3.3 },
    set_dps: 300,
    reference_dps_per_point: 0.31,
    weights: [
      { stat: 'healing_power', weight: 1, error: 0.01 },
      { stat: 'intellect', weight: 0.5, error: 0.01 },
    ],
  };
  const swapSlot = slot({
    swap_note: 'confirmed by the sim against Old Helm (id 5): kept the pick, 300.0 vs 290.0 set DPS',
    dps_delta: 10,
    sim_dps: 300,
    alternatives: [],
  });

  it('leaves a damage band with no healer view and DPS wording', () => {
    const file = fileWith([band({ slots: [swapSlot], reference_dps_per_point: 0.5 })]);
    const [info] = bandInfosFor(file, [20], 'alliance', depsWith());
    expect(info!.healer).toBeUndefined();
    expect(info!.unit).toBe('DPS');
    expect(info!.scaleNoteLine).toBe(bisCopy.weightsScaleNote('Agility'));
    expect(info!.rows[0]!.evidenceLine).toBe('Sim-checked against Old Helm: +10.0 DPS');
  });

  it('builds the healer view from the band and the file profile, and words every figure in HPS', () => {
    const profile: BisHealProfile = {
      id: 'onyxia-sized',
      label: 'Onyxia-sized tank hits and raid pulses',
      summary: '',
      notes: '',
      duration_sec: 300,
      damage_spread: 0.25,
      tank: { health: 9500, hit_damage: 1150, swing_seconds: 2, reason: '' },
      members: { health: 5000, reason: '' },
      pulse: { damage: 450, interval_seconds: 4, members: 3, reason: '' },
      sources: [],
    };
    const file = { ...fileWith([band({ ...healerFields, slots: [swapSlot] })]), heal_profile: profile };
    const [info] = bandInfosFor(file as BisFile, [20], 'alliance', depsWith({ spec: 'priest-holy' }));
    expect(info!.unit).toBe('HPS');
    expect(info!.healer?.figure).toBe('300.0');
    expect(info!.healer?.profileLabel).toBe('Onyxia-sized tank hits and raid pulses');
    expect(info!.scaleNoteLine).toContain('HPS per point');
    expect(info!.scaleNoteLine).not.toContain('DPS');
    expect(info!.rows[0]!.evidenceLine).toBe('Sim-checked against Old Helm: +10.0 HPS');
    expect(info!.rows[0]!.verifiedGlyphTitle).toBeUndefined();
  });

  it('words the parsed swap note and a no-value slot in HPS too', () => {
    const noDelta = slot({ swap_note: swapSlot.swap_note, dps_delta: null });
    const empty = slot({
      slot: 'neck',
      item_id: undefined as unknown as number,
      empty_reason: 'no_dps_value',
    });
    const file = fileWith([band({ ...healerFields, slots: [noDelta, empty] })]);
    const [info] = bandInfosFor(file, [20], 'alliance', depsWith({ spec: 'priest-holy' }));
    expect(info!.rows[0]!.evidenceLine).toBe('Sim-checked against Old Helm: 300.0 vs 290.0 HPS');
    expect(info!.rows.find((r) => r.slot === 'neck')?.emptyCopy).toBe(
      'Nothing sourced at this level helps your HPS',
    );
  });

  it('names a sim-verified pick with no swap note in HPS', () => {
    const file = fileWith([band({ ...healerFields, slots: [slot({ sim_dps: 281.3 })] })]);
    const [info] = bandInfosFor(file, [20], 'alliance', depsWith({ spec: 'priest-holy' }));
    expect(info!.rows[0]!.verifiedGlyphTitle).toBe('Confirmed by a full sim: 281.3 HPS with this item');
  });
});
