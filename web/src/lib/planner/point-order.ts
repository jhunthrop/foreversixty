// web/src/lib/planner/point-order.ts
// Point order's rows (spec 2026-11-08 section 4.5, decisions P4 and P7, ruling Q5): the spend
// order in rows of ten levels (10 to 19 ... 60), one socket per level. A spent socket names the
// talent and the rank it reached; past a character's level the points are planned ahead; the
// next socket names its level; the character's level line follows its row once the build
// reaches it. ui/PointOrder draws them.
import { plannerCopy } from './copy';
import { levelForIndex, pointsAvailableAt } from './derive';
import type { TalentIndex } from './rules';
import { FIRST_POINT_LEVEL, MAX_POINTS } from './types';

export interface PointOrderSlot {
  kind: 'spent' | 'planned' | 'empty' | 'next';
  name: string;
  iconSrc?: string;
  rank?: number;
}

export interface PointOrderView {
  rows: { label: string; slots: PointOrderSlot[] }[];
  /** The row the character's level line follows (-1: none). */
  capAfter: number;
  /** The level line, once the build reaches the character's level; '' otherwise. */
  capText: string;
}

const LAST_LEVEL = FIRST_POINT_LEVEL + MAX_POINTS - 1;

export function pointOrderRows(
  index: TalentIndex,
  order: readonly number[],
  characterLevel: number | null,
  characterName: string | null,
  iconUrl: (icon: string) => string,
): PointOrderView {
  const available = characterLevel === null ? MAX_POINTS : pointsAvailableAt(characterLevel);
  const nextShown = order.length < MAX_POINTS && order.length < available;
  const ranks = new Map<number, number>();
  const spent = order.map((id, i) => {
    const rank = (ranks.get(id) ?? 0) + 1;
    ranks.set(id, rank);
    const talent = index.byId.get(id);
    const level = levelForIndex(i);
    const slot: PointOrderSlot = {
      kind: characterLevel !== null && i >= available ? 'planned' : 'spent',
      name: plannerCopy.pointAt(level, talent?.name ?? '', rank),
      rank,
    };
    if (talent !== undefined) slot.iconSrc = iconUrl(talent.icon);
    return slot;
  });
  const rows: PointOrderView['rows'] = [];
  let capAfter = -1;
  for (let lo = FIRST_POINT_LEVEL; lo <= LAST_LEVEL; lo += 10) {
    const slots: PointOrderSlot[] = [];
    for (let level = lo; level <= Math.min(lo + 9, LAST_LEVEL); level += 1) {
      const i = level - FIRST_POINT_LEVEL;
      if (i < spent.length) slots.push(spent[i]);
      else if (i === spent.length && nextShown)
        slots.push({ kind: 'next', name: plannerCopy.nextPoint(level) });
      else slots.push({ kind: 'empty', name: plannerCopy.levelSocket(level) });
    }
    rows.push({ label: plannerCopy.decade(lo), slots });
    if (characterLevel !== null && lo <= characterLevel && characterLevel <= lo + 9)
      capAfter = rows.length - 1;
  }
  const capText =
    characterLevel !== null && characterName !== null && order.length >= available
      ? plannerCopy.levelLine(characterName, characterLevel)
      : '';
  return { rows, capAfter, capText };
}
