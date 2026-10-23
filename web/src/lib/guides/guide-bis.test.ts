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
  const hordeBand = {
    ...band(60, [pick({ slot: 'head', item_id: 1, item_name: 'Horde Crown' })]),
    faction: 'horde' as const,
    race: 'troll',
  };
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
      hordeBand,
    ),
    'alliance',
    DEPS,
  )!;
  const panel = (faction: string, level: number) =>
    guide.panels.find((p) => p.faction === faction && p.band === level)!;

  it('opens on the faction and the top band, and lists both factions', () => {
    expect(guide.defaultFaction).toBe('alliance');
    expect(guide.factions).toEqual(['alliance', 'horde']);
    expect(guide.defaultBand).toBe(60);
    expect(guide.bands.map((b) => b.label)).toEqual(['20 to 29', '60']);
  });

  it('names the race each faction band was ranked for', () => {
    expect(panel('alliance', 60).raceName).toBe('Gnome');
    expect(panel('horde', 60).raceName).toBe('Troll');
  });

  it('gives every wearable slot a row: the item, its quality, its icon and its source', () => {
    const top = panel('alliance', 60);
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
    expect(panel('horde', 60).rows.find((row) => row.slot === 'head')?.itemName).toBe('Horde Crown');
  });

  it('flags a pick the ranker did not sim-check, and drops the red question-mark icon', () => {
    const neck = panel('alliance', 60).rows.find((row) => row.slot === 'neck')!;
    expect(neck.notSimChecked).toBe(true);
    expect(neck.icon).toBeUndefined();
    expect(neck.quality).toBe(3);
  });

  it('says why a slot is empty, in the BiS page words', () => {
    const empty = panel('alliance', 60).rows.find((row) => row.slot === 'finger1')!;
    expect(empty.empty).toBe(true);
    expect(empty.emptyCopy).toBeTruthy();
  });

  it('anchors each panel on its faction band on the BiS page', () => {
    expect(panel('alliance', 20).bisAnchor).toBe('band-alliance-20');
    expect(panel('horde', 60).bisAnchor).toBe('band-horde-60');
  });

  it('is the pre-raid list only while no level 60 pick comes from a raid', () => {
    const noRaid = guideBisFrom(file(band(60, [pick({ slot: 'head', item_id: 1 })])), 'alliance', DEPS)!;
    expect(noRaid.panels[0]!.preRaid).toBe(true);
    expect(panel('alliance', 60).preRaid).toBe(false);
    expect(panel('alliance', 20).preRaid).toBe(false);
  });

  it('opens on the other faction when the file has no picks for the wanted one', () => {
    expect(guideBisFrom(file(hordeBand), 'alliance', DEPS)?.defaultFaction).toBe('horde');
    expect(guideBisFrom(file(band(60, [])), 'alliance', DEPS)?.factions).toEqual(['alliance']);
  });
});

describe('loadGuideBis for Frost Mage, from the real file', () => {
  const guide = loadGuideBis(BUILD, 'mage', 'mage-frost', ['gnome', 'troll']);

  it('has the five level ranges for both factions, opening on 60 for the first recommended race faction', () => {
    if (guide === undefined) throw new Error('mage-frost has no ranked file');
    expect(guide.defaultFaction).toBe('alliance');
    expect(guide.factions).toEqual(['alliance', 'horde']);
    expect(guide.defaultBand).toBe(60);
    expect(guide.bands.map((b) => b.label)).toEqual(['20 to 29', '30 to 39', '40 to 49', '50 to 59', '60']);
    expect(guide.panels).toHaveLength(10);
    expect(guide.panels.every((p) => p.rows.length === 17)).toBe(true);
  });

  it('names the same item in each slot as the BiS page does, for each faction', () => {
    if (guide === undefined) throw new Error('mage-frost has no ranked file');
    const bis = loadBisFile('mage-frost', BUILD)!;
    for (const faction of ['alliance', 'horde'] as const) {
      const raid = bis.bands.find((b) => b.band === 60 && b.faction === faction && b.preset === 'raid')!;
      const top = guide.panels.find((p) => p.faction === faction && p.band === 60)!;
      for (const slot of raid.slots.filter((s) => s.item_name !== undefined)) {
        expect(top.rows.find((row) => row.slot === slot.slot)?.itemName, `${faction} ${slot.slot}`).toBe(
          slot.item_name,
        );
      }
    }
  });

  it('has no ranked table for a spec without a published file', () => {
    expect(loadGuideBis(BUILD, 'mage', 'mage-nonexistent', ['gnome'])).toBeUndefined();
  });

  it('reads the band label from the shared copy', () => {
    expect(bisCopy.bandRangeLabel(60)).toBe('60');
  });
});
