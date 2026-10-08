// web/src/lib/bis/load.test.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { SLOTS } from '../planner/types';
import {
  bandEntry,
  bandLevels,
  changedSinceBand,
  filledSlots,
  fixtureSpecs,
  groupByClass,
  isEmptySlotRow,
  isMissingSlot,
  itemDetails,
  itemHoverModel,
  itemQualities,
  loadBisFile,
  loadLootFile,
  previousBandLevel,
  readSpecCatalog,
  sourceBadgeLabel,
  normaliseBisFile,
} from './load';
import type { BisBand, BisFile, BisSlot, SpecCatalogEntry } from './types';

const FIXTURE = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../data/fixtures/bis/hunter-marksmanship.json',
);

describe('readSpecCatalog', () => {
  it('reads the master written-spec list, hunter-marksmanship among them', () => {
    const catalog = readSpecCatalog();
    expect(catalog.length).toBeGreaterThan(0);
    expect(catalog.map((entry) => entry.spec)).toContain('hunter-marksmanship');
  });
});

describe('groupByClass', () => {
  it('groups specs by class, class order first-seen, specs in file order', () => {
    const catalog: SpecCatalogEntry[] = [
      {
        spec: 'druid-balance',
        class_slug: 'druid',
        spec_slug: 'balance',
        name: 'Balance',
        role: 'dps',
        tree_index: 0,
        reference_stat: 'spell_power',
        weight_stats: [],
      },
      {
        spec: 'hunter-marksmanship',
        class_slug: 'hunter',
        spec_slug: 'marksmanship',
        name: 'Marksmanship',
        role: 'dps',
        tree_index: 1,
        reference_stat: 'attack_power',
        weight_stats: [],
      },
      {
        spec: 'druid-feral',
        class_slug: 'druid',
        spec_slug: 'feral',
        name: 'Feral',
        role: 'dps',
        tree_index: 1,
        reference_stat: 'attack_power',
        weight_stats: [],
      },
    ];
    const grouped = groupByClass(catalog);
    expect(grouped.map((g) => g.classSlug)).toEqual(['druid', 'hunter']);
    expect(grouped[0].specs.map((s) => s.spec)).toEqual(['druid-balance', 'druid-feral']);
    expect(grouped[1].specs.map((s) => s.spec)).toEqual(['hunter-marksmanship']);
  });
});

describe('fixtureSpecs', () => {
  it('names the committed fixture, hunter-marksmanship', () => {
    expect(fixtureSpecs()).toContain('hunter-marksmanship');
  });
});

describe('loadBisFile', () => {
  it('falls back to the committed fixture when the real build has no bis/ directory yet', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291');
    expect(file).not.toBeNull();
    expect(file?.spec).toBe('hunter-marksmanship');
    expect(file?.bands.length).toBeGreaterThan(0);
  });

  it('returns null for a spec with neither a real file nor a fixture', () => {
    expect(loadBisFile('nosuch-spec', '1.60.1.70291')).toBeNull();
  });

  it('reads an empty "new at this band" list as an array, never null', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    for (const band of file.bands) expect(Array.isArray(band.new_at_band)).toBe(true);
  });

  it('reads coverage as an object on every band, never undefined', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    for (const band of file.bands) expect(typeof band.coverage).toBe('object');
  });
});

const baseBand: BisBand = {
  spec: 'hunter-marksmanship',
  band: 20,
  faction: 'horde',
  race: 'troll',
  talents: '',
  talent_points: 0,
  weights: [],
  slots: [],
  set_dps: 0,
  no_source_count: 0,
  new_at_band: [],
  weights_run_seconds: 0,
  verify_run_seconds: 0,
  coverage: {},
};

