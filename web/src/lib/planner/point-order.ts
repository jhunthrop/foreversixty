// web/src/lib/planner/point-order.ts
// Point order's rows (spec 2026-11-08 section 4.5, decisions P4 and P7, ruling Q5): the spend
// order in rows of ten levels (10 to 19 ... 60), one socket per level. A spent socket names the
// talent and the rank it reached, with the game's talent tooltip; past a character's level the
// points are planned ahead; the next socket names its level; the character's level line follows
// its row once the build reaches it. For a character, the sockets past its build carry the spec
// guide's next points (`ahead`, lane P7), dimmed, so the levels still to come read as a plan.
// ui/PointOrder draws them.
import type { OrderSlot } from '../ui/point-order';
import { plannerCopy } from './copy';
import { levelForIndex, pointsAvailableAt } from './derive';
import { canAddPoint, type TalentIndex } from './rules';
import { FIRST_POINT_LEVEL, MAX_POINTS } from './types';

export type PointOrderSlot = OrderSlot;

export interface PointOrderView {
  rows: { label: string; slots: PointOrderSlot[] }[];
  /** The row the character's level line follows (-1: none). */
  capAfter: number;
  /** The level line, once the build reaches the character's level; '' otherwise. */
  capText: string;
}

/** The guide's plan past a build: its ladder, the build's points taken out, then each point in the
 *  ladder's order that is legal on top of what comes before it, up to the 51st. */
export function aheadOf(index: TalentIndex, order: readonly number[], ladder: readonly number[]): number[] {
  const have = new Map<number, number>();
  for (const id of order) have.set(id, (have.get(id) ?? 0) + 1);
  const built = [...order];
  const ahead: number[] = [];
  for (const id of ladder) {
    if (built.length >= MAX_POINTS) break;
    const held = have.get(id) ?? 0;
    if (held > 0) {
      have.set(id, held - 1);
      continue;
    }
    if (!canAddPoint(index, built, id).ok) continue;
    built.push(id);
    ahead.push(id);
  }
  return ahead;
}

const LAST_LEVEL = FIRST_POINT_LEVEL + MAX_POINTS - 1;

/** The plan past the character's build: the guide's spec and its points. */
export interface PointOrderAhead {
  specName: string;
  points: readonly number[];
}

function talentSlot(
  index: TalentIndex,
  kind: 'spent' | 'planned' | 'ahead',
  id: number,
  position: number,
  rank: number,
  iconUrl: (icon: string) => string | null,
): PointOrderSlot {
  const talent = index.byId.get(id);
  const slot: PointOrderSlot = {
    kind,
    name: plannerCopy.pointAt(levelForIndex(position), talent?.name ?? '', rank),
    rank,
  };
  if (talent !== undefined) {
    slot.iconSrc = iconUrl(talent.icon) ?? undefined;
    slot.tip = {
      name: talent.name,
      rankLine: plannerCopy.rankOf(rank, talent.max_rank),
      now: talent.ranks[rank - 1]?.description ?? '',
    };
  }
  return slot;
}

export function pointOrderRows(
  index: TalentIndex,
  order: readonly number[],
  characterLevel: number | null,
  characterName: string | null,
  iconUrl: (icon: string) => string | null,
  ahead: PointOrderAhead | null = null,
): PointOrderView {
  const available = characterLevel === null ? MAX_POINTS : pointsAvailableAt(characterLevel);
  // A character's build continues in its guide's order; an empty build previews the guide's order (h-audit 2026-11-09).
  const plan = characterLevel !== null || order.length === 0 ? (ahead?.points ?? []) : [];
  const ranks = new Map<number, number>();
  const nextRank = (id: number): number => {
    const rank = (ranks.get(id) ?? 0) + 1;
    ranks.set(id, rank);
    return rank;
  };
  const spent = order.map((id, i) =>
    talentSlot(
      index,
      characterLevel !== null && i >= available ? 'planned' : 'spent',
      id,
      i,
      nextRank(id),
      iconUrl,
    ),
  );
  const planned = plan.map((id, i) =>
    talentSlot(index, 'ahead', id, order.length + i, nextRank(id), iconUrl),
  );
  const filled = [...spent, ...planned];
  const nextShown = planned.length === 0 && order.length < MAX_POINTS && order.length < available;
  const rows: PointOrderView['rows'] = [];
  let capAfter = -1;
  for (let lo = FIRST_POINT_LEVEL; lo <= LAST_LEVEL; lo += 10) {
    const slots: PointOrderSlot[] = [];
    for (let level = lo; level <= Math.min(lo + 9, LAST_LEVEL); level += 1) {
      const i = level - FIRST_POINT_LEVEL;
      if (i < filled.length) slots.push(filled[i]);
      else if (i === filled.length && nextShown)
        slots.push({ kind: 'next', name: plannerCopy.nextPoint(level) });
      else slots.push({ kind: 'empty', name: plannerCopy.levelSocket(level) });
    }
    rows.push({ label: plannerCopy.decade(lo), slots });
    if (characterLevel !== null && lo <= characterLevel && characterLevel <= lo + 9)
      capAfter = rows.length - 1;
  }
  const capText =
    characterLevel === null || characterName === null || order.length < available
      ? ''
      : planned.length > 0 && ahead !== null
        ? plannerCopy.aheadLine(characterName, characterLevel, ahead.specName)
        : plannerCopy.levelLine(characterName, characterLevel);
  return { rows, capAfter, capText };
}
