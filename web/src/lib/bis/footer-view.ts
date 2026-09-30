// web/src/lib/bis/footer-view.ts
// "Where to get it" and "New at this band" (bis rebuild spec §4.E) -- both computed from
// the same per-band `RowView`s `panel-view.ts`'s `bandInfosFor` already built for "The
// list", never re-templated from the mock's own illustrative sentence (the mock's specific
// wording is wrong for this spec's own real data in two places the spec calls out by name:
// the "comes at 30" trinket claim and the exact instance/rep list -- see panel-view.ts's
// `noSourcedItemCopyFor` for the first; this module is the second).
import { SLOTS, type Slot } from '../planner/types';
import { SLOT_DISPLAY_LABELS } from './slot-display-labels';
import { bisCopy, joinWithAnd } from './copy';
import type { BandInfo, RowView } from './panel-view';
import type { SourceCell } from './source-cell';

export interface SourceGroupItem {
  itemId: number;
  itemName: string;
}

export interface SourceGroup {
  label: string;
  items: SourceGroupItem[];
}

/** The COARSE name a source cell groups under -- an instance without its boss, a crafted
 *  profession without "Crafted:", a rep faction without its standing -- distinct from
 *  `describeSourceCell`'s own full row line, which is too specific to group by (two
 *  different bosses in the same dungeon must group under one instance name, not two). */
function sourceGroupKey(cell: SourceCell): string {
  switch (cell.kind) {
    case 'dungeon':
    case 'raid':
      return cell.instance;
    case 'crafted':
      return cell.name;
    case 'vendor':
      return cell.npc;
    case 'rep':
      return cell.faction;
    case 'zone':
    case 'world':
      return cell.place;
    case 'pvp':
      return `${cell.faction === 'alliance' ? 'Alliance' : 'Horde'} PvP`;
    case 'quest':
      return cell.questName;
    case 'world_drop':
      return bisCopy.worldDropGroupLabel;
    case 'unknown':
      return cell.label;
  }
}

/** "Where to get it" (spec §4.E) -- groups this band's own filled rows by their pick's
 *  resolved source (`sourceGroupKey`), sorted by count descending, ties alphabetical by the
 *  group's own label. `source_kind === 'world_drop'` always groups under
 *  `bisCopy.worldDropGroupLabel` regardless of what `sourceGroupKey` would otherwise derive
 *  (kept as `sourceGroupKey`'s own `world_drop` branch so there is exactly one place that
 *  rule lives). */
export function sourceGroupsFor(rows: readonly RowView[]): SourceGroup[] {
  const groups = new Map<string, SourceGroup>();
  for (const row of rows) {
    if (row.empty || row.sourceCell === undefined || row.itemId === undefined || row.itemName === undefined) {
      continue;
    }
    const label = sourceGroupKey(row.sourceCell);
    const item: SourceGroupItem = { itemId: row.itemId, itemName: row.itemName };
    const existing = groups.get(label);
    if (existing === undefined) groups.set(label, { label, items: [item] });
    else existing.items.push(item);
  }
  return [...groups.values()].sort(
    (a, b) => b.items.length - a.items.length || a.label.localeCompare(b.label),
  );
}

/** `"head: Brawler's Leather Hood"` -> `"head"` -- `BisBand.new_at_band`'s own shape, the
 *  slot key before the colon. */
function slotKeyFromEntry(entry: string): string {
  return entry.split(':')[0]!.trim();
}

/** The distinct dungeon instance names and rep faction names among `newAtBand`'s own
 *  slots, in the order they first appear there (spec §4.E) -- resolved against `rows`'
 *  already-computed `sourceCell`, not the raw pipeline `source` string (which can name a
 *  stale/looser place than the source cell's own `loot.json`-corrected one; see
 *  `source-cell.ts`). A slot `rows` has no entry for (should not happen on real data) is
 *  skipped rather than thrown on. */
function openedSourcesFor(
  newAtBand: readonly string[],
  rows: readonly RowView[],
): { dungeons: string[]; reps: string[] } {
  const bySlot = new Map(rows.map((row) => [row.slot, row]));
  const dungeons: string[] = [];
  const reps: string[] = [];
  for (const entry of newAtBand) {
    const row = bySlot.get(slotKeyFromEntry(entry));
    const cell = row?.sourceCell;
    if (cell === undefined) continue;
    if ((cell.kind === 'dungeon' || cell.kind === 'raid') && !dungeons.includes(cell.instance)) {
      dungeons.push(cell.instance);
    } else if (cell.kind === 'rep' && !reps.includes(cell.faction)) {
      reps.push(cell.faction);
    }
  }
  return { dungeons, reps };
}

function joinedOpenedSources(dungeons: readonly string[], reps: readonly string[]): string {
  return joinWithAnd([...dungeons, ...reps.map((faction) => `the ${faction} rewards`)]);
}

/** The first slot (canonical `SLOTS` order) that was empty this band and carries a real
 *  pick next band -- spec §4.E's "the first {slot} that helps" clause, only ever comparing
 *  THIS band against the very NEXT one (never a deeper scan -- that correction belongs to
 *  panel-view.ts's own "comes at N" search for an empty row's OWN sentence, a different
 *  question). `undefined` when nothing clears between the two. */
function firstSlotThatHelpsNextBand(current: BandInfo, next: BandInfo): string | undefined {
  const currentBySlot = new Map(current.rows.map((row) => [row.slot, row]));
  const nextBySlot = new Map(next.rows.map((row) => [row.slot, row]));
  for (const slot of SLOTS) {
    const wasEmpty = currentBySlot.get(slot)?.empty ?? false;
    const nowFilled = nextBySlot.get(slot)?.empty === false;
    if (wasEmpty && nowFilled) return (SLOT_DISPLAY_LABELS[slot as Slot] ?? slot).toLowerCase();
  }
  return undefined;
}

export interface NewAtBandLines {
  sentence1: string;
  /** `undefined` at the top band, which has no next band to preview (spec §4.E). */
  sentence2: string | undefined;
}

/** "New at this band"'s two sentences (spec §4.E). `currentNewAtBand` is this band's own
 *  `BisBand.new_at_band`; `next` and `nextNewAtBand` are the next `BisBand`/`BandInfo` pair
 *  in the same file, both `undefined` at the top band. */
export function newAtBandLinesFor(
  current: BandInfo,
  bandLow: number,
  currentNewAtBand: readonly string[],
  next: BandInfo | undefined,
  nextBandLabel: string | undefined,
  nextNewAtBand: readonly string[] | undefined,
): NewAtBandLines {
  if (current.bandIndex === 0) {
    return { sentence1: bisCopy.firstBandEveryPickNew, sentence2: undefined };
  }
  const { dungeons, reps } = openedSourcesFor(currentNewAtBand, current.rows);
  const sentence1 = bisCopy.reachingBandOpens(
    bandLow,
    joinedOpenedSources(dungeons, reps),
    currentNewAtBand.length,
    current.totalSlots,
  );
  if (next === undefined || nextBandLabel === undefined || nextNewAtBand === undefined) {
    return { sentence1, sentence2: undefined };
  }
  const nextOpened = openedSourcesFor(nextNewAtBand, next.rows);
  const sentence2 = bisCopy.nextBandLine(
    nextBandLabel,
    joinedOpenedSources(nextOpened.dungeons, nextOpened.reps),
    firstSlotThatHelpsNextBand(current, next),
  );
  return { sentence1, sentence2 };
}