describe('normaliseBisFile', () => {
  it('carries a healer band role, profile and metrics, and the file heal_profile, through untouched', () => {
    const metrics = { hps: 300, raw_hps: 360, overheal_pct: 0.16, mana_lasts_sec: 192, hpm: 3.3 };
    const healBand = { ...baseBand, role: 'healer' as const, profile: 'onyxia-sized', metrics };
    const profile = { id: 'onyxia-sized', label: 'Onyxia-sized tank hits and raid pulses' };
    const file = { bands: [healBand, baseBand], heal_profile: profile } as unknown as BisFile;
    const normalised = normaliseBisFile(file);
    expect(normalised.heal_profile).toEqual(profile);
    expect(normalised.bands[0]).toMatchObject({ role: 'healer', profile: 'onyxia-sized', metrics });
    expect(normalised.bands[1]!.role).toBe('dps');
    expect(normalised.bands[1]!.metrics).toBeNull();
  });

  it('reads hit_to_cap through, and defaults an absent key to null', () => {
    const withKey = { ...baseBand, hit_to_cap: { baseline: 3, specials: 6, white: 25 } };
    const file = { bands: [withKey, baseBand] } as unknown as BisFile;
    const [keyed, bare] = normaliseBisFile(file).bands;
    expect(keyed.hit_to_cap).toEqual({ baseline: 3, specials: 6, white: 25 });
    expect(bare.hit_to_cap).toBeNull();
  });

  it('defaults a missing coverage field to {} (a file published before guardrail A landed)', () => {
    const { coverage: _coverage, ...bandWithoutCoverage } = baseBand;
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [bandWithoutCoverage as unknown as BisBand],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].coverage).toEqual({});
  });

  it('defaults a literal null coverage (a Go nil map) to {}', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, coverage: null as unknown as Record<string, never> }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].coverage).toEqual({});
  });

  it('keeps a real coverage map unchanged', () => {
    const coverage = { head: { eligible: 12, sourced: 3 } };
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, coverage }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].coverage).toEqual(coverage);
  });

  it('defaults a missing reference_dps_per_point to null (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [baseBand],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].reference_dps_per_point).toBeNull();
  });

  it('keeps a real reference_dps_per_point unchanged', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, reference_dps_per_point: 0.0714 }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].reference_dps_per_point).toBe(0.0714);
  });

  it('defaults a missing weights_reason to null (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [baseBand],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].weights_reason).toBeNull();
  });

  it('keeps a real weights_reason unchanged', () => {
    const reason = 'reference stat spell_power measured -0.1893 ± 0.6199 DPS per point';
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, weights_reason: reason }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].weights_reason).toBe(reason);
  });

  it('defaults a missing set_dps_partial to false (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [baseBand],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].set_dps_partial).toBe(false);
  });

  it('keeps a real set_dps_partial true unchanged', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, set_dps_partial: true }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].set_dps_partial).toBe(true);
  });

  it('defaults a missing scale_reference_stat to null (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [baseBand],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].scale_reference_stat).toBeNull();
  });

  it('keeps a real scale_reference_stat unchanged', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, scale_reference_stat: 'agility' }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].scale_reference_stat).toBe('agility');
  });

  it('defaults a missing haste_scale_factor to null (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [baseBand],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].haste_scale_factor).toBeNull();
  });

  it('keeps a real haste_scale_factor unchanged', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, haste_scale_factor: 1.58 }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].haste_scale_factor).toBe(1.58);
  });

  it('defaults a missing haste_on_items to true (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [baseBand],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].haste_on_items).toBe(true);
  });

  it('keeps a real haste_on_items false unchanged', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, haste_on_items: false }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].haste_on_items).toBe(false);
  });

  const slotWithoutAlternatives: BisSlot = {
    slot: 'head',
    item_id: 1,
    item_name: 'Plain Helm',
    source: 'A Quest',
    source_kind: 'quest',
    score: 10,
    verified: true,
  };

  it('defaults a missing slot alternatives field to [] (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, slots: [slotWithoutAlternatives] }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].slots[0].alternatives).toEqual([]);
  });

  it('defaults a literal null slot alternatives (a Go nil slice) to []', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [
        {
          ...baseBand,
          slots: [{ ...slotWithoutAlternatives, alternatives: null as unknown as undefined }],
        },
      ],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].slots[0].alternatives).toEqual([]);
  });

  it('keeps a real slot alternatives list unchanged', () => {
    const alternatives = [
      {
        item_id: 2,
        item_name: 'Runner Up',
        source_kind: 'quest',
        source: 'A Quest',
        dps_delta: -2,
        verified: true,
      },
    ];
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, slots: [{ ...slotWithoutAlternatives, alternatives }] }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].slots[0].alternatives).toEqual(alternatives);
  });

  it('defaults a missing slot dps_delta to null (a file published before this lane landed)', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, slots: [slotWithoutAlternatives] }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].slots[0].dps_delta).toBeNull();
  });

  it('keeps a real slot dps_delta unchanged', () => {
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, slots: [{ ...slotWithoutAlternatives, dps_delta: 4.5 }] }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].slots[0].dps_delta).toBe(4.5);
  });

  // Spec addendum 2 (weight rail rows go per rating point): normaliseBisFile never touches
  // a band's `weights` array itself -- a rating row's new fields (`unit`, `rating_factor`,
  // `weight_per_percent`) and a non-rating row's absence of them both pass straight through,
  // the same discipline every other still-optional weight field (`error`, `insignificant`)
  // already gets here.
  it('keeps a real rating-family weight row’s unit/rating_factor/weight_per_percent unchanged', () => {
    const weights = [
      {
        stat: 'crit',
        weight: 1.44,
        error: 0.07,
        unit: 'rating' as const,
        rating_factor: 14,
        weight_per_percent: 20.16,
      },
    ];
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, weights }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].weights[0]).toEqual(weights[0]);
  });

  it('leaves a non-rating weight row with no unit/rating_factor/weight_per_percent keys at all', () => {
    const weights = [{ stat: 'agility', weight: 2, error: 0.1 }];
    const file = {
      spec: 'hunter-marksmanship',
      build: 'test',
      engine_version: 'test',
      generated_at: 'test',
      bands: [{ ...baseBand, weights }],
    };
    const normalised = normaliseBisFile(file);
    expect(normalised.bands[0].weights[0].unit).toBeUndefined();
    expect(normalised.bands[0].weights[0].rating_factor).toBeUndefined();
    expect(normalised.bands[0].weights[0].weight_per_percent).toBeUndefined();
  });
});

