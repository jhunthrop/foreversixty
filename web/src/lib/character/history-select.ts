// web/src/lib/character/history-select.ts
// Selection and hover (experience spec 2026-10-24-character-history 3.7 and 3.8) as plain
// values: which point is selected, which one the mouse previews, which list rows that marks,
// and how much of a list shows. Every function returns a new value.
import type { CharacterHistory, HistoryEncounter, HistoryFight } from './history-api';
import { fightKey } from './history-api';
import type { ScoreDay } from './history-chart';

/** Rows shown before `Show all`, and the most a list ever shows. */
export const COLLAPSED_ROWS = 8;
export const MAX_ROWS = 100;

export interface Selection {
  /** The selected point's index, oldest first. */
  selected: number;
  /** The point a mouse previews, or null. */
  hover: number | null;
}

export type Step = 'earlier' | 'later' | 'first' | 'newest';

function clamp(index: number, count: number): number {
  return Math.min(Math.max(index, 0), Math.max(count - 1, 0));
}

/** The newest point is selected first. */
export function initialSelection(count: number): Selection {
  return { selected: clamp(count - 1, count), hover: null };
}

export function select(selection: Selection, index: number, count: number): Selection {
  return { ...selection, selected: clamp(index, count) };
}

export function withHover(selection: Selection, index: number | null, count: number): Selection {
  return { ...selection, hover: index === null ? null : clamp(index, count) };
}

/** What the readout shows: the preview while the mouse is over the plot, else the selection. */
export function shownIndex(selection: Selection): number {
  return selection.hover ?? selection.selected;
}

export function step(selection: Selection, move: Step, count: number): Selection {
  const target: Record<Step, number> = {
    earlier: selection.selected - 1,
    later: selection.selected + 1,
    first: 0,
    newest: count - 1,
  };
  return select(selection, target[move], count);
}

/** The step buttons disable at the ends. */
export function canStep(selection: Selection, move: 'earlier' | 'later', count: number): boolean {
  return move === 'earlier' ? selection.selected > 0 : selection.selected < count - 1;
}

/** The point nearest an x position, for a click, a tap or a hover. */
export function nearestIndex(xs: readonly number[], x: number): number {
  let best = 0;
  xs.forEach((candidate, index) => {
    if (Math.abs(candidate - x) < Math.abs(xs[best] - x)) best = index;
  });
  return best;
}

/** The list rows the selected point marks: every fight of a score day, or the one kill. */
export function markedKeys(keysByPoint: readonly (readonly string[])[], selection: Selection): Set<string> {
  return new Set(keysByPoint[selection.selected] ?? []);
}

export function scoreMarkKeys(days: readonly ScoreDay[]): (readonly string[])[] {
  return days.map((day) => day.keys);
}

export function killMarkKeys(encounter: HistoryEncounter): (readonly string[])[] {
  return encounter.points.map((kill) => [fightKey(kill.report_id, kill.fight_index)]);
}

function fightsFor(history: CharacterHistory, keys: ReadonlySet<string>): HistoryFight[] {
  return history.fights.filter((fight) => keys.has(fightKey(fight.report_id, fight.fight_index)));
}

/** The fights behind the score line, newest first as the API sends them. */
export function scoreListFights(history: CharacterHistory): HistoryFight[] {
  return fightsFor(history, new Set(history.score.map((p) => fightKey(p.report_id, p.fight_index))));
}

/** The kills of one boss, newest first. */
export function killListFights(history: CharacterHistory, encounter: HistoryEncounter): HistoryFight[] {
  return fightsFor(history, new Set(killMarkKeys(encounter).flat()));
}

export type ListButton = 'none' | 'show-all' | 'show-fewer';

export interface ListWindow {
  rows: HistoryFight[];
  /** The count `Show all N` names: every row, up to 100. */
  total: number;
  button: ListButton;
  /** A marked row sits below the collapsed rows, so the list must open to show it. */
  needsExpand: boolean;
}

/** The first marked row's position in the list, or -1. */
export function firstMarkedRow(fights: readonly HistoryFight[], marked: ReadonlySet<string>): number {
  return fights.findIndex((fight) => marked.has(fightKey(fight.report_id, fight.fight_index)));
}

/** 8 rows, or all of them (up to 100) once expanded. `needsExpand` tells the component to
 *  expand when a selection marks a hidden row; collapsing afterwards stays collapsed. */
export function listWindow(
  fights: readonly HistoryFight[],
  expanded: boolean,
  marked: ReadonlySet<string>,
): ListWindow {
  const total = Math.min(fights.length, MAX_ROWS);
  const needsExpand = firstMarkedRow(fights, marked) >= COLLAPSED_ROWS;
  const open = expanded;
  const hasMore = total > COLLAPSED_ROWS;
  return {
    rows: fights.slice(0, open ? total : COLLAPSED_ROWS),
    total,
    button: !hasMore ? 'none' : open ? 'show-fewer' : 'show-all',
    needsExpand,
  };
}
