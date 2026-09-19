// web/src/lib/sim/loot.test.ts
import { describe, expect, it } from 'vitest';
import lootJson from '../../fixtures/planner/loot.json';
import {
  DEFAULT_OFF_KINDS,
  groupSources,
  isOpen,
  itemsOfBoss,
  itemsOfSource,
  professionSplit,
  sourceNameOf,
  sourcesByItem,
  type LootFile,
} from './loot';
import { BUILT_IN_PHASES } from './phase';

const file = lootJson as unknown as LootFile;
const phases = BUILT_IN_PHASES;

describe('sourcesByItem', () => {
  it('lists every source an item drops from, boss and trash alike', () => {
    const index = sourcesByItem(file);
    expect(index.get(16963)).toEqual(['raid:molten-core']);
    expect(index.get(13968)?.sort()).toEqual(
      ['dungeon:hall-of-thanes', 'raid:molten-core', 'rep:argent-dawn:exalted'].sort(),
    );
    expect(index.get(999999)).toBeUndefined();
  });
});

describe('groupSources', () => {
  it('groups by kind in picker order and labels each group', () => {
    const groups = groupSources(file);
    expect(groups.map((group) => group.kind)).toEqual([
      'raid',
      'dungeon',
      'world',
      'crafted',
      'rep',
      'pvp',
      'quest',
    ]);
    expect(groups[0].label).toBe('Raids');
    expect(groups[3].sources).toHaveLength(2);
  });
});

describe('itemsOfSource and itemsOfBoss', () => {
  it('rolls a raid up to every boss plus its trash, and a boss down to its own', () => {
    const raid = file.sources.find((source) => source.id === 'raid:molten-core')!;
    expect(itemsOfSource(raid).sort()).toEqual([12784, 13968, 16963, 19325]);
    expect(itemsOfBoss(raid, 'raid:molten-core:11502')).toEqual([12784, 19325]);
    expect(itemsOfBoss(raid, 'raid:molten-core:99999')).toEqual([]);
  });
});

describe('sourceNameOf', () => {
  it('names a source and a boss, which is what Candidate.SourceName carries', () => {
    expect(sourceNameOf(file, 'raid:molten-core')).toBe('Molten Core');
    expect(sourceNameOf(file, 'raid:molten-core:11502')).toBe('Ragnaros');
    expect(sourceNameOf(file, 'nothing')).toBe('');
  });
});

describe('isOpen', () => {
  it('gates an unreleased raid and passes everything without a phase', () => {
    const raid = file.sources.find((source) => source.id === 'raid:molten-core')!;
    const dungeon = file.sources.find((source) => source.id === 'dungeon:hall-of-thanes')!;
    expect(isOpen(phases, raid, new Date('2026-11-05T00:00:00Z'))).toBe(false);
    expect(isOpen(phases, raid, new Date('2026-12-09T00:00:00Z'))).toBe(true);
    expect(isOpen(phases, dungeon, new Date('2026-09-19T00:00:00Z'))).toBe(true);
  });

  it('never opens a source whose date is unknown ("later", contract 10.4)', () => {
    const world = file.sources.find((source) => source.id === 'world:azuregos')!;
    expect(world.opens).toBe('later');
    expect(world.zone_id).toBeUndefined();
    expect(isOpen(phases, world, new Date('2099-01-01T00:00:00Z'))).toBe(false);
  });
});

describe('professionSplit', () => {
  it('splits crafted into the character’s own and the rest', () => {
    const crafted = file.sources.filter((source) => source.kind === 'crafted');
    const split = professionSplit(crafted, ['blacksmithing']);
    expect(split.mine.map((source) => source.profession)).toEqual(['blacksmithing']);
    expect(split.other.map((source) => source.profession)).toEqual(['tailoring']);
  });

  it('puts everything in `other` when nothing recorded a profession', () => {
    const crafted = file.sources.filter((source) => source.kind === 'crafted');
    const split = professionSplit(crafted, undefined);
    expect(split.mine).toEqual([]);
    expect(split.other).toHaveLength(2);
  });
});

describe('DEFAULT_OFF_KINDS', () => {
  it('is quests and nothing else', () => {
    expect([...DEFAULT_OFF_KINDS]).toEqual(['quest']);
  });
});
