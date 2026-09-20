// web/src/lib/sim/drop-picks.test.ts
import { describe, expect, it } from 'vitest';
import { rowsFromPicks } from './drop-picks';
import type { LootFile } from './loot';
import type { Item } from '../planner/types';

function item(id: number, slot: string, name = `Item ${id}`): Item {
  return {
    id,
    name,
    icon: `fixture_${id}`,
    slot,
    quality: 4,
    required_level: 60,
    item_level: 70,
    armor: 0,
    stats: {},
    set_id: null,
    unique: false,
  };
}

const KNOWN_ID = 12784;
const UNKNOWN_ID = 21550;

const loot: LootFile = {
  sources: [
    {
      id: 'raid:molten-core',
      kind: 'raid',
      name: 'Molten Core',
      bosses: [
        {
          id: 'raid:molten-core:11502',
          name: 'Ragnaros',
          npc_id: 11502,
          items: [KNOWN_ID, UNKNOWN_ID],
        },
      ],
    },
  ],
};

const items = new Map<number, Item>([
  [KNOWN_ID, item(KNOWN_ID, 'main_hand', 'Arcanite Reaper')],
  [UNKNOWN_ID, item(UNKNOWN_ID, 'trinket', 'Idol of the White Stag')],
]);

const picked = ['raid:molten-core|raid:molten-core:11502'];

describe('rowsFromPicks', () => {
  it('ticks every drop when there is nothing to filter against (no simitems.json for this build)', () => {
    const rows = rowsFromPicks(picked, loot, items);
    expect(rows.map((row) => row.item.id).sort()).toEqual([KNOWN_ID, UNKNOWN_ID]);
    expect(rows.every((row) => row.known)).toBe(true);
    expect(rows.every((row) => row.checked)).toBe(true);
  });

  it('keeps a drop the engine does not know about in the list, unticked and marked unknown, rather than hiding it', () => {
    const known = new Set([KNOWN_ID]);
    const rows = rowsFromPicks(picked, loot, items, known);
    const byId = new Map(rows.map((row) => [row.item.id, row]));
    expect(byId.get(KNOWN_ID)).toMatchObject({ known: true, checked: true });
    expect(byId.get(UNKNOWN_ID)).toMatchObject({ known: false, checked: false });
    // Still present -- a boss that drops an item the engine cannot simulate must not look
    // like it drops nothing.
    expect(rows).toHaveLength(2);
  });
});
