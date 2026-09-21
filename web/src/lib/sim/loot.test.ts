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
  sourceLabel,
  sourceNameOf,
  sourcesByItem,
  type LootFile,
  type LootSource,
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

  /**
   * dps D35: a boss neither database names (`data/pipeline/loot/sources.py`'s own header,
   * "a boss the fork does not name is emitted with an empty name") used to surface as the
   * raw `<source id>:<npc id>` key everywhere this name is read, including the candidate's
   * own `source_name` on the wire -- a key a player never typed and should never read.
   */
  it('names an unresolved boss by its zone, never by its id', () => {
    const withUnnamedBoss: LootFile = {
      sources: [
        {
          id: 'dungeon:blackrock-spire',
          kind: 'dungeon',
          name: 'Blackrock Spire',
          bosses: [
            { id: 'dungeon:blackrock-spire:9568', name: 'Overlord Wyrmthalak', npc_id: 9568, items: [1] },
            { id: 'dungeon:blackrock-spire:175245', name: '', npc_id: 175245, items: [2] },
          ],
        },
      ],
    };
    expect(sourceNameOf(withUnnamedBoss, 'dungeon:blackrock-spire:9568')).toBe('Overlord Wyrmthalak');
    expect(sourceNameOf(withUnnamedBoss, 'dungeon:blackrock-spire:175245')).toBe(
      'Unnamed source in Blackrock Spire',
    );
  });
});

describe('sourceLabel', () => {
  function repSource(id: string, standing: string): LootSource {
    return { id, kind: 'rep', name: 'Cenarion Circle', faction_id: 609, standing, items: [1] };
  }

  /**
   * dps D30: the source picker listed four identical `Cenarion Circle` rows, one per
   * standing, with nothing to tell them apart -- `standing` already travels on the source,
   * this only needed showing.
   */
  it('appends the standing to a reputation source, so four standings read as four rows', () => {
    expect(sourceLabel(repSource('rep:cenarion-circle:friendly', 'friendly'))).toBe(
      'Cenarion Circle — Friendly',
    );
    expect(sourceLabel(repSource('rep:cenarion-circle:revered', 'revered'))).toBe(
      'Cenarion Circle — Revered',
    );
  });

  it('leaves every non-reputation source exactly as named', () => {
    const raid = file.sources.find((source) => source.id === 'raid:molten-core')!;
    expect(sourceLabel(raid)).toBe('Molten Core');
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
