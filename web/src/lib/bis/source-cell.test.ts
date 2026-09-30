// web/src/lib/bis/source-cell.test.ts
import { describe, expect, it } from 'vitest';
import type { LootFile } from '../sim/loot';
import { bisCopy } from './copy';
import { describeSourceCell, hasKnownSource, resolveSourceCell } from './source-cell';
import type { BisSlot, LootQuestsFile } from './types';

function slot(overrides: Partial<BisSlot> = {}): BisSlot {
  return {
    slot: 'head',
    item_id: 1,
    item_name: 'Test Item',
    source: 'Test source',
    source_kind: 'vendor',
    score: 1,
    verified: true,
    ...overrides,
  };
}

const LOOT: LootFile & LootQuestsFile = {
  sources: [
    {
      id: 'raid:molten-core',
      kind: 'raid',
      name: 'Molten Core',
      zone_id: 1,
      bosses: [{ id: 'raid:molten-core:1', name: 'Ragnaros', npc_id: 1, items: [10] }],
      trash: [11],
    },
    {
      id: 'dungeon:deadmines',
      kind: 'dungeon',
      name: 'The Deadmines',
      zone_id: 2,
      bosses: [],
      // item 82 also appears (with a real chance) as a boss drop in
      // `dungeon:the-deadmines` below -- see the
      // `findBestChanceSource`-covering tests: this trash-only, no-chance
      // entry must never win just because it comes first in the array.
      trash: [20, 82],
    },
    { id: 'crafted:tailoring', kind: 'crafted', name: 'Tailoring', profession: 'tailoring', items: [30] },
    { id: 'vendor:123', kind: 'vendor', name: 'Gorn One Eye', items: [40] },
    {
      id: 'rep:timbermaw:friendly',
      kind: 'rep',
      name: 'Timbermaw Hold',
      faction_id: 1,
      standing: 'friendly',
      items: [50],
    },
    { id: 'zone:1', kind: 'zone', name: 'Dun Morogh', zone_id: 1, items: [60] },
    {
      id: 'dungeon:the-deadmines',
      kind: 'dungeon',
      name: 'The Deadmines',
      zone_id: 3,
      bosses: [
        {
          id: 'dungeon:the-deadmines:1',
          name: 'Lord Serpentis',
          npc_id: 1,
          items: [80],
          item_chances: { '80': 40 },
        },
        {
          id: 'dungeon:the-deadmines:2',
          name: 'Mr. Smite',
          npc_id: 2,
          items: [82],
          item_chances: { '82': 15 },
        },
      ],
    },
    {
      id: 'world:some-creature',
      kind: 'world',
      name: 'Some Creature',
      items: [81],
      item_chances: { '81': 6 },
    },
    {
      id: 'world_drop:18-25',
      kind: 'world_drop',
      name: 'World drop',
      items: [90],
      level_min: 18,
      level_max: 25,
    },
    { id: 'world_drop:unknown', kind: 'world_drop', name: 'World drop', items: [91] },
    {
      id: 'pvp:rank-11:alliance',
      kind: 'pvp',
      name: 'Rank 11 (Alliance)',
      rank: 11,
      faction: 'alliance',
      items: [100],
    },
    { id: 'pvp:rank-9', kind: 'pvp', name: 'Rank 9', items: [101] },
  ],
  quests: {
    '70': [
      { quest_id: 1, name: 'A Faction Quest', faction: 'alliance', min_level: 10, level: 12 },
      { quest_id: 2, name: 'A Horde Quest', faction: 'horde', min_level: 10, level: 12 },
    ],
    '71': [{ quest_id: 3, name: 'A Neutral Quest', faction: 'both', min_level: 5, level: 8 }],
  },
};

describe('hasKnownSource', () => {
  it('is true for a slot with an item id', () => {
    expect(hasKnownSource(slot())).toBe(true);
  });

  it('is false for the pipeline’s own "no pick" shape', () => {
    expect(hasKnownSource({ slot: 'off_hand', verified: false } as BisSlot)).toBe(false);
  });
});

