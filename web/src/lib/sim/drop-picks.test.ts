// web/src/lib/sim/drop-picks.test.ts
import { describe, expect, it } from 'vitest';
import { pickedWithNothingTried, rowsFromPicks, sourceGroupVisibility, triedCount } from './drop-picks';
import type { LootFile, LootSource } from './loot';
import type { PhaseRow } from './phase';
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

/**
 * newcomer MAJOR (review.md:251-257): rowsFromPicks pushed straight into a plain array with
 * no candidateKey dedupe, so two ticked sources sharing an item produced two identical
 * candidate rows. Routed through candidates.ts's own `addRow`, which merges by
 * `candidateKey` and keeps the richer provenance (a `drop:` origin, a non-empty sourceName).
 */
describe('rowsFromPicks de-duplicates shared items across ticked sources', () => {
  const sharedLoot: LootFile = {
    sources: [
      {
        id: 'raid:molten-core',
        kind: 'raid',
        name: 'Molten Core',
        bosses: [{ id: 'raid:molten-core:11502', name: 'Ragnaros', npc_id: 11502, items: [KNOWN_ID] }],
      },
      {
        id: 'raid:onyxias-lair',
        kind: 'raid',
        name: "Onyxia's Lair",
        bosses: [{ id: 'raid:onyxias-lair:10184', name: 'Onyxia', npc_id: 10184, items: [KNOWN_ID] }],
      },
    ],
  };
  const bothPicked = ['raid:molten-core|raid:molten-core:11502', 'raid:onyxias-lair|raid:onyxias-lair:10184'];

  it('merges two ticked sources that share one item into a single row', () => {
    const rows = rowsFromPicks(bothPicked, sharedLoot, items);
    expect(rows).toHaveLength(1);
    const [row] = rows;
    expect(row.item.id).toBe(KNOWN_ID);
    expect(row.origin.startsWith('drop:')).toBe(true);
    expect(row.sourceName).not.toBe('');
  });
});

/**
 * An id present in loot.json's item list but present in neither this class's `items` map
 * nor the engine's `known` set is not something that will be tried -- the picker's badge
 * used to print the raw loot.json count, which is why a source advertising "2" could
 * contribute 0 simulated items (newcomer MAJOR, review.md:291-298; dps D34).
 */
describe('triedCount', () => {
  // Never added to `items` below -- stands in for an id loot.json lists that this class's
  // item file (data/builds/<build>/items/<class>.json) has no row for at all.
  const MISSING_ID = 999999;

  it('does not count an id missing from the item map', () => {
    expect(triedCount([MISSING_ID], items, null)).toBe(0);
  });

  it('does not count an id absent from a non-null known set', () => {
    const known = new Set<number>();
    expect(triedCount([KNOWN_ID], items, known)).toBe(0);
  });

  it('counts every id present in the item map when known is null (no simitems.json for this build)', () => {
    expect(triedCount([KNOWN_ID, UNKNOWN_ID, MISSING_ID], items, null)).toBe(2);
  });
});

/**
 * newcomer MAJOR (review.md:291-298) and dps D34: two ticked dungeons, one of which
 * contributes nothing tried, used to leave that dungeon unmentioned anywhere in the
 * result. This is the decision DropResults.svelte renders a row from.
 */
describe('pickedWithNothingTried', () => {
  const MISSING_ID = 999999;

  const bothSourcesLoot: LootFile = {
    sources: [
      {
        id: 'raid:molten-core',
        kind: 'raid',
        name: 'Molten Core',
        bosses: [{ id: 'raid:molten-core:11502', name: 'Ragnaros', npc_id: 11502, items: [KNOWN_ID] }],
      },
      {
        id: 'raid:onyxias-lair',
        kind: 'raid',
        name: "Onyxia's Lair",
        bosses: [{ id: 'raid:onyxias-lair:10184', name: 'Onyxia', npc_id: 10184, items: [MISSING_ID] }],
      },
    ],
  };
  const picks = ['raid:molten-core|raid:molten-core:11502', 'raid:onyxias-lair|raid:onyxias-lair:10184'];

  it('names a ticked boss pick that contributed nothing, and leaves out one that contributed something', () => {
    const untried = pickedWithNothingTried(picks, bothSourcesLoot, items, null);
    expect(untried).toEqual([{ key: 'raid:onyxias-lair|raid:onyxias-lair:10184', name: 'Onyxia' }]);
  });

  it('names a whole-source pick (no boss id) the same way', () => {
    const untried = pickedWithNothingTried(['raid:onyxias-lair|'], bothSourcesLoot, items, null);
    expect(untried).toEqual([{ key: 'raid:onyxias-lair|', name: "Onyxia's Lair" }]);
  });

  it('returns nothing when every ticked pick contributed at least one tried item', () => {
    const untried = pickedWithNothingTried([picks[0]], bothSourcesLoot, items, null);
    expect(untried).toEqual([]);
  });
});

/**
 * dps D33 (BLOCKER, review.md:334-342): a ticked kind whose every source is gated behind
 * an unopened phase used to render nothing at all -- no heading, no explanation. This is
 * the decision SourcePicker.svelte renders the "opens on" fallback from.
 */
describe('sourceGroupVisibility', () => {
  const at = new Date('2026-01-01T00:00:00Z');
  const gatedPhases: PhaseRow[] = [{ name: 'raids-1', start: '2026-12-09T00:00:00Z' }];

  function raidSource(id: string, opens?: string): LootSource {
    return { id, kind: 'raid', name: id, opens };
  }

  it('is allGated when every source of a ticked kind is behind an unopened phase', () => {
    const sources = [raidSource('raid:a', 'raids-1'), raidSource('raid:b', 'raids-1')];
    const visibility = sourceGroupVisibility('raid', sources, ['raid'], gatedPhases, false, at);
    expect(visibility).toMatchObject({ ticked: true, shown: [], allGated: true });
  });

  it('is not allGated when some sources of the kind are already open', () => {
    const sources = [raidSource('raid:a', 'raids-1'), raidSource('raid:b')];
    const visibility = sourceGroupVisibility('raid', sources, ['raid'], gatedPhases, false, at);
    expect(visibility.allGated).toBe(false);
    expect(visibility.shown.map((source) => source.id)).toEqual(['raid:b']);
  });

  it('is not allGated when the kind has no sources at all', () => {
    const visibility = sourceGroupVisibility('raid', [], ['raid'], gatedPhases, false, at);
    expect(visibility).toMatchObject({ ticked: true, shown: [], allGated: false });
  });

  it('is not allGated when the kind is not ticked', () => {
    const sources = [raidSource('raid:a', 'raids-1')];
    const visibility = sourceGroupVisibility('raid', sources, [], gatedPhases, false, at);
    expect(visibility).toMatchObject({ ticked: false, shown: [], allGated: false });
  });
});
