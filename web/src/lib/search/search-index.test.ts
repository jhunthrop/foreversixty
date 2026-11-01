import { describe, expect, it } from 'vitest';
import { ITEM_USES_SHARDS, buildSearchIndex, encodeSearchIndex } from './search-index';
import type { TalentFile } from '../planner/types';
import type { SpecCatalogEntry } from '../bis/types';

const cape = {
  id: 11623,
  name: 'Spritecaster Cape',
  icon: 'inv_misc_cape_20',
  quality: 3,
  slot: 'back',
  item_level: 52,
  required_level: 47,
};
const helm = {
  id: 16866,
  name: 'Helm of Wrath',
  icon: 'inv_helmet_71',
  quality: 4,
  slot: 'head',
  item_level: 76,
  required_level: 60,
};

function spec(overrides: Partial<SpecCatalogEntry & { icon: string }>): SpecCatalogEntry & { icon: string } {
  return {
    spec: 'shaman-elemental',
    class_slug: 'shaman',
    spec_slug: 'elemental',
    name: 'Elemental',
    role: 'dps',
    tree_index: 0,
    reference_stat: 'spell_power',
    weight_stats: [],
    icon: 'spell_nature_lightning',
    ...overrides,
  };
}

const shamanTalents: TalentFile = {
  build: 'b1',
  class_id: 7,
  class_slug: 'shaman',
  trees: [
    {
      id: 261,
      name: 'Elemental',
      position: 0,
      background: 'x',
      talents: [
        {
          id: 1,
          name: 'Convection',
          icon: 'spell_nature_wispsplode',
          max_rank: 5,
          tier: 0,
          column: 1,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
          spell_id: 1,
        },
      ],
    },
    {
      id: 263,
      name: 'Enhancement',
      position: 1,
      background: 'x',
      talents: [
        {
          id: 2,
          name: 'Ancestral Knowledge',
          icon: 'spell_shadow_grimward',
          max_rank: 5,
          tier: 0,
          column: 1,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
          spell_id: 2,
        },
      ],
    },
  ],
};

describe('buildSearchIndex', () => {
  const index = buildSearchIndex({
    build: 'b1',
    classItems: [
      { classSlug: 'mage', items: [cape] },
      { classSlug: 'shaman', items: [cape, helm] },
      { classSlug: 'warrior', items: [helm] },
    ],
    talents: [shamanTalents],
    specs: [
      spec({}),
      spec({
        spec: 'shaman-restoration',
        spec_slug: 'restoration',
        name: 'Restoration',
        role: 'healer',
        tree_index: 2,
      }),
    ],
    guidePaths: new Set(['shaman/elemental']),
  });

  it('carries each item once with every class whose file lists it', () => {
    expect(index.items).toEqual([
      { ...cape, classes: ['mage', 'shaman'] },
      { ...helm, classes: ['shaman', 'warrior'] },
    ]);
  });

  it('names a talent by its tree and the spec the tree is played as', () => {
    expect(index.talents).toEqual([
      {
        id: 1,
        name: 'Convection',
        icon: 'spell_nature_wispsplode',
        class: 'shaman',
        tree: 'Elemental',
        spec_name: 'Elemental',
        max_rank: 5,
      },
      // No written spec plays the Enhancement tree here: the tree's own name stands in.
      {
        id: 2,
        name: 'Ancestral Knowledge',
        icon: 'spell_shadow_grimward',
        class: 'shaman',
        tree: 'Enhancement',
        spec_name: 'Enhancement',
        max_rank: 5,
      },
    ]);
  });

  it('lists only the specs with a written guide', () => {
    expect(index.guides).toEqual([
      {
        class: 'shaman',
        spec: 'elemental',
        spec_name: 'Elemental',
        role: 'dps',
        icon: 'spell_nature_lightning',
      },
    ]);
  });

  it('publishes the shard count the item-uses files use', () => {
    expect(index.item_uses_shards).toBe(ITEM_USES_SHARDS);
  });
});

describe('encodeSearchIndex', () => {
  const index = buildSearchIndex({
    build: 'b1',
    classItems: [],
    talents: [],
    specs: [],
    guidePaths: new Set(),
  });

  it('writes the index as JSON under the ceiling', () => {
    expect(JSON.parse(encodeSearchIndex(index))).toEqual(index);
  });

  it('fails the build over the gzipped ceiling', () => {
    expect(() => encodeSearchIndex(index, 10)).toThrow(/over the 10-byte ceiling/);
  });
});
