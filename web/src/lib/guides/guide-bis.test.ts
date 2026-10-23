// web/src/lib/guides/guide-bis.test.ts
import activeBuild from '../../data/active-build.json';
import { describe, expect, it } from 'vitest';
import { bisCopy } from '../bis/copy';
import { loadBisFile, normaliseBisFile } from '../bis/load';
import type { ItemTooltipModel } from '../items/tooltip';
import type { LootFile } from '../sim/loot';
import type { BisBand, BisFile, BisSlot, LootQuestsFile } from '../bis/types';
import { guideBisFrom, loadGuideBis } from './guide-bis';

const BUILD = activeBuild.build;

function pick(overrides: Partial<BisSlot>): BisSlot {
  return {
    slot: 'head',
    item_id: 1,
    item_name: 'Test Crown',
    source: 'x',
    source_kind: 'vendor',
    verified: true,
    ...overrides,
  };
}

function band(level: number, slots: BisSlot[]): BisBand {
  return {
    spec: 'mage-frost',
    band: level,
    faction: 'alliance',
    race: 'gnome',
    talents: '',
    talent_points: 1,
    weights: [],
    slots,
    set_dps: 100,
    no_source_count: 0,
    coverage: {},
    new_at_band: [],
    weights_run_seconds: 0,
    verify_run_seconds: 0,
  };
}

function file(...bands: BisBand[]): BisFile {
  return normaliseBisFile({
    spec: 'mage-frost',
    build: BUILD,
    engine_version: 'test',
    generated_at: '2026-10-22T00:00:00Z',
    bands,
  });
}

const LOOT = {
  sources: [
    {
      id: 'raid:test',
      kind: 'raid',
      name: 'Test Raid',
      bosses: [{ id: 'raid:test:1', name: 'Boss', items: [2], item_chances: { '2': 12 } }],
    },
    { id: 'vendor:1', kind: 'vendor', name: 'Mokvar', items: [1] },
  ],
} as unknown as LootFile & Partial<LootQuestsFile>;

function model(id: number, quality: number, icon: string, typeLabel?: string): ItemTooltipModel {
  return {
    id,
    quality,
    icon,
    typeLabel,
    stats: [],
    itemLevel: 10,
    requiredLevel: 1,
  } as unknown as ItemTooltipModel;
}

const DEPS = {
  itemDetails: new Map(),
  loot: LOOT,
  tooltipFor: (id: number) =>
    id === 1
      ? model(1, 4, 'inv_crown', 'Cloth')
      : id === 2
        ? model(2, 3, 'inv_misc_questionmark')
        : undefined,
  spec: 'mage-frost',
};

describe('guideBisFrom', () => {
  const guide = guideBisFrom(
    file(
      band(20, [pick({ slot: 'head', item_id: 1 })]),
      band(60, [
        pick({ slot: 'head', item_id: 1 }),
        pick({
          slot: 'neck',
          item_id: 2,
          item_name: 'Raid Amulet',
          source_kind: 'raid',
          sim_status: 'not_in_sim',
        }),
      ]),
    ),
    'alliance',
    DEPS,
  )!;

  it('opens on the top band and names the race the band was ranked for', () => {
    expect(guide.defaultBand).toBe(60);
    expect(guide.faction).toBe('alliance');
    expect(guide.raceName).toBe('Gnome');
    expect(guide.bands.map((b) => b.label)).toEqual(['20 to 29', '60']);
  });

  it('gives every wearable slot a row: the item, its quality, its icon and its source', () => {
    const top = guide.bands.find((b) => b.band === 60)!;
    expect(top.rows).toHaveLength(17);
    const head = top.rows.find((row) => row.slot === 'head')!;
    expect(head).toMatchObject({
      slotLabel: 'Head',
      itemName: 'Test Crown',
      quality: 4,
      icon: 'inv_crown',
      sourceKind: 'vendor',
      notSimChecked: false,
    });
    expect(head.sourceDetail).toContain('Mokvar');
  });

  it('flags a pick the ranker did not sim-check, and drops the red question-mark icon', () => {
    const neck = guide.bands.find((b) => b.band === 60)!.rows.find((row) => row.slot === 'neck')!;
    expect(neck.notSimChecked).toBe(true);
    expect(neck.icon).toBeUndefined();
    expect(neck.quality).toBe(3);
  });

  it('says why a slot is empty, in the BiS page’s words', () => {
    const empty = guide.bands.find((b) => b.band === 60)!.rows.find((row) => row.slot === 'finger1')!;
    expect(empty.empty).toBe(true);
    expect(empty.emptyCopy).toBeTruthy();
  });

  it('anchors each band on the BiS page', () => {
    expect(guide.bands.map((b) => b.bisAnchor)).toEqual(['band-alliance-20', 'band-alliance-60']);
  });

  it('is the pre-raid list only while no level 60 pick comes from a raid', () => {
    const noRaid = guideBisFrom(file(band(60, [pick({ slot: 'head', item_id: 1 })])), 'alliance', DEPS)!;
    expect(noRaid.bands[0]!.preRaid).toBe(true);
    expect(guide.bands.find((b) => b.band === 60)!.preRaid).toBe(false);
    expect(guide.bands.find((b) => b.band === 20)!.preRaid).toBe(false);
  });

  it('has nothing for a faction the file has no band for', () => {
    expect(guideBisFrom(file(band(60, [])), 'horde', DEPS)).toBeUndefined();
  });
});

describe('loadGuideBis for Frost Mage, from the real file', () => {
  const guide = loadGuideBis(BUILD, 'mage', 'mage-frost', ['gnome', 'troll']);

  it('has the five level ranges, opening on 60, for the faction of the first recommended race', () => {
    if (guide === undefined) throw new Error('mage-frost has no ranked file');
    expect(guide.faction).toBe('alliance');
    expect(guide.defaultBand).toBe(60);
    expect(guide.bands.map((b) => b.label)).toEqual(['20 to 29', '30 to 39', '40 to 49', '50 to 59', '60']);
    expect(guide.bands.every((b) => b.rows.length === 17)).toBe(true);
  });

  it('names the same item in each slot as the BiS page does', () => {
    if (guide === undefined) throw new Error('mage-frost has no ranked file');
    const bis = loadBisFile('mage-frost', BUILD)!;
    const raid = bis.bands.find((b) => b.band === 60 && b.faction === 'alliance' && b.preset === 'raid')!;
    const top = guide.bands.find((b) => b.band === 60)!;
    for (const slot of raid.slots.filter((s) => s.item_name !== undefined)) {
      expect(top.rows.find((row) => row.slot === slot.slot)?.itemName, slot.slot).toBe(slot.item_name);
    }
  });

  it('has no ranked table for a spec without a published file', () => {
    expect(loadGuideBis(BUILD, 'mage', 'mage-nonexistent', ['gnome'])).toBeUndefined();
  });

  it('reads the band label from the shared copy', () => {
    expect(bisCopy.bandRangeLabel(60)).toBe('60');
  });
});
