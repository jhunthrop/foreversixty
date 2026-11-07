// web/src/lib/account/level-filter.ts
// The one low-level rule every character list shares (the header selector, /account): every
// character is imported whatever its level, but a list shows level 10 and up by default and
// folds the rest behind "Show low-level characters (N)". A character whose level is unknown is
// never folded, and a kept key (the current character, the main) always shows.
import { sharedText } from '../copy/shared';

/** Characters below this level are folded by default. */
export const LOW_LEVEL_BELOW = 10;

export const lowLevelCopy = {
  show: (count: number): string => sharedText('shared.lowLevel.show', { count }),
  hide: sharedText('shared.lowLevel.hide'),
} as const;

export function isLowLevel(item: { level?: number }): boolean {
  return item.level !== undefined && item.level < LOW_LEVEL_BELOW;
}

export interface LevelSplit<T> {
  /** What the list draws: everything when `showLow`, else everything not folded. */
  shown: T[];
  /** How many low-level items are folded away (0 when `showLow`). */
  hidden: number;
  /** How many low-level items the list has at all: whether a toggle is offered. */
  lowCount: number;
}

/** `items` with the low-level ones folded unless `showLow`; order is kept and nothing is mutated. */
export function splitByLevel<T extends { key: string; level?: number }>(
  items: readonly T[],
  showLow: boolean,
  keepKeys: readonly (string | null | undefined)[] = [],
): LevelSplit<T> {
  const foldable = (item: T): boolean => isLowLevel(item) && !keepKeys.includes(item.key);
  const lowCount = items.filter(foldable).length;
  if (showLow) return { shown: [...items], hidden: 0, lowCount };
  return { shown: items.filter((item) => !foldable(item)), hidden: lowCount, lowCount };
}