describe('bandLevels', () => {
  it('lists every band once, ascending, deduplicated across factions', () => {
    const file: BisFile = {
      spec: 'hunter-marksmanship',
      build: '1.60.1.70291',
      engine_version: 'e1',
      generated_at: 'test',
      bands: [
        { ...baseBand, band: 40, faction: 'horde' },
        { ...baseBand, band: 20, faction: 'alliance' },
        { ...baseBand, band: 40, faction: 'alliance' },
        { ...baseBand, band: 20, faction: 'horde' },
      ],
    };
    expect(bandLevels(file)).toEqual([20, 40]);
  });
});

describe('bandEntry', () => {
  it('finds the band+faction pair the contract keys bands by', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    const band30Horde = bandEntry(file, 30, 'horde');
    expect(band30Horde?.band).toBe(30);
    expect(band30Horde?.faction).toBe('horde');
    expect(bandEntry(file, 999, 'horde')).toBeUndefined();
  });

  it('gives band 30 real "new at this band" content against band 25 (the e2e fixture case)', () => {
    // The committed fixture, read directly: once the nightly publishes a real file for
    // this spec, loadBisFile prefers it and its band 30 may have nothing new.
    const fixture = normaliseBisFile(JSON.parse(readFileSync(FIXTURE, 'utf8')));
    const band30Horde = bandEntry(fixture, 30, 'horde')!;
    expect(band30Horde.new_at_band.length).toBeGreaterThan(0);
  });
});

describe('filledSlots', () => {
  it('returns all 17 planner slots, in the planner’s own order', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    const band = bandEntry(file, 20, 'alliance')!;
    const rows = filledSlots(band);
    expect(rows.map((row) => row.slot)).toEqual([...SLOTS]);
  });

  it('marks a slot the band has no pick for as missing, rather than dropping the row', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    const band = bandEntry(file, 20, 'alliance')!;
    const thin = { ...band, slots: band.slots.filter((slot) => slot.slot !== 'ranged') };
    const rows = filledSlots(thin);
    const ranged = rows.find((row) => row.slot === 'ranged')!;
    expect(isMissingSlot(ranged)).toBe(true);
  });
});

describe('itemQualities', () => {
  it('maps the real build’s own hunter items by id to quality', () => {
    const qualities = itemQualities('1.60.1.70291', 'hunter');
    expect(qualities.size).toBeGreaterThan(0);
  });

  it('returns an empty map for a class the build has no item file for', () => {
    expect(itemQualities('1.60.1.70291', 'not-a-class').size).toBe(0);
  });

  it('returns an empty map for a build that does not exist, rather than throwing', () => {
    expect(itemQualities('0.0.0.0', 'hunter').size).toBe(0);
  });
});

describe('sourceBadgeLabel', () => {
  function slot(source_kind: string): BisSlot {
    return {
      slot: 'head',
      item_id: 1,
      item_name: 'Test Item',
      source: 'Test source',
      source_kind,
      score: 1,
      verified: true,
    };
  }

  it('names a quest source by the band’s own faction, not the generic "Quests" label', () => {
    expect(sourceBadgeLabel(slot('quest'), 'alliance')).toBe('Alliance quest');
    expect(sourceBadgeLabel(slot('quest'), 'horde')).toBe('Horde quest');
  });

  it('uses the picker’s own badge word for every other kind', () => {
    expect(sourceBadgeLabel(slot('vendor'), 'alliance')).toBe('Vendors');
    expect(sourceBadgeLabel(slot('zone'), 'alliance')).toBe('Zone drops');
    expect(sourceBadgeLabel(slot('dungeon'), 'alliance')).toBe('Dungeons');
    expect(sourceBadgeLabel(slot('world_drop'), 'alliance')).toBe('World drops');
  });

  it('falls back to the raw kind for one loot.ts does not know, rather than throwing', () => {
    expect(sourceBadgeLabel(slot('mystery'), 'alliance')).toBe('mystery');
  });
});

