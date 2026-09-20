// web/src/lib/sim/drop-picks.ts
// The Droptimizer's own slice of bulk-store.svelte.ts, pulled into its own module because
// that file was pushing toward the 800-line cap: which sources are shown, which are picked,
// and the rows a pick turns into have no reason to sit beside the gear-candidate and
// run-envelope logic above them in the store. Every function here is pure -- the store still
// owns the state (`pickedBosses`, `showUpcoming`, `shownKinds`), these just compute the next
// value from it.
import { rowFor, uiSlotsOf, type CandidateRow, type Origin } from './candidates';
import { DEFAULT_OFF_KINDS, isOpen, itemsOfBoss, itemsOfSource, sourceNameOf } from './loot';
import type { LootFile, LootSource } from './loot';
import type { PhaseRow } from './phase';
import type { Item } from '../planner/types';

/** The kinds a freshly loaded build's sources open with: every kind but the off-by-default ones. */
export function initialShownKinds(loot: LootFile): string[] {
  return [...new Set(loot.sources.map((source) => source.kind))].filter(
    (kind) => !DEFAULT_OFF_KINDS.includes(kind),
  );
}

/** Toggles a `<source id>|<boss id or "">` pick. A raid card and one of its bosses are two keys. */
export function togglePick(picked: readonly string[], sourceId: string, bossId: string): string[] {
  const key = `${sourceId}|${bossId}`;
  return picked.includes(key) ? picked.filter((entry) => entry !== key) : [...picked, key];
}

export function toggleShownKind(shown: readonly string[], kind: string): string[] {
  return shown.includes(kind) ? shown.filter((entry) => entry !== kind) : [...shown, kind];
}

/** The picker's own filter: a kind the player has on, released unless asked otherwise. */
export function visibleSources(
  loot: LootFile,
  shownKinds: readonly string[],
  phases: readonly PhaseRow[],
  showUpcoming: boolean,
  at: Date,
): LootSource[] {
  return loot.sources.filter(
    (source) => shownKinds.includes(source.kind) && (showUpcoming || isOpen(phases, source, at)),
  );
}

/**
 * The Droptimizer's picks, as ticked rows with a `drop:` origin and nothing else -- gear
 * candidates from bags, bank or search never belong in `drops` mode (contract's validation
 * rule 3, `validateBulk`), so this is the whole of that tool's row list, not an addition to it.
 */
export function rowsFromPicks(
  picked: readonly string[],
  loot: LootFile,
  items: ReadonlyMap<number, Item>,
): CandidateRow[] {
  const rows: CandidateRow[] = [];
  for (const pick of picked) {
    const [sourceId, bossId] = pick.split('|');
    const source = loot.sources.find((entry) => entry.id === sourceId);
    if (source === undefined) continue;
    const pickedId = bossId === '' ? sourceId : bossId;
    const origin = `drop:${pickedId}` as Origin;
    // Contract 10.1 A6: the name is resolved once, here, and rides on the candidate.
    const sourceName = sourceNameOf(loot, pickedId);
    for (const itemId of bossId === '' ? itemsOfSource(source) : itemsOfBoss(source, bossId)) {
      const item = items.get(itemId);
      if (item === undefined) continue;
      for (const slot of uiSlotsOf(item)) {
        rows.push({ ...rowFor(item, slot, origin, sourceName), checked: true });
      }
    }
  }
  return rows;
}
