// web/src/lib/bis/source-cell.test.ts
import { describe, expect, it } from 'vitest';
import type { LootFile } from '../sim/loot';
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
      trash: [20],
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
      ],
    },
    {
      id: 'world:some-creature',
      kind: 'world',
      name: 'Some Creature',
      items: [81],
      item_chances: { '81': 6 },
    },
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
});

describe('describeSourceCell', () => {
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
    expect(describeSourceCell({ kind: 'unknown', label: 'Vendors' })).toBe('Vendors');
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