describe('isEmptySlotRow', () => {
  it('is true for a slot absent from the band’s own array (isMissingSlot)', () => {
    expect(isEmptySlotRow({ slot: 'ranged', missing: true })).toBe(true);
  });

  it('is true for a slot present but carrying no pick, the real pipeline’s own empty shape', () => {
    expect(isEmptySlotRow({ slot: 'off_hand', verified: false } as BisSlot)).toBe(true);
  });

  it('is false for a slot with a real pick', () => {
    const row: BisSlot = {
      slot: 'head',
      item_id: 1,
      item_name: 'Cap',
      source: 'Vendor: Someone',
      source_kind: 'vendor',
      score: 1,
      verified: true,
    };
    expect(isEmptySlotRow(row)).toBe(false);
  });
});

describe('itemDetails', () => {
  it('reads name, quality, item level, required level and icon off the real build’s item file', () => {
    const details = itemDetails('1.60.1.70291', 'hunter');
    expect(details.size).toBeGreaterThan(0);
    const [, detail] = [...details][0];
    expect(typeof detail.name).toBe('string');
    expect(typeof detail.quality).toBe('number');
    expect(typeof detail.item_level).toBe('number');
    expect(typeof detail.required_level).toBe('number');
    expect(typeof detail.icon).toBe('string');
  });

  it('returns an empty map for a build that does not exist, rather than throwing', () => {
    expect(itemDetails('0.0.0.0', 'hunter').size).toBe(0);
  });
});

describe('itemHoverModel', () => {
  it('takes the name from the BiS row (the caller’s own itemName), everything else from itemDetails', () => {
    const details = itemDetails('1.60.1.70291', 'hunter');
    const [id, detail] = [...details][0];
    const model = itemHoverModel(id, 'The BiS file’s own item name', details);
    expect(model).toEqual({
      name: 'The BiS file’s own item name',
      quality: detail.quality,
      itemLevel: detail.item_level,
      requiredLevel: detail.required_level,
      icon: detail.icon,
      stats: detail.stats,
    });
  });

  it('falls back to the BiS file’s own name for an id itemDetails does not know', () => {
    const model = itemHoverModel(999999999, 'Ye Olde Item', new Map());
    expect(model).toEqual({ name: 'Ye Olde Item', quality: 1, itemLevel: 0, requiredLevel: 0 });
  });
});

describe('loadLootFile', () => {
  it('reads the real build’s sources and quest map', () => {
    const loot = loadLootFile('1.60.1.70291');
    expect(loot.sources.length).toBeGreaterThan(0);
    expect(Object.keys(loot.quests).length).toBeGreaterThan(0);
  });

  it('returns an empty file for a build that does not exist, rather than throwing', () => {
    const loot = loadLootFile('0.0.0.0');
    expect(loot).toEqual({ sources: [], quests: {} });
  });
});

describe('previousBandLevel', () => {
  it('is undefined at the file’s first band', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    expect(previousBandLevel(file, 20)).toBeUndefined();
  });

  it('is the band immediately before, per bandLevels', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    expect(previousBandLevel(file, 30)).toBe(20);
  });

  it('is undefined for a band the file does not carry', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    expect(previousBandLevel(file, 999)).toBeUndefined();
  });
});

describe('changedSinceBand', () => {
  it('is undefined at the file’s first band', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    expect(changedSinceBand(file, 20, 'alliance')).toBeUndefined();
  });

  it('lists only the slots whose item id changed from the previous band, faction held constant', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    const changed = changedSinceBand(file, 30, 'horde')!;
    expect(changed.length).toBeGreaterThan(0);
    for (const entry of changed) {
      expect(entry.before?.item_id).not.toBe(entry.after?.item_id);
    }
  });

  it('does not list a slot whose pick is unchanged between bands', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70291')!;
    const changed = changedSinceBand(file, 30, 'horde')!;
    const trinket1Before = bandEntry(file, 20, 'horde')?.slots.find((s) => s.slot === 'trinket1');
    const trinket1After = bandEntry(file, 30, 'horde')?.slots.find((s) => s.slot === 'trinket1');
    if (trinket1Before?.item_id === trinket1After?.item_id) {
      expect(changed.some((entry) => entry.slot === 'trinket1')).toBe(false);
    }
  });
});
