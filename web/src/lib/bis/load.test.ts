// web/src/lib/bis/load.test.ts
import { describe, expect, it } from 'vitest';
import { SLOTS } from '../planner/types';
import {
  bandEntry,
  bandLevels,
  filledSlots,
  fixtureSpecs,
  groupByClass,
  isMissingSlot,
  itemQualities,
  loadBisFile,
  readSpecCatalog,
  sourceBadgeLabel,
} from './load';
import type { BisSlot, SpecCatalogEntry } from './types';

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
      { spec: 'druid-balance', class_slug: 'druid', spec_slug: 'balance', name: 'Balance', role: 'dps', tree_index: 0, reference_stat: 'spell_power', weight_stats: [] },
      { spec: 'hunter-marksmanship', class_slug: 'hunter', spec_slug: 'marksmanship', name: 'Marksmanship', role: 'dps', tree_index: 1, reference_stat: 'attack_power', weight_stats: [] },
      { spec: 'druid-feral', class_slug: 'druid', spec_slug: 'feral', name: 'Feral', role: 'dps', tree_index: 1, reference_stat: 'attack_power', weight_stats: [] },
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
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70009');
    expect(file).not.toBeNull();
    expect(file?.spec).toBe('hunter-marksmanship');
    expect(file?.bands.length).toBeGreaterThan(0);
  });

  it('returns null for a spec with neither a real file nor a fixture', () => {
    expect(loadBisFile('druid-balance', '1.60.1.70009')).toBeNull();
  });
});

describe('bandLevels', () => {
  it('lists every band once, ascending, deduplicated across factions', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70009')!;
    const levels = bandLevels(file);
    expect(levels).toEqual([10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60]);
  });
});

describe('bandEntry', () => {
  it('finds the band+faction pair the contract keys bands by', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70009')!;
    const band30Horde = bandEntry(file, 30, 'horde');
    expect(band30Horde?.band).toBe(30);
    expect(band30Horde?.faction).toBe('horde');
    expect(bandEntry(file, 999, 'horde')).toBeUndefined();
  });

  it('gives band 30 real "new at this band" content against band 25 (the e2e fixture case)', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70009')!;
    const band30Horde = bandEntry(file, 30, 'horde')!;
    expect(band30Horde.new_at_band.length).toBeGreaterThan(0);
  });
});

describe('filledSlots', () => {
  it('returns all 17 planner slots, in the planner’s own order', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70009')!;
    const band = bandEntry(file, 10, 'alliance')!;
    const rows = filledSlots(band);
    expect(rows.map((row) => row.slot)).toEqual([...SLOTS]);
  });

  it('marks a slot the band has no pick for as missing, rather than dropping the row', () => {
    const file = loadBisFile('hunter-marksmanship', '1.60.1.70009')!;
    const band = bandEntry(file, 10, 'alliance')!;
    const thin = { ...band, slots: band.slots.filter((slot) => slot.slot !== 'ranged') };
    const rows = filledSlots(thin);
    const ranged = rows.find((row) => row.slot === 'ranged')!;
    expect(isMissingSlot(ranged)).toBe(true);
  });
});

describe('itemQualities', () => {
  it('maps the real build’s own hunter items by id to quality', () => {
    const qualities = itemQualities('1.60.1.70009', 'hunter');
    expect(qualities.size).toBeGreaterThan(0);
  });

  it('returns an empty map for a class the build has no item file for', () => {
    expect(itemQualities('1.60.1.70009', 'not-a-class').size).toBe(0);
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
  });

  it('falls back to the raw kind for one loot.ts does not know, rather than throwing', () => {
    expect(sourceBadgeLabel(slot('mystery'), 'alliance')).toBe('mystery');
  });
});
