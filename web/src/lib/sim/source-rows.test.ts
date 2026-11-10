// web/src/lib/sim/source-rows.test.ts
import { describe, expect, it } from 'vitest';
import type { Item } from '../planner/types';
import type { LootSource } from './loot';
import {
  gatedLine,
  groupByFaction,
  pvpForFaction,
  showsSource,
  sourceItemStats,
  tickedSources,
} from './source-rows';

const item = (id: number, itemLevel: number): Item => ({ id, item_level: itemLevel }) as unknown as Item;
const source = (over: Partial<LootSource>): LootSource => ({ id: 'x', kind: 'rep', name: 'X', ...over });

describe('sourceItemStats', () => {
  it('counts the items this class has and takes the best item level among them', () => {
    const items = new Map([
      [1, item(1, 58)],
      [2, item(2, 63)],
    ]);
    expect(sourceItemStats([1, 2, 3], items, null)).toMatchObject({
      count: 2,
      bestIlvl: 63,
      best: { id: 2 },
    });
  });

  it('leaves out an item the engine does not know', () => {
    const items = new Map([[1, item(1, 58)]]);
    expect(sourceItemStats([1], items, new Set([2]))).toEqual({ count: 0, bestIlvl: 0, best: null });
  });
});

describe('groupByFaction', () => {
  it('folds standings under their faction, first seen first', () => {
    const rows = [
      source({ id: 'a', name: 'Timbermaw Hold', standing: 'honored' }),
      source({ id: 'b', name: 'Argent Dawn', standing: 'friendly' }),
      source({ id: 'c', name: 'Timbermaw Hold', standing: 'exalted' }),
    ];
    expect(groupByFaction(rows).map((group) => [group.name, group.standings.map((row) => row.id)])).toEqual([
      ['Timbermaw Hold', ['a', 'c']],
      ['Argent Dawn', ['b']],
    ]);
  });
});

describe('pvpForFaction', () => {
  const ranks = [
    source({ id: 'h', kind: 'pvp', faction: 'horde' }),
    source({ id: 'a', kind: 'pvp', faction: 'alliance' }),
    source({ id: 'n', kind: 'pvp' }),
  ];

  it('keeps the character’s side and any rank with no side', () => {
    expect(pvpForFaction(ranks, 'horde').map((row) => row.id)).toEqual(['h', 'n']);
  });

  it('keeps every rank when the side is unknown', () => {
    expect(pvpForFaction(ranks, null)).toHaveLength(3);
  });
});

describe('gatedLine', () => {
  it('names dated sources with their date and lists the rest once', () => {
    const rows = [source({ id: 'o', name: "Onyxia's Lair" }), source({ id: 'm', name: 'Molten Core' })];
    const gate = (row: LootSource): string => (row.id === 'o' ? 'First raids, 9 December 2026' : 'later');
    expect(gatedLine(rows, gate, 'later')).toEqual({
      dated: ["First raids, 9 December 2026: Onyxia's Lair"],
      undated: ['Molten Core'],
    });
  });
});

describe('showsSource', () => {
  it('hides an empty PvP rank or standing, never a raid or quest', () => {
    expect(showsSource('pvp', 0)).toBe(false);
    expect(showsSource('rep', 0)).toBe(false);
    expect(showsSource('raid', 0)).toBe(true);
    expect(showsSource('quest', 0)).toBe(true);
    expect(showsSource('pvp', 2)).toBe(true);
  });
});

describe('tickedSources', () => {
  it('counts each ticked source once, whole or by boss', () => {
    const rows = [source({ id: 'raid:mc' }), source({ id: 'raid:bwl' }), source({ id: 'raid:ony' })];
    expect(tickedSources(rows, ['raid:mc|', 'raid:mc|raid:mc:1', 'raid:ony|raid:ony:2', 'dungeon:x|'])).toBe(
      2,
    );
  });
});

describe('gatedLine dedupes', () => {
  it('names each source once under its date', () => {
    const rows = [
      source({ id: 'a', name: 'Blacksmithing' }),
      source({ id: 'b', name: 'Blacksmithing' }),
      source({ id: 'c', name: 'Leatherworking' }),
    ];
    expect(gatedLine(rows, () => 'First raids, 9 December 2026', 'later').dated).toEqual([
      'First raids, 9 December 2026: Blacksmithing, Leatherworking',
    ]);
  });
});
