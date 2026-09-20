// web/src/lib/sim/drop-picks.ts
// The Droptimizer's own slice of bulk-store.svelte.ts, pulled into its own module because
// that file was pushing toward the 800-line cap: which sources are shown, which are picked,
// and the rows a pick turns into have no reason to sit beside the gear-candidate and
// run-envelope logic above them in the store. Every function here is pure -- the store still
// owns the state (`pickedBosses`, `showUpcoming`, `shownKinds`), these just compute the next
// value from it.
import { addRow, rowFor, uiSlotsOf, type CandidateRow, type Origin } from './candidates';
import { DEFAULT_OFF_KINDS, isOpen, itemsOfBoss, itemsOfSource, sourceNameOf } from './loot';
import type { LootFile, LootSource } from './loot';
import type { PhaseRow } from './phase';
import { isKnownItem } from './sim-items';
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
 *
 * Every row is added through `addRow` (candidates.ts), not pushed straight into the array:
 * two ticked sources that share an item (the same boss picked twice under different names,
 * or two bosses dropping the same drop) would otherwise land as two character-for-character
 * identical rows, with no `candidateKey` dedupe to catch it (newcomer MAJOR, review.md:251-
 * 257). `addRow` merges same-key rows and keeps the richer provenance -- a `drop:` origin
 * and a non-empty `sourceName` both win.
 */
export function rowsFromPicks(
  picked: readonly string[],
  loot: LootFile,
  items: ReadonlyMap<number, Item>,
  /**
   * `sim-items.ts`'s known ids, `null` when the build ships no `simitems.json` -- the same
   * "nothing to filter against" default `isKnownItem` already gives. A boss can legitimately
   * drop an item the engine's database does not carry (the re-itemised raid tier is thin,
   * per loot.ts's own header); the row still shows, ticked, but disabled -- removing it
   * from the picker entirely would make the boss look like it drops nothing rather than
   * naming a known limitation.
   */
  known: ReadonlySet<number> | null = null,
): CandidateRow[] {
  let rows: CandidateRow[] = [];
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
      const itemKnown = isKnownItem(itemId, known);
      for (const slot of uiSlotsOf(item)) {
        rows = addRow(rows, { ...rowFor(item, slot, origin, sourceName, itemKnown), checked: itemKnown });
      }
    }
  }
  return rows;
}

/**
 * How many of these ids will actually be tried -- present in this class's item map AND
 * known to the engine (sim-items.ts's `isKnownItem`), the same two gates `rowsFromPicks`
 * above already applies per item. `SourcePicker.svelte`'s badge used to print loot.json's
 * raw item count instead, which is why a source advertising "2" could contribute 0
 * simulated items (newcomer MAJOR, review.md:291-298; dps D34).
 */
export function triedCount(
  itemIds: readonly number[],
  items: ReadonlyMap<number, Item>,
  known: ReadonlySet<number> | null,
): number {
  return itemIds.filter((id) => items.has(id) && isKnownItem(id, known)).length;
}

export interface UntriedPick {
  /** The `<source id>|<boss id or "">` key, the same shape a `picked` entry carries. */
  key: string;
  /** The source or boss name, resolved the same way `rowsFromPicks` resolves it. */
  name: string;
}

/**
 * Which ticked picks contributed zero tried items -- a ticked source or boss `rowsFromPicks`
 * produced no rows for at all, and so vanishes from a by-boss grouping of the result with
 * no trace it was ever asked for (newcomer MAJOR, review.md:291-298; dps D34).
 *
 * Says only what `triedCount` can prove: nothing from this pick was tried. It does not
 * guess why -- an id can fail one of `triedCount`'s two gates for two different reasons
 * this data cannot tell apart (absent from this class's item file, or absent from the
 * engine's `simitems.json`), and asserting either one specifically would be a claim the
 * data does not support.
 */
export function pickedWithNothingTried(
  picked: readonly string[],
  loot: LootFile,
  items: ReadonlyMap<number, Item>,
  known: ReadonlySet<number> | null,
): UntriedPick[] {
  const untried: UntriedPick[] = [];
  for (const pick of picked) {
    const [sourceId, bossId] = pick.split('|');
    const source = loot.sources.find((entry) => entry.id === sourceId);
    if (source === undefined) continue;
    const itemIds = bossId === '' ? itemsOfSource(source) : itemsOfBoss(source, bossId);
    if (triedCount(itemIds, items, known) > 0) continue;
    const pickedId = bossId === '' ? sourceId : bossId;
    untried.push({ key: pick, name: sourceNameOf(loot, pickedId) });
  }
  return untried;
}

export interface SourceGroupVisibility {
  /** Whether the player has this kind ticked in the top filter row. */
  ticked: boolean;
  /** The sources of this kind that pass today's show-upcoming rule. */
  shown: LootSource[];
  /** Ticked, has at least one source of this kind, and every one of them is gated shut. */
  allGated: boolean;
}

/**
 * Whether a ticked kind's picker group renders nothing not because it has no sources, but
 * because every one of them is gated behind a phase that has not opened -- the state that
 * used to fall straight out of `SourcePicker.svelte`'s `{#if shown.length > 0}` guard with
 * no heading and no explanation (dps D33, BLOCKER, review.md:334-342; newcomer MAJOR,
 * review.md:299-304, "ticking Raids produces an empty void").
 */
export function sourceGroupVisibility(
  kind: string,
  sources: readonly LootSource[],
  shownKinds: readonly string[],
  phases: readonly PhaseRow[],
  showUpcoming: boolean,
  at: Date,
): SourceGroupVisibility {
  const ticked = shownKinds.includes(kind);
  const shown = ticked ? sources.filter((source) => showUpcoming || isOpen(phases, source, at)) : [];
  return { ticked, shown, allGated: ticked && sources.length > 0 && shown.length === 0 };
}
