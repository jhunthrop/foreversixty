// web/src/lib/guild/report-view.ts
// The Report's orderings and caps (attendance spec 3.4): deaths most first, consumables worst
// first, loot by quality then name inside its boss, and the five biggest moves. Pure; the
// components only draw what these return.
import type { RaidConsumableRow, RaidDeathRow, RaidLootAward, RaidMoveRow } from './attendance-contract';
import { BIGGEST_CHANGES_ROWS } from './attendance-copy';
import { rosterOrder } from './raid-roster';

const keyOf = (row: { class: string; name: string }): { role: null; class: string; name: string } => ({
  role: null,
  class: row.class,
  name: row.name,
});

/** The game's sort: class order, then name. */
function gameOrder(a: { class: string; name: string }, b: { class: string; name: string }): number {
  return rosterOrder(keyOf(a), keyOf(b));
}

export const deathCount = (row: RaidDeathRow): number => row.on_wipes + row.on_kills;

/** Most deaths first, then the game's sort. */
export function deathOrder(rows: readonly RaidDeathRow[]): RaidDeathRow[] {
  return [...rows].sort((a, b) => deathCount(b) - deathCount(a) || gameOrder(a, b));
}

function usedNothing(row: RaidConsumableRow): boolean {
  return row.flask !== true && row.food !== true && (row.potions ?? 0) === 0;
}

/** 0: Short with nothing used; 1: Short; 2: the rest. */
function consumableRank(row: RaidConsumableRow): number {
  if (row.before !== 'short') return 2;
  return usedNothing(row) ? 0 : 1;
}

export function consumableOrder(rows: readonly RaidConsumableRow[]): RaidConsumableRow[] {
  return [...rows].sort((a, b) => consumableRank(a) - consumableRank(b) || gameOrder(a, b));
}

/** Quality (best first), then item name. */
export function lootOrder(awards: readonly RaidLootAward[]): RaidLootAward[] {
  return [...awards].sort((a, b) => b.quality - a.quality || a.item_name.localeCompare(b.item_name));
}

const moveSize = (row: RaidMoveRow): number => Math.abs(row.after - row.before);

/** The largest absolute moves, at most `BIGGEST_CHANGES_ROWS`; a flat move is no change. */
export function biggestMoves(rows: readonly RaidMoveRow[]): RaidMoveRow[] {
  return rows
    .filter((row) => moveSize(row) > 0)
    .sort((a, b) => moveSize(b) - moveSize(a) || gameOrder(a, b))
    .slice(0, BIGGEST_CHANGES_ROWS);
}

/** The raid page's absolute URL with the report anchor, for the copied text. */
export function reportUrlOf(location: Pick<Location, 'origin' | 'pathname'>): string {
  return `${location.origin}${location.pathname}#report`;
}
