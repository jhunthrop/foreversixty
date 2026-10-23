// web/src/lib/guides/build-diff.ts
// The one line under the raid build that names what it changes: "The raid build moves 5 points
// from A (2) and B (3) to C (2) and D (3)". Computed from the two builds' decoded trees, so it
// can never disagree with the trees drawn above it.
import type { DecodedTreeView } from './build-tree';

export interface BuildDiff {
  /** Talents that lose points in the raid build, with the points lost. */
  from: { name: string; points: number }[];
  /** Talents that gain points, with the points gained. */
  to: { name: string; points: number }[];
  /** Points moved: the total lost. */
  moved: number;
}

function ranksById(
  trees: readonly DecodedTreeView[],
): Map<number, { name: string; rank: number; order: number }> {
  const ranks = new Map<number, { name: string; rank: number; order: number }>();
  let order = 0;
  for (const tree of trees) {
    for (const talent of tree.talents) {
      ranks.set(talent.id, { name: talent.name, rank: talent.rank, order });
      order += 1;
    }
  }
  return ranks;
}

export function diffBuilds(
  leveling: readonly DecodedTreeView[],
  raid: readonly DecodedTreeView[],
): BuildDiff | undefined {
  const before = ranksById(leveling);
  const after = ranksById(raid);
  type Moved = { name: string; points: number; order: number };
  const from: Moved[] = [];
  const to: Moved[] = [];
  for (const [id, now] of after) {
    const was = before.get(id)?.rank ?? 0;
    if (now.rank > was) to.push({ name: now.name, points: now.rank - was, order: now.order });
    if (now.rank < was) from.push({ name: now.name, points: was - now.rank, order: now.order });
  }
  if (from.length === 0 && to.length === 0) return undefined;
  const strip = (list: Moved[]) =>
    list.sort((a, b) => a.order - b.order).map(({ name, points }) => ({ name, points }));
  return {
    from: strip(from),
    to: strip(to),
    moved: from.reduce((sum, entry) => sum + entry.points, 0),
  };
}

const LIST_CAP = 4;

function listOf(entries: readonly { name: string; points: number }[]): string {
  const shown = entries.slice(0, LIST_CAP).map((entry) => `${entry.name} (${entry.points})`);
  const rest = entries.length - shown.length;
  if (rest > 0) shown.push(`${rest} more`);
  if (shown.length <= 1) return shown.join('');
  return `${shown.slice(0, -1).join(', ')} and ${shown[shown.length - 1]}`;
}

function pointsText(count: number): string {
  return `${count} point${count === 1 ? '' : 's'}`;
}

/** The sentence, or `undefined` when the two builds are the same. */
export function buildDiffLine(diff: BuildDiff | undefined): string | undefined {
  if (diff === undefined) return undefined;
  if (diff.from.length > 0 && diff.to.length > 0) {
    return `The raid build moves ${pointsText(diff.moved)} from ${listOf(diff.from)} to ${listOf(diff.to)}.`;
  }
  if (diff.to.length > 0) {
    const added = diff.to.reduce((sum, entry) => sum + entry.points, 0);
    return `The raid build puts ${added} more point${added === 1 ? '' : 's'} into ${listOf(diff.to)}.`;
  }
  return `The raid build takes ${pointsText(diff.moved)} out of ${listOf(diff.from)}.`;
}