describe('resolveSourceCell', () => {
  it('resolves a quest to its band faction, name and level, not the quest’s own faction', () => {
    const cell = resolveSourceCell(slot({ item_id: 70, source_kind: 'quest' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'quest', questName: 'A Faction Quest', faction: 'alliance', level: 12 });
  });

  it('resolves a neutral ("both") quest under either faction panel using that panel’s faction', () => {
    const alliance = resolveSourceCell(
      slot({ item_id: 71, source_kind: 'quest' }),
      'alliance',
      LOOT,
      'fallback',
    );
    const horde = resolveSourceCell(slot({ item_id: 71, source_kind: 'quest' }), 'horde', LOOT, 'fallback');
    expect(alliance).toEqual({ kind: 'quest', questName: 'A Neutral Quest', faction: 'alliance', level: 8 });
    expect(horde).toEqual({ kind: 'quest', questName: 'A Neutral Quest', faction: 'horde', level: 8 });
  });

  it('falls back when loot.json has no quest entry for the item', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 999, source_kind: 'quest' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({ kind: 'unknown', label: 'fallback' });
  });

  it('resolves a raid drop to its instance and the boss that drops it', () => {
    const cell = resolveSourceCell(slot({ item_id: 10, source_kind: 'raid' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'raid', instance: 'Molten Core', boss: 'Ragnaros' });
  });

  it('resolves a raid trash item to its instance with no boss', () => {
    const cell = resolveSourceCell(slot({ item_id: 11, source_kind: 'raid' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'raid', instance: 'Molten Core', boss: undefined });
  });

  it('resolves a dungeon drop to its instance', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 20, source_kind: 'dungeon' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({ kind: 'dungeon', instance: 'The Deadmines', boss: undefined });
  });

  it('picks the higher-chance boss source when the same item names more than one dungeon source, never the first array match', () => {
    // Item 82 is a chance-less trash entry in `dungeon:deadmines` (which
    // comes FIRST in LOOT.sources) and a real, 15%-chance boss drop in
    // `dungeon:the-deadmines` (which comes after). findSource's plain
    // "first match" would have picked the trash-only source and reported
    // no boss and no chance at all -- wowhead-world-drops lane,
    // 2026-09-29, tenet 7's "never show an arbitrary trash mob when a
    // boss ... exists".
    const cell = resolveSourceCell(
      slot({ item_id: 82, source_kind: 'dungeon' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({
      kind: 'dungeon',
      instance: 'The Deadmines',
      boss: 'Mr. Smite',
      dropChance: 15,
    });
  });

  it('resolves a crafted item to its profession', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 30, source_kind: 'crafted' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({ kind: 'crafted', profession: 'Tailoring' });
  });

  it('resolves a vendor item to the npc', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 40, source_kind: 'vendor' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({ kind: 'vendor', npc: 'Gorn One Eye' });
  });

  it('resolves a rep item to the faction and standing', () => {
    const cell = resolveSourceCell(slot({ item_id: 50, source_kind: 'rep' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'rep', faction: 'Timbermaw Hold', standing: 'friendly' });
  });

  it('resolves a zone item to the zone', () => {
    const cell = resolveSourceCell(slot({ item_id: 60, source_kind: 'zone' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'zone', place: 'Dun Morogh' });
  });

  it('resolves a pvp item to its rank and faction, never the bare source name', () => {
    const cell = resolveSourceCell(slot({ item_id: 100, source_kind: 'pvp' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'pvp', rank: 11, faction: 'alliance' });
  });

  it('falls back for a pvp source missing rank or faction (should not happen on real data)', () => {
    const cell = resolveSourceCell(slot({ item_id: 101, source_kind: 'pvp' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'unknown', label: 'fallback' });
  });

  it('falls back for a kind this resolver does not special-case', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 1, source_kind: 'mystery' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({ kind: 'unknown', label: 'fallback' });
  });

  it('never throws for an item loot.json does not carry under that kind', () => {
    const cell = resolveSourceCell(slot({ item_id: 1, source_kind: 'raid' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'unknown', label: 'fallback' });
  });

  it('carries a boss drop’s own classic-db chance through as dropChance', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 80, source_kind: 'dungeon' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({
      kind: 'dungeon',
      instance: 'The Deadmines',
      boss: 'Lord Serpentis',
      dropChance: 40,
    });
  });

  it('carries a world source’s own classic-db chance through as dropChance', () => {
    const cell = resolveSourceCell(slot({ item_id: 81, source_kind: 'world' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'world', place: 'Some Creature', dropChance: 6 });
  });

  it('leaves dropChance undefined for a source with no classic-db chance data', () => {
    const cell = resolveSourceCell(slot({ item_id: 10, source_kind: 'raid' }), 'alliance', LOOT, 'fallback');
    expect(cell).toEqual({ kind: 'raid', instance: 'Molten Core', boss: 'Ragnaros' });
  });

  it('resolves a world_drop item to its own level range', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 90, source_kind: 'world_drop' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({ kind: 'world_drop', levelMin: 18, levelMax: 25 });
  });

  it('resolves a world_drop item with no known level range to undefined levels', () => {
    const cell = resolveSourceCell(
      slot({ item_id: 91, source_kind: 'world_drop' }),
      'alliance',
      LOOT,
      'fallback',
    );
    expect(cell).toEqual({ kind: 'world_drop', levelMin: undefined, levelMax: undefined });
  });
});

describe('describeSourceCell', () => {
  it('rounds a drop chance to a whole percent and says "<1%" below one', () => {
    expect(bisCopy.dropChanceLabel(1.6358, 'Blackwing Spellbinder')).toBe('2% from Blackwing Spellbinder');
    expect(bisCopy.dropChanceLabel(0.02, 'Deadmines trash')).toBe('<1% from Deadmines trash');
    expect(bisCopy.dropChanceLabel(20, 'Mr. Smite')).toBe('20% from Mr. Smite');
  });

  it('formats every kind as the one line the row shows', () => {
    expect(describeSourceCell({ kind: 'quest', questName: 'Foo', faction: 'alliance', level: 12 })).toBe(
      'Quest: Foo · Level 12',
    );
    expect(describeSourceCell({ kind: 'raid', instance: 'Molten Core', boss: 'Ragnaros' })).toBe(
      'Molten Core · Ragnaros',
    );
    expect(describeSourceCell({ kind: 'dungeon', instance: 'The Deadmines' })).toBe('The Deadmines');
    expect(describeSourceCell({ kind: 'crafted', profession: 'Tailoring' })).toBe('Crafted: Tailoring');
    expect(describeSourceCell({ kind: 'vendor', npc: 'Gorn One Eye' })).toBe('Vendor: Gorn One Eye');
    expect(describeSourceCell({ kind: 'rep', faction: 'Timbermaw Hold', standing: 'friendly' })).toBe(
      'Timbermaw Hold (friendly)',
    );
    expect(describeSourceCell({ kind: 'zone', place: 'Dun Morogh' })).toBe('Dun Morogh');
    expect(describeSourceCell({ kind: 'pvp', rank: 11, faction: 'alliance' })).toBe(
      'PvP rank 11 · Knight-Lieutenant · Alliance',
    );
    expect(describeSourceCell({ kind: 'unknown', label: 'Vendors' })).toBe('Vendors');
  });

  it('formats a world_drop cell with and without a known level range', () => {
    expect(describeSourceCell({ kind: 'world_drop', levelMin: 18, levelMax: 25 })).toBe(
      'World drop (BoE) · levels 18-25',
    );
    expect(describeSourceCell({ kind: 'world_drop' })).toBe('World drop (BoE)');
  });

  it('leads with the classic-db drop chance when the cell carries one', () => {
    expect(
      describeSourceCell({
        kind: 'dungeon',
        instance: 'The Deadmines',
        boss: 'Lord Serpentis',
        dropChance: 40,
      }),
    ).toBe('40% from Lord Serpentis');
    expect(describeSourceCell({ kind: 'world', place: 'Some Creature', dropChance: 6 })).toBe(
      '6% from Some Creature',
    );
  });

  it('falls back to the instance name for a chance on trash (no named boss)', () => {
    expect(describeSourceCell({ kind: 'dungeon', instance: 'The Deadmines', dropChance: 6 })).toBe(
      '6% from The Deadmines trash',
    );
  });
});
