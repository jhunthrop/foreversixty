import { describe, expect, it } from 'vitest';
import { buildItemUses } from './item-uses';
import { ITEM_USES_SHARDS } from './search-index';
import type { BisBand, BisFile, SpecCatalogEntry } from '../bis/types';
import type { LootFile } from '../sim/loot';

const CAPE = 11623;

function catalog(specKey: string, classSlug: string, name: string): SpecCatalogEntry {
  return {
    spec: specKey,
    class_slug: classSlug,
    spec_slug: name.toLowerCase(),
    name,
    role: 'dps',
    tree_index: 0,
    reference_stat: 'x',
    weight_stats: [],
  };
}

function band(
  specKey: string,
  level: number,
  preset: 'bare' | 'raid',
  faction: 'alliance' | 'horde',
  pick: number,
  alternatives: number[] = [],
): BisBand {
  return {
    spec: specKey,
    band: level,
    preset,
    faction,
    slots: [
      {
        slot: 'back',
        item_id: pick,
        item_name: '',
        source: '',
        source_kind: 'dungeon',
        verified: true,
        alternatives: alternatives.map((id) => ({
          item_id: id,
          item_name: '',
          source: '',
          source_kind: 'dungeon',
          dps_delta: -1,
        })),
      },
    ],
  } as unknown as BisBand;
}

function file(specKey: string, bands: BisBand[]): BisFile {
  return { spec: specKey, build: 'b1', engine_version: 'e', generated_at: 'g', bands };
}

const loot: LootFile = {
  sources: [
    {
      id: 'dungeon:blackrock-depths',
      kind: 'dungeon',
      name: 'Blackrock Depths',
      bosses: [
        { id: 'b1', name: 'High Interrogator Gerstahn', items: [CAPE] },
        { id: 'b2', name: 'Houndmaster Grebmar', items: [CAPE], item_chances: { [CAPE]: 12.5 } },
      ],
    },
  ],
} as unknown as LootFile;

const specs = [
  catalog('druid-balance', 'druid', 'Balance'),
  catalog('mage-arcane', 'mage', 'Arcane'),
  catalog('shaman-elemental', 'shaman', 'Elemental'),
];

describe('buildItemUses', () => {
  const shards = buildItemUses({
    build: 'b1',
    itemIds: new Set([CAPE, 1, 2]),
    loot,
    bisFiles: [
      file('shaman-elemental', [
        band('shaman-elemental', 60, 'raid', 'horde', 999, [CAPE]),
        band('shaman-elemental', 50, 'bare', 'horde', CAPE),
      ]),
      file('mage-arcane', [
        band('mage-arcane', 60, 'raid', 'alliance', 999, [5, CAPE]),
        band('mage-arcane', 60, 'raid', 'horde', 999, [5, CAPE]),
      ]),
      file('druid-balance', [
        band('druid-balance', 60, 'raid', 'horde', CAPE),
        band('druid-balance', 60, 'raid', 'alliance', 7),
      ]),
    ],
    specs,
  });
  const cape = shards[CAPE % ITEM_USES_SHARDS].items[String(CAPE)];

  it('writes one file per shard, each naming its own number', () => {
    expect(shards).toHaveLength(ITEM_USES_SHARDS);
    expect(shards.map((shard) => shard.shard)).toEqual([...Array(ITEM_USES_SHARDS).keys()]);
  });

  it("keeps each spec's best placement: a pick beats a runner-up even at a lower band", () => {
    expect(cape.specs).toEqual([
      {
        spec: 'druid-balance',
        class: 'druid',
        spec_name: 'Balance',
        band: 60,
        preset: 'raid',
        faction: 'horde',
        rank: 0,
      },
      {
        spec: 'shaman-elemental',
        class: 'shaman',
        spec_name: 'Elemental',
        band: 50,
        preset: 'bare',
        faction: 'horde',
        rank: 0,
      },
      // Both factions place it second: no faction is named.
      {
        spec: 'mage-arcane',
        class: 'mage',
        spec_name: 'Arcane',
        band: 60,
        preset: 'raid',
        faction: null,
        rank: 2,
      },
    ]);
  });

  it('lists the boss drops, a known chance first, an unknown one without a number', () => {
    expect(cape.sources).toEqual([
      { kind: 'dungeon', title: 'Houndmaster Grebmar', place: 'Blackrock Depths', chance: 12.5 },
      { kind: 'dungeon', title: 'High Interrogator Gerstahn', place: 'Blackrock Depths' },
    ]);
  });

  it('leaves out an item with no source and no list', () => {
    expect(shards[1].items['1']).toBeUndefined();
  });
});
