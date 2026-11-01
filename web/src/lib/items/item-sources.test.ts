import { describe, expect, it } from 'vitest';
import { itemSourceRows, sourceRowLine } from './item-sources';
import type { LootFile } from '../sim/loot';

const ID = 100;

function loot(file: Partial<LootFile>): LootFile {
  return { sources: [], quests: {}, ...file } as LootFile;
}

describe('itemSourceRows', () => {
  it('names a vendor by the vendor, with no chance', () => {
    const rows = itemSourceRows(
      ID,
      loot({
        sources: [{ id: 'vendor:1', kind: 'vendor', name: 'Meilosh', items: [ID] }] as LootFile['sources'],
      }),
    );
    expect(rows).toEqual([{ kind: 'vendor', title: 'Meilosh', primary: false, text: 'Meilosh' }]);
  });

  it('gives a quest its side and minimum level, ahead of a vendor', () => {
    const rows = itemSourceRows(
      ID,
      loot({
        sources: [{ id: 'vendor:1', kind: 'vendor', name: 'Meilosh', items: [ID] }] as LootFile['sources'],
        quests: {
          [ID]: [
            { quest_id: 1, name: 'Saving the Best for Last', faction: 'horde', min_level: 58, level: 60 },
          ],
        },
      }),
    );
    expect(rows[0]).toEqual({
      kind: 'quest',
      title: 'Saving the Best for Last',
      place: 'Horde',
      min_level: 58,
      primary: true,
      text: 'Saving the Best for Last (Horde)',
    });
  });

  it('collapses every world pool that lists the item into its own words, once', () => {
    const rows = itemSourceRows(
      ID,
      loot({
        sources: [
          { id: 'world:a', kind: 'world', name: 'Tanaris', items: [ID] },
          { id: 'world:b', kind: 'world', name: 'Feralas', items: [ID] },
          { id: 'world:c', kind: 'world', name: 'Silithus', items: [] },
        ] as LootFile['sources'],
      }),
    );
    expect(rows.map(sourceRowLine)).toEqual(['World drop']);
  });

  it('prints a known chance in the tooltip line form', () => {
    const [row] = itemSourceRows(
      ID,
      loot({
        sources: [
          {
            id: 'dungeon:x',
            kind: 'dungeon',
            name: 'Scholomance',
            bosses: [{ id: 'g', name: 'Gandling', items: [ID], item_chances: { [ID]: 0.4 } }],
          },
        ] as LootFile['sources'],
      }),
    );
    expect(sourceRowLine(row)).toBe('Scholomance — Gandling (<1%)');
  });
});
